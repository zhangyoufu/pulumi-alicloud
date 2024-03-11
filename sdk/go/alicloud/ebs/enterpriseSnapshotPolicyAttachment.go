// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a EBS Enterprise Snapshot Policy Attachment resource. Enterprise-level snapshot policy cloud disk binding relationship.
//
// For information about EBS Enterprise Snapshot Policy Attachment and how to use it, see [What is Enterprise Snapshot Policy Attachment](https://www.alibabacloud.com/help/en/).
//
// > **NOTE:** Available since v1.215.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ebs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultJkW46o, err := ecs.NewEcsDisk(ctx, "defaultJkW46o", &ecs.EcsDiskArgs{
//				Category:         pulumi.String("cloud_essd"),
//				Description:      pulumi.String("esp-attachment-test"),
//				ZoneId:           pulumi.String("cn-hangzhou-i"),
//				PerformanceLevel: pulumi.String("PL1"),
//				Size:             pulumi.Int(20),
//				DiskName:         pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultPE3jjR, err := ebs.NewEnterpriseSnapshotPolicy(ctx, "defaultPE3jjR", &ebs.EnterpriseSnapshotPolicyArgs{
//				Status: pulumi.String("DISABLED"),
//				Desc:   pulumi.String("DESC"),
//				Schedule: &ebs.EnterpriseSnapshotPolicyScheduleArgs{
//					CronExpression: pulumi.String("0 0 0 1 * ?"),
//				},
//				EnterpriseSnapshotPolicyName: pulumi.String(name),
//				TargetType:                   pulumi.String("DISK"),
//				RetainRule: &ebs.EnterpriseSnapshotPolicyRetainRuleArgs{
//					TimeInterval: pulumi.Int(120),
//					TimeUnit:     pulumi.String("DAYS"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ebs.NewEnterpriseSnapshotPolicyAttachment(ctx, "default", &ebs.EnterpriseSnapshotPolicyAttachmentArgs{
//				PolicyId: defaultPE3jjR.ID(),
//				DiskId:   defaultJkW46o.ID(),
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
// EBS Enterprise Snapshot Policy Attachment can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ebs/enterpriseSnapshotPolicyAttachment:EnterpriseSnapshotPolicyAttachment example <policy_id>:<disk_id>
// ```
type EnterpriseSnapshotPolicyAttachment struct {
	pulumi.CustomResourceState

	// Cloud Disk ID.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// the enterprise snapshot policy id.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
}

// NewEnterpriseSnapshotPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseSnapshotPolicyAttachment(ctx *pulumi.Context,
	name string, args *EnterpriseSnapshotPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*EnterpriseSnapshotPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnterpriseSnapshotPolicyAttachment
	err := ctx.RegisterResource("alicloud:ebs/enterpriseSnapshotPolicyAttachment:EnterpriseSnapshotPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseSnapshotPolicyAttachment gets an existing EnterpriseSnapshotPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseSnapshotPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseSnapshotPolicyAttachmentState, opts ...pulumi.ResourceOption) (*EnterpriseSnapshotPolicyAttachment, error) {
	var resource EnterpriseSnapshotPolicyAttachment
	err := ctx.ReadResource("alicloud:ebs/enterpriseSnapshotPolicyAttachment:EnterpriseSnapshotPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseSnapshotPolicyAttachment resources.
type enterpriseSnapshotPolicyAttachmentState struct {
	// Cloud Disk ID.
	DiskId *string `pulumi:"diskId"`
	// the enterprise snapshot policy id.
	PolicyId *string `pulumi:"policyId"`
}

type EnterpriseSnapshotPolicyAttachmentState struct {
	// Cloud Disk ID.
	DiskId pulumi.StringPtrInput
	// the enterprise snapshot policy id.
	PolicyId pulumi.StringPtrInput
}

func (EnterpriseSnapshotPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseSnapshotPolicyAttachmentState)(nil)).Elem()
}

