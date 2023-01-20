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
    public sealed class BucketLifecycleRule
    {
        /// <summary>
        /// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed (documented below).
        /// </summary>
        public readonly ImmutableArray<Outputs.BucketLifecycleRuleAbortMultipartUpload> AbortMultipartUploads;
        /// <summary>
        /// Specifies lifecycle rule status.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Specifies a period in the object's expire (documented below).
        /// </summary>
        public readonly ImmutableArray<Outputs.BucketLifecycleRuleExpiration> Expirations;
        /// <summary>
        /// Unique identifier for the rule. If omitted, OSS bucket will assign a unique name.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Specifies when noncurrent object versions expire (documented below).
        /// </summary>
        public readonly ImmutableArray<Outputs.BucketLifecycleRuleNoncurrentVersionExpiration> NoncurrentVersionExpirations;
        /// <summary>
        /// Specifies when noncurrent object versions transitions (documented below).
        /// </summary>
        public readonly ImmutableArray<Outputs.BucketLifecycleRuleNoncurrentVersionTransition> NoncurrentVersionTransitions;
        /// <summary>
        /// Object key prefix identifying one or more objects to which the rule applies. Default value is null, the rule applies to all objects in a bucket.
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// Specifies the time when an object is converted to the IA or archive storage class during a valid life cycle. (documented below).
        /// </summary>
        public readonly ImmutableArray<Outputs.BucketLifecycleRuleTransition> Transitions;

        [OutputConstructor]
        private BucketLifecycleRule(
            ImmutableArray<Outputs.BucketLifecycleRuleAbortMultipartUpload> abortMultipartUploads,

            bool enabled,

            ImmutableArray<Outputs.BucketLifecycleRuleExpiration> expirations,

            string? id,

            ImmutableArray<Outputs.BucketLifecycleRuleNoncurrentVersionExpiration> noncurrentVersionExpirations,

            ImmutableArray<Outputs.BucketLifecycleRuleNoncurrentVersionTransition> noncurrentVersionTransitions,

            string? prefix,

            ImmutableArray<Outputs.BucketLifecycleRuleTransition> transitions)
        {
            AbortMultipartUploads = abortMultipartUploads;
            Enabled = enabled;
            Expirations = expirations;
            Id = id;
            NoncurrentVersionExpirations = noncurrentVersionExpirations;
            NoncurrentVersionTransitions = noncurrentVersionTransitions;
            Prefix = prefix;
            Transitions = transitions;
        }
    }
}
