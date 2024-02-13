// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an ALB Server Group resource.
//
// For information about ALB Server Group and how to use it, see [What is Server Group](https://www.alibabacloud.com/help/en/slb/application-load-balancer/developer-reference/api-alb-2020-06-16-createservergroup).
//
// > **NOTE:** Available since v1.131.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/alb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
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
//			exampleResourceGroups, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			exampleZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("Instance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(exampleZones.Zones[0].Id),
//				CpuCoreCount:     pulumi.IntRef(1),
//				MemorySize:       pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex: pulumi.StringRef("^ubuntu_[0-9]+_[0-9]+_x64*"),
//				Owners:    pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleNetwork, err := vpc.NewNetwork(ctx, "exampleNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSwitch, err := vpc.NewSwitch(ctx, "exampleSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/16"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      *pulumi.String(exampleZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSecurityGroup, err := ecs.NewSecurityGroup(ctx, "exampleSecurityGroup", &ecs.SecurityGroupArgs{
//				Description: pulumi.String(name),
//				VpcId:       exampleNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			exampleInstance, err := ecs.NewInstance(ctx, "exampleInstance", &ecs.InstanceArgs{
//				AvailabilityZone: *pulumi.String(exampleZones.Zones[0].Id),
//				InstanceName:     pulumi.String(name),
//				ImageId:          *pulumi.String(exampleImages.Images[0].Id),
//				InstanceType:     *pulumi.String(exampleInstanceTypes.InstanceTypes[0].Id),
//				SecurityGroups: pulumi.StringArray{
//					exampleSecurityGroup.ID(),
//				},
//				VswitchId: exampleSwitch.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = alb.NewServerGroup(ctx, "exampleServerGroup", &alb.ServerGroupArgs{
//				Protocol:        pulumi.String("HTTP"),
//				VpcId:           exampleNetwork.ID(),
//				ServerGroupName: pulumi.String(name),
//				ResourceGroupId: *pulumi.String(exampleResourceGroups.Groups[0].Id),
//				StickySessionConfig: &alb.ServerGroupStickySessionConfigArgs{
//					StickySessionEnabled: pulumi.Bool(true),
//					Cookie:               pulumi.String("tf-example"),
//					StickySessionType:    pulumi.String("Server"),
//				},
//				HealthCheckConfig: &alb.ServerGroupHealthCheckConfigArgs{
//					HealthCheckConnectPort: pulumi.Int(46325),
//					HealthCheckEnabled:     pulumi.Bool(true),
//					HealthCheckHost:        pulumi.String("tf-example.com"),
//					HealthCheckCodes: pulumi.StringArray{
//						pulumi.String("http_2xx"),
//						pulumi.String("http_3xx"),
//						pulumi.String("http_4xx"),
//					},
//					HealthCheckHttpVersion: pulumi.String("HTTP1.1"),
//					HealthCheckInterval:    pulumi.Int(2),
//					HealthCheckMethod:      pulumi.String("HEAD"),
//					HealthCheckPath:        pulumi.String("/tf-example"),
//					HealthCheckProtocol:    pulumi.String("HTTP"),
//					HealthCheckTimeout:     pulumi.Int(5),
//					HealthyThreshold:       pulumi.Int(3),
//					UnhealthyThreshold:     pulumi.Int(3),
//				},
//				Servers: alb.ServerGroupServerArray{
//					&alb.ServerGroupServerArgs{
//						Description: pulumi.String(name),
//						Port:        pulumi.Int(80),
//						ServerId:    exampleInstance.ID(),
//						ServerIp:    exampleInstance.PrivateIp,
//						ServerType:  pulumi.String("Ecs"),
//						Weight:      pulumi.Int(10),
//					},
//				},
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//				},
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
// ALB Server Group can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:alb/serverGroup:ServerGroup example <id>
// ```
type ServerGroup struct {
	pulumi.CustomResourceState

	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The configuration of health checks. See `healthCheckConfig` below.
	HealthCheckConfig ServerGroupHealthCheckConfigOutput `pulumi:"healthCheckConfig"`
	// The server protocol. Valid values: `  HTTP `, `HTTPS`, `gRPC`. While `serverGroupType` is `Fc` this parameter will not take effect. From version 1.215.0, `protocol` can be set to `gRPC`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
	Scheduler pulumi.StringOutput `pulumi:"scheduler"`
	// The name of the server group.
	ServerGroupName pulumi.StringOutput `pulumi:"serverGroupName"`
	// The type of the server group. Default value: `Instance`. Valid values:
	ServerGroupType pulumi.StringOutput `pulumi:"serverGroupType"`
	// The backend servers. See `servers` below.
	Servers ServerGroupServerArrayOutput `pulumi:"servers"`
	// The status of the backend server.
	Status pulumi.StringOutput `pulumi:"status"`
	// The configuration of session persistence. See `stickySessionConfig` below.
	StickySessionConfig ServerGroupStickySessionConfigOutput `pulumi:"stickySessionConfig"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the VPC that you want to access. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewServerGroup registers a new resource with the given unique name, arguments, and options.
func NewServerGroup(ctx *pulumi.Context,
	name string, args *ServerGroupArgs, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HealthCheckConfig == nil {
		return nil, errors.New("invalid value for required argument 'HealthCheckConfig'")
	}
	if args.ServerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ServerGroupName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServerGroup
	err := ctx.RegisterResource("alicloud:alb/serverGroup:ServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerGroup gets an existing ServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerGroupState, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	var resource ServerGroup
	err := ctx.ReadResource("alicloud:alb/serverGroup:ServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerGroup resources.
type serverGroupState struct {
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The configuration of health checks. See `healthCheckConfig` below.
	HealthCheckConfig *ServerGroupHealthCheckConfig `pulumi:"healthCheckConfig"`
	// The server protocol. Valid values: `  HTTP `, `HTTPS`, `gRPC`. While `serverGroupType` is `Fc` this parameter will not take effect. From version 1.215.0, `protocol` can be set to `gRPC`.
	Protocol *string `pulumi:"protocol"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
	Scheduler *string `pulumi:"scheduler"`
	// The name of the server group.
	ServerGroupName *string `pulumi:"serverGroupName"`
	// The type of the server group. Default value: `Instance`. Valid values:
	ServerGroupType *string `pulumi:"serverGroupType"`
	// The backend servers. See `servers` below.
	Servers []ServerGroupServer `pulumi:"servers"`
	// The status of the backend server.
	Status *string `pulumi:"status"`
	// The configuration of session persistence. See `stickySessionConfig` below.
	StickySessionConfig *ServerGroupStickySessionConfig `pulumi:"stickySessionConfig"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC that you want to access. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
	VpcId *string `pulumi:"vpcId"`
}

type ServerGroupState struct {
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The configuration of health checks. See `healthCheckConfig` below.
	HealthCheckConfig ServerGroupHealthCheckConfigPtrInput
	// The server protocol. Valid values: `  HTTP `, `HTTPS`, `gRPC`. While `serverGroupType` is `Fc` this parameter will not take effect. From version 1.215.0, `protocol` can be set to `gRPC`.
	Protocol pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
	Scheduler pulumi.StringPtrInput
	// The name of the server group.
	ServerGroupName pulumi.StringPtrInput
	// The type of the server group. Default value: `Instance`. Valid values:
	ServerGroupType pulumi.StringPtrInput
	// The backend servers. See `servers` below.
	Servers ServerGroupServerArrayInput
	// The status of the backend server.
	Status pulumi.StringPtrInput
	// The configuration of session persistence. See `stickySessionConfig` below.
	StickySessionConfig ServerGroupStickySessionConfigPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VPC that you want to access. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
	VpcId pulumi.StringPtrInput
}

func (ServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupState)(nil)).Elem()
}

type serverGroupArgs struct {
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The configuration of health checks. See `healthCheckConfig` below.
	HealthCheckConfig ServerGroupHealthCheckConfig `pulumi:"healthCheckConfig"`
	// The server protocol. Valid values: `  HTTP `, `HTTPS`, `gRPC`. While `serverGroupType` is `Fc` this parameter will not take effect. From version 1.215.0, `protocol` can be set to `gRPC`.
	Protocol *string `pulumi:"protocol"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
	Scheduler *string `pulumi:"scheduler"`
	// The name of the server group.
	ServerGroupName string `pulumi:"serverGroupName"`
	// The type of the server group. Default value: `Instance`. Valid values:
	ServerGroupType *string `pulumi:"serverGroupType"`
	// The backend servers. See `servers` below.
	Servers []ServerGroupServer `pulumi:"servers"`
	// The configuration of session persistence. See `stickySessionConfig` below.
	StickySessionConfig *ServerGroupStickySessionConfig `pulumi:"stickySessionConfig"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC that you want to access. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a ServerGroup resource.
type ServerGroupArgs struct {
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The configuration of health checks. See `healthCheckConfig` below.
	HealthCheckConfig ServerGroupHealthCheckConfigInput
	// The server protocol. Valid values: `  HTTP `, `HTTPS`, `gRPC`. While `serverGroupType` is `Fc` this parameter will not take effect. From version 1.215.0, `protocol` can be set to `gRPC`.
	Protocol pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
	Scheduler pulumi.StringPtrInput
	// The name of the server group.
	ServerGroupName pulumi.StringInput
	// The type of the server group. Default value: `Instance`. Valid values:
	ServerGroupType pulumi.StringPtrInput
	// The backend servers. See `servers` below.
	Servers ServerGroupServerArrayInput
	// The configuration of session persistence. See `stickySessionConfig` below.
	StickySessionConfig ServerGroupStickySessionConfigPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VPC that you want to access. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
	VpcId pulumi.StringPtrInput
}

func (ServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupArgs)(nil)).Elem()
}

type ServerGroupInput interface {
	pulumi.Input

	ToServerGroupOutput() ServerGroupOutput
	ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput
}

func (*ServerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (i *ServerGroup) ToServerGroupOutput() ServerGroupOutput {
	return i.ToServerGroupOutputWithContext(context.Background())
}

func (i *ServerGroup) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupOutput)
}

// ServerGroupArrayInput is an input type that accepts ServerGroupArray and ServerGroupArrayOutput values.
// You can construct a concrete instance of `ServerGroupArrayInput` via:
//
//	ServerGroupArray{ ServerGroupArgs{...} }
type ServerGroupArrayInput interface {
	pulumi.Input

	ToServerGroupArrayOutput() ServerGroupArrayOutput
	ToServerGroupArrayOutputWithContext(context.Context) ServerGroupArrayOutput
}

type ServerGroupArray []ServerGroupInput

func (ServerGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroup)(nil)).Elem()
}

func (i ServerGroupArray) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return i.ToServerGroupArrayOutputWithContext(context.Background())
}

func (i ServerGroupArray) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupArrayOutput)
}

// ServerGroupMapInput is an input type that accepts ServerGroupMap and ServerGroupMapOutput values.
// You can construct a concrete instance of `ServerGroupMapInput` via:
//
//	ServerGroupMap{ "key": ServerGroupArgs{...} }
type ServerGroupMapInput interface {
	pulumi.Input

	ToServerGroupMapOutput() ServerGroupMapOutput
	ToServerGroupMapOutputWithContext(context.Context) ServerGroupMapOutput
}

type ServerGroupMap map[string]ServerGroupInput

func (ServerGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroup)(nil)).Elem()
}

func (i ServerGroupMap) ToServerGroupMapOutput() ServerGroupMapOutput {
	return i.ToServerGroupMapOutputWithContext(context.Background())
}

func (i ServerGroupMap) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupMapOutput)
}

type ServerGroupOutput struct{ *pulumi.OutputState }

func (ServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (o ServerGroupOutput) ToServerGroupOutput() ServerGroupOutput {
	return o
}

func (o ServerGroupOutput) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return o
}

// The dry run.
func (o ServerGroupOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The configuration of health checks. See `healthCheckConfig` below.
func (o ServerGroupOutput) HealthCheckConfig() ServerGroupHealthCheckConfigOutput {
	return o.ApplyT(func(v *ServerGroup) ServerGroupHealthCheckConfigOutput { return v.HealthCheckConfig }).(ServerGroupHealthCheckConfigOutput)
}

// The server protocol. Valid values: `  HTTP `, `HTTPS`, `gRPC`. While `serverGroupType` is `Fc` this parameter will not take effect. From version 1.215.0, `protocol` can be set to `gRPC`.
func (o ServerGroupOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o ServerGroupOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
func (o ServerGroupOutput) Scheduler() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Scheduler }).(pulumi.StringOutput)
}

// The name of the server group.
func (o ServerGroupOutput) ServerGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.ServerGroupName }).(pulumi.StringOutput)
}

// The type of the server group. Default value: `Instance`. Valid values:
func (o ServerGroupOutput) ServerGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.ServerGroupType }).(pulumi.StringOutput)
}

// The backend servers. See `servers` below.
func (o ServerGroupOutput) Servers() ServerGroupServerArrayOutput {
	return o.ApplyT(func(v *ServerGroup) ServerGroupServerArrayOutput { return v.Servers }).(ServerGroupServerArrayOutput)
}

// The status of the backend server.
func (o ServerGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The configuration of session persistence. See `stickySessionConfig` below.
func (o ServerGroupOutput) StickySessionConfig() ServerGroupStickySessionConfigOutput {
	return o.ApplyT(func(v *ServerGroup) ServerGroupStickySessionConfigOutput { return v.StickySessionConfig }).(ServerGroupStickySessionConfigOutput)
}

// A mapping of tags to assign to the resource.
func (o ServerGroupOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The ID of the VPC that you want to access. **NOTE:** This parameter takes effect when the `serverGroupType` parameter is set to `Instance` or `Ip`.
func (o ServerGroupOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

type ServerGroupArrayOutput struct{ *pulumi.OutputState }

func (ServerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroup)(nil)).Elem()
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) Index(i pulumi.IntInput) ServerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerGroup {
		return vs[0].([]*ServerGroup)[vs[1].(int)]
	}).(ServerGroupOutput)
}

type ServerGroupMapOutput struct{ *pulumi.OutputState }

func (ServerGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroup)(nil)).Elem()
}

func (o ServerGroupMapOutput) ToServerGroupMapOutput() ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) MapIndex(k pulumi.StringInput) ServerGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerGroup {
		return vs[0].(map[string]*ServerGroup)[vs[1].(string)]
	}).(ServerGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupInput)(nil)).Elem(), &ServerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupArrayInput)(nil)).Elem(), ServerGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupMapInput)(nil)).Elem(), ServerGroupMap{})
	pulumi.RegisterOutputType(ServerGroupOutput{})
	pulumi.RegisterOutputType(ServerGroupArrayOutput{})
	pulumi.RegisterOutputType(ServerGroupMapOutput{})
}
