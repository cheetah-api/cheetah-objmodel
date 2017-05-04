/*
 * Copyright (c) 2017 by cisco Systems, Inc.
 * All rights reserved.
 */
package util

// Standard packages
import (
    "fmt"
    "io"
    "log"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
    "github.com/golang/protobuf/jsonpb"
    "golang.org/x/net/websocket"
    "os"
)

// Cheetah packages
import (
    pb "gengo"
)

/* Get tats */
func StatsOperation(conn *grpc.ClientConn, wss_conn *websocket.Conn,
                    stats_type pb.APStatsType, time_interval uint32) {

    /* Create a NewAPStatisticsClient instance */
    c := pb.NewAPStatisticsClient(conn)

    /* Create a request */
    req := new(pb.APStatsRequest)
    req.StatsType = stats_type
    req.TimeInterval =  time_interval

    /* Create a request message */
    message := &pb.APStatsMsg{ }
    message.StatsRequest = []*pb.APStatsRequest{}
    message.StatsRequest = append(message.StatsRequest, req)

    stream, err := c.APStatsGet(context.Background(), message)
    if (err != nil) {
        fmt.Println("RPC Error %v", err)
        log.Fatal(err)
    }

    for {
        response, stream_err := stream.Recv()
        if stream_err != nil {
            if stream_err != io.EOF {
                fmt.Println("Client Recv Error %v", stream_err)
            }
            break
        }

        if response == nil {
            log.Println("No response from server")
            break
        }

        if response.ErrStatus == nil {
            log.Println("Null errStatus")
            break
        }

        if (response.ErrStatus.Status != pb.APErrorStatus_AP_SUCCESS) {
            //log.Println("Stats operation error: %s", response.String())
            continue
        }
        if (response.ErrStatus.Status == pb.APErrorStatus_AP_EINVAL) ||
           (response.ErrStatus.Status == pb.APErrorStatus_AP_NOT_AVAILABLE) {
            //log.Println("Stats operation error: %s", response.String())
            continue
        }

        marshaller := jsonpb.Marshaler{}

        data, _ := marshaller.MarshalToString(response)

        if (wss_conn != nil) {
            err = websocket.JSON.Send(wss_conn, data)
            if err != nil {
                fmt.Printf("Send failed: %s\n", err.Error())
                os.Exit(1)
            }
        } else {
            log.Println(data)
        }
    }
}
