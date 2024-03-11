// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecs Network Interfaces of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.123.1+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := ecs.GetEcsNetworkInterfaces(ctx, &ecs.GetEcsNetworkInterfacesArgs{
//				Ids: []string{
//					"eni-abcd1234",
//				},
//				NameRegex: pulumi.StringRef("tf-testAcc"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstEcsNetworkInterfaceId", example.Interfaces[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetEcsNetworkInterfaces(ctx *pulumi.Context, args *GetEcsNetworkInterfacesArgs, opts ...pulumi.InvokeOption) (*GetEcsNetworkInterfacesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEcsNetworkInterfacesResult
	err := ctx.Invoke("alicloud:ecs/getEcsNetworkInterfaces:getEcsNetworkInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEcsNetworkInterfaces.
type GetEcsNetworkInterfacesArgs struct {
	// A list of Network Interface IDs.
	Ids []string `pulumi:"ids"`
	// The instance id.
	InstanceId *string `pulumi:"instanceId"`
	// The network interface name.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name *string `pulumi:"name"`
	// A regex string to filter results by Network Interface name.
	NameRegex *string `pulumi:"nameRegex"`
	// The network interface name.
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The primary private IP address of the ENI.
	PrimaryIpAddress *string `pulumi:"primaryIpAddress"`
	// The primary private IP address of the ENI.
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp *string `pulumi:"privateIp"`
	// The resource group id.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The security group id.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// Whether the user of the elastic network card is a cloud product or a virtual vendor.
	ServiceManaged *bool `pulumi:"serviceManaged"`
	// The status of the ENI.
	Status *string `pulumi:"status"`
	// The tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// The type of the ENI.
	Type *string `pulumi:"type"`
	// The Vpc Id.
	VpcId *string `pulumi:"vpcId"`
	// The vswitch id.
	VswitchId *string `pulumi:"vswitchId"`
}

// A collection of values returned by getEcsNetworkInterfaces.
type GetEcsNetworkInterfacesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                             `pulumi:"id"`
	Ids        []string                           `pulumi:"ids"`
	InstanceId *string                            `pulumi:"instanceId"`
	Interfaces []GetEcsNetworkInterfacesInterface `pulumi:"interfaces"`
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name                 *string  `pulumi:"name"`
	NameRegex            *string  `pulumi:"nameRegex"`
	Names                []string `pulumi:"names"`
	NetworkInterfaceName *string  `pulumi:"networkInterfaceName"`
	OutputFile           *string  `pulumi:"outputFile"`
	PrimaryIpAddress     *string  `pulumi:"primaryIpAddress"`
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp       *string                `pulumi:"privateIp"`
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	SecurityGroupId *string                `pulumi:"securityGroupId"`
	ServiceManaged  *bool                  `pulumi:"serviceManaged"`
	Status          *string                `pulumi:"status"`
	Tags            map[string]interface{} `pulumi:"tags"`
	Type            *string                `pulumi:"type"`
	VpcId           *string                `pulumi:"vpcId"`
	VswitchId       *string                `pulumi:"vswitchId"`
}

func GetEcsNetworkInterfacesOutput(ctx *pulumi.Context, args GetEcsNetworkInterfacesOutputArgs, opts ...pulumi.InvokeOption) GetEcsNetworkInterfacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEcsNetworkInterfacesResult, error) {
			args := v.(GetEcsNetworkInterfacesArgs)
			r, err := GetEcsNetworkInterfaces(ctx, &args, opts...)
			var s GetEcsNetworkInterfacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEcsNetworkInterfacesResultOutput)
}

// A collection of arguments for invoking getEcsNetworkInterfaces.
type GetEcsNetworkInterfacesOutputArgs struct {
	// A list of Network Interface IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The instance id.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The network interface name.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A regex string to filter results by Network Interface name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The network interface name.
	NetworkInterfaceName pulumi.StringPtrInput `pulumi:"networkInterfaceName"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The primary private IP address of the ENI.
	PrimaryIpAddress pulumi.StringPtrInput `pulumi:"primaryIpAddress"`
	// The primary private IP address of the ENI.
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp pulumi.StringPtrInput `pulumi:"privateIp"`
	// The resource group id.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The security group id.
	SecurityGroupId pulumi.StringPtrInput `pulumi:"securityGroupId"`
	// Whether the user of the elastic network card is a cloud product or a virtual vendor.
	ServiceManaged pulumi.BoolPtrInput `pulumi:"serviceManaged"`
	// The status of the ENI.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The tags.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The type of the ENI.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// The Vpc Id.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// The vswitch id.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
}

func (GetEcsNetworkInterfacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsNetworkInterfacesArgs)(nil)).Elem()
}

// A collection of values returned by getEcsNetworkInterfaces.
type GetEcsNetworkInterfacesResultOutput struct{ *pulumi.OutputState }

func (GetEcsNetworkInterfacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsNetworkInterfacesResult)(nil)).Elem()
}

func (o GetEcsNetworkInterfacesResultOutput) ToGetEcsNetworkInterfacesResultOutput() GetEcsNetworkInterfacesResultOutput {
	return o
}

func (o GetEcsNetworkInterfacesResultOutput) ToGetEcsNetworkInterfacesResultOutputWithContext(ctx context.Context) GetEcsNetworkInterfacesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetEcsNetworkInterfacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) Interfaces() GetEcsNetworkInterfacesInterfaceArrayOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) []GetEcsNetworkInterfacesInterface { return v.Interfaces }).(GetEcsNetworkInterfacesInterfaceArrayOutput)
}

// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
func (o GetEcsNetworkInterfacesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) NetworkInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.NetworkInterfaceName }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) PrimaryIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.PrimaryIpAddress }).(pulumi.StringPtrOutput)
}

// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
func (o GetEcsNetworkInterfacesResultOutput) PrivateIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.PrivateIp }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) SecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.SecurityGroupId }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) ServiceManaged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *bool { return v.ServiceManaged }).(pulumi.BoolPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func (o GetEcsNetworkInterfacesResultOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsNetworkInterfacesResult) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEcsNetworkInterfacesResultOutput{})
}
