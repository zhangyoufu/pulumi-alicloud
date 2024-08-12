// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Threat Detection Oss Scan Config resource. Oss detection configuration.
//
// For information about Threat Detection Oss Scan Config and how to use it, see [What is Oss Scan Config](https://www.alibabacloud.com/help/zh/security-center/developer-reference/api-sas-2018-12-03-createossscanconfig/).
//
// > **NOTE:** Available since v1.214.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/threatdetection"
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
//			bucketRandom := _default.Result
//			default8j4t1R, err := oss.NewBucket(ctx, "default8j4t1R", &oss.BucketArgs{
//				Bucket:       pulumi.Sprintf("%v-1-%v", name, bucketRandom),
//				StorageClass: pulumi.String("Standard"),
//			})
//			if err != nil {
//				return err
//			}
//			default9HMqfT, err := oss.NewBucket(ctx, "default9HMqfT", &oss.BucketArgs{
//				Bucket:       pulumi.Sprintf("%v-2-%v", name, bucketRandom),
//				StorageClass: pulumi.String("Standard"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultxBXqFQ, err := oss.NewBucket(ctx, "defaultxBXqFQ", &oss.BucketArgs{
//				Bucket:       pulumi.Sprintf("%v-3-%v", name, bucketRandom),
//				StorageClass: pulumi.String("Standard"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucket(ctx, "defaulthZvCmR", &oss.BucketArgs{
//				Bucket:       pulumi.Sprintf("%v-4-%v", name, bucketRandom),
//				StorageClass: pulumi.String("Standard"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = threatdetection.NewOssScanConfig(ctx, "default", &threatdetection.OssScanConfigArgs{
//				KeySuffixLists: pulumi.StringArray{
//					pulumi.String(".jsp"),
//					pulumi.String(".php"),
//					pulumi.String(".k"),
//				},
//				ScanDayLists: pulumi.IntArray{
//					pulumi.Int(2),
//					pulumi.Int(5),
//					pulumi.Int(4),
//					pulumi.Int(3),
//				},
//				OssScanConfigName: pulumi.String(name),
//				EndTime:           pulumi.String("00:00:02"),
//				StartTime:         pulumi.String("00:00:01"),
//				Enable:            pulumi.Int(1),
//				AllKeyPrefix:      pulumi.Bool(false),
//				BucketNameLists: pulumi.StringArray{
//					default8j4t1R.Bucket,
//					default9HMqfT.Bucket,
//					defaultxBXqFQ.Bucket,
//				},
//				KeyPrefixLists: pulumi.StringArray{
//					pulumi.String("/root"),
//					pulumi.String("/usr"),
//					pulumi.String("/123"),
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
// Threat Detection Oss Scan Config can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:threatdetection/ossScanConfig:OssScanConfig example <id>
// ```
type OssScanConfig struct {
	pulumi.CustomResourceState

	// Match all prefixes.
	AllKeyPrefix pulumi.BoolOutput `pulumi:"allKeyPrefix"`
	// Bucket List.
	BucketNameLists pulumi.StringArrayOutput `pulumi:"bucketNameLists"`
	// Enable configuration.
	Enable pulumi.IntOutput `pulumi:"enable"`
	// End time, hours, minutes and seconds.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// File prefix list.
	KeyPrefixLists pulumi.StringArrayOutput `pulumi:"keyPrefixLists"`
	// File Suffix List.
	KeySuffixLists pulumi.StringArrayOutput `pulumi:"keySuffixLists"`
	// Configuration Name.
	OssScanConfigName pulumi.StringPtrOutput `pulumi:"ossScanConfigName"`
	// Scan cycle.
	ScanDayLists pulumi.IntArrayOutput `pulumi:"scanDayLists"`
	// Start time, hours, minutes and seconds.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
}

