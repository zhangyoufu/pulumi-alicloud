// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudsso

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Cloud SSO SCIM Server Credential resource.
//
// For information about Cloud SSO SCIM Server Credential and how to use it, see [What is Cloud SSO SCIM Server Credential](https://www.alibabacloud.com/help/en/cloudsso/latest/api-cloudsso-2021-05-15-createscimservercredential).
//
// > **NOTE:** Available since v1.138.0.
//
// > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
//
// ## Import
//
// Cloud SSO SCIM Server Credential can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cloudsso/scimServerCredential:ScimServerCredential example <directory_id>:<credential_id>
//
// ```
type ScimServerCredential struct {
	pulumi.CustomResourceState

	// The CredentialId of the resource.
	CredentialId pulumi.StringOutput `pulumi:"credentialId"`
	// The ID of the Directory.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// The Status of the resource. Valid values: `Disabled`, `Enabled`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewScimServerCredential registers a new resource with the given unique name, arguments, and options.
func NewScimServerCredential(ctx *pulumi.Context,
	name string, args *ScimServerCredentialArgs, opts ...pulumi.ResourceOption) (*ScimServerCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScimServerCredential
	err := ctx.RegisterResource("alicloud:cloudsso/scimServerCredential:ScimServerCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScimServerCredential gets an existing ScimServerCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScimServerCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScimServerCredentialState, opts ...pulumi.ResourceOption) (*ScimServerCredential, error) {
	var resource ScimServerCredential
	err := ctx.ReadResource("alicloud:cloudsso/scimServerCredential:ScimServerCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScimServerCredential resources.
type scimServerCredentialState struct {
	// The CredentialId of the resource.
	CredentialId *string `pulumi:"credentialId"`
	// The ID of the Directory.
	DirectoryId *string `pulumi:"directoryId"`
	// The Status of the resource. Valid values: `Disabled`, `Enabled`.
	Status *string `pulumi:"status"`
}

type ScimServerCredentialState struct {
	// The CredentialId of the resource.
	CredentialId pulumi.StringPtrInput
	// The ID of the Directory.
	DirectoryId pulumi.StringPtrInput
	// The Status of the resource. Valid values: `Disabled`, `Enabled`.
	Status pulumi.StringPtrInput
}

func (ScimServerCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*scimServerCredentialState)(nil)).Elem()
}

type scimServerCredentialArgs struct {
	// The ID of the Directory.
	DirectoryId string `pulumi:"directoryId"`
	// The Status of the resource. Valid values: `Disabled`, `Enabled`.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a ScimServerCredential resource.
type ScimServerCredentialArgs struct {
	// The ID of the Directory.
	DirectoryId pulumi.StringInput
	// The Status of the resource. Valid values: `Disabled`, `Enabled`.
	Status pulumi.StringPtrInput
}

func (ScimServerCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scimServerCredentialArgs)(nil)).Elem()
}

type ScimServerCredentialInput interface {
	pulumi.Input

	ToScimServerCredentialOutput() ScimServerCredentialOutput
	ToScimServerCredentialOutputWithContext(ctx context.Context) ScimServerCredentialOutput
}

func (*ScimServerCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**ScimServerCredential)(nil)).Elem()
}

func (i *ScimServerCredential) ToScimServerCredentialOutput() ScimServerCredentialOutput {
	return i.ToScimServerCredentialOutputWithContext(context.Background())
}

func (i *ScimServerCredential) ToScimServerCredentialOutputWithContext(ctx context.Context) ScimServerCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScimServerCredentialOutput)
}

func (i *ScimServerCredential) ToOutput(ctx context.Context) pulumix.Output[*ScimServerCredential] {
	return pulumix.Output[*ScimServerCredential]{
		OutputState: i.ToScimServerCredentialOutputWithContext(ctx).OutputState,
	}
}

// ScimServerCredentialArrayInput is an input type that accepts ScimServerCredentialArray and ScimServerCredentialArrayOutput values.
// You can construct a concrete instance of `ScimServerCredentialArrayInput` via:
//
//	ScimServerCredentialArray{ ScimServerCredentialArgs{...} }
type ScimServerCredentialArrayInput interface {
	pulumi.Input

	ToScimServerCredentialArrayOutput() ScimServerCredentialArrayOutput
	ToScimServerCredentialArrayOutputWithContext(context.Context) ScimServerCredentialArrayOutput
}

type ScimServerCredentialArray []ScimServerCredentialInput

func (ScimServerCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScimServerCredential)(nil)).Elem()
}

func (i ScimServerCredentialArray) ToScimServerCredentialArrayOutput() ScimServerCredentialArrayOutput {
	return i.ToScimServerCredentialArrayOutputWithContext(context.Background())
}

func (i ScimServerCredentialArray) ToScimServerCredentialArrayOutputWithContext(ctx context.Context) ScimServerCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScimServerCredentialArrayOutput)
}

