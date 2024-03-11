// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Gpdb
{
    public static class GetDbInstancePlans
    {
        /// <summary>
        /// This data source provides the Gpdb Db Instance Plans of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.189.0+.
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
        ///     var ids = AliCloud.Gpdb.GetDbInstancePlans.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Gpdb.GetDbInstancePlans.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///         NameRegex = "^my-DBInstancePlan",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["gpdbDbInstancePlanId1"] = ids.Apply(getDbInstancePlansResult =&gt; getDbInstancePlansResult.Plans[0]?.Id),
        ///         ["gpdbDbInstancePlanId2"] = nameRegex.Apply(getDbInstancePlansResult =&gt; getDbInstancePlansResult.Plans[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDbInstancePlansResult> InvokeAsync(GetDbInstancePlansArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDbInstancePlansResult>("alicloud:gpdb/getDbInstancePlans:getDbInstancePlans", args ?? new GetDbInstancePlansArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Gpdb Db Instance Plans of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.189.0+.
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
        ///     var ids = AliCloud.Gpdb.GetDbInstancePlans.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Gpdb.GetDbInstancePlans.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///         NameRegex = "^my-DBInstancePlan",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["gpdbDbInstancePlanId1"] = ids.Apply(getDbInstancePlansResult =&gt; getDbInstancePlansResult.Plans[0]?.Id),
        ///         ["gpdbDbInstancePlanId2"] = nameRegex.Apply(getDbInstancePlansResult =&gt; getDbInstancePlansResult.Plans[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDbInstancePlansResult> Invoke(GetDbInstancePlansInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDbInstancePlansResult>("alicloud:gpdb/getDbInstancePlans:getDbInstancePlans", args ?? new GetDbInstancePlansInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDbInstancePlansArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Database instance.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public string DbInstanceId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of DB Instance Plan IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by DB Instance Plan name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Plan scheduling type. Valid values: `Postpone`, `Regular`.
        /// </summary>
        [Input("planScheduleType")]
        public string? PlanScheduleType { get; set; }

        /// <summary>
        /// The type of the Plan. Valid values: `PauseResume`, `Resize`.
        /// </summary>
        [Input("planType")]
        public string? PlanType { get; set; }

        /// <summary>
        /// The Status of the Plan.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetDbInstancePlansArgs()
        {
        }
        public static new GetDbInstancePlansArgs Empty => new GetDbInstancePlansArgs();
    }

    public sealed class GetDbInstancePlansInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Database instance.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of DB Instance Plan IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by DB Instance Plan name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Plan scheduling type. Valid values: `Postpone`, `Regular`.
        /// </summary>
        [Input("planScheduleType")]
        public Input<string>? PlanScheduleType { get; set; }

        /// <summary>
        /// The type of the Plan. Valid values: `PauseResume`, `Resize`.
        /// </summary>
        [Input("planType")]
        public Input<string>? PlanType { get; set; }

        /// <summary>
        /// The Status of the Plan.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetDbInstancePlansInvokeArgs()
        {
        }
        public static new GetDbInstancePlansInvokeArgs Empty => new GetDbInstancePlansInvokeArgs();
    }


    [OutputType]
    public sealed class GetDbInstancePlansResult
    {
        public readonly string DbInstanceId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? PlanScheduleType;
        public readonly string? PlanType;
        public readonly ImmutableArray<Outputs.GetDbInstancePlansPlanResult> Plans;
        public readonly string? Status;

        [OutputConstructor]
        private GetDbInstancePlansResult(
            string dbInstanceId,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? planScheduleType,

            string? planType,

            ImmutableArray<Outputs.GetDbInstancePlansPlanResult> plans,

            string? status)
        {
            DbInstanceId = dbInstanceId;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PlanScheduleType = planScheduleType;
            PlanType = planType;
            Plans = plans;
            Status = status;
        }
    }
}
