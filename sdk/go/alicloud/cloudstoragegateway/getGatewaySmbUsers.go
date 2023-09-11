// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudstoragegateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides the Cloud Storage Gateway Gateway SMB Users of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.142.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudstoragegateway"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("default-NODELETING"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
//				VpcId: pulumi.StringRef(defaultNetworks.Ids[0]),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			example, err := cloudstoragegateway.NewStorageBundle(ctx, "example", &cloudstoragegateway.StorageBundleArgs{
//				StorageBundleName: pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGateway, err := cloudstoragegateway.NewGateway(ctx, "defaultGateway", &cloudstoragegateway.GatewayArgs{
//				Description:            pulumi.String("tf-acctestDesalone"),
//				GatewayClass:           pulumi.String("Standard"),
//				Type:                   pulumi.String("File"),
//				PaymentType:            pulumi.String("PayAsYouGo"),
//				VswitchId:              *pulumi.String(defaultSwitches.Ids[0]),
//				ReleaseAfterExpiration: pulumi.Bool(false),
//				PublicNetworkBandwidth: pulumi.Int(40),
//				StorageBundleId:        example.ID(),
//				Location:               pulumi.String("Cloud"),
//				GatewayName:            pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGatewaySmbUser, err := cloudstoragegateway.NewGatewaySmbUser(ctx, "defaultGatewaySmbUser", &cloudstoragegateway.GatewaySmbUserArgs{
//				Username:  pulumi.String("your_username"),
//				Password:  pulumi.String("password"),
//				GatewayId: defaultGateway.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ids := cloudstoragegateway.GetGatewaySmbUsersOutput(ctx, cloudstoragegateway.GetGatewaySmbUsersOutputArgs{
//				GatewayId: defaultGateway.ID(),
//				Ids: pulumi.StringArray{
//					defaultGatewaySmbUser.ID(),
//				},
//			}, nil)
//			ctx.Export("cloudStorageGatewayGatewaySmbUserId1", ids.ApplyT(func(ids cloudstoragegateway.GetGatewaySmbUsersResult) (*string, error) {
//				return &ids.Users[0].Id, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
func GetGatewaySmbUsers(ctx *pulumi.Context, args *GetGatewaySmbUsersArgs, opts ...pulumi.InvokeOption) (*GetGatewaySmbUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGatewaySmbUsersResult
	err := ctx.Invoke("alicloud:cloudstoragegateway/getGatewaySmbUsers:getGatewaySmbUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewaySmbUsers.
type GetGatewaySmbUsersArgs struct {
	// The Gateway ID.
	GatewayId string `pulumi:"gatewayId"`
	// A list of Gateway SMB User IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Gateway SMB username.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getGatewaySmbUsers.
type GetGatewaySmbUsersResult struct {
	GatewayId string `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                   `pulumi:"id"`
	Ids        []string                 `pulumi:"ids"`
	NameRegex  *string                  `pulumi:"nameRegex"`
	OutputFile *string                  `pulumi:"outputFile"`
	Users      []GetGatewaySmbUsersUser `pulumi:"users"`
}

func GetGatewaySmbUsersOutput(ctx *pulumi.Context, args GetGatewaySmbUsersOutputArgs, opts ...pulumi.InvokeOption) GetGatewaySmbUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGatewaySmbUsersResult, error) {
			args := v.(GetGatewaySmbUsersArgs)
			r, err := GetGatewaySmbUsers(ctx, &args, opts...)
			var s GetGatewaySmbUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGatewaySmbUsersResultOutput)
}

// A collection of arguments for invoking getGatewaySmbUsers.
type GetGatewaySmbUsersOutputArgs struct {
	// The Gateway ID.
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
	// A list of Gateway SMB User IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Gateway SMB username.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetGatewaySmbUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewaySmbUsersArgs)(nil)).Elem()
}

// A collection of values returned by getGatewaySmbUsers.
type GetGatewaySmbUsersResultOutput struct{ *pulumi.OutputState }

func (GetGatewaySmbUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewaySmbUsersResult)(nil)).Elem()
}

func (o GetGatewaySmbUsersResultOutput) ToGetGatewaySmbUsersResultOutput() GetGatewaySmbUsersResultOutput {
	return o
}

func (o GetGatewaySmbUsersResultOutput) ToGetGatewaySmbUsersResultOutputWithContext(ctx context.Context) GetGatewaySmbUsersResultOutput {
	return o
}

func (o GetGatewaySmbUsersResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetGatewaySmbUsersResult] {
	return pulumix.Output[GetGatewaySmbUsersResult]{
		OutputState: o.OutputState,
	}
}

func (o GetGatewaySmbUsersResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaySmbUsersResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGatewaySmbUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaySmbUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGatewaySmbUsersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGatewaySmbUsersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetGatewaySmbUsersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaySmbUsersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetGatewaySmbUsersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaySmbUsersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetGatewaySmbUsersResultOutput) Users() GetGatewaySmbUsersUserArrayOutput {
	return o.ApplyT(func(v GetGatewaySmbUsersResult) []GetGatewaySmbUsersUser { return v.Users }).(GetGatewaySmbUsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGatewaySmbUsersResultOutput{})
}
