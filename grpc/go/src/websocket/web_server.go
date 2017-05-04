/*
 * Copyright (c) 2017 by Cisco Systems, Inc.
 * All rights reserved.
 */
package main

/* Standard packages */
import (
    "flag"
    "fmt"
    "github.com/golang/protobuf/jsonpb"
    "github.com/gorilla/websocket"
    "github.com/influxdata/influxdb/client/v2"
    "log"
    "net"
    "net/http"
    "time"
)

/* Cheetah packages */
import (
    pb "gengo"
)

/* Command line arguments */
var (
    debug = flag.Bool("debug", false, "Enable debugging")
)

/* Debug related */
type Debug bool
var dbg Debug

func (d Debug) Println(a ...interface{}) {
    if d {
        fmt.Println(a...)
    }
}

/* Connected clients */
var clients = make(map[*websocket.Conn]bool)

// Configure the upgrader
var upgrader = websocket.Upgrader{}

/* DB client instance */
var db_client client.Client
var db_err error

/* DB names */
const iox_db_name = "apstats_iox"

func get_db_bp() client.BatchPoints {

    // Create a new point batch
    db_bp, err := client.NewBatchPoints(client.BatchPointsConfig{
                    Database:  iox_db_name,
                    Precision: "s",
    })

    if err != nil {
        log.Fatal(err)
        return nil
    }

    return db_bp
}

func ws_to_remote_host(ws *websocket.Conn) string {
    host, _, _ := net.SplitHostPort(ws.RemoteAddr().String())
    return host
}

func process_system_stats(ws *websocket.Conn, msg *pb.APSystemStatsMsgRsp) {
    //dbg.Println(*msg)
}

func process_memory_stats(ws *websocket.Conn, msg *pb.APMemoryStatsMsgRsp) {
    dbg.Println(*msg)

    /* Generate a new batchpoint */
    db_bp := get_db_bp()
    if db_bp == nil {
        return
    }

    tags := map[string]string{
                "AP":  fmt.Sprintf("%s", ws_to_remote_host(ws)),
    }
    // Create a point and add to batch
    fields := map[string]interface{}{
            "mem avail": msg.ProcMemInfo.GetAvailableKB(),
    }

    pt, err := client.NewPoint("memory", tags, fields, time.Now())
    if err != nil {
        log.Fatal(err)
    }
    db_bp.AddPoint(pt)

    // Write the batch
    if err := db_client.Write(db_bp); err != nil {
        log.Fatal(err)
    }
}

func process_interface_stats(ws *websocket.Conn, msg *pb.APInterfaceStatsMsgRsp) {
    dbg.Println(*msg)

    for _, ifEntry := range msg.GetInterfaces() {

        /* Generate a new batchpoint */
        db_bp := get_db_bp()
        if db_bp == nil {
            return
        }

        /* Create tags */
        tags := map[string]string{
                    "AP":  fmt.Sprintf("%s", ws_to_remote_host(ws)),
        }

        /* Create a point and add to batch */
        fields := map[string]interface{}{
                "RxPkts": ifEntry.GetRxPkts(),
                "RxBytes": ifEntry.GetRxBytes(),
                "TxPkts": ifEntry.GetTxPkts(),
                "TxBytes": ifEntry.GetTxBytes(),
                "RxDiscards": ifEntry.GetRxDiscards(),
        }

        pt, err := client.NewPoint("interface", tags, fields, time.Now())
        if err != nil {
            log.Fatal(err)
        }
        db_bp.AddPoint(pt)

        // Write the batch
        if err := db_client.Write(db_bp); err != nil {
            log.Fatal(err)
        }
    }
}

func process_routing_stats(ws *websocket.Conn, msg *pb.APRoutingStatsMsgRsp) {
    //dbg.Println(*msg)
}

func process_dns_stats(ws *websocket.Conn, msg *pb.APDNSStatsMsgRsp) {
    //dbg.Println(*msg)
}

func process_radio_stats(ws *websocket.Conn, msg *pb.APRadioStatsMsgRsp) {
    //dbg.Println(*msg)
}

func process_wlan_stats(ws *websocket.Conn, msg *pb.APWLANStatsMsgRsp) {
    //dbg.Println(*msg)
}

func process_client_stats(ws *websocket.Conn, msg *pb.APClientStatsMsgRsp) {
    //dbg.Println(*msg)
}


func process_msg(ws *websocket.Conn, msg *pb.APStatsMsgRsp) {
    switch  msg.MsgRsp.(type) {
      case *pb.APStatsMsgRsp_SystemStats:
        process_system_stats(ws, msg.GetSystemStats())
      case *pb.APStatsMsgRsp_MemoryStats:
        process_memory_stats(ws, msg.GetMemoryStats())
      case *pb.APStatsMsgRsp_InterfaceStats:
        process_interface_stats(ws, msg.GetInterfaceStats())
      case *pb.APStatsMsgRsp_RoutingStats:
        process_routing_stats(ws, msg.GetRoutingStats())
      case *pb.APStatsMsgRsp_DNSStats:
        process_dns_stats(ws, msg.GetDNSStats())
      case *pb.APStatsMsgRsp_RadioStats:
        process_radio_stats(ws, msg.GetRadioStats())
      case *pb.APStatsMsgRsp_WLANStats:
        process_wlan_stats(ws, msg.GetWLANStats())
      case *pb.APStatsMsgRsp_ClientStats:
        process_client_stats(ws, msg.GetClientStats())
      default:
        log.Println("Websocket server error: unknown type")
    }
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

    // Upgrade initial GET request to a websocket
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    // Make sure we close the connection when the function returns
    defer ws.Close()

    // Register our new client
    clients[ws] = true

    marshaller := jsonpb.Marshaler{}

    for {

        var msg *pb.APStatsMsgRsp = new(pb.APStatsMsgRsp)
        data, _ := marshaller.MarshalToString(msg)

        // Read in a new message as JSON and map it to a Message object
        err := ws.ReadJSON(&data)
        if err != nil {
            log.Printf("error: %v", err)
            delete(clients, ws)
            break
        }

        if err := jsonpb.UnmarshalString(data, msg); err != nil {
            panic(err)
        }
        if (msg.ErrStatus.Status != pb.APErrorStatus_AP_SUCCESS) {
            log.Println("Websocket server message error: bad msg")
            continue
        }
        if (msg.ErrStatus.Status == pb.APErrorStatus_AP_NOT_AVAILABLE) {
            log.Println("Websocket server message error: no records")
            continue
        }
        go process_msg(ws, msg)
    }
}

func db_init() {
    db_client, db_err = client.NewHTTPClient(client.HTTPConfig{
                            Addr: "http://localhost:8086"})
    if db_err != nil {
        fmt.Println("Error creating InfluxDB client: ", db_err.Error())
    }

    var create_op =  "CREATE DATABASE " + iox_db_name
    q := client.NewQuery(create_op, "", "")
    if response, err := db_client.Query(q); err == nil && response.Error() == nil {
        dbg.Println("Created apstats IOx DB")
    }
}

func main() {
    /* Parse any command line arguments */
    flag.Parse()

    if (*debug) {
        dbg = true
    }

    /* Initialize influxDB */
    db_init()

    /* Hook up websocker handler */
    //http.Handle("/ws", websocket.Handler(handler))
    http.HandleFunc("/ws", handleConnections)

    // Start the server on localhost port 8080 and log any errors
    fmt.Println("http server started on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
