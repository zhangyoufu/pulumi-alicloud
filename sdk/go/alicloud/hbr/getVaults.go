// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Hbr Vaults of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.129.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/hbr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := hbr.GetVaults(ctx, &hbr.GetVaultsArgs{
//				NameRegex: pulumi.StringRef("^my-Vault"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("hbrVaultId1", ids.Vaults[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetVaults(ctx *pulumi.Context, args *GetVaultsArgs, opts ...pulumi.InvokeOption) (*GetVaultsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVaultsResult
	err := ctx.Invoke("alicloud:hbr/getVaults:getVaults", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVaults.
type GetVaultsArgs struct {
	// A list of Vault IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Vault name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of Vault. Valid values: `CREATED`, `ERROR`, `UNKNOWN`.
	Status *string `pulumi:"status"`
	// The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
	VaultType *string `pulumi:"vaultType"`
}

// A collection of values returned by getVaults.
type GetVaultsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string           `pulumi:"id"`
	Ids        []string         `pulumi:"ids"`
	NameRegex  *string          `pulumi:"nameRegex"`
	Names      []string         `pulumi:"names"`
	OutputFile *string          `pulumi:"outputFile"`
	Status     *string          `pulumi:"status"`
	VaultType  *string          `pulumi:"vaultType"`
	Vaults     []GetVaultsVault `pulumi:"vaults"`
}

func GetVaultsOutput(ctx *pulumi.Context, args GetVaultsOutputArgs, opts ...pulumi.InvokeOption) GetVaultsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVaultsResult, error) {
			args := v.(GetVaultsArgs)
			r, err := GetVaults(ctx, &args, opts...)
			var s GetVaultsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVaultsResultOutput)
}

// A collection of arguments for invoking getVaults.
type GetVaultsOutputArgs struct {
	// A list of Vault IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Vault name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of Vault. Valid values: `CREATED`, `ERROR`, `UNKNOWN`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
	VaultType pulumi.StringPtrInput `pulumi:"vaultType"`
}

func (GetVaultsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVaultsArgs)(nil)).Elem()
}

// A collection of values returned by getVaults.
type GetVaultsResultOutput struct{ *pulumi.OutputState }

func (GetVaultsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVaultsResult)(nil)).Elem()
}

func (o GetVaultsResultOutput) ToGetVaultsResultOutput() GetVaultsResultOutput {
	return o
}

func (o GetVaultsResultOutput) ToGetVaultsResultOutputWithContext(ctx context.Context) GetVaultsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVaultsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVaultsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVaultsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVaultsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVaultsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVaultsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetVaultsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVaultsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetVaultsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVaultsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetVaultsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVaultsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetVaultsResultOutput) VaultType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVaultsResult) *string { return v.VaultType }).(pulumi.StringPtrOutput)
}

func (o GetVaultsResultOutput) Vaults() GetVaultsVaultArrayOutput {
	return o.ApplyT(func(v GetVaultsResult) []GetVaultsVault { return v.Vaults }).(GetVaultsVaultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVaultsResultOutput{})
}
