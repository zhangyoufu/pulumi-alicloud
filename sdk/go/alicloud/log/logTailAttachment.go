// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package log

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Logtail access service is a log collection agent provided by Log Service.
// You can use Logtail to collect logs from servers such as Alibaba Cloud Elastic
// Compute Service (ECS) instances in real time in the Log Service console. [Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm)
//
// This resource amis to attach one logtail configure to a machine group.
//
// > **NOTE:** One logtail configure can be attached to multiple machine groups and one machine group can attach several logtail configures.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/logtail_attachment.html.markdown.
type LogTailAttachment struct {
	pulumi.CustomResourceState

	// The Logtail configuration name, which is unique in the same project.
	LogtailConfigName pulumi.StringOutput `pulumi:"logtailConfigName"`
	// The machine group name, which is unique in the same project.
	MachineGroupName pulumi.StringOutput `pulumi:"machineGroupName"`
	// The project name to the log store belongs.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewLogTailAttachment registers a new resource with the given unique name, arguments, and options.
func NewLogTailAttachment(ctx *pulumi.Context,
	name string, args *LogTailAttachmentArgs, opts ...pulumi.ResourceOption) (*LogTailAttachment, error) {
	if args == nil || args.LogtailConfigName == nil {
		return nil, errors.New("missing required argument 'LogtailConfigName'")
	}
	if args == nil || args.MachineGroupName == nil {
		return nil, errors.New("missing required argument 'MachineGroupName'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil {
		args = &LogTailAttachmentArgs{}
	}
	var resource LogTailAttachment
	err := ctx.RegisterResource("alicloud:log/logTailAttachment:LogTailAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogTailAttachment gets an existing LogTailAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogTailAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogTailAttachmentState, opts ...pulumi.ResourceOption) (*LogTailAttachment, error) {
	var resource LogTailAttachment
	err := ctx.ReadResource("alicloud:log/logTailAttachment:LogTailAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogTailAttachment resources.
type logTailAttachmentState struct {
	// The Logtail configuration name, which is unique in the same project.
	LogtailConfigName *string `pulumi:"logtailConfigName"`
	// The machine group name, which is unique in the same project.
	MachineGroupName *string `pulumi:"machineGroupName"`
	// The project name to the log store belongs.
	Project *string `pulumi:"project"`
}

type LogTailAttachmentState struct {
	// The Logtail configuration name, which is unique in the same project.
	LogtailConfigName pulumi.StringPtrInput
	// The machine group name, which is unique in the same project.
	MachineGroupName pulumi.StringPtrInput
	// The project name to the log store belongs.
	Project pulumi.StringPtrInput
}

func (LogTailAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*logTailAttachmentState)(nil)).Elem()
}

type logTailAttachmentArgs struct {
	// The Logtail configuration name, which is unique in the same project.
	LogtailConfigName string `pulumi:"logtailConfigName"`
	// The machine group name, which is unique in the same project.
	MachineGroupName string `pulumi:"machineGroupName"`
	// The project name to the log store belongs.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a LogTailAttachment resource.
type LogTailAttachmentArgs struct {
	// The Logtail configuration name, which is unique in the same project.
	LogtailConfigName pulumi.StringInput
	// The machine group name, which is unique in the same project.
	MachineGroupName pulumi.StringInput
	// The project name to the log store belongs.
	Project pulumi.StringInput
}

func (LogTailAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logTailAttachmentArgs)(nil)).Elem()
}
