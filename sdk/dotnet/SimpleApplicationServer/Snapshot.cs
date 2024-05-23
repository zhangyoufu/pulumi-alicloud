// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.SimpleApplicationServer
{
    /// <summary>
    /// Provides a Simple Application Server Snapshot resource.
    /// 
    /// For information about Simple Application Server Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/doc-detail/190452.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.143.0.
    /// 
    /// ## Example Usage
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var @default = AliCloud.SimpleApplicationServer.GetImages.Invoke(new()
    ///     {
    ///         Platform = "Linux",
    ///     });
    /// 
    ///     var defaultGetServerPlans = AliCloud.SimpleApplicationServer.GetServerPlans.Invoke(new()
    ///     {
    ///         Platform = "Linux",
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.SimpleApplicationServer.Instance("default", new()
    ///     {
    ///         PaymentType = "Subscription",
    ///         PlanId = defaultGetServerPlans.Apply(getServerPlansResult =&gt; getServerPlansResult.Plans[0]?.Id),
    ///         InstanceName = name,
    ///         ImageId = @default.Apply(@default =&gt; @default.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id)),
    ///         Period = 1,
    ///         DataDiskSize = 100,
    ///     });
    /// 
    ///     var defaultGetServerDisks = AliCloud.SimpleApplicationServer.GetServerDisks.Invoke(new()
    ///     {
    ///         InstanceId = defaultInstance.Id,
    ///     });
    /// 
    ///     var defaultSnapshot = new AliCloud.SimpleApplicationServer.Snapshot("default", new()
    ///     {
    ///         DiskId = defaultGetServerDisks.Apply(getServerDisksResult =&gt; getServerDisksResult.Ids[0]),
    ///         SnapshotName = name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Simple Application Server Snapshot can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:simpleapplicationserver/snapshot:Snapshot example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:simpleapplicationserver/snapshot:Snapshot")]
    public partial class Snapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the disk.
        /// </summary>
        [Output("diskId")]
        public Output<string> DiskId { get; private set; } = null!;

        /// <summary>
        /// The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
        /// </summary>
        [Output("snapshotName")]
        public Output<string> SnapshotName { get; private set; } = null!;

        /// <summary>
        /// The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs args, CustomResourceOptions? options = null)
            : base("alicloud:simpleapplicationserver/snapshot:Snapshot", name, args ?? new SnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:simpleapplicationserver/snapshot:Snapshot", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Snapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snapshot Get(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new Snapshot(name, id, state, options);
        }
    }

    public sealed class SnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the disk.
        /// </summary>
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        /// <summary>
        /// The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
        /// </summary>
        [Input("snapshotName", required: true)]
        public Input<string> SnapshotName { get; set; } = null!;

        public SnapshotArgs()
        {
        }
        public static new SnapshotArgs Empty => new SnapshotArgs();
    }

    public sealed class SnapshotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the disk.
        /// </summary>
        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        /// <summary>
        /// The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
        /// </summary>
        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        /// <summary>
        /// The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SnapshotState()
        {
        }
        public static new SnapshotState Empty => new SnapshotState();
    }
}
