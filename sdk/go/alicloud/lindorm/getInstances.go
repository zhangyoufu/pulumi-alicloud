// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lindorm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Lindorm Instances of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.132.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/lindorm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := lindorm.GetInstances(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("lindormInstanceId1", ids.Instances[0].Id)
//			nameRegex, err := lindorm.GetInstances(ctx, &lindorm.GetInstancesArgs{
//				NameRegex: pulumi.StringRef("^my-Instance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("lindormInstanceId2", nameRegex.Instances[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:lindorm/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Instance IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Instance name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The query str, which can use `instanceName` keyword for fuzzy search.
	QueryStr *string `pulumi:"queryStr"`
	// The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
	Status *string `pulumi:"status"`
	// The support engine. Valid values: `1` to `7`.
	SupportEngine *int `pulumi:"supportEngine"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id            string                 `pulumi:"id"`
	Ids           []string               `pulumi:"ids"`
	Instances     []GetInstancesInstance `pulumi:"instances"`
	NameRegex     *string                `pulumi:"nameRegex"`
	Names         []string               `pulumi:"names"`
	OutputFile    *string                `pulumi:"outputFile"`
	QueryStr      *string                `pulumi:"queryStr"`
	Status        *string                `pulumi:"status"`
	SupportEngine *int                   `pulumi:"supportEngine"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			var s GetInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Instance IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Instance name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The query str, which can use `instanceName` keyword for fuzzy search.
	QueryStr pulumi.StringPtrInput `pulumi:"queryStr"`
	// The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The support engine. Valid values: `1` to `7`.
	SupportEngine pulumi.IntPtrInput `pulumi:"supportEngine"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) Instances() GetInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstance { return v.Instances }).(GetInstancesInstanceArrayOutput)
}

func (o GetInstancesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) QueryStr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.QueryStr }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) SupportEngine() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *int { return v.SupportEngine }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
