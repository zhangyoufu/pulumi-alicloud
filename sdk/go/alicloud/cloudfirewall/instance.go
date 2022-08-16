// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfirewall

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Firewall Instance resource.
//
// For information about Cloud Firewall Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/product/90174.htm).
//
// > **NOTE:** Available in v1.139.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudfirewall"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfirewall.NewInstance(ctx, "example", &cloudfirewall.InstanceArgs{
//				BandWidth:     pulumi.Int(10),
//				CfwLog:        pulumi.Bool(false),
//				CfwLogStorage: pulumi.Int(1000),
//				CfwService:    pulumi.Bool(false),
//				IpNumber:      pulumi.Int(20),
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(6),
//				Spec:          pulumi.String("premium_version"),
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
// Cloud Firewall Instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cloudfirewall/instance:Instance example <id>
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
	BandWidth pulumi.IntOutput `pulumi:"bandWidth"`
	// Whether to use log audit. Valid values: `true`, `false`.
	CfwLog pulumi.BoolOutput `pulumi:"cfwLog"`
	// The log storage capacity.
	CfwLogStorage pulumi.IntOutput `pulumi:"cfwLogStorage"`
	// Whether to use expert service. Valid values: `true`, `false`.
	CfwService pulumi.BoolOutput `pulumi:"cfwService"`
	// The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The end time.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The number of protected VPCs. Valid values between 2 and 500.
	FwVpcNumber pulumi.IntPtrOutput `pulumi:"fwVpcNumber"`
	// The number of assets.
	InstanceCount pulumi.IntPtrOutput `pulumi:"instanceCount"`
	// The number of public IPs that can be protected. Valid values: 20 to 4000.
	IpNumber pulumi.IntOutput `pulumi:"ipNumber"`
	// The logistics.
	Logistics pulumi.StringPtrOutput `pulumi:"logistics"`
	// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
	ModifyType pulumi.StringPtrOutput `pulumi:"modifyType"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The prepaid period. Valid values: `6`, `12`, `24`, `36`.
	Period pulumi.IntOutput `pulumi:"period"`
	// The release time.
	ReleaseTime pulumi.StringOutput `pulumi:"releaseTime"`
	// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`.
	RenewPeriod pulumi.IntPtrOutput `pulumi:"renewPeriod"`
	// Automatic renewal period unit. Valid values: `Month`,`Year`.
	RenewalDurationUnit pulumi.StringOutput `pulumi:"renewalDurationUnit"`
	// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
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
	if args.CfwLogStorage == nil {
		return nil, errors.New("invalid value for required argument 'CfwLogStorage'")
	}
	if args.CfwService == nil {
		return nil, errors.New("invalid value for required argument 'CfwService'")
	}
	if args.IpNumber == nil {
		return nil, errors.New("invalid value for required argument 'IpNumber'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.Period == nil {
		return nil, errors.New("invalid value for required argument 'Period'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
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
	// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
	BandWidth *int `pulumi:"bandWidth"`
	// Whether to use log audit. Valid values: `true`, `false`.
	CfwLog *bool `pulumi:"cfwLog"`
	// The log storage capacity.
	CfwLogStorage *int `pulumi:"cfwLogStorage"`
	// Whether to use expert service. Valid values: `true`, `false`.
	CfwService *bool `pulumi:"cfwService"`
	// The creation time.
	CreateTime *string `pulumi:"createTime"`
	// The end time.
	EndTime *string `pulumi:"endTime"`
	// The number of protected VPCs. Valid values between 2 and 500.
	FwVpcNumber *int `pulumi:"fwVpcNumber"`
	// The number of assets.
	InstanceCount *int `pulumi:"instanceCount"`
	// The number of public IPs that can be protected. Valid values: 20 to 4000.
	IpNumber *int `pulumi:"ipNumber"`
	// The logistics.
	Logistics *string `pulumi:"logistics"`
	// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
	ModifyType *string `pulumi:"modifyType"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The prepaid period. Valid values: `6`, `12`, `24`, `36`.
	Period *int `pulumi:"period"`
	// The release time.
	ReleaseTime *string `pulumi:"releaseTime"`
	// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`.
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Automatic renewal period unit. Valid values: `Month`,`Year`.
	RenewalDurationUnit *string `pulumi:"renewalDurationUnit"`
	// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
	Spec *string `pulumi:"spec"`
	// The status of Instance.
	Status *string `pulumi:"status"`
}

type InstanceState struct {
	// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
	BandWidth pulumi.IntPtrInput
	// Whether to use log audit. Valid values: `true`, `false`.
	CfwLog pulumi.BoolPtrInput
	// The log storage capacity.
	CfwLogStorage pulumi.IntPtrInput
	// Whether to use expert service. Valid values: `true`, `false`.
	CfwService pulumi.BoolPtrInput
	// The creation time.
	CreateTime pulumi.StringPtrInput
	// The end time.
	EndTime pulumi.StringPtrInput
	// The number of protected VPCs. Valid values between 2 and 500.
	FwVpcNumber pulumi.IntPtrInput
	// The number of assets.
	InstanceCount pulumi.IntPtrInput
	// The number of public IPs that can be protected. Valid values: 20 to 4000.
	IpNumber pulumi.IntPtrInput
	// The logistics.
	Logistics pulumi.StringPtrInput
	// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
	ModifyType pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType pulumi.StringPtrInput
	// The prepaid period. Valid values: `6`, `12`, `24`, `36`.
	Period pulumi.IntPtrInput
	// The release time.
	ReleaseTime pulumi.StringPtrInput
	// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`.
	RenewPeriod pulumi.IntPtrInput
	// Automatic renewal period unit. Valid values: `Month`,`Year`.
	RenewalDurationUnit pulumi.StringPtrInput
	// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
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
	// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
	BandWidth int `pulumi:"bandWidth"`
	// Whether to use log audit. Valid values: `true`, `false`.
	CfwLog bool `pulumi:"cfwLog"`
	// The log storage capacity.
	CfwLogStorage int `pulumi:"cfwLogStorage"`
	// Whether to use expert service. Valid values: `true`, `false`.
	CfwService bool `pulumi:"cfwService"`
	// The number of protected VPCs. Valid values between 2 and 500.
	FwVpcNumber *int `pulumi:"fwVpcNumber"`
	// The number of assets.
	InstanceCount *int `pulumi:"instanceCount"`
	// The number of public IPs that can be protected. Valid values: 20 to 4000.
	IpNumber int `pulumi:"ipNumber"`
	// The logistics.
	Logistics *string `pulumi:"logistics"`
	// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
	ModifyType *string `pulumi:"modifyType"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType string `pulumi:"paymentType"`
	// The prepaid period. Valid values: `6`, `12`, `24`, `36`.
	Period int `pulumi:"period"`
	// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`.
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// Current version. Valid values: `premiumVersion`, `enterpriseVersion`,`ultimateVersion`.
	Spec string `pulumi:"spec"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
	BandWidth pulumi.IntInput
	// Whether to use log audit. Valid values: `true`, `false`.
	CfwLog pulumi.BoolInput
	// The log storage capacity.
	CfwLogStorage pulumi.IntInput
	// Whether to use expert service. Valid values: `true`, `false`.
	CfwService pulumi.BoolInput
	// The number of protected VPCs. Valid values between 2 and 500.
	FwVpcNumber pulumi.IntPtrInput
	// The number of assets.
	InstanceCount pulumi.IntPtrInput
	// The number of public IPs that can be protected. Valid values: 20 to 4000.
	IpNumber pulumi.IntInput
	// The logistics.
	Logistics pulumi.StringPtrInput
	// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
	ModifyType pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType pulumi.StringInput
	// The prepaid period. Valid values: `6`, `12`, `24`, `36`.
	Period pulumi.IntInput
	// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`.
	RenewPeriod pulumi.IntPtrInput
	// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
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

// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
func (o InstanceOutput) BandWidth() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.BandWidth }).(pulumi.IntOutput)
}

