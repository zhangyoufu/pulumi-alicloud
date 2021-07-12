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
    public sealed class AggregateCompliancePackConfigRule
    {
        /// <summary>
        /// A list of parameter rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.AggregateCompliancePackConfigRuleConfigRuleParameter> ConfigRuleParameters;
        /// <summary>
        /// The Managed Rule Identifier.
        /// </summary>
        public readonly string ManagedRuleIdentifier;

        [OutputConstructor]
        private AggregateCompliancePackConfigRule(
            ImmutableArray<Outputs.AggregateCompliancePackConfigRuleConfigRuleParameter> configRuleParameters,

            string managedRuleIdentifier)
        {
            ConfigRuleParameters = configRuleParameters;
            ManagedRuleIdentifier = managedRuleIdentifier;
        }
    }
}
