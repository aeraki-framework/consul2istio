// Copyright Aeraki Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/aeraki-framework/consul2istio/pkg"
	"istio.io/pkg/log"
)

const (
	defaultConsulAddress = "127.0.0.1:8500"
)

func main() {
	consulAddress := flag.String("consulAddress", defaultConsulAddress, "Consul Address")
	flag.Parse()
	controller := pkg.NewController(*consulAddress)
	// Create the stop channel for all of the servers.
	stopChan := make(chan struct{}, 1)
	err := controller.Run(stopChan)
	if err != nil {
		log.Errorf("Fialed to run controller: %v", err)
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	stopChan <- struct{}{}
}
