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

package agent
import (
	log "github.com/sirupsen/logrus"
)

// Purpose of dbsync is to manage synchronizing DB -> agent runtime on startup
// After successful a successful startup + dbsync, the runtime memory is considered to be the source of truth
// and consequent resource modifications occur first in runtime and then written to DB

func (s *NimbessAgent) dbSync() error {
	// need to list resources from bottom up
	// check resources exist in Agent
	// check resources exist in dataplane
	// if DNE create resource in agent and DP (via single agent call)
	log.Info("Initiating Agent sync from DB")
	return nil
}