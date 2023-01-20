// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ebs
{
    public static class GetDiskReplicaPairs
    {
        /// <summary>
        /// This data source provides Ebs Disk Replica Pair available to the user.
        /// 
        /// &gt; **NOTE:** Available in 1.196.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.Ebs.GetDiskReplicaPairs.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             alicloud_ebs_disk_replica_pair.Default.Id,
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudEbsDiskReplicaPairExampleId"] = @default.Apply(getDiskReplicaPairsResult =&gt; getDiskReplicaPairsResult).Apply(@default =&gt; @default.Apply(getDiskReplicaPairsResult =&gt; getDiskReplicaPairsResult.Pairs[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDiskReplicaPairsResult> InvokeAsync(GetDiskReplicaPairsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDiskReplicaPairsResult>("alicloud:ebs/getDiskReplicaPairs:getDiskReplicaPairs", args ?? new GetDiskReplicaPairsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Ebs Disk Replica Pair available to the user.
        /// 
        /// &gt; **NOTE:** Available in 1.196.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.Ebs.GetDiskReplicaPairs.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             alicloud_ebs_disk_replica_pair.Default.Id,
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudEbsDiskReplicaPairExampleId"] = @default.Apply(getDiskReplicaPairsResult =&gt; getDiskReplicaPairsResult).Apply(@default =&gt; @default.Apply(getDiskReplicaPairsResult =&gt; getDiskReplicaPairsResult.Pairs[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDiskReplicaPairsResult> Invoke(GetDiskReplicaPairsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiskReplicaPairsResult>("alicloud:ebs/getDiskReplicaPairs:getDiskReplicaPairs", args ?? new GetDiskReplicaPairsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDiskReplicaPairsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Disk Replica Pair IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Consistent Replication Group ID, you can specify a consistent replication group ID to query the replication pairs within the group.
        /// </summary>
        [Input("replicaGroupId")]
        public string? ReplicaGroupId { get; set; }

        /// <summary>
        /// Get data for replication pairs where this Region is the production site or the disaster recovery site.
        /// </summary>
        [Input("site")]
        public string? Site { get; set; }

        public GetDiskReplicaPairsArgs()
        {
        }
        public static new GetDiskReplicaPairsArgs Empty => new GetDiskReplicaPairsArgs();
    }

    public sealed class GetDiskReplicaPairsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Disk Replica Pair IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Consistent Replication Group ID, you can specify a consistent replication group ID to query the replication pairs within the group.
        /// </summary>
        [Input("replicaGroupId")]
        public Input<string>? ReplicaGroupId { get; set; }

        /// <summary>
        /// Get data for replication pairs where this Region is the production site or the disaster recovery site.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        public GetDiskReplicaPairsInvokeArgs()
        {
        }
        public static new GetDiskReplicaPairsInvokeArgs Empty => new GetDiskReplicaPairsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDiskReplicaPairsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Disk Replica Pair IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of Disk Replica Pair Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDiskReplicaPairsPairResult> Pairs;
        public readonly string? ReplicaGroupId;
        public readonly string? Site;

        [OutputConstructor]
        private GetDiskReplicaPairsResult(
            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            ImmutableArray<Outputs.GetDiskReplicaPairsPairResult> pairs,

            string? replicaGroupId,

            string? site)
        {
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Pairs = pairs;
            ReplicaGroupId = replicaGroupId;
            Site = site;
        }
    }
}
