// Copyright 2020 Artificial Intelligence Redefined <dev+cogment@ai-r.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	"fmt"
	"github.com/cogment/cogment-cli/deployment"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

func runDeleteRegistryCmd(client *resty.Client, registryId string) error {
	resp, err := client.R().Delete("/docker-registries/" + registryId)

	if err != nil {
		log.Fatalf("%v", err)
	}

	if http.StatusNotFound == resp.StatusCode() {
		return fmt.Errorf("%s", "Registry ID not found")
	}

	return nil
}

func NewRegistryDeleteCommand(verbose bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete configured Docker Registry",
		Args:  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client, err := deployment.PlatformClient(verbose)
			if err != nil {
				log.Fatal(err)
			}

			err = runDeleteRegistryCmd(client, args[0])

			if err != nil {
				log.Fatalf("%v", err)
			}

			fmt.Println("Registry deleted")
		},
	}

	return cmd
}
