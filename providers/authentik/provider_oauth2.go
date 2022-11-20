// Copyright 2022 The Terraformer Authors.
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

package authentik

import (

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	api "goauthentik.io/api/v3"
)

var (
	ProviderOauth2AllowEmptyValues = []string{}
)

type ProviderOauth2Generator struct {
	AuthentikService
}

func (g ProviderOauth2Generator) createResources(oauth2Providers []*api.OAuth2Provider) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, providerOauth2 := range oauth2Providers {
		resourceId := string(providerOauth2.Pk)
		resourceName := providerOauth2.Name
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceId,
			resourceName,
			"authentik_provider_oauth2",
			"authentik",
			ProviderOauth2AllowEmptyValues,
		))
	}
	return resources
}
