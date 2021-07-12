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
    /// ## Import
    /// 
    /// Network Interfaces Attachment resource can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/networkInterfaceAttachment:NetworkInterfaceAttachment eni eni-abc123456789000:i-abc123456789000
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/networkInterfaceAttachment:NetworkInterfaceAttachment")]
    public partial class NetworkInterfaceAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The instance ID to attach.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The ENI ID to attach.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        [Output("trunkNetworkInstanceId")]
        public Output<string?> TrunkNetworkInstanceId { get; private set; } = null!;

        [Output("waitForNetworkConfigurationReady")]
        public Output<bool?> WaitForNetworkConfigurationReady { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterfaceAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterfaceAttachment(string name, NetworkInterfaceAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/networkInterfaceAttachment:NetworkInterfaceAttachment", name, args ?? new NetworkInterfaceAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterfaceAttachment(string name, Input<string> id, NetworkInterfaceAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/networkInterfaceAttachment:NetworkInterfaceAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterfaceAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterfaceAttachment Get(string name, Input<string> id, NetworkInterfaceAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterfaceAttachment(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance ID to attach.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The ENI ID to attach.
        /// </summary>
        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        [Input("trunkNetworkInstanceId")]
        public Input<string>? TrunkNetworkInstanceId { get; set; }

        [Input("waitForNetworkConfigurationReady")]
        public Input<bool>? WaitForNetworkConfigurationReady { get; set; }

        public NetworkInterfaceAttachmentArgs()
        {
        }
    }

    public sealed class NetworkInterfaceAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance ID to attach.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The ENI ID to attach.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        [Input("trunkNetworkInstanceId")]
        public Input<string>? TrunkNetworkInstanceId { get; set; }

        [Input("waitForNetworkConfigurationReady")]
        public Input<bool>? WaitForNetworkConfigurationReady { get; set; }

        public NetworkInterfaceAttachmentState()
        {
        }
    }
}
