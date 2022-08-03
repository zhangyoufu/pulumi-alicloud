// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **NOTE:** Available in v1.162.0+.
//
// > **NOTE:** The maximum number of entries per acl is 300.
//
// For information about acl entry attachment and how to use it, see [Configure an acl entry](https://www.alibabacloud.com/help/en/doc-detail/70023.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "terraformslbaclconfig"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		ipVersion := "ipv4"
// 		if param := cfg.Get("ipVersion"); param != "" {
// 			ipVersion = param
// 		}
// 		defaultAcl, err := slb.NewAcl(ctx, "defaultAcl", &slb.AclArgs{
// 			IpVersion: pulumi.String(ipVersion),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = slb.NewAclEntryAttachment(ctx, "defaultAclEntryAttachment", &slb.AclEntryAttachmentArgs{
// 			AclId:   defaultAcl.ID(),
// 			Entry:   pulumi.String("168.10.10.0/24"),
// 			Comment: pulumi.String("second"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Acl entry attachment can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:slb/aclEntryAttachment:AclEntryAttachment example <acl_id>:<entry>
// ```
type AclEntryAttachment struct {
	pulumi.CustomResourceState

	// The ID of the Acl.
	AclId pulumi.StringOutput `pulumi:"aclId"`
	// The comment of the entry.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The CIDR blocks.
	Entry pulumi.StringOutput `pulumi:"entry"`
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
	var resource AclEntryAttachment
	err := ctx.RegisterResource("alicloud:slb/aclEntryAttachment:AclEntryAttachment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:slb/aclEntryAttachment:AclEntryAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AclEntryAttachment resources.
type aclEntryAttachmentState struct {
	// The ID of the Acl.
	AclId *string `pulumi:"aclId"`
	// The comment of the entry.
	Comment *string `pulumi:"comment"`
	// The CIDR blocks.
	Entry *string `pulumi:"entry"`
}

type AclEntryAttachmentState struct {
	// The ID of the Acl.
	AclId pulumi.StringPtrInput
	// The comment of the entry.
	Comment pulumi.StringPtrInput
	// The CIDR blocks.
	Entry pulumi.StringPtrInput
}

func (AclEntryAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclEntryAttachmentState)(nil)).Elem()
}

type aclEntryAttachmentArgs struct {
	// The ID of the Acl.
	AclId string `pulumi:"aclId"`
	// The comment of the entry.
	Comment *string `pulumi:"comment"`
	// The CIDR blocks.
	Entry string `pulumi:"entry"`
}

// The set of arguments for constructing a AclEntryAttachment resource.
type AclEntryAttachmentArgs struct {
	// The ID of the Acl.
	AclId pulumi.StringInput
	// The comment of the entry.
	Comment pulumi.StringPtrInput
	// The CIDR blocks.
	Entry pulumi.StringInput
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
//          AclEntryAttachmentArray{ AclEntryAttachmentArgs{...} }
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
//          AclEntryAttachmentMap{ "key": AclEntryAttachmentArgs{...} }
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

// The ID of the Acl.
func (o AclEntryAttachmentOutput) AclId() pulumi.StringOutput {
	return o.ApplyT(func(v *AclEntryAttachment) pulumi.StringOutput { return v.AclId }).(pulumi.StringOutput)
}

// The comment of the entry.
func (o AclEntryAttachmentOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AclEntryAttachment) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// The CIDR blocks.
func (o AclEntryAttachmentOutput) Entry() pulumi.StringOutput {
	return o.ApplyT(func(v *AclEntryAttachment) pulumi.StringOutput { return v.Entry }).(pulumi.StringOutput)
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
