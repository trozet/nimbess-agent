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

package etcdv3

import (
	"context"
	"github.com/nimbess/nimbess-agent/pkg/etcdv3/model"
	"github.com/nimbess/nimbess-agent/pkg/network"
)

//CreatePort creates a DB entry for a port
func CreatePort(c Client, ctx context.Context, port network.Port, machineID string) error {
	portKey := model.PortKey{Name: port.PortName, AgentKey: model.AgentKey{MachineID: machineID}}
	portKV := &model.KVPair{Key: portKey, Value: port}
	return c.Create(ctx, portKV)
}

func DeletePort(c Client, ctx context.Context, port string, machineID string) error {
	portKey := model.PortKey{Name: port, AgentKey: model.AgentKey{MachineID: machineID}}
	return c.Delete(ctx, portKey)
}
