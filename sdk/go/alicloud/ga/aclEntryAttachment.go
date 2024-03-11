// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Acl entry attachment resource.
//
// For information about Global Accelerator (GA) Acl entry attachment and how to use it, see [What is Acl entry attachment](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-addentriestoacl).
//
// > **NOTE:** Available since v1.190.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultAcl, err := ga.NewAcl(ctx, "defaultAcl", &ga.AclArgs{
//				AclName:          pulumi.String("tf-example-value"),
//				AddressIpVersion: pulumi.String("IPv4"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ga.NewAclEntryAttachment(ctx, "defaultAclEntryAttachment", &ga.AclEntryAttachmentArgs{
//				AclId:            defaultAcl.ID(),
//				Entry:            pulumi.String("192.168.1.1/32"),
//				EntryDescription: pulumi.String("tf-example-value"),
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
// Global Accelerator (GA) Acl entry attachment can be imported using the id.Format to `<acl_id>:<entry>`, e.g.
//
// ```sh
// $ pulumi import alicloud:ga/aclEntryAttachment:AclEntryAttachment example your_acl_id:your_entry
// ```
type AclEntryAttachment struct {
	pulumi.CustomResourceState

	// The ID of the global acceleration instance.
	AclId pulumi.StringOutput `pulumi:"aclId"`
	// The IP address(192.168.XX.XX) or CIDR(10.0.XX.XX/24) block that you want to add to the network ACL.
	Entry pulumi.StringOutput `pulumi:"entry"`
	// The description of the entry. The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	EntryDescription pulumi.StringPtrOutput `pulumi:"entryDescription"`
	// The status of the network ACL.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAclEntryAttachment registers a new resource with the given unique name, arguments, and options.
func NewAclEntryAttachment(ctx *pulumi.Context,
	name string, args *AclEntryAttachmentArgs, opts ...pulumi.ResourceOption) (*AclEntryAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclId == nil {
		return nil, errors.New("invalid value for required argument 'AclId'")
	}
	if args.Entry == nil {
		return nil, errors.New("invalid value for required argument 'Entry'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AclEntryAttachment
	err := ctx.RegisterResource("alicloud:ga/aclEntryAttachment:AclEntryAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAclEntryAttachment gets an existing AclEntryAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAclEntryAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclEntryAttachmentState, opts ...pulumi.ResourceOption) (*AclEntryAttachment, error) {
	var resource AclEntryAttachment
	err := ctx.ReadResource("alicloud:ga/aclEntryAttachment:AclEntryAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AclEntryAttachment resources.
type aclEntryAttachmentState struct {
	// The ID of the global acceleration instance.
	AclId *string `pulumi:"aclId"`
	// The IP address(192.168.XX.XX) or CIDR(10.0.XX.XX/24) block that you want to add to the network ACL.
	Entry *string `pulumi:"entry"`
	// The description of the entry. The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	EntryDescription *string `pulumi:"entryDescription"`
	// The status of the network ACL.
	Status *string `pulumi:"status"`
}

type AclEntryAttachmentState struct {
	// The ID of the global acceleration instance.
	AclId pulumi.StringPtrInput
	// The IP address(192.168.XX.XX) or CIDR(10.0.XX.XX/24) block that you want to add to the network ACL.
	Entry pulumi.StringPtrInput
	// The description of the entry. The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	EntryDescription pulumi.StringPtrInput
	// The status of the network ACL.
	Status pulumi.StringPtrInput
}

func (AclEntryAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclEntryAttachmentState)(nil)).Elem()
}

type aclEntryAttachmentArgs struct {
	// The ID of the global acceleration instance.
	AclId string `pulumi:"aclId"`
	// The IP address(192.168.XX.XX) or CIDR(10.0.XX.XX/24) block that you want to add to the network ACL.
	Entry string `pulumi:"entry"`
	// The description of the entry. The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	EntryDescription *string `pulumi:"entryDescription"`
}

// The set of arguments for constructing a AclEntryAttachment resource.
type AclEntryAttachmentArgs struct {
	// The ID of the global acceleration instance.
	AclId pulumi.StringInput
	// The IP address(192.168.XX.XX) or CIDR(10.0.XX.XX/24) block that you want to add to the network ACL.
	Entry pulumi.StringInput
	// The description of the entry. The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	EntryDescription pulumi.StringPtrInput
}

func (AclEntryAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclEntryAttachmentArgs)(nil)).Elem()
}

type AclEntryAttachmentInput interface {
	pulumi.Input

	ToAclEntryAttachmentOutput() AclEntryAttachmentOutput
	ToAclEntryAttachmentOutputWithContext(ctx context.Context) AclEntryAttachmentOutput
}

func (*AclEntryAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**AclEntryAttachment)(nil)).Elem()
}

func (i *AclEntryAttachment) ToAclEntryAttachmentOutput() AclEntryAttachmentOutput {
	return i.ToAclEntryAttachmentOutputWithContext(context.Background())
}

func (i *AclEntryAttachment) ToAclEntryAttachmentOutputWithContext(ctx context.Context) AclEntryAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclEntryAttachmentOutput)
}

