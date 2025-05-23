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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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
var _ resource.Resource = &DeployTemplateResource{}

func NewDeployTemplateResource() resource.Resource {
	return &DeployTemplateResource{}
}

type DeployTemplateResource struct {
	client *cc.Client
}

func (r *DeployTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_deploy_template"
}

func (r *DeployTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Deploy Template.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"template_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of template to be provisioned").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"redeploy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flag to indicate whether the template should be redeployed. If set to `true`, template will be redeployed on every Terraform apply").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"force_push_template": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Force Push Template").String,
				Optional:            true,
			},
			"copying_config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Copy config from running into startup").String,
				Optional:            true,
			},
			"is_composite": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Composite template flag").String,
				Optional:            true,
			},
			"main_template_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Composite Template ID").String,
				Optional:            true,
			},
			"member_template_deployment_info": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Member Template Deployment Info").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"template_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Versioned Template ID").String,
							Required:            true,
						},
						"force_push_template": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Force Push Template").String,
							Optional:            true,
						},
						"is_composite": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Composite template flag").String,
							Optional:            true,
						},
						"copying_config": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Copy config from running into startup").String,
							Optional:            true,
						},
						"main_template_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Template ID").String,
							Optional:            true,
						},
						"target_info": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Target info to deploy template").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"host_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME").String,
										Optional:            true,
									},
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ID of device is required if targetType is MANAGED_DEVICE_UUID").String,
										Optional:            true,
									},
									"params": schema.MapAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Template params/values to be provisioned").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"resource_params": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Resource params to be provisioned").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Target type of device").AddStringEnumDescription("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME"),
													},
												},
												"scope": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Scope").String,
													Optional:            true,
												},
												"value": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Value").String,
													Optional:            true,
												},
											},
										},
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Target type of device").AddStringEnumDescription("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME"),
										},
									},
									"versioned_template_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Versioned template ID to be provisioned").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"target_info": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Target info to deploy template").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME").String,
							Optional:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of device is required if `type` is MANAGED_DEVICE_UUID").String,
							Optional:            true,
						},
						"params": schema.MapAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Template params/values to be provisioned").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"resource_params": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Resource params to be provisioned").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Target type of device").AddStringEnumDescription("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME"),
										},
									},
									"scope": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Scope").String,
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Value").String,
										Optional:            true,
									},
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Target type of device").AddStringEnumDescription("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME"),
							},
						},
						"versioned_template_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Versioned template ID to be provisioned").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DeployTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *DeployTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeployTemplate

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, DeployTemplate{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.TemplateId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

func (r *DeployTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeployTemplate

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	state.Redeploy = types.BoolValue(false)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *DeployTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeployTemplate

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *DeployTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeployTemplate

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
