// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dfs
{
    /// <summary>
    /// Provides a DFS Mount Point resource.
    /// 
    /// For information about DFS Mount Point and how to use it, see [What is Mount Point](https://www.alibabacloud.com/help/doc-detail/207144.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.140.0.
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
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultZones = AliCloud.Dfs.GetZones.Invoke();
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.4.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.4.0.0/24",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.ZoneId),
    ///     });
    /// 
    ///     var defaultFileSystem = new AliCloud.Dfs.FileSystem("defaultFileSystem", new()
    ///     {
    ///         StorageType = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Options[0]?.StorageType),
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.ZoneId),
    ///         ProtocolType = "HDFS",
    ///         Description = name,
    ///         FileSystemName = name,
    ///         ThroughputMode = "Standard",
    ///         SpaceCapacity = 1024,
    ///     });
    /// 
    ///     var defaultAccessGroup = new AliCloud.Dfs.AccessGroup("defaultAccessGroup", new()
    ///     {
    ///         AccessGroupName = name,
    ///         Description = name,
    ///         NetworkType = "VPC",
    ///     });
    /// 
    ///     var defaultMountPoint = new AliCloud.Dfs.MountPoint("defaultMountPoint", new()
    ///     {
    ///         Description = name,
    ///         VpcId = defaultNetwork.Id,
    ///         FileSystemId = defaultFileSystem.Id,
    ///         AccessGroupId = defaultAccessGroup.Id,
    ///         NetworkType = "VPC",
    ///         VswitchId = defaultSwitch.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DFS Mount Point can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:dfs/mountPoint:MountPoint example &lt;file_system_id&gt;:&lt;mount_point_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dfs/mountPoint:MountPoint")]
    public partial class MountPoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Access Group.
        /// </summary>
        [Output("accessGroupId")]
        public Output<string> AccessGroupId { get; private set; } = null!;

        /// <summary>
        /// The description of the Mount Point.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the File System.
        /// </summary>
        [Output("fileSystemId")]
        public Output<string> FileSystemId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Mount Point.
        /// </summary>
        [Output("mountPointId")]
        public Output<string> MountPointId { get; private set; } = null!;

        /// <summary>
        /// The network type of the Mount Point. Valid values: `VPC`.
        /// </summary>
        [Output("networkType")]
        public Output<string> NetworkType { get; private set; } = null!;

        /// <summary>
        /// The status of the Mount Point. Valid values: `Active`, `Inactive`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The vpc id.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The vswitch id.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a MountPoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MountPoint(string name, MountPointArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dfs/mountPoint:MountPoint", name, args ?? new MountPointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MountPoint(string name, Input<string> id, MountPointState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dfs/mountPoint:MountPoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MountPoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MountPoint Get(string name, Input<string> id, MountPointState? state = null, CustomResourceOptions? options = null)
        {
            return new MountPoint(name, id, state, options);
        }
    }

    public sealed class MountPointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Access Group.
        /// </summary>
        [Input("accessGroupId", required: true)]
        public Input<string> AccessGroupId { get; set; } = null!;

        /// <summary>
        /// The description of the Mount Point.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the File System.
        /// </summary>
        [Input("fileSystemId", required: true)]
        public Input<string> FileSystemId { get; set; } = null!;

        /// <summary>
        /// The network type of the Mount Point. Valid values: `VPC`.
        /// </summary>
        [Input("networkType", required: true)]
        public Input<string> NetworkType { get; set; } = null!;

        /// <summary>
        /// The status of the Mount Point. Valid values: `Active`, `Inactive`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The vpc id.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The vswitch id.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public MountPointArgs()
        {
        }
        public static new MountPointArgs Empty => new MountPointArgs();
    }

    public sealed class MountPointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Access Group.
        /// </summary>
        [Input("accessGroupId")]
        public Input<string>? AccessGroupId { get; set; }

        /// <summary>
        /// The description of the Mount Point.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the File System.
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        /// <summary>
        /// The ID of the Mount Point.
        /// </summary>
        [Input("mountPointId")]
        public Input<string>? MountPointId { get; set; }

        /// <summary>
        /// The network type of the Mount Point. Valid values: `VPC`.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The status of the Mount Point. Valid values: `Active`, `Inactive`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The vpc id.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The vswitch id.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public MountPointState()
        {
        }
        public static new MountPointState Empty => new MountPointState();
    }
}
