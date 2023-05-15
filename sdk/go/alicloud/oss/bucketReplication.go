// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an independent replication configuration resource for OSS bucket.
//
// For information about OSS replication and how to use it, see [What is cross-region replication](https://www.alibabacloud.com/help/doc-detail/31864.html) and [What is same-region replication](https://www.alibabacloud.com/help/doc-detail/254865.html).
//
// > **NOTE:** Available in v1.161.0+.
//
// ## Example Usage
//
// # Set bucket replication configuration
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oss.NewBucketReplication(ctx, "cross-region-replication", &oss.BucketReplicationArgs{
//				Action: pulumi.String("ALL"),
//				Bucket: pulumi.String("bucket-in-hangzhou"),
//				Destination: &oss.BucketReplicationDestinationArgs{
//					Bucket:   pulumi.String("bucket-in-beijing"),
//					Location: pulumi.String("oss-cn-beijing"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucketReplication(ctx, "same-region-replication", &oss.BucketReplicationArgs{
//				Action: pulumi.String("ALL"),
//				Bucket: pulumi.String("bucket-in-hangzhou"),
//				Destination: &oss.BucketReplicationDestinationArgs{
//					Bucket:   pulumi.String("bucket-in-hangzhou-1"),
//					Location: pulumi.String("oss-cn-hangzhou"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucketReplication(ctx, "replication-with-prefix", &oss.BucketReplicationArgs{
//				Action: pulumi.String("ALL"),
//				Bucket: pulumi.String("bucket-1"),
//				Destination: &oss.BucketReplicationDestinationArgs{
//					Bucket:   pulumi.String("bucket-2"),
//					Location: pulumi.String("oss-cn-hangzhou"),
//				},
//				HistoricalObjectReplication: pulumi.String("disabled"),
//				PrefixSet: &oss.BucketReplicationPrefixSetArgs{
//					Prefixes: pulumi.StringArray{
//						pulumi.String("prefix1/"),
//						pulumi.String("prefix2/"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucketReplication(ctx, "replication-with-specific-action", &oss.BucketReplicationArgs{
//				Action: pulumi.String("PUT"),
//				Bucket: pulumi.String("bucket-1"),
//				Destination: &oss.BucketReplicationDestinationArgs{
//					Bucket:   pulumi.String("bucket-2"),
//					Location: pulumi.String("oss-cn-hangzhou"),
//				},
//				HistoricalObjectReplication: pulumi.String("disabled"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucketReplication(ctx, "replication-with-kms-encryption", &oss.BucketReplicationArgs{
//				Action: pulumi.String("ALL"),
//				Bucket: pulumi.String("bucket-1"),
//				Destination: &oss.BucketReplicationDestinationArgs{
//					Bucket:   pulumi.String("bucket-2"),
//					Location: pulumi.String("oss-cn-hangzhou"),
//				},
//				EncryptionConfiguration: &oss.BucketReplicationEncryptionConfigurationArgs{
//					ReplicaKmsKeyId: pulumi.String("<your kms key id>"),
//				},
//				HistoricalObjectReplication: pulumi.String("disabled"),
//				SourceSelectionCriteria: &oss.BucketReplicationSourceSelectionCriteriaArgs{
//					SseKmsEncryptedObjects: &oss.BucketReplicationSourceSelectionCriteriaSseKmsEncryptedObjectsArgs{
//						Status: pulumi.String("Enabled"),
//					},
//				},
//				SyncRole: pulumi.String("<your ram role>"),
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
// ### Timeouts The `timeouts` block allows you to specify timeouts for certain actions* `delete` - (Defaults to 30 mins) Used when delete a data replication rule (until the data replication task is cleared).
type BucketReplication struct {
	pulumi.CustomResourceState

	// The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Specifies the destination for the rule(See the following block `destination`).
	Destination BucketReplicationDestinationOutput `pulumi:"destination"`
	// Specifies the encryption configuration for the objects replicated to the destination bucket(See the following block `encryptionConfiguration`).
	EncryptionConfiguration BucketReplicationEncryptionConfigurationPtrOutput `pulumi:"encryptionConfiguration"`
	// Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
	HistoricalObjectReplication pulumi.StringPtrOutput `pulumi:"historicalObjectReplication"`
	// The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket(See the following block `prefixSet`).
	PrefixSet BucketReplicationPrefixSetPtrOutput `pulumi:"prefixSet"`
	// Retrieves the progress of the data replication task. This status is returned only when the data replication task is in the doing state.
	Progress BucketReplicationProgressOutput `pulumi:"progress"`
	// The ID of the data replication rule.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Specifies other conditions used to filter the source objects to replicate(See the following block `sourceSelectionCriteria`).
	SourceSelectionCriteria BucketReplicationSourceSelectionCriteriaPtrOutput `pulumi:"sourceSelectionCriteria"`
	// Specifies whether to replicate objects encrypted by using SSE-KMS. Can be `Enabled` or `Disabled`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
	SyncRole pulumi.StringPtrOutput `pulumi:"syncRole"`
}

// NewBucketReplication registers a new resource with the given unique name, arguments, and options.
func NewBucketReplication(ctx *pulumi.Context,
	name string, args *BucketReplicationArgs, opts ...pulumi.ResourceOption) (*BucketReplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	var resource BucketReplication
	err := ctx.RegisterResource("alicloud:oss/bucketReplication:BucketReplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketReplication gets an existing BucketReplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketReplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketReplicationState, opts ...pulumi.ResourceOption) (*BucketReplication, error) {
	var resource BucketReplication
	err := ctx.ReadResource("alicloud:oss/bucketReplication:BucketReplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketReplication resources.
type bucketReplicationState struct {
	// The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
	Action *string `pulumi:"action"`
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// Specifies the destination for the rule(See the following block `destination`).
	Destination *BucketReplicationDestination `pulumi:"destination"`
	// Specifies the encryption configuration for the objects replicated to the destination bucket(See the following block `encryptionConfiguration`).
	EncryptionConfiguration *BucketReplicationEncryptionConfiguration `pulumi:"encryptionConfiguration"`
	// Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
	HistoricalObjectReplication *string `pulumi:"historicalObjectReplication"`
	// The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket(See the following block `prefixSet`).
	PrefixSet *BucketReplicationPrefixSet `pulumi:"prefixSet"`
	// Retrieves the progress of the data replication task. This status is returned only when the data replication task is in the doing state.
	Progress *BucketReplicationProgress `pulumi:"progress"`
	// The ID of the data replication rule.
	RuleId *string `pulumi:"ruleId"`
	// Specifies other conditions used to filter the source objects to replicate(See the following block `sourceSelectionCriteria`).
	SourceSelectionCriteria *BucketReplicationSourceSelectionCriteria `pulumi:"sourceSelectionCriteria"`
	// Specifies whether to replicate objects encrypted by using SSE-KMS. Can be `Enabled` or `Disabled`.
	Status *string `pulumi:"status"`
	// Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
	SyncRole *string `pulumi:"syncRole"`
}

type BucketReplicationState struct {
	// The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
	Action pulumi.StringPtrInput
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// Specifies the destination for the rule(See the following block `destination`).
	Destination BucketReplicationDestinationPtrInput
	// Specifies the encryption configuration for the objects replicated to the destination bucket(See the following block `encryptionConfiguration`).
	EncryptionConfiguration BucketReplicationEncryptionConfigurationPtrInput
	// Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
	HistoricalObjectReplication pulumi.StringPtrInput
	// The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket(See the following block `prefixSet`).
	PrefixSet BucketReplicationPrefixSetPtrInput
	// Retrieves the progress of the data replication task. This status is returned only when the data replication task is in the doing state.
	Progress BucketReplicationProgressPtrInput
	// The ID of the data replication rule.
	RuleId pulumi.StringPtrInput
	// Specifies other conditions used to filter the source objects to replicate(See the following block `sourceSelectionCriteria`).
	SourceSelectionCriteria BucketReplicationSourceSelectionCriteriaPtrInput
	// Specifies whether to replicate objects encrypted by using SSE-KMS. Can be `Enabled` or `Disabled`.
	Status pulumi.StringPtrInput
	// Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
	SyncRole pulumi.StringPtrInput
}

func (BucketReplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketReplicationState)(nil)).Elem()
}

type bucketReplicationArgs struct {
	// The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
	Action *string `pulumi:"action"`
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// Specifies the destination for the rule(See the following block `destination`).
	Destination BucketReplicationDestination `pulumi:"destination"`
	// Specifies the encryption configuration for the objects replicated to the destination bucket(See the following block `encryptionConfiguration`).
	EncryptionConfiguration *BucketReplicationEncryptionConfiguration `pulumi:"encryptionConfiguration"`
	// Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
	HistoricalObjectReplication *string `pulumi:"historicalObjectReplication"`
	// The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket(See the following block `prefixSet`).
	PrefixSet *BucketReplicationPrefixSet `pulumi:"prefixSet"`
	// Retrieves the progress of the data replication task. This status is returned only when the data replication task is in the doing state.
	Progress *BucketReplicationProgress `pulumi:"progress"`
	// Specifies other conditions used to filter the source objects to replicate(See the following block `sourceSelectionCriteria`).
	SourceSelectionCriteria *BucketReplicationSourceSelectionCriteria `pulumi:"sourceSelectionCriteria"`
	// Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
	SyncRole *string `pulumi:"syncRole"`
}

// The set of arguments for constructing a BucketReplication resource.
type BucketReplicationArgs struct {
	// The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
	Action pulumi.StringPtrInput
	// The name of the bucket.
	Bucket pulumi.StringInput
	// Specifies the destination for the rule(See the following block `destination`).
	Destination BucketReplicationDestinationInput
	// Specifies the encryption configuration for the objects replicated to the destination bucket(See the following block `encryptionConfiguration`).
	EncryptionConfiguration BucketReplicationEncryptionConfigurationPtrInput
	// Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
	HistoricalObjectReplication pulumi.StringPtrInput
	// The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket(See the following block `prefixSet`).
	PrefixSet BucketReplicationPrefixSetPtrInput
	// Retrieves the progress of the data replication task. This status is returned only when the data replication task is in the doing state.
	Progress BucketReplicationProgressPtrInput
	// Specifies other conditions used to filter the source objects to replicate(See the following block `sourceSelectionCriteria`).
	SourceSelectionCriteria BucketReplicationSourceSelectionCriteriaPtrInput
	// Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
	SyncRole pulumi.StringPtrInput
}

func (BucketReplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketReplicationArgs)(nil)).Elem()
}

type BucketReplicationInput interface {
	pulumi.Input

	ToBucketReplicationOutput() BucketReplicationOutput
	ToBucketReplicationOutputWithContext(ctx context.Context) BucketReplicationOutput
}

func (*BucketReplication) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketReplication)(nil)).Elem()
}

