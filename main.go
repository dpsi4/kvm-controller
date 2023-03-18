//
// main.go
// Copyright (C) 2023 Reid Miller <reidm@dpsi4.net>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const EnvAddress = "KVM_CONTROLLER_ADDRESS"
const EnvDevice = "KVM_CONTROLLER_DEVICE"
const KvmPort = 5000
const KvmAddress = "192.168.1.10"

var address, device string

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s for TeSmart KVMs", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Copyright 2023")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}
	var cliAddress = flag.String("address", "", "Address or Hostname of KVM to control")
	var cliDevice = flag.Int("device", -1, "Device number to control")
	flag.Parse()

	var parsedAddress string
	var parsedDevice int8
	var justExit = false

	// Address Parsing and precedence
	if *cliAddress != "" {
		parsedAddress = *cliAddress
	} else if os.Getenv(EnvAddress) != "" {
		parsedAddress = os.Getenv(EnvAddress)
	} else {
		parsedAddress = KvmAddress
	}

	// Device Parsing and precedence
	if *cliDevice > 0 {
		parsedDevice = int8(*cliDevice)
	} else if os.Getenv(EnvDevice) != "" {
		if v, err := strconv.ParseInt(os.Getenv(EnvDevice), 10, 8); err == nil {
			parsedDevice = int8(v)
		} else {
			fmt.Fprintln(os.Stderr, "Unable to parse DEVICE from Environment Variable!")
			justExit = true
		}
	} else {
		fmt.Fprintln(os.Stderr, "No Device Port given to control")
		justExit = true
	}

	if justExit {
		os.Exit(1)
	}

	var connection = net.JoinHostPort(parsedAddress, fmt.Sprintf("%d", KvmPort))
	kvmConnection, err := net.DialTimeout("tcp", connection, 1*time.Second)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Timeout during connect")
		os.Exit(1)
	}

	// Issue command to the KVM
	command := []byte{0xAA, 0xBB, 0x03, 0x01, byte(parsedDevice), 0xEE}
	kvmConnection.Write(command)
	kvmConnection.Close()
}
