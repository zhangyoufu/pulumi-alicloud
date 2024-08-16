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

// Log service data delivery management, this service provides the function of delivering data in logstore to oss product storage.
// [Refer to details](https://www.alibabacloud.com/help/en/log-service/latest/ship-logs-to-oss-new-version).
//
// > **NOTE:** Available in 1.187.0+
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
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
//			example, err := log.NewProject(ctx, "example", &log.ProjectArgs{
//				Name:        pulumi.Sprintf("terraform-example-%v", _default.Result),
//				Description: pulumi.String("terraform-example"),
//				Tags: pulumi.StringMap{
//					"Created": pulumi.String("TF"),
//					"For":     pulumi.String("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleStore, err := log.NewStore(ctx, "example", &log.StoreArgs{
//				Project:            example.Name,
//				Name:               pulumi.String("example-store"),
//				RetentionPeriod:    pulumi.Int(3650),
//				ShardCount:         pulumi.Int(3),
//				AutoSplit:          pulumi.Bool(true),
//				MaxSplitShardCount: pulumi.Int(60),
//				AppendMeta:         pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = log.NewOssExport(ctx, "example", &log.OssExportArgs{
//				ProjectName:    example.Name,
//				LogstoreName:   exampleStore.Name,
//				ExportName:     pulumi.String("terraform-example"),
//				DisplayName:    pulumi.String("terraform-example"),
//				Bucket:         pulumi.String("example-bucket"),
//				Prefix:         pulumi.String("root"),
//				Suffix:         pulumi.String(""),
//				BufferInterval: pulumi.Int(300),
//				BufferSize:     pulumi.Int(250),
//				CompressType:   pulumi.String("none"),
//				PathFormat:     pulumi.String("%Y/%m/%d/%H/%M"),
//				ContentType:    pulumi.String("json"),
//				JsonEnableTag:  pulumi.Bool(true),
//				RoleArn:        pulumi.String("role_arn_for_oss_write"),
//				LogReadRoleArn: pulumi.String("role_arn_for_sls_read"),
//				TimeZone:       pulumi.String("+0800"),
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
// Log oss export can be imported using the id or name, e.g.
//
// ```sh
// $ pulumi import alicloud:log/ossExport:OssExport example tf-log-project:tf-log-logstore:tf-log-export
// ```
type OssExport struct {
	pulumi.CustomResourceState

	// The name of the oss bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// How often is it delivered every interval.
	BufferInterval pulumi.IntOutput `pulumi:"bufferInterval"`
	// Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
	BufferSize pulumi.IntOutput `pulumi:"bufferSize"`
	// OSS data storage compression method, support: `none`, `snappy`, `zstd`, `gzip`. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
	CompressType pulumi.StringOutput `pulumi:"compressType"`
	// Configure columns when `contentType` is `parquet` or `orc`.
	ConfigColumns OssExportConfigColumnArrayOutput `pulumi:"configColumns"`
	// Storage format, only supports three types: `json`, `parquet`, `orc`, `csv`.
	// **According to the different format, please select the following parameters**
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// Field configuration in csv content_type.
	CsvConfigColumns pulumi.StringArrayOutput `pulumi:"csvConfigColumns"`
	// Separator configuration in csv content_type.
	CsvConfigDelimiter pulumi.StringPtrOutput `pulumi:"csvConfigDelimiter"`
	// escape in csv content_type.
	CsvConfigEscape pulumi.StringPtrOutput `pulumi:"csvConfigEscape"`
	// Indicates whether to write the field name to the CSV file, the default value is `false`.
	CsvConfigHeader pulumi.BoolPtrOutput `pulumi:"csvConfigHeader"`
	// lineFeed in csv content_type.
	CsvConfigLinefeed pulumi.StringPtrOutput `pulumi:"csvConfigLinefeed"`
	// Invalid field content in csv content_type.
	CsvConfigNull pulumi.StringPtrOutput `pulumi:"csvConfigNull"`
	// Escape character in csv content_type.
	CsvConfigQuote pulumi.StringPtrOutput `pulumi:"csvConfigQuote"`
	// The display name for oss export.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
	ExportName pulumi.StringOutput `pulumi:"exportName"`
	// The log from when to export to oss.
	FromTime pulumi.IntPtrOutput `pulumi:"fromTime"`
	// Whether to deliver the label when `contentType` = `json`.
	JsonEnableTag pulumi.BoolPtrOutput `pulumi:"jsonEnableTag"`
	// Used for logstore reading, the role should have log read policy, such as `acs:ram::13234:role/logrole`, if `logReadRoleArn` is not set, `roleArn` is used to read logstore.
	LogReadRoleArn pulumi.StringPtrOutput `pulumi:"logReadRoleArn"`
	// The name of the log logstore.
	LogstoreName pulumi.StringOutput `pulumi:"logstoreName"`
	// The OSS Bucket directory is dynamically generated according to the creation time of the export task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
	PathFormat pulumi.StringOutput `pulumi:"pathFormat"`
	// The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The name of the log project. It is the only in one Alicloud account.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// Used to write to oss bucket, the OSS Bucket owner creates the role mark which has the oss bucket write policy, such as `acs:ram::13234:role/logrole`.
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// The suffix for the objects in which the shipped data is stored.
	Suffix pulumi.StringPtrOutput `pulumi:"suffix"`
	// This time zone that is used to format the time, `+0800` e.g.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
}

