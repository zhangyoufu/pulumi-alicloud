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
    /// Provides an Alicloud Serverless App Engine (SAE) Application Load Balancer Attachment resource.
    /// 
    /// For information about Serverless App Engine (SAE) Load Balancer Internet Attachment and how to use it, see [alicloud.sae.LoadBalancerInternet](https://www.alibabacloud.com/help/en/sae/latest/bindslb).
    /// 
    /// &gt; **NOTE:** Available since v1.164.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
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
    ///         ImageUrl = "registry-vpc.cn-hangzhou.aliyuncs.com/lxepoo/apache-php5",
    ///         PackageType = "Image",
    ///         Jdk = "Open JDK 8",
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
    ///         AddressType = "internet",
    ///     });
    /// 
    ///     var defaultLoadBalancerInternet = new AliCloud.Sae.LoadBalancerInternet("defaultLoadBalancerInternet", new()
    ///     {
    ///         AppId = defaultApplication.Id,
    ///         InternetSlbId = defaultApplicationLoadBalancer.Id,
    ///         Internets = new[]
    ///         {
    ///             new AliCloud.Sae.Inputs.LoadBalancerInternetInternetArgs
    ///             {
    ///                 Protocol = "TCP",
    ///                 Port = 80,
    ///                 TargetPort = 8080,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The resource can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:sae/loadBalancerInternet:LoadBalancerInternet example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:sae/loadBalancerInternet:LoadBalancerInternet")]
    public partial class LoadBalancerInternet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The target application ID that needs to be bound to the SLB.
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// Use designated public network SLBs that have been purchased to support non-shared instances.
        /// </summary>
        [Output("internetIp")]
        public Output<string> InternetIp { get; private set; } = null!;

        /// <summary>
        /// The internet SLB ID.
        /// </summary>
        [Output("internetSlbId")]
        public Output<string?> InternetSlbId { get; private set; } = null!;

        /// <summary>
        /// The bound private network SLB. See `internet` below.
        /// </summary>
        [Output("internets")]
        public Output<ImmutableArray<Outputs.LoadBalancerInternetInternet>> Internets { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancerInternet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancerInternet(string name, LoadBalancerInternetArgs args, CustomResourceOptions? options = null)
            : base("alicloud:sae/loadBalancerInternet:LoadBalancerInternet", name, args ?? new LoadBalancerInternetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancerInternet(string name, Input<string> id, LoadBalancerInternetState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:sae/loadBalancerInternet:LoadBalancerInternet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancerInternet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancerInternet Get(string name, Input<string> id, LoadBalancerInternetState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadBalancerInternet(name, id, state, options);
        }
    }

    public sealed class LoadBalancerInternetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target application ID that needs to be bound to the SLB.
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// The internet SLB ID.
        /// </summary>
        [Input("internetSlbId")]
        public Input<string>? InternetSlbId { get; set; }

        [Input("internets", required: true)]
        private InputList<Inputs.LoadBalancerInternetInternetArgs>? _internets;

        /// <summary>
        /// The bound private network SLB. See `internet` below.
        /// </summary>
        public InputList<Inputs.LoadBalancerInternetInternetArgs> Internets
        {
            get => _internets ?? (_internets = new InputList<Inputs.LoadBalancerInternetInternetArgs>());
            set => _internets = value;
        }

        public LoadBalancerInternetArgs()
        {
        }
        public static new LoadBalancerInternetArgs Empty => new LoadBalancerInternetArgs();
    }

    public sealed class LoadBalancerInternetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target application ID that needs to be bound to the SLB.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// Use designated public network SLBs that have been purchased to support non-shared instances.
        /// </summary>
        [Input("internetIp")]
        public Input<string>? InternetIp { get; set; }

        /// <summary>
        /// The internet SLB ID.
        /// </summary>
        [Input("internetSlbId")]
        public Input<string>? InternetSlbId { get; set; }

        [Input("internets")]
        private InputList<Inputs.LoadBalancerInternetInternetGetArgs>? _internets;

        /// <summary>
        /// The bound private network SLB. See `internet` below.
        /// </summary>
        public InputList<Inputs.LoadBalancerInternetInternetGetArgs> Internets
        {
            get => _internets ?? (_internets = new InputList<Inputs.LoadBalancerInternetInternetGetArgs>());
            set => _internets = value;
        }

        public LoadBalancerInternetState()
        {
        }
        public static new LoadBalancerInternetState Empty => new LoadBalancerInternetState();
    }
}
