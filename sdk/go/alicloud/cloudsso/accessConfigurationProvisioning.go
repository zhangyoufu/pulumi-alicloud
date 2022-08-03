// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudsso

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud SSO Access Configuration Provisioning resource.
//
// For information about Cloud SSO Access Configuration Provisioning and how to use it, see [What is Access Configuration Provisioning](https://www.alibabacloud.com/help/en/doc-detail/266737.html).
//
// > **NOTE:** Available in v1.148.0+.
//
// ## Import
//
// Cloud SSO Access Configuration Provisioning can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cloudsso/accessConfigurationProvisioning:AccessConfigurationProvisioning example <directory_id>:<access_configuration_id>:<target_type>:<target_id>
// ```
type AccessConfigurationProvisioning struct {
	pulumi.CustomResourceState

	// The Access configuration ID.
	AccessConfigurationId pulumi.StringOutput `pulumi:"accessConfigurationId"`
	// The ID of the Directory.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// The status of the resource. Valid values: `Provisioned`, `ReprovisionRequired` and `DeprovisionFailed`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the target to create the resource range.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
	// The type of the resource range target to be accessed. Valid values: `RD-Account`.
	TargetType pulumi.StringOutput `pulumi:"targetType"`
}

// NewAccessConfigurationProvisioning registers a new resource with the given unique name, arguments, and options.
func NewAccessConfigurationProvisioning(ctx *pulumi.Context,
	name string, args *AccessConfigurationProvisioningArgs, opts ...pulumi.ResourceOption) (*AccessConfigurationProvisioning, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessConfigurationId == nil {
		return nil, errors.New("invalid value for required argument 'AccessConfigurationId'")
	}
	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	if args.TargetId == nil {
		return nil, errors.New("invalid value for required argument 'TargetId'")
	}
	if args.TargetType == nil {
		return nil, errors.New("invalid value for required argument 'TargetType'")
	}
	var resource AccessConfigurationProvisioning
	err := ctx.RegisterResource("alicloud:cloudsso/accessConfigurationProvisioning:AccessConfigurationProvisioning", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessConfigurationProvisioning gets an existing AccessConfigurationProvisioning resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessConfigurationProvisioning(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessConfigurationProvisioningState, opts ...pulumi.ResourceOption) (*AccessConfigurationProvisioning, error) {
	var resource AccessConfigurationProvisioning
	err := ctx.ReadResource("alicloud:cloudsso/accessConfigurationProvisioning:AccessConfigurationProvisioning", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessConfigurationProvisioning resources.
type accessConfigurationProvisioningState struct {
	// The Access configuration ID.
	AccessConfigurationId *string `pulumi:"accessConfigurationId"`
	// The ID of the Directory.
	DirectoryId *string `pulumi:"directoryId"`
	// The status of the resource. Valid values: `Provisioned`, `ReprovisionRequired` and `DeprovisionFailed`.
	Status *string `pulumi:"status"`
	// The ID of the target to create the resource range.
	TargetId *string `pulumi:"targetId"`
	// The type of the resource range target to be accessed. Valid values: `RD-Account`.
	TargetType *string `pulumi:"targetType"`
}

type AccessConfigurationProvisioningState struct {
	// The Access configuration ID.
	AccessConfigurationId pulumi.StringPtrInput
	// The ID of the Directory.
	DirectoryId pulumi.StringPtrInput
	// The status of the resource. Valid values: `Provisioned`, `ReprovisionRequired` and `DeprovisionFailed`.
	Status pulumi.StringPtrInput
	// The ID of the target to create the resource range.
	TargetId pulumi.StringPtrInput
	// The type of the resource range target to be accessed. Valid values: `RD-Account`.
	TargetType pulumi.StringPtrInput
}

func (AccessConfigurationProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessConfigurationProvisioningState)(nil)).Elem()
}

type accessConfigurationProvisioningArgs struct {
	// The Access configuration ID.
	AccessConfigurationId string `pulumi:"accessConfigurationId"`
	// The ID of the Directory.
	DirectoryId string `pulumi:"directoryId"`
	// The status of the resource. Valid values: `Provisioned`, `ReprovisionRequired` and `DeprovisionFailed`.
	Status *string `pulumi:"status"`
	// The ID of the target to create the resource range.
	TargetId string `pulumi:"targetId"`
	// The type of the resource range target to be accessed. Valid values: `RD-Account`.
	TargetType string `pulumi:"targetType"`
}

// The set of arguments for constructing a AccessConfigurationProvisioning resource.
type AccessConfigurationProvisioningArgs struct {
	// The Access configuration ID.
	AccessConfigurationId pulumi.StringInput
	// The ID of the Directory.
	DirectoryId pulumi.StringInput
	// The status of the resource. Valid values: `Provisioned`, `ReprovisionRequired` and `DeprovisionFailed`.
	Status pulumi.StringPtrInput
	// The ID of the target to create the resource range.
	TargetId pulumi.StringInput
	// The type of the resource range target to be accessed. Valid values: `RD-Account`.
	TargetType pulumi.StringInput
}

func (AccessConfigurationProvisioningArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessConfigurationProvisioningArgs)(nil)).Elem()
}