// NewOssExport registers a new resource with the given unique name, arguments, and options.
func NewOssExport(ctx *pulumi.Context,
	name string, args *OssExportArgs, opts ...pulumi.ResourceOption) (*OssExport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.BufferInterval == nil {
		return nil, errors.New("invalid value for required argument 'BufferInterval'")
	}
	if args.BufferSize == nil {
		return nil, errors.New("invalid value for required argument 'BufferSize'")
	}
	if args.ContentType == nil {
		return nil, errors.New("invalid value for required argument 'ContentType'")
	}
	if args.ExportName == nil {
		return nil, errors.New("invalid value for required argument 'ExportName'")
	}
	if args.LogstoreName == nil {
		return nil, errors.New("invalid value for required argument 'LogstoreName'")
	}
	if args.PathFormat == nil {
		return nil, errors.New("invalid value for required argument 'PathFormat'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.TimeZone == nil {
		return nil, errors.New("invalid value for required argument 'TimeZone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OssExport
	err := ctx.RegisterResource("alicloud:log/ossExport:OssExport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOssExport gets an existing OssExport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOssExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OssExportState, opts ...pulumi.ResourceOption) (*OssExport, error) {
	var resource OssExport
	err := ctx.ReadResource("alicloud:log/ossExport:OssExport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OssExport resources.
type ossExportState struct {
	// The name of the oss bucket.
	Bucket *string `pulumi:"bucket"`
	// How often is it delivered every interval.
	BufferInterval *int `pulumi:"bufferInterval"`
	// Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
	BufferSize *int `pulumi:"bufferSize"`
	// OSS data storage compression method, support: `none`, `snappy`, `zstd`, `gzip`. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
	CompressType *string `pulumi:"compressType"`
	// Configure columns when `contentType` is `parquet` or `orc`.
	ConfigColumns []OssExportConfigColumn `pulumi:"configColumns"`
	// Storage format, only supports three types: `json`, `parquet`, `orc`, `csv`.
	// **According to the different format, please select the following parameters**
	ContentType *string `pulumi:"contentType"`
	// Field configuration in csv content_type.
	CsvConfigColumns []string `pulumi:"csvConfigColumns"`
	// Separator configuration in csv content_type.
	CsvConfigDelimiter *string `pulumi:"csvConfigDelimiter"`
	// escape in csv content_type.
	CsvConfigEscape *string `pulumi:"csvConfigEscape"`
	// Indicates whether to write the field name to the CSV file, the default value is `false`.
	CsvConfigHeader *bool `pulumi:"csvConfigHeader"`
	// lineFeed in csv content_type.
	CsvConfigLinefeed *string `pulumi:"csvConfigLinefeed"`
	// Invalid field content in csv content_type.
	CsvConfigNull *string `pulumi:"csvConfigNull"`
	// Escape character in csv content_type.
	CsvConfigQuote *string `pulumi:"csvConfigQuote"`
	// The display name for oss export.
	DisplayName *string `pulumi:"displayName"`
	// Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
	ExportName *string `pulumi:"exportName"`
	// The log from when to export to oss.
	FromTime *int `pulumi:"fromTime"`
	// Whether to deliver the label when `contentType` = `json`.
	JsonEnableTag *bool `pulumi:"jsonEnableTag"`
	// Used for logstore reading, the role should have log read policy, such as `acs:ram::13234:role/logrole`, if `logReadRoleArn` is not set, `roleArn` is used to read logstore.
	LogReadRoleArn *string `pulumi:"logReadRoleArn"`
	// The name of the log logstore.
	LogstoreName *string `pulumi:"logstoreName"`
	// The OSS Bucket directory is dynamically generated according to the creation time of the export task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
	PathFormat *string `pulumi:"pathFormat"`
	// The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
	Prefix *string `pulumi:"prefix"`
	// The name of the log project. It is the only in one Alicloud account.
	ProjectName *string `pulumi:"projectName"`
	// Used to write to oss bucket, the OSS Bucket owner creates the role mark which has the oss bucket write policy, such as `acs:ram::13234:role/logrole`.
	RoleArn *string `pulumi:"roleArn"`
	// The suffix for the objects in which the shipped data is stored.
	Suffix *string `pulumi:"suffix"`
	// This time zone that is used to format the time, `+0800` e.g.
	TimeZone *string `pulumi:"timeZone"`
}

type OssExportState struct {
	// The name of the oss bucket.
	Bucket pulumi.StringPtrInput
	// How often is it delivered every interval.
	BufferInterval pulumi.IntPtrInput
	// Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
	BufferSize pulumi.IntPtrInput
	// OSS data storage compression method, support: `none`, `snappy`, `zstd`, `gzip`. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
	CompressType pulumi.StringPtrInput
	// Configure columns when `contentType` is `parquet` or `orc`.
	ConfigColumns OssExportConfigColumnArrayInput
	// Storage format, only supports three types: `json`, `parquet`, `orc`, `csv`.
	// **According to the different format, please select the following parameters**
	ContentType pulumi.StringPtrInput
	// Field configuration in csv content_type.
	CsvConfigColumns pulumi.StringArrayInput
	// Separator configuration in csv content_type.
	CsvConfigDelimiter pulumi.StringPtrInput
	// escape in csv content_type.
	CsvConfigEscape pulumi.StringPtrInput
	// Indicates whether to write the field name to the CSV file, the default value is `false`.
	CsvConfigHeader pulumi.BoolPtrInput
	// lineFeed in csv content_type.
	CsvConfigLinefeed pulumi.StringPtrInput
	// Invalid field content in csv content_type.
	CsvConfigNull pulumi.StringPtrInput
	// Escape character in csv content_type.
	CsvConfigQuote pulumi.StringPtrInput
	// The display name for oss export.
	DisplayName pulumi.StringPtrInput
	// Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
	ExportName pulumi.StringPtrInput
	// The log from when to export to oss.
	FromTime pulumi.IntPtrInput
	// Whether to deliver the label when `contentType` = `json`.
	JsonEnableTag pulumi.BoolPtrInput
	// Used for logstore reading, the role should have log read policy, such as `acs:ram::13234:role/logrole`, if `logReadRoleArn` is not set, `roleArn` is used to read logstore.
	LogReadRoleArn pulumi.StringPtrInput
	// The name of the log logstore.
	LogstoreName pulumi.StringPtrInput
	// The OSS Bucket directory is dynamically generated according to the creation time of the export task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
	PathFormat pulumi.StringPtrInput
	// The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
	Prefix pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account.
	ProjectName pulumi.StringPtrInput
	// Used to write to oss bucket, the OSS Bucket owner creates the role mark which has the oss bucket write policy, such as `acs:ram::13234:role/logrole`.
	RoleArn pulumi.StringPtrInput
	// The suffix for the objects in which the shipped data is stored.
	Suffix pulumi.StringPtrInput
	// This time zone that is used to format the time, `+0800` e.g.
	TimeZone pulumi.StringPtrInput
}

func (OssExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*ossExportState)(nil)).Elem()
}

type ossExportArgs struct {
	// The name of the oss bucket.
	Bucket string `pulumi:"bucket"`
	// How often is it delivered every interval.
	BufferInterval int `pulumi:"bufferInterval"`
	// Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
	BufferSize int `pulumi:"bufferSize"`
	// OSS data storage compression method, support: `none`, `snappy`, `zstd`, `gzip`. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
	CompressType *string `pulumi:"compressType"`
	// Configure columns when `contentType` is `parquet` or `orc`.
	ConfigColumns []OssExportConfigColumn `pulumi:"configColumns"`
	// Storage format, only supports three types: `json`, `parquet`, `orc`, `csv`.
	// **According to the different format, please select the following parameters**
	ContentType string `pulumi:"contentType"`
	// Field configuration in csv content_type.
	CsvConfigColumns []string `pulumi:"csvConfigColumns"`
	// Separator configuration in csv content_type.
	CsvConfigDelimiter *string `pulumi:"csvConfigDelimiter"`
	// escape in csv content_type.
	CsvConfigEscape *string `pulumi:"csvConfigEscape"`
	// Indicates whether to write the field name to the CSV file, the default value is `false`.
	CsvConfigHeader *bool `pulumi:"csvConfigHeader"`
	// lineFeed in csv content_type.
	CsvConfigLinefeed *string `pulumi:"csvConfigLinefeed"`
	// Invalid field content in csv content_type.
	CsvConfigNull *string `pulumi:"csvConfigNull"`
	// Escape character in csv content_type.
	CsvConfigQuote *string `pulumi:"csvConfigQuote"`
	// The display name for oss export.
	DisplayName *string `pulumi:"displayName"`
	// Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
	ExportName string `pulumi:"exportName"`
	// The log from when to export to oss.
	FromTime *int `pulumi:"fromTime"`
	// Whether to deliver the label when `contentType` = `json`.
	JsonEnableTag *bool `pulumi:"jsonEnableTag"`
	// Used for logstore reading, the role should have log read policy, such as `acs:ram::13234:role/logrole`, if `logReadRoleArn` is not set, `roleArn` is used to read logstore.
	LogReadRoleArn *string `pulumi:"logReadRoleArn"`
	// The name of the log logstore.
	LogstoreName string `pulumi:"logstoreName"`
	// The OSS Bucket directory is dynamically generated according to the creation time of the export task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
	PathFormat string `pulumi:"pathFormat"`
	// The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
	Prefix *string `pulumi:"prefix"`
	// The name of the log project. It is the only in one Alicloud account.
	ProjectName string `pulumi:"projectName"`
	// Used to write to oss bucket, the OSS Bucket owner creates the role mark which has the oss bucket write policy, such as `acs:ram::13234:role/logrole`.
	RoleArn *string `pulumi:"roleArn"`
	// The suffix for the objects in which the shipped data is stored.
	Suffix *string `pulumi:"suffix"`
	// This time zone that is used to format the time, `+0800` e.g.
	TimeZone string `pulumi:"timeZone"`
}

// The set of arguments for constructing a OssExport resource.
type OssExportArgs struct {
	// The name of the oss bucket.
	Bucket pulumi.StringInput
	// How often is it delivered every interval.
	BufferInterval pulumi.IntInput
	// Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
	BufferSize pulumi.IntInput
	// OSS data storage compression method, support: `none`, `snappy`, `zstd`, `gzip`. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
	CompressType pulumi.StringPtrInput
	// Configure columns when `contentType` is `parquet` or `orc`.
	ConfigColumns OssExportConfigColumnArrayInput
	// Storage format, only supports three types: `json`, `parquet`, `orc`, `csv`.
	// **According to the different format, please select the following parameters**
	ContentType pulumi.StringInput
	// Field configuration in csv content_type.
	CsvConfigColumns pulumi.StringArrayInput
	// Separator configuration in csv content_type.
	CsvConfigDelimiter pulumi.StringPtrInput
	// escape in csv content_type.
	CsvConfigEscape pulumi.StringPtrInput
	// Indicates whether to write the field name to the CSV file, the default value is `false`.
	CsvConfigHeader pulumi.BoolPtrInput
	// lineFeed in csv content_type.
	CsvConfigLinefeed pulumi.StringPtrInput
	// Invalid field content in csv content_type.
	CsvConfigNull pulumi.StringPtrInput
	// Escape character in csv content_type.
	CsvConfigQuote pulumi.StringPtrInput
	// The display name for oss export.
	DisplayName pulumi.StringPtrInput
	// Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
	ExportName pulumi.StringInput
	// The log from when to export to oss.
	FromTime pulumi.IntPtrInput
	// Whether to deliver the label when `contentType` = `json`.
	JsonEnableTag pulumi.BoolPtrInput
	// Used for logstore reading, the role should have log read policy, such as `acs:ram::13234:role/logrole`, if `logReadRoleArn` is not set, `roleArn` is used to read logstore.
	LogReadRoleArn pulumi.StringPtrInput
	// The name of the log logstore.
	LogstoreName pulumi.StringInput
	// The OSS Bucket directory is dynamically generated according to the creation time of the export task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
	PathFormat pulumi.StringInput
	// The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
	Prefix pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account.
	ProjectName pulumi.StringInput
	// Used to write to oss bucket, the OSS Bucket owner creates the role mark which has the oss bucket write policy, such as `acs:ram::13234:role/logrole`.
	RoleArn pulumi.StringPtrInput
	// The suffix for the objects in which the shipped data is stored.
	Suffix pulumi.StringPtrInput
	// This time zone that is used to format the time, `+0800` e.g.
	TimeZone pulumi.StringInput
}

func (OssExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ossExportArgs)(nil)).Elem()
}

type OssExportInput interface {
	pulumi.Input

	ToOssExportOutput() OssExportOutput
	ToOssExportOutputWithContext(ctx context.Context) OssExportOutput
}

func (*OssExport) ElementType() reflect.Type {
	return reflect.TypeOf((**OssExport)(nil)).Elem()
}

func (i *OssExport) ToOssExportOutput() OssExportOutput {
	return i.ToOssExportOutputWithContext(context.Background())
}

func (i *OssExport) ToOssExportOutputWithContext(ctx context.Context) OssExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OssExportOutput)
}

