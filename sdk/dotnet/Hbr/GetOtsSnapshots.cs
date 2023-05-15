// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr
{
    public static class GetOtsSnapshots
    {
        /// <summary>
        /// This data source provides the Hbr Ots Snapshots of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.164.0+.
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
        ///     var snapshots = AliCloud.Hbr.GetOtsSnapshots.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOtsSnapshotsResult> InvokeAsync(GetOtsSnapshotsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOtsSnapshotsResult>("alicloud:hbr/getOtsSnapshots:getOtsSnapshots", args ?? new GetOtsSnapshotsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Hbr Ots Snapshots of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.164.0+.
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
        ///     var snapshots = AliCloud.Hbr.GetOtsSnapshots.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOtsSnapshotsResult> Invoke(GetOtsSnapshotsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOtsSnapshotsResult>("alicloud:hbr/getOtsSnapshots:getOtsSnapshots", args ?? new GetOtsSnapshotsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOtsSnapshotsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The end time of the backup. This value must be a UNIX timestamp. Unit: milliseconds
        /// </summary>
        [Input("endTime")]
        public string? EndTime { get; set; }

        [Input("ids")]
        private List<string>? _ids;
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

        /// <summary>
        /// The start time of the backup snapshot. This value is a UNIX timestamp. Unit: seconds.
        /// </summary>
        [Input("startTime")]
        public string? StartTime { get; set; }

        public GetOtsSnapshotsArgs()
        {
        }
        public static new GetOtsSnapshotsArgs Empty => new GetOtsSnapshotsArgs();
    }

    public sealed class GetOtsSnapshotsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The end time of the backup. This value must be a UNIX timestamp. Unit: milliseconds
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;
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

        /// <summary>
        /// The start time of the backup snapshot. This value is a UNIX timestamp. Unit: seconds.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public GetOtsSnapshotsInvokeArgs()
        {
        }
        public static new GetOtsSnapshotsInvokeArgs Empty => new GetOtsSnapshotsInvokeArgs();
    }


    [OutputType]
    public sealed class GetOtsSnapshotsResult
    {
        public readonly string? EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetOtsSnapshotsSnapshotResult> Snapshots;
        public readonly string? StartTime;

        [OutputConstructor]
        private GetOtsSnapshotsResult(
            string? endTime,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            ImmutableArray<Outputs.GetOtsSnapshotsSnapshotResult> snapshots,

            string? startTime)
        {
            EndTime = endTime;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Snapshots = snapshots;
            StartTime = startTime;
        }
    }
}
