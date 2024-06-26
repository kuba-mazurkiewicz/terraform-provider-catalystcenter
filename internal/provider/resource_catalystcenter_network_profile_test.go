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
func TestAccCcNetworkProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_network_profile.test", "name", "Profile1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_network_profile.test", "type", "switching"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_network_profile.test", "templates.0.type", "cli.templates"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcNetworkProfilePrerequisitesConfig + testAccCcNetworkProfileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcNetworkProfilePrerequisitesConfig + testAccCcNetworkProfileConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_network_profile.test",
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
const testAccCcNetworkProfilePrerequisitesConfig = `
resource "catalystcenter_project" "test" {
  name        = "Project1"
}

resource "catalystcenter_template" "test" {
  project_id  = catalystcenter_project.test.id
  name        = "Template1"
  description = "My description"
  device_types = [
    {
      product_family = "Switches and Hubs"
      product_series = "Cisco Catalyst 9300 Series Switches"
      product_type   = "Cisco Catalyst 9300 Switch"
    }
  ]
  language         = "JINJA"
  software_type    = "IOS-XE"
  software_variant = "XE"
  software_version = "16.12.1a"
  template_content = "hostname SW1"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcNetworkProfileConfig_minimum() string {
	config := `resource "catalystcenter_network_profile" "test" {` + "\n"
	config += `	name = "Profile1"` + "\n"
	config += `	type = "switching"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcNetworkProfileConfig_all() string {
	config := `resource "catalystcenter_network_profile" "test" {` + "\n"
	config += `	name = "Profile1"` + "\n"
	config += `	type = "switching"` + "\n"
	config += `	templates = [{` + "\n"
	config += `	  type = "cli.templates"` + "\n"
	config += `	  attributes = [{` + "\n"
	config += `		template_id = catalystcenter_template.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
