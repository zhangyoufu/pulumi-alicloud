// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of DMS User Tenants in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in 1.161.0+
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Declare the data source
//			_default, err := dms.GetUserTenants(ctx, &dms.GetUserTenantsArgs{
//				Status: pulumi.StringRef("ACTIVE"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("tid", _default.Ids[0])
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetUserTenants(ctx *pulumi.Context, args *GetUserTenantsArgs, opts ...pulumi.InvokeOption) (*GetUserTenantsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserTenantsResult
	err := ctx.Invoke("alicloud:dms/getUserTenants:getUserTenants", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserTenants.
type GetUserTenantsArgs struct {
	// A list of DMS User Tenant IDs (TID).
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the user tenant.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getUserTenants.
type GetUserTenantsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of DMS User Tenant IDs (UID).
	Ids []string `pulumi:"ids"`
	// A list of DMS User Tenant names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The status of the user tenant.
	Status *string `pulumi:"status"`
	// A list of DMS User Tenants. Each element contains the following attributes:
	Tenants []GetUserTenantsTenant `pulumi:"tenants"`
}

func GetUserTenantsOutput(ctx *pulumi.Context, args GetUserTenantsOutputArgs, opts ...pulumi.InvokeOption) GetUserTenantsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserTenantsResult, error) {
			args := v.(GetUserTenantsArgs)
			r, err := GetUserTenants(ctx, &args, opts...)
			var s GetUserTenantsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserTenantsResultOutput)
}

// A collection of arguments for invoking getUserTenants.
type GetUserTenantsOutputArgs struct {
	// A list of DMS User Tenant IDs (TID).
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the user tenant.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetUserTenantsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserTenantsArgs)(nil)).Elem()
}

// A collection of values returned by getUserTenants.
type GetUserTenantsResultOutput struct{ *pulumi.OutputState }

func (GetUserTenantsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserTenantsResult)(nil)).Elem()
}

func (o GetUserTenantsResultOutput) ToGetUserTenantsResultOutput() GetUserTenantsResultOutput {
	return o
}

func (o GetUserTenantsResultOutput) ToGetUserTenantsResultOutputWithContext(ctx context.Context) GetUserTenantsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserTenantsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserTenantsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of DMS User Tenant IDs (UID).
func (o GetUserTenantsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUserTenantsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of DMS User Tenant names.
func (o GetUserTenantsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUserTenantsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetUserTenantsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserTenantsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The status of the user tenant.
func (o GetUserTenantsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserTenantsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A list of DMS User Tenants. Each element contains the following attributes:
func (o GetUserTenantsResultOutput) Tenants() GetUserTenantsTenantArrayOutput {
	return o.ApplyT(func(v GetUserTenantsResult) []GetUserTenantsTenant { return v.Tenants }).(GetUserTenantsTenantArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserTenantsResultOutput{})
}
