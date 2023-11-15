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

// The data transformation of the log service is a hosted, highly available, and scalable data processing service,
// which is widely applicable to scenarios such as data regularization, enrichment, distribution, aggregation, and index reconstruction.
// [Refer to details](https://www.alibabacloud.com/help/zh/doc-detail/125384.htm).
//
// > **NOTE:** Available in 1.120.0
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
//			example2, err := log.NewStore(ctx, "example2", &log.StoreArgs{
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
//			example3, err := log.NewStore(ctx, "example3", &log.StoreArgs{
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
//			_, err = log.NewEtl(ctx, "exampleEtl", &log.EtlArgs{
//				EtlName:         pulumi.String("terraform-example"),
//				Project:         exampleProject.Name,
//				DisplayName:     pulumi.String("terraform-example"),
//				Description:     pulumi.String("terraform-example"),
//				AccessKeyId:     pulumi.String("access_key_id"),
//				AccessKeySecret: pulumi.String("access_key_secret"),
//				Script:          pulumi.String("e_set('new','key')"),
//				Logstore:        exampleStore.Name,
//				EtlSinks: log.EtlEtlSinkArray{
//					&log.EtlEtlSinkArgs{
//						Name:            pulumi.String("target_name"),
//						AccessKeyId:     pulumi.String("example2_access_key_id"),
//						AccessKeySecret: pulumi.String("example2_access_key_secret"),
//						Endpoint:        pulumi.String("cn-hangzhou.log.aliyuncs.com"),
//						Project:         exampleProject.Name,
//						Logstore:        example2.Name,
//					},
//					&log.EtlEtlSinkArgs{
//						Name:            pulumi.String("target_name2"),
//						AccessKeyId:     pulumi.String("example3_access_key_id"),
//						AccessKeySecret: pulumi.String("example3_access_key_secret"),
//						Endpoint:        pulumi.String("cn-hangzhou.log.aliyuncs.com"),
//						Project:         exampleProject.Name,
//						Logstore:        example3.Name,
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
// ## Import
//
// Log etl can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:log/etl:Etl example tf-log-project:tf-log-etl-name
//
// ```
type Etl struct {
	pulumi.CustomResourceState

	// Delivery target logstore access key id.
	AccessKeyId pulumi.StringPtrOutput `pulumi:"accessKeyId"`
	// Delivery target logstore access key secret.
	AccessKeySecret pulumi.StringPtrOutput `pulumi:"accessKeySecret"`
	// The etl job create time.
	CreateTime pulumi.IntOutput `pulumi:"createTime"`
	// Description of the log etl job.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Log service etl job alias.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the log etl job.
	EtlName pulumi.StringOutput `pulumi:"etlName"`
	// Target logstore configuration for delivery after data processing.
	EtlSinks EtlEtlSinkArrayOutput `pulumi:"etlSinks"`
	// Log service etl type, the default value is `ETL`.
	EtlType pulumi.StringPtrOutput `pulumi:"etlType"`
	// The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
	FromTime pulumi.IntPtrOutput `pulumi:"fromTime"`
	// An KMS encrypts access key id used to a log etl job. If the `accessKeyId` is filled in, this field will be ignored.
	KmsEncryptedAccessKeyId pulumi.StringPtrOutput `pulumi:"kmsEncryptedAccessKeyId"`
	// An KMS encrypts access key secret used to a log etl job. If the `accessKeySecret` is filled in, this field will be ignored.
	KmsEncryptedAccessKeySecret pulumi.StringPtrOutput `pulumi:"kmsEncryptedAccessKeySecret"`
	// An KMS encryption context used to decrypt `kmsEncryptedAccessKeyId` before creating or updating an instance with `kmsEncryptedAccessKeyId`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
	KmsEncryptionAccessKeyIdContext pulumi.MapOutput `pulumi:"kmsEncryptionAccessKeyIdContext"`
	// An KMS encryption context used to decrypt `kmsEncryptedAccessKeySecret` before creating or updating an instance with `kmsEncryptedAccessKeySecret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
	KmsEncryptionAccessKeySecretContext pulumi.MapOutput `pulumi:"kmsEncryptionAccessKeySecretContext"`
	// ETL job last modified time.
	LastModifiedTime pulumi.IntOutput `pulumi:"lastModifiedTime"`
	// Delivery target logstore.
	Logstore pulumi.StringOutput `pulumi:"logstore"`
	// Advanced parameter configuration of processing operations.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The project where the target logstore is delivered.
	Project pulumi.StringOutput `pulumi:"project"`
	// Sts role info under delivery target logstore. `roleArn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// Job scheduling type, the default value is Resident.
	Schedule pulumi.StringPtrOutput `pulumi:"schedule"`
	// Processing operation grammar.
	Script pulumi.StringOutput `pulumi:"script"`
	// Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
	ToTime pulumi.IntPtrOutput `pulumi:"toTime"`
	// Log etl job version. the default value is `2`.
	Version pulumi.IntPtrOutput `pulumi:"version"`
}

