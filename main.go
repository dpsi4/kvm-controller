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
	"os"
	"strconv"
)

const EnvAddress = "KVM_CONTROLLER_ADDRESS"
const EnvDevice = "KVM_CONTROLLER_DEVICE"

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
	var parsedDevice int
	var justExit = false

	// Address Parsing and precedence
	if *cliAddress != "" {
		parsedAddress = *cliAddress
	} else if os.Getenv(EnvAddress) != "" {
		parsedAddress = os.Getenv(EnvAddress)
	} else {
		fmt.Fprintln(os.Stderr, "No Address given to control")
		justExit = true
	}

	// Device Parsing and precedence
	if *cliDevice > 0 {
		parsedDevice = *cliDevice
	} else if os.Getenv(EnvDevice) != "" {
		if v, err := strconv.ParseInt(os.Getenv(EnvDevice), 10, 64); err == nil {
			parsedDevice = int(v)
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
	fmt.Fprintln(os.Stderr, "Parsed Address: ", parsedAddress)
	fmt.Fprintln(os.Stderr, "Parsed Device: ", parsedDevice)
}
