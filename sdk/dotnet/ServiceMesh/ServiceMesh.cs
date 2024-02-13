// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceMesh
{
    /// <summary>
    /// Provides a Service Mesh Service Mesh resource.
    /// 
    /// For information about Service Mesh Service Mesh and how to use it, see [What is Service Mesh](https://www.alibabacloud.com/help/en/asm/developer-reference/api-servicemesh-2020-01-11-createservicemesh).
    /// 
    /// &gt; **NOTE:** Available since v1.138.0.
    /// 
    /// ## Import
    /// 
    /// Service Mesh Service Mesh can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:servicemesh/serviceMesh:ServiceMesh example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:servicemesh/serviceMesh:ServiceMesh")]
    public partial class ServiceMesh : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of clusters.
        /// </summary>
        [Output("clusterIds")]
        public Output<ImmutableArray<string>> ClusterIds { get; private set; } = null!;

        /// <summary>
        /// Cluster specification. The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`. Default to `standard`.
        /// </summary>
        [Output("clusterSpec")]
        public Output<string> ClusterSpec { get; private set; } = null!;

        /// <summary>
        /// Service grid creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Whether to customize Prometheus. Value:
        /// -'true': custom Prometheus.
        /// -'false': Do not customize Prometheus. Default value: 'false '.
        /// </summary>
        [Output("customizedPrometheus")]
        public Output<bool?> CustomizedPrometheus { get; private set; } = null!;

        /// <summary>
        /// Grid instance version type. Valid values: `Default` and `Pro`. Default: the standard. Pro: the Pro version.
        /// </summary>
        [Output("edition")]
        public Output<string> Edition { get; private set; } = null!;

        /// <summary>
        /// Data plane KubeAPI access capability. See `extra_configuration` below.
        /// </summary>
        [Output("extraConfiguration")]
        public Output<Outputs.ServiceMeshExtraConfiguration> ExtraConfiguration { get; private set; } = null!;

        /// <summary>
        /// Whether to forcibly delete the ASM instance. Value:
        /// -'true': force deletion of ASM instance
        /// -'false': no forced deletion of ASM instance. Default value: false.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// Load balancing information. See `load_balancer` below.
        /// </summary>
        [Output("loadBalancer")]
        public Output<Outputs.ServiceMeshLoadBalancer> LoadBalancer { get; private set; } = null!;

        /// <summary>
        /// Service grid configuration information. See `mesh_config` below.
        /// </summary>
        [Output("meshConfig")]
        public Output<Outputs.ServiceMeshMeshConfig> MeshConfig { get; private set; } = null!;

        /// <summary>
        /// Service grid network configuration information. See `network` below.
        /// </summary>
        [Output("network")]
        public Output<Outputs.ServiceMeshNetwork> Network { get; private set; } = null!;

        /// <summary>
        /// The Prometheus service address (in non-custom cases, use the ARMS address format).
        /// </summary>
        [Output("prometheusUrl")]
        public Output<string?> PrometheusUrl { get; private set; } = null!;

        /// <summary>
        /// ServiceMeshName.
        /// </summary>
        [Output("serviceMeshName")]
        public Output<string?> ServiceMeshName { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Service grid version number. The version of the resource. you can look up the version using alicloud_service_mesh_versions. Note: The version supports updating from v1.170.0, the relevant version can be obtained via istio_operator_version in `alicloud.servicemesh.getServiceMeshes`.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceMesh resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceMesh(string name, ServiceMeshArgs args, CustomResourceOptions? options = null)
            : base("alicloud:servicemesh/serviceMesh:ServiceMesh", name, args ?? new ServiceMeshArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceMesh(string name, Input<string> id, ServiceMeshState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:servicemesh/serviceMesh:ServiceMesh", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceMesh resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceMesh Get(string name, Input<string> id, ServiceMeshState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceMesh(name, id, state, options);
        }
    }

    public sealed class ServiceMeshArgs : global::Pulumi.ResourceArgs
    {
        [Input("clusterIds")]
        private InputList<string>? _clusterIds;

        /// <summary>
        /// List of clusters.
        /// </summary>
        public InputList<string> ClusterIds
        {
            get => _clusterIds ?? (_clusterIds = new InputList<string>());
            set => _clusterIds = value;
        }

        /// <summary>
        /// Cluster specification. The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`. Default to `standard`.
        /// </summary>
        [Input("clusterSpec")]
        public Input<string>? ClusterSpec { get; set; }

        /// <summary>
        /// Whether to customize Prometheus. Value:
        /// -'true': custom Prometheus.
        /// -'false': Do not customize Prometheus. Default value: 'false '.
        /// </summary>
        [Input("customizedPrometheus")]
        public Input<bool>? CustomizedPrometheus { get; set; }

        /// <summary>
        /// Grid instance version type. Valid values: `Default` and `Pro`. Default: the standard. Pro: the Pro version.
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// Data plane KubeAPI access capability. See `extra_configuration` below.
        /// </summary>
        [Input("extraConfiguration")]
        public Input<Inputs.ServiceMeshExtraConfigurationArgs>? ExtraConfiguration { get; set; }

        /// <summary>
        /// Whether to forcibly delete the ASM instance. Value:
        /// -'true': force deletion of ASM instance
        /// -'false': no forced deletion of ASM instance. Default value: false.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// Load balancing information. See `load_balancer` below.
        /// </summary>
        [Input("loadBalancer")]
        public Input<Inputs.ServiceMeshLoadBalancerArgs>? LoadBalancer { get; set; }

        /// <summary>
        /// Service grid configuration information. See `mesh_config` below.
        /// </summary>
        [Input("meshConfig")]
        public Input<Inputs.ServiceMeshMeshConfigArgs>? MeshConfig { get; set; }

        /// <summary>
        /// Service grid network configuration information. See `network` below.
        /// </summary>
        [Input("network", required: true)]
        public Input<Inputs.ServiceMeshNetworkArgs> Network { get; set; } = null!;

        /// <summary>
        /// The Prometheus service address (in non-custom cases, use the ARMS address format).
        /// </summary>
        [Input("prometheusUrl")]
        public Input<string>? PrometheusUrl { get; set; }

        /// <summary>
        /// ServiceMeshName.
        /// </summary>
        [Input("serviceMeshName")]
        public Input<string>? ServiceMeshName { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Service grid version number. The version of the resource. you can look up the version using alicloud_service_mesh_versions. Note: The version supports updating from v1.170.0, the relevant version can be obtained via istio_operator_version in `alicloud.servicemesh.getServiceMeshes`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ServiceMeshArgs()
        {
        }
        public static new ServiceMeshArgs Empty => new ServiceMeshArgs();
    }

    public sealed class ServiceMeshState : global::Pulumi.ResourceArgs
    {
        [Input("clusterIds")]
        private InputList<string>? _clusterIds;

        /// <summary>
        /// List of clusters.
        /// </summary>
        public InputList<string> ClusterIds
        {
            get => _clusterIds ?? (_clusterIds = new InputList<string>());
            set => _clusterIds = value;
        }

        /// <summary>
        /// Cluster specification. The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`. Default to `standard`.
        /// </summary>
        [Input("clusterSpec")]
        public Input<string>? ClusterSpec { get; set; }

        /// <summary>
        /// Service grid creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Whether to customize Prometheus. Value:
        /// -'true': custom Prometheus.
        /// -'false': Do not customize Prometheus. Default value: 'false '.
        /// </summary>
        [Input("customizedPrometheus")]
        public Input<bool>? CustomizedPrometheus { get; set; }

        /// <summary>
        /// Grid instance version type. Valid values: `Default` and `Pro`. Default: the standard. Pro: the Pro version.
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// Data plane KubeAPI access capability. See `extra_configuration` below.
        /// </summary>
        [Input("extraConfiguration")]
        public Input<Inputs.ServiceMeshExtraConfigurationGetArgs>? ExtraConfiguration { get; set; }

        /// <summary>
        /// Whether to forcibly delete the ASM instance. Value:
        /// -'true': force deletion of ASM instance
        /// -'false': no forced deletion of ASM instance. Default value: false.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// Load balancing information. See `load_balancer` below.
        /// </summary>
        [Input("loadBalancer")]
        public Input<Inputs.ServiceMeshLoadBalancerGetArgs>? LoadBalancer { get; set; }

        /// <summary>
        /// Service grid configuration information. See `mesh_config` below.
        /// </summary>
        [Input("meshConfig")]
        public Input<Inputs.ServiceMeshMeshConfigGetArgs>? MeshConfig { get; set; }

        /// <summary>
        /// Service grid network configuration information. See `network` below.
        /// </summary>
        [Input("network")]
        public Input<Inputs.ServiceMeshNetworkGetArgs>? Network { get; set; }

        /// <summary>
        /// The Prometheus service address (in non-custom cases, use the ARMS address format).
        /// </summary>
        [Input("prometheusUrl")]
        public Input<string>? PrometheusUrl { get; set; }

        /// <summary>
        /// ServiceMeshName.
        /// </summary>
        [Input("serviceMeshName")]
        public Input<string>? ServiceMeshName { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Service grid version number. The version of the resource. you can look up the version using alicloud_service_mesh_versions. Note: The version supports updating from v1.170.0, the relevant version can be obtained via istio_operator_version in `alicloud.servicemesh.getServiceMeshes`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ServiceMeshState()
        {
        }
        public static new ServiceMeshState Empty => new ServiceMeshState();
    }
}
