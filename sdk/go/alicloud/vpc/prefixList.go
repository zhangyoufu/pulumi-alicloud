// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vpc Prefix List resource. This resource is used to create a prefix list.
//
// For information about Vpc Prefix List and how to use it, see [What is Prefix List](https://www.alibabacloud.com/help/zh/virtual-private-cloud/latest/creatvpcprefixlist).
//
// > **NOTE:** Available in v1.182.0+.
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
//			defaultRg, err := resourcemanager.NewResourceGroup(ctx, "defaultRg", &resourcemanager.ResourceGroupArgs{
//				DisplayName:       pulumi.String("tf-testacc-chenyi"),
//				ResourceGroupName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = resourcemanager.NewResourceGroup(ctx, "changeRg", &resourcemanager.ResourceGroupArgs{
//				DisplayName:       pulumi.String("tf-testacc-chenyi-change"),
//				ResourceGroupName: pulumi.Sprintf("%v1", name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewPrefixList(ctx, "default", &vpc.PrefixListArgs{
//				MaxEntries:            pulumi.Int(50),
//				ResourceGroupId:       defaultRg.ID(),
//				PrefixListDescription: pulumi.String("test"),
//				IpVersion:             pulumi.String("IPV4"),
//				PrefixListName:        pulumi.String(name),
//				Entrys: vpc.PrefixListEntryArray{
//					&vpc.PrefixListEntryArgs{
//						Cidr:        pulumi.String("192.168.0.0/16"),
//						Description: pulumi.String("test"),
//					},
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
//
// ## Import
//
// Vpc Prefix List can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/prefixList:PrefixList example <id>
// ```
type PrefixList struct {
	pulumi.CustomResourceState

	// The time when the prefix list was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The CIDR address block list of the prefix list.See the following `Block Entrys`.
	Entrys PrefixListEntryArrayOutput `pulumi:"entrys"`
	// The IP version of the prefix list. Value:-**IPV4**:IPv4 version.-**IPV6**:IPv6 version.
	IpVersion pulumi.StringOutput `pulumi:"ipVersion"`
	// The maximum number of entries for CIDR address blocks in the prefix list.
	MaxEntries pulumi.IntOutput `pulumi:"maxEntries"`
	// The association list information of the prefix list.
	PrefixListAssociations PrefixListPrefixListAssociationArrayOutput `pulumi:"prefixListAssociations"`
	// The description of the prefix list.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	PrefixListDescription pulumi.StringPtrOutput `pulumi:"prefixListDescription"`
	// The ID of the query Prefix List.
	PrefixListId pulumi.StringOutput `pulumi:"prefixListId"`
	// The name of the prefix list. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
	PrefixListName pulumi.StringPtrOutput `pulumi:"prefixListName"`
	// The ID of the resource group to which the PrefixList belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The share type of the prefix list. Value:-**Shared**: indicates that the prefix list is a Shared prefix list.-Null: indicates that the prefix list is not a shared prefix list.
	ShareType pulumi.StringOutput `pulumi:"shareType"`
	// Resource attribute fields that represent the status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags of PrefixList.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewPrefixList registers a new resource with the given unique name, arguments, and options.
func NewPrefixList(ctx *pulumi.Context,
	name string, args *PrefixListArgs, opts ...pulumi.ResourceOption) (*PrefixList, error) {
	if args == nil {
		args = &PrefixListArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrefixList
	err := ctx.RegisterResource("alicloud:vpc/prefixList:PrefixList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrefixList gets an existing PrefixList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrefixList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrefixListState, opts ...pulumi.ResourceOption) (*PrefixList, error) {
	var resource PrefixList
	err := ctx.ReadResource("alicloud:vpc/prefixList:PrefixList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrefixList resources.
type prefixListState struct {
	// The time when the prefix list was created.
	CreateTime *string `pulumi:"createTime"`
	// The CIDR address block list of the prefix list.See the following `Block Entrys`.
	Entrys []PrefixListEntry `pulumi:"entrys"`
	// The IP version of the prefix list. Value:-**IPV4**:IPv4 version.-**IPV6**:IPv6 version.
	IpVersion *string `pulumi:"ipVersion"`
	// The maximum number of entries for CIDR address blocks in the prefix list.
	MaxEntries *int `pulumi:"maxEntries"`
	// The association list information of the prefix list.
	PrefixListAssociations []PrefixListPrefixListAssociation `pulumi:"prefixListAssociations"`
	// The description of the prefix list.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	PrefixListDescription *string `pulumi:"prefixListDescription"`
	// The ID of the query Prefix List.
	PrefixListId *string `pulumi:"prefixListId"`
	// The name of the prefix list. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
	PrefixListName *string `pulumi:"prefixListName"`
	// The ID of the resource group to which the PrefixList belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The share type of the prefix list. Value:-**Shared**: indicates that the prefix list is a Shared prefix list.-Null: indicates that the prefix list is not a shared prefix list.
	ShareType *string `pulumi:"shareType"`
	// Resource attribute fields that represent the status of the resource.
	Status *string `pulumi:"status"`
	// The tags of PrefixList.
	Tags map[string]interface{} `pulumi:"tags"`
}

type PrefixListState struct {
	// The time when the prefix list was created.
	CreateTime pulumi.StringPtrInput
	// The CIDR address block list of the prefix list.See the following `Block Entrys`.
	Entrys PrefixListEntryArrayInput
	// The IP version of the prefix list. Value:-**IPV4**:IPv4 version.-**IPV6**:IPv6 version.
	IpVersion pulumi.StringPtrInput
	// The maximum number of entries for CIDR address blocks in the prefix list.
	MaxEntries pulumi.IntPtrInput
	// The association list information of the prefix list.
	PrefixListAssociations PrefixListPrefixListAssociationArrayInput
	// The description of the prefix list.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	PrefixListDescription pulumi.StringPtrInput
	// The ID of the query Prefix List.
	PrefixListId pulumi.StringPtrInput
	// The name of the prefix list. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
	PrefixListName pulumi.StringPtrInput
	// The ID of the resource group to which the PrefixList belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The share type of the prefix list. Value:-**Shared**: indicates that the prefix list is a Shared prefix list.-Null: indicates that the prefix list is not a shared prefix list.
	ShareType pulumi.StringPtrInput
	// Resource attribute fields that represent the status of the resource.
	Status pulumi.StringPtrInput
	// The tags of PrefixList.
	Tags pulumi.MapInput
}

func (PrefixListState) ElementType() reflect.Type {
	return reflect.TypeOf((*prefixListState)(nil)).Elem()
}

type prefixListArgs struct {
	// The CIDR address block list of the prefix list.See the following `Block Entrys`.
	Entrys []PrefixListEntry `pulumi:"entrys"`
	// The IP version of the prefix list. Value:-**IPV4**:IPv4 version.-**IPV6**:IPv6 version.
	IpVersion *string `pulumi:"ipVersion"`
	// The maximum number of entries for CIDR address blocks in the prefix list.
	MaxEntries *int `pulumi:"maxEntries"`
	// The description of the prefix list.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	PrefixListDescription *string `pulumi:"prefixListDescription"`
	// The name of the prefix list. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
	PrefixListName *string `pulumi:"prefixListName"`
	// The ID of the resource group to which the PrefixList belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The tags of PrefixList.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a PrefixList resource.
type PrefixListArgs struct {
	// The CIDR address block list of the prefix list.See the following `Block Entrys`.
	Entrys PrefixListEntryArrayInput
	// The IP version of the prefix list. Value:-**IPV4**:IPv4 version.-**IPV6**:IPv6 version.
	IpVersion pulumi.StringPtrInput
	// The maximum number of entries for CIDR address blocks in the prefix list.
	MaxEntries pulumi.IntPtrInput
	// The description of the prefix list.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
	PrefixListDescription pulumi.StringPtrInput
	// The name of the prefix list. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
	PrefixListName pulumi.StringPtrInput
	// The ID of the resource group to which the PrefixList belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The tags of PrefixList.
	Tags pulumi.MapInput
}

func (PrefixListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prefixListArgs)(nil)).Elem()
}

type PrefixListInput interface {
	pulumi.Input

	ToPrefixListOutput() PrefixListOutput
	ToPrefixListOutputWithContext(ctx context.Context) PrefixListOutput
}

func (*PrefixList) ElementType() reflect.Type {
	return reflect.TypeOf((**PrefixList)(nil)).Elem()
}

func (i *PrefixList) ToPrefixListOutput() PrefixListOutput {
	return i.ToPrefixListOutputWithContext(context.Background())
}

func (i *PrefixList) ToPrefixListOutputWithContext(ctx context.Context) PrefixListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrefixListOutput)
}

// PrefixListArrayInput is an input type that accepts PrefixListArray and PrefixListArrayOutput values.
// You can construct a concrete instance of `PrefixListArrayInput` via:
//
//	PrefixListArray{ PrefixListArgs{...} }
type PrefixListArrayInput interface {
	pulumi.Input

	ToPrefixListArrayOutput() PrefixListArrayOutput
	ToPrefixListArrayOutputWithContext(context.Context) PrefixListArrayOutput
}

type PrefixListArray []PrefixListInput

func (PrefixListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrefixList)(nil)).Elem()
}

