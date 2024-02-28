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
    public sealed class GetConnectionsConnectionIpsecConfigResult
    {
        /// <summary>
        /// IPsec Auth algorithm.
        /// </summary>
        public readonly string? IpsecAuthAlg;
        /// <summary>
        /// IPsec Encript algorithm.
        /// </summary>
        public readonly string? IpsecEncAlg;
        /// <summary>
        /// IPsec lifetime.
        /// </summary>
        public readonly int? IpsecLifetime;
        /// <summary>
        /// DH Group.
        /// </summary>
        public readonly string? IpsecPfs;

        [OutputConstructor]
        private GetConnectionsConnectionIpsecConfigResult(
            string? ipsecAuthAlg,

            string? ipsecEncAlg,

            int? ipsecLifetime,

            string? ipsecPfs)
        {
            IpsecAuthAlg = ipsecAuthAlg;
            IpsecEncAlg = ipsecEncAlg;
            IpsecLifetime = ipsecLifetime;
            IpsecPfs = ipsecPfs;
        }
    }
}