type enterpriseSnapshotPolicyAttachmentArgs struct {
	// Cloud Disk ID.
	DiskId *string `pulumi:"diskId"`
	// the enterprise snapshot policy id.
	PolicyId string `pulumi:"policyId"`
}

// The set of arguments for constructing a EnterpriseSnapshotPolicyAttachment resource.
type EnterpriseSnapshotPolicyAttachmentArgs struct {
	// Cloud Disk ID.
	DiskId pulumi.StringPtrInput
	// the enterprise snapshot policy id.
	PolicyId pulumi.StringInput
}

func (EnterpriseSnapshotPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseSnapshotPolicyAttachmentArgs)(nil)).Elem()
}

type EnterpriseSnapshotPolicyAttachmentInput interface {
	pulumi.Input

	ToEnterpriseSnapshotPolicyAttachmentOutput() EnterpriseSnapshotPolicyAttachmentOutput
	ToEnterpriseSnapshotPolicyAttachmentOutputWithContext(ctx context.Context) EnterpriseSnapshotPolicyAttachmentOutput
}

func (*EnterpriseSnapshotPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseSnapshotPolicyAttachment)(nil)).Elem()
}

func (i *EnterpriseSnapshotPolicyAttachment) ToEnterpriseSnapshotPolicyAttachmentOutput() EnterpriseSnapshotPolicyAttachmentOutput {
	return i.ToEnterpriseSnapshotPolicyAttachmentOutputWithContext(context.Background())
}

func (i *EnterpriseSnapshotPolicyAttachment) ToEnterpriseSnapshotPolicyAttachmentOutputWithContext(ctx context.Context) EnterpriseSnapshotPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSnapshotPolicyAttachmentOutput)
}

// EnterpriseSnapshotPolicyAttachmentArrayInput is an input type that accepts EnterpriseSnapshotPolicyAttachmentArray and EnterpriseSnapshotPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `EnterpriseSnapshotPolicyAttachmentArrayInput` via:
//
//	EnterpriseSnapshotPolicyAttachmentArray{ EnterpriseSnapshotPolicyAttachmentArgs{...} }
type EnterpriseSnapshotPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToEnterpriseSnapshotPolicyAttachmentArrayOutput() EnterpriseSnapshotPolicyAttachmentArrayOutput
	ToEnterpriseSnapshotPolicyAttachmentArrayOutputWithContext(context.Context) EnterpriseSnapshotPolicyAttachmentArrayOutput
}

type EnterpriseSnapshotPolicyAttachmentArray []EnterpriseSnapshotPolicyAttachmentInput

func (EnterpriseSnapshotPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseSnapshotPolicyAttachment)(nil)).Elem()
}

func (i EnterpriseSnapshotPolicyAttachmentArray) ToEnterpriseSnapshotPolicyAttachmentArrayOutput() EnterpriseSnapshotPolicyAttachmentArrayOutput {
	return i.ToEnterpriseSnapshotPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i EnterpriseSnapshotPolicyAttachmentArray) ToEnterpriseSnapshotPolicyAttachmentArrayOutputWithContext(ctx context.Context) EnterpriseSnapshotPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSnapshotPolicyAttachmentArrayOutput)
}

// EnterpriseSnapshotPolicyAttachmentMapInput is an input type that accepts EnterpriseSnapshotPolicyAttachmentMap and EnterpriseSnapshotPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `EnterpriseSnapshotPolicyAttachmentMapInput` via:
//
//	EnterpriseSnapshotPolicyAttachmentMap{ "key": EnterpriseSnapshotPolicyAttachmentArgs{...} }
type EnterpriseSnapshotPolicyAttachmentMapInput interface {
	pulumi.Input

	ToEnterpriseSnapshotPolicyAttachmentMapOutput() EnterpriseSnapshotPolicyAttachmentMapOutput
	ToEnterpriseSnapshotPolicyAttachmentMapOutputWithContext(context.Context) EnterpriseSnapshotPolicyAttachmentMapOutput
}

type EnterpriseSnapshotPolicyAttachmentMap map[string]EnterpriseSnapshotPolicyAttachmentInput

