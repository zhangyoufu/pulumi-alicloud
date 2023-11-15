// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **DEPRECATED:**  This datasource has been deprecated from version `1.126.0`. Please use new datasource alicloud_eip_addresses.
//
// This data source provides a list of EIPs (Elastic IP address) owned by an Alibaba Cloud account.
//
// ## Example Usage
//
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
//			eipsDs, err := ecs.GetEips(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstEipId", eipsDs.Eips[0].Id)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: This function has been deprecated in favour of the getEipAddresses function
func GetEips(ctx *pulumi.Context, args *GetEipsArgs, opts ...pulumi.InvokeOption) (*GetEipsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEipsResult
	err := ctx.Invoke("alicloud:ecs/getEips:getEips", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEips.
type GetEipsArgs struct {
	AddressName            *string `pulumi:"addressName"`
	AssociatedInstanceId   *string `pulumi:"associatedInstanceId"`
	AssociatedInstanceType *string `pulumi:"associatedInstanceType"`
	DryRun                 *bool   `pulumi:"dryRun"`
	EnableDetails          *bool   `pulumi:"enableDetails"`
	// A list of EIP IDs.
	Ids                    []string `pulumi:"ids"`
	IncludeReservationData *bool    `pulumi:"includeReservationData"`
	// Public IP Address of the the EIP.
	IpAddress *string `pulumi:"ipAddress"`
	// A list of EIP public IP addresses.
	//
	// Deprecated: Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.
	IpAddresses []string `pulumi:"ipAddresses"`
	Isp         *string  `pulumi:"isp"`
	LockReason  *string  `pulumi:"lockReason"`
	NameRegex   *string  `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile  *string `pulumi:"outputFile"`
	PaymentType *string `pulumi:"paymentType"`
	// The Id of resource group which the eips belongs.
	ResourceGroupId   *string `pulumi:"resourceGroupId"`
	SegmentInstanceId *string `pulumi:"segmentInstanceId"`
	// EIP status. Possible values are: `Associating`, `Unassociating`, `InUse` and `Available`.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getEips.
type GetEipsResult struct {
	AddressName            *string          `pulumi:"addressName"`
	Addresses              []GetEipsAddress `pulumi:"addresses"`
	AssociatedInstanceId   *string          `pulumi:"associatedInstanceId"`
	AssociatedInstanceType *string          `pulumi:"associatedInstanceType"`
	DryRun                 *bool            `pulumi:"dryRun"`
	// A list of EIPs. Each element contains the following attributes:
	//
	// Deprecated: Field 'eips' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'addresses' instead.
	Eips          []GetEipsEip `pulumi:"eips"`
	EnableDetails *bool        `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Optional) A list of EIP IDs.
	Ids                    []string `pulumi:"ids"`
	IncludeReservationData *bool    `pulumi:"includeReservationData"`
	// Public IP Address of the the EIP.
	IpAddress *string `pulumi:"ipAddress"`
	// Deprecated: Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.
	IpAddresses []string `pulumi:"ipAddresses"`
	Isp         *string  `pulumi:"isp"`
	LockReason  *string  `pulumi:"lockReason"`
	NameRegex   *string  `pulumi:"nameRegex"`
	// (Optional) A list of EIP names.
	Names       []string `pulumi:"names"`
	OutputFile  *string  `pulumi:"outputFile"`
	PaymentType *string  `pulumi:"paymentType"`
	// The Id of resource group which the eips belongs.
	ResourceGroupId   *string `pulumi:"resourceGroupId"`
	SegmentInstanceId *string `pulumi:"segmentInstanceId"`
	// EIP status. Possible values are: `Associating`, `Unassociating`, `InUse` and `Available`.
	Status *string                `pulumi:"status"`
	Tags   map[string]interface{} `pulumi:"tags"`
}

func GetEipsOutput(ctx *pulumi.Context, args GetEipsOutputArgs, opts ...pulumi.InvokeOption) GetEipsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEipsResult, error) {
			args := v.(GetEipsArgs)
			r, err := GetEips(ctx, &args, opts...)
			var s GetEipsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEipsResultOutput)
}

