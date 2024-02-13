// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Ipv4 Cidr Block resource. VPC IPv4 additional network segment.
//
// For information about VPC Ipv4 Cidr Block and how to use it, see [What is Ipv4 Cidr Block](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/associatevpccidrblock).
//
// > **NOTE:** Available since v1.185.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultvpc, err := vpc.NewNetwork(ctx, "defaultvpc", &vpc.NetworkArgs{
//				Description: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewIpv4CidrBlock(ctx, "default", &vpc.Ipv4CidrBlockArgs{
//				SecondaryCidrBlock: pulumi.String("192.168.0.0/16"),
//				VpcId:              defaultvpc.ID(),
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
// VPC Ipv4 Cidr Block can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/ipv4CidrBlock:Ipv4CidrBlock example <vpc_id>:<secondary_cidr_block>
// ```
type Ipv4CidrBlock struct {
	pulumi.CustomResourceState

	// The IPv4 CIDR block. Take note of the following requirements:
	// * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
	// * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
	// * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
	// * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
	SecondaryCidrBlock pulumi.StringOutput `pulumi:"secondaryCidrBlock"`
	// The ID of the VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewIpv4CidrBlock registers a new resource with the given unique name, arguments, and options.
func NewIpv4CidrBlock(ctx *pulumi.Context,
	name string, args *Ipv4CidrBlockArgs, opts ...pulumi.ResourceOption) (*Ipv4CidrBlock, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecondaryCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'SecondaryCidrBlock'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ipv4CidrBlock
	err := ctx.RegisterResource("alicloud:vpc/ipv4CidrBlock:Ipv4CidrBlock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpv4CidrBlock gets an existing Ipv4CidrBlock resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpv4CidrBlock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipv4CidrBlockState, opts ...pulumi.ResourceOption) (*Ipv4CidrBlock, error) {
	var resource Ipv4CidrBlock
	err := ctx.ReadResource("alicloud:vpc/ipv4CidrBlock:Ipv4CidrBlock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipv4CidrBlock resources.
type ipv4CidrBlockState struct {
	// The IPv4 CIDR block. Take note of the following requirements:
	// * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
	// * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
	// * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
	// * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
	SecondaryCidrBlock *string `pulumi:"secondaryCidrBlock"`
	// The ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

type Ipv4CidrBlockState struct {
	// The IPv4 CIDR block. Take note of the following requirements:
	// * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
	// * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
	// * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
	// * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
	SecondaryCidrBlock pulumi.StringPtrInput
	// The ID of the VPC.
	VpcId pulumi.StringPtrInput
}

func (Ipv4CidrBlockState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv4CidrBlockState)(nil)).Elem()
}

type ipv4CidrBlockArgs struct {
	// The IPv4 CIDR block. Take note of the following requirements:
	// * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
	// * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
	// * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
	// * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
	SecondaryCidrBlock string `pulumi:"secondaryCidrBlock"`
	// The ID of the VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Ipv4CidrBlock resource.
type Ipv4CidrBlockArgs struct {
	// The IPv4 CIDR block. Take note of the following requirements:
	// * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
	// * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
	// * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
	// * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
	SecondaryCidrBlock pulumi.StringInput
	// The ID of the VPC.
	VpcId pulumi.StringInput
}

func (Ipv4CidrBlockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv4CidrBlockArgs)(nil)).Elem()
}

type Ipv4CidrBlockInput interface {
	pulumi.Input

	ToIpv4CidrBlockOutput() Ipv4CidrBlockOutput
	ToIpv4CidrBlockOutputWithContext(ctx context.Context) Ipv4CidrBlockOutput
}

func (*Ipv4CidrBlock) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv4CidrBlock)(nil)).Elem()
}

func (i *Ipv4CidrBlock) ToIpv4CidrBlockOutput() Ipv4CidrBlockOutput {
	return i.ToIpv4CidrBlockOutputWithContext(context.Background())
}

func (i *Ipv4CidrBlock) ToIpv4CidrBlockOutputWithContext(ctx context.Context) Ipv4CidrBlockOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4CidrBlockOutput)
}

// Ipv4CidrBlockArrayInput is an input type that accepts Ipv4CidrBlockArray and Ipv4CidrBlockArrayOutput values.
// You can construct a concrete instance of `Ipv4CidrBlockArrayInput` via:
//
//	Ipv4CidrBlockArray{ Ipv4CidrBlockArgs{...} }
type Ipv4CidrBlockArrayInput interface {
	pulumi.Input

	ToIpv4CidrBlockArrayOutput() Ipv4CidrBlockArrayOutput
	ToIpv4CidrBlockArrayOutputWithContext(context.Context) Ipv4CidrBlockArrayOutput
}

type Ipv4CidrBlockArray []Ipv4CidrBlockInput

