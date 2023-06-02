// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mse

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Microservice Engine (MSE) Engine Namespace resource.
//
// For information about Microservice Engine (MSE) Engine Namespace and how to use it, see [What is Engine Namespace](https://www.alibabacloud.com/help/zh/microservices-engine/latest/api-doc-mse-2019-05-31-api-doc-createenginenamespace).
//
// > **NOTE:** Available in v1.166.0+.
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
//			exampleZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleNetwork, err := vpc.NewNetwork(ctx, "exampleNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("172.17.3.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSwitch, err := vpc.NewSwitch(ctx, "exampleSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String("terraform-example"),
//				CidrBlock:   pulumi.String("172.17.3.0/24"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      *pulumi.String(exampleZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			exampleCluster, err := mse.NewCluster(ctx, "exampleCluster", &mse.ClusterArgs{
//				ClusterSpecification: pulumi.String("MSE_SC_1_2_60_c"),
//				ClusterType:          pulumi.String("Nacos-Ans"),
//				ClusterVersion:       pulumi.String("NACOS_2_0_0"),
//				InstanceCount:        pulumi.Int(1),
//				NetType:              pulumi.String("privatenet"),
//				PubNetworkFlow:       pulumi.String("1"),
//				ConnectionType:       pulumi.String("slb"),
//				ClusterAliasName:     pulumi.String("terraform-example"),
//				MseVersion:           pulumi.String("mse_dev"),
//				VswitchId:            exampleSwitch.ID(),
//				VpcId:                exampleNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mse.NewEngineNamespace(ctx, "exampleEngineNamespace", &mse.EngineNamespaceArgs{
//				ClusterId:         exampleCluster.ClusterId,
//				NamespaceShowName: pulumi.String("terraform-example"),
//				NamespaceId:       pulumi.String("terraform-example"),
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
// Microservice Engine (MSE) Engine Namespace can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:mse/engineNamespace:EngineNamespace example <cluster_id>:<namespace_id>
//
// ```
type EngineNamespace struct {
	pulumi.CustomResourceState

	// The language type of the returned information. Valid values: `zh`, `en`.
	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	// The id of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The id of Namespace.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// The name of the Engine Namespace.
	NamespaceShowName pulumi.StringOutput `pulumi:"namespaceShowName"`
}

// NewEngineNamespace registers a new resource with the given unique name, arguments, and options.
func NewEngineNamespace(ctx *pulumi.Context,
	name string, args *EngineNamespaceArgs, opts ...pulumi.ResourceOption) (*EngineNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	if args.NamespaceShowName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceShowName'")
	}
	var resource EngineNamespace
	err := ctx.RegisterResource("alicloud:mse/engineNamespace:EngineNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEngineNamespace gets an existing EngineNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEngineNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EngineNamespaceState, opts ...pulumi.ResourceOption) (*EngineNamespace, error) {
	var resource EngineNamespace
	err := ctx.ReadResource("alicloud:mse/engineNamespace:EngineNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EngineNamespace resources.
type engineNamespaceState struct {
	// The language type of the returned information. Valid values: `zh`, `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// The id of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The id of Namespace.
	NamespaceId *string `pulumi:"namespaceId"`
	// The name of the Engine Namespace.
	NamespaceShowName *string `pulumi:"namespaceShowName"`
}

type EngineNamespaceState struct {
	// The language type of the returned information. Valid values: `zh`, `en`.
	AcceptLanguage pulumi.StringPtrInput
	// The id of the cluster.
	ClusterId pulumi.StringPtrInput
	// The id of Namespace.
	NamespaceId pulumi.StringPtrInput
	// The name of the Engine Namespace.
	NamespaceShowName pulumi.StringPtrInput
}

func (EngineNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*engineNamespaceState)(nil)).Elem()
}

type engineNamespaceArgs struct {
	// The language type of the returned information. Valid values: `zh`, `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// The id of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// The id of Namespace.
	NamespaceId string `pulumi:"namespaceId"`
	// The name of the Engine Namespace.
	NamespaceShowName string `pulumi:"namespaceShowName"`
}

// The set of arguments for constructing a EngineNamespace resource.
type EngineNamespaceArgs struct {
	// The language type of the returned information. Valid values: `zh`, `en`.
	AcceptLanguage pulumi.StringPtrInput
	// The id of the cluster.
	ClusterId pulumi.StringInput
	// The id of Namespace.
	NamespaceId pulumi.StringInput
	// The name of the Engine Namespace.
	NamespaceShowName pulumi.StringInput
}

func (EngineNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*engineNamespaceArgs)(nil)).Elem()
}

type EngineNamespaceInput interface {
	pulumi.Input

	ToEngineNamespaceOutput() EngineNamespaceOutput
	ToEngineNamespaceOutputWithContext(ctx context.Context) EngineNamespaceOutput
}

func (*EngineNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineNamespace)(nil)).Elem()
}

func (i *EngineNamespace) ToEngineNamespaceOutput() EngineNamespaceOutput {
	return i.ToEngineNamespaceOutputWithContext(context.Background())
}

