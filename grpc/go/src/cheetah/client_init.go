/*
 * Copyright (c) 2016 by cisco Systems, Inc.
 * All rights reserved.
 */
package access_point

/* Standard packages */
import (
    "fmt"
    "log"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
)

/* Cheetah packages */
import (
    pb "gengo"
)

func ClientInit(conn *grpc.ClientConn) (int) {
    /* Setup a go-routine channel to synchronize both go-routines*/
    sync_chan := make(chan int)

    /* Setup the notification channel */
    go setupNotifChannel(conn, sync_chan)

    /* Wait for response 0: error. 1: all ok*/
    wait_resp := <- sync_chan
    if (wait_resp == 0) {
        fmt.Println("Client Error")
        return 0
    }

    /* Create a sLGlobalClient instance */
    globalClient := pb.NewAPGlobalClient(conn)

    /* Create a APGlobalsGetMsg */
    globalGetMsg := &pb.APGlobalsGetMsg {}

    /* RPC to Get the globals. */
    response, err := globalClient.APGlobalsGet(context.Background(),
        globalGetMsg)
    if (err != nil) {
        fmt.Println("Client Error %v", err)
        return 0
    }

    /* Print Server response */
    fmt.Println("Server response: ", response.ErrStatus.Status)
    fmt.Println("Max Radio Name Len  : ", response.MaxRadioNameLength)
    fmt.Println("Max Ssid Name Len   : ", response.MaxSsidNameLength)

    return wait_resp
}

func setupNotifChannel(conn *grpc.ClientConn, sync_chan chan int) {
    /* Create a sLGlobalClient instance */
    globalClient := pb.NewAPGlobalClient(conn)

    /* Create a APGlobalsGetMsg */
    initMsg := &pb.APInitMsg {
        MajorVer: uint32(pb.APVersion_AP_MAJOR_VERSION),
        MinorVer: uint32(pb.APVersion_AP_MINOR_VERSION),
        SubVer: uint32(pb.APVersion_AP_SUB_VERSION),
    }

    /* RPC to Init the notification channel */
    stream, err := globalClient.APGlobalInitNotif(context.Background(),
        initMsg)
    if err != nil {
        fmt.Println("Client Error %v", err)
        /*signal error*/
        sync_chan <- 0
        return
    }

    /* For ever read from stream */
    for {
        event, stream_err := stream.Recv()
        if (stream_err != nil) {
            fmt.Println("Client Recv Error %v", stream_err)
            break
        }

        switch event.EventType {
        case pb.APGlobalNotifType_AP_GLOBAL_EVENT_TYPE_VERSION:
            //initMsgRsp := event.GetInitRspMsg()
            /* Check Server event */
            if (event.ErrStatus.Status == pb.APErrorStatus_AP_SUCCESS) ||
               (event.ErrStatus.Status ==
                   pb.APErrorStatus_AP_INIT_STATE_CLEAR) ||
               (event.ErrStatus.Status ==
                   pb.APErrorStatus_AP_INIT_STATE_READY) {
                fmt.Printf("Server Returned %s\n",
                    event.ErrStatus.Status.String());
                /*fmt.Printf("Server Returned 0x%x, Version: %d.%d.%d\n",
                    event.ErrStatus.Status,
                    initMsgRsp.MajorVer, initMsgRsp.MinorVer, initMsgRsp.SubVer) */
                /*signal success, continue processing events from server*/
                sync_chan <- 1
            } else {
                log.Fatalf("Client Recv Error 0x%x", event.ErrStatus.Status)
            }

        case pb.APGlobalNotifType_AP_GLOBAL_EVENT_TYPE_ERROR:
            if (event.ErrStatus.Status == pb.APErrorStatus_AP_NOTIF_TERM) {
                log.Fatalf("Received notice to terminate. Client Takeover?\n")
            }
            log.Fatalf("Error not handled: 0x%x", event.ErrStatus.Status)

        case pb.APGlobalNotifType_AP_GLOBAL_EVENT_TYPE_HEARTBEAT:
            fmt.Printf("Received HeartBeat\n")

        default:
            log.Fatalf("Client Recv unknown event %s",
                event.EventType.String())
        }
    }

    log.Fatalf("Exiting")
}
