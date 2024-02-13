// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package actiontrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **DEPRECATED:**  This resource has been renamed to actiontrail.Trail from version 1.95.0.
//
// Provides a new resource to manage [Action Trail](https://www.alibabacloud.com/help/doc-detail/28804.htm).
//
// > **NOTE:** Available in 1.35.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/actiontrail"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := actiontrail.NewTrailDeprecated(ctx, "foo", &actiontrail.TrailDeprecatedArgs{
//				EventRw:       pulumi.String("Write-test"),
//				OssBucketName: pulumi.Any(alicloud_oss_bucket.Bucket.Id),
//				RoleName:      pulumi.Any(alicloud_ram_role_policy_attachment.Attach.Role_name),
//				OssKeyPrefix:  pulumi.String("at-product-account-audit-B"),
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
// Action trail can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:actiontrail/trailDeprecated:TrailDeprecated foo abc12345678
// ```
//
// Deprecated: Resource renamed to `Trail`
type TrailDeprecated struct {
	pulumi.CustomResourceState

	// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
	EventRw             pulumi.StringPtrOutput `pulumi:"eventRw"`
	IsOrganizationTrail pulumi.BoolPtrOutput   `pulumi:"isOrganizationTrail"`
	// Deprecated: Field 'mns_topic_arn' has been deprecated from version 1.118.0
	MnsTopicArn pulumi.StringPtrOutput `pulumi:"mnsTopicArn"`
	// The name of the trail to be created, which must be unique for an account.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
	OssBucketName pulumi.StringPtrOutput `pulumi:"ossBucketName"`
	// The prefix of the specified OSS bucket name. This parameter can be left empty.
	OssKeyPrefix    pulumi.StringPtrOutput `pulumi:"ossKeyPrefix"`
	OssWriteRoleArn pulumi.StringPtrOutput `pulumi:"ossWriteRoleArn"`
	// The RAM role in ActionTrail permitted by the user.
	//
	// Deprecated: Field 'role_name' has been deprecated from version 1.118.0
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// The unique ARN of the Log Service project.
	SlsProjectArn pulumi.StringPtrOutput `pulumi:"slsProjectArn"`
	// The unique ARN of the Log Service role.
	//
	// > **NOTE:** `slsProjectArn` and `slsWriteRoleArn` should be set or not set at the same time when actiontrail delivers logs.
	SlsWriteRoleArn pulumi.StringOutput    `pulumi:"slsWriteRoleArn"`
	Status          pulumi.StringPtrOutput `pulumi:"status"`
	TrailName       pulumi.StringOutput    `pulumi:"trailName"`
	TrailRegion     pulumi.StringOutput    `pulumi:"trailRegion"`
}

