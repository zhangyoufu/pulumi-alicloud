// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class ApplicationScalingRuleScalingRuleTimerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Start date. When the `begin_date` and `end_date` values are empty. it indicates long-term execution and is the default value.
        /// </summary>
        [Input("beginDate")]
        public Input<string>? BeginDate { get; set; }

        /// <summary>
        /// The End Date. When the `begin_date` and `end_date` values are empty. it indicates long-term execution and is the default value.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// The period in which a timed elastic scaling strategy is executed.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        [Input("schedules")]
        private InputList<Inputs.ApplicationScalingRuleScalingRuleTimerScheduleGetArgs>? _schedules;

        /// <summary>
        /// Resilient Scaling Strategy Trigger Timing. See `schedules` below.
        /// </summary>
        public InputList<Inputs.ApplicationScalingRuleScalingRuleTimerScheduleGetArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.ApplicationScalingRuleScalingRuleTimerScheduleGetArgs>());
            set => _schedules = value;
        }

        public ApplicationScalingRuleScalingRuleTimerGetArgs()
        {
        }
        public static new ApplicationScalingRuleScalingRuleTimerGetArgs Empty => new ApplicationScalingRuleScalingRuleTimerGetArgs();
    }
}
