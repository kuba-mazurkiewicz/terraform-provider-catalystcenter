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
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessEnterpriseSSIDDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessEnterpriseSSIDDataSource{}
)

func NewWirelessEnterpriseSSIDDataSource() datasource.DataSource {
	return &WirelessEnterpriseSSIDDataSource{}
}

type WirelessEnterpriseSSIDDataSource struct {
	client *cc.Client
}

func (d *WirelessEnterpriseSSIDDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_enterprise_ssid"
}

func (d *WirelessEnterpriseSSIDDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Wireless Enterprise SSID.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "SSID Name",
				Computed:            true,
			},
			"security_level": schema.StringAttribute{
				MarkdownDescription: "Security Level",
				Computed:            true,
			},
			"passphrase": schema.StringAttribute{
				MarkdownDescription: "Passphrase",
				Computed:            true,
			},
			"enable_fast_lane": schema.BoolAttribute{
				MarkdownDescription: "Enable FastLane",
				Computed:            true,
			},
			"enable_mac_filtering": schema.BoolAttribute{
				MarkdownDescription: "Enable MAC Filtering",
				Computed:            true,
			},
			"traffic_type": schema.StringAttribute{
				MarkdownDescription: "Traffic Type",
				Computed:            true,
			},
			"radio_policy": schema.StringAttribute{
				MarkdownDescription: "Radio Policy",
				Computed:            true,
			},
			"enable_broadcast_ssid": schema.BoolAttribute{
				MarkdownDescription: "Enable Broadcast SSID",
				Computed:            true,
			},
			"fast_transition": schema.StringAttribute{
				MarkdownDescription: "Fast Transition",
				Computed:            true,
			},
			"enable_session_time_out": schema.BoolAttribute{
				MarkdownDescription: "Enable Session Timeout",
				Computed:            true,
			},
			"session_time_out": schema.Int64Attribute{
				MarkdownDescription: "Session Time Out",
				Computed:            true,
			},
			"enable_client_exclusion": schema.BoolAttribute{
				MarkdownDescription: "Enable Client Exclusion",
				Computed:            true,
			},
			"client_exclusion_timeout": schema.Int64Attribute{
				MarkdownDescription: "Client Exclusion Timeout",
				Computed:            true,
			},
			"enable_basic_service_set_max_idle": schema.BoolAttribute{
				MarkdownDescription: "Enable Basic Service Set Max Idle",
				Computed:            true,
			},
			"basic_service_set_client_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "Basic Service Set Client Idle Timeout",
				Computed:            true,
			},
			"enable_directed_multicast_service": schema.BoolAttribute{
				MarkdownDescription: "Enable Directed Multicast Service",
				Computed:            true,
			},
			"enable_neighbor_list": schema.BoolAttribute{
				MarkdownDescription: "Enable Neighbor List",
				Computed:            true,
			},
			"mfp_client_protection": schema.StringAttribute{
				MarkdownDescription: "Mfp Client Protection",
				Computed:            true,
			},
			"nas_options": schema.SetAttribute{
				MarkdownDescription: "Nas Options",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"profile_name": schema.StringAttribute{
				MarkdownDescription: "Profile Name",
				Computed:            true,
			},
			"policy_profile_name": schema.StringAttribute{
				MarkdownDescription: "Policy Profile Name",
				Computed:            true,
			},
			"aaa_override": schema.BoolAttribute{
				MarkdownDescription: "AAA Override",
				Computed:            true,
			},
			"coverage_hole_detection_enable": schema.BoolAttribute{
				MarkdownDescription: "Coverage Hole Detection Enable",
				Computed:            true,
			},
			"protected_management_frame": schema.StringAttribute{
				MarkdownDescription: "Protected Management Frame",
				Computed:            true,
			},
			"multi_psk_settings": schema.ListNestedAttribute{
				MarkdownDescription: "Multi PSK Settings (Only applicable for SSID with PERSONAL auth type and PSK)",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"priority": schema.StringAttribute{
							MarkdownDescription: "Priority",
							Computed:            true,
						},
						"passphrase_type": schema.StringAttribute{
							MarkdownDescription: "Passphrase Type",
							Computed:            true,
						},
						"passphrase": schema.StringAttribute{
							MarkdownDescription: "Passphrase",
							Computed:            true,
						},
					},
				},
			},
			"client_rate_limit": schema.Int64Attribute{
				MarkdownDescription: "Client Rate Limit (in bits per second)",
				Computed:            true,
			},
			"auth_key_mgmt": schema.SetAttribute{
				MarkdownDescription: "Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"rsn_cipher_suite_gcmp256": schema.BoolAttribute{
				MarkdownDescription: "Rsn Cipher Suite Gcmp256",
				Computed:            true,
			},
			"rsn_cipher_suite_ccmp256": schema.BoolAttribute{
				MarkdownDescription: "Rsn Cipher Suite Ccmp256",
				Computed:            true,
			},
			"rsn_cipher_suite_gcmp128": schema.BoolAttribute{
				MarkdownDescription: "Rsn Cipher Suite Gcmp 128",
				Computed:            true,
			},
			"ghz6_policy_client_steering": schema.BoolAttribute{
				MarkdownDescription: "Ghz6 Policy Client Steering",
				Computed:            true,
			},
			"ghz24_policy": schema.StringAttribute{
				MarkdownDescription: "Ghz24 Policy",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessEnterpriseSSIDDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *WirelessEnterpriseSSIDDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessEnterpriseSSID

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?ssidName=" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
