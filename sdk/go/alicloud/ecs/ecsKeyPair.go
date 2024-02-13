// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECS Key Pair resource.
//
// For information about ECS Key Pair and how to use it, see [What is Key Pair](https://www.alibabacloud.com/help/en/doc-detail/51771.htm).
//
// > **NOTE:** Available in v1.121.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewEcsKeyPair(ctx, "example", &ecs.EcsKeyPairArgs{
//				KeyPairName: pulumi.String("key_pair_name"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewEcsKeyPair(ctx, "prefix", &ecs.EcsKeyPairArgs{
//				KeyNamePrefix: pulumi.String("terraform-test-key-pair-prefix"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewEcsKeyPair(ctx, "publickey", &ecs.EcsKeyPairArgs{
//				KeyPairName: pulumi.String("my_public_key"),
//				PublicKey:   pulumi.String("ssh-rsa AAAAB3Nza12345678qwertyuudsfsg"),
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
// ECS Key Pair can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ecs/ecsKeyPair:EcsKeyPair example <key_name>
// ```
type EcsKeyPair struct {
	pulumi.CustomResourceState

	// The finger print of the key pair.
	FingerPrint pulumi.StringOutput `pulumi:"fingerPrint"`
	// The key file.
	KeyFile pulumi.StringPtrOutput `pulumi:"keyFile"`
	// Field `keyName` has been deprecated from provider version 1.121.0. New field `keyPairName` instead.
	//
	// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
	KeyName       pulumi.StringOutput    `pulumi:"keyName"`
	KeyNamePrefix pulumi.StringPtrOutput `pulumi:"keyNamePrefix"`
	// The key pair's name. It is the only in one Alicloud account, the key pair's name. must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	KeyPairName pulumi.StringOutput `pulumi:"keyPairName"`
	// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// The Id of resource group which the key pair belongs.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	Tags            pulumi.MapOutput       `pulumi:"tags"`
}

// NewEcsKeyPair registers a new resource with the given unique name, arguments, and options.
func NewEcsKeyPair(ctx *pulumi.Context,
	name string, args *EcsKeyPairArgs, opts ...pulumi.ResourceOption) (*EcsKeyPair, error) {
	if args == nil {
		args = &EcsKeyPairArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EcsKeyPair
	err := ctx.RegisterResource("alicloud:ecs/ecsKeyPair:EcsKeyPair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcsKeyPair gets an existing EcsKeyPair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcsKeyPair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcsKeyPairState, opts ...pulumi.ResourceOption) (*EcsKeyPair, error) {
	var resource EcsKeyPair
	err := ctx.ReadResource("alicloud:ecs/ecsKeyPair:EcsKeyPair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcsKeyPair resources.
type ecsKeyPairState struct {
	// The finger print of the key pair.
	FingerPrint *string `pulumi:"fingerPrint"`
	// The key file.
	KeyFile *string `pulumi:"keyFile"`
	// Field `keyName` has been deprecated from provider version 1.121.0. New field `keyPairName` instead.
	//
	// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
	KeyName       *string `pulumi:"keyName"`
	KeyNamePrefix *string `pulumi:"keyNamePrefix"`
	// The key pair's name. It is the only in one Alicloud account, the key pair's name. must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	KeyPairName *string `pulumi:"keyPairName"`
	// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
	PublicKey *string `pulumi:"publicKey"`
	// The Id of resource group which the key pair belongs.
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	Tags            map[string]interface{} `pulumi:"tags"`
}

type EcsKeyPairState struct {
	// The finger print of the key pair.
	FingerPrint pulumi.StringPtrInput
	// The key file.
	KeyFile pulumi.StringPtrInput
	// Field `keyName` has been deprecated from provider version 1.121.0. New field `keyPairName` instead.
	//
	// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
	KeyName       pulumi.StringPtrInput
	KeyNamePrefix pulumi.StringPtrInput
	// The key pair's name. It is the only in one Alicloud account, the key pair's name. must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	KeyPairName pulumi.StringPtrInput
	// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
	PublicKey pulumi.StringPtrInput
	// The Id of resource group which the key pair belongs.
	ResourceGroupId pulumi.StringPtrInput
	Tags            pulumi.MapInput
}

func (EcsKeyPairState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsKeyPairState)(nil)).Elem()
}

type ecsKeyPairArgs struct {
	// The key file.
	KeyFile *string `pulumi:"keyFile"`
	// Field `keyName` has been deprecated from provider version 1.121.0. New field `keyPairName` instead.
	//
	// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
	KeyName       *string `pulumi:"keyName"`
	KeyNamePrefix *string `pulumi:"keyNamePrefix"`
	// The key pair's name. It is the only in one Alicloud account, the key pair's name. must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	KeyPairName *string `pulumi:"keyPairName"`
	// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
	PublicKey *string `pulumi:"publicKey"`
	// The Id of resource group which the key pair belongs.
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	Tags            map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a EcsKeyPair resource.
type EcsKeyPairArgs struct {
	// The key file.
	KeyFile pulumi.StringPtrInput
	// Field `keyName` has been deprecated from provider version 1.121.0. New field `keyPairName` instead.
	//
	// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
	KeyName       pulumi.StringPtrInput
	KeyNamePrefix pulumi.StringPtrInput
	// The key pair's name. It is the only in one Alicloud account, the key pair's name. must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	KeyPairName pulumi.StringPtrInput
	// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
	PublicKey pulumi.StringPtrInput
	// The Id of resource group which the key pair belongs.
	ResourceGroupId pulumi.StringPtrInput
	Tags            pulumi.MapInput
}

func (EcsKeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsKeyPairArgs)(nil)).Elem()
}

type EcsKeyPairInput interface {
	pulumi.Input

	ToEcsKeyPairOutput() EcsKeyPairOutput
	ToEcsKeyPairOutputWithContext(ctx context.Context) EcsKeyPairOutput
}

func (*EcsKeyPair) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsKeyPair)(nil)).Elem()
}