// OssExportArrayInput is an input type that accepts OssExportArray and OssExportArrayOutput values.
// You can construct a concrete instance of `OssExportArrayInput` via:
//
//	OssExportArray{ OssExportArgs{...} }
type OssExportArrayInput interface {
	pulumi.Input

	ToOssExportArrayOutput() OssExportArrayOutput
	ToOssExportArrayOutputWithContext(context.Context) OssExportArrayOutput
}

type OssExportArray []OssExportInput

func (OssExportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OssExport)(nil)).Elem()
}

func (i OssExportArray) ToOssExportArrayOutput() OssExportArrayOutput {
	return i.ToOssExportArrayOutputWithContext(context.Background())
}

func (i OssExportArray) ToOssExportArrayOutputWithContext(ctx context.Context) OssExportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OssExportArrayOutput)
}

// OssExportMapInput is an input type that accepts OssExportMap and OssExportMapOutput values.
// You can construct a concrete instance of `OssExportMapInput` via:
//
//	OssExportMap{ "key": OssExportArgs{...} }
type OssExportMapInput interface {
	pulumi.Input

	ToOssExportMapOutput() OssExportMapOutput
	ToOssExportMapOutputWithContext(context.Context) OssExportMapOutput
}

type OssExportMap map[string]OssExportInput

