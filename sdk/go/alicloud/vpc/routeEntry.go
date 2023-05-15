// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a route entry resource. A route entry represents a route item of one VPC route table.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(defaultZones.Zones[0].Id),
//				CpuCoreCount:     pulumi.IntRef(1),
//				MemorySize:       pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex:  pulumi.StringRef("^ubuntu_18.*64"),
//				MostRecent: pulumi.BoolRef(true),
//				Owners:     pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			cfg := config.New(ctx, "")
//			name := "RouteEntryConfig"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			fooNetwork, err := vpc.NewNetwork(ctx, "fooNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.1.0.0/21"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSwitch, err := vpc.NewSwitch(ctx, "fooSwitch", &vpc.SwitchArgs{
//				VpcId:       fooNetwork.ID(),
//				CidrBlock:   pulumi.String("10.1.1.0/24"),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			tfTestFoo, err := ecs.NewSecurityGroup(ctx, "tfTestFoo", &ecs.SecurityGroupArgs{
//				Description: pulumi.String("foo"),
//				VpcId:       fooNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewSecurityGroupRule(ctx, "ingress", &ecs.SecurityGroupRuleArgs{
//				Type:            pulumi.String("ingress"),
//				IpProtocol:      pulumi.String("tcp"),
//				NicType:         pulumi.String("intranet"),
//				Policy:          pulumi.String("accept"),
//				PortRange:       pulumi.String("22/22"),
//				Priority:        pulumi.Int(1),
//				SecurityGroupId: tfTestFoo.ID(),
//				CidrIp:          pulumi.String("0.0.0.0/0"),
//			})
//			if err != nil {
//				return err
//			}
//			fooInstance, err := ecs.NewInstance(ctx, "fooInstance", &ecs.InstanceArgs{
//				SecurityGroups: pulumi.StringArray{
//					tfTestFoo.ID(),
//				},
//				VswitchId:               fooSwitch.ID(),
//				InstanceChargeType:      pulumi.String("PostPaid"),
//				InstanceType:            *pulumi.String(defaultInstanceTypes.InstanceTypes[0].Id),
//				InternetChargeType:      pulumi.String("PayByTraffic"),
//				InternetMaxBandwidthOut: pulumi.Int(5),
//				SystemDiskCategory:      pulumi.String("cloud_efficiency"),
//				ImageId:                 *pulumi.String(defaultImages.Images[0].Id),
//				InstanceName:            pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewRouteEntry(ctx, "fooRouteEntry", &vpc.RouteEntryArgs{
//				RouteTableId:         fooNetwork.RouteTableId,
//				DestinationCidrblock: pulumi.String("172.11.1.1/32"),
//				NexthopType:          pulumi.String("Instance"),
//				NexthopId:            fooInstance.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Module Support
//
// You can use to the existing vpc module
// to create a VPC, several VSwitches and add several route entries one-click.
//
// ## Import
//
// Router entry can be imported using the id, e.g (formatted as<route_table_id:router_id:destination_cidrblock:nexthop_type:nexthop_id>).
//
// ```sh
//
//	$ pulumi import alicloud:vpc/routeEntry:RouteEntry example vtb-123456:vrt-123456:0.0.0.0/0:NatGateway:ngw-123456
//
// ```
type RouteEntry struct {
	pulumi.CustomResourceState

	// The RouteEntry's target network segment.
	DestinationCidrblock pulumi.StringPtrOutput `pulumi:"destinationCidrblock"`
	// The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	Name pulumi.StringOutput `pulumi:"name"`
	// The route entry's next hop. ECS instance ID or VPC router interface ID.
	NexthopId pulumi.StringPtrOutput `pulumi:"nexthopId"`
	// The next hop type. Available values:
	NexthopType pulumi.StringPtrOutput `pulumi:"nexthopType"`
	// The ID of the route table.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// This argument has been deprecated. Please use other arguments to launch a custom route entry.
	//
	// Deprecated: Attribute router_id has been deprecated and suggest removing it from your template.
	RouterId pulumi.StringOutput `pulumi:"routerId"`
}