func (i *EngineNamespace) ToEngineNamespaceOutputWithContext(ctx context.Context) EngineNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineNamespaceOutput)
}

// EngineNamespaceArrayInput is an input type that accepts EngineNamespaceArray and EngineNamespaceArrayOutput values.
// You can construct a concrete instance of `EngineNamespaceArrayInput` via:
//
//	EngineNamespaceArray{ EngineNamespaceArgs{...} }
type EngineNamespaceArrayInput interface {
	pulumi.Input

	ToEngineNamespaceArrayOutput() EngineNamespaceArrayOutput
	ToEngineNamespaceArrayOutputWithContext(context.Context) EngineNamespaceArrayOutput
}

type EngineNamespaceArray []EngineNamespaceInput

func (EngineNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EngineNamespace)(nil)).Elem()
}

func (i EngineNamespaceArray) ToEngineNamespaceArrayOutput() EngineNamespaceArrayOutput {
	return i.ToEngineNamespaceArrayOutputWithContext(context.Background())
}

func (i EngineNamespaceArray) ToEngineNamespaceArrayOutputWithContext(ctx context.Context) EngineNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineNamespaceArrayOutput)
}

// EngineNamespaceMapInput is an input type that accepts EngineNamespaceMap and EngineNamespaceMapOutput values.
// You can construct a concrete instance of `EngineNamespaceMapInput` via:
//
//	EngineNamespaceMap{ "key": EngineNamespaceArgs{...} }
type EngineNamespaceMapInput interface {
	pulumi.Input

	ToEngineNamespaceMapOutput() EngineNamespaceMapOutput
	ToEngineNamespaceMapOutputWithContext(context.Context) EngineNamespaceMapOutput
}

type EngineNamespaceMap map[string]EngineNamespaceInput

func (EngineNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EngineNamespace)(nil)).Elem()
}

func (i EngineNamespaceMap) ToEngineNamespaceMapOutput() EngineNamespaceMapOutput {
	return i.ToEngineNamespaceMapOutputWithContext(context.Background())
}

func (i EngineNamespaceMap) ToEngineNamespaceMapOutputWithContext(ctx context.Context) EngineNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineNamespaceMapOutput)
}

type EngineNamespaceOutput struct{ *pulumi.OutputState }

func (EngineNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineNamespace)(nil)).Elem()
}

func (o EngineNamespaceOutput) ToEngineNamespaceOutput() EngineNamespaceOutput {
	return o
}

func (o EngineNamespaceOutput) ToEngineNamespaceOutputWithContext(ctx context.Context) EngineNamespaceOutput {
	return o
}

// The language type of the returned information. Valid values: `zh`, `en`.
func (o EngineNamespaceOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EngineNamespace) pulumi.StringPtrOutput { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

// The id of the cluster.
func (o EngineNamespaceOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *EngineNamespace) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The id of Namespace.
func (o EngineNamespaceOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EngineNamespace) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// The name of the Engine Namespace.
func (o EngineNamespaceOutput) NamespaceShowName() pulumi.StringOutput {
	return o.ApplyT(func(v *EngineNamespace) pulumi.StringOutput { return v.NamespaceShowName }).(pulumi.StringOutput)
}

type EngineNamespaceArrayOutput struct{ *pulumi.OutputState }

func (EngineNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EngineNamespace)(nil)).Elem()
}

func (o EngineNamespaceArrayOutput) ToEngineNamespaceArrayOutput() EngineNamespaceArrayOutput {
	return o
}

func (o EngineNamespaceArrayOutput) ToEngineNamespaceArrayOutputWithContext(ctx context.Context) EngineNamespaceArrayOutput {
	return o
}

func (o EngineNamespaceArrayOutput) Index(i pulumi.IntInput) EngineNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EngineNamespace {
		return vs[0].([]*EngineNamespace)[vs[1].(int)]
	}).(EngineNamespaceOutput)
}

type EngineNamespaceMapOutput struct{ *pulumi.OutputState }

func (EngineNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EngineNamespace)(nil)).Elem()
}

func (o EngineNamespaceMapOutput) ToEngineNamespaceMapOutput() EngineNamespaceMapOutput {
	return o
}

func (o EngineNamespaceMapOutput) ToEngineNamespaceMapOutputWithContext(ctx context.Context) EngineNamespaceMapOutput {
	return o
}

func (o EngineNamespaceMapOutput) MapIndex(k pulumi.StringInput) EngineNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EngineNamespace {
		return vs[0].(map[string]*EngineNamespace)[vs[1].(string)]
	}).(EngineNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EngineNamespaceInput)(nil)).Elem(), &EngineNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*EngineNamespaceArrayInput)(nil)).Elem(), EngineNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EngineNamespaceMapInput)(nil)).Elem(), EngineNamespaceMap{})
	pulumi.RegisterOutputType(EngineNamespaceOutput{})
	pulumi.RegisterOutputType(EngineNamespaceArrayOutput{})
	pulumi.RegisterOutputType(EngineNamespaceMapOutput{})
}
