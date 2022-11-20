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
	"errors"
	"os"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type AuthentikProvider struct { //nolint
	terraformutils.Provider
	url      string
	token    string
	insecure bool
}

// Initializes all supported services
func (p *AuthentikProvider) Init(args []string) error {
	akUrl := os.Getenv("AUTHENTIK_API_URL")
	if akUrl == "" {
		return errors.New("set AUTHENTIK_API_URL env var")
	}
	p.url = akUrl

	token := os.Getenv("AUTHENTIK_API_TOKEN")
	if token == "" {
		return errors.New("set AUTHENTIK_API_TOKEN env var")
	}
	p.token = token

	insecure, err := strconv.ParseBool(os.Getenv("AUTHENTIK_INSECURE"))
	if err != nil {
		return errors.New("Parsing Error for AUTHENTIK_INSECURE")
	}
	p.insecure = insecure

	return nil
}

func (p *AuthentikProvider) GetName() string {
	return "authentik"
}

func (p *AuthentikProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"url":      cty.StringVal(p.url),
		"token":    cty.StringVal(p.token),
		"insecure": cty.BoolVal(p.insecure),
	})
}

func (p *AuthentikProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"url":      p.url,
		"token":    p.token,
		"insecure": p.insecure,
	})
	return nil
}

func (p *AuthentikProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"authentik_user":                          &UserGenerator{},
		"authentik_group":                         &GroupGenerator{},
		"authentik_outpost":                       &OutpostGenerator{},
		"authentik_event_rule":                    &EventRuleGenerator{},
		"authentik_policy_binding":                &PolicyBindingGenerator{},
		"authentik_policy_expression":             &PolicyExpressionGenerator{},
		"authentik_provider_ldap":                 &ProviderLdapGenerator{},
		"authentik_provider_oauth2":               &ProviderOauth2Generator{},
		"authentik_provider_saml":                 &ProviderSamlGenerator{},
		"authentik_provider_proxy":                &ProviderProxyGenerator{},
		"authentik_service_connection_kubernetes": &ServiceConnectionKubernetesGenerator{},
		"authentik_service_connection_docker":     &ServiceConnectionDockerGenerator{},
		"authentik_source_ldap":                   &SourceLdapGenerator{},
		"authentik_source_plex":                   &SourcePlexGenerator{},
		"authentik_source_saml":                   &SourceSamlGenerator{},
		"authentik_source_oauth":                  &SourceOauthGenerator{},
	}
}

func (p AuthentikProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p AuthentikProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
