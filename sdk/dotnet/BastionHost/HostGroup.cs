// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.BastionHost
{
    /// <summary>
    /// Provides a Bastion Host Host Group resource.
    /// 
    /// For information about Bastion Host Host Group and how to use it, see [What is Host Group](https://www.alibabacloud.com/help/en/doc-detail/204307.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.134.0+.
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
    ///         var example = new AliCloud.BastionHost.HostGroup("example", new AliCloud.BastionHost.HostGroupArgs
    ///         {
    ///             HostGroupName = "example_value",
    ///             InstanceId = "bastionhost-cn-tl3xxxxxxx",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Bastion Host Host Group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:bastionhost/hostGroup:HostGroup example &lt;instance_id&gt;:&lt;host_group_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:bastionhost/hostGroup:HostGroup")]
    public partial class HostGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Specify the New Host Group of Notes, Supports up to 500 Characters.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Host Group ID.
        /// </summary>
        [Output("hostGroupId")]
        public Output<string> HostGroupId { get; private set; } = null!;

        /// <summary>
        /// Specify the New Host Group Name, Supports up to 128 Characters.
        /// </summary>
        [Output("hostGroupName")]
        public Output<string> HostGroupName { get; private set; } = null!;

        /// <summary>
        /// Specify the New Host Group Where the Bastion Host ID of.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a HostGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostGroup(string name, HostGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:bastionhost/hostGroup:HostGroup", name, args ?? new HostGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostGroup(string name, Input<string> id, HostGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:bastionhost/hostGroup:HostGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HostGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostGroup Get(string name, Input<string> id, HostGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new HostGroup(name, id, state, options);
        }
    }

    public sealed class HostGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify the New Host Group of Notes, Supports up to 500 Characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Specify the New Host Group Name, Supports up to 128 Characters.
        /// </summary>
        [Input("hostGroupName", required: true)]
        public Input<string> HostGroupName { get; set; } = null!;

        /// <summary>
        /// Specify the New Host Group Where the Bastion Host ID of.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public HostGroupArgs()
        {
        }
    }

    public sealed class HostGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify the New Host Group of Notes, Supports up to 500 Characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Host Group ID.
        /// </summary>
        [Input("hostGroupId")]
        public Input<string>? HostGroupId { get; set; }

        /// <summary>
        /// Specify the New Host Group Name, Supports up to 128 Characters.
        /// </summary>
        [Input("hostGroupName")]
        public Input<string>? HostGroupName { get; set; }

        /// <summary>
        /// Specify the New Host Group Where the Bastion Host ID of.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public HostGroupState()
        {
        }
    }
}
