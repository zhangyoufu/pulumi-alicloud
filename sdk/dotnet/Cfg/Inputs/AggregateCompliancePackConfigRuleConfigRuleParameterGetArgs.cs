// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg.Inputs
{

    public sealed class AggregateCompliancePackConfigRuleConfigRuleParameterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Parameter Name.
        /// </summary>
        [Input("parameterName", required: true)]
        public Input<string> ParameterName { get; set; } = null!;

        /// <summary>
        /// The Parameter Value.
        /// </summary>
        [Input("parameterValue", required: true)]
        public Input<string> ParameterValue { get; set; } = null!;

        public AggregateCompliancePackConfigRuleConfigRuleParameterGetArgs()
        {
        }
    }
}
