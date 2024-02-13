// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nlb
{
    /// <summary>
    /// Provides a NLB Server Group resource.
    /// 
    /// For information about NLB Server Group and how to use it, see [What is Server Group](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createservergroup-nlb).
    /// 
    /// &gt; **NOTE:** Available since v1.186.0.
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
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.4.0.0/16",
    ///     });
    /// 
    ///     var defaultServerGroup = new AliCloud.Nlb.ServerGroup("defaultServerGroup", new()
    ///     {
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         ServerGroupName = name,
    ///         ServerGroupType = "Instance",
    ///         VpcId = defaultNetwork.Id,
    ///         Scheduler = "Wrr",
    ///         Protocol = "TCP",
    ///         ConnectionDrain = true,
    ///         ConnectionDrainTimeout = 60,
    ///         AddressIpVersion = "Ipv4",
    ///         HealthCheck = new AliCloud.Nlb.Inputs.ServerGroupHealthCheckArgs
    ///         {
    ///             HealthCheckEnabled = true,
    ///             HealthCheckType = "TCP",
    ///             HealthCheckConnectPort = 0,
    ///             HealthyThreshold = 2,
    ///             UnhealthyThreshold = 2,
    ///             HealthCheckConnectTimeout = 5,
    ///             HealthCheckInterval = 10,
    ///             HttpCheckMethod = "GET",
    ///             HealthCheckHttpCodes = new[]
    ///             {
    ///                 "http_2xx",
    ///                 "http_3xx",
    ///                 "http_4xx",
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "example" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// NLB Server Group can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:nlb/serverGroup:ServerGroup example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:nlb/serverGroup:ServerGroup")]
    public partial class ServerGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Protocol version. Value:
        /// - **ipv4**:IPv4 type.
        /// - **DualStack**: Double Stack type.
        /// </summary>
        [Output("addressIpVersion")]
        public Output<string> AddressIpVersion { get; private set; } = null!;

        /// <summary>
        /// Full port forwarding.
        /// </summary>
        [Output("anyPortEnabled")]
        public Output<bool> AnyPortEnabled { get; private set; } = null!;

        /// <summary>
        /// . Field 'connection_drain' has been deprecated from provider version 1.214.0. New field 'connection_drain_enabled' instead.
        /// </summary>
        [Output("connectionDrain")]
        public Output<bool> ConnectionDrain { get; private set; } = null!;

        /// <summary>
        /// Whether to open the connection gracefully interrupted. Value:
        /// - **true**: on.
        /// - **false**: closed.
        /// </summary>
        [Output("connectionDrainEnabled")]
        public Output<bool> ConnectionDrainEnabled { get; private set; } = null!;

        /// <summary>
        /// Set the connection elegant interrupt timeout. Unit: seconds. Valid values: **10** ~ **900**.
        /// </summary>
        [Output("connectionDrainTimeout")]
        public Output<int> ConnectionDrainTimeout { get; private set; } = null!;

        /// <summary>
        /// Health check configuration information. See `health_check` below.
        /// </summary>
        [Output("healthCheck")]
        public Output<Outputs.ServerGroupHealthCheck> HealthCheck { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the client address retention function. Value:
        /// - **true**: on.
        /// - **false**: closed.
        /// &gt; **NOTE:**  special instructions: When **AddressIPVersion** is of the **ipv4** type, the default value is **true**. **Addrestipversion** can only be **false** when the value of **ipv6** is **ipv6**, and can be **true** when supported by the underlying layer * *.
        /// </summary>
        [Output("preserveClientIpEnabled")]
        public Output<bool> PreserveClientIpEnabled { get; private set; } = null!;

        /// <summary>
        /// The backend Forwarding Protocol. Valid values: **TCP**, **UDP**, or **TCPSSL**.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group to which the server group belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// Scheduling algorithm. Value:
        /// - **Wrr**: Weighted polling. The higher the weight of the backend server, the higher the probability of being polled.
        /// - **Rr**: polls external requests are distributed to backend servers in sequence according to the access order. sch: Source IP hash: The same source address is scheduled to the same backend server.
        /// - **Tch**: Quadruple hash, based on the consistent hash of the Quad (source IP, Destination IP, source port, and destination port), the same stream is scheduled to the same backend server.
        /// - **Qch**: a QUIC ID hash that allows you to hash requests with the same QUIC ID to the same backend server.
        /// </summary>
        [Output("scheduler")]
        public Output<string> Scheduler { get; private set; } = null!;

        /// <summary>
        /// The name of the server group.
        /// </summary>
        [Output("serverGroupName")]
        public Output<string> ServerGroupName { get; private set; } = null!;

        /// <summary>
        /// Server group type. Value:
        /// - **Instance**: The server type. You can add **Ecs**, **Ens**, and **Eci** instances to the server group.
        /// - **Ip**: Ip address type. You can add Ip addresses to a server group of this type.
        /// </summary>
        [Output("serverGroupType")]
        public Output<string> ServerGroupType { get; private set; } = null!;

        /// <summary>
        /// Server group status. Value:
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Label.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC to which the server group belongs.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a ServerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerGroup(string name, ServerGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:nlb/serverGroup:ServerGroup", name, args ?? new ServerGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerGroup(string name, Input<string> id, ServerGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:nlb/serverGroup:ServerGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerGroup Get(string name, Input<string> id, ServerGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerGroup(name, id, state, options);
        }
    }

    public sealed class ServerGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Protocol version. Value:
        /// - **ipv4**:IPv4 type.
        /// - **DualStack**: Double Stack type.
        /// </summary>
        [Input("addressIpVersion")]
        public Input<string>? AddressIpVersion { get; set; }

        /// <summary>
        /// Full port forwarding.
        /// </summary>
        [Input("anyPortEnabled")]
        public Input<bool>? AnyPortEnabled { get; set; }

        /// <summary>
        /// . Field 'connection_drain' has been deprecated from provider version 1.214.0. New field 'connection_drain_enabled' instead.
        /// </summary>
        [Input("connectionDrain")]
        public Input<bool>? ConnectionDrain { get; set; }

        /// <summary>
        /// Whether to open the connection gracefully interrupted. Value:
        /// - **true**: on.
        /// - **false**: closed.
        /// </summary>
        [Input("connectionDrainEnabled")]
        public Input<bool>? ConnectionDrainEnabled { get; set; }

        /// <summary>
        /// Set the connection elegant interrupt timeout. Unit: seconds. Valid values: **10** ~ **900**.
        /// </summary>
        [Input("connectionDrainTimeout")]
        public Input<int>? ConnectionDrainTimeout { get; set; }

        /// <summary>
        /// Health check configuration information. See `health_check` below.
        /// </summary>
        [Input("healthCheck")]
        public Input<Inputs.ServerGroupHealthCheckArgs>? HealthCheck { get; set; }

        /// <summary>
        /// Whether to enable the client address retention function. Value:
        /// - **true**: on.
        /// - **false**: closed.
        /// &gt; **NOTE:**  special instructions: When **AddressIPVersion** is of the **ipv4** type, the default value is **true**. **Addrestipversion** can only be **false** when the value of **ipv6** is **ipv6**, and can be **true** when supported by the underlying layer * *.
        /// </summary>
        [Input("preserveClientIpEnabled")]
        public Input<bool>? PreserveClientIpEnabled { get; set; }

        /// <summary>
        /// The backend Forwarding Protocol. Valid values: **TCP**, **UDP**, or **TCPSSL**.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The ID of the resource group to which the server group belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Scheduling algorithm. Value:
        /// - **Wrr**: Weighted polling. The higher the weight of the backend server, the higher the probability of being polled.
        /// - **Rr**: polls external requests are distributed to backend servers in sequence according to the access order. sch: Source IP hash: The same source address is scheduled to the same backend server.
        /// - **Tch**: Quadruple hash, based on the consistent hash of the Quad (source IP, Destination IP, source port, and destination port), the same stream is scheduled to the same backend server.
        /// - **Qch**: a QUIC ID hash that allows you to hash requests with the same QUIC ID to the same backend server.
        /// </summary>
        [Input("scheduler")]
        public Input<string>? Scheduler { get; set; }

        /// <summary>
        /// The name of the server group.
        /// </summary>
        [Input("serverGroupName", required: true)]
        public Input<string> ServerGroupName { get; set; } = null!;

        /// <summary>
        /// Server group type. Value:
        /// - **Instance**: The server type. You can add **Ecs**, **Ens**, and **Eci** instances to the server group.
        /// - **Ip**: Ip address type. You can add Ip addresses to a server group of this type.
        /// </summary>
        [Input("serverGroupType")]
        public Input<string>? ServerGroupType { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Label.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VPC to which the server group belongs.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public ServerGroupArgs()
        {
        }
        public static new ServerGroupArgs Empty => new ServerGroupArgs();
    }

    public sealed class ServerGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Protocol version. Value:
        /// - **ipv4**:IPv4 type.
        /// - **DualStack**: Double Stack type.
        /// </summary>
        [Input("addressIpVersion")]
        public Input<string>? AddressIpVersion { get; set; }

        /// <summary>
        /// Full port forwarding.
        /// </summary>
        [Input("anyPortEnabled")]
        public Input<bool>? AnyPortEnabled { get; set; }

        /// <summary>
        /// . Field 'connection_drain' has been deprecated from provider version 1.214.0. New field 'connection_drain_enabled' instead.
        /// </summary>
        [Input("connectionDrain")]
        public Input<bool>? ConnectionDrain { get; set; }

        /// <summary>
        /// Whether to open the connection gracefully interrupted. Value:
        /// - **true**: on.
        /// - **false**: closed.
        /// </summary>
        [Input("connectionDrainEnabled")]
        public Input<bool>? ConnectionDrainEnabled { get; set; }

        /// <summary>
        /// Set the connection elegant interrupt timeout. Unit: seconds. Valid values: **10** ~ **900**.
        /// </summary>
        [Input("connectionDrainTimeout")]
        public Input<int>? ConnectionDrainTimeout { get; set; }

        /// <summary>
        /// Health check configuration information. See `health_check` below.
        /// </summary>
        [Input("healthCheck")]
        public Input<Inputs.ServerGroupHealthCheckGetArgs>? HealthCheck { get; set; }

        /// <summary>
        /// Whether to enable the client address retention function. Value:
        /// - **true**: on.
        /// - **false**: closed.
        /// &gt; **NOTE:**  special instructions: When **AddressIPVersion** is of the **ipv4** type, the default value is **true**. **Addrestipversion** can only be **false** when the value of **ipv6** is **ipv6**, and can be **true** when supported by the underlying layer * *.
        /// </summary>
        [Input("preserveClientIpEnabled")]
        public Input<bool>? PreserveClientIpEnabled { get; set; }

        /// <summary>
        /// The backend Forwarding Protocol. Valid values: **TCP**, **UDP**, or **TCPSSL**.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The ID of the resource group to which the server group belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Scheduling algorithm. Value:
        /// - **Wrr**: Weighted polling. The higher the weight of the backend server, the higher the probability of being polled.
        /// - **Rr**: polls external requests are distributed to backend servers in sequence according to the access order. sch: Source IP hash: The same source address is scheduled to the same backend server.
        /// - **Tch**: Quadruple hash, based on the consistent hash of the Quad (source IP, Destination IP, source port, and destination port), the same stream is scheduled to the same backend server.
        /// - **Qch**: a QUIC ID hash that allows you to hash requests with the same QUIC ID to the same backend server.
        /// </summary>
        [Input("scheduler")]
        public Input<string>? Scheduler { get; set; }

        /// <summary>
        /// The name of the server group.
        /// </summary>
        [Input("serverGroupName")]
        public Input<string>? ServerGroupName { get; set; }

        /// <summary>
        /// Server group type. Value:
        /// - **Instance**: The server type. You can add **Ecs**, **Ens**, and **Eci** instances to the server group.
        /// - **Ip**: Ip address type. You can add Ip addresses to a server group of this type.
        /// </summary>
        [Input("serverGroupType")]
        public Input<string>? ServerGroupType { get; set; }

        /// <summary>
        /// Server group status. Value:
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Label.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VPC to which the server group belongs.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ServerGroupState()
        {
        }
        public static new ServerGroupState Empty => new ServerGroupState();
    }
}
