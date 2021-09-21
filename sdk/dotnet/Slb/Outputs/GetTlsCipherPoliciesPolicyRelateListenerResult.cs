// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb.Outputs
{

    [OutputType]
    public sealed class GetTlsCipherPoliciesPolicyRelateListenerResult
    {
        /// <summary>
        /// The ID of SLB instance.
        /// </summary>
        public readonly string LoadBalancerId;
        /// <summary>
        /// Listening port. Valid value: 1 to 65535.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Snooping protocols. Valid values: `TCP`, `UDP`, `HTTP`, or `HTTPS`.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private GetTlsCipherPoliciesPolicyRelateListenerResult(
            string loadBalancerId,

            int port,

            string protocol)
        {
            LoadBalancerId = loadBalancerId;
            Port = port;
            Protocol = protocol;
        }
    }
}
