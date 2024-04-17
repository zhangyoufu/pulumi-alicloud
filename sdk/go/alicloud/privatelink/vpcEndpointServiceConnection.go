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

// Provides a Private Link Vpc Endpoint Connection resource. vpc endpoint connection.
//
// For information about Private Link Vpc Endpoint Connection and how to use it, see [What is Vpc Endpoint Connection](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-enablevpcendpointzoneconnection).
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
//			example, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleVpcEndpointService, err := privatelink.NewVpcEndpointService(ctx, "example", &privatelink.VpcEndpointServiceArgs{
//				ServiceDescription:   pulumi.String(name),
//				ConnectBandwidth:     pulumi.Int(103),
//				AutoAcceptConnection: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			exampleNetwork, err := vpc.NewNetwork(ctx, "example", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.0.0.0/8"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSwitch, err := vpc.NewSwitch(ctx, "example", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.1.0.0/16"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      pulumi.String(example.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSecurityGroup, err := ecs.NewSecurityGroup(ctx, "example", &ecs.SecurityGroupArgs{
//				Name:  pulumi.String(name),
//				VpcId: exampleNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			exampleApplicationLoadBalancer, err := slb.NewApplicationLoadBalancer(ctx, "example", &slb.ApplicationLoadBalancerArgs{
//				LoadBalancerName: pulumi.String(name),
//				VswitchId:        exampleSwitch.ID(),
//				LoadBalancerSpec: pulumi.String("slb.s2.small"),
//				AddressType:      pulumi.String("intranet"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleVpcEndpointServiceResource, err := privatelink.NewVpcEndpointServiceResource(ctx, "example", &privatelink.VpcEndpointServiceResourceArgs{
//				ServiceId:    exampleVpcEndpointService.ID(),
//				ResourceId:   exampleApplicationLoadBalancer.ID(),
//				ResourceType: pulumi.String("slb"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleVpcEndpoint, err := privatelink.NewVpcEndpoint(ctx, "example", &privatelink.VpcEndpointArgs{
//				ServiceId: exampleVpcEndpointServiceResource.ServiceId,
//				SecurityGroupIds: pulumi.StringArray{
//					exampleSecurityGroup.ID(),
//				},
//				VpcId:           exampleNetwork.ID(),
//				VpcEndpointName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = privatelink.NewVpcEndpointServiceConnection(ctx, "example", &privatelink.VpcEndpointServiceConnectionArgs{
//				EndpointId: exampleVpcEndpoint.ID(),
//				ServiceId:  exampleVpcEndpoint.ServiceId,
//				Bandwidth:  pulumi.Int(1024),
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
// Private Link Vpc Endpoint Connection can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection example <service_id>:<endpoint_id>
// ```
type VpcEndpointServiceConnection struct {
	pulumi.CustomResourceState

	// The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The endpoint ID.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// The endpoint service ID.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The state of the endpoint connection.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewVpcEndpointServiceConnection registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointServiceConnection(ctx *pulumi.Context,
	name string, args *VpcEndpointServiceConnectionArgs, opts ...pulumi.ResourceOption) (*VpcEndpointServiceConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointId'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcEndpointServiceConnection
	err := ctx.RegisterResource("alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointServiceConnection gets an existing VpcEndpointServiceConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointServiceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointServiceConnectionState, opts ...pulumi.ResourceOption) (*VpcEndpointServiceConnection, error) {
	var resource VpcEndpointServiceConnection
	err := ctx.ReadResource("alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointServiceConnection resources.
type vpcEndpointServiceConnectionState struct {
	// The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
	Bandwidth *int `pulumi:"bandwidth"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `pulumi:"dryRun"`
	// The endpoint ID.
	EndpointId *string `pulumi:"endpointId"`
	// The endpoint service ID.
	ServiceId *string `pulumi:"serviceId"`
	// The state of the endpoint connection.
	Status *string `pulumi:"status"`
}

type VpcEndpointServiceConnectionState struct {
	// The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
	Bandwidth pulumi.IntPtrInput
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun pulumi.BoolPtrInput
	// The endpoint ID.
	EndpointId pulumi.StringPtrInput
	// The endpoint service ID.
	ServiceId pulumi.StringPtrInput
	// The state of the endpoint connection.
	Status pulumi.StringPtrInput
}

func (VpcEndpointServiceConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceConnectionState)(nil)).Elem()
}

type vpcEndpointServiceConnectionArgs struct {
	// The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
	Bandwidth *int `pulumi:"bandwidth"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `pulumi:"dryRun"`
	// The endpoint ID.
	EndpointId string `pulumi:"endpointId"`
	// The endpoint service ID.
	ServiceId string `pulumi:"serviceId"`
}

// The set of arguments for constructing a VpcEndpointServiceConnection resource.
type VpcEndpointServiceConnectionArgs struct {
	// The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
	Bandwidth pulumi.IntPtrInput
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun pulumi.BoolPtrInput
	// The endpoint ID.
	EndpointId pulumi.StringInput
	// The endpoint service ID.
	ServiceId pulumi.StringInput
}

func (VpcEndpointServiceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceConnectionArgs)(nil)).Elem()
}

type VpcEndpointServiceConnectionInput interface {
	pulumi.Input

	ToVpcEndpointServiceConnectionOutput() VpcEndpointServiceConnectionOutput
	ToVpcEndpointServiceConnectionOutputWithContext(ctx context.Context) VpcEndpointServiceConnectionOutput
}

func (*VpcEndpointServiceConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServiceConnection)(nil)).Elem()
}

func (i *VpcEndpointServiceConnection) ToVpcEndpointServiceConnectionOutput() VpcEndpointServiceConnectionOutput {
	return i.ToVpcEndpointServiceConnectionOutputWithContext(context.Background())
}

