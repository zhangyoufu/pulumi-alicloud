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
    /// Provides a VPC Route Table resource. Currently, customized route tables are available in most regions apart from China (Beijing), China (Hangzhou), and China (Shenzhen) regions.
    /// 
    /// For information about VPC Route Table and how to use it, see [What is Route Table](https://www.alibabacloud.com/help/doc-detail/87057.htm).
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
    ///     var defaultVpc = new AliCloud.Vpc.Network("defaultVpc", new()
    ///     {
    ///         VpcName = name,
    ///     });
    /// 
    ///     var @default = new AliCloud.Vpc.RouteTable("default", new()
    ///     {
    ///         Description = "test-description",
    ///         VpcId = defaultVpc.Id,
    ///         RouteTableName = name,
    ///         AssociateType = "VSwitch",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPC Route Table can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/routeTable:RouteTable example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/routeTable:RouteTable")]
    public partial class RouteTable : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of cloud resource that is bound to the routing table. Value:
        /// - **VSwitch**: switch.
        /// - **Gateway**:IPv4 Gateway.
        /// </summary>
        [Output("associateType")]
        public Output<string> AssociateType { get; private set; } = null!;

        /// <summary>
        /// The creation time of the routing table.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the routing table.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the routing table.
        /// </summary>
        [Output("routeTableName")]
        public Output<string> RouteTableName { get; private set; } = null!;

        /// <summary>
        /// Routing table state.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tag.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of VPC.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a RouteTable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteTable(string name, RouteTableArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/routeTable:RouteTable", name, args ?? new RouteTableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteTable(string name, Input<string> id, RouteTableState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/routeTable:RouteTable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouteTable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteTable Get(string name, Input<string> id, RouteTableState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteTable(name, id, state, options);
        }
    }

    public sealed class RouteTableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of cloud resource that is bound to the routing table. Value:
        /// - **VSwitch**: switch.
        /// - **Gateway**:IPv4 Gateway.
        /// </summary>
        [Input("associateType")]
        public Input<string>? AssociateType { get; set; }

        /// <summary>
        /// Description of the routing table.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the routing table.
        /// </summary>
        [Input("routeTableName")]
        public Input<string>? RouteTableName { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of VPC.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public RouteTableArgs()
        {
        }
        public static new RouteTableArgs Empty => new RouteTableArgs();
    }

    public sealed class RouteTableState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of cloud resource that is bound to the routing table. Value:
        /// - **VSwitch**: switch.
        /// - **Gateway**:IPv4 Gateway.
        /// </summary>
        [Input("associateType")]
        public Input<string>? AssociateType { get; set; }

        /// <summary>
        /// The creation time of the routing table.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of the routing table.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resource group ID.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The name of the routing table.
        /// </summary>
        [Input("routeTableName")]
        public Input<string>? RouteTableName { get; set; }

        /// <summary>
        /// Routing table state.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of VPC.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public RouteTableState()
        {
        }
        public static new RouteTableState Empty => new RouteTableState();
    }
}
