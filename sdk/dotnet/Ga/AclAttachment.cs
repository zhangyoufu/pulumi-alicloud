// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    /// <summary>
    /// Provides a Global Accelerator (GA) Acl Attachment resource.
    /// 
    /// For information about Global Accelerator (GA) Acl Attachment and how to use it, see [What is Acl Attachment](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-associateaclswithlistener).
    /// 
    /// &gt; **NOTE:** Available since v1.150.0.
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
    ///     var defaultAccelerator = new AliCloud.Ga.Accelerator("defaultAccelerator", new()
    ///     {
    ///         Duration = 1,
    ///         AutoUseCoupon = true,
    ///         Spec = "1",
    ///     });
    /// 
    ///     var defaultBandwidthPackage = new AliCloud.Ga.BandwidthPackage("defaultBandwidthPackage", new()
    ///     {
    ///         Bandwidth = 100,
    ///         Type = "Basic",
    ///         BandwidthType = "Basic",
    ///         PaymentType = "PayAsYouGo",
    ///         BillingType = "PayBy95",
    ///         Ratio = 30,
    ///     });
    /// 
    ///     var defaultBandwidthPackageAttachment = new AliCloud.Ga.BandwidthPackageAttachment("defaultBandwidthPackageAttachment", new()
    ///     {
    ///         AcceleratorId = defaultAccelerator.Id,
    ///         BandwidthPackageId = defaultBandwidthPackage.Id,
    ///     });
    /// 
    ///     var defaultListener = new AliCloud.Ga.Listener("defaultListener", new()
    ///     {
    ///         AcceleratorId = defaultBandwidthPackageAttachment.AcceleratorId,
    ///         PortRanges = new[]
    ///         {
    ///             new AliCloud.Ga.Inputs.ListenerPortRangeArgs
    ///             {
    ///                 FromPort = 80,
    ///                 ToPort = 80,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultAcl = new AliCloud.Ga.Acl("defaultAcl", new()
    ///     {
    ///         AclName = "terraform-example",
    ///         AddressIpVersion = "IPv4",
    ///     });
    /// 
    ///     var defaultAclEntryAttachment = new AliCloud.Ga.AclEntryAttachment("defaultAclEntryAttachment", new()
    ///     {
    ///         AclId = defaultAcl.Id,
    ///         Entry = "192.168.1.1/32",
    ///         EntryDescription = "terraform-example",
    ///     });
    /// 
    ///     var defaultAclAttachment = new AliCloud.Ga.AclAttachment("defaultAclAttachment", new()
    ///     {
    ///         ListenerId = defaultListener.Id,
    ///         AclId = defaultAcl.Id,
    ///         AclType = "white",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Global Accelerator (GA) Acl Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ga/aclAttachment:AclAttachment example &lt;listener_id&gt;:&lt;acl_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/aclAttachment:AclAttachment")]
    public partial class AclAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of an ACL.
        /// </summary>
        [Output("aclId")]
        public Output<string> AclId { get; private set; } = null!;

        /// <summary>
        /// The type of the ACL. Valid values:
        /// </summary>
        [Output("aclType")]
        public Output<string> AclType { get; private set; } = null!;

        /// <summary>
        /// The dry run.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The ID of the listener.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// The status of the Acl Attachment.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a AclAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AclAttachment(string name, AclAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/aclAttachment:AclAttachment", name, args ?? new AclAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AclAttachment(string name, Input<string> id, AclAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/aclAttachment:AclAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AclAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AclAttachment Get(string name, Input<string> id, AclAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AclAttachment(name, id, state, options);
        }
    }

    public sealed class AclAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of an ACL.
        /// </summary>
        [Input("aclId", required: true)]
        public Input<string> AclId { get; set; } = null!;

        /// <summary>
        /// The type of the ACL. Valid values:
        /// </summary>
        [Input("aclType", required: true)]
        public Input<string> AclType { get; set; } = null!;

        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        public AclAttachmentArgs()
        {
        }
        public static new AclAttachmentArgs Empty => new AclAttachmentArgs();
    }

    public sealed class AclAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of an ACL.
        /// </summary>
        [Input("aclId")]
        public Input<string>? AclId { get; set; }

        /// <summary>
        /// The type of the ACL. Valid values:
        /// </summary>
        [Input("aclType")]
        public Input<string>? AclType { get; set; }

        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the listener.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// The status of the Acl Attachment.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AclAttachmentState()
        {
        }
        public static new AclAttachmentState Empty => new AclAttachmentState();
    }
}
