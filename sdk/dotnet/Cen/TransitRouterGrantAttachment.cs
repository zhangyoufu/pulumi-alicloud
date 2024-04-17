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
    /// Provides a Cloud Enterprise Network (CEN) Transit Router Grant Attachment resource.
    /// 
    /// For information about Cloud Enterprise Network (CEN) Transit Router Grant Attachment and how to use it, see [What is Transit Router Grant Attachment](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/grantinstancetotransitrouter).
    /// 
    /// &gt; **NOTE:** Available since v1.187.0.
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
    ///     var example = new AliCloud.Vpc.Network("example", new()
    ///     {
    ///         VpcName = "tf_example",
    ///         CidrBlock = "172.17.3.0/24",
    ///     });
    /// 
    ///     var exampleInstance = new AliCloud.Cen.Instance("example", new()
    ///     {
    ///         CenInstanceName = "tf_example",
    ///         Description = "an example for cen",
    ///     });
    /// 
    ///     var exampleTransitRouterGrantAttachment = new AliCloud.Cen.TransitRouterGrantAttachment("example", new()
    ///     {
    ///         CenId = exampleInstance.Id,
    ///         CenOwnerId = @default.Apply(@default =&gt; @default.Apply(getAccountResult =&gt; getAccountResult.Id)),
    ///         InstanceId = example.Id,
    ///         InstanceType = "VPC",
    ///         OrderType = "PayByCenOwner",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Cloud Enterprise Network (CEN) Transit Router Grant Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment example &lt;instance_type&gt;:&lt;instance_id&gt;:&lt;cen_owner_id&gt;:&lt;cen_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment")]
    public partial class TransitRouterGrantAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Alibaba Cloud account to which the CEN instance belongs.
        /// </summary>
        [Output("cenOwnerId")]
        public Output<string> CenOwnerId { get; private set; } = null!;

        /// <summary>
        /// The ID of the network instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
        /// </summary>
        [Output("orderType")]
        public Output<string> OrderType { get; private set; } = null!;


        /// <summary>
        /// Create a TransitRouterGrantAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransitRouterGrantAttachment(string name, TransitRouterGrantAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment", name, args ?? new TransitRouterGrantAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransitRouterGrantAttachment(string name, Input<string> id, TransitRouterGrantAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TransitRouterGrantAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransitRouterGrantAttachment Get(string name, Input<string> id, TransitRouterGrantAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new TransitRouterGrantAttachment(name, id, state, options);
        }
    }

    public sealed class TransitRouterGrantAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        /// <summary>
        /// The ID of the Alibaba Cloud account to which the CEN instance belongs.
        /// </summary>
        [Input("cenOwnerId", required: true)]
        public Input<string> CenOwnerId { get; set; } = null!;

        /// <summary>
        /// The ID of the network instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
        /// </summary>
        [Input("orderType")]
        public Input<string>? OrderType { get; set; }

        public TransitRouterGrantAttachmentArgs()
        {
        }
        public static new TransitRouterGrantAttachmentArgs Empty => new TransitRouterGrantAttachmentArgs();
    }

    public sealed class TransitRouterGrantAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The ID of the Alibaba Cloud account to which the CEN instance belongs.
        /// </summary>
        [Input("cenOwnerId")]
        public Input<string>? CenOwnerId { get; set; }

        /// <summary>
        /// The ID of the network instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
        /// </summary>
        [Input("orderType")]
        public Input<string>? OrderType { get; set; }

        public TransitRouterGrantAttachmentState()
        {
        }
        public static new TransitRouterGrantAttachmentState Empty => new TransitRouterGrantAttachmentState();
    }
}
