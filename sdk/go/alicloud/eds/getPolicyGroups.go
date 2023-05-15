// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecd Policy Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.130.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := eds.NewEcdPolicyGroup(ctx, "default", &eds.EcdPolicyGroupArgs{
//				PolicyGroupName: pulumi.String("my-policy-group"),
//				Clipboard:       pulumi.String("read"),
//				LocalDrive:      pulumi.String("read"),
//				UsbRedirect:     pulumi.String("off"),
//				Watermark:       pulumi.String("off"),
//				AuthorizeAccessPolicyRules: eds.EcdPolicyGroupAuthorizeAccessPolicyRuleArray{
//					&eds.EcdPolicyGroupAuthorizeAccessPolicyRuleArgs{
//						Description: pulumi.String("my-description1"),
//						CidrIp:      pulumi.String("1.2.3.45/24"),
//					},
//				},
//				AuthorizeSecurityPolicyRules: eds.EcdPolicyGroupAuthorizeSecurityPolicyRuleArray{
//					&eds.EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs{
//						Type:        pulumi.String("inflow"),
//						Policy:      pulumi.String("accept"),
//						Description: pulumi.String("my-description"),
//						PortRange:   pulumi.String("80/80"),
//						IpProtocol:  pulumi.String("TCP"),
//						Priority:    pulumi.String("1"),
//						CidrIp:      pulumi.String("1.2.3.4/24"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			nameRegex, err := eds.GetPolicyGroups(ctx, &eds.GetPolicyGroupsArgs{
//				NameRegex: pulumi.StringRef("^my-policy"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecdPolicyGroupId", nameRegex.Groups[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetPolicyGroups(ctx *pulumi.Context, args *GetPolicyGroupsArgs, opts ...pulumi.InvokeOption) (*GetPolicyGroupsResult, error) {
	var rv GetPolicyGroupsResult
	err := ctx.Invoke("alicloud:eds/getPolicyGroups:getPolicyGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyGroups.
type GetPolicyGroupsArgs struct {
	// A list of Policy Group IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Policy Group name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of policy.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getPolicyGroups.
type GetPolicyGroupsResult struct {
	Groups []GetPolicyGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}

func GetPolicyGroupsOutput(ctx *pulumi.Context, args GetPolicyGroupsOutputArgs, opts ...pulumi.InvokeOption) GetPolicyGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicyGroupsResult, error) {
			args := v.(GetPolicyGroupsArgs)
			r, err := GetPolicyGroups(ctx, &args, opts...)
			var s GetPolicyGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPolicyGroupsResultOutput)
}

// A collection of arguments for invoking getPolicyGroups.
type GetPolicyGroupsOutputArgs struct {
	// A list of Policy Group IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Policy Group name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of policy.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetPolicyGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getPolicyGroups.
type GetPolicyGroupsResultOutput struct{ *pulumi.OutputState }

func (GetPolicyGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyGroupsResult)(nil)).Elem()
}

func (o GetPolicyGroupsResultOutput) ToGetPolicyGroupsResultOutput() GetPolicyGroupsResultOutput {
	return o
}

func (o GetPolicyGroupsResultOutput) ToGetPolicyGroupsResultOutputWithContext(ctx context.Context) GetPolicyGroupsResultOutput {
	return o
}

func (o GetPolicyGroupsResultOutput) Groups() GetPolicyGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) []GetPolicyGroupsGroup { return v.Groups }).(GetPolicyGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicyGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPolicyGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetPolicyGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetPolicyGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetPolicyGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetPolicyGroupsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicyGroupsResultOutput{})
}
