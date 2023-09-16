// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"MyTest/internal/sdk"
	"MyTest/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ZoneResource{}
var _ resource.ResourceWithImportState = &ZoneResource{}

func NewZoneResource() resource.Resource {
	return &ZoneResource{}
}

// ZoneResource defines the resource implementation.
type ZoneResource struct {
	client *sdk.MyTest
}

// ZoneResourceModel describes the resource data model.
type ZoneResourceModel struct {
	AccountID     types.Int64        `tfsdk:"account_id"`
	Code          types.String       `tfsdk:"code"`
	Config        *ZoneVcenterConfig `tfsdk:"config"`
	Credential    *ZoneCredential    `tfsdk:"credential"`
	Description   types.String       `tfsdk:"description"`
	Enabled       types.Bool         `tfsdk:"enabled"`
	GroupID       types.Int64        `tfsdk:"group_id"`
	Groups        []ZoneGroups       `tfsdk:"groups"`
	ID            types.Int64        `tfsdk:"id"`
	Name          types.String       `tfsdk:"name"`
	ScalePriority types.Int64        `tfsdk:"scale_priority"`
	Visibility    types.String       `tfsdk:"visibility"`
	ZoneType      *ZoneZoneType      `tfsdk:"zone_type"`
}

func (r *ZoneResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone"
}

func (r *ZoneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Zone Resource",

		Attributes: map[string]schema.Attribute{
			"account_id": schema.Int64Attribute{
				Computed: true,
				Optional: true,
			},
			"code": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"api_url": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"appliance_url": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"datacenter": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			"credential": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `Map containing Credential ID. Setting ` + "`" + `type` + "`" + ` to ` + "`" + `local` + "`" + ` means use the values set in the local cloud config instead of associating a credential.`,
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: `Optional description field if you want to put more info there`,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"group_id": schema.Int64Attribute{
				Required:    true,
				Description: `Specifies which Server group this cloud should be assigned to`,
			},
			"groups": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"account_id": schema.Int64Attribute{
							Computed: true,
						},
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"scale_priority": schema.Int64Attribute{
				Computed: true,
				Optional: true,
			},
			"visibility": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"zone_type": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"code": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `Map containing the Cloud (zone) code name. See the zone-types API to fetch a list of all available Cloud (zone) types and their codes.`,
			},
		},
	}
}

func (r *ZoneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.MyTest)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.MyTest, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ZoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZoneResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	zone := *data.ToCreateSDKType()
	request := operations.AddCloudsRequestBody{
		Zone: zone,
	}
	res, err := r.client.Clouds.AddClouds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.AddClouds200ApplicationJSONObject == nil || res.AddClouds200ApplicationJSONObject.Zone == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.AddClouds200ApplicationJSONObject.Zone)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZoneResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueInt64()
	request := operations.GetCloudsRequest{
		ID: id,
	}
	res, err := r.client.Clouds.GetClouds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.GetClouds200ApplicationJSONObject == nil || res.GetClouds200ApplicationJSONObject.Zone == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.GetClouds200ApplicationJSONObject.Zone)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZoneResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var requestBody *operations.UpdateCloudsRequestBody
	zone := *data.ToUpdateSDKType()
	requestBody = &operations.UpdateCloudsRequestBody{
		Zone: zone,
	}
	id := data.ID.ValueInt64()
	request := operations.UpdateCloudsRequest{
		RequestBody: requestBody,
		ID:          id,
	}
	res, err := r.client.Clouds.UpdateClouds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.UpdateClouds200ApplicationJSONObject == nil || res.UpdateClouds200ApplicationJSONObject.Zone == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.UpdateClouds200ApplicationJSONObject.Zone)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZoneResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueInt64()
	request := operations.RemoveCloudsRequest{
		ID: id,
	}
	res, err := r.client.Clouds.RemoveClouds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *ZoneResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
