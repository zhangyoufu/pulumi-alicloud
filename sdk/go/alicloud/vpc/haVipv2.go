// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vpc Ha Vip resource. Highly available virtual IP
//
// For information about Vpc Ha Vip and how to use it, see [What is Ha Vip](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createhavip).
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
//			name := "tf-testacc-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultVpc, err := vpc.NewNetwork(ctx, "defaultVpc", &vpc.NetworkArgs{
//				Description: pulumi.String("tf-test-acc-vpc"),
//				VpcName:     pulumi.String(name),
//				CidrBlock:   pulumi.String("192.168.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultVswitch, err := vpc.NewSwitch(ctx, "defaultVswitch", &vpc.SwitchArgs{
//				VpcId:       defaultVpc.ID(),
//				CidrBlock:   pulumi.String("192.168.0.0/21"),
//				VswitchName: pulumi.Sprintf("%v1", name),
//				ZoneId:      pulumi.String(_default.Zones[0].Id),
//				Description: pulumi.String("tf-testacc-vswitch"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultRg, err := resourcemanager.NewResourceGroup(ctx, "defaultRg", &resourcemanager.ResourceGroupArgs{
//				DisplayName:       pulumi.String("tf-testacc-rg819"),
//				ResourceGroupName: pulumi.Sprintf("%v2", name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = resourcemanager.NewResourceGroup(ctx, "changeRg", &resourcemanager.ResourceGroupArgs{
//				DisplayName:       pulumi.String("tf-testacc-changerg670"),
//				ResourceGroupName: pulumi.Sprintf("%v3", name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewHaVipv2(ctx, "default", &vpc.HaVipv2Args{
//				Description:     pulumi.String("test"),
//				VswitchId:       defaultVswitch.ID(),
//				HaVipName:       pulumi.String(name),
//				IpAddress:       pulumi.String("192.168.1.101"),
//				ResourceGroupId: defaultRg.ID(),
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
// Vpc Ha Vip can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/haVipv2:HaVipv2 example <id>
// ```
type HaVipv2 struct {
	pulumi.CustomResourceState

	// EIP bound to HaVip.
	AssociatedEipAddresses pulumi.StringArrayOutput `pulumi:"associatedEipAddresses"`
	// The type of the instance that is bound to the HaVip. Value:-**EcsInstance**: ECS instance.-**NetworkInterface**: ENI instance.
	AssociatedInstanceType pulumi.StringOutput `pulumi:"associatedInstanceType"`
	// An ECS instance that is bound to HaVip.
	AssociatedInstances pulumi.StringArrayOutput `pulumi:"associatedInstances"`
	// The creation time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the HaVip instance. The length is 2 to 256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the resource.
	HaVipId pulumi.StringOutput `pulumi:"haVipId"`
	// The name of the HaVip instance.
	HaVipName pulumi.StringOutput `pulumi:"haVipName"`
	// Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
	//
	// Deprecated: Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
	HavipName pulumi.StringOutput `pulumi:"havipName"`
	// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The primary instance ID bound to HaVip.
	MasterInstanceId pulumi.StringOutput `pulumi:"masterInstanceId"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The status of this resource instance.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags of HaVip.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The VPC ID to which the HaVip instance belongs.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The switch ID to which the HaVip instance belongs.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewHaVipv2 registers a new resource with the given unique name, arguments, and options.
func NewHaVipv2(ctx *pulumi.Context,
	name string, args *HaVipv2Args, opts ...pulumi.ResourceOption) (*HaVipv2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HaVipv2
	err := ctx.RegisterResource("alicloud:vpc/haVipv2:HaVipv2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHaVipv2 gets an existing HaVipv2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHaVipv2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HaVipv2State, opts ...pulumi.ResourceOption) (*HaVipv2, error) {
	var resource HaVipv2
	err := ctx.ReadResource("alicloud:vpc/haVipv2:HaVipv2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HaVipv2 resources.
type haVipv2State struct {
	// EIP bound to HaVip.
	AssociatedEipAddresses []string `pulumi:"associatedEipAddresses"`
	// The type of the instance that is bound to the HaVip. Value:-**EcsInstance**: ECS instance.-**NetworkInterface**: ENI instance.
	AssociatedInstanceType *string `pulumi:"associatedInstanceType"`
	// An ECS instance that is bound to HaVip.
	AssociatedInstances []string `pulumi:"associatedInstances"`
	// The creation time of the resource.
	CreateTime *string `pulumi:"createTime"`
	// The description of the HaVip instance. The length is 2 to 256 characters.
	Description *string `pulumi:"description"`
	// The ID of the resource.
	HaVipId *string `pulumi:"haVipId"`
	// The name of the HaVip instance.
	HaVipName *string `pulumi:"haVipName"`
	// Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
	//
	// Deprecated: Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
	HavipName *string `pulumi:"havipName"`
	// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
	IpAddress *string `pulumi:"ipAddress"`
	// The primary instance ID bound to HaVip.
	MasterInstanceId *string `pulumi:"masterInstanceId"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of this resource instance.
	Status *string `pulumi:"status"`
	// The tags of HaVip.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID to which the HaVip instance belongs.
	VpcId *string `pulumi:"vpcId"`
	// The switch ID to which the HaVip instance belongs.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VswitchId *string `pulumi:"vswitchId"`
}

