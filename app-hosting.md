Application hosting
====================

# Introduction

This page provides some basic information for developing applications that could be deployed on the Cisco AP28xx/AP38xx Access Points using Cisco's IOx technology.

# IOx

IOx is Cisco’s implementation of “Fog Computing”. IOx enables hosting of applications and services developed by Cisco, it’s partners and third party developers in the network edge devices in a seamless fashion across diverse and disparate hardware platforms.

IOx provides a seamless application enablement framework and compute platform across various devices operating at the network edge with the ability to host applications and services, connecting them securely and reliably to applications in the cloud. The term Application enablement covers all life cycle aspects of applications including development, distribution, deployment, hosting, monitoring and management.

For more information about IOx please visit the following link.
[https://developer.cisco.com/site/iox/documents/developer-guide/]

# Setup

To develop an application for the AP, you need to download the Software Development Environment as described below.

## Local setup

1. Download IOx SDK for the Access Point
    [https://cisco.box.com/s/hdf81s7ij7z2isw1mk0hcgq957198udo]
2. Download Virtual Box
3. Import IOx SDK appliance downloaded above
4. Start image on Virtual Box
5. Open terminal

## IOx SDK setup

The default username to get access to the SDK is ioxsdk (password: cisco123)

## Enable IOx on the Access Point (AP2800, AP3800 series)

Please ensure that a gRPC server is running on the AP before you try to connect over gRPC.  For the time being this is done when IOX is enabled:

From the controller:
    config ap apphost apgroup <name> [enable | disable]

From the AP, verify that IOx is up:

    show running-config | include IOX
       IOX   : Enabled

Use the Access Point IP and Port 57777 to export to your docker container before running the gRPC client.

export SERVER_IP=10.32.170.208
export SERVER_PORT=57777

# Develop and test applications

## Create Python application
The official IOX Documentation for building Python applications is [here](https://developer.cisco.com/media/iox-dev-guide-11-28-16/docker/simple-python).

### Create application directory
```
   mkdir -p ~/apps/ap-python-client/devel
```

### Download Cheetah API

```
   cd ~/apps/ap-python-client/devel

   git clone ...
```

### Build development image
A devel Dockerfile that is suitable for downloading the necessary devel packages for python over gRPC/protobufs has been included for convenience and could be used. The base AP package that we need is base-ap3k.

```
# Go to application development directory
cd ~/apps/ap-python-client/devel

# Create the Dockerfile; To use sample do:
cp cheetah-objmodel/grpc/python/Dockerfile.devel Dockerfile

# Build the image; this will take a while
sudo docker build -t ap-python-client-devel .
```

### Run and test application
Once you have an image built, you could start a docker container and do some sanity testing before you deploy the image to the access point. You can run your image by issuing the following command:

```
sudo docker run -ti ap-python-client-devel /bin/sh
```

You could navigate to /opt/grpc and follow the tutorial for how to build your own application or run the available tests. Please ensure that the gRPC server is running first (see previous section).
You could do multiple iterations and test your code within the SDE before you even have to deploy to the platform for system testing.
Examples of what you could do:

```
cd /opt/grpc/python/src
python tutorial/quickstart.py
python -m unittest tests/test_cheetah
python -m unittest tests/test_cheetah.TestSuite_001_Statistics.test_003_get_dns_stats
```

### Build release image
Once you are happy with your application, you can prepare to build the image you will load on the access point. This step is necessary in order to cherry pick what you need to be on the released image and leave development packages and toolchains out of it.

```
# Go to application directory
cd ~/apps/ap-python-client

# Create a share directory where we can export content from the development package
sudo docker run -ti -v ${PWD}/share:/opt/share ap-python-client-devel /bin/sh

# From within the container, copy the site-packages to the share point.
#
# It is important to also copy any work that is not a build artifact and is
# required in the running image. For example, if you're adding your own python
# program you also need to copy it to /opt/share and also below from the Dockerfile.
(container shell)# cp -fr /usr/lib/python2.7/site-packages /opt/share
(container shell)# exit

# Copy "release" Dockerfile; If you edit it, you will notice that it copies
# the content of ./share to /usr/lib so that we pick up the site-packages
# that got built from the previous step.
#
# It is important to also copy any work that is not a build artifact and is
# required in the running image. For example, if you're adding your own python
# program you also need to copy within the Dockerfile.
#
cp devel/cheetah-objmodel/grpc/python/Dockerfile .

# Build the image; this is the one that will eventually get deployed
sudo docker build -t ap-python-client .
```

You could at this point verify that the "release" image will work as expected since it depends on what you have "exported". You can run your image by issuing the following command, notice the difference in the name of the image which now doesn't have -devel.

```
sudo docker run -ti ap-python-client /bin/sh
```

## Create Go application

For general instructions on running Go, including tutorials and general guidelines, go to the [grpc.io website](http://www.grpc.io/docs/quickstart/go.html)

### Create application directory
```
   mkdir -p ~/apps/ap-go-client/devel
```

### Download Cheetah API

```
   cd ~/apps/ap-go-client/devel

   git clone ...
```

### Build development image

A devel Dockerfile that is suitable for downloading the necessary devel packages for Go over gRPC/protobufs has been included for convenience and could be used. The base AP package that we need is ap3k/base-rootfs.

```
# Go to application development directory
cd ~/apps/ap-go-client/devel

wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
tar xzvf go1.8.linux-amd64.tar.gz

# Set ENV variables
export GOROOT=/home/ioxsdk/apps/ap-go-client/devel/go
export PATH=$PATH:$GOROOT/bin
export GOARCH=arm
export GOPATH=/home/ioxsdk/apps/ap-go-client/devel/cheetah-objmodel/grpc/go

# Use the following command to install gRPC
go get google.golang.org/grpc

# Next, install the protoc plugin for Go (optional if you have the Go bindings already)
go get -a github.com/golang/protobuf/protoc-gen-go

# Test/Build the tutorial

# You could use "go run" rather than "go build" to play with it before creating the binary
# This would require setting the SERVER_IP/PORT as below and changing $GOARCH to amd64
cd $GOPATH/src/tutorial
go build quickstart.go

# Copy the Go binary
cp quickstart ~/apps/ap-go-client/devel
cd ~/apps/ap-go-client

# Bring over the Dockerfile
cp devel/cheetah-objmodel/grpc/go/Dockerfile .

# Build the image; this is the one that will eventually get deployed
sudo docker build -t ap-go-client .

# You can test it first
sudo docker run -ti ap-go-client /bin/sh

/ #
/ # export SERVER_IP=10.32.170.82
/ # export SERVER_PORT=57777
/ # /opt/grpc/quickstart
Using SERVER IP PORT: 10.32.170.82:57777
Server Returned AP_SUCCESS
Server response:  AP_SUCCESS
Max Radio Name Len  :  16
Max Ssid Name Len   :  16
```

# Deploy application

## Package descriptor file
You need to create a package.yaml file that describes the application requirements. A basic yaml file for the AP has been added and could be used. The destination folder depends on the SDK version being used. See below.
In the text below the application folder is referred to as "ap-xxx-client" but it's any of "ap-python-client", "ap-go-client" or "ap-cpp-client" that you have built.

## Build the application package

Bring over package.yaml

```
# Go to the application folder
cd ~/apps/ap-xxx-client

# Create a conf directory
mkdir conf

# Bring over sample package.yaml and put it into the "conf" directory
cp devel/cheetah-objmodel/package.yaml conf
```

Build the image.

```
sudo ioxclient docker package –p ext2 ap-xxx-client ./conf
```

You will find the image in conf/package.tar.
You could now copy out the package and install it on the AP using the Fog Director or ioxclient itself.


## Application console access

Once the application has been successfully deployed on your device, you could access it in one of the following ways:

### Using ssh

From the Fog Director, under your device, find the "App Console Support" section.
The IP address it the one of your device. Use SSH_PORT 22. Download the PEM file.

```
ssh -p 22 -i <pemfile> appconsole@<IPaddress>
```

### Using ioxclient

From your SDK, you can use ioxclient to deploy, start and connect to the application console. This assumes you have created the proper ioxclient profile to connect to the device.

```
ioxsdk@ioxsdk-VirtualBox:~$ ioxclient application list
Currently active profile : default
Command Name: application-list
List of installed App :
1. ap-xxx-client ---> RUNNING

ioxsdk@ioxsdk-VirtualBox:~$ ioxclient application console ap-xxx-client
Currently active profile : default
Command Name: application-console
Console setup is complete..
Running command : [ssh -p 22 -i ap-xxx-client.pem appconsole@10.32.170.237]
Connected to domain ap-xxx-client
Escape character is ^]
Poky (Yocto Project Reference Distro) 1.7.2 AP-Host-Shell /dev/tty1
AP-Host-Shell login: root
root@AP-Host-Shell:~#
```
