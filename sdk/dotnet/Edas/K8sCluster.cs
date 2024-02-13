// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Edas
{
    /// <summary>
    /// Provides an EDAS K8s cluster resource. For information about EDAS K8s Cluster and how to use it, see[What is EDAS K8s Cluster](https://www.alibabacloud.com/help/en/doc-detail/85108.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.93.0.
    /// 
    /// ## Import
    /// 
    /// EDAS cluster can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:edas/k8sCluster:K8sCluster cluster cluster_id
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:edas/k8sCluster:K8sCluster")]
    public partial class K8sCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The import status of cluster: 
        /// `1`: success.
        /// `2`: failed.
        /// `3`: importing.
        /// `4`: deleted.
        /// </summary>
        [Output("clusterImportStatus")]
        public Output<int> ClusterImportStatus { get; private set; } = null!;

        /// <summary>
        /// The name of the cluster that you want to create.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// The type of the cluster that you want to create. Valid values only: 5: K8s cluster.
        /// </summary>
        [Output("clusterType")]
        public Output<int> ClusterType { get; private set; } = null!;

        /// <summary>
        /// The ID of the alicloud container service kubernetes cluster that you want to import.
        /// </summary>
        [Output("csClusterId")]
        public Output<string> CsClusterId { get; private set; } = null!;

        /// <summary>
        /// The ID of the namespace where you want to import. You can call the [ListUserDefineRegion](https://www.alibabacloud.com/help/en/doc-detail/149377.htm?spm=a2c63.p38356.879954.34.331054faK2yNvC#doc-api-Edas-ListUserDefineRegion) operation to query the namespace ID.
        /// </summary>
        [Output("namespaceId")]
        public Output<string?> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.
        /// </summary>
        [Output("networkMode")]
        public Output<int> NetworkMode { get; private set; } = null!;

        /// <summary>
        /// The ID of the Virtual Private Cloud (VPC) for the cluster.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a K8sCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public K8sCluster(string name, K8sClusterArgs args, CustomResourceOptions? options = null)
            : base("alicloud:edas/k8sCluster:K8sCluster", name, args ?? new K8sClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private K8sCluster(string name, Input<string> id, K8sClusterState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:edas/k8sCluster:K8sCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing K8sCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static K8sCluster Get(string name, Input<string> id, K8sClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new K8sCluster(name, id, state, options);
        }
    }

    public sealed class K8sClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the alicloud container service kubernetes cluster that you want to import.
        /// </summary>
        [Input("csClusterId", required: true)]
        public Input<string> CsClusterId { get; set; } = null!;

        /// <summary>
        /// The ID of the namespace where you want to import. You can call the [ListUserDefineRegion](https://www.alibabacloud.com/help/en/doc-detail/149377.htm?spm=a2c63.p38356.879954.34.331054faK2yNvC#doc-api-Edas-ListUserDefineRegion) operation to query the namespace ID.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        public K8sClusterArgs()
        {
        }
        public static new K8sClusterArgs Empty => new K8sClusterArgs();
    }

    public sealed class K8sClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The import status of cluster: 
        /// `1`: success.
        /// `2`: failed.
        /// `3`: importing.
        /// `4`: deleted.
        /// </summary>
        [Input("clusterImportStatus")]
        public Input<int>? ClusterImportStatus { get; set; }

        /// <summary>
        /// The name of the cluster that you want to create.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// The type of the cluster that you want to create. Valid values only: 5: K8s cluster.
        /// </summary>
        [Input("clusterType")]
        public Input<int>? ClusterType { get; set; }

        /// <summary>
        /// The ID of the alicloud container service kubernetes cluster that you want to import.
        /// </summary>
        [Input("csClusterId")]
        public Input<string>? CsClusterId { get; set; }

        /// <summary>
        /// The ID of the namespace where you want to import. You can call the [ListUserDefineRegion](https://www.alibabacloud.com/help/en/doc-detail/149377.htm?spm=a2c63.p38356.879954.34.331054faK2yNvC#doc-api-Edas-ListUserDefineRegion) operation to query the namespace ID.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.
        /// </summary>
        [Input("networkMode")]
        public Input<int>? NetworkMode { get; set; }

        /// <summary>
        /// The ID of the Virtual Private Cloud (VPC) for the cluster.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public K8sClusterState()
        {
        }
        public static new K8sClusterState Empty => new K8sClusterState();
    }
}
