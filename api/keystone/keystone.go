/*
Copyright 2017 The Kubernetes Authors.

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

package keystone

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
)

// Connect function authenticates in keystone using settings
// from env.
func Connect() (*gophercloud.ProviderClient, error) {
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	return openstack.AuthenticatedClient(opts)
}

// gophercloud didn't implement getResult.ExtractUser() call
// for IdentityV3 API thus we need to do all the job manually
// in order to get not only the token but the user info.
func extract(resp tokens.GetResult) (*Token, error) {
	if resp.Err != nil {
		return nil, resp.Err
	}

	obj := struct {
		Token
	}{}

	err := mapstructure.Decode(resp.Body, &obj)
	if err != nil {
		return nil, err
	}

	return &obj.Token, nil
}

// GetToken function requests information about the provided token
// from keystone. The function also returns user information
// which was encoded into the token.
// For additional details please refer to
// https://developer.openstack.org/api-ref/identity/v3/?expanded=validate-and-show-information-for-token-detail#validate-and-show-information-for-token
func GetToken(provider *gophercloud.ProviderClient, tokenID string) (*Token, error) {
	client := openstack.NewIdentityV3(provider)

	return extract(tokens.Get(client, tokenID))
}
