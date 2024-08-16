// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb.Outputs
{

    [OutputType]
    public sealed class GetAclsAclResult
    {
        /// <summary>
        /// A list of entry (IP addresses or CIDR blocks).  Each entry contains two sub-fields as `Entry Block` follows.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAclsAclEntryListResult> EntryLists;
        /// <summary>
        /// Acl ID.
        /// </summary>
        public readonly string Id;
        public readonly string IpVersion;
        /// <summary>
        /// Acl name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of listener are attached by the acl.  Each listener contains four sub-fields as `Listener Block` follows.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAclsAclRelatedListenerResult> RelatedListeners;
        /// <summary>
        /// The Id of resource group which acl belongs.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetAclsAclResult(
            ImmutableArray<Outputs.GetAclsAclEntryListResult> entryLists,

            string id,

            string ipVersion,

            string name,

            ImmutableArray<Outputs.GetAclsAclRelatedListenerResult> relatedListeners,

            string resourceGroupId,

            ImmutableDictionary<string, string>? tags)
        {
            EntryLists = entryLists;
            Id = id;
            IpVersion = ipVersion;
            Name = name;
            RelatedListeners = relatedListeners;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
        }
    }
}
