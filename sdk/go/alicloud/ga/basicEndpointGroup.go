// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Basic Endpoint Group resource.
//
// For information about Global Accelerator (GA) Basic Endpoint Group and how to use it, see [What is Basic Endpoint Group](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbasicendpointgroup).
//
// > **NOTE:** Available since v1.194.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			region := "cn-hangzhou"
//			if param := cfg.Get("region"); param != "" {
//				region = param
//			}
//			endpointGroupRegion := "cn-beijing"
//			if param := cfg.Get("endpointGroupRegion"); param != "" {
//				endpointGroupRegion = param
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("172.17.3.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String("terraform-example"),
//				CidrBlock:   pulumi.String("172.17.3.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultApplicationLoadBalancer, err := slb.NewApplicationLoadBalancer(ctx, "defaultApplicationLoadBalancer", &slb.ApplicationLoadBalancerArgs{
//				LoadBalancerName: pulumi.String("terraform-example"),
//				VswitchId:        defaultSwitch.ID(),
//				LoadBalancerSpec: pulumi.String("slb.s2.small"),
//				AddressType:      pulumi.String("intranet"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBasicAccelerator, err := ga.NewBasicAccelerator(ctx, "defaultBasicAccelerator", &ga.BasicAcceleratorArgs{
//				Duration:             pulumi.Int(1),
//				BasicAcceleratorName: pulumi.String("terraform-example"),
//				Description:          pulumi.String("terraform-example"),
//				BandwidthBillingType: pulumi.String("CDT"),
//				AutoUseCoupon:        pulumi.String("true"),
//				AutoPay:              pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ga.NewBasicEndpointGroup(ctx, "defaultBasicEndpointGroup", &ga.BasicEndpointGroupArgs{
//				AcceleratorId:          defaultBasicAccelerator.ID(),
//				EndpointGroupRegion:    pulumi.String(endpointGroupRegion),
//				EndpointType:           pulumi.String("SLB"),
//				EndpointAddress:        defaultApplicationLoadBalancer.ID(),
//				EndpointSubAddress:     pulumi.String("192.168.0.1"),
//				BasicEndpointGroupName: pulumi.String("terraform-example"),
//				Description:            pulumi.String("terraform-example"),
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
// Global Accelerator (GA) Basic Endpoint Group can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ga/basicEndpointGroup:BasicEndpointGroup example <id>
// ```
type BasicEndpointGroup struct {
	pulumi.CustomResourceState

	// The ID of the basic GA instance.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The name of the endpoint group. The `basicEndpointGroupName` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	BasicEndpointGroupName pulumi.StringPtrOutput `pulumi:"basicEndpointGroupName"`
	// The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The address of the endpoint.
	EndpointAddress pulumi.StringOutput `pulumi:"endpointAddress"`
	// The ID of the region where you want to create the endpoint group.
	EndpointGroupRegion pulumi.StringOutput `pulumi:"endpointGroupRegion"`
	// The sub address of the endpoint.
	EndpointSubAddress pulumi.StringOutput `pulumi:"endpointSubAddress"`
	// The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
	EndpointType pulumi.StringOutput `pulumi:"endpointType"`
	// The status of the Basic Endpoint Group.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewBasicEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewBasicEndpointGroup(ctx *pulumi.Context,
	name string, args *BasicEndpointGroupArgs, opts ...pulumi.ResourceOption) (*BasicEndpointGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	if args.EndpointGroupRegion == nil {
		return nil, errors.New("invalid value for required argument 'EndpointGroupRegion'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BasicEndpointGroup
	err := ctx.RegisterResource("alicloud:ga/basicEndpointGroup:BasicEndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBasicEndpointGroup gets an existing BasicEndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBasicEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BasicEndpointGroupState, opts ...pulumi.ResourceOption) (*BasicEndpointGroup, error) {
	var resource BasicEndpointGroup
	err := ctx.ReadResource("alicloud:ga/basicEndpointGroup:BasicEndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BasicEndpointGroup resources.
type basicEndpointGroupState struct {
	// The ID of the basic GA instance.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The name of the endpoint group. The `basicEndpointGroupName` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	BasicEndpointGroupName *string `pulumi:"basicEndpointGroupName"`
	// The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
	Description *string `pulumi:"description"`
	// The address of the endpoint.
	EndpointAddress *string `pulumi:"endpointAddress"`
	// The ID of the region where you want to create the endpoint group.
	EndpointGroupRegion *string `pulumi:"endpointGroupRegion"`
	// The sub address of the endpoint.
	EndpointSubAddress *string `pulumi:"endpointSubAddress"`
	// The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
	EndpointType *string `pulumi:"endpointType"`
	// The status of the Basic Endpoint Group.
	Status *string `pulumi:"status"`
}

type BasicEndpointGroupState struct {
	// The ID of the basic GA instance.
	AcceleratorId pulumi.StringPtrInput
	// The name of the endpoint group. The `basicEndpointGroupName` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	BasicEndpointGroupName pulumi.StringPtrInput
	// The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
	Description pulumi.StringPtrInput
	// The address of the endpoint.
	EndpointAddress pulumi.StringPtrInput
	// The ID of the region where you want to create the endpoint group.
	EndpointGroupRegion pulumi.StringPtrInput
	// The sub address of the endpoint.
	EndpointSubAddress pulumi.StringPtrInput
	// The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
	EndpointType pulumi.StringPtrInput
	// The status of the Basic Endpoint Group.
	Status pulumi.StringPtrInput
}

func (BasicEndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*basicEndpointGroupState)(nil)).Elem()
}

type basicEndpointGroupArgs struct {
	// The ID of the basic GA instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The name of the endpoint group. The `basicEndpointGroupName` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	BasicEndpointGroupName *string `pulumi:"basicEndpointGroupName"`
	// The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
	Description *string `pulumi:"description"`
	// The address of the endpoint.
	EndpointAddress *string `pulumi:"endpointAddress"`
	// The ID of the region where you want to create the endpoint group.
	EndpointGroupRegion string `pulumi:"endpointGroupRegion"`
	// The sub address of the endpoint.
	EndpointSubAddress *string `pulumi:"endpointSubAddress"`
	// The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
	EndpointType *string `pulumi:"endpointType"`
}

// The set of arguments for constructing a BasicEndpointGroup resource.
type BasicEndpointGroupArgs struct {
	// The ID of the basic GA instance.
	AcceleratorId pulumi.StringInput
	// The name of the endpoint group. The `basicEndpointGroupName` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	BasicEndpointGroupName pulumi.StringPtrInput
	// The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
	Description pulumi.StringPtrInput
	// The address of the endpoint.
	EndpointAddress pulumi.StringPtrInput
	// The ID of the region where you want to create the endpoint group.
	EndpointGroupRegion pulumi.StringInput
	// The sub address of the endpoint.
	EndpointSubAddress pulumi.StringPtrInput
	// The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
	EndpointType pulumi.StringPtrInput
}

func (BasicEndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*basicEndpointGroupArgs)(nil)).Elem()
}

type BasicEndpointGroupInput interface {
	pulumi.Input

	ToBasicEndpointGroupOutput() BasicEndpointGroupOutput
	ToBasicEndpointGroupOutputWithContext(ctx context.Context) BasicEndpointGroupOutput
}

func (*BasicEndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicEndpointGroup)(nil)).Elem()
}

