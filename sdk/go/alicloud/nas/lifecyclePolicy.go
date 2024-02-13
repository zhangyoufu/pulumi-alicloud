// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Network Attached Storage (NAS) Lifecycle Policy resource.
//
// For information about Network Attached Storage (NAS) Lifecycle Policy and how to use it, see [What is Lifecycle Policy](https://www.alibabacloud.com/help/en/doc-detail/169362.html).
//
// > **NOTE:** Available in v1.153.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleFileSystem, err := nas.NewFileSystem(ctx, "exampleFileSystem", &nas.FileSystemArgs{
//				ProtocolType: pulumi.String("NFS"),
//				StorageType:  pulumi.String("Capacity"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = nas.NewLifecyclePolicy(ctx, "exampleLifecyclePolicy", &nas.LifecyclePolicyArgs{
//				FileSystemId:        exampleFileSystem.ID(),
//				LifecyclePolicyName: pulumi.String("terraform-example"),
//				LifecycleRuleName:   pulumi.String("DEFAULT_ATIME_14"),
//				StorageType:         pulumi.String("InfrequentAccess"),
//				Paths: pulumi.StringArray{
//					pulumi.String("/"),
//				},
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
// Network Attached Storage (NAS) Lifecycle Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:nas/lifecyclePolicy:LifecyclePolicy example <file_system_id>:<lifecycle_policy_name>
// ```
type LifecyclePolicy struct {
	pulumi.CustomResourceState

	// The ID of the file system.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// The name of the lifecycle management policy.
	LifecyclePolicyName pulumi.StringOutput `pulumi:"lifecyclePolicyName"`
	// The rules in the lifecycle management policy. Valid values: `DEFAULT_ATIME_14`, `DEFAULT_ATIME_30`, `DEFAULT_ATIME_60`, `DEFAULT_ATIME_90`.
	LifecycleRuleName pulumi.StringOutput `pulumi:"lifecycleRuleName"`
	// The absolute path of the directory for which the lifecycle management policy is configured. Set a maximum of `10` path. The path value must be prefixed by a forward slash (/) and must be an existing path in the mount target.
	Paths pulumi.StringArrayOutput `pulumi:"paths"`
	// The storage type of the data that is dumped to the IA storage medium. Valid values: `InfrequentAccess`.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
}