// NewEtl registers a new resource with the given unique name, arguments, and options.
func NewEtl(ctx *pulumi.Context,
	name string, args *EtlArgs, opts ...pulumi.ResourceOption) (*Etl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.EtlName == nil {
		return nil, errors.New("invalid value for required argument 'EtlName'")
	}
	if args.EtlSinks == nil {
		return nil, errors.New("invalid value for required argument 'EtlSinks'")
	}
	if args.Logstore == nil {
		return nil, errors.New("invalid value for required argument 'Logstore'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Script == nil {
		return nil, errors.New("invalid value for required argument 'Script'")
	}
	if args.AccessKeyId != nil {
		args.AccessKeyId = pulumi.ToSecret(args.AccessKeyId).(pulumi.StringPtrInput)
	}
	if args.AccessKeySecret != nil {
		args.AccessKeySecret = pulumi.ToSecret(args.AccessKeySecret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessKeyId",
		"accessKeySecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Etl
	err := ctx.RegisterResource("alicloud:log/etl:Etl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEtl gets an existing Etl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEtl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EtlState, opts ...pulumi.ResourceOption) (*Etl, error) {
	var resource Etl
	err := ctx.ReadResource("alicloud:log/etl:Etl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Etl resources.
type etlState struct {
	// Delivery target logstore access key id.
	AccessKeyId *string `pulumi:"accessKeyId"`
	// Delivery target logstore access key secret.
	AccessKeySecret *string `pulumi:"accessKeySecret"`
	// The etl job create time.
	CreateTime *int `pulumi:"createTime"`
	// Description of the log etl job.
	Description *string `pulumi:"description"`
	// Log service etl job alias.
	DisplayName *string `pulumi:"displayName"`
	// The name of the log etl job.
	EtlName *string `pulumi:"etlName"`
	// Target logstore configuration for delivery after data processing.
	EtlSinks []EtlEtlSink `pulumi:"etlSinks"`
	// Log service etl type, the default value is `ETL`.
	EtlType *string `pulumi:"etlType"`
	// The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
	FromTime *int `pulumi:"fromTime"`
	// An KMS encrypts access key id used to a log etl job. If the `accessKeyId` is filled in, this field will be ignored.
	KmsEncryptedAccessKeyId *string `pulumi:"kmsEncryptedAccessKeyId"`
	// An KMS encrypts access key secret used to a log etl job. If the `accessKeySecret` is filled in, this field will be ignored.
	KmsEncryptedAccessKeySecret *string `pulumi:"kmsEncryptedAccessKeySecret"`
	// An KMS encryption context used to decrypt `kmsEncryptedAccessKeyId` before creating or updating an instance with `kmsEncryptedAccessKeyId`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
	KmsEncryptionAccessKeyIdContext map[string]interface{} `pulumi:"kmsEncryptionAccessKeyIdContext"`
	// An KMS encryption context used to decrypt `kmsEncryptedAccessKeySecret` before creating or updating an instance with `kmsEncryptedAccessKeySecret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
	KmsEncryptionAccessKeySecretContext map[string]interface{} `pulumi:"kmsEncryptionAccessKeySecretContext"`
	// ETL job last modified time.
	LastModifiedTime *int `pulumi:"lastModifiedTime"`
	// Delivery target logstore.
	Logstore *string `pulumi:"logstore"`
	// Advanced parameter configuration of processing operations.
	Parameters map[string]string `pulumi:"parameters"`
	// The project where the target logstore is delivered.
	Project *string `pulumi:"project"`
	// Sts role info under delivery target logstore. `roleArn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
	RoleArn *string `pulumi:"roleArn"`
	// Job scheduling type, the default value is Resident.
	Schedule *string `pulumi:"schedule"`
	// Processing operation grammar.
	Script *string `pulumi:"script"`
	// Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
	Status *string `pulumi:"status"`
	// Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
	ToTime *int `pulumi:"toTime"`
	// Log etl job version. the default value is `2`.
	Version *int `pulumi:"version"`
}

type EtlState struct {
	// Delivery target logstore access key id.
	AccessKeyId pulumi.StringPtrInput
	// Delivery target logstore access key secret.
	AccessKeySecret pulumi.StringPtrInput
	// The etl job create time.
	CreateTime pulumi.IntPtrInput
	// Description of the log etl job.
	Description pulumi.StringPtrInput
	// Log service etl job alias.
	DisplayName pulumi.StringPtrInput
	// The name of the log etl job.
	EtlName pulumi.StringPtrInput
	// Target logstore configuration for delivery after data processing.
	EtlSinks EtlEtlSinkArrayInput
	// Log service etl type, the default value is `ETL`.
	EtlType pulumi.StringPtrInput
	// The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
	FromTime pulumi.IntPtrInput
	// An KMS encrypts access key id used to a log etl job. If the `accessKeyId` is filled in, this field will be ignored.
	KmsEncryptedAccessKeyId pulumi.StringPtrInput
	// An KMS encrypts access key secret used to a log etl job. If the `accessKeySecret` is filled in, this field will be ignored.
	KmsEncryptedAccessKeySecret pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedAccessKeyId` before creating or updating an instance with `kmsEncryptedAccessKeyId`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
	KmsEncryptionAccessKeyIdContext pulumi.MapInput
	// An KMS encryption context used to decrypt `kmsEncryptedAccessKeySecret` before creating or updating an instance with `kmsEncryptedAccessKeySecret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
	KmsEncryptionAccessKeySecretContext pulumi.MapInput
	// ETL job last modified time.
	LastModifiedTime pulumi.IntPtrInput
	// Delivery target logstore.
	Logstore pulumi.StringPtrInput
	// Advanced parameter configuration of processing operations.
	Parameters pulumi.StringMapInput
	// The project where the target logstore is delivered.
	Project pulumi.StringPtrInput
	// Sts role info under delivery target logstore. `roleArn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
	RoleArn pulumi.StringPtrInput
	// Job scheduling type, the default value is Resident.
	Schedule pulumi.StringPtrInput
	// Processing operation grammar.
	Script pulumi.StringPtrInput
	// Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
	Status pulumi.StringPtrInput
	// Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
	ToTime pulumi.IntPtrInput
	// Log etl job version. the default value is `2`.
	Version pulumi.IntPtrInput
}

func (EtlState) ElementType() reflect.Type {
	return reflect.TypeOf((*etlState)(nil)).Elem()
}

type etlArgs struct {
	// Delivery target logstore access key id.
	AccessKeyId *string `pulumi:"accessKeyId"`
	// Delivery target logstore access key secret.
	AccessKeySecret *string `pulumi:"accessKeySecret"`
	// The etl job create time.
	CreateTime *int `pulumi:"createTime"`
	// Description of the log etl job.
	Description *string `pulumi:"description"`
	// Log service etl job alias.
	DisplayName string `pulumi:"displayName"`
	// The name of the log etl job.
	EtlName string `pulumi:"etlName"`
	// Target logstore configuration for delivery after data processing.
	EtlSinks []EtlEtlSink `pulumi:"etlSinks"`
	// Log service etl type, the default value is `ETL`.
	EtlType *string `pulumi:"etlType"`
	// The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
	FromTime *int `pulumi:"fromTime"`
	// An KMS encrypts access key id used to a log etl job. If the `accessKeyId` is filled in, this field will be ignored.
	KmsEncryptedAccessKeyId *string `pulumi:"kmsEncryptedAccessKeyId"`
	// An KMS encrypts access key secret used to a log etl job. If the `accessKeySecret` is filled in, this field will be ignored.
	KmsEncryptedAccessKeySecret *string `pulumi:"kmsEncryptedAccessKeySecret"`
	// An KMS encryption context used to decrypt `kmsEncryptedAccessKeyId` before creating or updating an instance with `kmsEncryptedAccessKeyId`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
	KmsEncryptionAccessKeyIdContext map[string]interface{} `pulumi:"kmsEncryptionAccessKeyIdContext"`
	// An KMS encryption context used to decrypt `kmsEncryptedAccessKeySecret` before creating or updating an instance with `kmsEncryptedAccessKeySecret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
	KmsEncryptionAccessKeySecretContext map[string]interface{} `pulumi:"kmsEncryptionAccessKeySecretContext"`
	// ETL job last modified time.
	LastModifiedTime *int `pulumi:"lastModifiedTime"`
	// Delivery target logstore.
	Logstore string `pulumi:"logstore"`
	// Advanced parameter configuration of processing operations.
	Parameters map[string]string `pulumi:"parameters"`
	// The project where the target logstore is delivered.
	Project string `pulumi:"project"`
	// Sts role info under delivery target logstore. `roleArn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
	RoleArn *string `pulumi:"roleArn"`
	// Job scheduling type, the default value is Resident.
	Schedule *string `pulumi:"schedule"`
	// Processing operation grammar.
	Script string `pulumi:"script"`
	// Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
	Status *string `pulumi:"status"`
	// Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
	ToTime *int `pulumi:"toTime"`
	// Log etl job version. the default value is `2`.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a Etl resource.
type EtlArgs struct {
	// Delivery target logstore access key id.
	AccessKeyId pulumi.StringPtrInput
	// Delivery target logstore access key secret.
	AccessKeySecret pulumi.StringPtrInput
	// The etl job create time.
	CreateTime pulumi.IntPtrInput
	// Description of the log etl job.
	Description pulumi.StringPtrInput
	// Log service etl job alias.
	DisplayName pulumi.StringInput
	// The name of the log etl job.
	EtlName pulumi.StringInput
	// Target logstore configuration for delivery after data processing.
	EtlSinks EtlEtlSinkArrayInput
	// Log service etl type, the default value is `ETL`.
	EtlType pulumi.StringPtrInput
	// The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
	FromTime pulumi.IntPtrInput
	// An KMS encrypts access key id used to a log etl job. If the `accessKeyId` is filled in, this field will be ignored.
	KmsEncryptedAccessKeyId pulumi.StringPtrInput
	// An KMS encrypts access key secret used to a log etl job. If the `accessKeySecret` is filled in, this field will be ignored.
	KmsEncryptedAccessKeySecret pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedAccessKeyId` before creating or updating an instance with `kmsEncryptedAccessKeyId`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
	KmsEncryptionAccessKeyIdContext pulumi.MapInput
	// An KMS encryption context used to decrypt `kmsEncryptedAccessKeySecret` before creating or updating an instance with `kmsEncryptedAccessKeySecret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
	KmsEncryptionAccessKeySecretContext pulumi.MapInput
	// ETL job last modified time.
	LastModifiedTime pulumi.IntPtrInput
	// Delivery target logstore.
	Logstore pulumi.StringInput
	// Advanced parameter configuration of processing operations.
	Parameters pulumi.StringMapInput
	// The project where the target logstore is delivered.
	Project pulumi.StringInput
	// Sts role info under delivery target logstore. `roleArn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
	RoleArn pulumi.StringPtrInput
	// Job scheduling type, the default value is Resident.
	Schedule pulumi.StringPtrInput
	// Processing operation grammar.
	Script pulumi.StringInput
	// Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
	Status pulumi.StringPtrInput
	// Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
	ToTime pulumi.IntPtrInput
	// Log etl job version. the default value is `2`.
	Version pulumi.IntPtrInput
}

func (EtlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*etlArgs)(nil)).Elem()
}

type EtlInput interface {
	pulumi.Input

	ToEtlOutput() EtlOutput
	ToEtlOutputWithContext(ctx context.Context) EtlOutput
}

func (*Etl) ElementType() reflect.Type {
	return reflect.TypeOf((**Etl)(nil)).Elem()
}

func (i *Etl) ToEtlOutput() EtlOutput {
	return i.ToEtlOutputWithContext(context.Background())
}

func (i *Etl) ToEtlOutputWithContext(ctx context.Context) EtlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtlOutput)
}

// EtlArrayInput is an input type that accepts EtlArray and EtlArrayOutput values.
// You can construct a concrete instance of `EtlArrayInput` via:
//
//	EtlArray{ EtlArgs{...} }
type EtlArrayInput interface {
	pulumi.Input

	ToEtlArrayOutput() EtlArrayOutput
	ToEtlArrayOutputWithContext(context.Context) EtlArrayOutput
}

type EtlArray []EtlInput

func (EtlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Etl)(nil)).Elem()
}

func (i EtlArray) ToEtlArrayOutput() EtlArrayOutput {
	return i.ToEtlArrayOutputWithContext(context.Background())
}

func (i EtlArray) ToEtlArrayOutputWithContext(ctx context.Context) EtlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtlArrayOutput)
}

