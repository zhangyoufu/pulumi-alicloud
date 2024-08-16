// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nlb.Outputs
{

    [OutputType]
    public sealed class GetServerGroupsGroupResult
    {
        /// <summary>
        /// The protocol version.
        /// </summary>
        public readonly string AddressIpVersion;
        /// <summary>
        /// Indicates whether connection draining is enabled.
        /// </summary>
        public readonly bool ConnectionDrain;
        /// <summary>
        /// The timeout period of connection draining. Unit: seconds.
        /// </summary>
        public readonly int ConnectionDrainTimeout;
        /// <summary>
        /// The configurations of health checks.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerGroupsGroupHealthCheckResult> HealthChecks;
        /// <summary>
        /// The ID of the Server Group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates whether client address retention is enabled.
        /// </summary>
        public readonly bool PreserveClientIpEnabled;
        /// <summary>
        /// The protocol used to forward requests to the backend servers.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The NLB instance.
        /// </summary>
        public readonly ImmutableArray<string> RelatedLoadBalancerIds;
        /// <summary>
        /// The ID of the resource group to which the security group belongs.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// The routing algorithm.
        /// </summary>
        public readonly string Scheduler;
        /// <summary>
        /// The number of server groups associated with the NLB instance.
        /// </summary>
        public readonly int ServerCount;
        /// <summary>
        /// The name of the server group.
        /// </summary>
        public readonly string ServerGroupName;
        /// <summary>
        /// The type of the server group.
        /// </summary>
        public readonly string ServerGroupType;
        /// <summary>
        /// The status of the server group.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The ID of the VPC to which the server group belongs.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetServerGroupsGroupResult(
            string addressIpVersion,

            bool connectionDrain,

            int connectionDrainTimeout,

            ImmutableArray<Outputs.GetServerGroupsGroupHealthCheckResult> healthChecks,

            string id,

            bool preserveClientIpEnabled,

            string protocol,

            ImmutableArray<string> relatedLoadBalancerIds,

            string resourceGroupId,

            string scheduler,

            int serverCount,

            string serverGroupName,

            string serverGroupType,

            string status,

            ImmutableDictionary<string, string> tags,

            string vpcId)
        {
            AddressIpVersion = addressIpVersion;
            ConnectionDrain = connectionDrain;
            ConnectionDrainTimeout = connectionDrainTimeout;
            HealthChecks = healthChecks;
            Id = id;
            PreserveClientIpEnabled = preserveClientIpEnabled;
            Protocol = protocol;
            RelatedLoadBalancerIds = relatedLoadBalancerIds;
            ResourceGroupId = resourceGroupId;
            Scheduler = scheduler;
            ServerCount = serverCount;
            ServerGroupName = serverGroupName;
            ServerGroupType = serverGroupType;
            Status = status;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
