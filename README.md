Cheetah API {#index}
========================

[TOC]

## Introduction

For Access Points based on the Cheetah operating system (like Barbados, etc), various applications could make use of the services provided by the access point, e.g, get information about the system, the radio(s), the clients connected, etc, or register and receive packets or accounting records that they are interested in analyzing. Such programming is exposed through the Cheetah API, which is very rich in nature.

Exposing the Cheetah Access Point API as a Google RPC (or GRPC), over Google protocol buffers (protobuf or GPB), enables customers to write their own applications, controllers, etc., whether on box or off box, in a rich set of languages including c++, python, GO, etc.

Please make sure to read app-hosting.md for the complete set of instructions for using the Cheetah API to build an application that could be hosted on the Access Points using IOx.

## Services

The Cheetah Access Point API is currently organized in a set of files that expose certain verticals e.g. statistics, packet capture, accounting records, etc.
In the initial release, the focus is to provide the following verticals:

* Initialization: This mainly handles global initialization, and sets up an event notification channel based on GRPC streaming mechanism.
* Statistics: This returns data from the access point regarding system resources, radio information, clients, configured WLANs, etc
* Packet capture: This allows an application to register for and receive a category of packets it's interested in (management or data).
* Accounting records: This allows an application to register for and receive a set of accounting records it's interested in. This functionality is under development.

The Cheetah API allows for GRPC unary functions in most cases, and GRPC streaming in other cases. The former can be rendered in both synchronous and asynchronous modes (depends on the language). The latter is used for continuous transmitting and/or receiving of objects in an asynchronous fashion. This is especially useful to boost performance in certain cases. Please refer to the GRPC website for more information: <http://grpc.io>

Each RPC usually takes a GRPC "message" or request, typically labeled (Something)Msg, example APInitMsg, which defines the parameters of the request, and return another "message", typically labeled (Something)MsgRsp as a response to the RPC request, example APInitMsgRsp.

Note that all files are annotated with detailed documentation.
The user of the API can use doxygen to render his/her own local documentation, refer to instructions under docs directory. The html generated documentation is broken up into sections that describe the messages, verticals, files, etc, and are very useful for quick reference.

Finally, please note that the API comes with:

## Tutorial

A quick start tutorial written in python and Go. The intent here is to get the user a jump-start on hooking up with the API. The reader is advized to try this next.

This can be found here:
```
grpc/python/src/tutorial/
grpc/go/src/tutorial/
```

## Websocket

A simple client/server Go implementation over websockets is also included and can be found here:

```
grpc/go/src/websocket/
```

A more detailed README.md file can be found there.

## Unit tests

A Python unittest regression suite that covers basic API sanities. It is also very useful and handy if someone wants to get some reference implementation for a certain use case.

This can be found here:

```
grpc/python/src/tests
```

To run the unit test regression, setup some Environment variables:

```
export SERVER_IP=192.168.122.192
export SERVER_PORT=57777
```

The above assumes that the IP address of the node is 192.168.122.192.

Run All tests:

```
python -m unittest -v test_cheetah
```

We hope that the above was useful quick overview about the Cheetah API. We recommend that the reader goes over the python quick tutorial first and then go over the .proto files under grpc/protos (or look at the generated .html pages, these are not kept in this repo, but can be auto-generated from this repo).

Any questions or concerns please contact:

cheetah-objmodel@cisco.com
