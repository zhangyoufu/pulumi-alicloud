// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss
{
    public static class GetBucketObjects
    {
        /// <summary>
        /// This data source provides the objects of an OSS bucket.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var bucketObjectsDs = AliCloud.Oss.GetBucketObjects.Invoke(new()
        ///     {
        ///         BucketName = "sample_bucket",
        ///         KeyRegex = "sample/sample_object.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstObjectKey"] = bucketObjectsDs.Apply(getBucketObjectsResult =&gt; getBucketObjectsResult.Objects[0]?.Key),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBucketObjectsResult> InvokeAsync(GetBucketObjectsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBucketObjectsResult>("alicloud:oss/getBucketObjects:getBucketObjects", args ?? new GetBucketObjectsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the objects of an OSS bucket.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var bucketObjectsDs = AliCloud.Oss.GetBucketObjects.Invoke(new()
        ///     {
        ///         BucketName = "sample_bucket",
        ///         KeyRegex = "sample/sample_object.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstObjectKey"] = bucketObjectsDs.Apply(getBucketObjectsResult =&gt; getBucketObjectsResult.Objects[0]?.Key),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBucketObjectsResult> Invoke(GetBucketObjectsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBucketObjectsResult>("alicloud:oss/getBucketObjects:getBucketObjects", args ?? new GetBucketObjectsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBucketObjectsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the bucket that contains the objects to find.
        /// </summary>
        [Input("bucketName", required: true)]
        public string BucketName { get; set; } = null!;

        /// <summary>
        /// Filter results by the given key prefix (such as "path/to/folder/logs-").
        /// </summary>
        [Input("keyPrefix")]
        public string? KeyPrefix { get; set; }

        /// <summary>
        /// A regex string to filter results by key.
        /// </summary>
        [Input("keyRegex")]
        public string? KeyRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetBucketObjectsArgs()
        {
        }
        public static new GetBucketObjectsArgs Empty => new GetBucketObjectsArgs();
    }

    public sealed class GetBucketObjectsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the bucket that contains the objects to find.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// Filter results by the given key prefix (such as "path/to/folder/logs-").
        /// </summary>
        [Input("keyPrefix")]
        public Input<string>? KeyPrefix { get; set; }

        /// <summary>
        /// A regex string to filter results by key.
        /// </summary>
        [Input("keyRegex")]
        public Input<string>? KeyRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetBucketObjectsInvokeArgs()
        {
        }
        public static new GetBucketObjectsInvokeArgs Empty => new GetBucketObjectsInvokeArgs();
    }


    [OutputType]
    public sealed class GetBucketObjectsResult
    {
        public readonly string BucketName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? KeyPrefix;
        public readonly string? KeyRegex;
        /// <summary>
        /// A list of bucket objects. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBucketObjectsObjectResult> Objects;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetBucketObjectsResult(
            string bucketName,

            string id,

            string? keyPrefix,

            string? keyRegex,

            ImmutableArray<Outputs.GetBucketObjectsObjectResult> objects,

            string? outputFile)
        {
            BucketName = bucketName;
            Id = id;
            KeyPrefix = keyPrefix;
            KeyRegex = keyRegex;
            Objects = objects;
            OutputFile = outputFile;
        }
    }
}
