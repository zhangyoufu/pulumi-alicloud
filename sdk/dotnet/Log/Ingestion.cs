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
    /// Log service ingestion, this service provides the function of importing logs of various data sources(OSS, MaxCompute) into logstore.
    /// [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
    /// 
    /// &gt; **NOTE:** Available in 1.161.0+
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
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
    ///     var @default = new Random.Index.Integer("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var example = new AliCloud.Log.Project("example", new()
    ///     {
    ///         ProjectName = $"terraform-example-{@default.Result}",
    ///         Description = "terraform-example",
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "example" },
    ///         },
    ///     });
    /// 
    ///     var exampleStore = new AliCloud.Log.Store("example", new()
    ///     {
    ///         ProjectName = example.ProjectName,
    ///         LogstoreName = "example-store",
    ///         RetentionPeriod = 3650,
    ///         ShardCount = 3,
    ///         AutoSplit = true,
    ///         MaxSplitShardCount = 60,
    ///         AppendMeta = true,
    ///     });
    /// 
    ///     var exampleIngestion = new AliCloud.Log.Ingestion("example", new()
    ///     {
    ///         Project = example.ProjectName,
    ///         Logstore = exampleStore.LogstoreName,
    ///         IngestionName = "terraform-example",
    ///         DisplayName = "terraform-example",
    ///         Description = "terraform-example",
    ///         Interval = "30m",
    ///         RunImmediately = true,
    ///         TimeZone = "+0800",
    ///         Source = @"        {
    ///           ""bucket"": ""bucket_name"",
    ///           ""compressionCodec"": ""none"",
    ///           ""encoding"": ""UTF-8"",
    ///           ""endpoint"": ""oss-cn-hangzhou-internal.aliyuncs.com"",
    ///           ""format"": {
    ///             ""escapeChar"": ""\\"",
    ///             ""fieldDelimiter"": "","",
    ///             ""fieldNames"": [],
    ///             ""firstRowAsHeader"": true,
    ///             ""maxLines"": 1,
    ///             ""quoteChar"": ""\"""",
    ///             ""skipLeadingRows"": 0,
    ///             ""timeField"": """",
    ///             ""type"": ""DelimitedText""
    ///           },
    ///           ""pattern"": """",
    ///           ""prefix"": ""test-prefix/"",
    ///           ""restoreObjectEnabled"": false,
    ///           ""roleARN"": ""acs:ram::1049446484210612:role/aliyunlogimportossrole"",
    ///           ""type"": ""AliyunOSS""
    ///         }
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Log ingestion can be imported using the id or name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:log/ingestion:Ingestion example tf-log-project:tf-log-logstore:ingestion_name
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:log/ingestion:Ingestion")]
    public partial class Ingestion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Ingestion job description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name displayed on the web page.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
        /// </summary>
        [Output("ingestionName")]
        public Output<string> IngestionName { get; private set; } = null!;

        /// <summary>
        /// Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
        /// </summary>
        [Output("interval")]
        public Output<string> Interval { get; private set; } = null!;

        /// <summary>
        /// The name of the target logstore.
        /// </summary>
        [Output("logstore")]
        public Output<string> Logstore { get; private set; } = null!;

        /// <summary>
        /// The name of the log project. It is the only in one Alicloud account.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
        /// </summary>
        [Output("runImmediately")]
        public Output<bool> RunImmediately { get; private set; } = null!;

        /// <summary>
        /// Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// Which time zone is the log time imported in, e.g. `+0800`.
        /// </summary>
        [Output("timeZone")]
        public Output<string?> TimeZone { get; private set; } = null!;


        /// <summary>
        /// Create a Ingestion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ingestion(string name, IngestionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:log/ingestion:Ingestion", name, args ?? new IngestionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ingestion(string name, Input<string> id, IngestionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:log/ingestion:Ingestion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ingestion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ingestion Get(string name, Input<string> id, IngestionState? state = null, CustomResourceOptions? options = null)
        {
            return new Ingestion(name, id, state, options);
        }
    }

    public sealed class IngestionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Ingestion job description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name displayed on the web page.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
        /// </summary>
        [Input("ingestionName", required: true)]
        public Input<string> IngestionName { get; set; } = null!;

        /// <summary>
        /// Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
        /// </summary>
        [Input("interval", required: true)]
        public Input<string> Interval { get; set; } = null!;

        /// <summary>
        /// The name of the target logstore.
        /// </summary>
        [Input("logstore", required: true)]
        public Input<string> Logstore { get; set; } = null!;

        /// <summary>
        /// The name of the log project. It is the only in one Alicloud account.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
        /// </summary>
        [Input("runImmediately", required: true)]
        public Input<bool> RunImmediately { get; set; } = null!;

        /// <summary>
        /// Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        /// <summary>
        /// Which time zone is the log time imported in, e.g. `+0800`.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public IngestionArgs()
        {
        }
        public static new IngestionArgs Empty => new IngestionArgs();
    }

    public sealed class IngestionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Ingestion job description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name displayed on the web page.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Ingestion job name, it can only contain lowercase letters, numbers, dashes `-` and underscores `_`. It must start and end with lowercase letters or numbers, and the name must be 2 to 128 characters long.
        /// </summary>
        [Input("ingestionName")]
        public Input<string>? IngestionName { get; set; }

        /// <summary>
        /// Task execution interval, support minute `m`, hour `h`, day `d`, for example 30 minutes `30m`.
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        /// <summary>
        /// The name of the target logstore.
        /// </summary>
        [Input("logstore")]
        public Input<string>? Logstore { get; set; }

        /// <summary>
        /// The name of the log project. It is the only in one Alicloud account.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Whether to run the ingestion job immediately, if false, wait for an interval before starting the ingestion.
        /// </summary>
        [Input("runImmediately")]
        public Input<bool>? RunImmediately { get; set; }

        /// <summary>
        /// Data source and data format details. [Refer to details](https://www.alibabacloud.com/help/en/doc-detail/147819.html).
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Which time zone is the log time imported in, e.g. `+0800`.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public IngestionState()
        {
        }
        public static new IngestionState Empty => new IngestionState();
    }
}
