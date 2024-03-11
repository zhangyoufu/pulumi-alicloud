// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Adb
{
    public static class GetResourceGroups
    {
        /// <summary>
        /// This data source provides Adb Resource Group available to the user.[What is Resource Group](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2019-03-15-describedbresourcegroup)
        /// 
        /// &gt; **NOTE:** Available since v1.195.0.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.Adb.GetResourceGroups.Invoke(new()
        ///     {
        ///         DbClusterId = "am-bp1a16357gty69185",
        ///         GroupName = "TESTOPENAPI",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudAdbResourceGroupExampleId"] = @default.Apply(@default =&gt; @default.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Groups[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetResourceGroupsResult> InvokeAsync(GetResourceGroupsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceGroupsResult>("alicloud:adb/getResourceGroups:getResourceGroups", args ?? new GetResourceGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Adb Resource Group available to the user.[What is Resource Group](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2019-03-15-describedbresourcegroup)
        /// 
        /// &gt; **NOTE:** Available since v1.195.0.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.Adb.GetResourceGroups.Invoke(new()
        ///     {
        ///         DbClusterId = "am-bp1a16357gty69185",
        ///         GroupName = "TESTOPENAPI",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudAdbResourceGroupExampleId"] = @default.Apply(@default =&gt; @default.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Groups[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetResourceGroupsResult> Invoke(GetResourceGroupsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceGroupsResult>("alicloud:adb/getResourceGroups:getResourceGroups", args ?? new GetResourceGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// DBClusterId
        /// </summary>
        [Input("dbClusterId", required: true)]
        public string DbClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the resource pool, which cannot exceed 64 bytes in length.
        /// </summary>
        [Input("groupName")]
        public string? GroupName { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of AnalyticDB for MySQL (ADB) Resource Group IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetResourceGroupsArgs()
        {
        }
        public static new GetResourceGroupsArgs Empty => new GetResourceGroupsArgs();
    }

    public sealed class GetResourceGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// DBClusterId
        /// </summary>
        [Input("dbClusterId", required: true)]
        public Input<string> DbClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the resource pool, which cannot exceed 64 bytes in length.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of AnalyticDB for MySQL (ADB) Resource Group IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetResourceGroupsInvokeArgs()
        {
        }
        public static new GetResourceGroupsInvokeArgs Empty => new GetResourceGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceGroupsResult
    {
        /// <summary>
        /// DB cluster id.
        /// </summary>
        public readonly string DbClusterId;
        /// <summary>
        /// The name of the resource pool.
        /// </summary>
        public readonly string? GroupName;
        /// <summary>
        /// A list of Resource Group Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetResourceGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetResourceGroupsResult(
            string dbClusterId,

            string? groupName,

            ImmutableArray<Outputs.GetResourceGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string? outputFile)
        {
            DbClusterId = dbClusterId;
            GroupName = groupName;
            Groups = groups;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
        }
    }
}
