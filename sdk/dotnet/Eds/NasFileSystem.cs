// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eds
{
    /// <summary>
    /// Provides a ECD Nas File System resource.
    /// 
    /// For information about ECD Nas File System and how to use it, see [What is Nas File System](https://www.alibabacloud.com/help/en/elastic-desktop-service/latest/api-reference-for-easy-use-1).
    /// 
    /// &gt; **NOTE:** Available since v1.141.0.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultRandomInteger = new Random.RandomInteger("defaultRandomInteger", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var defaultSimpleOfficeSite = new AliCloud.Eds.SimpleOfficeSite("defaultSimpleOfficeSite", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/12",
    ///         EnableAdminAccess = false,
    ///         DesktopAccessType = "Internet",
    ///         OfficeSiteName = defaultRandomInteger.Result.Apply(result =&gt; $"{name}-{result}"),
    ///     });
    /// 
    ///     var example = new AliCloud.Eds.NasFileSystem("example", new()
    ///     {
    ///         NasFileSystemName = name,
    ///         OfficeSiteId = defaultSimpleOfficeSite.Id,
    ///         Description = name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECD Nas File System can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:eds/nasFileSystem:NasFileSystem example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:eds/nasFileSystem:NasFileSystem")]
    public partial class NasFileSystem : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of nas file system.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The filesystem id of nas file system.
        /// </summary>
        [Output("fileSystemId")]
        public Output<string> FileSystemId { get; private set; } = null!;

        /// <summary>
        /// The domain of mount target.
        /// </summary>
        [Output("mountTargetDomain")]
        public Output<string> MountTargetDomain { get; private set; } = null!;

        /// <summary>
        /// The name of nas file system.
        /// </summary>
        [Output("nasFileSystemName")]
        public Output<string?> NasFileSystemName { get; private set; } = null!;

        /// <summary>
        /// The ID of office site.
        /// </summary>
        [Output("officeSiteId")]
        public Output<string> OfficeSiteId { get; private set; } = null!;

        /// <summary>
        /// The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
        /// </summary>
        [Output("reset")]
        public Output<bool?> Reset { get; private set; } = null!;

        /// <summary>
        /// The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a NasFileSystem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NasFileSystem(string name, NasFileSystemArgs args, CustomResourceOptions? options = null)
            : base("alicloud:eds/nasFileSystem:NasFileSystem", name, args ?? new NasFileSystemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NasFileSystem(string name, Input<string> id, NasFileSystemState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:eds/nasFileSystem:NasFileSystem", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NasFileSystem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NasFileSystem Get(string name, Input<string> id, NasFileSystemState? state = null, CustomResourceOptions? options = null)
        {
            return new NasFileSystem(name, id, state, options);
        }
    }

    public sealed class NasFileSystemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of nas file system.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The filesystem id of nas file system.
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        /// <summary>
        /// The domain of mount target.
        /// </summary>
        [Input("mountTargetDomain")]
        public Input<string>? MountTargetDomain { get; set; }

        /// <summary>
        /// The name of nas file system.
        /// </summary>
        [Input("nasFileSystemName")]
        public Input<string>? NasFileSystemName { get; set; }

        /// <summary>
        /// The ID of office site.
        /// </summary>
        [Input("officeSiteId", required: true)]
        public Input<string> OfficeSiteId { get; set; } = null!;

        /// <summary>
        /// The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
        /// </summary>
        [Input("reset")]
        public Input<bool>? Reset { get; set; }

        public NasFileSystemArgs()
        {
        }
        public static new NasFileSystemArgs Empty => new NasFileSystemArgs();
    }

    public sealed class NasFileSystemState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of nas file system.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The filesystem id of nas file system.
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        /// <summary>
        /// The domain of mount target.
        /// </summary>
        [Input("mountTargetDomain")]
        public Input<string>? MountTargetDomain { get; set; }

        /// <summary>
        /// The name of nas file system.
        /// </summary>
        [Input("nasFileSystemName")]
        public Input<string>? NasFileSystemName { get; set; }

        /// <summary>
        /// The ID of office site.
        /// </summary>
        [Input("officeSiteId")]
        public Input<string>? OfficeSiteId { get; set; }

        /// <summary>
        /// The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
        /// </summary>
        [Input("reset")]
        public Input<bool>? Reset { get; set; }

        /// <summary>
        /// The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public NasFileSystemState()
        {
        }
        public static new NasFileSystemState Empty => new NasFileSystemState();
    }
}
