// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vpc Ipv4 Gateway resource.
//
// For information about Vpc Ipv4 Gateway and how to use it, see [What is Ipv4 Gateway](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createipv4gateway).
//
// > **NOTE:** Available in v1.181.0+.
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
//	"fmt"
//
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
//			name := "tf-testacc-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultResourceGroup, err := resourcemanager.NewResourceGroup(ctx, "defaultResourceGroup", &resourcemanager.ResourceGroupArgs{
//				DisplayName:       pulumi.String("tf-testAcc-rg665"),
//				ResourceGroupName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = resourcemanager.NewResourceGroup(ctx, "modify", &resourcemanager.ResourceGroupArgs{
//				DisplayName:       pulumi.String("tf-testAcc-rg298"),
//				ResourceGroupName: pulumi.String(fmt.Sprintf("%v1", name)),
//			})
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(fmt.Sprintf("%v2", name)),
//				CidrBlock: pulumi.String("10.0.0.0/8"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewIpv4Gateway(ctx, "defaultIpv4Gateway", &vpc.Ipv4GatewayArgs{
//				Ipv4GatewayName:        pulumi.String(name),
//				Ipv4GatewayDescription: pulumi.String("tf-testAcc-Ipv4Gateway"),
//				ResourceGroupId:        defaultResourceGroup.ID(),
//				VpcId:                  defaultNetwork.ID(),
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
// Vpc Ipv4 Gateway can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/ipv4Gateway:Ipv4Gateway example <id>
//
// ```
type Ipv4Gateway struct {
	pulumi.CustomResourceState

	// The creation time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
	Ipv4GatewayDescription pulumi.StringPtrOutput `pulumi:"ipv4GatewayDescription"`
	// Resource primary key field.
	Ipv4GatewayId pulumi.StringOutput `pulumi:"ipv4GatewayId"`
	// The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	Ipv4GatewayName pulumi.StringPtrOutput `pulumi:"ipv4GatewayName"`
	// ID of the route table associated with IPv4 Gateway.
	Ipv4GatewayRouteTableId pulumi.StringOutput `pulumi:"ipv4GatewayRouteTableId"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags of the current resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewIpv4Gateway registers a new resource with the given unique name, arguments, and options.
func NewIpv4Gateway(ctx *pulumi.Context,
	name string, args *Ipv4GatewayArgs, opts ...pulumi.ResourceOption) (*Ipv4Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource Ipv4Gateway
	err := ctx.RegisterResource("alicloud:vpc/ipv4Gateway:Ipv4Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpv4Gateway gets an existing Ipv4Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpv4Gateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipv4GatewayState, opts ...pulumi.ResourceOption) (*Ipv4Gateway, error) {
	var resource Ipv4Gateway
	err := ctx.ReadResource("alicloud:vpc/ipv4Gateway:Ipv4Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipv4Gateway resources.
type ipv4GatewayState struct {
	// The creation time of the resource.
	CreateTime *string `pulumi:"createTime"`
	// Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
	DryRun *bool `pulumi:"dryRun"`
	// Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
	Enabled *bool `pulumi:"enabled"`
	// The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
	Ipv4GatewayDescription *string `pulumi:"ipv4GatewayDescription"`
	// Resource primary key field.
	Ipv4GatewayId *string `pulumi:"ipv4GatewayId"`
	// The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	Ipv4GatewayName *string `pulumi:"ipv4GatewayName"`
	// ID of the route table associated with IPv4 Gateway.
	Ipv4GatewayRouteTableId *string `pulumi:"ipv4GatewayRouteTableId"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The tags of the current resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
	VpcId *string `pulumi:"vpcId"`
}

type Ipv4GatewayState struct {
	// The creation time of the resource.
	CreateTime pulumi.StringPtrInput
	// Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
	DryRun pulumi.BoolPtrInput
	// Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
	Enabled pulumi.BoolPtrInput
	// The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
	Ipv4GatewayDescription pulumi.StringPtrInput
	// Resource primary key field.
	Ipv4GatewayId pulumi.StringPtrInput
	// The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	Ipv4GatewayName pulumi.StringPtrInput
	// ID of the route table associated with IPv4 Gateway.
	Ipv4GatewayRouteTableId pulumi.StringPtrInput
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The tags of the current resource.
	Tags pulumi.MapInput
	// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
	VpcId pulumi.StringPtrInput
}

func (Ipv4GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv4GatewayState)(nil)).Elem()
}

type ipv4GatewayArgs struct {
	// Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
	DryRun *bool `pulumi:"dryRun"`
	// Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
	Enabled *bool `pulumi:"enabled"`
	// The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
	Ipv4GatewayDescription *string `pulumi:"ipv4GatewayDescription"`
	// The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	Ipv4GatewayName *string `pulumi:"ipv4GatewayName"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The tags of the current resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Ipv4Gateway resource.
type Ipv4GatewayArgs struct {
	// Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
	DryRun pulumi.BoolPtrInput
	// Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
	Enabled pulumi.BoolPtrInput
	// The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
	Ipv4GatewayDescription pulumi.StringPtrInput
	// The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	Ipv4GatewayName pulumi.StringPtrInput
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The tags of the current resource.
	Tags pulumi.MapInput
	// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
	VpcId pulumi.StringInput
}

func (Ipv4GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv4GatewayArgs)(nil)).Elem()
}

type Ipv4GatewayInput interface {
	pulumi.Input

	ToIpv4GatewayOutput() Ipv4GatewayOutput
	ToIpv4GatewayOutputWithContext(ctx context.Context) Ipv4GatewayOutput
}

func (*Ipv4Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv4Gateway)(nil)).Elem()
}

func (i *Ipv4Gateway) ToIpv4GatewayOutput() Ipv4GatewayOutput {
	return i.ToIpv4GatewayOutputWithContext(context.Background())
}

