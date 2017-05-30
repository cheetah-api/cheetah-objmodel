/*
 * Copyright (c) 2017 by Cisco Systems, Inc.
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
    pb "gengo"
)

func main() {
    /* Parse any command line arguments */
    flag.Parse()

    done := make(chan bool)

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

    ///* System statistics */
    go util.StatsOperation(conn, nil, pb.APStatsType_AP_SYSTEM_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* Memory statistics */
    go util.StatsOperation(conn, nil, pb.APStatsType_AP_MEMORY_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* Interface statistics */
    go util.StatsOperation(conn, nil, pb.APStatsType_AP_INTERFACE_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* Routing statistics */
    go util.StatsOperation(conn, nil, pb.APStatsType_AP_ROUTING_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* DNS statistics */
    go util.StatsOperation(conn, nil, pb.APStatsType_AP_DNS_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* Radio statistics */
    go util.StatsOperation(conn, nil, pb.APStatsType_AP_RADIO_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* WLAN statistics */
    go util.StatsOperation(conn, nil, pb.APStatsType_AP_WLAN_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* Client statistics */
    go util.StatsOperation(conn, nil, pb.APStatsType_AP_CLIENT_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))


    <-done
    /*The process will exit here*/
}
