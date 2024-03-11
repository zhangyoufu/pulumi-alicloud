// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Private Link Vpc Endpoint Service User resource. Endpoint service user whitelist.
//
// For information about Private Link Vpc Endpoint Service User and how to use it, see [What is Vpc Endpoint Service User](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-addusertovpcendpointservice).
//
// > **NOTE:** Available since v1.110.0.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/privatelink"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tfexampleuser"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			exampleVpcEndpointService, err := privatelink.NewVpcEndpointService(ctx, "exampleVpcEndpointService", &privatelink.VpcEndpointServiceArgs{
//				ServiceDescription:   pulumi.String(name),
//				ConnectBandwidth:     pulumi.Int(103),
//				AutoAcceptConnection: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			exampleUser, err := ram.NewUser(ctx, "exampleUser", &ram.UserArgs{
//				DisplayName: pulumi.String("user_display_name"),
//				Mobile:      pulumi.String("86-18688888888"),
//				Email:       pulumi.String("hello.uuu@aaa.com"),
//				Comments:    pulumi.String("yoyoyo"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = privatelink.NewVpcEndpointServiceUser(ctx, "exampleVpcEndpointServiceUser", &privatelink.VpcEndpointServiceUserArgs{
//				ServiceId: exampleVpcEndpointService.ID(),
//				UserId:    exampleUser.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Private Link Vpc Endpoint Service User can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:privatelink/vpcEndpointServiceUser:VpcEndpointServiceUser example <service_id>:<user_id>
// ```
type VpcEndpointServiceUser struct {
	pulumi.CustomResourceState

	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The endpoint service ID.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The ID of the Alibaba Cloud account in the whitelist of the endpoint service.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewVpcEndpointServiceUser registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointServiceUser(ctx *pulumi.Context,
	name string, args *VpcEndpointServiceUserArgs, opts ...pulumi.ResourceOption) (*VpcEndpointServiceUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcEndpointServiceUser
	err := ctx.RegisterResource("alicloud:privatelink/vpcEndpointServiceUser:VpcEndpointServiceUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointServiceUser gets an existing VpcEndpointServiceUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointServiceUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointServiceUserState, opts ...pulumi.ResourceOption) (*VpcEndpointServiceUser, error) {
	var resource VpcEndpointServiceUser
	err := ctx.ReadResource("alicloud:privatelink/vpcEndpointServiceUser:VpcEndpointServiceUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointServiceUser resources.
type vpcEndpointServiceUserState struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `pulumi:"dryRun"`
	// The endpoint service ID.
	ServiceId *string `pulumi:"serviceId"`
	// The ID of the Alibaba Cloud account in the whitelist of the endpoint service.
	UserId *string `pulumi:"userId"`
}

type VpcEndpointServiceUserState struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun pulumi.BoolPtrInput
	// The endpoint service ID.
	ServiceId pulumi.StringPtrInput
	// The ID of the Alibaba Cloud account in the whitelist of the endpoint service.
	UserId pulumi.StringPtrInput
}

func (VpcEndpointServiceUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceUserState)(nil)).Elem()
}

type vpcEndpointServiceUserArgs struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `pulumi:"dryRun"`
	// The endpoint service ID.
	ServiceId string `pulumi:"serviceId"`
	// The ID of the Alibaba Cloud account in the whitelist of the endpoint service.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a VpcEndpointServiceUser resource.
type VpcEndpointServiceUserArgs struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun pulumi.BoolPtrInput
	// The endpoint service ID.
	ServiceId pulumi.StringInput
	// The ID of the Alibaba Cloud account in the whitelist of the endpoint service.
	UserId pulumi.StringInput
}

func (VpcEndpointServiceUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceUserArgs)(nil)).Elem()
}

type VpcEndpointServiceUserInput interface {
	pulumi.Input

	ToVpcEndpointServiceUserOutput() VpcEndpointServiceUserOutput
	ToVpcEndpointServiceUserOutputWithContext(ctx context.Context) VpcEndpointServiceUserOutput
}

func (*VpcEndpointServiceUser) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServiceUser)(nil)).Elem()
}

func (i *VpcEndpointServiceUser) ToVpcEndpointServiceUserOutput() VpcEndpointServiceUserOutput {
	return i.ToVpcEndpointServiceUserOutputWithContext(context.Background())
}

func (i *VpcEndpointServiceUser) ToVpcEndpointServiceUserOutputWithContext(ctx context.Context) VpcEndpointServiceUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceUserOutput)
}

