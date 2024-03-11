// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfirewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cloud Firewall Control Policies of the current Alibaba Cloud user.
//
// > **NOTE:** Available since v1.129.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudfirewall"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfirewall.GetControlPolicies(ctx, &cloudfirewall.GetControlPoliciesArgs{
//				Direction: "in",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetControlPolicies(ctx *pulumi.Context, args *GetControlPoliciesArgs, opts ...pulumi.InvokeOption) (*GetControlPoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetControlPoliciesResult
	err := ctx.Invoke("alicloud:cloudfirewall/getControlPolicies:getControlPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getControlPolicies.
type GetControlPoliciesArgs struct {
	// The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
	AclAction *string `pulumi:"aclAction"`
	// The unique ID of the access control policy.
	AclUuid *string `pulumi:"aclUuid"`
	// The description of the access control policy.
	Description *string `pulumi:"description"`
	// The destination address defined in the access control policy.
	Destination *string `pulumi:"destination"`
	// The direction of the traffic to which the access control policy applies. Valid values: `in`, `out`.
	Direction string `pulumi:"direction"`
	// The IP version of the address in the access control policy.
	IpVersion *string `pulumi:"ipVersion"`
	// The language of the content within the response. Valid values: `en`, `zh`.
	Lang *string `pulumi:"lang"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The type of the protocol in the access control policy. Valid values: If `direction` is  `in`, the valid value is `ANY`. If `direction` is `out`, the valid values are `ANY`, `TCP`, `UDP`, `ICMP`.
	Proto *string `pulumi:"proto"`
	// The source address in the access control policy.
	Source *string `pulumi:"source"`
}

// A collection of values returned by getControlPolicies.
type GetControlPoliciesResult struct {
	// The action that Cloud Firewall performs on the traffic.
	AclAction *string `pulumi:"aclAction"`
	// The unique ID of the access control policy.
	AclUuid *string `pulumi:"aclUuid"`
	// The description of the access control policy.
	Description *string `pulumi:"description"`
	// The destination address in the access control policy.
	Destination *string `pulumi:"destination"`
	// The direction of the traffic to which the access control policy applies.
	Direction string `pulumi:"direction"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Control Policy IDs.
	Ids        []string `pulumi:"ids"`
	IpVersion  *string  `pulumi:"ipVersion"`
	Lang       *string  `pulumi:"lang"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of Cloud Firewall Control Policies. Each element contains the following attributes:
	Policies []GetControlPoliciesPolicy `pulumi:"policies"`
	// The type of the protocol in the access control policy.
	Proto *string `pulumi:"proto"`
	// The source address in the access control policy.
	Source *string `pulumi:"source"`
}

func GetControlPoliciesOutput(ctx *pulumi.Context, args GetControlPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetControlPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetControlPoliciesResult, error) {
			args := v.(GetControlPoliciesArgs)
			r, err := GetControlPolicies(ctx, &args, opts...)
			var s GetControlPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetControlPoliciesResultOutput)
}

// A collection of arguments for invoking getControlPolicies.
type GetControlPoliciesOutputArgs struct {
	// The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
	AclAction pulumi.StringPtrInput `pulumi:"aclAction"`
	// The unique ID of the access control policy.
	AclUuid pulumi.StringPtrInput `pulumi:"aclUuid"`
	// The description of the access control policy.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The destination address defined in the access control policy.
	Destination pulumi.StringPtrInput `pulumi:"destination"`
	// The direction of the traffic to which the access control policy applies. Valid values: `in`, `out`.
	Direction pulumi.StringInput `pulumi:"direction"`
	// The IP version of the address in the access control policy.
	IpVersion pulumi.StringPtrInput `pulumi:"ipVersion"`
	// The language of the content within the response. Valid values: `en`, `zh`.
	Lang pulumi.StringPtrInput `pulumi:"lang"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The type of the protocol in the access control policy. Valid values: If `direction` is  `in`, the valid value is `ANY`. If `direction` is `out`, the valid values are `ANY`, `TCP`, `UDP`, `ICMP`.
	Proto pulumi.StringPtrInput `pulumi:"proto"`
	// The source address in the access control policy.
	Source pulumi.StringPtrInput `pulumi:"source"`
}

func (GetControlPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetControlPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getControlPolicies.
type GetControlPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetControlPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetControlPoliciesResult)(nil)).Elem()
}

func (o GetControlPoliciesResultOutput) ToGetControlPoliciesResultOutput() GetControlPoliciesResultOutput {
	return o
}

func (o GetControlPoliciesResultOutput) ToGetControlPoliciesResultOutputWithContext(ctx context.Context) GetControlPoliciesResultOutput {
	return o
}

// The action that Cloud Firewall performs on the traffic.
func (o GetControlPoliciesResultOutput) AclAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.AclAction }).(pulumi.StringPtrOutput)
}

// The unique ID of the access control policy.
func (o GetControlPoliciesResultOutput) AclUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.AclUuid }).(pulumi.StringPtrOutput)
}

// The description of the access control policy.
func (o GetControlPoliciesResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination address in the access control policy.
func (o GetControlPoliciesResultOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

// The direction of the traffic to which the access control policy applies.
func (o GetControlPoliciesResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) string { return v.Direction }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetControlPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Control Policy IDs.
func (o GetControlPoliciesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetControlPoliciesResultOutput) IpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.IpVersion }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of Cloud Firewall Control Policies. Each element contains the following attributes:
func (o GetControlPoliciesResultOutput) Policies() GetControlPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) []GetControlPoliciesPolicy { return v.Policies }).(GetControlPoliciesPolicyArrayOutput)
}

// The type of the protocol in the access control policy.
func (o GetControlPoliciesResultOutput) Proto() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Proto }).(pulumi.StringPtrOutput)
}

// The source address in the access control policy.
func (o GetControlPoliciesResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetControlPoliciesResultOutput{})
}
