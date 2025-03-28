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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PnPDeviceClaimSiteResource{}

func NewPnPDeviceClaimSiteResource() resource.Resource {
	return &PnPDeviceClaimSiteResource{}
}

type PnPDeviceClaimSiteResource struct {
	client *cc.Client
}

func (r *PnPDeviceClaimSiteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pnp_device_claim_site"
}

func (r *PnPDeviceClaimSiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a PnP Device Claim Site.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The device ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The site ID").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The device type").AddStringEnumDescription("Default", "StackSwitch", "AccessPoint", "Sensor", "CatalystWLC", "MobilityExpress").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Default", "StackSwitch", "AccessPoint", "Sensor", "CatalystWLC", "MobilityExpress"),
				},
			},
			"image_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Image ID. Required if `type` is `Default` or `StackSwitch`.").String,
				Optional:            true,
			},
			"image_skip": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Skip image provisioning.").String,
				Optional:            true,
			},
			"config_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Config (temaplate) ID. Required if `type` is `Default` or `StackSwitch`.").String,
				Optional:            true,
			},
			"config_parameters": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of config (temaplate) parameters.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of config parameter.").String,
							Required:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Value of config parameter.").String,
							Optional:            true,
						},
					},
				},
			},
			"rf_profile": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("RF profile. Required if `type` is `AccessPoint`.").String,
				Optional:            true,
			},
			"static_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static IP address. Required if `type` is `CatalystWLC` or `MobilityExpress`.").String,
				Optional:            true,
			},
			"subnet_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Subnet mask. Required if `type` is `CatalystWLC` or `MobilityExpress`.").String,
				Optional:            true,
			},
			"gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Gateway IP. Required if `type` is `CatalystWLC` or `MobilityExpress`.").String,
				Optional:            true,
			},
			"vlan_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Vlan ID. Required for Catalyst 9800 WLC.").String,
				Optional:            true,
			},
			"ip_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP interface name. Required for Catalyst 9800 WLC.").String,
				Optional:            true,
			},
			"sensor_profile": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Sensor profile. Required if `type` is `Sensor`.").String,
				Optional:            true,
			},
		},
	}
}

func (r *PnPDeviceClaimSiteResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *PnPDeviceClaimSiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PnPDeviceClaimSite

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, PnPDeviceClaimSite{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.DeviceId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *PnPDeviceClaimSiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PnPDeviceClaimSite

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *PnPDeviceClaimSiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PnPDeviceClaimSite

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx, state)
	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *PnPDeviceClaimSiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PnPDeviceClaimSite

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
