// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Public Ip Address Pool Cidr Block resource.
// > **NOTE:** Only users who have the required permissions can use the IP address pool feature of Elastic IP Address (EIP). To apply for the required permissions, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket).
//
// For information about VPC Public Ip Address Pool Cidr Block and how to use it, see [What is Public Ip Address Pool Cidr Block](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/429100).
//
// > **NOTE:** Available since v1.189.0.
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
//			defaultPublicIpAddressPool, err := vpc.NewPublicIpAddressPool(ctx, "defaultPublicIpAddressPool", &vpc.PublicIpAddressPoolArgs{
//				PublicIpAddressPoolName: pulumi.String(name),
//				Description:             pulumi.String(name),
//				Isp:                     pulumi.String("BGP"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewPublicIpAddressPoolCidrBlock(ctx, "defaultPublicIpAddressPoolCidrBlock", &vpc.PublicIpAddressPoolCidrBlockArgs{
//				PublicIpAddressPoolId: defaultPublicIpAddressPool.ID(),
//				CidrBlock:             pulumi.String("47.118.126.0/25"),
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
// VPC Public Ip Address Pool Cidr Block can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/publicIpAddressPoolCidrBlock:PublicIpAddressPoolCidrBlock example <public_ip_address_pool_id>:<cidr_block>
//
// ```
type PublicIpAddressPoolCidrBlock struct {
	pulumi.CustomResourceState

	// The CIDR block.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The creation time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The ID of the VPC Public IP address pool.
	PublicIpAddressPoolId pulumi.StringOutput `pulumi:"publicIpAddressPoolId"`
	// The status of the VPC Public Ip Address Pool Cidr Block.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewPublicIpAddressPoolCidrBlock registers a new resource with the given unique name, arguments, and options.
func NewPublicIpAddressPoolCidrBlock(ctx *pulumi.Context,
	name string, args *PublicIpAddressPoolCidrBlockArgs, opts ...pulumi.ResourceOption) (*PublicIpAddressPoolCidrBlock, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicIpAddressPoolId == nil {
		return nil, errors.New("invalid value for required argument 'PublicIpAddressPoolId'")
	}
	var resource PublicIpAddressPoolCidrBlock
	err := ctx.RegisterResource("alicloud:vpc/publicIpAddressPoolCidrBlock:PublicIpAddressPoolCidrBlock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIpAddressPoolCidrBlock gets an existing PublicIpAddressPoolCidrBlock resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIpAddressPoolCidrBlock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIpAddressPoolCidrBlockState, opts ...pulumi.ResourceOption) (*PublicIpAddressPoolCidrBlock, error) {
	var resource PublicIpAddressPoolCidrBlock
	err := ctx.ReadResource("alicloud:vpc/publicIpAddressPoolCidrBlock:PublicIpAddressPoolCidrBlock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIpAddressPoolCidrBlock resources.
type publicIpAddressPoolCidrBlockState struct {
	// The CIDR block.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The creation time of the resource.
	CreateTime *string `pulumi:"createTime"`
	// The ID of the VPC Public IP address pool.
	PublicIpAddressPoolId *string `pulumi:"publicIpAddressPoolId"`
	// The status of the VPC Public Ip Address Pool Cidr Block.
	Status *string `pulumi:"status"`
}

type PublicIpAddressPoolCidrBlockState struct {
	// The CIDR block.
	CidrBlock pulumi.StringPtrInput
	// The creation time of the resource.
	CreateTime pulumi.StringPtrInput
	// The ID of the VPC Public IP address pool.
	PublicIpAddressPoolId pulumi.StringPtrInput
	// The status of the VPC Public Ip Address Pool Cidr Block.
	Status pulumi.StringPtrInput
}

func (PublicIpAddressPoolCidrBlockState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpAddressPoolCidrBlockState)(nil)).Elem()
}

type publicIpAddressPoolCidrBlockArgs struct {
	// The CIDR block.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The ID of the VPC Public IP address pool.
	PublicIpAddressPoolId string `pulumi:"publicIpAddressPoolId"`
}

// The set of arguments for constructing a PublicIpAddressPoolCidrBlock resource.
type PublicIpAddressPoolCidrBlockArgs struct {
	// The CIDR block.
	CidrBlock pulumi.StringPtrInput
	// The ID of the VPC Public IP address pool.
	PublicIpAddressPoolId pulumi.StringInput
}

func (PublicIpAddressPoolCidrBlockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpAddressPoolCidrBlockArgs)(nil)).Elem()
}