func (i *EcsKeyPair) ToEcsKeyPairOutput() EcsKeyPairOutput {
	return i.ToEcsKeyPairOutputWithContext(context.Background())
}

func (i *EcsKeyPair) ToEcsKeyPairOutputWithContext(ctx context.Context) EcsKeyPairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsKeyPairOutput)
}

// EcsKeyPairArrayInput is an input type that accepts EcsKeyPairArray and EcsKeyPairArrayOutput values.
// You can construct a concrete instance of `EcsKeyPairArrayInput` via:
//
//	EcsKeyPairArray{ EcsKeyPairArgs{...} }
type EcsKeyPairArrayInput interface {
	pulumi.Input

	ToEcsKeyPairArrayOutput() EcsKeyPairArrayOutput
	ToEcsKeyPairArrayOutputWithContext(context.Context) EcsKeyPairArrayOutput
}

type EcsKeyPairArray []EcsKeyPairInput

func (EcsKeyPairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsKeyPair)(nil)).Elem()
}

func (i EcsKeyPairArray) ToEcsKeyPairArrayOutput() EcsKeyPairArrayOutput {
	return i.ToEcsKeyPairArrayOutputWithContext(context.Background())
}

func (i EcsKeyPairArray) ToEcsKeyPairArrayOutputWithContext(ctx context.Context) EcsKeyPairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsKeyPairArrayOutput)
}

// EcsKeyPairMapInput is an input type that accepts EcsKeyPairMap and EcsKeyPairMapOutput values.
// You can construct a concrete instance of `EcsKeyPairMapInput` via:
//
//	EcsKeyPairMap{ "key": EcsKeyPairArgs{...} }
type EcsKeyPairMapInput interface {
	pulumi.Input

	ToEcsKeyPairMapOutput() EcsKeyPairMapOutput
	ToEcsKeyPairMapOutputWithContext(context.Context) EcsKeyPairMapOutput
}

type EcsKeyPairMap map[string]EcsKeyPairInput

