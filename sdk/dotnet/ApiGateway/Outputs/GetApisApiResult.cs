// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class GetApisApiResult
    {
        /// <summary>
        /// The ID of the API.
        /// </summary>
        public readonly string ApiId;
        /// <summary>
        /// The description of the API.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of the API group.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// The name of the API group.
        /// </summary>
        public readonly string GroupName;
        public readonly string Id;
        /// <summary>
        /// The name of the API.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The region ID of the API.
        /// </summary>
        public readonly string RegionId;

        [OutputConstructor]
        private GetApisApiResult(
            string apiId,

            string description,

            string groupId,

            string groupName,

            string id,

            string name,

            string regionId)
        {
            ApiId = apiId;
            Description = description;
            GroupId = groupId;
            GroupName = groupName;
            Id = id;
            Name = name;
            RegionId = regionId;
        }
    }
}
