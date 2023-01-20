// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Outputs
{

    [OutputType]
    public sealed class HybridMonitorSlsTaskSlsProcessConfigFilterFilter
    {
        /// <summary>
        /// The method that is used to filter logs imported from Log Service. Valid values: `&gt;`, `&gt;=`, `=`, `&lt;=`, `&lt;`, `!=`, `contain`, `notContain`.
        /// </summary>
        public readonly string? Operator;
        /// <summary>
        /// The name of the key that is used to aggregate logs imported from Log Service.
        /// </summary>
        public readonly string? SlsKeyName;
        /// <summary>
        /// The value of the key that is used to filter logs imported from Log Service.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private HybridMonitorSlsTaskSlsProcessConfigFilterFilter(
            string? @operator,

            string? slsKeyName,

            string? value)
        {
            Operator = @operator;
            SlsKeyName = slsKeyName;
            Value = value;
        }
    }
}
