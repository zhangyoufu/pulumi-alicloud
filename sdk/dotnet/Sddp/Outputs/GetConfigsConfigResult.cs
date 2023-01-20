// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sddp.Outputs
{

    [OutputType]
    public sealed class GetConfigsConfigResult
    {
        /// <summary>
        /// Abnormal Alarm General Configuration Module by Using the Encoding.Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// Configure the Number.
        /// </summary>
        public readonly string ConfigId;
        /// <summary>
        /// Default Value.
        /// </summary>
        public readonly string DefaultValue;
        /// <summary>
        /// Abnormal Alarm General Description of the Configuration Item.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of the Config.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetConfigsConfigResult(
            string code,

            string configId,

            string defaultValue,

            string description,

            string id,

            string value)
        {
            Code = code;
            ConfigId = configId;
            DefaultValue = defaultValue;
            Description = description;
            Id = id;
            Value = value;
        }
    }
}
