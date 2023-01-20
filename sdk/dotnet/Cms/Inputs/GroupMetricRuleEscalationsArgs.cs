// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Inputs
{

    public sealed class GroupMetricRuleEscalationsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The critical level.
        /// </summary>
        [Input("critical")]
        public Input<Inputs.GroupMetricRuleEscalationsCriticalArgs>? Critical { get; set; }

        /// <summary>
        /// The info level.
        /// </summary>
        [Input("info")]
        public Input<Inputs.GroupMetricRuleEscalationsInfoArgs>? Info { get; set; }

        /// <summary>
        /// The warn level.
        /// </summary>
        [Input("warn")]
        public Input<Inputs.GroupMetricRuleEscalationsWarnArgs>? Warn { get; set; }

        public GroupMetricRuleEscalationsArgs()
        {
        }
        public static new GroupMetricRuleEscalationsArgs Empty => new GroupMetricRuleEscalationsArgs();
    }
}
