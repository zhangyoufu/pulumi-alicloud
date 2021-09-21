// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr
{
    public static class GetEcsBackupPlans
    {
        /// <summary>
        /// This data source provides the Hbr EcsBackupPlans of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.132.0+.
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
        ///         var ids = Output.Create(AliCloud.Hbr.GetEcsBackupPlans.InvokeAsync(new AliCloud.Hbr.GetEcsBackupPlansArgs
        ///         {
        ///             NameRegex = "^my-EcsBackupPlan",
        ///         }));
        ///         this.HbrEcsBackupPlanId = ids.Apply(ids =&gt; ids.Plans[0].Id);
        ///     }
        /// 
        ///     [Output("hbrEcsBackupPlanId")]
        ///     public Output&lt;string&gt; HbrEcsBackupPlanId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEcsBackupPlansResult> InvokeAsync(GetEcsBackupPlansArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEcsBackupPlansResult>("alicloud:hbr/getEcsBackupPlans:getEcsBackupPlans", args ?? new GetEcsBackupPlansArgs(), options.WithVersion());
    }


    public sealed class GetEcsBackupPlansArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of EcsBackupPlan IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of ECS instance.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// A regex string to filter results by EcsBackupPlan name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of Backup vault.
        /// </summary>
        [Input("vaultId")]
        public string? VaultId { get; set; }

        public GetEcsBackupPlansArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEcsBackupPlansResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? InstanceId;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetEcsBackupPlansPlanResult> Plans;
        public readonly string? VaultId;

        [OutputConstructor]
        private GetEcsBackupPlansResult(
            string id,

            ImmutableArray<string> ids,

            string? instanceId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetEcsBackupPlansPlanResult> plans,

            string? vaultId)
        {
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Plans = plans;
            VaultId = vaultId;
        }
    }
}
