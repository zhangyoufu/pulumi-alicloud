// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetMongoInstancesInstance struct {
	AvailabilityZone string                           `pulumi:"availabilityZone"`
	ChargeType       string                           `pulumi:"chargeType"`
	CreationTime     string                           `pulumi:"creationTime"`
	Engine           string                           `pulumi:"engine"`
	EngineVersion    string                           `pulumi:"engineVersion"`
	ExpirationTime   string                           `pulumi:"expirationTime"`
	Id               string                           `pulumi:"id"`
	InstanceClass    string                           `pulumi:"instanceClass"`
	InstanceType     string                           `pulumi:"instanceType"`
	LockMode         string                           `pulumi:"lockMode"`
	Mongos           []GetMongoInstancesInstanceMongo `pulumi:"mongos"`
	Name             string                           `pulumi:"name"`
	NetworkType      string                           `pulumi:"networkType"`
	RegionId         string                           `pulumi:"regionId"`
	Replication      string                           `pulumi:"replication"`
	Shards           []GetMongoInstancesInstanceShard `pulumi:"shards"`
	Status           string                           `pulumi:"status"`
	Storage          int                              `pulumi:"storage"`
	Tags             map[string]string                `pulumi:"tags"`
}

// GetMongoInstancesInstanceInput is an input type that accepts GetMongoInstancesInstanceArgs and GetMongoInstancesInstanceOutput values.
// You can construct a concrete instance of `GetMongoInstancesInstanceInput` via:
//
//	GetMongoInstancesInstanceArgs{...}
type GetMongoInstancesInstanceInput interface {
	pulumi.Input

	ToGetMongoInstancesInstanceOutput() GetMongoInstancesInstanceOutput
	ToGetMongoInstancesInstanceOutputWithContext(context.Context) GetMongoInstancesInstanceOutput
}

type GetMongoInstancesInstanceArgs struct {
	AvailabilityZone pulumi.StringInput                       `pulumi:"availabilityZone"`
	ChargeType       pulumi.StringInput                       `pulumi:"chargeType"`
	CreationTime     pulumi.StringInput                       `pulumi:"creationTime"`
	Engine           pulumi.StringInput                       `pulumi:"engine"`
	EngineVersion    pulumi.StringInput                       `pulumi:"engineVersion"`
	ExpirationTime   pulumi.StringInput                       `pulumi:"expirationTime"`
	Id               pulumi.StringInput                       `pulumi:"id"`
	InstanceClass    pulumi.StringInput                       `pulumi:"instanceClass"`
	InstanceType     pulumi.StringInput                       `pulumi:"instanceType"`
	LockMode         pulumi.StringInput                       `pulumi:"lockMode"`
	Mongos           GetMongoInstancesInstanceMongoArrayInput `pulumi:"mongos"`
	Name             pulumi.StringInput                       `pulumi:"name"`
	NetworkType      pulumi.StringInput                       `pulumi:"networkType"`
	RegionId         pulumi.StringInput                       `pulumi:"regionId"`
	Replication      pulumi.StringInput                       `pulumi:"replication"`
	Shards           GetMongoInstancesInstanceShardArrayInput `pulumi:"shards"`
	Status           pulumi.StringInput                       `pulumi:"status"`
	Storage          pulumi.IntInput                          `pulumi:"storage"`
	Tags             pulumi.StringMapInput                    `pulumi:"tags"`
}

func (GetMongoInstancesInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMongoInstancesInstance)(nil)).Elem()
}

func (i GetMongoInstancesInstanceArgs) ToGetMongoInstancesInstanceOutput() GetMongoInstancesInstanceOutput {
	return i.ToGetMongoInstancesInstanceOutputWithContext(context.Background())
}

func (i GetMongoInstancesInstanceArgs) ToGetMongoInstancesInstanceOutputWithContext(ctx context.Context) GetMongoInstancesInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMongoInstancesInstanceOutput)
}

// GetMongoInstancesInstanceArrayInput is an input type that accepts GetMongoInstancesInstanceArray and GetMongoInstancesInstanceArrayOutput values.
// You can construct a concrete instance of `GetMongoInstancesInstanceArrayInput` via:
//
//	GetMongoInstancesInstanceArray{ GetMongoInstancesInstanceArgs{...} }
type GetMongoInstancesInstanceArrayInput interface {
	pulumi.Input

	ToGetMongoInstancesInstanceArrayOutput() GetMongoInstancesInstanceArrayOutput
	ToGetMongoInstancesInstanceArrayOutputWithContext(context.Context) GetMongoInstancesInstanceArrayOutput
}

type GetMongoInstancesInstanceArray []GetMongoInstancesInstanceInput

func (GetMongoInstancesInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMongoInstancesInstance)(nil)).Elem()
}

