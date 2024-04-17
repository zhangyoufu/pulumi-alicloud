// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.RocketMQ
{
    /// <summary>
    /// Provides an ONS group resource.
    /// 
    /// For more information about how to use it, see [RocketMQ Group Management API](https://www.alibabacloud.com/help/doc-detail/29616.html).
    /// 
    /// &gt; **NOTE:** Available in 1.53.0+
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
    ///     var name = config.Get("name") ?? "GID-tf-example";
    ///     var groupName = config.Get("groupName") ?? "GID-tf-example";
    ///     var @default = new Random.Index.Integer("default", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.RocketMQ.Instance("default", new()
    ///     {
    ///         Name = $"{name}-{@default.Result}",
    ///     });
    /// 
    ///     var defaultGroup = new AliCloud.RocketMQ.Group("default", new()
    ///     {
    ///         GroupName = groupName,
    ///         InstanceId = defaultInstance.Id,
    ///         Remark = "dafault_ons_group_remark",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ONS GROUP can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:rocketmq/group:Group group MQ_INST_1234567890_Baso1234567:GID-onsGroupDemo
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rocketmq/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Replaced by `group_name` after version 1.98.0.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// Name of the group. Two groups on a single instance cannot have the same name. A `group_name` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
        /// </summary>
        [Output("groupType")]
        public Output<string?> GroupType { get; private set; } = null!;

        /// <summary>
        /// ID of the ONS Instance that owns the groups.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
        /// </summary>
        [Output("readEnable")]
        public Output<bool?> ReadEnable { get; private set; } = null!;

        /// <summary>
        /// This attribute is a concise description of group. The length cannot exceed 256.
        /// </summary>
        [Output("remark")]
        public Output<string?> Remark { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:rocketmq/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rocketmq/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Replaced by `group_name` after version 1.98.0.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// Name of the group. Two groups on a single instance cannot have the same name. A `group_name` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
        /// </summary>
        [Input("groupType")]
        public Input<string>? GroupType { get; set; }

        /// <summary>
        /// ID of the ONS Instance that owns the groups.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
        /// </summary>
        [Input("readEnable")]
        public Input<bool>? ReadEnable { get; set; }

        /// <summary>
        /// This attribute is a concise description of group. The length cannot exceed 256.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Replaced by `group_name` after version 1.98.0.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// Name of the group. Two groups on a single instance cannot have the same name. A `group_name` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
        /// </summary>
        [Input("groupType")]
        public Input<string>? GroupType { get; set; }

        /// <summary>
        /// ID of the ONS Instance that owns the groups.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
        /// </summary>
        [Input("readEnable")]
        public Input<bool>? ReadEnable { get; set; }

        /// <summary>
        /// This attribute is a concise description of group. The length cannot exceed 256.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}
