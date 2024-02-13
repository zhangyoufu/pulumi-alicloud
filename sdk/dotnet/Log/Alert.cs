// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log
{
    /// <summary>
    /// Log alert is a unit of log service, which is used to monitor and alert the user's logstore status information.
    /// Log Service enables you to configure alerts based on the charts in a dashboard to monitor the service status in real time.
    /// 
    /// For information about SLS Alert and how to use it, see [SLS Alert Overview](https://www.alibabacloud.com/help/en/doc-detail/209202.html)
    /// 
    /// &gt; **NOTE:** Available in 1.78.0
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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var exampleProject = new AliCloud.Log.Project("exampleProject", new()
    ///     {
    ///         Description = "terraform-example",
    ///     });
    /// 
    ///     var exampleStore = new AliCloud.Log.Store("exampleStore", new()
    ///     {
    ///         Project = exampleProject.Name,
    ///         RetentionPeriod = 3650,
    ///         ShardCount = 3,
    ///         AutoSplit = true,
    ///         MaxSplitShardCount = 60,
    ///         AppendMeta = true,
    ///     });
    /// 
    ///     var exampleAlert = new AliCloud.Log.Alert("exampleAlert", new()
    ///     {
    ///         ProjectName = exampleProject.Name,
    ///         AlertName = "example-alert",
    ///         AlertDisplayname = "example-alert",
    ///         Condition = "count&gt; 100",
    ///         Dashboard = "example-dashboard",
    ///         Schedule = new AliCloud.Log.Inputs.AlertScheduleArgs
    ///         {
    ///             Type = "FixedRate",
    ///             Interval = "5m",
    ///             Hour = 0,
    ///             DayOfWeek = 0,
    ///             Delay = 0,
    ///             RunImmediately = false,
    ///         },
    ///         QueryLists = new[]
    ///         {
    ///             new AliCloud.Log.Inputs.AlertQueryListArgs
    ///             {
    ///                 Logstore = exampleStore.Name,
    ///                 ChartTitle = "chart_title",
    ///                 Start = "-60s",
    ///                 End = "20s",
    ///                 Query = "* AND aliyun",
    ///             },
    ///         },
    ///         NotificationLists = new[]
    ///         {
    ///             new AliCloud.Log.Inputs.AlertNotificationListArgs
    ///             {
    ///                 Type = "SMS",
    ///                 MobileLists = new[]
    ///                 {
    ///                     "12345678",
    ///                     "87654321",
    ///                 },
    ///                 Content = "alert content",
    ///             },
    ///             new AliCloud.Log.Inputs.AlertNotificationListArgs
    ///             {
    ///                 Type = "Email",
    ///                 EmailLists = new[]
    ///                 {
    ///                     "aliyun@alibaba-inc.com",
    ///                     "tf-example@123.com",
    ///                 },
    ///                 Content = "alert content",
    ///             },
    ///             new AliCloud.Log.Inputs.AlertNotificationListArgs
    ///             {
    ///                 Type = "DingTalk",
    ///                 ServiceUri = "www.aliyun.com",
    ///                 Content = "alert content",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Basic Usage for new alert
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var exampleProject = new AliCloud.Log.Project("exampleProject", new()
    ///     {
    ///         Description = "terraform-example",
    ///     });
    /// 
    ///     var exampleStore = new AliCloud.Log.Store("exampleStore", new()
    ///     {
    ///         Project = exampleProject.Name,
    ///         RetentionPeriod = 3650,
    ///         ShardCount = 3,
    ///         AutoSplit = true,
    ///         MaxSplitShardCount = 60,
    ///         AppendMeta = true,
    ///     });
    /// 
    ///     var example_2 = new AliCloud.Log.Alert("example-2", new()
    ///     {
    ///         Version = "2.0",
    ///         Type = "default",
    ///         ProjectName = exampleProject.Name,
    ///         AlertName = "example-alert",
    ///         AlertDisplayname = "example-alert",
    ///         MuteUntil = 1632486684,
    ///         NoDataFire = false,
    ///         NoDataSeverity = 8,
    ///         SendResolved = true,
    ///         AutoAnnotation = true,
    ///         Dashboard = "example-dashboard",
    ///         Schedule = new AliCloud.Log.Inputs.AlertScheduleArgs
    ///         {
    ///             Type = "FixedRate",
    ///             Interval = "5m",
    ///             Hour = 0,
    ///             DayOfWeek = 0,
    ///             Delay = 0,
    ///             RunImmediately = false,
    ///         },
    ///         QueryLists = new[]
    ///         {
    ///             new AliCloud.Log.Inputs.AlertQueryListArgs
    ///             {
    ///                 Store = exampleStore.Name,
    ///                 StoreType = "log",
    ///                 Project = exampleProject.Name,
    ///                 Region = "cn-heyuan",
    ///                 ChartTitle = "chart_title",
    ///                 Start = "-60s",
    ///                 End = "20s",
    ///                 Query = "* AND aliyun | select count(1) as cnt",
    ///                 PowerSqlMode = "auto",
    ///             },
    ///             new AliCloud.Log.Inputs.AlertQueryListArgs
    ///             {
    ///                 Store = exampleStore.Name,
    ///                 StoreType = "log",
    ///                 Project = exampleProject.Name,
    ///                 Region = "cn-heyuan",
    ///                 ChartTitle = "chart_title",
    ///                 Start = "-60s",
    ///                 End = "20s",
    ///                 Query = "error | select count(1) as error_cnt",
    ///                 PowerSqlMode = "enable",
    ///             },
    ///         },
    ///         Labels = new[]
    ///         {
    ///             new AliCloud.Log.Inputs.AlertLabelArgs
    ///             {
    ///                 Key = "env",
    ///                 Value = "test",
    ///             },
    ///         },
    ///         Annotations = new[]
    ///         {
    ///             new AliCloud.Log.Inputs.AlertAnnotationArgs
    ///             {
    ///                 Key = "title",
    ///                 Value = "alert title",
    ///             },
    ///             new AliCloud.Log.Inputs.AlertAnnotationArgs
    ///             {
    ///                 Key = "desc",
    ///                 Value = "alert desc",
    ///             },
    ///             new AliCloud.Log.Inputs.AlertAnnotationArgs
    ///             {
    ///                 Key = "test_key",
    ///                 Value = "test value",
    ///             },
    ///         },
    ///         GroupConfiguration = new AliCloud.Log.Inputs.AlertGroupConfigurationArgs
    ///         {
    ///             Type = "custom",
    ///             Fields = new[]
    ///             {
    ///                 "cnt",
    ///             },
    ///         },
    ///         PolicyConfiguration = new AliCloud.Log.Inputs.AlertPolicyConfigurationArgs
    ///         {
    ///             AlertPolicyId = "sls.bultin",
    ///             ActionPolicyId = "sls_test_action",
    ///             RepeatInterval = "4h",
    ///         },
    ///         SeverityConfigurations = new[]
    ///         {
    ///             new AliCloud.Log.Inputs.AlertSeverityConfigurationArgs
    ///             {
    ///                 Severity = 8,
    ///                 EvalCondition = 
    ///                 {
    ///                     { "condition", "cnt &gt; 3" },
    ///                     { "count_condition", "__count__ &gt; 3" },
    ///                 },
    ///             },
    ///             new AliCloud.Log.Inputs.AlertSeverityConfigurationArgs
    ///             {
    ///                 Severity = 6,
    ///                 EvalCondition = 
    ///                 {
    ///                     { "condition", "" },
    ///                     { "count_condition", "__count__ &gt; 0" },
    ///                 },
    ///             },
    ///             new AliCloud.Log.Inputs.AlertSeverityConfigurationArgs
    ///             {
    ///                 Severity = 2,
    ///                 EvalCondition = 
    ///                 {
    ///                     { "condition", "" },
    ///                     { "count_condition", "" },
    ///                 },
    ///             },
    ///         },
    ///         JoinConfigurations = new[]
    ///         {
    ///             new AliCloud.Log.Inputs.AlertJoinConfigurationArgs
    ///             {
    ///                 Type = "cross_join",
    ///                 Condition = "",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Basic Usage for alert template
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var exampleProject = new AliCloud.Log.Project("exampleProject", new()
    ///     {
    ///         Description = "terraform-example",
    ///     });
    /// 
    ///     var exampleStore = new AliCloud.Log.Store("exampleStore", new()
    ///     {
    ///         Project = exampleProject.Name,
    ///         RetentionPeriod = 3650,
    ///         ShardCount = 3,
    ///         AutoSplit = true,
    ///         MaxSplitShardCount = 60,
    ///         AppendMeta = true,
    ///     });
    /// 
    ///     var example_3 = new AliCloud.Log.Alert("example-3", new()
    ///     {
    ///         Version = "2.0",
    ///         Type = "tpl",
    ///         ProjectName = exampleProject.Name,
    ///         AlertName = "example-alert",
    ///         AlertDisplayname = "example-alert",
    ///         MuteUntil = 1632486684,
    ///         Schedule = new AliCloud.Log.Inputs.AlertScheduleArgs
    ///         {
    ///             Type = "FixedRate",
    ///             Interval = "5m",
    ///             Hour = 0,
    ///             DayOfWeek = 0,
    ///             Delay = 0,
    ///             RunImmediately = false,
    ///         },
    ///         TemplateConfiguration = new AliCloud.Log.Inputs.AlertTemplateConfigurationArgs
    ///         {
    ///             Id = "sls.app.sls_ack.node.down",
    ///             Type = "sys",
    ///             Lang = "cn",
    ///             Annotations = null,
    ///             Tokens = 
    ///             {
    ///                 { "interval_minute", "5" },
    ///                 { "default.action_policy", "sls.app.ack.builtin" },
    ///                 { "default.severity", "6" },
    ///                 { "sendResolved", "false" },
    ///                 { "default.project", exampleProject.Name },
    ///                 { "default.logstore", "k8s-event" },
    ///                 { "default.repeatInterval", "4h" },
    ///                 { "trigger_threshold", "1" },
    ///                 { "default.clusterId", "example-cluster-id" },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Log alert can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:log/alert:Alert example tf-log:tf-log-alert
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:log/alert:Alert")]
    public partial class Alert : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Alert description.
        /// </summary>
        [Output("alertDescription")]
        public Output<string?> AlertDescription { get; private set; } = null!;

        /// <summary>
        /// Alert displayname.
        /// </summary>
        [Output("alertDisplayname")]
        public Output<string> AlertDisplayname { get; private set; } = null!;

        /// <summary>
        /// Name of logstore for configuring alarm service.
        /// </summary>
        [Output("alertName")]
        public Output<string> AlertName { get; private set; } = null!;

        /// <summary>
        /// Alert template annotations.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableArray<Outputs.AlertAnnotation>> Annotations { get; private set; } = null!;

        /// <summary>
        /// whether to add automatic annotation, default is false.
        /// </summary>
        [Output("autoAnnotation")]
        public Output<bool?> AutoAnnotation { get; private set; } = null!;

        /// <summary>
        /// Join condition.
        /// </summary>
        [Output("condition")]
        public Output<string?> Condition { get; private set; } = null!;

        [Output("dashboard")]
        public Output<string?> Dashboard { get; private set; } = null!;

        /// <summary>
        /// Group configuration for new alert.
        /// </summary>
        [Output("groupConfiguration")]
        public Output<Outputs.AlertGroupConfiguration?> GroupConfiguration { get; private set; } = null!;

        /// <summary>
        /// Join configuration for different queries.
        /// </summary>
        [Output("joinConfigurations")]
        public Output<ImmutableArray<Outputs.AlertJoinConfiguration>> JoinConfigurations { get; private set; } = null!;

        /// <summary>
        /// Labels for new alert.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.AlertLabel>> Labels { get; private set; } = null!;

        /// <summary>
        /// Timestamp, notifications before closing again.
        /// </summary>
        [Output("muteUntil")]
        public Output<int> MuteUntil { get; private set; } = null!;

        /// <summary>
        /// Switch for whether new alert fires when no data happens, default is false.
        /// </summary>
        [Output("noDataFire")]
        public Output<bool?> NoDataFire { get; private set; } = null!;

        /// <summary>
        /// when no data happens, the severity of new alert.
        /// </summary>
        [Output("noDataSeverity")]
        public Output<int?> NoDataSeverity { get; private set; } = null!;

        /// <summary>
        /// Alarm information notification list, Deprecated from 1.161.0+.
        /// </summary>
        [Output("notificationLists")]
        public Output<ImmutableArray<Outputs.AlertNotificationList>> NotificationLists { get; private set; } = null!;

        /// <summary>
        /// Notification threshold, which is not notified until the number of triggers is reached. The default is 1, Deprecated from 1.161.0+.
        /// </summary>
        [Output("notifyThreshold")]
        public Output<int?> NotifyThreshold { get; private set; } = null!;

        /// <summary>
        /// Policy configuration for new alert.
        /// </summary>
        [Output("policyConfiguration")]
        public Output<Outputs.AlertPolicyConfiguration?> PolicyConfiguration { get; private set; } = null!;

        /// <summary>
        /// The project name.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// Multiple conditions for configured alarm query.
        /// </summary>
        [Output("queryLists")]
        public Output<ImmutableArray<Outputs.AlertQueryList>> QueryLists { get; private set; } = null!;

        /// <summary>
        /// schedule for alert.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.AlertSchedule?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Execution interval. 60 seconds minimum, such as 60s, 1h. Deprecated from 1.176.0+. use interval in schedule.
        /// </summary>
        [Output("scheduleInterval")]
        public Output<string> ScheduleInterval { get; private set; } = null!;

        /// <summary>
        /// Default FixedRate. No need to configure this parameter. Deprecated from 1.176.0+. use type in schedule.
        /// </summary>
        [Output("scheduleType")]
        public Output<string> ScheduleType { get; private set; } = null!;

        /// <summary>
        /// when new alert is resolved, whether to notify, default is false.
        /// </summary>
        [Output("sendResolved")]
        public Output<bool?> SendResolved { get; private set; } = null!;

        /// <summary>
        /// Severity configuration for new alert.
        /// </summary>
        [Output("severityConfigurations")]
        public Output<ImmutableArray<Outputs.AlertSeverityConfiguration>> SeverityConfigurations { get; private set; } = null!;

        /// <summary>
        /// Template configuration for alert, when `type` is `tpl`.
        /// </summary>
        [Output("templateConfiguration")]
        public Output<Outputs.AlertTemplateConfiguration?> TemplateConfiguration { get; private set; } = null!;

        /// <summary>
        /// Evaluation threshold, alert will not fire until the number of triggers is reached. The default is 1.
        /// </summary>
        [Output("threshold")]
        public Output<int> Threshold { get; private set; } = null!;

        /// <summary>
        /// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h, Deprecated from 1.161.0+.
        /// </summary>
        [Output("throttling")]
        public Output<string?> Throttling { get; private set; } = null!;

        /// <summary>
        /// including FixedRate,Hourly,Daily,Weekly,Cron.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// The version of alert, new alert is 2.0.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Alert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Alert(string name, AlertArgs args, CustomResourceOptions? options = null)
            : base("alicloud:log/alert:Alert", name, args ?? new AlertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Alert(string name, Input<string> id, AlertState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:log/alert:Alert", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Alert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Alert Get(string name, Input<string> id, AlertState? state = null, CustomResourceOptions? options = null)
        {
            return new Alert(name, id, state, options);
        }
    }

    public sealed class AlertArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alert description.
        /// </summary>
        [Input("alertDescription")]
        public Input<string>? AlertDescription { get; set; }

        /// <summary>
        /// Alert displayname.
        /// </summary>
        [Input("alertDisplayname", required: true)]
        public Input<string> AlertDisplayname { get; set; } = null!;

        /// <summary>
        /// Name of logstore for configuring alarm service.
        /// </summary>
        [Input("alertName", required: true)]
        public Input<string> AlertName { get; set; } = null!;

        [Input("annotations")]
        private InputList<Inputs.AlertAnnotationArgs>? _annotations;

        /// <summary>
        /// Alert template annotations.
        /// </summary>
        public InputList<Inputs.AlertAnnotationArgs> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<Inputs.AlertAnnotationArgs>());
            set => _annotations = value;
        }

        /// <summary>
        /// whether to add automatic annotation, default is false.
        /// </summary>
        [Input("autoAnnotation")]
        public Input<bool>? AutoAnnotation { get; set; }

        /// <summary>
        /// Join condition.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        [Input("dashboard")]
        public Input<string>? Dashboard { get; set; }

        /// <summary>
        /// Group configuration for new alert.
        /// </summary>
        [Input("groupConfiguration")]
        public Input<Inputs.AlertGroupConfigurationArgs>? GroupConfiguration { get; set; }

        [Input("joinConfigurations")]
        private InputList<Inputs.AlertJoinConfigurationArgs>? _joinConfigurations;

        /// <summary>
        /// Join configuration for different queries.
        /// </summary>
        public InputList<Inputs.AlertJoinConfigurationArgs> JoinConfigurations
        {
            get => _joinConfigurations ?? (_joinConfigurations = new InputList<Inputs.AlertJoinConfigurationArgs>());
            set => _joinConfigurations = value;
        }

        [Input("labels")]
        private InputList<Inputs.AlertLabelArgs>? _labels;

        /// <summary>
        /// Labels for new alert.
        /// </summary>
        public InputList<Inputs.AlertLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.AlertLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Timestamp, notifications before closing again.
        /// </summary>
        [Input("muteUntil")]
        public Input<int>? MuteUntil { get; set; }

        /// <summary>
        /// Switch for whether new alert fires when no data happens, default is false.
        /// </summary>
        [Input("noDataFire")]
        public Input<bool>? NoDataFire { get; set; }

        /// <summary>
        /// when no data happens, the severity of new alert.
        /// </summary>
        [Input("noDataSeverity")]
        public Input<int>? NoDataSeverity { get; set; }

        [Input("notificationLists")]
        private InputList<Inputs.AlertNotificationListArgs>? _notificationLists;

        /// <summary>
        /// Alarm information notification list, Deprecated from 1.161.0+.
        /// </summary>
        [Obsolete(@"Deprecated from 1.161.0+, use policy_configuration for notification")]
        public InputList<Inputs.AlertNotificationListArgs> NotificationLists
        {
            get => _notificationLists ?? (_notificationLists = new InputList<Inputs.AlertNotificationListArgs>());
            set => _notificationLists = value;
        }

        /// <summary>
        /// Notification threshold, which is not notified until the number of triggers is reached. The default is 1, Deprecated from 1.161.0+.
        /// </summary>
        [Input("notifyThreshold")]
        public Input<int>? NotifyThreshold { get; set; }

        /// <summary>
        /// Policy configuration for new alert.
        /// </summary>
        [Input("policyConfiguration")]
        public Input<Inputs.AlertPolicyConfigurationArgs>? PolicyConfiguration { get; set; }

        /// <summary>
        /// The project name.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        [Input("queryLists")]
        private InputList<Inputs.AlertQueryListArgs>? _queryLists;

        /// <summary>
        /// Multiple conditions for configured alarm query.
        /// </summary>
        public InputList<Inputs.AlertQueryListArgs> QueryLists
        {
            get => _queryLists ?? (_queryLists = new InputList<Inputs.AlertQueryListArgs>());
            set => _queryLists = value;
        }

        /// <summary>
        /// schedule for alert.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.AlertScheduleArgs>? Schedule { get; set; }

        /// <summary>
        /// Execution interval. 60 seconds minimum, such as 60s, 1h. Deprecated from 1.176.0+. use interval in schedule.
        /// </summary>
        [Input("scheduleInterval")]
        public Input<string>? ScheduleInterval { get; set; }

        /// <summary>
        /// Default FixedRate. No need to configure this parameter. Deprecated from 1.176.0+. use type in schedule.
        /// </summary>
        [Input("scheduleType")]
        public Input<string>? ScheduleType { get; set; }

        /// <summary>
        /// when new alert is resolved, whether to notify, default is false.
        /// </summary>
        [Input("sendResolved")]
        public Input<bool>? SendResolved { get; set; }

        [Input("severityConfigurations")]
        private InputList<Inputs.AlertSeverityConfigurationArgs>? _severityConfigurations;

        /// <summary>
        /// Severity configuration for new alert.
        /// </summary>
        public InputList<Inputs.AlertSeverityConfigurationArgs> SeverityConfigurations
        {
            get => _severityConfigurations ?? (_severityConfigurations = new InputList<Inputs.AlertSeverityConfigurationArgs>());
            set => _severityConfigurations = value;
        }

        /// <summary>
        /// Template configuration for alert, when `type` is `tpl`.
        /// </summary>
        [Input("templateConfiguration")]
        public Input<Inputs.AlertTemplateConfigurationArgs>? TemplateConfiguration { get; set; }

        /// <summary>
        /// Evaluation threshold, alert will not fire until the number of triggers is reached. The default is 1.
        /// </summary>
        [Input("threshold")]
        public Input<int>? Threshold { get; set; }

        /// <summary>
        /// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h, Deprecated from 1.161.0+.
        /// </summary>
        [Input("throttling")]
        public Input<string>? Throttling { get; set; }

        /// <summary>
        /// including FixedRate,Hourly,Daily,Weekly,Cron.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The version of alert, new alert is 2.0.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AlertArgs()
        {
        }
        public static new AlertArgs Empty => new AlertArgs();
    }

    public sealed class AlertState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alert description.
        /// </summary>
        [Input("alertDescription")]
        public Input<string>? AlertDescription { get; set; }

        /// <summary>
        /// Alert displayname.
        /// </summary>
        [Input("alertDisplayname")]
        public Input<string>? AlertDisplayname { get; set; }

        /// <summary>
        /// Name of logstore for configuring alarm service.
        /// </summary>
        [Input("alertName")]
        public Input<string>? AlertName { get; set; }

        [Input("annotations")]
        private InputList<Inputs.AlertAnnotationGetArgs>? _annotations;

        /// <summary>
        /// Alert template annotations.
        /// </summary>
        public InputList<Inputs.AlertAnnotationGetArgs> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<Inputs.AlertAnnotationGetArgs>());
            set => _annotations = value;
        }

        /// <summary>
        /// whether to add automatic annotation, default is false.
        /// </summary>
        [Input("autoAnnotation")]
        public Input<bool>? AutoAnnotation { get; set; }

        /// <summary>
        /// Join condition.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        [Input("dashboard")]
        public Input<string>? Dashboard { get; set; }

        /// <summary>
        /// Group configuration for new alert.
        /// </summary>
        [Input("groupConfiguration")]
        public Input<Inputs.AlertGroupConfigurationGetArgs>? GroupConfiguration { get; set; }

        [Input("joinConfigurations")]
        private InputList<Inputs.AlertJoinConfigurationGetArgs>? _joinConfigurations;

        /// <summary>
        /// Join configuration for different queries.
        /// </summary>
        public InputList<Inputs.AlertJoinConfigurationGetArgs> JoinConfigurations
        {
            get => _joinConfigurations ?? (_joinConfigurations = new InputList<Inputs.AlertJoinConfigurationGetArgs>());
            set => _joinConfigurations = value;
        }

        [Input("labels")]
        private InputList<Inputs.AlertLabelGetArgs>? _labels;

        /// <summary>
        /// Labels for new alert.
        /// </summary>
        public InputList<Inputs.AlertLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.AlertLabelGetArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Timestamp, notifications before closing again.
        /// </summary>
        [Input("muteUntil")]
        public Input<int>? MuteUntil { get; set; }

        /// <summary>
        /// Switch for whether new alert fires when no data happens, default is false.
        /// </summary>
        [Input("noDataFire")]
        public Input<bool>? NoDataFire { get; set; }

        /// <summary>
        /// when no data happens, the severity of new alert.
        /// </summary>
        [Input("noDataSeverity")]
        public Input<int>? NoDataSeverity { get; set; }

        [Input("notificationLists")]
        private InputList<Inputs.AlertNotificationListGetArgs>? _notificationLists;

        /// <summary>
        /// Alarm information notification list, Deprecated from 1.161.0+.
        /// </summary>
        [Obsolete(@"Deprecated from 1.161.0+, use policy_configuration for notification")]
        public InputList<Inputs.AlertNotificationListGetArgs> NotificationLists
        {
            get => _notificationLists ?? (_notificationLists = new InputList<Inputs.AlertNotificationListGetArgs>());
            set => _notificationLists = value;
        }

        /// <summary>
        /// Notification threshold, which is not notified until the number of triggers is reached. The default is 1, Deprecated from 1.161.0+.
        /// </summary>
        [Input("notifyThreshold")]
        public Input<int>? NotifyThreshold { get; set; }

        /// <summary>
        /// Policy configuration for new alert.
        /// </summary>
        [Input("policyConfiguration")]
        public Input<Inputs.AlertPolicyConfigurationGetArgs>? PolicyConfiguration { get; set; }

        /// <summary>
        /// The project name.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("queryLists")]
        private InputList<Inputs.AlertQueryListGetArgs>? _queryLists;

        /// <summary>
        /// Multiple conditions for configured alarm query.
        /// </summary>
        public InputList<Inputs.AlertQueryListGetArgs> QueryLists
        {
            get => _queryLists ?? (_queryLists = new InputList<Inputs.AlertQueryListGetArgs>());
            set => _queryLists = value;
        }

        /// <summary>
        /// schedule for alert.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.AlertScheduleGetArgs>? Schedule { get; set; }

        /// <summary>
        /// Execution interval. 60 seconds minimum, such as 60s, 1h. Deprecated from 1.176.0+. use interval in schedule.
        /// </summary>
        [Input("scheduleInterval")]
        public Input<string>? ScheduleInterval { get; set; }

        /// <summary>
        /// Default FixedRate. No need to configure this parameter. Deprecated from 1.176.0+. use type in schedule.
        /// </summary>
        [Input("scheduleType")]
        public Input<string>? ScheduleType { get; set; }

        /// <summary>
        /// when new alert is resolved, whether to notify, default is false.
        /// </summary>
        [Input("sendResolved")]
        public Input<bool>? SendResolved { get; set; }

        [Input("severityConfigurations")]
        private InputList<Inputs.AlertSeverityConfigurationGetArgs>? _severityConfigurations;

        /// <summary>
        /// Severity configuration for new alert.
        /// </summary>
        public InputList<Inputs.AlertSeverityConfigurationGetArgs> SeverityConfigurations
        {
            get => _severityConfigurations ?? (_severityConfigurations = new InputList<Inputs.AlertSeverityConfigurationGetArgs>());
            set => _severityConfigurations = value;
        }

        /// <summary>
        /// Template configuration for alert, when `type` is `tpl`.
        /// </summary>
        [Input("templateConfiguration")]
        public Input<Inputs.AlertTemplateConfigurationGetArgs>? TemplateConfiguration { get; set; }

        /// <summary>
        /// Evaluation threshold, alert will not fire until the number of triggers is reached. The default is 1.
        /// </summary>
        [Input("threshold")]
        public Input<int>? Threshold { get; set; }

        /// <summary>
        /// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h, Deprecated from 1.161.0+.
        /// </summary>
        [Input("throttling")]
        public Input<string>? Throttling { get; set; }

        /// <summary>
        /// including FixedRate,Hourly,Daily,Weekly,Cron.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The version of alert, new alert is 2.0.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AlertState()
        {
        }
        public static new AlertState Empty => new AlertState();
    }
}
