// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ResourceManager
{
    /// <summary>
    /// Provides a Resource Manager Folder resource. A folder is an organizational unit in a resource directory. You can use folders to build an organizational structure for resources.
    /// For information about Resource Manager Foler and how to use it, see [What is Resource Manager Folder](https://www.alibabacloud.com/help/en/doc-detail/111221.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.82.0.
    /// 
    /// &gt; **NOTE:** A maximum of five levels of folders can be created under the root folder.
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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var example = new AliCloud.ResourceManager.Folder("example", new()
    ///     {
    ///         FolderName = @default.Result.Apply(result =&gt; $"{name}-{result}"),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Resource Manager Folder can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:resourcemanager/folder:Folder example fd-u8B321****	
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:resourcemanager/folder:Folder")]
    public partial class Folder : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
        /// </summary>
        [Output("folderName")]
        public Output<string> FolderName { get; private set; } = null!;

        /// <summary>
        /// The ID of the parent folder. If not set, the system default value will be used.
        /// </summary>
        [Output("parentFolderId")]
        public Output<string> ParentFolderId { get; private set; } = null!;


        /// <summary>
        /// Create a Folder resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Folder(string name, FolderArgs args, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/folder:Folder", name, args ?? new FolderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Folder(string name, Input<string> id, FolderState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/folder:Folder", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Folder resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Folder Get(string name, Input<string> id, FolderState? state = null, CustomResourceOptions? options = null)
        {
            return new Folder(name, id, state, options);
        }
    }

    public sealed class FolderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
        /// </summary>
        [Input("folderName", required: true)]
        public Input<string> FolderName { get; set; } = null!;

        /// <summary>
        /// The ID of the parent folder. If not set, the system default value will be used.
        /// </summary>
        [Input("parentFolderId")]
        public Input<string>? ParentFolderId { get; set; }

        public FolderArgs()
        {
        }
        public static new FolderArgs Empty => new FolderArgs();
    }

    public sealed class FolderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
        /// </summary>
        [Input("folderName")]
        public Input<string>? FolderName { get; set; }

        /// <summary>
        /// The ID of the parent folder. If not set, the system default value will be used.
        /// </summary>
        [Input("parentFolderId")]
        public Input<string>? ParentFolderId { get; set; }

        public FolderState()
        {
        }
        public static new FolderState Empty => new FolderState();
    }
}
