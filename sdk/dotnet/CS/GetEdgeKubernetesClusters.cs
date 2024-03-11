// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    public static class GetEdgeKubernetesClusters
    {
        /// <summary>
        /// This data source provides a list Container Service Edge Kubernetes Clusters on Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.103.0+
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var k8sClusters = AliCloud.CS.GetEdgeKubernetesClusters.Invoke(new()
        ///     {
        ///         NameRegex = "my-first-k8s",
        ///         OutputFile = "my-first-k8s-json",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["output"] = k8sClusters.Apply(getEdgeKubernetesClustersResult =&gt; getEdgeKubernetesClustersResult.Clusters),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetEdgeKubernetesClustersResult> InvokeAsync(GetEdgeKubernetesClustersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEdgeKubernetesClustersResult>("alicloud:cs/getEdgeKubernetesClusters:getEdgeKubernetesClusters", args ?? new GetEdgeKubernetesClustersArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list Container Service Edge Kubernetes Clusters on Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.103.0+
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var k8sClusters = AliCloud.CS.GetEdgeKubernetesClusters.Invoke(new()
        ///     {
        ///         NameRegex = "my-first-k8s",
        ///         OutputFile = "my-first-k8s-json",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["output"] = k8sClusters.Apply(getEdgeKubernetesClustersResult =&gt; getEdgeKubernetesClustersResult.Clusters),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetEdgeKubernetesClustersResult> Invoke(GetEdgeKubernetesClustersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEdgeKubernetesClustersResult>("alicloud:cs/getEdgeKubernetesClusters:getEdgeKubernetesClusters", args ?? new GetEdgeKubernetesClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEdgeKubernetesClustersArgs : global::Pulumi.InvokeArgs
    {
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// Cluster IDs to filter.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by cluster name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetEdgeKubernetesClustersArgs()
        {
        }
        public static new GetEdgeKubernetesClustersArgs Empty => new GetEdgeKubernetesClustersArgs();
    }

    public sealed class GetEdgeKubernetesClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// Cluster IDs to filter.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by cluster name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetEdgeKubernetesClustersInvokeArgs()
        {
        }
        public static new GetEdgeKubernetesClustersInvokeArgs Empty => new GetEdgeKubernetesClustersInvokeArgs();
    }


    [OutputType]
    public sealed class GetEdgeKubernetesClustersResult
    {
        /// <summary>
        /// A list of matched Kubernetes clusters. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEdgeKubernetesClustersClusterResult> Clusters;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of matched Kubernetes clusters' ids.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of matched Kubernetes clusters' names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetEdgeKubernetesClustersResult(
            ImmutableArray<Outputs.GetEdgeKubernetesClustersClusterResult> clusters,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            Clusters = clusters;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}