func (OssExportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OssExport)(nil)).Elem()
}

func (i OssExportMap) ToOssExportMapOutput() OssExportMapOutput {
	return i.ToOssExportMapOutputWithContext(context.Background())
}

func (i OssExportMap) ToOssExportMapOutputWithContext(ctx context.Context) OssExportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OssExportMapOutput)
}

type OssExportOutput struct{ *pulumi.OutputState }

func (OssExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OssExport)(nil)).Elem()
}

func (o OssExportOutput) ToOssExportOutput() OssExportOutput {
	return o
}

func (o OssExportOutput) ToOssExportOutputWithContext(ctx context.Context) OssExportOutput {
	return o
}

// The name of the oss bucket.
func (o OssExportOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// How often is it delivered every interval.
func (o OssExportOutput) BufferInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *OssExport) pulumi.IntOutput { return v.BufferInterval }).(pulumi.IntOutput)
}

// Automatically control the creation interval of delivery tasks and set the upper limit of an OSS object size (calculated in uncompressed), unit: `MB`.
func (o OssExportOutput) BufferSize() pulumi.IntOutput {
	return o.ApplyT(func(v *OssExport) pulumi.IntOutput { return v.BufferSize }).(pulumi.IntOutput)
}

// OSS data storage compression method, support: `none`, `snappy`, `zstd`, `gzip`. Among them, none means that the original data is not compressed, and snappy means that the data is compressed using the snappy algorithm, which can reduce the storage space usage of the `OSS Bucket`.
func (o OssExportOutput) CompressType() pulumi.StringOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringOutput { return v.CompressType }).(pulumi.StringOutput)
}