// A collection of arguments for invoking getEips.
type GetEipsOutputArgs struct {
	AddressName            pulumi.StringPtrInput `pulumi:"addressName"`
	AssociatedInstanceId   pulumi.StringPtrInput `pulumi:"associatedInstanceId"`
	AssociatedInstanceType pulumi.StringPtrInput `pulumi:"associatedInstanceType"`
	DryRun                 pulumi.BoolPtrInput   `pulumi:"dryRun"`
	EnableDetails          pulumi.BoolPtrInput   `pulumi:"enableDetails"`
	// A list of EIP IDs.
	Ids                    pulumi.StringArrayInput `pulumi:"ids"`
	IncludeReservationData pulumi.BoolPtrInput     `pulumi:"includeReservationData"`
	// Public IP Address of the the EIP.
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// A list of EIP public IP addresses.
	//
	// Deprecated: Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.
	IpAddresses pulumi.StringArrayInput `pulumi:"ipAddresses"`
	Isp         pulumi.StringPtrInput   `pulumi:"isp"`
	LockReason  pulumi.StringPtrInput   `pulumi:"lockReason"`
	NameRegex   pulumi.StringPtrInput   `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile  pulumi.StringPtrInput `pulumi:"outputFile"`
	PaymentType pulumi.StringPtrInput `pulumi:"paymentType"`
	// The Id of resource group which the eips belongs.
	ResourceGroupId   pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	SegmentInstanceId pulumi.StringPtrInput `pulumi:"segmentInstanceId"`
	// EIP status. Possible values are: `Associating`, `Unassociating`, `InUse` and `Available`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetEipsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEipsArgs)(nil)).Elem()
}

// A collection of values returned by getEips.
type GetEipsResultOutput struct{ *pulumi.OutputState }

func (GetEipsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEipsResult)(nil)).Elem()
}

func (o GetEipsResultOutput) ToGetEipsResultOutput() GetEipsResultOutput {
	return o
}

func (o GetEipsResultOutput) ToGetEipsResultOutputWithContext(ctx context.Context) GetEipsResultOutput {
	return o
}

func (o GetEipsResultOutput) AddressName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.AddressName }).(pulumi.StringPtrOutput)
}

func (o GetEipsResultOutput) Addresses() GetEipsAddressArrayOutput {
	return o.ApplyT(func(v GetEipsResult) []GetEipsAddress { return v.Addresses }).(GetEipsAddressArrayOutput)
}

func (o GetEipsResultOutput) AssociatedInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.AssociatedInstanceId }).(pulumi.StringPtrOutput)
}

func (o GetEipsResultOutput) AssociatedInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.AssociatedInstanceType }).(pulumi.StringPtrOutput)
}

func (o GetEipsResultOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *bool { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// A list of EIPs. Each element contains the following attributes:
//
// Deprecated: Field 'eips' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'addresses' instead.
func (o GetEipsResultOutput) Eips() GetEipsEipArrayOutput {
	return o.ApplyT(func(v GetEipsResult) []GetEipsEip { return v.Eips }).(GetEipsEipArrayOutput)
}

func (o GetEipsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEipsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipsResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Optional) A list of EIP IDs.
func (o GetEipsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEipsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEipsResultOutput) IncludeReservationData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *bool { return v.IncludeReservationData }).(pulumi.BoolPtrOutput)
}

// Public IP Address of the the EIP.
func (o GetEipsResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// Deprecated: Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.
func (o GetEipsResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEipsResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o GetEipsResultOutput) Isp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.Isp }).(pulumi.StringPtrOutput)
}

func (o GetEipsResultOutput) LockReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.LockReason }).(pulumi.StringPtrOutput)
}

func (o GetEipsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// (Optional) A list of EIP names.
func (o GetEipsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEipsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEipsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEipsResultOutput) PaymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.PaymentType }).(pulumi.StringPtrOutput)
}

// The Id of resource group which the eips belongs.
func (o GetEipsResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetEipsResultOutput) SegmentInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.SegmentInstanceId }).(pulumi.StringPtrOutput)
}

// EIP status. Possible values are: `Associating`, `Unassociating`, `InUse` and `Available`.
func (o GetEipsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetEipsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetEipsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEipsResultOutput{})
}
