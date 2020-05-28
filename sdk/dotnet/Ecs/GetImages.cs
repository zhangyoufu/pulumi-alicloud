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
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var imagesDs = Output.Create(AliCloud.Ecs.GetImages.InvokeAsync(new AliCloud.Ecs.GetImagesArgs
        ///         {
        ///             NameRegex = "^centos_6",
        ///             Owners = "system",
        ///         }));
        ///         this.FirstImageId = imagesDs.Apply(imagesDs =&gt; imagesDs.Images[0].Id);
        ///     }
        /// 
        ///     [Output("firstImageId")]
        ///     public Output&lt;string&gt; FirstImageId { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetImagesResult> InvokeAsync(GetImagesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImagesResult>("alicloud:ecs/getImages:getImages", args ?? new GetImagesArgs(), options.WithVersion());
    }


    public sealed class GetImagesArgs : Pulumi.InvokeArgs
    {
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

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
        /// </summary>
        [Input("owners")]
        public string? Owners { get; set; }

        public GetImagesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetImagesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of image IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of images. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetImagesImageResult> Images;
        public readonly bool? MostRecent;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly string? Owners;

        [OutputConstructor]
        private GetImagesResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetImagesImageResult> images,

            bool? mostRecent,

            string? nameRegex,

            string? outputFile,

            string? owners)
        {
            Id = id;
            Ids = ids;
            Images = images;
            MostRecent = mostRecent;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            Owners = owners;
        }
    }
}