func (i ScimServerCredentialArray) ToOutput(ctx context.Context) pulumix.Output[[]*ScimServerCredential] {
	return pulumix.Output[[]*ScimServerCredential]{
		OutputState: i.ToScimServerCredentialArrayOutputWithContext(ctx).OutputState,
	}
}

// ScimServerCredentialMapInput is an input type that accepts ScimServerCredentialMap and ScimServerCredentialMapOutput values.
// You can construct a concrete instance of `ScimServerCredentialMapInput` via:
//
//	ScimServerCredentialMap{ "key": ScimServerCredentialArgs{...} }
type ScimServerCredentialMapInput interface {
	pulumi.Input

	ToScimServerCredentialMapOutput() ScimServerCredentialMapOutput
	ToScimServerCredentialMapOutputWithContext(context.Context) ScimServerCredentialMapOutput
}

type ScimServerCredentialMap map[string]ScimServerCredentialInput

func (ScimServerCredentialMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScimServerCredential)(nil)).Elem()
}

func (i ScimServerCredentialMap) ToScimServerCredentialMapOutput() ScimServerCredentialMapOutput {
	return i.ToScimServerCredentialMapOutputWithContext(context.Background())
}

func (i ScimServerCredentialMap) ToScimServerCredentialMapOutputWithContext(ctx context.Context) ScimServerCredentialMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScimServerCredentialMapOutput)
}

func (i ScimServerCredentialMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ScimServerCredential] {
	return pulumix.Output[map[string]*ScimServerCredential]{
		OutputState: i.ToScimServerCredentialMapOutputWithContext(ctx).OutputState,
	}
}

type ScimServerCredentialOutput struct{ *pulumi.OutputState }

func (ScimServerCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScimServerCredential)(nil)).Elem()
}

func (o ScimServerCredentialOutput) ToScimServerCredentialOutput() ScimServerCredentialOutput {
	return o
}

func (o ScimServerCredentialOutput) ToScimServerCredentialOutputWithContext(ctx context.Context) ScimServerCredentialOutput {
	return o
}

func (o ScimServerCredentialOutput) ToOutput(ctx context.Context) pulumix.Output[*ScimServerCredential] {
	return pulumix.Output[*ScimServerCredential]{
		OutputState: o.OutputState,
	}
}

// The CredentialId of the resource.
func (o ScimServerCredentialOutput) CredentialId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScimServerCredential) pulumi.StringOutput { return v.CredentialId }).(pulumi.StringOutput)
}

// The ID of the Directory.
func (o ScimServerCredentialOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScimServerCredential) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

// The Status of the resource. Valid values: `Disabled`, `Enabled`.
func (o ScimServerCredentialOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ScimServerCredential) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ScimServerCredentialArrayOutput struct{ *pulumi.OutputState }

func (ScimServerCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScimServerCredential)(nil)).Elem()
}

func (o ScimServerCredentialArrayOutput) ToScimServerCredentialArrayOutput() ScimServerCredentialArrayOutput {
	return o
}

func (o ScimServerCredentialArrayOutput) ToScimServerCredentialArrayOutputWithContext(ctx context.Context) ScimServerCredentialArrayOutput {
	return o
}

func (o ScimServerCredentialArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ScimServerCredential] {
	return pulumix.Output[[]*ScimServerCredential]{
		OutputState: o.OutputState,
	}
}

func (o ScimServerCredentialArrayOutput) Index(i pulumi.IntInput) ScimServerCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScimServerCredential {
		return vs[0].([]*ScimServerCredential)[vs[1].(int)]
	}).(ScimServerCredentialOutput)
}

type ScimServerCredentialMapOutput struct{ *pulumi.OutputState }

func (ScimServerCredentialMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScimServerCredential)(nil)).Elem()
}

func (o ScimServerCredentialMapOutput) ToScimServerCredentialMapOutput() ScimServerCredentialMapOutput {
	return o
}

func (o ScimServerCredentialMapOutput) ToScimServerCredentialMapOutputWithContext(ctx context.Context) ScimServerCredentialMapOutput {
	return o
}

func (o ScimServerCredentialMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ScimServerCredential] {
	return pulumix.Output[map[string]*ScimServerCredential]{
		OutputState: o.OutputState,
	}
}

func (o ScimServerCredentialMapOutput) MapIndex(k pulumi.StringInput) ScimServerCredentialOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScimServerCredential {
		return vs[0].(map[string]*ScimServerCredential)[vs[1].(string)]
	}).(ScimServerCredentialOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScimServerCredentialInput)(nil)).Elem(), &ScimServerCredential{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScimServerCredentialArrayInput)(nil)).Elem(), ScimServerCredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScimServerCredentialMapInput)(nil)).Elem(), ScimServerCredentialMap{})
	pulumi.RegisterOutputType(ScimServerCredentialOutput{})
	pulumi.RegisterOutputType(ScimServerCredentialArrayOutput{})
	pulumi.RegisterOutputType(ScimServerCredentialMapOutput{})
}
