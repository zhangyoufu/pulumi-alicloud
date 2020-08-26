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
    /// &gt; **NOTE:** Available in 1.78.0
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
    ///         var exampleProject = new AliCloud.Log.Project("exampleProject", new AliCloud.Log.ProjectArgs
    ///         {
    ///             Description = "create by terraform",
    ///         });
    ///         var exampleStore = new AliCloud.Log.Store("exampleStore", new AliCloud.Log.StoreArgs
    ///         {
    ///             AppendMeta = true,
    ///             AutoSplit = true,
    ///             MaxSplitShardCount = 60,
    ///             Project = exampleProject.Name,
    ///             RetentionPeriod = 3650,
    ///             ShardCount = 3,
    ///         });
    ///         var exampleAlert = new AliCloud.Log.Alert("exampleAlert", new AliCloud.Log.AlertArgs
    ///         {
    ///             AlertDisplayname = "tf-test-alert-displayname",
    ///             AlertName = "tf-test-alert",
    ///             Condition = "count&gt; 100",
    ///             Dashboard = "tf-test-dashboard",
    ///             NotificationLists = 
    ///             {
    ///                 new AliCloud.Log.Inputs.AlertNotificationListArgs
    ///                 {
    ///                     Content = "alert content",
    ///                     MobileLists = 
    ///                     {
    ///                         "12345678",
    ///                         "87654321",
    ///                     },
    ///                     Type = "SMS",
    ///                 },
    ///                 new AliCloud.Log.Inputs.AlertNotificationListArgs
    ///                 {
    ///                     Content = "alert content",
    ///                     EmailLists = 
    ///                     {
    ///                         "aliyun@alibaba-inc.com",
    ///                         "tf-test@123.com",
    ///                     },
    ///                     Type = "Email",
    ///                 },
    ///                 new AliCloud.Log.Inputs.AlertNotificationListArgs
    ///                 {
    ///                     Content = "alert content",
    ///                     ServiceUri = "www.aliyun.com",
    ///                     Type = "DingTalk",
    ///                 },
    ///             },
    ///             ProjectName = exampleProject.Name,
    ///             QueryLists = 
    ///             {
    ///                 new AliCloud.Log.Inputs.AlertQueryListArgs
    ///                 {
    ///                     ChartTitle = "chart_title",
    ///                     End = "20s",
    ///                     Logstore = "tf-test-logstore",
    ///                     Query = "* AND aliyun",
    ///                     Start = "-60s",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Alert : Pulumi.CustomResource
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
        /// Conditional expression, such as: count&gt; 100.
        /// </summary>
        [Output("condition")]
        public Output<string> Condition { get; private set; } = null!;

        [Output("dashboard")]
        public Output<string> Dashboard { get; private set; } = null!;

        /// <summary>
        /// Timestamp, notifications before closing again.
        /// </summary>
        [Output("muteUntil")]
        public Output<int?> MuteUntil { get; private set; } = null!;

        /// <summary>
        /// Alarm information notification list.
        /// </summary>
        [Output("notificationLists")]
        public Output<ImmutableArray<Outputs.AlertNotificationList>> NotificationLists { get; private set; } = null!;

        /// <summary>
        /// Notification threshold, which is not notified until the number of triggers is reached. The default is 1.
        /// </summary>
        [Output("notifyThreshold")]
        public Output<int?> NotifyThreshold { get; private set; } = null!;

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
        /// Execution interval. 60 seconds minimum, such as 60s, 1h.
        /// </summary>
        [Output("scheduleInterval")]
        public Output<string?> ScheduleInterval { get; private set; } = null!;

        /// <summary>
        /// Default FixedRate. No need to configure this parameter.
        /// </summary>
        [Output("scheduleType")]
        public Output<string?> ScheduleType { get; private set; } = null!;

        /// <summary>
        /// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h.
        /// </summary>
        [Output("throttling")]
        public Output<string?> Throttling { get; private set; } = null!;


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

    public sealed class AlertArgs : Pulumi.ResourceArgs
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

        /// <summary>
        /// Conditional expression, such as: count&gt; 100.
        /// </summary>
        [Input("condition", required: true)]
        public Input<string> Condition { get; set; } = null!;

        [Input("dashboard", required: true)]
        public Input<string> Dashboard { get; set; } = null!;

        /// <summary>
        /// Timestamp, notifications before closing again.
        /// </summary>
        [Input("muteUntil")]
        public Input<int>? MuteUntil { get; set; }

        [Input("notificationLists", required: true)]
        private InputList<Inputs.AlertNotificationListArgs>? _notificationLists;

        /// <summary>
        /// Alarm information notification list.
        /// </summary>
        public InputList<Inputs.AlertNotificationListArgs> NotificationLists
        {
            get => _notificationLists ?? (_notificationLists = new InputList<Inputs.AlertNotificationListArgs>());
            set => _notificationLists = value;
        }

        /// <summary>
        /// Notification threshold, which is not notified until the number of triggers is reached. The default is 1.
        /// </summary>
        [Input("notifyThreshold")]
        public Input<int>? NotifyThreshold { get; set; }

        /// <summary>
        /// The project name.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        [Input("queryLists", required: true)]
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
        /// Execution interval. 60 seconds minimum, such as 60s, 1h.
        /// </summary>
        [Input("scheduleInterval")]
        public Input<string>? ScheduleInterval { get; set; }

        /// <summary>
        /// Default FixedRate. No need to configure this parameter.
        /// </summary>
        [Input("scheduleType")]
        public Input<string>? ScheduleType { get; set; }

        /// <summary>
        /// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h.
        /// </summary>
        [Input("throttling")]
        public Input<string>? Throttling { get; set; }

        public AlertArgs()
        {
        }
    }

    public sealed class AlertState : Pulumi.ResourceArgs
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

        /// <summary>
        /// Conditional expression, such as: count&gt; 100.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        [Input("dashboard")]
        public Input<string>? Dashboard { get; set; }

        /// <summary>
        /// Timestamp, notifications before closing again.
        /// </summary>
        [Input("muteUntil")]
        public Input<int>? MuteUntil { get; set; }

        [Input("notificationLists")]
        private InputList<Inputs.AlertNotificationListGetArgs>? _notificationLists;

        /// <summary>
        /// Alarm information notification list.
        /// </summary>
        public InputList<Inputs.AlertNotificationListGetArgs> NotificationLists
        {
            get => _notificationLists ?? (_notificationLists = new InputList<Inputs.AlertNotificationListGetArgs>());
            set => _notificationLists = value;
        }

        /// <summary>
        /// Notification threshold, which is not notified until the number of triggers is reached. The default is 1.
        /// </summary>
        [Input("notifyThreshold")]
        public Input<int>? NotifyThreshold { get; set; }

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
        /// Execution interval. 60 seconds minimum, such as 60s, 1h.
        /// </summary>
        [Input("scheduleInterval")]
        public Input<string>? ScheduleInterval { get; set; }

        /// <summary>
        /// Default FixedRate. No need to configure this parameter.
        /// </summary>
        [Input("scheduleType")]
        public Input<string>? ScheduleType { get; set; }

        /// <summary>
        /// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h.
        /// </summary>
        [Input("throttling")]
        public Input<string>? Throttling { get; set; }

        public AlertState()
        {
        }
    }
}
