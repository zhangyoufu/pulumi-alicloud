// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess
{
    /// <summary>
    /// Provides a ESS schedule resource.
    /// 
    /// For information about ess schedule task, see [Scheduled Tasks](https://www.alibabacloud.com/help/en/auto-scaling/latest/createscheduledtask).
    /// 
    /// &gt; **NOTE:** Available since v1.60.0.
    /// 
    /// ## Module Support
    /// 
    /// You can use to the existing autoscaling-rule module
    /// to create scheduled task, different type rules and alarm task one-click.
    /// 
    /// ## Import
    /// 
    /// ESS schedule task can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ess/scheduledTask:ScheduledTask example abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ess/scheduledTask:ScheduledTask")]
    public partial class ScheduledTask : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the scheduled task, which is 2-200 characters (English or Chinese) long.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The expected number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group. **NOTE:** You must specify the `DesiredCapacity` parameter when you create the scaling group.
        /// </summary>
        [Output("desiredCapacity")]
        public Output<int?> DesiredCapacity { get; private set; } = null!;

        /// <summary>
        /// The time period during which a failed scheduled task is retried. Unit: seconds. Valid values: 0 to 21600. Default value: 600
        /// </summary>
        [Output("launchExpirationTime")]
        public Output<int?> LaunchExpirationTime { get; private set; } = null!;

        /// <summary>
        /// The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. 
        /// The time must be in UTC. You cannot enter a time point later than 90 days from the date of scheduled task creation.
        /// If the `recurrence_type` parameter is specified, the task is executed repeatedly at the time specified by LaunchTime.
        /// Otherwise, the task is only executed once at the date and time specified by LaunchTime.
        /// </summary>
        [Output("launchTime")]
        public Output<string?> LaunchTime { get; private set; } = null!;

        /// <summary>
        /// The maximum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
        /// </summary>
        [Output("maxValue")]
        public Output<int?> MaxValue { get; private set; } = null!;

        /// <summary>
        /// The minimum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
        /// </summary>
        [Output("minValue")]
        public Output<int?> MinValue { get; private set; } = null!;

        /// <summary>
        /// Specifies the end time after which the scheduled task is no longer repeated. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. 
        /// The time must be in UTC. You cannot enter a time point later than 365 days from the date of scheduled task creation. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time.
        /// </summary>
        [Output("recurrenceEndTime")]
        public Output<string> RecurrenceEndTime { get; private set; } = null!;

        /// <summary>
        /// Specifies the recurrence type of the scheduled task. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. Valid values:
        /// - Daily: The scheduled task is executed once every specified number of days.
        /// - Weekly: The scheduled task is executed on each specified day of a week.
        /// - Monthly: The scheduled task is executed on each specified day of a month.
        /// - Cron: (Available in 1.60.0+) The scheduled task is executed based on the specified cron expression.
        /// </summary>
        [Output("recurrenceType")]
        public Output<string> RecurrenceType { get; private set; } = null!;

        /// <summary>
        /// Specifies how often a scheduled task recurs. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. The valid value depends on `recurrence_type`
        /// - Daily: You can enter one value. Valid values: 1 to 31.
        /// - Weekly: You can enter multiple values and separate them with commas (,). For example, the values 0 to 6 correspond to the days of the week in sequence from Sunday to Saturday.
        /// - Monthly: You can enter two values in A-B format. Valid values of A and B: 1 to 31. The value of B must be greater than or equal to the value of A.
        /// - Cron: You can enter a cron expression which is written in UTC and consists of five fields: minute, hour, day of month (date), month, and day of week. The expression can contain wildcard characters including commas (,), question marks (?), hyphens (-), asterisks (*), number signs (#), forward slashes (/), and the L and W letters.
        /// </summary>
        [Output("recurrenceValue")]
        public Output<string> RecurrenceValue { get; private set; } = null!;

        /// <summary>
        /// The ID of the scaling group where the number of instances is modified when the scheduled task is triggered. After the `ScalingGroupId` parameter is specified, the scaling method of the scheduled task is to specify the number of instances in a scaling group. You must specify at least one of the following parameters: `MinValue`, `MaxValue`, and `DesiredCapacity`. **NOTE:** You cannot specify `scheduled_action` and `scaling_group_id` at the same time.
        /// </summary>
        [Output("scalingGroupId")]
        public Output<string> ScalingGroupId { get; private set; } = null!;

        /// <summary>
        /// The operation to be performed when a scheduled task is triggered. Enter the unique identifier of a scaling rule. **NOTE:** You cannot specify `scheduled_action` and `scaling_group_id` at the same time.
        /// </summary>
        [Output("scheduledAction")]
        public Output<string?> ScheduledAction { get; private set; } = null!;

        /// <summary>
        /// Display name of the scheduled task, which must be 2-40 characters (English or Chinese) long.
        /// </summary>
        [Output("scheduledTaskName")]
        public Output<string?> ScheduledTaskName { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to start the scheduled task. Default to true.
        /// </summary>
        [Output("taskEnabled")]
        public Output<bool?> TaskEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a ScheduledTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScheduledTask(string name, ScheduledTaskArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:ess/scheduledTask:ScheduledTask", name, args ?? new ScheduledTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScheduledTask(string name, Input<string> id, ScheduledTaskState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ess/scheduledTask:ScheduledTask", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScheduledTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScheduledTask Get(string name, Input<string> id, ScheduledTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new ScheduledTask(name, id, state, options);
        }
    }

    public sealed class ScheduledTaskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the scheduled task, which is 2-200 characters (English or Chinese) long.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expected number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group. **NOTE:** You must specify the `DesiredCapacity` parameter when you create the scaling group.
        /// </summary>
        [Input("desiredCapacity")]
        public Input<int>? DesiredCapacity { get; set; }

        /// <summary>
        /// The time period during which a failed scheduled task is retried. Unit: seconds. Valid values: 0 to 21600. Default value: 600
        /// </summary>
        [Input("launchExpirationTime")]
        public Input<int>? LaunchExpirationTime { get; set; }

        /// <summary>
        /// The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. 
        /// The time must be in UTC. You cannot enter a time point later than 90 days from the date of scheduled task creation.
        /// If the `recurrence_type` parameter is specified, the task is executed repeatedly at the time specified by LaunchTime.
        /// Otherwise, the task is only executed once at the date and time specified by LaunchTime.
        /// </summary>
        [Input("launchTime")]
        public Input<string>? LaunchTime { get; set; }

        /// <summary>
        /// The maximum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
        /// </summary>
        [Input("maxValue")]
        public Input<int>? MaxValue { get; set; }

        /// <summary>
        /// The minimum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
        /// </summary>
        [Input("minValue")]
        public Input<int>? MinValue { get; set; }

        /// <summary>
        /// Specifies the end time after which the scheduled task is no longer repeated. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. 
        /// The time must be in UTC. You cannot enter a time point later than 365 days from the date of scheduled task creation. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time.
        /// </summary>
        [Input("recurrenceEndTime")]
        public Input<string>? RecurrenceEndTime { get; set; }

        /// <summary>
        /// Specifies the recurrence type of the scheduled task. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. Valid values:
        /// - Daily: The scheduled task is executed once every specified number of days.
        /// - Weekly: The scheduled task is executed on each specified day of a week.
        /// - Monthly: The scheduled task is executed on each specified day of a month.
        /// - Cron: (Available in 1.60.0+) The scheduled task is executed based on the specified cron expression.
        /// </summary>
        [Input("recurrenceType")]
        public Input<string>? RecurrenceType { get; set; }

        /// <summary>
        /// Specifies how often a scheduled task recurs. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. The valid value depends on `recurrence_type`
        /// - Daily: You can enter one value. Valid values: 1 to 31.
        /// - Weekly: You can enter multiple values and separate them with commas (,). For example, the values 0 to 6 correspond to the days of the week in sequence from Sunday to Saturday.
        /// - Monthly: You can enter two values in A-B format. Valid values of A and B: 1 to 31. The value of B must be greater than or equal to the value of A.
        /// - Cron: You can enter a cron expression which is written in UTC and consists of five fields: minute, hour, day of month (date), month, and day of week. The expression can contain wildcard characters including commas (,), question marks (?), hyphens (-), asterisks (*), number signs (#), forward slashes (/), and the L and W letters.
        /// </summary>
        [Input("recurrenceValue")]
        public Input<string>? RecurrenceValue { get; set; }

        /// <summary>
        /// The ID of the scaling group where the number of instances is modified when the scheduled task is triggered. After the `ScalingGroupId` parameter is specified, the scaling method of the scheduled task is to specify the number of instances in a scaling group. You must specify at least one of the following parameters: `MinValue`, `MaxValue`, and `DesiredCapacity`. **NOTE:** You cannot specify `scheduled_action` and `scaling_group_id` at the same time.
        /// </summary>
        [Input("scalingGroupId")]
        public Input<string>? ScalingGroupId { get; set; }

        /// <summary>
        /// The operation to be performed when a scheduled task is triggered. Enter the unique identifier of a scaling rule. **NOTE:** You cannot specify `scheduled_action` and `scaling_group_id` at the same time.
        /// </summary>
        [Input("scheduledAction")]
        public Input<string>? ScheduledAction { get; set; }

        /// <summary>
        /// Display name of the scheduled task, which must be 2-40 characters (English or Chinese) long.
        /// </summary>
        [Input("scheduledTaskName")]
        public Input<string>? ScheduledTaskName { get; set; }

        /// <summary>
        /// Specifies whether to start the scheduled task. Default to true.
        /// </summary>
        [Input("taskEnabled")]
        public Input<bool>? TaskEnabled { get; set; }

        public ScheduledTaskArgs()
        {
        }
        public static new ScheduledTaskArgs Empty => new ScheduledTaskArgs();
    }

    public sealed class ScheduledTaskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the scheduled task, which is 2-200 characters (English or Chinese) long.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expected number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group. **NOTE:** You must specify the `DesiredCapacity` parameter when you create the scaling group.
        /// </summary>
        [Input("desiredCapacity")]
        public Input<int>? DesiredCapacity { get; set; }

        /// <summary>
        /// The time period during which a failed scheduled task is retried. Unit: seconds. Valid values: 0 to 21600. Default value: 600
        /// </summary>
        [Input("launchExpirationTime")]
        public Input<int>? LaunchExpirationTime { get; set; }

        /// <summary>
        /// The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. 
        /// The time must be in UTC. You cannot enter a time point later than 90 days from the date of scheduled task creation.
        /// If the `recurrence_type` parameter is specified, the task is executed repeatedly at the time specified by LaunchTime.
        /// Otherwise, the task is only executed once at the date and time specified by LaunchTime.
        /// </summary>
        [Input("launchTime")]
        public Input<string>? LaunchTime { get; set; }

        /// <summary>
        /// The maximum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
        /// </summary>
        [Input("maxValue")]
        public Input<int>? MaxValue { get; set; }

        /// <summary>
        /// The minimum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
        /// </summary>
        [Input("minValue")]
        public Input<int>? MinValue { get; set; }

        /// <summary>
        /// Specifies the end time after which the scheduled task is no longer repeated. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. 
        /// The time must be in UTC. You cannot enter a time point later than 365 days from the date of scheduled task creation. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time.
        /// </summary>
        [Input("recurrenceEndTime")]
        public Input<string>? RecurrenceEndTime { get; set; }

        /// <summary>
        /// Specifies the recurrence type of the scheduled task. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. Valid values:
        /// - Daily: The scheduled task is executed once every specified number of days.
        /// - Weekly: The scheduled task is executed on each specified day of a week.
        /// - Monthly: The scheduled task is executed on each specified day of a month.
        /// - Cron: (Available in 1.60.0+) The scheduled task is executed based on the specified cron expression.
        /// </summary>
        [Input("recurrenceType")]
        public Input<string>? RecurrenceType { get; set; }

        /// <summary>
        /// Specifies how often a scheduled task recurs. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. The valid value depends on `recurrence_type`
        /// - Daily: You can enter one value. Valid values: 1 to 31.
        /// - Weekly: You can enter multiple values and separate them with commas (,). For example, the values 0 to 6 correspond to the days of the week in sequence from Sunday to Saturday.
        /// - Monthly: You can enter two values in A-B format. Valid values of A and B: 1 to 31. The value of B must be greater than or equal to the value of A.
        /// - Cron: You can enter a cron expression which is written in UTC and consists of five fields: minute, hour, day of month (date), month, and day of week. The expression can contain wildcard characters including commas (,), question marks (?), hyphens (-), asterisks (*), number signs (#), forward slashes (/), and the L and W letters.
        /// </summary>
        [Input("recurrenceValue")]
        public Input<string>? RecurrenceValue { get; set; }

        /// <summary>
        /// The ID of the scaling group where the number of instances is modified when the scheduled task is triggered. After the `ScalingGroupId` parameter is specified, the scaling method of the scheduled task is to specify the number of instances in a scaling group. You must specify at least one of the following parameters: `MinValue`, `MaxValue`, and `DesiredCapacity`. **NOTE:** You cannot specify `scheduled_action` and `scaling_group_id` at the same time.
        /// </summary>
        [Input("scalingGroupId")]
        public Input<string>? ScalingGroupId { get; set; }

        /// <summary>
        /// The operation to be performed when a scheduled task is triggered. Enter the unique identifier of a scaling rule. **NOTE:** You cannot specify `scheduled_action` and `scaling_group_id` at the same time.
        /// </summary>
        [Input("scheduledAction")]
        public Input<string>? ScheduledAction { get; set; }

        /// <summary>
        /// Display name of the scheduled task, which must be 2-40 characters (English or Chinese) long.
        /// </summary>
        [Input("scheduledTaskName")]
        public Input<string>? ScheduledTaskName { get; set; }

        /// <summary>
        /// Specifies whether to start the scheduled task. Default to true.
        /// </summary>
        [Input("taskEnabled")]
        public Input<bool>? TaskEnabled { get; set; }

        public ScheduledTaskState()
        {
        }
        public static new ScheduledTaskState Empty => new ScheduledTaskState();
    }
}