// VpcEndpointServiceUserArrayInput is an input type that accepts VpcEndpointServiceUserArray and VpcEndpointServiceUserArrayOutput values.
// You can construct a concrete instance of `VpcEndpointServiceUserArrayInput` via:
//
//	VpcEndpointServiceUserArray{ VpcEndpointServiceUserArgs{...} }
type VpcEndpointServiceUserArrayInput interface {
	pulumi.Input

	ToVpcEndpointServiceUserArrayOutput() VpcEndpointServiceUserArrayOutput
	ToVpcEndpointServiceUserArrayOutputWithContext(context.Context) VpcEndpointServiceUserArrayOutput
}

type VpcEndpointServiceUserArray []VpcEndpointServiceUserInput

func (VpcEndpointServiceUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointServiceUser)(nil)).Elem()
}

func (i VpcEndpointServiceUserArray) ToVpcEndpointServiceUserArrayOutput() VpcEndpointServiceUserArrayOutput {
	return i.ToVpcEndpointServiceUserArrayOutputWithContext(context.Background())
}

func (i VpcEndpointServiceUserArray) ToVpcEndpointServiceUserArrayOutputWithContext(ctx context.Context) VpcEndpointServiceUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceUserArrayOutput)
}

// VpcEndpointServiceUserMapInput is an input type that accepts VpcEndpointServiceUserMap and VpcEndpointServiceUserMapOutput values.
// You can construct a concrete instance of `VpcEndpointServiceUserMapInput` via:
//
//	VpcEndpointServiceUserMap{ "key": VpcEndpointServiceUserArgs{...} }
type VpcEndpointServiceUserMapInput interface {
	pulumi.Input

	ToVpcEndpointServiceUserMapOutput() VpcEndpointServiceUserMapOutput
	ToVpcEndpointServiceUserMapOutputWithContext(context.Context) VpcEndpointServiceUserMapOutput
}

type VpcEndpointServiceUserMap map[string]VpcEndpointServiceUserInput

func (VpcEndpointServiceUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointServiceUser)(nil)).Elem()
}

func (i VpcEndpointServiceUserMap) ToVpcEndpointServiceUserMapOutput() VpcEndpointServiceUserMapOutput {
	return i.ToVpcEndpointServiceUserMapOutputWithContext(context.Background())
}

func (i VpcEndpointServiceUserMap) ToVpcEndpointServiceUserMapOutputWithContext(ctx context.Context) VpcEndpointServiceUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceUserMapOutput)
}

type VpcEndpointServiceUserOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServiceUser)(nil)).Elem()
}

func (o VpcEndpointServiceUserOutput) ToVpcEndpointServiceUserOutput() VpcEndpointServiceUserOutput {
	return o
}

func (o VpcEndpointServiceUserOutput) ToVpcEndpointServiceUserOutputWithContext(ctx context.Context) VpcEndpointServiceUserOutput {
	return o
}

// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
func (o VpcEndpointServiceUserOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcEndpointServiceUser) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The endpoint service ID.
func (o VpcEndpointServiceUserOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceUser) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The ID of the Alibaba Cloud account in the whitelist of the endpoint service.
func (o VpcEndpointServiceUserOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceUser) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type VpcEndpointServiceUserArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointServiceUser)(nil)).Elem()
}

func (o VpcEndpointServiceUserArrayOutput) ToVpcEndpointServiceUserArrayOutput() VpcEndpointServiceUserArrayOutput {
	return o
}

func (o VpcEndpointServiceUserArrayOutput) ToVpcEndpointServiceUserArrayOutputWithContext(ctx context.Context) VpcEndpointServiceUserArrayOutput {
	return o
}

func (o VpcEndpointServiceUserArrayOutput) Index(i pulumi.IntInput) VpcEndpointServiceUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointServiceUser {
		return vs[0].([]*VpcEndpointServiceUser)[vs[1].(int)]
	}).(VpcEndpointServiceUserOutput)
}

type VpcEndpointServiceUserMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointServiceUser)(nil)).Elem()
}

func (o VpcEndpointServiceUserMapOutput) ToVpcEndpointServiceUserMapOutput() VpcEndpointServiceUserMapOutput {
	return o
}

func (o VpcEndpointServiceUserMapOutput) ToVpcEndpointServiceUserMapOutputWithContext(ctx context.Context) VpcEndpointServiceUserMapOutput {
	return o
}

func (o VpcEndpointServiceUserMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointServiceUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointServiceUser {
		return vs[0].(map[string]*VpcEndpointServiceUser)[vs[1].(string)]
	}).(VpcEndpointServiceUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceUserInput)(nil)).Elem(), &VpcEndpointServiceUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceUserArrayInput)(nil)).Elem(), VpcEndpointServiceUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceUserMapInput)(nil)).Elem(), VpcEndpointServiceUserMap{})
	pulumi.RegisterOutputType(VpcEndpointServiceUserOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceUserArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceUserMapOutput{})
}
