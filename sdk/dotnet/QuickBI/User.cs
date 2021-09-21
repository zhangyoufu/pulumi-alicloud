// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.QuickBI
{
    /// <summary>
    /// Provides a Quick BI User resource.
    /// 
    /// For information about Quick BI User and how to use it, see [What is User](https://www.alibabacloud.com/help/doc-detail/33813.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.136.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new AliCloud.QuickBI.User("example", new AliCloud.QuickBI.UserArgs
    ///         {
    ///             AccountName = "example_value",
    ///             AdminUser = false,
    ///             AuthAdminUser = false,
    ///             NickName = "example_value",
    ///             UserType = "Analyst",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Quick BI User can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:quickbi/user:User example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:quickbi/user:User")]
    public partial class User : Pulumi.CustomResource
    {
        /// <summary>
        /// Alibaba Cloud account ID.
        /// </summary>
        [Output("accountId")]
        public Output<string?> AccountId { get; private set; } = null!;

        /// <summary>
        /// An Alibaba Cloud account, Alibaba Cloud name.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// Whether it is the administrator. Valid values: `true` and `false`.
        /// </summary>
        [Output("adminUser")]
        public Output<bool> AdminUser { get; private set; } = null!;

        /// <summary>
        /// Whether this is a permissions administrator. Valid values: `false`, `true`.
        /// </summary>
        [Output("authAdminUser")]
        public Output<bool> AuthAdminUser { get; private set; } = null!;

        /// <summary>
        /// The nickname of the user.
        /// </summary>
        [Output("nickName")]
        public Output<string> NickName { get; private set; } = null!;

        /// <summary>
        /// The members of the organization of the type of role separately. Valid values: `Analyst`, `Developer` and `Visitor`.
        /// </summary>
        [Output("userType")]
        public Output<string> UserType { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("alicloud:quickbi/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:quickbi/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alibaba Cloud account ID.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// An Alibaba Cloud account, Alibaba Cloud name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Whether it is the administrator. Valid values: `true` and `false`.
        /// </summary>
        [Input("adminUser", required: true)]
        public Input<bool> AdminUser { get; set; } = null!;

        /// <summary>
        /// Whether this is a permissions administrator. Valid values: `false`, `true`.
        /// </summary>
        [Input("authAdminUser", required: true)]
        public Input<bool> AuthAdminUser { get; set; } = null!;

        /// <summary>
        /// The nickname of the user.
        /// </summary>
        [Input("nickName", required: true)]
        public Input<string> NickName { get; set; } = null!;

        /// <summary>
        /// The members of the organization of the type of role separately. Valid values: `Analyst`, `Developer` and `Visitor`.
        /// </summary>
        [Input("userType", required: true)]
        public Input<string> UserType { get; set; } = null!;

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alibaba Cloud account ID.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// An Alibaba Cloud account, Alibaba Cloud name.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Whether it is the administrator. Valid values: `true` and `false`.
        /// </summary>
        [Input("adminUser")]
        public Input<bool>? AdminUser { get; set; }

        /// <summary>
        /// Whether this is a permissions administrator. Valid values: `false`, `true`.
        /// </summary>
        [Input("authAdminUser")]
        public Input<bool>? AuthAdminUser { get; set; }

        /// <summary>
        /// The nickname of the user.
        /// </summary>
        [Input("nickName")]
        public Input<string>? NickName { get; set; }

        /// <summary>
        /// The members of the organization of the type of role separately. Valid values: `Analyst`, `Developer` and `Visitor`.
        /// </summary>
        [Input("userType")]
        public Input<string>? UserType { get; set; }

        public UserState()
        {
        }
    }
}
