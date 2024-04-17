// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sddp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Data Security Center Instance resource.
//
// For information about Data Security Center Instance and how to use it, see [What is Instance](https://help.aliyun.com/product/88674.html).
//
// > **NOTE:** Available in v1.136.0+.
//
// > **NOTE:** The Data Security Center Instance is not support in the international site.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/sddp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sddp.NewInstance(ctx, "default", &sddp.InstanceArgs{
//				PaymentType: pulumi.String("Subscription"),
//				SddpVersion: pulumi.String("version_company"),
//				SdCbool:     pulumi.String("yes"),
//				Period:      pulumi.Int(1),
//				Sdc:         pulumi.String("3"),
//				UdCbool:     pulumi.String("yes"),
//				Udc:         pulumi.String("2000"),
//				Dataphin:    pulumi.String("yes"),
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
// Data Security Center Instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:sddp/instance:Instance example <id>
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Whether the required RAM authorization is configured.
	Authed pulumi.BoolOutput `pulumi:"authed"`
	// The dataphin. Valid values: `yes`,`no`.
	Dataphin pulumi.StringPtrOutput `pulumi:"dataphin"`
	// The dataphin count. Valid values: 1 to 20.
	DataphinCount pulumi.StringPtrOutput `pulumi:"dataphinCount"`
	// The number of instances.
	InstanceNum pulumi.StringOutput `pulumi:"instanceNum"`
	// The logistics.
	Logistics pulumi.StringPtrOutput `pulumi:"logistics"`
	// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute a update operation.
	ModifyType pulumi.StringPtrOutput `pulumi:"modifyType"`
	// Whether the authorized MaxCompute (ODPS) assets.
	OdpsSet pulumi.BoolOutput `pulumi:"odpsSet"`
	// Whether the authorized oss assets.
	OssBucketSet pulumi.BoolOutput `pulumi:"ossBucketSet"`
	// The OSS storage capacity.
	OssSize pulumi.StringOutput `pulumi:"ossSize"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The Prepaid period. Valid values: `1`, `2`, `3`, `6`,`12`,`24`.
	Period pulumi.IntOutput `pulumi:"period"`
	// Whether the authorized rds assets.
	RdsSet pulumi.BoolOutput `pulumi:"rdsSet"`
	// The remaining days of the protection period of the assets in the current login account.
	RemainDays pulumi.StringOutput `pulumi:"remainDays"`
	// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`,
	RenewPeriod pulumi.IntPtrOutput `pulumi:"renewPeriod"`
	// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
	RenewalStatus pulumi.StringOutput `pulumi:"renewalStatus"`
	// Whether to use the database. Valid values:`yes`,`no`.
	SdCbool pulumi.StringOutput `pulumi:"sdCbool"`
	// The number of instances.
	Sdc pulumi.StringOutput `pulumi:"sdc"`
	// The sddp version. Valid values: `versionAudit`,`versionCompany`,`versionDlp`.
	SddpVersion pulumi.StringOutput `pulumi:"sddpVersion"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Whether to use OSS. Valid values: `yes`,`no`.
	UdCbool pulumi.StringOutput `pulumi:"udCbool"`
	// OSS Size.
	Udc pulumi.StringOutput `pulumi:"udc"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.Period == nil {
		return nil, errors.New("invalid value for required argument 'Period'")
	}
	if args.SdCbool == nil {
		return nil, errors.New("invalid value for required argument 'SdCbool'")
	}
	if args.Sdc == nil {
		return nil, errors.New("invalid value for required argument 'Sdc'")
	}
	if args.SddpVersion == nil {
		return nil, errors.New("invalid value for required argument 'SddpVersion'")
	}
	if args.UdCbool == nil {
		return nil, errors.New("invalid value for required argument 'UdCbool'")
	}
	if args.Udc == nil {
		return nil, errors.New("invalid value for required argument 'Udc'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("alicloud:sddp/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:sddp/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Whether the required RAM authorization is configured.
	Authed *bool `pulumi:"authed"`
	// The dataphin. Valid values: `yes`,`no`.
	Dataphin *string `pulumi:"dataphin"`
	// The dataphin count. Valid values: 1 to 20.
	DataphinCount *string `pulumi:"dataphinCount"`
	// The number of instances.
	InstanceNum *string `pulumi:"instanceNum"`
	// The logistics.
	Logistics *string `pulumi:"logistics"`
	// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute a update operation.
	ModifyType *string `pulumi:"modifyType"`
	// Whether the authorized MaxCompute (ODPS) assets.
	OdpsSet *bool `pulumi:"odpsSet"`
	// Whether the authorized oss assets.
	OssBucketSet *bool `pulumi:"ossBucketSet"`
	// The OSS storage capacity.
	OssSize *string `pulumi:"ossSize"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The Prepaid period. Valid values: `1`, `2`, `3`, `6`,`12`,`24`.
	Period *int `pulumi:"period"`
	// Whether the authorized rds assets.
	RdsSet *bool `pulumi:"rdsSet"`
	// The remaining days of the protection period of the assets in the current login account.
	RemainDays *string `pulumi:"remainDays"`
	// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`,
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// Whether to use the database. Valid values:`yes`,`no`.
	SdCbool *string `pulumi:"sdCbool"`
	// The number of instances.
	Sdc *string `pulumi:"sdc"`
	// The sddp version. Valid values: `versionAudit`,`versionCompany`,`versionDlp`.
	SddpVersion *string `pulumi:"sddpVersion"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// Whether to use OSS. Valid values: `yes`,`no`.
	UdCbool *string `pulumi:"udCbool"`
	// OSS Size.
	Udc *string `pulumi:"udc"`
}

type InstanceState struct {
	// Whether the required RAM authorization is configured.
	Authed pulumi.BoolPtrInput
	// The dataphin. Valid values: `yes`,`no`.
	Dataphin pulumi.StringPtrInput
	// The dataphin count. Valid values: 1 to 20.
	DataphinCount pulumi.StringPtrInput
	// The number of instances.
	InstanceNum pulumi.StringPtrInput
	// The logistics.
	Logistics pulumi.StringPtrInput
	// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute a update operation.
	ModifyType pulumi.StringPtrInput
	// Whether the authorized MaxCompute (ODPS) assets.
	OdpsSet pulumi.BoolPtrInput
	// Whether the authorized oss assets.
	OssBucketSet pulumi.BoolPtrInput
	// The OSS storage capacity.
	OssSize pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType pulumi.StringPtrInput
	// The Prepaid period. Valid values: `1`, `2`, `3`, `6`,`12`,`24`.
	Period pulumi.IntPtrInput
	// Whether the authorized rds assets.
	RdsSet pulumi.BoolPtrInput
	// The remaining days of the protection period of the assets in the current login account.
	RemainDays pulumi.StringPtrInput
	// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`,
	RenewPeriod pulumi.IntPtrInput
	// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
	RenewalStatus pulumi.StringPtrInput
	// Whether to use the database. Valid values:`yes`,`no`.
	SdCbool pulumi.StringPtrInput
	// The number of instances.
	Sdc pulumi.StringPtrInput
	// The sddp version. Valid values: `versionAudit`,`versionCompany`,`versionDlp`.
	SddpVersion pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// Whether to use OSS. Valid values: `yes`,`no`.
	UdCbool pulumi.StringPtrInput
	// OSS Size.
	Udc pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The dataphin. Valid values: `yes`,`no`.
	Dataphin *string `pulumi:"dataphin"`
	// The dataphin count. Valid values: 1 to 20.
	DataphinCount *string `pulumi:"dataphinCount"`
	// The logistics.
	Logistics *string `pulumi:"logistics"`
	// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute a update operation.
	ModifyType *string `pulumi:"modifyType"`
	// The OSS storage capacity.
	OssSize *string `pulumi:"ossSize"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType string `pulumi:"paymentType"`
	// The Prepaid period. Valid values: `1`, `2`, `3`, `6`,`12`,`24`.
	Period int `pulumi:"period"`
	// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`,
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// Whether to use the database. Valid values:`yes`,`no`.
	SdCbool string `pulumi:"sdCbool"`
	// The number of instances.
	Sdc string `pulumi:"sdc"`
	// The sddp version. Valid values: `versionAudit`,`versionCompany`,`versionDlp`.
	SddpVersion string `pulumi:"sddpVersion"`
	// Whether to use OSS. Valid values: `yes`,`no`.
	UdCbool string `pulumi:"udCbool"`
	// OSS Size.
	Udc string `pulumi:"udc"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The dataphin. Valid values: `yes`,`no`.
	Dataphin pulumi.StringPtrInput
	// The dataphin count. Valid values: 1 to 20.
	DataphinCount pulumi.StringPtrInput
	// The logistics.
	Logistics pulumi.StringPtrInput
	// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute a update operation.
	ModifyType pulumi.StringPtrInput
	// The OSS storage capacity.
	OssSize pulumi.StringPtrInput
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType pulumi.StringInput
	// The Prepaid period. Valid values: `1`, `2`, `3`, `6`,`12`,`24`.
	Period pulumi.IntInput
	// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`,
	RenewPeriod pulumi.IntPtrInput
	// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
	RenewalStatus pulumi.StringPtrInput
	// Whether to use the database. Valid values:`yes`,`no`.
	SdCbool pulumi.StringInput
	// The number of instances.
	Sdc pulumi.StringInput
	// The sddp version. Valid values: `versionAudit`,`versionCompany`,`versionDlp`.
	SddpVersion pulumi.StringInput
	// Whether to use OSS. Valid values: `yes`,`no`.
	UdCbool pulumi.StringInput
	// OSS Size.
	Udc pulumi.StringInput
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

// Whether the required RAM authorization is configured.
func (o InstanceOutput) Authed() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.Authed }).(pulumi.BoolOutput)
}

