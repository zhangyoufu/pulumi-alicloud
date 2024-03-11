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

// Provides a VPC Ipv6 Address resource.
//
// For information about VPC Ipv6 Address and how to use it, see [What is Ipv6 Address](https://next.api.alibabacloud.com/document/Vpc/2016-04-28/AllocateIpv6Address).
//
// > **NOTE:** Available since v1.216.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
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
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := vpc.NewNetwork(ctx, "vpc", &vpc.NetworkArgs{
//				Ipv6Isp:    pulumi.String("BGP"),
//				CidrBlock:  pulumi.String("172.168.0.0/16"),
//				EnableIpv6: pulumi.Bool(true),
//				VpcName:    pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			vswich, err := vpc.NewSwitch(ctx, "vswich", &vpc.SwitchArgs{
//				VpcId:             vpc.ID(),
//				CidrBlock:         pulumi.String("172.168.0.0/24"),
//				ZoneId:            *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName:       pulumi.String(name),
//				Ipv6CidrBlockMask: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewIpv6Address(ctx, "defaultIpv6Address", &vpc.Ipv6AddressArgs{
//				ResourceGroupId:        *pulumi.String(defaultResourceGroups.Ids[0]),
//				VswitchId:              vswich.ID(),
//				Ipv6AddressDescription: pulumi.String("create_description"),
//				Ipv6AddressName:        pulumi.String(name),
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
// VPC Ipv6 Address can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/ipv6Address:Ipv6Address example <id>
// ```
type Ipv6Address struct {
	pulumi.CustomResourceState

	// The creation time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// IPv6 address.
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
	// The description of the IPv6 Address. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
	Ipv6AddressDescription pulumi.StringOutput `pulumi:"ipv6AddressDescription"`
	// The name of the IPv6 Address. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
	Ipv6AddressName pulumi.StringPtrOutput `pulumi:"ipv6AddressName"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The status of the resource.  Available, Pending and Deleting.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags for the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The VSwitchId of the IPv6 address.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewIpv6Address registers a new resource with the given unique name, arguments, and options.
func NewIpv6Address(ctx *pulumi.Context,
	name string, args *Ipv6AddressArgs, opts ...pulumi.ResourceOption) (*Ipv6Address, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ipv6Address
	err := ctx.RegisterResource("alicloud:vpc/ipv6Address:Ipv6Address", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpv6Address gets an existing Ipv6Address resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpv6Address(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipv6AddressState, opts ...pulumi.ResourceOption) (*Ipv6Address, error) {
	var resource Ipv6Address
	err := ctx.ReadResource("alicloud:vpc/ipv6Address:Ipv6Address", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipv6Address resources.
type ipv6AddressState struct {
	// The creation time of the resource.
	CreateTime *string `pulumi:"createTime"`
	// IPv6 address.
	Ipv6Address *string `pulumi:"ipv6Address"`
	// The description of the IPv6 Address. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
	Ipv6AddressDescription *string `pulumi:"ipv6AddressDescription"`
	// The name of the IPv6 Address. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
	Ipv6AddressName *string `pulumi:"ipv6AddressName"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the resource.  Available, Pending and Deleting.
	Status *string `pulumi:"status"`
	// The tags for the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VSwitchId of the IPv6 address.
	VswitchId *string `pulumi:"vswitchId"`
}

type Ipv6AddressState struct {
	// The creation time of the resource.
	CreateTime pulumi.StringPtrInput
	// IPv6 address.
	Ipv6Address pulumi.StringPtrInput
	// The description of the IPv6 Address. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
	Ipv6AddressDescription pulumi.StringPtrInput
	// The name of the IPv6 Address. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
	Ipv6AddressName pulumi.StringPtrInput
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The status of the resource.  Available, Pending and Deleting.
	Status pulumi.StringPtrInput
	// The tags for the resource.
	Tags pulumi.MapInput
	// The VSwitchId of the IPv6 address.
	VswitchId pulumi.StringPtrInput
}

func (Ipv6AddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6AddressState)(nil)).Elem()
}

