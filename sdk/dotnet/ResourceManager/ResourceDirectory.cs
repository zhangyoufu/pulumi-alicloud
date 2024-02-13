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
    /// Provides a Resource Manager Resource Directory resource. Resource Directory enables you to establish an organizational structure for the resources used by applications of your enterprise. You can plan, build, and manage the resources in a centralized manner by using only one resource directory.
    /// 
    /// For information about Resource Manager Resource Directory and how to use it, see [What is Resource Manager Resource Directory](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.84.0.
    /// 
    /// &gt; **NOTE:** An account can only be used to enable a resource directory after it passes enterprise real-name verification. An account that only passed individual real-name verification cannot be used to enable a resource directory.
    /// 
    /// &gt; **NOTE:** Before you destroy the resource, make sure that the following requirements are met:
    ///   - All member accounts must be removed from the resource directory.
    ///   - All folders except the root folder must be deleted from the resource directory.
    /// 
    /// ## Import
    /// 
    /// Resource Manager Resource Directory can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:resourcemanager/resourceDirectory:ResourceDirectory example rd-s3****
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:resourcemanager/resourceDirectory:ResourceDirectory")]
    public partial class ResourceDirectory : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the master account.
        /// </summary>
        [Output("masterAccountId")]
        public Output<string> MasterAccountId { get; private set; } = null!;

        /// <summary>
        /// The name of the master account.
        /// </summary>
        [Output("masterAccountName")]
        public Output<string> MasterAccountName { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable the member deletion feature. Valid values:`Enabled` and `Disabled`.
        /// </summary>
        [Output("memberDeletionStatus")]
        public Output<string> MemberDeletionStatus { get; private set; } = null!;

        /// <summary>
        /// The ID of the root folder.
        /// </summary>
        [Output("rootFolderId")]
        public Output<string> RootFolderId { get; private set; } = null!;

        /// <summary>
        /// The status of control policy. Valid values:`Enabled` and `Disabled`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceDirectory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceDirectory(string name, ResourceDirectoryArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/resourceDirectory:ResourceDirectory", name, args ?? new ResourceDirectoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceDirectory(string name, Input<string> id, ResourceDirectoryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/resourceDirectory:ResourceDirectory", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceDirectory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceDirectory Get(string name, Input<string> id, ResourceDirectoryState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceDirectory(name, id, state, options);
        }
    }

    public sealed class ResourceDirectoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable the member deletion feature. Valid values:`Enabled` and `Disabled`.
        /// </summary>
        [Input("memberDeletionStatus")]
        public Input<string>? MemberDeletionStatus { get; set; }

        /// <summary>
        /// The status of control policy. Valid values:`Enabled` and `Disabled`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ResourceDirectoryArgs()
        {
        }
        public static new ResourceDirectoryArgs Empty => new ResourceDirectoryArgs();
    }

    public sealed class ResourceDirectoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the master account.
        /// </summary>
        [Input("masterAccountId")]
        public Input<string>? MasterAccountId { get; set; }

        /// <summary>
        /// The name of the master account.
        /// </summary>
        [Input("masterAccountName")]
        public Input<string>? MasterAccountName { get; set; }

        /// <summary>
        /// Specifies whether to enable the member deletion feature. Valid values:`Enabled` and `Disabled`.
        /// </summary>
        [Input("memberDeletionStatus")]
        public Input<string>? MemberDeletionStatus { get; set; }

        /// <summary>
        /// The ID of the root folder.
        /// </summary>
        [Input("rootFolderId")]
        public Input<string>? RootFolderId { get; set; }

        /// <summary>
        /// The status of control policy. Valid values:`Enabled` and `Disabled`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ResourceDirectoryState()
        {
        }
        public static new ResourceDirectoryState Empty => new ResourceDirectoryState();
    }
}
