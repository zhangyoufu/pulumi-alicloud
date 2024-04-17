// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds
{
    /// <summary>
    /// Provide RDS cluster instance endpoint public connection resources.
    /// 
    /// Information about RDS MySQL cluster endpoint address and how to use it, see [What is RDS DB Instance Endpoint Address](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-createdbinstanceendpointaddress).
    /// 
    /// &gt; **NOTE:** Available since v1.204.0.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var @default = AliCloud.Rds.GetZones.Invoke(new()
    ///     {
    ///         Engine = "MySQL",
    ///         EngineVersion = "8.0",
    ///         InstanceChargeType = "PostPaid",
    ///         Category = "cluster",
    ///         DbInstanceStorageType = "cloud_essd",
    ///     });
    /// 
    ///     var defaultGetInstanceClasses = AliCloud.Rds.GetInstanceClasses.Invoke(new()
    ///     {
    ///         ZoneId = @default.Apply(getZonesResult =&gt; getZonesResult.Ids[0]),
    ///         Engine = "MySQL",
    ///         EngineVersion = "8.0",
    ///         Category = "cluster",
    ///         DbInstanceStorageType = "cloud_essd",
    ///         InstanceChargeType = "PostPaid",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("default", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Ids[0])),
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("default", new()
    ///     {
    ///         Name = name,
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Rds.Instance("default", new()
    ///     {
    ///         Engine = "MySQL",
    ///         EngineVersion = "8.0",
    ///         InstanceType = defaultGetInstanceClasses.Apply(getInstanceClassesResult =&gt; getInstanceClassesResult.InstanceClasses[0]?.InstanceClass),
    ///         InstanceStorage = defaultGetInstanceClasses.Apply(getInstanceClassesResult =&gt; getInstanceClassesResult.InstanceClasses[0]?.StorageRange?.Min),
    ///         InstanceChargeType = "Postpaid",
    ///         InstanceName = name,
    ///         VswitchId = defaultSwitch.Id,
    ///         MonitoringPeriod = 60,
    ///         DbInstanceStorageType = "cloud_essd",
    ///         SecurityGroupIds = new[]
    ///         {
    ///             defaultSecurityGroup.Id,
    ///         },
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Ids[0])),
    ///         ZoneIdSlaveA = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Ids[0])),
    ///     });
    /// 
    ///     var defaultDbNode = new AliCloud.Rds.DbNode("default", new()
    ///     {
    ///         DbInstanceId = defaultInstance.Id,
    ///         ClassCode = defaultInstance.InstanceType,
    ///         ZoneId = defaultSwitch.ZoneId,
    ///     });
    /// 
    ///     var defaultDbInstanceEndpoint = new AliCloud.Rds.DbInstanceEndpoint("default", new()
    ///     {
    ///         DbInstanceId = defaultDbNode.DbInstanceId,
    ///         VpcId = defaultNetwork.Id,
    ///         VswitchId = defaultInstance.VswitchId,
    ///         ConnectionStringPrefix = "example",
    ///         Port = "3306",
    ///         DbInstanceEndpointDescription = name,
    ///         NodeItems = new[]
    ///         {
    ///             new AliCloud.Rds.Inputs.DbInstanceEndpointNodeItemArgs
    ///             {
    ///                 NodeId = defaultDbNode.NodeId,
    ///                 Weight = 25,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultDbInstanceEndpointAddress = new AliCloud.Rds.DbInstanceEndpointAddress("default", new()
    ///     {
    ///         DbInstanceId = defaultInstance.Id,
    ///         DbInstanceEndpointId = defaultDbInstanceEndpoint.DbInstanceEndpointId,
    ///         ConnectionStringPrefix = "tf-example-prefix",
    ///         Port = "3306",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// RDS database endpoint public address feature can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:rds/dbInstanceEndpointAddress:DbInstanceEndpointAddress example &lt;db_instance_id&gt;:&lt;db_instance_endpoint_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rds/dbInstanceEndpointAddress:DbInstanceEndpointAddress")]
    public partial class DbInstanceEndpointAddress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The endpoint of the instance.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// The prefix of the public endpoint.
        /// </summary>
        [Output("connectionStringPrefix")]
        public Output<string> ConnectionStringPrefix { get; private set; } = null!;

        /// <summary>
        /// The Endpoint ID of the instance.
        /// </summary>
        [Output("dbInstanceEndpointId")]
        public Output<string> DbInstanceEndpointId { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// The IP address of the endpoint.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The type of the IP address.
        /// </summary>
        [Output("ipType")]
        public Output<string> IpType { get; private set; } = null!;

        /// <summary>
        /// The port number of the public endpoint.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;


        /// <summary>
        /// Create a DbInstanceEndpointAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DbInstanceEndpointAddress(string name, DbInstanceEndpointAddressArgs args, CustomResourceOptions? options = null)
            : base("alicloud:rds/dbInstanceEndpointAddress:DbInstanceEndpointAddress", name, args ?? new DbInstanceEndpointAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DbInstanceEndpointAddress(string name, Input<string> id, DbInstanceEndpointAddressState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rds/dbInstanceEndpointAddress:DbInstanceEndpointAddress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DbInstanceEndpointAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DbInstanceEndpointAddress Get(string name, Input<string> id, DbInstanceEndpointAddressState? state = null, CustomResourceOptions? options = null)
        {
            return new DbInstanceEndpointAddress(name, id, state, options);
        }
    }

    public sealed class DbInstanceEndpointAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The prefix of the public endpoint.
        /// </summary>
        [Input("connectionStringPrefix", required: true)]
        public Input<string> ConnectionStringPrefix { get; set; } = null!;

        /// <summary>
        /// The Endpoint ID of the instance.
        /// </summary>
        [Input("dbInstanceEndpointId", required: true)]
        public Input<string> DbInstanceEndpointId { get; set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        /// <summary>
        /// The port number of the public endpoint.
        /// </summary>
        [Input("port", required: true)]
        public Input<string> Port { get; set; } = null!;

        public DbInstanceEndpointAddressArgs()
        {
        }
        public static new DbInstanceEndpointAddressArgs Empty => new DbInstanceEndpointAddressArgs();
    }

    public sealed class DbInstanceEndpointAddressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint of the instance.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The prefix of the public endpoint.
        /// </summary>
        [Input("connectionStringPrefix")]
        public Input<string>? ConnectionStringPrefix { get; set; }

        /// <summary>
        /// The Endpoint ID of the instance.
        /// </summary>
        [Input("dbInstanceEndpointId")]
        public Input<string>? DbInstanceEndpointId { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// The IP address of the endpoint.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The type of the IP address.
        /// </summary>
        [Input("ipType")]
        public Input<string>? IpType { get; set; }

        /// <summary>
        /// The port number of the public endpoint.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        public DbInstanceEndpointAddressState()
        {
        }
        public static new DbInstanceEndpointAddressState Empty => new DbInstanceEndpointAddressState();
    }
}
