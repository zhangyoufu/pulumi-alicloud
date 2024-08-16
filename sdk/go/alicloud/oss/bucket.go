// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a oss bucket and set its attribution.
//
// > **NOTE:** The bucket namespace is shared by all users of the OSS system. Please set bucket name as unique as possible.
//
// > **NOTE:** Available since v1.2.0.
//
// ## Example Usage
//
// # Private Bucket
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Max: 99999,
//				Min: 10000,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucket(ctx, "bucket-acl", &oss.BucketArgs{
//				Bucket: pulumi.Sprintf("example-value-%v", _default.Result),
//				Acl:    pulumi.String("private"),
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
// # Static Website
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Max: 99999,
//				Min: 10000,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucket(ctx, "bucket-website", &oss.BucketArgs{
//				Bucket: pulumi.Sprintf("example-value-%v", _default.Result),
//				Website: &oss.BucketWebsiteArgs{
//					IndexDocument: pulumi.String("index.html"),
//					ErrorDocument: pulumi.String("error.html"),
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
// # Enable Logging
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Max: 99999,
//				Min: 10000,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucket(ctx, "bucket-target", &oss.BucketArgs{
//				Bucket: pulumi.Sprintf("example-value-%v", _default.Result),
//				Acl:    pulumi.String("public-read"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucket(ctx, "bucket-logging", &oss.BucketArgs{
//				Bucket: pulumi.Sprintf("example-logging-%v", _default.Result),
//				Logging: &oss.BucketLoggingTypeArgs{
//					TargetBucket: bucket_target.ID(),
//					TargetPrefix: pulumi.String("log/"),
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
// # Referer configuration
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Max: 99999,
//				Min: 10000,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucket(ctx, "bucket-referer", &oss.BucketArgs{
//				Bucket: pulumi.Sprintf("example-value-%v", _default.Result),
//				Acl:    pulumi.String("private"),
//				RefererConfig: &oss.BucketRefererConfigArgs{
//					AllowEmpty: pulumi.Bool(false),
//					Referers: pulumi.StringArray{
//						pulumi.String("http://www.aliyun.com"),
//						pulumi.String("https://www.aliyun.com"),
//					},
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
// # Set lifecycle rule
//
// ## Import
//
// OSS bucket can be imported using the bucket name, e.g.
//
// ```sh
// $ pulumi import alicloud:oss/bucket:Bucket bucket bucket-12345678
// ```
type Bucket struct {
	pulumi.CustomResourceState

	// A access monitor status of a bucket. See `accessMonitor` below.
	AccessMonitor BucketAccessMonitorTypeOutput `pulumi:"accessMonitor"`
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". This property has been deprecated since 1.220.0, please use the resource `oss.BucketAcl` instead.
	//
	// Deprecated: Field 'acl' has been deprecated since provider version 1.220.0. New resource 'alicloud_oss_bucket_acl' instead.
	Acl    pulumi.StringOutput    `pulumi:"acl"`
	Bucket pulumi.StringPtrOutput `pulumi:"bucket"`
	// A rule of  [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm). The items of core rule are no more than 10 for every OSS bucket. See `corsRule` below.
	CorsRules BucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// The creation date of the bucket.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The extranet access endpoint of the bucket.
	ExtranetEndpoint pulumi.StringOutput `pulumi:"extranetEndpoint"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// The intranet access endpoint of the bucket.
	IntranetEndpoint pulumi.StringOutput `pulumi:"intranetEndpoint"`
	// A boolean that indicates lifecycle rules allow prefix overlap.
	LifecycleRuleAllowSameActionOverlap pulumi.BoolPtrOutput `pulumi:"lifecycleRuleAllowSameActionOverlap"`
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm). See `lifecycleRule` below.
	LifecycleRules BucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// The location of the bucket.
	Location pulumi.StringOutput `pulumi:"location"`
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm). See `logging` below.
	Logging BucketLoggingTypePtrOutput `pulumi:"logging"`
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable pulumi.BoolPtrOutput `pulumi:"loggingIsenable"`
	// The bucket owner.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketPolicy` instead.
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType pulumi.StringPtrOutput `pulumi:"redundancyType"`
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketReferer` instead. See `refererConfig` below.
	RefererConfig BucketRefererConfigPtrOutput `pulumi:"refererConfig"`
	// The ID of the resource group to which the bucket belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// A configuration of server-side encryption. See `serverSideEncryptionRule` below.
	ServerSideEncryptionRule BucketServerSideEncryptionRulePtrOutput `pulumi:"serverSideEncryptionRule"`
	// The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA", "Archive", "ColdArchive" and "DeepColdArchive". Defaults to "Standard". "ColdArchive" is available since 1.203.0. "DeepColdArchive" is available since 1.209.0.
	StorageClass pulumi.StringPtrOutput `pulumi:"storageClass"`
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A transfer acceleration status of a bucket. See `transferAcceleration` below.
	TransferAcceleration BucketTransferAccelerationTypePtrOutput `pulumi:"transferAcceleration"`
	// A state of versioning. See `versioning` below.
	Versioning BucketVersioningTypePtrOutput `pulumi:"versioning"`
	// A website configuration. See `website` below.
	Website BucketWebsitePtrOutput `pulumi:"website"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		args = &BucketArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bucket
	err := ctx.RegisterResource("alicloud:oss/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("alicloud:oss/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// A access monitor status of a bucket. See `accessMonitor` below.
	AccessMonitor *BucketAccessMonitorType `pulumi:"accessMonitor"`
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". This property has been deprecated since 1.220.0, please use the resource `oss.BucketAcl` instead.
	//
	// Deprecated: Field 'acl' has been deprecated since provider version 1.220.0. New resource 'alicloud_oss_bucket_acl' instead.
	Acl    *string `pulumi:"acl"`
	Bucket *string `pulumi:"bucket"`
	// A rule of  [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm). The items of core rule are no more than 10 for every OSS bucket. See `corsRule` below.
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// The creation date of the bucket.
	CreationDate *string `pulumi:"creationDate"`
	// The extranet access endpoint of the bucket.
	ExtranetEndpoint *string `pulumi:"extranetEndpoint"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The intranet access endpoint of the bucket.
	IntranetEndpoint *string `pulumi:"intranetEndpoint"`
	// A boolean that indicates lifecycle rules allow prefix overlap.
	LifecycleRuleAllowSameActionOverlap *bool `pulumi:"lifecycleRuleAllowSameActionOverlap"`
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm). See `lifecycleRule` below.
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The location of the bucket.
	Location *string `pulumi:"location"`
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm). See `logging` below.
	Logging *BucketLoggingType `pulumi:"logging"`
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable *bool `pulumi:"loggingIsenable"`
	// The bucket owner.
	Owner *string `pulumi:"owner"`
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketPolicy` instead.
	Policy *string `pulumi:"policy"`
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType *string `pulumi:"redundancyType"`
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketReferer` instead. See `refererConfig` below.
	RefererConfig *BucketRefererConfig `pulumi:"refererConfig"`
	// The ID of the resource group to which the bucket belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A configuration of server-side encryption. See `serverSideEncryptionRule` below.
	ServerSideEncryptionRule *BucketServerSideEncryptionRule `pulumi:"serverSideEncryptionRule"`
	// The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA", "Archive", "ColdArchive" and "DeepColdArchive". Defaults to "Standard". "ColdArchive" is available since 1.203.0. "DeepColdArchive" is available since 1.209.0.
	StorageClass *string `pulumi:"storageClass"`
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags map[string]string `pulumi:"tags"`
	// A transfer acceleration status of a bucket. See `transferAcceleration` below.
	TransferAcceleration *BucketTransferAccelerationType `pulumi:"transferAcceleration"`
	// A state of versioning. See `versioning` below.
	Versioning *BucketVersioningType `pulumi:"versioning"`
	// A website configuration. See `website` below.
	Website *BucketWebsite `pulumi:"website"`
}

type BucketState struct {
	// A access monitor status of a bucket. See `accessMonitor` below.
	AccessMonitor BucketAccessMonitorTypePtrInput
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". This property has been deprecated since 1.220.0, please use the resource `oss.BucketAcl` instead.
	//
	// Deprecated: Field 'acl' has been deprecated since provider version 1.220.0. New resource 'alicloud_oss_bucket_acl' instead.
	Acl    pulumi.StringPtrInput
	Bucket pulumi.StringPtrInput
	// A rule of  [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm). The items of core rule are no more than 10 for every OSS bucket. See `corsRule` below.
	CorsRules BucketCorsRuleArrayInput
	// The creation date of the bucket.
	CreationDate pulumi.StringPtrInput
	// The extranet access endpoint of the bucket.
	ExtranetEndpoint pulumi.StringPtrInput
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy pulumi.BoolPtrInput
	// The intranet access endpoint of the bucket.
	IntranetEndpoint pulumi.StringPtrInput
	// A boolean that indicates lifecycle rules allow prefix overlap.
	LifecycleRuleAllowSameActionOverlap pulumi.BoolPtrInput
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm). See `lifecycleRule` below.
	LifecycleRules BucketLifecycleRuleArrayInput
	// The location of the bucket.
	Location pulumi.StringPtrInput
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm). See `logging` below.
	Logging BucketLoggingTypePtrInput
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable pulumi.BoolPtrInput
	// The bucket owner.
	Owner pulumi.StringPtrInput
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketPolicy` instead.
	Policy pulumi.StringPtrInput
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType pulumi.StringPtrInput
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketReferer` instead. See `refererConfig` below.
	RefererConfig BucketRefererConfigPtrInput
	// The ID of the resource group to which the bucket belongs.
	ResourceGroupId pulumi.StringPtrInput
	// A configuration of server-side encryption. See `serverSideEncryptionRule` below.
	ServerSideEncryptionRule BucketServerSideEncryptionRulePtrInput
	// The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA", "Archive", "ColdArchive" and "DeepColdArchive". Defaults to "Standard". "ColdArchive" is available since 1.203.0. "DeepColdArchive" is available since 1.209.0.
	StorageClass pulumi.StringPtrInput
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags pulumi.StringMapInput
	// A transfer acceleration status of a bucket. See `transferAcceleration` below.
	TransferAcceleration BucketTransferAccelerationTypePtrInput
	// A state of versioning. See `versioning` below.
	Versioning BucketVersioningTypePtrInput
	// A website configuration. See `website` below.
	Website BucketWebsitePtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// A access monitor status of a bucket. See `accessMonitor` below.
	AccessMonitor *BucketAccessMonitorType `pulumi:"accessMonitor"`
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". This property has been deprecated since 1.220.0, please use the resource `oss.BucketAcl` instead.
	//
	// Deprecated: Field 'acl' has been deprecated since provider version 1.220.0. New resource 'alicloud_oss_bucket_acl' instead.
	Acl    *string `pulumi:"acl"`
	Bucket *string `pulumi:"bucket"`
	// A rule of  [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm). The items of core rule are no more than 10 for every OSS bucket. See `corsRule` below.
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A boolean that indicates lifecycle rules allow prefix overlap.
	LifecycleRuleAllowSameActionOverlap *bool `pulumi:"lifecycleRuleAllowSameActionOverlap"`
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm). See `lifecycleRule` below.
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm). See `logging` below.
	Logging *BucketLoggingType `pulumi:"logging"`
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable *bool `pulumi:"loggingIsenable"`
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketPolicy` instead.
	Policy *string `pulumi:"policy"`
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType *string `pulumi:"redundancyType"`
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketReferer` instead. See `refererConfig` below.
	RefererConfig *BucketRefererConfig `pulumi:"refererConfig"`
	// The ID of the resource group to which the bucket belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A configuration of server-side encryption. See `serverSideEncryptionRule` below.
	ServerSideEncryptionRule *BucketServerSideEncryptionRule `pulumi:"serverSideEncryptionRule"`
	// The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA", "Archive", "ColdArchive" and "DeepColdArchive". Defaults to "Standard". "ColdArchive" is available since 1.203.0. "DeepColdArchive" is available since 1.209.0.
	StorageClass *string `pulumi:"storageClass"`
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags map[string]string `pulumi:"tags"`
	// A transfer acceleration status of a bucket. See `transferAcceleration` below.
	TransferAcceleration *BucketTransferAccelerationType `pulumi:"transferAcceleration"`
	// A state of versioning. See `versioning` below.
	Versioning *BucketVersioningType `pulumi:"versioning"`
	// A website configuration. See `website` below.
	Website *BucketWebsite `pulumi:"website"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// A access monitor status of a bucket. See `accessMonitor` below.
	AccessMonitor BucketAccessMonitorTypePtrInput
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". This property has been deprecated since 1.220.0, please use the resource `oss.BucketAcl` instead.
	//
	// Deprecated: Field 'acl' has been deprecated since provider version 1.220.0. New resource 'alicloud_oss_bucket_acl' instead.
	Acl    pulumi.StringPtrInput
	Bucket pulumi.StringPtrInput
	// A rule of  [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm). The items of core rule are no more than 10 for every OSS bucket. See `corsRule` below.
	CorsRules BucketCorsRuleArrayInput
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy pulumi.BoolPtrInput
	// A boolean that indicates lifecycle rules allow prefix overlap.
	LifecycleRuleAllowSameActionOverlap pulumi.BoolPtrInput
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm). See `lifecycleRule` below.
	LifecycleRules BucketLifecycleRuleArrayInput
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm). See `logging` below.
	Logging BucketLoggingTypePtrInput
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable pulumi.BoolPtrInput
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketPolicy` instead.
	Policy pulumi.StringPtrInput
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType pulumi.StringPtrInput
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketReferer` instead. See `refererConfig` below.
	RefererConfig BucketRefererConfigPtrInput
	// The ID of the resource group to which the bucket belongs.
	ResourceGroupId pulumi.StringPtrInput
	// A configuration of server-side encryption. See `serverSideEncryptionRule` below.
	ServerSideEncryptionRule BucketServerSideEncryptionRulePtrInput
	// The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA", "Archive", "ColdArchive" and "DeepColdArchive". Defaults to "Standard". "ColdArchive" is available since 1.203.0. "DeepColdArchive" is available since 1.209.0.
	StorageClass pulumi.StringPtrInput
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags pulumi.StringMapInput
	// A transfer acceleration status of a bucket. See `transferAcceleration` below.
	TransferAcceleration BucketTransferAccelerationTypePtrInput
	// A state of versioning. See `versioning` below.
	Versioning BucketVersioningTypePtrInput
	// A website configuration. See `website` below.
	Website BucketWebsitePtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

