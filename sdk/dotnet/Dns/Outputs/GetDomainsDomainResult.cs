// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dns.Outputs
{

    [OutputType]
    public sealed class GetDomainsDomainResult
    {
        /// <summary>
        /// Specifies whether the domain is from Alibaba Cloud or not.
        /// </summary>
        public readonly bool AliDomain;
        /// <summary>
        /// DNS list of the domain in the analysis system.
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// ID of the domain.
        /// </summary>
        public readonly string DomainId;
        /// <summary>
        /// Name of the domain.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// Id of group that contains the domain.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// Name of group that contains the domain.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Cloud analysis product ID.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Punycode of the Chinese domain.
        /// </summary>
        public readonly string PunyCode;
        /// <summary>
        /// Cloud analysis version code.
        /// </summary>
        public readonly string VersionCode;

        [OutputConstructor]
        private GetDomainsDomainResult(
            bool aliDomain,

            ImmutableArray<string> dnsServers,

            string domainId,

            string domainName,

            string groupId,

            string groupName,

            string instanceId,

            string punyCode,

            string versionCode)
        {
            AliDomain = aliDomain;
            DnsServers = dnsServers;
            DomainId = domainId;
            DomainName = domainName;
            GroupId = groupId;
            GroupName = groupName;
            InstanceId = instanceId;
            PunyCode = punyCode;
            VersionCode = versionCode;
        }
    }
}
