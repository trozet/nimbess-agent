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
	typePipeline = reflect.TypeOf(Pipeline{})
)

// We cant directly use NimbessPipeline here because it contains pointer references
// We need to translate NimbessPipeline into an etcd pipeline we can store
type Pipeline struct {
	Name        string
	Modules     []network.PipelineModule
	MetaKey     string
}

type PipelineKey struct {
	AgentKey
	Name string
}

func (key PipelineKey) defaultDeletePath() (string, error) {
	return key.defaultPath()
}

func (key PipelineKey) defaultPath() (string, error) {
	if key.Name == "" {
		return "", errors.ErrorInsufficientIdentifiers{Name: "Name"}
	}
	return fmt.Sprintf("/nimbess/agents/%s/pipelines/%s", key.MachineID, key.Name), nil
}

func (key PipelineKey) valueType() (reflect.Type, error) {
	return typePipeline, nil
}

func (key PipelineKey) String() string {
	return fmt.Sprintf("Name(name=%s)", key.Name)
}

func (key PipelineKey) KeyToDefaultDeletePath() (string, error) {
	return key.defaultPath()
}
