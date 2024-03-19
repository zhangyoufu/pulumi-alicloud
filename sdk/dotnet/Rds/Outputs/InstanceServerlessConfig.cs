// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds.Outputs
{

    [OutputType]
    public sealed class InstanceServerlessConfig
    {
        public readonly bool? AutoPause;
        /// <summary>
        /// The maximum number of RDS Capacity Units (RCUs). The value of this parameter must be greater than or equal to `min_capacity` and only supports passing integers. Valid values:
        /// - MySQL: 1~8
        /// - SQLServer: 2~8
        /// - PostgreSQL: 1~12
        /// </summary>
        public readonly double MaxCapacity;
        /// <summary>
        /// The minimum number of RCUs. The value of this parameter must be less than or equal to `max_capacity`. Valid values:
        /// - MySQL: 0.5~8
        /// - SQLServer: 2~8 \(Supports integers only\).
        /// - PostgreSQL: 0.5~12
        /// </summary>
        public readonly double MinCapacity;
        /// <summary>
        /// Specifies whether to enable the forced scaling feature for the serverless instance. Valid values:
        /// - true: enables the feature.
        /// - false: disables the feature. This is the default value.
        /// &gt; - Only MySQL Serverless instances need to set this parameter. After enabling this parameter, there will be a flash break within 1 minute when the instance is forced to expand or shrink. Please use it with caution according to the actual situation.
        /// &gt; - The elastic scaling of an instance RCU usually takes effect immediately, but in some special circumstances (such as during large transaction execution), it is not possible to complete scaling immediately. In this case, this parameter can be enabled to force scaling.
        /// </summary>
        public readonly bool? SwitchForce;

        [OutputConstructor]
        private InstanceServerlessConfig(
            bool? autoPause,

            double maxCapacity,

            double minCapacity,

            bool? switchForce)
        {
            AutoPause = autoPause;
            MaxCapacity = maxCapacity;
            MinCapacity = minCapacity;
            SwitchForce = switchForce;
        }
    }
}
