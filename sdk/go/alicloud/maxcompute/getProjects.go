// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package maxcompute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Max Compute Project available to the user.[What is Project](https://help.aliyun.com/document_detail/473479.html)
//
// > **NOTE:** Available in 1.196.0+
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/maxcompute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf_testaccmp"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultProject, err := maxcompute.NewProject(ctx, "defaultProject", &maxcompute.ProjectArgs{
//				DefaultQuota: pulumi.String("默认后付费Quota"),
//				ProjectName:  pulumi.String(name),
//				Comment:      pulumi.String(name),
//				ProductType:  pulumi.String("PAYASYOUGO"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultProjects := defaultProject.ID().ApplyT(func(id string) (maxcompute.GetProjectsResult, error) {
//				return maxcompute.GetProjectsOutput(ctx, maxcompute.GetProjectsOutputArgs{
//					Ids: []string{
//						id,
//					},
//					NameRegex: defaultProject.Name,
//				}, nil), nil
//			}).(maxcompute.GetProjectsResultOutput)
//			ctx.Export("alicloudMaxcomputeProjectExampleId", defaultProjects.ApplyT(func(defaultProjects maxcompute.GetProjectsResult) (*string, error) {
//				return &defaultProjects.Projects[0].Id, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetProjects(ctx *pulumi.Context, args *GetProjectsArgs, opts ...pulumi.InvokeOption) (*GetProjectsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectsResult
	err := ctx.Invoke("alicloud:maxcompute/getProjects:getProjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjects.
type GetProjectsArgs struct {
	// A list of Project IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Group Metric Rule name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getProjects.
type GetProjectsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Project IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of name of Projects.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of Project Entries. Each element contains the following attributes:
	Projects []GetProjectsProject `pulumi:"projects"`
}

func GetProjectsOutput(ctx *pulumi.Context, args GetProjectsOutputArgs, opts ...pulumi.InvokeOption) GetProjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectsResult, error) {
			args := v.(GetProjectsArgs)
			r, err := GetProjects(ctx, &args, opts...)
			var s GetProjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectsResultOutput)
}

// A collection of arguments for invoking getProjects.
type GetProjectsOutputArgs struct {
	// A list of Project IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Group Metric Rule name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetProjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsArgs)(nil)).Elem()
}

// A collection of values returned by getProjects.
type GetProjectsResultOutput struct{ *pulumi.OutputState }

func (GetProjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsResult)(nil)).Elem()
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutput() GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutputWithContext(ctx context.Context) GetProjectsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Project IDs.
func (o GetProjectsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetProjectsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of name of Projects.
func (o GetProjectsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetProjectsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of Project Entries. Each element contains the following attributes:
func (o GetProjectsResultOutput) Projects() GetProjectsProjectArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []GetProjectsProject { return v.Projects }).(GetProjectsProjectArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectsResultOutput{})
}
