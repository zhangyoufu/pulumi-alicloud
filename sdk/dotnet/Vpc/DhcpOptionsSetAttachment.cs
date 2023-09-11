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
    /// Provides a VPC Dhcp Options Set Attachment resource.
    /// 
    /// For information about VPC Dhcp Options Set and how to use it, see [What is Dhcp Options Set](https://www.alibabacloud.com/help/doc-detail/174112.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.153.0.
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
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var exampleNetwork = new AliCloud.Vpc.Network("exampleNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.4.0.0/16",
    ///     });
    /// 
    ///     var exampleDhcpOptionsSet = new AliCloud.Vpc.DhcpOptionsSet("exampleDhcpOptionsSet", new()
    ///     {
    ///         DhcpOptionsSetName = name,
    ///         DhcpOptionsSetDescription = name,
    ///         DomainName = "example.com",
    ///         DomainNameServers = "100.100.2.136",
    ///     });
    /// 
    ///     var exampleDhcpOptionsSetAttachment = new AliCloud.Vpc.DhcpOptionsSetAttachment("exampleDhcpOptionsSetAttachment", new()
    ///     {
    ///         VpcId = exampleNetwork.Id,
    ///         DhcpOptionsSetId = exampleDhcpOptionsSet.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPC Dhcp Options Set Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment")]
    public partial class DhcpOptionsSetAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the DHCP options set.
        /// </summary>
        [Output("dhcpOptionsSetId")]
        public Output<string> DhcpOptionsSetId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to precheck this request only. Default values: `false`. Valid values:
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The status of the VPC network that is associated with the DHCP options set.  Valid values: `InUse` or `Pending`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC network that is to be associated with the DHCP options set..
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a DhcpOptionsSetAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DhcpOptionsSetAttachment(string name, DhcpOptionsSetAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment", name, args ?? new DhcpOptionsSetAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DhcpOptionsSetAttachment(string name, Input<string> id, DhcpOptionsSetAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DhcpOptionsSetAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DhcpOptionsSetAttachment Get(string name, Input<string> id, DhcpOptionsSetAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new DhcpOptionsSetAttachment(name, id, state, options);
        }
    }

    public sealed class DhcpOptionsSetAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the DHCP options set.
        /// </summary>
        [Input("dhcpOptionsSetId", required: true)]
        public Input<string> DhcpOptionsSetId { get; set; } = null!;

        /// <summary>
        /// Specifies whether to precheck this request only. Default values: `false`. Valid values:
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the VPC network that is to be associated with the DHCP options set..
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public DhcpOptionsSetAttachmentArgs()
        {
        }
        public static new DhcpOptionsSetAttachmentArgs Empty => new DhcpOptionsSetAttachmentArgs();
    }

    public sealed class DhcpOptionsSetAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the DHCP options set.
        /// </summary>
        [Input("dhcpOptionsSetId")]
        public Input<string>? DhcpOptionsSetId { get; set; }

        /// <summary>
        /// Specifies whether to precheck this request only. Default values: `false`. Valid values:
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The status of the VPC network that is associated with the DHCP options set.  Valid values: `InUse` or `Pending`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the VPC network that is to be associated with the DHCP options set..
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public DhcpOptionsSetAttachmentState()
        {
        }
        public static new DhcpOptionsSetAttachmentState Empty => new DhcpOptionsSetAttachmentState();
    }
}