// NewRouteEntry registers a new resource with the given unique name, arguments, and options.
func NewRouteEntry(ctx *pulumi.Context,
	name string, args *RouteEntryArgs, opts ...pulumi.ResourceOption) (*RouteEntry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableId'")
	}
	var resource RouteEntry
	err := ctx.RegisterResource("alicloud:vpc/routeEntry:RouteEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteEntry gets an existing RouteEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteEntryState, opts ...pulumi.ResourceOption) (*RouteEntry, error) {
	var resource RouteEntry
	err := ctx.ReadResource("alicloud:vpc/routeEntry:RouteEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteEntry resources.
type routeEntryState struct {
	// The RouteEntry's target network segment.
	DestinationCidrblock *string `pulumi:"destinationCidrblock"`
	// The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	Name *string `pulumi:"name"`
	// The route entry's next hop. ECS instance ID or VPC router interface ID.
	NexthopId *string `pulumi:"nexthopId"`
	// The next hop type. Available values:
	NexthopType *string `pulumi:"nexthopType"`
	// The ID of the route table.
	RouteTableId *string `pulumi:"routeTableId"`
	// This argument has been deprecated. Please use other arguments to launch a custom route entry.
	//
	// Deprecated: Attribute router_id has been deprecated and suggest removing it from your template.
	RouterId *string `pulumi:"routerId"`
}

type RouteEntryState struct {
	// The RouteEntry's target network segment.
	DestinationCidrblock pulumi.StringPtrInput
	// The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	Name pulumi.StringPtrInput
	// The route entry's next hop. ECS instance ID or VPC router interface ID.
	NexthopId pulumi.StringPtrInput
	// The next hop type. Available values:
	NexthopType pulumi.StringPtrInput
	// The ID of the route table.
	RouteTableId pulumi.StringPtrInput
	// This argument has been deprecated. Please use other arguments to launch a custom route entry.
	//
	// Deprecated: Attribute router_id has been deprecated and suggest removing it from your template.
	RouterId pulumi.StringPtrInput
}

func (RouteEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeEntryState)(nil)).Elem()
}

type routeEntryArgs struct {
	// The RouteEntry's target network segment.
	DestinationCidrblock *string `pulumi:"destinationCidrblock"`
	// The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	Name *string `pulumi:"name"`
	// The route entry's next hop. ECS instance ID or VPC router interface ID.
	NexthopId *string `pulumi:"nexthopId"`
	// The next hop type. Available values:
	NexthopType *string `pulumi:"nexthopType"`
	// The ID of the route table.
	RouteTableId string `pulumi:"routeTableId"`
	// This argument has been deprecated. Please use other arguments to launch a custom route entry.
	//
	// Deprecated: Attribute router_id has been deprecated and suggest removing it from your template.
	RouterId *string `pulumi:"routerId"`
}

// The set of arguments for constructing a RouteEntry resource.
type RouteEntryArgs struct {
	// The RouteEntry's target network segment.
	DestinationCidrblock pulumi.StringPtrInput
	// The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	Name pulumi.StringPtrInput
	// The route entry's next hop. ECS instance ID or VPC router interface ID.
	NexthopId pulumi.StringPtrInput
	// The next hop type. Available values:
	NexthopType pulumi.StringPtrInput
	// The ID of the route table.
	RouteTableId pulumi.StringInput
	// This argument has been deprecated. Please use other arguments to launch a custom route entry.
	//
	// Deprecated: Attribute router_id has been deprecated and suggest removing it from your template.
	RouterId pulumi.StringPtrInput
}

func (RouteEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeEntryArgs)(nil)).Elem()
}

type RouteEntryInput interface {
	pulumi.Input

	ToRouteEntryOutput() RouteEntryOutput
	ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput
}

func (*RouteEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteEntry)(nil)).Elem()
}

func (i *RouteEntry) ToRouteEntryOutput() RouteEntryOutput {
	return i.ToRouteEntryOutputWithContext(context.Background())
}

func (i *RouteEntry) ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryOutput)
}

// RouteEntryArrayInput is an input type that accepts RouteEntryArray and RouteEntryArrayOutput values.
// You can construct a concrete instance of `RouteEntryArrayInput` via:
//
//	RouteEntryArray{ RouteEntryArgs{...} }
type RouteEntryArrayInput interface {
	pulumi.Input

	ToRouteEntryArrayOutput() RouteEntryArrayOutput
	ToRouteEntryArrayOutputWithContext(context.Context) RouteEntryArrayOutput
}

