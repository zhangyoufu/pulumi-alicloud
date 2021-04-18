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
    /// ## Example Usage
    /// 
    /// Private Bucket
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_acl = new AliCloud.Oss.Bucket("bucket-acl", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Acl = "private",
    ///             Bucket = "bucket-170309-acl",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Static Website
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_website = new AliCloud.Oss.Bucket("bucket-website", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Bucket = "bucket-170309-website",
    ///             Website = new AliCloud.Oss.Inputs.BucketWebsiteArgs
    ///             {
    ///                 ErrorDocument = "error.html",
    ///                 IndexDocument = "index.html",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Enable Logging
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_target = new AliCloud.Oss.Bucket("bucket-target", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Bucket = "bucket-170309-acl",
    ///             Acl = "public-read",
    ///         });
    ///         var bucket_logging = new AliCloud.Oss.Bucket("bucket-logging", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Bucket = "bucket-170309-logging",
    ///             Logging = new AliCloud.Oss.Inputs.BucketLoggingArgs
    ///             {
    ///                 TargetBucket = bucket_target.Id,
    ///                 TargetPrefix = "log/",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Referer configuration
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_referer = new AliCloud.Oss.Bucket("bucket-referer", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Acl = "private",
    ///             Bucket = "bucket-170309-referer",
    ///             RefererConfig = new AliCloud.Oss.Inputs.BucketRefererConfigArgs
    ///             {
    ///                 AllowEmpty = false,
    ///                 Referers = 
    ///                 {
    ///                     "http://www.aliyun.com",
    ///                     "https://www.aliyun.com",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Set lifecycle rule
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_lifecycle = new AliCloud.Oss.Bucket("bucket-lifecycle", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Acl = "public-read",
    ///             Bucket = "bucket-170309-lifecycle",
    ///             LifecycleRules = 
    ///             {
    ///                 new AliCloud.Oss.Inputs.BucketLifecycleRuleArgs
    ///                 {
    ///                     AbortMultipartUploads = 
    ///                     {
    ///                         new AliCloud.Oss.Inputs.BucketLifecycleRuleAbortMultipartUploadArgs
    ///                         {
    ///                             Days = 128,
    ///                         },
    ///                     },
    ///                     Enabled = true,
    ///                     Id = "rule-abort-multipart-upload",
    ///                     Prefix = "path3/",
    ///                 },
    ///             },
    ///         });
    ///         var bucket_versioning_lifecycle = new AliCloud.Oss.Bucket("bucket-versioning-lifecycle", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Acl = "private",
    ///             Bucket = "bucket-170309-lifecycle",
    ///             LifecycleRules = 
    ///             {
    ///                 new AliCloud.Oss.Inputs.BucketLifecycleRuleArgs
    ///                 {
    ///                     Enabled = true,
    ///                     Expirations = 
    ///                     {
    ///                         new AliCloud.Oss.Inputs.BucketLifecycleRuleExpirationArgs
    ///                         {
    ///                             ExpiredObjectDeleteMarker = true,
    ///                         },
    ///                     },
    ///                     Id = "rule-versioning",
    ///                     NoncurrentVersionExpirations = 
    ///                     {
    ///                         new AliCloud.Oss.Inputs.BucketLifecycleRuleNoncurrentVersionExpirationArgs
    ///                         {
    ///                             Days = 240,
    ///                         },
    ///                     },
    ///                     NoncurrentVersionTransitions = 
    ///                     {
    ///                         new AliCloud.Oss.Inputs.BucketLifecycleRuleNoncurrentVersionTransitionArgs
    ///                         {
    ///                             Days = 180,
    ///                             StorageClass = "Archive",
    ///                         },
    ///                         new AliCloud.Oss.Inputs.BucketLifecycleRuleNoncurrentVersionTransitionArgs
    ///                         {
    ///                             Days = 60,
    ///                             StorageClass = "IA",
    ///                         },
    ///                     },
    ///                     Prefix = "path1/",
    ///                 },
    ///             },
    ///             Versioning = new AliCloud.Oss.Inputs.BucketVersioningArgs
    ///             {
    ///                 Status = "Enabled",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Set bucket policy
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_policy = new AliCloud.Oss.Bucket("bucket-policy", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Acl = "private",
    ///             Bucket = "bucket-170309-policy",
    ///             Policy = @"  {""Statement"":
    ///       [{""Action"":
    ///           [""oss:PutObject"", ""oss:GetObject"", ""oss:DeleteBucket""],
    ///         ""Effect"":""Allow"",
    ///         ""Resource"":
    ///             [""acs:oss:*:*:*""]}],
    ///    ""Version"":""1""}
    ///   
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// IA Bucket
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_storageclass = new AliCloud.Oss.Bucket("bucket-storageclass", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Bucket = "bucket-170309-storageclass",
    ///             StorageClass = "IA",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Set bucket server-side encryption rule
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_sserule = new AliCloud.Oss.Bucket("bucket-sserule", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Acl = "private",
    ///             Bucket = "bucket-170309-sserule",
    ///             ServerSideEncryptionRule = new AliCloud.Oss.Inputs.BucketServerSideEncryptionRuleArgs
    ///             {
    ///                 KmsMasterKeyId = "your kms key id",
    ///                 SseAlgorithm = "KMS",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Set bucket tags
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_tags = new AliCloud.Oss.Bucket("bucket-tags", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Acl = "private",
    ///             Bucket = "bucket-170309-tags",
    ///             Tags = 
    ///             {
    ///                 { "key1", "value1" },
    ///                 { "key2", "value2" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Enable bucket versioning
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_versioning = new AliCloud.Oss.Bucket("bucket-versioning", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Acl = "private",
    ///             Bucket = "bucket-170309-versioning",
    ///             Versioning = new AliCloud.Oss.Inputs.BucketVersioningArgs
    ///             {
    ///                 Status = "Enabled",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Set bucket redundancy type
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucket_redundancytype = new AliCloud.Oss.Bucket("bucket-redundancytype", new AliCloud.Oss.BucketArgs
    ///         {
    ///             Bucket = "bucket_name",
    ///             RedundancyType = "ZRS",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// OSS bucket can be imported using the bucket name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:oss/bucket:Bucket bucket bucket-12345678
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oss/bucket:Bucket")]
    public partial class Bucket : Pulumi.CustomResource
    {
        /// <summary>
        /// The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".
        /// </summary>
        [Output("acl")]
        public Output<string?> Acl { get; private set; } = null!;

        [Output("bucket")]
        public Output<string?> BucketName { get; private set; } = null!;

        /// <summary>
        /// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
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
        /// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
        /// </summary>
        [Output("lifecycleRules")]
        public Output<ImmutableArray<Outputs.BucketLifecycleRule>> LifecycleRules { get; private set; } = null!;

        /// <summary>
        /// The location of the bucket.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
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
        /// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
        /// </summary>
        [Output("refererConfig")]
        public Output<Outputs.BucketRefererConfig?> RefererConfig { get; private set; } = null!;

        /// <summary>
        /// A configuration of server-side encryption (documented below).
        /// </summary>
        [Output("serverSideEncryptionRule")]
        public Output<Outputs.BucketServerSideEncryptionRule?> ServerSideEncryptionRule { get; private set; } = null!;

        /// <summary>
        /// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`.
        /// </summary>
        [Output("storageClass")]
        public Output<string?> StorageClass { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A state of versioning (documented below).
        /// </summary>
        [Output("versioning")]
        public Output<Outputs.BucketVersioning?> Versioning { get; private set; } = null!;

        /// <summary>
        /// A website object(documented below).
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

    public sealed class BucketArgs : Pulumi.ResourceArgs
    {
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
        /// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
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

        [Input("lifecycleRules")]
        private InputList<Inputs.BucketLifecycleRuleArgs>? _lifecycleRules;

        /// <summary>
        /// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
        /// </summary>
        public InputList<Inputs.BucketLifecycleRuleArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.BucketLifecycleRuleArgs>());
            set => _lifecycleRules = value;
        }

        /// <summary>
        /// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
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
        /// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
        /// </summary>
        [Input("refererConfig")]
        public Input<Inputs.BucketRefererConfigArgs>? RefererConfig { get; set; }

        /// <summary>
        /// A configuration of server-side encryption (documented below).
        /// </summary>
        [Input("serverSideEncryptionRule")]
        public Input<Inputs.BucketServerSideEncryptionRuleArgs>? ServerSideEncryptionRule { get; set; }

        /// <summary>
        /// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`.
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
        /// A state of versioning (documented below).
        /// </summary>
        [Input("versioning")]
        public Input<Inputs.BucketVersioningArgs>? Versioning { get; set; }

        /// <summary>
        /// A website object(documented below).
        /// </summary>
        [Input("website")]
        public Input<Inputs.BucketWebsiteArgs>? Website { get; set; }

        public BucketArgs()
        {
        }
    }

    public sealed class BucketState : Pulumi.ResourceArgs
    {
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
        /// A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.
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

        [Input("lifecycleRules")]
        private InputList<Inputs.BucketLifecycleRuleGetArgs>? _lifecycleRules;

        /// <summary>
        /// A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).
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
        /// A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).
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
        /// The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).
        /// </summary>
        [Input("refererConfig")]
        public Input<Inputs.BucketRefererConfigGetArgs>? RefererConfig { get; set; }

        /// <summary>
        /// A configuration of server-side encryption (documented below).
        /// </summary>
        [Input("serverSideEncryptionRule")]
        public Input<Inputs.BucketServerSideEncryptionRuleGetArgs>? ServerSideEncryptionRule { get; set; }

        /// <summary>
        /// Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: `IA`, `Archive`.
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
        /// A state of versioning (documented below).
        /// </summary>
        [Input("versioning")]
        public Input<Inputs.BucketVersioningGetArgs>? Versioning { get; set; }

        /// <summary>
        /// A website object(documented below).
        /// </summary>
        [Input("website")]
        public Input<Inputs.BucketWebsiteGetArgs>? Website { get; set; }

        public BucketState()
        {
        }
    }
}
