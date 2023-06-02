// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc.Inputs
{

    public sealed class NetworkAclEgressAclEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The network of the destination address.
        /// </summary>
        [Input("destinationCidrIp")]
        public Input<string>? DestinationCidrIp { get; set; }

        /// <summary>
        /// Name of the outbound rule entry.The name must be 1 to 128 characters in length and cannot start with http:// or https.
        /// </summary>
        [Input("networkAclEntryName")]
        public Input<string>? NetworkAclEntryName { get; set; }

        /// <summary>
        /// Authorization policy. Value:
        /// - accept: Allow.
        /// - drop: Refused.
        /// - accept: Allow.
        /// - drop: Refused.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The destination port range of the outbound rule.When the Protocol type of the outbound rule is all, icmp, or gre, the port range is - 1/-1, indicating that the port is not restricted.When the Protocol type of the outbound rule is tcp or udp, the port range is 1 to 65535, and the format is 1/200 or 80/80, indicating port 1 to port 200 or port 80.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The protocol type. Value:
        /// - icmp: Network Control Message Protocol.
        /// - gre: Generic Routing Encapsulation Protocol.
        /// - tcp: Transmission Control Protocol.
        /// - udp: User Datagram Protocol.
        /// - all: Supports all protocols.
        /// 
        /// - icmp: Network Control Message Protocol.
        /// - gre: Generic Routing Encapsulation Protocol.
        /// - tcp: Transmission Control Protocol.
        /// - udp: User Datagram Protocol.
        /// - all: Supports all protocols.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public NetworkAclEgressAclEntryArgs()
        {
        }
        public static new NetworkAclEgressAclEntryArgs Empty => new NetworkAclEgressAclEntryArgs();
    }
}
