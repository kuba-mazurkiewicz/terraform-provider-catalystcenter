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

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcUser(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_user.test", "first_name", "john"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_user.test", "last_name", "doe"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_user.test", "username", "johndoe"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_user.test", "email", "john.doe@example.com"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcUserPrerequisitesConfig + testAccDataSourceCcUserConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcUserPrerequisitesConfig = `
data "catalystcenter_role" "test" {
  name = "OBSERVER-ROLE"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcUserConfig() string {
	config := `resource "catalystcenter_user" "test" {` + "\n"
	config += `	first_name = "john"` + "\n"
	config += `	last_name = "doe"` + "\n"
	config += `	username = "johndoe"` + "\n"
	config += `	password = "C1sco1357"` + "\n"
	config += `	email = "john.doe@example.com"` + "\n"
	config += `	role_ids = [data.catalystcenter_role.test.id]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_user" "test" {
			id = catalystcenter_user.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
