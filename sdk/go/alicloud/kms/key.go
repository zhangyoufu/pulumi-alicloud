// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A kms key can help user to protect data security in the transmission process. For information about Alikms Key and how to use it, see [What is Resource Alikms Key](https://www.alibabacloud.com/help/doc-detail/28947.htm).
//
// > **NOTE:** Available in v1.85.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewKey(ctx, "key", &kms.KeyArgs{
// 			Description:         pulumi.String("Hello KMS"),
// 			PendingWindowInDays: pulumi.Int(7),
// 			Status:              pulumi.String("Enabled"),
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
// Alikms key can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:kms/key:Key example abc123456
// ```
type Key struct {
	pulumi.CustomResourceState

	// The Alicloud Resource Name (ARN) of the key.
	// * `creationDate` -The date and time when the CMK was created. The time is displayed in UTC.
	// * `creator` -The creator of the CMK.
	// * `deleteDate` -The scheduled date to delete CMK. The time is displayed in UTC. This value is returned only when the KeyState value is PendingDeletion.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies whether to enable automatic key rotation. Default:"Disabled".
	AutomaticRotation pulumi.StringPtrOutput `pulumi:"automaticRotation"`
	CreationDate      pulumi.StringOutput    `pulumi:"creationDate"`
	Creator           pulumi.StringOutput    `pulumi:"creator"`
	DeleteDate        pulumi.StringOutput    `pulumi:"deleteDate"`
	// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
	//
	// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
	DeletionWindowInDays pulumi.IntOutput `pulumi:"deletionWindowInDays"`
	// The description of the key as viewed in Alicloud console.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
	//
	// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// The type of the CMK.
	KeySpec pulumi.StringOutput `pulumi:"keySpec"`
	// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
	//
	// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
	KeyState pulumi.StringOutput `pulumi:"keyState"`
	// Specifies the usage of CMK. Currently, default to `ENCRYPT/DECRYPT`, indicating that CMK is used for encryption and decryption.
	KeyUsage pulumi.StringPtrOutput `pulumi:"keyUsage"`
	// The date and time the last rotation was performed. The time is displayed in UTC.
	LastRotationDate pulumi.StringOutput `pulumi:"lastRotationDate"`
	// The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.
	MaterialExpireTime pulumi.StringOutput `pulumi:"materialExpireTime"`
	// The time the next rotation is scheduled for execution.
	NextRotationDate pulumi.StringOutput `pulumi:"nextRotationDate"`
	// The source of the key material for the CMK. Defaults to "Aliyun_KMS".
	Origin pulumi.StringPtrOutput `pulumi:"origin"`
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	PendingWindowInDays pulumi.IntOutput `pulumi:"pendingWindowInDays"`
	// The ID of the current primary key version of the symmetric CMK.
	PrimaryKeyVersion pulumi.StringOutput `pulumi:"primaryKeyVersion"`
	// The protection level of the CMK. Defaults to "SOFTWARE".
	ProtectionLevel pulumi.StringPtrOutput `pulumi:"protectionLevel"`
	// The period of automatic key rotation. Unit: seconds.
	RotationInterval pulumi.StringPtrOutput `pulumi:"rotationInterval"`
	// The status of CMK. Defaults to Enabled. Valid Values: `Disabled`, `Enabled`, `PendingDeletion`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		args = &KeyArgs{}
	}

	var resource Key
	err := ctx.RegisterResource("alicloud:kms/key:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("alicloud:kms/key:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
	// The Alicloud Resource Name (ARN) of the key.
	// * `creationDate` -The date and time when the CMK was created. The time is displayed in UTC.
	// * `creator` -The creator of the CMK.
	// * `deleteDate` -The scheduled date to delete CMK. The time is displayed in UTC. This value is returned only when the KeyState value is PendingDeletion.
	Arn *string `pulumi:"arn"`
	// Specifies whether to enable automatic key rotation. Default:"Disabled".
	AutomaticRotation *string `pulumi:"automaticRotation"`
	CreationDate      *string `pulumi:"creationDate"`
	Creator           *string `pulumi:"creator"`
	DeleteDate        *string `pulumi:"deleteDate"`
	// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
	//
	// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// The description of the key as viewed in Alicloud console.
	Description *string `pulumi:"description"`
	// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
	//
	// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The type of the CMK.
	KeySpec *string `pulumi:"keySpec"`
	// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
	//
	// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
	KeyState *string `pulumi:"keyState"`
	// Specifies the usage of CMK. Currently, default to `ENCRYPT/DECRYPT`, indicating that CMK is used for encryption and decryption.
	KeyUsage *string `pulumi:"keyUsage"`
	// The date and time the last rotation was performed. The time is displayed in UTC.
	LastRotationDate *string `pulumi:"lastRotationDate"`
	// The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.
	MaterialExpireTime *string `pulumi:"materialExpireTime"`
	// The time the next rotation is scheduled for execution.
	NextRotationDate *string `pulumi:"nextRotationDate"`
	// The source of the key material for the CMK. Defaults to "Aliyun_KMS".
	Origin *string `pulumi:"origin"`
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	PendingWindowInDays *int `pulumi:"pendingWindowInDays"`
	// The ID of the current primary key version of the symmetric CMK.
	PrimaryKeyVersion *string `pulumi:"primaryKeyVersion"`
	// The protection level of the CMK. Defaults to "SOFTWARE".
	ProtectionLevel *string `pulumi:"protectionLevel"`
	// The period of automatic key rotation. Unit: seconds.
	RotationInterval *string `pulumi:"rotationInterval"`
	// The status of CMK. Defaults to Enabled. Valid Values: `Disabled`, `Enabled`, `PendingDeletion`.
	Status *string `pulumi:"status"`
}