type PublicIpAddressPoolCidrBlockInput interface {
	pulumi.Input

	ToPublicIpAddressPoolCidrBlockOutput() PublicIpAddressPoolCidrBlockOutput
	ToPublicIpAddressPoolCidrBlockOutputWithContext(ctx context.Context) PublicIpAddressPoolCidrBlockOutput
}

func (*PublicIpAddressPoolCidrBlock) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpAddressPoolCidrBlock)(nil)).Elem()
}

func (i *PublicIpAddressPoolCidrBlock) ToPublicIpAddressPoolCidrBlockOutput() PublicIpAddressPoolCidrBlockOutput {
	return i.ToPublicIpAddressPoolCidrBlockOutputWithContext(context.Background())
}

func (i *PublicIpAddressPoolCidrBlock) ToPublicIpAddressPoolCidrBlockOutputWithContext(ctx context.Context) PublicIpAddressPoolCidrBlockOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpAddressPoolCidrBlockOutput)
}

// PublicIpAddressPoolCidrBlockArrayInput is an input type that accepts PublicIpAddressPoolCidrBlockArray and PublicIpAddressPoolCidrBlockArrayOutput values.
// You can construct a concrete instance of `PublicIpAddressPoolCidrBlockArrayInput` via:
//
//	PublicIpAddressPoolCidrBlockArray{ PublicIpAddressPoolCidrBlockArgs{...} }
type PublicIpAddressPoolCidrBlockArrayInput interface {
	pulumi.Input

	ToPublicIpAddressPoolCidrBlockArrayOutput() PublicIpAddressPoolCidrBlockArrayOutput
	ToPublicIpAddressPoolCidrBlockArrayOutputWithContext(context.Context) PublicIpAddressPoolCidrBlockArrayOutput
}

type PublicIpAddressPoolCidrBlockArray []PublicIpAddressPoolCidrBlockInput

func (PublicIpAddressPoolCidrBlockArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicIpAddressPoolCidrBlock)(nil)).Elem()
}

func (i PublicIpAddressPoolCidrBlockArray) ToPublicIpAddressPoolCidrBlockArrayOutput() PublicIpAddressPoolCidrBlockArrayOutput {
	return i.ToPublicIpAddressPoolCidrBlockArrayOutputWithContext(context.Background())
}

func (i PublicIpAddressPoolCidrBlockArray) ToPublicIpAddressPoolCidrBlockArrayOutputWithContext(ctx context.Context) PublicIpAddressPoolCidrBlockArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpAddressPoolCidrBlockArrayOutput)
}

// PublicIpAddressPoolCidrBlockMapInput is an input type that accepts PublicIpAddressPoolCidrBlockMap and PublicIpAddressPoolCidrBlockMapOutput values.
// You can construct a concrete instance of `PublicIpAddressPoolCidrBlockMapInput` via:
//
//	PublicIpAddressPoolCidrBlockMap{ "key": PublicIpAddressPoolCidrBlockArgs{...} }
type PublicIpAddressPoolCidrBlockMapInput interface {
	pulumi.Input

	ToPublicIpAddressPoolCidrBlockMapOutput() PublicIpAddressPoolCidrBlockMapOutput
	ToPublicIpAddressPoolCidrBlockMapOutputWithContext(context.Context) PublicIpAddressPoolCidrBlockMapOutput
}

type PublicIpAddressPoolCidrBlockMap map[string]PublicIpAddressPoolCidrBlockInput

func (PublicIpAddressPoolCidrBlockMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicIpAddressPoolCidrBlock)(nil)).Elem()
}

func (i PublicIpAddressPoolCidrBlockMap) ToPublicIpAddressPoolCidrBlockMapOutput() PublicIpAddressPoolCidrBlockMapOutput {
	return i.ToPublicIpAddressPoolCidrBlockMapOutputWithContext(context.Background())
}

