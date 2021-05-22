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
	"errors"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/aeraki-framework/consul2istio/pkg/serviceregistry/consul"

	"github.com/aeraki-framework/consul2istio/pkg"

	_ "time/tzdata"

	"istio.io/pkg/log"
)

const (
	sitConsulAddress     = "10.4.45.40:8500"
	defaultConsulAddress = "localhost:30395"
)

var consulAddress string

func init() {
	if err := parseFlags(); err != nil {
		log.Fatalf("parse flags error: %v", err)
	}
}

func main() {
	server := pkg.NewServer(consulAddress)

	// Create the stop channel for all of the servers.
	stopChan := make(chan struct{}, 1)
	if err := server.Run(stopChan); err != nil {
		log.Errorf("Failed to run controller: %v", err)
		return
	}

	// server graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	stopChan <- struct{}{}
}

func parseFlags() error {
	flag.StringVar(&consulAddress, "consulAddress", sitConsulAddress, "Consul Address")
	flag.StringVar(&consul.ServiceZone, "consulZone", "", "Consul Data Center")
	flag.Parse()

	if consul.ServiceZone == "" {
		return errors.New("please set consul zone with --consulZone parameter")
	}

	return nil
}
