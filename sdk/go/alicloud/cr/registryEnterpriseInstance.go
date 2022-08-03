// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource will help you to manager Container Registry Enterprise Edition instances.
//
// For information about Container Registry Enterprise Edition instances and how to use it, see [Create a Instance](https://www.alibabacloud.com/help/en/doc-detail/208144.htm)
//
// > **NOTE:** Available in v1.124.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cr.NewRegistryEnterpriseInstance(ctx, "my-instance", &cr.RegistryEnterpriseInstanceArgs{
// 			InstanceName:  pulumi.String("test"),
// 			InstanceType:  pulumi.String("Advanced"),
// 			PaymentType:   pulumi.String("Subscription"),
// 			Period:        pulumi.Int(1),
// 			RenewPeriod:   pulumi.Int(1),
// 			RenewalStatus: pulumi.String("AutoRenewal"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Container Registry Enterprise Edition instance can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import alicloud:cr/registryEnterpriseInstance:RegistryEnterpriseInstance default cri-test
// ```
type RegistryEnterpriseInstance struct {
	pulumi.CustomResourceState

	// Time of Container Registry Enterprise Edition instance creation.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Name of your customized oss bucket. Use this bucket as instance storage if set.
	CustomOssBucket pulumi.StringPtrOutput `pulumi:"customOssBucket"`
	// Time of Container Registry Enterprise Edition instance expiration.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Name of Container Registry Enterprise Edition instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
	PaymentType pulumi.StringPtrOutput `pulumi:"paymentType"`
	// Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
	RenewPeriod pulumi.IntPtrOutput `pulumi:"renewPeriod"`
	// Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
	RenewalStatus pulumi.StringPtrOutput `pulumi:"renewalStatus"`
	// Status of Container Registry Enterprise Edition instance.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewRegistryEnterpriseInstance registers a new resource with the given unique name, arguments, and options.
func NewRegistryEnterpriseInstance(ctx *pulumi.Context,
	name string, args *RegistryEnterpriseInstanceArgs, opts ...pulumi.ResourceOption) (*RegistryEnterpriseInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	var resource RegistryEnterpriseInstance
	err := ctx.RegisterResource("alicloud:cr/registryEnterpriseInstance:RegistryEnterpriseInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryEnterpriseInstance gets an existing RegistryEnterpriseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryEnterpriseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryEnterpriseInstanceState, opts ...pulumi.ResourceOption) (*RegistryEnterpriseInstance, error) {
	var resource RegistryEnterpriseInstance
	err := ctx.ReadResource("alicloud:cr/registryEnterpriseInstance:RegistryEnterpriseInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryEnterpriseInstance resources.
type registryEnterpriseInstanceState struct {
	// Time of Container Registry Enterprise Edition instance creation.
	CreatedTime *string `pulumi:"createdTime"`
	// Name of your customized oss bucket. Use this bucket as instance storage if set.
	CustomOssBucket *string `pulumi:"customOssBucket"`
	// Time of Container Registry Enterprise Edition instance expiration.
	EndTime *string `pulumi:"endTime"`
	// Name of Container Registry Enterprise Edition instance.
	InstanceName *string `pulumi:"instanceName"`
	// Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
	InstanceType *string `pulumi:"instanceType"`
	// An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
	Password *string `pulumi:"password"`
	// Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
	Period *int `pulumi:"period"`
	// Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// Status of Container Registry Enterprise Edition instance.
	Status *string `pulumi:"status"`
}

type RegistryEnterpriseInstanceState struct {
	// Time of Container Registry Enterprise Edition instance creation.
	CreatedTime pulumi.StringPtrInput
	// Name of your customized oss bucket. Use this bucket as instance storage if set.
	CustomOssBucket pulumi.StringPtrInput
	// Time of Container Registry Enterprise Edition instance expiration.
	EndTime pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition instance.
	InstanceName pulumi.StringPtrInput
	// Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
	InstanceType pulumi.StringPtrInput
	// An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
	Password pulumi.StringPtrInput
	// Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
	PaymentType pulumi.StringPtrInput
	// Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
	Period pulumi.IntPtrInput
	// Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
	RenewPeriod pulumi.IntPtrInput
	// Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
	RenewalStatus pulumi.StringPtrInput
	// Status of Container Registry Enterprise Edition instance.
	Status pulumi.StringPtrInput
}

func (RegistryEnterpriseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnterpriseInstanceState)(nil)).Elem()
}

type registryEnterpriseInstanceArgs struct {
	// Name of your customized oss bucket. Use this bucket as instance storage if set.
	CustomOssBucket *string `pulumi:"customOssBucket"`
	// Name of Container Registry Enterprise Edition instance.
	InstanceName string `pulumi:"instanceName"`
	// Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
	InstanceType string `pulumi:"instanceType"`
	// An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
	Password *string `pulumi:"password"`
	// Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
	Period *int `pulumi:"period"`
	// Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
}

// The set of arguments for constructing a RegistryEnterpriseInstance resource.
type RegistryEnterpriseInstanceArgs struct {
	// Name of your customized oss bucket. Use this bucket as instance storage if set.
	CustomOssBucket pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition instance.
	InstanceName pulumi.StringInput
	// Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
	InstanceType pulumi.StringInput
	// An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
	Password pulumi.StringPtrInput
	// Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
	PaymentType pulumi.StringPtrInput
	// Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
	Period pulumi.IntPtrInput
	// Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
	RenewPeriod pulumi.IntPtrInput
	// Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
	RenewalStatus pulumi.StringPtrInput
}

func (RegistryEnterpriseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnterpriseInstanceArgs)(nil)).Elem()
}

