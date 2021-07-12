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
    /// ## Import
    /// 
    /// Serverless Kubernetes cluster can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cs/serverlessKubernetes:ServerlessKubernetes main ce4273f9156874b46bb
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cs/serverlessKubernetes:ServerlessKubernetes")]
    public partial class ServerlessKubernetes : Pulumi.CustomResource
    {
        /// <summary>
        /// ) You can specific network plugin,log component,ingress component and so on.Detailed below.
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
        /// Whether enable the deletion protection or not.
        /// - true: Enable deletion protection.
        /// - false: Disable deletion protection.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

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
        /// The cluster api server load balance instance specification, default `slb.s1.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
        /// </summary>
        [Output("loadBalancerSpec")]
        public Output<string?> LoadBalancerSpec { get; private set; } = null!;

        /// <summary>
        /// Enable log service, Valid value `SLS`.
        /// </summary>
        [Output("loggingType")]
        public Output<string?> LoggingType { get; private set; } = null!;

        /// <summary>
        /// Name of the ACK add-on. The name must match one of the names returned by [DescribeAddons](https://help.aliyun.com/document_detail/171524.html).
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
        /// (Optional, ForceNew) Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `service_discovery_types`.
        /// </summary>
        [Output("privateZone")]
        public Output<bool?> PrivateZone { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

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
        /// (Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availability_zone` specified.
        /// </summary>
        [Output("vswitchId")]
        public Output<string?> VswitchId { get; private set; } = null!;

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

    public sealed class ServerlessKubernetesArgs : Pulumi.ResourceArgs
    {
        [Input("addons")]
        private InputList<Inputs.ServerlessKubernetesAddonArgs>? _addons;

        /// <summary>
        /// ) You can specific network plugin,log component,ingress component and so on.Detailed below.
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
        /// Whether enable the deletion protection or not.
        /// - true: Enable deletion protection.
        /// - false: Disable deletion protection.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

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
        /// The cluster api server load balance instance specification, default `slb.s1.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
        /// </summary>
        [Input("loadBalancerSpec")]
        public Input<string>? LoadBalancerSpec { get; set; }

        /// <summary>
        /// Enable log service, Valid value `SLS`.
        /// </summary>
        [Input("loggingType")]
        public Input<string>? LoggingType { get; set; }

        /// <summary>
        /// Name of the ACK add-on. The name must match one of the names returned by [DescribeAddons](https://help.aliyun.com/document_detail/171524.html).
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
        /// (Optional, ForceNew) Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `service_discovery_types`.
        /// </summary>
        [Input("privateZone")]
        public Input<bool>? PrivateZone { get; set; }

        /// <summary>
        /// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

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
        /// (Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availability_zone` specified.
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
    }

    public sealed class ServerlessKubernetesState : Pulumi.ResourceArgs
    {
        [Input("addons")]
        private InputList<Inputs.ServerlessKubernetesAddonGetArgs>? _addons;

        /// <summary>
        /// ) You can specific network plugin,log component,ingress component and so on.Detailed below.
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
        /// Whether enable the deletion protection or not.
        /// - true: Enable deletion protection.
        /// - false: Disable deletion protection.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

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
        /// The cluster api server load balance instance specification, default `slb.s1.small`. For more information on how to select a LB instance specification, see [SLB instance overview](https://help.aliyun.com/document_detail/85931.html).
        /// </summary>
        [Input("loadBalancerSpec")]
        public Input<string>? LoadBalancerSpec { get; set; }

        /// <summary>
        /// Enable log service, Valid value `SLS`.
        /// </summary>
        [Input("loggingType")]
        public Input<string>? LoggingType { get; set; }

        /// <summary>
        /// Name of the ACK add-on. The name must match one of the names returned by [DescribeAddons](https://help.aliyun.com/document_detail/171524.html).
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
        /// (Optional, ForceNew) Has been deprecated from provider version 1.123.1. `PrivateZone` is used as the enumeration value of `service_discovery_types`.
        /// </summary>
        [Input("privateZone")]
        public Input<bool>? PrivateZone { get; set; }

        /// <summary>
        /// The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

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
        /// (Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which `availability_zone` specified.
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
    }
}
