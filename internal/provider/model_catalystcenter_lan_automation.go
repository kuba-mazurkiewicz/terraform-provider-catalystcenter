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
	"context"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type LANAutomation struct {
	Id                                types.String           `tfsdk:"id"`
	DiscoveredDeviceSiteNameHierarchy types.String           `tfsdk:"discovered_device_site_name_hierarchy"`
	PrimaryDeviceManagementIpAddress  types.String           `tfsdk:"primary_device_management_ip_address"`
	PeerDeviceManagementIpAddress     types.String           `tfsdk:"peer_device_management_ip_address"`
	PrimaryDeviceInterfaceNames       types.Set              `tfsdk:"primary_device_interface_names"`
	IpPools                           []LANAutomationIpPools `tfsdk:"ip_pools"`
	MulticastEnabled                  types.Bool             `tfsdk:"multicast_enabled"`
	HostNamePrefix                    types.String           `tfsdk:"host_name_prefix"`
	HostNameFileId                    types.String           `tfsdk:"host_name_file_id"`
	IsisDomainPassword                types.String           `tfsdk:"isis_domain_password"`
	RedistributeIsisToBgp             types.Bool             `tfsdk:"redistribute_isis_to_bgp"`
}

type LANAutomationIpPools struct {
	IpPoolName types.String `tfsdk:"ip_pool_name"`
	IpPoolRole types.String `tfsdk:"ip_pool_role"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data LANAutomation) getPath() string {
	return "/dna/intent/api/v1/lan-automation"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data LANAutomation) toBody(ctx context.Context, state LANAutomation) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DiscoveredDeviceSiteNameHierarchy.IsNull() {
		body, _ = sjson.Set(body, "0.discoveredDeviceSiteNameHierarchy", data.DiscoveredDeviceSiteNameHierarchy.ValueString())
	}
	if !data.PrimaryDeviceManagementIpAddress.IsNull() {
		body, _ = sjson.Set(body, "0.primaryDeviceManagmentIPAddress", data.PrimaryDeviceManagementIpAddress.ValueString())
	}
	if !data.PeerDeviceManagementIpAddress.IsNull() {
		body, _ = sjson.Set(body, "0.peerDeviceManagmentIPAddress", data.PeerDeviceManagementIpAddress.ValueString())
	}
	if !data.PrimaryDeviceInterfaceNames.IsNull() {
		var values []string
		data.PrimaryDeviceInterfaceNames.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.primaryDeviceInterfaceNames", values)
	}
	if len(data.IpPools) > 0 {
		body, _ = sjson.Set(body, "0.ipPools", []interface{}{})
		for _, item := range data.IpPools {
			itemBody := ""
			if !item.IpPoolName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipPoolName", item.IpPoolName.ValueString())
			}
			if !item.IpPoolRole.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipPoolRole", item.IpPoolRole.ValueString())
			}
			body, _ = sjson.SetRaw(body, "0.ipPools.-1", itemBody)
		}
	}
	if !data.MulticastEnabled.IsNull() {
		body, _ = sjson.Set(body, "0.multicastEnabled", data.MulticastEnabled.ValueBool())
	}
	if !data.HostNamePrefix.IsNull() {
		body, _ = sjson.Set(body, "0.hostNamePrefix", data.HostNamePrefix.ValueString())
	}
	if !data.HostNameFileId.IsNull() {
		body, _ = sjson.Set(body, "0.hostNameFileId", data.HostNameFileId.ValueString())
	}
	if !data.IsisDomainPassword.IsNull() {
		body, _ = sjson.Set(body, "0.isisDomainPwd", data.IsisDomainPassword.ValueString())
	}
	if !data.RedistributeIsisToBgp.IsNull() {
		body, _ = sjson.Set(body, "0.redistributeIsisToBgp", data.RedistributeIsisToBgp.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *LANAutomation) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.discoveredDeviceSiteNameHierarchy"); value.Exists() {
		data.DiscoveredDeviceSiteNameHierarchy = types.StringValue(value.String())
	} else {
		data.DiscoveredDeviceSiteNameHierarchy = types.StringNull()
	}
	if value := res.Get("response.0.primaryDeviceManagmentIPAddress"); value.Exists() {
		data.PrimaryDeviceManagementIpAddress = types.StringValue(value.String())
	} else {
		data.PrimaryDeviceManagementIpAddress = types.StringNull()
	}
	if value := res.Get("response.0.peerDeviceManagmentIPAddress"); value.Exists() {
		data.PeerDeviceManagementIpAddress = types.StringValue(value.String())
	} else {
		data.PeerDeviceManagementIpAddress = types.StringNull()
	}
	if value := res.Get("response.0.primaryDeviceInterfaceNames"); value.Exists() && len(value.Array()) > 0 {
		data.PrimaryDeviceInterfaceNames = helpers.GetStringSet(value.Array())
	} else {
		data.PrimaryDeviceInterfaceNames = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.ipPools"); value.Exists() && len(value.Array()) > 0 {
		data.IpPools = make([]LANAutomationIpPools, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LANAutomationIpPools{}
			if cValue := v.Get("ipPoolName"); cValue.Exists() {
				item.IpPoolName = types.StringValue(cValue.String())
			} else {
				item.IpPoolName = types.StringNull()
			}
			if cValue := v.Get("ipPoolRole"); cValue.Exists() {
				item.IpPoolRole = types.StringValue(cValue.String())
			} else {
				item.IpPoolRole = types.StringNull()
			}
			data.IpPools = append(data.IpPools, item)
			return true
		})
	}
	if value := res.Get("response.0.multicatEnabled"); value.Exists() {
		data.MulticastEnabled = types.BoolValue(value.Bool())
	} else {
		data.MulticastEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.hostNamePrefix"); value.Exists() {
		data.HostNamePrefix = types.StringValue(value.String())
	} else {
		data.HostNamePrefix = types.StringNull()
	}
	if value := res.Get("response.0.hostNameFileId"); value.Exists() {
		data.HostNameFileId = types.StringValue(value.String())
	} else {
		data.HostNameFileId = types.StringNull()
	}
	if value := res.Get("response.0.isisDomainPwd"); value.Exists() {
		data.IsisDomainPassword = types.StringValue(value.String())
	} else {
		data.IsisDomainPassword = types.StringNull()
	}
	if value := res.Get("response.0.redistributeIsisToBgp"); value.Exists() {
		data.RedistributeIsisToBgp = types.BoolValue(value.Bool())
	} else {
		data.RedistributeIsisToBgp = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *LANAutomation) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.discoveredDeviceSiteNameHierarchy"); value.Exists() && !data.DiscoveredDeviceSiteNameHierarchy.IsNull() {
		data.DiscoveredDeviceSiteNameHierarchy = types.StringValue(value.String())
	} else {
		data.DiscoveredDeviceSiteNameHierarchy = types.StringNull()
	}
	if value := res.Get("response.0.primaryDeviceManagmentIPAddress"); value.Exists() && !data.PrimaryDeviceManagementIpAddress.IsNull() {
		data.PrimaryDeviceManagementIpAddress = types.StringValue(value.String())
	} else {
		data.PrimaryDeviceManagementIpAddress = types.StringNull()
	}
	if value := res.Get("response.0.peerDeviceManagmentIPAddress"); value.Exists() && !data.PeerDeviceManagementIpAddress.IsNull() {
		data.PeerDeviceManagementIpAddress = types.StringValue(value.String())
	} else {
		data.PeerDeviceManagementIpAddress = types.StringNull()
	}
	if value := res.Get("response.0.primaryDeviceInterfaceNames"); value.Exists() && !data.PrimaryDeviceInterfaceNames.IsNull() {
		data.PrimaryDeviceInterfaceNames = helpers.GetStringSet(value.Array())
	} else {
		data.PrimaryDeviceInterfaceNames = types.SetNull(types.StringType)
	}
	for i := range data.IpPools {
		keys := [...]string{"ipPoolName", "ipPoolRole"}
		keyValues := [...]string{data.IpPools[i].IpPoolName.ValueString(), data.IpPools[i].IpPoolRole.ValueString()}

		var r gjson.Result
		res.Get("response.0.ipPools").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("ipPoolName"); value.Exists() && !data.IpPools[i].IpPoolName.IsNull() {
			data.IpPools[i].IpPoolName = types.StringValue(value.String())
		} else {
			data.IpPools[i].IpPoolName = types.StringNull()
		}
		if value := r.Get("ipPoolRole"); value.Exists() && !data.IpPools[i].IpPoolRole.IsNull() {
			data.IpPools[i].IpPoolRole = types.StringValue(value.String())
		} else {
			data.IpPools[i].IpPoolRole = types.StringNull()
		}
	}
	if value := res.Get("response.0.multicatEnabled"); value.Exists() && !data.MulticastEnabled.IsNull() {
		data.MulticastEnabled = types.BoolValue(value.Bool())
	} else {
		data.MulticastEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.hostNamePrefix"); value.Exists() && !data.HostNamePrefix.IsNull() {
		data.HostNamePrefix = types.StringValue(value.String())
	} else {
		data.HostNamePrefix = types.StringNull()
	}
	if value := res.Get("response.0.hostNameFileId"); value.Exists() && !data.HostNameFileId.IsNull() {
		data.HostNameFileId = types.StringValue(value.String())
	} else {
		data.HostNameFileId = types.StringNull()
	}
	if value := res.Get("response.0.isisDomainPwd"); value.Exists() && !data.IsisDomainPassword.IsNull() {
		data.IsisDomainPassword = types.StringValue(value.String())
	} else {
		data.IsisDomainPassword = types.StringNull()
	}
	if value := res.Get("response.0.redistributeIsisToBgp"); value.Exists() && !data.RedistributeIsisToBgp.IsNull() {
		data.RedistributeIsisToBgp = types.BoolValue(value.Bool())
	} else {
		data.RedistributeIsisToBgp = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *LANAutomation) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DiscoveredDeviceSiteNameHierarchy.IsNull() {
		return false
	}
	if !data.PrimaryDeviceManagementIpAddress.IsNull() {
		return false
	}
	if !data.PeerDeviceManagementIpAddress.IsNull() {
		return false
	}
	if !data.PrimaryDeviceInterfaceNames.IsNull() {
		return false
	}
	if len(data.IpPools) > 0 {
		return false
	}
	if !data.MulticastEnabled.IsNull() {
		return false
	}
	if !data.HostNamePrefix.IsNull() {
		return false
	}
	if !data.HostNameFileId.IsNull() {
		return false
	}
	if !data.IsisDomainPassword.IsNull() {
		return false
	}
	if !data.RedistributeIsisToBgp.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
