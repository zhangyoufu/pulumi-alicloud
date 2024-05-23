// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class ManagedKubernetesDeleteOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The deletion mode of the cluster. Different resources may have different default behavior, see `resource_type` for details. Valid values:
        /// </summary>
        [Input("deleteMode")]
        public Input<string>? DeleteMode { get; set; }

        /// <summary>
        /// The type of resources that are created by cluster. Valid values:
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public ManagedKubernetesDeleteOptionArgs()
        {
        }
        public static new ManagedKubernetesDeleteOptionArgs Empty => new ManagedKubernetesDeleteOptionArgs();
    }
}
