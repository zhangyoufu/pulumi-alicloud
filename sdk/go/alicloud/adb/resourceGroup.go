// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package adb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Adb Resource Group resource.
//
// For information about Adb Resource Group and how to use it, see [What is Adb Resource Group](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/latest/api-doc-adb-2019-03-15-api-doc-createdbresourcegroup).
//
// > **NOTE:** Available since v1.195.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/adb"
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
//			defaultZones, err := adb.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, &resourcemanager.GetResourceGroupsArgs{
//				Status: pulumi.StringRef("OK"),
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
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultDBCluster, err := adb.NewDBCluster(ctx, "defaultDBCluster", &adb.DBClusterArgs{
//				ComputeResource:   pulumi.String("48Core192GBNEW"),
//				DbClusterCategory: pulumi.String("MixedStorage"),
//				DbClusterVersion:  pulumi.String("3.0"),
//				DbNodeClass:       pulumi.String("E32"),
//				DbNodeCount:       pulumi.Int(1),
//				DbNodeStorage:     pulumi.Int(100),
//				Description:       pulumi.String(name),
//				ElasticIoResource: pulumi.Int(1),
//				MaintainTime:      pulumi.String("04:00Z-05:00Z"),
//				Mode:              pulumi.String("flexible"),
//				PaymentType:       pulumi.String("PayAsYouGo"),
//				ResourceGroupId:   *pulumi.String(defaultResourceGroups.Ids[0]),
//				SecurityIps: pulumi.StringArray{
//					pulumi.String("10.168.1.12"),
//					pulumi.String("10.168.1.11"),
//				},
//				VpcId:     defaultNetwork.ID(),
//				VswitchId: defaultSwitch.ID(),
//				ZoneId:    *pulumi.String(defaultZones.Zones[0].Id),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = adb.NewResourceGroup(ctx, "defaultResourceGroup", &adb.ResourceGroupArgs{
//				GroupName:   pulumi.String("TF_EXAMPLE"),
//				GroupType:   pulumi.String("batch"),
//				NodeNum:     pulumi.Int(1),
//				DbClusterId: defaultDBCluster.ID(),
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
// Adb Resource Group can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:adb/resourceGroup:ResourceGroup example <db_cluster_id>:<group_name>
// ```
type ResourceGroup struct {
	pulumi.CustomResourceState

	// Creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// DB cluster id.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// Query type, value description:
	// * **etl**: Batch query mode.
	// * **interactive**: interactive Query mode.
	// * **default_type**: the default query mode.
	GroupType pulumi.StringOutput `pulumi:"groupType"`
	// The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
	NodeNum pulumi.IntOutput `pulumi:"nodeNum"`
	// Update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Binding User.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewResourceGroup(ctx *pulumi.Context,
	name string, args *ResourceGroupArgs, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbClusterId == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterId'")
	}
	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceGroup
	err := ctx.RegisterResource("alicloud:adb/resourceGroup:ResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGroup gets an existing ResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupState, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	var resource ResourceGroup
	err := ctx.ReadResource("alicloud:adb/resourceGroup:ResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGroup resources.
type resourceGroupState struct {
	// Creation time.
	CreateTime *string `pulumi:"createTime"`
	// DB cluster id.
	DbClusterId *string `pulumi:"dbClusterId"`
	// The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
	GroupName *string `pulumi:"groupName"`
	// Query type, value description:
	// * **etl**: Batch query mode.
	// * **interactive**: interactive Query mode.
	// * **default_type**: the default query mode.
	GroupType *string `pulumi:"groupType"`
	// The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
	NodeNum *int `pulumi:"nodeNum"`
	// Update time.
	UpdateTime *string `pulumi:"updateTime"`
	// Binding User.
	User *string `pulumi:"user"`
}

type ResourceGroupState struct {
	// Creation time.
	CreateTime pulumi.StringPtrInput
	// DB cluster id.
	DbClusterId pulumi.StringPtrInput
	// The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
	GroupName pulumi.StringPtrInput
	// Query type, value description:
	// * **etl**: Batch query mode.
	// * **interactive**: interactive Query mode.
	// * **default_type**: the default query mode.
	GroupType pulumi.StringPtrInput
	// The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
	NodeNum pulumi.IntPtrInput
	// Update time.
	UpdateTime pulumi.StringPtrInput
	// Binding User.
	User pulumi.StringPtrInput
}

func (ResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupState)(nil)).Elem()
}

type resourceGroupArgs struct {
	// DB cluster id.
	DbClusterId string `pulumi:"dbClusterId"`
	// The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
	GroupName string `pulumi:"groupName"`
	// Query type, value description:
	// * **etl**: Batch query mode.
	// * **interactive**: interactive Query mode.
	// * **default_type**: the default query mode.
	GroupType *string `pulumi:"groupType"`
	// The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
	NodeNum *int `pulumi:"nodeNum"`
}

// The set of arguments for constructing a ResourceGroup resource.
type ResourceGroupArgs struct {
	// DB cluster id.
	DbClusterId pulumi.StringInput
	// The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
	GroupName pulumi.StringInput
	// Query type, value description:
	// * **etl**: Batch query mode.
	// * **interactive**: interactive Query mode.
	// * **default_type**: the default query mode.
	GroupType pulumi.StringPtrInput
	// The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
	NodeNum pulumi.IntPtrInput
}

func (ResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupArgs)(nil)).Elem()
}