func (EnterpriseSnapshotPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseSnapshotPolicyAttachment)(nil)).Elem()
}

func (i EnterpriseSnapshotPolicyAttachmentMap) ToEnterpriseSnapshotPolicyAttachmentMapOutput() EnterpriseSnapshotPolicyAttachmentMapOutput {
	return i.ToEnterpriseSnapshotPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i EnterpriseSnapshotPolicyAttachmentMap) ToEnterpriseSnapshotPolicyAttachmentMapOutputWithContext(ctx context.Context) EnterpriseSnapshotPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseSnapshotPolicyAttachmentMapOutput)
}

type EnterpriseSnapshotPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (EnterpriseSnapshotPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseSnapshotPolicyAttachment)(nil)).Elem()
}

func (o EnterpriseSnapshotPolicyAttachmentOutput) ToEnterpriseSnapshotPolicyAttachmentOutput() EnterpriseSnapshotPolicyAttachmentOutput {
	return o
}

func (o EnterpriseSnapshotPolicyAttachmentOutput) ToEnterpriseSnapshotPolicyAttachmentOutputWithContext(ctx context.Context) EnterpriseSnapshotPolicyAttachmentOutput {
	return o
}

// Cloud Disk ID.
func (o EnterpriseSnapshotPolicyAttachmentOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseSnapshotPolicyAttachment) pulumi.StringOutput { return v.DiskId }).(pulumi.StringOutput)
}

// the enterprise snapshot policy id.
func (o EnterpriseSnapshotPolicyAttachmentOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseSnapshotPolicyAttachment) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

type EnterpriseSnapshotPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseSnapshotPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseSnapshotPolicyAttachment)(nil)).Elem()
}

func (o EnterpriseSnapshotPolicyAttachmentArrayOutput) ToEnterpriseSnapshotPolicyAttachmentArrayOutput() EnterpriseSnapshotPolicyAttachmentArrayOutput {
	return o
}

func (o EnterpriseSnapshotPolicyAttachmentArrayOutput) ToEnterpriseSnapshotPolicyAttachmentArrayOutputWithContext(ctx context.Context) EnterpriseSnapshotPolicyAttachmentArrayOutput {
	return o
}

func (o EnterpriseSnapshotPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) EnterpriseSnapshotPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnterpriseSnapshotPolicyAttachment {
		return vs[0].([]*EnterpriseSnapshotPolicyAttachment)[vs[1].(int)]
	}).(EnterpriseSnapshotPolicyAttachmentOutput)
}

type EnterpriseSnapshotPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (EnterpriseSnapshotPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseSnapshotPolicyAttachment)(nil)).Elem()
}

func (o EnterpriseSnapshotPolicyAttachmentMapOutput) ToEnterpriseSnapshotPolicyAttachmentMapOutput() EnterpriseSnapshotPolicyAttachmentMapOutput {
	return o
}

func (o EnterpriseSnapshotPolicyAttachmentMapOutput) ToEnterpriseSnapshotPolicyAttachmentMapOutputWithContext(ctx context.Context) EnterpriseSnapshotPolicyAttachmentMapOutput {
	return o
}

func (o EnterpriseSnapshotPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) EnterpriseSnapshotPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnterpriseSnapshotPolicyAttachment {
		return vs[0].(map[string]*EnterpriseSnapshotPolicyAttachment)[vs[1].(string)]
	}).(EnterpriseSnapshotPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseSnapshotPolicyAttachmentInput)(nil)).Elem(), &EnterpriseSnapshotPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseSnapshotPolicyAttachmentArrayInput)(nil)).Elem(), EnterpriseSnapshotPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseSnapshotPolicyAttachmentMapInput)(nil)).Elem(), EnterpriseSnapshotPolicyAttachmentMap{})
	pulumi.RegisterOutputType(EnterpriseSnapshotPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(EnterpriseSnapshotPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseSnapshotPolicyAttachmentMapOutput{})
}
