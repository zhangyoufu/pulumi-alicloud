// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECS Prefix List resource.
//
// For information about ECS Prefix List and how to use it, see [What is Prefix List.](https://www.alibabacloud.com/help/en/doc-detail/207969.html).
//
// > **NOTE:** Available in v1.152.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewEcsPrefixList(ctx, "default", &ecs.EcsPrefixListArgs{
//				AddressFamily: pulumi.String("IPv4"),
//				Description:   pulumi.String("description"),
//				Entries: ecs.EcsPrefixListEntryArray{
//					&ecs.EcsPrefixListEntryArgs{
//						Cidr:        pulumi.String("192.168.0.0/24"),
//						Description: pulumi.String("description"),
//					},
//				},
//				MaxEntries:     pulumi.Int(2),
//				PrefixListName: pulumi.String("tftest"),
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
// ECS Prefix List can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ecs/ecsPrefixList:EcsPrefixList example <id>
// ```
type EcsPrefixList struct {
	pulumi.CustomResourceState

	// The IP address family. Valid values: `IPv4`,`IPv6`.
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Entry. The details see Block `entry`.
	Entries EcsPrefixListEntryArrayOutput `pulumi:"entries"`
	// The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
	MaxEntries pulumi.IntOutput `pulumi:"maxEntries"`
	// The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	PrefixListName pulumi.StringOutput `pulumi:"prefixListName"`
}

// NewEcsPrefixList registers a new resource with the given unique name, arguments, and options.
func NewEcsPrefixList(ctx *pulumi.Context,
	name string, args *EcsPrefixListArgs, opts ...pulumi.ResourceOption) (*EcsPrefixList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'AddressFamily'")
	}
	if args.Entries == nil {
		return nil, errors.New("invalid value for required argument 'Entries'")
	}
	if args.MaxEntries == nil {
		return nil, errors.New("invalid value for required argument 'MaxEntries'")
	}
	if args.PrefixListName == nil {
		return nil, errors.New("invalid value for required argument 'PrefixListName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EcsPrefixList
	err := ctx.RegisterResource("alicloud:ecs/ecsPrefixList:EcsPrefixList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcsPrefixList gets an existing EcsPrefixList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcsPrefixList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcsPrefixListState, opts ...pulumi.ResourceOption) (*EcsPrefixList, error) {
	var resource EcsPrefixList
	err := ctx.ReadResource("alicloud:ecs/ecsPrefixList:EcsPrefixList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcsPrefixList resources.
type ecsPrefixListState struct {
	// The IP address family. Valid values: `IPv4`,`IPv6`.
	AddressFamily *string `pulumi:"addressFamily"`
	// The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// The Entry. The details see Block `entry`.
	Entries []EcsPrefixListEntry `pulumi:"entries"`
	// The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
	MaxEntries *int `pulumi:"maxEntries"`
	// The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	PrefixListName *string `pulumi:"prefixListName"`
}

type EcsPrefixListState struct {
	// The IP address family. Valid values: `IPv4`,`IPv6`.
	AddressFamily pulumi.StringPtrInput
	// The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// The Entry. The details see Block `entry`.
	Entries EcsPrefixListEntryArrayInput
	// The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
	MaxEntries pulumi.IntPtrInput
	// The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	PrefixListName pulumi.StringPtrInput
}

func (EcsPrefixListState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsPrefixListState)(nil)).Elem()
}

type ecsPrefixListArgs struct {
	// The IP address family. Valid values: `IPv4`,`IPv6`.
	AddressFamily string `pulumi:"addressFamily"`
	// The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// The Entry. The details see Block `entry`.
	Entries []EcsPrefixListEntry `pulumi:"entries"`
	// The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
	MaxEntries int `pulumi:"maxEntries"`
	// The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	PrefixListName string `pulumi:"prefixListName"`
}

// The set of arguments for constructing a EcsPrefixList resource.
type EcsPrefixListArgs struct {
	// The IP address family. Valid values: `IPv4`,`IPv6`.
	AddressFamily pulumi.StringInput
	// The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// The Entry. The details see Block `entry`.
	Entries EcsPrefixListEntryArrayInput
	// The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
	MaxEntries pulumi.IntInput
	// The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	PrefixListName pulumi.StringInput
}

func (EcsPrefixListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsPrefixListArgs)(nil)).Elem()
}

type EcsPrefixListInput interface {
	pulumi.Input

	ToEcsPrefixListOutput() EcsPrefixListOutput
	ToEcsPrefixListOutputWithContext(ctx context.Context) EcsPrefixListOutput
}

func (*EcsPrefixList) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsPrefixList)(nil)).Elem()
}

func (i *EcsPrefixList) ToEcsPrefixListOutput() EcsPrefixListOutput {
	return i.ToEcsPrefixListOutputWithContext(context.Background())
}

func (i *EcsPrefixList) ToEcsPrefixListOutputWithContext(ctx context.Context) EcsPrefixListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsPrefixListOutput)
}

