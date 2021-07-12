// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg.Inputs
{

    public sealed class CompliancePackConfigRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("configRuleParameters", required: true)]
        private InputList<Inputs.CompliancePackConfigRuleConfigRuleParameterGetArgs>? _configRuleParameters;

        /// <summary>
        /// A list of Config Rule Parameters.
        /// </summary>
        public InputList<Inputs.CompliancePackConfigRuleConfigRuleParameterGetArgs> ConfigRuleParameters
        {
            get => _configRuleParameters ?? (_configRuleParameters = new InputList<Inputs.CompliancePackConfigRuleConfigRuleParameterGetArgs>());
            set => _configRuleParameters = value;
        }

        /// <summary>
        /// The Managed Rule Identifier.
        /// </summary>
        [Input("managedRuleIdentifier", required: true)]
        public Input<string> ManagedRuleIdentifier { get; set; } = null!;

        public CompliancePackConfigRuleGetArgs()
        {
        }
    }
}
