Python Quick Tutorial {#tutorial}
=====================

## Table of Contents
- [Server Setup](#server)
- [Client Setup](#client)
- [Running the tutorial](#quick)
- [Generate gRPC Code](#gen)
- [Initialize the client server connection](#init)

### <a name='server'></a>Server Setup

On the server side, GRPC is already enabled as part of the Access Point bring up sequence.

TBD: do we want to make the GRPC port configurable from the WLC config or the AP config?
     For now assume a well known port.

We also need to configure a server IP address. The IP address is assigned by the system

TBD: ...


### <a name='client'></a>Client Setup

On the client side, the very first thing we need to do is set the server IP address and gRPC port.

Set the server address and port number as environment variables with the
following example command (this is assuming you are in bash shell):

```
    $ export SERVER_IP=192.168.122.192
    $ export SERVER_PORT=57344
```


Note: For IOX applications this might be exposed by other TBD means.

The above assumes that the IP address of the node is 192.168.122.192.
This completes all the setup needed to start writing some code! Hop into
your python interpreter and try out some of the commands to get familiar
with the API.

### <a name='quick'></a>Running the tutorial

The following basic tutorial will walk you through getting started with the Access Point API.
This may require some initial python and GRPC setup, which will be explained below. For now,
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

    | Script | Description |
    | ------ | ----------- |
    | quickstart.py  | The full tutorial example |
    | client_init.py | Used to setup the client-server connection |

As shown in quickstart.py, the first thing to do is to setup the GRPC channel:

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
        #
        # # Set the client version number based on the current proto files' version
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
Notice that the client_init definition above takes a GRPC stub as an argument.
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

        # The main reason we spawn a thread here, is that we dedicate a GRPC
        # channel to listen on Global asynchronous events/notifications.
        # This thread will be handling these event notifications.
        t = threading.Thread(target = global_thread, args=(stub, event))
        t.start()
```
