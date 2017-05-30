Python Quick Tutorial {#tutorial}
=====================

## Table of Contents
- [Server Setup](#server)
- [Client Setup](#client)
- [Running the tutorial](#quick)
- [Generate gRPC Code](#gen)
- [Initialize the client server connection](#init)
- [How to get statistics](#stats)

### <a name='server'></a>Server Setup

On the server side, gRPC is already enabled as part of the Access Point bring up sequence.
The server IP is that assigned to the wired interface of the access point (wired0) and the
port is statically set to 57777.

### <a name='client'></a>Client Setup

On the client side, the very first thing we need to do is set the server IP address and gRPC port.

If your application is a hosted one and started by IOx then the variables below are being populated automatically by the infrastructure and you can skip this step.

Set the server address and port number as environment variables with the following example command (this is assuming you are in bash shell):

```
    $ export SERVER_IP=192.168.122.192
    $ export SERVER_PORT=57777
```

The above assumes that the IP address of the node is 192.168.122.192.

This completes all the setup needed to start writing some code! Hop into
your python interpreter and try out some of the commands to get familiar
with the API.

### <a name='quick'></a>Running the tutorial

The following basic tutorial will walk you through getting started with the Cheetah Access Point API.
This may require some initial python and gRPC setup, which will be explained below. For now,
if you already have passed this setup step, run the example:

```
    cd grpc/python/src
    python tutorial/quickstart.py
```

The following sections explain the details of the above example tutorial.

#### <a name='gen'></a>Generate gRPC Code (optional in this example)

If you are not familiar with gRPC, we recommend you refer to gRPC's
documentation before beginning with our tutorial: [gRPC Docs](http://www.grpc.io/docs/)

You should have received all of the protobuf files required to run the
Access Point API. In order to generate the gRPC client side code stubs in python,
run the following command (you may have to adjust the path to the proto files and the
output according to your requirements):

**For convenience, these files are also committed in this repo under
  grpc/python/src/genpy/ (so you can skip this step).**

```
    $ protoc -I ../../protos --python_out=. --grpc_out=./genpy/ \
             --plugin=protoc-gen-grpc=`which grpc_python_plugin` \
             ../../protos/*.proto
```

This generates the code stubs that we will now utilize to create a client.
The files are recognizable from the "_pb2" that is appended to the name of the
proto files they were generated from (example: ap_global_init.py).

#### <a name='init'></a>Initialize the client server connection

In order to follow this quick tutorial, it is best to open the files in `grpc/python/src/tutorial/`

|Script | Description|
|:----: | :---------:|
|quickstart.py  | The full tutorial example|
|client_init.py | Used to setup the client-server connection|
|stats.py       | Sample code to start a stats collection thread |

As shown in quickstart.py, the first thing to do is to setup the gRPC channel:

```python
    server_ip, server_port = util.get_server_ip_port()
    channel = implementations.insecure_channel(server_ip, server_port)
```

Once connected, we need to handshake the API version number with the server.
The same RPC call also sets up an asynchronous stream of notifications from the server.
The first notification would be the response to our version number message i.e. APInitMsg,
as an APGlobalNotif event with type AP_GLOBAL_EVENT_TYPE_VERSION. This can be done by calling:

```python
    APGlobalInitNotif(init_msg, Timeout)
```

The above function takes the client major, minor and sub version numbers and sends
them to the Access Point service to get a handshake on the API version number.
More on this below.

The following code snippets are copied from file client_init.py

```python
    def client_init(stub, event):
        #
        # Create APInitMsg to handshake the version number with the server.
        # The Server will allow/deny access based on the version number.
        # The same RPC is used to setup a notification channel for global
        # events coming from the server.

        # Set the client version number based on the current proto files' version
        init_msg = ap_global_pb2.APInitMsg()
        init_msg.MajorVer = ap_version_pb2.AP_MAJOR_VERSION
        init_msg.MinorVer = ap_version_pb2.AP_MINOR_VERSION
        init_msg.SubVer = ap_version_pb2.AP_SUB_VERSION

        # Set a very large timeout, as we will "for ever" loop listening on
        # notifications from the server
        Timeout = 365*24*60*60 # Seconds

        # This for loop will never end unless the server closes the session
        for response in stub.APGlobalInitNotif(init_msg, Timeout):
            if response.EventType == ap_global_pb2.AP_GLOBAL_EVENT_TYPE_VERSION:
                if (ap_common_types_pb2.APErrorStatus.AP_SUCCESS ==
                        response.ErrStatus.Status):
```

The above python definition also handles other events such as errors and heartbeats.
Notice that the client_init definition above takes a gRPC stub as an argument.
This is typically created through:

```python
    # Create the gRPC stub.
    stub = ap_global_pb2.beta_create_APGlobal_stub(channel)
```

Since the above client_init function would never return, it is best to spawn
it as a thread, and run it in the background. In python, we do so by calling a threading event:

```python
    #
    # Spawn a thread for global events
    #
    def global_init(channel):
        # Create the gRPC stub.
        stub = ap_global_pb2.beta_create_APGlobal_stub(channel)

        # Create a thread sync event. This will be used to order thread execution
        event = threading.Event()

        # The main reason we spawn a thread here, is that we dedicate a gRPC
        # channel to listen on Global asynchronous events/notifications.
        # This thread will be handling these event notifications.
        t = threading.Thread(target = global_thread, args=(stub, event))
        t.start()
```

#### <a name='stats'></a>How to get statistics

The Cheetah API provides for a way to get statistics in both a unary and stream type of operations. The input required by the API is the type of statistics required (system, memory, interface, etc) and the cadence/time interval at which the gRPC server should push the statistics.

```python
    message APStatsRequest {

        // The type to be retrieved
        APStatsType StatsType = 1;

        // The time interval (cadence) that the server should use to push statistics.
        // If set to 0 the server will respond only once
        // For positive values, the connection will remain open and the server will be
        // pushing statistics of this category every TimeInterval seconds 
        uint32 TimeInterval = 2;
    }
```

The code below will initiate the retrieval of system statistics every 6 seconds.

```python
   # System Stats every 6 seconds
   t1=stats.stats_operations(channel, ap_stats_pb2.AP_SYSTEM_STATS, 6)
```

This function will start and keep a thread to receive this stream of messages back from the server. You can see that in stats.py:
 
```python
def stats_operations(channel, stats_type, time_interval):
    t = threading.Thread(target = stats_thread, args=(channel, stats_type, time_interval))
    t.start()
    return t
```

The stats_thread itself will initialize the get request:

```python
    # Get the system level stats. Create APStatsMsg
    stats_msg = ap_stats_pb2.APStatsMsg()

    #
    # Add system stats to the list, once every 5 seconds
    #
    stats = stats_msg.StatsRequest.add()
    #system_stats.StatsType = ap_stats_pb2.AP_SYSTEM_STATS
    stats.StatsType = stats_type
    stats.TimeInterval = time_interval
```

and then loop and process the messages coming back:

```python
    for response in stub.APStatsGet(stats_msg, Timeout):
      if (response.ErrStatus.Status ==
            ap_common_types_pb2.APErrorStatus.AP_SUCCESS):
            if response.HasField("SystemStats"):
                util.print_system_stats(response.SystemStats)
            elif response.HasField("MemoryStats"):
                util.print_memory_stats(response.MemoryStats)
        else:
            print "Stats response error 0x%x" %(response.ErrStatus.Status)
            os._exit(0)

```
