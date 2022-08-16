// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a MongoDB Sharding Network Public Address resource.
//
// For information about MongoDB Sharding Network Public Address and how to use it, see [What is Sharding Network Public Address](https://www.alibabacloud.com/help/doc-detail/67602.html).
//
// > **NOTE:** Available in v1.149.0+.
//
// > **NOTE:** This operation supports sharded cluster instances only.
//
// ## Import
//
// MongoDB Sharding Network Public Address can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:mongodb/shardingNetworkPublicAddress:ShardingNetworkPublicAddress example <db_instance_id>:<node_id>
//
// ```
type ShardingNetworkPublicAddress struct {
	pulumi.CustomResourceState

	// The ID of the instance.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// The endpoint of the instance.
	NetworkAddresses ShardingNetworkPublicAddressNetworkAddressArrayOutput `pulumi:"networkAddresses"`
	// The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
	NodeId pulumi.StringOutput `pulumi:"nodeId"`
}

// NewShardingNetworkPublicAddress registers a new resource with the given unique name, arguments, and options.
func NewShardingNetworkPublicAddress(ctx *pulumi.Context,
	name string, args *ShardingNetworkPublicAddressArgs, opts ...pulumi.ResourceOption) (*ShardingNetworkPublicAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	if args.NodeId == nil {
		return nil, errors.New("invalid value for required argument 'NodeId'")
	}
	var resource ShardingNetworkPublicAddress
	err := ctx.RegisterResource("alicloud:mongodb/shardingNetworkPublicAddress:ShardingNetworkPublicAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShardingNetworkPublicAddress gets an existing ShardingNetworkPublicAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShardingNetworkPublicAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShardingNetworkPublicAddressState, opts ...pulumi.ResourceOption) (*ShardingNetworkPublicAddress, error) {
	var resource ShardingNetworkPublicAddress
	err := ctx.ReadResource("alicloud:mongodb/shardingNetworkPublicAddress:ShardingNetworkPublicAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShardingNetworkPublicAddress resources.
type shardingNetworkPublicAddressState struct {
	// The ID of the instance.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// The endpoint of the instance.
	NetworkAddresses []ShardingNetworkPublicAddressNetworkAddress `pulumi:"networkAddresses"`
	// The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
	NodeId *string `pulumi:"nodeId"`
}

type ShardingNetworkPublicAddressState struct {
	// The ID of the instance.
	DbInstanceId pulumi.StringPtrInput
	// The endpoint of the instance.
	NetworkAddresses ShardingNetworkPublicAddressNetworkAddressArrayInput
	// The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
	NodeId pulumi.StringPtrInput
}

func (ShardingNetworkPublicAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*shardingNetworkPublicAddressState)(nil)).Elem()
}

type shardingNetworkPublicAddressArgs struct {
	// The ID of the instance.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
	NodeId string `pulumi:"nodeId"`
}

// The set of arguments for constructing a ShardingNetworkPublicAddress resource.
type ShardingNetworkPublicAddressArgs struct {
	// The ID of the instance.
	DbInstanceId pulumi.StringInput
	// The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
	NodeId pulumi.StringInput
}

func (ShardingNetworkPublicAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shardingNetworkPublicAddressArgs)(nil)).Elem()
}

type ShardingNetworkPublicAddressInput interface {
	pulumi.Input

	ToShardingNetworkPublicAddressOutput() ShardingNetworkPublicAddressOutput
	ToShardingNetworkPublicAddressOutputWithContext(ctx context.Context) ShardingNetworkPublicAddressOutput
}

func (*ShardingNetworkPublicAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**ShardingNetworkPublicAddress)(nil)).Elem()
}

func (i *ShardingNetworkPublicAddress) ToShardingNetworkPublicAddressOutput() ShardingNetworkPublicAddressOutput {
	return i.ToShardingNetworkPublicAddressOutputWithContext(context.Background())
}

func (i *ShardingNetworkPublicAddress) ToShardingNetworkPublicAddressOutputWithContext(ctx context.Context) ShardingNetworkPublicAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingNetworkPublicAddressOutput)
}

// ShardingNetworkPublicAddressArrayInput is an input type that accepts ShardingNetworkPublicAddressArray and ShardingNetworkPublicAddressArrayOutput values.
// You can construct a concrete instance of `ShardingNetworkPublicAddressArrayInput` via:
//
//	ShardingNetworkPublicAddressArray{ ShardingNetworkPublicAddressArgs{...} }
type ShardingNetworkPublicAddressArrayInput interface {
	pulumi.Input

	ToShardingNetworkPublicAddressArrayOutput() ShardingNetworkPublicAddressArrayOutput
	ToShardingNetworkPublicAddressArrayOutputWithContext(context.Context) ShardingNetworkPublicAddressArrayOutput
}

type ShardingNetworkPublicAddressArray []ShardingNetworkPublicAddressInput

func (ShardingNetworkPublicAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShardingNetworkPublicAddress)(nil)).Elem()
}

func (i ShardingNetworkPublicAddressArray) ToShardingNetworkPublicAddressArrayOutput() ShardingNetworkPublicAddressArrayOutput {
	return i.ToShardingNetworkPublicAddressArrayOutputWithContext(context.Background())
}

