// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Threat Detection Anti Brute Force Rule resource.
//
// For information about Threat Detection Anti Brute Force Rule and how to use it, see [What is Anti Brute Force Rule](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createantibruteforcerule).
//
// > **NOTE:** Available since v1.195.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/threatdetection"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := threatdetection.NewAntiBruteForceRule(ctx, "default", &threatdetection.AntiBruteForceRuleArgs{
//				AntiBruteForceRuleName: pulumi.String("apispec_example"),
//				FailCount:              pulumi.Int(80),
//				ForbiddenTime:          pulumi.Int(360),
//				Span:                   pulumi.Int(10),
//				UuidLists: pulumi.StringArray{
//					pulumi.String("032b618f-b220-4a0d-bd37-fbdc6ef58b6a"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Threat Detection Anti Brute Force Rule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:threatdetection/antiBruteForceRule:AntiBruteForceRule example <id>
//
// ```
type AntiBruteForceRule struct {
	pulumi.CustomResourceState

	// The ID of the defense rule.
	AntiBruteForceRuleId pulumi.StringOutput `pulumi:"antiBruteForceRuleId"`
	// The name of the defense rule.
	AntiBruteForceRuleName pulumi.StringOutput `pulumi:"antiBruteForceRuleName"`
	// Specifies whether to set the defense rule as the default rule.
	DefaultRule pulumi.BoolOutput `pulumi:"defaultRule"`
	// The threshold for the number of failed user logins when the brute-force defense rule takes effect.
	FailCount pulumi.IntOutput `pulumi:"failCount"`
	// The period of time during which logons from an account are not allowed. Unit: minutes.
	ForbiddenTime pulumi.IntOutput `pulumi:"forbiddenTime"`
	// The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
	Span pulumi.IntOutput `pulumi:"span"`
	// An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
	UuidLists pulumi.StringArrayOutput `pulumi:"uuidLists"`
}

// NewAntiBruteForceRule registers a new resource with the given unique name, arguments, and options.
func NewAntiBruteForceRule(ctx *pulumi.Context,
	name string, args *AntiBruteForceRuleArgs, opts ...pulumi.ResourceOption) (*AntiBruteForceRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AntiBruteForceRuleName == nil {
		return nil, errors.New("invalid value for required argument 'AntiBruteForceRuleName'")
	}
	if args.FailCount == nil {
		return nil, errors.New("invalid value for required argument 'FailCount'")
	}
	if args.ForbiddenTime == nil {
		return nil, errors.New("invalid value for required argument 'ForbiddenTime'")
	}
	if args.Span == nil {
		return nil, errors.New("invalid value for required argument 'Span'")
	}
	if args.UuidLists == nil {
		return nil, errors.New("invalid value for required argument 'UuidLists'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AntiBruteForceRule
	err := ctx.RegisterResource("alicloud:threatdetection/antiBruteForceRule:AntiBruteForceRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAntiBruteForceRule gets an existing AntiBruteForceRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAntiBruteForceRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AntiBruteForceRuleState, opts ...pulumi.ResourceOption) (*AntiBruteForceRule, error) {
	var resource AntiBruteForceRule
	err := ctx.ReadResource("alicloud:threatdetection/antiBruteForceRule:AntiBruteForceRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AntiBruteForceRule resources.
type antiBruteForceRuleState struct {
	// The ID of the defense rule.
	AntiBruteForceRuleId *string `pulumi:"antiBruteForceRuleId"`
	// The name of the defense rule.
	AntiBruteForceRuleName *string `pulumi:"antiBruteForceRuleName"`
	// Specifies whether to set the defense rule as the default rule.
	DefaultRule *bool `pulumi:"defaultRule"`
	// The threshold for the number of failed user logins when the brute-force defense rule takes effect.
	FailCount *int `pulumi:"failCount"`
	// The period of time during which logons from an account are not allowed. Unit: minutes.
	ForbiddenTime *int `pulumi:"forbiddenTime"`
	// The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
	Span *int `pulumi:"span"`
	// An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
	UuidLists []string `pulumi:"uuidLists"`
}

type AntiBruteForceRuleState struct {
	// The ID of the defense rule.
	AntiBruteForceRuleId pulumi.StringPtrInput
	// The name of the defense rule.
	AntiBruteForceRuleName pulumi.StringPtrInput
	// Specifies whether to set the defense rule as the default rule.
	DefaultRule pulumi.BoolPtrInput
	// The threshold for the number of failed user logins when the brute-force defense rule takes effect.
	FailCount pulumi.IntPtrInput
	// The period of time during which logons from an account are not allowed. Unit: minutes.
	ForbiddenTime pulumi.IntPtrInput
	// The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
	Span pulumi.IntPtrInput
	// An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
	UuidLists pulumi.StringArrayInput
}

func (AntiBruteForceRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*antiBruteForceRuleState)(nil)).Elem()
}

type antiBruteForceRuleArgs struct {
	// The name of the defense rule.
	AntiBruteForceRuleName string `pulumi:"antiBruteForceRuleName"`
	// Specifies whether to set the defense rule as the default rule.
	DefaultRule *bool `pulumi:"defaultRule"`
	// The threshold for the number of failed user logins when the brute-force defense rule takes effect.
	FailCount int `pulumi:"failCount"`
	// The period of time during which logons from an account are not allowed. Unit: minutes.
	ForbiddenTime int `pulumi:"forbiddenTime"`
	// The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
	Span int `pulumi:"span"`
	// An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
	UuidLists []string `pulumi:"uuidLists"`
}

// The set of arguments for constructing a AntiBruteForceRule resource.
type AntiBruteForceRuleArgs struct {
	// The name of the defense rule.
	AntiBruteForceRuleName pulumi.StringInput
	// Specifies whether to set the defense rule as the default rule.
	DefaultRule pulumi.BoolPtrInput
	// The threshold for the number of failed user logins when the brute-force defense rule takes effect.
	FailCount pulumi.IntInput
	// The period of time during which logons from an account are not allowed. Unit: minutes.
	ForbiddenTime pulumi.IntInput
	// The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
	Span pulumi.IntInput
	// An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
	UuidLists pulumi.StringArrayInput
}

func (AntiBruteForceRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*antiBruteForceRuleArgs)(nil)).Elem()
}

type AntiBruteForceRuleInput interface {
	pulumi.Input

	ToAntiBruteForceRuleOutput() AntiBruteForceRuleOutput
	ToAntiBruteForceRuleOutputWithContext(ctx context.Context) AntiBruteForceRuleOutput
}

func (*AntiBruteForceRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AntiBruteForceRule)(nil)).Elem()
}

