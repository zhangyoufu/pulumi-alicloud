// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// Provides a Cen Child Instance Route Entry To Attachment resource.
    /// 
    /// For information about Cen Child Instance Route Entry To Attachment and how to use it, see [What is Child Instance Route Entry To Attachment](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createcenchildinstancerouteentrytoattachment).
    /// 
    /// &gt; **NOTE:** Available since v1.195.0.
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
    ///     var @default = AliCloud.Cen.GetTransitRouterAvailableResources.Invoke();
    /// 
    ///     var masterZone = @default.Apply(@default =&gt; @default.Apply(getTransitRouterAvailableResourcesResult =&gt; getTransitRouterAvailableResourcesResult.Resources[0]?.MasterZones[0]));
    /// 
    ///     var slaveZone = @default.Apply(@default =&gt; @default.Apply(getTransitRouterAvailableResourcesResult =&gt; getTransitRouterAvailableResourcesResult.Resources[0]?.SlaveZones[1]));
    /// 
    ///     var example = new AliCloud.Vpc.Network("example", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "192.168.0.0/16",
    ///     });
    /// 
    ///     var exampleMaster = new AliCloud.Vpc.Switch("example_master", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "192.168.1.0/24",
    ///         VpcId = example.Id,
    ///         ZoneId = masterZone,
    ///     });
    /// 
    ///     var exampleSlave = new AliCloud.Vpc.Switch("example_slave", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "192.168.2.0/24",
    ///         VpcId = example.Id,
    ///         ZoneId = slaveZone,
    ///     });
    /// 
    ///     var exampleInstance = new AliCloud.Cen.Instance("example", new()
    ///     {
    ///         CenInstanceName = name,
    ///         ProtectionLevel = "REDUCED",
    ///     });
    /// 
    ///     var exampleTransitRouter = new AliCloud.Cen.TransitRouter("example", new()
    ///     {
    ///         TransitRouterName = name,
    ///         CenId = exampleInstance.Id,
    ///     });
    /// 
    ///     var exampleTransitRouterVpcAttachment = new AliCloud.Cen.TransitRouterVpcAttachment("example", new()
    ///     {
    ///         CenId = exampleInstance.Id,
    ///         TransitRouterId = exampleTransitRouter.TransitRouterId,
    ///         VpcId = example.Id,
    ///         ZoneMappings = new[]
    ///         {
    ///             new AliCloud.Cen.Inputs.TransitRouterVpcAttachmentZoneMappingArgs
    ///             {
    ///                 ZoneId = masterZone,
    ///                 VswitchId = exampleMaster.Id,
    ///             },
    ///             new AliCloud.Cen.Inputs.TransitRouterVpcAttachmentZoneMappingArgs
    ///             {
    ///                 ZoneId = slaveZone,
    ///                 VswitchId = exampleSlave.Id,
    ///             },
    ///         },
    ///         TransitRouterAttachmentName = name,
    ///         TransitRouterAttachmentDescription = name,
    ///     });
    /// 
    ///     var exampleRouteTable = new AliCloud.Vpc.RouteTable("example", new()
    ///     {
    ///         VpcId = example.Id,
    ///         RouteTableName = name,
    ///         Description = name,
    ///     });
    /// 
    ///     var exampleChildInstanceRouteEntryToAttachment = new AliCloud.Cen.ChildInstanceRouteEntryToAttachment("example", new()
    ///     {
    ///         TransitRouterAttachmentId = exampleTransitRouterVpcAttachment.TransitRouterAttachmentId,
    ///         CenId = exampleInstance.Id,
    ///         DestinationCidrBlock = "10.0.0.0/24",
    ///         ChildInstanceRouteTableId = exampleRouteTable.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Cen Child Instance Route Entry To Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment example &lt;cen_id&gt;:&lt;child_instance_route_table_id&gt;:&lt;transit_router_attachment_id&gt;:&lt;destination_cidr_block&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment")]
    public partial class ChildInstanceRouteEntryToAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;

        /// <summary>
        /// The first ID of the resource
        /// </summary>
        [Output("childInstanceRouteTableId")]
        public Output<string> ChildInstanceRouteTableId { get; private set; } = null!;

        /// <summary>
        /// DestinationCidrBlock
        /// </summary>
        [Output("destinationCidrBlock")]
        public Output<string> DestinationCidrBlock { get; private set; } = null!;

        /// <summary>
        /// Whether to perform pre-check on this request, including permission and instance status verification.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// ServiceType
        /// </summary>
        [Output("serviceType")]
        public Output<string> ServiceType { get; private set; } = null!;

        /// <summary>
        /// The status of the resource
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// TransitRouterAttachmentId
        /// </summary>
        [Output("transitRouterAttachmentId")]
        public Output<string> TransitRouterAttachmentId { get; private set; } = null!;


        /// <summary>
        /// Create a ChildInstanceRouteEntryToAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ChildInstanceRouteEntryToAttachment(string name, ChildInstanceRouteEntryToAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment", name, args ?? new ChildInstanceRouteEntryToAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ChildInstanceRouteEntryToAttachment(string name, Input<string> id, ChildInstanceRouteEntryToAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ChildInstanceRouteEntryToAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ChildInstanceRouteEntryToAttachment Get(string name, Input<string> id, ChildInstanceRouteEntryToAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ChildInstanceRouteEntryToAttachment(name, id, state, options);
        }
    }

    public sealed class ChildInstanceRouteEntryToAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        /// <summary>
        /// The first ID of the resource
        /// </summary>
        [Input("childInstanceRouteTableId", required: true)]
        public Input<string> ChildInstanceRouteTableId { get; set; } = null!;

        /// <summary>
        /// DestinationCidrBlock
        /// </summary>
        [Input("destinationCidrBlock", required: true)]
        public Input<string> DestinationCidrBlock { get; set; } = null!;

        /// <summary>
        /// Whether to perform pre-check on this request, including permission and instance status verification.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// TransitRouterAttachmentId
        /// </summary>
        [Input("transitRouterAttachmentId", required: true)]
        public Input<string> TransitRouterAttachmentId { get; set; } = null!;

        public ChildInstanceRouteEntryToAttachmentArgs()
        {
        }
        public static new ChildInstanceRouteEntryToAttachmentArgs Empty => new ChildInstanceRouteEntryToAttachmentArgs();
    }

    public sealed class ChildInstanceRouteEntryToAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The first ID of the resource
        /// </summary>
        [Input("childInstanceRouteTableId")]
        public Input<string>? ChildInstanceRouteTableId { get; set; }

        /// <summary>
        /// DestinationCidrBlock
        /// </summary>
        [Input("destinationCidrBlock")]
        public Input<string>? DestinationCidrBlock { get; set; }

        /// <summary>
        /// Whether to perform pre-check on this request, including permission and instance status verification.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// ServiceType
        /// </summary>
        [Input("serviceType")]
        public Input<string>? ServiceType { get; set; }

        /// <summary>
        /// The status of the resource
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// TransitRouterAttachmentId
        /// </summary>
        [Input("transitRouterAttachmentId")]
        public Input<string>? TransitRouterAttachmentId { get; set; }

        public ChildInstanceRouteEntryToAttachmentState()
        {
        }
        public static new ChildInstanceRouteEntryToAttachmentState Empty => new ChildInstanceRouteEntryToAttachmentState();
    }
}