func (i GetMongoInstancesInstanceArray) ToGetMongoInstancesInstanceArrayOutput() GetMongoInstancesInstanceArrayOutput {
	return i.ToGetMongoInstancesInstanceArrayOutputWithContext(context.Background())
}

func (i GetMongoInstancesInstanceArray) ToGetMongoInstancesInstanceArrayOutputWithContext(ctx context.Context) GetMongoInstancesInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMongoInstancesInstanceArrayOutput)
}

type GetMongoInstancesInstanceOutput struct{ *pulumi.OutputState }

func (GetMongoInstancesInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMongoInstancesInstance)(nil)).Elem()
}

func (o GetMongoInstancesInstanceOutput) ToGetMongoInstancesInstanceOutput() GetMongoInstancesInstanceOutput {
	return o
}

func (o GetMongoInstancesInstanceOutput) ToGetMongoInstancesInstanceOutputWithContext(ctx context.Context) GetMongoInstancesInstanceOutput {
	return o
}

func (o GetMongoInstancesInstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.ChargeType }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.Engine }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.EngineVersion }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) InstanceClass() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.InstanceClass }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) LockMode() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.LockMode }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) Mongos() GetMongoInstancesInstanceMongoArrayOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) []GetMongoInstancesInstanceMongo { return v.Mongos }).(GetMongoInstancesInstanceMongoArrayOutput)
}

func (o GetMongoInstancesInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.NetworkType }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.RegionId }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) Replication() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.Replication }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) Shards() GetMongoInstancesInstanceShardArrayOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) []GetMongoInstancesInstanceShard { return v.Shards }).(GetMongoInstancesInstanceShardArrayOutput)
}

func (o GetMongoInstancesInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceOutput) Storage() pulumi.IntOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) int { return v.Storage }).(pulumi.IntOutput)
}

func (o GetMongoInstancesInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetMongoInstancesInstance) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type GetMongoInstancesInstanceArrayOutput struct{ *pulumi.OutputState }

func (GetMongoInstancesInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMongoInstancesInstance)(nil)).Elem()
}

func (o GetMongoInstancesInstanceArrayOutput) ToGetMongoInstancesInstanceArrayOutput() GetMongoInstancesInstanceArrayOutput {
	return o
}

func (o GetMongoInstancesInstanceArrayOutput) ToGetMongoInstancesInstanceArrayOutputWithContext(ctx context.Context) GetMongoInstancesInstanceArrayOutput {
	return o
}

func (o GetMongoInstancesInstanceArrayOutput) Index(i pulumi.IntInput) GetMongoInstancesInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetMongoInstancesInstance {
		return vs[0].([]GetMongoInstancesInstance)[vs[1].(int)]
	}).(GetMongoInstancesInstanceOutput)
}

type GetMongoInstancesInstanceMongo struct {
	Class       string `pulumi:"class"`
	Description string `pulumi:"description"`
	NodeId      string `pulumi:"nodeId"`
}

// GetMongoInstancesInstanceMongoInput is an input type that accepts GetMongoInstancesInstanceMongoArgs and GetMongoInstancesInstanceMongoOutput values.
// You can construct a concrete instance of `GetMongoInstancesInstanceMongoInput` via:
//
//	GetMongoInstancesInstanceMongoArgs{...}
type GetMongoInstancesInstanceMongoInput interface {
	pulumi.Input

	ToGetMongoInstancesInstanceMongoOutput() GetMongoInstancesInstanceMongoOutput
	ToGetMongoInstancesInstanceMongoOutputWithContext(context.Context) GetMongoInstancesInstanceMongoOutput
}

type GetMongoInstancesInstanceMongoArgs struct {
	Class       pulumi.StringInput `pulumi:"class"`
	Description pulumi.StringInput `pulumi:"description"`
	NodeId      pulumi.StringInput `pulumi:"nodeId"`
}

func (GetMongoInstancesInstanceMongoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMongoInstancesInstanceMongo)(nil)).Elem()
}

func (i GetMongoInstancesInstanceMongoArgs) ToGetMongoInstancesInstanceMongoOutput() GetMongoInstancesInstanceMongoOutput {
	return i.ToGetMongoInstancesInstanceMongoOutputWithContext(context.Background())
}

func (i GetMongoInstancesInstanceMongoArgs) ToGetMongoInstancesInstanceMongoOutputWithContext(ctx context.Context) GetMongoInstancesInstanceMongoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMongoInstancesInstanceMongoOutput)
}