type AccessConfigurationProvisioningInput interface {
	pulumi.Input

	ToAccessConfigurationProvisioningOutput() AccessConfigurationProvisioningOutput
	ToAccessConfigurationProvisioningOutputWithContext(ctx context.Context) AccessConfigurationProvisioningOutput
}

func (*AccessConfigurationProvisioning) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessConfigurationProvisioning)(nil)).Elem()
}

func (i *AccessConfigurationProvisioning) ToAccessConfigurationProvisioningOutput() AccessConfigurationProvisioningOutput {
	return i.ToAccessConfigurationProvisioningOutputWithContext(context.Background())
}

func (i *AccessConfigurationProvisioning) ToAccessConfigurationProvisioningOutputWithContext(ctx context.Context) AccessConfigurationProvisioningOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessConfigurationProvisioningOutput)
}

// AccessConfigurationProvisioningArrayInput is an input type that accepts AccessConfigurationProvisioningArray and AccessConfigurationProvisioningArrayOutput values.
// You can construct a concrete instance of `AccessConfigurationProvisioningArrayInput` via:
//
//          AccessConfigurationProvisioningArray{ AccessConfigurationProvisioningArgs{...} }
type AccessConfigurationProvisioningArrayInput interface {
	pulumi.Input

	ToAccessConfigurationProvisioningArrayOutput() AccessConfigurationProvisioningArrayOutput
	ToAccessConfigurationProvisioningArrayOutputWithContext(context.Context) AccessConfigurationProvisioningArrayOutput
}

type AccessConfigurationProvisioningArray []AccessConfigurationProvisioningInput

func (AccessConfigurationProvisioningArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessConfigurationProvisioning)(nil)).Elem()
}

func (i AccessConfigurationProvisioningArray) ToAccessConfigurationProvisioningArrayOutput() AccessConfigurationProvisioningArrayOutput {
	return i.ToAccessConfigurationProvisioningArrayOutputWithContext(context.Background())
}

func (i AccessConfigurationProvisioningArray) ToAccessConfigurationProvisioningArrayOutputWithContext(ctx context.Context) AccessConfigurationProvisioningArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessConfigurationProvisioningArrayOutput)
}

// AccessConfigurationProvisioningMapInput is an input type that accepts AccessConfigurationProvisioningMap and AccessConfigurationProvisioningMapOutput values.
// You can construct a concrete instance of `AccessConfigurationProvisioningMapInput` via:
//
//          AccessConfigurationProvisioningMap{ "key": AccessConfigurationProvisioningArgs{...} }
type AccessConfigurationProvisioningMapInput interface {
	pulumi.Input

	ToAccessConfigurationProvisioningMapOutput() AccessConfigurationProvisioningMapOutput
	ToAccessConfigurationProvisioningMapOutputWithContext(context.Context) AccessConfigurationProvisioningMapOutput
}

type AccessConfigurationProvisioningMap map[string]AccessConfigurationProvisioningInput

func (AccessConfigurationProvisioningMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessConfigurationProvisioning)(nil)).Elem()
}

func (i AccessConfigurationProvisioningMap) ToAccessConfigurationProvisioningMapOutput() AccessConfigurationProvisioningMapOutput {
	return i.ToAccessConfigurationProvisioningMapOutputWithContext(context.Background())
}

