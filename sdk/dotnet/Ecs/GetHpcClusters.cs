// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static class GetHpcClusters
    {
        /// <summary>
        /// This data source provides the Ecs Hpc Clusters of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.116.0+.
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
        ///     var example = AliCloud.Ecs.GetHpcClusters.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "hpc-bp1i09xxxxxxxx",
        ///         },
        ///         NameRegex = "tf-testAcc",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstEcsHpcClusterId"] = example.Apply(getHpcClustersResult =&gt; getHpcClustersResult.Clusters[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetHpcClustersResult> InvokeAsync(GetHpcClustersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHpcClustersResult>("alicloud:ecs/getHpcClusters:getHpcClusters", args ?? new GetHpcClustersArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ecs Hpc Clusters of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.116.0+.
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
        ///     var example = AliCloud.Ecs.GetHpcClusters.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "hpc-bp1i09xxxxxxxx",
        ///         },
        ///         NameRegex = "tf-testAcc",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstEcsHpcClusterId"] = example.Apply(getHpcClustersResult =&gt; getHpcClustersResult.Clusters[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetHpcClustersResult> Invoke(GetHpcClustersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHpcClustersResult>("alicloud:ecs/getHpcClusters:getHpcClusters", args ?? new GetHpcClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHpcClustersArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Hpc Cluster IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Hpc Cluster name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetHpcClustersArgs()
        {
        }
        public static new GetHpcClustersArgs Empty => new GetHpcClustersArgs();
    }

    public sealed class GetHpcClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Hpc Cluster IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Hpc Cluster name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetHpcClustersInvokeArgs()
        {
        }
        public static new GetHpcClustersInvokeArgs Empty => new GetHpcClustersInvokeArgs();
    }


    [OutputType]
    public sealed class GetHpcClustersResult
    {
        public readonly ImmutableArray<Outputs.GetHpcClustersClusterResult> Clusters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetHpcClustersResult(
            ImmutableArray<Outputs.GetHpcClustersClusterResult> clusters,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            Clusters = clusters;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}
