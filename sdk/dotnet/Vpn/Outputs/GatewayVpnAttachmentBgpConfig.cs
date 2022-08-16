// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpn.Outputs
{

    [OutputType]
    public sealed class GatewayVpnAttachmentBgpConfig
    {
        /// <summary>
        /// Whether to enable BGP.
        /// </summary>
        public readonly bool? Enable;
        /// <summary>
        /// The ASN on the Alibaba Cloud side.
        /// </summary>
        public readonly int? LocalAsn;
        /// <summary>
        /// The BGP IP address on the Alibaba Cloud side.
        /// </summary>
        public readonly string? LocalBgpIp;
        /// <summary>
        /// The CIDR block of the IPsec tunnel. The CIDR block belongs to 169.254.0.0/16. The mask of the CIDR block is 30 bits in length.
        /// </summary>
        public readonly string? TunnelCidr;

        [OutputConstructor]
        private GatewayVpnAttachmentBgpConfig(
            bool? enable,

            int? localAsn,

            string? localBgpIp,

            string? tunnelCidr)
        {
            Enable = enable;
            LocalAsn = localAsn;
            LocalBgpIp = localBgpIp;
            TunnelCidr = tunnelCidr;
        }
    }
}
