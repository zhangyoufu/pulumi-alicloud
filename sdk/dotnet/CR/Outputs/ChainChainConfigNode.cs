// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CR.Outputs
{

    [OutputType]
    public sealed class ChainChainConfigNode
    {
        /// <summary>
        /// Whether to enable the delivery chain node. Valid values: `true`, `false`.
        /// </summary>
        public readonly bool? Enable;
        /// <summary>
        /// The configuration of delivery chain node. See `node_config` below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ChainChainConfigNodeNodeConfig> NodeConfigs;
        /// <summary>
        /// The name of delivery chain node.
        /// </summary>
        public readonly string? NodeName;

        [OutputConstructor]
        private ChainChainConfigNode(
            bool? enable,

            ImmutableArray<Outputs.ChainChainConfigNodeNodeConfig> nodeConfigs,

            string? nodeName)
        {
            Enable = enable;
            NodeConfigs = nodeConfigs;
            NodeName = nodeName;
        }
    }
}