func (i *BasicEndpointGroup) ToBasicEndpointGroupOutput() BasicEndpointGroupOutput {
	return i.ToBasicEndpointGroupOutputWithContext(context.Background())
}

func (i *BasicEndpointGroup) ToBasicEndpointGroupOutputWithContext(ctx context.Context) BasicEndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicEndpointGroupOutput)
}

// BasicEndpointGroupArrayInput is an input type that accepts BasicEndpointGroupArray and BasicEndpointGroupArrayOutput values.
// You can construct a concrete instance of `BasicEndpointGroupArrayInput` via:
//
//	BasicEndpointGroupArray{ BasicEndpointGroupArgs{...} }
type BasicEndpointGroupArrayInput interface {
	pulumi.Input

	ToBasicEndpointGroupArrayOutput() BasicEndpointGroupArrayOutput
	ToBasicEndpointGroupArrayOutputWithContext(context.Context) BasicEndpointGroupArrayOutput
}

type BasicEndpointGroupArray []BasicEndpointGroupInput

func (BasicEndpointGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicEndpointGroup)(nil)).Elem()
}

func (i BasicEndpointGroupArray) ToBasicEndpointGroupArrayOutput() BasicEndpointGroupArrayOutput {
	return i.ToBasicEndpointGroupArrayOutputWithContext(context.Background())
}

func (i BasicEndpointGroupArray) ToBasicEndpointGroupArrayOutputWithContext(ctx context.Context) BasicEndpointGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicEndpointGroupArrayOutput)
}

// BasicEndpointGroupMapInput is an input type that accepts BasicEndpointGroupMap and BasicEndpointGroupMapOutput values.
// You can construct a concrete instance of `BasicEndpointGroupMapInput` via:
//
//	BasicEndpointGroupMap{ "key": BasicEndpointGroupArgs{...} }
type BasicEndpointGroupMapInput interface {
	pulumi.Input

	ToBasicEndpointGroupMapOutput() BasicEndpointGroupMapOutput
	ToBasicEndpointGroupMapOutputWithContext(context.Context) BasicEndpointGroupMapOutput
}

