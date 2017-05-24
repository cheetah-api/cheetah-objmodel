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
    "golang.org/x/net/websocket"
    "flag"
    "os"
)

/* Cheetah packages */
import (
    "cheetah"
    "util"
    pb "gengo"
)

/* Command line arguments */
var (
    web_server = flag.String("web-ip", "", "Websocket server IP")
    web_port = flag.String("web-port", "", "Websocket server port")
)

func initWssClient() (conn *websocket.Conn) {
    fmt.Println("Starting Web Client")
    http_address := ""

    log.Println("Initializing connection to Web server ...")

    if (*web_server != "") {
        fmt.Printf("Web Server IP: %s\n", *web_server)
        if (*web_port != "") {
            log.Printf("Web Server Port: %s\n", *web_port)
        } else {
            log.Printf("Please provide Websocket Port\n")
            os.Exit(1)
        }
        http_address = fmt.Sprintf("%s:%s", *web_server, *web_port)
    } else {
        /* Get HTTP Server IP and Port from Env */
        http_server, http_port := util.GetHTTPServerIPPort()
        http_address = fmt.Sprintf("%s:%s", http_server, http_port)
    }

    log.Printf("Initializing connection to %s: ", http_address)
    wss_conn, err := websocket.Dial(fmt.Sprintf("ws://%s/ws", http_address), "",
                              fmt.Sprintf("http://%s/", http_address))
    if err != nil {
        log.Printf("Dial failed: %s\n", err.Error())
        os.Exit(1)
    }

    return wss_conn
}

func main() {
    /* Parse any command line arguments */
    flag.Parse()

    done := make(chan bool)

    /* Initialize with web server */
    wss_conn := initWssClient()

    log.Println("Initializing connection to gRPC server ...")
    /* Get GRPC Server IP and Port from Env */
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
    go util.StatsOperation(conn, wss_conn, pb.APStatsType_AP_SYSTEM_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* Memory statistics */
    go util.StatsOperation(conn, wss_conn, pb.APStatsType_AP_MEMORY_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* Interface statistics */
    go util.StatsOperation(conn, wss_conn, pb.APStatsType_AP_INTERFACE_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* Routing statistics */
    go util.StatsOperation(conn, wss_conn, pb.APStatsType_AP_ROUTING_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* DNS statistics */
    go util.StatsOperation(conn, wss_conn, pb.APStatsType_AP_DNS_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* Radio statistics */
    go util.StatsOperation(conn, wss_conn, pb.APStatsType_AP_RADIO_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* WLAN statistics */
    go util.StatsOperation(conn, wss_conn, pb.APStatsType_AP_WLAN_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))

    /* Client statistics */
    go util.StatsOperation(conn, wss_conn, pb.APStatsType_AP_CLIENT_STATS,
                           uint32(pb.StatsTimeInterval_AP_STATS_MIN_TIME_INTERVAL))


    <-done
    /*The process will exit here*/
}
