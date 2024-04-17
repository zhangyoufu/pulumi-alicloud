# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NotificationArgs', 'Notification']

@pulumi.input_type
class NotificationArgs:
    def __init__(__self__, *,
                 notification_arn: pulumi.Input[str],
                 notification_types: pulumi.Input[Sequence[pulumi.Input[str]]],
                 scaling_group_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a Notification resource.
        :param pulumi.Input[str] notification_arn: The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
               * region: the region ID of the scaling group. For more information, see `Regions and zones`
               * account-id: the ID of your account.
               * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_types: The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
        :param pulumi.Input[str] scaling_group_id: The ID of the Auto Scaling group.
        """
        pulumi.set(__self__, "notification_arn", notification_arn)
        pulumi.set(__self__, "notification_types", notification_types)
        pulumi.set(__self__, "scaling_group_id", scaling_group_id)

    @property
    @pulumi.getter(name="notificationArn")
    def notification_arn(self) -> pulumi.Input[str]:
        """
        The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
        * region: the region ID of the scaling group. For more information, see `Regions and zones`
        * account-id: the ID of your account.
        * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
        """
        return pulumi.get(self, "notification_arn")

    @notification_arn.setter
    def notification_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "notification_arn", value)

    @property
    @pulumi.getter(name="notificationTypes")
    def notification_types(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
        """
        return pulumi.get(self, "notification_types")

    @notification_types.setter
    def notification_types(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "notification_types", value)

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Input[str]:
        """
        The ID of the Auto Scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @scaling_group_id.setter
    def scaling_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scaling_group_id", value)


@pulumi.input_type
class _NotificationState:
    def __init__(__self__, *,
                 notification_arn: Optional[pulumi.Input[str]] = None,
                 notification_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Notification resources.
        :param pulumi.Input[str] notification_arn: The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
               * region: the region ID of the scaling group. For more information, see `Regions and zones`
               * account-id: the ID of your account.
               * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_types: The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
        :param pulumi.Input[str] scaling_group_id: The ID of the Auto Scaling group.
        """
        if notification_arn is not None:
            pulumi.set(__self__, "notification_arn", notification_arn)
        if notification_types is not None:
            pulumi.set(__self__, "notification_types", notification_types)
        if scaling_group_id is not None:
            pulumi.set(__self__, "scaling_group_id", scaling_group_id)

    @property
    @pulumi.getter(name="notificationArn")
    def notification_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
        * region: the region ID of the scaling group. For more information, see `Regions and zones`
        * account-id: the ID of your account.
        * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
        """
        return pulumi.get(self, "notification_arn")

    @notification_arn.setter
    def notification_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_arn", value)

    @property
    @pulumi.getter(name="notificationTypes")
    def notification_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
        """
        return pulumi.get(self, "notification_types")

    @notification_types.setter
    def notification_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notification_types", value)

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Auto Scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @scaling_group_id.setter
    def scaling_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_group_id", value)


class Notification(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 notification_arn: Optional[pulumi.Input[str]] = None,
                 notification_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ESS notification resource. More about Ess notification, see [Autoscaling Notification](https://www.alibabacloud.com/help/doc-detail/71114.htm).

        > **NOTE:** Available since v1.55.0.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_integer = random.index.Integer("default",
            min=10000,
            max=99999)
        my_name = f"{name}-{default_integer['result']}"
        default = alicloud.get_regions(current=True)
        default_get_account = alicloud.get_account()
        default_get_zones = alicloud.get_zones(available_disk_category="cloud_efficiency",
            available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("default",
            vpc_name=my_name,
            cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            zone_id=default_get_zones.zones[0].id,
            vswitch_name=my_name)
        default_scaling_group = alicloud.ess.ScalingGroup("default",
            min_size=1,
            max_size=1,
            scaling_group_name=my_name,
            removal_policies=[
                "OldestInstance",
                "NewestInstance",
            ],
            vswitch_ids=[default_switch.id])
        default_queue = alicloud.mns.Queue("default", name=my_name)
        default_notification = alicloud.ess.Notification("default",
            scaling_group_id=default_scaling_group.id,
            notification_types=[
                "AUTOSCALING:SCALE_OUT_SUCCESS",
                "AUTOSCALING:SCALE_OUT_ERROR",
            ],
            notification_arn=default_queue.name.apply(lambda name: f"acs:ess:{default.regions[0].id}:{default_get_account.id}:queue/{name}"))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Ess notification can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ess/notification:Notification example 'scaling_group_id:notification_arn'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] notification_arn: The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
               * region: the region ID of the scaling group. For more information, see `Regions and zones`
               * account-id: the ID of your account.
               * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_types: The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
        :param pulumi.Input[str] scaling_group_id: The ID of the Auto Scaling group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NotificationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ESS notification resource. More about Ess notification, see [Autoscaling Notification](https://www.alibabacloud.com/help/doc-detail/71114.htm).

        > **NOTE:** Available since v1.55.0.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_integer = random.index.Integer("default",
            min=10000,
            max=99999)
        my_name = f"{name}-{default_integer['result']}"
        default = alicloud.get_regions(current=True)
        default_get_account = alicloud.get_account()
        default_get_zones = alicloud.get_zones(available_disk_category="cloud_efficiency",
            available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("default",
            vpc_name=my_name,
            cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            zone_id=default_get_zones.zones[0].id,
            vswitch_name=my_name)
        default_scaling_group = alicloud.ess.ScalingGroup("default",
            min_size=1,
            max_size=1,
            scaling_group_name=my_name,
            removal_policies=[
                "OldestInstance",
                "NewestInstance",
            ],
            vswitch_ids=[default_switch.id])
        default_queue = alicloud.mns.Queue("default", name=my_name)
        default_notification = alicloud.ess.Notification("default",
            scaling_group_id=default_scaling_group.id,
            notification_types=[
                "AUTOSCALING:SCALE_OUT_SUCCESS",
                "AUTOSCALING:SCALE_OUT_ERROR",
            ],
            notification_arn=default_queue.name.apply(lambda name: f"acs:ess:{default.regions[0].id}:{default_get_account.id}:queue/{name}"))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Ess notification can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ess/notification:Notification example 'scaling_group_id:notification_arn'
        ```

        :param str resource_name: The name of the resource.
        :param NotificationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NotificationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 notification_arn: Optional[pulumi.Input[str]] = None,
                 notification_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NotificationArgs.__new__(NotificationArgs)

            if notification_arn is None and not opts.urn:
                raise TypeError("Missing required property 'notification_arn'")
            __props__.__dict__["notification_arn"] = notification_arn
            if notification_types is None and not opts.urn:
                raise TypeError("Missing required property 'notification_types'")
            __props__.__dict__["notification_types"] = notification_types
            if scaling_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'scaling_group_id'")
            __props__.__dict__["scaling_group_id"] = scaling_group_id
        super(Notification, __self__).__init__(
            'alicloud:ess/notification:Notification',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            notification_arn: Optional[pulumi.Input[str]] = None,
            notification_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            scaling_group_id: Optional[pulumi.Input[str]] = None) -> 'Notification':
        """
        Get an existing Notification resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] notification_arn: The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
               * region: the region ID of the scaling group. For more information, see `Regions and zones`
               * account-id: the ID of your account.
               * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_types: The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
        :param pulumi.Input[str] scaling_group_id: The ID of the Auto Scaling group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NotificationState.__new__(_NotificationState)

        __props__.__dict__["notification_arn"] = notification_arn
        __props__.__dict__["notification_types"] = notification_types
        __props__.__dict__["scaling_group_id"] = scaling_group_id
        return Notification(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="notificationArn")
    def notification_arn(self) -> pulumi.Output[str]:
        """
        The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in `acs:ess:{region}:{account-id}:{resource-relative-id}` format.
        * region: the region ID of the scaling group. For more information, see `Regions and zones`
        * account-id: the ID of your account.
        * resource-relative-id: the notification method. Valid values : `cloudmonitor`, MNS queue: `queue/{queuename}`, Replace the queuename with the specific MNS queue name, MNS topic: `topic/{topicname}`, Replace the topicname with the specific MNS topic name.
        """
        return pulumi.get(self, "notification_arn")

    @property
    @pulumi.getter(name="notificationTypes")
    def notification_types(self) -> pulumi.Output[Sequence[str]]:
        """
        The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'.
        """
        return pulumi.get(self, "notification_types")

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the Auto Scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

