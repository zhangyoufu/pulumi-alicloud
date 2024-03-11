// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Sag SnatEntry resource. This topic describes how to add a SNAT entry to enable the SNAT function. The SNAT function can hide internal IP addresses and resolve private IP address conflicts. With this function, on-premises sites can access internal IP addresses, but cannot be accessed by internal IP addresses. If you do not add a SNAT entry, on-premises sites can access each other only when all related IP addresses do not conflict.
//
// For information about Sag SnatEntry and how to use it, see [What is Sag SnatEntry](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/addsnatentry).
//
// > **NOTE:** Available since v1.61.0.
//
// > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rocketmq"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			sagId := "sag-9bifk***"
//			if param := cfg.Get("sagId"); param != "" {
//				sagId = param
//			}
//			_, err := rocketmq.NewSnatEntry(ctx, "default", &rocketmq.SnatEntryArgs{
//				SagId:     pulumi.String(sagId),
//				CidrBlock: pulumi.String("192.168.7.0/24"),
//				SnatIp:    pulumi.String("192.0.0.2"),
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
// The Sag SnatEntry can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:rocketmq/snatEntry:SnatEntry example sag-abc123456:snat-abc123456
// ```
type SnatEntry struct {
	pulumi.CustomResourceState

	// The destination CIDR block.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The ID of the SAG instance.
	SagId pulumi.StringOutput `pulumi:"sagId"`
	// The public IP address.
	SnatIp pulumi.StringOutput `pulumi:"snatIp"`
}

// NewSnatEntry registers a new resource with the given unique name, arguments, and options.
func NewSnatEntry(ctx *pulumi.Context,
	name string, args *SnatEntryArgs, opts ...pulumi.ResourceOption) (*SnatEntry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.SagId == nil {
		return nil, errors.New("invalid value for required argument 'SagId'")
	}
	if args.SnatIp == nil {
		return nil, errors.New("invalid value for required argument 'SnatIp'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SnatEntry
	err := ctx.RegisterResource("alicloud:rocketmq/snatEntry:SnatEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnatEntry gets an existing SnatEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnatEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnatEntryState, opts ...pulumi.ResourceOption) (*SnatEntry, error) {
	var resource SnatEntry
	err := ctx.ReadResource("alicloud:rocketmq/snatEntry:SnatEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnatEntry resources.
type snatEntryState struct {
	// The destination CIDR block.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The ID of the SAG instance.
	SagId *string `pulumi:"sagId"`
	// The public IP address.
	SnatIp *string `pulumi:"snatIp"`
}

type SnatEntryState struct {
	// The destination CIDR block.
	CidrBlock pulumi.StringPtrInput
	// The ID of the SAG instance.
	SagId pulumi.StringPtrInput
	// The public IP address.
	SnatIp pulumi.StringPtrInput
}

func (SnatEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*snatEntryState)(nil)).Elem()
}

type snatEntryArgs struct {
	// The destination CIDR block.
	CidrBlock string `pulumi:"cidrBlock"`
	// The ID of the SAG instance.
	SagId string `pulumi:"sagId"`
	// The public IP address.
	SnatIp string `pulumi:"snatIp"`
}

// The set of arguments for constructing a SnatEntry resource.
type SnatEntryArgs struct {
	// The destination CIDR block.
	CidrBlock pulumi.StringInput
	// The ID of the SAG instance.
	SagId pulumi.StringInput
	// The public IP address.
	SnatIp pulumi.StringInput
}

func (SnatEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snatEntryArgs)(nil)).Elem()
}

type SnatEntryInput interface {
	pulumi.Input

	ToSnatEntryOutput() SnatEntryOutput
	ToSnatEntryOutputWithContext(ctx context.Context) SnatEntryOutput
}

func (*SnatEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**SnatEntry)(nil)).Elem()
}

func (i *SnatEntry) ToSnatEntryOutput() SnatEntryOutput {
	return i.ToSnatEntryOutputWithContext(context.Background())
}

func (i *SnatEntry) ToSnatEntryOutputWithContext(ctx context.Context) SnatEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatEntryOutput)
}

