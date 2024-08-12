# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ServerBackupPlanArgs', 'ServerBackupPlan']

@pulumi.input_type
class ServerBackupPlanArgs:
    def __init__(__self__, *,
                 details: pulumi.Input[Sequence[pulumi.Input['ServerBackupPlanDetailArgs']]],
                 ecs_server_backup_plan_name: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 retention: pulumi.Input[int],
                 schedule: pulumi.Input[str],
                 cross_account_role_name: Optional[pulumi.Input[str]] = None,
                 cross_account_type: Optional[pulumi.Input[str]] = None,
                 cross_account_user_id: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ServerBackupPlan resource.
        :param pulumi.Input[Sequence[pulumi.Input['ServerBackupPlanDetailArgs']]] details: ECS server backup plan details.
        :param pulumi.Input[str] ecs_server_backup_plan_name: The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        :param pulumi.Input[str] instance_id: The ID of ECS instance.
        :param pulumi.Input[int] retention: Backup retention days, the minimum is 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: `I|{startTime}|{interval}`
               * `startTime` Backup start time, UNIX time, in seconds.
        :param pulumi.Input[str] cross_account_role_name: The role name created in the original account RAM backup by the cross account managed by the current account.
        :param pulumi.Input[str] cross_account_type: The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
        :param pulumi.Input[int] cross_account_user_id: The original account ID of the cross account backup managed by the current account.
        :param pulumi.Input[bool] disabled: Whether to disable the backup task. Valid values: `true`, `false`.
        """
        pulumi.set(__self__, "details", details)
        pulumi.set(__self__, "ecs_server_backup_plan_name", ecs_server_backup_plan_name)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "retention", retention)
        pulumi.set(__self__, "schedule", schedule)
        if cross_account_role_name is not None:
            pulumi.set(__self__, "cross_account_role_name", cross_account_role_name)
        if cross_account_type is not None:
            pulumi.set(__self__, "cross_account_type", cross_account_type)
        if cross_account_user_id is not None:
            pulumi.set(__self__, "cross_account_user_id", cross_account_user_id)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter
    def details(self) -> pulumi.Input[Sequence[pulumi.Input['ServerBackupPlanDetailArgs']]]:
        """
        ECS server backup plan details.
        """
        return pulumi.get(self, "details")

    @details.setter
    def details(self, value: pulumi.Input[Sequence[pulumi.Input['ServerBackupPlanDetailArgs']]]):
        pulumi.set(self, "details", value)

    @property
    @pulumi.getter(name="ecsServerBackupPlanName")
    def ecs_server_backup_plan_name(self) -> pulumi.Input[str]:
        """
        The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        """
        return pulumi.get(self, "ecs_server_backup_plan_name")

    @ecs_server_backup_plan_name.setter
    def ecs_server_backup_plan_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "ecs_server_backup_plan_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of ECS instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def retention(self) -> pulumi.Input[int]:
        """
        Backup retention days, the minimum is 1.
        """
        return pulumi.get(self, "retention")

    @retention.setter
    def retention(self, value: pulumi.Input[int]):
        pulumi.set(self, "retention", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input[str]:
        """
        Backup strategy. Optional format: `I|{startTime}|{interval}`
        * `startTime` Backup start time, UNIX time, in seconds.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="crossAccountRoleName")
    def cross_account_role_name(self) -> Optional[pulumi.Input[str]]:
        """
        The role name created in the original account RAM backup by the cross account managed by the current account.
        """
        return pulumi.get(self, "cross_account_role_name")

    @cross_account_role_name.setter
    def cross_account_role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cross_account_role_name", value)

    @property
    @pulumi.getter(name="crossAccountType")
    def cross_account_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
        """
        return pulumi.get(self, "cross_account_type")

    @cross_account_type.setter
    def cross_account_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cross_account_type", value)

    @property
    @pulumi.getter(name="crossAccountUserId")
    def cross_account_user_id(self) -> Optional[pulumi.Input[int]]:
        """
        The original account ID of the cross account backup managed by the current account.
        """
        return pulumi.get(self, "cross_account_user_id")

    @cross_account_user_id.setter
    def cross_account_user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cross_account_user_id", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to disable the backup task. Valid values: `true`, `false`.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)


@pulumi.input_type
class _ServerBackupPlanState:
    def __init__(__self__, *,
                 cross_account_role_name: Optional[pulumi.Input[str]] = None,
                 cross_account_type: Optional[pulumi.Input[str]] = None,
                 cross_account_user_id: Optional[pulumi.Input[int]] = None,
                 details: Optional[pulumi.Input[Sequence[pulumi.Input['ServerBackupPlanDetailArgs']]]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 ecs_server_backup_plan_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 retention: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServerBackupPlan resources.
        :param pulumi.Input[str] cross_account_role_name: The role name created in the original account RAM backup by the cross account managed by the current account.
        :param pulumi.Input[str] cross_account_type: The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
        :param pulumi.Input[int] cross_account_user_id: The original account ID of the cross account backup managed by the current account.
        :param pulumi.Input[Sequence[pulumi.Input['ServerBackupPlanDetailArgs']]] details: ECS server backup plan details.
        :param pulumi.Input[bool] disabled: Whether to disable the backup task. Valid values: `true`, `false`.
        :param pulumi.Input[str] ecs_server_backup_plan_name: The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        :param pulumi.Input[str] instance_id: The ID of ECS instance.
        :param pulumi.Input[int] retention: Backup retention days, the minimum is 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: `I|{startTime}|{interval}`
               * `startTime` Backup start time, UNIX time, in seconds.
        """
        if cross_account_role_name is not None:
            pulumi.set(__self__, "cross_account_role_name", cross_account_role_name)
        if cross_account_type is not None:
            pulumi.set(__self__, "cross_account_type", cross_account_type)
        if cross_account_user_id is not None:
            pulumi.set(__self__, "cross_account_user_id", cross_account_user_id)
        if details is not None:
            pulumi.set(__self__, "details", details)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if ecs_server_backup_plan_name is not None:
            pulumi.set(__self__, "ecs_server_backup_plan_name", ecs_server_backup_plan_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if retention is not None:
            pulumi.set(__self__, "retention", retention)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)

    @property
    @pulumi.getter(name="crossAccountRoleName")
    def cross_account_role_name(self) -> Optional[pulumi.Input[str]]:
        """
        The role name created in the original account RAM backup by the cross account managed by the current account.
        """
        return pulumi.get(self, "cross_account_role_name")

    @cross_account_role_name.setter
    def cross_account_role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cross_account_role_name", value)

    @property
    @pulumi.getter(name="crossAccountType")
    def cross_account_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
        """
        return pulumi.get(self, "cross_account_type")

    @cross_account_type.setter
    def cross_account_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cross_account_type", value)

    @property
    @pulumi.getter(name="crossAccountUserId")
    def cross_account_user_id(self) -> Optional[pulumi.Input[int]]:
        """
        The original account ID of the cross account backup managed by the current account.
        """
        return pulumi.get(self, "cross_account_user_id")

    @cross_account_user_id.setter
    def cross_account_user_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cross_account_user_id", value)

    @property
    @pulumi.getter
    def details(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerBackupPlanDetailArgs']]]]:
        """
        ECS server backup plan details.
        """
        return pulumi.get(self, "details")

    @details.setter
    def details(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerBackupPlanDetailArgs']]]]):
        pulumi.set(self, "details", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to disable the backup task. Valid values: `true`, `false`.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="ecsServerBackupPlanName")
    def ecs_server_backup_plan_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        """
        return pulumi.get(self, "ecs_server_backup_plan_name")

    @ecs_server_backup_plan_name.setter
    def ecs_server_backup_plan_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ecs_server_backup_plan_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of ECS instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def retention(self) -> Optional[pulumi.Input[int]]:
        """
        Backup retention days, the minimum is 1.
        """
        return pulumi.get(self, "retention")

    @retention.setter
    def retention(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        Backup strategy. Optional format: `I|{startTime}|{interval}`
        * `startTime` Backup start time, UNIX time, in seconds.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)


class ServerBackupPlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cross_account_role_name: Optional[pulumi.Input[str]] = None,
                 cross_account_type: Optional[pulumi.Input[str]] = None,
                 cross_account_user_id: Optional[pulumi.Input[int]] = None,
                 details: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServerBackupPlanDetailArgs', 'ServerBackupPlanDetailArgsDict']]]]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 ecs_server_backup_plan_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 retention: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Hybrid Backup Recovery (HBR) Server Backup Plan resource.

        For information about Hybrid Backup Recovery (HBR) Server Backup Plan and how to use it, see [What is Server Backup Plan](https://www.alibabacloud.com/help/doc-detail/211140.htm).

        > **NOTE:** Available in v1.142.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.get_zones(available_resource_creation="Instance")
        example_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=example.zones[0].id,
            cpu_core_count=1,
            memory_size=2)
        example_get_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            owners="system")
        example_network = alicloud.vpc.Network("example",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        example_switch = alicloud.vpc.Switch("example",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=example_network.id,
            zone_id=example.zones[0].id)
        example_security_group = alicloud.ecs.SecurityGroup("example",
            name="terraform-example",
            vpc_id=example_network.id)
        example_instance = alicloud.ecs.Instance("example",
            image_id=example_get_images.images[0].id,
            instance_type=example_get_instance_types.instance_types[0].id,
            availability_zone=example.zones[0].id,
            security_groups=[example_security_group.id],
            instance_name="terraform-example",
            internet_charge_type="PayByBandwidth",
            vswitch_id=example_switch.id)
        example_server_backup_plan = alicloud.hbr.ServerBackupPlan("example",
            ecs_server_backup_plan_name="terraform-example",
            instance_id=example_instance.id,
            schedule="I|1602673264|PT2H",
            retention=1,
            details=[{
                "app_consistent": True,
                "snapshot_group": True,
            }],
            disabled=False)
        ```

        ## Import

        Hybrid Backup Recovery (HBR) Server Backup Plan can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:hbr/serverBackupPlan:ServerBackupPlan example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cross_account_role_name: The role name created in the original account RAM backup by the cross account managed by the current account.
        :param pulumi.Input[str] cross_account_type: The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
        :param pulumi.Input[int] cross_account_user_id: The original account ID of the cross account backup managed by the current account.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ServerBackupPlanDetailArgs', 'ServerBackupPlanDetailArgsDict']]]] details: ECS server backup plan details.
        :param pulumi.Input[bool] disabled: Whether to disable the backup task. Valid values: `true`, `false`.
        :param pulumi.Input[str] ecs_server_backup_plan_name: The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        :param pulumi.Input[str] instance_id: The ID of ECS instance.
        :param pulumi.Input[int] retention: Backup retention days, the minimum is 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: `I|{startTime}|{interval}`
               * `startTime` Backup start time, UNIX time, in seconds.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerBackupPlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Hybrid Backup Recovery (HBR) Server Backup Plan resource.

        For information about Hybrid Backup Recovery (HBR) Server Backup Plan and how to use it, see [What is Server Backup Plan](https://www.alibabacloud.com/help/doc-detail/211140.htm).

        > **NOTE:** Available in v1.142.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.get_zones(available_resource_creation="Instance")
        example_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=example.zones[0].id,
            cpu_core_count=1,
            memory_size=2)
        example_get_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            owners="system")
        example_network = alicloud.vpc.Network("example",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        example_switch = alicloud.vpc.Switch("example",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=example_network.id,
            zone_id=example.zones[0].id)
        example_security_group = alicloud.ecs.SecurityGroup("example",
            name="terraform-example",
            vpc_id=example_network.id)
        example_instance = alicloud.ecs.Instance("example",
            image_id=example_get_images.images[0].id,
            instance_type=example_get_instance_types.instance_types[0].id,
            availability_zone=example.zones[0].id,
            security_groups=[example_security_group.id],
            instance_name="terraform-example",
            internet_charge_type="PayByBandwidth",
            vswitch_id=example_switch.id)
        example_server_backup_plan = alicloud.hbr.ServerBackupPlan("example",
            ecs_server_backup_plan_name="terraform-example",
            instance_id=example_instance.id,
            schedule="I|1602673264|PT2H",
            retention=1,
            details=[{
                "app_consistent": True,
                "snapshot_group": True,
            }],
            disabled=False)
        ```

        ## Import

        Hybrid Backup Recovery (HBR) Server Backup Plan can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:hbr/serverBackupPlan:ServerBackupPlan example <id>
        ```

        :param str resource_name: The name of the resource.
        :param ServerBackupPlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerBackupPlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cross_account_role_name: Optional[pulumi.Input[str]] = None,
                 cross_account_type: Optional[pulumi.Input[str]] = None,
                 cross_account_user_id: Optional[pulumi.Input[int]] = None,
                 details: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServerBackupPlanDetailArgs', 'ServerBackupPlanDetailArgsDict']]]]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 ecs_server_backup_plan_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 retention: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerBackupPlanArgs.__new__(ServerBackupPlanArgs)

            __props__.__dict__["cross_account_role_name"] = cross_account_role_name
            __props__.__dict__["cross_account_type"] = cross_account_type
            __props__.__dict__["cross_account_user_id"] = cross_account_user_id
            if details is None and not opts.urn:
                raise TypeError("Missing required property 'details'")
            __props__.__dict__["details"] = details
            __props__.__dict__["disabled"] = disabled
            if ecs_server_backup_plan_name is None and not opts.urn:
                raise TypeError("Missing required property 'ecs_server_backup_plan_name'")
            __props__.__dict__["ecs_server_backup_plan_name"] = ecs_server_backup_plan_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if retention is None and not opts.urn:
                raise TypeError("Missing required property 'retention'")
            __props__.__dict__["retention"] = retention
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
        super(ServerBackupPlan, __self__).__init__(
            'alicloud:hbr/serverBackupPlan:ServerBackupPlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cross_account_role_name: Optional[pulumi.Input[str]] = None,
            cross_account_type: Optional[pulumi.Input[str]] = None,
            cross_account_user_id: Optional[pulumi.Input[int]] = None,
            details: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServerBackupPlanDetailArgs', 'ServerBackupPlanDetailArgsDict']]]]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            ecs_server_backup_plan_name: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            retention: Optional[pulumi.Input[int]] = None,
            schedule: Optional[pulumi.Input[str]] = None) -> 'ServerBackupPlan':
        """
        Get an existing ServerBackupPlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cross_account_role_name: The role name created in the original account RAM backup by the cross account managed by the current account.
        :param pulumi.Input[str] cross_account_type: The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
        :param pulumi.Input[int] cross_account_user_id: The original account ID of the cross account backup managed by the current account.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ServerBackupPlanDetailArgs', 'ServerBackupPlanDetailArgsDict']]]] details: ECS server backup plan details.
        :param pulumi.Input[bool] disabled: Whether to disable the backup task. Valid values: `true`, `false`.
        :param pulumi.Input[str] ecs_server_backup_plan_name: The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        :param pulumi.Input[str] instance_id: The ID of ECS instance.
        :param pulumi.Input[int] retention: Backup retention days, the minimum is 1.
        :param pulumi.Input[str] schedule: Backup strategy. Optional format: `I|{startTime}|{interval}`
               * `startTime` Backup start time, UNIX time, in seconds.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerBackupPlanState.__new__(_ServerBackupPlanState)

        __props__.__dict__["cross_account_role_name"] = cross_account_role_name
        __props__.__dict__["cross_account_type"] = cross_account_type
        __props__.__dict__["cross_account_user_id"] = cross_account_user_id
        __props__.__dict__["details"] = details
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["ecs_server_backup_plan_name"] = ecs_server_backup_plan_name
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["retention"] = retention
        __props__.__dict__["schedule"] = schedule
        return ServerBackupPlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="crossAccountRoleName")
    def cross_account_role_name(self) -> pulumi.Output[Optional[str]]:
        """
        The role name created in the original account RAM backup by the cross account managed by the current account.
        """
        return pulumi.get(self, "cross_account_role_name")

    @property
    @pulumi.getter(name="crossAccountType")
    def cross_account_type(self) -> pulumi.Output[str]:
        """
        The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
        """
        return pulumi.get(self, "cross_account_type")

    @property
    @pulumi.getter(name="crossAccountUserId")
    def cross_account_user_id(self) -> pulumi.Output[Optional[int]]:
        """
        The original account ID of the cross account backup managed by the current account.
        """
        return pulumi.get(self, "cross_account_user_id")

    @property
    @pulumi.getter
    def details(self) -> pulumi.Output[Sequence['outputs.ServerBackupPlanDetail']]:
        """
        ECS server backup plan details.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Whether to disable the backup task. Valid values: `true`, `false`.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="ecsServerBackupPlanName")
    def ecs_server_backup_plan_name(self) -> pulumi.Output[str]:
        """
        The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
        """
        return pulumi.get(self, "ecs_server_backup_plan_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of ECS instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def retention(self) -> pulumi.Output[int]:
        """
        Backup retention days, the minimum is 1.
        """
        return pulumi.get(self, "retention")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[str]:
        """
        Backup strategy. Optional format: `I|{startTime}|{interval}`
        * `startTime` Backup start time, UNIX time, in seconds.
        """
        return pulumi.get(self, "schedule")

