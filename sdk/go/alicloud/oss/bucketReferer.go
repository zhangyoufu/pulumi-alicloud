// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a OSS Bucket Referer resource. Bucket Referer configuration (Hotlink protection).
//
// For information about OSS Bucket Referer and how to use it, see [What is Bucket Referer](https://www.alibabacloud.com/help/en/oss/user-guide/hotlink-protection).
//
// > **NOTE:** Available since v1.220.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
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
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Min: 10000,
//				Max: 99999,
//			})
//			if err != nil {
//				return err
//			}
//			createBucket, err := oss.NewBucket(ctx, "CreateBucket", &oss.BucketArgs{
//				StorageClass: pulumi.String("Standard"),
//				Bucket:       pulumi.String(fmt.Sprintf("%v-%v", name, _default.Result)),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucketReferer(ctx, "default", &oss.BucketRefererArgs{
//				AllowEmptyReferer: pulumi.Bool(true),
//				RefererBlacklists: pulumi.StringArray{
//					pulumi.String("*.forbidden.com"),
//				},
//				Bucket:                   createBucket.Bucket,
//				TruncatePath:             pulumi.Bool(false),
//				AllowTruncateQueryString: pulumi.Bool(true),
//				RefererLists: pulumi.StringArray{
//					pulumi.String("*.aliyun.com"),
//					pulumi.String("*.example.com"),
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
// OSS Bucket Referer can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:oss/bucketReferer:BucketReferer example <id>
// ```
type BucketReferer struct {
	pulumi.CustomResourceState

	// Whether to allow empty Referer request headers.
	AllowEmptyReferer pulumi.BoolOutput `pulumi:"allowEmptyReferer"`
	// Specifies whether to truncate the query string in the URL when the Referer is matched. Valid values: true, false.
	AllowTruncateQueryString pulumi.BoolOutput `pulumi:"allowTruncateQueryString"`
	// Name of the Bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The container that holds the Referer blacklist.
	RefererBlacklists pulumi.StringArrayOutput `pulumi:"refererBlacklists"`
	// The container that holds the Referer whitelist.
	RefererLists pulumi.StringArrayOutput `pulumi:"refererLists"`
	// Specifies whether to truncate the path and parts that follow the path in the URL when the Referer is matched. Valid values: true, false. If TruncatePath is set to true, the value of AllowTruncateQueryString must be also true because the query string follows the path component. When the path is truncated, the query string is also truncated.
	TruncatePath pulumi.BoolPtrOutput `pulumi:"truncatePath"`
}

