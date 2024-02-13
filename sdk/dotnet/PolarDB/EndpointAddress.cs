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
    /// Provides a PolarDB endpoint address resource to allocate an Internet endpoint address string for PolarDB instance.
    /// 
    /// &gt; **NOTE:** Available since v1.68.0. Each PolarDB instance will allocate a intranet connection string automatically and its prefix is Cluster ID.
    ///  To avoid unnecessary conflict, please specified a internet connection prefix before applying the resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultNodeClasses = AliCloud.PolarDB.GetNodeClasses.Invoke(new()
    ///     {
    ///         DbType = "MySQL",
    ///         DbVersion = "8.0",
    ///         PayType = "PostPaid",
    ///         Category = "Normal",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = "terraform-example",
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = defaultNodeClasses.Apply(getNodeClassesResult =&gt; getNodeClassesResult.Classes[0]?.ZoneId),
    ///         VswitchName = "terraform-example",
    ///     });
    /// 
    ///     var defaultCluster = new AliCloud.PolarDB.Cluster("defaultCluster", new()
    ///     {
    ///         DbType = "MySQL",
    ///         DbVersion = "8.0",
    ///         DbNodeClass = defaultNodeClasses.Apply(getNodeClassesResult =&gt; getNodeClassesResult.Classes[0]?.SupportedEngines[0]?.AvailableResources[0]?.DbNodeClass),
    ///         PayType = "PostPaid",
    ///         VswitchId = defaultSwitch.Id,
    ///         Description = "terraform-example",
    ///     });
    /// 
    ///     var defaultEndpoints = AliCloud.PolarDB.GetEndpoints.Invoke(new()
    ///     {
    ///         DbClusterId = defaultCluster.Id,
    ///     });
    /// 
    ///     var defaultEndpointAddress = new AliCloud.PolarDB.EndpointAddress("defaultEndpointAddress", new()
    ///     {
    ///         DbClusterId = defaultCluster.Id,
    ///         DbEndpointId = defaultEndpoints.Apply(getEndpointsResult =&gt; getEndpointsResult.Endpoints[0]?.DbEndpointId),
    ///         ConnectionPrefix = "polardbexample",
    ///         NetType = "Public",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// PolarDB endpoint address can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:polardb/endpointAddress:EndpointAddress example pc-abc123456:pe-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:polardb/endpointAddress:EndpointAddress")]
    public partial class EndpointAddress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
        /// </summary>
        [Output("connectionPrefix")]
        public Output<string> ConnectionPrefix { get; private set; } = null!;

        /// <summary>
        /// Connection cluster or endpoint string.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// The Id of cluster that can run database.
        /// </summary>
        [Output("dbClusterId")]
        public Output<string> DbClusterId { get; private set; } = null!;

        /// <summary>
        /// The Id of endpoint that can run database.
        /// </summary>
        [Output("dbEndpointId")]
        public Output<string> DbEndpointId { get; private set; } = null!;

        /// <summary>
        /// The ip address of connection string.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
        /// </summary>
        [Output("netType")]
        public Output<string?> NetType { get; private set; } = null!;

        /// <summary>
        /// Port of the specified endpoint. Valid values: 3000 to 5999.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointAddress(string name, EndpointAddressArgs args, CustomResourceOptions? options = null)
            : base("alicloud:polardb/endpointAddress:EndpointAddress", name, args ?? new EndpointAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointAddress(string name, Input<string> id, EndpointAddressState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:polardb/endpointAddress:EndpointAddress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EndpointAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointAddress Get(string name, Input<string> id, EndpointAddressState? state = null, CustomResourceOptions? options = null)
        {
            return new EndpointAddress(name, id, state, options);
        }
    }

    public sealed class EndpointAddressArgs : global::Pulumi.ResourceArgs
    {
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
        /// The Id of endpoint that can run database.
        /// </summary>
        [Input("dbEndpointId", required: true)]
        public Input<string> DbEndpointId { get; set; } = null!;

        /// <summary>
        /// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
        /// </summary>
        [Input("netType")]
        public Input<string>? NetType { get; set; }

        /// <summary>
        /// Port of the specified endpoint. Valid values: 3000 to 5999.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        public EndpointAddressArgs()
        {
        }
        public static new EndpointAddressArgs Empty => new EndpointAddressArgs();
    }

    public sealed class EndpointAddressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
        /// </summary>
        [Input("connectionPrefix")]
        public Input<string>? ConnectionPrefix { get; set; }

        /// <summary>
        /// Connection cluster or endpoint string.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The Id of cluster that can run database.
        /// </summary>
        [Input("dbClusterId")]
        public Input<string>? DbClusterId { get; set; }

        /// <summary>
        /// The Id of endpoint that can run database.
        /// </summary>
        [Input("dbEndpointId")]
        public Input<string>? DbEndpointId { get; set; }

        /// <summary>
        /// The ip address of connection string.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
        /// </summary>
        [Input("netType")]
        public Input<string>? NetType { get; set; }

        /// <summary>
        /// Port of the specified endpoint. Valid values: 3000 to 5999.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        public EndpointAddressState()
        {
        }
        public static new EndpointAddressState Empty => new EndpointAddressState();
    }
}
