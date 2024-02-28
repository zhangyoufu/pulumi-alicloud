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
    public sealed class GetConnectionsConnectionTunnelOptionsSpecificationResult
    {
        /// <summary>
        /// Use the VPN customer gateway ID as the search key.
        /// </summary>
        public readonly string CustomerGatewayId;
        /// <summary>
        /// Wether enable Dpd detection.
        /// </summary>
        public readonly bool EnableDpd;
        /// <summary>
        /// enable nat traversal.
        /// </summary>
        public readonly bool EnableNatTraversal;
        public readonly string InternetIp;
        /// <summary>
        /// The role of Tunnel.
        /// </summary>
        public readonly string Role;
        public readonly string State;
        /// <summary>
        /// The negotiation status of the BGP routing protocol. Valid values: `success`, `false`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The bgp config of Tunnel.
        /// </summary>
        public readonly Outputs.GetConnectionsConnectionTunnelOptionsSpecificationTunnelBgpConfigResult TunnelBgpConfig;
        public readonly string TunnelId;
        /// <summary>
        /// The configuration of Phase 1 negotiations in Tunnel.
        /// </summary>
        public readonly Outputs.GetConnectionsConnectionTunnelOptionsSpecificationTunnelIkeConfigResult TunnelIkeConfig;
        /// <summary>
        /// IPsec configuration in Tunnel.
        /// </summary>
        public readonly Outputs.GetConnectionsConnectionTunnelOptionsSpecificationTunnelIpsecConfigResult TunnelIpsecConfig;
        public readonly string ZoneNo;

        [OutputConstructor]
        private GetConnectionsConnectionTunnelOptionsSpecificationResult(
            string customerGatewayId,

            bool enableDpd,

            bool enableNatTraversal,

            string internetIp,

            string role,

            string state,

            string status,

            Outputs.GetConnectionsConnectionTunnelOptionsSpecificationTunnelBgpConfigResult tunnelBgpConfig,

            string tunnelId,

            Outputs.GetConnectionsConnectionTunnelOptionsSpecificationTunnelIkeConfigResult tunnelIkeConfig,

            Outputs.GetConnectionsConnectionTunnelOptionsSpecificationTunnelIpsecConfigResult tunnelIpsecConfig,

            string zoneNo)
        {
            CustomerGatewayId = customerGatewayId;
            EnableDpd = enableDpd;
            EnableNatTraversal = enableNatTraversal;
            InternetIp = internetIp;
            Role = role;
            State = state;
            Status = status;
            TunnelBgpConfig = tunnelBgpConfig;
            TunnelId = tunnelId;
            TunnelIkeConfig = tunnelIkeConfig;
            TunnelIpsecConfig = tunnelIpsecConfig;
            ZoneNo = zoneNo;
        }
    }
}
