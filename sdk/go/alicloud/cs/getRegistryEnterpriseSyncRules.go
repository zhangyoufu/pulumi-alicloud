// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list Container Registry Enterprise Edition sync rules on Alibaba Cloud.
//
// > **NOTE:** Available in v1.90.0+
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Declare the data source
//			mySyncRules, err := cs.GetRegistryEnterpriseSyncRules(ctx, &cs.GetRegistryEnterpriseSyncRulesArgs{
//				InstanceId:       "cri-xxx",
//				NamespaceName:    pulumi.StringRef("test-namespace"),
//				RepoName:         pulumi.StringRef("test-repo"),
//				TargetInstanceId: pulumi.StringRef("cri-yyy"),
//				NameRegex:        pulumi.StringRef("test-rule"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			var splat0 []*string
//			for _, val0 := range mySyncRules.Rules {
//				splat0 = append(splat0, val0.Id)
//			}
//			ctx.Export("output", splat0)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetRegistryEnterpriseSyncRules(ctx *pulumi.Context, args *GetRegistryEnterpriseSyncRulesArgs, opts ...pulumi.InvokeOption) (*GetRegistryEnterpriseSyncRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRegistryEnterpriseSyncRulesResult
	err := ctx.Invoke("alicloud:cs/getRegistryEnterpriseSyncRules:getRegistryEnterpriseSyncRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistryEnterpriseSyncRules.
type GetRegistryEnterpriseSyncRulesArgs struct {
	// A list of ids to filter results by sync rule id.
	Ids []string `pulumi:"ids"`
	// ID of Container Registry Enterprise Edition local instance.
	InstanceId string `pulumi:"instanceId"`
	// A regex string to filter results by sync rule name.
	NameRegex *string `pulumi:"nameRegex"`
	// Name of Container Registry Enterprise Edition local namespace.
	NamespaceName *string `pulumi:"namespaceName"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// Name of Container Registry Enterprise Edition local repo.
	RepoName *string `pulumi:"repoName"`
	// ID of Container Registry Enterprise Edition target instance.
	TargetInstanceId *string `pulumi:"targetInstanceId"`
}

// A collection of values returned by getRegistryEnterpriseSyncRules.
type GetRegistryEnterpriseSyncRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of matched Container Registry Enterprise Edition sync rules. Its element is a sync rule uuid.
	Ids []string `pulumi:"ids"`
	// ID of Container Registry Enterprise Edition local instance.
	InstanceId string  `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	// A list of sync rule names.
	Names []string `pulumi:"names"`
	// Name of Container Registry Enterprise Edition local namespace.
	NamespaceName *string `pulumi:"namespaceName"`
	OutputFile    *string `pulumi:"outputFile"`
	// Name of Container Registry Enterprise Edition local repo.
	RepoName *string `pulumi:"repoName"`
	// A list of matched Container Registry Enterprise Edition sync rules. Each element contains the following attributes:
	Rules []GetRegistryEnterpriseSyncRulesRule `pulumi:"rules"`
	// ID of Container Registry Enterprise Edition target instance.
	TargetInstanceId *string `pulumi:"targetInstanceId"`
}

func GetRegistryEnterpriseSyncRulesOutput(ctx *pulumi.Context, args GetRegistryEnterpriseSyncRulesOutputArgs, opts ...pulumi.InvokeOption) GetRegistryEnterpriseSyncRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegistryEnterpriseSyncRulesResult, error) {
			args := v.(GetRegistryEnterpriseSyncRulesArgs)
			r, err := GetRegistryEnterpriseSyncRules(ctx, &args, opts...)
			var s GetRegistryEnterpriseSyncRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegistryEnterpriseSyncRulesResultOutput)
}

// A collection of arguments for invoking getRegistryEnterpriseSyncRules.
type GetRegistryEnterpriseSyncRulesOutputArgs struct {
	// A list of ids to filter results by sync rule id.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// ID of Container Registry Enterprise Edition local instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A regex string to filter results by sync rule name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// Name of Container Registry Enterprise Edition local namespace.
	NamespaceName pulumi.StringPtrInput `pulumi:"namespaceName"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Name of Container Registry Enterprise Edition local repo.
	RepoName pulumi.StringPtrInput `pulumi:"repoName"`
	// ID of Container Registry Enterprise Edition target instance.
	TargetInstanceId pulumi.StringPtrInput `pulumi:"targetInstanceId"`
}

func (GetRegistryEnterpriseSyncRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryEnterpriseSyncRulesArgs)(nil)).Elem()
}

// A collection of values returned by getRegistryEnterpriseSyncRules.
type GetRegistryEnterpriseSyncRulesResultOutput struct{ *pulumi.OutputState }

func (GetRegistryEnterpriseSyncRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryEnterpriseSyncRulesResult)(nil)).Elem()
}

func (o GetRegistryEnterpriseSyncRulesResultOutput) ToGetRegistryEnterpriseSyncRulesResultOutput() GetRegistryEnterpriseSyncRulesResultOutput {
	return o
}

func (o GetRegistryEnterpriseSyncRulesResultOutput) ToGetRegistryEnterpriseSyncRulesResultOutputWithContext(ctx context.Context) GetRegistryEnterpriseSyncRulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRegistryEnterpriseSyncRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegistryEnterpriseSyncRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of matched Container Registry Enterprise Edition sync rules. Its element is a sync rule uuid.
func (o GetRegistryEnterpriseSyncRulesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRegistryEnterpriseSyncRulesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// ID of Container Registry Enterprise Edition local instance.
func (o GetRegistryEnterpriseSyncRulesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegistryEnterpriseSyncRulesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetRegistryEnterpriseSyncRulesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryEnterpriseSyncRulesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of sync rule names.
func (o GetRegistryEnterpriseSyncRulesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRegistryEnterpriseSyncRulesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

// Name of Container Registry Enterprise Edition local namespace.
func (o GetRegistryEnterpriseSyncRulesResultOutput) NamespaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryEnterpriseSyncRulesResult) *string { return v.NamespaceName }).(pulumi.StringPtrOutput)
}

func (o GetRegistryEnterpriseSyncRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryEnterpriseSyncRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// Name of Container Registry Enterprise Edition local repo.
func (o GetRegistryEnterpriseSyncRulesResultOutput) RepoName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryEnterpriseSyncRulesResult) *string { return v.RepoName }).(pulumi.StringPtrOutput)
}

// A list of matched Container Registry Enterprise Edition sync rules. Each element contains the following attributes:
func (o GetRegistryEnterpriseSyncRulesResultOutput) Rules() GetRegistryEnterpriseSyncRulesRuleArrayOutput {
	return o.ApplyT(func(v GetRegistryEnterpriseSyncRulesResult) []GetRegistryEnterpriseSyncRulesRule { return v.Rules }).(GetRegistryEnterpriseSyncRulesRuleArrayOutput)
}

// ID of Container Registry Enterprise Edition target instance.
func (o GetRegistryEnterpriseSyncRulesResultOutput) TargetInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryEnterpriseSyncRulesResult) *string { return v.TargetInstanceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegistryEnterpriseSyncRulesResultOutput{})
}