type KeyState struct {
	// The Alicloud Resource Name (ARN) of the key.
	// * `creationDate` -The date and time when the CMK was created. The time is displayed in UTC.
	// * `creator` -The creator of the CMK.
	// * `deleteDate` -The scheduled date to delete CMK. The time is displayed in UTC. This value is returned only when the KeyState value is PendingDeletion.
	Arn pulumi.StringPtrInput
	// Specifies whether to enable automatic key rotation. Default:"Disabled".
	AutomaticRotation pulumi.StringPtrInput
	CreationDate      pulumi.StringPtrInput
	Creator           pulumi.StringPtrInput
	DeleteDate        pulumi.StringPtrInput
	// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
	//
	// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
	DeletionWindowInDays pulumi.IntPtrInput
	// The description of the key as viewed in Alicloud console.
	Description pulumi.StringPtrInput
	// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
	//
	// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
	IsEnabled pulumi.BoolPtrInput
	// The type of the CMK.
	KeySpec pulumi.StringPtrInput
	// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
	//
	// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
	KeyState pulumi.StringPtrInput
	// Specifies the usage of CMK. Currently, default to `ENCRYPT/DECRYPT`, indicating that CMK is used for encryption and decryption.
	KeyUsage pulumi.StringPtrInput
	// The date and time the last rotation was performed. The time is displayed in UTC.
	LastRotationDate pulumi.StringPtrInput
	// The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.
	MaterialExpireTime pulumi.StringPtrInput
	// The time the next rotation is scheduled for execution.
	NextRotationDate pulumi.StringPtrInput
	// The source of the key material for the CMK. Defaults to "Aliyun_KMS".
	Origin pulumi.StringPtrInput
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	PendingWindowInDays pulumi.IntPtrInput
	// The ID of the current primary key version of the symmetric CMK.
	PrimaryKeyVersion pulumi.StringPtrInput
	// The protection level of the CMK. Defaults to "SOFTWARE".
	ProtectionLevel pulumi.StringPtrInput
	// The period of automatic key rotation. Unit: seconds.
	RotationInterval pulumi.StringPtrInput
	// The status of CMK. Defaults to Enabled. Valid Values: `Disabled`, `Enabled`, `PendingDeletion`.
	Status pulumi.StringPtrInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Specifies whether to enable automatic key rotation. Default:"Disabled".
	AutomaticRotation *string `pulumi:"automaticRotation"`
	// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
	//
	// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// The description of the key as viewed in Alicloud console.
	Description *string `pulumi:"description"`
	// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
	//
	// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The type of the CMK.
	KeySpec *string `pulumi:"keySpec"`
	// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
	//
	// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
	KeyState *string `pulumi:"keyState"`
	// Specifies the usage of CMK. Currently, default to `ENCRYPT/DECRYPT`, indicating that CMK is used for encryption and decryption.
	KeyUsage *string `pulumi:"keyUsage"`
	// The source of the key material for the CMK. Defaults to "Aliyun_KMS".
	Origin *string `pulumi:"origin"`
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	PendingWindowInDays *int `pulumi:"pendingWindowInDays"`
	// The protection level of the CMK. Defaults to "SOFTWARE".
	ProtectionLevel *string `pulumi:"protectionLevel"`
	// The period of automatic key rotation. Unit: seconds.
	RotationInterval *string `pulumi:"rotationInterval"`
	// The status of CMK. Defaults to Enabled. Valid Values: `Disabled`, `Enabled`, `PendingDeletion`.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Specifies whether to enable automatic key rotation. Default:"Disabled".
	AutomaticRotation pulumi.StringPtrInput
	// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
	//
	// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
	DeletionWindowInDays pulumi.IntPtrInput
	// The description of the key as viewed in Alicloud console.
	Description pulumi.StringPtrInput
	// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
	//
	// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
	IsEnabled pulumi.BoolPtrInput
	// The type of the CMK.
	KeySpec pulumi.StringPtrInput
	// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
	//
	// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
	KeyState pulumi.StringPtrInput
	// Specifies the usage of CMK. Currently, default to `ENCRYPT/DECRYPT`, indicating that CMK is used for encryption and decryption.
	KeyUsage pulumi.StringPtrInput
	// The source of the key material for the CMK. Defaults to "Aliyun_KMS".
	Origin pulumi.StringPtrInput
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	PendingWindowInDays pulumi.IntPtrInput
	// The protection level of the CMK. Defaults to "SOFTWARE".
	ProtectionLevel pulumi.StringPtrInput
	// The period of automatic key rotation. Unit: seconds.
	RotationInterval pulumi.StringPtrInput
	// The status of CMK. Defaults to Enabled. Valid Values: `Disabled`, `Enabled`, `PendingDeletion`.
	Status pulumi.StringPtrInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((*Key)(nil))
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

func (i *Key) ToKeyPtrOutput() KeyPtrOutput {
	return i.ToKeyPtrOutputWithContext(context.Background())
}

func (i *Key) ToKeyPtrOutputWithContext(ctx context.Context) KeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPtrOutput)
}

