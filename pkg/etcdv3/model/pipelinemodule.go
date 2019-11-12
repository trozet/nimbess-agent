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
	v1 "github.com/nimbess/stargazer/pkg/crd/api/unp/v1"
	"reflect"
	"strings"

	"github.com/nimbess/stargazer/pkg/errors"
)

var (
	typePipelineModule = reflect.TypeOf(PipelineModule{})
	typeIngressPort = reflect.TypeOf(IngressPort{})
	typeSwitch = reflect.TypeOf(Switch{})
	typeURLFilter = reflect.TypeOf(URLFilter{})
)

// Represents base modules who are members of a pipeline and cannot be connected to any other pipelines
type PipelineModule network.Module

// Pipeline specific module types
type IngressPort struct {
	PipelineModule
	PortName string
}
// We dont store L2FIB in etcd Switch model because we treat the L2FIB of the meta pipeline as the single source of
// truth
type Switch PipelineModule

type URLFilter struct {
	PipelineModule
	v1.URLFilter
}

type PipelineModuleKey struct {
	PipelineKey
	ModuleName string
}

type MetaPipelineModuleKey PipelineModuleKey

func (key PipelineModuleKey) defaultDeletePath() (string, error) {
	return key.defaultPath()
}

func (key PipelineModuleKey) defaultPath() (string, error) {
	if key.Name == "" {
		return "", errors.ErrorInsufficientIdentifiers{Name: "Name"}
	}
	return fmt.Sprintf("/nimbess/agents/%s/pipelines/%s/modules/%s", key.MachineID, key.Name, key.ModuleName), nil
}

func (key MetaPipelineModuleKey) defaultPath() (string, error) {
	if key.Name == "" {
		return "", errors.ErrorInsufficientIdentifiers{Name: "Name"}
	}
	return fmt.Sprintf("/nimbess/agents/%s/metapipelines/%s/modules/%s", key.MachineID, key.Name, key.ModuleName), nil
}

func (key PipelineModuleKey) valueType() (reflect.Type, error) {
	// TODO fix this in the future
	// Only way to really tell what kind of module this is via name, which for now is unique, but may not be in future
	n := key.ModuleName
	moduleNameMap := map[string]reflect.Type {
		"Switch": typeSwitch,
		"ingress": typeIngressPort,
		"URLFilter": typeURLFilter,
	}
	for modName, modType := range moduleNameMap {
		if strings.Contains(n, modName) {
			return modType, nil
		}
	}

	return typePipelineModule, fmt.Errorf("unknown mod type %s", key.ModuleName)
}

func (key PipelineModuleKey) String() string {
	return fmt.Sprintf("Name(name=%s)", key.Name)
}

func (key PipelineModuleKey) KeyToDefaultDeletePath() (string, error) {
	return key.defaultPath()
}

