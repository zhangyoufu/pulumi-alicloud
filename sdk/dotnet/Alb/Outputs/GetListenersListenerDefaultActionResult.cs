// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Outputs
{

    [OutputType]
    public sealed class GetListenersListenerDefaultActionResult
    {
        /// <summary>
        /// The configuration of the forwarding rule action. This parameter is required if the Type parameter is set to FowardGroup.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListenersListenerDefaultActionForwardGroupConfigResult> ForwardGroupConfigs;
        /// <summary>
        /// Action Type. The value is set to ForwardGroup. It indicates that requests are forwarded to multiple vServer groups.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetListenersListenerDefaultActionResult(
            ImmutableArray<Outputs.GetListenersListenerDefaultActionForwardGroupConfigResult> forwardGroupConfigs,

            string type)
        {
            ForwardGroupConfigs = forwardGroupConfigs;
            Type = type;
        }
    }
}
