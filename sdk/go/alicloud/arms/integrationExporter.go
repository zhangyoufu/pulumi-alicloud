// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package arms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Application Real-Time Monitoring Service (ARMS) Integration Exporter resource.
//
// For information about Application Real-Time Monitoring Service (ARMS) Integration Exporter and how to use it, see [What is Integration Exporter](https://www.alibabacloud.com/help/en/application-real-time-monitoring-service/latest/api-doc-arms-2019-08-08-api-doc-addprometheusintegration).
//
// > **NOTE:** Available since v1.203.0.
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
//				Tags: pulumi.AnyMap{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("Prometheus"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = arms.NewIntegrationExporter(ctx, "defaultIntegrationExporter", &arms.IntegrationExporterArgs{
//				ClusterId:       defaultPrometheus.ID(),
//				IntegrationType: pulumi.String("kafka"),
//				Param:           pulumi.String("{\"tls_insecure-skip-tls-verify\":\"none=tls.insecure-skip-tls-verify\",\"tls_enabled\":\"none=tls.enabled\",\"sasl_mechanism\":\"\",\"name\":\"kafka1\",\"sasl_enabled\":\"none=sasl.enabled\",\"ip_ports\":\"abc:888\",\"scrape_interval\":30,\"version\":\"0.10.1.0\"}"),
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
// Application Real-Time Monitoring Service (ARMS) Integration Exporter can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:arms/integrationExporter:IntegrationExporter example <cluster_id>:<integration_type>:<instance_id>
//
// ```
type IntegrationExporter struct {
	pulumi.CustomResourceState

	// The ID of the Prometheus instance.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The ID of the Integration Exporter instance.
	InstanceId pulumi.IntOutput `pulumi:"instanceId"`
	// The type of prometheus integration.
	IntegrationType pulumi.StringOutput `pulumi:"integrationType"`
	// Exporter configuration parameter json string.
	Param pulumi.StringOutput `pulumi:"param"`
}

// NewIntegrationExporter registers a new resource with the given unique name, arguments, and options.
func NewIntegrationExporter(ctx *pulumi.Context,
	name string, args *IntegrationExporterArgs, opts ...pulumi.ResourceOption) (*IntegrationExporter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.IntegrationType == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationType'")
	}
	if args.Param == nil {
		return nil, errors.New("invalid value for required argument 'Param'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IntegrationExporter
	err := ctx.RegisterResource("alicloud:arms/integrationExporter:IntegrationExporter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationExporter gets an existing IntegrationExporter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationExporter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationExporterState, opts ...pulumi.ResourceOption) (*IntegrationExporter, error) {
	var resource IntegrationExporter
	err := ctx.ReadResource("alicloud:arms/integrationExporter:IntegrationExporter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationExporter resources.
type integrationExporterState struct {
	// The ID of the Prometheus instance.
	ClusterId *string `pulumi:"clusterId"`
	// The ID of the Integration Exporter instance.
	InstanceId *int `pulumi:"instanceId"`
	// The type of prometheus integration.
	IntegrationType *string `pulumi:"integrationType"`
	// Exporter configuration parameter json string.
	Param *string `pulumi:"param"`
}

type IntegrationExporterState struct {
	// The ID of the Prometheus instance.
	ClusterId pulumi.StringPtrInput
	// The ID of the Integration Exporter instance.
	InstanceId pulumi.IntPtrInput
	// The type of prometheus integration.
	IntegrationType pulumi.StringPtrInput
	// Exporter configuration parameter json string.
	Param pulumi.StringPtrInput
}

func (IntegrationExporterState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationExporterState)(nil)).Elem()
}

type integrationExporterArgs struct {
	// The ID of the Prometheus instance.
	ClusterId string `pulumi:"clusterId"`
	// The type of prometheus integration.
	IntegrationType string `pulumi:"integrationType"`
	// Exporter configuration parameter json string.
	Param string `pulumi:"param"`
}

// The set of arguments for constructing a IntegrationExporter resource.
type IntegrationExporterArgs struct {
	// The ID of the Prometheus instance.
	ClusterId pulumi.StringInput
	// The type of prometheus integration.
	IntegrationType pulumi.StringInput
	// Exporter configuration parameter json string.
	Param pulumi.StringInput
}

func (IntegrationExporterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationExporterArgs)(nil)).Elem()
}

type IntegrationExporterInput interface {
	pulumi.Input

	ToIntegrationExporterOutput() IntegrationExporterOutput
	ToIntegrationExporterOutputWithContext(ctx context.Context) IntegrationExporterOutput
}

func (*IntegrationExporter) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationExporter)(nil)).Elem()
}

func (i *IntegrationExporter) ToIntegrationExporterOutput() IntegrationExporterOutput {
	return i.ToIntegrationExporterOutputWithContext(context.Background())
}

func (i *IntegrationExporter) ToIntegrationExporterOutputWithContext(ctx context.Context) IntegrationExporterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationExporterOutput)
}

func (i *IntegrationExporter) ToOutput(ctx context.Context) pulumix.Output[*IntegrationExporter] {
	return pulumix.Output[*IntegrationExporter]{
		OutputState: i.ToIntegrationExporterOutputWithContext(ctx).OutputState,
	}
}

// IntegrationExporterArrayInput is an input type that accepts IntegrationExporterArray and IntegrationExporterArrayOutput values.
// You can construct a concrete instance of `IntegrationExporterArrayInput` via:
//
//	IntegrationExporterArray{ IntegrationExporterArgs{...} }
type IntegrationExporterArrayInput interface {
	pulumi.Input

	ToIntegrationExporterArrayOutput() IntegrationExporterArrayOutput
	ToIntegrationExporterArrayOutputWithContext(context.Context) IntegrationExporterArrayOutput
}

