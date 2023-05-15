// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ddos
{
    public static class GetDdosCooDomainResources
    {
        /// <summary>
        /// This data source provides the Ddoscoo Domain Resources of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.123.0+.
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
        ///     var example = AliCloud.Ddos.GetDdosCooDomainResources.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "tftestacc1234.abc",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstDdoscooDomainResourceId"] = example.Apply(getDdosCooDomainResourcesResult =&gt; getDdosCooDomainResourcesResult.Resources[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDdosCooDomainResourcesResult> InvokeAsync(GetDdosCooDomainResourcesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDdosCooDomainResourcesResult>("alicloud:ddos/getDdosCooDomainResources:getDdosCooDomainResources", args ?? new GetDdosCooDomainResourcesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ddoscoo Domain Resources of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.123.0+.
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
        ///     var example = AliCloud.Ddos.GetDdosCooDomainResources.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "tftestacc1234.abc",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstDdoscooDomainResourceId"] = example.Apply(getDdosCooDomainResourcesResult =&gt; getDdosCooDomainResourcesResult.Resources[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDdosCooDomainResourcesResult> Invoke(GetDdosCooDomainResourcesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDdosCooDomainResourcesResult>("alicloud:ddos/getDdosCooDomainResources:getDdosCooDomainResources", args ?? new GetDdosCooDomainResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDdosCooDomainResourcesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Domain Resource IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("instanceIds")]
        private List<string>? _instanceIds;

        /// <summary>
        /// A list ID of instance that you want to associate.
        /// </summary>
        public List<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new List<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Match the pattern.
        /// </summary>
        [Input("queryDomainPattern")]
        public string? QueryDomainPattern { get; set; }

        public GetDdosCooDomainResourcesArgs()
        {
        }
        public static new GetDdosCooDomainResourcesArgs Empty => new GetDdosCooDomainResourcesArgs();
    }

    public sealed class GetDdosCooDomainResourcesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Domain Resource IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// A list ID of instance that you want to associate.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Match the pattern.
        /// </summary>
        [Input("queryDomainPattern")]
        public Input<string>? QueryDomainPattern { get; set; }

        public GetDdosCooDomainResourcesInvokeArgs()
        {
        }
        public static new GetDdosCooDomainResourcesInvokeArgs Empty => new GetDdosCooDomainResourcesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDdosCooDomainResourcesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<string> InstanceIds;
        public readonly string? OutputFile;
        public readonly string? QueryDomainPattern;
        public readonly ImmutableArray<Outputs.GetDdosCooDomainResourcesResourceResult> Resources;

        [OutputConstructor]
        private GetDdosCooDomainResourcesResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<string> instanceIds,

            string? outputFile,

            string? queryDomainPattern,

            ImmutableArray<Outputs.GetDdosCooDomainResourcesResourceResult> resources)
        {
            Id = id;
            Ids = ids;
            InstanceIds = instanceIds;
            OutputFile = outputFile;
            QueryDomainPattern = queryDomainPattern;
            Resources = resources;
        }
    }
}