type RegistryEnterpriseInstanceInput interface {
	pulumi.Input

	ToRegistryEnterpriseInstanceOutput() RegistryEnterpriseInstanceOutput
	ToRegistryEnterpriseInstanceOutputWithContext(ctx context.Context) RegistryEnterpriseInstanceOutput
}

func (*RegistryEnterpriseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnterpriseInstance)(nil)).Elem()
}

func (i *RegistryEnterpriseInstance) ToRegistryEnterpriseInstanceOutput() RegistryEnterpriseInstanceOutput {
	return i.ToRegistryEnterpriseInstanceOutputWithContext(context.Background())
}

func (i *RegistryEnterpriseInstance) ToRegistryEnterpriseInstanceOutputWithContext(ctx context.Context) RegistryEnterpriseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseInstanceOutput)
}

// RegistryEnterpriseInstanceArrayInput is an input type that accepts RegistryEnterpriseInstanceArray and RegistryEnterpriseInstanceArrayOutput values.
// You can construct a concrete instance of `RegistryEnterpriseInstanceArrayInput` via:
//
//          RegistryEnterpriseInstanceArray{ RegistryEnterpriseInstanceArgs{...} }
type RegistryEnterpriseInstanceArrayInput interface {
	pulumi.Input

	ToRegistryEnterpriseInstanceArrayOutput() RegistryEnterpriseInstanceArrayOutput
	ToRegistryEnterpriseInstanceArrayOutputWithContext(context.Context) RegistryEnterpriseInstanceArrayOutput
}

type RegistryEnterpriseInstanceArray []RegistryEnterpriseInstanceInput

func (RegistryEnterpriseInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryEnterpriseInstance)(nil)).Elem()
}

func (i RegistryEnterpriseInstanceArray) ToRegistryEnterpriseInstanceArrayOutput() RegistryEnterpriseInstanceArrayOutput {
	return i.ToRegistryEnterpriseInstanceArrayOutputWithContext(context.Background())
}

func (i RegistryEnterpriseInstanceArray) ToRegistryEnterpriseInstanceArrayOutputWithContext(ctx context.Context) RegistryEnterpriseInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseInstanceArrayOutput)
}

// RegistryEnterpriseInstanceMapInput is an input type that accepts RegistryEnterpriseInstanceMap and RegistryEnterpriseInstanceMapOutput values.
// You can construct a concrete instance of `RegistryEnterpriseInstanceMapInput` via:
//
//          RegistryEnterpriseInstanceMap{ "key": RegistryEnterpriseInstanceArgs{...} }
type RegistryEnterpriseInstanceMapInput interface {
	pulumi.Input

	ToRegistryEnterpriseInstanceMapOutput() RegistryEnterpriseInstanceMapOutput
	ToRegistryEnterpriseInstanceMapOutputWithContext(context.Context) RegistryEnterpriseInstanceMapOutput
}

type RegistryEnterpriseInstanceMap map[string]RegistryEnterpriseInstanceInput

func (RegistryEnterpriseInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryEnterpriseInstance)(nil)).Elem()
}

func (i RegistryEnterpriseInstanceMap) ToRegistryEnterpriseInstanceMapOutput() RegistryEnterpriseInstanceMapOutput {
	return i.ToRegistryEnterpriseInstanceMapOutputWithContext(context.Background())
}

