// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    /// <summary>
    /// This resource will help you to manager a Serverless Kubernetes Cluster, see [What is serverless kubernetes](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/developer-reference/create-a-dedicated-kubernetes-cluster-that-supports-sandboxed-containers). The cluster is same as container service created by web console.
    /// 
    /// &gt; **NOTE:** Available since v1.58.0.
    /// 
    /// &gt; **NOTE:** Serverless Kubernetes cluster only supports VPC network and it can access internet while creating kubernetes cluster.
    /// A Nat Gateway and configuring a SNAT for it can ensure one VPC network access internet. If there is no nat gateway in the
    /// VPC, you can set `new_nat_gateway` to "true" to create one automatically.
    /// 
    /// &gt; **NOTE:** Creating serverless kubernetes cluster need to install several packages and it will cost about 5 minutes. Please be patient.
    /// 
    /// &gt; **NOTE:** The provider supports to download kube config, client certificate, client key and cluster ca certificate
    /// after creating cluster successfully, and you can put them into the specified location, like '~/.kube/config'.
    /// 
    /// &gt; **NOTE:** If you want to manage serverless Kubernetes, you can use Kubernetes Provider.
    /// 
    /// &gt; **NOTE:** You need to activate several other products and confirm Authorization Policy used by Container Service before using this resource.
    /// Please refer to the `Authorization management` and `Cluster management` sections in the [Document Center](https://www.alibabacloud.com/help/doc-detail/86488.htm).
    /// 
    /// &gt; **NOTE:** From version 1.162.0, support for creating professional serverless cluster.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "ask-example";
    ///     var @default = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.1.0.0/21",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("default", new()
    ///     {
    ///         VswitchName = name,
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "10.1.1.0/24",
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
    ///     });
    /// 
    ///     var serverless = new AliCloud.CS.ServerlessKubernetes("serverless", new()
    ///     {
    ///         NamePrefix = name,
    ///         ClusterSpec = "ack.pro.small",
    ///         VpcId = defaultNetwork.Id,
    ///         VswitchIds = new[]
    ///         {
    ///             defaultSwitch.Id,
    ///         },
    ///         NewNatGateway = true,
    ///         EndpointPublicAccessEnabled = true,
    ///         DeletionProtection = false,
    ///         LoadBalancerSpec = "slb.s2.small",
    ///         TimeZone = "Asia/Shanghai",
    ///         ServiceCidr = "172.21.0.0/20",
    ///         ServiceDiscoveryTypes = new[]
    ///         {
    ///             "PrivateZone",
    ///         },
    ///         LoggingType = "SLS",
    ///         Tags = 
    ///         {
    ///             { "k-aa", "v-aa" },
    ///             { "k-bb", "v-aa" },
    ///         },
    ///         Addons = new[]
    ///         {
    ///             new AliCloud.CS.Inputs.ServerlessKubernetesAddonArgs
    ///             {
    ///                 Name = "alb-ingress-controller",
    ///             },
    ///             new AliCloud.CS.Inputs.ServerlessKubernetesAddonArgs
    ///             {
    ///                 Name = "metrics-server",
    ///             },
    ///             new AliCloud.CS.Inputs.ServerlessKubernetesAddonArgs
    ///             {
    ///                 Name = "knative",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Serverless Kubernetes cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of `pulumi preview`.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cs/serverlessKubernetes:ServerlessKubernetes main ce4273f9156874b46bb
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cs/serverlessKubernetes:ServerlessKubernetes")]
    public partial class ServerlessKubernetes : global::Pulumi.CustomResource
    {
        /// <summary>
        /// You can specific network plugin,log component,ingress component and so on. See `addons` below.
        /// </summary>
        [Output("addons")]
        public Output<ImmutableArray<Outputs.ServerlessKubernetesAddon>> Addons { get; private set; } = null!;

        /// <summary>
        /// The path of client certificate, like `~/.kube/client-cert.pem`.
        /// </summary>
        [Output("clientCert")]
        public Output<string?> ClientCert { get; private set; } = null!;

        /// <summary>
        /// The path of client key, like `~/.kube/client-key.pem`.
        /// </summary>
        [Output("clientKey")]
        public Output<string?> ClientKey { get; private set; } = null!;

        /// <summary>
        /// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
        /// </summary>
        [Output("clusterCaCert")]
        public Output<string?> ClusterCaCert { get; private set; } = null!;

        /// <summary>
        /// The cluster specifications of serverless kubernetes cluster, which can be empty. Valid values:
        /// - ack.standard: Standard serverless clusters.
        /// - ack.pro.small: Professional serverless clusters.
        /// </summary>
        [Output("clusterSpec")]
        public Output<string> ClusterSpec { get; private set; } = null!;

        /// <summary>
        /// whether to create a v2 version cluster.
        /// 
        /// *Removed params*
        /// </summary>
        [Output("createV2Cluster")]
        public Output<bool> CreateV2Cluster { get; private set; } = null!;

        /// <summary>
        /// Whether enable the deletion protection or not.
        /// - true: Enable deletion protection.
        /// - false: Disable deletion protection.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// Whether to enable cluster to support RRSA for version 1.22.3+. Default to `false`. Once the RRSA function is turned on, it is not allowed to turn off. If your cluster has enabled this function, please manually modify your tf file and add the rrsa configuration to the file, learn more [RAM Roles for Service Accounts](https://www.alibabacloud.com/help/zh/container-service-for-kubernetes/latest/use-rrsa-to-enforce-access-control).
        /// </summary>
        [Output("enableRrsa")]
        public Output<bool?> EnableRrsa { get; private set; } = null!;

        /// <summary>
        /// Whether to create internet  eip for API Server. Default to false.
        /// </summary>
        [Output("endpointPublicAccessEnabled")]
        public Output<bool?> EndpointPublicAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// Default false, when you want to change `vpc_id` and `vswitch_id`, you have to set this field to true, then the cluster will be recreated.
        /// </summary>
        [Output("forceUpdate")]
        public Output<bool?> ForceUpdate { get; private set; } = null!;

        /// <summary>
        /// The path of kube config, like `~/.kube/config`.
        /// </summary>
        [Output("kubeConfig")]
        public Output<string?> KubeConfig { get; private set; } = null!;

        /// <summary>
        /// The cluster api server load balance instance specification, default `slb.s2.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
        /// </summary>
        [Output("loadBalancerSpec")]
        public Output<string> LoadBalancerSpec { get; private set; } = null!;

        /// <summary>
        /// Enable log service, Valid value `SLS`.
        /// </summary>
        [Output("loggingType")]
        public Output<string?> LoggingType { get; private set; } = null!;

        /// <summary>
        /// The kubernetes cluster's name. It is the only in one Alicloud account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// Whether to create a new nat gateway while creating kubernetes cluster. SNAT must be configured when a new VPC is automatically created. Default is `true`.
        /// </summary>
        [Output("newNatGateway")]
        public Output<bool?> NewNatGateway { get; private set; } = null!;

        /// <summary>
        /// Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `service_discovery_types`.
        /// </summary>
        [Output("privateZone")]
        public Output<bool?> PrivateZone { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        [Output("retainResources")]
        public Output<ImmutableArray<string>> RetainResources { get; private set; } = null!;

        /// <summary>
        /// Nested attribute containing RRSA related data for your cluster. See `rrsa_metadata` below.
        /// </summary>
        [Output("rrsaMetadata")]
        public Output<Outputs.ServerlessKubernetesRrsaMetadata> RrsaMetadata { get; private set; } = null!;

        /// <summary>
        /// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// CIDR block of the service network. The specified CIDR block cannot overlap with that of the VPC or those of the ACK clusters that are deployed in the VPC. The CIDR block cannot be modified after the cluster is created.
        /// </summary>
        [Output("serviceCidr")]
        public Output<string?> ServiceCidr { get; private set; } = null!;

        /// <summary>
        /// Service discovery type. If the value is empty, it means that service discovery is not enabled. Valid values are `CoreDNS` and `PrivateZone`.
        /// </summary>
        [Output("serviceDiscoveryTypes")]
        public Output<ImmutableArray<string>> ServiceDiscoveryTypes { get; private set; } = null!;

        /// <summary>
        /// If you use an existing SLS project, you must specify `sls_project_name`.
        /// </summary>
        [Output("slsProjectName")]
        public Output<string> SlsProjectName { get; private set; } = null!;

        /// <summary>
        /// Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The time zone of the cluster.
        /// </summary>
        [Output("timeZone")]
        public Output<string> TimeZone { get; private set; } = null!;

        /// <summary>
        /// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availability_zone` specified.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The vswitches where new kubernetes cluster will be located.
        /// </summary>
        [Output("vswitchIds")]
        public Output<ImmutableArray<string>> VswitchIds { get; private set; } = null!;

        /// <summary>
        /// When creating a cluster using automatic VPC creation, you need to specify the zone where the VPC is located.
        /// </summary>
        [Output("zoneId")]
        public Output<string?> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a ServerlessKubernetes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerlessKubernetes(string name, ServerlessKubernetesArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cs/serverlessKubernetes:ServerlessKubernetes", name, args ?? new ServerlessKubernetesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerlessKubernetes(string name, Input<string> id, ServerlessKubernetesState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cs/serverlessKubernetes:ServerlessKubernetes", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerlessKubernetes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerlessKubernetes Get(string name, Input<string> id, ServerlessKubernetesState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerlessKubernetes(name, id, state, options);
        }
    }

    public sealed class ServerlessKubernetesArgs : global::Pulumi.ResourceArgs
    {
        [Input("addons")]
        private InputList<Inputs.ServerlessKubernetesAddonArgs>? _addons;

        /// <summary>
        /// You can specific network plugin,log component,ingress component and so on. See `addons` below.
        /// </summary>
        public InputList<Inputs.ServerlessKubernetesAddonArgs> Addons
        {
            get => _addons ?? (_addons = new InputList<Inputs.ServerlessKubernetesAddonArgs>());
            set => _addons = value;
        }

        /// <summary>
        /// The path of client certificate, like `~/.kube/client-cert.pem`.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// The path of client key, like `~/.kube/client-key.pem`.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
        /// </summary>
        [Input("clusterCaCert")]
        public Input<string>? ClusterCaCert { get; set; }

        /// <summary>
        /// The cluster specifications of serverless kubernetes cluster, which can be empty. Valid values:
        /// - ack.standard: Standard serverless clusters.
        /// - ack.pro.small: Professional serverless clusters.
        /// </summary>
        [Input("clusterSpec")]
        public Input<string>? ClusterSpec { get; set; }

        /// <summary>
        /// whether to create a v2 version cluster.
        /// 
        /// *Removed params*
        /// </summary>
        [Input("createV2Cluster")]
        public Input<bool>? CreateV2Cluster { get; set; }

        /// <summary>
        /// Whether enable the deletion protection or not.
        /// - true: Enable deletion protection.
        /// - false: Disable deletion protection.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Whether to enable cluster to support RRSA for version 1.22.3+. Default to `false`. Once the RRSA function is turned on, it is not allowed to turn off. If your cluster has enabled this function, please manually modify your tf file and add the rrsa configuration to the file, learn more [RAM Roles for Service Accounts](https://www.alibabacloud.com/help/zh/container-service-for-kubernetes/latest/use-rrsa-to-enforce-access-control).
        /// </summary>
        [Input("enableRrsa")]
        public Input<bool>? EnableRrsa { get; set; }

        /// <summary>
        /// Whether to create internet  eip for API Server. Default to false.
        /// </summary>
        [Input("endpointPublicAccessEnabled")]
        public Input<bool>? EndpointPublicAccessEnabled { get; set; }

        /// <summary>
        /// Default false, when you want to change `vpc_id` and `vswitch_id`, you have to set this field to true, then the cluster will be recreated.
        /// </summary>
        [Input("forceUpdate")]
        public Input<bool>? ForceUpdate { get; set; }

        /// <summary>
        /// The path of kube config, like `~/.kube/config`.
        /// </summary>
        [Input("kubeConfig")]
        public Input<string>? KubeConfig { get; set; }

        /// <summary>
        /// The cluster api server load balance instance specification, default `slb.s2.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
        /// </summary>
        [Input("loadBalancerSpec")]
        public Input<string>? LoadBalancerSpec { get; set; }

        /// <summary>
        /// Enable log service, Valid value `SLS`.
        /// </summary>
        [Input("loggingType")]
        public Input<string>? LoggingType { get; set; }

        /// <summary>
        /// The kubernetes cluster's name. It is the only in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Whether to create a new nat gateway while creating kubernetes cluster. SNAT must be configured when a new VPC is automatically created. Default is `true`.
        /// </summary>
        [Input("newNatGateway")]
        public Input<bool>? NewNatGateway { get; set; }

        /// <summary>
        /// Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `service_discovery_types`.
        /// </summary>
        [Input("privateZone")]
        public Input<bool>? PrivateZone { get; set; }

        /// <summary>
        /// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("retainResources")]
        private InputList<string>? _retainResources;
        public InputList<string> RetainResources
        {
            get => _retainResources ?? (_retainResources = new InputList<string>());
            set => _retainResources = value;
        }

        /// <summary>
        /// Nested attribute containing RRSA related data for your cluster. See `rrsa_metadata` below.
        /// </summary>
        [Input("rrsaMetadata")]
        public Input<Inputs.ServerlessKubernetesRrsaMetadataArgs>? RrsaMetadata { get; set; }

        /// <summary>
        /// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// CIDR block of the service network. The specified CIDR block cannot overlap with that of the VPC or those of the ACK clusters that are deployed in the VPC. The CIDR block cannot be modified after the cluster is created.
        /// </summary>
        [Input("serviceCidr")]
        public Input<string>? ServiceCidr { get; set; }

        [Input("serviceDiscoveryTypes")]
        private InputList<string>? _serviceDiscoveryTypes;

        /// <summary>
        /// Service discovery type. If the value is empty, it means that service discovery is not enabled. Valid values are `CoreDNS` and `PrivateZone`.
        /// </summary>
        public InputList<string> ServiceDiscoveryTypes
        {
            get => _serviceDiscoveryTypes ?? (_serviceDiscoveryTypes = new InputList<string>());
            set => _serviceDiscoveryTypes = value;
        }

        /// <summary>
        /// If you use an existing SLS project, you must specify `sls_project_name`.
        /// </summary>
        [Input("slsProjectName")]
        public Input<string>? SlsProjectName { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The time zone of the cluster.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availability_zone` specified.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        [Input("vswitchIds")]
        private InputList<string>? _vswitchIds;

        /// <summary>
        /// The vswitches where new kubernetes cluster will be located.
        /// </summary>
        public InputList<string> VswitchIds
        {
            get => _vswitchIds ?? (_vswitchIds = new InputList<string>());
            set => _vswitchIds = value;
        }

        /// <summary>
        /// When creating a cluster using automatic VPC creation, you need to specify the zone where the VPC is located.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ServerlessKubernetesArgs()
        {
        }
        public static new ServerlessKubernetesArgs Empty => new ServerlessKubernetesArgs();
    }

    public sealed class ServerlessKubernetesState : global::Pulumi.ResourceArgs
    {
        [Input("addons")]
        private InputList<Inputs.ServerlessKubernetesAddonGetArgs>? _addons;

        /// <summary>
        /// You can specific network plugin,log component,ingress component and so on. See `addons` below.
        /// </summary>
        public InputList<Inputs.ServerlessKubernetesAddonGetArgs> Addons
        {
            get => _addons ?? (_addons = new InputList<Inputs.ServerlessKubernetesAddonGetArgs>());
            set => _addons = value;
        }

        /// <summary>
        /// The path of client certificate, like `~/.kube/client-cert.pem`.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// The path of client key, like `~/.kube/client-key.pem`.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// The path of cluster ca certificate, like `~/.kube/cluster-ca-cert.pem`
        /// </summary>
        [Input("clusterCaCert")]
        public Input<string>? ClusterCaCert { get; set; }

        /// <summary>
        /// The cluster specifications of serverless kubernetes cluster, which can be empty. Valid values:
        /// - ack.standard: Standard serverless clusters.
        /// - ack.pro.small: Professional serverless clusters.
        /// </summary>
        [Input("clusterSpec")]
        public Input<string>? ClusterSpec { get; set; }

        /// <summary>
        /// whether to create a v2 version cluster.
        /// 
        /// *Removed params*
        /// </summary>
        [Input("createV2Cluster")]
        public Input<bool>? CreateV2Cluster { get; set; }

        /// <summary>
        /// Whether enable the deletion protection or not.
        /// - true: Enable deletion protection.
        /// - false: Disable deletion protection.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Whether to enable cluster to support RRSA for version 1.22.3+. Default to `false`. Once the RRSA function is turned on, it is not allowed to turn off. If your cluster has enabled this function, please manually modify your tf file and add the rrsa configuration to the file, learn more [RAM Roles for Service Accounts](https://www.alibabacloud.com/help/zh/container-service-for-kubernetes/latest/use-rrsa-to-enforce-access-control).
        /// </summary>
        [Input("enableRrsa")]
        public Input<bool>? EnableRrsa { get; set; }

        /// <summary>
        /// Whether to create internet  eip for API Server. Default to false.
        /// </summary>
        [Input("endpointPublicAccessEnabled")]
        public Input<bool>? EndpointPublicAccessEnabled { get; set; }

        /// <summary>
        /// Default false, when you want to change `vpc_id` and `vswitch_id`, you have to set this field to true, then the cluster will be recreated.
        /// </summary>
        [Input("forceUpdate")]
        public Input<bool>? ForceUpdate { get; set; }

        /// <summary>
        /// The path of kube config, like `~/.kube/config`.
        /// </summary>
        [Input("kubeConfig")]
        public Input<string>? KubeConfig { get; set; }

        /// <summary>
        /// The cluster api server load balance instance specification, default `slb.s2.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
        /// </summary>
        [Input("loadBalancerSpec")]
        public Input<string>? LoadBalancerSpec { get; set; }

        /// <summary>
        /// Enable log service, Valid value `SLS`.
        /// </summary>
        [Input("loggingType")]
        public Input<string>? LoggingType { get; set; }

        /// <summary>
        /// The kubernetes cluster's name. It is the only in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Whether to create a new nat gateway while creating kubernetes cluster. SNAT must be configured when a new VPC is automatically created. Default is `true`.
        /// </summary>
        [Input("newNatGateway")]
        public Input<bool>? NewNatGateway { get; set; }

        /// <summary>
        /// Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `service_discovery_types`.
        /// </summary>
        [Input("privateZone")]
        public Input<bool>? PrivateZone { get; set; }

        /// <summary>
        /// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("retainResources")]
        private InputList<string>? _retainResources;
        public InputList<string> RetainResources
        {
            get => _retainResources ?? (_retainResources = new InputList<string>());
            set => _retainResources = value;
        }

        /// <summary>
        /// Nested attribute containing RRSA related data for your cluster. See `rrsa_metadata` below.
        /// </summary>
        [Input("rrsaMetadata")]
        public Input<Inputs.ServerlessKubernetesRrsaMetadataGetArgs>? RrsaMetadata { get; set; }

        /// <summary>
        /// The ID of the security group to which the ECS instances in the cluster belong. If it is not specified, a new Security group will be built.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// CIDR block of the service network. The specified CIDR block cannot overlap with that of the VPC or those of the ACK clusters that are deployed in the VPC. The CIDR block cannot be modified after the cluster is created.
        /// </summary>
        [Input("serviceCidr")]
        public Input<string>? ServiceCidr { get; set; }

        [Input("serviceDiscoveryTypes")]
        private InputList<string>? _serviceDiscoveryTypes;

        /// <summary>
        /// Service discovery type. If the value is empty, it means that service discovery is not enabled. Valid values are `CoreDNS` and `PrivateZone`.
        /// </summary>
        public InputList<string> ServiceDiscoveryTypes
        {
            get => _serviceDiscoveryTypes ?? (_serviceDiscoveryTypes = new InputList<string>());
            set => _serviceDiscoveryTypes = value;
        }

        /// <summary>
        /// If you use an existing SLS project, you must specify `sls_project_name`.
        /// </summary>
        [Input("slsProjectName")]
        public Input<string>? SlsProjectName { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Default nil, A map of tags assigned to the kubernetes cluster and work nodes.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The time zone of the cluster.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC  will be built.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availability_zone` specified.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        [Input("vswitchIds")]
        private InputList<string>? _vswitchIds;

        /// <summary>
        /// The vswitches where new kubernetes cluster will be located.
        /// </summary>
        public InputList<string> VswitchIds
        {
            get => _vswitchIds ?? (_vswitchIds = new InputList<string>());
            set => _vswitchIds = value;
        }

        /// <summary>
        /// When creating a cluster using automatic VPC creation, you need to specify the zone where the VPC is located.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ServerlessKubernetesState()
        {
        }
        public static new ServerlessKubernetesState Empty => new ServerlessKubernetesState();
    }
}
