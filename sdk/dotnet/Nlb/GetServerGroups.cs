// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nlb
{
    public static class GetServerGroups
    {
        /// <summary>
        /// This data source provides the Nlb Server Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.186.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Nlb.GetServerGroups.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Nlb.GetServerGroups.Invoke(new()
        ///     {
        ///         NameRegex = "^my-ServerGroup",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["nlbServerGroupId1"] = ids.Apply(getServerGroupsResult =&gt; getServerGroupsResult.Groups[0]?.Id),
        ///         ["nlbServerGroupId2"] = nameRegex.Apply(getServerGroupsResult =&gt; getServerGroupsResult.Groups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServerGroupsResult> InvokeAsync(GetServerGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerGroupsResult>("alicloud:nlb/getServerGroups:getServerGroups", args ?? new GetServerGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Nlb Server Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.186.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Nlb.GetServerGroups.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Nlb.GetServerGroups.Invoke(new()
        ///     {
        ///         NameRegex = "^my-ServerGroup",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["nlbServerGroupId1"] = ids.Apply(getServerGroupsResult =&gt; getServerGroupsResult.Groups[0]?.Id),
        ///         ["nlbServerGroupId2"] = nameRegex.Apply(getServerGroupsResult =&gt; getServerGroupsResult.Groups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServerGroupsResult> Invoke(GetServerGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerGroupsResult>("alicloud:nlb/getServerGroups:getServerGroups", args ?? new GetServerGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerGroupsArgs : global::Pulumi.InvokeArgs
    {
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

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group to which the security group belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("serverGroupNames")]
        private List<string>? _serverGroupNames;

        /// <summary>
        /// The names of the server groups to be queried.
        /// </summary>
        public List<string> ServerGroupNames
        {
            get => _serverGroupNames ?? (_serverGroupNames = new List<string>());
            set => _serverGroupNames = value;
        }

        /// <summary>
        /// The type of the server group.
        /// </summary>
        [Input("serverGroupType")]
        public string? ServerGroupType { get; set; }

        /// <summary>
        /// The status of the server group.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

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

        public GetServerGroupsArgs()
        {
        }
        public static new GetServerGroupsArgs Empty => new GetServerGroupsArgs();
    }

    public sealed class GetServerGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Server Group IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Server Group name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group to which the security group belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("serverGroupNames")]
        private InputList<string>? _serverGroupNames;

        /// <summary>
        /// The names of the server groups to be queried.
        /// </summary>
        public InputList<string> ServerGroupNames
        {
            get => _serverGroupNames ?? (_serverGroupNames = new InputList<string>());
            set => _serverGroupNames = value;
        }

        /// <summary>
        /// The type of the server group.
        /// </summary>
        [Input("serverGroupType")]
        public Input<string>? ServerGroupType { get; set; }

        /// <summary>
        /// The status of the server group.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetServerGroupsInvokeArgs()
        {
        }
        public static new GetServerGroupsInvokeArgs Empty => new GetServerGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerGroupsResult
    {
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
        public readonly ImmutableArray<string> ServerGroupNames;
        public readonly string? ServerGroupType;
        public readonly string? Status;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetServerGroupsResult(
            ImmutableArray<Outputs.GetServerGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? resourceGroupId,

            ImmutableArray<string> serverGroupNames,

            string? serverGroupType,

            string? status,

            ImmutableDictionary<string, object>? tags)
        {
            Groups = groups;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            ServerGroupNames = serverGroupNames;
            ServerGroupType = serverGroupType;
            Status = status;
            Tags = tags;
        }
    }
}
