// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sls.Inputs
{

    public sealed class AlertConfigurationPolicyConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionPolicyId")]
        public Input<string>? ActionPolicyId { get; set; }

        [Input("alertPolicyId")]
        public Input<string>? AlertPolicyId { get; set; }

        [Input("repeatInterval")]
        public Input<string>? RepeatInterval { get; set; }

        public AlertConfigurationPolicyConfigurationGetArgs()
        {
        }
        public static new AlertConfigurationPolicyConfigurationGetArgs Empty => new AlertConfigurationPolicyConfigurationGetArgs();
    }
}