// EtlMapInput is an input type that accepts EtlMap and EtlMapOutput values.
// You can construct a concrete instance of `EtlMapInput` via:
//
//	EtlMap{ "key": EtlArgs{...} }
type EtlMapInput interface {
	pulumi.Input

	ToEtlMapOutput() EtlMapOutput
	ToEtlMapOutputWithContext(context.Context) EtlMapOutput
}

type EtlMap map[string]EtlInput

func (EtlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Etl)(nil)).Elem()
}

func (i EtlMap) ToEtlMapOutput() EtlMapOutput {
	return i.ToEtlMapOutputWithContext(context.Background())
}

func (i EtlMap) ToEtlMapOutputWithContext(ctx context.Context) EtlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtlMapOutput)
}

type EtlOutput struct{ *pulumi.OutputState }

func (EtlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Etl)(nil)).Elem()
}

func (o EtlOutput) ToEtlOutput() EtlOutput {
	return o
}

func (o EtlOutput) ToEtlOutputWithContext(ctx context.Context) EtlOutput {
	return o
}

// Delivery target logstore access key id.
func (o EtlOutput) AccessKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringPtrOutput { return v.AccessKeyId }).(pulumi.StringPtrOutput)
}

// Delivery target logstore access key secret.
func (o EtlOutput) AccessKeySecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringPtrOutput { return v.AccessKeySecret }).(pulumi.StringPtrOutput)
}

