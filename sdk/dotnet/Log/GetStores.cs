// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log
{
    public static class GetStores
    {
        /// <summary>
        /// This data source provides the Log Stores of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.126.0+.
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
        ///     var example = AliCloud.Log.GetStores.Invoke(new()
        ///     {
        ///         Project = "the_project_name",
        ///         Ids = new[]
        ///         {
        ///             "the_store_name",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstLogStoreId"] = example.Apply(getStoresResult =&gt; getStoresResult.Stores[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetStoresResult> InvokeAsync(GetStoresArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStoresResult>("alicloud:log/getStores:getStores", args ?? new GetStoresArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Log Stores of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.126.0+.
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
        ///     var example = AliCloud.Log.GetStores.Invoke(new()
        ///     {
        ///         Project = "the_project_name",
        ///         Ids = new[]
        ///         {
        ///             "the_store_name",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstLogStoreId"] = example.Apply(getStoresResult =&gt; getStoresResult.Stores[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetStoresResult> Invoke(GetStoresInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStoresResult>("alicloud:log/getStores:getStores", args ?? new GetStoresInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStoresArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of store IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by store name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetStoresArgs()
        {
        }
        public static new GetStoresArgs Empty => new GetStoresArgs();
    }

    public sealed class GetStoresInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of store IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by store name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public GetStoresInvokeArgs()
        {
        }
        public static new GetStoresInvokeArgs Empty => new GetStoresInvokeArgs();
    }


    [OutputType]
    public sealed class GetStoresResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string Project;
        public readonly ImmutableArray<Outputs.GetStoresStoreResult> Stores;

        [OutputConstructor]
        private GetStoresResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string project,

            ImmutableArray<Outputs.GetStoresStoreResult> stores)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Project = project;
            Stores = stores;
        }
    }
}