func (i PublicIpAddressPoolCidrBlockMap) ToPublicIpAddressPoolCidrBlockMapOutputWithContext(ctx context.Context) PublicIpAddressPoolCidrBlockMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpAddressPoolCidrBlockMapOutput)
}

type PublicIpAddressPoolCidrBlockOutput struct{ *pulumi.OutputState }

func (PublicIpAddressPoolCidrBlockOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpAddressPoolCidrBlock)(nil)).Elem()
}

func (o PublicIpAddressPoolCidrBlockOutput) ToPublicIpAddressPoolCidrBlockOutput() PublicIpAddressPoolCidrBlockOutput {
	return o
}

func (o PublicIpAddressPoolCidrBlockOutput) ToPublicIpAddressPoolCidrBlockOutputWithContext(ctx context.Context) PublicIpAddressPoolCidrBlockOutput {
	return o
}

// The CIDR block.
func (o PublicIpAddressPoolCidrBlockOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIpAddressPoolCidrBlock) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// The creation time of the resource.
func (o PublicIpAddressPoolCidrBlockOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIpAddressPoolCidrBlock) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The ID of the VPC Public IP address pool.
func (o PublicIpAddressPoolCidrBlockOutput) PublicIpAddressPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIpAddressPoolCidrBlock) pulumi.StringOutput { return v.PublicIpAddressPoolId }).(pulumi.StringOutput)
}

// The status of the VPC Public Ip Address Pool Cidr Block.
func (o PublicIpAddressPoolCidrBlockOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIpAddressPoolCidrBlock) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type PublicIpAddressPoolCidrBlockArrayOutput struct{ *pulumi.OutputState }

func (PublicIpAddressPoolCidrBlockArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicIpAddressPoolCidrBlock)(nil)).Elem()
}

func (o PublicIpAddressPoolCidrBlockArrayOutput) ToPublicIpAddressPoolCidrBlockArrayOutput() PublicIpAddressPoolCidrBlockArrayOutput {
	return o
}

func (o PublicIpAddressPoolCidrBlockArrayOutput) ToPublicIpAddressPoolCidrBlockArrayOutputWithContext(ctx context.Context) PublicIpAddressPoolCidrBlockArrayOutput {
	return o
}

func (o PublicIpAddressPoolCidrBlockArrayOutput) Index(i pulumi.IntInput) PublicIpAddressPoolCidrBlockOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PublicIpAddressPoolCidrBlock {
		return vs[0].([]*PublicIpAddressPoolCidrBlock)[vs[1].(int)]
	}).(PublicIpAddressPoolCidrBlockOutput)
}

type PublicIpAddressPoolCidrBlockMapOutput struct{ *pulumi.OutputState }

func (PublicIpAddressPoolCidrBlockMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicIpAddressPoolCidrBlock)(nil)).Elem()
}

func (o PublicIpAddressPoolCidrBlockMapOutput) ToPublicIpAddressPoolCidrBlockMapOutput() PublicIpAddressPoolCidrBlockMapOutput {
	return o
}

func (o PublicIpAddressPoolCidrBlockMapOutput) ToPublicIpAddressPoolCidrBlockMapOutputWithContext(ctx context.Context) PublicIpAddressPoolCidrBlockMapOutput {
	return o
}

func (o PublicIpAddressPoolCidrBlockMapOutput) MapIndex(k pulumi.StringInput) PublicIpAddressPoolCidrBlockOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PublicIpAddressPoolCidrBlock {
		return vs[0].(map[string]*PublicIpAddressPoolCidrBlock)[vs[1].(string)]
	}).(PublicIpAddressPoolCidrBlockOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpAddressPoolCidrBlockInput)(nil)).Elem(), &PublicIpAddressPoolCidrBlock{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpAddressPoolCidrBlockArrayInput)(nil)).Elem(), PublicIpAddressPoolCidrBlockArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpAddressPoolCidrBlockMapInput)(nil)).Elem(), PublicIpAddressPoolCidrBlockMap{})
	pulumi.RegisterOutputType(PublicIpAddressPoolCidrBlockOutput{})
	pulumi.RegisterOutputType(PublicIpAddressPoolCidrBlockArrayOutput{})
	pulumi.RegisterOutputType(PublicIpAddressPoolCidrBlockMapOutput{})
}
