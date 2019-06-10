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

// Network Pipeline implementation for Nimbess Agent

package agent

import (
	"errors"
	"fmt"
	"sync"

	"github.com/nimbess/nimbess-agent/pkg/drivers"
	"github.com/nimbess/nimbess-agent/pkg/network"

	log "github.com/sirupsen/logrus"
)

type NimbessPipeline struct {
	Mu *sync.Mutex
	Modules []*network.Module
	Driver drivers.Driver
	EgressPorts []*network.EgressPort
}

func(p *NimbessPipeline) GetModule(name string) *network.Module {
	for _, mod := range p.Modules {
		if mod.Name == name {
			return mod
		}
	}

	return nil
}

func(p *NimbessPipeline) AddModule(module *network.Module, pModule *network.Module, nModule *network.Module) error {
	log.Infof("Adding module %v to pipeline", module)
	//Check if module already exists in pipeline
	if p.GetModule(module.Name) != nil {
		return errors.New(fmt.Sprintf("Module %s already exists in pipeline", module.Name))
	}
	if err := p.Driver.AddModule(module, pModule, nModule); err != nil {
		return err
	}

	return nil
}

// AddPort adds a port to a pipeline.
// It returns a pointer to a EgressPort module to update other pipelines with
func (s *NimbessPipeline) AddPort(port network.Port) (*network.EgressPort, error) {
	if !virtual {
		return nil, errors.New("Physical ports are not yet supported")
	}
	logging
}