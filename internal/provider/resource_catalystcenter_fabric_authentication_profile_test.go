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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcFabricAuthenticationProfile(t *testing.T) {
	if os.Getenv("FABRIC") == "" {
		t.Skip("skipping test, set environment variable FABRIC")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_authentication_profile.test", "authentication_template_name", "No Authentication"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcFabricAuthenticationProfilePrerequisitesConfig + testAccCcFabricAuthenticationProfileConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_fabric_authentication_profile.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccCcFabricAuthenticationProfilePrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}

resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  authentication_profile_name = "No Authentication"
  pub_sub_enabled             = false
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcFabricAuthenticationProfileConfig_minimum() string {
	config := `resource "catalystcenter_fabric_authentication_profile" "test" {` + "\n"
	config += `	site_name_hierarchy = "${catalystcenter_area.test.parent_name}/${catalystcenter_area.test.name}"` + "\n"
	config += `	authentication_template_name = "No Authentication"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcFabricAuthenticationProfileConfig_all() string {
	config := `resource "catalystcenter_fabric_authentication_profile" "test" {` + "\n"
	config += `	site_name_hierarchy = "${catalystcenter_area.test.parent_name}/${catalystcenter_area.test.name}"` + "\n"
	config += `	authentication_template_name = "No Authentication"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
