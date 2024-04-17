// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log
{
    /// <summary>
    /// The data transformation of the log service is a hosted, highly available, and scalable data processing service,
    /// which is widely applicable to scenarios such as data regularization, enrichment, distribution, aggregation, and index reconstruction.
    /// [Refer to details](https://www.alibabacloud.com/help/zh/doc-detail/125384.htm).
    /// 
    /// &gt; **NOTE:** Available in 1.120.0
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Random.Index.Integer("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var example = new AliCloud.Log.Project("example", new()
    ///     {
    ///         Name = $"terraform-example-{@default.Result}",
    ///         Description = "terraform-example",
    ///     });
    /// 
    ///     var exampleStore = new AliCloud.Log.Store("example", new()
    ///     {
    ///         Project = example.Name,
    ///         Name = "example-store",
    ///         RetentionPeriod = 3650,
    ///         ShardCount = 3,
    ///         AutoSplit = true,
    ///         MaxSplitShardCount = 60,
    ///         AppendMeta = true,
    ///     });
    /// 
    ///     var example2 = new AliCloud.Log.Store("example2", new()
    ///     {
    ///         Project = example.Name,
    ///         Name = "example-store2",
    ///         RetentionPeriod = 3650,
    ///         ShardCount = 3,
    ///         AutoSplit = true,
    ///         MaxSplitShardCount = 60,
    ///         AppendMeta = true,
    ///     });
    /// 
    ///     var example3 = new AliCloud.Log.Store("example3", new()
    ///     {
    ///         Project = example.Name,
    ///         Name = "example-store3",
    ///         RetentionPeriod = 3650,
    ///         ShardCount = 3,
    ///         AutoSplit = true,
    ///         MaxSplitShardCount = 60,
    ///         AppendMeta = true,
    ///     });
    /// 
    ///     var exampleEtl = new AliCloud.Log.Etl("example", new()
    ///     {
    ///         EtlName = "terraform-example",
    ///         Project = example.Name,
    ///         DisplayName = "terraform-example",
    ///         Description = "terraform-example",
    ///         AccessKeyId = "access_key_id",
    ///         AccessKeySecret = "access_key_secret",
    ///         Script = "e_set('new','key')",
    ///         Logstore = exampleStore.Name,
    ///         EtlSinks = new[]
    ///         {
    ///             new AliCloud.Log.Inputs.EtlEtlSinkArgs
    ///             {
    ///                 Name = "target_name",
    ///                 AccessKeyId = "example2_access_key_id",
    ///                 AccessKeySecret = "example2_access_key_secret",
    ///                 Endpoint = "cn-hangzhou.log.aliyuncs.com",
    ///                 Project = example.Name,
    ///                 Logstore = example2.Name,
    ///             },
    ///             new AliCloud.Log.Inputs.EtlEtlSinkArgs
    ///             {
    ///                 Name = "target_name2",
    ///                 AccessKeyId = "example3_access_key_id",
    ///                 AccessKeySecret = "example3_access_key_secret",
    ///                 Endpoint = "cn-hangzhou.log.aliyuncs.com",
    ///                 Project = example.Name,
    ///                 Logstore = example3.Name,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Log etl can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:log/etl:Etl example tf-log-project:tf-log-etl-name
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:log/etl:Etl")]
    public partial class Etl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Delivery target logstore access key id.
        /// </summary>
        [Output("accessKeyId")]
        public Output<string?> AccessKeyId { get; private set; } = null!;

        /// <summary>
        /// Delivery target logstore access key secret.
        /// </summary>
        [Output("accessKeySecret")]
        public Output<string?> AccessKeySecret { get; private set; } = null!;

        /// <summary>
        /// The etl job create time.
        /// </summary>
        [Output("createTime")]
        public Output<int> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the log etl job.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Log service etl job alias.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the log etl job.
        /// </summary>
        [Output("etlName")]
        public Output<string> EtlName { get; private set; } = null!;

        /// <summary>
        /// Target logstore configuration for delivery after data processing.
        /// </summary>
        [Output("etlSinks")]
        public Output<ImmutableArray<Outputs.EtlEtlSink>> EtlSinks { get; private set; } = null!;

        /// <summary>
        /// Log service etl type, the default value is `ETL`.
        /// </summary>
        [Output("etlType")]
        public Output<string?> EtlType { get; private set; } = null!;

        /// <summary>
        /// The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
        /// </summary>
        [Output("fromTime")]
        public Output<int?> FromTime { get; private set; } = null!;

        /// <summary>
        /// An KMS encrypts access key id used to a log etl job. If the `access_key_id` is filled in, this field will be ignored.
        /// </summary>
        [Output("kmsEncryptedAccessKeyId")]
        public Output<string?> KmsEncryptedAccessKeyId { get; private set; } = null!;

        /// <summary>
        /// An KMS encrypts access key secret used to a log etl job. If the `access_key_secret` is filled in, this field will be ignored.
        /// </summary>
        [Output("kmsEncryptedAccessKeySecret")]
        public Output<string?> KmsEncryptedAccessKeySecret { get; private set; } = null!;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_access_key_id` before creating or updating an instance with `kms_encrypted_access_key_id`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set. When it is changed, the instance will reboot to make the change take effect.
        /// </summary>
        [Output("kmsEncryptionAccessKeyIdContext")]
        public Output<ImmutableDictionary<string, object>?> KmsEncryptionAccessKeyIdContext { get; private set; } = null!;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_access_key_secret` before creating or updating an instance with `kms_encrypted_access_key_secret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set. When it is changed, the instance will reboot to make the change take effect.
        /// </summary>
        [Output("kmsEncryptionAccessKeySecretContext")]
        public Output<ImmutableDictionary<string, object>?> KmsEncryptionAccessKeySecretContext { get; private set; } = null!;

        /// <summary>
        /// ETL job last modified time.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<int> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// Delivery target logstore.
        /// </summary>
        [Output("logstore")]
        public Output<string> Logstore { get; private set; } = null!;

        /// <summary>
        /// Advanced parameter configuration of processing operations.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// The project where the target logstore is delivered.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Sts role info under delivery target logstore. `role_arn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Job scheduling type, the default value is Resident.
        /// </summary>
        [Output("schedule")]
        public Output<string?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Processing operation grammar.
        /// </summary>
        [Output("script")]
        public Output<string> Script { get; private set; } = null!;

        /// <summary>
        /// Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
        /// </summary>
        [Output("toTime")]
        public Output<int?> ToTime { get; private set; } = null!;

        /// <summary>
        /// Log etl job version. the default value is `2`.
        /// </summary>
        [Output("version")]
        public Output<int?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Etl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Etl(string name, EtlArgs args, CustomResourceOptions? options = null)
            : base("alicloud:log/etl:Etl", name, args ?? new EtlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Etl(string name, Input<string> id, EtlState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:log/etl:Etl", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "accessKeyId",
                    "accessKeySecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Etl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Etl Get(string name, Input<string> id, EtlState? state = null, CustomResourceOptions? options = null)
        {
            return new Etl(name, id, state, options);
        }
    }

    public sealed class EtlArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessKeyId")]
        private Input<string>? _accessKeyId;

        /// <summary>
        /// Delivery target logstore access key id.
        /// </summary>
        public Input<string>? AccessKeyId
        {
            get => _accessKeyId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKeyId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("accessKeySecret")]
        private Input<string>? _accessKeySecret;

        /// <summary>
        /// Delivery target logstore access key secret.
        /// </summary>
        public Input<string>? AccessKeySecret
        {
            get => _accessKeySecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKeySecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The etl job create time.
        /// </summary>
        [Input("createTime")]
        public Input<int>? CreateTime { get; set; }

        /// <summary>
        /// Description of the log etl job.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Log service etl job alias.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The name of the log etl job.
        /// </summary>
        [Input("etlName", required: true)]
        public Input<string> EtlName { get; set; } = null!;

        [Input("etlSinks", required: true)]
        private InputList<Inputs.EtlEtlSinkArgs>? _etlSinks;

        /// <summary>
        /// Target logstore configuration for delivery after data processing.
        /// </summary>
        public InputList<Inputs.EtlEtlSinkArgs> EtlSinks
        {
            get => _etlSinks ?? (_etlSinks = new InputList<Inputs.EtlEtlSinkArgs>());
            set => _etlSinks = value;
        }

        /// <summary>
        /// Log service etl type, the default value is `ETL`.
        /// </summary>
        [Input("etlType")]
        public Input<string>? EtlType { get; set; }

        /// <summary>
        /// The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
        /// </summary>
        [Input("fromTime")]
        public Input<int>? FromTime { get; set; }

        /// <summary>
        /// An KMS encrypts access key id used to a log etl job. If the `access_key_id` is filled in, this field will be ignored.
        /// </summary>
        [Input("kmsEncryptedAccessKeyId")]
        public Input<string>? KmsEncryptedAccessKeyId { get; set; }

        /// <summary>
        /// An KMS encrypts access key secret used to a log etl job. If the `access_key_secret` is filled in, this field will be ignored.
        /// </summary>
        [Input("kmsEncryptedAccessKeySecret")]
        public Input<string>? KmsEncryptedAccessKeySecret { get; set; }

        [Input("kmsEncryptionAccessKeyIdContext")]
        private InputMap<object>? _kmsEncryptionAccessKeyIdContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_access_key_id` before creating or updating an instance with `kms_encrypted_access_key_id`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set. When it is changed, the instance will reboot to make the change take effect.
        /// </summary>
        public InputMap<object> KmsEncryptionAccessKeyIdContext
        {
            get => _kmsEncryptionAccessKeyIdContext ?? (_kmsEncryptionAccessKeyIdContext = new InputMap<object>());
            set => _kmsEncryptionAccessKeyIdContext = value;
        }

        [Input("kmsEncryptionAccessKeySecretContext")]
        private InputMap<object>? _kmsEncryptionAccessKeySecretContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_access_key_secret` before creating or updating an instance with `kms_encrypted_access_key_secret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set. When it is changed, the instance will reboot to make the change take effect.
        /// </summary>
        public InputMap<object> KmsEncryptionAccessKeySecretContext
        {
            get => _kmsEncryptionAccessKeySecretContext ?? (_kmsEncryptionAccessKeySecretContext = new InputMap<object>());
            set => _kmsEncryptionAccessKeySecretContext = value;
        }

        /// <summary>
        /// ETL job last modified time.
        /// </summary>
        [Input("lastModifiedTime")]
        public Input<int>? LastModifiedTime { get; set; }

        /// <summary>
        /// Delivery target logstore.
        /// </summary>
        [Input("logstore", required: true)]
        public Input<string> Logstore { get; set; } = null!;

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Advanced parameter configuration of processing operations.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The project where the target logstore is delivered.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Sts role info under delivery target logstore. `role_arn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Job scheduling type, the default value is Resident.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Processing operation grammar.
        /// </summary>
        [Input("script", required: true)]
        public Input<string> Script { get; set; } = null!;

        /// <summary>
        /// Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
        /// </summary>
        [Input("toTime")]
        public Input<int>? ToTime { get; set; }

        /// <summary>
        /// Log etl job version. the default value is `2`.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public EtlArgs()
        {
        }
        public static new EtlArgs Empty => new EtlArgs();
    }

    public sealed class EtlState : global::Pulumi.ResourceArgs
    {
        [Input("accessKeyId")]
        private Input<string>? _accessKeyId;

        /// <summary>
        /// Delivery target logstore access key id.
        /// </summary>
        public Input<string>? AccessKeyId
        {
            get => _accessKeyId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKeyId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("accessKeySecret")]
        private Input<string>? _accessKeySecret;

        /// <summary>
        /// Delivery target logstore access key secret.
        /// </summary>
        public Input<string>? AccessKeySecret
        {
            get => _accessKeySecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKeySecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The etl job create time.
        /// </summary>
        [Input("createTime")]
        public Input<int>? CreateTime { get; set; }

        /// <summary>
        /// Description of the log etl job.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Log service etl job alias.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the log etl job.
        /// </summary>
        [Input("etlName")]
        public Input<string>? EtlName { get; set; }

        [Input("etlSinks")]
        private InputList<Inputs.EtlEtlSinkGetArgs>? _etlSinks;

        /// <summary>
        /// Target logstore configuration for delivery after data processing.
        /// </summary>
        public InputList<Inputs.EtlEtlSinkGetArgs> EtlSinks
        {
            get => _etlSinks ?? (_etlSinks = new InputList<Inputs.EtlEtlSinkGetArgs>());
            set => _etlSinks = value;
        }

        /// <summary>
        /// Log service etl type, the default value is `ETL`.
        /// </summary>
        [Input("etlType")]
        public Input<string>? EtlType { get; set; }

        /// <summary>
        /// The start time of the processing job, if not set the value is 0, indicates to start processing from the oldest data.
        /// </summary>
        [Input("fromTime")]
        public Input<int>? FromTime { get; set; }

        /// <summary>
        /// An KMS encrypts access key id used to a log etl job. If the `access_key_id` is filled in, this field will be ignored.
        /// </summary>
        [Input("kmsEncryptedAccessKeyId")]
        public Input<string>? KmsEncryptedAccessKeyId { get; set; }

        /// <summary>
        /// An KMS encrypts access key secret used to a log etl job. If the `access_key_secret` is filled in, this field will be ignored.
        /// </summary>
        [Input("kmsEncryptedAccessKeySecret")]
        public Input<string>? KmsEncryptedAccessKeySecret { get; set; }

        [Input("kmsEncryptionAccessKeyIdContext")]
        private InputMap<object>? _kmsEncryptionAccessKeyIdContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_access_key_id` before creating or updating an instance with `kms_encrypted_access_key_id`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set. When it is changed, the instance will reboot to make the change take effect.
        /// </summary>
        public InputMap<object> KmsEncryptionAccessKeyIdContext
        {
            get => _kmsEncryptionAccessKeyIdContext ?? (_kmsEncryptionAccessKeyIdContext = new InputMap<object>());
            set => _kmsEncryptionAccessKeyIdContext = value;
        }

        [Input("kmsEncryptionAccessKeySecretContext")]
        private InputMap<object>? _kmsEncryptionAccessKeySecretContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_access_key_secret` before creating or updating an instance with `kms_encrypted_access_key_secret`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set. When it is changed, the instance will reboot to make the change take effect.
        /// </summary>
        public InputMap<object> KmsEncryptionAccessKeySecretContext
        {
            get => _kmsEncryptionAccessKeySecretContext ?? (_kmsEncryptionAccessKeySecretContext = new InputMap<object>());
            set => _kmsEncryptionAccessKeySecretContext = value;
        }

        /// <summary>
        /// ETL job last modified time.
        /// </summary>
        [Input("lastModifiedTime")]
        public Input<int>? LastModifiedTime { get; set; }

        /// <summary>
        /// Delivery target logstore.
        /// </summary>
        [Input("logstore")]
        public Input<string>? Logstore { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Advanced parameter configuration of processing operations.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// The project where the target logstore is delivered.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Sts role info under delivery target logstore. `role_arn` and `(access_key_id, access_key_secret)` fill in at most one. If you do not fill in both, then you must fill in `(kms_encrypted_access_key_id, kms_encrypted_access_key_secret, kms_encryption_access_key_id_context, kms_encryption_access_key_secret_context)` to use KMS to get the key pair.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Job scheduling type, the default value is Resident.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Processing operation grammar.
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        /// <summary>
        /// Log project tags. the default value is RUNNING, Only 4 values are supported: `STARTING`，`RUNNING`，`STOPPING`，`STOPPED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Deadline of processing job, if not set the value is 0, indicates that new data will be processed continuously.
        /// </summary>
        [Input("toTime")]
        public Input<int>? ToTime { get; set; }

        /// <summary>
        /// Log etl job version. the default value is `2`.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public EtlState()
        {
        }
        public static new EtlState Empty => new EtlState();
    }
}
