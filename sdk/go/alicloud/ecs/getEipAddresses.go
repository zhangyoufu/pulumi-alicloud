// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Eip Addresses of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.126.0+.
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
//			example, err := ecs.GetEipAddresses(ctx, &ecs.GetEipAddressesArgs{
//				Ids: []string{
//					"eip-bp1jvx5ki6c********",
//				},
//				NameRegex: pulumi.StringRef("the_resource_name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstEipAddressId", example.Addresses[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetEipAddresses(ctx *pulumi.Context, args *GetEipAddressesArgs, opts ...pulumi.InvokeOption) (*GetEipAddressesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEipAddressesResult
	err := ctx.Invoke("alicloud:ecs/getEipAddresses:getEipAddresses", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEipAddresses.
type GetEipAddressesArgs struct {
	// The name of the EIP.
	AddressName *string `pulumi:"addressName"`
	// The associated instance id.
	AssociatedInstanceId *string `pulumi:"associatedInstanceId"`
	// The associated instance type.
	AssociatedInstanceType *string `pulumi:"associatedInstanceType"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// Default to `true`. Set it to `false` can hide the `tags` to output.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Address IDs.
	Ids []string `pulumi:"ids"`
	// The include reservation data. Valid values: `BGP` and `BGP_PRO`.
	IncludeReservationData *bool `pulumi:"includeReservationData"`
	// The IP address of the EIP.
	IpAddress *string `pulumi:"ipAddress"`
	// Deprecated: Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.
	IpAddresses []string `pulumi:"ipAddresses"`
	// The Internet service provider (ISP).
	Isp *string `pulumi:"isp"`
	// The lock reason.
	LockReason *string `pulumi:"lockReason"`
	// A regex string to filter results by Address name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The billing method of the EIP.
	PaymentType *string `pulumi:"paymentType"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The IDs of the contiguous EIPs.
	SegmentInstanceId *string `pulumi:"segmentInstanceId"`
	// The status of the EIP.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getEipAddresses.
type GetEipAddressesResult struct {
	AddressName            *string                  `pulumi:"addressName"`
	Addresses              []GetEipAddressesAddress `pulumi:"addresses"`
	AssociatedInstanceId   *string                  `pulumi:"associatedInstanceId"`
	AssociatedInstanceType *string                  `pulumi:"associatedInstanceType"`
	DryRun                 *bool                    `pulumi:"dryRun"`
	// Deprecated: Field 'eips' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'addresses' instead.
	Eips          []GetEipAddressesEip `pulumi:"eips"`
	EnableDetails *bool                `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string   `pulumi:"id"`
	Ids                    []string `pulumi:"ids"`
	IncludeReservationData *bool    `pulumi:"includeReservationData"`
	IpAddress              *string  `pulumi:"ipAddress"`
	// Deprecated: Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.
	IpAddresses       []string               `pulumi:"ipAddresses"`
	Isp               *string                `pulumi:"isp"`
	LockReason        *string                `pulumi:"lockReason"`
	NameRegex         *string                `pulumi:"nameRegex"`
	Names             []string               `pulumi:"names"`
	OutputFile        *string                `pulumi:"outputFile"`
	PaymentType       *string                `pulumi:"paymentType"`
	ResourceGroupId   *string                `pulumi:"resourceGroupId"`
	SegmentInstanceId *string                `pulumi:"segmentInstanceId"`
	Status            *string                `pulumi:"status"`
	Tags              map[string]interface{} `pulumi:"tags"`
}

func GetEipAddressesOutput(ctx *pulumi.Context, args GetEipAddressesOutputArgs, opts ...pulumi.InvokeOption) GetEipAddressesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEipAddressesResult, error) {
			args := v.(GetEipAddressesArgs)
			r, err := GetEipAddresses(ctx, &args, opts...)
			var s GetEipAddressesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEipAddressesResultOutput)
}

// A collection of arguments for invoking getEipAddresses.
type GetEipAddressesOutputArgs struct {
	// The name of the EIP.
	AddressName pulumi.StringPtrInput `pulumi:"addressName"`
	// The associated instance id.
	AssociatedInstanceId pulumi.StringPtrInput `pulumi:"associatedInstanceId"`
	// The associated instance type.
	AssociatedInstanceType pulumi.StringPtrInput `pulumi:"associatedInstanceType"`
	// The dry run.
	DryRun pulumi.BoolPtrInput `pulumi:"dryRun"`
	// Default to `true`. Set it to `false` can hide the `tags` to output.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Address IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The include reservation data. Valid values: `BGP` and `BGP_PRO`.
	IncludeReservationData pulumi.BoolPtrInput `pulumi:"includeReservationData"`
	// The IP address of the EIP.
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// Deprecated: Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.
	IpAddresses pulumi.StringArrayInput `pulumi:"ipAddresses"`
	// The Internet service provider (ISP).
	Isp pulumi.StringPtrInput `pulumi:"isp"`
	// The lock reason.
	LockReason pulumi.StringPtrInput `pulumi:"lockReason"`
	// A regex string to filter results by Address name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The billing method of the EIP.
	PaymentType pulumi.StringPtrInput `pulumi:"paymentType"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The IDs of the contiguous EIPs.
	SegmentInstanceId pulumi.StringPtrInput `pulumi:"segmentInstanceId"`
	// The status of the EIP.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetEipAddressesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEipAddressesArgs)(nil)).Elem()
}

// A collection of values returned by getEipAddresses.
type GetEipAddressesResultOutput struct{ *pulumi.OutputState }

func (GetEipAddressesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEipAddressesResult)(nil)).Elem()
}

func (o GetEipAddressesResultOutput) ToGetEipAddressesResultOutput() GetEipAddressesResultOutput {
	return o
}

func (o GetEipAddressesResultOutput) ToGetEipAddressesResultOutputWithContext(ctx context.Context) GetEipAddressesResultOutput {
	return o
}

func (o GetEipAddressesResultOutput) AddressName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.AddressName }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) Addresses() GetEipAddressesAddressArrayOutput {
	return o.ApplyT(func(v GetEipAddressesResult) []GetEipAddressesAddress { return v.Addresses }).(GetEipAddressesAddressArrayOutput)
}

func (o GetEipAddressesResultOutput) AssociatedInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.AssociatedInstanceId }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) AssociatedInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.AssociatedInstanceType }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *bool { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// Deprecated: Field 'eips' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'addresses' instead.
func (o GetEipAddressesResultOutput) Eips() GetEipAddressesEipArrayOutput {
	return o.ApplyT(func(v GetEipAddressesResult) []GetEipAddressesEip { return v.Eips }).(GetEipAddressesEipArrayOutput)
}

func (o GetEipAddressesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEipAddressesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEipAddressesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEipAddressesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEipAddressesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEipAddressesResultOutput) IncludeReservationData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *bool { return v.IncludeReservationData }).(pulumi.BoolPtrOutput)
}

func (o GetEipAddressesResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// Deprecated: Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.
func (o GetEipAddressesResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEipAddressesResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o GetEipAddressesResultOutput) Isp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.Isp }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) LockReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.LockReason }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEipAddressesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEipAddressesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) PaymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.PaymentType }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) SegmentInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.SegmentInstanceId }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEipAddressesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetEipAddressesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetEipAddressesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEipAddressesResultOutput{})
}
