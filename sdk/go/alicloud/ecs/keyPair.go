// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a key pair resource.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ecs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.NewKeyPair(ctx, "basic", &ecs.KeyPairArgs{
// 			KeyName: pulumi.String("terraform-test-key-pair"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecs.NewKeyPair(ctx, "prefix", &ecs.KeyPairArgs{
// 			KeyNamePrefix: pulumi.String("terraform-test-key-pair-prefix"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecs.NewKeyPair(ctx, "publickey", &ecs.KeyPairArgs{
// 			KeyName:   pulumi.String("my_public_key"),
// 			PublicKey: pulumi.String("ssh-rsa AAAAB3Nza12345678qwertyuudsfsg"),
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
// Key pair can be imported using the name, e.g.
//
// ```sh
//  $ pulumi import alicloud:ecs/keyPair:KeyPair example my_public_key
// ```
type KeyPair struct {
	pulumi.CustomResourceState

	FingerPrint pulumi.StringOutput `pulumi:"fingerPrint"`
	// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
	KeyFile pulumi.StringPtrOutput `pulumi:"keyFile"`
	// The key pair's name. It is the only in one Alicloud account.
	KeyName       pulumi.StringOutput    `pulumi:"keyName"`
	KeyNamePrefix pulumi.StringPtrOutput `pulumi:"keyNamePrefix"`
	// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// The Id of resource group which the key pair belongs.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	Tags            pulumi.MapOutput       `pulumi:"tags"`
}

// NewKeyPair registers a new resource with the given unique name, arguments, and options.
func NewKeyPair(ctx *pulumi.Context,
	name string, args *KeyPairArgs, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	if args == nil {
		args = &KeyPairArgs{}
	}

	var resource KeyPair
	err := ctx.RegisterResource("alicloud:ecs/keyPair:KeyPair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyPair gets an existing KeyPair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyPair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyPairState, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	var resource KeyPair
	err := ctx.ReadResource("alicloud:ecs/keyPair:KeyPair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyPair resources.
type keyPairState struct {
	FingerPrint *string `pulumi:"fingerPrint"`
	// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
	KeyFile *string `pulumi:"keyFile"`
	// The key pair's name. It is the only in one Alicloud account.
	KeyName       *string `pulumi:"keyName"`
	KeyNamePrefix *string `pulumi:"keyNamePrefix"`
	// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
	PublicKey *string `pulumi:"publicKey"`
	// The Id of resource group which the key pair belongs.
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	Tags            map[string]interface{} `pulumi:"tags"`
}

type KeyPairState struct {
	FingerPrint pulumi.StringPtrInput
	// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
	KeyFile pulumi.StringPtrInput
	// The key pair's name. It is the only in one Alicloud account.
	KeyName       pulumi.StringPtrInput
	KeyNamePrefix pulumi.StringPtrInput
	// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
	PublicKey pulumi.StringPtrInput
	// The Id of resource group which the key pair belongs.
	ResourceGroupId pulumi.StringPtrInput
	Tags            pulumi.MapInput
}

func (KeyPairState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairState)(nil)).Elem()
}

type keyPairArgs struct {
	// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
	KeyFile *string `pulumi:"keyFile"`
	// The key pair's name. It is the only in one Alicloud account.
	KeyName       *string `pulumi:"keyName"`
	KeyNamePrefix *string `pulumi:"keyNamePrefix"`
	// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
	PublicKey *string `pulumi:"publicKey"`
	// The Id of resource group which the key pair belongs.
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	Tags            map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a KeyPair resource.
type KeyPairArgs struct {
	// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
	KeyFile pulumi.StringPtrInput
	// The key pair's name. It is the only in one Alicloud account.
	KeyName       pulumi.StringPtrInput
	KeyNamePrefix pulumi.StringPtrInput
	// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
	PublicKey pulumi.StringPtrInput
	// The Id of resource group which the key pair belongs.
	ResourceGroupId pulumi.StringPtrInput
	Tags            pulumi.MapInput
}

func (KeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairArgs)(nil)).Elem()
}

type KeyPairInput interface {
	pulumi.Input

	ToKeyPairOutput() KeyPairOutput
	ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput
}

func (*KeyPair) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPair)(nil))
}

func (i *KeyPair) ToKeyPairOutput() KeyPairOutput {
	return i.ToKeyPairOutputWithContext(context.Background())
}

func (i *KeyPair) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairOutput)
}

