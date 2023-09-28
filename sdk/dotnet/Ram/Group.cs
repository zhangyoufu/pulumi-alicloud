// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ram
{
    /// <summary>
    /// Provides a RAM Group resource.
    /// 
    /// &gt; **NOTE:** When you want to destroy this resource forcefully(means remove all the relationships associated with it automatically and then destroy it) without set `force`  with `true` at beginning, you need add `force = true` to configuration file and run `pulumi preview`, then you can delete resource forcefully.
    /// 
    /// &gt; **NOTE:** Available since v1.0.0+.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a new RAM Group.
    ///     var @group = new AliCloud.Ram.Group("group", new()
    ///     {
    ///         Comments = "this is a group comments.",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RAM group can be imported using the id or name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ram/group:Group example my-group
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ram/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment of the RAM group. This parameter can have a string of 1 to 128 characters.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// This parameter is used for resource destroy. Default value is `false`.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// Name of the RAM group. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:ram/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ram/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment of the RAM group. This parameter can have a string of 1 to 128 characters.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// This parameter is used for resource destroy. Default value is `false`.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// Name of the RAM group. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment of the RAM group. This parameter can have a string of 1 to 128 characters.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// This parameter is used for resource destroy. Default value is `false`.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// Name of the RAM group. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}