func (i *BucketReplication) ToBucketReplicationOutput() BucketReplicationOutput {
	return i.ToBucketReplicationOutputWithContext(context.Background())
}

func (i *BucketReplication) ToBucketReplicationOutputWithContext(ctx context.Context) BucketReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationOutput)
}

// BucketReplicationArrayInput is an input type that accepts BucketReplicationArray and BucketReplicationArrayOutput values.
// You can construct a concrete instance of `BucketReplicationArrayInput` via:
//
//	BucketReplicationArray{ BucketReplicationArgs{...} }
type BucketReplicationArrayInput interface {
	pulumi.Input

	ToBucketReplicationArrayOutput() BucketReplicationArrayOutput
	ToBucketReplicationArrayOutputWithContext(context.Context) BucketReplicationArrayOutput
}

type BucketReplicationArray []BucketReplicationInput

func (BucketReplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketReplication)(nil)).Elem()
}

func (i BucketReplicationArray) ToBucketReplicationArrayOutput() BucketReplicationArrayOutput {
	return i.ToBucketReplicationArrayOutputWithContext(context.Background())
}

func (i BucketReplicationArray) ToBucketReplicationArrayOutputWithContext(ctx context.Context) BucketReplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationArrayOutput)
}

// BucketReplicationMapInput is an input type that accepts BucketReplicationMap and BucketReplicationMapOutput values.
// You can construct a concrete instance of `BucketReplicationMapInput` via:
//
//	BucketReplicationMap{ "key": BucketReplicationArgs{...} }
type BucketReplicationMapInput interface {
	pulumi.Input

	ToBucketReplicationMapOutput() BucketReplicationMapOutput
	ToBucketReplicationMapOutputWithContext(context.Context) BucketReplicationMapOutput
}

