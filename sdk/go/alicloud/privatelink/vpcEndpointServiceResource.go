// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Private Link Vpc Endpoint Service Resource resource. Endpoint service resource.
//
// For information about Private Link Vpc Endpoint Service Resource and how to use it, see [What is Vpc Endpoint Service Resource](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-attachresourcetovpcendpointservice).
//
// > **NOTE:** Available since v1.110.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/privatelink"
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
//			name := "tf_example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			exampleZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleVpcEndpointService, err := privatelink.NewVpcEndpointService(ctx, "exampleVpcEndpointService", &privatelink.VpcEndpointServiceArgs{
//				ServiceDescription:   pulumi.String(name),
//				ConnectBandwidth:     pulumi.Int(103),
//				AutoAcceptConnection: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			exampleNetwork, err := vpc.NewNetwork(ctx, "exampleNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.0.0.0/8"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSwitch, err := vpc.NewSwitch(ctx, "exampleSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.1.0.0/16"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      *pulumi.String(exampleZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSecurityGroup, err := ecs.NewSecurityGroup(ctx, "exampleSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: exampleNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			exampleApplicationLoadBalancer, err := slb.NewApplicationLoadBalancer(ctx, "exampleApplicationLoadBalancer", &slb.ApplicationLoadBalancerArgs{
//				LoadBalancerName: pulumi.String(name),
//				VswitchId:        exampleSwitch.ID(),
//				LoadBalancerSpec: pulumi.String("slb.s2.small"),
//				AddressType:      pulumi.String("intranet"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = privatelink.NewVpcEndpoint(ctx, "exampleVpcEndpoint", &privatelink.VpcEndpointArgs{
//				ServiceId: exampleVpcEndpointService.ID(),
//				SecurityGroupIds: pulumi.StringArray{
//					exampleSecurityGroup.ID(),
//				},
//				VpcId:           exampleNetwork.ID(),
//				VpcEndpointName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = privatelink.NewVpcEndpointServiceResource(ctx, "exampleVpcEndpointServiceResource", &privatelink.VpcEndpointServiceResourceArgs{
//				ServiceId:    exampleVpcEndpointService.ID(),
//				ResourceId:   exampleApplicationLoadBalancer.ID(),
//				ResourceType: pulumi.String("slb"),
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
// Private Link Vpc Endpoint Service Resource can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource example <service_id>:<resource_id>:<zone_id>
// ```
type VpcEndpointServiceResource struct {
	pulumi.CustomResourceState

	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The service resource ID.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Service resource type, value:
	// - **slb**: indicates that the service resource type is Classic Load Balancer (CLB).
	// - **alb**: indicates that the service resource type is Application Load Balancer (ALB).
	// - **nlb**: indicates that the service resource type is Network Load Balancer (NLB).
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// The endpoint service ID.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The ID of the zone to which the service resource belongs. (valid when the resource type is nlb/alb).
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewVpcEndpointServiceResource registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointServiceResource(ctx *pulumi.Context,
	name string, args *VpcEndpointServiceResourceArgs, opts ...pulumi.ResourceOption) (*VpcEndpointServiceResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcEndpointServiceResource
	err := ctx.RegisterResource("alicloud:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointServiceResource gets an existing VpcEndpointServiceResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointServiceResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointServiceResourceState, opts ...pulumi.ResourceOption) (*VpcEndpointServiceResource, error) {
	var resource VpcEndpointServiceResource
	err := ctx.ReadResource("alicloud:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointServiceResource resources.
type vpcEndpointServiceResourceState struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `pulumi:"dryRun"`
	// The service resource ID.
	ResourceId *string `pulumi:"resourceId"`
	// Service resource type, value:
	// - **slb**: indicates that the service resource type is Classic Load Balancer (CLB).
	// - **alb**: indicates that the service resource type is Application Load Balancer (ALB).
	// - **nlb**: indicates that the service resource type is Network Load Balancer (NLB).
	ResourceType *string `pulumi:"resourceType"`
	// The endpoint service ID.
	ServiceId *string `pulumi:"serviceId"`
	// The ID of the zone to which the service resource belongs. (valid when the resource type is nlb/alb).
	ZoneId *string `pulumi:"zoneId"`
}

type VpcEndpointServiceResourceState struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun pulumi.BoolPtrInput
	// The service resource ID.
	ResourceId pulumi.StringPtrInput
	// Service resource type, value:
	// - **slb**: indicates that the service resource type is Classic Load Balancer (CLB).
	// - **alb**: indicates that the service resource type is Application Load Balancer (ALB).
	// - **nlb**: indicates that the service resource type is Network Load Balancer (NLB).
	ResourceType pulumi.StringPtrInput
	// The endpoint service ID.
	ServiceId pulumi.StringPtrInput
	// The ID of the zone to which the service resource belongs. (valid when the resource type is nlb/alb).
	ZoneId pulumi.StringPtrInput
}

func (VpcEndpointServiceResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceResourceState)(nil)).Elem()
}

