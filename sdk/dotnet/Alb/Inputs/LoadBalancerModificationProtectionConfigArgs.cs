// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class LoadBalancerModificationProtectionConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reason for modification protection. This parameter must be 2 to 128 characters in length, and can contain letters, digits, periods, underscores, and hyphens. The reason must start with a letter. **Note:** This parameter takes effect only when `status` is set to `ConsoleProtection`.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// Specifies whether to enable the configuration read-only mode for the ALB instance. Valid values: `NonProtection` and `ConsoleProtection`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public LoadBalancerModificationProtectionConfigArgs()
        {
        }
        public static new LoadBalancerModificationProtectionConfigArgs Empty => new LoadBalancerModificationProtectionConfigArgs();
    }
}