func (EcsKeyPairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsKeyPair)(nil)).Elem()
}

func (i EcsKeyPairMap) ToEcsKeyPairMapOutput() EcsKeyPairMapOutput {
	return i.ToEcsKeyPairMapOutputWithContext(context.Background())
}

func (i EcsKeyPairMap) ToEcsKeyPairMapOutputWithContext(ctx context.Context) EcsKeyPairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsKeyPairMapOutput)
}

type EcsKeyPairOutput struct{ *pulumi.OutputState }

func (EcsKeyPairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsKeyPair)(nil)).Elem()
}

func (o EcsKeyPairOutput) ToEcsKeyPairOutput() EcsKeyPairOutput {
	return o
}

func (o EcsKeyPairOutput) ToEcsKeyPairOutputWithContext(ctx context.Context) EcsKeyPairOutput {
	return o
}

// The finger print of the key pair.
func (o EcsKeyPairOutput) FingerPrint() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsKeyPair) pulumi.StringOutput { return v.FingerPrint }).(pulumi.StringOutput)
}

// The key file.
func (o EcsKeyPairOutput) KeyFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsKeyPair) pulumi.StringPtrOutput { return v.KeyFile }).(pulumi.StringPtrOutput)
}

// Field `keyName` has been deprecated from provider version 1.121.0. New field `keyPairName` instead.
//
// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
func (o EcsKeyPairOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsKeyPair) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

func (o EcsKeyPairOutput) KeyNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsKeyPair) pulumi.StringPtrOutput { return v.KeyNamePrefix }).(pulumi.StringPtrOutput)
}

// The key pair's name. It is the only in one Alicloud account, the key pair's name. must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
func (o EcsKeyPairOutput) KeyPairName() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsKeyPair) pulumi.StringOutput { return v.KeyPairName }).(pulumi.StringOutput)
}

// You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resourceGroupId` is the key pair belongs.
func (o EcsKeyPairOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsKeyPair) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// The Id of resource group which the key pair belongs.
func (o EcsKeyPairOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsKeyPair) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o EcsKeyPairOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *EcsKeyPair) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type EcsKeyPairArrayOutput struct{ *pulumi.OutputState }

func (EcsKeyPairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsKeyPair)(nil)).Elem()
}

func (o EcsKeyPairArrayOutput) ToEcsKeyPairArrayOutput() EcsKeyPairArrayOutput {
	return o
}

func (o EcsKeyPairArrayOutput) ToEcsKeyPairArrayOutputWithContext(ctx context.Context) EcsKeyPairArrayOutput {
	return o
}

func (o EcsKeyPairArrayOutput) Index(i pulumi.IntInput) EcsKeyPairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EcsKeyPair {
		return vs[0].([]*EcsKeyPair)[vs[1].(int)]
	}).(EcsKeyPairOutput)
}

type EcsKeyPairMapOutput struct{ *pulumi.OutputState }

func (EcsKeyPairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsKeyPair)(nil)).Elem()
}

func (o EcsKeyPairMapOutput) ToEcsKeyPairMapOutput() EcsKeyPairMapOutput {
	return o
}

func (o EcsKeyPairMapOutput) ToEcsKeyPairMapOutputWithContext(ctx context.Context) EcsKeyPairMapOutput {
	return o
}

func (o EcsKeyPairMapOutput) MapIndex(k pulumi.StringInput) EcsKeyPairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EcsKeyPair {
		return vs[0].(map[string]*EcsKeyPair)[vs[1].(string)]
	}).(EcsKeyPairOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EcsKeyPairInput)(nil)).Elem(), &EcsKeyPair{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsKeyPairArrayInput)(nil)).Elem(), EcsKeyPairArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsKeyPairMapInput)(nil)).Elem(), EcsKeyPairMap{})
	pulumi.RegisterOutputType(EcsKeyPairOutput{})
	pulumi.RegisterOutputType(EcsKeyPairArrayOutput{})
	pulumi.RegisterOutputType(EcsKeyPairMapOutput{})
}
