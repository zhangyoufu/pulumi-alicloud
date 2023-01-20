// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class ListenerQuicConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// There Is a Need to Correlate the QuIC Listener ID. The Https Listener, in Effect at the Time. quicupgradeenabled True When Required.
        /// </summary>
        [Input("quicListenerId")]
        public Input<string>? QuicListenerId { get; set; }

        /// <summary>
        /// Indicates Whether to Enable the QuIC Upgrade.
        /// </summary>
        [Input("quicUpgradeEnabled")]
        public Input<bool>? QuicUpgradeEnabled { get; set; }

        public ListenerQuicConfigArgs()
        {
        }
        public static new ListenerQuicConfigArgs Empty => new ListenerQuicConfigArgs();
    }
}