// Whether to use log audit. Valid values: `true`, `false`.
func (o InstanceOutput) CfwLog() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.CfwLog }).(pulumi.BoolOutput)
}

// The log storage capacity.
func (o InstanceOutput) CfwLogStorage() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.CfwLogStorage }).(pulumi.IntOutput)
}

// Whether to use expert service. Valid values: `true`, `false`.
func (o InstanceOutput) CfwService() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.CfwService }).(pulumi.BoolOutput)
}

// The creation time.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The end time.
func (o InstanceOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// The number of protected VPCs. Valid values between 2 and 500.
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

// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute an update operation.
func (o InstanceOutput) ModifyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ModifyType }).(pulumi.StringPtrOutput)
}

// The payment type of the resource. Valid values: `Subscription`.
func (o InstanceOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The prepaid period. Valid values: `6`, `12`, `24`, `36`.
func (o InstanceOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Period }).(pulumi.IntOutput)
}

// The release time.
func (o InstanceOutput) ReleaseTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ReleaseTime }).(pulumi.StringOutput)
}

// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`.
func (o InstanceOutput) RenewPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.RenewPeriod }).(pulumi.IntPtrOutput)
}

// Automatic renewal period unit. Valid values: `Month`,`Year`.
func (o InstanceOutput) RenewalDurationUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.RenewalDurationUnit }).(pulumi.StringOutput)
}

// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
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
