// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Log service ingestion, this service provides the function of importing logs of various data sources(OSS, MaxCompute) into logstore.
// [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
//
// > **NOTE:** Available in 1.161.0+
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewRandomInteger(ctx, "default", &random.RandomIntegerArgs{
//				Max: pulumi.Int(99999),
//				Min: pulumi.Int(10000),
//			})
//			if err != nil {
//				return err
//			}
//			exampleProject, err := log.NewProject(ctx, "exampleProject", &log.ProjectArgs{
//				Description: pulumi.String("terraform-example"),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleStore, err := log.NewStore(ctx, "exampleStore", &log.StoreArgs{
//				Project:            exampleProject.Name,
//				RetentionPeriod:    pulumi.Int(3650),
//				ShardCount:         pulumi.Int(3),
//				AutoSplit:          pulumi.Bool(true),
//				MaxSplitShardCount: pulumi.Int(60),
//				AppendMeta:         pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = log.NewIngestion(ctx, "exampleIngestion", &log.IngestionArgs{
//				Project:        exampleProject.Name,
//				Logstore:       exampleStore.Name,
//				IngestionName:  pulumi.String("terraform-example"),
//				DisplayName:    pulumi.String("terraform-example"),
//				Description:    pulumi.String("terraform-example"),
//				Interval:       pulumi.String("30m"),
//				RunImmediately: pulumi.Bool(true),
//				TimeZone:       pulumi.String("+0800"),
//				Source: pulumi.String(`        {
//	          "bucket": "bucket_name",
//	          "compressionCodec": "none",
//	          "encoding": "UTF-8",
//	          "endpoint": "oss-cn-hangzhou-internal.aliyuncs.com",
//	          "format": {
//	            "escapeChar": "\\",
//	            "fieldDelimiter": ",",
//	            "fieldNames": [],
//	            "firstRowAsHeader": true,
//	            "maxLines": 1,
//	            "quoteChar": "\"",
//	            "skipLeadingRows": 0,
//	            "timeField": "",
//	            "type": "DelimitedText"
//	          },
//	          "pattern": "",
//	          "prefix": "test-prefix/",
//	          "restoreObjectEnabled": false,
//	          "roleARN": "acs:ram::1049446484210612:role/aliyunlogimportossrole",
//	          "type": "AliyunOSS"
//	        }
//
// `),
//
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
// Log ingestion can be imported using the id or name, e.g.
//
// ```sh
// $ pulumi import alicloud:log/ingestion:Ingestion example tf-log-project:tf-log-logstore:ingestion_name
// ```
type Ingestion struct {
	pulumi.CustomResourceState

	// Ingestion job description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name displayed on the web page.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
	IngestionName pulumi.StringOutput `pulumi:"ingestionName"`
	// Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
	Interval pulumi.StringOutput `pulumi:"interval"`
	// The name of the target logstore.
	Logstore pulumi.StringOutput `pulumi:"logstore"`
	// The name of the log project. It is the only in one Alicloud account.
	Project pulumi.StringOutput `pulumi:"project"`
	// Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
	RunImmediately pulumi.BoolOutput `pulumi:"runImmediately"`
	// Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
	Source pulumi.StringOutput `pulumi:"source"`
	// Which time zone is the log time imported in, e.g. `+0800`.
	TimeZone pulumi.StringPtrOutput `pulumi:"timeZone"`
}

// NewIngestion registers a new resource with the given unique name, arguments, and options.
func NewIngestion(ctx *pulumi.Context,
	name string, args *IngestionArgs, opts ...pulumi.ResourceOption) (*Ingestion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.IngestionName == nil {
		return nil, errors.New("invalid value for required argument 'IngestionName'")
	}
	if args.Interval == nil {
		return nil, errors.New("invalid value for required argument 'Interval'")
	}
	if args.Logstore == nil {
		return nil, errors.New("invalid value for required argument 'Logstore'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.RunImmediately == nil {
		return nil, errors.New("invalid value for required argument 'RunImmediately'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ingestion
	err := ctx.RegisterResource("alicloud:log/ingestion:Ingestion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIngestion gets an existing Ingestion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIngestion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IngestionState, opts ...pulumi.ResourceOption) (*Ingestion, error) {
	var resource Ingestion
	err := ctx.ReadResource("alicloud:log/ingestion:Ingestion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ingestion resources.
type ingestionState struct {
	// Ingestion job description.
	Description *string `pulumi:"description"`
	// The name displayed on the web page.
	DisplayName *string `pulumi:"displayName"`
	// Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
	IngestionName *string `pulumi:"ingestionName"`
	// Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
	Interval *string `pulumi:"interval"`
	// The name of the target logstore.
	Logstore *string `pulumi:"logstore"`
	// The name of the log project. It is the only in one Alicloud account.
	Project *string `pulumi:"project"`
	// Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
	RunImmediately *bool `pulumi:"runImmediately"`
	// Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
	Source *string `pulumi:"source"`
	// Which time zone is the log time imported in, e.g. `+0800`.
	TimeZone *string `pulumi:"timeZone"`
}

type IngestionState struct {
	// Ingestion job description.
	Description pulumi.StringPtrInput
	// The name displayed on the web page.
	DisplayName pulumi.StringPtrInput
	// Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
	IngestionName pulumi.StringPtrInput
	// Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
	Interval pulumi.StringPtrInput
	// The name of the target logstore.
	Logstore pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account.
	Project pulumi.StringPtrInput
	// Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
	RunImmediately pulumi.BoolPtrInput
	// Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
	Source pulumi.StringPtrInput
	// Which time zone is the log time imported in, e.g. `+0800`.
	TimeZone pulumi.StringPtrInput
}

func (IngestionState) ElementType() reflect.Type {
	return reflect.TypeOf((*ingestionState)(nil)).Elem()
}

type ingestionArgs struct {
	// Ingestion job description.
	Description *string `pulumi:"description"`
	// The name displayed on the web page.
	DisplayName string `pulumi:"displayName"`
	// Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
	IngestionName string `pulumi:"ingestionName"`
	// Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
	Interval string `pulumi:"interval"`
	// The name of the target logstore.
	Logstore string `pulumi:"logstore"`
	// The name of the log project. It is the only in one Alicloud account.
	Project string `pulumi:"project"`
	// Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
	RunImmediately bool `pulumi:"runImmediately"`
	// Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
	Source string `pulumi:"source"`
	// Which time zone is the log time imported in, e.g. `+0800`.
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a Ingestion resource.
type IngestionArgs struct {
	// Ingestion job description.
	Description pulumi.StringPtrInput
	// The name displayed on the web page.
	DisplayName pulumi.StringInput
	// Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
	IngestionName pulumi.StringInput
	// Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
	Interval pulumi.StringInput
	// The name of the target logstore.
	Logstore pulumi.StringInput
	// The name of the log project. It is the only in one Alicloud account.
	Project pulumi.StringInput
	// Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
	RunImmediately pulumi.BoolInput
	// Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
	Source pulumi.StringInput
	// Which time zone is the log time imported in, e.g. `+0800`.
	TimeZone pulumi.StringPtrInput
}

func (IngestionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ingestionArgs)(nil)).Elem()
}

type IngestionInput interface {
	pulumi.Input

	ToIngestionOutput() IngestionOutput
	ToIngestionOutputWithContext(ctx context.Context) IngestionOutput
}

func (*Ingestion) ElementType() reflect.Type {
	return reflect.TypeOf((**Ingestion)(nil)).Elem()
}

func (i *Ingestion) ToIngestionOutput() IngestionOutput {
	return i.ToIngestionOutputWithContext(context.Background())
}

func (i *Ingestion) ToIngestionOutputWithContext(ctx context.Context) IngestionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionOutput)
}

// IngestionArrayInput is an input type that accepts IngestionArray and IngestionArrayOutput values.
// You can construct a concrete instance of `IngestionArrayInput` via:
//
//	IngestionArray{ IngestionArgs{...} }
type IngestionArrayInput interface {
	pulumi.Input

	ToIngestionArrayOutput() IngestionArrayOutput
	ToIngestionArrayOutputWithContext(context.Context) IngestionArrayOutput
}

type IngestionArray []IngestionInput

func (IngestionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ingestion)(nil)).Elem()
}

func (i IngestionArray) ToIngestionArrayOutput() IngestionArrayOutput {
	return i.ToIngestionArrayOutputWithContext(context.Background())
}

func (i IngestionArray) ToIngestionArrayOutputWithContext(ctx context.Context) IngestionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionArrayOutput)
}

// IngestionMapInput is an input type that accepts IngestionMap and IngestionMapOutput values.
// You can construct a concrete instance of `IngestionMapInput` via:
//
//	IngestionMap{ "key": IngestionArgs{...} }
type IngestionMapInput interface {
	pulumi.Input

	ToIngestionMapOutput() IngestionMapOutput
	ToIngestionMapOutputWithContext(context.Context) IngestionMapOutput
}

type IngestionMap map[string]IngestionInput

func (IngestionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ingestion)(nil)).Elem()
}

