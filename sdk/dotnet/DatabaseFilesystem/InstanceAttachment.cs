// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.DatabaseFilesystem
{
    /// <summary>
    /// Provides a DBFS Instance Attachment resource.
    /// 
    /// For information about DBFS Instance Attachment and how to use it, see [What is Instance Attachment](https://help.aliyun.com/document_detail/149726.html).
    /// 
    /// &gt; **NOTE:** Available in v1.156.0+.
    /// 
    /// ## Example Usage
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
    ///     var defaultNetworks = AliCloud.Vpc.GetNetworks.Invoke(new()
    ///     {
    ///         NameRegex = "default-NODELETING",
    ///     });
    /// 
    ///     var zoneId = "cn-hangzhou-i";
    /// 
    ///     var defaultSwitches = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = zoneId,
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new()
    ///     {
    ///         Description = "tf test",
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///     });
    /// 
    ///     var defaultImages = AliCloud.Ecs.GetImages.Invoke(new()
    ///     {
    ///         Owners = "system",
    ///         NameRegex = "^centos_8",
    ///         MostRecent = true,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Ecs.Instance("defaultInstance", new()
    ///     {
    ///         ImageId = defaultImages.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///         InstanceName = @var.Name,
    ///         InstanceType = "ecs.g7se.large",
    ///         AvailabilityZone = zoneId,
    ///         VswitchId = defaultSwitches.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]),
    ///         SystemDiskCategory = "cloud_essd",
    ///         SecurityGroups = new[]
    ///         {
    ///             defaultSecurityGroup.Id,
    ///         },
    ///     });
    /// 
    ///     var defaultDatabasefilesystem_instanceInstance = new AliCloud.DatabaseFilesystem.Instance("defaultDatabasefilesystem/instanceInstance", new()
    ///     {
    ///         Category = "standard",
    ///         ZoneId = defaultInstance.AvailabilityZone,
    ///         PerformanceLevel = "PL1",
    ///         InstanceName = @var.Name,
    ///         Size = 100,
    ///     });
    /// 
    ///     var example = new AliCloud.DatabaseFilesystem.InstanceAttachment("example", new()
    ///     {
    ///         EcsId = defaultInstance.Id,
    ///         InstanceId = defaultDatabasefilesystem / instanceInstance.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DBFS Instance Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:databasefilesystem/instanceAttachment:InstanceAttachment example &lt;instance_id&gt;:&lt;ecs_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:databasefilesystem/instanceAttachment:InstanceAttachment")]
    public partial class InstanceAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the ECS instance.
        /// </summary>
        [Output("ecsId")]
        public Output<string> EcsId { get; private set; } = null!;

        /// <summary>
        /// The ID of the database file system.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The status of Database file system. Valid values: `attached`, `attaching`, `unattached`, `detaching`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceAttachment(string name, InstanceAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:databasefilesystem/instanceAttachment:InstanceAttachment", name, args ?? new InstanceAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceAttachment(string name, Input<string> id, InstanceAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:databasefilesystem/instanceAttachment:InstanceAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceAttachment Get(string name, Input<string> id, InstanceAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceAttachment(name, id, state, options);
        }
    }

    public sealed class InstanceAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the ECS instance.
        /// </summary>
        [Input("ecsId", required: true)]
        public Input<string> EcsId { get; set; } = null!;

        /// <summary>
        /// The ID of the database file system.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public InstanceAttachmentArgs()
        {
        }
        public static new InstanceAttachmentArgs Empty => new InstanceAttachmentArgs();
    }

    public sealed class InstanceAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the ECS instance.
        /// </summary>
        [Input("ecsId")]
        public Input<string>? EcsId { get; set; }

        /// <summary>
        /// The ID of the database file system.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The status of Database file system. Valid values: `attached`, `attaching`, `unattached`, `detaching`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public InstanceAttachmentState()
        {
        }
        public static new InstanceAttachmentState Empty => new InstanceAttachmentState();
    }
}