func (i *VpcEndpointServiceConnection) ToVpcEndpointServiceConnectionOutputWithContext(ctx context.Context) VpcEndpointServiceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceConnectionOutput)
}

// VpcEndpointServiceConnectionArrayInput is an input type that accepts VpcEndpointServiceConnectionArray and VpcEndpointServiceConnectionArrayOutput values.
// You can construct a concrete instance of `VpcEndpointServiceConnectionArrayInput` via:
//
//	VpcEndpointServiceConnectionArray{ VpcEndpointServiceConnectionArgs{...} }
type VpcEndpointServiceConnectionArrayInput interface {
	pulumi.Input

	ToVpcEndpointServiceConnectionArrayOutput() VpcEndpointServiceConnectionArrayOutput
	ToVpcEndpointServiceConnectionArrayOutputWithContext(context.Context) VpcEndpointServiceConnectionArrayOutput
}

type VpcEndpointServiceConnectionArray []VpcEndpointServiceConnectionInput

func (VpcEndpointServiceConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointServiceConnection)(nil)).Elem()
}

func (i VpcEndpointServiceConnectionArray) ToVpcEndpointServiceConnectionArrayOutput() VpcEndpointServiceConnectionArrayOutput {
	return i.ToVpcEndpointServiceConnectionArrayOutputWithContext(context.Background())
}

func (i VpcEndpointServiceConnectionArray) ToVpcEndpointServiceConnectionArrayOutputWithContext(ctx context.Context) VpcEndpointServiceConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceConnectionArrayOutput)
}

// VpcEndpointServiceConnectionMapInput is an input type that accepts VpcEndpointServiceConnectionMap and VpcEndpointServiceConnectionMapOutput values.
// You can construct a concrete instance of `VpcEndpointServiceConnectionMapInput` via:
//
//	VpcEndpointServiceConnectionMap{ "key": VpcEndpointServiceConnectionArgs{...} }
type VpcEndpointServiceConnectionMapInput interface {
	pulumi.Input

	ToVpcEndpointServiceConnectionMapOutput() VpcEndpointServiceConnectionMapOutput
	ToVpcEndpointServiceConnectionMapOutputWithContext(context.Context) VpcEndpointServiceConnectionMapOutput
}

type VpcEndpointServiceConnectionMap map[string]VpcEndpointServiceConnectionInput

func (VpcEndpointServiceConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointServiceConnection)(nil)).Elem()
}

func (i VpcEndpointServiceConnectionMap) ToVpcEndpointServiceConnectionMapOutput() VpcEndpointServiceConnectionMapOutput {
	return i.ToVpcEndpointServiceConnectionMapOutputWithContext(context.Background())
}

func (i VpcEndpointServiceConnectionMap) ToVpcEndpointServiceConnectionMapOutputWithContext(ctx context.Context) VpcEndpointServiceConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceConnectionMapOutput)
}

type VpcEndpointServiceConnectionOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServiceConnection)(nil)).Elem()
}

func (o VpcEndpointServiceConnectionOutput) ToVpcEndpointServiceConnectionOutput() VpcEndpointServiceConnectionOutput {
	return o
}

func (o VpcEndpointServiceConnectionOutput) ToVpcEndpointServiceConnectionOutputWithContext(ctx context.Context) VpcEndpointServiceConnectionOutput {
	return o
}

// The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
func (o VpcEndpointServiceConnectionOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcEndpointServiceConnection) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
func (o VpcEndpointServiceConnectionOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcEndpointServiceConnection) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The endpoint ID.
func (o VpcEndpointServiceConnectionOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceConnection) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

// The endpoint service ID.
func (o VpcEndpointServiceConnectionOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceConnection) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The state of the endpoint connection.
func (o VpcEndpointServiceConnectionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceConnection) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type VpcEndpointServiceConnectionArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointServiceConnection)(nil)).Elem()
}

func (o VpcEndpointServiceConnectionArrayOutput) ToVpcEndpointServiceConnectionArrayOutput() VpcEndpointServiceConnectionArrayOutput {
	return o
}

func (o VpcEndpointServiceConnectionArrayOutput) ToVpcEndpointServiceConnectionArrayOutputWithContext(ctx context.Context) VpcEndpointServiceConnectionArrayOutput {
	return o
}

func (o VpcEndpointServiceConnectionArrayOutput) Index(i pulumi.IntInput) VpcEndpointServiceConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointServiceConnection {
		return vs[0].([]*VpcEndpointServiceConnection)[vs[1].(int)]
	}).(VpcEndpointServiceConnectionOutput)
}

type VpcEndpointServiceConnectionMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointServiceConnection)(nil)).Elem()
}

func (o VpcEndpointServiceConnectionMapOutput) ToVpcEndpointServiceConnectionMapOutput() VpcEndpointServiceConnectionMapOutput {
	return o
}

func (o VpcEndpointServiceConnectionMapOutput) ToVpcEndpointServiceConnectionMapOutputWithContext(ctx context.Context) VpcEndpointServiceConnectionMapOutput {
	return o
}

func (o VpcEndpointServiceConnectionMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointServiceConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointServiceConnection {
		return vs[0].(map[string]*VpcEndpointServiceConnection)[vs[1].(string)]
	}).(VpcEndpointServiceConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceConnectionInput)(nil)).Elem(), &VpcEndpointServiceConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceConnectionArrayInput)(nil)).Elem(), VpcEndpointServiceConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceConnectionMapInput)(nil)).Elem(), VpcEndpointServiceConnectionMap{})
	pulumi.RegisterOutputType(VpcEndpointServiceConnectionOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceConnectionArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceConnectionMapOutput{})
}
