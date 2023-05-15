// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cloud Architect Design Tools (BPStudio) Applications of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.192.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/bp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := bp.GetStudioApplications(ctx, &bp.GetStudioApplicationsArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("bpStudioApplicationId1", ids.Applications[0].Id)
//			nameRegex, err := bp.GetStudioApplications(ctx, &bp.GetStudioApplicationsArgs{
//				NameRegex: pulumi.StringRef("^my-Application"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("bpStudioApplicationId2", nameRegex.Applications[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetStudioApplications(ctx *pulumi.Context, args *GetStudioApplicationsArgs, opts ...pulumi.InvokeOption) (*GetStudioApplicationsResult, error) {
	var rv GetStudioApplicationsResult
	err := ctx.Invoke("alicloud:bp/getStudioApplications:getStudioApplications", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStudioApplications.
type GetStudioApplicationsArgs struct {
	// A list of Application IDs.
	Ids []string `pulumi:"ids"`
	// The keyword of the Application.
	Keyword    *string `pulumi:"keyword"`
	MaxResults *int    `pulumi:"maxResults"`
	// A regex string to filter results by Application name.
	NameRegex *string `pulumi:"nameRegex"`
	// The order type of the Application. Valid values:
	OrderType *int `pulumi:"orderType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the Application. Valid values: `success`, `release`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getStudioApplications.
type GetStudioApplicationsResult struct {
	// A list of Cloud Architect Design Tools (BPStudio) Applications. Each element contains the following attributes:
	Applications []GetStudioApplicationsApplication `pulumi:"applications"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	Keyword    *string  `pulumi:"keyword"`
	MaxResults *int     `pulumi:"maxResults"`
	NameRegex  *string  `pulumi:"nameRegex"`
	// A list of Application names.
	Names      []string `pulumi:"names"`
	OrderType  *int     `pulumi:"orderType"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the Application.
	Status *string `pulumi:"status"`
}

func GetStudioApplicationsOutput(ctx *pulumi.Context, args GetStudioApplicationsOutputArgs, opts ...pulumi.InvokeOption) GetStudioApplicationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStudioApplicationsResult, error) {
			args := v.(GetStudioApplicationsArgs)
			r, err := GetStudioApplications(ctx, &args, opts...)
			var s GetStudioApplicationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStudioApplicationsResultOutput)
}

// A collection of arguments for invoking getStudioApplications.
type GetStudioApplicationsOutputArgs struct {
	// A list of Application IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The keyword of the Application.
	Keyword    pulumi.StringPtrInput `pulumi:"keyword"`
	MaxResults pulumi.IntPtrInput    `pulumi:"maxResults"`
	// A regex string to filter results by Application name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The order type of the Application. Valid values:
	OrderType pulumi.IntPtrInput `pulumi:"orderType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The status of the Application. Valid values: `success`, `release`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetStudioApplicationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStudioApplicationsArgs)(nil)).Elem()
}

// A collection of values returned by getStudioApplications.
type GetStudioApplicationsResultOutput struct{ *pulumi.OutputState }

func (GetStudioApplicationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStudioApplicationsResult)(nil)).Elem()
}

func (o GetStudioApplicationsResultOutput) ToGetStudioApplicationsResultOutput() GetStudioApplicationsResultOutput {
	return o
}

func (o GetStudioApplicationsResultOutput) ToGetStudioApplicationsResultOutputWithContext(ctx context.Context) GetStudioApplicationsResultOutput {
	return o
}

// A list of Cloud Architect Design Tools (BPStudio) Applications. Each element contains the following attributes:
func (o GetStudioApplicationsResultOutput) Applications() GetStudioApplicationsApplicationArrayOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) []GetStudioApplicationsApplication { return v.Applications }).(GetStudioApplicationsApplicationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStudioApplicationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStudioApplicationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetStudioApplicationsResultOutput) Keyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) *string { return v.Keyword }).(pulumi.StringPtrOutput)
}

func (o GetStudioApplicationsResultOutput) MaxResults() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) *int { return v.MaxResults }).(pulumi.IntPtrOutput)
}

func (o GetStudioApplicationsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of Application names.
func (o GetStudioApplicationsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetStudioApplicationsResultOutput) OrderType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) *int { return v.OrderType }).(pulumi.IntPtrOutput)
}

func (o GetStudioApplicationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ID of the resource group.
func (o GetStudioApplicationsResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The status of the Application.
func (o GetStudioApplicationsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStudioApplicationsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStudioApplicationsResultOutput{})
}
