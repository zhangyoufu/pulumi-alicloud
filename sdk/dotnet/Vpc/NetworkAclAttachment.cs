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
    /// Provides a network acl attachment resource to associate network acls to vswitches.
    /// 
    /// &gt; **NOTE:** Available in 1.44.0+. Currently, the resource are only available in Hongkong(cn-hongkong), India(ap-south-1), and Indonesia(ap-southeast-1) regions.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/network_acl_attachment.html.markdown.
    /// </summary>
    public partial class NetworkAclAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the network acl, the field can't be changed.
        /// </summary>
        [Output("networkAclId")]
        public Output<string> NetworkAclId { get; private set; } = null!;

        /// <summary>
        /// List of the resources associated with the network acl. The details see Block Resources.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.NetworkAclAttachmentResources>> Resources { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkAclAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkAclAttachment(string name, NetworkAclAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/networkAclAttachment:NetworkAclAttachment", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private NetworkAclAttachment(string name, Input<string> id, NetworkAclAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/networkAclAttachment:NetworkAclAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkAclAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkAclAttachment Get(string name, Input<string> id, NetworkAclAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkAclAttachment(name, id, state, options);
        }
    }

    public sealed class NetworkAclAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the network acl, the field can't be changed.
        /// </summary>
        [Input("networkAclId", required: true)]
        public Input<string> NetworkAclId { get; set; } = null!;

        [Input("resources", required: true)]
        private InputList<Inputs.NetworkAclAttachmentResourcesArgs>? _resources;

        /// <summary>
        /// List of the resources associated with the network acl. The details see Block Resources.
        /// </summary>
        public InputList<Inputs.NetworkAclAttachmentResourcesArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.NetworkAclAttachmentResourcesArgs>());
            set => _resources = value;
        }

        public NetworkAclAttachmentArgs()
        {
        }
    }

    public sealed class NetworkAclAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the network acl, the field can't be changed.
        /// </summary>
        [Input("networkAclId")]
        public Input<string>? NetworkAclId { get; set; }

        [Input("resources")]
        private InputList<Inputs.NetworkAclAttachmentResourcesGetArgs>? _resources;

        /// <summary>
        /// List of the resources associated with the network acl. The details see Block Resources.
        /// </summary>
        public InputList<Inputs.NetworkAclAttachmentResourcesGetArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.NetworkAclAttachmentResourcesGetArgs>());
            set => _resources = value;
        }

        public NetworkAclAttachmentState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class NetworkAclAttachmentResourcesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource id that the network acl will associate with.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The resource id that the network acl will associate with. Only support `VSwitch` now.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        public NetworkAclAttachmentResourcesArgs()
        {
        }
    }

    public sealed class NetworkAclAttachmentResourcesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource id that the network acl will associate with.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The resource id that the network acl will associate with. Only support `VSwitch` now.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        public NetworkAclAttachmentResourcesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class NetworkAclAttachmentResources
    {
        /// <summary>
        /// The resource id that the network acl will associate with.
        /// </summary>
        public readonly string ResourceId;
        /// <summary>
        /// The resource id that the network acl will associate with. Only support `VSwitch` now.
        /// </summary>
        public readonly string ResourceType;

        [OutputConstructor]
        private NetworkAclAttachmentResources(
            string resourceId,
            string resourceType)
        {
            ResourceId = resourceId;
            ResourceType = resourceType;
        }
    }
    }
}
