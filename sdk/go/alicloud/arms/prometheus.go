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

// Provides a Application Real-Time Monitoring Service (ARMS) Prometheus resource.
//
// For information about Application Real-Time Monitoring Service (ARMS) Prometheus and how to use it, see [What is Prometheus](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-createprometheusinstance).
//
// > **NOTE:** Available since v1.203.0.
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
//			name := "tf_example"
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
//			_, err = arms.NewPrometheus(ctx, "defaultPrometheus", &arms.PrometheusArgs{
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
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Application Real-Time Monitoring Service (ARMS) Prometheus can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:arms/prometheus:Prometheus example <id>
// ```
type Prometheus struct {
	pulumi.CustomResourceState

	// The ID of the Kubernetes cluster. This parameter is required, if you set `clusterType` to `aliyun-cs`.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The name of the created cluster. This parameter is required, if you set `clusterType` to `remote-write`, `ecs` or `global-view`.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The type of the Prometheus instance. Valid values: `remote-write`, `ecs`, `global-view`, `aliyun-cs`.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// The ID of the Grafana dedicated instance. When using the shared version of Grafana, you can set `grafanaInstanceId` to `free`.
	GrafanaInstanceId pulumi.StringOutput `pulumi:"grafanaInstanceId"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The ID of the security group. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	SecurityGroupId pulumi.StringPtrOutput `pulumi:"securityGroupId"`
	// The child instance json string of the globalView instance.
	SubClustersJson pulumi.StringPtrOutput `pulumi:"subClustersJson"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the VPC. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
	// The ID of the VSwitch. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
}

// NewPrometheus registers a new resource with the given unique name, arguments, and options.
func NewPrometheus(ctx *pulumi.Context,
	name string, args *PrometheusArgs, opts ...pulumi.ResourceOption) (*Prometheus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterType == nil {
		return nil, errors.New("invalid value for required argument 'ClusterType'")
	}
	if args.GrafanaInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'GrafanaInstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Prometheus
	err := ctx.RegisterResource("alicloud:arms/prometheus:Prometheus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrometheus gets an existing Prometheus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrometheus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrometheusState, opts ...pulumi.ResourceOption) (*Prometheus, error) {
	var resource Prometheus
	err := ctx.ReadResource("alicloud:arms/prometheus:Prometheus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Prometheus resources.
type prometheusState struct {
	// The ID of the Kubernetes cluster. This parameter is required, if you set `clusterType` to `aliyun-cs`.
	ClusterId *string `pulumi:"clusterId"`
	// The name of the created cluster. This parameter is required, if you set `clusterType` to `remote-write`, `ecs` or `global-view`.
	ClusterName *string `pulumi:"clusterName"`
	// The type of the Prometheus instance. Valid values: `remote-write`, `ecs`, `global-view`, `aliyun-cs`.
	ClusterType *string `pulumi:"clusterType"`
	// The ID of the Grafana dedicated instance. When using the shared version of Grafana, you can set `grafanaInstanceId` to `free`.
	GrafanaInstanceId *string `pulumi:"grafanaInstanceId"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The ID of the security group. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The child instance json string of the globalView instance.
	SubClustersJson *string `pulumi:"subClustersJson"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	VpcId *string `pulumi:"vpcId"`
	// The ID of the VSwitch. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	VswitchId *string `pulumi:"vswitchId"`
}

type PrometheusState struct {
	// The ID of the Kubernetes cluster. This parameter is required, if you set `clusterType` to `aliyun-cs`.
	ClusterId pulumi.StringPtrInput
	// The name of the created cluster. This parameter is required, if you set `clusterType` to `remote-write`, `ecs` or `global-view`.
	ClusterName pulumi.StringPtrInput
	// The type of the Prometheus instance. Valid values: `remote-write`, `ecs`, `global-view`, `aliyun-cs`.
	ClusterType pulumi.StringPtrInput
	// The ID of the Grafana dedicated instance. When using the shared version of Grafana, you can set `grafanaInstanceId` to `free`.
	GrafanaInstanceId pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The ID of the security group. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	SecurityGroupId pulumi.StringPtrInput
	// The child instance json string of the globalView instance.
	SubClustersJson pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VPC. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	VpcId pulumi.StringPtrInput
	// The ID of the VSwitch. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	VswitchId pulumi.StringPtrInput
}

func (PrometheusState) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusState)(nil)).Elem()
}

