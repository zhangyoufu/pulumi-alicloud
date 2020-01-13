// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Creates a custom image. You can then use a custom image to create ECS instances (RunInstances) or change the system disk for an existing instance (ReplaceSystemDisk).
    /// 
    /// &gt; **NOTE:**  If you want to create a template from an ECS instance, you can specify the instance ID (InstanceId) to create a custom image. You must make sure that the status of the specified instance is Running or Stopped. After a successful invocation, each disk of the specified instance has a new snapshot created.
    /// 
    /// &gt; **NOTE:**  If you want to create a custom image based on the system disk of your ECS instance, you can specify one of the system disk snapshots (SnapshotId) to create a custom image. However, the specified snapshot cannot be created on or before July 15, 2013.
    /// 
    /// &gt; **NOTE:**  If you want to combine snapshots of multiple disks into an image template, you can specify DiskDeviceMapping to create a custom image.
    /// 
    /// &gt; **NOTE:**  Available in 1.64.0+
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/image.html.markdown.
    /// </summary>
    public partial class Image : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x86_64`.
        /// </summary>
        [Output("architecture")]
        public Output<string?> Architecture { get; private set; } = null!;

        /// <summary>
        /// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Description of the system with disks and snapshots under the image.
        /// </summary>
        [Output("diskDeviceMappings")]
        public Output<ImmutableArray<Outputs.ImageDiskDeviceMappings>> DiskDeviceMappings { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to force delete the custom image, Default is `false`. 
        /// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
        /// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
        /// </summary>
        [Output("imageName")]
        public Output<string> ImageName { get; private set; } = null!;

        /// <summary>
        /// The instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string?> InstanceId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the operating system platform of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `CentOS`, `Ubuntu`, `SUSE`, `OpenSUSE`, `RedHat`, `Debian`, `CoreOS`, `Aliyun Linux`, `Windows Server 2003`, `Windows Server 2008`, `Windows Server 2012`, `Windows 7`, Default is `Others Linux`, `Customized Linux`.
        /// </summary>
        [Output("platform")]
        public Output<string?> Platform { get; private set; } = null!;

        /// <summary>
        /// The ID of the enterprise resource group to which a custom image belongs
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// Specifies a snapshot that is used to create a combined custom image.
        /// </summary>
        [Output("snapshotId")]
        public Output<string?> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// The tag value of an image. The value of N ranges from 1 to 20.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Image resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Image(string name, ImageArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/image:Image", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Image(string name, Input<string> id, ImageState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/image:Image", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Image resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Image Get(string name, Input<string> id, ImageState? state = null, CustomResourceOptions? options = null)
        {
            return new Image(name, id, state, options);
        }
    }

    public sealed class ImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x86_64`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("diskDeviceMappings")]
        private InputList<Inputs.ImageDiskDeviceMappingsArgs>? _diskDeviceMappings;

        /// <summary>
        /// Description of the system with disks and snapshots under the image.
        /// </summary>
        public InputList<Inputs.ImageDiskDeviceMappingsArgs> DiskDeviceMappings
        {
            get => _diskDeviceMappings ?? (_diskDeviceMappings = new InputList<Inputs.ImageDiskDeviceMappingsArgs>());
            set => _diskDeviceMappings = value;
        }

        /// <summary>
        /// Indicates whether to force delete the custom image, Default is `false`. 
        /// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
        /// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// The instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the operating system platform of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `CentOS`, `Ubuntu`, `SUSE`, `OpenSUSE`, `RedHat`, `Debian`, `CoreOS`, `Aliyun Linux`, `Windows Server 2003`, `Windows Server 2008`, `Windows Server 2012`, `Windows 7`, Default is `Others Linux`, `Customized Linux`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The ID of the enterprise resource group to which a custom image belongs
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Specifies a snapshot that is used to create a combined custom image.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag value of an image. The value of N ranges from 1 to 20.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ImageArgs()
        {
        }
    }

    public sealed class ImageState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x86_64`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("diskDeviceMappings")]
        private InputList<Inputs.ImageDiskDeviceMappingsGetArgs>? _diskDeviceMappings;

        /// <summary>
        /// Description of the system with disks and snapshots under the image.
        /// </summary>
        public InputList<Inputs.ImageDiskDeviceMappingsGetArgs> DiskDeviceMappings
        {
            get => _diskDeviceMappings ?? (_diskDeviceMappings = new InputList<Inputs.ImageDiskDeviceMappingsGetArgs>());
            set => _diskDeviceMappings = value;
        }

        /// <summary>
        /// Indicates whether to force delete the custom image, Default is `false`. 
        /// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
        /// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// The instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the operating system platform of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `CentOS`, `Ubuntu`, `SUSE`, `OpenSUSE`, `RedHat`, `Debian`, `CoreOS`, `Aliyun Linux`, `Windows Server 2003`, `Windows Server 2008`, `Windows Server 2012`, `Windows 7`, Default is `Others Linux`, `Customized Linux`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The ID of the enterprise resource group to which a custom image belongs
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// Specifies a snapshot that is used to create a combined custom image.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag value of an image. The value of N ranges from 1 to 20.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ImageState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ImageDiskDeviceMappingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of a disk in the combined custom image. Value range: /dev/xvda to /dev/xvdz.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Specifies the type of a disk in the combined custom image. If you specify this parameter, you can use a data disk snapshot as the data source of a system disk for creating an image. If it is not specified, the disk type is determined by the corresponding snapshot. Valid values: `system`, `data`,
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// Specifies the size of a disk in the combined custom image, in GiB. Value range: 5 to 2000.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Specifies a snapshot that is used to create a combined custom image.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public ImageDiskDeviceMappingsArgs()
        {
        }
    }

    public sealed class ImageDiskDeviceMappingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of a disk in the combined custom image. Value range: /dev/xvda to /dev/xvdz.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Specifies the type of a disk in the combined custom image. If you specify this parameter, you can use a data disk snapshot as the data source of a system disk for creating an image. If it is not specified, the disk type is determined by the corresponding snapshot. Valid values: `system`, `data`,
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// Specifies the size of a disk in the combined custom image, in GiB. Value range: 5 to 2000.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Specifies a snapshot that is used to create a combined custom image.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public ImageDiskDeviceMappingsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ImageDiskDeviceMappings
    {
        /// <summary>
        /// Specifies the name of a disk in the combined custom image. Value range: /dev/xvda to /dev/xvdz.
        /// </summary>
        public readonly string Device;
        /// <summary>
        /// Specifies the type of a disk in the combined custom image. If you specify this parameter, you can use a data disk snapshot as the data source of a system disk for creating an image. If it is not specified, the disk type is determined by the corresponding snapshot. Valid values: `system`, `data`,
        /// </summary>
        public readonly string DiskType;
        /// <summary>
        /// Specifies the size of a disk in the combined custom image, in GiB. Value range: 5 to 2000.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Specifies a snapshot that is used to create a combined custom image.
        /// </summary>
        public readonly string SnapshotId;

        [OutputConstructor]
        private ImageDiskDeviceMappings(
            string device,
            string diskType,
            int size,
            string snapshotId)
        {
            Device = device;
            DiskType = diskType;
            Size = size;
            SnapshotId = snapshotId;
        }
    }
    }
}
