// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create a oss bucket and set its attribution.
//
// > **NOTE:** The bucket namespace is shared by all users of the OSS system. Please set bucket name as unique as possible.
//
// ## Example Usage
//
// Private Bucket
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_acl", &oss.BucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Bucket: pulumi.String("bucket-170309-acl"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Static Website
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_website", &oss.BucketArgs{
// 			Bucket: pulumi.String("bucket-170309-website"),
// 			Website: &oss.BucketWebsiteArgs{
// 				ErrorDocument: pulumi.String("error.html"),
// 				IndexDocument: pulumi.String("index.html"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Enable Logging
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_target", &oss.BucketArgs{
// 			Bucket: pulumi.String("bucket-170309-acl"),
// 			Acl:    pulumi.String("public-read"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = oss.NewBucket(ctx, "bucket_logging", &oss.BucketArgs{
// 			Bucket: pulumi.String("bucket-170309-logging"),
// 			Logging: &oss.BucketLoggingArgs{
// 				TargetBucket: bucket_target.ID(),
// 				TargetPrefix: pulumi.String("log/"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Referer configuration
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_referer", &oss.BucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Bucket: pulumi.String("bucket-170309-referer"),
// 			RefererConfig: &oss.BucketRefererConfigArgs{
// 				AllowEmpty: pulumi.Bool(false),
// 				Referers: pulumi.StringArray{
// 					pulumi.String("http://www.aliyun.com"),
// 					pulumi.String("https://www.aliyun.com"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Set lifecycle rule
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_lifecycle", &oss.BucketArgs{
// 			Acl:    pulumi.String("public-read"),
// 			Bucket: pulumi.String("bucket-170309-lifecycle"),
// 			LifecycleRules: oss.BucketLifecycleRuleArray{
// 				&oss.BucketLifecycleRuleArgs{
// 					Enabled: pulumi.Bool(true),
// 					Id:      pulumi.String("rule-days-transition"),
// 					Prefix:  pulumi.String("path3/"),
// 					Transitions: oss.BucketLifecycleRuleTransitionArray{
// 						&oss.BucketLifecycleRuleTransitionArgs{
// 							CreatedBeforeDate: pulumi.String("2020-11-11"),
// 							StorageClass:      pulumi.String("IA"),
// 						},
// 						&oss.BucketLifecycleRuleTransitionArgs{
// 							CreatedBeforeDate: pulumi.String("2021-11-11"),
// 							StorageClass:      pulumi.String("Archive"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Set bucket policy
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_policy", &oss.BucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Bucket: pulumi.String("bucket-170309-policy"),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v", "  {\"Statement\":\n", "      [{\"Action\":\n", "          [\"oss:PutObject\", \"oss:GetObject\", \"oss:DeleteBucket\"],\n", "        \"Effect\":\"Allow\",\n", "        \"Resource\":\n", "            [\"acs:oss:*:*:*\"]}],\n", "   \"Version\":\"1\"}\n", "  \n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// IA Bucket
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_storageclass", &oss.BucketArgs{
// 			Bucket:       pulumi.String("bucket-170309-storageclass"),
// 			StorageClass: pulumi.String("IA"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Set bucket server-side encryption rule
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_sserule", &oss.BucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Bucket: pulumi.String("bucket-170309-sserule"),
// 			ServerSideEncryptionRule: &oss.BucketServerSideEncryptionRuleArgs{
// 				KmsMasterKeyId: pulumi.String("your kms key id"),
// 				SseAlgorithm:   pulumi.String("KMS"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Set bucket tags
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_tags", &oss.BucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Bucket: pulumi.String("bucket-170309-tags"),
// 			Tags: pulumi.StringMap{
// 				"key1": pulumi.String("value1"),
// 				"key2": pulumi.String("value2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Enable bucket versioning
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_versioning", &oss.BucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Bucket: pulumi.String("bucket-170309-versioning"),
// 			Versioning: &oss.BucketVersioningArgs{
// 				Status: pulumi.String("Enabled"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Set bucket redundancy type
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oss.NewBucket(ctx, "bucket_redundancytype", &oss.BucketArgs{
// 			Bucket:         pulumi.String("bucket_name"),
// 			RedundancyType: pulumi.String("ZRS"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// OSS bucket can be imported using the bucket name, e.g.
//
// ```sh
//  $ pulumi import alicloud:oss/bucket:Bucket bucket bucket-12345678
// ```
type Bucket struct {
	pulumi.CustomResourceState

	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
	Acl    pulumi.StringPtrOutput `pulumi:"acl"`
	Bucket pulumi.StringPtrOutput `pulumi:"bucket"`
	// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
	CorsRules BucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// The creation date of the bucket.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The extranet access endpoint of the bucket.
	ExtranetEndpoint pulumi.StringOutput `pulumi:"extranetEndpoint"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// The intranet access endpoint of the bucket.
	IntranetEndpoint pulumi.StringOutput `pulumi:"intranetEndpoint"`
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
	LifecycleRules BucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// The location of the bucket.
	Location pulumi.StringOutput `pulumi:"location"`
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
	Logging BucketLoggingPtrOutput `pulumi:"logging"`
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable pulumi.BoolPtrOutput `pulumi:"loggingIsenable"`
	// The bucket owner.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType pulumi.StringPtrOutput `pulumi:"redundancyType"`
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
	RefererConfig BucketRefererConfigPtrOutput `pulumi:"refererConfig"`
	// A configuration of server-side encryption (documented below).
	ServerSideEncryptionRule BucketServerSideEncryptionRulePtrOutput `pulumi:"serverSideEncryptionRule"`
	// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`.
	StorageClass pulumi.StringPtrOutput `pulumi:"storageClass"`
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// A state of versioning (documented below).
	Versioning BucketVersioningPtrOutput `pulumi:"versioning"`
	// A website object(documented below).
	Website BucketWebsitePtrOutput `pulumi:"website"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		args = &BucketArgs{}
	}

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
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
	Acl    *string `pulumi:"acl"`
	Bucket *string `pulumi:"bucket"`
	// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// The creation date of the bucket.
	CreationDate *string `pulumi:"creationDate"`
	// The extranet access endpoint of the bucket.
	ExtranetEndpoint *string `pulumi:"extranetEndpoint"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The intranet access endpoint of the bucket.
	IntranetEndpoint *string `pulumi:"intranetEndpoint"`
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The location of the bucket.
	Location *string `pulumi:"location"`
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
	Logging *BucketLogging `pulumi:"logging"`
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable *bool `pulumi:"loggingIsenable"`
	// The bucket owner.
	Owner *string `pulumi:"owner"`
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
	Policy *string `pulumi:"policy"`
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType *string `pulumi:"redundancyType"`
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
	RefererConfig *BucketRefererConfig `pulumi:"refererConfig"`
	// A configuration of server-side encryption (documented below).
	ServerSideEncryptionRule *BucketServerSideEncryptionRule `pulumi:"serverSideEncryptionRule"`
	// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`.
	StorageClass *string `pulumi:"storageClass"`
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags map[string]interface{} `pulumi:"tags"`
	// A state of versioning (documented below).
	Versioning *BucketVersioning `pulumi:"versioning"`
	// A website object(documented below).
	Website *BucketWebsite `pulumi:"website"`
}

type BucketState struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
	Acl    pulumi.StringPtrInput
	Bucket pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
	CorsRules BucketCorsRuleArrayInput
	// The creation date of the bucket.
	CreationDate pulumi.StringPtrInput
	// The extranet access endpoint of the bucket.
	ExtranetEndpoint pulumi.StringPtrInput
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy pulumi.BoolPtrInput
	// The intranet access endpoint of the bucket.
	IntranetEndpoint pulumi.StringPtrInput
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
	LifecycleRules BucketLifecycleRuleArrayInput
	// The location of the bucket.
	Location pulumi.StringPtrInput
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
	Logging BucketLoggingPtrInput
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable pulumi.BoolPtrInput
	// The bucket owner.
	Owner pulumi.StringPtrInput
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
	Policy pulumi.StringPtrInput
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType pulumi.StringPtrInput
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
	RefererConfig BucketRefererConfigPtrInput
	// A configuration of server-side encryption (documented below).
	ServerSideEncryptionRule BucketServerSideEncryptionRulePtrInput
	// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`.
	StorageClass pulumi.StringPtrInput
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags pulumi.MapInput
	// A state of versioning (documented below).
	Versioning BucketVersioningPtrInput
	// A website object(documented below).
	Website BucketWebsitePtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
	Acl    *string `pulumi:"acl"`
	Bucket *string `pulumi:"bucket"`
	// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
	Logging *BucketLogging `pulumi:"logging"`
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable *bool `pulumi:"loggingIsenable"`
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
	Policy *string `pulumi:"policy"`
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType *string `pulumi:"redundancyType"`
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
	RefererConfig *BucketRefererConfig `pulumi:"refererConfig"`
	// A configuration of server-side encryption (documented below).
	ServerSideEncryptionRule *BucketServerSideEncryptionRule `pulumi:"serverSideEncryptionRule"`
	// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`.
	StorageClass *string `pulumi:"storageClass"`
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags map[string]interface{} `pulumi:"tags"`
	// A state of versioning (documented below).
	Versioning *BucketVersioning `pulumi:"versioning"`
	// A website object(documented below).
	Website *BucketWebsite `pulumi:"website"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
	Acl    pulumi.StringPtrInput
	Bucket pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
	CorsRules BucketCorsRuleArrayInput
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
	ForceDestroy pulumi.BoolPtrInput
	// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
	LifecycleRules BucketLifecycleRuleArrayInput
	// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
	Logging BucketLoggingPtrInput
	// The flag of using logging enable container. Defaults true.
	//
	// Deprecated: Deprecated from 1.37.0. When `logging` is set, the bucket logging will be able.
	LoggingIsenable pulumi.BoolPtrInput
	// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
	Policy pulumi.StringPtrInput
	// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
	RedundancyType pulumi.StringPtrInput
	// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
	RefererConfig BucketRefererConfigPtrInput
	// A configuration of server-side encryption (documented below).
	ServerSideEncryptionRule BucketServerSideEncryptionRulePtrInput
	// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`.
	StorageClass pulumi.StringPtrInput
	// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
	Tags pulumi.MapInput
	// A state of versioning (documented below).
	Versioning BucketVersioningPtrInput
	// A website object(documented below).
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
	return reflect.TypeOf((*Bucket)(nil))
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

func (i *Bucket) ToBucketPtrOutput() BucketPtrOutput {
	return i.ToBucketPtrOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketPtrOutputWithContext(ctx context.Context) BucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPtrOutput)
}

type BucketPtrInput interface {
	pulumi.Input

	ToBucketPtrOutput() BucketPtrOutput
	ToBucketPtrOutputWithContext(ctx context.Context) BucketPtrOutput
}

type bucketPtrType BucketArgs

func (*bucketPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil))
}

