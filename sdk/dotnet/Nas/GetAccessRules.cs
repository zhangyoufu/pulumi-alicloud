// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nas
{
    public static class GetAccessRules
    {
        /// <summary>
        /// This data source provides AccessRule available to the user.
        /// 
        /// &gt; NOTE: Available in 1.35.0+
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccessRulesResult> InvokeAsync(GetAccessRulesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccessRulesResult>("alicloud:nas/getAccessRules:getAccessRules", args ?? new GetAccessRulesArgs(), options.WithVersion());
    }


    public sealed class GetAccessRulesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter results by a specific AccessGroupName.
        /// </summary>
        [Input("accessGroupName", required: true)]
        public string AccessGroupName { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter results by a specific RWAccess. 
        /// </summary>
        [Input("rwAccess")]
        public string? RwAccess { get; set; }

        /// <summary>
        /// Filter results by a specific SourceCidrIp. 
        /// </summary>
        [Input("sourceCidrIp")]
        public string? SourceCidrIp { get; set; }

        /// <summary>
        /// Filter results by a specific UserAccess. 
        /// </summary>
        [Input("userAccess")]
        public string? UserAccess { get; set; }

        public GetAccessRulesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccessRulesResult
    {
        public readonly string AccessGroupName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of rule IDs, Each element set to `access_rule_id` (Each element formats as `&lt;access_group_name&gt;:&lt;access rule id&gt;` before 1.53.0).
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of AccessRules. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessRulesRuleResult> Rules;
        /// <summary>
        /// RWAccess of the AccessRule.
        /// </summary>
        public readonly string? RwAccess;
        /// <summary>
        /// SourceCidrIp of the AccessRule.
        /// </summary>
        public readonly string? SourceCidrIp;
        /// <summary>
        /// UserAccess of the AccessRule
        /// </summary>
        public readonly string? UserAccess;

        [OutputConstructor]
        private GetAccessRulesResult(
            string accessGroupName,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            ImmutableArray<Outputs.GetAccessRulesRuleResult> rules,

            string? rwAccess,

            string? sourceCidrIp,

            string? userAccess)
        {
            AccessGroupName = accessGroupName;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Rules = rules;
            RwAccess = rwAccess;
            SourceCidrIp = sourceCidrIp;
            UserAccess = userAccess;
        }
    }
}
