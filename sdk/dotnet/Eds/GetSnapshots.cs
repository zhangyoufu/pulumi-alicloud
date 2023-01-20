// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eds
{
    public static class GetSnapshots
    {
        /// <summary>
        /// This data source provides the Ecd Snapshots of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.169.0+.
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
        ///     var ids = AliCloud.Eds.GetSnapshots.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ecdSnapshotId1"] = ids.Apply(getSnapshotsResult =&gt; getSnapshotsResult.Snapshots[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSnapshotsResult> InvokeAsync(GetSnapshotsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotsResult>("alicloud:eds/getSnapshots:getSnapshots", args ?? new GetSnapshotsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ecd Snapshots of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.169.0+.
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
        ///     var ids = AliCloud.Eds.GetSnapshots.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ecdSnapshotId1"] = ids.Apply(getSnapshotsResult =&gt; getSnapshotsResult.Snapshots[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSnapshotsResult> Invoke(GetSnapshotsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotsResult>("alicloud:eds/getSnapshots:getSnapshots", args ?? new GetSnapshotsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the cloud desktop to which the snapshot belongs.
        /// </summary>
        [Input("desktopId")]
        public string? DesktopId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Snapshot IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Snapshot name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the snapshot.
        /// </summary>
        [Input("snapshotId")]
        public string? SnapshotId { get; set; }

        public GetSnapshotsArgs()
        {
        }
        public static new GetSnapshotsArgs Empty => new GetSnapshotsArgs();
    }

    public sealed class GetSnapshotsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the cloud desktop to which the snapshot belongs.
        /// </summary>
        [Input("desktopId")]
        public Input<string>? DesktopId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Snapshot IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Snapshot name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ID of the snapshot.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public GetSnapshotsInvokeArgs()
        {
        }
        public static new GetSnapshotsInvokeArgs Empty => new GetSnapshotsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotsResult
    {
        public readonly string? DesktopId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? SnapshotId;
        public readonly ImmutableArray<Outputs.GetSnapshotsSnapshotResult> Snapshots;

        [OutputConstructor]
        private GetSnapshotsResult(
            string? desktopId,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? snapshotId,

            ImmutableArray<Outputs.GetSnapshotsSnapshotResult> snapshots)
        {
            DesktopId = desktopId;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            SnapshotId = snapshotId;
            Snapshots = snapshots;
        }
    }
}
