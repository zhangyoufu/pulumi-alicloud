// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package arms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Application Real-Time Monitoring Service (ARMS) Remote Write resource.
//
// For information about Application Real-Time Monitoring Service (ARMS) Remote Write and how to use it, see [What is Remote Write](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-addprometheusremotewrite).
//
// > **NOTE:** Available since v1.204.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/arms"
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
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      defaultZones.Zones[len(defaultZones.Zones)-1].Id,
//			})
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: defaultNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultPrometheus, err := arms.NewPrometheus(ctx, "defaultPrometheus", &arms.PrometheusArgs{
//				ClusterType:       pulumi.String("ecs"),
//				GrafanaInstanceId: pulumi.String("free"),
//				VpcId:             defaultNetwork.ID(),
//				VswitchId:         defaultSwitch.ID(),
//				SecurityGroupId:   defaultSecurityGroup.ID(),
//				ClusterName: defaultNetwork.ID().ApplyT(func(id string) (string, error) {
//					return fmt.Sprintf("%v-%v", name, id), nil
//				}).(pulumi.StringOutput),
//				ResourceGroupId: *pulumi.String(defaultResourceGroups.Groups[0].Id),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("Prometheus"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = arms.NewRemoteWrite(ctx, "defaultRemoteWrite", &arms.RemoteWriteArgs{
//				ClusterId: defaultPrometheus.ID(),
//				RemoteWriteYaml: pulumi.String(`remote_write:
//   - name: ArmsRemoteWrite
//     url: http://47.96.227.137:8080/prometheus/xxx/yyy/cn-hangzhou/api/v3/write
//     basic_auth: {username: 666, password: '******'}
//     write_relabel_configs:
//   - source_labels: [instance_id]
//     separator: ;
//     regex: si-6e2ca86444db4e55a7c1
//     replacement: $1
//     action: keep
//
// `),
//
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
// Application Real-Time Monitoring Service (ARMS) Remote Write can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:arms/remoteWrite:RemoteWrite example <cluster_id>:<remote_write_name>
// ```
type RemoteWrite struct {
	pulumi.CustomResourceState

	// The ID of the Prometheus instance.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The name of the Remote Write configuration item.
	RemoteWriteName pulumi.StringOutput `pulumi:"remoteWriteName"`
	// The details of the Remote Write configuration item. Specify the value in the YAML format.
	RemoteWriteYaml pulumi.StringOutput `pulumi:"remoteWriteYaml"`
}

// NewRemoteWrite registers a new resource with the given unique name, arguments, and options.
func NewRemoteWrite(ctx *pulumi.Context,
	name string, args *RemoteWriteArgs, opts ...pulumi.ResourceOption) (*RemoteWrite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.RemoteWriteYaml == nil {
		return nil, errors.New("invalid value for required argument 'RemoteWriteYaml'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RemoteWrite
	err := ctx.RegisterResource("alicloud:arms/remoteWrite:RemoteWrite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemoteWrite gets an existing RemoteWrite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemoteWrite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteWriteState, opts ...pulumi.ResourceOption) (*RemoteWrite, error) {
	var resource RemoteWrite
	err := ctx.ReadResource("alicloud:arms/remoteWrite:RemoteWrite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemoteWrite resources.
type remoteWriteState struct {
	// The ID of the Prometheus instance.
	ClusterId *string `pulumi:"clusterId"`
	// The name of the Remote Write configuration item.
	RemoteWriteName *string `pulumi:"remoteWriteName"`
	// The details of the Remote Write configuration item. Specify the value in the YAML format.
	RemoteWriteYaml *string `pulumi:"remoteWriteYaml"`
}

type RemoteWriteState struct {
	// The ID of the Prometheus instance.
	ClusterId pulumi.StringPtrInput
	// The name of the Remote Write configuration item.
	RemoteWriteName pulumi.StringPtrInput
	// The details of the Remote Write configuration item. Specify the value in the YAML format.
	RemoteWriteYaml pulumi.StringPtrInput
}

func (RemoteWriteState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteWriteState)(nil)).Elem()
}

type remoteWriteArgs struct {
	// The ID of the Prometheus instance.
	ClusterId string `pulumi:"clusterId"`
	// The details of the Remote Write configuration item. Specify the value in the YAML format.
	RemoteWriteYaml string `pulumi:"remoteWriteYaml"`
}

// The set of arguments for constructing a RemoteWrite resource.
type RemoteWriteArgs struct {
	// The ID of the Prometheus instance.
	ClusterId pulumi.StringInput
	// The details of the Remote Write configuration item. Specify the value in the YAML format.
	RemoteWriteYaml pulumi.StringInput
}

func (RemoteWriteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteWriteArgs)(nil)).Elem()
}

type RemoteWriteInput interface {
	pulumi.Input

	ToRemoteWriteOutput() RemoteWriteOutput
	ToRemoteWriteOutputWithContext(ctx context.Context) RemoteWriteOutput
}

func (*RemoteWrite) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteWrite)(nil)).Elem()
}

