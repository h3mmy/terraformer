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
	EventRuleAllowEmptyValues = []string{}
)

type EventRuleGenerator struct {
	AuthentikService
}

func (g EventRuleGenerator) createResources(eventRules []*api.EventRule) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, eventRule := range eventRules {
		resourceId := string(eventRule.Pk)
		resourceName := eventRule.EventRulename
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceId,
			resourceName,
			"authentik_event_rule",
			"authentik",
			EventRuleAllowEmptyValues,
		))
	}
	return resources
}
