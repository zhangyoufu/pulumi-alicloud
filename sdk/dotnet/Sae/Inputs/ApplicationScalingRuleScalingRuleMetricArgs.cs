// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class ApplicationScalingRuleScalingRuleMetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of instances applied.
        /// </summary>
        [Input("maxReplicas")]
        public Input<int>? MaxReplicas { get; set; }

        [Input("metrics")]
        private InputList<Inputs.ApplicationScalingRuleScalingRuleMetricMetricArgs>? _metrics;

        /// <summary>
        /// Indicator rule configuration. See `metrics` below.
        /// </summary>
        public InputList<Inputs.ApplicationScalingRuleScalingRuleMetricMetricArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.ApplicationScalingRuleScalingRuleMetricMetricArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// Minimum number of instances applied.
        /// </summary>
        [Input("minReplicas")]
        public Input<int>? MinReplicas { get; set; }

        /// <summary>
        /// Apply shrink rules. See `scale_down_rules` below.
        /// </summary>
        [Input("scaleDownRules")]
        public Input<Inputs.ApplicationScalingRuleScalingRuleMetricScaleDownRulesArgs>? ScaleDownRules { get; set; }

        /// <summary>
        /// Apply expansion rules. See `scale_up_rules` below.
        /// </summary>
        [Input("scaleUpRules")]
        public Input<Inputs.ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs>? ScaleUpRules { get; set; }

        public ApplicationScalingRuleScalingRuleMetricArgs()
        {
        }
        public static new ApplicationScalingRuleScalingRuleMetricArgs Empty => new ApplicationScalingRuleScalingRuleMetricArgs();
    }
}