// BucketArrayInput is an input type that accepts BucketArray and BucketArrayOutput values.
// You can construct a concrete instance of `BucketArrayInput` via:
//
//	BucketArray{ BucketArgs{...} }
type BucketArrayInput interface {
	pulumi.Input

	ToBucketArrayOutput() BucketArrayOutput
	ToBucketArrayOutputWithContext(context.Context) BucketArrayOutput
}

type BucketArray []BucketInput

func (BucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (i BucketArray) ToBucketArrayOutput() BucketArrayOutput {
	return i.ToBucketArrayOutputWithContext(context.Background())
}

func (i BucketArray) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketArrayOutput)
}

// BucketMapInput is an input type that accepts BucketMap and BucketMapOutput values.
// You can construct a concrete instance of `BucketMapInput` via:
//
//	BucketMap{ "key": BucketArgs{...} }
type BucketMapInput interface {
	pulumi.Input

	ToBucketMapOutput() BucketMapOutput
	ToBucketMapOutputWithContext(context.Context) BucketMapOutput
}

type BucketMap map[string]BucketInput

func (BucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (i BucketMap) ToBucketMapOutput() BucketMapOutput {
	return i.ToBucketMapOutputWithContext(context.Background())
}

func (i BucketMap) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMapOutput)
}

