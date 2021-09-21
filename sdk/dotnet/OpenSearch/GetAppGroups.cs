// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.OpenSearch
{
    public static class GetAppGroups
    {
        /// <summary>
        /// This data source provides the Open Search App Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.136.0+.
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
        ///         var config = new Config();
        ///         var name = config.Get("name") ?? "tf_testacc";
        ///         var defaultAppGroup = new AliCloud.OpenSearch.AppGroup("defaultAppGroup", new AliCloud.OpenSearch.AppGroupArgs
        ///         {
        ///             AppGroupName = name,
        ///             PaymentType = "PayAsYouGo",
        ///             Type = "standard",
        ///             Quota = new AliCloud.OpenSearch.Inputs.AppGroupQuotaArgs
        ///             {
        ///                 DocSize = 1,
        ///                 ComputeResource = 20,
        ///                 Spec = "opensearch.share.common",
        ///             },
        ///         });
        ///         var defaultAppGroups = defaultAppGroup.Id.Apply(id =&gt; AliCloud.OpenSearch.GetAppGroups.InvokeAsync(new AliCloud.OpenSearch.GetAppGroupsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 id,
        ///             },
        ///         }));
        ///         this.AppGroups = defaultAppGroups.Apply(defaultAppGroups =&gt; defaultAppGroups.Groups);
        ///     }
        /// 
        ///     [Output("appGroups")]
        ///     public Output&lt;string&gt; AppGroups { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppGroupsResult> InvokeAsync(GetAppGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppGroupsResult>("alicloud:opensearch/getAppGroups:getAppGroups", args ?? new GetAppGroupsArgs(), options.WithVersion());
    }


    public sealed class GetAppGroupsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of App Group IDs. Its element value is same as App Group Name.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The Instance ID.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A regex string to filter results by App Group name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The Resource Group ID.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        /// <summary>
        /// Application type. Valid Values: `standard`, `enhanced`.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetAppGroupsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppGroupsResult
    {
        public readonly bool? EnableDetails;
        public readonly ImmutableArray<Outputs.GetAppGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? InstanceId;
        public readonly string? Name;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? ResourceGroupId;
        public readonly string? Type;

        [OutputConstructor]
        private GetAppGroupsResult(
            bool? enableDetails,

            ImmutableArray<Outputs.GetAppGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string? instanceId,

            string? name,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? resourceGroupId,

            string? type)
        {
            EnableDetails = enableDetails;
            Groups = groups;
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            Name = name;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            Type = type;
        }
    }
}