func (i *RemoteWrite) ToRemoteWriteOutput() RemoteWriteOutput {
	return i.ToRemoteWriteOutputWithContext(context.Background())
}

func (i *RemoteWrite) ToRemoteWriteOutputWithContext(ctx context.Context) RemoteWriteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteWriteOutput)
}

// RemoteWriteArrayInput is an input type that accepts RemoteWriteArray and RemoteWriteArrayOutput values.
// You can construct a concrete instance of `RemoteWriteArrayInput` via:
//
//	RemoteWriteArray{ RemoteWriteArgs{...} }
type RemoteWriteArrayInput interface {
	pulumi.Input

	ToRemoteWriteArrayOutput() RemoteWriteArrayOutput
	ToRemoteWriteArrayOutputWithContext(context.Context) RemoteWriteArrayOutput
}

type RemoteWriteArray []RemoteWriteInput

func (RemoteWriteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteWrite)(nil)).Elem()
}

func (i RemoteWriteArray) ToRemoteWriteArrayOutput() RemoteWriteArrayOutput {
	return i.ToRemoteWriteArrayOutputWithContext(context.Background())
}

func (i RemoteWriteArray) ToRemoteWriteArrayOutputWithContext(ctx context.Context) RemoteWriteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteWriteArrayOutput)
}

// RemoteWriteMapInput is an input type that accepts RemoteWriteMap and RemoteWriteMapOutput values.
// You can construct a concrete instance of `RemoteWriteMapInput` via:
//
//	RemoteWriteMap{ "key": RemoteWriteArgs{...} }
type RemoteWriteMapInput interface {
	pulumi.Input

	ToRemoteWriteMapOutput() RemoteWriteMapOutput
	ToRemoteWriteMapOutputWithContext(context.Context) RemoteWriteMapOutput
}

type RemoteWriteMap map[string]RemoteWriteInput

func (RemoteWriteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteWrite)(nil)).Elem()
}

func (i RemoteWriteMap) ToRemoteWriteMapOutput() RemoteWriteMapOutput {
	return i.ToRemoteWriteMapOutputWithContext(context.Background())
}

func (i RemoteWriteMap) ToRemoteWriteMapOutputWithContext(ctx context.Context) RemoteWriteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteWriteMapOutput)
}

type RemoteWriteOutput struct{ *pulumi.OutputState }

func (RemoteWriteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteWrite)(nil)).Elem()
}

func (o RemoteWriteOutput) ToRemoteWriteOutput() RemoteWriteOutput {
	return o
}

func (o RemoteWriteOutput) ToRemoteWriteOutputWithContext(ctx context.Context) RemoteWriteOutput {
	return o
}

// The ID of the Prometheus instance.
func (o RemoteWriteOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteWrite) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The name of the Remote Write configuration item.
func (o RemoteWriteOutput) RemoteWriteName() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteWrite) pulumi.StringOutput { return v.RemoteWriteName }).(pulumi.StringOutput)
}

// The details of the Remote Write configuration item. Specify the value in the YAML format.
func (o RemoteWriteOutput) RemoteWriteYaml() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteWrite) pulumi.StringOutput { return v.RemoteWriteYaml }).(pulumi.StringOutput)
}

type RemoteWriteArrayOutput struct{ *pulumi.OutputState }

func (RemoteWriteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteWrite)(nil)).Elem()
}

func (o RemoteWriteArrayOutput) ToRemoteWriteArrayOutput() RemoteWriteArrayOutput {
	return o
}

func (o RemoteWriteArrayOutput) ToRemoteWriteArrayOutputWithContext(ctx context.Context) RemoteWriteArrayOutput {
	return o
}

func (o RemoteWriteArrayOutput) Index(i pulumi.IntInput) RemoteWriteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RemoteWrite {
		return vs[0].([]*RemoteWrite)[vs[1].(int)]
	}).(RemoteWriteOutput)
}

type RemoteWriteMapOutput struct{ *pulumi.OutputState }

func (RemoteWriteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteWrite)(nil)).Elem()
}

func (o RemoteWriteMapOutput) ToRemoteWriteMapOutput() RemoteWriteMapOutput {
	return o
}

func (o RemoteWriteMapOutput) ToRemoteWriteMapOutputWithContext(ctx context.Context) RemoteWriteMapOutput {
	return o
}

func (o RemoteWriteMapOutput) MapIndex(k pulumi.StringInput) RemoteWriteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RemoteWrite {
		return vs[0].(map[string]*RemoteWrite)[vs[1].(string)]
	}).(RemoteWriteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteWriteInput)(nil)).Elem(), &RemoteWrite{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteWriteArrayInput)(nil)).Elem(), RemoteWriteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteWriteMapInput)(nil)).Elem(), RemoteWriteMap{})
	pulumi.RegisterOutputType(RemoteWriteOutput{})
	pulumi.RegisterOutputType(RemoteWriteArrayOutput{})
	pulumi.RegisterOutputType(RemoteWriteMapOutput{})
}