// AclEntryAttachmentArrayInput is an input type that accepts AclEntryAttachmentArray and AclEntryAttachmentArrayOutput values.
// You can construct a concrete instance of `AclEntryAttachmentArrayInput` via:
//
//	AclEntryAttachmentArray{ AclEntryAttachmentArgs{...} }
type AclEntryAttachmentArrayInput interface {
	pulumi.Input

	ToAclEntryAttachmentArrayOutput() AclEntryAttachmentArrayOutput
	ToAclEntryAttachmentArrayOutputWithContext(context.Context) AclEntryAttachmentArrayOutput
}

type AclEntryAttachmentArray []AclEntryAttachmentInput

func (AclEntryAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclEntryAttachment)(nil)).Elem()
}

func (i AclEntryAttachmentArray) ToAclEntryAttachmentArrayOutput() AclEntryAttachmentArrayOutput {
	return i.ToAclEntryAttachmentArrayOutputWithContext(context.Background())
}

func (i AclEntryAttachmentArray) ToAclEntryAttachmentArrayOutputWithContext(ctx context.Context) AclEntryAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclEntryAttachmentArrayOutput)
}

// AclEntryAttachmentMapInput is an input type that accepts AclEntryAttachmentMap and AclEntryAttachmentMapOutput values.
// You can construct a concrete instance of `AclEntryAttachmentMapInput` via:
//
//	AclEntryAttachmentMap{ "key": AclEntryAttachmentArgs{...} }
type AclEntryAttachmentMapInput interface {
	pulumi.Input

	ToAclEntryAttachmentMapOutput() AclEntryAttachmentMapOutput
	ToAclEntryAttachmentMapOutputWithContext(context.Context) AclEntryAttachmentMapOutput
}

type AclEntryAttachmentMap map[string]AclEntryAttachmentInput

func (AclEntryAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclEntryAttachment)(nil)).Elem()
}

func (i AclEntryAttachmentMap) ToAclEntryAttachmentMapOutput() AclEntryAttachmentMapOutput {
	return i.ToAclEntryAttachmentMapOutputWithContext(context.Background())
}

func (i AclEntryAttachmentMap) ToAclEntryAttachmentMapOutputWithContext(ctx context.Context) AclEntryAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclEntryAttachmentMapOutput)
}

type AclEntryAttachmentOutput struct{ *pulumi.OutputState }

func (AclEntryAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AclEntryAttachment)(nil)).Elem()
}

func (o AclEntryAttachmentOutput) ToAclEntryAttachmentOutput() AclEntryAttachmentOutput {
	return o
}

func (o AclEntryAttachmentOutput) ToAclEntryAttachmentOutputWithContext(ctx context.Context) AclEntryAttachmentOutput {
	return o
}

// The ID of the global acceleration instance.
func (o AclEntryAttachmentOutput) AclId() pulumi.StringOutput {
	return o.ApplyT(func(v *AclEntryAttachment) pulumi.StringOutput { return v.AclId }).(pulumi.StringOutput)
}

// The IP address(192.168.XX.XX) or CIDR(10.0.XX.XX/24) block that you want to add to the network ACL.
func (o AclEntryAttachmentOutput) Entry() pulumi.StringOutput {
	return o.ApplyT(func(v *AclEntryAttachment) pulumi.StringOutput { return v.Entry }).(pulumi.StringOutput)
}

// The description of the entry. The description must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
func (o AclEntryAttachmentOutput) EntryDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AclEntryAttachment) pulumi.StringPtrOutput { return v.EntryDescription }).(pulumi.StringPtrOutput)
}

// The status of the network ACL.
func (o AclEntryAttachmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AclEntryAttachment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AclEntryAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AclEntryAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclEntryAttachment)(nil)).Elem()
}

func (o AclEntryAttachmentArrayOutput) ToAclEntryAttachmentArrayOutput() AclEntryAttachmentArrayOutput {
	return o
}

func (o AclEntryAttachmentArrayOutput) ToAclEntryAttachmentArrayOutputWithContext(ctx context.Context) AclEntryAttachmentArrayOutput {
	return o
}

func (o AclEntryAttachmentArrayOutput) Index(i pulumi.IntInput) AclEntryAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AclEntryAttachment {
		return vs[0].([]*AclEntryAttachment)[vs[1].(int)]
	}).(AclEntryAttachmentOutput)
}

type AclEntryAttachmentMapOutput struct{ *pulumi.OutputState }

func (AclEntryAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclEntryAttachment)(nil)).Elem()
}

func (o AclEntryAttachmentMapOutput) ToAclEntryAttachmentMapOutput() AclEntryAttachmentMapOutput {
	return o
}

func (o AclEntryAttachmentMapOutput) ToAclEntryAttachmentMapOutputWithContext(ctx context.Context) AclEntryAttachmentMapOutput {
	return o
}

func (o AclEntryAttachmentMapOutput) MapIndex(k pulumi.StringInput) AclEntryAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AclEntryAttachment {
		return vs[0].(map[string]*AclEntryAttachment)[vs[1].(string)]
	}).(AclEntryAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclEntryAttachmentInput)(nil)).Elem(), &AclEntryAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclEntryAttachmentArrayInput)(nil)).Elem(), AclEntryAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclEntryAttachmentMapInput)(nil)).Elem(), AclEntryAttachmentMap{})
	pulumi.RegisterOutputType(AclEntryAttachmentOutput{})
	pulumi.RegisterOutputType(AclEntryAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AclEntryAttachmentMapOutput{})
}
