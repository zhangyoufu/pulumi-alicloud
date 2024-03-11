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
    /// Provides a resource to create a oss bucket and set its attribution.
    /// 
    /// &gt; **NOTE:** The bucket namespace is shared by all users of the OSS system. Please set bucket name as unique as possible.
    /// 
    /// &gt; **NOTE:** Available since v1.2.0.
    /// 
    /// ## Example Usage
    /// 
    /// Private Bucket
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
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var bucket_acl = new AliCloud.Oss.Bucket("bucket-acl", new()
    ///     {
    ///         Acl = "private",
    ///         BucketName = @default.Result.Apply(result =&gt; $"example-value-{result}"),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// Static Website
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
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var bucket_website = new AliCloud.Oss.Bucket("bucket-website", new()
    ///     {
    ///         BucketName = @default.Result.Apply(result =&gt; $"example-value-{result}"),
    ///         Website = new AliCloud.Oss.Inputs.BucketWebsiteArgs
    ///         {
    ///             ErrorDocument = "error.html",
    ///             IndexDocument = "index.html",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// Enable Logging
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
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var bucket_target = new AliCloud.Oss.Bucket("bucket-target", new()
    ///     {
    ///         BucketName = @default.Result.Apply(result =&gt; $"example-value-{result}"),
    ///         Acl = "public-read",
    ///     });
    /// 
    ///     var bucket_logging = new AliCloud.Oss.Bucket("bucket-logging", new()
    ///     {
    ///         BucketName = @default.Result.Apply(result =&gt; $"example-logging-{result}"),
    ///         Logging = new AliCloud.Oss.Inputs.BucketLoggingArgs
    ///         {
    ///             TargetBucket = bucket_target.Id,
    ///             TargetPrefix = "log/",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// Referer configuration
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
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var bucket_referer = new AliCloud.Oss.Bucket("bucket-referer", new()
    ///     {
    ///         Acl = "private",
    ///         BucketName = @default.Result.Apply(result =&gt; $"example-value-{result}"),
    ///         RefererConfig = new AliCloud.Oss.Inputs.BucketRefererConfigArgs
    ///         {
    ///             AllowEmpty = false,
    ///             Referers = new[]
    ///             {
    ///                 "http://www.aliyun.com",
    ///                 "https://www.aliyun.com",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// Set lifecycle rule
    /// 
    /// ## Import
    /// 
    /// OSS bucket can be imported using the bucket name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:oss/bucket:Bucket bucket bucket-12345678
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oss/bucket:Bucket")]
    public partial class Bucket : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A access monitor status of a bucket. See `access_monitor` below.
        /// </summary>
        [Output("accessMonitor")]
        public Output<Outputs.BucketAccessMonitor> AccessMonitor { get; private set; } = null!;

        /// <summary>
        /// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
        /// </summary>
        [Output("acl")]
        public Output<string?> Acl { get; private set; } = null!;

        [Output("bucket")]
        public Output<string?> BucketName { get; private set; } = null!;

        /// <summary>
        /// A rule of  [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm). The items of core rule are no more than 10 for every OSS bucket. See `cors_rule` below.
        /// </summary>
        [Output("corsRules")]
        public Output<ImmutableArray<Outputs.BucketCorsRule>> CorsRules { get; private set; } = null!;

        /// <summary>
        /// The creation date of the bucket.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The extranet access endpoint of the bucket.
        /// </summary>
        [Output("extranetEndpoint")]
        public Output<string> ExtranetEndpoint { get; private set; } = null!;

        /// <summary>
        /// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// The intranet access endpoint of the bucket.
        /// </summary>
        [Output("intranetEndpoint")]
        public Output<string> IntranetEndpoint { get; private set; } = null!;

        /// <summary>
        /// A boolean that indicates lifecycle rules allow prefix overlap.
        /// </summary>
        [Output("lifecycleRuleAllowSameActionOverlap")]
        public Output<bool?> LifecycleRuleAllowSameActionOverlap { get; private set; } = null!;

        /// <summary>
        /// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm). See `lifecycle_rule` below.
        /// </summary>
        [Output("lifecycleRules")]
        public Output<ImmutableArray<Outputs.BucketLifecycleRule>> LifecycleRules { get; private set; } = null!;

        /// <summary>
        /// The location of the bucket.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm). See `logging` below.
        /// </summary>
        [Output("logging")]
        public Output<Outputs.BucketLogging?> Logging { get; private set; } = null!;

        /// <summary>
        /// The flag of using logging enable container. Defaults true.
        /// </summary>
        [Output("loggingIsenable")]
        public Output<bool?> LoggingIsenable { get; private set; } = null!;

        /// <summary>
        /// The bucket owner.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
        /// </summary>
        [Output("policy")]
        public Output<string?> Policy { get; private set; } = null!;

        /// <summary>
        /// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
        /// </summary>
        [Output("redundancyType")]
        public Output<string?> RedundancyType { get; private set; } = null!;

        /// <summary>
        /// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm). See `referer_config` below.
        /// </summary>
        [Output("refererConfig")]
        public Output<Outputs.BucketRefererConfig?> RefererConfig { get; private set; } = null!;

        /// <summary>
        /// A configuration of server-side encryption. See `server_side_encryption_rule` below.
        /// </summary>
        [Output("serverSideEncryptionRule")]
        public Output<Outputs.BucketServerSideEncryptionRule?> ServerSideEncryptionRule { get; private set; } = null!;

        /// <summary>
        /// The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA", "Archive", "ColdArchive" and "DeepColdArchive". Defaults to "Standard". "ColdArchive" is available since 1.203.0. "DeepColdArchive" is available since 1.209.0.
        /// </summary>
        [Output("storageClass")]
        public Output<string?> StorageClass { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A transfer acceleration status of a bucket. See `transfer_acceleration` below.
        /// </summary>
        [Output("transferAcceleration")]
        public Output<Outputs.BucketTransferAcceleration?> TransferAcceleration { get; private set; } = null!;

        /// <summary>
        /// A state of versioning. See `versioning` below.
        /// </summary>
        [Output("versioning")]
        public Output<Outputs.BucketVersioning?> Versioning { get; private set; } = null!;

        /// <summary>
        /// A website configuration. See `website` below.
        /// </summary>
        [Output("website")]
        public Output<Outputs.BucketWebsite?> Website { get; private set; } = null!;


        /// <summary>
        /// Create a Bucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bucket(string name, BucketArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:oss/bucket:Bucket", name, args ?? new BucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Bucket(string name, Input<string> id, BucketState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:oss/bucket:Bucket", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Bucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Bucket Get(string name, Input<string> id, BucketState? state = null, CustomResourceOptions? options = null)
        {
            return new Bucket(name, id, state, options);
        }
    }

    public sealed class BucketArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A access monitor status of a bucket. See `access_monitor` below.
        /// </summary>
        [Input("accessMonitor")]
        public Input<Inputs.BucketAccessMonitorArgs>? AccessMonitor { get; set; }

        /// <summary>
        /// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        [Input("bucket")]
        public Input<string>? BucketName { get; set; }

        [Input("corsRules")]
        private InputList<Inputs.BucketCorsRuleArgs>? _corsRules;

        /// <summary>
        /// A rule of  [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm). The items of core rule are no more than 10 for every OSS bucket. See `cors_rule` below.
        /// </summary>
        public InputList<Inputs.BucketCorsRuleArgs> CorsRules
        {
            get => _corsRules ?? (_corsRules = new InputList<Inputs.BucketCorsRuleArgs>());
            set => _corsRules = value;
        }

        /// <summary>
        /// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// A boolean that indicates lifecycle rules allow prefix overlap.
        /// </summary>
        [Input("lifecycleRuleAllowSameActionOverlap")]
        public Input<bool>? LifecycleRuleAllowSameActionOverlap { get; set; }

        [Input("lifecycleRules")]
        private InputList<Inputs.BucketLifecycleRuleArgs>? _lifecycleRules;

        /// <summary>
        /// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm). See `lifecycle_rule` below.
        /// </summary>
        public InputList<Inputs.BucketLifecycleRuleArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.BucketLifecycleRuleArgs>());
            set => _lifecycleRules = value;
        }

        /// <summary>
        /// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm). See `logging` below.
        /// </summary>
        [Input("logging")]
        public Input<Inputs.BucketLoggingArgs>? Logging { get; set; }

        /// <summary>
        /// The flag of using logging enable container. Defaults true.
        /// </summary>
        [Input("loggingIsenable")]
        public Input<bool>? LoggingIsenable { get; set; }

        /// <summary>
        /// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
        /// </summary>
        [Input("redundancyType")]
        public Input<string>? RedundancyType { get; set; }

        /// <summary>
        /// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm). See `referer_config` below.
        /// </summary>
        [Input("refererConfig")]
        public Input<Inputs.BucketRefererConfigArgs>? RefererConfig { get; set; }

        /// <summary>
        /// A configuration of server-side encryption. See `server_side_encryption_rule` below.
        /// </summary>
        [Input("serverSideEncryptionRule")]
        public Input<Inputs.BucketServerSideEncryptionRuleArgs>? ServerSideEncryptionRule { get; set; }

        /// <summary>
        /// The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA", "Archive", "ColdArchive" and "DeepColdArchive". Defaults to "Standard". "ColdArchive" is available since 1.203.0. "DeepColdArchive" is available since 1.209.0.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// A transfer acceleration status of a bucket. See `transfer_acceleration` below.
        /// </summary>
        [Input("transferAcceleration")]
        public Input<Inputs.BucketTransferAccelerationArgs>? TransferAcceleration { get; set; }

        /// <summary>
        /// A state of versioning. See `versioning` below.
        /// </summary>
        [Input("versioning")]
        public Input<Inputs.BucketVersioningArgs>? Versioning { get; set; }

        /// <summary>
        /// A website configuration. See `website` below.
        /// </summary>
        [Input("website")]
        public Input<Inputs.BucketWebsiteArgs>? Website { get; set; }

        public BucketArgs()
        {
        }
        public static new BucketArgs Empty => new BucketArgs();
    }

    public sealed class BucketState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A access monitor status of a bucket. See `access_monitor` below.
        /// </summary>
        [Input("accessMonitor")]
        public Input<Inputs.BucketAccessMonitorGetArgs>? AccessMonitor { get; set; }

        /// <summary>
        /// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        [Input("bucket")]
        public Input<string>? BucketName { get; set; }

        [Input("corsRules")]
        private InputList<Inputs.BucketCorsRuleGetArgs>? _corsRules;

        /// <summary>
        /// A rule of  [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm). The items of core rule are no more than 10 for every OSS bucket. See `cors_rule` below.
        /// </summary>
        public InputList<Inputs.BucketCorsRuleGetArgs> CorsRules
        {
            get => _corsRules ?? (_corsRules = new InputList<Inputs.BucketCorsRuleGetArgs>());
            set => _corsRules = value;
        }

        /// <summary>
        /// The creation date of the bucket.
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        /// <summary>
        /// The extranet access endpoint of the bucket.
        /// </summary>
        [Input("extranetEndpoint")]
        public Input<string>? ExtranetEndpoint { get; set; }

        /// <summary>
        /// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false".
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// The intranet access endpoint of the bucket.
        /// </summary>
        [Input("intranetEndpoint")]
        public Input<string>? IntranetEndpoint { get; set; }

        /// <summary>
        /// A boolean that indicates lifecycle rules allow prefix overlap.
        /// </summary>
        [Input("lifecycleRuleAllowSameActionOverlap")]
        public Input<bool>? LifecycleRuleAllowSameActionOverlap { get; set; }

        [Input("lifecycleRules")]
        private InputList<Inputs.BucketLifecycleRuleGetArgs>? _lifecycleRules;

        /// <summary>
        /// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm). See `lifecycle_rule` below.
        /// </summary>
        public InputList<Inputs.BucketLifecycleRuleGetArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.BucketLifecycleRuleGetArgs>());
            set => _lifecycleRules = value;
        }

        /// <summary>
        /// The location of the bucket.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm). See `logging` below.
        /// </summary>
        [Input("logging")]
        public Input<Inputs.BucketLoggingGetArgs>? Logging { get; set; }

        /// <summary>
        /// The flag of using logging enable container. Defaults true.
        /// </summary>
        [Input("loggingIsenable")]
        public Input<bool>? LoggingIsenable { get; set; }

        /// <summary>
        /// The bucket owner.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The [redundancy type](https://www.alibabacloud.com/help/doc-detail/90589.htm) to enable. Can be "LRS", and "ZRS". Defaults to "LRS".
        /// </summary>
        [Input("redundancyType")]
        public Input<string>? RedundancyType { get; set; }

        /// <summary>
        /// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm). See `referer_config` below.
        /// </summary>
        [Input("refererConfig")]
        public Input<Inputs.BucketRefererConfigGetArgs>? RefererConfig { get; set; }

        /// <summary>
        /// A configuration of server-side encryption. See `server_side_encryption_rule` below.
        /// </summary>
        [Input("serverSideEncryptionRule")]
        public Input<Inputs.BucketServerSideEncryptionRuleGetArgs>? ServerSideEncryptionRule { get; set; }

        /// <summary>
        /// The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA", "Archive", "ColdArchive" and "DeepColdArchive". Defaults to "Standard". "ColdArchive" is available since 1.203.0. "DeepColdArchive" is available since 1.209.0.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// A transfer acceleration status of a bucket. See `transfer_acceleration` below.
        /// </summary>
        [Input("transferAcceleration")]
        public Input<Inputs.BucketTransferAccelerationGetArgs>? TransferAcceleration { get; set; }

        /// <summary>
        /// A state of versioning. See `versioning` below.
        /// </summary>
        [Input("versioning")]
        public Input<Inputs.BucketVersioningGetArgs>? Versioning { get; set; }

        /// <summary>
        /// A website configuration. See `website` below.
        /// </summary>
        [Input("website")]
        public Input<Inputs.BucketWebsiteGetArgs>? Website { get; set; }

        public BucketState()
        {
        }
        public static new BucketState Empty => new BucketState();
    }
}
