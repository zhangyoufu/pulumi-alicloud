// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of Security Groups in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/security_groups.html.markdown.
        /// </summary>
        [Obsolete("Use GetSecurityGroups.InvokeAsync() instead")]
        public static Task<GetSecurityGroupsResult> GetSecurityGroups(GetSecurityGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupsResult>("alicloud:ecs/getSecurityGroups:getSecurityGroups", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetSecurityGroups
    {
        /// <summary>
        /// This data source provides a list of Security Groups in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/security_groups.html.markdown.
        /// </summary>
        public static Task<GetSecurityGroupsResult> InvokeAsync(GetSecurityGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupsResult>("alicloud:ecs/getSecurityGroups:getSecurityGroups", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSecurityGroupsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Security Group IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter the resulting security groups by their names.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The Id of resource group which the security_group belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags assigned to the ECS instances. It must be in the format:
        /// ```
        /// data "alicloud.ecs.getSecurityGroups" "taggedSecurityGroups" {
        /// tags = {
        /// tagKey1 = "tagValue1",
        /// tagKey2 = "tagValue2"
        /// }
        /// }
        /// ```
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// Used to retrieve security groups that belong to the specified VPC ID.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetSecurityGroupsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSecurityGroupsResult
    {
        /// <summary>
        /// A list of Security Groups. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSecurityGroupsGroupsResult> Groups;
        /// <summary>
        /// A list of Security Group IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of Security Group names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The Id of resource group which the security_group belongs.
        /// </summary>
        public readonly string? ResourceGroupId;
        /// <summary>
        /// A map of tags assigned to the ECS instance.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// The ID of the VPC that owns the security group.
        /// </summary>
        public readonly string? VpcId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSecurityGroupsResult(
            ImmutableArray<Outputs.GetSecurityGroupsGroupsResult> groups,
            ImmutableArray<string> ids,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string? resourceGroupId,
            ImmutableDictionary<string, object>? tags,
            string? vpcId,
            string id)
        {
            Groups = groups;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
            VpcId = vpcId;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetSecurityGroupsGroupsResult
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
        /// ```
        /// data "alicloud.ecs.getSecurityGroups" "taggedSecurityGroups" {
        /// tags = {
        /// tagKey1 = "tagValue1",
        /// tagKey2 = "tagValue2"
        /// }
        /// }
        /// ```
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// Used to retrieve security groups that belong to the specified VPC ID.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetSecurityGroupsGroupsResult(
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
}