type ipv6AddressArgs struct {
	// The description of the IPv6 Address. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
	Ipv6AddressDescription *string `pulumi:"ipv6AddressDescription"`
	// The name of the IPv6 Address. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
	Ipv6AddressName *string `pulumi:"ipv6AddressName"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The tags for the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VSwitchId of the IPv6 address.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a Ipv6Address resource.
type Ipv6AddressArgs struct {
	// The description of the IPv6 Address. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
	Ipv6AddressDescription pulumi.StringPtrInput
	// The name of the IPv6 Address. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
	Ipv6AddressName pulumi.StringPtrInput
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The tags for the resource.
	Tags pulumi.MapInput
	// The VSwitchId of the IPv6 address.
	VswitchId pulumi.StringInput
}

func (Ipv6AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6AddressArgs)(nil)).Elem()
}

type Ipv6AddressInput interface {
	pulumi.Input

	ToIpv6AddressOutput() Ipv6AddressOutput
	ToIpv6AddressOutputWithContext(ctx context.Context) Ipv6AddressOutput
}

func (*Ipv6Address) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6Address)(nil)).Elem()
}

func (i *Ipv6Address) ToIpv6AddressOutput() Ipv6AddressOutput {
	return i.ToIpv6AddressOutputWithContext(context.Background())
}

func (i *Ipv6Address) ToIpv6AddressOutputWithContext(ctx context.Context) Ipv6AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6AddressOutput)
}

// Ipv6AddressArrayInput is an input type that accepts Ipv6AddressArray and Ipv6AddressArrayOutput values.
// You can construct a concrete instance of `Ipv6AddressArrayInput` via:
//
//	Ipv6AddressArray{ Ipv6AddressArgs{...} }
type Ipv6AddressArrayInput interface {
	pulumi.Input

	ToIpv6AddressArrayOutput() Ipv6AddressArrayOutput
	ToIpv6AddressArrayOutputWithContext(context.Context) Ipv6AddressArrayOutput
}

type Ipv6AddressArray []Ipv6AddressInput

func (Ipv6AddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv6Address)(nil)).Elem()
}

func (i Ipv6AddressArray) ToIpv6AddressArrayOutput() Ipv6AddressArrayOutput {
	return i.ToIpv6AddressArrayOutputWithContext(context.Background())
}

func (i Ipv6AddressArray) ToIpv6AddressArrayOutputWithContext(ctx context.Context) Ipv6AddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6AddressArrayOutput)
}

// Ipv6AddressMapInput is an input type that accepts Ipv6AddressMap and Ipv6AddressMapOutput values.
// You can construct a concrete instance of `Ipv6AddressMapInput` via:
//
//	Ipv6AddressMap{ "key": Ipv6AddressArgs{...} }
type Ipv6AddressMapInput interface {
	pulumi.Input

	ToIpv6AddressMapOutput() Ipv6AddressMapOutput
	ToIpv6AddressMapOutputWithContext(context.Context) Ipv6AddressMapOutput
}

type Ipv6AddressMap map[string]Ipv6AddressInput

func (Ipv6AddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv6Address)(nil)).Elem()
}

func (i Ipv6AddressMap) ToIpv6AddressMapOutput() Ipv6AddressMapOutput {
	return i.ToIpv6AddressMapOutputWithContext(context.Background())
}

func (i Ipv6AddressMap) ToIpv6AddressMapOutputWithContext(ctx context.Context) Ipv6AddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6AddressMapOutput)
}

type Ipv6AddressOutput struct{ *pulumi.OutputState }

func (Ipv6AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6Address)(nil)).Elem()
}

func (o Ipv6AddressOutput) ToIpv6AddressOutput() Ipv6AddressOutput {
	return o
}

func (o Ipv6AddressOutput) ToIpv6AddressOutputWithContext(ctx context.Context) Ipv6AddressOutput {
	return o
}

// The creation time of the resource.
func (o Ipv6AddressOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6Address) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// IPv6 address.
func (o Ipv6AddressOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6Address) pulumi.StringOutput { return v.Ipv6Address }).(pulumi.StringOutput)
}

// The description of the IPv6 Address. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
func (o Ipv6AddressOutput) Ipv6AddressDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6Address) pulumi.StringOutput { return v.Ipv6AddressDescription }).(pulumi.StringOutput)
}

// The name of the IPv6 Address. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
func (o Ipv6AddressOutput) Ipv6AddressName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6Address) pulumi.StringPtrOutput { return v.Ipv6AddressName }).(pulumi.StringPtrOutput)
}

// The ID of the resource group to which the instance belongs.
func (o Ipv6AddressOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6Address) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of the resource.  Available, Pending and Deleting.
func (o Ipv6AddressOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6Address) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags for the resource.
func (o Ipv6AddressOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Ipv6Address) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The VSwitchId of the IPv6 address.
func (o Ipv6AddressOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv6Address) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type Ipv6AddressArrayOutput struct{ *pulumi.OutputState }

func (Ipv6AddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv6Address)(nil)).Elem()
}

func (o Ipv6AddressArrayOutput) ToIpv6AddressArrayOutput() Ipv6AddressArrayOutput {
	return o
}

func (o Ipv6AddressArrayOutput) ToIpv6AddressArrayOutputWithContext(ctx context.Context) Ipv6AddressArrayOutput {
	return o
}

func (o Ipv6AddressArrayOutput) Index(i pulumi.IntInput) Ipv6AddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipv6Address {
		return vs[0].([]*Ipv6Address)[vs[1].(int)]
	}).(Ipv6AddressOutput)
}

type Ipv6AddressMapOutput struct{ *pulumi.OutputState }

func (Ipv6AddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv6Address)(nil)).Elem()
}

func (o Ipv6AddressMapOutput) ToIpv6AddressMapOutput() Ipv6AddressMapOutput {
	return o
}

func (o Ipv6AddressMapOutput) ToIpv6AddressMapOutputWithContext(ctx context.Context) Ipv6AddressMapOutput {
	return o
}

func (o Ipv6AddressMapOutput) MapIndex(k pulumi.StringInput) Ipv6AddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipv6Address {
		return vs[0].(map[string]*Ipv6Address)[vs[1].(string)]
	}).(Ipv6AddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6AddressInput)(nil)).Elem(), &Ipv6Address{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6AddressArrayInput)(nil)).Elem(), Ipv6AddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6AddressMapInput)(nil)).Elem(), Ipv6AddressMap{})
	pulumi.RegisterOutputType(Ipv6AddressOutput{})
	pulumi.RegisterOutputType(Ipv6AddressArrayOutput{})
	pulumi.RegisterOutputType(Ipv6AddressMapOutput{})
}
