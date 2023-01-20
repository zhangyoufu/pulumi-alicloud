// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Adb.Outputs
{

    [OutputType]
    public sealed class GetResourceGroupsGroupResult
    {
        /// <summary>
        /// Creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// DBClusterId
        /// </summary>
        public readonly string DbClusterId;
        /// <summary>
        /// The name of the resource pool, which cannot exceed 64 bytes in length.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Query type, value description:
        /// * **etl**: Batch query mode.
        /// * **interactive**: interactive Query mode
        /// * **default_type**: the default query mode.
        /// </summary>
        public readonly string GroupType;
        /// <summary>
        /// The `key` of the resource supplied above.The value is formulated as `&lt;db_cluster_id&gt;:&lt;group_name&gt;`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
        /// </summary>
        public readonly int NodeNum;
        /// <summary>
        /// Binding User.
        /// </summary>
        public readonly string User;

        [OutputConstructor]
        private GetResourceGroupsGroupResult(
            string createTime,

            string dbClusterId,

            string groupName,

            string groupType,

            string id,

            int nodeNum,

            string user)
        {
            CreateTime = createTime;
            DbClusterId = dbClusterId;
            GroupName = groupName;
            GroupType = groupType;
            Id = id;
            NodeNum = nodeNum;
            User = user;
        }
    }
}
