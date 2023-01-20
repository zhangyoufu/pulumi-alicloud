// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceMesh
{
    /// <summary>
    /// Provides a Service Mesh UserPermission resource.
    /// 
    /// For information about Service Mesh User Permission and how to use it, see [What is User Permission](https://help.aliyun.com/document_detail/171622.html).
    /// 
    /// &gt; **NOTE:** Available in v1.174.0+.
    /// 
    /// ## Import
    /// 
    /// Service Mesh User Permission can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:servicemesh/userPermission:UserPermission example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:servicemesh/userPermission:UserPermission")]
    public partial class UserPermission : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.UserPermissionPermission>> Permissions { get; private set; } = null!;

        /// <summary>
        /// The configuration of the Load Balancer. See the following `Block load_balancer`.
        /// </summary>
        [Output("subAccountUserId")]
        public Output<string> SubAccountUserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserPermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserPermission(string name, UserPermissionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:servicemesh/userPermission:UserPermission", name, args ?? new UserPermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserPermission(string name, Input<string> id, UserPermissionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:servicemesh/userPermission:UserPermission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserPermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserPermission Get(string name, Input<string> id, UserPermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new UserPermission(name, id, state, options);
        }
    }

    public sealed class UserPermissionArgs : global::Pulumi.ResourceArgs
    {
        [Input("permissions")]
        private InputList<Inputs.UserPermissionPermissionArgs>? _permissions;

        /// <summary>
        /// List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
        /// </summary>
        public InputList<Inputs.UserPermissionPermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.UserPermissionPermissionArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// The configuration of the Load Balancer. See the following `Block load_balancer`.
        /// </summary>
        [Input("subAccountUserId", required: true)]
        public Input<string> SubAccountUserId { get; set; } = null!;

        public UserPermissionArgs()
        {
        }
        public static new UserPermissionArgs Empty => new UserPermissionArgs();
    }

    public sealed class UserPermissionState : global::Pulumi.ResourceArgs
    {
        [Input("permissions")]
        private InputList<Inputs.UserPermissionPermissionGetArgs>? _permissions;

        /// <summary>
        /// List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
        /// </summary>
        public InputList<Inputs.UserPermissionPermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.UserPermissionPermissionGetArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// The configuration of the Load Balancer. See the following `Block load_balancer`.
        /// </summary>
        [Input("subAccountUserId")]
        public Input<string>? SubAccountUserId { get; set; }

        public UserPermissionState()
        {
        }
        public static new UserPermissionState Empty => new UserPermissionState();
    }
}
