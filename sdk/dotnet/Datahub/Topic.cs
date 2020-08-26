// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Datahub
{
    /// <summary>
    /// The topic is the basic unit of Datahub data source and is used to define one kind of data or stream. It contains a set of subscriptions. You can manage the datahub source of an application by using topics. [Refer to details](https://help.aliyun.com/document_detail/47440.html).
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// - BLob Topic
    /// resource "alicloud_datahub_topic" "example" {
    ///   name         = "tf_datahub_topic"
    ///   project_name = "tf_datahub_project"
    ///   record_type  = "BLOB"
    ///   shard_count  = 3
    ///   life_cycle   = 7
    ///   comment      = "created by terraform"
    /// }
    /// 
    /// resource "alicloud_datahub_topic" "example" {
    ///   name         = "tf_datahub_topic"
    ///   project_name = "tf_datahub_project"
    ///   record_type  = "TUPLE"
    ///   record_schema = {
    ///     bigint_field    = "BIGINT"
    ///     timestamp_field = "TIMESTAMP"
    ///     string_field    = "STRING"
    ///     double_field    = "DOUBLE"
    ///     boolean_field   = "BOOLEAN"
    ///   }
    ///   shard_count = 3
    ///   life_cycle  = 7
    ///   comment     = "created by terraform"
    /// }
    /// </summary>
    public partial class Topic : Pulumi.CustomResource
    {
        /// <summary>
        /// Comment of the datahub topic. It cannot be longer than 255 characters.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Create time of the datahub topic. It is a human-readable string rather than 64-bits UTC.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Last modify time of the datahub topic. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
        /// </summary>
        [Output("lastModifyTime")]
        public Output<string> LastModifyTime { get; private set; } = null!;

        /// <summary>
        /// How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
        /// </summary>
        [Output("lifeCycle")]
        public Output<int?> LifeCycle { get; private set; } = null!;

        /// <summary>
        /// The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the datahub project that this topic belongs to. It is case-insensitive.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
        /// - BIGINT
        /// - STRING
        /// - BOOLEAN
        /// - DOUBLE
        /// - TIMESTAMP
        /// </summary>
        [Output("recordSchema")]
        public Output<ImmutableDictionary<string, object>?> RecordSchema { get; private set; } = null!;

        /// <summary>
        /// The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.
        /// </summary>
        [Output("recordType")]
        public Output<string?> RecordType { get; private set; } = null!;

        /// <summary>
        /// The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
        /// </summary>
        [Output("shardCount")]
        public Output<int?> ShardCount { get; private set; } = null!;


        /// <summary>
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs args, CustomResourceOptions? options = null)
            : base("alicloud:datahub/topic:Topic", name, args ?? new TopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:datahub/topic:Topic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, state, options);
        }
    }

    public sealed class TopicArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment of the datahub topic. It cannot be longer than 255 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
        /// </summary>
        [Input("lifeCycle")]
        public Input<int>? LifeCycle { get; set; }

        /// <summary>
        /// The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the datahub project that this topic belongs to. It is case-insensitive.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        [Input("recordSchema")]
        private InputMap<object>? _recordSchema;

        /// <summary>
        /// Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
        /// - BIGINT
        /// - STRING
        /// - BOOLEAN
        /// - DOUBLE
        /// - TIMESTAMP
        /// </summary>
        public InputMap<object> RecordSchema
        {
            get => _recordSchema ?? (_recordSchema = new InputMap<object>());
            set => _recordSchema = value;
        }

        /// <summary>
        /// The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.
        /// </summary>
        [Input("recordType")]
        public Input<string>? RecordType { get; set; }

        /// <summary>
        /// The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        public TopicArgs()
        {
        }
    }

    public sealed class TopicState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment of the datahub topic. It cannot be longer than 255 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Create time of the datahub topic. It is a human-readable string rather than 64-bits UTC.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Last modify time of the datahub topic. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
        /// </summary>
        [Input("lastModifyTime")]
        public Input<string>? LastModifyTime { get; set; }

        /// <summary>
        /// How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
        /// </summary>
        [Input("lifeCycle")]
        public Input<int>? LifeCycle { get; set; }

        /// <summary>
        /// The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the datahub project that this topic belongs to. It is case-insensitive.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("recordSchema")]
        private InputMap<object>? _recordSchema;

        /// <summary>
        /// Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
        /// - BIGINT
        /// - STRING
        /// - BOOLEAN
        /// - DOUBLE
        /// - TIMESTAMP
        /// </summary>
        public InputMap<object> RecordSchema
        {
            get => _recordSchema ?? (_recordSchema = new InputMap<object>());
            set => _recordSchema = value;
        }

        /// <summary>
        /// The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.
        /// </summary>
        [Input("recordType")]
        public Input<string>? RecordType { get; set; }

        /// <summary>
        /// The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        public TopicState()
        {
        }
    }
}
