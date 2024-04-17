// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mse

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Microservice Engine (MSE) Znode resource.
//
// For information about Microservice Engine (MSE) Znode and how to use it, see [What is Znode](https://help.aliyun.com/document_detail/393622.html).
//
// > **NOTE:** Available in v1.162.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mse"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleNetwork, err := vpc.NewNetwork(ctx, "example", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("172.17.3.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSwitch, err := vpc.NewSwitch(ctx, "example", &vpc.SwitchArgs{
//				VswitchName: pulumi.String("terraform-example"),
//				CidrBlock:   pulumi.String("172.17.3.0/24"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      pulumi.String(example.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			exampleCluster, err := mse.NewCluster(ctx, "example", &mse.ClusterArgs{
//				ClusterSpecification: pulumi.String("MSE_SC_1_2_60_c"),
//				ClusterType:          pulumi.String("ZooKeeper"),
//				ClusterVersion:       pulumi.String("ZooKeeper_3_8_0"),
//				InstanceCount:        pulumi.Int(1),
//				NetType:              pulumi.String("privatenet"),
//				PubNetworkFlow:       pulumi.String("1"),
//				AclEntryLists: pulumi.StringArray{
//					pulumi.String("127.0.0.1/32"),
//				},
//				ClusterAliasName: pulumi.String("terraform-example"),
//				MseVersion:       pulumi.String("mse_dev"),
//				VswitchId:        exampleSwitch.ID(),
//				VpcId:            exampleNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mse.NewZnode(ctx, "example", &mse.ZnodeArgs{
//				ClusterId: exampleCluster.ClusterId,
//				Data:      pulumi.String("terraform-example"),
//				Path:      pulumi.String("/example"),
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
// Microservice Engine (MSE) Znode can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:mse/znode:Znode example <cluster_id>:<path>
// ```
type Znode struct {
	pulumi.CustomResourceState

	// The language type of the returned information. Valid values: `zh` or `en`.
	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	// The ID of the Cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The Node data.
	Data pulumi.StringPtrOutput `pulumi:"data"`
	// The Node path. The value must start with a forward slash (/).
	Path pulumi.StringOutput `pulumi:"path"`
}

// NewZnode registers a new resource with the given unique name, arguments, and options.
func NewZnode(ctx *pulumi.Context,
	name string, args *ZnodeArgs, opts ...pulumi.ResourceOption) (*Znode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Znode
	err := ctx.RegisterResource("alicloud:mse/znode:Znode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZnode gets an existing Znode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZnode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZnodeState, opts ...pulumi.ResourceOption) (*Znode, error) {
	var resource Znode
	err := ctx.ReadResource("alicloud:mse/znode:Znode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Znode resources.
type znodeState struct {
	// The language type of the returned information. Valid values: `zh` or `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// The ID of the Cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The Node data.
	Data *string `pulumi:"data"`
	// The Node path. The value must start with a forward slash (/).
	Path *string `pulumi:"path"`
}

type ZnodeState struct {
	// The language type of the returned information. Valid values: `zh` or `en`.
	AcceptLanguage pulumi.StringPtrInput
	// The ID of the Cluster.
	ClusterId pulumi.StringPtrInput
	// The Node data.
	Data pulumi.StringPtrInput
	// The Node path. The value must start with a forward slash (/).
	Path pulumi.StringPtrInput
}

func (ZnodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*znodeState)(nil)).Elem()
}

type znodeArgs struct {
	// The language type of the returned information. Valid values: `zh` or `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// The ID of the Cluster.
	ClusterId string `pulumi:"clusterId"`
	// The Node data.
	Data *string `pulumi:"data"`
	// The Node path. The value must start with a forward slash (/).
	Path string `pulumi:"path"`
}

// The set of arguments for constructing a Znode resource.
type ZnodeArgs struct {
	// The language type of the returned information. Valid values: `zh` or `en`.
	AcceptLanguage pulumi.StringPtrInput
	// The ID of the Cluster.
	ClusterId pulumi.StringInput
	// The Node data.
	Data pulumi.StringPtrInput
	// The Node path. The value must start with a forward slash (/).
	Path pulumi.StringInput
}

func (ZnodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*znodeArgs)(nil)).Elem()
}

type ZnodeInput interface {
	pulumi.Input

	ToZnodeOutput() ZnodeOutput
	ToZnodeOutputWithContext(ctx context.Context) ZnodeOutput
}

