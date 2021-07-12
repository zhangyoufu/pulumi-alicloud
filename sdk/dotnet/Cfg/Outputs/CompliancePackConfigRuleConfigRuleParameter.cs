// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg.Outputs
{

    [OutputType]
    public sealed class CompliancePackConfigRuleConfigRuleParameter
    {
        /// <summary>
        /// The parameter name.
        /// </summary>
        public readonly string ParameterName;
        /// <summary>
        /// The parameter value.
        /// </summary>
        public readonly string? ParameterValue;

        [OutputConstructor]
        private CompliancePackConfigRuleConfigRuleParameter(
            string parameterName,

            string? parameterValue)
        {
            ParameterName = parameterName;
            ParameterValue = parameterValue;
        }
    }
}