// Configure columns when `contentType` is `parquet` or `orc`.
func (o OssExportOutput) ConfigColumns() OssExportConfigColumnArrayOutput {
	return o.ApplyT(func(v *OssExport) OssExportConfigColumnArrayOutput { return v.ConfigColumns }).(OssExportConfigColumnArrayOutput)
}

// Storage format, only supports three types: `json`, `parquet`, `orc`, `csv`.
// **According to the different format, please select the following parameters**
func (o OssExportOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

// Field configuration in csv content_type.
func (o OssExportOutput) CsvConfigColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringArrayOutput { return v.CsvConfigColumns }).(pulumi.StringArrayOutput)
}

// Separator configuration in csv content_type.
func (o OssExportOutput) CsvConfigDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringPtrOutput { return v.CsvConfigDelimiter }).(pulumi.StringPtrOutput)
}

// escape in csv content_type.
func (o OssExportOutput) CsvConfigEscape() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringPtrOutput { return v.CsvConfigEscape }).(pulumi.StringPtrOutput)
}

// Indicates whether to write the field name to the CSV file, the default value is `false`.
func (o OssExportOutput) CsvConfigHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.BoolPtrOutput { return v.CsvConfigHeader }).(pulumi.BoolPtrOutput)
}

// lineFeed in csv content_type.
func (o OssExportOutput) CsvConfigLinefeed() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringPtrOutput { return v.CsvConfigLinefeed }).(pulumi.StringPtrOutput)
}

