// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.DBS
{
    public static class GetBackupPlans
    {
        /// <summary>
        /// This data source provides the Dbs Backup Plans of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.185.0+.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var ids = AliCloud.DBS.GetBackupPlans.Invoke();
        /// 
        ///     var nameRegex = AliCloud.DBS.GetBackupPlans.Invoke(new()
        ///     {
        ///         NameRegex = "^my-BackupPlan",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dbsBackupPlanId1"] = ids.Apply(getBackupPlansResult =&gt; getBackupPlansResult.Plans[0]?.Id),
        ///         ["dbsBackupPlanId2"] = nameRegex.Apply(getBackupPlansResult =&gt; getBackupPlansResult.Plans[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetBackupPlansResult> InvokeAsync(GetBackupPlansArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBackupPlansResult>("alicloud:dbs/getBackupPlans:getBackupPlans", args ?? new GetBackupPlansArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Dbs Backup Plans of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.185.0+.
        /// 
        /// ## Example Usage
        /// 
        /// Basic Usage
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
        ///     var ids = AliCloud.DBS.GetBackupPlans.Invoke();
        /// 
        ///     var nameRegex = AliCloud.DBS.GetBackupPlans.Invoke(new()
        ///     {
        ///         NameRegex = "^my-BackupPlan",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dbsBackupPlanId1"] = ids.Apply(getBackupPlansResult =&gt; getBackupPlansResult.Plans[0]?.Id),
        ///         ["dbsBackupPlanId2"] = nameRegex.Apply(getBackupPlansResult =&gt; getBackupPlansResult.Plans[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetBackupPlansResult> Invoke(GetBackupPlansInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackupPlansResult>("alicloud:dbs/getBackupPlans:getBackupPlans", args ?? new GetBackupPlansInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupPlansArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("backupPlanName")]
        public string? BackupPlanName { get; set; }

        /// <summary>
        /// Default to `true`. Set it to `false` can hide the `payment_type` to output.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Backup Plan IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Backup Plan name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetBackupPlansArgs()
        {
        }
        public static new GetBackupPlansArgs Empty => new GetBackupPlansArgs();
    }

    public sealed class GetBackupPlansInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("backupPlanName")]
        public Input<string>? BackupPlanName { get; set; }

        /// <summary>
        /// Default to `true`. Set it to `false` can hide the `payment_type` to output.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Backup Plan IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Backup Plan name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetBackupPlansInvokeArgs()
        {
        }
        public static new GetBackupPlansInvokeArgs Empty => new GetBackupPlansInvokeArgs();
    }


    [OutputType]
    public sealed class GetBackupPlansResult
    {
        public readonly string? BackupPlanName;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly ImmutableArray<Outputs.GetBackupPlansPlanResult> Plans;
        public readonly string? Status;

        [OutputConstructor]
        private GetBackupPlansResult(
            string? backupPlanName,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            ImmutableArray<Outputs.GetBackupPlansPlanResult> plans,

            string? status)
        {
            BackupPlanName = backupPlanName;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            Plans = plans;
            Status = status;
        }
    }
}
