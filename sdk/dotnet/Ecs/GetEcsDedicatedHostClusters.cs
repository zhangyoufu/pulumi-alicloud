// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static class GetEcsDedicatedHostClusters
    {
        /// <summary>
        /// This data source provides the Ecs Dedicated Host Clusters of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.146.0+.
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
        ///     var ids = AliCloud.Ecs.GetEcsDedicatedHostClusters.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Ecs.GetEcsDedicatedHostClusters.Invoke(new()
        ///     {
        ///         NameRegex = "^my-DedicatedHostCluster",
        ///     });
        /// 
        ///     var zoneId = AliCloud.Ecs.GetEcsDedicatedHostClusters.Invoke(new()
        ///     {
        ///         ZoneId = "example_value",
        ///     });
        /// 
        ///     var clusterName = AliCloud.Ecs.GetEcsDedicatedHostClusters.Invoke(new()
        ///     {
        ///         DedicatedHostClusterName = "example_value",
        ///     });
        /// 
        ///     var clusterIds = AliCloud.Ecs.GetEcsDedicatedHostClusters.Invoke(new()
        ///     {
        ///         DedicatedHostClusterIds = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ecsDedicatedHostClusterId1"] = ids.Apply(getEcsDedicatedHostClustersResult =&gt; getEcsDedicatedHostClustersResult.Clusters[0]?.Id),
        ///         ["ecsDedicatedHostClusterId2"] = nameRegex.Apply(getEcsDedicatedHostClustersResult =&gt; getEcsDedicatedHostClustersResult.Clusters[0]?.Id),
        ///         ["ecsDedicatedHostClusterId3"] = zoneId.Apply(getEcsDedicatedHostClustersResult =&gt; getEcsDedicatedHostClustersResult.Clusters[0]?.Id),
        ///         ["ecsDedicatedHostClusterId4"] = clusterName.Apply(getEcsDedicatedHostClustersResult =&gt; getEcsDedicatedHostClustersResult.Clusters[0]?.Id),
        ///         ["ecsDedicatedHostClusterId5"] = clusterIds.Apply(getEcsDedicatedHostClustersResult =&gt; getEcsDedicatedHostClustersResult.Clusters[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEcsDedicatedHostClustersResult> InvokeAsync(GetEcsDedicatedHostClustersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEcsDedicatedHostClustersResult>("alicloud:ecs/getEcsDedicatedHostClusters:getEcsDedicatedHostClusters", args ?? new GetEcsDedicatedHostClustersArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ecs Dedicated Host Clusters of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.146.0+.
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
        ///     var ids = AliCloud.Ecs.GetEcsDedicatedHostClusters.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Ecs.GetEcsDedicatedHostClusters.Invoke(new()
        ///     {
        ///         NameRegex = "^my-DedicatedHostCluster",
        ///     });
        /// 
        ///     var zoneId = AliCloud.Ecs.GetEcsDedicatedHostClusters.Invoke(new()
        ///     {
        ///         ZoneId = "example_value",
        ///     });
        /// 
        ///     var clusterName = AliCloud.Ecs.GetEcsDedicatedHostClusters.Invoke(new()
        ///     {
        ///         DedicatedHostClusterName = "example_value",
        ///     });
        /// 
        ///     var clusterIds = AliCloud.Ecs.GetEcsDedicatedHostClusters.Invoke(new()
        ///     {
        ///         DedicatedHostClusterIds = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ecsDedicatedHostClusterId1"] = ids.Apply(getEcsDedicatedHostClustersResult =&gt; getEcsDedicatedHostClustersResult.Clusters[0]?.Id),
        ///         ["ecsDedicatedHostClusterId2"] = nameRegex.Apply(getEcsDedicatedHostClustersResult =&gt; getEcsDedicatedHostClustersResult.Clusters[0]?.Id),
        ///         ["ecsDedicatedHostClusterId3"] = zoneId.Apply(getEcsDedicatedHostClustersResult =&gt; getEcsDedicatedHostClustersResult.Clusters[0]?.Id),
        ///         ["ecsDedicatedHostClusterId4"] = clusterName.Apply(getEcsDedicatedHostClustersResult =&gt; getEcsDedicatedHostClustersResult.Clusters[0]?.Id),
        ///         ["ecsDedicatedHostClusterId5"] = clusterIds.Apply(getEcsDedicatedHostClustersResult =&gt; getEcsDedicatedHostClustersResult.Clusters[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEcsDedicatedHostClustersResult> Invoke(GetEcsDedicatedHostClustersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEcsDedicatedHostClustersResult>("alicloud:ecs/getEcsDedicatedHostClusters:getEcsDedicatedHostClusters", args ?? new GetEcsDedicatedHostClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEcsDedicatedHostClustersArgs : global::Pulumi.InvokeArgs
    {
        [Input("dedicatedHostClusterIds")]
        private List<string>? _dedicatedHostClusterIds;

        /// <summary>
        /// The IDs of dedicated host clusters.
        /// </summary>
        public List<string> DedicatedHostClusterIds
        {
            get => _dedicatedHostClusterIds ?? (_dedicatedHostClusterIds = new List<string>());
            set => _dedicatedHostClusterIds = value;
        }

        /// <summary>
        /// The name of the dedicated host cluster.
        /// </summary>
        [Input("dedicatedHostClusterName")]
        public string? DedicatedHostClusterName { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Dedicated Host Cluster IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Dedicated Host Cluster name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

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
        /// The zone ID of the dedicated host cluster.
        /// </summary>
        [Input("zoneId")]
        public string? ZoneId { get; set; }

        public GetEcsDedicatedHostClustersArgs()
        {
        }
        public static new GetEcsDedicatedHostClustersArgs Empty => new GetEcsDedicatedHostClustersArgs();
    }

    public sealed class GetEcsDedicatedHostClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dedicatedHostClusterIds")]
        private InputList<string>? _dedicatedHostClusterIds;

        /// <summary>
        /// The IDs of dedicated host clusters.
        /// </summary>
        public InputList<string> DedicatedHostClusterIds
        {
            get => _dedicatedHostClusterIds ?? (_dedicatedHostClusterIds = new InputList<string>());
            set => _dedicatedHostClusterIds = value;
        }

        /// <summary>
        /// The name of the dedicated host cluster.
        /// </summary>
        [Input("dedicatedHostClusterName")]
        public Input<string>? DedicatedHostClusterName { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Dedicated Host Cluster IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Dedicated Host Cluster name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

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
        /// The zone ID of the dedicated host cluster.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public GetEcsDedicatedHostClustersInvokeArgs()
        {
        }
        public static new GetEcsDedicatedHostClustersInvokeArgs Empty => new GetEcsDedicatedHostClustersInvokeArgs();
    }


    [OutputType]
    public sealed class GetEcsDedicatedHostClustersResult
    {
        public readonly ImmutableArray<Outputs.GetEcsDedicatedHostClustersClusterResult> Clusters;
        public readonly ImmutableArray<string> DedicatedHostClusterIds;
        public readonly string? DedicatedHostClusterName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? ZoneId;

        [OutputConstructor]
        private GetEcsDedicatedHostClustersResult(
            ImmutableArray<Outputs.GetEcsDedicatedHostClustersClusterResult> clusters,

            ImmutableArray<string> dedicatedHostClusterIds,

            string? dedicatedHostClusterName,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableDictionary<string, object>? tags,

            string? zoneId)
        {
            Clusters = clusters;
            DedicatedHostClusterIds = dedicatedHostClusterIds;
            DedicatedHostClusterName = dedicatedHostClusterName;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Tags = tags;
            ZoneId = zoneId;
        }
    }
}
