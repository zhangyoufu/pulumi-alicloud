// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vod.Outputs
{

    [OutputType]
    public sealed class GetDomainsDomainResult
    {
        /// <summary>
        /// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
        /// </summary>
        public readonly string Cname;
        /// <summary>
        /// The description of the domain name for CDN.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The domain name for CDN.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
        /// </summary>
        public readonly string GmtCreated;
        /// <summary>
        /// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
        /// </summary>
        public readonly string GmtModified;
        /// <summary>
        /// The ID of the Domain. Its value is same as Queue Name.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates whether the domain name for CDN is in a sandbox environment.
        /// </summary>
        public readonly string SandBox;
        /// <summary>
        /// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainsDomainSourceResult> Sources;
        /// <summary>
        /// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
        /// </summary>
        public readonly string SslProtocol;
        /// <summary>
        /// The status of the resource.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetDomainsDomainResult(
            string cname,

            string description,

            string domainName,

            string gmtCreated,

            string gmtModified,

            string id,

            string sandBox,

            ImmutableArray<Outputs.GetDomainsDomainSourceResult> sources,

            string sslProtocol,

            string status)
        {
            Cname = cname;
            Description = description;
            DomainName = domainName;
            GmtCreated = gmtCreated;
            GmtModified = gmtModified;
            Id = id;
            SandBox = sandBox;
            Sources = sources;
            SslProtocol = sslProtocol;
            Status = status;
        }
    }
}
