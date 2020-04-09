// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ram
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of RAM Roles in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ram_roles.html.markdown.
        /// </summary>
        [Obsolete("Use GetRoles.InvokeAsync() instead")]
        public static Task<GetRolesResult> GetRoles(GetRolesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRolesResult>("alicloud:ram/getRoles:getRoles", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetRoles
    {
        /// <summary>
        /// This data source provides a list of RAM Roles in an Alibaba Cloud account according to the specified filters.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ram_roles.html.markdown.
        /// </summary>
        public static Task<GetRolesResult> InvokeAsync(GetRolesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRolesResult>("alicloud:ram/getRoles:getRoles", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetRolesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of ram role IDs. 
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by the role name.
        /// * `ids` (Optional, Available 1.53.0+) - A list of ram role IDs.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter results by a specific policy name. If you set this parameter without setting `policy_type`, the later will be automatically set to `System`. The resulting roles will be attached to the specified policy.
        /// </summary>
        [Input("policyName")]
        public string? PolicyName { get; set; }

        /// <summary>
        /// Filter results by a specific policy type. Valid values are `Custom` and `System`. If you set this parameter, you must set `policy_name` as well.
        /// </summary>
        [Input("policyType")]
        public string? PolicyType { get; set; }

        public GetRolesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRolesResult
    {
        /// <summary>
        /// A list of ram role IDs. 
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of ram role names. 
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? PolicyName;
        public readonly string? PolicyType;
        /// <summary>
        /// A list of roles. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRolesRolesResult> Roles;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRolesResult(
            ImmutableArray<string> ids,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string? policyName,
            string? policyType,
            ImmutableArray<Outputs.GetRolesRolesResult> roles,
            string id)
        {
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PolicyName = policyName;
            PolicyType = policyType;
            Roles = roles;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetRolesRolesResult
    {
        /// <summary>
        /// Resource descriptor of the role.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Authorization strategy of the role. This parameter is deprecated and replaced by `document`.
        /// </summary>
        public readonly string AssumeRolePolicyDocument;
        /// <summary>
        /// Creation date of the role.
        /// </summary>
        public readonly string CreateDate;
        /// <summary>
        /// Description of the role.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Authorization strategy of the role.
        /// </summary>
        public readonly string Document;
        /// <summary>
        /// Id of the role.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the role.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Update date of the role.
        /// </summary>
        public readonly string UpdateDate;

        [OutputConstructor]
        private GetRolesRolesResult(
            string arn,
            string assumeRolePolicyDocument,
            string createDate,
            string description,
            string document,
            string id,
            string name,
            string updateDate)
        {
            Arn = arn;
            AssumeRolePolicyDocument = assumeRolePolicyDocument;
            CreateDate = createDate;
            Description = description;
            Document = document;
            Id = id;
            Name = name;
            UpdateDate = updateDate;
        }
    }
    }
}
