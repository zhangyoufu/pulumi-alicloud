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
    /// &gt; **NOTE:** Available since v1.80.0.
    /// **NOTE:** After v1.80.0 and before v1.121.0, you can only use this resource to manage the custom endpoint. Since v1.121.0, you also can import the primary endpoint and the cluster endpoint, to modify their ssl status and so on.
    /// **NOTE:** The primary endpoint and the default cluster endpoint can not be created or deleted manually.
    /// 
    /// ## Example Usage
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
    ///     var @default = AliCloud.PolarDB.GetNodeClasses.Invoke(new()
    ///     {
    ///         DbType = "MySQL",
    ///         DbVersion = "8.0",
    ///         PayType = "PostPaid",
    ///         Category = "Normal",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
    ///     {
    ///         VpcName = "terraform-example",
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("default", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getNodeClassesResult =&gt; getNodeClassesResult.Classes[0]?.ZoneId)),
    ///         VswitchName = "terraform-example",
    ///     });
    /// 
    ///     var defaultCluster = new AliCloud.PolarDB.Cluster("default", new()
    ///     {
    ///         DbType = "MySQL",
    ///         DbVersion = "8.0",
    ///         DbNodeClass = @default.Apply(@default =&gt; @default.Apply(getNodeClassesResult =&gt; getNodeClassesResult.Classes[0]?.SupportedEngines[0]?.AvailableResources[0]?.DbNodeClass)),
    ///         PayType = "PostPaid",
    ///         VswitchId = defaultSwitch.Id,
    ///         Description = "terraform-example",
    ///     });
    /// 
    ///     var defaultEndpoint = new AliCloud.PolarDB.Endpoint("default", new()
    ///     {
    ///         DbClusterId = defaultCluster.Id,
    ///         EndpointType = "Custom",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// PolarDB endpoint can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:polardb/endpoint:Endpoint example pc-abc123456:pe-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:polardb/endpoint:Endpoint")]
    public partial class Endpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
        /// </summary>
        [Output("autoAddNewNodes")]
        public Output<string> AutoAddNewNodes { get; private set; } = null!;

        /// <summary>
        /// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
        /// </summary>
        [Output("connectionPrefix")]
        public Output<string> ConnectionPrefix { get; private set; } = null!;

        /// <summary>
        /// The Id of cluster that can run database.
        /// </summary>
        [Output("dbClusterId")]
        public Output<string> DbClusterId { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Output("dbEndpointDescription")]
        public Output<string?> DbEndpointDescription { get; private set; } = null!;

        /// <summary>
        /// (Available since v1.161.0) The ID of the cluster endpoint.
        /// </summary>
        [Output("dbEndpointId")]
        public Output<string> DbEndpointId { get; private set; } = null!;

        /// <summary>
        /// The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
        /// </summary>
        [Output("endpointConfig")]
        public Output<ImmutableDictionary<string, object>> EndpointConfig { get; private set; } = null!;

        /// <summary>
        /// Type of the endpoint. Before v1.121.0, it only can be `Custom`. since v1.121.0, `Custom`, `Cluster`, `Primary` are valid, default to `Custom`. However when creating a new endpoint, it also only can be `Custom`.
        /// </summary>
        [Output("endpointType")]
        public Output<string?> EndpointType { get; private set; } = null!;

        /// <summary>
        /// The network type of the endpoint address.
        /// </summary>
        [Output("netType")]
        public Output<string?> NetType { get; private set; } = null!;

        /// <summary>
        /// Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
        /// </summary>
        [Output("nodes")]
        public Output<ImmutableArray<string>> Nodes { get; private set; } = null!;

        /// <summary>
        /// Port of the specified endpoint. Valid values: 3000 to 5999.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
        /// </summary>
        [Output("readWriteMode")]
        public Output<string> ReadWriteMode { get; private set; } = null!;

        /// <summary>
        /// Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
        /// </summary>
        [Output("sslAutoRotate")]
        public Output<string?> SslAutoRotate { get; private set; } = null!;

        /// <summary>
        /// Specifies SSL certificate download link.  
        /// **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
        /// For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
        /// </summary>
        [Output("sslCertificateUrl")]
        public Output<string> SslCertificateUrl { get; private set; } = null!;

        /// <summary>
        /// (Available since v1.121.0) The SSL connection string.
        /// </summary>
        [Output("sslConnectionString")]
        public Output<string> SslConnectionString { get; private set; } = null!;

        /// <summary>
        /// Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
        /// </summary>
        [Output("sslEnabled")]
        public Output<string?> SslEnabled { get; private set; } = null!;

        /// <summary>
        /// (Available since v1.121.0) The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
        /// </summary>
        [Output("sslExpireTime")]
        public Output<string> SslExpireTime { get; private set; } = null!;


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

    public sealed class EndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
        /// </summary>
        [Input("autoAddNewNodes")]
        public Input<string>? AutoAddNewNodes { get; set; }

        /// <summary>
        /// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
        /// </summary>
        [Input("connectionPrefix")]
        public Input<string>? ConnectionPrefix { get; set; }

        /// <summary>
        /// The Id of cluster that can run database.
        /// </summary>
        [Input("dbClusterId", required: true)]
        public Input<string> DbClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Input("dbEndpointDescription")]
        public Input<string>? DbEndpointDescription { get; set; }

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
        /// Type of the endpoint. Before v1.121.0, it only can be `Custom`. since v1.121.0, `Custom`, `Cluster`, `Primary` are valid, default to `Custom`. However when creating a new endpoint, it also only can be `Custom`.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// The network type of the endpoint address.
        /// </summary>
        [Input("netType")]
        public Input<string>? NetType { get; set; }

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
        /// Port of the specified endpoint. Valid values: 3000 to 5999.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
        /// </summary>
        [Input("readWriteMode")]
        public Input<string>? ReadWriteMode { get; set; }

        /// <summary>
        /// Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
        /// </summary>
        [Input("sslAutoRotate")]
        public Input<string>? SslAutoRotate { get; set; }

        /// <summary>
        /// Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
        /// </summary>
        [Input("sslEnabled")]
        public Input<string>? SslEnabled { get; set; }

        public EndpointArgs()
        {
        }
        public static new EndpointArgs Empty => new EndpointArgs();
    }

    public sealed class EndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
        /// </summary>
        [Input("autoAddNewNodes")]
        public Input<string>? AutoAddNewNodes { get; set; }

        /// <summary>
        /// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
        /// </summary>
        [Input("connectionPrefix")]
        public Input<string>? ConnectionPrefix { get; set; }

        /// <summary>
        /// The Id of cluster that can run database.
        /// </summary>
        [Input("dbClusterId")]
        public Input<string>? DbClusterId { get; set; }

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Input("dbEndpointDescription")]
        public Input<string>? DbEndpointDescription { get; set; }

        /// <summary>
        /// (Available since v1.161.0) The ID of the cluster endpoint.
        /// </summary>
        [Input("dbEndpointId")]
        public Input<string>? DbEndpointId { get; set; }

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
        /// Type of the endpoint. Before v1.121.0, it only can be `Custom`. since v1.121.0, `Custom`, `Cluster`, `Primary` are valid, default to `Custom`. However when creating a new endpoint, it also only can be `Custom`.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// The network type of the endpoint address.
        /// </summary>
        [Input("netType")]
        public Input<string>? NetType { get; set; }

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
        /// Port of the specified endpoint. Valid values: 3000 to 5999.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
        /// </summary>
        [Input("readWriteMode")]
        public Input<string>? ReadWriteMode { get; set; }

        /// <summary>
        /// Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
        /// </summary>
        [Input("sslAutoRotate")]
        public Input<string>? SslAutoRotate { get; set; }

        /// <summary>
        /// Specifies SSL certificate download link.  
        /// **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
        /// For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
        /// </summary>
        [Input("sslCertificateUrl")]
        public Input<string>? SslCertificateUrl { get; set; }

        /// <summary>
        /// (Available since v1.121.0) The SSL connection string.
        /// </summary>
        [Input("sslConnectionString")]
        public Input<string>? SslConnectionString { get; set; }

        /// <summary>
        /// Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
        /// </summary>
        [Input("sslEnabled")]
        public Input<string>? SslEnabled { get; set; }

        /// <summary>
        /// (Available since v1.121.0) The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
        /// </summary>
        [Input("sslExpireTime")]
        public Input<string>? SslExpireTime { get; set; }

        public EndpointState()
        {
        }
        public static new EndpointState Empty => new EndpointState();
    }
}
