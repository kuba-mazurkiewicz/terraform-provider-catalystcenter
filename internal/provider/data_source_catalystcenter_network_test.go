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
func TestAccDataSourceCcNetwork(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "dhcp_servers.0", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "domain_name", "cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "primary_dns_server", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "secondary_dns_server", "1.2.3.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "syslog_servers.0", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "catalyst_center_as_syslog_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "snmp_servers.0", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "catalyst_center_as_snmp_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "netflow_collector", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "netflow_collector_port", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "ntp_servers.0", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network.test", "timezone", "Europe/Vienna"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcNetworkPrerequisitesConfig + testAccDataSourceCcNetworkConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

//template:end testAccDataSource

//template:begin testPrerequisites
const testAccDataSourceCcNetworkPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}

`

//template:end testPrerequisites

//template:begin testAccDataSourceConfig
func testAccDataSourceCcNetworkConfig() string {
	config := `resource "catalystcenter_network" "test" {` + "\n"
	config += `	site_id = catalystcenter_area.test.id` + "\n"
	config += `	dhcp_servers = ["1.2.3.4"]` + "\n"
	config += `	domain_name = "cisco.com"` + "\n"
	config += `	primary_dns_server = "1.2.3.4"` + "\n"
	config += `	secondary_dns_server = "1.2.3.5"` + "\n"
	config += `	syslog_servers = ["1.2.3.4"]` + "\n"
	config += `	catalyst_center_as_syslog_server = false` + "\n"
	config += `	snmp_servers = ["1.2.3.4"]` + "\n"
	config += `	catalyst_center_as_snmp_server = false` + "\n"
	config += `	netflow_collector = "1.2.3.4"` + "\n"
	config += `	netflow_collector_port = 1234` + "\n"
	config += `	ntp_servers = ["1.2.3.4"]` + "\n"
	config += `	timezone = "Europe/Vienna"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_network" "test" {
			id = catalystcenter_network.test.id
			site_id = catalystcenter_area.test.id
		}
	`
	return config
}

//template:end testAccDataSourceConfig