// Invalid field content in csv content_type.
func (o OssExportOutput) CsvConfigNull() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringPtrOutput { return v.CsvConfigNull }).(pulumi.StringPtrOutput)
}

// Escape character in csv content_type.
func (o OssExportOutput) CsvConfigQuote() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringPtrOutput { return v.CsvConfigQuote }).(pulumi.StringPtrOutput)
}

// The display name for oss export.
func (o OssExportOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Delivery configuration name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
func (o OssExportOutput) ExportName() pulumi.StringOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringOutput { return v.ExportName }).(pulumi.StringOutput)
}

// The log from when to export to oss.
func (o OssExportOutput) FromTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.IntPtrOutput { return v.FromTime }).(pulumi.IntPtrOutput)
}

// Whether to deliver the label when `contentType` = `json`.
func (o OssExportOutput) JsonEnableTag() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.BoolPtrOutput { return v.JsonEnableTag }).(pulumi.BoolPtrOutput)
}

// Used for logstore reading, the role should have log read policy, such as `acs:ram::13234:role/logrole`, if `logReadRoleArn` is not set, `roleArn` is used to read logstore.
func (o OssExportOutput) LogReadRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringPtrOutput { return v.LogReadRoleArn }).(pulumi.StringPtrOutput)
}

