/*
 * Copyright (c) 2017 by cisco Systems, Inc.
 * All rights reserved.
 */
package util

// Standard packages
import (
    "fmt"
    "log"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
)

// Cheetah packages
import (
    pb "gengo"
)

/* Get SystemStats */
func SystemStatsOperation(conn *grpc.ClientConn) {

    /* Create a NewAPStatisticsClient instance */
    c := pb.NewAPStatisticsClient(conn)

    /* Create a request message */
    message := &pb.APStatsGetMsg{}

    response, err := c.APSystemStatsGet(context.Background(), message)
    if (err != nil) {
        log.Fatal(err)
    }

    if (response.ErrStatus.Status != pb.APErrorStatus_AP_SUCCESS) {
        log.Fatalf("System stats operation error: %s", response.String())
    }


    fmt.Printf("---System statistics----\n")
    fmt.Printf("AP ID       : %s\n", response.ID)
    fmt.Printf("Serial No   : %s\n", response.SerialNumber)
    fmt.Printf("Product ID  : %s\n", response.ProductId)
    fmt.Printf("Uptime      : %d\n", response.Uptime)
    fmt.Printf("When        : %s\n", response.When)
}