// NewTrailDeprecated registers a new resource with the given unique name, arguments, and options.
func NewTrailDeprecated(ctx *pulumi.Context,
	name string, args *TrailDeprecatedArgs, opts ...pulumi.ResourceOption) (*TrailDeprecated, error) {
	if args == nil {
		args = &TrailDeprecatedArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TrailDeprecated
	err := ctx.RegisterResource("alicloud:actiontrail/trailDeprecated:TrailDeprecated", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrailDeprecated gets an existing TrailDeprecated resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrailDeprecated(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrailDeprecatedState, opts ...pulumi.ResourceOption) (*TrailDeprecated, error) {
	var resource TrailDeprecated
	err := ctx.ReadResource("alicloud:actiontrail/trailDeprecated:TrailDeprecated", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrailDeprecated resources.
type trailDeprecatedState struct {
	// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
	EventRw             *string `pulumi:"eventRw"`
	IsOrganizationTrail *bool   `pulumi:"isOrganizationTrail"`
	// Deprecated: Field 'mns_topic_arn' has been deprecated from version 1.118.0
	MnsTopicArn *string `pulumi:"mnsTopicArn"`
	// The name of the trail to be created, which must be unique for an account.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
	Name *string `pulumi:"name"`
	// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
	OssBucketName *string `pulumi:"ossBucketName"`
	// The prefix of the specified OSS bucket name. This parameter can be left empty.
	OssKeyPrefix    *string `pulumi:"ossKeyPrefix"`
	OssWriteRoleArn *string `pulumi:"ossWriteRoleArn"`
	// The RAM role in ActionTrail permitted by the user.
	//
	// Deprecated: Field 'role_name' has been deprecated from version 1.118.0
	RoleName *string `pulumi:"roleName"`
	// The unique ARN of the Log Service project.
	SlsProjectArn *string `pulumi:"slsProjectArn"`
	// The unique ARN of the Log Service role.
	//
	// > **NOTE:** `slsProjectArn` and `slsWriteRoleArn` should be set or not set at the same time when actiontrail delivers logs.
	SlsWriteRoleArn *string `pulumi:"slsWriteRoleArn"`
	Status          *string `pulumi:"status"`
	TrailName       *string `pulumi:"trailName"`
	TrailRegion     *string `pulumi:"trailRegion"`
}

type TrailDeprecatedState struct {
	// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
	EventRw             pulumi.StringPtrInput
	IsOrganizationTrail pulumi.BoolPtrInput
	// Deprecated: Field 'mns_topic_arn' has been deprecated from version 1.118.0
	MnsTopicArn pulumi.StringPtrInput
	// The name of the trail to be created, which must be unique for an account.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
	Name pulumi.StringPtrInput
	// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
	OssBucketName pulumi.StringPtrInput
	// The prefix of the specified OSS bucket name. This parameter can be left empty.
	OssKeyPrefix    pulumi.StringPtrInput
	OssWriteRoleArn pulumi.StringPtrInput
	// The RAM role in ActionTrail permitted by the user.
	//
	// Deprecated: Field 'role_name' has been deprecated from version 1.118.0
	RoleName pulumi.StringPtrInput
	// The unique ARN of the Log Service project.
	SlsProjectArn pulumi.StringPtrInput
	// The unique ARN of the Log Service role.
	//
	// > **NOTE:** `slsProjectArn` and `slsWriteRoleArn` should be set or not set at the same time when actiontrail delivers logs.
	SlsWriteRoleArn pulumi.StringPtrInput
	Status          pulumi.StringPtrInput
	TrailName       pulumi.StringPtrInput
	TrailRegion     pulumi.StringPtrInput
}

func (TrailDeprecatedState) ElementType() reflect.Type {
	return reflect.TypeOf((*trailDeprecatedState)(nil)).Elem()
}

type trailDeprecatedArgs struct {
	// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
	EventRw             *string `pulumi:"eventRw"`
	IsOrganizationTrail *bool   `pulumi:"isOrganizationTrail"`
	// Deprecated: Field 'mns_topic_arn' has been deprecated from version 1.118.0
	MnsTopicArn *string `pulumi:"mnsTopicArn"`
	// The name of the trail to be created, which must be unique for an account.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
	Name *string `pulumi:"name"`
	// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
	OssBucketName *string `pulumi:"ossBucketName"`
	// The prefix of the specified OSS bucket name. This parameter can be left empty.
	OssKeyPrefix    *string `pulumi:"ossKeyPrefix"`
	OssWriteRoleArn *string `pulumi:"ossWriteRoleArn"`
	// The RAM role in ActionTrail permitted by the user.
	//
	// Deprecated: Field 'role_name' has been deprecated from version 1.118.0
	RoleName *string `pulumi:"roleName"`
	// The unique ARN of the Log Service project.
	SlsProjectArn *string `pulumi:"slsProjectArn"`
	// The unique ARN of the Log Service role.
	//
	// > **NOTE:** `slsProjectArn` and `slsWriteRoleArn` should be set or not set at the same time when actiontrail delivers logs.
	SlsWriteRoleArn *string `pulumi:"slsWriteRoleArn"`
	Status          *string `pulumi:"status"`
	TrailName       *string `pulumi:"trailName"`
	TrailRegion     *string `pulumi:"trailRegion"`
}

// The set of arguments for constructing a TrailDeprecated resource.
type TrailDeprecatedArgs struct {
	// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
	EventRw             pulumi.StringPtrInput
	IsOrganizationTrail pulumi.BoolPtrInput
	// Deprecated: Field 'mns_topic_arn' has been deprecated from version 1.118.0
	MnsTopicArn pulumi.StringPtrInput
	// The name of the trail to be created, which must be unique for an account.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
	Name pulumi.StringPtrInput
	// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
	OssBucketName pulumi.StringPtrInput
	// The prefix of the specified OSS bucket name. This parameter can be left empty.
	OssKeyPrefix    pulumi.StringPtrInput
	OssWriteRoleArn pulumi.StringPtrInput
	// The RAM role in ActionTrail permitted by the user.
	//
	// Deprecated: Field 'role_name' has been deprecated from version 1.118.0
	RoleName pulumi.StringPtrInput
	// The unique ARN of the Log Service project.
	SlsProjectArn pulumi.StringPtrInput
	// The unique ARN of the Log Service role.
	//
	// > **NOTE:** `slsProjectArn` and `slsWriteRoleArn` should be set or not set at the same time when actiontrail delivers logs.
	SlsWriteRoleArn pulumi.StringPtrInput
	Status          pulumi.StringPtrInput
	TrailName       pulumi.StringPtrInput
	TrailRegion     pulumi.StringPtrInput
}

func (TrailDeprecatedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trailDeprecatedArgs)(nil)).Elem()
}

type TrailDeprecatedInput interface {
	pulumi.Input

	ToTrailDeprecatedOutput() TrailDeprecatedOutput
	ToTrailDeprecatedOutputWithContext(ctx context.Context) TrailDeprecatedOutput
}

func (*TrailDeprecated) ElementType() reflect.Type {
	return reflect.TypeOf((**TrailDeprecated)(nil)).Elem()
}

func (i *TrailDeprecated) ToTrailDeprecatedOutput() TrailDeprecatedOutput {
	return i.ToTrailDeprecatedOutputWithContext(context.Background())
}

func (i *TrailDeprecated) ToTrailDeprecatedOutputWithContext(ctx context.Context) TrailDeprecatedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailDeprecatedOutput)
}

// TrailDeprecatedArrayInput is an input type that accepts TrailDeprecatedArray and TrailDeprecatedArrayOutput values.
// You can construct a concrete instance of `TrailDeprecatedArrayInput` via:
//
//	TrailDeprecatedArray{ TrailDeprecatedArgs{...} }
type TrailDeprecatedArrayInput interface {
	pulumi.Input

	ToTrailDeprecatedArrayOutput() TrailDeprecatedArrayOutput
	ToTrailDeprecatedArrayOutputWithContext(context.Context) TrailDeprecatedArrayOutput
}

type TrailDeprecatedArray []TrailDeprecatedInput

func (TrailDeprecatedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrailDeprecated)(nil)).Elem()
}

