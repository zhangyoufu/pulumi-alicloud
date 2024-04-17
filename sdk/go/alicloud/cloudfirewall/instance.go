// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfirewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Firewall Instance resource.
//
// For information about Cloud Firewall Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/product/90174.htm).
//
// > **NOTE:** Available since v1.139.0.
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
//			_, err := cloudfirewall.NewInstance(ctx, "default", &cloudfirewall.InstanceArgs{
//				PaymentType:   pulumi.String("PayAsYouGo"),
//				Spec:          pulumi.String("ultimate_version"),
//				IpNumber:      pulumi.Int(400),
//				BandWidth:     pulumi.Int(200),
//				CfwLog:        pulumi.Bool(true),
//				CfwLogStorage: pulumi.Int(1000),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Cloud Firewall Instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cloudfirewall/instance:Instance example <id>
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The number of multi account. It will be ignored when `cfwAccount = false`.
	AccountNumber pulumi.IntPtrOutput `pulumi:"accountNumber"`
	// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
	BandWidth pulumi.IntOutput `pulumi:"bandWidth"`
	// Whether to use multi-account. Valid values: `true`, `false`.
	CfwAccount pulumi.BoolPtrOutput `pulumi:"cfwAccount"`
	// Whether to use log audit. Valid values: `true`, `false`.
	CfwLog pulumi.BoolOutput `pulumi:"cfwLog"`
	// The log storage capacity. It will be ignored when `cfwLog = false`.
	CfwLogStorage pulumi.IntPtrOutput `pulumi:"cfwLogStorage"`
	// The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The end time.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The number of protected VPCs. It will be ignored when `spec = "premiumVersion"`. Valid values between 2 and 500.
	FwVpcNumber pulumi.IntPtrOutput `pulumi:"fwVpcNumber"`
	// The number of assets.
	InstanceCount pulumi.IntPtrOutput `pulumi:"instanceCount"`
	// The number of public IPs that can be protected. Valid values: 20 to 4000.
	IpNumber pulumi.IntOutput `pulumi:"ipNumber"`
	// The logistics.
	Logistics pulumi.StringPtrOutput `pulumi:"logistics"`
	// The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
	ModifyType pulumi.StringPtrOutput `pulumi:"modifyType"`
	// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `paymentType` can be set to `PayAsYouGo`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `paymentType` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The release time.
	ReleaseTime pulumi.StringOutput `pulumi:"releaseTime"`
	// Automatic renewal period. Attribute `renewPeriod` has been deprecated since 1.209.1. Using `renewalDuration` instead.
	//
	// Deprecated: Attribute 'renew_period' has been deprecated since 1.209.1. Using 'renewal_duration' instead.
	RenewPeriod pulumi.IntOutput `pulumi:"renewPeriod"`
	// Auto-Renewal Duration. It is required under the condition that `renewalStatus` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
	// **NOTE:** `renewalDuration` takes effect only if `paymentType` is set to `Subscription`, and `renewalStatus` is set to `AutoRenewal`.
	RenewalDuration pulumi.IntOutput `pulumi:"renewalDuration"`
	// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
	RenewalDurationUnit pulumi.StringPtrOutput `pulumi:"renewalDurationUnit"`
	// Whether to renew an instance automatically or not. Default to "ManualRenewal".
	RenewalStatus pulumi.StringOutput `pulumi:"renewalStatus"`
	// Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
	Spec pulumi.StringOutput `pulumi:"spec"`
	// The status of Instance.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BandWidth == nil {
		return nil, errors.New("invalid value for required argument 'BandWidth'")
	}
	if args.CfwLog == nil {
		return nil, errors.New("invalid value for required argument 'CfwLog'")
	}
	if args.IpNumber == nil {
		return nil, errors.New("invalid value for required argument 'IpNumber'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("alicloud:cloudfirewall/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:cloudfirewall/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The number of multi account. It will be ignored when `cfwAccount = false`.
	AccountNumber *int `pulumi:"accountNumber"`
	// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
	BandWidth *int `pulumi:"bandWidth"`
	// Whether to use multi-account. Valid values: `true`, `false`.
	CfwAccount *bool `pulumi:"cfwAccount"`
	// Whether to use log audit. Valid values: `true`, `false`.
	CfwLog *bool `pulumi:"cfwLog"`
	// The log storage capacity. It will be ignored when `cfwLog = false`.
	CfwLogStorage *int `pulumi:"cfwLogStorage"`
	// The creation time.
	CreateTime *string `pulumi:"createTime"`
	// The end time.
	EndTime *string `pulumi:"endTime"`
	// The number of protected VPCs. It will be ignored when `spec = "premiumVersion"`. Valid values between 2 and 500.
	FwVpcNumber *int `pulumi:"fwVpcNumber"`
	// The number of assets.
	InstanceCount *int `pulumi:"instanceCount"`
	// The number of public IPs that can be protected. Valid values: 20 to 4000.
	IpNumber *int `pulumi:"ipNumber"`
	// The logistics.
	Logistics *string `pulumi:"logistics"`
	// The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
	ModifyType *string `pulumi:"modifyType"`
	// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `paymentType` can be set to `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `paymentType` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
	Period *int `pulumi:"period"`
	// The release time.
	ReleaseTime *string `pulumi:"releaseTime"`
	// Automatic renewal period. Attribute `renewPeriod` has been deprecated since 1.209.1. Using `renewalDuration` instead.
	//
	// Deprecated: Attribute 'renew_period' has been deprecated since 1.209.1. Using 'renewal_duration' instead.
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Auto-Renewal Duration. It is required under the condition that `renewalStatus` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
	// **NOTE:** `renewalDuration` takes effect only if `paymentType` is set to `Subscription`, and `renewalStatus` is set to `AutoRenewal`.
	RenewalDuration *int `pulumi:"renewalDuration"`
	// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
	RenewalDurationUnit *string `pulumi:"renewalDurationUnit"`
	// Whether to renew an instance automatically or not. Default to "ManualRenewal".
	RenewalStatus *string `pulumi:"renewalStatus"`
	// Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
	Spec *string `pulumi:"spec"`
	// The status of Instance.
	Status *string `pulumi:"status"`
}

type InstanceState struct {
	// The number of multi account. It will be ignored when `cfwAccount = false`.
	AccountNumber pulumi.IntPtrInput
	// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
	BandWidth pulumi.IntPtrInput
	// Whether to use multi-account. Valid values: `true`, `false`.
	CfwAccount pulumi.BoolPtrInput
	// Whether to use log audit. Valid values: `true`, `false`.
	CfwLog pulumi.BoolPtrInput
	// The log storage capacity. It will be ignored when `cfwLog = false`.
	CfwLogStorage pulumi.IntPtrInput
	// The creation time.
	CreateTime pulumi.StringPtrInput
	// The end time.
	EndTime pulumi.StringPtrInput
	// The number of protected VPCs. It will be ignored when `spec = "premiumVersion"`. Valid values between 2 and 500.
	FwVpcNumber pulumi.IntPtrInput
	// The number of assets.
	InstanceCount pulumi.IntPtrInput
	// The number of public IPs that can be protected. Valid values: 20 to 4000.
	IpNumber pulumi.IntPtrInput
	// The logistics.
	Logistics pulumi.StringPtrInput
	// The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
	ModifyType pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `paymentType` can be set to `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `paymentType` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
	Period pulumi.IntPtrInput
	// The release time.
	ReleaseTime pulumi.StringPtrInput
	// Automatic renewal period. Attribute `renewPeriod` has been deprecated since 1.209.1. Using `renewalDuration` instead.
	//
	// Deprecated: Attribute 'renew_period' has been deprecated since 1.209.1. Using 'renewal_duration' instead.
	RenewPeriod pulumi.IntPtrInput
	// Auto-Renewal Duration. It is required under the condition that `renewalStatus` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
	// **NOTE:** `renewalDuration` takes effect only if `paymentType` is set to `Subscription`, and `renewalStatus` is set to `AutoRenewal`.
	RenewalDuration pulumi.IntPtrInput
	// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
	RenewalDurationUnit pulumi.StringPtrInput
	// Whether to renew an instance automatically or not. Default to "ManualRenewal".
	RenewalStatus pulumi.StringPtrInput
	// Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
	Spec pulumi.StringPtrInput
	// The status of Instance.
	Status pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The number of multi account. It will be ignored when `cfwAccount = false`.
	AccountNumber *int `pulumi:"accountNumber"`
	// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
	BandWidth int `pulumi:"bandWidth"`
	// Whether to use multi-account. Valid values: `true`, `false`.
	CfwAccount *bool `pulumi:"cfwAccount"`
	// Whether to use log audit. Valid values: `true`, `false`.
	CfwLog bool `pulumi:"cfwLog"`
	// The log storage capacity. It will be ignored when `cfwLog = false`.
	CfwLogStorage *int `pulumi:"cfwLogStorage"`
	// The number of protected VPCs. It will be ignored when `spec = "premiumVersion"`. Valid values between 2 and 500.
	FwVpcNumber *int `pulumi:"fwVpcNumber"`
	// The number of assets.
	InstanceCount *int `pulumi:"instanceCount"`
	// The number of public IPs that can be protected. Valid values: 20 to 4000.
	IpNumber int `pulumi:"ipNumber"`
	// The logistics.
	Logistics *string `pulumi:"logistics"`
	// The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
	ModifyType *string `pulumi:"modifyType"`
	// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `paymentType` can be set to `PayAsYouGo`.
	PaymentType string `pulumi:"paymentType"`
	// The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `paymentType` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
	Period *int `pulumi:"period"`
	// Automatic renewal period. Attribute `renewPeriod` has been deprecated since 1.209.1. Using `renewalDuration` instead.
	//
	// Deprecated: Attribute 'renew_period' has been deprecated since 1.209.1. Using 'renewal_duration' instead.
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Auto-Renewal Duration. It is required under the condition that `renewalStatus` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
	// **NOTE:** `renewalDuration` takes effect only if `paymentType` is set to `Subscription`, and `renewalStatus` is set to `AutoRenewal`.
	RenewalDuration *int `pulumi:"renewalDuration"`
	// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
	RenewalDurationUnit *string `pulumi:"renewalDurationUnit"`
	// Whether to renew an instance automatically or not. Default to "ManualRenewal".
	RenewalStatus *string `pulumi:"renewalStatus"`
	// Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
	Spec string `pulumi:"spec"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The number of multi account. It will be ignored when `cfwAccount = false`.
	AccountNumber pulumi.IntPtrInput
	// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
	BandWidth pulumi.IntInput
	// Whether to use multi-account. Valid values: `true`, `false`.
	CfwAccount pulumi.BoolPtrInput
	// Whether to use log audit. Valid values: `true`, `false`.
	CfwLog pulumi.BoolInput
	// The log storage capacity. It will be ignored when `cfwLog = false`.
	CfwLogStorage pulumi.IntPtrInput
	// The number of protected VPCs. It will be ignored when `spec = "premiumVersion"`. Valid values between 2 and 500.
	FwVpcNumber pulumi.IntPtrInput
	// The number of assets.
	InstanceCount pulumi.IntPtrInput
	// The number of public IPs that can be protected. Valid values: 20 to 4000.
	IpNumber pulumi.IntInput
	// The logistics.
	Logistics pulumi.StringPtrInput
	// The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
	ModifyType pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `paymentType` can be set to `PayAsYouGo`.
	PaymentType pulumi.StringInput
	// The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `paymentType` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
	Period pulumi.IntPtrInput
	// Automatic renewal period. Attribute `renewPeriod` has been deprecated since 1.209.1. Using `renewalDuration` instead.
	//
	// Deprecated: Attribute 'renew_period' has been deprecated since 1.209.1. Using 'renewal_duration' instead.
	RenewPeriod pulumi.IntPtrInput
	// Auto-Renewal Duration. It is required under the condition that `renewalStatus` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
	// **NOTE:** `renewalDuration` takes effect only if `paymentType` is set to `Subscription`, and `renewalStatus` is set to `AutoRenewal`.
	RenewalDuration pulumi.IntPtrInput
	// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
	RenewalDurationUnit pulumi.StringPtrInput
	// Whether to renew an instance automatically or not. Default to "ManualRenewal".
	RenewalStatus pulumi.StringPtrInput
	// Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
	Spec pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// The number of multi account. It will be ignored when `cfwAccount = false`.
func (o InstanceOutput) AccountNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.AccountNumber }).(pulumi.IntPtrOutput)
}

// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
func (o InstanceOutput) BandWidth() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.BandWidth }).(pulumi.IntOutput)
}

// Whether to use multi-account. Valid values: `true`, `false`.
func (o InstanceOutput) CfwAccount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.CfwAccount }).(pulumi.BoolPtrOutput)
}

// Whether to use log audit. Valid values: `true`, `false`.
func (o InstanceOutput) CfwLog() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.CfwLog }).(pulumi.BoolOutput)
}

// The log storage capacity. It will be ignored when `cfwLog = false`.
func (o InstanceOutput) CfwLogStorage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.CfwLogStorage }).(pulumi.IntPtrOutput)
}

// The creation time.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The end time.
func (o InstanceOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// The number of protected VPCs. It will be ignored when `spec = "premiumVersion"`. Valid values between 2 and 500.
func (o InstanceOutput) FwVpcNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.FwVpcNumber }).(pulumi.IntPtrOutput)
}

// The number of assets.
func (o InstanceOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

// The number of public IPs that can be protected. Valid values: 20 to 4000.
func (o InstanceOutput) IpNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.IpNumber }).(pulumi.IntOutput)
}

// The logistics.
func (o InstanceOutput) Logistics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Logistics }).(pulumi.StringPtrOutput)
}

// The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
func (o InstanceOutput) ModifyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ModifyType }).(pulumi.StringPtrOutput)
}

// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `paymentType` can be set to `PayAsYouGo`.
func (o InstanceOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `paymentType` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
func (o InstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The release time.
func (o InstanceOutput) ReleaseTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ReleaseTime }).(pulumi.StringOutput)
}

// Automatic renewal period. Attribute `renewPeriod` has been deprecated since 1.209.1. Using `renewalDuration` instead.
//
// Deprecated: Attribute 'renew_period' has been deprecated since 1.209.1. Using 'renewal_duration' instead.
func (o InstanceOutput) RenewPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.RenewPeriod }).(pulumi.IntOutput)
}

// Auto-Renewal Duration. It is required under the condition that `renewalStatus` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
// **NOTE:** `renewalDuration` takes effect only if `paymentType` is set to `Subscription`, and `renewalStatus` is set to `AutoRenewal`.
func (o InstanceOutput) RenewalDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.RenewalDuration }).(pulumi.IntOutput)
}

// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
func (o InstanceOutput) RenewalDurationUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.RenewalDurationUnit }).(pulumi.StringPtrOutput)
}

// Whether to renew an instance automatically or not. Default to "ManualRenewal".
func (o InstanceOutput) RenewalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.RenewalStatus }).(pulumi.StringOutput)
}

// Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
func (o InstanceOutput) Spec() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Spec }).(pulumi.StringOutput)
}

// The status of Instance.
func (o InstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
