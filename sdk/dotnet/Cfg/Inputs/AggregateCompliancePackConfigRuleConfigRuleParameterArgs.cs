// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg.Inputs
{

    public sealed class AggregateCompliancePackConfigRuleConfigRuleParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Parameter Name.
        /// </summary>
        [Input("parameterName")]
        public Input<string>? ParameterName { get; set; }

        /// <summary>
        /// The Parameter Value.
        /// </summary>
        [Input("parameterValue")]
        public Input<string>? ParameterValue { get; set; }

        public AggregateCompliancePackConfigRuleConfigRuleParameterArgs()
        {
        }
        public static new AggregateCompliancePackConfigRuleConfigRuleParameterArgs Empty => new AggregateCompliancePackConfigRuleConfigRuleParameterArgs();
    }
}