func (i PrefixListArray) ToPrefixListArrayOutput() PrefixListArrayOutput {
	return i.ToPrefixListArrayOutputWithContext(context.Background())
}

func (i PrefixListArray) ToPrefixListArrayOutputWithContext(ctx context.Context) PrefixListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrefixListArrayOutput)
}

// PrefixListMapInput is an input type that accepts PrefixListMap and PrefixListMapOutput values.
// You can construct a concrete instance of `PrefixListMapInput` via:
//
//	PrefixListMap{ "key": PrefixListArgs{...} }
type PrefixListMapInput interface {
	pulumi.Input

	ToPrefixListMapOutput() PrefixListMapOutput
	ToPrefixListMapOutputWithContext(context.Context) PrefixListMapOutput
}

type PrefixListMap map[string]PrefixListInput

func (PrefixListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrefixList)(nil)).Elem()
}

func (i PrefixListMap) ToPrefixListMapOutput() PrefixListMapOutput {
	return i.ToPrefixListMapOutputWithContext(context.Background())
}

func (i PrefixListMap) ToPrefixListMapOutputWithContext(ctx context.Context) PrefixListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrefixListMapOutput)
}

type PrefixListOutput struct{ *pulumi.OutputState }

func (PrefixListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrefixList)(nil)).Elem()
}