// The etl job create time.
func (o EtlOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Etl) pulumi.IntOutput { return v.CreateTime }).(pulumi.IntOutput)
}

// Description of the log etl job.
func (o EtlOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Log service etl job alias.
func (o EtlOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The name of the log etl job.
func (o EtlOutput) EtlName() pulumi.StringOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringOutput { return v.EtlName }).(pulumi.StringOutput)
}

// Target logstore configuration for delivery after data processing.
func (o EtlOutput) EtlSinks() EtlEtlSinkArrayOutput {
	return o.ApplyT(func(v *Etl) EtlEtlSinkArrayOutput { return v.EtlSinks }).(EtlEtlSinkArrayOutput)
}

// Log service etl type, the default value is `ETL`.
func (o EtlOutput) EtlType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringPtrOutput { return v.EtlType }).(pulumi.StringPtrOutput)
}

// The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
func (o EtlOutput) FromTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.IntPtrOutput { return v.FromTime }).(pulumi.IntPtrOutput)
}

// An KMS encrypts access key id used to a log etl job. If the `accessKeyId` is filled in, this field will be ignored.
func (o EtlOutput) KmsEncryptedAccessKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringPtrOutput { return v.KmsEncryptedAccessKeyId }).(pulumi.StringPtrOutput)
}