type BucketReplicationMap map[string]BucketReplicationInput

func (BucketReplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketReplication)(nil)).Elem()
}

func (i BucketReplicationMap) ToBucketReplicationMapOutput() BucketReplicationMapOutput {
	return i.ToBucketReplicationMapOutputWithContext(context.Background())
}

func (i BucketReplicationMap) ToBucketReplicationMapOutputWithContext(ctx context.Context) BucketReplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationMapOutput)
}

type BucketReplicationOutput struct{ *pulumi.OutputState }

func (BucketReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketReplication)(nil)).Elem()
}

func (o BucketReplicationOutput) ToBucketReplicationOutput() BucketReplicationOutput {
	return o
}

func (o BucketReplicationOutput) ToBucketReplicationOutputWithContext(ctx context.Context) BucketReplicationOutput {
	return o
}

// The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
func (o BucketReplicationOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketReplication) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

// The name of the bucket.
func (o BucketReplicationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketReplication) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Specifies the destination for the rule(See the following block `destination`).
func (o BucketReplicationOutput) Destination() BucketReplicationDestinationOutput {
	return o.ApplyT(func(v *BucketReplication) BucketReplicationDestinationOutput { return v.Destination }).(BucketReplicationDestinationOutput)
}

// Specifies the encryption configuration for the objects replicated to the destination bucket(See the following block `encryptionConfiguration`).
func (o BucketReplicationOutput) EncryptionConfiguration() BucketReplicationEncryptionConfigurationPtrOutput {
	return o.ApplyT(func(v *BucketReplication) BucketReplicationEncryptionConfigurationPtrOutput {
		return v.EncryptionConfiguration
	}).(BucketReplicationEncryptionConfigurationPtrOutput)
}

// Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
func (o BucketReplicationOutput) HistoricalObjectReplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketReplication) pulumi.StringPtrOutput { return v.HistoricalObjectReplication }).(pulumi.StringPtrOutput)
}

// The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket(See the following block `prefixSet`).
func (o BucketReplicationOutput) PrefixSet() BucketReplicationPrefixSetPtrOutput {
	return o.ApplyT(func(v *BucketReplication) BucketReplicationPrefixSetPtrOutput { return v.PrefixSet }).(BucketReplicationPrefixSetPtrOutput)
}

// Retrieves the progress of the data replication task. This status is returned only when the data replication task is in the doing state.
func (o BucketReplicationOutput) Progress() BucketReplicationProgressOutput {
	return o.ApplyT(func(v *BucketReplication) BucketReplicationProgressOutput { return v.Progress }).(BucketReplicationProgressOutput)
}

// The ID of the data replication rule.
func (o BucketReplicationOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketReplication) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// Specifies other conditions used to filter the source objects to replicate(See the following block `sourceSelectionCriteria`).
func (o BucketReplicationOutput) SourceSelectionCriteria() BucketReplicationSourceSelectionCriteriaPtrOutput {
	return o.ApplyT(func(v *BucketReplication) BucketReplicationSourceSelectionCriteriaPtrOutput {
		return v.SourceSelectionCriteria
	}).(BucketReplicationSourceSelectionCriteriaPtrOutput)
}

// Specifies whether to replicate objects encrypted by using SSE-KMS. Can be `Enabled` or `Disabled`.
func (o BucketReplicationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketReplication) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
func (o BucketReplicationOutput) SyncRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketReplication) pulumi.StringPtrOutput { return v.SyncRole }).(pulumi.StringPtrOutput)
}

type BucketReplicationArrayOutput struct{ *pulumi.OutputState }

func (BucketReplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketReplication)(nil)).Elem()
}

func (o BucketReplicationArrayOutput) ToBucketReplicationArrayOutput() BucketReplicationArrayOutput {
	return o
}

func (o BucketReplicationArrayOutput) ToBucketReplicationArrayOutputWithContext(ctx context.Context) BucketReplicationArrayOutput {
	return o
}

func (o BucketReplicationArrayOutput) Index(i pulumi.IntInput) BucketReplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketReplication {
		return vs[0].([]*BucketReplication)[vs[1].(int)]
	}).(BucketReplicationOutput)
}

type BucketReplicationMapOutput struct{ *pulumi.OutputState }

func (BucketReplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketReplication)(nil)).Elem()
}

func (o BucketReplicationMapOutput) ToBucketReplicationMapOutput() BucketReplicationMapOutput {
	return o
}

func (o BucketReplicationMapOutput) ToBucketReplicationMapOutputWithContext(ctx context.Context) BucketReplicationMapOutput {
	return o
}

func (o BucketReplicationMapOutput) MapIndex(k pulumi.StringInput) BucketReplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketReplication {
		return vs[0].(map[string]*BucketReplication)[vs[1].(string)]
	}).(BucketReplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketReplicationInput)(nil)).Elem(), &BucketReplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketReplicationArrayInput)(nil)).Elem(), BucketReplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketReplicationMapInput)(nil)).Elem(), BucketReplicationMap{})
	pulumi.RegisterOutputType(BucketReplicationOutput{})
	pulumi.RegisterOutputType(BucketReplicationArrayOutput{})
	pulumi.RegisterOutputType(BucketReplicationMapOutput{})
}
