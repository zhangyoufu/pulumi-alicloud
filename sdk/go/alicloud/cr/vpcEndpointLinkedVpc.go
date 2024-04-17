// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CR Vpc Endpoint Linked Vpc resource.
//
// For information about CR Vpc Endpoint Linked Vpc and how to use it, see [What is Vpc Endpoint Linked Vpc](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createinstancevpcendpointlinkedvpc).
//
// > **NOTE:** Available since v1.199.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cr"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      pulumi.String(_default.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultRegistryEnterpriseInstance, err := cr.NewRegistryEnterpriseInstance(ctx, "default", &cr.RegistryEnterpriseInstanceArgs{
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(1),
//				RenewPeriod:   pulumi.Int(0),
//				RenewalStatus: pulumi.String("ManualRenewal"),
//				InstanceType:  pulumi.String("Advanced"),
//				InstanceName:  pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cr.NewVpcEndpointLinkedVpc(ctx, "default", &cr.VpcEndpointLinkedVpcArgs{
//				InstanceId:                  defaultRegistryEnterpriseInstance.ID(),
//				VpcId:                       defaultNetwork.ID(),
//				VswitchId:                   defaultSwitch.ID(),
//				ModuleName:                  pulumi.String("Registry"),
//				EnableCreateDnsRecordInPvzt: pulumi.Bool(true),
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
// CR Vpc Endpoint Linked Vpc can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cr/vpcEndpointLinkedVpc:VpcEndpointLinkedVpc example <instance_id>:<vpc_id>:<vswitch_id>:<module_name>
// ```
type VpcEndpointLinkedVpc struct {
	pulumi.CustomResourceState

	// Specifies whether to automatically create an Alibaba Cloud DNS PrivateZone record. Valid Values:
	EnableCreateDnsRecordInPvzt pulumi.BoolPtrOutput `pulumi:"enableCreateDnsRecordInPvzt"`
	// The ID of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The name of the module that you want to access. Valid Values:
	ModuleName pulumi.StringOutput `pulumi:"moduleName"`
	// The status of the Vpc Endpoint Linked Vpc.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The ID of the vSwitch.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewVpcEndpointLinkedVpc registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointLinkedVpc(ctx *pulumi.Context,
	name string, args *VpcEndpointLinkedVpcArgs, opts ...pulumi.ResourceOption) (*VpcEndpointLinkedVpc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.ModuleName == nil {
		return nil, errors.New("invalid value for required argument 'ModuleName'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcEndpointLinkedVpc
	err := ctx.RegisterResource("alicloud:cr/vpcEndpointLinkedVpc:VpcEndpointLinkedVpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointLinkedVpc gets an existing VpcEndpointLinkedVpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointLinkedVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointLinkedVpcState, opts ...pulumi.ResourceOption) (*VpcEndpointLinkedVpc, error) {
	var resource VpcEndpointLinkedVpc
	err := ctx.ReadResource("alicloud:cr/vpcEndpointLinkedVpc:VpcEndpointLinkedVpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointLinkedVpc resources.
type vpcEndpointLinkedVpcState struct {
	// Specifies whether to automatically create an Alibaba Cloud DNS PrivateZone record. Valid Values:
	EnableCreateDnsRecordInPvzt *bool `pulumi:"enableCreateDnsRecordInPvzt"`
	// The ID of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the module that you want to access. Valid Values:
	ModuleName *string `pulumi:"moduleName"`
	// The status of the Vpc Endpoint Linked Vpc.
	Status *string `pulumi:"status"`
	// The ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
	// The ID of the vSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

type VpcEndpointLinkedVpcState struct {
	// Specifies whether to automatically create an Alibaba Cloud DNS PrivateZone record. Valid Values:
	EnableCreateDnsRecordInPvzt pulumi.BoolPtrInput
	// The ID of the instance.
	InstanceId pulumi.StringPtrInput
	// The name of the module that you want to access. Valid Values:
	ModuleName pulumi.StringPtrInput
	// The status of the Vpc Endpoint Linked Vpc.
	Status pulumi.StringPtrInput
	// The ID of the VPC.
	VpcId pulumi.StringPtrInput
	// The ID of the vSwitch.
	VswitchId pulumi.StringPtrInput
}

func (VpcEndpointLinkedVpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointLinkedVpcState)(nil)).Elem()
}

type vpcEndpointLinkedVpcArgs struct {
	// Specifies whether to automatically create an Alibaba Cloud DNS PrivateZone record. Valid Values:
	EnableCreateDnsRecordInPvzt *bool `pulumi:"enableCreateDnsRecordInPvzt"`
	// The ID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// The name of the module that you want to access. Valid Values:
	ModuleName string `pulumi:"moduleName"`
	// The ID of the VPC.
	VpcId string `pulumi:"vpcId"`
	// The ID of the vSwitch.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a VpcEndpointLinkedVpc resource.
type VpcEndpointLinkedVpcArgs struct {
	// Specifies whether to automatically create an Alibaba Cloud DNS PrivateZone record. Valid Values:
	EnableCreateDnsRecordInPvzt pulumi.BoolPtrInput
	// The ID of the instance.
	InstanceId pulumi.StringInput
	// The name of the module that you want to access. Valid Values:
	ModuleName pulumi.StringInput
	// The ID of the VPC.
	VpcId pulumi.StringInput
	// The ID of the vSwitch.
	VswitchId pulumi.StringInput
}

func (VpcEndpointLinkedVpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointLinkedVpcArgs)(nil)).Elem()
}

type VpcEndpointLinkedVpcInput interface {
	pulumi.Input

	ToVpcEndpointLinkedVpcOutput() VpcEndpointLinkedVpcOutput
	ToVpcEndpointLinkedVpcOutputWithContext(ctx context.Context) VpcEndpointLinkedVpcOutput
}

func (*VpcEndpointLinkedVpc) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointLinkedVpc)(nil)).Elem()
}