// NewLifecyclePolicy registers a new resource with the given unique name, arguments, and options.
func NewLifecyclePolicy(ctx *pulumi.Context,
	name string, args *LifecyclePolicyArgs, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	if args.LifecyclePolicyName == nil {
		return nil, errors.New("invalid value for required argument 'LifecyclePolicyName'")
	}
	if args.LifecycleRuleName == nil {
		return nil, errors.New("invalid value for required argument 'LifecycleRuleName'")
	}
	if args.Paths == nil {
		return nil, errors.New("invalid value for required argument 'Paths'")
	}
	if args.StorageType == nil {
		return nil, errors.New("invalid value for required argument 'StorageType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LifecyclePolicy
	err := ctx.RegisterResource("alicloud:nas/lifecyclePolicy:LifecyclePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLifecyclePolicy gets an existing LifecyclePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLifecyclePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LifecyclePolicyState, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	var resource LifecyclePolicy
	err := ctx.ReadResource("alicloud:nas/lifecyclePolicy:LifecyclePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LifecyclePolicy resources.
type lifecyclePolicyState struct {
	// The ID of the file system.
	FileSystemId *string `pulumi:"fileSystemId"`
	// The name of the lifecycle management policy.
	LifecyclePolicyName *string `pulumi:"lifecyclePolicyName"`
	// The rules in the lifecycle management policy. Valid values: `DEFAULT_ATIME_14`, `DEFAULT_ATIME_30`, `DEFAULT_ATIME_60`, `DEFAULT_ATIME_90`.
	LifecycleRuleName *string `pulumi:"lifecycleRuleName"`
	// The absolute path of the directory for which the lifecycle management policy is configured. Set a maximum of `10` path. The path value must be prefixed by a forward slash (/) and must be an existing path in the mount target.
	Paths []string `pulumi:"paths"`
	// The storage type of the data that is dumped to the IA storage medium. Valid values: `InfrequentAccess`.
	StorageType *string `pulumi:"storageType"`
}

type LifecyclePolicyState struct {
	// The ID of the file system.
	FileSystemId pulumi.StringPtrInput
	// The name of the lifecycle management policy.
	LifecyclePolicyName pulumi.StringPtrInput
	// The rules in the lifecycle management policy. Valid values: `DEFAULT_ATIME_14`, `DEFAULT_ATIME_30`, `DEFAULT_ATIME_60`, `DEFAULT_ATIME_90`.
	LifecycleRuleName pulumi.StringPtrInput
	// The absolute path of the directory for which the lifecycle management policy is configured. Set a maximum of `10` path. The path value must be prefixed by a forward slash (/) and must be an existing path in the mount target.
	Paths pulumi.StringArrayInput
	// The storage type of the data that is dumped to the IA storage medium. Valid values: `InfrequentAccess`.
	StorageType pulumi.StringPtrInput
}

func (LifecyclePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyState)(nil)).Elem()
}

type lifecyclePolicyArgs struct {
	// The ID of the file system.
	FileSystemId string `pulumi:"fileSystemId"`
	// The name of the lifecycle management policy.
	LifecyclePolicyName string `pulumi:"lifecyclePolicyName"`
	// The rules in the lifecycle management policy. Valid values: `DEFAULT_ATIME_14`, `DEFAULT_ATIME_30`, `DEFAULT_ATIME_60`, `DEFAULT_ATIME_90`.
	LifecycleRuleName string `pulumi:"lifecycleRuleName"`
	// The absolute path of the directory for which the lifecycle management policy is configured. Set a maximum of `10` path. The path value must be prefixed by a forward slash (/) and must be an existing path in the mount target.
	Paths []string `pulumi:"paths"`
	// The storage type of the data that is dumped to the IA storage medium. Valid values: `InfrequentAccess`.
	StorageType string `pulumi:"storageType"`
}

// The set of arguments for constructing a LifecyclePolicy resource.
type LifecyclePolicyArgs struct {
	// The ID of the file system.
	FileSystemId pulumi.StringInput
	// The name of the lifecycle management policy.
	LifecyclePolicyName pulumi.StringInput
	// The rules in the lifecycle management policy. Valid values: `DEFAULT_ATIME_14`, `DEFAULT_ATIME_30`, `DEFAULT_ATIME_60`, `DEFAULT_ATIME_90`.
	LifecycleRuleName pulumi.StringInput
	// The absolute path of the directory for which the lifecycle management policy is configured. Set a maximum of `10` path. The path value must be prefixed by a forward slash (/) and must be an existing path in the mount target.
	Paths pulumi.StringArrayInput
	// The storage type of the data that is dumped to the IA storage medium. Valid values: `InfrequentAccess`.
	StorageType pulumi.StringInput
}

func (LifecyclePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyArgs)(nil)).Elem()
}

type LifecyclePolicyInput interface {
	pulumi.Input

	ToLifecyclePolicyOutput() LifecyclePolicyOutput
	ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput
}

func (*LifecyclePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**LifecyclePolicy)(nil)).Elem()
}

func (i *LifecyclePolicy) ToLifecyclePolicyOutput() LifecyclePolicyOutput {
	return i.ToLifecyclePolicyOutputWithContext(context.Background())
}

func (i *LifecyclePolicy) ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyOutput)
}

// LifecyclePolicyArrayInput is an input type that accepts LifecyclePolicyArray and LifecyclePolicyArrayOutput values.
// You can construct a concrete instance of `LifecyclePolicyArrayInput` via:
//
//	LifecyclePolicyArray{ LifecyclePolicyArgs{...} }
type LifecyclePolicyArrayInput interface {
	pulumi.Input

	ToLifecyclePolicyArrayOutput() LifecyclePolicyArrayOutput
	ToLifecyclePolicyArrayOutputWithContext(context.Context) LifecyclePolicyArrayOutput
}

type LifecyclePolicyArray []LifecyclePolicyInput

func (LifecyclePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LifecyclePolicy)(nil)).Elem()
}

func (i LifecyclePolicyArray) ToLifecyclePolicyArrayOutput() LifecyclePolicyArrayOutput {
	return i.ToLifecyclePolicyArrayOutputWithContext(context.Background())
}

func (i LifecyclePolicyArray) ToLifecyclePolicyArrayOutputWithContext(ctx context.Context) LifecyclePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyArrayOutput)
}