// NewOssScanConfig registers a new resource with the given unique name, arguments, and options.
func NewOssScanConfig(ctx *pulumi.Context,
	name string, args *OssScanConfigArgs, opts ...pulumi.ResourceOption) (*OssScanConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketNameLists == nil {
		return nil, errors.New("invalid value for required argument 'BucketNameLists'")
	}
	if args.Enable == nil {
		return nil, errors.New("invalid value for required argument 'Enable'")
	}
	if args.EndTime == nil {
		return nil, errors.New("invalid value for required argument 'EndTime'")
	}
	if args.KeySuffixLists == nil {
		return nil, errors.New("invalid value for required argument 'KeySuffixLists'")
	}
	if args.ScanDayLists == nil {
		return nil, errors.New("invalid value for required argument 'ScanDayLists'")
	}
	if args.StartTime == nil {
		return nil, errors.New("invalid value for required argument 'StartTime'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OssScanConfig
	err := ctx.RegisterResource("alicloud:threatdetection/ossScanConfig:OssScanConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOssScanConfig gets an existing OssScanConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOssScanConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OssScanConfigState, opts ...pulumi.ResourceOption) (*OssScanConfig, error) {
	var resource OssScanConfig
	err := ctx.ReadResource("alicloud:threatdetection/ossScanConfig:OssScanConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OssScanConfig resources.
type ossScanConfigState struct {
	// Match all prefixes.
	AllKeyPrefix *bool `pulumi:"allKeyPrefix"`
	// Bucket List.
	BucketNameLists []string `pulumi:"bucketNameLists"`
	// Enable configuration.
	Enable *int `pulumi:"enable"`
	// End time, hours, minutes and seconds.
	EndTime *string `pulumi:"endTime"`
	// File prefix list.
	KeyPrefixLists []string `pulumi:"keyPrefixLists"`
	// File Suffix List.
	KeySuffixLists []string `pulumi:"keySuffixLists"`
	// Configuration Name.
	OssScanConfigName *string `pulumi:"ossScanConfigName"`
	// Scan cycle.
	ScanDayLists []int `pulumi:"scanDayLists"`
	// Start time, hours, minutes and seconds.
	StartTime *string `pulumi:"startTime"`
}

type OssScanConfigState struct {
	// Match all prefixes.
	AllKeyPrefix pulumi.BoolPtrInput
	// Bucket List.
	BucketNameLists pulumi.StringArrayInput
	// Enable configuration.
	Enable pulumi.IntPtrInput
	// End time, hours, minutes and seconds.
	EndTime pulumi.StringPtrInput
	// File prefix list.
	KeyPrefixLists pulumi.StringArrayInput
	// File Suffix List.
	KeySuffixLists pulumi.StringArrayInput
	// Configuration Name.
	OssScanConfigName pulumi.StringPtrInput
	// Scan cycle.
	ScanDayLists pulumi.IntArrayInput
	// Start time, hours, minutes and seconds.
	StartTime pulumi.StringPtrInput
}

func (OssScanConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*ossScanConfigState)(nil)).Elem()
}

type ossScanConfigArgs struct {
	// Match all prefixes.
	AllKeyPrefix *bool `pulumi:"allKeyPrefix"`
	// Bucket List.
	BucketNameLists []string `pulumi:"bucketNameLists"`
	// Enable configuration.
	Enable int `pulumi:"enable"`
	// End time, hours, minutes and seconds.
	EndTime string `pulumi:"endTime"`
	// File prefix list.
	KeyPrefixLists []string `pulumi:"keyPrefixLists"`
	// File Suffix List.
	KeySuffixLists []string `pulumi:"keySuffixLists"`
	// Configuration Name.
	OssScanConfigName *string `pulumi:"ossScanConfigName"`
	// Scan cycle.
	ScanDayLists []int `pulumi:"scanDayLists"`
	// Start time, hours, minutes and seconds.
	StartTime string `pulumi:"startTime"`
}

// The set of arguments for constructing a OssScanConfig resource.
type OssScanConfigArgs struct {
	// Match all prefixes.
	AllKeyPrefix pulumi.BoolPtrInput
	// Bucket List.
	BucketNameLists pulumi.StringArrayInput
	// Enable configuration.
	Enable pulumi.IntInput
	// End time, hours, minutes and seconds.
	EndTime pulumi.StringInput
	// File prefix list.
	KeyPrefixLists pulumi.StringArrayInput
	// File Suffix List.
	KeySuffixLists pulumi.StringArrayInput
	// Configuration Name.
	OssScanConfigName pulumi.StringPtrInput
	// Scan cycle.
	ScanDayLists pulumi.IntArrayInput
	// Start time, hours, minutes and seconds.
	StartTime pulumi.StringInput
}

func (OssScanConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ossScanConfigArgs)(nil)).Elem()
}

type OssScanConfigInput interface {
	pulumi.Input

	ToOssScanConfigOutput() OssScanConfigOutput
	ToOssScanConfigOutputWithContext(ctx context.Context) OssScanConfigOutput
}

func (*OssScanConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**OssScanConfig)(nil)).Elem()
}

func (i *OssScanConfig) ToOssScanConfigOutput() OssScanConfigOutput {
	return i.ToOssScanConfigOutputWithContext(context.Background())
}

func (i *OssScanConfig) ToOssScanConfigOutputWithContext(ctx context.Context) OssScanConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OssScanConfigOutput)
}

// OssScanConfigArrayInput is an input type that accepts OssScanConfigArray and OssScanConfigArrayOutput values.
// You can construct a concrete instance of `OssScanConfigArrayInput` via:
//
//	OssScanConfigArray{ OssScanConfigArgs{...} }
type OssScanConfigArrayInput interface {
	pulumi.Input

	ToOssScanConfigArrayOutput() OssScanConfigArrayOutput
	ToOssScanConfigArrayOutputWithContext(context.Context) OssScanConfigArrayOutput
}

type OssScanConfigArray []OssScanConfigInput

func (OssScanConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OssScanConfig)(nil)).Elem()
}

func (i OssScanConfigArray) ToOssScanConfigArrayOutput() OssScanConfigArrayOutput {
	return i.ToOssScanConfigArrayOutputWithContext(context.Background())
}

