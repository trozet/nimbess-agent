// Copyright (c) 2019 Red Hat and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Modules define network function types

// Package network contains all rendered network oriented resources for Nimbess
package network

// Gate is a virtual port for a module for traffic to ingress or egress
type Gate uint64

// Module represents a generic Network Function.
// It contains ingress gates and egress gates for traffic flow, which map to the next connected module.
type Module struct {
	Id     string
	Name   string
	IGates map[Gate]Module
	EGates map[Gate]Module
}

// Switch represents an L2 network switch.
// Extends Module and defines L2FIB table which holds a map of map to output Gates
type Switch struct {
	Module
	L2FIB map[string]Gate
}

type Port struct {
	Virtual bool
	DPDK	bool
	Name string
	NamesSpace string
	IpAddr	string
}

type IngressPort struct {
	Module
	Port
}

type EgressPort struct {
	Module
	Port
}