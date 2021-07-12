// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Kms.Outputs
{

    [OutputType]
    public sealed class GetKeyVersionsVersionResult
    {
        /// <summary>
        /// Date and time when the key version was created (UTC time).
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// (Removed from v1.124.4) It has been removed and using `create_time` instead.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// ID of the KMS KeyVersion resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The id of kms key.
        /// </summary>
        public readonly string KeyId;
        /// <summary>
        /// ID of the key version.
        /// </summary>
        public readonly string KeyVersionId;

        [OutputConstructor]
        private GetKeyVersionsVersionResult(
            string createTime,

            string creationDate,

            string id,

            string keyId,

            string keyVersionId)
        {
            CreateTime = createTime;
            CreationDate = creationDate;
            Id = id;
            KeyId = keyId;
            KeyVersionId = keyVersionId;
        }
    }
}
