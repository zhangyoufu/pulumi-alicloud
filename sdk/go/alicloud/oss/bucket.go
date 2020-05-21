// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create a oss bucket and set its attribution.
//
// > **NOTE:** The bucket namespace is shared by all users of the OSS system. Please set bucket name as unique as possible.
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
