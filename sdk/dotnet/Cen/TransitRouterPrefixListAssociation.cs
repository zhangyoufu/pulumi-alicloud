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
    /// Provides a Cloud Enterprise Network (CEN) Transit Router Prefix List Association resource.
    /// 
    /// For information about Cloud Enterprise Network (CEN) Transit Router Prefix List Association and how to use it, see [What is Transit Router Prefix List Association](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/createtransitrouterprefixlistassociation).
    /// 
    /// &gt; **NOTE:** Available since v1.188.0.
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
    ///     var @default = AliCloud.GetAccount.Invoke();
    /// 
    ///     var examplePrefixList = new AliCloud.Vpc.PrefixList("examplePrefixList", new()
    ///     {
    ///         Entrys = new[]
    ///         {
    ///             new AliCloud.Vpc.Inputs.PrefixListEntryArgs
    ///             {
    ///                 Cidr = "192.168.0.0/16",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleInstance = new AliCloud.Cen.Instance("exampleInstance", new()
    ///     {
    ///         CenInstanceName = "tf_example",
    ///         Description = "an example for cen",
    ///     });
    /// 
    ///     var exampleTransitRouter = new AliCloud.Cen.TransitRouter("exampleTransitRouter", new()
    ///     {
    ///         TransitRouterName = "tf_example",
    ///         CenId = exampleInstance.Id,
    ///     });
    /// 
    ///     var exampleTransitRouterRouteTable = new AliCloud.Cen.TransitRouterRouteTable("exampleTransitRouterRouteTable", new()
    ///     {
    ///         TransitRouterId = exampleTransitRouter.TransitRouterId,
    ///     });
    /// 
    ///     var exampleTransitRouterPrefixListAssociation = new AliCloud.Cen.TransitRouterPrefixListAssociation("exampleTransitRouterPrefixListAssociation", new()
    ///     {
    ///         PrefixListId = examplePrefixList.Id,
    ///         TransitRouterId = exampleTransitRouter.TransitRouterId,
    ///         TransitRouterTableId = exampleTransitRouterRouteTable.TransitRouterRouteTableId,
    ///         NextHop = "BlackHole",
    ///         NextHopType = "BlackHole",
    ///         OwnerUid = @default.Apply(@default =&gt; @default.Apply(getAccountResult =&gt; getAccountResult.Id)),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Cloud Enterprise Network (CEN) Transit Router Prefix List Association can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cen/transitRouterPrefixListAssociation:TransitRouterPrefixListAssociation default &lt;prefix_list_id&gt;:&lt;transit_router_id&gt;:&lt;transit_router_table_id&gt;:&lt;next_hop&gt;.
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/transitRouterPrefixListAssociation:TransitRouterPrefixListAssociation")]
    public partial class TransitRouterPrefixListAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the next hop. **NOTE:** If `next_hop` is set to `BlackHole`, you must set this parameter to `BlackHole`.
        /// </summary>
        [Output("nextHop")]
        public Output<string> NextHop { get; private set; } = null!;

        /// <summary>
        /// The type of the next hop. Valid values:
        /// </summary>
        [Output("nextHopType")]
        public Output<string> NextHopType { get; private set; } = null!;

        /// <summary>
        /// The ID of the Alibaba Cloud account to which the prefix list belongs.
        /// </summary>
        [Output("ownerUid")]
        public Output<int> OwnerUid { get; private set; } = null!;

        /// <summary>
        /// The ID of the prefix list.
        /// </summary>
        [Output("prefixListId")]
        public Output<string> PrefixListId { get; private set; } = null!;

        /// <summary>
        /// The status of the prefix list.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Output("transitRouterId")]
        public Output<string> TransitRouterId { get; private set; } = null!;

        /// <summary>
        /// The ID of the route table of the transit router.
        /// </summary>
        [Output("transitRouterTableId")]
        public Output<string> TransitRouterTableId { get; private set; } = null!;


        /// <summary>
        /// Create a TransitRouterPrefixListAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransitRouterPrefixListAssociation(string name, TransitRouterPrefixListAssociationArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/transitRouterPrefixListAssociation:TransitRouterPrefixListAssociation", name, args ?? new TransitRouterPrefixListAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransitRouterPrefixListAssociation(string name, Input<string> id, TransitRouterPrefixListAssociationState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/transitRouterPrefixListAssociation:TransitRouterPrefixListAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TransitRouterPrefixListAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransitRouterPrefixListAssociation Get(string name, Input<string> id, TransitRouterPrefixListAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new TransitRouterPrefixListAssociation(name, id, state, options);
        }
    }

    public sealed class TransitRouterPrefixListAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the next hop. **NOTE:** If `next_hop` is set to `BlackHole`, you must set this parameter to `BlackHole`.
        /// </summary>
        [Input("nextHop", required: true)]
        public Input<string> NextHop { get; set; } = null!;

        /// <summary>
        /// The type of the next hop. Valid values:
        /// </summary>
        [Input("nextHopType")]
        public Input<string>? NextHopType { get; set; }

        /// <summary>
        /// The ID of the Alibaba Cloud account to which the prefix list belongs.
        /// </summary>
        [Input("ownerUid")]
        public Input<int>? OwnerUid { get; set; }

        /// <summary>
        /// The ID of the prefix list.
        /// </summary>
        [Input("prefixListId", required: true)]
        public Input<string> PrefixListId { get; set; } = null!;

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Input("transitRouterId", required: true)]
        public Input<string> TransitRouterId { get; set; } = null!;

        /// <summary>
        /// The ID of the route table of the transit router.
        /// </summary>
        [Input("transitRouterTableId", required: true)]
        public Input<string> TransitRouterTableId { get; set; } = null!;

        public TransitRouterPrefixListAssociationArgs()
        {
        }
        public static new TransitRouterPrefixListAssociationArgs Empty => new TransitRouterPrefixListAssociationArgs();
    }

    public sealed class TransitRouterPrefixListAssociationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the next hop. **NOTE:** If `next_hop` is set to `BlackHole`, you must set this parameter to `BlackHole`.
        /// </summary>
        [Input("nextHop")]
        public Input<string>? NextHop { get; set; }

        /// <summary>
        /// The type of the next hop. Valid values:
        /// </summary>
        [Input("nextHopType")]
        public Input<string>? NextHopType { get; set; }

        /// <summary>
        /// The ID of the Alibaba Cloud account to which the prefix list belongs.
        /// </summary>
        [Input("ownerUid")]
        public Input<int>? OwnerUid { get; set; }

        /// <summary>
        /// The ID of the prefix list.
        /// </summary>
        [Input("prefixListId")]
        public Input<string>? PrefixListId { get; set; }

        /// <summary>
        /// The status of the prefix list.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Input("transitRouterId")]
        public Input<string>? TransitRouterId { get; set; }

        /// <summary>
        /// The ID of the route table of the transit router.
        /// </summary>
        [Input("transitRouterTableId")]
        public Input<string>? TransitRouterTableId { get; set; }

        public TransitRouterPrefixListAssociationState()
        {
        }
        public static new TransitRouterPrefixListAssociationState Empty => new TransitRouterPrefixListAssociationState();
    }
}
