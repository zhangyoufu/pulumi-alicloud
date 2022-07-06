// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eds.Outputs
{

    [OutputType]
    public sealed class GetAdConnectorDirectoriesDirectoryAdConnectorResult
    {
        /// <summary>
        /// The address of AD connector.
        /// </summary>
        public readonly string AdConnectorAddress;
        /// <summary>
        /// The status of connector.
        /// </summary>
        public readonly string ConnectorStatus;
        /// <summary>
        /// The ID of the network interface.
        /// </summary>
        public readonly string NetworkInterfaceId;
        /// <summary>
        /// The AD Connector specifications.
        /// </summary>
        public readonly string Specification;
        /// <summary>
        /// The AD Connector control trust password.
        /// </summary>
        public readonly string TrustKey;
        /// <summary>
        /// The ID of VSwitch.
        /// </summary>
        public readonly string VswitchId;

        [OutputConstructor]
        private GetAdConnectorDirectoriesDirectoryAdConnectorResult(
            string adConnectorAddress,

            string connectorStatus,

            string networkInterfaceId,

            string specification,

            string trustKey,

            string vswitchId)
        {
            AdConnectorAddress = adConnectorAddress;
            ConnectorStatus = connectorStatus;
            NetworkInterfaceId = networkInterfaceId;
            Specification = specification;
            TrustKey = trustKey;
            VswitchId = vswitchId;
        }
    }
}
