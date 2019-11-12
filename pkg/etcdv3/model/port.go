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

package model

import (
	"fmt"
	"github.com/nimbess/nimbess-agent/pkg/network"
	"reflect"

	"github.com/nimbess/stargazer/pkg/errors"
)
var (
	typePort = reflect.TypeOf(network.Port{})
)

type PortKey struct {
	AgentKey
	Name string
}

func (key PortKey) defaultDeletePath() (string, error) {
	return key.defaultPath()
}

func (key PortKey) defaultPath() (string, error) {
	if key.Name == "" {
		return "", errors.ErrorInsufficientIdentifiers{Name: "Name"}
	}
	return fmt.Sprintf("/nimbess/agents/%s/ports/%s", key.MachineID, key.Name), nil
}

func (key PortKey) valueType() (reflect.Type, error) {
	return typePipeline, nil
}

func (key PortKey) String() string {
	return fmt.Sprintf("Name(name=%s)", key.Name)
}

func (key PortKey) KeyToDefaultDeletePath() (string, error) {
	return key.defaultPath()
}
