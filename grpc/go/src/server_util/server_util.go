/*
 * Copyright (c) 2017 by cisco Systems, Inc.
 * All rights reserved.
 */
package server_util

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	context "golang.org/x/net/context"
	"google.golang.org/grpc/peer"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetFieldIndex(keys []string, field string) int {
	var index int

	index = -1
	for i, v := range keys {
		if v == field {
			index = i
			break
		}
	}

	return index
}

func GetFieldsFromFile(filename string, keys []string, sep string) map[string]string {
	var count int
	var fmap map[string]string

	// open a file
	if file, err := os.Open(filename); err == nil {
		// make sure it gets closed
		defer file.Close()

		// create a new reader and read the file line by line
		bf := bufio.NewReader(file)
		fmap = make(map[string]string)
		for {
			line, isPrefix, errno := bf.ReadLine()

			// loop termination - EOF
			if errno == io.EOF {
				break
			}

			// loop termination - error
			if errno != nil {
				log.Fatal(errno)
			}

			// loop termination - error
			if isPrefix {
				log.Fatal("Error: Unexpected long line reading", file.Name())
			}

			fields := strings.SplitN(string(line), sep, 2)
			if GetFieldIndex(keys, strings.TrimSpace(fields[0])) != -1 {
				fmap[strings.TrimSpace(fields[0])] = strings.TrimSpace(fields[1])
				count++
				if count == len(keys) {
					break
				}
			}
		}
	} else {
		return nil
	}

	return fmap
}

func ReadFileAsString(filename string) string {
	var data []byte
	var err error

	//read file
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}

	return string(data)
}

func GetIfStats(ifname string, field string) (uint64, error) {
	dir := "/sys/class/net/" + ifname + "/statistics/"

	stat, err := exec.Command("cat", dir+field).Output()
	if err != nil {
		log.Fatal(err)
	}

	sstat := strings.Trim(string(stat), "\n")
	return strconv.ParseUint(sstat, 10, 32)
}

func FD_SET(p *syscall.FdSet, i int) {
	p.Bits[i/64] |= 1 << uint(i) % 64
}

func FD_ISSET(p *syscall.FdSet, i int) bool {
	return (p.Bits[i/64] & (1 << uint(i) % 64)) != 0
}

func FD_ZERO(p *syscall.FdSet) {
	for i := range p.Bits {
		p.Bits[i] = 0
	}
}

/*
 * Get address from context
 */
func GetAddressFromCtx(ctx context.Context) (string, bool) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", false
	}

	if pr.Addr == net.Addr(nil) {
		return "", false
	}

	return pr.Addr.String(), true
}

func GetWiredInterfaceName() string {
	ifname := os.Getenv("WIRED_IF")
	if ifname == "" {
		ifname = "wired0" /* default wired interface */
	} else {
		fmt.Printf("Using WIRED INTERFACE: %s\n", ifname)
	}

	return ifname
}

func GetCaptureInterfaceName() string {
	ifname := os.Getenv("CAPTURE_IF")
	if ifname == "" {
		ifname = "aptrace0" /* default capture interface */
	} else {
		fmt.Printf("Using CAPTURE INTERFACE: %s\n", ifname)
	}

	return ifname
}