// The name of the log logstore.
func (o OssExportOutput) LogstoreName() pulumi.StringOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringOutput { return v.LogstoreName }).(pulumi.StringOutput)
}

// The OSS Bucket directory is dynamically generated according to the creation time of the export task, it cannot start with a forward slash `/`, the default value is `%Y/%m/%d/%H/%M`.
func (o OssExportOutput) PathFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringOutput { return v.PathFormat }).(pulumi.StringOutput)
}

// The data synchronized from Log Service to OSS will be stored in this directory of Bucket.
func (o OssExportOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringPtrOutput { return v.Prefix }).(pulumi.StringPtrOutput)
}

// The name of the log project. It is the only in one Alicloud account.
func (o OssExportOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// Used to write to oss bucket, the OSS Bucket owner creates the role mark which has the oss bucket write policy, such as `acs:ram::13234:role/logrole`.
func (o OssExportOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringPtrOutput { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// The suffix for the objects in which the shipped data is stored.
func (o OssExportOutput) Suffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringPtrOutput { return v.Suffix }).(pulumi.StringPtrOutput)
}

// This time zone that is used to format the time, `+0800` e.g.
func (o OssExportOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *OssExport) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

type OssExportArrayOutput struct{ *pulumi.OutputState }

func (OssExportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OssExport)(nil)).Elem()
}

func (o OssExportArrayOutput) ToOssExportArrayOutput() OssExportArrayOutput {
	return o
}

func (o OssExportArrayOutput) ToOssExportArrayOutputWithContext(ctx context.Context) OssExportArrayOutput {
	return o
}

func (o OssExportArrayOutput) Index(i pulumi.IntInput) OssExportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OssExport {
		return vs[0].([]*OssExport)[vs[1].(int)]
	}).(OssExportOutput)
}

type OssExportMapOutput struct{ *pulumi.OutputState }

func (OssExportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OssExport)(nil)).Elem()
}

func (o OssExportMapOutput) ToOssExportMapOutput() OssExportMapOutput {
	return o
}

func (o OssExportMapOutput) ToOssExportMapOutputWithContext(ctx context.Context) OssExportMapOutput {
	return o
}

func (o OssExportMapOutput) MapIndex(k pulumi.StringInput) OssExportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OssExport {
		return vs[0].(map[string]*OssExport)[vs[1].(string)]
	}).(OssExportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OssExportInput)(nil)).Elem(), &OssExport{})
	pulumi.RegisterInputType(reflect.TypeOf((*OssExportArrayInput)(nil)).Elem(), OssExportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OssExportMapInput)(nil)).Elem(), OssExportMap{})
	pulumi.RegisterOutputType(OssExportOutput{})
	pulumi.RegisterOutputType(OssExportArrayOutput{})
	pulumi.RegisterOutputType(OssExportMapOutput{})
}
