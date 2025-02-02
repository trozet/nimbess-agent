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

// Package bess contains the BESS data plane driver for Nimbess Agent
package bess

import (
	"fmt"

	"github.com/nimbess/nimbess-agent/pkg/drivers"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Driver represents the BESS driver for Nimbess Agent
type Driver struct {
	drivers.DriverConfig
}

// Connect is used to setup gRPC connection with the data plane.
func (d *Driver) Connect() *grpc.ClientConn {
	log.Info("Connecting to BESS")
	// TODO change this to unix socket connection
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", d.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	log.Info("Connected to BESS")
	return conn
}
