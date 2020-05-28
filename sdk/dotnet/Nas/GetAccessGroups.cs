// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nas
{
    public static class GetAccessGroups
    {
        /// <summary>
        /// This data source provides user-available access groups. Use when you can create mount points
        /// 
        /// &gt; NOTE: Available in 1.35.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ag = Output.Create(AliCloud.Nas.GetAccessGroups.InvokeAsync(new AliCloud.Nas.GetAccessGroupsArgs
        ///         {
        ///             Description = "tf-testAccAccessGroupsdatasource",
        ///             NameRegex = "^foo",
        ///             Type = "Classic",
        ///         }));
        ///         this.AlicloudNasAccessGroupsId = ag.Apply(ag =&gt; ag.Groups[0].Id);
        ///     }
        /// 
        ///     [Output("alicloudNasAccessGroupsId")]
        ///     public Output&lt;string&gt; AlicloudNasAccessGroupsId { get; set; }
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccessGroupsResult> InvokeAsync(GetAccessGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccessGroupsResult>("alicloud:nas/getAccessGroups:getAccessGroups", args ?? new GetAccessGroupsArgs(), options.WithVersion());
    }


    public sealed class GetAccessGroupsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter results by a specific Description.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// A regex string to filter AccessGroups by name. 
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter results by a specific AccessGroupType.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetAccessGroupsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccessGroupsResult
    {
        /// <summary>
        /// Destription of the AccessGroup.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A list of AccessGroups. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of AccessGroup IDs, the value is set to `names` .
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of AccessGroup names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// AccessGroupType of the AccessGroup.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetAccessGroupsResult(
            string? description,

            ImmutableArray<Outputs.GetAccessGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? type)
        {
            Description = description;
            Groups = groups;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Type = type;
        }
    }
}
