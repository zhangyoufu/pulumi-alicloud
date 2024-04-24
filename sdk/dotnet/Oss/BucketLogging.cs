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
    /// Provides a OSS Bucket Logging resource. After you enable and configure logging for a bucket, Object Storage Service (OSS) generates log objects based on a predefined naming convention. This way, access logs are generated and stored in the specified bucket on an hourly basis.
    /// 
    /// For information about OSS Bucket Logging and how to use it, see [What is Bucket Logging](https://www.alibabacloud.com/help/en/oss/developer-reference/putbucketlogging).
    /// 
    /// &gt; **NOTE:** Available since v1.222.0.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var @default = new Random.Index.Integer("default", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var createBucket = new AliCloud.Oss.Bucket("CreateBucket", new()
    ///     {
    ///         StorageClass = "Standard",
    ///         BucketName = $"{name}-{@default.Result}",
    ///     });
    /// 
    ///     var defaultBucketLogging = new AliCloud.Oss.BucketLogging("default", new()
    ///     {
    ///         Bucket = createBucket.BucketName,
    ///         TargetBucket = createBucket.BucketName,
    ///         TargetPrefix = "log/",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OSS Bucket Logging can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:oss/bucketLogging:BucketLogging example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oss/bucketLogging:BucketLogging")]
    public partial class BucketLogging : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The bucket that stores access logs.
        /// </summary>
        [Output("targetBucket")]
        public Output<string> TargetBucket { get; private set; } = null!;

        /// <summary>
        /// The prefix of the saved log objects. This element can be left empty.
        /// </summary>
        [Output("targetPrefix")]
        public Output<string?> TargetPrefix { get; private set; } = null!;


        /// <summary>
        /// Create a BucketLogging resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketLogging(string name, BucketLoggingArgs args, CustomResourceOptions? options = null)
            : base("alicloud:oss/bucketLogging:BucketLogging", name, args ?? new BucketLoggingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketLogging(string name, Input<string> id, BucketLoggingState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:oss/bucketLogging:BucketLogging", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketLogging resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketLogging Get(string name, Input<string> id, BucketLoggingState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketLogging(name, id, state, options);
        }
    }

    public sealed class BucketLoggingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The bucket that stores access logs.
        /// </summary>
        [Input("targetBucket", required: true)]
        public Input<string> TargetBucket { get; set; } = null!;

        /// <summary>
        /// The prefix of the saved log objects. This element can be left empty.
        /// </summary>
        [Input("targetPrefix")]
        public Input<string>? TargetPrefix { get; set; }

        public BucketLoggingArgs()
        {
        }
        public static new BucketLoggingArgs Empty => new BucketLoggingArgs();
    }

    public sealed class BucketLoggingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The bucket that stores access logs.
        /// </summary>
        [Input("targetBucket")]
        public Input<string>? TargetBucket { get; set; }

        /// <summary>
        /// The prefix of the saved log objects. This element can be left empty.
        /// </summary>
        [Input("targetPrefix")]
        public Input<string>? TargetPrefix { get; set; }

        public BucketLoggingState()
        {
        }
        public static new BucketLoggingState Empty => new BucketLoggingState();
    }
}
