// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Export a custom image to the OSS bucket in the same region as the custom image.
    /// 
    /// &gt; **NOTE:** If you create an ECS instance using a mirror image and create a system disk snapshot again, exporting a custom image created from the system disk snapshot is not supported.
    /// 
    /// &gt; **NOTE:** Support for exporting custom images that include data disk snapshot information in the image. The number of data disks cannot exceed 4 and the maximum capacity of a single data disk cannot exceed 500 GiB.
    /// 
    /// &gt; **NOTE:** Before exporting the image, you must authorize the cloud server ECS official service account to write OSS permissions through RAM.
    /// 
    /// &gt; **NOTE:** Available in 1.68.0+.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new AliCloud.Ecs.ImageExport("default", new AliCloud.Ecs.ImageExportArgs
    ///         {
    ///             ImageId = "m-bp1gxy***",
    ///             OssBucket = "ecsimageexportconfig",
    ///             OssPrefix = "ecsExport",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Attributes Reference0
    /// 
    ///  The following attributes are exported:
    /// 
    /// * `id` - ID of the image.
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/imageExport:ImageExport")]
    public partial class ImageExport : Pulumi.CustomResource
    {
        /// <summary>
        /// The source image ID.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        /// <summary>
        /// Save the exported OSS bucket.
        /// </summary>
        [Output("ossBucket")]
        public Output<string> OssBucket { get; private set; } = null!;

        /// <summary>
        /// The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        /// </summary>
        [Output("ossPrefix")]
        public Output<string?> OssPrefix { get; private set; } = null!;


        /// <summary>
        /// Create a ImageExport resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageExport(string name, ImageExportArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/imageExport:ImageExport", name, args ?? new ImageExportArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImageExport(string name, Input<string> id, ImageExportState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/imageExport:ImageExport", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ImageExport resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageExport Get(string name, Input<string> id, ImageExportState? state = null, CustomResourceOptions? options = null)
        {
            return new ImageExport(name, id, state, options);
        }
    }

    public sealed class ImageExportArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source image ID.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        /// <summary>
        /// Save the exported OSS bucket.
        /// </summary>
        [Input("ossBucket", required: true)]
        public Input<string> OssBucket { get; set; } = null!;

        /// <summary>
        /// The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        /// </summary>
        [Input("ossPrefix")]
        public Input<string>? OssPrefix { get; set; }

        public ImageExportArgs()
        {
        }
    }

    public sealed class ImageExportState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source image ID.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Save the exported OSS bucket.
        /// </summary>
        [Input("ossBucket")]
        public Input<string>? OssBucket { get; set; }

        /// <summary>
        /// The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        /// </summary>
        [Input("ossPrefix")]
        public Input<string>? OssPrefix { get; set; }

        public ImageExportState()
        {
        }
    }
}
