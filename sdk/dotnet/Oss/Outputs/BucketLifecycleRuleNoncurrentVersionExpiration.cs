// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss.Outputs
{

    [OutputType]
    public sealed class BucketLifecycleRuleNoncurrentVersionExpiration
    {
        /// <summary>
        /// Specifies the number of days after object creation when the specific rule action takes effect.
        /// 
        /// `NOTE`: One and only one of "created_before_date" and "days" can be specified in one abort_multipart_upload configuration.
        /// </summary>
        public readonly int Days;

        [OutputConstructor]
        private BucketLifecycleRuleNoncurrentVersionExpiration(int days)
        {
            Days = days;
        }
    }
}