func (i *VpcEndpointLinkedVpc) ToVpcEndpointLinkedVpcOutput() VpcEndpointLinkedVpcOutput {
	return i.ToVpcEndpointLinkedVpcOutputWithContext(context.Background())
}

func (i *VpcEndpointLinkedVpc) ToVpcEndpointLinkedVpcOutputWithContext(ctx context.Context) VpcEndpointLinkedVpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointLinkedVpcOutput)
}

// VpcEndpointLinkedVpcArrayInput is an input type that accepts VpcEndpointLinkedVpcArray and VpcEndpointLinkedVpcArrayOutput values.
// You can construct a concrete instance of `VpcEndpointLinkedVpcArrayInput` via:
//
//	VpcEndpointLinkedVpcArray{ VpcEndpointLinkedVpcArgs{...} }
type VpcEndpointLinkedVpcArrayInput interface {
	pulumi.Input

	ToVpcEndpointLinkedVpcArrayOutput() VpcEndpointLinkedVpcArrayOutput
	ToVpcEndpointLinkedVpcArrayOutputWithContext(context.Context) VpcEndpointLinkedVpcArrayOutput
}

type VpcEndpointLinkedVpcArray []VpcEndpointLinkedVpcInput

func (VpcEndpointLinkedVpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointLinkedVpc)(nil)).Elem()
}

func (i VpcEndpointLinkedVpcArray) ToVpcEndpointLinkedVpcArrayOutput() VpcEndpointLinkedVpcArrayOutput {
	return i.ToVpcEndpointLinkedVpcArrayOutputWithContext(context.Background())
}

func (i VpcEndpointLinkedVpcArray) ToVpcEndpointLinkedVpcArrayOutputWithContext(ctx context.Context) VpcEndpointLinkedVpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointLinkedVpcArrayOutput)
}

// VpcEndpointLinkedVpcMapInput is an input type that accepts VpcEndpointLinkedVpcMap and VpcEndpointLinkedVpcMapOutput values.
// You can construct a concrete instance of `VpcEndpointLinkedVpcMapInput` via:
//
//	VpcEndpointLinkedVpcMap{ "key": VpcEndpointLinkedVpcArgs{...} }
type VpcEndpointLinkedVpcMapInput interface {
	pulumi.Input

	ToVpcEndpointLinkedVpcMapOutput() VpcEndpointLinkedVpcMapOutput
	ToVpcEndpointLinkedVpcMapOutputWithContext(context.Context) VpcEndpointLinkedVpcMapOutput
}

type VpcEndpointLinkedVpcMap map[string]VpcEndpointLinkedVpcInput

func (VpcEndpointLinkedVpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointLinkedVpc)(nil)).Elem()
}

func (i VpcEndpointLinkedVpcMap) ToVpcEndpointLinkedVpcMapOutput() VpcEndpointLinkedVpcMapOutput {
	return i.ToVpcEndpointLinkedVpcMapOutputWithContext(context.Background())
}

func (i VpcEndpointLinkedVpcMap) ToVpcEndpointLinkedVpcMapOutputWithContext(ctx context.Context) VpcEndpointLinkedVpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointLinkedVpcMapOutput)
}

type VpcEndpointLinkedVpcOutput struct{ *pulumi.OutputState }

func (VpcEndpointLinkedVpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointLinkedVpc)(nil)).Elem()
}

func (o VpcEndpointLinkedVpcOutput) ToVpcEndpointLinkedVpcOutput() VpcEndpointLinkedVpcOutput {
	return o
}

func (o VpcEndpointLinkedVpcOutput) ToVpcEndpointLinkedVpcOutputWithContext(ctx context.Context) VpcEndpointLinkedVpcOutput {
	return o
}

// Specifies whether to automatically create an Alibaba Cloud DNS PrivateZone record. Valid Values:
func (o VpcEndpointLinkedVpcOutput) EnableCreateDnsRecordInPvzt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcEndpointLinkedVpc) pulumi.BoolPtrOutput { return v.EnableCreateDnsRecordInPvzt }).(pulumi.BoolPtrOutput)
}

// The ID of the instance.
func (o VpcEndpointLinkedVpcOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointLinkedVpc) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The name of the module that you want to access. Valid Values:
func (o VpcEndpointLinkedVpcOutput) ModuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointLinkedVpc) pulumi.StringOutput { return v.ModuleName }).(pulumi.StringOutput)
}

// The status of the Vpc Endpoint Linked Vpc.
func (o VpcEndpointLinkedVpcOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointLinkedVpc) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the VPC.
func (o VpcEndpointLinkedVpcOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointLinkedVpc) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The ID of the vSwitch.
func (o VpcEndpointLinkedVpcOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointLinkedVpc) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type VpcEndpointLinkedVpcArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointLinkedVpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointLinkedVpc)(nil)).Elem()
}

func (o VpcEndpointLinkedVpcArrayOutput) ToVpcEndpointLinkedVpcArrayOutput() VpcEndpointLinkedVpcArrayOutput {
	return o
}

func (o VpcEndpointLinkedVpcArrayOutput) ToVpcEndpointLinkedVpcArrayOutputWithContext(ctx context.Context) VpcEndpointLinkedVpcArrayOutput {
	return o
}

func (o VpcEndpointLinkedVpcArrayOutput) Index(i pulumi.IntInput) VpcEndpointLinkedVpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointLinkedVpc {
		return vs[0].([]*VpcEndpointLinkedVpc)[vs[1].(int)]
	}).(VpcEndpointLinkedVpcOutput)
}

type VpcEndpointLinkedVpcMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointLinkedVpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointLinkedVpc)(nil)).Elem()
}

func (o VpcEndpointLinkedVpcMapOutput) ToVpcEndpointLinkedVpcMapOutput() VpcEndpointLinkedVpcMapOutput {
	return o
}

func (o VpcEndpointLinkedVpcMapOutput) ToVpcEndpointLinkedVpcMapOutputWithContext(ctx context.Context) VpcEndpointLinkedVpcMapOutput {
	return o
}

func (o VpcEndpointLinkedVpcMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointLinkedVpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointLinkedVpc {
		return vs[0].(map[string]*VpcEndpointLinkedVpc)[vs[1].(string)]
	}).(VpcEndpointLinkedVpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointLinkedVpcInput)(nil)).Elem(), &VpcEndpointLinkedVpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointLinkedVpcArrayInput)(nil)).Elem(), VpcEndpointLinkedVpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointLinkedVpcMapInput)(nil)).Elem(), VpcEndpointLinkedVpcMap{})
	pulumi.RegisterOutputType(VpcEndpointLinkedVpcOutput{})
	pulumi.RegisterOutputType(VpcEndpointLinkedVpcArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointLinkedVpcMapOutput{})
}