func (*Znode) ElementType() reflect.Type {
	return reflect.TypeOf((**Znode)(nil)).Elem()
}

func (i *Znode) ToZnodeOutput() ZnodeOutput {
	return i.ToZnodeOutputWithContext(context.Background())
}

func (i *Znode) ToZnodeOutputWithContext(ctx context.Context) ZnodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZnodeOutput)
}

// ZnodeArrayInput is an input type that accepts ZnodeArray and ZnodeArrayOutput values.
// You can construct a concrete instance of `ZnodeArrayInput` via:
//
//	ZnodeArray{ ZnodeArgs{...} }
type ZnodeArrayInput interface {
	pulumi.Input

	ToZnodeArrayOutput() ZnodeArrayOutput
	ToZnodeArrayOutputWithContext(context.Context) ZnodeArrayOutput
}

type ZnodeArray []ZnodeInput

func (ZnodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Znode)(nil)).Elem()
}

func (i ZnodeArray) ToZnodeArrayOutput() ZnodeArrayOutput {
	return i.ToZnodeArrayOutputWithContext(context.Background())
}

func (i ZnodeArray) ToZnodeArrayOutputWithContext(ctx context.Context) ZnodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZnodeArrayOutput)
}

// ZnodeMapInput is an input type that accepts ZnodeMap and ZnodeMapOutput values.
// You can construct a concrete instance of `ZnodeMapInput` via:
//
//	ZnodeMap{ "key": ZnodeArgs{...} }
type ZnodeMapInput interface {
	pulumi.Input

	ToZnodeMapOutput() ZnodeMapOutput
	ToZnodeMapOutputWithContext(context.Context) ZnodeMapOutput
}

type ZnodeMap map[string]ZnodeInput

func (ZnodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Znode)(nil)).Elem()
}

func (i ZnodeMap) ToZnodeMapOutput() ZnodeMapOutput {
	return i.ToZnodeMapOutputWithContext(context.Background())
}

func (i ZnodeMap) ToZnodeMapOutputWithContext(ctx context.Context) ZnodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZnodeMapOutput)
}

type ZnodeOutput struct{ *pulumi.OutputState }

func (ZnodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Znode)(nil)).Elem()
}

func (o ZnodeOutput) ToZnodeOutput() ZnodeOutput {
	return o
}

func (o ZnodeOutput) ToZnodeOutputWithContext(ctx context.Context) ZnodeOutput {
	return o
}

// The language type of the returned information. Valid values: `zh` or `en`.
func (o ZnodeOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Znode) pulumi.StringPtrOutput { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

// The ID of the Cluster.
func (o ZnodeOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Znode) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The Node data.
func (o ZnodeOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Znode) pulumi.StringPtrOutput { return v.Data }).(pulumi.StringPtrOutput)
}

// The Node path. The value must start with a forward slash (/).
func (o ZnodeOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Znode) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

type ZnodeArrayOutput struct{ *pulumi.OutputState }

func (ZnodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Znode)(nil)).Elem()
}

func (o ZnodeArrayOutput) ToZnodeArrayOutput() ZnodeArrayOutput {
	return o
}

func (o ZnodeArrayOutput) ToZnodeArrayOutputWithContext(ctx context.Context) ZnodeArrayOutput {
	return o
}

func (o ZnodeArrayOutput) Index(i pulumi.IntInput) ZnodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Znode {
		return vs[0].([]*Znode)[vs[1].(int)]
	}).(ZnodeOutput)
}

type ZnodeMapOutput struct{ *pulumi.OutputState }

func (ZnodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Znode)(nil)).Elem()
}

func (o ZnodeMapOutput) ToZnodeMapOutput() ZnodeMapOutput {
	return o
}

func (o ZnodeMapOutput) ToZnodeMapOutputWithContext(ctx context.Context) ZnodeMapOutput {
	return o
}

func (o ZnodeMapOutput) MapIndex(k pulumi.StringInput) ZnodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Znode {
		return vs[0].(map[string]*Znode)[vs[1].(string)]
	}).(ZnodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZnodeInput)(nil)).Elem(), &Znode{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZnodeArrayInput)(nil)).Elem(), ZnodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZnodeMapInput)(nil)).Elem(), ZnodeMap{})
	pulumi.RegisterOutputType(ZnodeOutput{})
	pulumi.RegisterOutputType(ZnodeArrayOutput{})
	pulumi.RegisterOutputType(ZnodeMapOutput{})
}
