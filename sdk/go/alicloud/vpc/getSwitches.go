// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of VSwitches owned by an Alibaba Cloud account.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "vswitchDatasourceName"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := alicloud.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := vpc.NewNetwork(ctx, "vpc", &vpc.NetworkArgs{
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//				VpcName:   pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			vswitch, err := vpc.NewSwitch(ctx, "vswitch", &vpc.SwitchArgs{
//				AvailabilityZone: *pulumi.String(defaultZones.Zones[0].Id),
//				CidrBlock:        pulumi.String("172.16.0.0/24"),
//				VpcId:            vpc.ID(),
//				VswitchName:      pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_ = vpc.GetSwitchesOutput(ctx, vpc.GetSwitchesOutputArgs{
//				NameRegex: vswitch.VswitchName,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetSwitches(ctx *pulumi.Context, args *GetSwitchesArgs, opts ...pulumi.InvokeOption) (*GetSwitchesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSwitchesResult
	err := ctx.Invoke("alicloud:vpc/getSwitches:getSwitches", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSwitches.
type GetSwitchesArgs struct {
	// Filter results by a specific CIDR block. For example: "172.16.0.0/12".
	CidrBlock *string `pulumi:"cidrBlock"`
	// Specifies whether to precheck this request only. Valid values: `true` and `false`.
	DryRun *bool `pulumi:"dryRun"`
	// A list of VSwitch IDs.
	Ids []string `pulumi:"ids"`
	// Indicate whether the VSwitch is created by the system.
	IsDefault *bool `pulumi:"isDefault"`
	// A regex string to filter results by name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The Id of resource group which VSWitch belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The route table ID of the VSwitch.
	RouteTableId *string `pulumi:"routeTableId"`
	// The status of the VSwitch. Valid values: `Available` and `Pending`.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the VPC that owns the VSwitch.
	VpcId *string `pulumi:"vpcId"`
	// The name of the VSwitch.
	VswitchName *string `pulumi:"vswitchName"`
	// The VSwitch owner id.
	VswitchOwnerId *int `pulumi:"vswitchOwnerId"`
	// The availability zone of the VSwitch.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getSwitches.
type GetSwitchesResult struct {
	// CIDR block of the VSwitch.
	CidrBlock *string `pulumi:"cidrBlock"`
	DryRun    *bool   `pulumi:"dryRun"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of VSwitch IDs.
	Ids []string `pulumi:"ids"`
	// Whether the VSwitch is the default one in the region.
	IsDefault *bool   `pulumi:"isDefault"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of VSwitch names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The resource group ID of the VSwitch.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The route table ID of the VSwitch.
	RouteTableId *string `pulumi:"routeTableId"`
	// The status of the VSwitch.
	Status *string `pulumi:"status"`
	// The Tags of the VSwitch.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the VPC that owns the VSwitch.
	VpcId *string `pulumi:"vpcId"`
	// Name of the VSwitch.
	VswitchName    *string `pulumi:"vswitchName"`
	VswitchOwnerId *int    `pulumi:"vswitchOwnerId"`
	// A list of VSwitches. Each element contains the following attributes:
	Vswitches []GetSwitchesVswitch `pulumi:"vswitches"`
	// ID of the availability zone where the VSwitch is located.
	ZoneId *string `pulumi:"zoneId"`
}

func GetSwitchesOutput(ctx *pulumi.Context, args GetSwitchesOutputArgs, opts ...pulumi.InvokeOption) GetSwitchesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSwitchesResult, error) {
			args := v.(GetSwitchesArgs)
			r, err := GetSwitches(ctx, &args, opts...)
			var s GetSwitchesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSwitchesResultOutput)
}

// A collection of arguments for invoking getSwitches.
type GetSwitchesOutputArgs struct {
	// Filter results by a specific CIDR block. For example: "172.16.0.0/12".
	CidrBlock pulumi.StringPtrInput `pulumi:"cidrBlock"`
	// Specifies whether to precheck this request only. Valid values: `true` and `false`.
	DryRun pulumi.BoolPtrInput `pulumi:"dryRun"`
	// A list of VSwitch IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Indicate whether the VSwitch is created by the system.
	IsDefault pulumi.BoolPtrInput `pulumi:"isDefault"`
	// A regex string to filter results by name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Id of resource group which VSWitch belongs.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The route table ID of the VSwitch.
	RouteTableId pulumi.StringPtrInput `pulumi:"routeTableId"`
	// The status of the VSwitch. Valid values: `Available` and `Pending`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// ID of the VPC that owns the VSwitch.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// The name of the VSwitch.
	VswitchName pulumi.StringPtrInput `pulumi:"vswitchName"`
	// The VSwitch owner id.
	VswitchOwnerId pulumi.IntPtrInput `pulumi:"vswitchOwnerId"`
	// The availability zone of the VSwitch.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (GetSwitchesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSwitchesArgs)(nil)).Elem()
}

// A collection of values returned by getSwitches.
type GetSwitchesResultOutput struct{ *pulumi.OutputState }

func (GetSwitchesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSwitchesResult)(nil)).Elem()
}

func (o GetSwitchesResultOutput) ToGetSwitchesResultOutput() GetSwitchesResultOutput {
	return o
}

func (o GetSwitchesResultOutput) ToGetSwitchesResultOutputWithContext(ctx context.Context) GetSwitchesResultOutput {
	return o
}

// CIDR block of the VSwitch.
func (o GetSwitchesResultOutput) CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *string { return v.CidrBlock }).(pulumi.StringPtrOutput)
}

func (o GetSwitchesResultOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *bool { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSwitchesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSwitchesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of VSwitch IDs.
func (o GetSwitchesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSwitchesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// Whether the VSwitch is the default one in the region.
func (o GetSwitchesResultOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o GetSwitchesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of VSwitch names.
func (o GetSwitchesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSwitchesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetSwitchesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The resource group ID of the VSwitch.
func (o GetSwitchesResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The route table ID of the VSwitch.
func (o GetSwitchesResultOutput) RouteTableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *string { return v.RouteTableId }).(pulumi.StringPtrOutput)
}

// The status of the VSwitch.
func (o GetSwitchesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The Tags of the VSwitch.
func (o GetSwitchesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetSwitchesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// ID of the VPC that owns the VSwitch.
func (o GetSwitchesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// Name of the VSwitch.
func (o GetSwitchesResultOutput) VswitchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *string { return v.VswitchName }).(pulumi.StringPtrOutput)
}

func (o GetSwitchesResultOutput) VswitchOwnerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *int { return v.VswitchOwnerId }).(pulumi.IntPtrOutput)
}

// A list of VSwitches. Each element contains the following attributes:
func (o GetSwitchesResultOutput) Vswitches() GetSwitchesVswitchArrayOutput {
	return o.ApplyT(func(v GetSwitchesResult) []GetSwitchesVswitch { return v.Vswitches }).(GetSwitchesVswitchArrayOutput)
}

// ID of the availability zone where the VSwitch is located.
func (o GetSwitchesResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSwitchesResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSwitchesResultOutput{})
}