type ResourceGroupInput interface {
	pulumi.Input

	ToResourceGroupOutput() ResourceGroupOutput
	ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput
}

func (*ResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroup)(nil)).Elem()
}

func (i *ResourceGroup) ToResourceGroupOutput() ResourceGroupOutput {
	return i.ToResourceGroupOutputWithContext(context.Background())
}

func (i *ResourceGroup) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupOutput)
}

// ResourceGroupArrayInput is an input type that accepts ResourceGroupArray and ResourceGroupArrayOutput values.
// You can construct a concrete instance of `ResourceGroupArrayInput` via:
//
//	ResourceGroupArray{ ResourceGroupArgs{...} }
type ResourceGroupArrayInput interface {
	pulumi.Input

	ToResourceGroupArrayOutput() ResourceGroupArrayOutput
	ToResourceGroupArrayOutputWithContext(context.Context) ResourceGroupArrayOutput
}

type ResourceGroupArray []ResourceGroupInput

func (ResourceGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceGroup)(nil)).Elem()
}

func (i ResourceGroupArray) ToResourceGroupArrayOutput() ResourceGroupArrayOutput {
	return i.ToResourceGroupArrayOutputWithContext(context.Background())
}

func (i ResourceGroupArray) ToResourceGroupArrayOutputWithContext(ctx context.Context) ResourceGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupArrayOutput)
}

// ResourceGroupMapInput is an input type that accepts ResourceGroupMap and ResourceGroupMapOutput values.
// You can construct a concrete instance of `ResourceGroupMapInput` via:
//
//	ResourceGroupMap{ "key": ResourceGroupArgs{...} }
type ResourceGroupMapInput interface {
	pulumi.Input

	ToResourceGroupMapOutput() ResourceGroupMapOutput
	ToResourceGroupMapOutputWithContext(context.Context) ResourceGroupMapOutput
}

type ResourceGroupMap map[string]ResourceGroupInput

func (ResourceGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceGroup)(nil)).Elem()
}

func (i ResourceGroupMap) ToResourceGroupMapOutput() ResourceGroupMapOutput {
	return i.ToResourceGroupMapOutputWithContext(context.Background())
}

func (i ResourceGroupMap) ToResourceGroupMapOutputWithContext(ctx context.Context) ResourceGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupMapOutput)
}

type ResourceGroupOutput struct{ *pulumi.OutputState }

func (ResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroup)(nil)).Elem()
}

func (o ResourceGroupOutput) ToResourceGroupOutput() ResourceGroupOutput {
	return o
}

func (o ResourceGroupOutput) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return o
}

// Creation time.
func (o ResourceGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// DB cluster id.
func (o ResourceGroupOutput) DbClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.DbClusterId }).(pulumi.StringOutput)
}

// The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
func (o ResourceGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// Query type, value description:
// * **etl**: Batch query mode.
// * **interactive**: interactive Query mode.
// * **default_type**: the default query mode.
func (o ResourceGroupOutput) GroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.GroupType }).(pulumi.StringOutput)
}

// The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
func (o ResourceGroupOutput) NodeNum() pulumi.IntOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.IntOutput { return v.NodeNum }).(pulumi.IntOutput)
}

// Update time.
func (o ResourceGroupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Binding User.
func (o ResourceGroupOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type ResourceGroupArrayOutput struct{ *pulumi.OutputState }

func (ResourceGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceGroup)(nil)).Elem()
}

func (o ResourceGroupArrayOutput) ToResourceGroupArrayOutput() ResourceGroupArrayOutput {
	return o
}

func (o ResourceGroupArrayOutput) ToResourceGroupArrayOutputWithContext(ctx context.Context) ResourceGroupArrayOutput {
	return o
}

func (o ResourceGroupArrayOutput) Index(i pulumi.IntInput) ResourceGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceGroup {
		return vs[0].([]*ResourceGroup)[vs[1].(int)]
	}).(ResourceGroupOutput)
}

type ResourceGroupMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceGroup)(nil)).Elem()
}

func (o ResourceGroupMapOutput) ToResourceGroupMapOutput() ResourceGroupMapOutput {
	return o
}

func (o ResourceGroupMapOutput) ToResourceGroupMapOutputWithContext(ctx context.Context) ResourceGroupMapOutput {
	return o
}

func (o ResourceGroupMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceGroup {
		return vs[0].(map[string]*ResourceGroup)[vs[1].(string)]
	}).(ResourceGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupInput)(nil)).Elem(), &ResourceGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupArrayInput)(nil)).Elem(), ResourceGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupMapInput)(nil)).Elem(), ResourceGroupMap{})
	pulumi.RegisterOutputType(ResourceGroupOutput{})
	pulumi.RegisterOutputType(ResourceGroupArrayOutput{})
	pulumi.RegisterOutputType(ResourceGroupMapOutput{})
}