type BasicEndpointGroupMap map[string]BasicEndpointGroupInput

func (BasicEndpointGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicEndpointGroup)(nil)).Elem()
}

func (i BasicEndpointGroupMap) ToBasicEndpointGroupMapOutput() BasicEndpointGroupMapOutput {
	return i.ToBasicEndpointGroupMapOutputWithContext(context.Background())
}

func (i BasicEndpointGroupMap) ToBasicEndpointGroupMapOutputWithContext(ctx context.Context) BasicEndpointGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicEndpointGroupMapOutput)
}

type BasicEndpointGroupOutput struct{ *pulumi.OutputState }

func (BasicEndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicEndpointGroup)(nil)).Elem()
}

func (o BasicEndpointGroupOutput) ToBasicEndpointGroupOutput() BasicEndpointGroupOutput {
	return o
}

func (o BasicEndpointGroupOutput) ToBasicEndpointGroupOutputWithContext(ctx context.Context) BasicEndpointGroupOutput {
	return o
}

// The ID of the basic GA instance.
func (o BasicEndpointGroupOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpointGroup) pulumi.StringOutput { return v.AcceleratorId }).(pulumi.StringOutput)
}

// The name of the endpoint group. The `basicEndpointGroupName` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
func (o BasicEndpointGroupOutput) BasicEndpointGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicEndpointGroup) pulumi.StringPtrOutput { return v.BasicEndpointGroupName }).(pulumi.StringPtrOutput)
}

// The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
func (o BasicEndpointGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicEndpointGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The address of the endpoint.
func (o BasicEndpointGroupOutput) EndpointAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpointGroup) pulumi.StringOutput { return v.EndpointAddress }).(pulumi.StringOutput)
}

// The ID of the region where you want to create the endpoint group.
func (o BasicEndpointGroupOutput) EndpointGroupRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpointGroup) pulumi.StringOutput { return v.EndpointGroupRegion }).(pulumi.StringOutput)
}

// The sub address of the endpoint.
func (o BasicEndpointGroupOutput) EndpointSubAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpointGroup) pulumi.StringOutput { return v.EndpointSubAddress }).(pulumi.StringOutput)
}

// The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
func (o BasicEndpointGroupOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpointGroup) pulumi.StringOutput { return v.EndpointType }).(pulumi.StringOutput)
}

// The status of the Basic Endpoint Group.
func (o BasicEndpointGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicEndpointGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type BasicEndpointGroupArrayOutput struct{ *pulumi.OutputState }

func (BasicEndpointGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicEndpointGroup)(nil)).Elem()
}

func (o BasicEndpointGroupArrayOutput) ToBasicEndpointGroupArrayOutput() BasicEndpointGroupArrayOutput {
	return o
}

func (o BasicEndpointGroupArrayOutput) ToBasicEndpointGroupArrayOutputWithContext(ctx context.Context) BasicEndpointGroupArrayOutput {
	return o
}

func (o BasicEndpointGroupArrayOutput) Index(i pulumi.IntInput) BasicEndpointGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BasicEndpointGroup {
		return vs[0].([]*BasicEndpointGroup)[vs[1].(int)]
	}).(BasicEndpointGroupOutput)
}

type BasicEndpointGroupMapOutput struct{ *pulumi.OutputState }

func (BasicEndpointGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicEndpointGroup)(nil)).Elem()
}

func (o BasicEndpointGroupMapOutput) ToBasicEndpointGroupMapOutput() BasicEndpointGroupMapOutput {
	return o
}

func (o BasicEndpointGroupMapOutput) ToBasicEndpointGroupMapOutputWithContext(ctx context.Context) BasicEndpointGroupMapOutput {
	return o
}

func (o BasicEndpointGroupMapOutput) MapIndex(k pulumi.StringInput) BasicEndpointGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BasicEndpointGroup {
		return vs[0].(map[string]*BasicEndpointGroup)[vs[1].(string)]
	}).(BasicEndpointGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BasicEndpointGroupInput)(nil)).Elem(), &BasicEndpointGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicEndpointGroupArrayInput)(nil)).Elem(), BasicEndpointGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicEndpointGroupMapInput)(nil)).Elem(), BasicEndpointGroupMap{})
	pulumi.RegisterOutputType(BasicEndpointGroupOutput{})
	pulumi.RegisterOutputType(BasicEndpointGroupArrayOutput{})
	pulumi.RegisterOutputType(BasicEndpointGroupMapOutput{})
}
