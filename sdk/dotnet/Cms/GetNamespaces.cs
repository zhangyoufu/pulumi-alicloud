// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms
{
    public static class GetNamespaces
    {
        /// <summary>
        /// This data source provides the Cms Namespaces of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.171.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var ids = AliCloud.Cms.GetNamespaces.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cmsNamespaceId1"] = ids.Apply(getNamespacesResult =&gt; getNamespacesResult.Namespaces[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNamespacesResult> InvokeAsync(GetNamespacesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNamespacesResult>("alicloud:cms/getNamespaces:getNamespaces", args ?? new GetNamespacesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Cms Namespaces of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.171.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var ids = AliCloud.Cms.GetNamespaces.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cmsNamespaceId1"] = ids.Apply(getNamespacesResult =&gt; getNamespacesResult.Namespaces[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNamespacesResult> Invoke(GetNamespacesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNamespacesResult>("alicloud:cms/getNamespaces:getNamespaces", args ?? new GetNamespacesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNamespacesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Namespace IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The keywords of the `namespace` or `description` of the namespace.
        /// </summary>
        [Input("keyword")]
        public string? Keyword { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        public GetNamespacesArgs()
        {
        }
        public static new GetNamespacesArgs Empty => new GetNamespacesArgs();
    }

    public sealed class GetNamespacesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Namespace IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The keywords of the `namespace` or `description` of the namespace.
        /// </summary>
        [Input("keyword")]
        public Input<string>? Keyword { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        public GetNamespacesInvokeArgs()
        {
        }
        public static new GetNamespacesInvokeArgs Empty => new GetNamespacesInvokeArgs();
    }


    [OutputType]
    public sealed class GetNamespacesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? Keyword;
        public readonly ImmutableArray<Outputs.GetNamespacesNamespaceResult> Namespaces;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;

        [OutputConstructor]
        private GetNamespacesResult(
            string id,

            ImmutableArray<string> ids,

            string? keyword,

            ImmutableArray<Outputs.GetNamespacesNamespaceResult> namespaces,

            string? outputFile,

            int? pageNumber,

            int? pageSize)
        {
            Id = id;
            Ids = ids;
            Keyword = keyword;
            Namespaces = namespaces;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
        }
    }
}
