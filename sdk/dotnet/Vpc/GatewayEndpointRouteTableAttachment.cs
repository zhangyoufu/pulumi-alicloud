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
    /// Provides a VPC Gateway Endpoint Route Table Attachment resource. VPC gateway node association route.
    /// 
    /// For information about VPC Gateway Endpoint Route Table Attachment and how to use it, see [What is Gateway Endpoint Route Table Attachment](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/311148).
    /// 
    /// &gt; **NOTE:** Available since v1.208.0.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaulteVpc = new AliCloud.Vpc.Network("defaulteVpc", new()
    ///     {
    ///         Description = "test",
    ///     });
    /// 
    ///     var defaultGE = new AliCloud.Vpc.GatewayEndpoint("defaultGE", new()
    ///     {
    ///         ServiceName = "com.aliyun.cn-hangzhou.oss",
    ///         PolicyDocument = @"        {
    ///           ""Version"": ""1"",
    ///           ""Statement"": [{
    ///             ""Effect"": ""Allow"",
    ///             ""Resource"": [""*""],
    ///             ""Action"": [""*""],
    ///             ""Principal"": [""*""]
    ///           }]
    ///         }
    /// ",
    ///         VpcId = defaulteVpc.Id,
    ///         GatewayEndpointDescrption = "test-gateway-endpoint",
    ///         GatewayEndpointName = $"{name}1",
    ///     });
    /// 
    ///     var defaultRT = new AliCloud.Vpc.RouteTable("defaultRT", new()
    ///     {
    ///         VpcId = defaulteVpc.Id,
    ///         RouteTableName = $"{name}2",
    ///     });
    /// 
    ///     var @default = new AliCloud.Vpc.GatewayEndpointRouteTableAttachment("default", new()
    ///     {
    ///         GatewayEndpointId = defaultGE.Id,
    ///         RouteTableId = defaultRT.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// VPC Gateway Endpoint Route Table Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/gatewayEndpointRouteTableAttachment:GatewayEndpointRouteTableAttachment example &lt;gateway_endpoint_id&gt;:&lt;route_table_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/gatewayEndpointRouteTableAttachment:GatewayEndpointRouteTableAttachment")]
    public partial class GatewayEndpointRouteTableAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the gateway endpoint instance to which you want to associate the route table.
        /// </summary>
        [Output("gatewayEndpointId")]
        public Output<string> GatewayEndpointId { get; private set; } = null!;

        /// <summary>
        /// Routing table ID.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;

        /// <summary>
        /// Status of the gateway endpoint.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayEndpointRouteTableAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayEndpointRouteTableAttachment(string name, GatewayEndpointRouteTableAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/gatewayEndpointRouteTableAttachment:GatewayEndpointRouteTableAttachment", name, args ?? new GatewayEndpointRouteTableAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayEndpointRouteTableAttachment(string name, Input<string> id, GatewayEndpointRouteTableAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/gatewayEndpointRouteTableAttachment:GatewayEndpointRouteTableAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayEndpointRouteTableAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayEndpointRouteTableAttachment Get(string name, Input<string> id, GatewayEndpointRouteTableAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayEndpointRouteTableAttachment(name, id, state, options);
        }
    }

    public sealed class GatewayEndpointRouteTableAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the gateway endpoint instance to which you want to associate the route table.
        /// </summary>
        [Input("gatewayEndpointId", required: true)]
        public Input<string> GatewayEndpointId { get; set; } = null!;

        /// <summary>
        /// Routing table ID.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        public GatewayEndpointRouteTableAttachmentArgs()
        {
        }
        public static new GatewayEndpointRouteTableAttachmentArgs Empty => new GatewayEndpointRouteTableAttachmentArgs();
    }

    public sealed class GatewayEndpointRouteTableAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the gateway endpoint instance to which you want to associate the route table.
        /// </summary>
        [Input("gatewayEndpointId")]
        public Input<string>? GatewayEndpointId { get; set; }

        /// <summary>
        /// Routing table ID.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        /// <summary>
        /// Status of the gateway endpoint.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GatewayEndpointRouteTableAttachmentState()
        {
        }
        public static new GatewayEndpointRouteTableAttachmentState Empty => new GatewayEndpointRouteTableAttachmentState();
    }
}