// An KMS encrypts access key secret used to a log etl job. If the `accessKeySecret` is filled in, this field will be ignored.
func (o EtlOutput) KmsEncryptedAccessKeySecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringPtrOutput { return v.KmsEncryptedAccessKeySecret }).(pulumi.StringPtrOutput)
}

// An KMS encryption context used to decrypt `kmsEncryptedAccessKeyId` before creating or updating an instance with `kmsEncryptedAccessKeyId`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
func (o EtlOutput) KmsEncryptionAccessKeyIdContext() pulumi.MapOutput {
	return o.ApplyT(func(v *Etl) pulumi.MapOutput { return v.KmsEncryptionAccessKeyIdContext }).(pulumi.MapOutput)
}

// An KMS encryption context used to decrypt `kmsEncryptedAccessKeySecret` before creating or updating an instance with `kmsEncryptedAccessKeySecret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set. When it is changed, the instance will reboot to make the change take effect.
func (o EtlOutput) KmsEncryptionAccessKeySecretContext() pulumi.MapOutput {
	return o.ApplyT(func(v *Etl) pulumi.MapOutput { return v.KmsEncryptionAccessKeySecretContext }).(pulumi.MapOutput)
}

// ETL job last modified time.
func (o EtlOutput) LastModifiedTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Etl) pulumi.IntOutput { return v.LastModifiedTime }).(pulumi.IntOutput)
}

