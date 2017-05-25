/*
 * Copyright (c) 2017 by cisco Systems, Inc.
 * All rights reserved.
 */
package server_util

import (
    "bufio"
    "io"
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "strconv"
    "strings"
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

func GetFieldsFromFile(filename string, keys []string) map[string]string {
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

            fields := strings.SplitN(string(line), ":", 2)
            if GetFieldIndex(keys, strings.TrimSpace(fields[0])) != -1 {
                fmap[strings.TrimSpace(fields[0])] = strings.TrimSpace(fields[1])
                count++
                if (count == len(keys)) {
                    break
                }
            }
        }
    } else {
        log.Fatal(err)
    }

    return fmap
}

func ReadFileAsString(filename string) string {
    var data []byte
    var err error

    //read file
    data, err = ioutil.ReadFile(filename)
    check(err)

    return string(data)
}

func GetIfStats(ifname string, field string) (uint64, error) {
    dir := "/sys/class/net/" + ifname + "/statistics/"

    stat, err := exec.Command("cat", dir + field).Output()
    if err != nil {
        log.Fatal(err)
    }

    sstat := strings.Trim(string(stat), "\n")
    return strconv.ParseUint(sstat, 10, 32)
}
