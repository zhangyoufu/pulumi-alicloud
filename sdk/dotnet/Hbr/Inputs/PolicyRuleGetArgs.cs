// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr.Inputs
{

    public sealed class PolicyRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This parameter is required only when the value of **RuleType** is **TRANSITION. The minimum value is 30, and the Retention-ArchiveDays needs to be greater than or equal to 60.
        /// </summary>
        [Input("archiveDays")]
        public Input<int>? ArchiveDays { get; set; }

        /// <summary>
        /// This parameter is required only when the **RuleType** value is **BACKUP. Backup Type.
        /// </summary>
        [Input("backupType")]
        public Input<string>? BackupType { get; set; }

        /// <summary>
        /// This parameter is required only when **RuleType** is set to **BACKUP**.
        /// </summary>
        [Input("keepLatestSnapshots")]
        public Input<int>? KeepLatestSnapshots { get; set; }

        /// <summary>
        /// Only when the **RuleType** value is.
        /// </summary>
        [Input("replicationRegionId")]
        public Input<string>? ReplicationRegionId { get; set; }

        /// <summary>
        /// Retention time, in days.
        /// </summary>
        [Input("retention")]
        public Input<int>? Retention { get; set; }

        [Input("retentionRules")]
        private InputList<Inputs.PolicyRuleRetentionRuleGetArgs>? _retentionRules;

        /// <summary>
        /// This parameter is required only when the value of **RuleType** is **TRANSITION**. See `retention_rules` below.
        /// </summary>
        public InputList<Inputs.PolicyRuleRetentionRuleGetArgs> RetentionRules
        {
            get => _retentionRules ?? (_retentionRules = new InputList<Inputs.PolicyRuleRetentionRuleGetArgs>());
            set => _retentionRules = value;
        }

        /// <summary>
        /// Rule ID.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// Rule Type.
        /// </summary>
        [Input("ruleType", required: true)]
        public Input<string> RuleType { get; set; } = null!;

        /// <summary>
        /// This parameter is required only if you set the **RuleType** parameter to **BACKUP**. This parameter specifies the backup schedule settings. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is complete. For example, `I|1631685600|P1D` specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.  *   startTime: the time at which the system starts to run a backup job. The time must follow the UNIX time format. Unit: seconds. *   interval: the interval at which the system runs a backup job. The interval must follow the ISO 8601 standard. For example, PT1H specifies an interval of one hour. P1D specifies an interval of one day.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Vault ID.
        /// </summary>
        [Input("vaultId")]
        public Input<string>? VaultId { get; set; }

        public PolicyRuleGetArgs()
        {
        }
        public static new PolicyRuleGetArgs Empty => new PolicyRuleGetArgs();
    }
}
