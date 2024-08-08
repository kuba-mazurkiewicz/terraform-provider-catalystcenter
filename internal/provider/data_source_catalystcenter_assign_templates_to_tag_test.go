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
func TestAccDataSourceCcAssignTemplatesToTag(t *testing.T) {
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcAssignTemplatesToTagPrerequisitesConfig + testAccDataSourceCcAssignTemplatesToTagConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcAssignTemplatesToTagPrerequisitesConfig = `
resource "catalystcenter_tag" "test" {
  name        = "Tag1"
  description = "Tag1 Description"
  system_tag  = false
}

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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcAssignTemplatesToTagConfig() string {
	config := `resource "catalystcenter_assign_templates_to_tag" "test" {` + "\n"
	config += `	tag_id = catalystcenter_tag.test.id` + "\n"
	config += `	template_ids = [catalystcenter_template.test.id]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_assign_templates_to_tag" "test" {
			tag_id = catalystcenter_tag.test.id
			depends_on = [catalystcenter_assign_templates_to_tag.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