// LifecyclePolicyMapInput is an input type that accepts LifecyclePolicyMap and LifecyclePolicyMapOutput values.
// You can construct a concrete instance of `LifecyclePolicyMapInput` via:
//
//	LifecyclePolicyMap{ "key": LifecyclePolicyArgs{...} }
type LifecyclePolicyMapInput interface {
	pulumi.Input

	ToLifecyclePolicyMapOutput() LifecyclePolicyMapOutput
	ToLifecyclePolicyMapOutputWithContext(context.Context) LifecyclePolicyMapOutput
}

type LifecyclePolicyMap map[string]LifecyclePolicyInput

func (LifecyclePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LifecyclePolicy)(nil)).Elem()
}

func (i LifecyclePolicyMap) ToLifecyclePolicyMapOutput() LifecyclePolicyMapOutput {
	return i.ToLifecyclePolicyMapOutputWithContext(context.Background())
}

func (i LifecyclePolicyMap) ToLifecyclePolicyMapOutputWithContext(ctx context.Context) LifecyclePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyMapOutput)
}

type LifecyclePolicyOutput struct{ *pulumi.OutputState }

func (LifecyclePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LifecyclePolicy)(nil)).Elem()
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyOutput() LifecyclePolicyOutput {
	return o
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput {
	return o
}

// The ID of the file system.
func (o LifecyclePolicyOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// The name of the lifecycle management policy.
func (o LifecyclePolicyOutput) LifecyclePolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.LifecyclePolicyName }).(pulumi.StringOutput)
}

// The rules in the lifecycle management policy. Valid values: `DEFAULT_ATIME_14`, `DEFAULT_ATIME_30`, `DEFAULT_ATIME_60`, `DEFAULT_ATIME_90`.
func (o LifecyclePolicyOutput) LifecycleRuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.LifecycleRuleName }).(pulumi.StringOutput)
}

// The absolute path of the directory for which the lifecycle management policy is configured. Set a maximum of `10` path. The path value must be prefixed by a forward slash (/) and must be an existing path in the mount target.
func (o LifecyclePolicyOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringArrayOutput { return v.Paths }).(pulumi.StringArrayOutput)
}

// The storage type of the data that is dumped to the IA storage medium. Valid values: `InfrequentAccess`.
func (o LifecyclePolicyOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LifecyclePolicy) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

type LifecyclePolicyArrayOutput struct{ *pulumi.OutputState }

func (LifecyclePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LifecyclePolicy)(nil)).Elem()
}

func (o LifecyclePolicyArrayOutput) ToLifecyclePolicyArrayOutput() LifecyclePolicyArrayOutput {
	return o
}

func (o LifecyclePolicyArrayOutput) ToLifecyclePolicyArrayOutputWithContext(ctx context.Context) LifecyclePolicyArrayOutput {
	return o
}

func (o LifecyclePolicyArrayOutput) Index(i pulumi.IntInput) LifecyclePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LifecyclePolicy {
		return vs[0].([]*LifecyclePolicy)[vs[1].(int)]
	}).(LifecyclePolicyOutput)
}

type LifecyclePolicyMapOutput struct{ *pulumi.OutputState }

func (LifecyclePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LifecyclePolicy)(nil)).Elem()
}

func (o LifecyclePolicyMapOutput) ToLifecyclePolicyMapOutput() LifecyclePolicyMapOutput {
	return o
}

func (o LifecyclePolicyMapOutput) ToLifecyclePolicyMapOutputWithContext(ctx context.Context) LifecyclePolicyMapOutput {
	return o
}

func (o LifecyclePolicyMapOutput) MapIndex(k pulumi.StringInput) LifecyclePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LifecyclePolicy {
		return vs[0].(map[string]*LifecyclePolicy)[vs[1].(string)]
	}).(LifecyclePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LifecyclePolicyInput)(nil)).Elem(), &LifecyclePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*LifecyclePolicyArrayInput)(nil)).Elem(), LifecyclePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LifecyclePolicyMapInput)(nil)).Elem(), LifecyclePolicyMap{})
	pulumi.RegisterOutputType(LifecyclePolicyOutput{})
	pulumi.RegisterOutputType(LifecyclePolicyArrayOutput{})
	pulumi.RegisterOutputType(LifecyclePolicyMapOutput{})
}
