// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.DirectMail
{
    public static class GetTags
    {
        /// <summary>
        /// This data source provides the Direct Mail Tags of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.144.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.DirectMail.GetTags.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.DirectMail.GetTags.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Tag",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["directMailTagId1"] = ids.Apply(getTagsResult =&gt; getTagsResult.Tags[0]?.Id),
        ///         ["directMailTagId2"] = nameRegex.Apply(getTagsResult =&gt; getTagsResult.Tags[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTagsResult> InvokeAsync(GetTagsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTagsResult>("alicloud:directmail/getTags:getTags", args ?? new GetTagsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Direct Mail Tags of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.144.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.DirectMail.GetTags.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.DirectMail.GetTags.Invoke(new()
        ///     {
        ///         NameRegex = "^my-Tag",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["directMailTagId1"] = ids.Apply(getTagsResult =&gt; getTagsResult.Tags[0]?.Id),
        ///         ["directMailTagId2"] = nameRegex.Apply(getTagsResult =&gt; getTagsResult.Tags[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTagsResult> Invoke(GetTagsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagsResult>("alicloud:directmail/getTags:getTags", args ?? new GetTagsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Tag IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Tag name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetTagsArgs()
        {
        }
        public static new GetTagsArgs Empty => new GetTagsArgs();
    }

    public sealed class GetTagsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Tag IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Tag name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetTagsInvokeArgs()
        {
        }
        public static new GetTagsInvokeArgs Empty => new GetTagsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTagsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetTagsTagResult> Tags;

        [OutputConstructor]
        private GetTagsResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetTagsTagResult> tags)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Tags = tags;
        }
    }
}
