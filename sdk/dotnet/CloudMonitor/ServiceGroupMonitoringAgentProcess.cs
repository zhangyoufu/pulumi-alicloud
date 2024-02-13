// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudMonitor
{
    /// <summary>
    /// Provides a Cloud Monitor Service Group Monitoring Agent Process resource.
    /// 
    /// For information about Cloud Monitor Service Group Monitoring Agent Process and how to use it, see [What is Group Monitoring Agent Process](https://www.alibabacloud.com/help/en/cms/developer-reference/api-cms-2019-01-01-creategroupmonitoringagentprocess).
    /// 
    /// &gt; **NOTE:** Available since v1.212.0.
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
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultAlarmContactGroup = new AliCloud.Cms.AlarmContactGroup("defaultAlarmContactGroup", new()
    ///     {
    ///         AlarmContactGroupName = name,
    ///         Contacts = new[]
    ///         {
    ///             "user",
    ///             "user1",
    ///             "user2",
    ///         },
    ///     });
    /// 
    ///     var defaultMonitorGroup = new AliCloud.Cms.MonitorGroup("defaultMonitorGroup", new()
    ///     {
    ///         MonitorGroupName = name,
    ///         ContactGroups = new[]
    ///         {
    ///             defaultAlarmContactGroup.Id,
    ///         },
    ///     });
    /// 
    ///     var defaultServiceGroupMonitoringAgentProcess = new AliCloud.CloudMonitor.ServiceGroupMonitoringAgentProcess("defaultServiceGroupMonitoringAgentProcess", new()
    ///     {
    ///         GroupId = defaultMonitorGroup.Id,
    ///         ProcessName = name,
    ///         MatchExpressFilterRelation = "or",
    ///         MatchExpresses = new[]
    ///         {
    ///             new AliCloud.CloudMonitor.Inputs.ServiceGroupMonitoringAgentProcessMatchExpressArgs
    ///             {
    ///                 Name = name,
    ///                 Value = "*",
    ///                 Function = "all",
    ///             },
    ///         },
    ///         AlertConfigs = new[]
    ///         {
    ///             new AliCloud.CloudMonitor.Inputs.ServiceGroupMonitoringAgentProcessAlertConfigArgs
    ///             {
    ///                 EscalationsLevel = "critical",
    ///                 ComparisonOperator = "GreaterThanOrEqualToThreshold",
    ///                 Statistics = "Average",
    ///                 Threshold = "20",
    ///                 Times = "100",
    ///                 EffectiveInterval = "00:00-22:59",
    ///                 SilenceTime = 85800,
    ///                 Webhook = "https://www.aliyun.com",
    ///                 TargetLists = new[]
    ///                 {
    ///                     new AliCloud.CloudMonitor.Inputs.ServiceGroupMonitoringAgentProcessAlertConfigTargetListArgs
    ///                     {
    ///                         TargetListId = "1",
    ///                         JsonParams = "{}",
    ///                         Level = "WARN",
    ///                         Arn = "acs:mns:cn-hangzhou:120886317861****:/queues/test123/message",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Monitor Service Group Monitoring Agent Process can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cloudmonitor/serviceGroupMonitoringAgentProcess:ServiceGroupMonitoringAgentProcess example &lt;group_id&gt;:&lt;group_monitoring_agent_process_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cloudmonitor/serviceGroupMonitoringAgentProcess:ServiceGroupMonitoringAgentProcess")]
    public partial class ServiceGroupMonitoringAgentProcess : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The alert rule configurations. See `alert_config` below.
        /// </summary>
        [Output("alertConfigs")]
        public Output<ImmutableArray<Outputs.ServiceGroupMonitoringAgentProcessAlertConfig>> AlertConfigs { get; private set; } = null!;

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Group Monitoring Agent Process.
        /// </summary>
        [Output("groupMonitoringAgentProcessId")]
        public Output<string> GroupMonitoringAgentProcessId { get; private set; } = null!;

        /// <summary>
        /// The logical operator used between conditional expressions that are used to match instances. Valid values: `all`, `and`, `or`.
        /// </summary>
        [Output("matchExpressFilterRelation")]
        public Output<string> MatchExpressFilterRelation { get; private set; } = null!;

        /// <summary>
        /// The expressions used to match instances. See `match_express` below.
        /// </summary>
        [Output("matchExpresses")]
        public Output<ImmutableArray<Outputs.ServiceGroupMonitoringAgentProcessMatchExpress>> MatchExpresses { get; private set; } = null!;

        /// <summary>
        /// The name of the process.
        /// </summary>
        [Output("processName")]
        public Output<string> ProcessName { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceGroupMonitoringAgentProcess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceGroupMonitoringAgentProcess(string name, ServiceGroupMonitoringAgentProcessArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cloudmonitor/serviceGroupMonitoringAgentProcess:ServiceGroupMonitoringAgentProcess", name, args ?? new ServiceGroupMonitoringAgentProcessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceGroupMonitoringAgentProcess(string name, Input<string> id, ServiceGroupMonitoringAgentProcessState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cloudmonitor/serviceGroupMonitoringAgentProcess:ServiceGroupMonitoringAgentProcess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceGroupMonitoringAgentProcess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceGroupMonitoringAgentProcess Get(string name, Input<string> id, ServiceGroupMonitoringAgentProcessState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceGroupMonitoringAgentProcess(name, id, state, options);
        }
    }

    public sealed class ServiceGroupMonitoringAgentProcessArgs : global::Pulumi.ResourceArgs
    {
        [Input("alertConfigs", required: true)]
        private InputList<Inputs.ServiceGroupMonitoringAgentProcessAlertConfigArgs>? _alertConfigs;

        /// <summary>
        /// The alert rule configurations. See `alert_config` below.
        /// </summary>
        public InputList<Inputs.ServiceGroupMonitoringAgentProcessAlertConfigArgs> AlertConfigs
        {
            get => _alertConfigs ?? (_alertConfigs = new InputList<Inputs.ServiceGroupMonitoringAgentProcessAlertConfigArgs>());
            set => _alertConfigs = value;
        }

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The logical operator used between conditional expressions that are used to match instances. Valid values: `all`, `and`, `or`.
        /// </summary>
        [Input("matchExpressFilterRelation")]
        public Input<string>? MatchExpressFilterRelation { get; set; }

        [Input("matchExpresses")]
        private InputList<Inputs.ServiceGroupMonitoringAgentProcessMatchExpressArgs>? _matchExpresses;

        /// <summary>
        /// The expressions used to match instances. See `match_express` below.
        /// </summary>
        public InputList<Inputs.ServiceGroupMonitoringAgentProcessMatchExpressArgs> MatchExpresses
        {
            get => _matchExpresses ?? (_matchExpresses = new InputList<Inputs.ServiceGroupMonitoringAgentProcessMatchExpressArgs>());
            set => _matchExpresses = value;
        }

        /// <summary>
        /// The name of the process.
        /// </summary>
        [Input("processName", required: true)]
        public Input<string> ProcessName { get; set; } = null!;

        public ServiceGroupMonitoringAgentProcessArgs()
        {
        }
        public static new ServiceGroupMonitoringAgentProcessArgs Empty => new ServiceGroupMonitoringAgentProcessArgs();
    }

    public sealed class ServiceGroupMonitoringAgentProcessState : global::Pulumi.ResourceArgs
    {
        [Input("alertConfigs")]
        private InputList<Inputs.ServiceGroupMonitoringAgentProcessAlertConfigGetArgs>? _alertConfigs;

        /// <summary>
        /// The alert rule configurations. See `alert_config` below.
        /// </summary>
        public InputList<Inputs.ServiceGroupMonitoringAgentProcessAlertConfigGetArgs> AlertConfigs
        {
            get => _alertConfigs ?? (_alertConfigs = new InputList<Inputs.ServiceGroupMonitoringAgentProcessAlertConfigGetArgs>());
            set => _alertConfigs = value;
        }

        /// <summary>
        /// The ID of the application group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The ID of the Group Monitoring Agent Process.
        /// </summary>
        [Input("groupMonitoringAgentProcessId")]
        public Input<string>? GroupMonitoringAgentProcessId { get; set; }

        /// <summary>
        /// The logical operator used between conditional expressions that are used to match instances. Valid values: `all`, `and`, `or`.
        /// </summary>
        [Input("matchExpressFilterRelation")]
        public Input<string>? MatchExpressFilterRelation { get; set; }

        [Input("matchExpresses")]
        private InputList<Inputs.ServiceGroupMonitoringAgentProcessMatchExpressGetArgs>? _matchExpresses;

        /// <summary>
        /// The expressions used to match instances. See `match_express` below.
        /// </summary>
        public InputList<Inputs.ServiceGroupMonitoringAgentProcessMatchExpressGetArgs> MatchExpresses
        {
            get => _matchExpresses ?? (_matchExpresses = new InputList<Inputs.ServiceGroupMonitoringAgentProcessMatchExpressGetArgs>());
            set => _matchExpresses = value;
        }

        /// <summary>
        /// The name of the process.
        /// </summary>
        [Input("processName")]
        public Input<string>? ProcessName { get; set; }

        public ServiceGroupMonitoringAgentProcessState()
        {
        }
        public static new ServiceGroupMonitoringAgentProcessState Empty => new ServiceGroupMonitoringAgentProcessState();
    }
}