func (i *AntiBruteForceRule) ToAntiBruteForceRuleOutput() AntiBruteForceRuleOutput {
	return i.ToAntiBruteForceRuleOutputWithContext(context.Background())
}

func (i *AntiBruteForceRule) ToAntiBruteForceRuleOutputWithContext(ctx context.Context) AntiBruteForceRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AntiBruteForceRuleOutput)
}

// AntiBruteForceRuleArrayInput is an input type that accepts AntiBruteForceRuleArray and AntiBruteForceRuleArrayOutput values.
// You can construct a concrete instance of `AntiBruteForceRuleArrayInput` via:
//
//	AntiBruteForceRuleArray{ AntiBruteForceRuleArgs{...} }
type AntiBruteForceRuleArrayInput interface {
	pulumi.Input

	ToAntiBruteForceRuleArrayOutput() AntiBruteForceRuleArrayOutput
	ToAntiBruteForceRuleArrayOutputWithContext(context.Context) AntiBruteForceRuleArrayOutput
}

type AntiBruteForceRuleArray []AntiBruteForceRuleInput

func (AntiBruteForceRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AntiBruteForceRule)(nil)).Elem()
}

func (i AntiBruteForceRuleArray) ToAntiBruteForceRuleArrayOutput() AntiBruteForceRuleArrayOutput {
	return i.ToAntiBruteForceRuleArrayOutputWithContext(context.Background())
}

func (i AntiBruteForceRuleArray) ToAntiBruteForceRuleArrayOutputWithContext(ctx context.Context) AntiBruteForceRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AntiBruteForceRuleArrayOutput)
}

// AntiBruteForceRuleMapInput is an input type that accepts AntiBruteForceRuleMap and AntiBruteForceRuleMapOutput values.
// You can construct a concrete instance of `AntiBruteForceRuleMapInput` via:
//
//	AntiBruteForceRuleMap{ "key": AntiBruteForceRuleArgs{...} }
type AntiBruteForceRuleMapInput interface {
	pulumi.Input

	ToAntiBruteForceRuleMapOutput() AntiBruteForceRuleMapOutput
	ToAntiBruteForceRuleMapOutputWithContext(context.Context) AntiBruteForceRuleMapOutput
}

type AntiBruteForceRuleMap map[string]AntiBruteForceRuleInput

func (AntiBruteForceRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AntiBruteForceRule)(nil)).Elem()
}

func (i AntiBruteForceRuleMap) ToAntiBruteForceRuleMapOutput() AntiBruteForceRuleMapOutput {
	return i.ToAntiBruteForceRuleMapOutputWithContext(context.Background())
}

