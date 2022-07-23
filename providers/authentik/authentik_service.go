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
	"fmt"
	"net/http"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	httptransport "github.com/go-openapi/runtime/client"
	api "goauthentik.io/api/v3"
)

type AuthentikService struct { //nolint
	terraformutils.Service
}

// APIClient Hold the API Client and any relevant configuration
type APIClient struct {
	client *api.APIClient
}

func (s *AuthentikService) generateClient() *APIClient {

	apiURL := s.Args["url"].(string)
	token := s.Args["token"].(string)
	insecure := s.Args["insecure"].(bool)

	akURL, err := url.Parse(apiURL)
	if err != nil {
		panic(err)
	}

	config := api.NewConfiguration()
	config.UserAgent = fmt.Sprintf("authentik-terraformer")
	config.Host = akURL.Host
	config.Scheme = akURL.Scheme

	config.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	config.HTTPClient = &http.Client{
		Transport: GetTLSTransport(insecure),
	}
	apiClient := api.NewAPIClient(config)


	return &APIClient{
		client: apiClient,
	}
}

// GetTLSTransport Get a TLS transport instance, that skips verification if configured via environment variables.
func GetTLSTransport(insecure bool) http.RoundTripper {
	tlsTransport, err := httptransport.TLSTransport(httptransport.TLSClientOptions{
		InsecureSkipVerify: insecure,
	})
	if err != nil {
		panic(err)
	}
	return tlsTransport
}
