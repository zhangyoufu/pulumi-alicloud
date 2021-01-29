// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PolarDB
{
    /// <summary>
    /// Provides a PolarDB endpoint resource to manage custom endpoint of PolarDB cluster.
    /// 
    /// &gt; **NOTE:** Available in v1.80.0+. Only used to manage PolarDB MySQL custom cluster endpoint.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var creation = config.Get("creation") ?? "PolarDB";
    ///         var name = config.Get("name") ?? "polardbconnectionbasic";
    ///         var defaultZones = Output.Create(AliCloud.GetZones.InvokeAsync(new AliCloud.GetZonesArgs
    ///         {
    ///             AvailableResourceCreation = creation,
    ///         }));
    ///         var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/16",
    ///         });
    ///         var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new AliCloud.Vpc.SwitchArgs
    ///         {
    ///             VpcId = defaultNetwork.Id,
    ///             CidrBlock = "172.16.0.0/24",
    ///             AvailabilityZone = defaultZones.Apply(defaultZones =&gt; defaultZones.Zones[0].Id),
    ///         });
    ///         var defaultCluster = new AliCloud.PolarDB.Cluster("defaultCluster", new AliCloud.PolarDB.ClusterArgs
    ///         {
    ///             DbType = "MySQL",
    ///             DbVersion = "8.0",
    ///             PayType = "PostPaid",
    ///             DbNodeClass = "polar.mysql.x4.large",
    ///             VswitchId = defaultSwitch.Id,
    ///             Description = name,
    ///         });
    ///         var endpoint = new AliCloud.PolarDB.Endpoint("endpoint", new AliCloud.PolarDB.EndpointArgs
    ///         {
    ///             DbClusterId = defaultCluster.Id,
    ///             EndpointType = "Custom",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// PolarDB endpoint can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:polardb/endpoint:Endpoint example pc-abc123456:pe-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:polardb/endpoint:Endpoint")]
    public partial class Endpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. Default to `Disable`.
        /// </summary>
        [Output("autoAddNewNodes")]
        public Output<string?> AutoAddNewNodes { get; private set; } = null!;

        /// <summary>
        /// The Id of cluster that can run database.
        /// </summary>
        [Output("dbClusterId")]
        public Output<string> DbClusterId { get; private set; } = null!;

        /// <summary>
        /// The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
        /// </summary>
        [Output("endpointConfig")]
        public Output<ImmutableDictionary<string, object>> EndpointConfig { get; private set; } = null!;

        /// <summary>
        /// Type of endpoint. Valid value: `Custom`. Currently supported only `Custom`.
        /// </summary>
        [Output("endpointType")]
        public Output<string> EndpointType { get; private set; } = null!;

        /// <summary>
        /// Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
        /// </summary>
        [Output("nodes")]
        public Output<ImmutableArray<string>> Nodes { get; private set; } = null!;

        /// <summary>
        /// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. Default to `ReadOnly`.
        /// </summary>
        [Output("readWriteMode")]
        public Output<string?> ReadWriteMode { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("alicloud:polardb/endpoint:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:polardb/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, state, options);
        }
    }

    public sealed class EndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. Default to `Disable`.
        /// </summary>
        [Input("autoAddNewNodes")]
        public Input<string>? AutoAddNewNodes { get; set; }

        /// <summary>
        /// The Id of cluster that can run database.
        /// </summary>
        [Input("dbClusterId", required: true)]
        public Input<string> DbClusterId { get; set; } = null!;

        [Input("endpointConfig")]
        private InputMap<object>? _endpointConfig;

        /// <summary>
        /// The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
        /// </summary>
        public InputMap<object> EndpointConfig
        {
            get => _endpointConfig ?? (_endpointConfig = new InputMap<object>());
            set => _endpointConfig = value;
        }

        /// <summary>
        /// Type of endpoint. Valid value: `Custom`. Currently supported only `Custom`.
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        [Input("nodes")]
        private InputList<string>? _nodes;

        /// <summary>
        /// Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
        /// </summary>
        public InputList<string> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<string>());
            set => _nodes = value;
        }

        /// <summary>
        /// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. Default to `ReadOnly`.
        /// </summary>
        [Input("readWriteMode")]
        public Input<string>? ReadWriteMode { get; set; }

        public EndpointArgs()
        {
        }
    }

    public sealed class EndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. Default to `Disable`.
        /// </summary>
        [Input("autoAddNewNodes")]
        public Input<string>? AutoAddNewNodes { get; set; }

        /// <summary>
        /// The Id of cluster that can run database.
        /// </summary>
        [Input("dbClusterId")]
        public Input<string>? DbClusterId { get; set; }

        [Input("endpointConfig")]
        private InputMap<object>? _endpointConfig;

        /// <summary>
        /// The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
        /// </summary>
        public InputMap<object> EndpointConfig
        {
            get => _endpointConfig ?? (_endpointConfig = new InputMap<object>());
            set => _endpointConfig = value;
        }

        /// <summary>
        /// Type of endpoint. Valid value: `Custom`. Currently supported only `Custom`.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        [Input("nodes")]
        private InputList<string>? _nodes;

        /// <summary>
        /// Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
        /// </summary>
        public InputList<string> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<string>());
            set => _nodes = value;
        }

        /// <summary>
        /// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. Default to `ReadOnly`.
        /// </summary>
        [Input("readWriteMode")]
        public Input<string>? ReadWriteMode { get; set; }

        public EndpointState()
        {
        }
    }
}