// Delivery target logstore.
func (o EtlOutput) Logstore() pulumi.StringOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringOutput { return v.Logstore }).(pulumi.StringOutput)
}

// Advanced parameter configuration of processing operations.
func (o EtlOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

// The project where the target logstore is delivered.
func (o EtlOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Sts role info under delivery target logstore. `roleArn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
func (o EtlOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringPtrOutput { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// Job scheduling type, the default value is Resident.
func (o EtlOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringPtrOutput { return v.Schedule }).(pulumi.StringPtrOutput)
}

// Processing operation grammar.
func (o EtlOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringOutput { return v.Script }).(pulumi.StringOutput)
}

// Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
func (o EtlOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Etl) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
func (o EtlOutput) ToTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.IntPtrOutput { return v.ToTime }).(pulumi.IntPtrOutput)
}

// Log etl job version. the default value is `2`.
func (o EtlOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Etl) pulumi.IntPtrOutput { return v.Version }).(pulumi.IntPtrOutput)
}

type EtlArrayOutput struct{ *pulumi.OutputState }

func (EtlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Etl)(nil)).Elem()
}

func (o EtlArrayOutput) ToEtlArrayOutput() EtlArrayOutput {
	return o
}

func (o EtlArrayOutput) ToEtlArrayOutputWithContext(ctx context.Context) EtlArrayOutput {
	return o
}

func (o EtlArrayOutput) Index(i pulumi.IntInput) EtlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Etl {
		return vs[0].([]*Etl)[vs[1].(int)]
	}).(EtlOutput)
}

type EtlMapOutput struct{ *pulumi.OutputState }

func (EtlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Etl)(nil)).Elem()
}

func (o EtlMapOutput) ToEtlMapOutput() EtlMapOutput {
	return o
}

func (o EtlMapOutput) ToEtlMapOutputWithContext(ctx context.Context) EtlMapOutput {
	return o
}

func (o EtlMapOutput) MapIndex(k pulumi.StringInput) EtlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Etl {
		return vs[0].(map[string]*Etl)[vs[1].(string)]
	}).(EtlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EtlInput)(nil)).Elem(), &Etl{})
	pulumi.RegisterInputType(reflect.TypeOf((*EtlArrayInput)(nil)).Elem(), EtlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EtlMapInput)(nil)).Elem(), EtlMap{})
	pulumi.RegisterOutputType(EtlOutput{})
	pulumi.RegisterOutputType(EtlArrayOutput{})
	pulumi.RegisterOutputType(EtlMapOutput{})
}
