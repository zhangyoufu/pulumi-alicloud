// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Microservice Engine (MSE) Gateway resource.
//
// For information about Microservice Engine (MSE) Gateway and how to use it, see [What is Gateway](https://help.aliyun.com/document_detail/347638.html).
//
// > **NOTE:** Available in v1.157.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mse"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultZones, err := alicloud.GetZones(ctx, &GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("default-NODELETING"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
//				VpcId:  pulumi.StringRef(defaultNetworks.Ids[0]),
//				ZoneId: pulumi.StringRef(defaultZones.Zones[0].Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = mse.NewGateway(ctx, "example", &mse.GatewayArgs{
//				GatewayName:     pulumi.String("example_value"),
//				Replica:         pulumi.Int(2),
//				Spec:            pulumi.String("MSE_GTW_2_4_200_c"),
//				VswitchId:       pulumi.String(defaultSwitches.Ids[0]),
//				BackupVswitchId: pulumi.String(defaultSwitches.Ids[1]),
//				VpcId:           pulumi.String(defaultNetworks.Ids[0]),
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
// Microservice Engine (MSE) Gateway can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:mse/gateway:Gateway example <id>
//
// ```
type Gateway struct {
	pulumi.CustomResourceState

	// The backup vswitch id.
	BackupVswitchId pulumi.StringPtrOutput `pulumi:"backupVswitchId"`
	// Whether to delete the SLB purchased on behalf of the gateway at the same time.
	DeleteSlb pulumi.BoolPtrOutput `pulumi:"deleteSlb"`
	// Whether the enterprise security group type.
	EnterpriseSecurityGroup pulumi.BoolPtrOutput `pulumi:"enterpriseSecurityGroup"`
	// The name of the Gateway .
	GatewayName pulumi.StringPtrOutput `pulumi:"gatewayName"`
	// Public network SLB specifications.
	InternetSlbSpec pulumi.StringPtrOutput `pulumi:"internetSlbSpec"`
	// Number of Gateway Nodes.
	Replica pulumi.IntOutput `pulumi:"replica"`
	// A list of gateway Slb.
	SlbLists GatewaySlbListArrayOutput `pulumi:"slbLists"`
	// Private network SLB specifications.
	SlbSpec pulumi.StringPtrOutput `pulumi:"slbSpec"`
	// Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
	Spec pulumi.StringOutput `pulumi:"spec"`
	// The status of the gateway.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the vpc.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The ID of the vswitch.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Replica == nil {
		return nil, errors.New("invalid value for required argument 'Replica'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource Gateway
	err := ctx.RegisterResource("alicloud:mse/gateway:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("alicloud:mse/gateway:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
	// The backup vswitch id.
	BackupVswitchId *string `pulumi:"backupVswitchId"`
	// Whether to delete the SLB purchased on behalf of the gateway at the same time.
	DeleteSlb *bool `pulumi:"deleteSlb"`
	// Whether the enterprise security group type.
	EnterpriseSecurityGroup *bool `pulumi:"enterpriseSecurityGroup"`
	// The name of the Gateway .
	GatewayName *string `pulumi:"gatewayName"`
	// Public network SLB specifications.
	InternetSlbSpec *string `pulumi:"internetSlbSpec"`
	// Number of Gateway Nodes.
	Replica *int `pulumi:"replica"`
	// A list of gateway Slb.
	SlbLists []GatewaySlbList `pulumi:"slbLists"`
	// Private network SLB specifications.
	SlbSpec *string `pulumi:"slbSpec"`
	// Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
	Spec *string `pulumi:"spec"`
	// The status of the gateway.
	Status *string `pulumi:"status"`
	// The ID of the vpc.
	VpcId *string `pulumi:"vpcId"`
	// The ID of the vswitch.
	VswitchId *string `pulumi:"vswitchId"`
}

type GatewayState struct {
	// The backup vswitch id.
	BackupVswitchId pulumi.StringPtrInput
	// Whether to delete the SLB purchased on behalf of the gateway at the same time.
	DeleteSlb pulumi.BoolPtrInput
	// Whether the enterprise security group type.
	EnterpriseSecurityGroup pulumi.BoolPtrInput
	// The name of the Gateway .
	GatewayName pulumi.StringPtrInput
	// Public network SLB specifications.
	InternetSlbSpec pulumi.StringPtrInput
	// Number of Gateway Nodes.
	Replica pulumi.IntPtrInput
	// A list of gateway Slb.
	SlbLists GatewaySlbListArrayInput
	// Private network SLB specifications.
	SlbSpec pulumi.StringPtrInput
	// Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
	Spec pulumi.StringPtrInput
	// The status of the gateway.
	Status pulumi.StringPtrInput
	// The ID of the vpc.
	VpcId pulumi.StringPtrInput
	// The ID of the vswitch.
	VswitchId pulumi.StringPtrInput
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// The backup vswitch id.
	BackupVswitchId *string `pulumi:"backupVswitchId"`
	// Whether to delete the SLB purchased on behalf of the gateway at the same time.
	DeleteSlb *bool `pulumi:"deleteSlb"`
	// Whether the enterprise security group type.
	EnterpriseSecurityGroup *bool `pulumi:"enterpriseSecurityGroup"`
	// The name of the Gateway .
	GatewayName *string `pulumi:"gatewayName"`
	// Public network SLB specifications.
	InternetSlbSpec *string `pulumi:"internetSlbSpec"`
	// Number of Gateway Nodes.
	Replica int `pulumi:"replica"`
	// Private network SLB specifications.
	SlbSpec *string `pulumi:"slbSpec"`
	// Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
	Spec string `pulumi:"spec"`
	// The ID of the vpc.
	VpcId string `pulumi:"vpcId"`
	// The ID of the vswitch.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// The backup vswitch id.
	BackupVswitchId pulumi.StringPtrInput
	// Whether to delete the SLB purchased on behalf of the gateway at the same time.
	DeleteSlb pulumi.BoolPtrInput
	// Whether the enterprise security group type.
	EnterpriseSecurityGroup pulumi.BoolPtrInput
	// The name of the Gateway .
	GatewayName pulumi.StringPtrInput
	// Public network SLB specifications.
	InternetSlbSpec pulumi.StringPtrInput
	// Number of Gateway Nodes.
	Replica pulumi.IntInput
	// Private network SLB specifications.
	SlbSpec pulumi.StringPtrInput
	// Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
	Spec pulumi.StringInput
	// The ID of the vpc.
	VpcId pulumi.StringInput
	// The ID of the vswitch.
	VswitchId pulumi.StringInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

// GatewayArrayInput is an input type that accepts GatewayArray and GatewayArrayOutput values.
// You can construct a concrete instance of `GatewayArrayInput` via:
//
//	GatewayArray{ GatewayArgs{...} }
type GatewayArrayInput interface {
	pulumi.Input

	ToGatewayArrayOutput() GatewayArrayOutput
	ToGatewayArrayOutputWithContext(context.Context) GatewayArrayOutput
}

type GatewayArray []GatewayInput

func (GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (i GatewayArray) ToGatewayArrayOutput() GatewayArrayOutput {
	return i.ToGatewayArrayOutputWithContext(context.Background())
}

func (i GatewayArray) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayArrayOutput)
}

// GatewayMapInput is an input type that accepts GatewayMap and GatewayMapOutput values.
// You can construct a concrete instance of `GatewayMapInput` via:
//
//	GatewayMap{ "key": GatewayArgs{...} }
type GatewayMapInput interface {
	pulumi.Input

	ToGatewayMapOutput() GatewayMapOutput
	ToGatewayMapOutputWithContext(context.Context) GatewayMapOutput
}

type GatewayMap map[string]GatewayInput

func (GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (i GatewayMap) ToGatewayMapOutput() GatewayMapOutput {
	return i.ToGatewayMapOutputWithContext(context.Background())
}

func (i GatewayMap) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayMapOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

// The backup vswitch id.
func (o GatewayOutput) BackupVswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.BackupVswitchId }).(pulumi.StringPtrOutput)
}