type vpcEndpointServiceResourceArgs struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `pulumi:"dryRun"`
	// The service resource ID.
	ResourceId string `pulumi:"resourceId"`
	// Service resource type, value:
	// - **slb**: indicates that the service resource type is Classic Load Balancer (CLB).
	// - **alb**: indicates that the service resource type is Application Load Balancer (ALB).
	// - **nlb**: indicates that the service resource type is Network Load Balancer (NLB).
	ResourceType string `pulumi:"resourceType"`
	// The endpoint service ID.
	ServiceId string `pulumi:"serviceId"`
	// The ID of the zone to which the service resource belongs. (valid when the resource type is nlb/alb).
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a VpcEndpointServiceResource resource.
type VpcEndpointServiceResourceArgs struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun pulumi.BoolPtrInput
	// The service resource ID.
	ResourceId pulumi.StringInput
	// Service resource type, value:
	// - **slb**: indicates that the service resource type is Classic Load Balancer (CLB).
	// - **alb**: indicates that the service resource type is Application Load Balancer (ALB).
	// - **nlb**: indicates that the service resource type is Network Load Balancer (NLB).
	ResourceType pulumi.StringInput
	// The endpoint service ID.
	ServiceId pulumi.StringInput
	// The ID of the zone to which the service resource belongs. (valid when the resource type is nlb/alb).
	ZoneId pulumi.StringPtrInput
}

func (VpcEndpointServiceResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceResourceArgs)(nil)).Elem()
}

type VpcEndpointServiceResourceInput interface {
	pulumi.Input

	ToVpcEndpointServiceResourceOutput() VpcEndpointServiceResourceOutput
	ToVpcEndpointServiceResourceOutputWithContext(ctx context.Context) VpcEndpointServiceResourceOutput
}

func (*VpcEndpointServiceResource) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServiceResource)(nil)).Elem()
}

func (i *VpcEndpointServiceResource) ToVpcEndpointServiceResourceOutput() VpcEndpointServiceResourceOutput {
	return i.ToVpcEndpointServiceResourceOutputWithContext(context.Background())
}

func (i *VpcEndpointServiceResource) ToVpcEndpointServiceResourceOutputWithContext(ctx context.Context) VpcEndpointServiceResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceResourceOutput)
}

// VpcEndpointServiceResourceArrayInput is an input type that accepts VpcEndpointServiceResourceArray and VpcEndpointServiceResourceArrayOutput values.
// You can construct a concrete instance of `VpcEndpointServiceResourceArrayInput` via:
//
//	VpcEndpointServiceResourceArray{ VpcEndpointServiceResourceArgs{...} }
type VpcEndpointServiceResourceArrayInput interface {
	pulumi.Input

	ToVpcEndpointServiceResourceArrayOutput() VpcEndpointServiceResourceArrayOutput
	ToVpcEndpointServiceResourceArrayOutputWithContext(context.Context) VpcEndpointServiceResourceArrayOutput
}

type VpcEndpointServiceResourceArray []VpcEndpointServiceResourceInput

func (VpcEndpointServiceResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointServiceResource)(nil)).Elem()
}

func (i VpcEndpointServiceResourceArray) ToVpcEndpointServiceResourceArrayOutput() VpcEndpointServiceResourceArrayOutput {
	return i.ToVpcEndpointServiceResourceArrayOutputWithContext(context.Background())
}

func (i VpcEndpointServiceResourceArray) ToVpcEndpointServiceResourceArrayOutputWithContext(ctx context.Context) VpcEndpointServiceResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceResourceArrayOutput)
}

