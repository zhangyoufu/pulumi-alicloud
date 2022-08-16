// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides CEN flow logs available to the user.
//
// > **NOTE:** Available in 1.78.0+
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.GetFlowlogs(ctx, &cen.GetFlowlogsArgs{
//				Ids: []string{
//					"flowlog-tig1xxxxx",
//				},
//				NameRegex: pulumi.StringRef("^foo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstCenFlowlogId", data.Alicloud_cen_instances.Default.Flowlogs[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetFlowlogs(ctx *pulumi.Context, args *GetFlowlogsArgs, opts ...pulumi.InvokeOption) (*GetFlowlogsResult, error) {
	var rv GetFlowlogsResult
	err := ctx.Invoke("alicloud:cen/getFlowlogs:getFlowlogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlowlogs.
type GetFlowlogsArgs struct {
	// The ID of the CEN Instance.
	CenId *string `pulumi:"cenId"`
	// The description of flowlog.
	Description *string `pulumi:"description"`
	// A list of CEN flow log IDs.
	Ids []string `pulumi:"ids"`
	// The name of the log store which is in the  `projectName` SLS project.
	LogStoreName *string `pulumi:"logStoreName"`
	// A regex string to filter CEN flow logs by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The name of the SLS project.
	ProjectName *string `pulumi:"projectName"`
	// The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
	Status *string `pulumi:"status"`
}

// A collection of values returned by getFlowlogs.
type GetFlowlogsResult struct {
	// The ID of the CEN Instance.
	CenId *string `pulumi:"cenId"`
	// The description of flowlog.
	Description *string              `pulumi:"description"`
	Flowlogs    []GetFlowlogsFlowlog `pulumi:"flowlogs"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of CEN flow log IDs.
	Ids []string `pulumi:"ids"`
	// The name of the log store which is in the  `projectName` SLS project.
	LogStoreName *string `pulumi:"logStoreName"`
	NameRegex    *string `pulumi:"nameRegex"`
	// A list of CEN flow log names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The name of the SLS project.
	ProjectName *string `pulumi:"projectName"`
	// The status of flowlog.
	Status *string `pulumi:"status"`
}

func GetFlowlogsOutput(ctx *pulumi.Context, args GetFlowlogsOutputArgs, opts ...pulumi.InvokeOption) GetFlowlogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFlowlogsResult, error) {
			args := v.(GetFlowlogsArgs)
			r, err := GetFlowlogs(ctx, &args, opts...)
			var s GetFlowlogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFlowlogsResultOutput)
}

// A collection of arguments for invoking getFlowlogs.
type GetFlowlogsOutputArgs struct {
	// The ID of the CEN Instance.
	CenId pulumi.StringPtrInput `pulumi:"cenId"`
	// The description of flowlog.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// A list of CEN flow log IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The name of the log store which is in the  `projectName` SLS project.
	LogStoreName pulumi.StringPtrInput `pulumi:"logStoreName"`
	// A regex string to filter CEN flow logs by name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of the SLS project.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetFlowlogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlowlogsArgs)(nil)).Elem()
}

// A collection of values returned by getFlowlogs.
type GetFlowlogsResultOutput struct{ *pulumi.OutputState }

func (GetFlowlogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlowlogsResult)(nil)).Elem()
}

func (o GetFlowlogsResultOutput) ToGetFlowlogsResultOutput() GetFlowlogsResultOutput {
	return o
}

func (o GetFlowlogsResultOutput) ToGetFlowlogsResultOutputWithContext(ctx context.Context) GetFlowlogsResultOutput {
	return o
}

// The ID of the CEN Instance.
func (o GetFlowlogsResultOutput) CenId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlowlogsResult) *string { return v.CenId }).(pulumi.StringPtrOutput)
}

// The description of flowlog.
func (o GetFlowlogsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlowlogsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetFlowlogsResultOutput) Flowlogs() GetFlowlogsFlowlogArrayOutput {
	return o.ApplyT(func(v GetFlowlogsResult) []GetFlowlogsFlowlog { return v.Flowlogs }).(GetFlowlogsFlowlogArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFlowlogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlowlogsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of CEN flow log IDs.
func (o GetFlowlogsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFlowlogsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The name of the log store which is in the  `projectName` SLS project.
func (o GetFlowlogsResultOutput) LogStoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlowlogsResult) *string { return v.LogStoreName }).(pulumi.StringPtrOutput)
}

func (o GetFlowlogsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlowlogsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of CEN flow log names.
func (o GetFlowlogsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFlowlogsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetFlowlogsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlowlogsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The name of the SLS project.
func (o GetFlowlogsResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlowlogsResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The status of flowlog.
func (o GetFlowlogsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlowlogsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFlowlogsResultOutput{})
}
