// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae
{
    /// <summary>
    /// Provides a Serverless App Engine (SAE) Ingress resource.
    /// 
    /// For information about Serverless App Engine (SAE) Ingress and how to use it, see [What is Ingress](https://www.alibabacloud.com/help/en/sae/latest/createingress).
    /// 
    /// &gt; **NOTE:** Available since v1.137.0.
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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var defaultRandomInteger = new Random.RandomInteger("defaultRandomInteger", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.4.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.4.0.0/24",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultNamespace = new AliCloud.Sae.Namespace("defaultNamespace", new()
    ///     {
    ///         NamespaceId = Output.Tuple(defaultRegions, defaultRandomInteger.Result).Apply(values =&gt;
    ///         {
    ///             var defaultRegions = values.Item1;
    ///             var result = values.Item2;
    ///             return $"{defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}:example{result}";
    ///         }),
    ///         NamespaceName = name,
    ///         NamespaceDescription = name,
    ///         EnableMicroRegistration = false,
    ///     });
    /// 
    ///     var defaultApplication = new AliCloud.Sae.Application("defaultApplication", new()
    ///     {
    ///         AppDescription = name,
    ///         AppName = name,
    ///         NamespaceId = defaultNamespace.Id,
    ///         ImageUrl = $"registry-vpc.{defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}.aliyuncs.com/sae-demo-image/consumer:1.0",
    ///         PackageType = "Image",
    ///         SecurityGroupId = defaultSecurityGroup.Id,
    ///         VpcId = defaultNetwork.Id,
    ///         VswitchId = defaultSwitch.Id,
    ///         Timezone = "Asia/Beijing",
    ///         Replicas = 5,
    ///         Cpu = 500,
    ///         Memory = 2048,
    ///     });
    /// 
    ///     var defaultApplicationLoadBalancer = new AliCloud.Slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer", new()
    ///     {
    ///         LoadBalancerName = name,
    ///         VswitchId = defaultSwitch.Id,
    ///         LoadBalancerSpec = "slb.s2.small",
    ///         AddressType = "intranet",
    ///     });
    /// 
    ///     var defaultIngress = new AliCloud.Sae.Ingress("defaultIngress", new()
    ///     {
    ///         SlbId = defaultApplicationLoadBalancer.Id,
    ///         NamespaceId = defaultNamespace.Id,
    ///         ListenerPort = 80,
    ///         Rules = new[]
    ///         {
    ///             new AliCloud.Sae.Inputs.IngressRuleArgs
    ///             {
    ///                 AppId = defaultApplication.Id,
    ///                 ContainerPort = 443,
    ///                 Domain = "www.alicloud.com",
    ///                 AppName = defaultApplication.AppName,
    ///                 Path = "/",
    ///             },
    ///         },
    ///         DefaultRule = new AliCloud.Sae.Inputs.IngressDefaultRuleArgs
    ///         {
    ///             AppId = defaultApplication.Id,
    ///             ContainerPort = 443,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Serverless App Engine (SAE) Ingress can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:sae/ingress:Ingress example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:sae/ingress:Ingress")]
    public partial class Ingress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The certificate ID of the HTTPS listener. The `cert_id` takes effect only when `load_balance_type` is set to `clb`.
        /// </summary>
        [Output("certId")]
        public Output<string?> CertId { get; private set; } = null!;

        /// <summary>
        /// The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `cert_ids` takes effect only when `load_balance_type` is set to `alb`.
        /// </summary>
        [Output("certIds")]
        public Output<string?> CertIds { get; private set; } = null!;

        /// <summary>
        /// Default Rule. See `default_rule` below.
        /// </summary>
        [Output("defaultRule")]
        public Output<Outputs.IngressDefaultRule?> DefaultRule { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// SLB listening port.
        /// </summary>
        [Output("listenerPort")]
        public Output<int> ListenerPort { get; private set; } = null!;

        /// <summary>
        /// The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
        /// </summary>
        [Output("listenerProtocol")]
        public Output<string> ListenerProtocol { get; private set; } = null!;

        /// <summary>
        /// The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
        /// </summary>
        [Output("loadBalanceType")]
        public Output<string> LoadBalanceType { get; private set; } = null!;

        /// <summary>
        /// The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.IngressRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// SLB ID.
        /// </summary>
        [Output("slbId")]
        public Output<string> SlbId { get; private set; } = null!;


        /// <summary>
        /// Create a Ingress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ingress(string name, IngressArgs args, CustomResourceOptions? options = null)
            : base("alicloud:sae/ingress:Ingress", name, args ?? new IngressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ingress(string name, Input<string> id, IngressState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:sae/ingress:Ingress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ingress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ingress Get(string name, Input<string> id, IngressState? state = null, CustomResourceOptions? options = null)
        {
            return new Ingress(name, id, state, options);
        }
    }

    public sealed class IngressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate ID of the HTTPS listener. The `cert_id` takes effect only when `load_balance_type` is set to `clb`.
        /// </summary>
        [Input("certId")]
        public Input<string>? CertId { get; set; }

        /// <summary>
        /// The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `cert_ids` takes effect only when `load_balance_type` is set to `alb`.
        /// </summary>
        [Input("certIds")]
        public Input<string>? CertIds { get; set; }

        /// <summary>
        /// Default Rule. See `default_rule` below.
        /// </summary>
        [Input("defaultRule")]
        public Input<Inputs.IngressDefaultRuleArgs>? DefaultRule { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// SLB listening port.
        /// </summary>
        [Input("listenerPort", required: true)]
        public Input<int> ListenerPort { get; set; } = null!;

        /// <summary>
        /// The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
        /// </summary>
        [Input("listenerProtocol")]
        public Input<string>? ListenerProtocol { get; set; }

        /// <summary>
        /// The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
        /// </summary>
        [Input("loadBalanceType")]
        public Input<string>? LoadBalanceType { get; set; }

        /// <summary>
        /// The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        [Input("rules", required: true)]
        private InputList<Inputs.IngressRuleArgs>? _rules;

        /// <summary>
        /// Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
        /// </summary>
        public InputList<Inputs.IngressRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.IngressRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// SLB ID.
        /// </summary>
        [Input("slbId", required: true)]
        public Input<string> SlbId { get; set; } = null!;

        public IngressArgs()
        {
        }
        public static new IngressArgs Empty => new IngressArgs();
    }

    public sealed class IngressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate ID of the HTTPS listener. The `cert_id` takes effect only when `load_balance_type` is set to `clb`.
        /// </summary>
        [Input("certId")]
        public Input<string>? CertId { get; set; }

        /// <summary>
        /// The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `cert_ids` takes effect only when `load_balance_type` is set to `alb`.
        /// </summary>
        [Input("certIds")]
        public Input<string>? CertIds { get; set; }

        /// <summary>
        /// Default Rule. See `default_rule` below.
        /// </summary>
        [Input("defaultRule")]
        public Input<Inputs.IngressDefaultRuleGetArgs>? DefaultRule { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// SLB listening port.
        /// </summary>
        [Input("listenerPort")]
        public Input<int>? ListenerPort { get; set; }

        /// <summary>
        /// The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
        /// </summary>
        [Input("listenerProtocol")]
        public Input<string>? ListenerProtocol { get; set; }

        /// <summary>
        /// The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
        /// </summary>
        [Input("loadBalanceType")]
        public Input<string>? LoadBalanceType { get; set; }

        /// <summary>
        /// The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        [Input("rules")]
        private InputList<Inputs.IngressRuleGetArgs>? _rules;

        /// <summary>
        /// Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
        /// </summary>
        public InputList<Inputs.IngressRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.IngressRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// SLB ID.
        /// </summary>
        [Input("slbId")]
        public Input<string>? SlbId { get; set; }

        public IngressState()
        {
        }
        public static new IngressState Empty => new IngressState();
    }
}
