// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a MongoDB Audit Policy resource.
//
// For information about MongoDB Audit Policy and how to use it, see [What is Audit Policy](https://www.alibabacloud.com/help/doc-detail/131941.html).
//
// > **NOTE:** Available since v1.148.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mongodb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
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
//			defaultZones, err := mongodb.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			index := len(defaultZones.Zones) - 1
//			zoneId := defaultZones.Zones[index].Id
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("172.17.3.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("172.17.3.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      *pulumi.String(zoneId),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := mongodb.NewInstance(ctx, "defaultInstance", &mongodb.InstanceArgs{
//				EngineVersion:     pulumi.String("4.2"),
//				DbInstanceClass:   pulumi.String("dds.mongo.mid"),
//				DbInstanceStorage: pulumi.Int(10),
//				VswitchId:         defaultSwitch.ID(),
//				SecurityIpLists: pulumi.StringArray{
//					pulumi.String("10.168.1.12"),
//					pulumi.String("100.69.7.112"),
//				},
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mongodb.NewAuditPolicy(ctx, "defaultAuditPolicy", &mongodb.AuditPolicyArgs{
//				DbInstanceId: defaultInstance.ID(),
//				AuditStatus:  pulumi.String("disabled"),
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
// MongoDB Audit Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:mongodb/auditPolicy:AuditPolicy example <db_instance_id>
// ```
type AuditPolicy struct {
	pulumi.CustomResourceState

	// The status of the audit log. Valid values: `disabled`, `enable`.
	AuditStatus pulumi.StringOutput `pulumi:"auditStatus"`
	// The ID of the instance.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
	StoragePeriod pulumi.IntPtrOutput `pulumi:"storagePeriod"`
}

// NewAuditPolicy registers a new resource with the given unique name, arguments, and options.
func NewAuditPolicy(ctx *pulumi.Context,
	name string, args *AuditPolicyArgs, opts ...pulumi.ResourceOption) (*AuditPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuditStatus == nil {
		return nil, errors.New("invalid value for required argument 'AuditStatus'")
	}
	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuditPolicy
	err := ctx.RegisterResource("alicloud:mongodb/auditPolicy:AuditPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuditPolicy gets an existing AuditPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuditPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuditPolicyState, opts ...pulumi.ResourceOption) (*AuditPolicy, error) {
	var resource AuditPolicy
	err := ctx.ReadResource("alicloud:mongodb/auditPolicy:AuditPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuditPolicy resources.
type auditPolicyState struct {
	// The status of the audit log. Valid values: `disabled`, `enable`.
	AuditStatus *string `pulumi:"auditStatus"`
	// The ID of the instance.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
	StoragePeriod *int `pulumi:"storagePeriod"`
}

type AuditPolicyState struct {
	// The status of the audit log. Valid values: `disabled`, `enable`.
	AuditStatus pulumi.StringPtrInput
	// The ID of the instance.
	DbInstanceId pulumi.StringPtrInput
	// The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
	StoragePeriod pulumi.IntPtrInput
}

func (AuditPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*auditPolicyState)(nil)).Elem()
}

type auditPolicyArgs struct {
	// The status of the audit log. Valid values: `disabled`, `enable`.
	AuditStatus string `pulumi:"auditStatus"`
	// The ID of the instance.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
	StoragePeriod *int `pulumi:"storagePeriod"`
}

// The set of arguments for constructing a AuditPolicy resource.
type AuditPolicyArgs struct {
	// The status of the audit log. Valid values: `disabled`, `enable`.
	AuditStatus pulumi.StringInput
	// The ID of the instance.
	DbInstanceId pulumi.StringInput
	// The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
	StoragePeriod pulumi.IntPtrInput
}

func (AuditPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*auditPolicyArgs)(nil)).Elem()
}

type AuditPolicyInput interface {
	pulumi.Input

	ToAuditPolicyOutput() AuditPolicyOutput
	ToAuditPolicyOutputWithContext(ctx context.Context) AuditPolicyOutput
}

func (*AuditPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditPolicy)(nil)).Elem()
}

func (i *AuditPolicy) ToAuditPolicyOutput() AuditPolicyOutput {
	return i.ToAuditPolicyOutputWithContext(context.Background())
}

func (i *AuditPolicy) ToAuditPolicyOutputWithContext(ctx context.Context) AuditPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditPolicyOutput)
}

