# TeSmart KVM Controller CLI
CLI controller for the TeSmart KVMs Network capabilities


## Usage

    kvm-controller -address <Hostname/IP> -device <KVM Device Port Number>

## Environment Variables
Both the `-address` and `-device` parameters can be controlled using
Environment Variables.

    KVM_CONTROLLER_ADDRESS = 192.168.1.10
    KVM_CONTROLLER_DEVICE = 7

## Default Value for Network Address
Omission of the `KVM_CONTROLLER_ADDRESS` variable or `-address` command line switch
will result in the use of the default IP address of the KVM, a simple convenience.

