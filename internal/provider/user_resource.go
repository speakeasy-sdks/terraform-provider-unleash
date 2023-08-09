// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"terraform/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"terraform/internal/sdk/pkg/models/operations"
	"terraform/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &UserResource{}
var _ resource.ResourceWithImportState = &UserResource{}

func NewUserResource() resource.Resource {
	return &UserResource{}
}

// UserResource defines the resource implementation.
type UserResource struct {
	client *sdk.SDK
}

// UserResourceModel describes the resource data model.
type UserResourceModel struct {
	AccountType   types.String             `tfsdk:"account_type"`
	CreatedAt     types.String             `tfsdk:"created_at"`
	Email         types.String             `tfsdk:"email"`
	EmailSent     types.Bool               `tfsdk:"email_sent"`
	ID            types.Int64              `tfsdk:"id"`
	ImageURL      types.String             `tfsdk:"image_url"`
	InviteLink    types.String             `tfsdk:"invite_link"`
	IsAPI         types.Bool               `tfsdk:"is_api"`
	LoginAttempts types.Int64              `tfsdk:"login_attempts"`
	Name          types.String             `tfsdk:"name"`
	Password      types.String             `tfsdk:"password"`
	Permissions   []types.String           `tfsdk:"permissions"`
	RootRole      CreateUserSchemaRootRole `tfsdk:"root_role"`
	SeenAt        types.String             `tfsdk:"seen_at"`
	SendEmail     types.Bool               `tfsdk:"send_email"`
	Username      types.String             `tfsdk:"username"`
}

func (r *UserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *UserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "User Resource",

		Attributes: map[string]schema.Attribute{
			"account_type": schema.StringAttribute{
				Computed:    true,
				Description: `A user is either an actual User or a Service Account`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The user was created at this time`,
			},
			"email": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Email of the user`,
			},
			"email_sent": schema.BoolAttribute{
				Computed:    true,
				Description: `Is the welcome email sent to the user or not`,
			},
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: `The user id`,
			},
			"image_url": schema.StringAttribute{
				Computed:    true,
				Description: `URL used for the userprofile image`,
			},
			"invite_link": schema.StringAttribute{
				Computed:    true,
				Description: `If the user is actively inviting other users, this is the link that can be shared with other users`,
			},
			"is_api": schema.BoolAttribute{
				Computed:    true,
				Description: `(Deprecated): Used internally to know which operations the user should be allowed to perform`,
			},
			"login_attempts": schema.Int64Attribute{
				Computed:    true,
				Description: `How many unsuccessful attempts at logging in has the user made`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Name of the user`,
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Description: `Password for the user`,
			},
			"permissions": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `Deprecated`,
			},
			"root_role": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"integer": schema.Int64Attribute{
						Computed: true,
						Optional: true,
					},
					"create_user_response_schema_root_role_2": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Admin",
								"Editor",
								"Viewer",
								"Owner",
								"Member",
							),
						},
						MarkdownDescription: `must be one of ["Admin", "Editor", "Viewer", "Owner", "Member"]` + "\n" +
							`Which [root role](https://docs.getunleash.io/reference/rbac#standard-roles) this user is assigned. Usually a numeric role ID, but can be a string when returning newly created user with an explicit string role.`,
					},
					"create_user_schema_root_role_2": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Admin",
								"Editor",
								"Viewer",
								"Owner",
								"Member",
							),
						},
						MarkdownDescription: `must be one of ["Admin", "Editor", "Viewer", "Owner", "Member"]` + "\n" +
							`The role to assign to the user. Can be either the role's ID or its unique name.`,
					},
				},
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
				Description: `Which [root role](https://docs.getunleash.io/reference/rbac#standard-roles) this user is assigned. Usually a numeric role ID, but can be a string when returning newly created user with an explicit string role.`,
			},
			"seen_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The last time this user logged in`,
			},
			"send_email": schema.BoolAttribute{
				Optional:    true,
				Description: `Whether to send a welcome email with a login link to the user or not. Defaults to ` + "`" + `true` + "`" + `.`,
			},
			"username": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `A unique username for the user`,
			},
		},
	}
}

func (r *UserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *UserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *UserResourceModel
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

	request := *data.ToCreateSDKType()
	res, err := r.client.Users.CreateUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.CreateUserResponseSchema == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.CreateUserResponseSchema)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *UserResourceModel
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

	id := data.ID.ValueString()
	request := operations.GetUserRequest{
		ID: id,
	}
	res, err := r.client.Users.GetUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	if res.UserSchema == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.UserSchema)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *UserResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()
	updateUserSchema := *data.ToUpdateSDKType()
	request := operations.UpdateUserRequest{
		ID:               id,
		UpdateUserSchema: updateUserSchema,
	}
	res, err := r.client.Users.UpdateUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	if res.CreateUserResponseSchema == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.CreateUserResponseSchema)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *UserResourceModel
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

	id := data.ID.ValueString()
	request := operations.DeleteUserRequest{
		ID: id,
	}
	res, err := r.client.Users.DeleteUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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

func (r *UserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
