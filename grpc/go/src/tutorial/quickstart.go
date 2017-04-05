/*
 * Copyright (c) 2016 by cisco Systems, Inc.
 * All rights reserved.
 */
package main

/* Standard packages */
import (
    "fmt"
    "log"
    "google.golang.org/grpc"
    "flag"
)

/* Cheetah packages */
import (
    "cheetah"
    "util"
    /* pb "gengo" */
)

func main() {
    /* Parse any command line arguments */
    flag.Parse()

    /* Get Server IP and Port from Env */
    server,port := util.GetServerIPPort()
    address := fmt.Sprintf("%s:%s", server, port)

    /* Setup the connection with the server */
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    /* Initialize and handshake with server */
    if access_point.ClientInit(conn) == 0 {
        log.Fatalf("ClientInit error")
        return
    }

    /* Issue an RPC to get the system statistics */
    util.SystemStatsOperation(conn)

    /*The process will exit here*/
}
