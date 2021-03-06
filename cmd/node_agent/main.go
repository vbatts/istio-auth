// Copyright 2017 Istio Authors
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
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"istio.io/auth/cmd/node_agent/na"
)

var (
	naConfig na.Config

	rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			runNodeAgent()
		},
	}
)

func init() {
	flags := rootCmd.Flags()

	flags.StringVar(&naConfig.ServiceIdentity, "service-identity", "", "Service Identity the node agent is managing")
	flags.StringVar(&naConfig.ServiceIdentityOrg, "org", "Juju org", "Organization for the cert")
	flags.IntVar(&naConfig.RSAKeySize, "key-size", 1024, "Size of generated private key")
	flags.StringVar(&naConfig.NodeIdentityCertFile, "na-cert", "", "Node Agent identity cert file")
	flags.StringVar(&naConfig.NodeIdentityPrivateKeyFile, "na-key", "", "Node identity private key file")
	flags.StringVar(&naConfig.IstioCAAddress, "ca-address", "127.0.0.1", "Istio CA address")
	flags.StringVar(&naConfig.ServiceIdentityDir, "cert-dir", "./", "Certificate directory")
	flags.StringVar(&naConfig.RootCACertFile, "root-cert", "", "Root Certificate file")
	flags.IntVar(&naConfig.Env, "env", na.ONPREM, "Node Environment : onprem | gcp")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Error(err)
		os.Exit(-1)
	}
}

func runNodeAgent() {
	nodeAgent := na.NewNodeAgent(&naConfig)
	glog.Infof("Starting Node Agent")
	nodeAgent.Start()
}
