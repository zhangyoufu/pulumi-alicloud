// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ecs

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a key pair resource.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/key_pair.html.markdown.
type KeyPair struct {
	pulumi.CustomResourceState

	FingerPrint pulumi.StringOutput `pulumi:"fingerPrint"`
	// The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.
	KeyFile pulumi.StringPtrOutput `pulumi:"keyFile"`
	// The key pair's name. It is the only in one Alicloud account.
	KeyName       pulumi.StringOutput    `pulumi:"keyName"`
	KeyNamePrefix pulumi.StringPtrOutput `pulumi:"keyNamePrefix"`
	// You can import an existing public key and using Alicloud key pair to manage it.
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
	// You can import an existing public key and using Alicloud key pair to manage it.
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
	// You can import an existing public key and using Alicloud key pair to manage it.
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
	// You can import an existing public key and using Alicloud key pair to manage it.
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
	// You can import an existing public key and using Alicloud key pair to manage it.
	PublicKey pulumi.StringPtrInput
	// The Id of resource group which the key pair belongs.
	ResourceGroupId pulumi.StringPtrInput
	Tags            pulumi.MapInput
}

func (KeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairArgs)(nil)).Elem()
}
