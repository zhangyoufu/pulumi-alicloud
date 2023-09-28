// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class ApplicationCustomHostAliasV2GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name or hostname.
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// The IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public ApplicationCustomHostAliasV2GetArgs()
        {
        }
        public static new ApplicationCustomHostAliasV2GetArgs Empty => new ApplicationCustomHostAliasV2GetArgs();
    }
}
