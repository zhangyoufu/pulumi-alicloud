// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static class GetImages
    {
        /// <summary>
        /// This data source provides available image resources. It contains user's private images, system images provided by Alibaba Cloud, 
        /// other public images and the ones available on the image market. 
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
        ///     var imagesDs = AliCloud.Ecs.GetImages.Invoke(new()
        ///     {
        ///         NameRegex = "^centos_6",
        ///         Owners = "system",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstImageId"] = imagesDs.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetImagesResult> InvokeAsync(GetImagesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImagesResult>("alicloud:ecs/getImages:getImages", args ?? new GetImagesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides available image resources. It contains user's private images, system images provided by Alibaba Cloud, 
        /// other public images and the ones available on the image market. 
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
        ///     var imagesDs = AliCloud.Ecs.GetImages.Invoke(new()
        ///     {
        ///         NameRegex = "^centos_6",
        ///         Owners = "system",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstImageId"] = imagesDs.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetImagesResult> Invoke(GetImagesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImagesResult>("alicloud:ecs/getImages:getImages", args ?? new GetImagesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImagesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The scenario in which the image will be used. Default value: `CreateEcs`. Valid values:
        /// </summary>
        [Input("actionType")]
        public string? ActionType { get; set; }

        /// <summary>
        /// The image architecture. Valid values: `i386` and `x86_64`.
        /// </summary>
        [Input("architecture")]
        public string? Architecture { get; set; }

        /// <summary>
        /// Specifies whether the image is running on an ECS instance. Default value: `false`. Valid values:
        /// </summary>
        [Input("dryRun")]
        public bool? DryRun { get; set; }

        /// <summary>
        /// The name of the image family. You can set this parameter to query images of the specified image family. This parameter is empty by default.
        /// </summary>
        [Input("imageFamily")]
        public string? ImageFamily { get; set; }

        /// <summary>
        /// The ID of the image.
        /// </summary>
        [Input("imageId")]
        public string? ImageId { get; set; }

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("imageName")]
        public string? ImageName { get; set; }

        /// <summary>
        /// The ID of the Alibaba Cloud account to which the image belongs. This parameter takes effect only when you query shared images or community images.
        /// </summary>
        [Input("imageOwnerId")]
        public string? ImageOwnerId { get; set; }

        /// <summary>
        /// The instance type for which the image can be used.
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

        /// <summary>
        /// Specifies whether the image supports cloud-init.
        /// </summary>
        [Input("isSupportCloudInit")]
        public bool? IsSupportCloudInit { get; set; }

        /// <summary>
        /// Specifies whether the image can be used on I/O optimized instances.
        /// </summary>
        [Input("isSupportIoOptimized")]
        public bool? IsSupportIoOptimized { get; set; }

        /// <summary>
        /// If more than one result are returned, select the most recent one.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        /// <summary>
        /// A regex string to filter resulting images by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The operating system type of the image. Valid values: `windows` and `linux`.
        /// </summary>
        [Input("osType")]
        public string? OsType { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// 
        /// &gt; **NOTE:** At least one of the `name_regex`, `most_recent` and `owners` must be set.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
        /// </summary>
        [Input("owners")]
        public string? Owners { get; set; }

        /// <summary>
        /// The ID of the resource group to which the custom image belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        /// <summary>
        /// The ID of the snapshot used to create the custom image.
        /// </summary>
        [Input("snapshotId")]
        public string? SnapshotId { get; set; }

        /// <summary>
        /// The status of the image. The following values are available, Separate multiple parameter values by using commas (,). Default value: `Available`. Valid values:
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies whether to check the validity of the request without actually making the request. Valid values:
        /// </summary>
        [Input("usage")]
        public string? Usage { get; set; }

        public GetImagesArgs()
        {
        }
        public static new GetImagesArgs Empty => new GetImagesArgs();
    }

    public sealed class GetImagesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The scenario in which the image will be used. Default value: `CreateEcs`. Valid values:
        /// </summary>
        [Input("actionType")]
        public Input<string>? ActionType { get; set; }

        /// <summary>
        /// The image architecture. Valid values: `i386` and `x86_64`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// Specifies whether the image is running on an ECS instance. Default value: `false`. Valid values:
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The name of the image family. You can set this parameter to query images of the specified image family. This parameter is empty by default.
        /// </summary>
        [Input("imageFamily")]
        public Input<string>? ImageFamily { get; set; }

        /// <summary>
        /// The ID of the image.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// The ID of the Alibaba Cloud account to which the image belongs. This parameter takes effect only when you query shared images or community images.
        /// </summary>
        [Input("imageOwnerId")]
        public Input<string>? ImageOwnerId { get; set; }

        /// <summary>
        /// The instance type for which the image can be used.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Specifies whether the image supports cloud-init.
        /// </summary>
        [Input("isSupportCloudInit")]
        public Input<bool>? IsSupportCloudInit { get; set; }

        /// <summary>
        /// Specifies whether the image can be used on I/O optimized instances.
        /// </summary>
        [Input("isSupportIoOptimized")]
        public Input<bool>? IsSupportIoOptimized { get; set; }

        /// <summary>
        /// If more than one result are returned, select the most recent one.
        /// </summary>
        [Input("mostRecent")]
        public Input<bool>? MostRecent { get; set; }

        /// <summary>
        /// A regex string to filter resulting images by name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// The operating system type of the image. Valid values: `windows` and `linux`.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// 
        /// &gt; **NOTE:** At least one of the `name_regex`, `most_recent` and `owners` must be set.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
        /// </summary>
        [Input("owners")]
        public Input<string>? Owners { get; set; }

        /// <summary>
        /// The ID of the resource group to which the custom image belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The ID of the snapshot used to create the custom image.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// The status of the image. The following values are available, Separate multiple parameter values by using commas (,). Default value: `Available`. Valid values:
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies whether to check the validity of the request without actually making the request. Valid values:
        /// </summary>
        [Input("usage")]
        public Input<string>? Usage { get; set; }

        public GetImagesInvokeArgs()
        {
        }
        public static new GetImagesInvokeArgs Empty => new GetImagesInvokeArgs();
    }


    [OutputType]
    public sealed class GetImagesResult
    {
        public readonly string? ActionType;
        /// <summary>
        /// Platform type of the image system: i386 or x86_64.
        /// </summary>
        public readonly string? Architecture;
        public readonly bool? DryRun;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of image IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? ImageFamily;
        public readonly string? ImageId;
        public readonly string? ImageName;
        public readonly string? ImageOwnerId;
        /// <summary>
        /// A list of images. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetImagesImageResult> Images;
        public readonly string? InstanceType;
        public readonly bool? IsSupportCloudInit;
        public readonly bool? IsSupportIoOptimized;
        public readonly bool? MostRecent;
        public readonly string? NameRegex;
        public readonly string? OsType;
        public readonly string? OutputFile;
        public readonly string? Owners;
        public readonly string? ResourceGroupId;
        /// <summary>
        /// Snapshot ID.
        /// </summary>
        public readonly string? SnapshotId;
        /// <summary>
        /// Status of the image. Possible values: `UnAvailable`, `Available`, `Creating` and `CreateFailed`.
        /// </summary>
        public readonly string? Status;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? Usage;

        [OutputConstructor]
        private GetImagesResult(
            string? actionType,

            string? architecture,

            bool? dryRun,

            string id,

            ImmutableArray<string> ids,

            string? imageFamily,

            string? imageId,

            string? imageName,

            string? imageOwnerId,

            ImmutableArray<Outputs.GetImagesImageResult> images,

            string? instanceType,

            bool? isSupportCloudInit,

            bool? isSupportIoOptimized,

            bool? mostRecent,

            string? nameRegex,

            string? osType,

            string? outputFile,

            string? owners,

            string? resourceGroupId,

            string? snapshotId,

            string? status,

            ImmutableDictionary<string, object>? tags,

            string? usage)
        {
            ActionType = actionType;
            Architecture = architecture;
            DryRun = dryRun;
            Id = id;
            Ids = ids;
            ImageFamily = imageFamily;
            ImageId = imageId;
            ImageName = imageName;
            ImageOwnerId = imageOwnerId;
            Images = images;
            InstanceType = instanceType;
            IsSupportCloudInit = isSupportCloudInit;
            IsSupportIoOptimized = isSupportIoOptimized;
            MostRecent = mostRecent;
            NameRegex = nameRegex;
            OsType = osType;
            OutputFile = outputFile;
            Owners = owners;
            ResourceGroupId = resourceGroupId;
            SnapshotId = snapshotId;
            Status = status;
            Tags = tags;
            Usage = usage;
        }
    }
}
