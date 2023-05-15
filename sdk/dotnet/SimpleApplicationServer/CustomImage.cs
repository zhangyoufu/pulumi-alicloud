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
    /// Provides a Simple Application Server Custom Image resource.
    /// 
    /// For information about Simple Application Server Custom Image and how to use it, see [What is Custom Image](https://www.alibabacloud.com/help/zh/doc-detail/333535.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.143.0+.
    /// 
    /// ## Import
    /// 
    /// Simple Application Server Custom Image can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:simpleapplicationserver/customImage:CustomImage example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:simpleapplicationserver/customImage:CustomImage")]
    public partial class CustomImage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource. The name must be `2` to `128` characters in length. It must start with a letter or a number. It can contain letters, digits, colons (:), underscores (_) and hyphens (-).
        /// </summary>
        [Output("customImageName")]
        public Output<string> CustomImageName { get; private set; } = null!;

        /// <summary>
        /// The description of the Custom Image.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The Shared status of the Custom Image. Valid values: `Share`, `UnShare`.
        /// 
        /// **NOTE:** The `status` will be automatically change to `UnShare` when the resource is deleted, please operate with caution.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the system snapshot.
        /// </summary>
        [Output("systemSnapshotId")]
        public Output<string> SystemSnapshotId { get; private set; } = null!;


        /// <summary>
        /// Create a CustomImage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomImage(string name, CustomImageArgs args, CustomResourceOptions? options = null)
            : base("alicloud:simpleapplicationserver/customImage:CustomImage", name, args ?? new CustomImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomImage(string name, Input<string> id, CustomImageState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:simpleapplicationserver/customImage:CustomImage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomImage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomImage Get(string name, Input<string> id, CustomImageState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomImage(name, id, state, options);
        }
    }

    public sealed class CustomImageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the resource. The name must be `2` to `128` characters in length. It must start with a letter or a number. It can contain letters, digits, colons (:), underscores (_) and hyphens (-).
        /// </summary>
        [Input("customImageName", required: true)]
        public Input<string> CustomImageName { get; set; } = null!;

        /// <summary>
        /// The description of the Custom Image.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The Shared status of the Custom Image. Valid values: `Share`, `UnShare`.
        /// 
        /// **NOTE:** The `status` will be automatically change to `UnShare` when the resource is deleted, please operate with caution.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the system snapshot.
        /// </summary>
        [Input("systemSnapshotId", required: true)]
        public Input<string> SystemSnapshotId { get; set; } = null!;

        public CustomImageArgs()
        {
        }
        public static new CustomImageArgs Empty => new CustomImageArgs();
    }

    public sealed class CustomImageState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the resource. The name must be `2` to `128` characters in length. It must start with a letter or a number. It can contain letters, digits, colons (:), underscores (_) and hyphens (-).
        /// </summary>
        [Input("customImageName")]
        public Input<string>? CustomImageName { get; set; }

        /// <summary>
        /// The description of the Custom Image.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The Shared status of the Custom Image. Valid values: `Share`, `UnShare`.
        /// 
        /// **NOTE:** The `status` will be automatically change to `UnShare` when the resource is deleted, please operate with caution.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the system snapshot.
        /// </summary>
        [Input("systemSnapshotId")]
        public Input<string>? SystemSnapshotId { get; set; }

        public CustomImageState()
        {
        }
        public static new CustomImageState Empty => new CustomImageState();
    }
}