type HaVipv2State struct {
	// EIP bound to HaVip.
	AssociatedEipAddresses pulumi.StringArrayInput
	// The type of the instance that is bound to the HaVip. Value:-**EcsInstance**: ECS instance.-**NetworkInterface**: ENI instance.
	AssociatedInstanceType pulumi.StringPtrInput
	// An ECS instance that is bound to HaVip.
	AssociatedInstances pulumi.StringArrayInput
	// The creation time of the resource.
	CreateTime pulumi.StringPtrInput
	// The description of the HaVip instance. The length is 2 to 256 characters.
	Description pulumi.StringPtrInput
	// The ID of the resource.
	HaVipId pulumi.StringPtrInput
	// The name of the HaVip instance.
	HaVipName pulumi.StringPtrInput
	// Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
	//
	// Deprecated: Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
	HavipName pulumi.StringPtrInput
	// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
	IpAddress pulumi.StringPtrInput
	// The primary instance ID bound to HaVip.
	MasterInstanceId pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The status of this resource instance.
	Status pulumi.StringPtrInput
	// The tags of HaVip.
	Tags pulumi.StringMapInput
	// The VPC ID to which the HaVip instance belongs.
	VpcId pulumi.StringPtrInput
	// The switch ID to which the HaVip instance belongs.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VswitchId pulumi.StringPtrInput
}

func (HaVipv2State) ElementType() reflect.Type {
	return reflect.TypeOf((*haVipv2State)(nil)).Elem()
}

