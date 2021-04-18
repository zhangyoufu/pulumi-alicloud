// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss.Inputs
{

    public sealed class BucketLifecycleRuleExpirationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the time before which the rules take effect. The date must conform to the ISO8601 format and always be UTC 00:00. For example: 2002-10-11T00:00:00.000Z indicates that parts created before 2002-10-11T00:00:00.000Z are deleted, and parts created after this time (including this time) are not deleted.
        /// </summary>
        [Input("createdBeforeDate")]
        public Input<string>? CreatedBeforeDate { get; set; }

        /// <summary>
        /// Specifies the date after which you want the corresponding action to take effect. The value obeys ISO8601 format like `2017-03-09`.
        /// </summary>
        [Input("date")]
        public Input<string>? Date { get; set; }

        /// <summary>
        /// Specifies the number of days noncurrent object versions transition.
        /// </summary>
        [Input("days")]
        public Input<int>? Days { get; set; }

        [Input("expiredObjectDeleteMarker")]
        public Input<bool>? ExpiredObjectDeleteMarker { get; set; }

        public BucketLifecycleRuleExpirationArgs()
        {
        }
    }
}
