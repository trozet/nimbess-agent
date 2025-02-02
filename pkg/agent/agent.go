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

// Agent runtime

package agent

import (
	"fmt"
	"net"
	"sync"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/nimbess/nimbess-agent/pkg/drivers"
	"github.com/nimbess/nimbess-agent/pkg/proto/cni"
)

// NimbessAgent represents the agent runtime server.
// It contains a loadable runtime data plane driver to manage the data plane.
// It includes a mutex used to handle locking between driver and agent events to force a
// single-processed event pipeline.
type NimbessAgent struct {
	Mu     *sync.Mutex
	Config *NimbessConfig
	Driver drivers.Driver
}

// Add implements CNI Add Handler.
// It returns a CNI Reply to be sent to the Nimbess CNI client.
func (s *NimbessAgent) Add(ctx context.Context, req *cni.CNIRequest) (*cni.CNIReply, error) {
	//TODO implement
	return &cni.CNIReply{}, nil
}

// Delete implements CNI Delete Handler.
// It returns a CNI Reply to be sent to the Nimbess CNI client.
func (s *NimbessAgent) Delete(ctx context.Context, req *cni.CNIRequest) (*cni.CNIReply, error) {
	//TODO implement
	return &cni.CNIReply{}, nil
}

// Run starts up the main Agent daemon.
func (s *NimbessAgent) Run() error {
	log.Info("Starting Nimbess Agent...")
	log.Info("Connecting to Data Plane")
	dpConn := s.Driver.Connect()
	defer dpConn.Close()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.Port))
	if err != nil {
		log.Errorf("Failed to listen on port: %d", s.Config.Port)
		return err
	}
	log.Info("Starting Nimbess gRPC server...")
	grpcServer := grpc.NewServer()
	cni.RegisterRemoteCNIServer(grpcServer, s)
	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