func (i RegistryEnterpriseInstanceMap) ToRegistryEnterpriseInstanceMapOutputWithContext(ctx context.Context) RegistryEnterpriseInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseInstanceMapOutput)
}

type RegistryEnterpriseInstanceOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnterpriseInstance)(nil)).Elem()
}

func (o RegistryEnterpriseInstanceOutput) ToRegistryEnterpriseInstanceOutput() RegistryEnterpriseInstanceOutput {
	return o
}

func (o RegistryEnterpriseInstanceOutput) ToRegistryEnterpriseInstanceOutputWithContext(ctx context.Context) RegistryEnterpriseInstanceOutput {
	return o
}

// Time of Container Registry Enterprise Edition instance creation.
func (o RegistryEnterpriseInstanceOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Name of your customized oss bucket. Use this bucket as instance storage if set.
func (o RegistryEnterpriseInstanceOutput) CustomOssBucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.StringPtrOutput { return v.CustomOssBucket }).(pulumi.StringPtrOutput)
}

// Time of Container Registry Enterprise Edition instance expiration.
func (o RegistryEnterpriseInstanceOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// Name of Container Registry Enterprise Edition instance.
func (o RegistryEnterpriseInstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
func (o RegistryEnterpriseInstanceOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
func (o RegistryEnterpriseInstanceOutput) KmsEncryptedPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.StringPtrOutput { return v.KmsEncryptedPassword }).(pulumi.StringPtrOutput)
}

// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
func (o RegistryEnterpriseInstanceOutput) KmsEncryptionContext() pulumi.MapOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.MapOutput { return v.KmsEncryptionContext }).(pulumi.MapOutput)
}

// The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
func (o RegistryEnterpriseInstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
func (o RegistryEnterpriseInstanceOutput) PaymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.StringPtrOutput { return v.PaymentType }).(pulumi.StringPtrOutput)
}

// Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
func (o RegistryEnterpriseInstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
func (o RegistryEnterpriseInstanceOutput) RenewPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.IntPtrOutput { return v.RenewPeriod }).(pulumi.IntPtrOutput)
}

// Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
func (o RegistryEnterpriseInstanceOutput) RenewalStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.StringPtrOutput { return v.RenewalStatus }).(pulumi.StringPtrOutput)
}

// Status of Container Registry Enterprise Edition instance.
func (o RegistryEnterpriseInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type RegistryEnterpriseInstanceArrayOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryEnterpriseInstance)(nil)).Elem()
}

func (o RegistryEnterpriseInstanceArrayOutput) ToRegistryEnterpriseInstanceArrayOutput() RegistryEnterpriseInstanceArrayOutput {
	return o
}

func (o RegistryEnterpriseInstanceArrayOutput) ToRegistryEnterpriseInstanceArrayOutputWithContext(ctx context.Context) RegistryEnterpriseInstanceArrayOutput {
	return o
}

func (o RegistryEnterpriseInstanceArrayOutput) Index(i pulumi.IntInput) RegistryEnterpriseInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegistryEnterpriseInstance {
		return vs[0].([]*RegistryEnterpriseInstance)[vs[1].(int)]
	}).(RegistryEnterpriseInstanceOutput)
}

type RegistryEnterpriseInstanceMapOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryEnterpriseInstance)(nil)).Elem()
}

func (o RegistryEnterpriseInstanceMapOutput) ToRegistryEnterpriseInstanceMapOutput() RegistryEnterpriseInstanceMapOutput {
	return o
}

func (o RegistryEnterpriseInstanceMapOutput) ToRegistryEnterpriseInstanceMapOutputWithContext(ctx context.Context) RegistryEnterpriseInstanceMapOutput {
	return o
}

func (o RegistryEnterpriseInstanceMapOutput) MapIndex(k pulumi.StringInput) RegistryEnterpriseInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegistryEnterpriseInstance {
		return vs[0].(map[string]*RegistryEnterpriseInstance)[vs[1].(string)]
	}).(RegistryEnterpriseInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseInstanceInput)(nil)).Elem(), &RegistryEnterpriseInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseInstanceArrayInput)(nil)).Elem(), RegistryEnterpriseInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseInstanceMapInput)(nil)).Elem(), RegistryEnterpriseInstanceMap{})
	pulumi.RegisterOutputType(RegistryEnterpriseInstanceOutput{})
	pulumi.RegisterOutputType(RegistryEnterpriseInstanceArrayOutput{})
	pulumi.RegisterOutputType(RegistryEnterpriseInstanceMapOutput{})
}
