// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg.Inputs
{

    public sealed class CompliancePackConfigRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("configRuleParameters")]
        private InputList<Inputs.CompliancePackConfigRuleConfigRuleParameterArgs>? _configRuleParameters;

        /// <summary>
        /// A list of Config Rule Parameters.
        /// </summary>
        public InputList<Inputs.CompliancePackConfigRuleConfigRuleParameterArgs> ConfigRuleParameters
        {
            get => _configRuleParameters ?? (_configRuleParameters = new InputList<Inputs.CompliancePackConfigRuleConfigRuleParameterArgs>());
            set => _configRuleParameters = value;
        }

        /// <summary>
        /// The Managed Rule Identifier.
        /// </summary>
        [Input("managedRuleIdentifier", required: true)]
        public Input<string> ManagedRuleIdentifier { get; set; } = null!;

        public CompliancePackConfigRuleArgs()
        {
        }
        public static new CompliancePackConfigRuleArgs Empty => new CompliancePackConfigRuleArgs();
    }
}