// GetMongoInstancesInstanceMongoArrayInput is an input type that accepts GetMongoInstancesInstanceMongoArray and GetMongoInstancesInstanceMongoArrayOutput values.
// You can construct a concrete instance of `GetMongoInstancesInstanceMongoArrayInput` via:
//
//	GetMongoInstancesInstanceMongoArray{ GetMongoInstancesInstanceMongoArgs{...} }
type GetMongoInstancesInstanceMongoArrayInput interface {
	pulumi.Input

	ToGetMongoInstancesInstanceMongoArrayOutput() GetMongoInstancesInstanceMongoArrayOutput
	ToGetMongoInstancesInstanceMongoArrayOutputWithContext(context.Context) GetMongoInstancesInstanceMongoArrayOutput
}

type GetMongoInstancesInstanceMongoArray []GetMongoInstancesInstanceMongoInput

func (GetMongoInstancesInstanceMongoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMongoInstancesInstanceMongo)(nil)).Elem()
}

func (i GetMongoInstancesInstanceMongoArray) ToGetMongoInstancesInstanceMongoArrayOutput() GetMongoInstancesInstanceMongoArrayOutput {
	return i.ToGetMongoInstancesInstanceMongoArrayOutputWithContext(context.Background())
}

func (i GetMongoInstancesInstanceMongoArray) ToGetMongoInstancesInstanceMongoArrayOutputWithContext(ctx context.Context) GetMongoInstancesInstanceMongoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMongoInstancesInstanceMongoArrayOutput)
}

type GetMongoInstancesInstanceMongoOutput struct{ *pulumi.OutputState }

func (GetMongoInstancesInstanceMongoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMongoInstancesInstanceMongo)(nil)).Elem()
}

func (o GetMongoInstancesInstanceMongoOutput) ToGetMongoInstancesInstanceMongoOutput() GetMongoInstancesInstanceMongoOutput {
	return o
}

func (o GetMongoInstancesInstanceMongoOutput) ToGetMongoInstancesInstanceMongoOutputWithContext(ctx context.Context) GetMongoInstancesInstanceMongoOutput {
	return o
}

func (o GetMongoInstancesInstanceMongoOutput) Class() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstanceMongo) string { return v.Class }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceMongoOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstanceMongo) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceMongoOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstanceMongo) string { return v.NodeId }).(pulumi.StringOutput)
}

type GetMongoInstancesInstanceMongoArrayOutput struct{ *pulumi.OutputState }

func (GetMongoInstancesInstanceMongoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMongoInstancesInstanceMongo)(nil)).Elem()
}

func (o GetMongoInstancesInstanceMongoArrayOutput) ToGetMongoInstancesInstanceMongoArrayOutput() GetMongoInstancesInstanceMongoArrayOutput {
	return o
}

func (o GetMongoInstancesInstanceMongoArrayOutput) ToGetMongoInstancesInstanceMongoArrayOutputWithContext(ctx context.Context) GetMongoInstancesInstanceMongoArrayOutput {
	return o
}

func (o GetMongoInstancesInstanceMongoArrayOutput) Index(i pulumi.IntInput) GetMongoInstancesInstanceMongoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetMongoInstancesInstanceMongo {
		return vs[0].([]GetMongoInstancesInstanceMongo)[vs[1].(int)]
	}).(GetMongoInstancesInstanceMongoOutput)
}

type GetMongoInstancesInstanceShard struct {
	Class       string `pulumi:"class"`
	Description string `pulumi:"description"`
	NodeId      string `pulumi:"nodeId"`
	Storage     int    `pulumi:"storage"`
}

// GetMongoInstancesInstanceShardInput is an input type that accepts GetMongoInstancesInstanceShardArgs and GetMongoInstancesInstanceShardOutput values.
// You can construct a concrete instance of `GetMongoInstancesInstanceShardInput` via:
//
//	GetMongoInstancesInstanceShardArgs{...}
type GetMongoInstancesInstanceShardInput interface {
	pulumi.Input

	ToGetMongoInstancesInstanceShardOutput() GetMongoInstancesInstanceShardOutput
	ToGetMongoInstancesInstanceShardOutputWithContext(context.Context) GetMongoInstancesInstanceShardOutput
}

type GetMongoInstancesInstanceShardArgs struct {
	Class       pulumi.StringInput `pulumi:"class"`
	Description pulumi.StringInput `pulumi:"description"`
	NodeId      pulumi.StringInput `pulumi:"nodeId"`
	Storage     pulumi.IntInput    `pulumi:"storage"`
}

func (GetMongoInstancesInstanceShardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMongoInstancesInstanceShard)(nil)).Elem()
}

func (i GetMongoInstancesInstanceShardArgs) ToGetMongoInstancesInstanceShardOutput() GetMongoInstancesInstanceShardOutput {
	return i.ToGetMongoInstancesInstanceShardOutputWithContext(context.Background())
}

func (i GetMongoInstancesInstanceShardArgs) ToGetMongoInstancesInstanceShardOutputWithContext(ctx context.Context) GetMongoInstancesInstanceShardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMongoInstancesInstanceShardOutput)
}