type haVipv2Args struct {
	// The description of the HaVip instance. The length is 2 to 256 characters.
	Description *string `pulumi:"description"`
	// The name of the HaVip instance.
	HaVipName *string `pulumi:"haVipName"`
	// Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
	//
	// Deprecated: Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
	HavipName *string `pulumi:"havipName"`
	// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
	IpAddress *string `pulumi:"ipAddress"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The tags of HaVip.
	Tags map[string]string `pulumi:"tags"`
	// The switch ID to which the HaVip instance belongs.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a HaVipv2 resource.
type HaVipv2Args struct {
	// The description of the HaVip instance. The length is 2 to 256 characters.
	Description pulumi.StringPtrInput
	// The name of the HaVip instance.
	HaVipName pulumi.StringPtrInput
	// Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
	//
	// Deprecated: Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
	HavipName pulumi.StringPtrInput
	// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
	IpAddress pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The tags of HaVip.
	Tags pulumi.StringMapInput
	// The switch ID to which the HaVip instance belongs.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VswitchId pulumi.StringInput
}

func (HaVipv2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*haVipv2Args)(nil)).Elem()
}

type HaVipv2Input interface {
	pulumi.Input

	ToHaVipv2Output() HaVipv2Output
	ToHaVipv2OutputWithContext(ctx context.Context) HaVipv2Output
}

func (*HaVipv2) ElementType() reflect.Type {
	return reflect.TypeOf((**HaVipv2)(nil)).Elem()
}

func (i *HaVipv2) ToHaVipv2Output() HaVipv2Output {
	return i.ToHaVipv2OutputWithContext(context.Background())
}

func (i *HaVipv2) ToHaVipv2OutputWithContext(ctx context.Context) HaVipv2Output {
	return pulumi.ToOutputWithContext(ctx, i).(HaVipv2Output)
}

// HaVipv2ArrayInput is an input type that accepts HaVipv2Array and HaVipv2ArrayOutput values.
// You can construct a concrete instance of `HaVipv2ArrayInput` via:
//
//	HaVipv2Array{ HaVipv2Args{...} }
type HaVipv2ArrayInput interface {
	pulumi.Input

	ToHaVipv2ArrayOutput() HaVipv2ArrayOutput
	ToHaVipv2ArrayOutputWithContext(context.Context) HaVipv2ArrayOutput
}

type HaVipv2Array []HaVipv2Input

func (HaVipv2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HaVipv2)(nil)).Elem()
}

func (i HaVipv2Array) ToHaVipv2ArrayOutput() HaVipv2ArrayOutput {
	return i.ToHaVipv2ArrayOutputWithContext(context.Background())
}

func (i HaVipv2Array) ToHaVipv2ArrayOutputWithContext(ctx context.Context) HaVipv2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaVipv2ArrayOutput)
}

// HaVipv2MapInput is an input type that accepts HaVipv2Map and HaVipv2MapOutput values.
// You can construct a concrete instance of `HaVipv2MapInput` via:
//
//	HaVipv2Map{ "key": HaVipv2Args{...} }
type HaVipv2MapInput interface {
	pulumi.Input

	ToHaVipv2MapOutput() HaVipv2MapOutput
	ToHaVipv2MapOutputWithContext(context.Context) HaVipv2MapOutput
}

type HaVipv2Map map[string]HaVipv2Input

func (HaVipv2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HaVipv2)(nil)).Elem()
}

func (i HaVipv2Map) ToHaVipv2MapOutput() HaVipv2MapOutput {
	return i.ToHaVipv2MapOutputWithContext(context.Background())
}

func (i HaVipv2Map) ToHaVipv2MapOutputWithContext(ctx context.Context) HaVipv2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaVipv2MapOutput)
}

type HaVipv2Output struct{ *pulumi.OutputState }

func (HaVipv2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**HaVipv2)(nil)).Elem()
}

func (o HaVipv2Output) ToHaVipv2Output() HaVipv2Output {
	return o
}

func (o HaVipv2Output) ToHaVipv2OutputWithContext(ctx context.Context) HaVipv2Output {
	return o
}

// EIP bound to HaVip.
func (o HaVipv2Output) AssociatedEipAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringArrayOutput { return v.AssociatedEipAddresses }).(pulumi.StringArrayOutput)
}

// The type of the instance that is bound to the HaVip. Value:-**EcsInstance**: ECS instance.-**NetworkInterface**: ENI instance.
func (o HaVipv2Output) AssociatedInstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.AssociatedInstanceType }).(pulumi.StringOutput)
}

// An ECS instance that is bound to HaVip.
func (o HaVipv2Output) AssociatedInstances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringArrayOutput { return v.AssociatedInstances }).(pulumi.StringArrayOutput)
}

// The creation time of the resource.
func (o HaVipv2Output) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the HaVip instance. The length is 2 to 256 characters.
func (o HaVipv2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the resource.
func (o HaVipv2Output) HaVipId() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.HaVipId }).(pulumi.StringOutput)
}

// The name of the HaVip instance.
func (o HaVipv2Output) HaVipName() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.HaVipName }).(pulumi.StringOutput)
}

// Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
//
// Deprecated: Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
func (o HaVipv2Output) HavipName() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.HavipName }).(pulumi.StringOutput)
}

// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
func (o HaVipv2Output) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The primary instance ID bound to HaVip.
func (o HaVipv2Output) MasterInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.MasterInstanceId }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o HaVipv2Output) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of this resource instance.
func (o HaVipv2Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags of HaVip.
func (o HaVipv2Output) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The VPC ID to which the HaVip instance belongs.
func (o HaVipv2Output) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The switch ID to which the HaVip instance belongs.
//
// The following arguments will be discarded. Please use new fields as soon as possible:
func (o HaVipv2Output) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVipv2) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type HaVipv2ArrayOutput struct{ *pulumi.OutputState }

func (HaVipv2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HaVipv2)(nil)).Elem()
}

func (o HaVipv2ArrayOutput) ToHaVipv2ArrayOutput() HaVipv2ArrayOutput {
	return o
}

func (o HaVipv2ArrayOutput) ToHaVipv2ArrayOutputWithContext(ctx context.Context) HaVipv2ArrayOutput {
	return o
}

func (o HaVipv2ArrayOutput) Index(i pulumi.IntInput) HaVipv2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HaVipv2 {
		return vs[0].([]*HaVipv2)[vs[1].(int)]
	}).(HaVipv2Output)
}

type HaVipv2MapOutput struct{ *pulumi.OutputState }

func (HaVipv2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HaVipv2)(nil)).Elem()
}

func (o HaVipv2MapOutput) ToHaVipv2MapOutput() HaVipv2MapOutput {
	return o
}

func (o HaVipv2MapOutput) ToHaVipv2MapOutputWithContext(ctx context.Context) HaVipv2MapOutput {
	return o
}

func (o HaVipv2MapOutput) MapIndex(k pulumi.StringInput) HaVipv2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HaVipv2 {
		return vs[0].(map[string]*HaVipv2)[vs[1].(string)]
	}).(HaVipv2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HaVipv2Input)(nil)).Elem(), &HaVipv2{})
	pulumi.RegisterInputType(reflect.TypeOf((*HaVipv2ArrayInput)(nil)).Elem(), HaVipv2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*HaVipv2MapInput)(nil)).Elem(), HaVipv2Map{})
	pulumi.RegisterOutputType(HaVipv2Output{})
	pulumi.RegisterOutputType(HaVipv2ArrayOutput{})
	pulumi.RegisterOutputType(HaVipv2MapOutput{})
}
