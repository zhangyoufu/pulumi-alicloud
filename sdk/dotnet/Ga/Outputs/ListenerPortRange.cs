// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga.Outputs
{

    [OutputType]
    public sealed class ListenerPortRange
    {
        /// <summary>
        /// The initial listening port used to receive requests and forward them to terminal nodes.
        /// </summary>
        public readonly int FromPort;
        /// <summary>
        /// The end listening port used to receive requests and forward them to terminal nodes.
        /// </summary>
        public readonly int ToPort;

        [OutputConstructor]
        private ListenerPortRange(
            int fromPort,

            int toPort)
        {
            FromPort = fromPort;
            ToPort = toPort;
        }
    }
}
