// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//template:end imports

//template:begin testAccDataSource
func TestAccDataSourceCcCredentialsSNMPv3(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_credentials_snmpv3.test", "description", "My SNMPv3 credentials"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_credentials_snmpv3.test", "username", "user1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_credentials_snmpv3.test", "privacy_type", "AES128"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_credentials_snmpv3.test", "auth_type", "SHA"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_credentials_snmpv3.test", "snmp_mode", "AUTHPRIV"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcCredentialsSNMPv3Config(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

//template:end testAccDataSource

//template:begin testPrerequisites
//template:end testPrerequisites

//template:begin testAccDataSourceConfig
func testAccDataSourceCcCredentialsSNMPv3Config() string {
	config := `resource "catalystcenter_credentials_snmpv3" "test" {` + "\n"
	config += `	description = "My SNMPv3 credentials"` + "\n"
	config += `	username = "user1"` + "\n"
	config += `	privacy_type = "AES128"` + "\n"
	config += `	privacy_password = "password1"` + "\n"
	config += `	auth_type = "SHA"` + "\n"
	config += `	auth_password = "password1"` + "\n"
	config += `	snmp_mode = "AUTHPRIV"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_credentials_snmpv3" "test" {
			id = catalystcenter_credentials_snmpv3.test.id
		}
	`
	return config
}

//template:end testAccDataSourceConfig