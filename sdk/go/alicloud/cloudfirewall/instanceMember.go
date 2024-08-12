// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfirewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Firewall Instance Member resource.
//
// For information about Cloud Firewall Instance Member and how to use it, see [What is Instance Member](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createloadbalancer).
//
// > **NOTE:** Available in v1.194.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudfirewall"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "AliyunTerraform"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Min: 10000,
//				Max: 99999,
//			})
//			if err != nil {
//				return err
//			}
//			defaultAccount, err := resourcemanager.NewAccount(ctx, "default", &resourcemanager.AccountArgs{
//				DisplayName: pulumi.Sprintf("%v-%v", name, _default.Result),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudfirewall.NewInstanceMember(ctx, "default", &cloudfirewall.InstanceMemberArgs{
//				MemberDesc: pulumi.Sprintf("%v-%v", name, _default.Result),
//				MemberUid:  defaultAccount.ID(),
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
// Cloud Firewall Instance Member can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cloudfirewall/instanceMember:InstanceMember example <id>
// ```
type InstanceMember struct {
	pulumi.CustomResourceState

	// When the cloud firewall member account was added.> use second-level timestamp format.
	CreateTime pulumi.IntOutput `pulumi:"createTime"`
	// Remarks of cloud firewall member accounts.
	MemberDesc pulumi.StringPtrOutput `pulumi:"memberDesc"`
	// The name of the cloud firewall member account.
	MemberDisplayName pulumi.StringOutput `pulumi:"memberDisplayName"`
	// The UID of the cloud firewall member account.
	MemberUid pulumi.StringOutput `pulumi:"memberUid"`
	// The last modification time of the cloud firewall member account.> use second-level timestamp format.
	ModifyTime pulumi.IntOutput `pulumi:"modifyTime"`
	// The resource attribute field that represents the resource status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewInstanceMember registers a new resource with the given unique name, arguments, and options.
func NewInstanceMember(ctx *pulumi.Context,
	name string, args *InstanceMemberArgs, opts ...pulumi.ResourceOption) (*InstanceMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MemberUid == nil {
		return nil, errors.New("invalid value for required argument 'MemberUid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceMember
	err := ctx.RegisterResource("alicloud:cloudfirewall/instanceMember:InstanceMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceMember gets an existing InstanceMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceMemberState, opts ...pulumi.ResourceOption) (*InstanceMember, error) {
	var resource InstanceMember
	err := ctx.ReadResource("alicloud:cloudfirewall/instanceMember:InstanceMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceMember resources.
type instanceMemberState struct {
	// When the cloud firewall member account was added.> use second-level timestamp format.
	CreateTime *int `pulumi:"createTime"`
	// Remarks of cloud firewall member accounts.
	MemberDesc *string `pulumi:"memberDesc"`
	// The name of the cloud firewall member account.
	MemberDisplayName *string `pulumi:"memberDisplayName"`
	// The UID of the cloud firewall member account.
	MemberUid *string `pulumi:"memberUid"`
	// The last modification time of the cloud firewall member account.> use second-level timestamp format.
	ModifyTime *int `pulumi:"modifyTime"`
	// The resource attribute field that represents the resource status.
	Status *string `pulumi:"status"`
}

type InstanceMemberState struct {
	// When the cloud firewall member account was added.> use second-level timestamp format.
	CreateTime pulumi.IntPtrInput
	// Remarks of cloud firewall member accounts.
	MemberDesc pulumi.StringPtrInput
	// The name of the cloud firewall member account.
	MemberDisplayName pulumi.StringPtrInput
	// The UID of the cloud firewall member account.
	MemberUid pulumi.StringPtrInput
	// The last modification time of the cloud firewall member account.> use second-level timestamp format.
	ModifyTime pulumi.IntPtrInput
	// The resource attribute field that represents the resource status.
	Status pulumi.StringPtrInput
}

func (InstanceMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceMemberState)(nil)).Elem()
}

type instanceMemberArgs struct {
	// Remarks of cloud firewall member accounts.
	MemberDesc *string `pulumi:"memberDesc"`
	// The UID of the cloud firewall member account.
	MemberUid string `pulumi:"memberUid"`
}

// The set of arguments for constructing a InstanceMember resource.
type InstanceMemberArgs struct {
	// Remarks of cloud firewall member accounts.
	MemberDesc pulumi.StringPtrInput
	// The UID of the cloud firewall member account.
	MemberUid pulumi.StringInput
}

func (InstanceMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceMemberArgs)(nil)).Elem()
}

type InstanceMemberInput interface {
	pulumi.Input

	ToInstanceMemberOutput() InstanceMemberOutput
	ToInstanceMemberOutputWithContext(ctx context.Context) InstanceMemberOutput
}

func (*InstanceMember) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceMember)(nil)).Elem()
}

func (i *InstanceMember) ToInstanceMemberOutput() InstanceMemberOutput {
	return i.ToInstanceMemberOutputWithContext(context.Background())
}

func (i *InstanceMember) ToInstanceMemberOutputWithContext(ctx context.Context) InstanceMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemberOutput)
}

