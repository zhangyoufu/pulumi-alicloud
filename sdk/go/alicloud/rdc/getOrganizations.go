// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rdc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Rdc Organizations of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.137.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rdc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-testAccOrganizations-Organizations"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := rdc.NewOrganization(ctx, "default", &rdc.OrganizationArgs{
//				OrganizationName: pulumi.String(name),
//				Source:           pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			ids := rdc.GetOrganizationsOutput(ctx, rdc.GetOrganizationsOutputArgs{
//				Ids: pulumi.StringArray{
//					_default.ID(),
//				},
//			}, nil)
//			ctx.Export("rdcOrganizationId1", ids.ApplyT(func(ids rdc.GetOrganizationsResult) (string, error) {
//				return ids.Id, nil
//			}).(pulumi.StringOutput))
//			nameRegex, err := rdc.GetOrganizations(ctx, &rdc.GetOrganizationsArgs{
//				NameRegex: pulumi.StringRef("^my-Organization"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("rdcOrganizationId2", nameRegex.Id)
//			return nil
//		})
//	}
//
// ```
func GetOrganizations(ctx *pulumi.Context, args *GetOrganizationsArgs, opts ...pulumi.InvokeOption) (*GetOrganizationsResult, error) {
	var rv GetOrganizationsResult
	err := ctx.Invoke("alicloud:rdc/getOrganizations:getOrganizations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganizations.
type GetOrganizationsArgs struct {
	// A list of Organization IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Organization name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// User pk, not required, only required when the ak used by the calling interface is inconsistent with the user pk
	RealPk *string `pulumi:"realPk"`
}

// A collection of values returned by getOrganizations.
type GetOrganizationsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string                         `pulumi:"id"`
	Ids           []string                       `pulumi:"ids"`
	NameRegex     *string                        `pulumi:"nameRegex"`
	Names         []string                       `pulumi:"names"`
	Organizations []GetOrganizationsOrganization `pulumi:"organizations"`
	OutputFile    *string                        `pulumi:"outputFile"`
	RealPk        *string                        `pulumi:"realPk"`
}

func GetOrganizationsOutput(ctx *pulumi.Context, args GetOrganizationsOutputArgs, opts ...pulumi.InvokeOption) GetOrganizationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrganizationsResult, error) {
			args := v.(GetOrganizationsArgs)
			r, err := GetOrganizations(ctx, &args, opts...)
			var s GetOrganizationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrganizationsResultOutput)
}

// A collection of arguments for invoking getOrganizations.
type GetOrganizationsOutputArgs struct {
	// A list of Organization IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Organization name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// User pk, not required, only required when the ak used by the calling interface is inconsistent with the user pk
	RealPk pulumi.StringPtrInput `pulumi:"realPk"`
}

func (GetOrganizationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationsArgs)(nil)).Elem()
}

// A collection of values returned by getOrganizations.
type GetOrganizationsResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationsResult)(nil)).Elem()
}

func (o GetOrganizationsResultOutput) ToGetOrganizationsResultOutput() GetOrganizationsResultOutput {
	return o
}

func (o GetOrganizationsResultOutput) ToGetOrganizationsResultOutputWithContext(ctx context.Context) GetOrganizationsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrganizationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOrganizationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOrganizationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetOrganizationsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrganizationsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetOrganizationsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOrganizationsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetOrganizationsResultOutput) Organizations() GetOrganizationsOrganizationArrayOutput {
	return o.ApplyT(func(v GetOrganizationsResult) []GetOrganizationsOrganization { return v.Organizations }).(GetOrganizationsOrganizationArrayOutput)
}

func (o GetOrganizationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrganizationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetOrganizationsResultOutput) RealPk() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrganizationsResult) *string { return v.RealPk }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationsResultOutput{})
}
