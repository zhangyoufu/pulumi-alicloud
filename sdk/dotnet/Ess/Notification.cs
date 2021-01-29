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
    /// Provides a ESS notification resource. More about Ess notification, see [Autoscaling Notification](https://www.alibabacloud.com/help/doc-detail/71114.htm).
    /// 
    /// &gt; **NOTE:** Available in 1.55.0+
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var name = config.Get("name") ?? "tf-testAccEssNotification-%d";
    ///         var defaultRegions = Output.Create(AliCloud.GetRegions.InvokeAsync(new AliCloud.GetRegionsArgs
    ///         {
    ///             Current = true,
    ///         }));
    ///         var defaultAccount = Output.Create(AliCloud.GetAccount.InvokeAsync());
    ///         var defaultZones = Output.Create(AliCloud.GetZones.InvokeAsync(new AliCloud.GetZonesArgs
    ///         {
    ///             AvailableDiskCategory = "cloud_efficiency",
    ///             AvailableResourceCreation = "VSwitch",
    ///         }));
    ///         var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/16",
    ///         });
    ///         var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new AliCloud.Vpc.SwitchArgs
    ///         {
    ///             VpcId = defaultNetwork.Id,
    ///             CidrBlock = "172.16.0.0/24",
    ///             AvailabilityZone = defaultZones.Apply(defaultZones =&gt; defaultZones.Zones[0].Id),
    ///         });
    ///         var defaultScalingGroup = new AliCloud.Ess.ScalingGroup("defaultScalingGroup", new AliCloud.Ess.ScalingGroupArgs
    ///         {
    ///             MinSize = 1,
    ///             MaxSize = 1,
    ///             ScalingGroupName = name,
    ///             RemovalPolicies = 
    ///             {
    ///                 "OldestInstance",
    ///                 "NewestInstance",
    ///             },
    ///             VswitchIds = 
    ///             {
    ///                 defaultSwitch.Id,
    ///             },
    ///         });
    ///         var defaultQueue = new AliCloud.Mns.Queue("defaultQueue", new AliCloud.Mns.QueueArgs
    ///         {
    ///         });
    ///         var defaultNotification = new AliCloud.Ess.Notification("defaultNotification", new AliCloud.Ess.NotificationArgs
    ///         {
    ///             ScalingGroupId = defaultScalingGroup.Id,
    ///             NotificationTypes = 
    ///             {
    ///                 "AUTOSCALING:SCALE_OUT_SUCCESS",
    ///                 "AUTOSCALING:SCALE_OUT_ERROR",
    ///             },
    ///             NotificationArn = Output.Tuple(defaultRegions, defaultAccount, defaultQueue.Name).Apply(values =&gt;
    ///             {
    ///                 var defaultRegions = values.Item1;
    ///                 var defaultAccount = values.Item2;
    ///                 var name = values.Item3;
    ///                 return $"acs:ess:{defaultRegions.Regions[0].Id}:{defaultAccount.Id}:queue/{name}";
    ///             }),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Ess notification can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ess/notification:Notification example 'scaling_group_id:notification_arn'
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ess/notification:Notification")]
    public partial class Notification : Pulumi.CustomResource
    {
        /// <summary>
        /// The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
        /// * region: the region ID of the scaling group. For more information, see `Regions and zones`
        /// * account-id: the ID of your account.
        /// * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
        /// </summary>
        [Output("notificationArn")]
        public Output<string> NotificationArn { get; private set; } = null!;

        /// <summary>
        /// The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
        /// </summary>
        [Output("notificationTypes")]
        public Output<ImmutableArray<string>> NotificationTypes { get; private set; } = null!;

        /// <summary>
        /// The ID of the Auto Scaling group.
        /// </summary>
        [Output("scalingGroupId")]
        public Output<string> ScalingGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a Notification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Notification(string name, NotificationArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ess/notification:Notification", name, args ?? new NotificationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Notification(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ess/notification:Notification", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Notification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Notification Get(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
        {
            return new Notification(name, id, state, options);
        }
    }

    public sealed class NotificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
        /// * region: the region ID of the scaling group. For more information, see `Regions and zones`
        /// * account-id: the ID of your account.
        /// * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
        /// </summary>
        [Input("notificationArn", required: true)]
        public Input<string> NotificationArn { get; set; } = null!;

        [Input("notificationTypes", required: true)]
        private InputList<string>? _notificationTypes;

        /// <summary>
        /// The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
        /// </summary>
        public InputList<string> NotificationTypes
        {
            get => _notificationTypes ?? (_notificationTypes = new InputList<string>());
            set => _notificationTypes = value;
        }

        /// <summary>
        /// The ID of the Auto Scaling group.
        /// </summary>
        [Input("scalingGroupId", required: true)]
        public Input<string> ScalingGroupId { get; set; } = null!;

        public NotificationArgs()
        {
        }
    }

    public sealed class NotificationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
        /// * region: the region ID of the scaling group. For more information, see `Regions and zones`
        /// * account-id: the ID of your account.
        /// * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
        /// </summary>
        [Input("notificationArn")]
        public Input<string>? NotificationArn { get; set; }

        [Input("notificationTypes")]
        private InputList<string>? _notificationTypes;

        /// <summary>
        /// The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
        /// </summary>
        public InputList<string> NotificationTypes
        {
            get => _notificationTypes ?? (_notificationTypes = new InputList<string>());
            set => _notificationTypes = value;
        }

        /// <summary>
        /// The ID of the Auto Scaling group.
        /// </summary>
        [Input("scalingGroupId")]
        public Input<string>? ScalingGroupId { get; set; }

        public NotificationState()
        {
        }
    }
}
