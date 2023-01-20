// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ThreatDetection.Outputs
{

    [OutputType]
    public sealed class GetHoneypotNodesNodeResult
    {
        /// <summary>
        /// Whether to allow honeypot access to the external network. Value:-**true**: Allow-**false**: Disabled
        /// </summary>
        public readonly bool AllowHoneypotAccessInternet;
        /// <summary>
        /// Number of probes available.
        /// </summary>
        public readonly int AvailableProbeNum;
        public readonly string CreateTime;
        /// <summary>
        /// The ID of the Honeypot management node.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Honeypot management node id.
        /// </summary>
        public readonly string NodeId;
        /// <summary>
        /// The name of the management node.
        /// </summary>
        public readonly string NodeName;
        /// <summary>
        /// Release the collection of network segments.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupProbeIpLists;
        public readonly int Status;

        [OutputConstructor]
        private GetHoneypotNodesNodeResult(
            bool allowHoneypotAccessInternet,

            int availableProbeNum,

            string createTime,

            string id,

            string nodeId,

            string nodeName,

            ImmutableArray<string> securityGroupProbeIpLists,

            int status)
        {
            AllowHoneypotAccessInternet = allowHoneypotAccessInternet;
            AvailableProbeNum = availableProbeNum;
            CreateTime = createTime;
            Id = id;
            NodeId = nodeId;
            NodeName = nodeName;
            SecurityGroupProbeIpLists = securityGroupProbeIpLists;
            Status = status;
        }
    }
}
