// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sls.Inputs
{

    public sealed class AlertConfigurationJoinConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public AlertConfigurationJoinConfigurationArgs()
        {
        }
        public static new AlertConfigurationJoinConfigurationArgs Empty => new AlertConfigurationJoinConfigurationArgs();
    }
}
