# Requirements

## Goal

Build a light weight VPC(virtual private cloud) on a single Linux host using Go.

## Functional Requirements.

version1

1. create/delete VPC
2. create/delete subnets
3. hosts in a subnet can communiate with each other
4. hosts in different subnets can communicate with each other
5. create virtual hosts in subnets and allow them to run applications such as web server
6. Enable outbound internet connectivity from vittual hosts
7. isolate different VPC network

## Non-Functional Requirements

1. support idempotency(never create duplicate VPC)
2. 
## Out of Scope (version 1)
1. Multi-host networking
2. VXLAN and EVPN
3. eBPF-based packet processing
4. Inbound internet access to virtual hosts

