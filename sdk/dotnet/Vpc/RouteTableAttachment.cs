// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
    /// Provides a VPC Route Table Attachment resource. Routing table associated resource type.
    /// 
    /// For information about VPC Route Table Attachment and how to use it, see [What is Route Table Attachment](https://www.alibabacloud.com/help/doc-detail/174112.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.194.0.
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
    ///     var foo = new AliCloud.Vpc.Network("foo", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/12",
    ///         VpcName = name,
    ///     });
    /// 
    ///     var @default = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var fooSwitch = new AliCloud.Vpc.Switch("foo", new()
    ///     {
    ///         VpcId = foo.Id,
    ///         CidrBlock = "172.16.0.0/21",
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var fooRouteTable = new AliCloud.Vpc.RouteTable("foo", new()
    ///     {
    ///         VpcId = foo.Id,
    ///         RouteTableName = name,
    ///         Description = "route_table_attachment",
    ///     });
    /// 
    ///     var fooRouteTableAttachment = new AliCloud.Vpc.RouteTableAttachment("foo", new()
    ///     {
    ///         VswitchId = fooSwitch.Id,
    ///         RouteTableId = fooRouteTable.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPC Route Table Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/routeTableAttachment:RouteTableAttachment example &lt;route_table_id&gt;:&lt;vswitch_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/routeTableAttachment:RouteTableAttachment")]
    public partial class RouteTableAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the route table to be bound to the switch.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the switch to bind the route table.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a RouteTableAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteTableAttachment(string name, RouteTableAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/routeTableAttachment:RouteTableAttachment", name, args ?? new RouteTableAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteTableAttachment(string name, Input<string> id, RouteTableAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/routeTableAttachment:RouteTableAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouteTableAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteTableAttachment Get(string name, Input<string> id, RouteTableAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteTableAttachment(name, id, state, options);
        }
    }

    public sealed class RouteTableAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the route table to be bound to the switch.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        /// <summary>
        /// The ID of the switch to bind the route table.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public RouteTableAttachmentArgs()
        {
        }
        public static new RouteTableAttachmentArgs Empty => new RouteTableAttachmentArgs();
    }

    public sealed class RouteTableAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the route table to be bound to the switch.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the switch to bind the route table.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public RouteTableAttachmentState()
        {
        }
        public static new RouteTableAttachmentState Empty => new RouteTableAttachmentState();
    }
}
