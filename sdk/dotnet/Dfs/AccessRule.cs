// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dfs
{
    /// <summary>
    /// Provides a DFS Access Rule resource.
    /// 
    /// For information about DFS Access Rule and how to use it, see [What is Access Rule](https://www.alibabacloud.com/help/doc-detail/207144.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.140.0.
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
    ///     var name = config.Get("name") ?? "example_name";
    ///     var defaultAccessGroup = new AliCloud.Dfs.AccessGroup("defaultAccessGroup", new()
    ///     {
    ///         NetworkType = "VPC",
    ///         AccessGroupName = name,
    ///         Description = name,
    ///     });
    /// 
    ///     var defaultAccessRule = new AliCloud.Dfs.AccessRule("defaultAccessRule", new()
    ///     {
    ///         NetworkSegment = "192.0.2.0/24",
    ///         AccessGroupId = defaultAccessGroup.Id,
    ///         Description = name,
    ///         RwAccessType = "RDWR",
    ///         Priority = 10,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DFS Access Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:dfs/accessRule:AccessRule example &lt;access_group_id&gt;:&lt;access_rule_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dfs/accessRule:AccessRule")]
    public partial class AccessRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID of Access Group.
        /// </summary>
        [Output("accessGroupId")]
        public Output<string> AccessGroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Access Rule.
        /// </summary>
        [Output("accessRuleId")]
        public Output<string> AccessRuleId { get; private set; } = null!;

        /// <summary>
        /// The Description of the Access Rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The NetworkSegment of the Access Rule.
        /// </summary>
        [Output("networkSegment")]
        public Output<string> NetworkSegment { get; private set; } = null!;

        /// <summary>
        /// The Priority of the Access Rule. Valid values: `1` to `100`. **NOTE:** When multiple rules are matched by the same authorized object, the high-priority rule takes effect. `1` is the highest priority.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
        /// </summary>
        [Output("rwAccessType")]
        public Output<string> RwAccessType { get; private set; } = null!;


        /// <summary>
        /// Create a AccessRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessRule(string name, AccessRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dfs/accessRule:AccessRule", name, args ?? new AccessRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessRule(string name, Input<string> id, AccessRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dfs/accessRule:AccessRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccessRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessRule Get(string name, Input<string> id, AccessRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessRule(name, id, state, options);
        }
    }

    public sealed class AccessRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of Access Group.
        /// </summary>
        [Input("accessGroupId", required: true)]
        public Input<string> AccessGroupId { get; set; } = null!;

        /// <summary>
        /// The Description of the Access Rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The NetworkSegment of the Access Rule.
        /// </summary>
        [Input("networkSegment", required: true)]
        public Input<string> NetworkSegment { get; set; } = null!;

        /// <summary>
        /// The Priority of the Access Rule. Valid values: `1` to `100`. **NOTE:** When multiple rules are matched by the same authorized object, the high-priority rule takes effect. `1` is the highest priority.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// The RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
        /// </summary>
        [Input("rwAccessType", required: true)]
        public Input<string> RwAccessType { get; set; } = null!;

        public AccessRuleArgs()
        {
        }
        public static new AccessRuleArgs Empty => new AccessRuleArgs();
    }

    public sealed class AccessRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of Access Group.
        /// </summary>
        [Input("accessGroupId")]
        public Input<string>? AccessGroupId { get; set; }

        /// <summary>
        /// The ID of the Access Rule.
        /// </summary>
        [Input("accessRuleId")]
        public Input<string>? AccessRuleId { get; set; }

        /// <summary>
        /// The Description of the Access Rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The NetworkSegment of the Access Rule.
        /// </summary>
        [Input("networkSegment")]
        public Input<string>? NetworkSegment { get; set; }

        /// <summary>
        /// The Priority of the Access Rule. Valid values: `1` to `100`. **NOTE:** When multiple rules are matched by the same authorized object, the high-priority rule takes effect. `1` is the highest priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
        /// </summary>
        [Input("rwAccessType")]
        public Input<string>? RwAccessType { get; set; }

        public AccessRuleState()
        {
        }
        public static new AccessRuleState Empty => new AccessRuleState();
    }
}
