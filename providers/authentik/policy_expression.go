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
	PolicyExpressionAllowEmptyValues = []string{}
)

type PolicyExpressionGenerator struct {
	AuthentikService
}

func (g PolicyExpressionGenerator) createResources(policyExpressions []*api.ExpressionPolicy) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, policyExpression := range policyExpressions {
		resourceId := string(policyExpression.Pk)
		resourceName := policyExpression.VerboseName
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceId,
			resourceName,
			"authentik_policy_expression",
			"authentik",
			PolicyExpressionAllowEmptyValues,
		))
	}
	return resources
}
