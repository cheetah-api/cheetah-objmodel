Sample client/server example using websockets {#websockets}
==============================================

## Table of Contents
- [Introduction](#intro)
- [Web client description](#client)
- [Web server description](#server)
- [Install Grafana on Ubuntu](#ubuntu)
- [Install Grafana on Centos](#centos)

### <a name='intro'></a>Introduction

 The provided example code demonstrates how to create an application
 that collects information from the AP and sends them to a server that
 adds them to a database like InfluxDB and graphs them out using Grafana.

   Server side                                         Client (AP) side
   -----------------------                      ------------------------
   Grafana                                               gRPC Server
       |__ InfluxDB                                           |
             |__ WebServer    <--websocket-->      WebClient__|
                                  (JSON)

### <a name='client'></a>Web client desciption

The web client performs the following tasks:
 - Connect to the gRPC server
 - Connect to the Web server
 - Request a stream of statistics to be pushed by the gRPC server
 - Convert to JSON
 - Send to web server

### <a name='server'></a>Web server desciption

The web server performs the following tasks:
 - Initialize the Influx DB
 - Listen for websocket connections

For each incoming connection:
 - Convert from JSON to protobuf
 - Process each message and insert data in to the DB

### <a name='ubuntu'></a>Install Grafana/InfluxDB on Ubuntu
  http://www.andremiller.net/content/grafana-and-influxdb-quickstart-on-ubuntu

### <a name='centos'></a>Install Grafana/InfluxDB on Centos
  http://vmkdaily.ghost.io/influxdb-and-grafana-on-centos/

