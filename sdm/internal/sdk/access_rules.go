// Copyright 2020 StrongDM Inc
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
//

// This file was generated by protogen. DO NOT EDIT.
package sdm

// AccessRules is exposed as a string to Terraform to make the transformation as transparent as possible.
// The Go SDK uses structures to expose AccessRules as idiomatically as possible.
type AccessRules string

func convertAccessRulesToPorcelain(rules string) (AccessRules, error) {
	return AccessRules(rules), nil
}

func convertAccessRulesToPlumbing(rules AccessRules) string {
	return string(rules)
}