func (o PrefixListOutput) ToPrefixListOutput() PrefixListOutput {
	return o
}

func (o PrefixListOutput) ToPrefixListOutputWithContext(ctx context.Context) PrefixListOutput {
	return o
}

// The time when the prefix list was created.
func (o PrefixListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PrefixList) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The CIDR address block list of the prefix list.See the following `Block Entrys`.
func (o PrefixListOutput) Entrys() PrefixListEntryArrayOutput {
	return o.ApplyT(func(v *PrefixList) PrefixListEntryArrayOutput { return v.Entrys }).(PrefixListEntryArrayOutput)
}

// The IP version of the prefix list. Value:-**IPV4**:IPv4 version.-**IPV6**:IPv6 version.
func (o PrefixListOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PrefixList) pulumi.StringOutput { return v.IpVersion }).(pulumi.StringOutput)
}

// The maximum number of entries for CIDR address blocks in the prefix list.
func (o PrefixListOutput) MaxEntries() pulumi.IntOutput {
	return o.ApplyT(func(v *PrefixList) pulumi.IntOutput { return v.MaxEntries }).(pulumi.IntOutput)
}

// The association list information of the prefix list.
func (o PrefixListOutput) PrefixListAssociations() PrefixListPrefixListAssociationArrayOutput {
	return o.ApplyT(func(v *PrefixList) PrefixListPrefixListAssociationArrayOutput { return v.PrefixListAssociations }).(PrefixListPrefixListAssociationArrayOutput)
}

// The description of the prefix list.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
func (o PrefixListOutput) PrefixListDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrefixList) pulumi.StringPtrOutput { return v.PrefixListDescription }).(pulumi.StringPtrOutput)
}

// The ID of the query Prefix List.
func (o PrefixListOutput) PrefixListId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrefixList) pulumi.StringOutput { return v.PrefixListId }).(pulumi.StringOutput)
}

// The name of the prefix list. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
func (o PrefixListOutput) PrefixListName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrefixList) pulumi.StringPtrOutput { return v.PrefixListName }).(pulumi.StringPtrOutput)
}

// The ID of the resource group to which the PrefixList belongs.
func (o PrefixListOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrefixList) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The share type of the prefix list. Value:-**Shared**: indicates that the prefix list is a Shared prefix list.-Null: indicates that the prefix list is not a shared prefix list.
func (o PrefixListOutput) ShareType() pulumi.StringOutput {
	return o.ApplyT(func(v *PrefixList) pulumi.StringOutput { return v.ShareType }).(pulumi.StringOutput)
}

// Resource attribute fields that represent the status of the resource.
func (o PrefixListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *PrefixList) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags of PrefixList.
func (o PrefixListOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *PrefixList) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type PrefixListArrayOutput struct{ *pulumi.OutputState }

func (PrefixListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrefixList)(nil)).Elem()
}

func (o PrefixListArrayOutput) ToPrefixListArrayOutput() PrefixListArrayOutput {
	return o
}

func (o PrefixListArrayOutput) ToPrefixListArrayOutputWithContext(ctx context.Context) PrefixListArrayOutput {
	return o
}

func (o PrefixListArrayOutput) Index(i pulumi.IntInput) PrefixListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrefixList {
		return vs[0].([]*PrefixList)[vs[1].(int)]
	}).(PrefixListOutput)
}

type PrefixListMapOutput struct{ *pulumi.OutputState }

func (PrefixListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrefixList)(nil)).Elem()
}

func (o PrefixListMapOutput) ToPrefixListMapOutput() PrefixListMapOutput {
	return o
}

func (o PrefixListMapOutput) ToPrefixListMapOutputWithContext(ctx context.Context) PrefixListMapOutput {
	return o
}

func (o PrefixListMapOutput) MapIndex(k pulumi.StringInput) PrefixListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrefixList {
		return vs[0].(map[string]*PrefixList)[vs[1].(string)]
	}).(PrefixListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrefixListInput)(nil)).Elem(), &PrefixList{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrefixListArrayInput)(nil)).Elem(), PrefixListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrefixListMapInput)(nil)).Elem(), PrefixListMap{})
	pulumi.RegisterOutputType(PrefixListOutput{})
	pulumi.RegisterOutputType(PrefixListArrayOutput{})
	pulumi.RegisterOutputType(PrefixListMapOutput{})
}