func (i OssScanConfigArray) ToOssScanConfigArrayOutputWithContext(ctx context.Context) OssScanConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OssScanConfigArrayOutput)
}

// OssScanConfigMapInput is an input type that accepts OssScanConfigMap and OssScanConfigMapOutput values.
// You can construct a concrete instance of `OssScanConfigMapInput` via:
//
//	OssScanConfigMap{ "key": OssScanConfigArgs{...} }
type OssScanConfigMapInput interface {
	pulumi.Input

	ToOssScanConfigMapOutput() OssScanConfigMapOutput
	ToOssScanConfigMapOutputWithContext(context.Context) OssScanConfigMapOutput
}

type OssScanConfigMap map[string]OssScanConfigInput

func (OssScanConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OssScanConfig)(nil)).Elem()
}

func (i OssScanConfigMap) ToOssScanConfigMapOutput() OssScanConfigMapOutput {
	return i.ToOssScanConfigMapOutputWithContext(context.Background())
}

func (i OssScanConfigMap) ToOssScanConfigMapOutputWithContext(ctx context.Context) OssScanConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OssScanConfigMapOutput)
}

type OssScanConfigOutput struct{ *pulumi.OutputState }

func (OssScanConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OssScanConfig)(nil)).Elem()
}

func (o OssScanConfigOutput) ToOssScanConfigOutput() OssScanConfigOutput {
	return o
}

func (o OssScanConfigOutput) ToOssScanConfigOutputWithContext(ctx context.Context) OssScanConfigOutput {
	return o
}

// Match all prefixes.
func (o OssScanConfigOutput) AllKeyPrefix() pulumi.BoolOutput {
	return o.ApplyT(func(v *OssScanConfig) pulumi.BoolOutput { return v.AllKeyPrefix }).(pulumi.BoolOutput)
}

// Bucket List.
func (o OssScanConfigOutput) BucketNameLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OssScanConfig) pulumi.StringArrayOutput { return v.BucketNameLists }).(pulumi.StringArrayOutput)
}

// Enable configuration.
func (o OssScanConfigOutput) Enable() pulumi.IntOutput {
	return o.ApplyT(func(v *OssScanConfig) pulumi.IntOutput { return v.Enable }).(pulumi.IntOutput)
}

// End time, hours, minutes and seconds.
func (o OssScanConfigOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OssScanConfig) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// File prefix list.
func (o OssScanConfigOutput) KeyPrefixLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OssScanConfig) pulumi.StringArrayOutput { return v.KeyPrefixLists }).(pulumi.StringArrayOutput)
}

// File Suffix List.
func (o OssScanConfigOutput) KeySuffixLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OssScanConfig) pulumi.StringArrayOutput { return v.KeySuffixLists }).(pulumi.StringArrayOutput)
}

// Configuration Name.
func (o OssScanConfigOutput) OssScanConfigName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssScanConfig) pulumi.StringPtrOutput { return v.OssScanConfigName }).(pulumi.StringPtrOutput)
}

// Scan cycle.
func (o OssScanConfigOutput) ScanDayLists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *OssScanConfig) pulumi.IntArrayOutput { return v.ScanDayLists }).(pulumi.IntArrayOutput)
}

// Start time, hours, minutes and seconds.
func (o OssScanConfigOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OssScanConfig) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

type OssScanConfigArrayOutput struct{ *pulumi.OutputState }

func (OssScanConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OssScanConfig)(nil)).Elem()
}

func (o OssScanConfigArrayOutput) ToOssScanConfigArrayOutput() OssScanConfigArrayOutput {
	return o
}

func (o OssScanConfigArrayOutput) ToOssScanConfigArrayOutputWithContext(ctx context.Context) OssScanConfigArrayOutput {
	return o
}

func (o OssScanConfigArrayOutput) Index(i pulumi.IntInput) OssScanConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OssScanConfig {
		return vs[0].([]*OssScanConfig)[vs[1].(int)]
	}).(OssScanConfigOutput)
}

type OssScanConfigMapOutput struct{ *pulumi.OutputState }

func (OssScanConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OssScanConfig)(nil)).Elem()
}

func (o OssScanConfigMapOutput) ToOssScanConfigMapOutput() OssScanConfigMapOutput {
	return o
}

func (o OssScanConfigMapOutput) ToOssScanConfigMapOutputWithContext(ctx context.Context) OssScanConfigMapOutput {
	return o
}

func (o OssScanConfigMapOutput) MapIndex(k pulumi.StringInput) OssScanConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OssScanConfig {
		return vs[0].(map[string]*OssScanConfig)[vs[1].(string)]
	}).(OssScanConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OssScanConfigInput)(nil)).Elem(), &OssScanConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*OssScanConfigArrayInput)(nil)).Elem(), OssScanConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OssScanConfigMapInput)(nil)).Elem(), OssScanConfigMap{})
	pulumi.RegisterOutputType(OssScanConfigOutput{})
	pulumi.RegisterOutputType(OssScanConfigArrayOutput{})
	pulumi.RegisterOutputType(OssScanConfigMapOutput{})
}