func (i *KeyPair) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return i.ToKeyPairPtrOutputWithContext(context.Background())
}

func (i *KeyPair) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairPtrOutput)
}

type KeyPairPtrInput interface {
	pulumi.Input

	ToKeyPairPtrOutput() KeyPairPtrOutput
	ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput
}

type keyPairPtrType KeyPairArgs

func (*keyPairPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPair)(nil))
}

func (i *keyPairPtrType) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return i.ToKeyPairPtrOutputWithContext(context.Background())
}

func (i *keyPairPtrType) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairPtrOutput)
}

// KeyPairArrayInput is an input type that accepts KeyPairArray and KeyPairArrayOutput values.
// You can construct a concrete instance of `KeyPairArrayInput` via:
//
//          KeyPairArray{ KeyPairArgs{...} }
type KeyPairArrayInput interface {
	pulumi.Input

	ToKeyPairArrayOutput() KeyPairArrayOutput
	ToKeyPairArrayOutputWithContext(context.Context) KeyPairArrayOutput
}

type KeyPairArray []KeyPairInput

func (KeyPairArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*KeyPair)(nil))
}

func (i KeyPairArray) ToKeyPairArrayOutput() KeyPairArrayOutput {
	return i.ToKeyPairArrayOutputWithContext(context.Background())
}

func (i KeyPairArray) ToKeyPairArrayOutputWithContext(ctx context.Context) KeyPairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairArrayOutput)
}

// KeyPairMapInput is an input type that accepts KeyPairMap and KeyPairMapOutput values.
// You can construct a concrete instance of `KeyPairMapInput` via:
//
//          KeyPairMap{ "key": KeyPairArgs{...} }
type KeyPairMapInput interface {
	pulumi.Input

	ToKeyPairMapOutput() KeyPairMapOutput
	ToKeyPairMapOutputWithContext(context.Context) KeyPairMapOutput
}

type KeyPairMap map[string]KeyPairInput

func (KeyPairMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*KeyPair)(nil))
}

func (i KeyPairMap) ToKeyPairMapOutput() KeyPairMapOutput {
	return i.ToKeyPairMapOutputWithContext(context.Background())
}

func (i KeyPairMap) ToKeyPairMapOutputWithContext(ctx context.Context) KeyPairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairMapOutput)
}

type KeyPairOutput struct {
	*pulumi.OutputState
}

func (KeyPairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPair)(nil))
}

func (o KeyPairOutput) ToKeyPairOutput() KeyPairOutput {
	return o
}

func (o KeyPairOutput) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return o
}

func (o KeyPairOutput) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return o.ToKeyPairPtrOutputWithContext(context.Background())
}

func (o KeyPairOutput) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return o.ApplyT(func(v KeyPair) *KeyPair {
		return &v
	}).(KeyPairPtrOutput)
}

type KeyPairPtrOutput struct {
	*pulumi.OutputState
}

func (KeyPairPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPair)(nil))
}

func (o KeyPairPtrOutput) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return o
}

func (o KeyPairPtrOutput) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return o
}

type KeyPairArrayOutput struct{ *pulumi.OutputState }

func (KeyPairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyPair)(nil))
}

func (o KeyPairArrayOutput) ToKeyPairArrayOutput() KeyPairArrayOutput {
	return o
}

func (o KeyPairArrayOutput) ToKeyPairArrayOutputWithContext(ctx context.Context) KeyPairArrayOutput {
	return o
}

func (o KeyPairArrayOutput) Index(i pulumi.IntInput) KeyPairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyPair {
		return vs[0].([]KeyPair)[vs[1].(int)]
	}).(KeyPairOutput)
}

type KeyPairMapOutput struct{ *pulumi.OutputState }

func (KeyPairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KeyPair)(nil))
}

func (o KeyPairMapOutput) ToKeyPairMapOutput() KeyPairMapOutput {
	return o
}

func (o KeyPairMapOutput) ToKeyPairMapOutputWithContext(ctx context.Context) KeyPairMapOutput {
	return o
}

func (o KeyPairMapOutput) MapIndex(k pulumi.StringInput) KeyPairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KeyPair {
		return vs[0].(map[string]KeyPair)[vs[1].(string)]
	}).(KeyPairOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyPairOutput{})
	pulumi.RegisterOutputType(KeyPairPtrOutput{})
	pulumi.RegisterOutputType(KeyPairArrayOutput{})
	pulumi.RegisterOutputType(KeyPairMapOutput{})
}
