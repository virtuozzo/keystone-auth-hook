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

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"

	"k8s.io/keystone-auth-hook/api/keystone"
)

// authWebhook responds to authentication requests from K8s
type authWebhook struct {
}

// Create new authWebhook with requested connection parameters.
func newAuthWebhook() *authWebhook {
	return &authWebhook{}
}

// ServeHTTP verifies the incoming token in keystone and
// sends user's info if the token is valid.
func (hook *authWebhook) ServeHTTP(dst http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		dst.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	resp := &Payload{}
	err := json.NewDecoder(req.Body).Decode(resp)
	if err != nil {
		glog.Errorf("Unable to decode the request: %s", err)
		dst.WriteHeader(http.StatusInternalServerError)
		return
	}
	glog.Infof("Incoming request: %s", resp)
	defer req.Body.Close()

	provider, err := keystone.Connect()
	if err != nil {
		glog.Errorf("Unable to connect to Keystone: %s", err)
		dst.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := keystone.GetToken(provider, resp.Spec.Token)
	if err != nil {
		glog.Errorf("Failed to verify the token: %s", err)
		dst.WriteHeader(http.StatusUnauthorized)
		return
	}

	resp.Status.Authenticated = true
	resp.Status.User.Username = token.User.Name
	resp.Status.User.UID = token.User.ID

	glog.Infof("Outgoing response: %s", resp)

	dst.Header().Set("Content-Type", "application/json")
	json.NewEncoder(dst).Encode(resp)
	// FIXME: check encoding error?
}
