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
    /// Provides a Cloud Monitor Service Hybrid Monitor Sls Task resource.
    /// 
    /// For information about Cloud Monitor Service Hybrid Monitor Sls Task and how to use it, see [What is Hybrid Monitor Sls Task](https://www.alibabacloud.com/help/en/cloudmonitor/latest/createhybridmonitortask).
    /// 
    /// &gt; **NOTE:** Available since v1.179.0.
    /// 
    /// ## Import
    /// 
    /// Cloud Monitor Service Hybrid Monitor Sls Task can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cms/hybridMonitorSlsTask:HybridMonitorSlsTask example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cms/hybridMonitorSlsTask:HybridMonitorSlsTask")]
    public partial class HybridMonitorSlsTask : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The label of the monitoring task. See `attach_labels` below.
        /// </summary>
        [Output("attachLabels")]
        public Output<ImmutableArray<Outputs.HybridMonitorSlsTaskAttachLabel>> AttachLabels { get; private set; } = null!;

        /// <summary>
        /// The interval at which metrics are collected. Valid values: `15`, `60`(default value). Unit: seconds.
        /// </summary>
        [Output("collectInterval")]
        public Output<int> CollectInterval { get; private set; } = null!;

        /// <summary>
        /// The type of the collection target, enter the name of the Logstore group.
        /// </summary>
        [Output("collectTargetType")]
        public Output<string> CollectTargetType { get; private set; } = null!;

        /// <summary>
        /// The description of the metric import task.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;

        /// <summary>
        /// The configurations of the logs that are imported from Log Service. See `sls_process_config` below.
        /// </summary>
        [Output("slsProcessConfig")]
        public Output<Outputs.HybridMonitorSlsTaskSlsProcessConfig> SlsProcessConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the metric import task, enter the name of the metric for logs imported from Log Service.
        /// </summary>
        [Output("taskName")]
        public Output<string> TaskName { get; private set; } = null!;


        /// <summary>
        /// Create a HybridMonitorSlsTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HybridMonitorSlsTask(string name, HybridMonitorSlsTaskArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cms/hybridMonitorSlsTask:HybridMonitorSlsTask", name, args ?? new HybridMonitorSlsTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HybridMonitorSlsTask(string name, Input<string> id, HybridMonitorSlsTaskState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cms/hybridMonitorSlsTask:HybridMonitorSlsTask", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HybridMonitorSlsTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HybridMonitorSlsTask Get(string name, Input<string> id, HybridMonitorSlsTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new HybridMonitorSlsTask(name, id, state, options);
        }
    }

    public sealed class HybridMonitorSlsTaskArgs : global::Pulumi.ResourceArgs
    {
        [Input("attachLabels")]
        private InputList<Inputs.HybridMonitorSlsTaskAttachLabelArgs>? _attachLabels;

        /// <summary>
        /// The label of the monitoring task. See `attach_labels` below.
        /// </summary>
        public InputList<Inputs.HybridMonitorSlsTaskAttachLabelArgs> AttachLabels
        {
            get => _attachLabels ?? (_attachLabels = new InputList<Inputs.HybridMonitorSlsTaskAttachLabelArgs>());
            set => _attachLabels = value;
        }

        /// <summary>
        /// The interval at which metrics are collected. Valid values: `15`, `60`(default value). Unit: seconds.
        /// </summary>
        [Input("collectInterval")]
        public Input<int>? CollectInterval { get; set; }

        /// <summary>
        /// The type of the collection target, enter the name of the Logstore group.
        /// </summary>
        [Input("collectTargetType", required: true)]
        public Input<string> CollectTargetType { get; set; } = null!;

        /// <summary>
        /// The description of the metric import task.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// The configurations of the logs that are imported from Log Service. See `sls_process_config` below.
        /// </summary>
        [Input("slsProcessConfig", required: true)]
        public Input<Inputs.HybridMonitorSlsTaskSlsProcessConfigArgs> SlsProcessConfig { get; set; } = null!;

        /// <summary>
        /// The name of the metric import task, enter the name of the metric for logs imported from Log Service.
        /// </summary>
        [Input("taskName", required: true)]
        public Input<string> TaskName { get; set; } = null!;

        public HybridMonitorSlsTaskArgs()
        {
        }
        public static new HybridMonitorSlsTaskArgs Empty => new HybridMonitorSlsTaskArgs();
    }

    public sealed class HybridMonitorSlsTaskState : global::Pulumi.ResourceArgs
    {
        [Input("attachLabels")]
        private InputList<Inputs.HybridMonitorSlsTaskAttachLabelGetArgs>? _attachLabels;

        /// <summary>
        /// The label of the monitoring task. See `attach_labels` below.
        /// </summary>
        public InputList<Inputs.HybridMonitorSlsTaskAttachLabelGetArgs> AttachLabels
        {
            get => _attachLabels ?? (_attachLabels = new InputList<Inputs.HybridMonitorSlsTaskAttachLabelGetArgs>());
            set => _attachLabels = value;
        }

        /// <summary>
        /// The interval at which metrics are collected. Valid values: `15`, `60`(default value). Unit: seconds.
        /// </summary>
        [Input("collectInterval")]
        public Input<int>? CollectInterval { get; set; }

        /// <summary>
        /// The type of the collection target, enter the name of the Logstore group.
        /// </summary>
        [Input("collectTargetType")]
        public Input<string>? CollectTargetType { get; set; }

        /// <summary>
        /// The description of the metric import task.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The configurations of the logs that are imported from Log Service. See `sls_process_config` below.
        /// </summary>
        [Input("slsProcessConfig")]
        public Input<Inputs.HybridMonitorSlsTaskSlsProcessConfigGetArgs>? SlsProcessConfig { get; set; }

        /// <summary>
        /// The name of the metric import task, enter the name of the metric for logs imported from Log Service.
        /// </summary>
        [Input("taskName")]
        public Input<string>? TaskName { get; set; }

        public HybridMonitorSlsTaskState()
        {
        }
        public static new HybridMonitorSlsTaskState Empty => new HybridMonitorSlsTaskState();
    }
}