// EcsPrefixListArrayInput is an input type that accepts EcsPrefixListArray and EcsPrefixListArrayOutput values.
// You can construct a concrete instance of `EcsPrefixListArrayInput` via:
//
//	EcsPrefixListArray{ EcsPrefixListArgs{...} }
type EcsPrefixListArrayInput interface {
	pulumi.Input

	ToEcsPrefixListArrayOutput() EcsPrefixListArrayOutput
	ToEcsPrefixListArrayOutputWithContext(context.Context) EcsPrefixListArrayOutput
}

type EcsPrefixListArray []EcsPrefixListInput

func (EcsPrefixListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsPrefixList)(nil)).Elem()
}

func (i EcsPrefixListArray) ToEcsPrefixListArrayOutput() EcsPrefixListArrayOutput {
	return i.ToEcsPrefixListArrayOutputWithContext(context.Background())
}

func (i EcsPrefixListArray) ToEcsPrefixListArrayOutputWithContext(ctx context.Context) EcsPrefixListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsPrefixListArrayOutput)
}

// EcsPrefixListMapInput is an input type that accepts EcsPrefixListMap and EcsPrefixListMapOutput values.
// You can construct a concrete instance of `EcsPrefixListMapInput` via:
//
//	EcsPrefixListMap{ "key": EcsPrefixListArgs{...} }
type EcsPrefixListMapInput interface {
	pulumi.Input

	ToEcsPrefixListMapOutput() EcsPrefixListMapOutput
	ToEcsPrefixListMapOutputWithContext(context.Context) EcsPrefixListMapOutput
}

type EcsPrefixListMap map[string]EcsPrefixListInput

func (EcsPrefixListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsPrefixList)(nil)).Elem()
}

func (i EcsPrefixListMap) ToEcsPrefixListMapOutput() EcsPrefixListMapOutput {
	return i.ToEcsPrefixListMapOutputWithContext(context.Background())
}

func (i EcsPrefixListMap) ToEcsPrefixListMapOutputWithContext(ctx context.Context) EcsPrefixListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsPrefixListMapOutput)
}

type EcsPrefixListOutput struct{ *pulumi.OutputState }

func (EcsPrefixListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsPrefixList)(nil)).Elem()
}

func (o EcsPrefixListOutput) ToEcsPrefixListOutput() EcsPrefixListOutput {
	return o
}

func (o EcsPrefixListOutput) ToEcsPrefixListOutputWithContext(ctx context.Context) EcsPrefixListOutput {
	return o
}

// The IP address family. Valid values: `IPv4`,`IPv6`.
func (o EcsPrefixListOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsPrefixList) pulumi.StringOutput { return v.AddressFamily }).(pulumi.StringOutput)
}

// The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
func (o EcsPrefixListOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsPrefixList) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Entry. The details see Block `entry`.
func (o EcsPrefixListOutput) Entries() EcsPrefixListEntryArrayOutput {
	return o.ApplyT(func(v *EcsPrefixList) EcsPrefixListEntryArrayOutput { return v.Entries }).(EcsPrefixListEntryArrayOutput)
}

// The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
func (o EcsPrefixListOutput) MaxEntries() pulumi.IntOutput {
	return o.ApplyT(func(v *EcsPrefixList) pulumi.IntOutput { return v.MaxEntries }).(pulumi.IntOutput)
}

// The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
func (o EcsPrefixListOutput) PrefixListName() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsPrefixList) pulumi.StringOutput { return v.PrefixListName }).(pulumi.StringOutput)
}

type EcsPrefixListArrayOutput struct{ *pulumi.OutputState }

func (EcsPrefixListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsPrefixList)(nil)).Elem()
}

func (o EcsPrefixListArrayOutput) ToEcsPrefixListArrayOutput() EcsPrefixListArrayOutput {
	return o
}

func (o EcsPrefixListArrayOutput) ToEcsPrefixListArrayOutputWithContext(ctx context.Context) EcsPrefixListArrayOutput {
	return o
}

func (o EcsPrefixListArrayOutput) Index(i pulumi.IntInput) EcsPrefixListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EcsPrefixList {
		return vs[0].([]*EcsPrefixList)[vs[1].(int)]
	}).(EcsPrefixListOutput)
}

type EcsPrefixListMapOutput struct{ *pulumi.OutputState }

func (EcsPrefixListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsPrefixList)(nil)).Elem()
}

func (o EcsPrefixListMapOutput) ToEcsPrefixListMapOutput() EcsPrefixListMapOutput {
	return o
}

func (o EcsPrefixListMapOutput) ToEcsPrefixListMapOutputWithContext(ctx context.Context) EcsPrefixListMapOutput {
	return o
}

func (o EcsPrefixListMapOutput) MapIndex(k pulumi.StringInput) EcsPrefixListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EcsPrefixList {
		return vs[0].(map[string]*EcsPrefixList)[vs[1].(string)]
	}).(EcsPrefixListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EcsPrefixListInput)(nil)).Elem(), &EcsPrefixList{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsPrefixListArrayInput)(nil)).Elem(), EcsPrefixListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsPrefixListMapInput)(nil)).Elem(), EcsPrefixListMap{})
	pulumi.RegisterOutputType(EcsPrefixListOutput{})
	pulumi.RegisterOutputType(EcsPrefixListArrayOutput{})
	pulumi.RegisterOutputType(EcsPrefixListMapOutput{})
}
