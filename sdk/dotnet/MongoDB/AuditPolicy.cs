// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MongoDB
{
    /// <summary>
    /// Provides a MongoDB Audit Policy resource.
    /// 
    /// For information about MongoDB Audit Policy and how to use it, see [What is Audit Policy](https://www.alibabacloud.com/help/doc-detail/131941.html).
    /// 
    /// &gt; **NOTE:** Available since v1.148.0.
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
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultZones = AliCloud.MongoDB.GetZones.Invoke();
    /// 
    ///     var index = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones).Length.Apply(length =&gt; length - 1);
    /// 
    ///     var zoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones)[index].Id;
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.17.3.0/24",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "172.17.3.0/24",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = zoneId,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.MongoDB.Instance("defaultInstance", new()
    ///     {
    ///         EngineVersion = "4.2",
    ///         DbInstanceClass = "dds.mongo.mid",
    ///         DbInstanceStorage = 10,
    ///         VswitchId = defaultSwitch.Id,
    ///         SecurityIpLists = new[]
    ///         {
    ///             "10.168.1.12",
    ///             "100.69.7.112",
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "example" },
    ///         },
    ///     });
    /// 
    ///     var defaultAuditPolicy = new AliCloud.MongoDB.AuditPolicy("defaultAuditPolicy", new()
    ///     {
    ///         DbInstanceId = defaultInstance.Id,
    ///         AuditStatus = "disabled",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// MongoDB Audit Policy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:mongodb/auditPolicy:AuditPolicy example &lt;db_instance_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:mongodb/auditPolicy:AuditPolicy")]
    public partial class AuditPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The status of the audit log. Valid values: `disabled`, `enable`.
        /// </summary>
        [Output("auditStatus")]
        public Output<string> AuditStatus { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
        /// </summary>
        [Output("storagePeriod")]
        public Output<int?> StoragePeriod { get; private set; } = null!;


        /// <summary>
        /// Create a AuditPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuditPolicy(string name, AuditPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:mongodb/auditPolicy:AuditPolicy", name, args ?? new AuditPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuditPolicy(string name, Input<string> id, AuditPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:mongodb/auditPolicy:AuditPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuditPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuditPolicy Get(string name, Input<string> id, AuditPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AuditPolicy(name, id, state, options);
        }
    }

    public sealed class AuditPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The status of the audit log. Valid values: `disabled`, `enable`.
        /// </summary>
        [Input("auditStatus", required: true)]
        public Input<string> AuditStatus { get; set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        /// <summary>
        /// The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
        /// </summary>
        [Input("storagePeriod")]
        public Input<int>? StoragePeriod { get; set; }

        public AuditPolicyArgs()
        {
        }
        public static new AuditPolicyArgs Empty => new AuditPolicyArgs();
    }

    public sealed class AuditPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The status of the audit log. Valid values: `disabled`, `enable`.
        /// </summary>
        [Input("auditStatus")]
        public Input<string>? AuditStatus { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
        /// </summary>
        [Input("storagePeriod")]
        public Input<int>? StoragePeriod { get; set; }

        public AuditPolicyState()
        {
        }
        public static new AuditPolicyState Empty => new AuditPolicyState();
    }
}