// VpcEndpointServiceResourceMapInput is an input type that accepts VpcEndpointServiceResourceMap and VpcEndpointServiceResourceMapOutput values.
// You can construct a concrete instance of `VpcEndpointServiceResourceMapInput` via:
//
//	VpcEndpointServiceResourceMap{ "key": VpcEndpointServiceResourceArgs{...} }
type VpcEndpointServiceResourceMapInput interface {
	pulumi.Input

	ToVpcEndpointServiceResourceMapOutput() VpcEndpointServiceResourceMapOutput
	ToVpcEndpointServiceResourceMapOutputWithContext(context.Context) VpcEndpointServiceResourceMapOutput
}

type VpcEndpointServiceResourceMap map[string]VpcEndpointServiceResourceInput

func (VpcEndpointServiceResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointServiceResource)(nil)).Elem()
}

func (i VpcEndpointServiceResourceMap) ToVpcEndpointServiceResourceMapOutput() VpcEndpointServiceResourceMapOutput {
	return i.ToVpcEndpointServiceResourceMapOutputWithContext(context.Background())
}

func (i VpcEndpointServiceResourceMap) ToVpcEndpointServiceResourceMapOutputWithContext(ctx context.Context) VpcEndpointServiceResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceResourceMapOutput)
}

type VpcEndpointServiceResourceOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServiceResource)(nil)).Elem()
}

func (o VpcEndpointServiceResourceOutput) ToVpcEndpointServiceResourceOutput() VpcEndpointServiceResourceOutput {
	return o
}

func (o VpcEndpointServiceResourceOutput) ToVpcEndpointServiceResourceOutputWithContext(ctx context.Context) VpcEndpointServiceResourceOutput {
	return o
}

// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
func (o VpcEndpointServiceResourceOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcEndpointServiceResource) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The service resource ID.
func (o VpcEndpointServiceResourceOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceResource) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Service resource type, value:
// - **slb**: indicates that the service resource type is Classic Load Balancer (CLB).
// - **alb**: indicates that the service resource type is Application Load Balancer (ALB).
// - **nlb**: indicates that the service resource type is Network Load Balancer (NLB).
func (o VpcEndpointServiceResourceOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceResource) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// The endpoint service ID.
func (o VpcEndpointServiceResourceOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceResource) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The ID of the zone to which the service resource belongs. (valid when the resource type is nlb/alb).
func (o VpcEndpointServiceResourceOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceResource) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type VpcEndpointServiceResourceArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointServiceResource)(nil)).Elem()
}

func (o VpcEndpointServiceResourceArrayOutput) ToVpcEndpointServiceResourceArrayOutput() VpcEndpointServiceResourceArrayOutput {
	return o
}

func (o VpcEndpointServiceResourceArrayOutput) ToVpcEndpointServiceResourceArrayOutputWithContext(ctx context.Context) VpcEndpointServiceResourceArrayOutput {
	return o
}

func (o VpcEndpointServiceResourceArrayOutput) Index(i pulumi.IntInput) VpcEndpointServiceResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointServiceResource {
		return vs[0].([]*VpcEndpointServiceResource)[vs[1].(int)]
	}).(VpcEndpointServiceResourceOutput)
}

type VpcEndpointServiceResourceMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointServiceResource)(nil)).Elem()
}

func (o VpcEndpointServiceResourceMapOutput) ToVpcEndpointServiceResourceMapOutput() VpcEndpointServiceResourceMapOutput {
	return o
}

func (o VpcEndpointServiceResourceMapOutput) ToVpcEndpointServiceResourceMapOutputWithContext(ctx context.Context) VpcEndpointServiceResourceMapOutput {
	return o
}

func (o VpcEndpointServiceResourceMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointServiceResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointServiceResource {
		return vs[0].(map[string]*VpcEndpointServiceResource)[vs[1].(string)]
	}).(VpcEndpointServiceResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceResourceInput)(nil)).Elem(), &VpcEndpointServiceResource{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceResourceArrayInput)(nil)).Elem(), VpcEndpointServiceResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceResourceMapInput)(nil)).Elem(), VpcEndpointServiceResourceMap{})
	pulumi.RegisterOutputType(VpcEndpointServiceResourceOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceResourceArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceResourceMapOutput{})
}
