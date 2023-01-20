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
    public sealed class GetConfigurationRecordersRecorderResult
    {
        /// <summary>
        /// The ID of the Alicloud account.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The ID of the Config Configuration Recorder. Value as the `account_id`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enterprise version configuration audit enabled status.
        /// </summary>
        public readonly string OrganizationEnableStatus;
        /// <summary>
        /// The ID of the Enterprise management account.
        /// </summary>
        public readonly int OrganizationMasterId;
        /// <summary>
        /// A list of resource types to be monitored.
        /// </summary>
        public readonly ImmutableArray<string> ResourceTypes;
        /// <summary>
        /// Status of resource monitoring.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetConfigurationRecordersRecorderResult(
            string accountId,

            string id,

            string organizationEnableStatus,

            int organizationMasterId,

            ImmutableArray<string> resourceTypes,

            string status)
        {
            AccountId = accountId;
            Id = id;
            OrganizationEnableStatus = organizationEnableStatus;
            OrganizationMasterId = organizationMasterId;
            ResourceTypes = resourceTypes;
            Status = status;
        }
    }
}
