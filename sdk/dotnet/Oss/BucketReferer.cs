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
    /// Provides a OSS Bucket Referer resource. Bucket Referer configuration.
    /// 
    /// For information about OSS Bucket Referer and how to use it, see [What is Bucket Referer](https://www.alibabacloud.com/help/en/oss/user-guide/hotlink-protection).
    /// 
    /// &gt; **NOTE:** Available since v1.220.0.
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
    ///     var defaultBucketReferer = new AliCloud.Oss.BucketReferer("default", new()
    ///     {
    ///         AllowEmptyReferer = true,
    ///         RefererBlacklists = new[]
    ///         {
    ///             "*.forbidden.com",
    ///         },
    ///         Bucket = createBucket.BucketName,
    ///         TruncatePath = false,
    ///         AllowTruncateQueryString = true,
    ///         RefererLists = new[]
    ///         {
    ///             "*.aliyun.com",
    ///             "*.example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// OSS Bucket Referer can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:oss/bucketReferer:BucketReferer example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oss/bucketReferer:BucketReferer")]
    public partial class BucketReferer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to allow empty Referer request headers.
        /// </summary>
        [Output("allowEmptyReferer")]
        public Output<bool> AllowEmptyReferer { get; private set; } = null!;

        /// <summary>
        /// Allow phase request parameters.
        /// </summary>
        [Output("allowTruncateQueryString")]
        public Output<bool> AllowTruncateQueryString { get; private set; } = null!;

        /// <summary>
        /// Name of the Bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The container that holds the Referer blacklist.
        /// </summary>
        [Output("refererBlacklists")]
        public Output<ImmutableArray<string>> RefererBlacklists { get; private set; } = null!;

        /// <summary>
        /// The container that holds the Referer whitelist.
        /// </summary>
        [Output("refererLists")]
        public Output<ImmutableArray<string>> RefererLists { get; private set; } = null!;

        /// <summary>
        /// Name of the bucket.
        /// </summary>
        [Output("truncatePath")]
        public Output<bool?> TruncatePath { get; private set; } = null!;


        /// <summary>
        /// Create a BucketReferer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketReferer(string name, BucketRefererArgs args, CustomResourceOptions? options = null)
            : base("alicloud:oss/bucketReferer:BucketReferer", name, args ?? new BucketRefererArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketReferer(string name, Input<string> id, BucketRefererState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:oss/bucketReferer:BucketReferer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketReferer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketReferer Get(string name, Input<string> id, BucketRefererState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketReferer(name, id, state, options);
        }
    }

    public sealed class BucketRefererArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to allow empty Referer request headers.
        /// </summary>
        [Input("allowEmptyReferer", required: true)]
        public Input<bool> AllowEmptyReferer { get; set; } = null!;

        /// <summary>
        /// Allow phase request parameters.
        /// </summary>
        [Input("allowTruncateQueryString")]
        public Input<bool>? AllowTruncateQueryString { get; set; }

        /// <summary>
        /// Name of the Bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("refererBlacklists")]
        private InputList<string>? _refererBlacklists;

        /// <summary>
        /// The container that holds the Referer blacklist.
        /// </summary>
        public InputList<string> RefererBlacklists
        {
            get => _refererBlacklists ?? (_refererBlacklists = new InputList<string>());
            set => _refererBlacklists = value;
        }

        [Input("refererLists")]
        private InputList<string>? _refererLists;

        /// <summary>
        /// The container that holds the Referer whitelist.
        /// </summary>
        public InputList<string> RefererLists
        {
            get => _refererLists ?? (_refererLists = new InputList<string>());
            set => _refererLists = value;
        }

        /// <summary>
        /// Name of the bucket.
        /// </summary>
        [Input("truncatePath")]
        public Input<bool>? TruncatePath { get; set; }

        public BucketRefererArgs()
        {
        }
        public static new BucketRefererArgs Empty => new BucketRefererArgs();
    }

    public sealed class BucketRefererState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to allow empty Referer request headers.
        /// </summary>
        [Input("allowEmptyReferer")]
        public Input<bool>? AllowEmptyReferer { get; set; }

        /// <summary>
        /// Allow phase request parameters.
        /// </summary>
        [Input("allowTruncateQueryString")]
        public Input<bool>? AllowTruncateQueryString { get; set; }

        /// <summary>
        /// Name of the Bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("refererBlacklists")]
        private InputList<string>? _refererBlacklists;

        /// <summary>
        /// The container that holds the Referer blacklist.
        /// </summary>
        public InputList<string> RefererBlacklists
        {
            get => _refererBlacklists ?? (_refererBlacklists = new InputList<string>());
            set => _refererBlacklists = value;
        }

        [Input("refererLists")]
        private InputList<string>? _refererLists;

        /// <summary>
        /// The container that holds the Referer whitelist.
        /// </summary>
        public InputList<string> RefererLists
        {
            get => _refererLists ?? (_refererLists = new InputList<string>());
            set => _refererLists = value;
        }

        /// <summary>
        /// Name of the bucket.
        /// </summary>
        [Input("truncatePath")]
        public Input<bool>? TruncatePath { get; set; }

        public BucketRefererState()
        {
        }
        public static new BucketRefererState Empty => new BucketRefererState();
    }
}
