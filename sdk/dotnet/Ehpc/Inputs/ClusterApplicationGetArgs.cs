// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ehpc.Inputs
{

    public sealed class ClusterApplicationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The tag of the software.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public ClusterApplicationGetArgs()
        {
        }
        public static new ClusterApplicationGetArgs Empty => new ClusterApplicationGetArgs();
    }
}
