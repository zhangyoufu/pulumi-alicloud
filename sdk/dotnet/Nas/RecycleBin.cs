// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nas
{
    /// <summary>
    /// Provides a Network Attached Storage (NAS) Recycle Bin resource.
    /// 
    /// For information about Network Attached Storage (NAS) Recycle Bin and how to use it, see [What is Recycle Bin](https://www.alibabacloud.com/help/en/doc-detail/264185.html).
    /// 
    /// &gt; **NOTE:** Available in v1.155.0+.
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
    ///     var example = AliCloud.Nas.GetZones.Invoke(new()
    ///     {
    ///         FileSystemType = "standard",
    ///     });
    /// 
    ///     var exampleFileSystem = new AliCloud.Nas.FileSystem("example", new()
    ///     {
    ///         ProtocolType = "NFS",
    ///         StorageType = "Performance",
    ///         Description = "terraform-example",
    ///         EncryptType = 1,
    ///         ZoneId = example.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.ZoneId),
    ///     });
    /// 
    ///     var exampleRecycleBin = new AliCloud.Nas.RecycleBin("example", new()
    ///     {
    ///         FileSystemId = exampleFileSystem.Id,
    ///         ReservedDays = 3,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Network Attached Storage (NAS) Recycle Bin can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:nas/recycleBin:RecycleBin example &lt;file_system_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:nas/recycleBin:RecycleBin")]
    public partial class RecycleBin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the file system for which you want to enable the recycle bin feature.
        /// </summary>
        [Output("fileSystemId")]
        public Output<string> FileSystemId { get; private set; } = null!;

        /// <summary>
        /// The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
        /// </summary>
        [Output("reservedDays")]
        public Output<int> ReservedDays { get; private set; } = null!;

        /// <summary>
        /// The status of the recycle bin.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a RecycleBin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RecycleBin(string name, RecycleBinArgs args, CustomResourceOptions? options = null)
            : base("alicloud:nas/recycleBin:RecycleBin", name, args ?? new RecycleBinArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RecycleBin(string name, Input<string> id, RecycleBinState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:nas/recycleBin:RecycleBin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RecycleBin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RecycleBin Get(string name, Input<string> id, RecycleBinState? state = null, CustomResourceOptions? options = null)
        {
            return new RecycleBin(name, id, state, options);
        }
    }

    public sealed class RecycleBinArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the file system for which you want to enable the recycle bin feature.
        /// </summary>
        [Input("fileSystemId", required: true)]
        public Input<string> FileSystemId { get; set; } = null!;

        /// <summary>
        /// The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
        /// </summary>
        [Input("reservedDays")]
        public Input<int>? ReservedDays { get; set; }

        public RecycleBinArgs()
        {
        }
        public static new RecycleBinArgs Empty => new RecycleBinArgs();
    }

    public sealed class RecycleBinState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the file system for which you want to enable the recycle bin feature.
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        /// <summary>
        /// The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
        /// </summary>
        [Input("reservedDays")]
        public Input<int>? ReservedDays { get; set; }

        /// <summary>
        /// The status of the recycle bin.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RecycleBinState()
        {
        }
        public static new RecycleBinState Empty => new RecycleBinState();
    }
}