func (i *Ipv4Gateway) ToIpv4GatewayOutputWithContext(ctx context.Context) Ipv4GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4GatewayOutput)
}

// Ipv4GatewayArrayInput is an input type that accepts Ipv4GatewayArray and Ipv4GatewayArrayOutput values.
// You can construct a concrete instance of `Ipv4GatewayArrayInput` via:
//
//	Ipv4GatewayArray{ Ipv4GatewayArgs{...} }
type Ipv4GatewayArrayInput interface {
	pulumi.Input

	ToIpv4GatewayArrayOutput() Ipv4GatewayArrayOutput
	ToIpv4GatewayArrayOutputWithContext(context.Context) Ipv4GatewayArrayOutput
}

type Ipv4GatewayArray []Ipv4GatewayInput

func (Ipv4GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv4Gateway)(nil)).Elem()
}

func (i Ipv4GatewayArray) ToIpv4GatewayArrayOutput() Ipv4GatewayArrayOutput {
	return i.ToIpv4GatewayArrayOutputWithContext(context.Background())
}

func (i Ipv4GatewayArray) ToIpv4GatewayArrayOutputWithContext(ctx context.Context) Ipv4GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4GatewayArrayOutput)
}

// Ipv4GatewayMapInput is an input type that accepts Ipv4GatewayMap and Ipv4GatewayMapOutput values.
// You can construct a concrete instance of `Ipv4GatewayMapInput` via:
//
//	Ipv4GatewayMap{ "key": Ipv4GatewayArgs{...} }
type Ipv4GatewayMapInput interface {
	pulumi.Input

	ToIpv4GatewayMapOutput() Ipv4GatewayMapOutput
	ToIpv4GatewayMapOutputWithContext(context.Context) Ipv4GatewayMapOutput
}

type Ipv4GatewayMap map[string]Ipv4GatewayInput

func (Ipv4GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv4Gateway)(nil)).Elem()
}

func (i Ipv4GatewayMap) ToIpv4GatewayMapOutput() Ipv4GatewayMapOutput {
	return i.ToIpv4GatewayMapOutputWithContext(context.Background())
}

func (i Ipv4GatewayMap) ToIpv4GatewayMapOutputWithContext(ctx context.Context) Ipv4GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4GatewayMapOutput)
}

type Ipv4GatewayOutput struct{ *pulumi.OutputState }

func (Ipv4GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv4Gateway)(nil)).Elem()
}

func (o Ipv4GatewayOutput) ToIpv4GatewayOutput() Ipv4GatewayOutput {
	return o
}

func (o Ipv4GatewayOutput) ToIpv4GatewayOutputWithContext(ctx context.Context) Ipv4GatewayOutput {
	return o
}

// The creation time of the resource.
func (o Ipv4GatewayOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
func (o Ipv4GatewayOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
func (o Ipv4GatewayOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
func (o Ipv4GatewayOutput) Ipv4GatewayDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringPtrOutput { return v.Ipv4GatewayDescription }).(pulumi.StringPtrOutput)
}

// Resource primary key field.
func (o Ipv4GatewayOutput) Ipv4GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringOutput { return v.Ipv4GatewayId }).(pulumi.StringOutput)
}

// The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
func (o Ipv4GatewayOutput) Ipv4GatewayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringPtrOutput { return v.Ipv4GatewayName }).(pulumi.StringPtrOutput)
}

// ID of the route table associated with IPv4 Gateway.
func (o Ipv4GatewayOutput) Ipv4GatewayRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringOutput { return v.Ipv4GatewayRouteTableId }).(pulumi.StringOutput)
}

// The ID of the resource group to which the instance belongs.
func (o Ipv4GatewayOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of the resource.
func (o Ipv4GatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags of the current resource.
func (o Ipv4GatewayOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
func (o Ipv4GatewayOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type Ipv4GatewayArrayOutput struct{ *pulumi.OutputState }

func (Ipv4GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv4Gateway)(nil)).Elem()
}

func (o Ipv4GatewayArrayOutput) ToIpv4GatewayArrayOutput() Ipv4GatewayArrayOutput {
	return o
}

func (o Ipv4GatewayArrayOutput) ToIpv4GatewayArrayOutputWithContext(ctx context.Context) Ipv4GatewayArrayOutput {
	return o
}

func (o Ipv4GatewayArrayOutput) Index(i pulumi.IntInput) Ipv4GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipv4Gateway {
		return vs[0].([]*Ipv4Gateway)[vs[1].(int)]
	}).(Ipv4GatewayOutput)
}

type Ipv4GatewayMapOutput struct{ *pulumi.OutputState }

func (Ipv4GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv4Gateway)(nil)).Elem()
}

func (o Ipv4GatewayMapOutput) ToIpv4GatewayMapOutput() Ipv4GatewayMapOutput {
	return o
}

func (o Ipv4GatewayMapOutput) ToIpv4GatewayMapOutputWithContext(ctx context.Context) Ipv4GatewayMapOutput {
	return o
}

func (o Ipv4GatewayMapOutput) MapIndex(k pulumi.StringInput) Ipv4GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipv4Gateway {
		return vs[0].(map[string]*Ipv4Gateway)[vs[1].(string)]
	}).(Ipv4GatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4GatewayInput)(nil)).Elem(), &Ipv4Gateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4GatewayArrayInput)(nil)).Elem(), Ipv4GatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4GatewayMapInput)(nil)).Elem(), Ipv4GatewayMap{})
	pulumi.RegisterOutputType(Ipv4GatewayOutput{})
	pulumi.RegisterOutputType(Ipv4GatewayArrayOutput{})
	pulumi.RegisterOutputType(Ipv4GatewayMapOutput{})
}
