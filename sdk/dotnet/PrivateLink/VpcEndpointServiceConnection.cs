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
    /// Provides a Private Link Vpc Endpoint Connection resource. vpc endpoint connection.
    /// 
    /// For information about Private Link Vpc Endpoint Connection and how to use it, see [What is Vpc Endpoint Connection](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-enablevpcendpointzoneconnection).
    /// 
    /// &gt; **NOTE:** Available since v1.110.0.
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
    ///     var example = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var exampleVpcEndpointService = new AliCloud.PrivateLink.VpcEndpointService("example", new()
    ///     {
    ///         ServiceDescription = name,
    ///         ConnectBandwidth = 103,
    ///         AutoAcceptConnection = false,
    ///     });
    /// 
    ///     var exampleNetwork = new AliCloud.Vpc.Network("example", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.0.0.0/8",
    ///     });
    /// 
    ///     var exampleSwitch = new AliCloud.Vpc.Switch("example", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.1.0.0/16",
    ///         VpcId = exampleNetwork.Id,
    ///         ZoneId = example.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var exampleSecurityGroup = new AliCloud.Ecs.SecurityGroup("example", new()
    ///     {
    ///         Name = name,
    ///         VpcId = exampleNetwork.Id,
    ///     });
    /// 
    ///     var exampleApplicationLoadBalancer = new AliCloud.Slb.ApplicationLoadBalancer("example", new()
    ///     {
    ///         LoadBalancerName = name,
    ///         VswitchId = exampleSwitch.Id,
    ///         LoadBalancerSpec = "slb.s2.small",
    ///         AddressType = "intranet",
    ///     });
    /// 
    ///     var exampleVpcEndpointServiceResource = new AliCloud.PrivateLink.VpcEndpointServiceResource("example", new()
    ///     {
    ///         ServiceId = exampleVpcEndpointService.Id,
    ///         ResourceId = exampleApplicationLoadBalancer.Id,
    ///         ResourceType = "slb",
    ///     });
    /// 
    ///     var exampleVpcEndpoint = new AliCloud.PrivateLink.VpcEndpoint("example", new()
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
    ///     var exampleVpcEndpointServiceConnection = new AliCloud.PrivateLink.VpcEndpointServiceConnection("example", new()
    ///     {
    ///         EndpointId = exampleVpcEndpoint.Id,
    ///         ServiceId = exampleVpcEndpoint.ServiceId,
    ///         Bandwidth = 1024,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Private Link Vpc Endpoint Connection can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection example &lt;service_id&gt;:&lt;endpoint_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection")]
    public partial class VpcEndpointServiceConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

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
        /// The endpoint service ID.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// The state of the endpoint connection.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpointServiceConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpointServiceConnection(string name, VpcEndpointServiceConnectionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection", name, args ?? new VpcEndpointServiceConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpointServiceConnection(string name, Input<string> id, VpcEndpointServiceConnectionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpointServiceConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpointServiceConnection Get(string name, Input<string> id, VpcEndpointServiceConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpointServiceConnection(name, id, state, options);
        }
    }

    public sealed class VpcEndpointServiceConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

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
        /// The endpoint service ID.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public VpcEndpointServiceConnectionArgs()
        {
        }
        public static new VpcEndpointServiceConnectionArgs Empty => new VpcEndpointServiceConnectionArgs();
    }

    public sealed class VpcEndpointServiceConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

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
        /// The endpoint service ID.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// The state of the endpoint connection.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public VpcEndpointServiceConnectionState()
        {
        }
        public static new VpcEndpointServiceConnectionState Empty => new VpcEndpointServiceConnectionState();
    }
}