func (i *bucketPtrType) ToBucketPtrOutput() BucketPtrOutput {
	return i.ToBucketPtrOutputWithContext(context.Background())
}

func (i *bucketPtrType) ToBucketPtrOutputWithContext(ctx context.Context) BucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPtrOutput)
}

// BucketArrayInput is an input type that accepts BucketArray and BucketArrayOutput values.
// You can construct a concrete instance of `BucketArrayInput` via:
//
//          BucketArray{ BucketArgs{...} }
type BucketArrayInput interface {
	pulumi.Input

	ToBucketArrayOutput() BucketArrayOutput
	ToBucketArrayOutputWithContext(context.Context) BucketArrayOutput
}

type BucketArray []BucketInput

func (BucketArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Bucket)(nil))
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
//          BucketMap{ "key": BucketArgs{...} }
type BucketMapInput interface {
	pulumi.Input

	ToBucketMapOutput() BucketMapOutput
	ToBucketMapOutputWithContext(context.Context) BucketMapOutput
}

type BucketMap map[string]BucketInput

func (BucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Bucket)(nil))
}

func (i BucketMap) ToBucketMapOutput() BucketMapOutput {
	return i.ToBucketMapOutputWithContext(context.Background())
}

func (i BucketMap) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMapOutput)
}