func (i AntiBruteForceRuleMap) ToAntiBruteForceRuleMapOutputWithContext(ctx context.Context) AntiBruteForceRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AntiBruteForceRuleMapOutput)
}

type AntiBruteForceRuleOutput struct{ *pulumi.OutputState }

func (AntiBruteForceRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AntiBruteForceRule)(nil)).Elem()
}

func (o AntiBruteForceRuleOutput) ToAntiBruteForceRuleOutput() AntiBruteForceRuleOutput {
	return o
}

func (o AntiBruteForceRuleOutput) ToAntiBruteForceRuleOutputWithContext(ctx context.Context) AntiBruteForceRuleOutput {
	return o
}

// The ID of the defense rule.
func (o AntiBruteForceRuleOutput) AntiBruteForceRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiBruteForceRule) pulumi.StringOutput { return v.AntiBruteForceRuleId }).(pulumi.StringOutput)
}

// The name of the defense rule.
func (o AntiBruteForceRuleOutput) AntiBruteForceRuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiBruteForceRule) pulumi.StringOutput { return v.AntiBruteForceRuleName }).(pulumi.StringOutput)
}

// Specifies whether to set the defense rule as the default rule.
func (o AntiBruteForceRuleOutput) DefaultRule() pulumi.BoolOutput {
	return o.ApplyT(func(v *AntiBruteForceRule) pulumi.BoolOutput { return v.DefaultRule }).(pulumi.BoolOutput)
}

// The threshold for the number of failed user logins when the brute-force defense rule takes effect.
func (o AntiBruteForceRuleOutput) FailCount() pulumi.IntOutput {
	return o.ApplyT(func(v *AntiBruteForceRule) pulumi.IntOutput { return v.FailCount }).(pulumi.IntOutput)
}

// The period of time during which logons from an account are not allowed. Unit: minutes.
func (o AntiBruteForceRuleOutput) ForbiddenTime() pulumi.IntOutput {
	return o.ApplyT(func(v *AntiBruteForceRule) pulumi.IntOutput { return v.ForbiddenTime }).(pulumi.IntOutput)
}

// The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
func (o AntiBruteForceRuleOutput) Span() pulumi.IntOutput {
	return o.ApplyT(func(v *AntiBruteForceRule) pulumi.IntOutput { return v.Span }).(pulumi.IntOutput)
}

// An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
func (o AntiBruteForceRuleOutput) UuidLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AntiBruteForceRule) pulumi.StringArrayOutput { return v.UuidLists }).(pulumi.StringArrayOutput)
}

type AntiBruteForceRuleArrayOutput struct{ *pulumi.OutputState }

func (AntiBruteForceRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AntiBruteForceRule)(nil)).Elem()
}

func (o AntiBruteForceRuleArrayOutput) ToAntiBruteForceRuleArrayOutput() AntiBruteForceRuleArrayOutput {
	return o
}

func (o AntiBruteForceRuleArrayOutput) ToAntiBruteForceRuleArrayOutputWithContext(ctx context.Context) AntiBruteForceRuleArrayOutput {
	return o
}

func (o AntiBruteForceRuleArrayOutput) Index(i pulumi.IntInput) AntiBruteForceRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AntiBruteForceRule {
		return vs[0].([]*AntiBruteForceRule)[vs[1].(int)]
	}).(AntiBruteForceRuleOutput)
}

type AntiBruteForceRuleMapOutput struct{ *pulumi.OutputState }

func (AntiBruteForceRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AntiBruteForceRule)(nil)).Elem()
}

func (o AntiBruteForceRuleMapOutput) ToAntiBruteForceRuleMapOutput() AntiBruteForceRuleMapOutput {
	return o
}

func (o AntiBruteForceRuleMapOutput) ToAntiBruteForceRuleMapOutputWithContext(ctx context.Context) AntiBruteForceRuleMapOutput {
	return o
}

func (o AntiBruteForceRuleMapOutput) MapIndex(k pulumi.StringInput) AntiBruteForceRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AntiBruteForceRule {
		return vs[0].(map[string]*AntiBruteForceRule)[vs[1].(string)]
	}).(AntiBruteForceRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AntiBruteForceRuleInput)(nil)).Elem(), &AntiBruteForceRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AntiBruteForceRuleArrayInput)(nil)).Elem(), AntiBruteForceRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AntiBruteForceRuleMapInput)(nil)).Elem(), AntiBruteForceRuleMap{})
	pulumi.RegisterOutputType(AntiBruteForceRuleOutput{})
	pulumi.RegisterOutputType(AntiBruteForceRuleArrayOutput{})
	pulumi.RegisterOutputType(AntiBruteForceRuleMapOutput{})
}