func (i IngestionMap) ToIngestionMapOutput() IngestionMapOutput {
	return i.ToIngestionMapOutputWithContext(context.Background())
}

func (i IngestionMap) ToIngestionMapOutputWithContext(ctx context.Context) IngestionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionMapOutput)
}

type IngestionOutput struct{ *pulumi.OutputState }

func (IngestionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ingestion)(nil)).Elem()
}

func (o IngestionOutput) ToIngestionOutput() IngestionOutput {
	return o
}

func (o IngestionOutput) ToIngestionOutputWithContext(ctx context.Context) IngestionOutput {
	return o
}

// Ingestion job description.
func (o IngestionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name displayed on the web page.
func (o IngestionOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
func (o IngestionOutput) IngestionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.IngestionName }).(pulumi.StringOutput)
}

// Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
func (o IngestionOutput) Interval() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.Interval }).(pulumi.StringOutput)
}

// The name of the target logstore.
func (o IngestionOutput) Logstore() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.Logstore }).(pulumi.StringOutput)
}

// The name of the log project. It is the only in one Alicloud account.
func (o IngestionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
func (o IngestionOutput) RunImmediately() pulumi.BoolOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.BoolOutput { return v.RunImmediately }).(pulumi.BoolOutput)
}

// Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
func (o IngestionOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Which time zone is the log time imported in, e.g. `+0800`.
func (o IngestionOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ingestion) pulumi.StringPtrOutput { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type IngestionArrayOutput struct{ *pulumi.OutputState }

func (IngestionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ingestion)(nil)).Elem()
}

