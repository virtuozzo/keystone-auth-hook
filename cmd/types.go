/*
Copyright (c) 2016-2017 Parallels International GmbH.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import "encoding/json"

// Payload represents request and response package of
// the authentication webhook.
type Payload struct {
	// k8s API version. For additional details visit
	// https://kubernetes.io/docs/concepts/overview/kubernetes-api/
	APIVersion string `json:"apiVersion"`

	// Current operation name.
	Kind string `json:"kind"`

	Spec   Spec   `json:"spec"`
	Status Status `json:"status"`
}

// Spec represents request parameters.
type Spec struct {
	// Token issued by keystone.
	Token string `json:"token"`
}

// Status represents auth operation result.
type Status struct {
	// Authentication result.
	Authenticated bool `json:"authenticated"`

	User User `json:"user"`
}

// Provides basic user information required for
// successful authentication.
type User struct {
	Username string   `json:"username,omitempty"`
	UID      string   `json:"uid,omitempty"`
	Groups   []string `json:"groups,omitempty"`
}

func (p *Payload) String() string {
	b, _ := json.MarshalIndent(p, "", "  ")
	return string(b)
}
