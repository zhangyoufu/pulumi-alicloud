// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Provides a ECS Image resource.
    /// 
    /// &gt; **NOTE:**  If you want to create a template from an ECS instance, you can specify the instance ID (InstanceId) to create a custom image. You must make sure that the status of the specified instance is Running or Stopped. After a successful invocation, each disk of the specified instance has a new snapshot created.
    /// 
    /// &gt; **NOTE:**  If you want to create a custom image based on the system disk of your ECS instance, you can specify one of the system disk snapshots (SnapshotId) to create a custom image. However, the specified snapshot cannot be created on or before July 15, 2013.
    /// 
    /// &gt; **NOTE:**  If you want to combine snapshots of multiple disks into an image template, you can specify DiskDeviceMapping to create a custom image.
    /// 
    /// For information about ECS Image and how to use it, see [What is Image](https://www.alibabacloud.com/help/en/ecs/developer-reference/api-ecs-2014-05-26-createimage).
    /// 
    /// &gt; **NOTE:** Available since v1.64.0.
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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "Instance",
    ///     });
    /// 
    ///     var defaultGetInstanceTypes = AliCloud.Ecs.GetInstanceTypes.Invoke(new()
    ///     {
    ///         InstanceTypeFamily = "ecs.sn1ne",
    ///     });
    /// 
    ///     var defaultGetImages = AliCloud.Ecs.GetImages.Invoke(new()
    ///     {
    ///         NameRegex = "^ubuntu_18.*64",
    ///         Owners = "system",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
    ///     {
    ///         VpcName = "terraform-example",
    ///         CidrBlock = "172.17.3.0/24",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("default", new()
    ///     {
    ///         VswitchName = "terraform-example",
    ///         CidrBlock = "172.17.3.0/24",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("default", new()
    ///     {
    ///         Name = "terraform-example",
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Ecs.Instance("default", new()
    ///     {
    ///         AvailabilityZone = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
    ///         InstanceName = "terraform-example",
    ///         SecurityGroups = new[]
    ///         {
    ///             defaultSecurityGroup.Id,
    ///         },
    ///         VswitchId = defaultSwitch.Id,
    ///         InstanceType = defaultGetInstanceTypes.Apply(getInstanceTypesResult =&gt; getInstanceTypesResult.Ids[0]),
    ///         ImageId = defaultGetImages.Apply(getImagesResult =&gt; getImagesResult.Ids[0]),
    ///         InternetMaxBandwidthOut = 10,
    ///     });
    /// 
    ///     var defaultGetResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var defaultInteger = new Random.Index.Integer("default", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var defaultImage = new AliCloud.Ecs.Image("default", new()
    ///     {
    ///         InstanceId = defaultInstance.Id,
    ///         ImageName = $"terraform-example-{defaultInteger.Result}",
    ///         Description = "terraform-example",
    ///         Architecture = "x86_64",
    ///         ResourceGroupId = defaultGetResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         Tags = 
    ///         {
    ///             { "FinanceDept", "FinanceDeptJoshua" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECS Image can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ecs/image:Image example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/image:Image")]
    public partial class Image : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The system architecture of the system disk. If you specify a data disk snapshot to create the system disk of the custom image, you must use Architecture to specify the system architecture of the system disk. Valid values: `i386`, `x86\_64`, `arm64`. Default value: `x86\_64`.
        /// </summary>
        [Output("architecture")]
        public Output<string?> Architecture { get; private set; } = null!;

        /// <summary>
        /// The new boot mode of the image. Valid values:
        /// 
        /// *   BIOS: Basic Input/Output System (BIOS)
        /// 
        /// *   UEFI: Unified Extensible Firmware Interface (UEFI)
        /// 
        /// *   UEFI-Preferred: BIOS and UEFI
        /// 
        /// &gt; **NOTE:**   Before you change the boot mode, we recommend that you obtain the boot modes supported by the image. If you specify an unsupported boot mode for the image, ECS instances that use the image cannot start as expected. If you do not know which boot modes are supported by the image, we recommend that you use the image check feature to perform a check. For information about the image check feature, see [Overview](https://www.alibabacloud.com/help/en/doc-detail/439819.html).
        /// 
        /// &gt; **NOTE:**   For information about the UEFI-Preferred boot mode, see [Best practices for ECS instance boot modes](https://www.alibabacloud.com/help/en/doc-detail/2244655.html).
        /// </summary>
        [Output("bootMode")]
        public Output<string> BootMode { get; private set; } = null!;

        /// <summary>
        /// The create time
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Not the public attribute and it used to automatically delete dependence snapshots while deleting the image.
        /// </summary>
        [Output("deleteAutoSnapshot")]
        public Output<bool?> DeleteAutoSnapshot { get; private set; } = null!;

        /// <summary>
        /// The new description of the custom image. The description must be 2 to 256 characters in length It cannot start with `http://` or `https://`. This parameter is empty by default, which specifies that the original description is retained.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The mode in which to check the custom image. If you do not specify this parameter, the image is not checked. Only the standard check mode is supported.
        /// 
        /// &gt; **NOTE:**   This parameter is supported for most Linux and Windows operating system versions. For information about image check items and operating system limits for image check, see [Overview of image check](https://www.alibabacloud.com/help/en/doc-detail/439819.html) and [Operating system limits for image check](https://www.alibabacloud.com/help/en/doc-detail/475800.html).
        /// </summary>
        [Output("detectionStrategy")]
        public Output<string?> DetectionStrategy { get; private set; } = null!;

        /// <summary>
        /// Snapshot information for the image See `disk_device_mapping` below.
        /// </summary>
        [Output("diskDeviceMappings")]
        public Output<ImmutableArray<Outputs.ImageDiskDeviceMapping>> DiskDeviceMappings { get; private set; } = null!;

        /// <summary>
        /// Features See `features` below.
        /// </summary>
        [Output("features")]
        public Output<Outputs.ImageFeatures> Features { get; private set; } = null!;

        /// <summary>
        /// Whether to perform forced deletion. Value range:
        /// - true: forcibly deletes the custom image, ignoring whether the current image is used by other instances.
        /// - false: The custom image is deleted normally. Before deleting the custom image, check whether the current image is used by other instances.
        /// 
        /// Default value: false
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// The name of the image family. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with acs: or aliyun. It cannot contain http:// or https://. It can contain letters, digits, periods (.), colons (:), underscores (\_), and hyphens (-). By default, this parameter is empty.
        /// </summary>
        [Output("imageFamily")]
        public Output<string?> ImageFamily { get; private set; } = null!;

        /// <summary>
        /// The name of the custom image. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with acs: or aliyun. It cannot contain http:// or https://. It can contain letters, digits, periods (.), colons (:), underscores (\_), and hyphens (-). By default, this parameter is empty. In this case, the original name is retained.
        /// </summary>
        [Output("imageName")]
        public Output<string> ImageName { get; private set; } = null!;

        /// <summary>
        /// The image version.
        /// 
        /// &gt; **NOTE:**  If you specify an instance by configuring `InstanceId`, and the instance uses an Alibaba Cloud Marketplace image or a custom image that is created from an Alibaba Cloud Marketplace image, you must leave this parameter empty or set this parameter to the value of ImageVersion of the instance.
        /// </summary>
        [Output("imageVersion")]
        public Output<string?> ImageVersion { get; private set; } = null!;

        /// <summary>
        /// The instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string?> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The type of the license that is used to activate the operating system after the image is imported. Set the value to BYOL. BYOL: The license that comes with the source operating system is used. When you use the BYOL license, make sure that your license key is supported by Alibaba Cloud.
        /// </summary>
        [Output("licenseType")]
        public Output<string?> LicenseType { get; private set; } = null!;

        /// <summary>
        /// . Field 'name' has been deprecated from provider version 1.227.0. New field 'image_name' instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The operating system distribution for the system disk in the custom image. If you specify a data disk snapshot to create the system disk of the custom image, use Platform to specify the operating system distribution for the system disk. Valid values: `Aliyun`, `Anolis`, `CentOS`, `Ubuntu`, `CoreOS`, `SUSE`, `Debian`, `OpenSUSE`, `FreeBSD`, `RedHat`, `Kylin`, `UOS`, `Fedora`, `Fedora CoreOS`, `CentOS Stream`, `AlmaLinux`, `Rocky Linux`, `Gentoo`, `Customized Linux`, `Others Linux`, `Windows Server 2022`, `Windows Server 2019`, `Windows Server 2016`, `Windows Server 2012`, `Windows Server 2008`, `Windows Server 2003`. Default value: `Others Linux`.
        /// </summary>
        [Output("platform")]
        public Output<string> Platform { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group to which to assign the custom image. If you do not specify this parameter, the image is assigned to the default resource group.
        /// 
        /// &gt; **NOTE:**   If you call the CreateImage operation as a Resource Access Management (RAM) user who does not have the permissions to manage the default resource group and do not specify `ResourceGroupId`, the `Forbbiden: User not authorized to operate on the specified resource` error message is returned. You must specify the ID of a resource group that the RAM user has the permissions to manage or grant the RAM user the permissions to manage the default resource group before you call the CreateImage operation again.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the snapshot that you want to use to create the custom image.
        /// </summary>
        [Output("snapshotId")]
        public Output<string?> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// The status of the image. By default, if you do not specify this parameter, only images in the Available state are returned.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tag
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Image resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Image(string name, ImageArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/image:Image", name, args ?? new ImageArgs(), MakeResourceOptions(options, ""))
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

    public sealed class ImageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The system architecture of the system disk. If you specify a data disk snapshot to create the system disk of the custom image, you must use Architecture to specify the system architecture of the system disk. Valid values: `i386`, `x86\_64`, `arm64`. Default value: `x86\_64`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The new boot mode of the image. Valid values:
        /// 
        /// *   BIOS: Basic Input/Output System (BIOS)
        /// 
        /// *   UEFI: Unified Extensible Firmware Interface (UEFI)
        /// 
        /// *   UEFI-Preferred: BIOS and UEFI
        /// 
        /// &gt; **NOTE:**   Before you change the boot mode, we recommend that you obtain the boot modes supported by the image. If you specify an unsupported boot mode for the image, ECS instances that use the image cannot start as expected. If you do not know which boot modes are supported by the image, we recommend that you use the image check feature to perform a check. For information about the image check feature, see [Overview](https://www.alibabacloud.com/help/en/doc-detail/439819.html).
        /// 
        /// &gt; **NOTE:**   For information about the UEFI-Preferred boot mode, see [Best practices for ECS instance boot modes](https://www.alibabacloud.com/help/en/doc-detail/2244655.html).
        /// </summary>
        [Input("bootMode")]
        public Input<string>? BootMode { get; set; }

        /// <summary>
        /// Not the public attribute and it used to automatically delete dependence snapshots while deleting the image.
        /// </summary>
        [Input("deleteAutoSnapshot")]
        public Input<bool>? DeleteAutoSnapshot { get; set; }

        /// <summary>
        /// The new description of the custom image. The description must be 2 to 256 characters in length It cannot start with `http://` or `https://`. This parameter is empty by default, which specifies that the original description is retained.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The mode in which to check the custom image. If you do not specify this parameter, the image is not checked. Only the standard check mode is supported.
        /// 
        /// &gt; **NOTE:**   This parameter is supported for most Linux and Windows operating system versions. For information about image check items and operating system limits for image check, see [Overview of image check](https://www.alibabacloud.com/help/en/doc-detail/439819.html) and [Operating system limits for image check](https://www.alibabacloud.com/help/en/doc-detail/475800.html).
        /// </summary>
        [Input("detectionStrategy")]
        public Input<string>? DetectionStrategy { get; set; }

        [Input("diskDeviceMappings")]
        private InputList<Inputs.ImageDiskDeviceMappingArgs>? _diskDeviceMappings;

        /// <summary>
        /// Snapshot information for the image See `disk_device_mapping` below.
        /// </summary>
        public InputList<Inputs.ImageDiskDeviceMappingArgs> DiskDeviceMappings
        {
            get => _diskDeviceMappings ?? (_diskDeviceMappings = new InputList<Inputs.ImageDiskDeviceMappingArgs>());
            set => _diskDeviceMappings = value;
        }

        /// <summary>
        /// Features See `features` below.
        /// </summary>
        [Input("features")]
        public Input<Inputs.ImageFeaturesArgs>? Features { get; set; }

        /// <summary>
        /// Whether to perform forced deletion. Value range:
        /// - true: forcibly deletes the custom image, ignoring whether the current image is used by other instances.
        /// - false: The custom image is deleted normally. Before deleting the custom image, check whether the current image is used by other instances.
        /// 
        /// Default value: false
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The name of the image family. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with acs: or aliyun. It cannot contain http:// or https://. It can contain letters, digits, periods (.), colons (:), underscores (\_), and hyphens (-). By default, this parameter is empty.
        /// </summary>
        [Input("imageFamily")]
        public Input<string>? ImageFamily { get; set; }

        /// <summary>
        /// The name of the custom image. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with acs: or aliyun. It cannot contain http:// or https://. It can contain letters, digits, periods (.), colons (:), underscores (\_), and hyphens (-). By default, this parameter is empty. In this case, the original name is retained.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// The image version.
        /// 
        /// &gt; **NOTE:**  If you specify an instance by configuring `InstanceId`, and the instance uses an Alibaba Cloud Marketplace image or a custom image that is created from an Alibaba Cloud Marketplace image, you must leave this parameter empty or set this parameter to the value of ImageVersion of the instance.
        /// </summary>
        [Input("imageVersion")]
        public Input<string>? ImageVersion { get; set; }

        /// <summary>
        /// The instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The type of the license that is used to activate the operating system after the image is imported. Set the value to BYOL. BYOL: The license that comes with the source operating system is used. When you use the BYOL license, make sure that your license key is supported by Alibaba Cloud.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// . Field 'name' has been deprecated from provider version 1.227.0. New field 'image_name' instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The operating system distribution for the system disk in the custom image. If you specify a data disk snapshot to create the system disk of the custom image, use Platform to specify the operating system distribution for the system disk. Valid values: `Aliyun`, `Anolis`, `CentOS`, `Ubuntu`, `CoreOS`, `SUSE`, `Debian`, `OpenSUSE`, `FreeBSD`, `RedHat`, `Kylin`, `UOS`, `Fedora`, `Fedora CoreOS`, `CentOS Stream`, `AlmaLinux`, `Rocky Linux`, `Gentoo`, `Customized Linux`, `Others Linux`, `Windows Server 2022`, `Windows Server 2019`, `Windows Server 2016`, `Windows Server 2012`, `Windows Server 2008`, `Windows Server 2003`. Default value: `Others Linux`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The ID of the resource group to which to assign the custom image. If you do not specify this parameter, the image is assigned to the default resource group.
        /// 
        /// &gt; **NOTE:**   If you call the CreateImage operation as a Resource Access Management (RAM) user who does not have the permissions to manage the default resource group and do not specify `ResourceGroupId`, the `Forbbiden: User not authorized to operate on the specified resource` error message is returned. You must specify the ID of a resource group that the RAM user has the permissions to manage or grant the RAM user the permissions to manage the default resource group before you call the CreateImage operation again.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The ID of the snapshot that you want to use to create the custom image.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ImageArgs()
        {
        }
        public static new ImageArgs Empty => new ImageArgs();
    }

    public sealed class ImageState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The system architecture of the system disk. If you specify a data disk snapshot to create the system disk of the custom image, you must use Architecture to specify the system architecture of the system disk. Valid values: `i386`, `x86\_64`, `arm64`. Default value: `x86\_64`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The new boot mode of the image. Valid values:
        /// 
        /// *   BIOS: Basic Input/Output System (BIOS)
        /// 
        /// *   UEFI: Unified Extensible Firmware Interface (UEFI)
        /// 
        /// *   UEFI-Preferred: BIOS and UEFI
        /// 
        /// &gt; **NOTE:**   Before you change the boot mode, we recommend that you obtain the boot modes supported by the image. If you specify an unsupported boot mode for the image, ECS instances that use the image cannot start as expected. If you do not know which boot modes are supported by the image, we recommend that you use the image check feature to perform a check. For information about the image check feature, see [Overview](https://www.alibabacloud.com/help/en/doc-detail/439819.html).
        /// 
        /// &gt; **NOTE:**   For information about the UEFI-Preferred boot mode, see [Best practices for ECS instance boot modes](https://www.alibabacloud.com/help/en/doc-detail/2244655.html).
        /// </summary>
        [Input("bootMode")]
        public Input<string>? BootMode { get; set; }

        /// <summary>
        /// The create time
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Not the public attribute and it used to automatically delete dependence snapshots while deleting the image.
        /// </summary>
        [Input("deleteAutoSnapshot")]
        public Input<bool>? DeleteAutoSnapshot { get; set; }

        /// <summary>
        /// The new description of the custom image. The description must be 2 to 256 characters in length It cannot start with `http://` or `https://`. This parameter is empty by default, which specifies that the original description is retained.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The mode in which to check the custom image. If you do not specify this parameter, the image is not checked. Only the standard check mode is supported.
        /// 
        /// &gt; **NOTE:**   This parameter is supported for most Linux and Windows operating system versions. For information about image check items and operating system limits for image check, see [Overview of image check](https://www.alibabacloud.com/help/en/doc-detail/439819.html) and [Operating system limits for image check](https://www.alibabacloud.com/help/en/doc-detail/475800.html).
        /// </summary>
        [Input("detectionStrategy")]
        public Input<string>? DetectionStrategy { get; set; }

        [Input("diskDeviceMappings")]
        private InputList<Inputs.ImageDiskDeviceMappingGetArgs>? _diskDeviceMappings;

        /// <summary>
        /// Snapshot information for the image See `disk_device_mapping` below.
        /// </summary>
        public InputList<Inputs.ImageDiskDeviceMappingGetArgs> DiskDeviceMappings
        {
            get => _diskDeviceMappings ?? (_diskDeviceMappings = new InputList<Inputs.ImageDiskDeviceMappingGetArgs>());
            set => _diskDeviceMappings = value;
        }

        /// <summary>
        /// Features See `features` below.
        /// </summary>
        [Input("features")]
        public Input<Inputs.ImageFeaturesGetArgs>? Features { get; set; }

        /// <summary>
        /// Whether to perform forced deletion. Value range:
        /// - true: forcibly deletes the custom image, ignoring whether the current image is used by other instances.
        /// - false: The custom image is deleted normally. Before deleting the custom image, check whether the current image is used by other instances.
        /// 
        /// Default value: false
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The name of the image family. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with acs: or aliyun. It cannot contain http:// or https://. It can contain letters, digits, periods (.), colons (:), underscores (\_), and hyphens (-). By default, this parameter is empty.
        /// </summary>
        [Input("imageFamily")]
        public Input<string>? ImageFamily { get; set; }

        /// <summary>
        /// The name of the custom image. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with acs: or aliyun. It cannot contain http:// or https://. It can contain letters, digits, periods (.), colons (:), underscores (\_), and hyphens (-). By default, this parameter is empty. In this case, the original name is retained.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// The image version.
        /// 
        /// &gt; **NOTE:**  If you specify an instance by configuring `InstanceId`, and the instance uses an Alibaba Cloud Marketplace image or a custom image that is created from an Alibaba Cloud Marketplace image, you must leave this parameter empty or set this parameter to the value of ImageVersion of the instance.
        /// </summary>
        [Input("imageVersion")]
        public Input<string>? ImageVersion { get; set; }

        /// <summary>
        /// The instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The type of the license that is used to activate the operating system after the image is imported. Set the value to BYOL. BYOL: The license that comes with the source operating system is used. When you use the BYOL license, make sure that your license key is supported by Alibaba Cloud.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// . Field 'name' has been deprecated from provider version 1.227.0. New field 'image_name' instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The operating system distribution for the system disk in the custom image. If you specify a data disk snapshot to create the system disk of the custom image, use Platform to specify the operating system distribution for the system disk. Valid values: `Aliyun`, `Anolis`, `CentOS`, `Ubuntu`, `CoreOS`, `SUSE`, `Debian`, `OpenSUSE`, `FreeBSD`, `RedHat`, `Kylin`, `UOS`, `Fedora`, `Fedora CoreOS`, `CentOS Stream`, `AlmaLinux`, `Rocky Linux`, `Gentoo`, `Customized Linux`, `Others Linux`, `Windows Server 2022`, `Windows Server 2019`, `Windows Server 2016`, `Windows Server 2012`, `Windows Server 2008`, `Windows Server 2003`. Default value: `Others Linux`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The ID of the resource group to which to assign the custom image. If you do not specify this parameter, the image is assigned to the default resource group.
        /// 
        /// &gt; **NOTE:**   If you call the CreateImage operation as a Resource Access Management (RAM) user who does not have the permissions to manage the default resource group and do not specify `ResourceGroupId`, the `Forbbiden: User not authorized to operate on the specified resource` error message is returned. You must specify the ID of a resource group that the RAM user has the permissions to manage or grant the RAM user the permissions to manage the default resource group before you call the CreateImage operation again.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The ID of the snapshot that you want to use to create the custom image.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// The status of the image. By default, if you do not specify this parameter, only images in the Available state are returned.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ImageState()
        {
        }
        public static new ImageState Empty => new ImageState();
    }
}