// InstanceMemberArrayInput is an input type that accepts InstanceMemberArray and InstanceMemberArrayOutput values.
// You can construct a concrete instance of `InstanceMemberArrayInput` via:
//
//	InstanceMemberArray{ InstanceMemberArgs{...} }
type InstanceMemberArrayInput interface {
	pulumi.Input

	ToInstanceMemberArrayOutput() InstanceMemberArrayOutput
	ToInstanceMemberArrayOutputWithContext(context.Context) InstanceMemberArrayOutput
}

type InstanceMemberArray []InstanceMemberInput

func (InstanceMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceMember)(nil)).Elem()
}

func (i InstanceMemberArray) ToInstanceMemberArrayOutput() InstanceMemberArrayOutput {
	return i.ToInstanceMemberArrayOutputWithContext(context.Background())
}

func (i InstanceMemberArray) ToInstanceMemberArrayOutputWithContext(ctx context.Context) InstanceMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemberArrayOutput)
}

// InstanceMemberMapInput is an input type that accepts InstanceMemberMap and InstanceMemberMapOutput values.
// You can construct a concrete instance of `InstanceMemberMapInput` via:
//
//	InstanceMemberMap{ "key": InstanceMemberArgs{...} }
type InstanceMemberMapInput interface {
	pulumi.Input

	ToInstanceMemberMapOutput() InstanceMemberMapOutput
	ToInstanceMemberMapOutputWithContext(context.Context) InstanceMemberMapOutput
}

type InstanceMemberMap map[string]InstanceMemberInput

func (InstanceMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceMember)(nil)).Elem()
}

func (i InstanceMemberMap) ToInstanceMemberMapOutput() InstanceMemberMapOutput {
	return i.ToInstanceMemberMapOutputWithContext(context.Background())
}

func (i InstanceMemberMap) ToInstanceMemberMapOutputWithContext(ctx context.Context) InstanceMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemberMapOutput)
}

type InstanceMemberOutput struct{ *pulumi.OutputState }

func (InstanceMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceMember)(nil)).Elem()
}

func (o InstanceMemberOutput) ToInstanceMemberOutput() InstanceMemberOutput {
	return o
}

func (o InstanceMemberOutput) ToInstanceMemberOutputWithContext(ctx context.Context) InstanceMemberOutput {
	return o
}

// When the cloud firewall member account was added.> use second-level timestamp format.
func (o InstanceMemberOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceMember) pulumi.IntOutput { return v.CreateTime }).(pulumi.IntOutput)
}

// Remarks of cloud firewall member accounts.
func (o InstanceMemberOutput) MemberDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceMember) pulumi.StringPtrOutput { return v.MemberDesc }).(pulumi.StringPtrOutput)
}

// The name of the cloud firewall member account.
func (o InstanceMemberOutput) MemberDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceMember) pulumi.StringOutput { return v.MemberDisplayName }).(pulumi.StringOutput)
}

// The UID of the cloud firewall member account.
func (o InstanceMemberOutput) MemberUid() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceMember) pulumi.StringOutput { return v.MemberUid }).(pulumi.StringOutput)
}

// The last modification time of the cloud firewall member account.> use second-level timestamp format.
func (o InstanceMemberOutput) ModifyTime() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceMember) pulumi.IntOutput { return v.ModifyTime }).(pulumi.IntOutput)
}

// The resource attribute field that represents the resource status.
func (o InstanceMemberOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceMember) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type InstanceMemberArrayOutput struct{ *pulumi.OutputState }

func (InstanceMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceMember)(nil)).Elem()
}

func (o InstanceMemberArrayOutput) ToInstanceMemberArrayOutput() InstanceMemberArrayOutput {
	return o
}

func (o InstanceMemberArrayOutput) ToInstanceMemberArrayOutputWithContext(ctx context.Context) InstanceMemberArrayOutput {
	return o
}

func (o InstanceMemberArrayOutput) Index(i pulumi.IntInput) InstanceMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceMember {
		return vs[0].([]*InstanceMember)[vs[1].(int)]
	}).(InstanceMemberOutput)
}

type InstanceMemberMapOutput struct{ *pulumi.OutputState }

func (InstanceMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceMember)(nil)).Elem()
}

func (o InstanceMemberMapOutput) ToInstanceMemberMapOutput() InstanceMemberMapOutput {
	return o
}

func (o InstanceMemberMapOutput) ToInstanceMemberMapOutputWithContext(ctx context.Context) InstanceMemberMapOutput {
	return o
}

func (o InstanceMemberMapOutput) MapIndex(k pulumi.StringInput) InstanceMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceMember {
		return vs[0].(map[string]*InstanceMember)[vs[1].(string)]
	}).(InstanceMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMemberInput)(nil)).Elem(), &InstanceMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMemberArrayInput)(nil)).Elem(), InstanceMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMemberMapInput)(nil)).Elem(), InstanceMemberMap{})
	pulumi.RegisterOutputType(InstanceMemberOutput{})
	pulumi.RegisterOutputType(InstanceMemberArrayOutput{})
	pulumi.RegisterOutputType(InstanceMemberMapOutput{})
}
