// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb
{
    public static class GetServerGroups
    {
        /// <summary>
        /// This data source provides the Alb Server Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.131.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ids = Output.Create(AliCloud.Alb.GetServerGroups.InvokeAsync());
        ///         this.AlbServerGroupId1 = ids.Apply(ids =&gt; ids.Groups[0].Id);
        ///         var nameRegex = Output.Create(AliCloud.Alb.GetServerGroups.InvokeAsync(new AliCloud.Alb.GetServerGroupsArgs
        ///         {
        ///             NameRegex = "^my-ServerGroup",
        ///         }));
        ///         this.AlbServerGroupId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Groups[0].Id);
        ///     }
        /// 
        ///     [Output("albServerGroupId1")]
        ///     public Output&lt;string&gt; AlbServerGroupId1 { get; set; }
        ///     [Output("albServerGroupId2")]
        ///     public Output&lt;string&gt; AlbServerGroupId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServerGroupsResult> InvokeAsync(GetServerGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerGroupsResult>("alicloud:alb/getServerGroups:getServerGroups", args ?? new GetServerGroupsArgs(), options.WithVersion());
    }


    public sealed class GetServerGroupsArgs : Pulumi.InvokeArgs
    {
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Server Group IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Server Group name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("serverGroupIds")]
        private List<string>? _serverGroupIds;

        /// <summary>
        /// The server group ids.
        /// </summary>
        public List<string> ServerGroupIds
        {
            get => _serverGroupIds ?? (_serverGroupIds = new List<string>());
            set => _serverGroupIds = value;
        }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("serverGroupName")]
        public string? ServerGroupName { get; set; }

        /// <summary>
        /// The status of the resource. Valid values: `Provisioning`, `Available` and `Configuring`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VPC that you want to access.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetServerGroupsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerGroupsResult
    {
        public readonly bool? EnableDetails;
        public readonly ImmutableArray<Outputs.GetServerGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? ResourceGroupId;
        public readonly ImmutableArray<string> ServerGroupIds;
        public readonly string? ServerGroupName;
        public readonly string? Status;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? VpcId;

        [OutputConstructor]
        private GetServerGroupsResult(
            bool? enableDetails,

            ImmutableArray<Outputs.GetServerGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? resourceGroupId,

            ImmutableArray<string> serverGroupIds,

            string? serverGroupName,

            string? status,

            ImmutableDictionary<string, object>? tags,

            string? vpcId)
        {
            EnableDetails = enableDetails;
            Groups = groups;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            ServerGroupIds = serverGroupIds;
            ServerGroupName = serverGroupName;
            Status = status;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
