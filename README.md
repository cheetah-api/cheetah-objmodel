Cheetah API {#index}
========================

[TOC]

## Introduction

For Access Points based on the Cheetah operating system (like Barbados, etc), various applications could make use of the services provided by the access point, e.g, get information about the system, the radio(s), the clients connected, etc, or register and receive packets or accounting records that they are interested in analyzing. Such programming is exposed through the Cheetah API, which is very rich in nature.

Exposing the Cheetah Access Point API as a Google RPC (or GRPC), over Google protocol buffers (protobuf or GPB), enables customers to write their own applications, controllers, etc., whether on box or off box, in a rich set of languages including c++, python, GO, etc.

## Services

The Cheetah Access Point API is currently organized in a set of files that expose certain verticals e.g. statistics, packet capture, accounting records, etc.
In the initial release, the focus is to provide the following verticals:

* Initialization: This mainly handles global initialization, and sets up an event notification channel based on GRPC streaming mechanism.
* Statistics: This returns data from the access point regarding system resources, radio information, clients, configured WLANs, etc
* Packet capture: This allows an application to register for and receive a category of packets it's interested in (management or data)
* Accounting records: This allows an application to register for and receive a set of accounting records it's interested in
* More functions may be added in the future.

Each function vertical, e.g. statistics vertical, declares a "template" set of RPCs that is more or less consistently followed throughout other verticals. Some of these template RPCs are explained here:

* (Vertical)Get(): This is mainly used to query certain capabilities for that vertical.
* (Vertical)GetStats(): This is mainly used to query vertical specific statistics.
* (Vertical)RegOp(): This is mainly used to Register/Unregister/EoF, which basically notifies the server about interest in the vertical, no interest, and end of file (EoF), respectively. The EoF marker is especially useful on replay of objects in certain restart scenarios.
* (Vertical)(Object)Op(): This is mainly used to add, delete, update objects. The convention used for add and update, is that, object 'adds' may fail if the object already exists, whereas update can create or simply override the object if it exists.
* (Vertical)(Object)Get(): This is mainly used to retrieve an object or a set of objects.
* Stream(): This is mainly a GRPC "streaming" version of the non-streaming version of the function.
* Notif(): This is mainly a streaming notification function, e.g. asynchronous BFD session state events' streaming.

The Cheetah API allows for GRPC unary functions in most cases, and GRPC streaming in other cases. The former can be rendered in both synchronous and asynchronous modes (depends on the language). The latter is used for continuous transmitting and/or receiving of objects in an asynchronous fashion. This is especially useful to boost performance in certain cases. Please refer to the GRPC website for more information: <http://grpc.io>
In addition, certain RPCs may also allow for batching e.g. creating a number of routes in a single RPC call (in a batch).

Each RPC usually takes a GRPC "message" or request, typically labeled (Something)Msg, example APInitMsg, which defines the parameters of the request, and return another "message", typically labeled (Something)MsgRsp as a response to the RPC request, example APInitMsgRsp.

Note that all files are annotated with detailed documentation.
The user of the API can use doxygen to render his/her own local documentation, refer to instructions under docs directory. The html generated documentation is broken up into sections that describe the messages, verticals, files, etc, and are very useful for quick reference.

Finally, please note that the API comes with:

## Tutorial

A quick start tutorial written in python. The intent here is to get the user a jump-start on hooking up with the API. The reader is advized to try this next.

This can be found here:
```
grpc/python/src/tutorial/
```

## Unit tests

A Python unittest regression suite that covers basic API sanities. It is also very useful and handy if someone wants to get some reference implementation for a certain use case.

This can be found here:

```
grpc/python/src/tests
```

To run the unit test regression, setup some Environment variables:

```
export SERVER_IP=192.168.122.192
export SERVER_PORT=57344
```

Run All tests:

```
python -m unittest -v tests.test_cheetah
```

We hope that the above was useful quick overview about the Cheetah API. We recommend that the reader goes over the python quick tutorial first and then go over the .proto files under grpc/protos (or look at the generated .html pages, these are not kept in this repo, but can be auto-generated from this repo).
