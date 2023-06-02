// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Acl resource.
//
// For information about Global Accelerator (GA) Acl and how to use it, see [What is Acl](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-doc-ga-2019-11-20-api-doc-createacl).
//
// > **NOTE:** Available in v1.150.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ga.NewAcl(ctx, "default", &ga.AclArgs{
//				AclEntries: ga.AclAclEntryArray{
//					&ga.AclAclEntryArgs{
//						Entry:            pulumi.String("192.168.1.0/24"),
//						EntryDescription: pulumi.String("tf-test1"),
//					},
//				},
//				AclName:          pulumi.String("tf-testAccAcl"),
//				AddressIpVersion: pulumi.String("IPv4"),
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
// Global Accelerator (GA) Acl can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ga/acl:Acl example <id>
//
// ```
type Acl struct {
	pulumi.CustomResourceState

	// The entries of the Acl. See the following `Block aclEntries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
	//
	// Deprecated: Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.
	AclEntries AclAclEntryArrayOutput `pulumi:"aclEntries"`
	// The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
	AclName pulumi.StringPtrOutput `pulumi:"aclName"`
	// The IP version. Valid values: `IPv4` and `IPv6`.
	AddressIpVersion pulumi.StringOutput `pulumi:"addressIpVersion"`
	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAcl registers a new resource with the given unique name, arguments, and options.
func NewAcl(ctx *pulumi.Context,
	name string, args *AclArgs, opts ...pulumi.ResourceOption) (*Acl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressIpVersion == nil {
		return nil, errors.New("invalid value for required argument 'AddressIpVersion'")
	}
	var resource Acl
	err := ctx.RegisterResource("alicloud:ga/acl:Acl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAcl gets an existing Acl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclState, opts ...pulumi.ResourceOption) (*Acl, error) {
	var resource Acl
	err := ctx.ReadResource("alicloud:ga/acl:Acl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Acl resources.
type aclState struct {
	// The entries of the Acl. See the following `Block aclEntries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
	//
	// Deprecated: Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.
	AclEntries []AclAclEntry `pulumi:"aclEntries"`
	// The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
	AclName *string `pulumi:"aclName"`
	// The IP version. Valid values: `IPv4` and `IPv6`.
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The status of the resource.
	Status *string `pulumi:"status"`
}

type AclState struct {
	// The entries of the Acl. See the following `Block aclEntries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
	//
	// Deprecated: Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.
	AclEntries AclAclEntryArrayInput
	// The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
	AclName pulumi.StringPtrInput
	// The IP version. Valid values: `IPv4` and `IPv6`.
	AddressIpVersion pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
}

func (AclState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclState)(nil)).Elem()
}

type aclArgs struct {
	// The entries of the Acl. See the following `Block aclEntries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
	//
	// Deprecated: Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.
	AclEntries []AclAclEntry `pulumi:"aclEntries"`
	// The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
	AclName *string `pulumi:"aclName"`
	// The IP version. Valid values: `IPv4` and `IPv6`.
	AddressIpVersion string `pulumi:"addressIpVersion"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
}

// The set of arguments for constructing a Acl resource.
type AclArgs struct {
	// The entries of the Acl. See the following `Block aclEntries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
	//
	// Deprecated: Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.
	AclEntries AclAclEntryArrayInput
	// The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
	AclName pulumi.StringPtrInput
	// The IP version. Valid values: `IPv4` and `IPv6`.
	AddressIpVersion pulumi.StringInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
}

func (AclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclArgs)(nil)).Elem()
}

type AclInput interface {
	pulumi.Input

	ToAclOutput() AclOutput
	ToAclOutputWithContext(ctx context.Context) AclOutput
}

func (*Acl) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (i *Acl) ToAclOutput() AclOutput {
	return i.ToAclOutputWithContext(context.Background())
}

func (i *Acl) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclOutput)
}

// AclArrayInput is an input type that accepts AclArray and AclArrayOutput values.
// You can construct a concrete instance of `AclArrayInput` via:
//
//	AclArray{ AclArgs{...} }
type AclArrayInput interface {
	pulumi.Input

	ToAclArrayOutput() AclArrayOutput
	ToAclArrayOutputWithContext(context.Context) AclArrayOutput
}

type AclArray []AclInput

func (AclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (i AclArray) ToAclArrayOutput() AclArrayOutput {
	return i.ToAclArrayOutputWithContext(context.Background())
}

func (i AclArray) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclArrayOutput)
}

// AclMapInput is an input type that accepts AclMap and AclMapOutput values.
// You can construct a concrete instance of `AclMapInput` via:
//
//	AclMap{ "key": AclArgs{...} }
type AclMapInput interface {
	pulumi.Input

	ToAclMapOutput() AclMapOutput
	ToAclMapOutputWithContext(context.Context) AclMapOutput
}

type AclMap map[string]AclInput

func (AclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (i AclMap) ToAclMapOutput() AclMapOutput {
	return i.ToAclMapOutputWithContext(context.Background())
}

func (i AclMap) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclMapOutput)
}

type AclOutput struct{ *pulumi.OutputState }

func (AclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (o AclOutput) ToAclOutput() AclOutput {
	return o
}

func (o AclOutput) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return o
}

// The entries of the Acl. See the following `Block aclEntries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
//
// Deprecated: Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.
func (o AclOutput) AclEntries() AclAclEntryArrayOutput {
	return o.ApplyT(func(v *Acl) AclAclEntryArrayOutput { return v.AclEntries }).(AclAclEntryArrayOutput)
}

// The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
func (o AclOutput) AclName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringPtrOutput { return v.AclName }).(pulumi.StringPtrOutput)
}

// The IP version. Valid values: `IPv4` and `IPv6`.
func (o AclOutput) AddressIpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.AddressIpVersion }).(pulumi.StringOutput)
}

// The dry run.
func (o AclOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The status of the resource.
func (o AclOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AclArrayOutput struct{ *pulumi.OutputState }

func (AclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (o AclArrayOutput) ToAclArrayOutput() AclArrayOutput {
	return o
}

func (o AclArrayOutput) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return o
}

func (o AclArrayOutput) Index(i pulumi.IntInput) AclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Acl {
		return vs[0].([]*Acl)[vs[1].(int)]
	}).(AclOutput)
}

type AclMapOutput struct{ *pulumi.OutputState }

func (AclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (o AclMapOutput) ToAclMapOutput() AclMapOutput {
	return o
}

func (o AclMapOutput) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return o
}

func (o AclMapOutput) MapIndex(k pulumi.StringInput) AclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Acl {
		return vs[0].(map[string]*Acl)[vs[1].(string)]
	}).(AclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclInput)(nil)).Elem(), &Acl{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclArrayInput)(nil)).Elem(), AclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclMapInput)(nil)).Elem(), AclMap{})
	pulumi.RegisterOutputType(AclOutput{})
	pulumi.RegisterOutputType(AclArrayOutput{})
	pulumi.RegisterOutputType(AclMapOutput{})
}