type prometheusArgs struct {
	// The ID of the Kubernetes cluster. This parameter is required, if you set `clusterType` to `aliyun-cs`.
	ClusterId *string `pulumi:"clusterId"`
	// The name of the created cluster. This parameter is required, if you set `clusterType` to `remote-write`, `ecs` or `global-view`.
	ClusterName *string `pulumi:"clusterName"`
	// The type of the Prometheus instance. Valid values: `remote-write`, `ecs`, `global-view`, `aliyun-cs`.
	ClusterType string `pulumi:"clusterType"`
	// The ID of the Grafana dedicated instance. When using the shared version of Grafana, you can set `grafanaInstanceId` to `free`.
	GrafanaInstanceId string `pulumi:"grafanaInstanceId"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The ID of the security group. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The child instance json string of the globalView instance.
	SubClustersJson *string `pulumi:"subClustersJson"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	VpcId *string `pulumi:"vpcId"`
	// The ID of the VSwitch. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	VswitchId *string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a Prometheus resource.
type PrometheusArgs struct {
	// The ID of the Kubernetes cluster. This parameter is required, if you set `clusterType` to `aliyun-cs`.
	ClusterId pulumi.StringPtrInput
	// The name of the created cluster. This parameter is required, if you set `clusterType` to `remote-write`, `ecs` or `global-view`.
	ClusterName pulumi.StringPtrInput
	// The type of the Prometheus instance. Valid values: `remote-write`, `ecs`, `global-view`, `aliyun-cs`.
	ClusterType pulumi.StringInput
	// The ID of the Grafana dedicated instance. When using the shared version of Grafana, you can set `grafanaInstanceId` to `free`.
	GrafanaInstanceId pulumi.StringInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The ID of the security group. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	SecurityGroupId pulumi.StringPtrInput
	// The child instance json string of the globalView instance.
	SubClustersJson pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VPC. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	VpcId pulumi.StringPtrInput
	// The ID of the VSwitch. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
	VswitchId pulumi.StringPtrInput
}

func (PrometheusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusArgs)(nil)).Elem()
}

type PrometheusInput interface {
	pulumi.Input

	ToPrometheusOutput() PrometheusOutput
	ToPrometheusOutputWithContext(ctx context.Context) PrometheusOutput
}

func (*Prometheus) ElementType() reflect.Type {
	return reflect.TypeOf((**Prometheus)(nil)).Elem()
}

func (i *Prometheus) ToPrometheusOutput() PrometheusOutput {
	return i.ToPrometheusOutputWithContext(context.Background())
}

func (i *Prometheus) ToPrometheusOutputWithContext(ctx context.Context) PrometheusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusOutput)
}

// PrometheusArrayInput is an input type that accepts PrometheusArray and PrometheusArrayOutput values.
// You can construct a concrete instance of `PrometheusArrayInput` via:
//
//	PrometheusArray{ PrometheusArgs{...} }
type PrometheusArrayInput interface {
	pulumi.Input

	ToPrometheusArrayOutput() PrometheusArrayOutput
	ToPrometheusArrayOutputWithContext(context.Context) PrometheusArrayOutput
}

type PrometheusArray []PrometheusInput

func (PrometheusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Prometheus)(nil)).Elem()
}

func (i PrometheusArray) ToPrometheusArrayOutput() PrometheusArrayOutput {
	return i.ToPrometheusArrayOutputWithContext(context.Background())
}

func (i PrometheusArray) ToPrometheusArrayOutputWithContext(ctx context.Context) PrometheusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusArrayOutput)
}