// The dataphin. Valid values: `yes`,`no`.
func (o InstanceOutput) Dataphin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Dataphin }).(pulumi.StringPtrOutput)
}

// The dataphin count. Valid values: 1 to 20.
func (o InstanceOutput) DataphinCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.DataphinCount }).(pulumi.StringPtrOutput)
}

// The number of instances.
func (o InstanceOutput) InstanceNum() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceNum }).(pulumi.StringOutput)
}

// The logistics.
func (o InstanceOutput) Logistics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Logistics }).(pulumi.StringPtrOutput)
}

// The modify type. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modifyType` is required when you execute a update operation.
func (o InstanceOutput) ModifyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ModifyType }).(pulumi.StringPtrOutput)
}

// Whether the authorized MaxCompute (ODPS) assets.
func (o InstanceOutput) OdpsSet() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.OdpsSet }).(pulumi.BoolOutput)
}

// Whether the authorized oss assets.
func (o InstanceOutput) OssBucketSet() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.OssBucketSet }).(pulumi.BoolOutput)
}

// The OSS storage capacity.
func (o InstanceOutput) OssSize() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.OssSize }).(pulumi.StringOutput)
}

// The payment type of the resource. Valid values: `Subscription`.
func (o InstanceOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The Prepaid period. Valid values: `1`, `2`, `3`, `6`,`12`,`24`.
func (o InstanceOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Period }).(pulumi.IntOutput)
}