// SnatEntryArrayInput is an input type that accepts SnatEntryArray and SnatEntryArrayOutput values.
// You can construct a concrete instance of `SnatEntryArrayInput` via:
//
//	SnatEntryArray{ SnatEntryArgs{...} }
type SnatEntryArrayInput interface {
	pulumi.Input

	ToSnatEntryArrayOutput() SnatEntryArrayOutput
	ToSnatEntryArrayOutputWithContext(context.Context) SnatEntryArrayOutput
}

type SnatEntryArray []SnatEntryInput

func (SnatEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnatEntry)(nil)).Elem()
}

func (i SnatEntryArray) ToSnatEntryArrayOutput() SnatEntryArrayOutput {
	return i.ToSnatEntryArrayOutputWithContext(context.Background())
}

func (i SnatEntryArray) ToSnatEntryArrayOutputWithContext(ctx context.Context) SnatEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatEntryArrayOutput)
}

// SnatEntryMapInput is an input type that accepts SnatEntryMap and SnatEntryMapOutput values.
// You can construct a concrete instance of `SnatEntryMapInput` via:
//
//	SnatEntryMap{ "key": SnatEntryArgs{...} }
type SnatEntryMapInput interface {
	pulumi.Input

	ToSnatEntryMapOutput() SnatEntryMapOutput
	ToSnatEntryMapOutputWithContext(context.Context) SnatEntryMapOutput
}

type SnatEntryMap map[string]SnatEntryInput

func (SnatEntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnatEntry)(nil)).Elem()
}

func (i SnatEntryMap) ToSnatEntryMapOutput() SnatEntryMapOutput {
	return i.ToSnatEntryMapOutputWithContext(context.Background())
}

func (i SnatEntryMap) ToSnatEntryMapOutputWithContext(ctx context.Context) SnatEntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatEntryMapOutput)
}

type SnatEntryOutput struct{ *pulumi.OutputState }

func (SnatEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnatEntry)(nil)).Elem()
}

func (o SnatEntryOutput) ToSnatEntryOutput() SnatEntryOutput {
	return o
}

func (o SnatEntryOutput) ToSnatEntryOutputWithContext(ctx context.Context) SnatEntryOutput {
	return o
}

// The destination CIDR block.
func (o SnatEntryOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *SnatEntry) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// The ID of the SAG instance.
func (o SnatEntryOutput) SagId() pulumi.StringOutput {
	return o.ApplyT(func(v *SnatEntry) pulumi.StringOutput { return v.SagId }).(pulumi.StringOutput)
}

// The public IP address.
func (o SnatEntryOutput) SnatIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SnatEntry) pulumi.StringOutput { return v.SnatIp }).(pulumi.StringOutput)
}

type SnatEntryArrayOutput struct{ *pulumi.OutputState }

func (SnatEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnatEntry)(nil)).Elem()
}

func (o SnatEntryArrayOutput) ToSnatEntryArrayOutput() SnatEntryArrayOutput {
	return o
}

func (o SnatEntryArrayOutput) ToSnatEntryArrayOutputWithContext(ctx context.Context) SnatEntryArrayOutput {
	return o
}

func (o SnatEntryArrayOutput) Index(i pulumi.IntInput) SnatEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnatEntry {
		return vs[0].([]*SnatEntry)[vs[1].(int)]
	}).(SnatEntryOutput)
}

type SnatEntryMapOutput struct{ *pulumi.OutputState }

func (SnatEntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnatEntry)(nil)).Elem()
}

func (o SnatEntryMapOutput) ToSnatEntryMapOutput() SnatEntryMapOutput {
	return o
}

func (o SnatEntryMapOutput) ToSnatEntryMapOutputWithContext(ctx context.Context) SnatEntryMapOutput {
	return o
}

func (o SnatEntryMapOutput) MapIndex(k pulumi.StringInput) SnatEntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnatEntry {
		return vs[0].(map[string]*SnatEntry)[vs[1].(string)]
	}).(SnatEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnatEntryInput)(nil)).Elem(), &SnatEntry{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnatEntryArrayInput)(nil)).Elem(), SnatEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnatEntryMapInput)(nil)).Elem(), SnatEntryMap{})
	pulumi.RegisterOutputType(SnatEntryOutput{})
	pulumi.RegisterOutputType(SnatEntryArrayOutput{})
	pulumi.RegisterOutputType(SnatEntryMapOutput{})
}
