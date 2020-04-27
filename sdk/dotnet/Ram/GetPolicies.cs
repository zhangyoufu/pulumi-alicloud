// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ram
{
    public static class GetPolicies
    {
        /// <summary>
        /// This data source provides a list of RAM policies in an Alibaba Cloud account according to the specified filters.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPoliciesResult> InvokeAsync(GetPoliciesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPoliciesResult>("alicloud:ram/getPolicies:getPolicies", args ?? new GetPoliciesArgs(), options.WithVersion());
    }


    public sealed class GetPoliciesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter results by a specific group name. Returned policies are attached to the specified group.
        /// </summary>
        [Input("groupName")]
        public string? GroupName { get; set; }

        /// <summary>
        /// A regex string to filter resulting policies by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter results by a specific role name. Returned policies are attached to the specified role.
        /// </summary>
        [Input("roleName")]
        public string? RoleName { get; set; }

        /// <summary>
        /// Filter results by a specific policy type. Valid values are `Custom` and `System`.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// Filter results by a specific user name. Returned policies are attached to the specified user.
        /// </summary>
        [Input("userName")]
        public string? UserName { get; set; }

        public GetPoliciesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPoliciesResult
    {
        public readonly string? GroupName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of ram group names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of policies. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPoliciesPolicyResult> Policies;
        public readonly string? RoleName;
        /// <summary>
        /// Type of the policy.
        /// </summary>
        public readonly string? Type;
        public readonly string? UserName;

        [OutputConstructor]
        private GetPoliciesResult(
            string? groupName,

            string id,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetPoliciesPolicyResult> policies,

            string? roleName,

            string? type,

            string? userName)
        {
            GroupName = groupName;
            Id = id;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Policies = policies;
            RoleName = roleName;
            Type = type;
            UserName = userName;
        }
    }
}