type RouteEntryArray []RouteEntryInput

func (RouteEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteEntry)(nil)).Elem()
}

func (i RouteEntryArray) ToRouteEntryArrayOutput() RouteEntryArrayOutput {
	return i.ToRouteEntryArrayOutputWithContext(context.Background())
}

func (i RouteEntryArray) ToRouteEntryArrayOutputWithContext(ctx context.Context) RouteEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryArrayOutput)
}

// RouteEntryMapInput is an input type that accepts RouteEntryMap and RouteEntryMapOutput values.
// You can construct a concrete instance of `RouteEntryMapInput` via:
//
//	RouteEntryMap{ "key": RouteEntryArgs{...} }
type RouteEntryMapInput interface {
	pulumi.Input

	ToRouteEntryMapOutput() RouteEntryMapOutput
	ToRouteEntryMapOutputWithContext(context.Context) RouteEntryMapOutput
}

type RouteEntryMap map[string]RouteEntryInput

func (RouteEntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteEntry)(nil)).Elem()
}

func (i RouteEntryMap) ToRouteEntryMapOutput() RouteEntryMapOutput {
	return i.ToRouteEntryMapOutputWithContext(context.Background())
}

func (i RouteEntryMap) ToRouteEntryMapOutputWithContext(ctx context.Context) RouteEntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryMapOutput)
}

type RouteEntryOutput struct{ *pulumi.OutputState }

func (RouteEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteEntry)(nil)).Elem()
}

func (o RouteEntryOutput) ToRouteEntryOutput() RouteEntryOutput {
	return o
}

func (o RouteEntryOutput) ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput {
	return o
}

// The RouteEntry's target network segment.
func (o RouteEntryOutput) DestinationCidrblock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringPtrOutput { return v.DestinationCidrblock }).(pulumi.StringPtrOutput)
}

// The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
func (o RouteEntryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The route entry's next hop. ECS instance ID or VPC router interface ID.
func (o RouteEntryOutput) NexthopId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringPtrOutput { return v.NexthopId }).(pulumi.StringPtrOutput)
}

// The next hop type. Available values:
func (o RouteEntryOutput) NexthopType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringPtrOutput { return v.NexthopType }).(pulumi.StringPtrOutput)
}

// The ID of the route table.
func (o RouteEntryOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.RouteTableId }).(pulumi.StringOutput)
}

// This argument has been deprecated. Please use other arguments to launch a custom route entry.
//
// Deprecated: Attribute router_id has been deprecated and suggest removing it from your template.
func (o RouteEntryOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.RouterId }).(pulumi.StringOutput)
}

type RouteEntryArrayOutput struct{ *pulumi.OutputState }

func (RouteEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteEntry)(nil)).Elem()
}

func (o RouteEntryArrayOutput) ToRouteEntryArrayOutput() RouteEntryArrayOutput {
	return o
}

func (o RouteEntryArrayOutput) ToRouteEntryArrayOutputWithContext(ctx context.Context) RouteEntryArrayOutput {
	return o
}

func (o RouteEntryArrayOutput) Index(i pulumi.IntInput) RouteEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouteEntry {
		return vs[0].([]*RouteEntry)[vs[1].(int)]
	}).(RouteEntryOutput)
}

type RouteEntryMapOutput struct{ *pulumi.OutputState }

func (RouteEntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteEntry)(nil)).Elem()
}

func (o RouteEntryMapOutput) ToRouteEntryMapOutput() RouteEntryMapOutput {
	return o
}

func (o RouteEntryMapOutput) ToRouteEntryMapOutputWithContext(ctx context.Context) RouteEntryMapOutput {
	return o
}

func (o RouteEntryMapOutput) MapIndex(k pulumi.StringInput) RouteEntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouteEntry {
		return vs[0].(map[string]*RouteEntry)[vs[1].(string)]
	}).(RouteEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryInput)(nil)).Elem(), &RouteEntry{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryArrayInput)(nil)).Elem(), RouteEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryMapInput)(nil)).Elem(), RouteEntryMap{})
	pulumi.RegisterOutputType(RouteEntryOutput{})
	pulumi.RegisterOutputType(RouteEntryArrayOutput{})
	pulumi.RegisterOutputType(RouteEntryMapOutput{})
}
