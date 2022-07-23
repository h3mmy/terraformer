// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"errors"
	"os"

	authentik_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/authentik"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

// Initializes Authentik Provider
func newCmdAuthentikImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authentik",
		Short: "Import current state to Terraform configuration from Authentik",
		Long:  "Import current state to Terraform configuration from Authentik",
		RunE: func(cmd *cobra.Command, args []string) error {
			url := os.Getenv("AUTHENTIK_API_URL")
			if len(url) == 0 {
				return errors.New("API Url for Authentik must be set through `AUTHENTIK_API_URL` env var")
			}
			token := os.Getenv("AUTHENTIK_API_TOKEN")
			if len(token) == 0 {
				return errors.New("API Token for Auht0 must be set through `AUTHENTIK_API_TOKEN` env var")
			}
			insecure := os.Getenv("AUTHENTIK_INSECURE")
			if len(insecure) == 0 {
				return errors.New("Authentik Insecure Flag for Auhthentik must be set through `AUTHENTIK_INSECURE` env var")
			}

			provider := newAuthentikProvider()
			err := Import(provider, options, []string{url, token, token})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newAuthentikProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "action", "action=name1:name2:name3")
	return cmd
}

func newAuthentikProvider() terraformutils.ProviderGenerator {
	return &authentik_terraforming.AuthentikProvider{}
}
