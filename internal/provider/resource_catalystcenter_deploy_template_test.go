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
func TestAccCcDeployTemplate(t *testing.T) {
	if os.Getenv("TEMPLATES") == "" {
		t.Skip("skipping test, set environment variable TEMPLATES")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_deploy_template.test", "force_push_template", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_deploy_template.test", "copying_config", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_deploy_template.test", "is_composite", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_deploy_template.test", "target_info.0.host_name", "SW01"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_deploy_template.test", "target_info.0.type", "MANAGED_DEVICE_HOSTNAME"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcDeployTemplatePrerequisitesConfig + testAccCcDeployTemplateConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccCcDeployTemplatePrerequisitesConfig = `
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

resource "catalystcenter_template_version" "example" {
  template_id = catalystcenter_template.test.id
  comments    = "New Version"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcDeployTemplateConfig_minimum() string {
	config := `resource "catalystcenter_deploy_template" "test" {` + "\n"
	config += `	template_id = catalystcenter_template_version.example.id` + "\n"
	config += `	target_info = [{` + "\n"
	config += `	  type = "MANAGED_DEVICE_HOSTNAME"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcDeployTemplateConfig_all() string {
	config := `resource "catalystcenter_deploy_template" "test" {` + "\n"
	config += `	template_id = catalystcenter_template_version.example.id` + "\n"
	config += `	force_push_template = false` + "\n"
	config += `	copying_config = true` + "\n"
	config += `	is_composite = false` + "\n"
	config += `	target_info = [{` + "\n"
	config += `	  host_name = "SW01"` + "\n"
	config += `	  type = "MANAGED_DEVICE_HOSTNAME"` + "\n"
	config += `	  versioned_template_id = catalystcenter_template_version.example.id` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