type BucketOutput struct {
	*pulumi.OutputState
}

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Bucket)(nil))
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

func (o BucketOutput) ToBucketPtrOutput() BucketPtrOutput {
	return o.ToBucketPtrOutputWithContext(context.Background())
}

func (o BucketOutput) ToBucketPtrOutputWithContext(ctx context.Context) BucketPtrOutput {
	return o.ApplyT(func(v Bucket) *Bucket {
		return &v
	}).(BucketPtrOutput)
}

type BucketPtrOutput struct {
	*pulumi.OutputState
}

func (BucketPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil))
}

func (o BucketPtrOutput) ToBucketPtrOutput() BucketPtrOutput {
	return o
}

func (o BucketPtrOutput) ToBucketPtrOutputWithContext(ctx context.Context) BucketPtrOutput {
	return o
}

type BucketArrayOutput struct{ *pulumi.OutputState }

func (BucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Bucket)(nil))
}

func (o BucketArrayOutput) ToBucketArrayOutput() BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) Index(i pulumi.IntInput) BucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Bucket {
		return vs[0].([]Bucket)[vs[1].(int)]
	}).(BucketOutput)
}

type BucketMapOutput struct{ *pulumi.OutputState }

func (BucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Bucket)(nil))
}

func (o BucketMapOutput) ToBucketMapOutput() BucketMapOutput {
	return o
}

func (o BucketMapOutput) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return o
}

func (o BucketMapOutput) MapIndex(k pulumi.StringInput) BucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Bucket {
		return vs[0].(map[string]Bucket)[vs[1].(string)]
	}).(BucketOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketOutput{})
	pulumi.RegisterOutputType(BucketPtrOutput{})
	pulumi.RegisterOutputType(BucketArrayOutput{})
	pulumi.RegisterOutputType(BucketMapOutput{})
}
