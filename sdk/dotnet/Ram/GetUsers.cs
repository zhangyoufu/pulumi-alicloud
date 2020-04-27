// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ram
{
    public static class GetUsers
    {
        /// <summary>
        /// This data source provides a list of RAM users in an Alibaba Cloud account according to the specified filters.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUsersResult> InvokeAsync(GetUsersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUsersResult>("alicloud:ram/getUsers:getUsers", args ?? new GetUsersArgs(), options.WithVersion());
    }


    public sealed class GetUsersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter results by a specific group name. Returned users are in the specified group. 
        /// </summary>
        [Input("groupName")]
        public string? GroupName { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// - A list of ram user IDs. 
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter resulting users by their names.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter results by a specific policy name. If you set this parameter without setting `policy_type`, the later will be automatically set to `System`. Returned users are attached to the specified policy.
        /// </summary>
        [Input("policyName")]
        public string? PolicyName { get; set; }

        /// <summary>
        /// Filter results by a specific policy type. Valid values are `Custom` and `System`. If you set this parameter, you must set `policy_name` as well.
        /// </summary>
        [Input("policyType")]
        public string? PolicyType { get; set; }

        public GetUsersArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUsersResult
    {
        public readonly string? GroupName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of ram user IDs. 
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of ram user names. 
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? PolicyName;
        public readonly string? PolicyType;
        /// <summary>
        /// A list of users. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsersUserResult> Users;

        [OutputConstructor]
        private GetUsersResult(
            string? groupName,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? policyName,

            string? policyType,

            ImmutableArray<Outputs.GetUsersUserResult> users)
        {
            GroupName = groupName;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PolicyName = policyName;
            PolicyType = policyType;
            Users = users;
        }
    }
}
