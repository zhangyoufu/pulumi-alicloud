// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PrivateLink
{
    /// <summary>
    /// Provides a Private Link Vpc Endpoint Zone resource.
    /// 
    /// For information about Private Link Vpc Endpoint Zone and how to use it, see [What is Vpc Endpoint Zone](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-addzonetovpcendpoint).
    /// 
    /// &gt; **NOTE:** Available since v1.111.0.
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
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var exampleZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var exampleVpcEndpointService = new AliCloud.PrivateLink.VpcEndpointService("exampleVpcEndpointService", new()
    ///     {
    ///         ServiceDescription = name,
    ///         ConnectBandwidth = 103,
    ///         AutoAcceptConnection = false,
    ///     });
    /// 
    ///     var exampleNetwork = new AliCloud.Vpc.Network("exampleNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.0.0.0/8",
    ///     });
    /// 
    ///     var exampleSwitch = new AliCloud.Vpc.Switch("exampleSwitch", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.1.0.0/16",
    ///         VpcId = exampleNetwork.Id,
    ///         ZoneId = exampleZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var exampleSecurityGroup = new AliCloud.Ecs.SecurityGroup("exampleSecurityGroup", new()
    ///     {
    ///         VpcId = exampleNetwork.Id,
    ///     });
    /// 
    ///     var exampleApplicationLoadBalancer = new AliCloud.Slb.ApplicationLoadBalancer("exampleApplicationLoadBalancer", new()
    ///     {
    ///         LoadBalancerName = name,
    ///         VswitchId = exampleSwitch.Id,
    ///         LoadBalancerSpec = "slb.s2.small",
    ///         AddressType = "intranet",
    ///     });
    /// 
    ///     var exampleVpcEndpointServiceResource = new AliCloud.PrivateLink.VpcEndpointServiceResource("exampleVpcEndpointServiceResource", new()
    ///     {
    ///         ServiceId = exampleVpcEndpointService.Id,
    ///         ResourceId = exampleApplicationLoadBalancer.Id,
    ///         ResourceType = "slb",
    ///     });
    /// 
    ///     var exampleVpcEndpoint = new AliCloud.PrivateLink.VpcEndpoint("exampleVpcEndpoint", new()
    ///     {
    ///         ServiceId = exampleVpcEndpointServiceResource.ServiceId,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             exampleSecurityGroup.Id,
    ///         },
    ///         VpcId = exampleNetwork.Id,
    ///         VpcEndpointName = name,
    ///     });
    /// 
    ///     var exampleVpcEndpointZone = new AliCloud.PrivateLink.VpcEndpointZone("exampleVpcEndpointZone", new()
    ///     {
    ///         EndpointId = exampleVpcEndpoint.Id,
    ///         VswitchId = exampleSwitch.Id,
    ///         ZoneId = exampleZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Private Link Vpc Endpoint Zone can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:privatelink/vpcEndpointZone:VpcEndpointZone example &lt;endpoint_id&gt;:&lt;zone_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:privatelink/vpcEndpointZone:VpcEndpointZone")]
    public partial class VpcEndpointZone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
        /// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
        /// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The endpoint ID.
        /// </summary>
        [Output("endpointId")]
        public Output<string> EndpointId { get; private set; } = null!;

        /// <summary>
        /// The IP address of the endpoint ENI.
        /// </summary>
        [Output("eniIp")]
        public Output<string> EniIp { get; private set; } = null!;

        /// <summary>
        /// The state of the zone.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the vSwitch in the zone. .
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The zone ID.
        /// </summary>
        [Output("zoneId")]
        public Output<string?> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpointZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpointZone(string name, VpcEndpointZoneArgs args, CustomResourceOptions? options = null)
            : base("alicloud:privatelink/vpcEndpointZone:VpcEndpointZone", name, args ?? new VpcEndpointZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpointZone(string name, Input<string> id, VpcEndpointZoneState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:privatelink/vpcEndpointZone:VpcEndpointZone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpointZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpointZone Get(string name, Input<string> id, VpcEndpointZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpointZone(name, id, state, options);
        }
    }

    public sealed class VpcEndpointZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
        /// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
        /// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The endpoint ID.
        /// </summary>
        [Input("endpointId", required: true)]
        public Input<string> EndpointId { get; set; } = null!;

        /// <summary>
        /// The IP address of the endpoint ENI.
        /// </summary>
        [Input("eniIp")]
        public Input<string>? EniIp { get; set; }

        /// <summary>
        /// The ID of the vSwitch in the zone. .
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        /// <summary>
        /// The zone ID.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public VpcEndpointZoneArgs()
        {
        }
        public static new VpcEndpointZoneArgs Empty => new VpcEndpointZoneArgs();
    }

    public sealed class VpcEndpointZoneState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
        /// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
        /// - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The endpoint ID.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// The IP address of the endpoint ENI.
        /// </summary>
        [Input("eniIp")]
        public Input<string>? EniIp { get; set; }

        /// <summary>
        /// The state of the zone.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the vSwitch in the zone. .
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The zone ID.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public VpcEndpointZoneState()
        {
        }
        public static new VpcEndpointZoneState Empty => new VpcEndpointZoneState();
    }
}
