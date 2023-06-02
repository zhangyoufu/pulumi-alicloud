// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss
{
    /// <summary>
    /// Provides an independent replication configuration resource for OSS bucket.
    /// 
    /// For information about OSS replication and how to use it, see [What is cross-region replication](https://www.alibabacloud.com/help/doc-detail/31864.html) and [What is same-region replication](https://www.alibabacloud.com/help/doc-detail/254865.html).
    /// 
    /// &gt; **NOTE:** Available in v1.161.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Set bucket replication configuration
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var bucketSrc = new AliCloud.Oss.Bucket("bucketSrc", new()
    ///     {
    ///         BucketName = @default.Result.Apply(result =&gt; $"example-src-{result}"),
    ///     });
    /// 
    ///     var bucketDest = new AliCloud.Oss.Bucket("bucketDest", new()
    ///     {
    ///         BucketName = @default.Result.Apply(result =&gt; $"example-dest-{result}"),
    ///     });
    /// 
    ///     var role = new AliCloud.Ram.Role("role", new()
    ///     {
    ///         Document = @"		{
    /// 		  ""Statement"": [
    /// 			{
    /// 			  ""Action"": ""sts:AssumeRole"",
    /// 			  ""Effect"": ""Allow"",
    /// 			  ""Principal"": {
    /// 				""Service"": [
    /// 				  ""oss.aliyuncs.com""
    /// 				]
    /// 			  }
    /// 			}
    /// 		  ],
    /// 		  ""Version"": ""1""
    /// 		}
    /// ",
    ///         Description = "this is a test",
    ///         Force = true,
    ///     });
    /// 
    ///     var policy = new AliCloud.Ram.Policy("policy", new()
    ///     {
    ///         PolicyName = @default.Result.Apply(result =&gt; $"example-policy-{result}"),
    ///         PolicyDocument = @"		{
    /// 		  ""Statement"": [
    /// 			{
    /// 			  ""Action"": [
    /// 				""*""
    /// 			  ],
    /// 			  ""Effect"": ""Allow"",
    /// 			  ""Resource"": [
    /// 				""*""
    /// 			  ]
    /// 			}
    /// 		  ],
    /// 			""Version"": ""1""
    /// 		}
    /// ",
    ///         Description = "this is a policy test",
    ///         Force = true,
    ///     });
    /// 
    ///     var attach = new AliCloud.Ram.RolePolicyAttachment("attach", new()
    ///     {
    ///         PolicyName = policy.Name,
    ///         PolicyType = policy.Type,
    ///         RoleName = role.Name,
    ///     });
    /// 
    ///     var key = new AliCloud.Kms.Key("key", new()
    ///     {
    ///         Description = "Hello KMS",
    ///         PendingWindowInDays = 7,
    ///         Status = "Enabled",
    ///     });
    /// 
    ///     var cross_region_replication = new AliCloud.Oss.BucketReplication("cross-region-replication", new()
    ///     {
    ///         Bucket = bucketSrc.Id,
    ///         Action = "PUT,DELETE",
    ///         HistoricalObjectReplication = "enabled",
    ///         PrefixSet = new AliCloud.Oss.Inputs.BucketReplicationPrefixSetArgs
    ///         {
    ///             Prefixes = new[]
    ///             {
    ///                 "prefix1/",
    ///                 "prefix2/",
    ///             },
    ///         },
    ///         Destination = new AliCloud.Oss.Inputs.BucketReplicationDestinationArgs
    ///         {
    ///             Bucket = bucketDest.Id,
    ///             Location = bucketDest.Location,
    ///         },
    ///         SyncRole = role.Name,
    ///         EncryptionConfiguration = new AliCloud.Oss.Inputs.BucketReplicationEncryptionConfigurationArgs
    ///         {
    ///             ReplicaKmsKeyId = key.Id,
    ///         },
    ///         SourceSelectionCriteria = new AliCloud.Oss.Inputs.BucketReplicationSourceSelectionCriteriaArgs
    ///         {
    ///             SseKmsEncryptedObjects = new AliCloud.Oss.Inputs.BucketReplicationSourceSelectionCriteriaSseKmsEncryptedObjectsArgs
    ///             {
    ///                 Status = "Enabled",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ### Timeouts The `timeouts` block allows you to specify timeouts for certain actions* `delete` - (Defaults to 30 mins) Used when delete a data replication rule (until the data replication task is cleared).
    /// </summary>
    [AliCloudResourceType("alicloud:oss/bucketReplication:BucketReplication")]
    public partial class BucketReplication : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Specifies the destination for the rule(See the following block `destination`).
        /// </summary>
        [Output("destination")]
        public Output<Outputs.BucketReplicationDestination> Destination { get; private set; } = null!;

        /// <summary>
        /// Specifies the encryption configuration for the objects replicated to the destination bucket(See the following block `encryption_configuration`).
        /// </summary>
        [Output("encryptionConfiguration")]
        public Output<Outputs.BucketReplicationEncryptionConfiguration?> EncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
        /// </summary>
        [Output("historicalObjectReplication")]
        public Output<string?> HistoricalObjectReplication { get; private set; } = null!;

        /// <summary>
        /// The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket(See the following block `prefix_set`).
        /// </summary>
        [Output("prefixSet")]
        public Output<Outputs.BucketReplicationPrefixSet?> PrefixSet { get; private set; } = null!;

        /// <summary>
        /// Retrieves the progress of the data replication task. This status is returned only when the data replication task is in the doing state.
        /// </summary>
        [Output("progress")]
        public Output<Outputs.BucketReplicationProgress> Progress { get; private set; } = null!;

        /// <summary>
        /// The ID of the data replication rule.
        /// </summary>
        [Output("ruleId")]
        public Output<string> RuleId { get; private set; } = null!;

        /// <summary>
        /// Specifies other conditions used to filter the source objects to replicate(See the following block `source_selection_criteria`).
        /// </summary>
        [Output("sourceSelectionCriteria")]
        public Output<Outputs.BucketReplicationSourceSelectionCriteria?> SourceSelectionCriteria { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to replicate objects encrypted by using SSE-KMS. Can be `Enabled` or `Disabled`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
        /// </summary>
        [Output("syncRole")]
        public Output<string?> SyncRole { get; private set; } = null!;


        /// <summary>
        /// Create a BucketReplication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketReplication(string name, BucketReplicationArgs args, CustomResourceOptions? options = null)
            : base("alicloud:oss/bucketReplication:BucketReplication", name, args ?? new BucketReplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketReplication(string name, Input<string> id, BucketReplicationState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:oss/bucketReplication:BucketReplication", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BucketReplication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketReplication Get(string name, Input<string> id, BucketReplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketReplication(name, id, state, options);
        }
    }

    public sealed class BucketReplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Specifies the destination for the rule(See the following block `destination`).
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.BucketReplicationDestinationArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Specifies the encryption configuration for the objects replicated to the destination bucket(See the following block `encryption_configuration`).
        /// </summary>
        [Input("encryptionConfiguration")]
        public Input<Inputs.BucketReplicationEncryptionConfigurationArgs>? EncryptionConfiguration { get; set; }

        /// <summary>
        /// Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
        /// </summary>
        [Input("historicalObjectReplication")]
        public Input<string>? HistoricalObjectReplication { get; set; }

        /// <summary>
        /// The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket(See the following block `prefix_set`).
        /// </summary>
        [Input("prefixSet")]
        public Input<Inputs.BucketReplicationPrefixSetArgs>? PrefixSet { get; set; }

        /// <summary>
        /// Retrieves the progress of the data replication task. This status is returned only when the data replication task is in the doing state.
        /// </summary>
        [Input("progress")]
        public Input<Inputs.BucketReplicationProgressArgs>? Progress { get; set; }

        /// <summary>
        /// Specifies other conditions used to filter the source objects to replicate(See the following block `source_selection_criteria`).
        /// </summary>
        [Input("sourceSelectionCriteria")]
        public Input<Inputs.BucketReplicationSourceSelectionCriteriaArgs>? SourceSelectionCriteria { get; set; }

        /// <summary>
        /// Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
        /// </summary>
        [Input("syncRole")]
        public Input<string>? SyncRole { get; set; }

        public BucketReplicationArgs()
        {
        }
        public static new BucketReplicationArgs Empty => new BucketReplicationArgs();
    }

    public sealed class BucketReplicationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operations that can be synchronized to the destination bucket. You can set action to one or more of the following operation types. Valid values: `ALL`(contains PUT, DELETE, and ABORT), `PUT`, `DELETE` and `ABORT`. Defaults to `ALL`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Specifies the destination for the rule(See the following block `destination`).
        /// </summary>
        [Input("destination")]
        public Input<Inputs.BucketReplicationDestinationGetArgs>? Destination { get; set; }

        /// <summary>
        /// Specifies the encryption configuration for the objects replicated to the destination bucket(See the following block `encryption_configuration`).
        /// </summary>
        [Input("encryptionConfiguration")]
        public Input<Inputs.BucketReplicationEncryptionConfigurationGetArgs>? EncryptionConfiguration { get; set; }

        /// <summary>
        /// Specifies whether to replicate historical data from the source bucket to the destination bucket before data replication is enabled. Can be `enabled` or `disabled`. Defaults to `enabled`.
        /// </summary>
        [Input("historicalObjectReplication")]
        public Input<string>? HistoricalObjectReplication { get; set; }

        /// <summary>
        /// The prefixes used to specify the object to replicate. Only objects that match the prefix are replicated to the destination bucket(See the following block `prefix_set`).
        /// </summary>
        [Input("prefixSet")]
        public Input<Inputs.BucketReplicationPrefixSetGetArgs>? PrefixSet { get; set; }

        /// <summary>
        /// Retrieves the progress of the data replication task. This status is returned only when the data replication task is in the doing state.
        /// </summary>
        [Input("progress")]
        public Input<Inputs.BucketReplicationProgressGetArgs>? Progress { get; set; }

        /// <summary>
        /// The ID of the data replication rule.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// Specifies other conditions used to filter the source objects to replicate(See the following block `source_selection_criteria`).
        /// </summary>
        [Input("sourceSelectionCriteria")]
        public Input<Inputs.BucketReplicationSourceSelectionCriteriaGetArgs>? SourceSelectionCriteria { get; set; }

        /// <summary>
        /// Specifies whether to replicate objects encrypted by using SSE-KMS. Can be `Enabled` or `Disabled`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the role that you authorize OSS to use to replicate data. If SSE-KMS is specified to encrypt the objects replicated to the destination bucket, it must be specified.
        /// </summary>
        [Input("syncRole")]
        public Input<string>? SyncRole { get; set; }

        public BucketReplicationState()
        {
        }
        public static new BucketReplicationState Empty => new BucketReplicationState();
    }
}