// Whether the authorized rds assets.
func (o InstanceOutput) RdsSet() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.RdsSet }).(pulumi.BoolOutput)
}

// The remaining days of the protection period of the assets in the current login account.
func (o InstanceOutput) RemainDays() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.RemainDays }).(pulumi.StringOutput)
}

// Automatic renewal period. **NOTE:** The `renewPeriod` is required under the condition that renewalStatus is `AutoRenewal`,
func (o InstanceOutput) RenewPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.RenewPeriod }).(pulumi.IntPtrOutput)
}

// Automatic renewal status. Valid values: `AutoRenewal`,`ManualRenewal`. Default Value: `ManualRenewal`.
func (o InstanceOutput) RenewalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.RenewalStatus }).(pulumi.StringOutput)
}

// Whether to use the database. Valid values:`yes`,`no`.
func (o InstanceOutput) SdCbool() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SdCbool }).(pulumi.StringOutput)
}

// The number of instances.
func (o InstanceOutput) Sdc() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Sdc }).(pulumi.StringOutput)
}

// The sddp version. Valid values: `versionAudit`,`versionCompany`,`versionDlp`.
func (o InstanceOutput) SddpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SddpVersion }).(pulumi.StringOutput)
}

// The status of the resource.
func (o InstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Whether to use OSS. Valid values: `yes`,`no`.
func (o InstanceOutput) UdCbool() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.UdCbool }).(pulumi.StringOutput)
}

// OSS Size.
func (o InstanceOutput) Udc() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Udc }).(pulumi.StringOutput)
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
