// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides the acls in the region.
        /// 
        /// 
        /// ## Entry Block
        /// 
        /// The entry mapping supports the following:
        /// 
        /// * `entry`   - An IP addresses or CIDR blocks.
        /// * `comment` - the comment of the entry.
        /// 
        /// ## Listener Block
        /// 
        /// The Listener mapping supports the following:
        /// 
        /// * `load_balancer_id` - the id of load balancer instance, the listener belongs to.
        /// * `frontend_port` - the listener port.
        /// * `protocol`      - the listener protocol (such as tcp/udp/http/https, etc).
        /// * `acl_type`      - the type of acl (such as white/black).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_acls.html.markdown.
        /// </summary>
        [Obsolete("Use GetAcls.InvokeAsync() instead")]
        public static Task<GetAclsResult> GetAcls(GetAclsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAclsResult>("alicloud:slb/getAcls:getAcls", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetAcls
    {
        /// <summary>
        /// This data source provides the acls in the region.
        /// 
        /// 
        /// ## Entry Block
        /// 
        /// The entry mapping supports the following:
        /// 
        /// * `entry`   - An IP addresses or CIDR blocks.
        /// * `comment` - the comment of the entry.
        /// 
        /// ## Listener Block
        /// 
        /// The Listener mapping supports the following:
        /// 
        /// * `load_balancer_id` - the id of load balancer instance, the listener belongs to.
        /// * `frontend_port` - the listener port.
        /// * `protocol`      - the listener protocol (such as tcp/udp/http/https, etc).
        /// * `acl_type`      - the type of acl (such as white/black).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_acls.html.markdown.
        /// </summary>
        public static Task<GetAclsResult> InvokeAsync(GetAclsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAclsResult>("alicloud:slb/getAcls:getAcls", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetAclsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of acls IDs to filter results.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by acl name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The Id of resource group which acl belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetAclsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAclsResult
    {
        /// <summary>
        /// A list of SLB  acls. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAclsAclsResult> Acls;
        /// <summary>
        /// A list of SLB acls IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of SLB acls names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// Resource group ID.
        /// </summary>
        public readonly string? ResourceGroupId;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAclsResult(
            ImmutableArray<Outputs.GetAclsAclsResult> acls,
            ImmutableArray<string> ids,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string? resourceGroupId,
            ImmutableDictionary<string, object>? tags,
            string id)
        {
            Acls = acls;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetAclsAclsEntryListsResult
    {
        public readonly string Comment;
        public readonly string Entry;

        [OutputConstructor]
        private GetAclsAclsEntryListsResult(
            string comment,
            string entry)
        {
            Comment = comment;
            Entry = entry;
        }
    }

    [OutputType]
    public sealed class GetAclsAclsRelatedListenersResult
    {
        public readonly string AclType;
        public readonly int FrontendPort;
        public readonly string LoadBalancerId;
        public readonly string Protocol;

        [OutputConstructor]
        private GetAclsAclsRelatedListenersResult(
            string aclType,
            int frontendPort,
            string loadBalancerId,
            string protocol)
        {
            AclType = aclType;
            FrontendPort = frontendPort;
            LoadBalancerId = loadBalancerId;
            Protocol = protocol;
        }
    }

    [OutputType]
    public sealed class GetAclsAclsResult
    {
        /// <summary>
        /// A list of entry (IP addresses or CIDR blocks).  Each entry contains two sub-fields as `Entry Block` follows.
        /// </summary>
        public readonly ImmutableArray<GetAclsAclsEntryListsResult> EntryLists;
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
        public readonly ImmutableArray<GetAclsAclsRelatedListenersResult> RelatedListeners;
        /// <summary>
        /// The Id of resource group which acl belongs.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetAclsAclsResult(
            ImmutableArray<GetAclsAclsEntryListsResult> entryLists,
            string id,
            string ipVersion,
            string name,
            ImmutableArray<GetAclsAclsRelatedListenersResult> relatedListeners,
            string resourceGroupId,
            ImmutableDictionary<string, object>? tags)
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
}