func (o IngestionArrayOutput) ToIngestionArrayOutput() IngestionArrayOutput {
	return o
}

func (o IngestionArrayOutput) ToIngestionArrayOutputWithContext(ctx context.Context) IngestionArrayOutput {
	return o
}

func (o IngestionArrayOutput) Index(i pulumi.IntInput) IngestionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ingestion {
		return vs[0].([]*Ingestion)[vs[1].(int)]
	}).(IngestionOutput)
}

type IngestionMapOutput struct{ *pulumi.OutputState }

func (IngestionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ingestion)(nil)).Elem()
}

func (o IngestionMapOutput) ToIngestionMapOutput() IngestionMapOutput {
	return o
}

func (o IngestionMapOutput) ToIngestionMapOutputWithContext(ctx context.Context) IngestionMapOutput {
	return o
}

func (o IngestionMapOutput) MapIndex(k pulumi.StringInput) IngestionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ingestion {
		return vs[0].(map[string]*Ingestion)[vs[1].(string)]
	}).(IngestionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IngestionInput)(nil)).Elem(), &Ingestion{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngestionArrayInput)(nil)).Elem(), IngestionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngestionMapInput)(nil)).Elem(), IngestionMap{})
	pulumi.RegisterOutputType(IngestionOutput{})
	pulumi.RegisterOutputType(IngestionArrayOutput{})
	pulumi.RegisterOutputType(IngestionMapOutput{})
}