func (i ShardingNetworkPublicAddressArray) ToShardingNetworkPublicAddressArrayOutputWithContext(ctx context.Context) ShardingNetworkPublicAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingNetworkPublicAddressArrayOutput)
}

// ShardingNetworkPublicAddressMapInput is an input type that accepts ShardingNetworkPublicAddressMap and ShardingNetworkPublicAddressMapOutput values.
// You can construct a concrete instance of `ShardingNetworkPublicAddressMapInput` via:
//
//	ShardingNetworkPublicAddressMap{ "key": ShardingNetworkPublicAddressArgs{...} }
type ShardingNetworkPublicAddressMapInput interface {
	pulumi.Input

	ToShardingNetworkPublicAddressMapOutput() ShardingNetworkPublicAddressMapOutput
	ToShardingNetworkPublicAddressMapOutputWithContext(context.Context) ShardingNetworkPublicAddressMapOutput
}

type ShardingNetworkPublicAddressMap map[string]ShardingNetworkPublicAddressInput

func (ShardingNetworkPublicAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShardingNetworkPublicAddress)(nil)).Elem()
}

func (i ShardingNetworkPublicAddressMap) ToShardingNetworkPublicAddressMapOutput() ShardingNetworkPublicAddressMapOutput {
	return i.ToShardingNetworkPublicAddressMapOutputWithContext(context.Background())
}

func (i ShardingNetworkPublicAddressMap) ToShardingNetworkPublicAddressMapOutputWithContext(ctx context.Context) ShardingNetworkPublicAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingNetworkPublicAddressMapOutput)
}

type ShardingNetworkPublicAddressOutput struct{ *pulumi.OutputState }

func (ShardingNetworkPublicAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShardingNetworkPublicAddress)(nil)).Elem()
}

func (o ShardingNetworkPublicAddressOutput) ToShardingNetworkPublicAddressOutput() ShardingNetworkPublicAddressOutput {
	return o
}

func (o ShardingNetworkPublicAddressOutput) ToShardingNetworkPublicAddressOutputWithContext(ctx context.Context) ShardingNetworkPublicAddressOutput {
	return o
}

// The ID of the instance.
func (o ShardingNetworkPublicAddressOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingNetworkPublicAddress) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The endpoint of the instance.
func (o ShardingNetworkPublicAddressOutput) NetworkAddresses() ShardingNetworkPublicAddressNetworkAddressArrayOutput {
	return o.ApplyT(func(v *ShardingNetworkPublicAddress) ShardingNetworkPublicAddressNetworkAddressArrayOutput {
		return v.NetworkAddresses
	}).(ShardingNetworkPublicAddressNetworkAddressArrayOutput)
}

// The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
func (o ShardingNetworkPublicAddressOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v *ShardingNetworkPublicAddress) pulumi.StringOutput { return v.NodeId }).(pulumi.StringOutput)
}

type ShardingNetworkPublicAddressArrayOutput struct{ *pulumi.OutputState }

func (ShardingNetworkPublicAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShardingNetworkPublicAddress)(nil)).Elem()
}

func (o ShardingNetworkPublicAddressArrayOutput) ToShardingNetworkPublicAddressArrayOutput() ShardingNetworkPublicAddressArrayOutput {
	return o
}

func (o ShardingNetworkPublicAddressArrayOutput) ToShardingNetworkPublicAddressArrayOutputWithContext(ctx context.Context) ShardingNetworkPublicAddressArrayOutput {
	return o
}

func (o ShardingNetworkPublicAddressArrayOutput) Index(i pulumi.IntInput) ShardingNetworkPublicAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ShardingNetworkPublicAddress {
		return vs[0].([]*ShardingNetworkPublicAddress)[vs[1].(int)]
	}).(ShardingNetworkPublicAddressOutput)
}

type ShardingNetworkPublicAddressMapOutput struct{ *pulumi.OutputState }

func (ShardingNetworkPublicAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShardingNetworkPublicAddress)(nil)).Elem()
}

func (o ShardingNetworkPublicAddressMapOutput) ToShardingNetworkPublicAddressMapOutput() ShardingNetworkPublicAddressMapOutput {
	return o
}

func (o ShardingNetworkPublicAddressMapOutput) ToShardingNetworkPublicAddressMapOutputWithContext(ctx context.Context) ShardingNetworkPublicAddressMapOutput {
	return o
}

func (o ShardingNetworkPublicAddressMapOutput) MapIndex(k pulumi.StringInput) ShardingNetworkPublicAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ShardingNetworkPublicAddress {
		return vs[0].(map[string]*ShardingNetworkPublicAddress)[vs[1].(string)]
	}).(ShardingNetworkPublicAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShardingNetworkPublicAddressInput)(nil)).Elem(), &ShardingNetworkPublicAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShardingNetworkPublicAddressArrayInput)(nil)).Elem(), ShardingNetworkPublicAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShardingNetworkPublicAddressMapInput)(nil)).Elem(), ShardingNetworkPublicAddressMap{})
	pulumi.RegisterOutputType(ShardingNetworkPublicAddressOutput{})
	pulumi.RegisterOutputType(ShardingNetworkPublicAddressArrayOutput{})
	pulumi.RegisterOutputType(ShardingNetworkPublicAddressMapOutput{})
}