// NewBucketReferer registers a new resource with the given unique name, arguments, and options.
func NewBucketReferer(ctx *pulumi.Context,
	name string, args *BucketRefererArgs, opts ...pulumi.ResourceOption) (*BucketReferer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowEmptyReferer == nil {
		return nil, errors.New("invalid value for required argument 'AllowEmptyReferer'")
	}
	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketReferer
	err := ctx.RegisterResource("alicloud:oss/bucketReferer:BucketReferer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketReferer gets an existing BucketReferer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketReferer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketRefererState, opts ...pulumi.ResourceOption) (*BucketReferer, error) {
	var resource BucketReferer
	err := ctx.ReadResource("alicloud:oss/bucketReferer:BucketReferer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketReferer resources.
type bucketRefererState struct {
	// Whether to allow empty Referer request headers.
	AllowEmptyReferer *bool `pulumi:"allowEmptyReferer"`
	// Specifies whether to truncate the query string in the URL when the Referer is matched. Valid values: true, false.
	AllowTruncateQueryString *bool `pulumi:"allowTruncateQueryString"`
	// Name of the Bucket.
	Bucket *string `pulumi:"bucket"`
	// The container that holds the Referer blacklist.
	RefererBlacklists []string `pulumi:"refererBlacklists"`
	// The container that holds the Referer whitelist.
	RefererLists []string `pulumi:"refererLists"`
	// Specifies whether to truncate the path and parts that follow the path in the URL when the Referer is matched. Valid values: true, false. If TruncatePath is set to true, the value of AllowTruncateQueryString must be also true because the query string follows the path component. When the path is truncated, the query string is also truncated.
	TruncatePath *bool `pulumi:"truncatePath"`
}

type BucketRefererState struct {
	// Whether to allow empty Referer request headers.
	AllowEmptyReferer pulumi.BoolPtrInput
	// Specifies whether to truncate the query string in the URL when the Referer is matched. Valid values: true, false.
	AllowTruncateQueryString pulumi.BoolPtrInput
	// Name of the Bucket.
	Bucket pulumi.StringPtrInput
	// The container that holds the Referer blacklist.
	RefererBlacklists pulumi.StringArrayInput
	// The container that holds the Referer whitelist.
	RefererLists pulumi.StringArrayInput
	// Specifies whether to truncate the path and parts that follow the path in the URL when the Referer is matched. Valid values: true, false. If TruncatePath is set to true, the value of AllowTruncateQueryString must be also true because the query string follows the path component. When the path is truncated, the query string is also truncated.
	TruncatePath pulumi.BoolPtrInput
}

func (BucketRefererState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketRefererState)(nil)).Elem()
}

type bucketRefererArgs struct {
	// Whether to allow empty Referer request headers.
	AllowEmptyReferer bool `pulumi:"allowEmptyReferer"`
	// Specifies whether to truncate the query string in the URL when the Referer is matched. Valid values: true, false.
	AllowTruncateQueryString *bool `pulumi:"allowTruncateQueryString"`
	// Name of the Bucket.
	Bucket string `pulumi:"bucket"`
	// The container that holds the Referer blacklist.
	RefererBlacklists []string `pulumi:"refererBlacklists"`
	// The container that holds the Referer whitelist.
	RefererLists []string `pulumi:"refererLists"`
	// Specifies whether to truncate the path and parts that follow the path in the URL when the Referer is matched. Valid values: true, false. If TruncatePath is set to true, the value of AllowTruncateQueryString must be also true because the query string follows the path component. When the path is truncated, the query string is also truncated.
	TruncatePath *bool `pulumi:"truncatePath"`
}

// The set of arguments for constructing a BucketReferer resource.
type BucketRefererArgs struct {
	// Whether to allow empty Referer request headers.
	AllowEmptyReferer pulumi.BoolInput
	// Specifies whether to truncate the query string in the URL when the Referer is matched. Valid values: true, false.
	AllowTruncateQueryString pulumi.BoolPtrInput
	// Name of the Bucket.
	Bucket pulumi.StringInput
	// The container that holds the Referer blacklist.
	RefererBlacklists pulumi.StringArrayInput
	// The container that holds the Referer whitelist.
	RefererLists pulumi.StringArrayInput
	// Specifies whether to truncate the path and parts that follow the path in the URL when the Referer is matched. Valid values: true, false. If TruncatePath is set to true, the value of AllowTruncateQueryString must be also true because the query string follows the path component. When the path is truncated, the query string is also truncated.
	TruncatePath pulumi.BoolPtrInput
}

func (BucketRefererArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketRefererArgs)(nil)).Elem()
}

type BucketRefererInput interface {
	pulumi.Input

	ToBucketRefererOutput() BucketRefererOutput
	ToBucketRefererOutputWithContext(ctx context.Context) BucketRefererOutput
}

func (*BucketReferer) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketReferer)(nil)).Elem()
}

func (i *BucketReferer) ToBucketRefererOutput() BucketRefererOutput {
	return i.ToBucketRefererOutputWithContext(context.Background())
}

func (i *BucketReferer) ToBucketRefererOutputWithContext(ctx context.Context) BucketRefererOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketRefererOutput)
}

// BucketRefererArrayInput is an input type that accepts BucketRefererArray and BucketRefererArrayOutput values.
// You can construct a concrete instance of `BucketRefererArrayInput` via:
//
//	BucketRefererArray{ BucketRefererArgs{...} }
type BucketRefererArrayInput interface {
	pulumi.Input

	ToBucketRefererArrayOutput() BucketRefererArrayOutput
	ToBucketRefererArrayOutputWithContext(context.Context) BucketRefererArrayOutput
}

type BucketRefererArray []BucketRefererInput

func (BucketRefererArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketReferer)(nil)).Elem()
}

func (i BucketRefererArray) ToBucketRefererArrayOutput() BucketRefererArrayOutput {
	return i.ToBucketRefererArrayOutputWithContext(context.Background())
}

func (i BucketRefererArray) ToBucketRefererArrayOutputWithContext(ctx context.Context) BucketRefererArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketRefererArrayOutput)
}

