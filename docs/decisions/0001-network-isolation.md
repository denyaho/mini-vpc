# ADR-0001: Network Isolation

## Status
Proposed

## Context
This project is going to craete VPC with two isolated subnets.
Want to avoid overlapping IPs

## Options
- Linux Network Namespace
- Linux VRF
- Open vSwitch

## Decision
Use Linux Network Namespace

## Rationale
- Suitable for studying Linux Networking
- Supports isolated routing tables