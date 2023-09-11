// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ens

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides the Ens Key Pairs of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.133.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ens"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			nameRegex, err := ens.GetKeyPairs(ctx, &ens.GetKeyPairsArgs{
//				Version:   "example_value",
//				NameRegex: pulumi.StringRef("^my-KeyPair"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ensKeyPairId1", nameRegex.Pairs[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetKeyPairs(ctx *pulumi.Context, args *GetKeyPairsArgs, opts ...pulumi.InvokeOption) (*GetKeyPairsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKeyPairsResult
	err := ctx.Invoke("alicloud:ens/getKeyPairs:getKeyPairs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKeyPairs.
type GetKeyPairsArgs struct {
	// The name of the key pair.
	KeyPairName *string `pulumi:"keyPairName"`
	// A regex string to filter results by Key Pair name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The version number.
	Version string `pulumi:"version"`
}

// A collection of values returned by getKeyPairs.
type GetKeyPairsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string            `pulumi:"id"`
	Ids         []string          `pulumi:"ids"`
	KeyPairName *string           `pulumi:"keyPairName"`
	NameRegex   *string           `pulumi:"nameRegex"`
	Names       []string          `pulumi:"names"`
	OutputFile  *string           `pulumi:"outputFile"`
	Pairs       []GetKeyPairsPair `pulumi:"pairs"`
	Version     string            `pulumi:"version"`
}

func GetKeyPairsOutput(ctx *pulumi.Context, args GetKeyPairsOutputArgs, opts ...pulumi.InvokeOption) GetKeyPairsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKeyPairsResult, error) {
			args := v.(GetKeyPairsArgs)
			r, err := GetKeyPairs(ctx, &args, opts...)
			var s GetKeyPairsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKeyPairsResultOutput)
}

// A collection of arguments for invoking getKeyPairs.
type GetKeyPairsOutputArgs struct {
	// The name of the key pair.
	KeyPairName pulumi.StringPtrInput `pulumi:"keyPairName"`
	// A regex string to filter results by Key Pair name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The version number.
	Version pulumi.StringInput `pulumi:"version"`
}

func (GetKeyPairsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyPairsArgs)(nil)).Elem()
}

// A collection of values returned by getKeyPairs.
type GetKeyPairsResultOutput struct{ *pulumi.OutputState }

func (GetKeyPairsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyPairsResult)(nil)).Elem()
}

func (o GetKeyPairsResultOutput) ToGetKeyPairsResultOutput() GetKeyPairsResultOutput {
	return o
}

func (o GetKeyPairsResultOutput) ToGetKeyPairsResultOutputWithContext(ctx context.Context) GetKeyPairsResultOutput {
	return o
}

func (o GetKeyPairsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetKeyPairsResult] {
	return pulumix.Output[GetKeyPairsResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetKeyPairsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyPairsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetKeyPairsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKeyPairsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetKeyPairsResultOutput) KeyPairName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKeyPairsResult) *string { return v.KeyPairName }).(pulumi.StringPtrOutput)
}

func (o GetKeyPairsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKeyPairsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetKeyPairsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKeyPairsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetKeyPairsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKeyPairsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetKeyPairsResultOutput) Pairs() GetKeyPairsPairArrayOutput {
	return o.ApplyT(func(v GetKeyPairsResult) []GetKeyPairsPair { return v.Pairs }).(GetKeyPairsPairArrayOutput)
}

func (o GetKeyPairsResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyPairsResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKeyPairsResultOutput{})
}