// AuditPolicyArrayInput is an input type that accepts AuditPolicyArray and AuditPolicyArrayOutput values.
// You can construct a concrete instance of `AuditPolicyArrayInput` via:
//
//	AuditPolicyArray{ AuditPolicyArgs{...} }
type AuditPolicyArrayInput interface {
	pulumi.Input

	ToAuditPolicyArrayOutput() AuditPolicyArrayOutput
	ToAuditPolicyArrayOutputWithContext(context.Context) AuditPolicyArrayOutput
}

type AuditPolicyArray []AuditPolicyInput

func (AuditPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuditPolicy)(nil)).Elem()
}

func (i AuditPolicyArray) ToAuditPolicyArrayOutput() AuditPolicyArrayOutput {
	return i.ToAuditPolicyArrayOutputWithContext(context.Background())
}

func (i AuditPolicyArray) ToAuditPolicyArrayOutputWithContext(ctx context.Context) AuditPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditPolicyArrayOutput)
}

// AuditPolicyMapInput is an input type that accepts AuditPolicyMap and AuditPolicyMapOutput values.
// You can construct a concrete instance of `AuditPolicyMapInput` via:
//
//	AuditPolicyMap{ "key": AuditPolicyArgs{...} }
type AuditPolicyMapInput interface {
	pulumi.Input

	ToAuditPolicyMapOutput() AuditPolicyMapOutput
	ToAuditPolicyMapOutputWithContext(context.Context) AuditPolicyMapOutput
}

type AuditPolicyMap map[string]AuditPolicyInput

func (AuditPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuditPolicy)(nil)).Elem()
}

func (i AuditPolicyMap) ToAuditPolicyMapOutput() AuditPolicyMapOutput {
	return i.ToAuditPolicyMapOutputWithContext(context.Background())
}

func (i AuditPolicyMap) ToAuditPolicyMapOutputWithContext(ctx context.Context) AuditPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditPolicyMapOutput)
}

type AuditPolicyOutput struct{ *pulumi.OutputState }

func (AuditPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditPolicy)(nil)).Elem()
}

func (o AuditPolicyOutput) ToAuditPolicyOutput() AuditPolicyOutput {
	return o
}

func (o AuditPolicyOutput) ToAuditPolicyOutputWithContext(ctx context.Context) AuditPolicyOutput {
	return o
}

// The status of the audit log. Valid values: `disabled`, `enable`.
func (o AuditPolicyOutput) AuditStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditPolicy) pulumi.StringOutput { return v.AuditStatus }).(pulumi.StringOutput)
}

// The ID of the instance.
func (o AuditPolicyOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditPolicy) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
func (o AuditPolicyOutput) StoragePeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuditPolicy) pulumi.IntPtrOutput { return v.StoragePeriod }).(pulumi.IntPtrOutput)
}

type AuditPolicyArrayOutput struct{ *pulumi.OutputState }

func (AuditPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuditPolicy)(nil)).Elem()
}

func (o AuditPolicyArrayOutput) ToAuditPolicyArrayOutput() AuditPolicyArrayOutput {
	return o
}

func (o AuditPolicyArrayOutput) ToAuditPolicyArrayOutputWithContext(ctx context.Context) AuditPolicyArrayOutput {
	return o
}

func (o AuditPolicyArrayOutput) Index(i pulumi.IntInput) AuditPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuditPolicy {
		return vs[0].([]*AuditPolicy)[vs[1].(int)]
	}).(AuditPolicyOutput)
}

type AuditPolicyMapOutput struct{ *pulumi.OutputState }

func (AuditPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuditPolicy)(nil)).Elem()
}

func (o AuditPolicyMapOutput) ToAuditPolicyMapOutput() AuditPolicyMapOutput {
	return o
}

func (o AuditPolicyMapOutput) ToAuditPolicyMapOutputWithContext(ctx context.Context) AuditPolicyMapOutput {
	return o
}

func (o AuditPolicyMapOutput) MapIndex(k pulumi.StringInput) AuditPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuditPolicy {
		return vs[0].(map[string]*AuditPolicy)[vs[1].(string)]
	}).(AuditPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditPolicyInput)(nil)).Elem(), &AuditPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditPolicyArrayInput)(nil)).Elem(), AuditPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditPolicyMapInput)(nil)).Elem(), AuditPolicyMap{})
	pulumi.RegisterOutputType(AuditPolicyOutput{})
	pulumi.RegisterOutputType(AuditPolicyArrayOutput{})
	pulumi.RegisterOutputType(AuditPolicyMapOutput{})
}