// BucketRefererMapInput is an input type that accepts BucketRefererMap and BucketRefererMapOutput values.
// You can construct a concrete instance of `BucketRefererMapInput` via:
//
//	BucketRefererMap{ "key": BucketRefererArgs{...} }
type BucketRefererMapInput interface {
	pulumi.Input

	ToBucketRefererMapOutput() BucketRefererMapOutput
	ToBucketRefererMapOutputWithContext(context.Context) BucketRefererMapOutput
}

type BucketRefererMap map[string]BucketRefererInput

func (BucketRefererMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketReferer)(nil)).Elem()
}

func (i BucketRefererMap) ToBucketRefererMapOutput() BucketRefererMapOutput {
	return i.ToBucketRefererMapOutputWithContext(context.Background())
}

func (i BucketRefererMap) ToBucketRefererMapOutputWithContext(ctx context.Context) BucketRefererMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketRefererMapOutput)
}

type BucketRefererOutput struct{ *pulumi.OutputState }

func (BucketRefererOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketReferer)(nil)).Elem()
}

func (o BucketRefererOutput) ToBucketRefererOutput() BucketRefererOutput {
	return o
}

func (o BucketRefererOutput) ToBucketRefererOutputWithContext(ctx context.Context) BucketRefererOutput {
	return o
}

// Whether to allow empty Referer request headers.
func (o BucketRefererOutput) AllowEmptyReferer() pulumi.BoolOutput {
	return o.ApplyT(func(v *BucketReferer) pulumi.BoolOutput { return v.AllowEmptyReferer }).(pulumi.BoolOutput)
}

// Specifies whether to truncate the query string in the URL when the Referer is matched. Valid values: true, false.
func (o BucketRefererOutput) AllowTruncateQueryString() pulumi.BoolOutput {
	return o.ApplyT(func(v *BucketReferer) pulumi.BoolOutput { return v.AllowTruncateQueryString }).(pulumi.BoolOutput)
}

// Name of the Bucket.
func (o BucketRefererOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketReferer) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The container that holds the Referer blacklist.
func (o BucketRefererOutput) RefererBlacklists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BucketReferer) pulumi.StringArrayOutput { return v.RefererBlacklists }).(pulumi.StringArrayOutput)
}

// The container that holds the Referer whitelist.
func (o BucketRefererOutput) RefererLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BucketReferer) pulumi.StringArrayOutput { return v.RefererLists }).(pulumi.StringArrayOutput)
}

// Specifies whether to truncate the path and parts that follow the path in the URL when the Referer is matched. Valid values: true, false. If TruncatePath is set to true, the value of AllowTruncateQueryString must be also true because the query string follows the path component. When the path is truncated, the query string is also truncated.
func (o BucketRefererOutput) TruncatePath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketReferer) pulumi.BoolPtrOutput { return v.TruncatePath }).(pulumi.BoolPtrOutput)
}

type BucketRefererArrayOutput struct{ *pulumi.OutputState }

func (BucketRefererArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketReferer)(nil)).Elem()
}

func (o BucketRefererArrayOutput) ToBucketRefererArrayOutput() BucketRefererArrayOutput {
	return o
}

func (o BucketRefererArrayOutput) ToBucketRefererArrayOutputWithContext(ctx context.Context) BucketRefererArrayOutput {
	return o
}

func (o BucketRefererArrayOutput) Index(i pulumi.IntInput) BucketRefererOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketReferer {
		return vs[0].([]*BucketReferer)[vs[1].(int)]
	}).(BucketRefererOutput)
}

type BucketRefererMapOutput struct{ *pulumi.OutputState }

func (BucketRefererMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketReferer)(nil)).Elem()
}

func (o BucketRefererMapOutput) ToBucketRefererMapOutput() BucketRefererMapOutput {
	return o
}

func (o BucketRefererMapOutput) ToBucketRefererMapOutputWithContext(ctx context.Context) BucketRefererMapOutput {
	return o
}

func (o BucketRefererMapOutput) MapIndex(k pulumi.StringInput) BucketRefererOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketReferer {
		return vs[0].(map[string]*BucketReferer)[vs[1].(string)]
	}).(BucketRefererOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketRefererInput)(nil)).Elem(), &BucketReferer{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketRefererArrayInput)(nil)).Elem(), BucketRefererArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketRefererMapInput)(nil)).Elem(), BucketRefererMap{})
	pulumi.RegisterOutputType(BucketRefererOutput{})
	pulumi.RegisterOutputType(BucketRefererArrayOutput{})
	pulumi.RegisterOutputType(BucketRefererMapOutput{})
}
