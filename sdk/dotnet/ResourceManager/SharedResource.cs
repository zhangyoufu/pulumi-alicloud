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
    /// Provides a Resource Manager Shared Resource resource.
    /// 
    /// For information about Resource Manager Shared Resource and how to use it, see [What is Shared Resource](https://www.alibabacloud.com/help/en/resource-management/latest/api-resourcesharing-2020-01-10-associateresourceshare).
    /// 
    /// &gt; **NOTE:** Available since v1.111.0.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tfexample";
    ///     var exampleZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var exampleNetwork = new AliCloud.Vpc.Network("exampleNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "192.168.0.0/16",
    ///     });
    /// 
    ///     var exampleSwitch = new AliCloud.Vpc.Switch("exampleSwitch", new()
    ///     {
    ///         ZoneId = exampleZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         CidrBlock = "192.168.0.0/16",
    ///         VpcId = exampleNetwork.Id,
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var exampleResourceShare = new AliCloud.ResourceManager.ResourceShare("exampleResourceShare", new()
    ///     {
    ///         ResourceShareName = name,
    ///     });
    /// 
    ///     var exampleSharedResource = new AliCloud.ResourceManager.SharedResource("exampleSharedResource", new()
    ///     {
    ///         ResourceId = exampleSwitch.Id,
    ///         ResourceShareId = exampleResourceShare.Id,
    ///         ResourceType = "VSwitch",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Resource Manager Shared Resource can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:resourcemanager/sharedResource:SharedResource example &lt;resource_share_id&gt;:&lt;resource_id&gt;:&lt;resource_type&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:resourcemanager/sharedResource:SharedResource")]
    public partial class SharedResource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID need shared.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The resource share ID of resource manager.
        /// </summary>
        [Output("resourceShareId")]
        public Output<string> ResourceShareId { get; private set; } = null!;

        /// <summary>
        /// The resource type of should shared. Valid values:
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        /// <summary>
        /// The status of the Shared Resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a SharedResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SharedResource(string name, SharedResourceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/sharedResource:SharedResource", name, args ?? new SharedResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SharedResource(string name, Input<string> id, SharedResourceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/sharedResource:SharedResource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SharedResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SharedResource Get(string name, Input<string> id, SharedResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new SharedResource(name, id, state, options);
        }
    }

    public sealed class SharedResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID need shared.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The resource share ID of resource manager.
        /// </summary>
        [Input("resourceShareId", required: true)]
        public Input<string> ResourceShareId { get; set; } = null!;

        /// <summary>
        /// The resource type of should shared. Valid values:
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        public SharedResourceArgs()
        {
        }
        public static new SharedResourceArgs Empty => new SharedResourceArgs();
    }

    public sealed class SharedResourceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID need shared.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The resource share ID of resource manager.
        /// </summary>
        [Input("resourceShareId")]
        public Input<string>? ResourceShareId { get; set; }

        /// <summary>
        /// The resource type of should shared. Valid values:
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// The status of the Shared Resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SharedResourceState()
        {
        }
        public static new SharedResourceState Empty => new SharedResourceState();
    }
}