func (Ipv4CidrBlockArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv4CidrBlock)(nil)).Elem()
}

func (i Ipv4CidrBlockArray) ToIpv4CidrBlockArrayOutput() Ipv4CidrBlockArrayOutput {
	return i.ToIpv4CidrBlockArrayOutputWithContext(context.Background())
}

func (i Ipv4CidrBlockArray) ToIpv4CidrBlockArrayOutputWithContext(ctx context.Context) Ipv4CidrBlockArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4CidrBlockArrayOutput)
}

// Ipv4CidrBlockMapInput is an input type that accepts Ipv4CidrBlockMap and Ipv4CidrBlockMapOutput values.
// You can construct a concrete instance of `Ipv4CidrBlockMapInput` via:
//
//	Ipv4CidrBlockMap{ "key": Ipv4CidrBlockArgs{...} }
type Ipv4CidrBlockMapInput interface {
	pulumi.Input

	ToIpv4CidrBlockMapOutput() Ipv4CidrBlockMapOutput
	ToIpv4CidrBlockMapOutputWithContext(context.Context) Ipv4CidrBlockMapOutput
}

type Ipv4CidrBlockMap map[string]Ipv4CidrBlockInput

func (Ipv4CidrBlockMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv4CidrBlock)(nil)).Elem()
}

func (i Ipv4CidrBlockMap) ToIpv4CidrBlockMapOutput() Ipv4CidrBlockMapOutput {
	return i.ToIpv4CidrBlockMapOutputWithContext(context.Background())
}

func (i Ipv4CidrBlockMap) ToIpv4CidrBlockMapOutputWithContext(ctx context.Context) Ipv4CidrBlockMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4CidrBlockMapOutput)
}

type Ipv4CidrBlockOutput struct{ *pulumi.OutputState }

func (Ipv4CidrBlockOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv4CidrBlock)(nil)).Elem()
}

func (o Ipv4CidrBlockOutput) ToIpv4CidrBlockOutput() Ipv4CidrBlockOutput {
	return o
}

func (o Ipv4CidrBlockOutput) ToIpv4CidrBlockOutputWithContext(ctx context.Context) Ipv4CidrBlockOutput {
	return o
}

// The IPv4 CIDR block. Take note of the following requirements:
// * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
// * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
// * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
// * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
func (o Ipv4CidrBlockOutput) SecondaryCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4CidrBlock) pulumi.StringOutput { return v.SecondaryCidrBlock }).(pulumi.StringOutput)
}

// The ID of the VPC.
func (o Ipv4CidrBlockOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4CidrBlock) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type Ipv4CidrBlockArrayOutput struct{ *pulumi.OutputState }

func (Ipv4CidrBlockArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv4CidrBlock)(nil)).Elem()
}

func (o Ipv4CidrBlockArrayOutput) ToIpv4CidrBlockArrayOutput() Ipv4CidrBlockArrayOutput {
	return o
}

func (o Ipv4CidrBlockArrayOutput) ToIpv4CidrBlockArrayOutputWithContext(ctx context.Context) Ipv4CidrBlockArrayOutput {
	return o
}

func (o Ipv4CidrBlockArrayOutput) Index(i pulumi.IntInput) Ipv4CidrBlockOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipv4CidrBlock {
		return vs[0].([]*Ipv4CidrBlock)[vs[1].(int)]
	}).(Ipv4CidrBlockOutput)
}

type Ipv4CidrBlockMapOutput struct{ *pulumi.OutputState }

func (Ipv4CidrBlockMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv4CidrBlock)(nil)).Elem()
}

func (o Ipv4CidrBlockMapOutput) ToIpv4CidrBlockMapOutput() Ipv4CidrBlockMapOutput {
	return o
}

func (o Ipv4CidrBlockMapOutput) ToIpv4CidrBlockMapOutputWithContext(ctx context.Context) Ipv4CidrBlockMapOutput {
	return o
}

func (o Ipv4CidrBlockMapOutput) MapIndex(k pulumi.StringInput) Ipv4CidrBlockOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipv4CidrBlock {
		return vs[0].(map[string]*Ipv4CidrBlock)[vs[1].(string)]
	}).(Ipv4CidrBlockOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4CidrBlockInput)(nil)).Elem(), &Ipv4CidrBlock{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4CidrBlockArrayInput)(nil)).Elem(), Ipv4CidrBlockArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4CidrBlockMapInput)(nil)).Elem(), Ipv4CidrBlockMap{})
	pulumi.RegisterOutputType(Ipv4CidrBlockOutput{})
	pulumi.RegisterOutputType(Ipv4CidrBlockArrayOutput{})
	pulumi.RegisterOutputType(Ipv4CidrBlockMapOutput{})
}
