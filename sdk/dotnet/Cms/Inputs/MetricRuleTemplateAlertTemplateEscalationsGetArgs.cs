// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Inputs
{

    public sealed class MetricRuleTemplateAlertTemplateEscalationsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The condition for triggering critical-level alerts. See the following `Block critical`.
        /// </summary>
        [Input("critical")]
        public Input<Inputs.MetricRuleTemplateAlertTemplateEscalationsCriticalGetArgs>? Critical { get; set; }

        /// <summary>
        /// The condition for triggering info-level alerts. See the following `Block info`.
        /// </summary>
        [Input("info")]
        public Input<Inputs.MetricRuleTemplateAlertTemplateEscalationsInfoGetArgs>? Info { get; set; }

        /// <summary>
        /// The condition for triggering warn-level alerts. See the following `Block warn`.
        /// </summary>
        [Input("warn")]
        public Input<Inputs.MetricRuleTemplateAlertTemplateEscalationsWarnGetArgs>? Warn { get; set; }

        public MetricRuleTemplateAlertTemplateEscalationsGetArgs()
        {
        }
    }
}
