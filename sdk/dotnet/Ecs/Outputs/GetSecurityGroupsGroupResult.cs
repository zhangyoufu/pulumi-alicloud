// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs.Outputs
{

    [OutputType]
    public sealed class GetSecurityGroupsGroupResult
    {
        /// <summary>
        /// Creation time of the security group.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The description of the security group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of the security group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether to allow inner network access.
        /// </summary>
        public readonly bool InnerAccess;
        /// <summary>
        /// The name of the security group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Id of resource group which the security_group belongs.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// The type of the security group.
        /// </summary>
        public readonly string SecurityGroupType;
        /// <summary>
        /// A map of tags assigned to the ECS instances. It must be in the format:
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var taggedSecurityGroups = AliCloud.Ecs.GetSecurityGroups.Invoke(new()
        ///     {
        ///         Tags = 
        ///         {
        ///             { "tagKey1", "tagValue1" },
        ///             { "tagKey2", "tagValue2" },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// Used to retrieve security groups that belong to the specified VPC ID.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetSecurityGroupsGroupResult(
            string creationTime,

            string description,

            string id,

            bool innerAccess,

            string name,

            string resourceGroupId,

            string securityGroupType,

            ImmutableDictionary<string, object>? tags,

            string vpcId)
        {
            CreationTime = creationTime;
            Description = description;
            Id = id;
            InnerAccess = innerAccess;
            Name = name;
            ResourceGroupId = resourceGroupId;
            SecurityGroupType = securityGroupType;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