func (i AccessConfigurationProvisioningMap) ToAccessConfigurationProvisioningMapOutputWithContext(ctx context.Context) AccessConfigurationProvisioningMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessConfigurationProvisioningMapOutput)
}

type AccessConfigurationProvisioningOutput struct{ *pulumi.OutputState }

func (AccessConfigurationProvisioningOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessConfigurationProvisioning)(nil)).Elem()
}

func (o AccessConfigurationProvisioningOutput) ToAccessConfigurationProvisioningOutput() AccessConfigurationProvisioningOutput {
	return o
}

func (o AccessConfigurationProvisioningOutput) ToAccessConfigurationProvisioningOutputWithContext(ctx context.Context) AccessConfigurationProvisioningOutput {
	return o
}

// The Access configuration ID.
func (o AccessConfigurationProvisioningOutput) AccessConfigurationId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConfigurationProvisioning) pulumi.StringOutput { return v.AccessConfigurationId }).(pulumi.StringOutput)
}

// The ID of the Directory.
func (o AccessConfigurationProvisioningOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConfigurationProvisioning) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

// The status of the resource. Valid values: `Provisioned`, `ReprovisionRequired` and `DeprovisionFailed`.
func (o AccessConfigurationProvisioningOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConfigurationProvisioning) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the target to create the resource range.
func (o AccessConfigurationProvisioningOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConfigurationProvisioning) pulumi.StringOutput { return v.TargetId }).(pulumi.StringOutput)
}

// The type of the resource range target to be accessed. Valid values: `RD-Account`.
func (o AccessConfigurationProvisioningOutput) TargetType() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConfigurationProvisioning) pulumi.StringOutput { return v.TargetType }).(pulumi.StringOutput)
}

type AccessConfigurationProvisioningArrayOutput struct{ *pulumi.OutputState }

func (AccessConfigurationProvisioningArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessConfigurationProvisioning)(nil)).Elem()
}

func (o AccessConfigurationProvisioningArrayOutput) ToAccessConfigurationProvisioningArrayOutput() AccessConfigurationProvisioningArrayOutput {
	return o
}

func (o AccessConfigurationProvisioningArrayOutput) ToAccessConfigurationProvisioningArrayOutputWithContext(ctx context.Context) AccessConfigurationProvisioningArrayOutput {
	return o
}

func (o AccessConfigurationProvisioningArrayOutput) Index(i pulumi.IntInput) AccessConfigurationProvisioningOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessConfigurationProvisioning {
		return vs[0].([]*AccessConfigurationProvisioning)[vs[1].(int)]
	}).(AccessConfigurationProvisioningOutput)
}

type AccessConfigurationProvisioningMapOutput struct{ *pulumi.OutputState }

func (AccessConfigurationProvisioningMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessConfigurationProvisioning)(nil)).Elem()
}

func (o AccessConfigurationProvisioningMapOutput) ToAccessConfigurationProvisioningMapOutput() AccessConfigurationProvisioningMapOutput {
	return o
}

func (o AccessConfigurationProvisioningMapOutput) ToAccessConfigurationProvisioningMapOutputWithContext(ctx context.Context) AccessConfigurationProvisioningMapOutput {
	return o
}

func (o AccessConfigurationProvisioningMapOutput) MapIndex(k pulumi.StringInput) AccessConfigurationProvisioningOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessConfigurationProvisioning {
		return vs[0].(map[string]*AccessConfigurationProvisioning)[vs[1].(string)]
	}).(AccessConfigurationProvisioningOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessConfigurationProvisioningInput)(nil)).Elem(), &AccessConfigurationProvisioning{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessConfigurationProvisioningArrayInput)(nil)).Elem(), AccessConfigurationProvisioningArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessConfigurationProvisioningMapInput)(nil)).Elem(), AccessConfigurationProvisioningMap{})
	pulumi.RegisterOutputType(AccessConfigurationProvisioningOutput{})
	pulumi.RegisterOutputType(AccessConfigurationProvisioningArrayOutput{})
	pulumi.RegisterOutputType(AccessConfigurationProvisioningMapOutput{})
}
