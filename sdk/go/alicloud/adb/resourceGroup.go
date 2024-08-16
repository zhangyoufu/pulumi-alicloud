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

// Provides a AnalyticDB for MySQL (ADB) Resource Group resource.
//
// For information about AnalyticDB for MySQL (ADB) Resource Group and how to use it, see [What is Resource Group](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/latest/api-doc-adb-2019-03-15-api-doc-createdbresourcegroup).
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
//			_default, err := adb.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetResourceGroups, err := resourcemanager.GetResourceGroups(ctx, &resourcemanager.GetResourceGroupsArgs{
//				Status: pulumi.StringRef("OK"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				ZoneId:      pulumi.String(_default.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultDBCluster, err := adb.NewDBCluster(ctx, "default", &adb.DBClusterArgs{
//				ComputeResource:   pulumi.String("48Core192GB"),
//				DbClusterCategory: pulumi.String("MixedStorage"),
//				DbClusterVersion:  pulumi.String("3.0"),
//				DbNodeClass:       pulumi.String("E32"),
//				DbNodeStorage:     pulumi.Int(100),
//				Description:       pulumi.String(name),
//				ElasticIoResource: pulumi.Int(1),
//				MaintainTime:      pulumi.String("04:00Z-05:00Z"),
//				Mode:              pulumi.String("flexible"),
//				PaymentType:       pulumi.String("PayAsYouGo"),
//				ResourceGroupId:   pulumi.String(defaultGetResourceGroups.Ids[0]),
//				SecurityIps: pulumi.StringArray{
//					pulumi.String("10.168.1.12"),
//					pulumi.String("10.168.1.11"),
//				},
//				VpcId:     defaultNetwork.ID(),
//				VswitchId: defaultSwitch.ID(),
//				ZoneId:    pulumi.String(_default.Zones[0].Id),
//				Tags: pulumi.StringMap{
//					"Created": pulumi.String("TF"),
//					"For":     pulumi.String("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = adb.NewResourceGroup(ctx, "default", &adb.ResourceGroupArgs{
//				GroupName:   pulumi.String("TF_EXAMPLE"),
//				GroupType:   pulumi.String("batch"),
//				NodeNum:     pulumi.Int(0),
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

	// The time when the resource group was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The ID of the DBCluster.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// The name of the resource group. The `groupName` can be up to 255 characters in length and can contain digits, uppercase letters, hyphens (-), and underscores (_). It must start with a digit or uppercase letter.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The query execution mode. Default value: `interactive`. Valid values: `interactive`, `batch`.
	GroupType pulumi.StringOutput `pulumi:"groupType"`
	// The number of nodes.
	NodeNum pulumi.IntPtrOutput `pulumi:"nodeNum"`
	// The time when the resource group was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The database accounts that are associated with the resource group.
	User pulumi.StringOutput `pulumi:"user"`
	// The database accounts with which to associate the resource group.
	Users pulumi.StringArrayOutput `pulumi:"users"`
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
	// The time when the resource group was created.
	CreateTime *string `pulumi:"createTime"`
	// The ID of the DBCluster.
	DbClusterId *string `pulumi:"dbClusterId"`
	// The name of the resource group. The `groupName` can be up to 255 characters in length and can contain digits, uppercase letters, hyphens (-), and underscores (_). It must start with a digit or uppercase letter.
	GroupName *string `pulumi:"groupName"`
	// The query execution mode. Default value: `interactive`. Valid values: `interactive`, `batch`.
	GroupType *string `pulumi:"groupType"`
	// The number of nodes.
	NodeNum *int `pulumi:"nodeNum"`
	// The time when the resource group was updated.
	UpdateTime *string `pulumi:"updateTime"`
	// The database accounts that are associated with the resource group.
	User *string `pulumi:"user"`
	// The database accounts with which to associate the resource group.
	Users []string `pulumi:"users"`
}

type ResourceGroupState struct {
	// The time when the resource group was created.
	CreateTime pulumi.StringPtrInput
	// The ID of the DBCluster.
	DbClusterId pulumi.StringPtrInput
	// The name of the resource group. The `groupName` can be up to 255 characters in length and can contain digits, uppercase letters, hyphens (-), and underscores (_). It must start with a digit or uppercase letter.
	GroupName pulumi.StringPtrInput
	// The query execution mode. Default value: `interactive`. Valid values: `interactive`, `batch`.
	GroupType pulumi.StringPtrInput
	// The number of nodes.
	NodeNum pulumi.IntPtrInput
	// The time when the resource group was updated.
	UpdateTime pulumi.StringPtrInput
	// The database accounts that are associated with the resource group.
	User pulumi.StringPtrInput
	// The database accounts with which to associate the resource group.
	Users pulumi.StringArrayInput
}

func (ResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupState)(nil)).Elem()
}

type resourceGroupArgs struct {
	// The ID of the DBCluster.
	DbClusterId string `pulumi:"dbClusterId"`
	// The name of the resource group. The `groupName` can be up to 255 characters in length and can contain digits, uppercase letters, hyphens (-), and underscores (_). It must start with a digit or uppercase letter.
	GroupName string `pulumi:"groupName"`
	// The query execution mode. Default value: `interactive`. Valid values: `interactive`, `batch`.
	GroupType *string `pulumi:"groupType"`
	// The number of nodes.
	NodeNum *int `pulumi:"nodeNum"`
	// The database accounts with which to associate the resource group.
	Users []string `pulumi:"users"`
}

// The set of arguments for constructing a ResourceGroup resource.
type ResourceGroupArgs struct {
	// The ID of the DBCluster.
	DbClusterId pulumi.StringInput
	// The name of the resource group. The `groupName` can be up to 255 characters in length and can contain digits, uppercase letters, hyphens (-), and underscores (_). It must start with a digit or uppercase letter.
	GroupName pulumi.StringInput
	// The query execution mode. Default value: `interactive`. Valid values: `interactive`, `batch`.
	GroupType pulumi.StringPtrInput
	// The number of nodes.
	NodeNum pulumi.IntPtrInput
	// The database accounts with which to associate the resource group.
	Users pulumi.StringArrayInput
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

// The time when the resource group was created.
func (o ResourceGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The ID of the DBCluster.
func (o ResourceGroupOutput) DbClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.DbClusterId }).(pulumi.StringOutput)
}

// The name of the resource group. The `groupName` can be up to 255 characters in length and can contain digits, uppercase letters, hyphens (-), and underscores (_). It must start with a digit or uppercase letter.
func (o ResourceGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// The query execution mode. Default value: `interactive`. Valid values: `interactive`, `batch`.
func (o ResourceGroupOutput) GroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.GroupType }).(pulumi.StringOutput)
}

// The number of nodes.
func (o ResourceGroupOutput) NodeNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.IntPtrOutput { return v.NodeNum }).(pulumi.IntPtrOutput)
}

// The time when the resource group was updated.
func (o ResourceGroupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The database accounts that are associated with the resource group.
func (o ResourceGroupOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// The database accounts with which to associate the resource group.
func (o ResourceGroupOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
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
