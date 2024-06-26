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
type DeviceDetail struct {
	Id                     types.String `tfsdk:"id"`
	PolicyTagName          types.String `tfsdk:"policy_tag_name"`
	NwDeviceRole           types.String `tfsdk:"nw_device_role"`
	SerialNumber           types.String `tfsdk:"serial_number"`
	NwDeviceName           types.String `tfsdk:"nw_device_name"`
	DeviceGroupHierarchyId types.String `tfsdk:"device_group_hierarchy_id"`
	Cpu                    types.String `tfsdk:"cpu"`
	NwDeviceId             types.String `tfsdk:"nw_device_id"`
	SiteHierarchyGraphId   types.String `tfsdk:"site_hierarchy_graph_id"`
	NwDeviceFamily         types.String `tfsdk:"nw_device_family"`
	MacAddress             types.String `tfsdk:"mac_address"`
	DeviceSeries           types.String `tfsdk:"device_series"`
	CollectionStatus       types.String `tfsdk:"collection_status"`
	MaintenanceMode        types.Bool   `tfsdk:"maintenance_mode"`
	SoftwareVersion        types.String `tfsdk:"software_version"`
	TagIdList              types.Set    `tfsdk:"tag_id_list"`
	OverallHealth          types.Int64  `tfsdk:"overall_health"`
	ManagementIpAddress    types.String `tfsdk:"management_ip_address"`
	Memory                 types.String `tfsdk:"memory"`
	CommunicationState     types.String `tfsdk:"communication_state"`
	NwDeviceType           types.String `tfsdk:"nw_device_type"`
	PlatformId             types.String `tfsdk:"platform_id"`
	Location               types.String `tfsdk:"location"`
	HaStatus               types.String `tfsdk:"ha_status"`
	OsType                 types.String `tfsdk:"os_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data DeviceDetail) getPath() string {
	return "/dna/intent/api/v1/device-detail"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data DeviceDetail) toBody(ctx context.Context, state DeviceDetail) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.PolicyTagName.IsNull() {
		body, _ = sjson.Set(body, "response.policyTagName", data.PolicyTagName.ValueString())
	}
	if !data.NwDeviceRole.IsNull() {
		body, _ = sjson.Set(body, "response.nwDeviceRole", data.NwDeviceRole.ValueString())
	}
	if !data.SerialNumber.IsNull() {
		body, _ = sjson.Set(body, "response.serialNumber", data.SerialNumber.ValueString())
	}
	if !data.NwDeviceName.IsNull() {
		body, _ = sjson.Set(body, "response.nwDeviceName", data.NwDeviceName.ValueString())
	}
	if !data.DeviceGroupHierarchyId.IsNull() {
		body, _ = sjson.Set(body, "response.deviceGroupHierarchyId", data.DeviceGroupHierarchyId.ValueString())
	}
	if !data.Cpu.IsNull() {
		body, _ = sjson.Set(body, "response.cpu", data.Cpu.ValueString())
	}
	if !data.NwDeviceId.IsNull() {
		body, _ = sjson.Set(body, "response.nwDeviceId", data.NwDeviceId.ValueString())
	}
	if !data.SiteHierarchyGraphId.IsNull() {
		body, _ = sjson.Set(body, "response.siteHierarchyGraphId", data.SiteHierarchyGraphId.ValueString())
	}
	if !data.NwDeviceFamily.IsNull() {
		body, _ = sjson.Set(body, "response.nwDeviceFamily", data.NwDeviceFamily.ValueString())
	}
	if !data.MacAddress.IsNull() {
		body, _ = sjson.Set(body, "response.macAddress", data.MacAddress.ValueString())
	}
	if !data.DeviceSeries.IsNull() {
		body, _ = sjson.Set(body, "response.deviceSeries", data.DeviceSeries.ValueString())
	}
	if !data.CollectionStatus.IsNull() {
		body, _ = sjson.Set(body, "response.collectionStatus", data.CollectionStatus.ValueString())
	}
	if !data.MaintenanceMode.IsNull() {
		body, _ = sjson.Set(body, "response.maintenanceMode", data.MaintenanceMode.ValueBool())
	}
	if !data.SoftwareVersion.IsNull() {
		body, _ = sjson.Set(body, "response.softwareVersion", data.SoftwareVersion.ValueString())
	}
	if !data.TagIdList.IsNull() {
		var values []string
		data.TagIdList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "response.tagIdList", values)
	}
	if !data.OverallHealth.IsNull() {
		body, _ = sjson.Set(body, "response.overallHealth", data.OverallHealth.ValueInt64())
	}
	if !data.ManagementIpAddress.IsNull() {
		body, _ = sjson.Set(body, "response.managementIpAddr", data.ManagementIpAddress.ValueString())
	}
	if !data.Memory.IsNull() {
		body, _ = sjson.Set(body, "response.memory", data.Memory.ValueString())
	}
	if !data.CommunicationState.IsNull() {
		body, _ = sjson.Set(body, "response.communicationState", data.CommunicationState.ValueString())
	}
	if !data.NwDeviceType.IsNull() {
		body, _ = sjson.Set(body, "response.nwDeviceType", data.NwDeviceType.ValueString())
	}
	if !data.PlatformId.IsNull() {
		body, _ = sjson.Set(body, "response.platformId", data.PlatformId.ValueString())
	}
	if !data.Location.IsNull() {
		body, _ = sjson.Set(body, "response.location", data.Location.ValueString())
	}
	if !data.HaStatus.IsNull() {
		body, _ = sjson.Set(body, "response.haStatus", data.HaStatus.ValueString())
	}
	if !data.OsType.IsNull() {
		body, _ = sjson.Set(body, "response.osType", data.OsType.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *DeviceDetail) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.policyTagName"); value.Exists() {
		data.PolicyTagName = types.StringValue(value.String())
	} else {
		data.PolicyTagName = types.StringNull()
	}
	if value := res.Get("response.nwDeviceRole"); value.Exists() {
		data.NwDeviceRole = types.StringValue(value.String())
	} else {
		data.NwDeviceRole = types.StringNull()
	}
	if value := res.Get("response.serialNumber"); value.Exists() {
		data.SerialNumber = types.StringValue(value.String())
	} else {
		data.SerialNumber = types.StringNull()
	}
	if value := res.Get("response.nwDeviceName"); value.Exists() {
		data.NwDeviceName = types.StringValue(value.String())
	} else {
		data.NwDeviceName = types.StringNull()
	}
	if value := res.Get("response.deviceGroupHierarchyId"); value.Exists() {
		data.DeviceGroupHierarchyId = types.StringValue(value.String())
	} else {
		data.DeviceGroupHierarchyId = types.StringNull()
	}
	if value := res.Get("response.cpu"); value.Exists() {
		data.Cpu = types.StringValue(value.String())
	} else {
		data.Cpu = types.StringNull()
	}
	if value := res.Get("response.nwDeviceId"); value.Exists() {
		data.NwDeviceId = types.StringValue(value.String())
	} else {
		data.NwDeviceId = types.StringNull()
	}
	if value := res.Get("response.siteHierarchyGraphId"); value.Exists() {
		data.SiteHierarchyGraphId = types.StringValue(value.String())
	} else {
		data.SiteHierarchyGraphId = types.StringNull()
	}
	if value := res.Get("response.nwDeviceFamily"); value.Exists() {
		data.NwDeviceFamily = types.StringValue(value.String())
	} else {
		data.NwDeviceFamily = types.StringNull()
	}
	if value := res.Get("response.macAddress"); value.Exists() {
		data.MacAddress = types.StringValue(value.String())
	} else {
		data.MacAddress = types.StringNull()
	}
	if value := res.Get("response.deviceSeries"); value.Exists() {
		data.DeviceSeries = types.StringValue(value.String())
	} else {
		data.DeviceSeries = types.StringNull()
	}
	if value := res.Get("response.collectionStatus"); value.Exists() {
		data.CollectionStatus = types.StringValue(value.String())
	} else {
		data.CollectionStatus = types.StringNull()
	}
	if value := res.Get("response.maintenanceMode"); value.Exists() {
		data.MaintenanceMode = types.BoolValue(value.Bool())
	} else {
		data.MaintenanceMode = types.BoolNull()
	}
	if value := res.Get("response.softwareVersion"); value.Exists() {
		data.SoftwareVersion = types.StringValue(value.String())
	} else {
		data.SoftwareVersion = types.StringNull()
	}
	if value := res.Get("response.tagIdList"); value.Exists() && len(value.Array()) > 0 {
		data.TagIdList = helpers.GetStringSet(value.Array())
	} else {
		data.TagIdList = types.SetNull(types.StringType)
	}
	if value := res.Get("response.overallHealth"); value.Exists() {
		data.OverallHealth = types.Int64Value(value.Int())
	} else {
		data.OverallHealth = types.Int64Null()
	}
	if value := res.Get("response.managementIpAddr"); value.Exists() {
		data.ManagementIpAddress = types.StringValue(value.String())
	} else {
		data.ManagementIpAddress = types.StringNull()
	}
	if value := res.Get("response.memory"); value.Exists() {
		data.Memory = types.StringValue(value.String())
	} else {
		data.Memory = types.StringNull()
	}
	if value := res.Get("response.communicationState"); value.Exists() {
		data.CommunicationState = types.StringValue(value.String())
	} else {
		data.CommunicationState = types.StringNull()
	}
	if value := res.Get("response.nwDeviceType"); value.Exists() {
		data.NwDeviceType = types.StringValue(value.String())
	} else {
		data.NwDeviceType = types.StringNull()
	}
	if value := res.Get("response.platformId"); value.Exists() {
		data.PlatformId = types.StringValue(value.String())
	} else {
		data.PlatformId = types.StringNull()
	}
	if value := res.Get("response.location"); value.Exists() {
		data.Location = types.StringValue(value.String())
	} else {
		data.Location = types.StringNull()
	}
	if value := res.Get("response.haStatus"); value.Exists() {
		data.HaStatus = types.StringValue(value.String())
	} else {
		data.HaStatus = types.StringNull()
	}
	if value := res.Get("response.osType"); value.Exists() {
		data.OsType = types.StringValue(value.String())
	} else {
		data.OsType = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *DeviceDetail) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.policyTagName"); value.Exists() && !data.PolicyTagName.IsNull() {
		data.PolicyTagName = types.StringValue(value.String())
	} else {
		data.PolicyTagName = types.StringNull()
	}
	if value := res.Get("response.nwDeviceRole"); value.Exists() && !data.NwDeviceRole.IsNull() {
		data.NwDeviceRole = types.StringValue(value.String())
	} else {
		data.NwDeviceRole = types.StringNull()
	}
	if value := res.Get("response.serialNumber"); value.Exists() && !data.SerialNumber.IsNull() {
		data.SerialNumber = types.StringValue(value.String())
	} else {
		data.SerialNumber = types.StringNull()
	}
	if value := res.Get("response.nwDeviceName"); value.Exists() && !data.NwDeviceName.IsNull() {
		data.NwDeviceName = types.StringValue(value.String())
	} else {
		data.NwDeviceName = types.StringNull()
	}
	if value := res.Get("response.deviceGroupHierarchyId"); value.Exists() && !data.DeviceGroupHierarchyId.IsNull() {
		data.DeviceGroupHierarchyId = types.StringValue(value.String())
	} else {
		data.DeviceGroupHierarchyId = types.StringNull()
	}
	if value := res.Get("response.cpu"); value.Exists() && !data.Cpu.IsNull() {
		data.Cpu = types.StringValue(value.String())
	} else {
		data.Cpu = types.StringNull()
	}
	if value := res.Get("response.nwDeviceId"); value.Exists() && !data.NwDeviceId.IsNull() {
		data.NwDeviceId = types.StringValue(value.String())
	} else {
		data.NwDeviceId = types.StringNull()
	}
	if value := res.Get("response.siteHierarchyGraphId"); value.Exists() && !data.SiteHierarchyGraphId.IsNull() {
		data.SiteHierarchyGraphId = types.StringValue(value.String())
	} else {
		data.SiteHierarchyGraphId = types.StringNull()
	}
	if value := res.Get("response.nwDeviceFamily"); value.Exists() && !data.NwDeviceFamily.IsNull() {
		data.NwDeviceFamily = types.StringValue(value.String())
	} else {
		data.NwDeviceFamily = types.StringNull()
	}
	if value := res.Get("response.macAddress"); value.Exists() && !data.MacAddress.IsNull() {
		data.MacAddress = types.StringValue(value.String())
	} else {
		data.MacAddress = types.StringNull()
	}
	if value := res.Get("response.deviceSeries"); value.Exists() && !data.DeviceSeries.IsNull() {
		data.DeviceSeries = types.StringValue(value.String())
	} else {
		data.DeviceSeries = types.StringNull()
	}
	if value := res.Get("response.collectionStatus"); value.Exists() && !data.CollectionStatus.IsNull() {
		data.CollectionStatus = types.StringValue(value.String())
	} else {
		data.CollectionStatus = types.StringNull()
	}
	if value := res.Get("response.maintenanceMode"); value.Exists() && !data.MaintenanceMode.IsNull() {
		data.MaintenanceMode = types.BoolValue(value.Bool())
	} else {
		data.MaintenanceMode = types.BoolNull()
	}
	if value := res.Get("response.softwareVersion"); value.Exists() && !data.SoftwareVersion.IsNull() {
		data.SoftwareVersion = types.StringValue(value.String())
	} else {
		data.SoftwareVersion = types.StringNull()
	}
	if value := res.Get("response.tagIdList"); value.Exists() && !data.TagIdList.IsNull() {
		data.TagIdList = helpers.GetStringSet(value.Array())
	} else {
		data.TagIdList = types.SetNull(types.StringType)
	}
	if value := res.Get("response.overallHealth"); value.Exists() && !data.OverallHealth.IsNull() {
		data.OverallHealth = types.Int64Value(value.Int())
	} else {
		data.OverallHealth = types.Int64Null()
	}
	if value := res.Get("response.managementIpAddr"); value.Exists() && !data.ManagementIpAddress.IsNull() {
		data.ManagementIpAddress = types.StringValue(value.String())
	} else {
		data.ManagementIpAddress = types.StringNull()
	}
	if value := res.Get("response.memory"); value.Exists() && !data.Memory.IsNull() {
		data.Memory = types.StringValue(value.String())
	} else {
		data.Memory = types.StringNull()
	}
	if value := res.Get("response.communicationState"); value.Exists() && !data.CommunicationState.IsNull() {
		data.CommunicationState = types.StringValue(value.String())
	} else {
		data.CommunicationState = types.StringNull()
	}
	if value := res.Get("response.nwDeviceType"); value.Exists() && !data.NwDeviceType.IsNull() {
		data.NwDeviceType = types.StringValue(value.String())
	} else {
		data.NwDeviceType = types.StringNull()
	}
	if value := res.Get("response.platformId"); value.Exists() && !data.PlatformId.IsNull() {
		data.PlatformId = types.StringValue(value.String())
	} else {
		data.PlatformId = types.StringNull()
	}
	if value := res.Get("response.location"); value.Exists() && !data.Location.IsNull() {
		data.Location = types.StringValue(value.String())
	} else {
		data.Location = types.StringNull()
	}
	if value := res.Get("response.haStatus"); value.Exists() && !data.HaStatus.IsNull() {
		data.HaStatus = types.StringValue(value.String())
	} else {
		data.HaStatus = types.StringNull()
	}
	if value := res.Get("response.osType"); value.Exists() && !data.OsType.IsNull() {
		data.OsType = types.StringValue(value.String())
	} else {
		data.OsType = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *DeviceDetail) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.PolicyTagName.IsNull() {
		return false
	}
	if !data.NwDeviceRole.IsNull() {
		return false
	}
	if !data.SerialNumber.IsNull() {
		return false
	}
	if !data.NwDeviceName.IsNull() {
		return false
	}
	if !data.DeviceGroupHierarchyId.IsNull() {
		return false
	}
	if !data.Cpu.IsNull() {
		return false
	}
	if !data.NwDeviceId.IsNull() {
		return false
	}
	if !data.SiteHierarchyGraphId.IsNull() {
		return false
	}
	if !data.NwDeviceFamily.IsNull() {
		return false
	}
	if !data.MacAddress.IsNull() {
		return false
	}
	if !data.DeviceSeries.IsNull() {
		return false
	}
	if !data.CollectionStatus.IsNull() {
		return false
	}
	if !data.MaintenanceMode.IsNull() {
		return false
	}
	if !data.SoftwareVersion.IsNull() {
		return false
	}
	if !data.TagIdList.IsNull() {
		return false
	}
	if !data.OverallHealth.IsNull() {
		return false
	}
	if !data.ManagementIpAddress.IsNull() {
		return false
	}
	if !data.Memory.IsNull() {
		return false
	}
	if !data.CommunicationState.IsNull() {
		return false
	}
	if !data.NwDeviceType.IsNull() {
		return false
	}
	if !data.PlatformId.IsNull() {
		return false
	}
	if !data.Location.IsNull() {
		return false
	}
	if !data.HaStatus.IsNull() {
		return false
	}
	if !data.OsType.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
