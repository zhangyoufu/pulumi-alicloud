// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide RDS cluster instance to increase node resources.
// > **NOTE:** Available in 1.202.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-testaccrdsdbnodes"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := rds.GetZones(ctx, &rds.GetZonesArgs{
//				Engine:                pulumi.StringRef("MySQL"),
//				EngineVersion:         pulumi.StringRef("8.0"),
//				InstanceChargeType:    pulumi.StringRef("PostPaid"),
//				Category:              pulumi.StringRef("cluster"),
//				DbInstanceStorageType: pulumi.StringRef("cloud_essd"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstanceClasses, err := rds.GetInstanceClasses(ctx, &rds.GetInstanceClassesArgs{
//				ZoneId:                pulumi.StringRef(defaultZones.Zones[0].Id),
//				Engine:                pulumi.StringRef("MySQL"),
//				EngineVersion:         pulumi.StringRef("8.0"),
//				Category:              pulumi.StringRef("cluster"),
//				DbInstanceStorageType: pulumi.StringRef("cloud_essd"),
//				InstanceChargeType:    pulumi.StringRef("PostPaid"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef(fmt.Sprintf("^default-NODELETING$")),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
//				VpcId:  pulumi.StringRef(defaultNetworks.Ids[0]),
//				ZoneId: pulumi.StringRef(defaultZones.Ids[0]),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := rds.NewInstance(ctx, "defaultInstance", &rds.InstanceArgs{
//				Engine:                pulumi.String("MySQL"),
//				EngineVersion:         pulumi.String("8.0"),
//				DbInstanceStorageType: pulumi.String("cloud_essd"),
//				InstanceType:          *pulumi.String(defaultInstanceClasses.InstanceClasses[0].InstanceClass),
//				InstanceStorage:       *pulumi.String(defaultInstanceClasses.InstanceClasses[0].StorageRange.Min),
//				VswitchId:             *pulumi.String(defaultSwitches.Ids[0]),
//				InstanceName:          pulumi.String(name),
//				ZoneId:                *pulumi.String(defaultZones.Ids[0]),
//				ZoneIdSlaveA:          *pulumi.String(defaultZones.Ids[0]),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewDbNode(ctx, "node", &rds.DbNodeArgs{
//				DbInstanceId: defaultInstance.ID(),
//				ClassCode:    defaultInstance.InstanceType,
//				ZoneId:       defaultInstance.ZoneId,
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
// RDS MySQL database cluster node agent function can be imported using id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:rds/dbNode:DbNode example <db_instance_id>:<node_id>
//
// ```
type DbNode struct {
	pulumi.CustomResourceState

	// The specification information of the node.
	ClassCode pulumi.StringOutput `pulumi:"classCode"`
	// The Id of instance that can run database.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// The ID of the node.
	NodeId pulumi.StringOutput `pulumi:"nodeId"`
	// The region ID of the node.
	NodeRegionId pulumi.StringOutput `pulumi:"nodeRegionId"`
	// The role of node.
	NodeRole pulumi.StringOutput `pulumi:"nodeRole"`
	// The zone ID of the node.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewDbNode registers a new resource with the given unique name, arguments, and options.
func NewDbNode(ctx *pulumi.Context,
	name string, args *DbNodeArgs, opts ...pulumi.ResourceOption) (*DbNode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClassCode == nil {
		return nil, errors.New("invalid value for required argument 'ClassCode'")
	}
	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource DbNode
	err := ctx.RegisterResource("alicloud:rds/dbNode:DbNode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbNode gets an existing DbNode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbNodeState, opts ...pulumi.ResourceOption) (*DbNode, error) {
	var resource DbNode
	err := ctx.ReadResource("alicloud:rds/dbNode:DbNode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbNode resources.
type dbNodeState struct {
	// The specification information of the node.
	ClassCode *string `pulumi:"classCode"`
	// The Id of instance that can run database.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// The ID of the node.
	NodeId *string `pulumi:"nodeId"`
	// The region ID of the node.
	NodeRegionId *string `pulumi:"nodeRegionId"`
	// The role of node.
	NodeRole *string `pulumi:"nodeRole"`
	// The zone ID of the node.
	ZoneId *string `pulumi:"zoneId"`
}

type DbNodeState struct {
	// The specification information of the node.
	ClassCode pulumi.StringPtrInput
	// The Id of instance that can run database.
	DbInstanceId pulumi.StringPtrInput
	// The ID of the node.
	NodeId pulumi.StringPtrInput
	// The region ID of the node.
	NodeRegionId pulumi.StringPtrInput
	// The role of node.
	NodeRole pulumi.StringPtrInput
	// The zone ID of the node.
	ZoneId pulumi.StringPtrInput
}

func (DbNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbNodeState)(nil)).Elem()
}

type dbNodeArgs struct {
	// The specification information of the node.
	ClassCode string `pulumi:"classCode"`
	// The Id of instance that can run database.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// The zone ID of the node.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a DbNode resource.
type DbNodeArgs struct {
	// The specification information of the node.
	ClassCode pulumi.StringInput
	// The Id of instance that can run database.
	DbInstanceId pulumi.StringInput
	// The zone ID of the node.
	ZoneId pulumi.StringInput
}

func (DbNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbNodeArgs)(nil)).Elem()
}

type DbNodeInput interface {
	pulumi.Input

	ToDbNodeOutput() DbNodeOutput
	ToDbNodeOutputWithContext(ctx context.Context) DbNodeOutput
}

func (*DbNode) ElementType() reflect.Type {
	return reflect.TypeOf((**DbNode)(nil)).Elem()
}

func (i *DbNode) ToDbNodeOutput() DbNodeOutput {
	return i.ToDbNodeOutputWithContext(context.Background())
}

func (i *DbNode) ToDbNodeOutputWithContext(ctx context.Context) DbNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbNodeOutput)
}

