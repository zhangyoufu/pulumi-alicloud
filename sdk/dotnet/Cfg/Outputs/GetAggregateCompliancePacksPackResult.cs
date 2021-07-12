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
    public sealed class GetAggregateCompliancePacksPackResult
    {
        /// <summary>
        /// The Aliyun User Id.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// -The Aggregate Compliance Package Name.
        /// </summary>
        public readonly string AggregateCompliancePackName;
        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        public readonly string AggregatorCompliancePackId;
        /// <summary>
        /// The template ID of the Compliance Package.
        /// </summary>
        public readonly string CompliancePackTemplateId;
        /// <summary>
        /// A list of The Aggregate Compliance Package Rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAggregateCompliancePacksPackConfigRuleResult> ConfigRules;
        /// <summary>
        /// The description of aggregate compliance pack.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of the Aggregate Compliance Pack.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Risk Level.
        /// </summary>
        public readonly int RiskLevel;
        /// <summary>
        /// The status of the resource.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetAggregateCompliancePacksPackResult(
            string accountId,

            string aggregateCompliancePackName,

            string aggregatorCompliancePackId,

            string compliancePackTemplateId,

            ImmutableArray<Outputs.GetAggregateCompliancePacksPackConfigRuleResult> configRules,

            string description,

            string id,

            int riskLevel,

            string status)
        {
            AccountId = accountId;
            AggregateCompliancePackName = aggregateCompliancePackName;
            AggregatorCompliancePackId = aggregatorCompliancePackId;
            CompliancePackTemplateId = compliancePackTemplateId;
            ConfigRules = configRules;
            Description = description;
            Id = id;
            RiskLevel = riskLevel;
            Status = status;
        }
    }
}
