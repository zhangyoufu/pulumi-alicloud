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
    /// Provides a OSS Bucket Public Access Block resource. Blocking public access at the bucket-level.
    /// 
    /// For information about OSS Bucket Public Access Block and how to use it, see [What is Bucket Public Access Block](https://www.alibabacloud.com/help/en/oss/developer-reference/putbucketpublicaccessblock).
    /// 
    /// &gt; **NOTE:** Available since v1.224.0.
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
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var createBucket = new AliCloud.Oss.Bucket("CreateBucket", new()
    ///     {
    ///         StorageClass = "Standard",
    ///         BucketName = name,
    ///     });
    /// 
    ///     var @default = new AliCloud.Oss.BucketPublicAccessBlock("default", new()
    ///     {
    ///         Bucket = createBucket.BucketName,
    ///         BlockPublicAccess = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OSS Bucket Public Access Block can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:oss/bucketPublicAccessBlock:BucketPublicAccessBlock example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oss/bucketPublicAccessBlock:BucketPublicAccessBlock")]
    public partial class BucketPublicAccessBlock : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether AlibabaCloud OSS should block public bucket policies and ACL for this bucket.
        /// </summary>
        [Output("blockPublicAccess")]
        public Output<bool> BlockPublicAccess { get; private set; } = null!;

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;


        /// <summary>
        /// Create a BucketPublicAccessBlock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketPublicAccessBlock(string name, BucketPublicAccessBlockArgs args, CustomResourceOptions? options = null)
            : base("alicloud:oss/bucketPublicAccessBlock:BucketPublicAccessBlock", name, args ?? new BucketPublicAccessBlockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketPublicAccessBlock(string name, Input<string> id, BucketPublicAccessBlockState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:oss/bucketPublicAccessBlock:BucketPublicAccessBlock", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketPublicAccessBlock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketPublicAccessBlock Get(string name, Input<string> id, BucketPublicAccessBlockState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketPublicAccessBlock(name, id, state, options);
        }
    }

    public sealed class BucketPublicAccessBlockArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether AlibabaCloud OSS should block public bucket policies and ACL for this bucket.
        /// </summary>
        [Input("blockPublicAccess", required: true)]
        public Input<bool> BlockPublicAccess { get; set; } = null!;

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        public BucketPublicAccessBlockArgs()
        {
        }
        public static new BucketPublicAccessBlockArgs Empty => new BucketPublicAccessBlockArgs();
    }

    public sealed class BucketPublicAccessBlockState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether AlibabaCloud OSS should block public bucket policies and ACL for this bucket.
        /// </summary>
        [Input("blockPublicAccess")]
        public Input<bool>? BlockPublicAccess { get; set; }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        public BucketPublicAccessBlockState()
        {
        }
        public static new BucketPublicAccessBlockState Empty => new BucketPublicAccessBlockState();
    }
}
