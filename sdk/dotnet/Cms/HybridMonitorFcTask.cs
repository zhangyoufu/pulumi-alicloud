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
    /// Provides a Cloud Monitor Service Hybrid Monitor Fc Task resource.
    /// 
    /// For information about Cloud Monitor Service Hybrid Monitor Fc Task and how to use it, see [What is Hybrid Monitor Fc Task](https://www.alibabacloud.com/help/en/cloudmonitor/latest/createhybridmonitortask).
    /// 
    /// &gt; **NOTE:** Available since v1.179.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultAccount = AliCloud.GetAccount.Invoke();
    /// 
    ///     var defaultNamespace = new AliCloud.Cms.Namespace("defaultNamespace", new()
    ///     {
    ///         Description = name,
    ///         NamespaceName = name,
    ///         Specification = "cms.s1.large",
    ///     });
    /// 
    ///     var defaultHybridMonitorFcTask = new AliCloud.Cms.HybridMonitorFcTask("defaultHybridMonitorFcTask", new()
    ///     {
    ///         Namespace = defaultNamespace.Id,
    ///         YarmConfig = @"products:
    /// - namespace: acs_ecs_dashboard
    ///   metric_info:
    ///   - metric_list:
    ///     - cpu_total
    ///     - cpu_idle
    ///     - diskusage_utilization
    ///     - CPUUtilization
    ///     - DiskReadBPS
    ///     - InternetOut
    ///     - IntranetOut
    ///     - cpu_system
    /// - namespace: acs_rds_dashboard
    ///   metric_info:
    ///   - metric_list:
    ///     - MySQL_QPS
    ///     - MySQL_TPS
    /// ",
    ///         TargetUserId = defaultAccount.Apply(getAccountResult =&gt; getAccountResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Monitor Service Hybrid Monitor Fc Task can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cms/hybridMonitorFcTask:HybridMonitorFcTask example &lt;hybrid_monitor_fc_task_id&gt;:&lt;namespace&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cms/hybridMonitorFcTask:HybridMonitorFcTask")]
    public partial class HybridMonitorFcTask : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the monitoring task.
        /// </summary>
        [Output("hybridMonitorFcTaskId")]
        public Output<string> HybridMonitorFcTaskId { get; private set; } = null!;

        /// <summary>
        /// the namespace of the Alibaba Cloud service.
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;

        /// <summary>
        /// The ID of the member account. If you call API operations by using a management account, you can connect the Alibaba Cloud services that are activated for a member account in Resource Directory to Hybrid Cloud Monitoring. You can use Resource Directory to monitor Alibaba Cloud services across enterprise accounts.
        /// </summary>
        [Output("targetUserId")]
        public Output<string> TargetUserId { get; private set; } = null!;

        /// <summary>
        /// The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
        /// </summary>
        [Output("yarmConfig")]
        public Output<string> YarmConfig { get; private set; } = null!;


        /// <summary>
        /// Create a HybridMonitorFcTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HybridMonitorFcTask(string name, HybridMonitorFcTaskArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cms/hybridMonitorFcTask:HybridMonitorFcTask", name, args ?? new HybridMonitorFcTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HybridMonitorFcTask(string name, Input<string> id, HybridMonitorFcTaskState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cms/hybridMonitorFcTask:HybridMonitorFcTask", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HybridMonitorFcTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HybridMonitorFcTask Get(string name, Input<string> id, HybridMonitorFcTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new HybridMonitorFcTask(name, id, state, options);
        }
    }

    public sealed class HybridMonitorFcTaskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the namespace of the Alibaba Cloud service.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// The ID of the member account. If you call API operations by using a management account, you can connect the Alibaba Cloud services that are activated for a member account in Resource Directory to Hybrid Cloud Monitoring. You can use Resource Directory to monitor Alibaba Cloud services across enterprise accounts.
        /// </summary>
        [Input("targetUserId")]
        public Input<string>? TargetUserId { get; set; }

        /// <summary>
        /// The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
        /// </summary>
        [Input("yarmConfig", required: true)]
        public Input<string> YarmConfig { get; set; } = null!;

        public HybridMonitorFcTaskArgs()
        {
        }
        public static new HybridMonitorFcTaskArgs Empty => new HybridMonitorFcTaskArgs();
    }

    public sealed class HybridMonitorFcTaskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the monitoring task.
        /// </summary>
        [Input("hybridMonitorFcTaskId")]
        public Input<string>? HybridMonitorFcTaskId { get; set; }

        /// <summary>
        /// the namespace of the Alibaba Cloud service.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The ID of the member account. If you call API operations by using a management account, you can connect the Alibaba Cloud services that are activated for a member account in Resource Directory to Hybrid Cloud Monitoring. You can use Resource Directory to monitor Alibaba Cloud services across enterprise accounts.
        /// </summary>
        [Input("targetUserId")]
        public Input<string>? TargetUserId { get; set; }

        /// <summary>
        /// The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
        /// </summary>
        [Input("yarmConfig")]
        public Input<string>? YarmConfig { get; set; }

        public HybridMonitorFcTaskState()
        {
        }
        public static new HybridMonitorFcTaskState Empty => new HybridMonitorFcTaskState();
    }
}