// PrometheusMapInput is an input type that accepts PrometheusMap and PrometheusMapOutput values.
// You can construct a concrete instance of `PrometheusMapInput` via:
//
//	PrometheusMap{ "key": PrometheusArgs{...} }
type PrometheusMapInput interface {
	pulumi.Input

	ToPrometheusMapOutput() PrometheusMapOutput
	ToPrometheusMapOutputWithContext(context.Context) PrometheusMapOutput
}

type PrometheusMap map[string]PrometheusInput

func (PrometheusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Prometheus)(nil)).Elem()
}

func (i PrometheusMap) ToPrometheusMapOutput() PrometheusMapOutput {
	return i.ToPrometheusMapOutputWithContext(context.Background())
}

func (i PrometheusMap) ToPrometheusMapOutputWithContext(ctx context.Context) PrometheusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusMapOutput)
}

type PrometheusOutput struct{ *pulumi.OutputState }

func (PrometheusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Prometheus)(nil)).Elem()
}

func (o PrometheusOutput) ToPrometheusOutput() PrometheusOutput {
	return o
}

func (o PrometheusOutput) ToPrometheusOutputWithContext(ctx context.Context) PrometheusOutput {
	return o
}

// The ID of the Kubernetes cluster. This parameter is required, if you set `clusterType` to `aliyun-cs`.
func (o PrometheusOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The name of the created cluster. This parameter is required, if you set `clusterType` to `remote-write`, `ecs` or `global-view`.
func (o PrometheusOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// The type of the Prometheus instance. Valid values: `remote-write`, `ecs`, `global-view`, `aliyun-cs`.
func (o PrometheusOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// The ID of the Grafana dedicated instance. When using the shared version of Grafana, you can set `grafanaInstanceId` to `free`.
func (o PrometheusOutput) GrafanaInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringOutput { return v.GrafanaInstanceId }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o PrometheusOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The ID of the security group. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
func (o PrometheusOutput) SecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringPtrOutput { return v.SecurityGroupId }).(pulumi.StringPtrOutput)
}

// The child instance json string of the globalView instance.
func (o PrometheusOutput) SubClustersJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringPtrOutput { return v.SubClustersJson }).(pulumi.StringPtrOutput)
}

// A mapping of tags to assign to the resource.
func (o PrometheusOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The ID of the VPC. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
func (o PrometheusOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

// The ID of the VSwitch. This parameter is required, if you set `clusterType` to `ecs` or `aliyun-cs`(ASK instance).
func (o PrometheusOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringPtrOutput { return v.VswitchId }).(pulumi.StringPtrOutput)
}

type PrometheusArrayOutput struct{ *pulumi.OutputState }

func (PrometheusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Prometheus)(nil)).Elem()
}

func (o PrometheusArrayOutput) ToPrometheusArrayOutput() PrometheusArrayOutput {
	return o
}

func (o PrometheusArrayOutput) ToPrometheusArrayOutputWithContext(ctx context.Context) PrometheusArrayOutput {
	return o
}

func (o PrometheusArrayOutput) Index(i pulumi.IntInput) PrometheusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Prometheus {
		return vs[0].([]*Prometheus)[vs[1].(int)]
	}).(PrometheusOutput)
}

type PrometheusMapOutput struct{ *pulumi.OutputState }

func (PrometheusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Prometheus)(nil)).Elem()
}

func (o PrometheusMapOutput) ToPrometheusMapOutput() PrometheusMapOutput {
	return o
}

func (o PrometheusMapOutput) ToPrometheusMapOutputWithContext(ctx context.Context) PrometheusMapOutput {
	return o
}

func (o PrometheusMapOutput) MapIndex(k pulumi.StringInput) PrometheusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Prometheus {
		return vs[0].(map[string]*Prometheus)[vs[1].(string)]
	}).(PrometheusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusInput)(nil)).Elem(), &Prometheus{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusArrayInput)(nil)).Elem(), PrometheusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusMapInput)(nil)).Elem(), PrometheusMap{})
	pulumi.RegisterOutputType(PrometheusOutput{})
	pulumi.RegisterOutputType(PrometheusArrayOutput{})
	pulumi.RegisterOutputType(PrometheusMapOutput{})
}
