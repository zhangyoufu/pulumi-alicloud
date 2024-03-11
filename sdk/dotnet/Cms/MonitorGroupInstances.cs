// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms
{
    /// <summary>
    /// Provides a Cloud Monitor Service Monitor Group Instances resource.
    /// 
    /// For information about Cloud Monitor Service Monitor Group Instances and how to use it, see [What is Monitor Group Instances](https://www.alibabacloud.com/help/en/cloudmonitor/latest/createmonitorgroupinstances).
    /// 
    /// &gt; **NOTE:** Available since v1.115.0.
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
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "192.168.0.0/16",
    ///     });
    /// 
    ///     var defaultMonitorGroup = new AliCloud.Cms.MonitorGroup("defaultMonitorGroup", new()
    ///     {
    ///         MonitorGroupName = name,
    ///     });
    /// 
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var example = new AliCloud.Cms.MonitorGroupInstances("example", new()
    ///     {
    ///         GroupId = defaultMonitorGroup.Id,
    ///         Instances = new[]
    ///         {
    ///             new AliCloud.Cms.Inputs.MonitorGroupInstancesInstanceArgs
    ///             {
    ///                 InstanceId = defaultNetwork.Id,
    ///                 InstanceName = name,
    ///                 RegionId = defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id),
    ///                 Category = "vpc",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Cloud Monitor Service Monitor Group Instances can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cms/monitorGroupInstances:MonitorGroupInstances example &lt;group_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cms/monitorGroupInstances:MonitorGroupInstances")]
    public partial class MonitorGroupInstances : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of Cms Group.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// Instance information added to the Cms Group. See `instances` below.
        /// </summary>
        [Output("instances")]
        public Output<ImmutableArray<Outputs.MonitorGroupInstancesInstance>> Instances { get; private set; } = null!;


        /// <summary>
        /// Create a MonitorGroupInstances resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MonitorGroupInstances(string name, MonitorGroupInstancesArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cms/monitorGroupInstances:MonitorGroupInstances", name, args ?? new MonitorGroupInstancesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MonitorGroupInstances(string name, Input<string> id, MonitorGroupInstancesState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cms/monitorGroupInstances:MonitorGroupInstances", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MonitorGroupInstances resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MonitorGroupInstances Get(string name, Input<string> id, MonitorGroupInstancesState? state = null, CustomResourceOptions? options = null)
        {
            return new MonitorGroupInstances(name, id, state, options);
        }
    }

    public sealed class MonitorGroupInstancesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of Cms Group.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("instances", required: true)]
        private InputList<Inputs.MonitorGroupInstancesInstanceArgs>? _instances;

        /// <summary>
        /// Instance information added to the Cms Group. See `instances` below.
        /// </summary>
        public InputList<Inputs.MonitorGroupInstancesInstanceArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.MonitorGroupInstancesInstanceArgs>());
            set => _instances = value;
        }

        public MonitorGroupInstancesArgs()
        {
        }
        public static new MonitorGroupInstancesArgs Empty => new MonitorGroupInstancesArgs();
    }

    public sealed class MonitorGroupInstancesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of Cms Group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("instances")]
        private InputList<Inputs.MonitorGroupInstancesInstanceGetArgs>? _instances;

        /// <summary>
        /// Instance information added to the Cms Group. See `instances` below.
        /// </summary>
        public InputList<Inputs.MonitorGroupInstancesInstanceGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.MonitorGroupInstancesInstanceGetArgs>());
            set => _instances = value;
        }

        public MonitorGroupInstancesState()
        {
        }
        public static new MonitorGroupInstancesState Empty => new MonitorGroupInstancesState();
    }
}
