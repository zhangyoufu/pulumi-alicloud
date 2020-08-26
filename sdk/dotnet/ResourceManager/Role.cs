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
    /// Provides a Resource Manager role resource. Members are resource containers in the resource directory, which can physically isolate resources to form an independent resource grouping unit. You can create members in the resource folder to manage them in a unified manner.
    /// For information about Resource Manager role and how to use it, see [What is Resource Manager role](https://www.alibabacloud.com/help/en/doc-detail/111231.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.82.0+.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Add a Resource Manager role.
    ///         var example = new AliCloud.ResourceManager.Role("example", new AliCloud.ResourceManager.RoleArgs
    ///         {
    ///             AssumeRolePolicyDocument = @"     {
    ///           ""Statement"": [
    ///                {
    ///                     ""Action"": ""sts:AssumeRole"",
    ///                     ""Effect"": ""Allow"",
    ///                     ""Principal"": {
    ///                         ""RAM"":[
    ///                                 ""acs:ram::103755469187****:root""，
    ///                                 ""acs:ram::104408977069****:root""
    ///                         ]
    ///                     }
    ///                 }
    ///           ],
    ///           ""Version"": ""1""
    ///      }
    /// 	 
    /// ",
    ///             RoleName = "testrd",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Role : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource descriptor of the role.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The content of the permissions strategy that plays a role.
        /// </summary>
        [Output("assumeRolePolicyDocument")]
        public Output<string> AssumeRolePolicyDocument { get; private set; } = null!;

        /// <summary>
        /// Role creation time.
        /// </summary>
        [Output("createDate")]
        public Output<string> CreateDate { get; private set; } = null!;

        /// <summary>
        /// The description of the Resource Manager role.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
        /// </summary>
        [Output("maxSessionDuration")]
        public Output<int?> MaxSessionDuration { get; private set; } = null!;

        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;

        /// <summary>
        /// Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots "." and dashes "-".
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

        /// <summary>
        /// Role update time.
        /// </summary>
        [Output("updateDate")]
        public Output<string> UpdateDate { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/role:Role", name, args ?? new RoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/role:Role", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
        {
            return new Role(name, id, state, options);
        }
    }

    public sealed class RoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content of the permissions strategy that plays a role.
        /// </summary>
        [Input("assumeRolePolicyDocument", required: true)]
        public Input<string> AssumeRolePolicyDocument { get; set; } = null!;

        /// <summary>
        /// The description of the Resource Manager role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
        /// </summary>
        [Input("maxSessionDuration")]
        public Input<int>? MaxSessionDuration { get; set; }

        /// <summary>
        /// Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots "." and dashes "-".
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        public RoleArgs()
        {
        }
    }

    public sealed class RoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource descriptor of the role.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The content of the permissions strategy that plays a role.
        /// </summary>
        [Input("assumeRolePolicyDocument")]
        public Input<string>? AssumeRolePolicyDocument { get; set; }

        /// <summary>
        /// Role creation time.
        /// </summary>
        [Input("createDate")]
        public Input<string>? CreateDate { get; set; }

        /// <summary>
        /// The description of the Resource Manager role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
        /// </summary>
        [Input("maxSessionDuration")]
        public Input<int>? MaxSessionDuration { get; set; }

        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        /// <summary>
        /// Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots "." and dashes "-".
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// Role update time.
        /// </summary>
        [Input("updateDate")]
        public Input<string>? UpdateDate { get; set; }

        public RoleState()
        {
        }
    }
}