func (i TrailDeprecatedArray) ToTrailDeprecatedArrayOutput() TrailDeprecatedArrayOutput {
	return i.ToTrailDeprecatedArrayOutputWithContext(context.Background())
}

func (i TrailDeprecatedArray) ToTrailDeprecatedArrayOutputWithContext(ctx context.Context) TrailDeprecatedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailDeprecatedArrayOutput)
}

// TrailDeprecatedMapInput is an input type that accepts TrailDeprecatedMap and TrailDeprecatedMapOutput values.
// You can construct a concrete instance of `TrailDeprecatedMapInput` via:
//
//	TrailDeprecatedMap{ "key": TrailDeprecatedArgs{...} }
type TrailDeprecatedMapInput interface {
	pulumi.Input

	ToTrailDeprecatedMapOutput() TrailDeprecatedMapOutput
	ToTrailDeprecatedMapOutputWithContext(context.Context) TrailDeprecatedMapOutput
}

type TrailDeprecatedMap map[string]TrailDeprecatedInput

func (TrailDeprecatedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrailDeprecated)(nil)).Elem()
}

func (i TrailDeprecatedMap) ToTrailDeprecatedMapOutput() TrailDeprecatedMapOutput {
	return i.ToTrailDeprecatedMapOutputWithContext(context.Background())
}

func (i TrailDeprecatedMap) ToTrailDeprecatedMapOutputWithContext(ctx context.Context) TrailDeprecatedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailDeprecatedMapOutput)
}

type TrailDeprecatedOutput struct{ *pulumi.OutputState }

func (TrailDeprecatedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrailDeprecated)(nil)).Elem()
}

func (o TrailDeprecatedOutput) ToTrailDeprecatedOutput() TrailDeprecatedOutput {
	return o
}

func (o TrailDeprecatedOutput) ToTrailDeprecatedOutputWithContext(ctx context.Context) TrailDeprecatedOutput {
	return o
}

// Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
func (o TrailDeprecatedOutput) EventRw() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringPtrOutput { return v.EventRw }).(pulumi.StringPtrOutput)
}

func (o TrailDeprecatedOutput) IsOrganizationTrail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.BoolPtrOutput { return v.IsOrganizationTrail }).(pulumi.BoolPtrOutput)
}

// Deprecated: Field 'mns_topic_arn' has been deprecated from version 1.118.0
func (o TrailDeprecatedOutput) MnsTopicArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringPtrOutput { return v.MnsTopicArn }).(pulumi.StringPtrOutput)
}

// The name of the trail to be created, which must be unique for an account.
//
// Deprecated: Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead.
func (o TrailDeprecatedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
func (o TrailDeprecatedOutput) OssBucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringPtrOutput { return v.OssBucketName }).(pulumi.StringPtrOutput)
}

// The prefix of the specified OSS bucket name. This parameter can be left empty.
func (o TrailDeprecatedOutput) OssKeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringPtrOutput { return v.OssKeyPrefix }).(pulumi.StringPtrOutput)
}

func (o TrailDeprecatedOutput) OssWriteRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringPtrOutput { return v.OssWriteRoleArn }).(pulumi.StringPtrOutput)
}

// The RAM role in ActionTrail permitted by the user.
//
// Deprecated: Field 'role_name' has been deprecated from version 1.118.0
func (o TrailDeprecatedOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

// The unique ARN of the Log Service project.
func (o TrailDeprecatedOutput) SlsProjectArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringPtrOutput { return v.SlsProjectArn }).(pulumi.StringPtrOutput)
}

// The unique ARN of the Log Service role.
//
// > **NOTE:** `slsProjectArn` and `slsWriteRoleArn` should be set or not set at the same time when actiontrail delivers logs.
func (o TrailDeprecatedOutput) SlsWriteRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringOutput { return v.SlsWriteRoleArn }).(pulumi.StringOutput)
}

func (o TrailDeprecatedOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TrailDeprecatedOutput) TrailName() pulumi.StringOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringOutput { return v.TrailName }).(pulumi.StringOutput)
}

func (o TrailDeprecatedOutput) TrailRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *TrailDeprecated) pulumi.StringOutput { return v.TrailRegion }).(pulumi.StringOutput)
}

type TrailDeprecatedArrayOutput struct{ *pulumi.OutputState }

func (TrailDeprecatedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrailDeprecated)(nil)).Elem()
}

func (o TrailDeprecatedArrayOutput) ToTrailDeprecatedArrayOutput() TrailDeprecatedArrayOutput {
	return o
}

func (o TrailDeprecatedArrayOutput) ToTrailDeprecatedArrayOutputWithContext(ctx context.Context) TrailDeprecatedArrayOutput {
	return o
}

func (o TrailDeprecatedArrayOutput) Index(i pulumi.IntInput) TrailDeprecatedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrailDeprecated {
		return vs[0].([]*TrailDeprecated)[vs[1].(int)]
	}).(TrailDeprecatedOutput)
}

type TrailDeprecatedMapOutput struct{ *pulumi.OutputState }

func (TrailDeprecatedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrailDeprecated)(nil)).Elem()
}

func (o TrailDeprecatedMapOutput) ToTrailDeprecatedMapOutput() TrailDeprecatedMapOutput {
	return o
}

func (o TrailDeprecatedMapOutput) ToTrailDeprecatedMapOutputWithContext(ctx context.Context) TrailDeprecatedMapOutput {
	return o
}

func (o TrailDeprecatedMapOutput) MapIndex(k pulumi.StringInput) TrailDeprecatedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrailDeprecated {
		return vs[0].(map[string]*TrailDeprecated)[vs[1].(string)]
	}).(TrailDeprecatedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrailDeprecatedInput)(nil)).Elem(), &TrailDeprecated{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailDeprecatedArrayInput)(nil)).Elem(), TrailDeprecatedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailDeprecatedMapInput)(nil)).Elem(), TrailDeprecatedMap{})
	pulumi.RegisterOutputType(TrailDeprecatedOutput{})
	pulumi.RegisterOutputType(TrailDeprecatedArrayOutput{})
	pulumi.RegisterOutputType(TrailDeprecatedMapOutput{})
}