type KeyPtrInput interface {
	pulumi.Input

	ToKeyPtrOutput() KeyPtrOutput
	ToKeyPtrOutputWithContext(ctx context.Context) KeyPtrOutput
}

type keyPtrType KeyArgs

func (*keyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil))
}

func (i *keyPtrType) ToKeyPtrOutput() KeyPtrOutput {
	return i.ToKeyPtrOutputWithContext(context.Background())
}

func (i *keyPtrType) ToKeyPtrOutputWithContext(ctx context.Context) KeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPtrOutput)
}

// KeyArrayInput is an input type that accepts KeyArray and KeyArrayOutput values.
// You can construct a concrete instance of `KeyArrayInput` via:
//
//          KeyArray{ KeyArgs{...} }
type KeyArrayInput interface {
	pulumi.Input

	ToKeyArrayOutput() KeyArrayOutput
	ToKeyArrayOutputWithContext(context.Context) KeyArrayOutput
}

type KeyArray []KeyInput

func (KeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Key)(nil))
}

func (i KeyArray) ToKeyArrayOutput() KeyArrayOutput {
	return i.ToKeyArrayOutputWithContext(context.Background())
}

func (i KeyArray) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyArrayOutput)
}

// KeyMapInput is an input type that accepts KeyMap and KeyMapOutput values.
// You can construct a concrete instance of `KeyMapInput` via:
//
//          KeyMap{ "key": KeyArgs{...} }
type KeyMapInput interface {
	pulumi.Input

	ToKeyMapOutput() KeyMapOutput
	ToKeyMapOutputWithContext(context.Context) KeyMapOutput
}

type KeyMap map[string]KeyInput

func (KeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Key)(nil))
}

func (i KeyMap) ToKeyMapOutput() KeyMapOutput {
	return i.ToKeyMapOutputWithContext(context.Background())
}

func (i KeyMap) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyMapOutput)
}

type KeyOutput struct {
	*pulumi.OutputState
}

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Key)(nil))
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

func (o KeyOutput) ToKeyPtrOutput() KeyPtrOutput {
	return o.ToKeyPtrOutputWithContext(context.Background())
}

func (o KeyOutput) ToKeyPtrOutputWithContext(ctx context.Context) KeyPtrOutput {
	return o.ApplyT(func(v Key) *Key {
		return &v
	}).(KeyPtrOutput)
}

type KeyPtrOutput struct {
	*pulumi.OutputState
}

func (KeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil))
}

func (o KeyPtrOutput) ToKeyPtrOutput() KeyPtrOutput {
	return o
}

func (o KeyPtrOutput) ToKeyPtrOutputWithContext(ctx context.Context) KeyPtrOutput {
	return o
}

type KeyArrayOutput struct{ *pulumi.OutputState }

func (KeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Key)(nil))
}

func (o KeyArrayOutput) ToKeyArrayOutput() KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) Index(i pulumi.IntInput) KeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Key {
		return vs[0].([]Key)[vs[1].(int)]
	}).(KeyOutput)
}

type KeyMapOutput struct{ *pulumi.OutputState }

func (KeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Key)(nil))
}

func (o KeyMapOutput) ToKeyMapOutput() KeyMapOutput {
	return o
}

func (o KeyMapOutput) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return o
}

func (o KeyMapOutput) MapIndex(k pulumi.StringInput) KeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Key {
		return vs[0].(map[string]Key)[vs[1].(string)]
	}).(KeyOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyOutput{})
	pulumi.RegisterOutputType(KeyPtrOutput{})
	pulumi.RegisterOutputType(KeyArrayOutput{})
	pulumi.RegisterOutputType(KeyMapOutput{})
}