// DbNodeArrayInput is an input type that accepts DbNodeArray and DbNodeArrayOutput values.
// You can construct a concrete instance of `DbNodeArrayInput` via:
//
//	DbNodeArray{ DbNodeArgs{...} }
type DbNodeArrayInput interface {
	pulumi.Input

	ToDbNodeArrayOutput() DbNodeArrayOutput
	ToDbNodeArrayOutputWithContext(context.Context) DbNodeArrayOutput
}

type DbNodeArray []DbNodeInput

func (DbNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbNode)(nil)).Elem()
}

func (i DbNodeArray) ToDbNodeArrayOutput() DbNodeArrayOutput {
	return i.ToDbNodeArrayOutputWithContext(context.Background())
}

func (i DbNodeArray) ToDbNodeArrayOutputWithContext(ctx context.Context) DbNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbNodeArrayOutput)
}

// DbNodeMapInput is an input type that accepts DbNodeMap and DbNodeMapOutput values.
// You can construct a concrete instance of `DbNodeMapInput` via:
//
//	DbNodeMap{ "key": DbNodeArgs{...} }
type DbNodeMapInput interface {
	pulumi.Input

	ToDbNodeMapOutput() DbNodeMapOutput
	ToDbNodeMapOutputWithContext(context.Context) DbNodeMapOutput
}

type DbNodeMap map[string]DbNodeInput

func (DbNodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbNode)(nil)).Elem()
}

func (i DbNodeMap) ToDbNodeMapOutput() DbNodeMapOutput {
	return i.ToDbNodeMapOutputWithContext(context.Background())
}

func (i DbNodeMap) ToDbNodeMapOutputWithContext(ctx context.Context) DbNodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbNodeMapOutput)
}

type DbNodeOutput struct{ *pulumi.OutputState }

func (DbNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbNode)(nil)).Elem()
}

func (o DbNodeOutput) ToDbNodeOutput() DbNodeOutput {
	return o
}

func (o DbNodeOutput) ToDbNodeOutputWithContext(ctx context.Context) DbNodeOutput {
	return o
}

// The specification information of the node.
func (o DbNodeOutput) ClassCode() pulumi.StringOutput {
	return o.ApplyT(func(v *DbNode) pulumi.StringOutput { return v.ClassCode }).(pulumi.StringOutput)
}

// The Id of instance that can run database.
func (o DbNodeOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbNode) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The ID of the node.
func (o DbNodeOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbNode) pulumi.StringOutput { return v.NodeId }).(pulumi.StringOutput)
}

// The region ID of the node.
func (o DbNodeOutput) NodeRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbNode) pulumi.StringOutput { return v.NodeRegionId }).(pulumi.StringOutput)
}

// The role of node.
func (o DbNodeOutput) NodeRole() pulumi.StringOutput {
	return o.ApplyT(func(v *DbNode) pulumi.StringOutput { return v.NodeRole }).(pulumi.StringOutput)
}

// The zone ID of the node.
func (o DbNodeOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbNode) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type DbNodeArrayOutput struct{ *pulumi.OutputState }

func (DbNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbNode)(nil)).Elem()
}

func (o DbNodeArrayOutput) ToDbNodeArrayOutput() DbNodeArrayOutput {
	return o
}

func (o DbNodeArrayOutput) ToDbNodeArrayOutputWithContext(ctx context.Context) DbNodeArrayOutput {
	return o
}

func (o DbNodeArrayOutput) Index(i pulumi.IntInput) DbNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DbNode {
		return vs[0].([]*DbNode)[vs[1].(int)]
	}).(DbNodeOutput)
}

type DbNodeMapOutput struct{ *pulumi.OutputState }

func (DbNodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbNode)(nil)).Elem()
}

func (o DbNodeMapOutput) ToDbNodeMapOutput() DbNodeMapOutput {
	return o
}

func (o DbNodeMapOutput) ToDbNodeMapOutputWithContext(ctx context.Context) DbNodeMapOutput {
	return o
}

func (o DbNodeMapOutput) MapIndex(k pulumi.StringInput) DbNodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DbNode {
		return vs[0].(map[string]*DbNode)[vs[1].(string)]
	}).(DbNodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbNodeInput)(nil)).Elem(), &DbNode{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbNodeArrayInput)(nil)).Elem(), DbNodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbNodeMapInput)(nil)).Elem(), DbNodeMap{})
	pulumi.RegisterOutputType(DbNodeOutput{})
	pulumi.RegisterOutputType(DbNodeArrayOutput{})
	pulumi.RegisterOutputType(DbNodeMapOutput{})
}
