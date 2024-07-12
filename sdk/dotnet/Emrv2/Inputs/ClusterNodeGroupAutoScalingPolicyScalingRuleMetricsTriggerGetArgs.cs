// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Emrv2.Inputs
{

    public sealed class ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditionLogicOperator")]
        public Input<string>? ConditionLogicOperator { get; set; }

        [Input("conditions")]
        private InputList<Inputs.ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerConditionGetArgs>? _conditions;
        public InputList<Inputs.ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerConditionGetArgs>());
            set => _conditions = value;
        }

        [Input("coolDownInterval")]
        public Input<int>? CoolDownInterval { get; set; }

        [Input("evaluationCount", required: true)]
        public Input<int> EvaluationCount { get; set; } = null!;

        [Input("timeConstraints")]
        private InputList<Inputs.ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerTimeConstraintGetArgs>? _timeConstraints;
        public InputList<Inputs.ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerTimeConstraintGetArgs> TimeConstraints
        {
            get => _timeConstraints ?? (_timeConstraints = new InputList<Inputs.ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerTimeConstraintGetArgs>());
            set => _timeConstraints = value;
        }

        [Input("timeWindow", required: true)]
        public Input<int> TimeWindow { get; set; } = null!;

        public ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerGetArgs()
        {
        }
        public static new ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerGetArgs Empty => new ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerGetArgs();
    }
}
