// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dts
{
    public static class GetMigrationJobs
    {
        /// <summary>
        /// This data source provides the Dts Migration Jobs of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.157.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Dts.GetMigrationJobs.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "dts_job_id",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dtsMigrationJobId1"] = ids.Apply(getMigrationJobsResult =&gt; getMigrationJobsResult.Jobs[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMigrationJobsResult> InvokeAsync(GetMigrationJobsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMigrationJobsResult>("alicloud:dts/getMigrationJobs:getMigrationJobs", args ?? new GetMigrationJobsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Dts Migration Jobs of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.157.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ids = AliCloud.Dts.GetMigrationJobs.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "dts_job_id",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dtsMigrationJobId1"] = ids.Apply(getMigrationJobsResult =&gt; getMigrationJobsResult.Jobs[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMigrationJobsResult> Invoke(GetMigrationJobsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMigrationJobsResult>("alicloud:dts/getMigrationJobs:getMigrationJobs", args ?? new GetMigrationJobsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMigrationJobsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Synchronization Job IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Migration Job name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetMigrationJobsArgs()
        {
        }
        public static new GetMigrationJobsArgs Empty => new GetMigrationJobsArgs();
    }

    public sealed class GetMigrationJobsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Synchronization Job IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Migration Job name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetMigrationJobsInvokeArgs()
        {
        }
        public static new GetMigrationJobsInvokeArgs Empty => new GetMigrationJobsInvokeArgs();
    }


    [OutputType]
    public sealed class GetMigrationJobsResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetMigrationJobsJobResult> Jobs;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetMigrationJobsResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetMigrationJobsJobResult> jobs,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            Jobs = jobs;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}