// GetMongoInstancesInstanceShardArrayInput is an input type that accepts GetMongoInstancesInstanceShardArray and GetMongoInstancesInstanceShardArrayOutput values.
// You can construct a concrete instance of `GetMongoInstancesInstanceShardArrayInput` via:
//
//	GetMongoInstancesInstanceShardArray{ GetMongoInstancesInstanceShardArgs{...} }
type GetMongoInstancesInstanceShardArrayInput interface {
	pulumi.Input

	ToGetMongoInstancesInstanceShardArrayOutput() GetMongoInstancesInstanceShardArrayOutput
	ToGetMongoInstancesInstanceShardArrayOutputWithContext(context.Context) GetMongoInstancesInstanceShardArrayOutput
}

type GetMongoInstancesInstanceShardArray []GetMongoInstancesInstanceShardInput

func (GetMongoInstancesInstanceShardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMongoInstancesInstanceShard)(nil)).Elem()
}

func (i GetMongoInstancesInstanceShardArray) ToGetMongoInstancesInstanceShardArrayOutput() GetMongoInstancesInstanceShardArrayOutput {
	return i.ToGetMongoInstancesInstanceShardArrayOutputWithContext(context.Background())
}

func (i GetMongoInstancesInstanceShardArray) ToGetMongoInstancesInstanceShardArrayOutputWithContext(ctx context.Context) GetMongoInstancesInstanceShardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMongoInstancesInstanceShardArrayOutput)
}

type GetMongoInstancesInstanceShardOutput struct{ *pulumi.OutputState }

func (GetMongoInstancesInstanceShardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMongoInstancesInstanceShard)(nil)).Elem()
}

func (o GetMongoInstancesInstanceShardOutput) ToGetMongoInstancesInstanceShardOutput() GetMongoInstancesInstanceShardOutput {
	return o
}

func (o GetMongoInstancesInstanceShardOutput) ToGetMongoInstancesInstanceShardOutputWithContext(ctx context.Context) GetMongoInstancesInstanceShardOutput {
	return o
}

func (o GetMongoInstancesInstanceShardOutput) Class() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstanceShard) string { return v.Class }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceShardOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstanceShard) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceShardOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMongoInstancesInstanceShard) string { return v.NodeId }).(pulumi.StringOutput)
}

func (o GetMongoInstancesInstanceShardOutput) Storage() pulumi.IntOutput {
	return o.ApplyT(func(v GetMongoInstancesInstanceShard) int { return v.Storage }).(pulumi.IntOutput)
}

type GetMongoInstancesInstanceShardArrayOutput struct{ *pulumi.OutputState }

func (GetMongoInstancesInstanceShardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMongoInstancesInstanceShard)(nil)).Elem()
}

func (o GetMongoInstancesInstanceShardArrayOutput) ToGetMongoInstancesInstanceShardArrayOutput() GetMongoInstancesInstanceShardArrayOutput {
	return o
}

func (o GetMongoInstancesInstanceShardArrayOutput) ToGetMongoInstancesInstanceShardArrayOutputWithContext(ctx context.Context) GetMongoInstancesInstanceShardArrayOutput {
	return o
}

func (o GetMongoInstancesInstanceShardArrayOutput) Index(i pulumi.IntInput) GetMongoInstancesInstanceShardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetMongoInstancesInstanceShard {
		return vs[0].([]GetMongoInstancesInstanceShard)[vs[1].(int)]
	}).(GetMongoInstancesInstanceShardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetMongoInstancesInstanceInput)(nil)).Elem(), GetMongoInstancesInstanceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetMongoInstancesInstanceArrayInput)(nil)).Elem(), GetMongoInstancesInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetMongoInstancesInstanceMongoInput)(nil)).Elem(), GetMongoInstancesInstanceMongoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetMongoInstancesInstanceMongoArrayInput)(nil)).Elem(), GetMongoInstancesInstanceMongoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetMongoInstancesInstanceShardInput)(nil)).Elem(), GetMongoInstancesInstanceShardArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetMongoInstancesInstanceShardArrayInput)(nil)).Elem(), GetMongoInstancesInstanceShardArray{})
	pulumi.RegisterOutputType(GetMongoInstancesInstanceOutput{})
	pulumi.RegisterOutputType(GetMongoInstancesInstanceArrayOutput{})
	pulumi.RegisterOutputType(GetMongoInstancesInstanceMongoOutput{})
	pulumi.RegisterOutputType(GetMongoInstancesInstanceMongoArrayOutput{})
	pulumi.RegisterOutputType(GetMongoInstancesInstanceShardOutput{})
	pulumi.RegisterOutputType(GetMongoInstancesInstanceShardArrayOutput{})
}