// Whether to delete the SLB purchased on behalf of the gateway at the same time.
func (o GatewayOutput) DeleteSlb() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.BoolPtrOutput { return v.DeleteSlb }).(pulumi.BoolPtrOutput)
}

// Whether the enterprise security group type.
func (o GatewayOutput) EnterpriseSecurityGroup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.BoolPtrOutput { return v.EnterpriseSecurityGroup }).(pulumi.BoolPtrOutput)
}

// The name of the Gateway .
func (o GatewayOutput) GatewayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.GatewayName }).(pulumi.StringPtrOutput)
}

// Public network SLB specifications.
func (o GatewayOutput) InternetSlbSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.InternetSlbSpec }).(pulumi.StringPtrOutput)
}

// Number of Gateway Nodes.
func (o GatewayOutput) Replica() pulumi.IntOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntOutput { return v.Replica }).(pulumi.IntOutput)
}

// A list of gateway Slb.
func (o GatewayOutput) SlbLists() GatewaySlbListArrayOutput {
	return o.ApplyT(func(v *Gateway) GatewaySlbListArrayOutput { return v.SlbLists }).(GatewaySlbListArrayOutput)
}

// Private network SLB specifications.
func (o GatewayOutput) SlbSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.SlbSpec }).(pulumi.StringPtrOutput)
}

// Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
func (o GatewayOutput) Spec() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Spec }).(pulumi.StringOutput)
}

// The status of the gateway.
func (o GatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the vpc.
func (o GatewayOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The ID of the vswitch.
func (o GatewayOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type GatewayArrayOutput struct{ *pulumi.OutputState }

func (GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (o GatewayArrayOutput) ToGatewayArrayOutput() GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) Index(i pulumi.IntInput) GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].([]*Gateway)[vs[1].(int)]
	}).(GatewayOutput)
}

type GatewayMapOutput struct{ *pulumi.OutputState }

func (GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (o GatewayMapOutput) ToGatewayMapOutput() GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) MapIndex(k pulumi.StringInput) GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].(map[string]*Gateway)[vs[1].(string)]
	}).(GatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayInput)(nil)).Elem(), &Gateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayArrayInput)(nil)).Elem(), GatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayMapInput)(nil)).Elem(), GatewayMap{})
	pulumi.RegisterOutputType(GatewayOutput{})
	pulumi.RegisterOutputType(GatewayArrayOutput{})
	pulumi.RegisterOutputType(GatewayMapOutput{})
}