type IntegrationExporterArray []IntegrationExporterInput

func (IntegrationExporterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationExporter)(nil)).Elem()
}

func (i IntegrationExporterArray) ToIntegrationExporterArrayOutput() IntegrationExporterArrayOutput {
	return i.ToIntegrationExporterArrayOutputWithContext(context.Background())
}

func (i IntegrationExporterArray) ToIntegrationExporterArrayOutputWithContext(ctx context.Context) IntegrationExporterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationExporterArrayOutput)
}

func (i IntegrationExporterArray) ToOutput(ctx context.Context) pulumix.Output[[]*IntegrationExporter] {
	return pulumix.Output[[]*IntegrationExporter]{
		OutputState: i.ToIntegrationExporterArrayOutputWithContext(ctx).OutputState,
	}
}

// IntegrationExporterMapInput is an input type that accepts IntegrationExporterMap and IntegrationExporterMapOutput values.
// You can construct a concrete instance of `IntegrationExporterMapInput` via:
//
//	IntegrationExporterMap{ "key": IntegrationExporterArgs{...} }
type IntegrationExporterMapInput interface {
	pulumi.Input

	ToIntegrationExporterMapOutput() IntegrationExporterMapOutput
	ToIntegrationExporterMapOutputWithContext(context.Context) IntegrationExporterMapOutput
}

type IntegrationExporterMap map[string]IntegrationExporterInput

func (IntegrationExporterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationExporter)(nil)).Elem()
}

func (i IntegrationExporterMap) ToIntegrationExporterMapOutput() IntegrationExporterMapOutput {
	return i.ToIntegrationExporterMapOutputWithContext(context.Background())
}

func (i IntegrationExporterMap) ToIntegrationExporterMapOutputWithContext(ctx context.Context) IntegrationExporterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationExporterMapOutput)
}

func (i IntegrationExporterMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IntegrationExporter] {
	return pulumix.Output[map[string]*IntegrationExporter]{
		OutputState: i.ToIntegrationExporterMapOutputWithContext(ctx).OutputState,
	}
}

type IntegrationExporterOutput struct{ *pulumi.OutputState }

func (IntegrationExporterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationExporter)(nil)).Elem()
}

func (o IntegrationExporterOutput) ToIntegrationExporterOutput() IntegrationExporterOutput {
	return o
}

func (o IntegrationExporterOutput) ToIntegrationExporterOutputWithContext(ctx context.Context) IntegrationExporterOutput {
	return o
}

func (o IntegrationExporterOutput) ToOutput(ctx context.Context) pulumix.Output[*IntegrationExporter] {
	return pulumix.Output[*IntegrationExporter]{
		OutputState: o.OutputState,
	}
}

// The ID of the Prometheus instance.
func (o IntegrationExporterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationExporter) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The ID of the Integration Exporter instance.
func (o IntegrationExporterOutput) InstanceId() pulumi.IntOutput {
	return o.ApplyT(func(v *IntegrationExporter) pulumi.IntOutput { return v.InstanceId }).(pulumi.IntOutput)
}

// The type of prometheus integration.
func (o IntegrationExporterOutput) IntegrationType() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationExporter) pulumi.StringOutput { return v.IntegrationType }).(pulumi.StringOutput)
}

// Exporter configuration parameter json string.
func (o IntegrationExporterOutput) Param() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationExporter) pulumi.StringOutput { return v.Param }).(pulumi.StringOutput)
}

type IntegrationExporterArrayOutput struct{ *pulumi.OutputState }

func (IntegrationExporterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IntegrationExporter)(nil)).Elem()
}

func (o IntegrationExporterArrayOutput) ToIntegrationExporterArrayOutput() IntegrationExporterArrayOutput {
	return o
}

func (o IntegrationExporterArrayOutput) ToIntegrationExporterArrayOutputWithContext(ctx context.Context) IntegrationExporterArrayOutput {
	return o
}

func (o IntegrationExporterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IntegrationExporter] {
	return pulumix.Output[[]*IntegrationExporter]{
		OutputState: o.OutputState,
	}
}

func (o IntegrationExporterArrayOutput) Index(i pulumi.IntInput) IntegrationExporterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IntegrationExporter {
		return vs[0].([]*IntegrationExporter)[vs[1].(int)]
	}).(IntegrationExporterOutput)
}

type IntegrationExporterMapOutput struct{ *pulumi.OutputState }

func (IntegrationExporterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IntegrationExporter)(nil)).Elem()
}

func (o IntegrationExporterMapOutput) ToIntegrationExporterMapOutput() IntegrationExporterMapOutput {
	return o
}

func (o IntegrationExporterMapOutput) ToIntegrationExporterMapOutputWithContext(ctx context.Context) IntegrationExporterMapOutput {
	return o
}

func (o IntegrationExporterMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IntegrationExporter] {
	return pulumix.Output[map[string]*IntegrationExporter]{
		OutputState: o.OutputState,
	}
}

func (o IntegrationExporterMapOutput) MapIndex(k pulumi.StringInput) IntegrationExporterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IntegrationExporter {
		return vs[0].(map[string]*IntegrationExporter)[vs[1].(string)]
	}).(IntegrationExporterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationExporterInput)(nil)).Elem(), &IntegrationExporter{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationExporterArrayInput)(nil)).Elem(), IntegrationExporterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationExporterMapInput)(nil)).Elem(), IntegrationExporterMap{})
	pulumi.RegisterOutputType(IntegrationExporterOutput{})
	pulumi.RegisterOutputType(IntegrationExporterArrayOutput{})
	pulumi.RegisterOutputType(IntegrationExporterMapOutput{})
}
