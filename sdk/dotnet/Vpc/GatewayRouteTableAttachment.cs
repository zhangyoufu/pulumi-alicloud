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
    /// Provides a VPC Gateway Route Table Attachment resource.
    /// 
    /// For information about VPC Gateway Route Table Attachment and how to use it, see [What is Gateway Route Table Attachment](https://www.alibabacloud.com/help/doc-detail/174112.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.194.0+.
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
    ///     var exampleNetwork = new AliCloud.Vpc.Network("exampleNetwork", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/12",
    ///         VpcName = "terraform-example",
    ///     });
    /// 
    ///     var exampleRouteTable = new AliCloud.Vpc.RouteTable("exampleRouteTable", new()
    ///     {
    ///         VpcId = exampleNetwork.Id,
    ///         RouteTableName = "terraform-example",
    ///         Description = "terraform-example",
    ///         AssociateType = "Gateway",
    ///     });
    /// 
    ///     var exampleIpv4Gateway = new AliCloud.Vpc.Ipv4Gateway("exampleIpv4Gateway", new()
    ///     {
    ///         Ipv4GatewayName = "terraform-example",
    ///         VpcId = exampleNetwork.Id,
    ///         Enabled = true,
    ///     });
    /// 
    ///     var exampleGatewayRouteTableAttachment = new AliCloud.Vpc.GatewayRouteTableAttachment("exampleGatewayRouteTableAttachment", new()
    ///     {
    ///         Ipv4GatewayId = exampleIpv4Gateway.Id,
    ///         RouteTableId = exampleRouteTable.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// VPC Gateway Route Table Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment example &lt;route_table_id&gt;:&lt;ipv4_gateway_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment")]
    public partial class GatewayRouteTableAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to only precheck this request. Default value: `false`.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The ID of the IPv4 Gateway instance.
        /// </summary>
        [Output("ipv4GatewayId")]
        public Output<string> Ipv4GatewayId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Gateway route table to be bound.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;

        /// <summary>
        /// The status of the IPv4 Gateway instance. Value:
        /// - **Creating**: The function is being created.
        /// - **Created**: Created and available.
        /// - **Modifying**: is being modified.
        /// - **Deleting**: Deleting.
        /// - **Deleted**: Deleted.
        /// - **Activating**: enabled.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayRouteTableAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayRouteTableAttachment(string name, GatewayRouteTableAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment", name, args ?? new GatewayRouteTableAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayRouteTableAttachment(string name, Input<string> id, GatewayRouteTableAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayRouteTableAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayRouteTableAttachment Get(string name, Input<string> id, GatewayRouteTableAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayRouteTableAttachment(name, id, state, options);
        }
    }

    public sealed class GatewayRouteTableAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to only precheck this request. Default value: `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the IPv4 Gateway instance.
        /// </summary>
        [Input("ipv4GatewayId", required: true)]
        public Input<string> Ipv4GatewayId { get; set; } = null!;

        /// <summary>
        /// The ID of the Gateway route table to be bound.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        public GatewayRouteTableAttachmentArgs()
        {
        }
        public static new GatewayRouteTableAttachmentArgs Empty => new GatewayRouteTableAttachmentArgs();
    }

    public sealed class GatewayRouteTableAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Specifies whether to only precheck this request. Default value: `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the IPv4 Gateway instance.
        /// </summary>
        [Input("ipv4GatewayId")]
        public Input<string>? Ipv4GatewayId { get; set; }

        /// <summary>
        /// The ID of the Gateway route table to be bound.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        /// <summary>
        /// The status of the IPv4 Gateway instance. Value:
        /// - **Creating**: The function is being created.
        /// - **Created**: Created and available.
        /// - **Modifying**: is being modified.
        /// - **Deleting**: Deleting.
        /// - **Deleted**: Deleted.
        /// - **Activating**: enabled.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GatewayRouteTableAttachmentState()
        {
        }
        public static new GatewayRouteTableAttachmentState Empty => new GatewayRouteTableAttachmentState();
    }
}
