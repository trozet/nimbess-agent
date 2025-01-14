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

// Data Plane driver definition

package drivers

import (
	"google.golang.org/grpc"
)

// Driver represents an abstract data plane driver type.
type Driver interface {
	// TODO define more methods for a generic Data Plane driver
	Connect() *grpc.ClientConn
}

// DriverConfig represents the generic driver configuration required by Driver.
type DriverConfig struct {
	NetworkMode string
	MacLearn    bool
	TunnelMode  bool
	FIBSize     int64
	Port        int
	PCIDevices  []string
	WorkerCores []int64
}