type BucketOutput struct{ *pulumi.OutputState }

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

// A access monitor status of a bucket. See `accessMonitor` below.
func (o BucketOutput) AccessMonitor() BucketAccessMonitorTypeOutput {
	return o.ApplyT(func(v *Bucket) BucketAccessMonitorTypeOutput { return v.AccessMonitor }).(BucketAccessMonitorTypeOutput)
}

// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". This property has been deprecated since 1.220.0, please use the resource `oss.BucketAcl` instead.
//
// Deprecated: Field 'acl' has been deprecated since provider version 1.220.0. New resource 'alicloud_oss_bucket_acl' instead.
func (o BucketOutput) Acl() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Acl }).(pulumi.StringOutput)
}

func (o BucketOutput) Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.Bucket }).(pulumi.StringPtrOutput)
}

// A rule of  [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm). The items of core rule are no more than 10 for every OSS bucket. See `corsRule` below.
func (o BucketOutput) CorsRules() BucketCorsRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketCorsRuleArrayOutput { return v.CorsRules }).(BucketCorsRuleArrayOutput)
}

// The creation date of the bucket.
func (o BucketOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The extranet access endpoint of the bucket.
func (o BucketOutput) ExtranetEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.ExtranetEndpoint }).(pulumi.StringOutput)
}

// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
func (o BucketOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// The intranet access endpoint of the bucket.
func (o BucketOutput) IntranetEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.IntranetEndpoint }).(pulumi.StringOutput)
}

// A boolean that indicates lifecycle rules allow prefix overlap.
func (o BucketOutput) LifecycleRuleAllowSameActionOverlap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.LifecycleRuleAllowSameActionOverlap }).(pulumi.BoolPtrOutput)
}

// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm). See `lifecycleRule` below.
func (o BucketOutput) LifecycleRules() BucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketLifecycleRuleArrayOutput { return v.LifecycleRules }).(BucketLifecycleRuleArrayOutput)
}

// The location of the bucket.
func (o BucketOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm). See `logging` below.
func (o BucketOutput) Logging() BucketLoggingTypePtrOutput {
	return o.ApplyT(func(v *Bucket) BucketLoggingTypePtrOutput { return v.Logging }).(BucketLoggingTypePtrOutput)
}

// The flag of using logging enable container. Defaults true.
//
// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
func (o BucketOutput) LoggingIsenable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.LoggingIsenable }).(pulumi.BoolPtrOutput)
}

// The bucket owner.
func (o BucketOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketPolicy` instead.
func (o BucketOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.Policy }).(pulumi.StringPtrOutput)
}

// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
func (o BucketOutput) RedundancyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.RedundancyType }).(pulumi.StringPtrOutput)
}

// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm). This property has been deprecated since 1.220.0, please use the resource `oss.BucketReferer` instead. See `refererConfig` below.
func (o BucketOutput) RefererConfig() BucketRefererConfigPtrOutput {
	return o.ApplyT(func(v *Bucket) BucketRefererConfigPtrOutput { return v.RefererConfig }).(BucketRefererConfigPtrOutput)
}

// The ID of the resource group to which the bucket belongs.
func (o BucketOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// A configuration of server-side encryption. See `serverSideEncryptionRule` below.
func (o BucketOutput) ServerSideEncryptionRule() BucketServerSideEncryptionRulePtrOutput {
	return o.ApplyT(func(v *Bucket) BucketServerSideEncryptionRulePtrOutput { return v.ServerSideEncryptionRule }).(BucketServerSideEncryptionRulePtrOutput)
}

// The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA", "Archive", "ColdArchive" and "DeepColdArchive". Defaults to "Standard". "ColdArchive" is available since 1.203.0. "DeepColdArchive" is available since 1.209.0.
func (o BucketOutput) StorageClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.StorageClass }).(pulumi.StringPtrOutput)
}

// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
func (o BucketOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A transfer acceleration status of a bucket. See `transferAcceleration` below.
func (o BucketOutput) TransferAcceleration() BucketTransferAccelerationTypePtrOutput {
	return o.ApplyT(func(v *Bucket) BucketTransferAccelerationTypePtrOutput { return v.TransferAcceleration }).(BucketTransferAccelerationTypePtrOutput)
}

// A state of versioning. See `versioning` below.
func (o BucketOutput) Versioning() BucketVersioningTypePtrOutput {
	return o.ApplyT(func(v *Bucket) BucketVersioningTypePtrOutput { return v.Versioning }).(BucketVersioningTypePtrOutput)
}

// A website configuration. See `website` below.
func (o BucketOutput) Website() BucketWebsitePtrOutput {
	return o.ApplyT(func(v *Bucket) BucketWebsitePtrOutput { return v.Website }).(BucketWebsitePtrOutput)
}

type BucketArrayOutput struct{ *pulumi.OutputState }

func (BucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (o BucketArrayOutput) ToBucketArrayOutput() BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) Index(i pulumi.IntInput) BucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].([]*Bucket)[vs[1].(int)]
	}).(BucketOutput)
}

type BucketMapOutput struct{ *pulumi.OutputState }

func (BucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (o BucketMapOutput) ToBucketMapOutput() BucketMapOutput {
	return o
}

func (o BucketMapOutput) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return o
}

func (o BucketMapOutput) MapIndex(k pulumi.StringInput) BucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].(map[string]*Bucket)[vs[1].(string)]
	}).(BucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInput)(nil)).Elem(), &Bucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketArrayInput)(nil)).Elem(), BucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketMapInput)(nil)).Elem(), BucketMap{})
	pulumi.RegisterOutputType(BucketOutput{})
	pulumi.RegisterOutputType(BucketArrayOutput{})
	pulumi.RegisterOutputType(BucketMapOutput{})
}
