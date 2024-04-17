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

__all__ = ['DbInstancePlanArgs', 'DbInstancePlan']

@pulumi.input_type
class DbInstancePlanArgs:
    def __init__(__self__, *,
                 db_instance_id: pulumi.Input[str],
                 db_instance_plan_name: pulumi.Input[str],
                 plan_configs: pulumi.Input[Sequence[pulumi.Input['DbInstancePlanPlanConfigArgs']]],
                 plan_schedule_type: pulumi.Input[str],
                 plan_type: pulumi.Input[str],
                 plan_desc: Optional[pulumi.Input[str]] = None,
                 plan_end_date: Optional[pulumi.Input[str]] = None,
                 plan_start_date: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DbInstancePlan resource.
        :param pulumi.Input[str] db_instance_id: The ID of the GPDB instance.
        :param pulumi.Input[str] db_instance_plan_name: The name of the Plan.
        :param pulumi.Input[Sequence[pulumi.Input['DbInstancePlanPlanConfigArgs']]] plan_configs: The execution information of the plan. See `plan_config` below.
        :param pulumi.Input[str] plan_schedule_type: The execution mode of the plan. Valid values: `Postpone`, `Regular`.
        :param pulumi.Input[str] plan_type: The type of the Plan. Valid values: `PauseResume`, `Resize`.
        :param pulumi.Input[str] plan_desc: The description of the Plan.
        :param pulumi.Input[str] plan_end_date: The end time of the Plan.
        :param pulumi.Input[str] plan_start_date: The start time of the Plan.
        :param pulumi.Input[str] status: The Status of the Plan. Valid values: `active`, `cancel`.
        """
        pulumi.set(__self__, "db_instance_id", db_instance_id)
        pulumi.set(__self__, "db_instance_plan_name", db_instance_plan_name)
        pulumi.set(__self__, "plan_configs", plan_configs)
        pulumi.set(__self__, "plan_schedule_type", plan_schedule_type)
        pulumi.set(__self__, "plan_type", plan_type)
        if plan_desc is not None:
            pulumi.set(__self__, "plan_desc", plan_desc)
        if plan_end_date is not None:
            pulumi.set(__self__, "plan_end_date", plan_end_date)
        if plan_start_date is not None:
            pulumi.set(__self__, "plan_start_date", plan_start_date)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the GPDB instance.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter(name="dbInstancePlanName")
    def db_instance_plan_name(self) -> pulumi.Input[str]:
        """
        The name of the Plan.
        """
        return pulumi.get(self, "db_instance_plan_name")

    @db_instance_plan_name.setter
    def db_instance_plan_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_instance_plan_name", value)

    @property
    @pulumi.getter(name="planConfigs")
    def plan_configs(self) -> pulumi.Input[Sequence[pulumi.Input['DbInstancePlanPlanConfigArgs']]]:
        """
        The execution information of the plan. See `plan_config` below.
        """
        return pulumi.get(self, "plan_configs")

    @plan_configs.setter
    def plan_configs(self, value: pulumi.Input[Sequence[pulumi.Input['DbInstancePlanPlanConfigArgs']]]):
        pulumi.set(self, "plan_configs", value)

    @property
    @pulumi.getter(name="planScheduleType")
    def plan_schedule_type(self) -> pulumi.Input[str]:
        """
        The execution mode of the plan. Valid values: `Postpone`, `Regular`.
        """
        return pulumi.get(self, "plan_schedule_type")

    @plan_schedule_type.setter
    def plan_schedule_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "plan_schedule_type", value)

    @property
    @pulumi.getter(name="planType")
    def plan_type(self) -> pulumi.Input[str]:
        """
        The type of the Plan. Valid values: `PauseResume`, `Resize`.
        """
        return pulumi.get(self, "plan_type")

    @plan_type.setter
    def plan_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "plan_type", value)

    @property
    @pulumi.getter(name="planDesc")
    def plan_desc(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Plan.
        """
        return pulumi.get(self, "plan_desc")

    @plan_desc.setter
    def plan_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_desc", value)

    @property
    @pulumi.getter(name="planEndDate")
    def plan_end_date(self) -> Optional[pulumi.Input[str]]:
        """
        The end time of the Plan.
        """
        return pulumi.get(self, "plan_end_date")

    @plan_end_date.setter
    def plan_end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_end_date", value)

    @property
    @pulumi.getter(name="planStartDate")
    def plan_start_date(self) -> Optional[pulumi.Input[str]]:
        """
        The start time of the Plan.
        """
        return pulumi.get(self, "plan_start_date")

    @plan_start_date.setter
    def plan_start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_start_date", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The Status of the Plan. Valid values: `active`, `cancel`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _DbInstancePlanState:
    def __init__(__self__, *,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 db_instance_plan_name: Optional[pulumi.Input[str]] = None,
                 plan_configs: Optional[pulumi.Input[Sequence[pulumi.Input['DbInstancePlanPlanConfigArgs']]]] = None,
                 plan_desc: Optional[pulumi.Input[str]] = None,
                 plan_end_date: Optional[pulumi.Input[str]] = None,
                 plan_id: Optional[pulumi.Input[str]] = None,
                 plan_schedule_type: Optional[pulumi.Input[str]] = None,
                 plan_start_date: Optional[pulumi.Input[str]] = None,
                 plan_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DbInstancePlan resources.
        :param pulumi.Input[str] db_instance_id: The ID of the GPDB instance.
        :param pulumi.Input[str] db_instance_plan_name: The name of the Plan.
        :param pulumi.Input[Sequence[pulumi.Input['DbInstancePlanPlanConfigArgs']]] plan_configs: The execution information of the plan. See `plan_config` below.
        :param pulumi.Input[str] plan_desc: The description of the Plan.
        :param pulumi.Input[str] plan_end_date: The end time of the Plan.
        :param pulumi.Input[str] plan_id: The ID of the plan.
        :param pulumi.Input[str] plan_schedule_type: The execution mode of the plan. Valid values: `Postpone`, `Regular`.
        :param pulumi.Input[str] plan_start_date: The start time of the Plan.
        :param pulumi.Input[str] plan_type: The type of the Plan. Valid values: `PauseResume`, `Resize`.
        :param pulumi.Input[str] status: The Status of the Plan. Valid values: `active`, `cancel`.
        """
        if db_instance_id is not None:
            pulumi.set(__self__, "db_instance_id", db_instance_id)
        if db_instance_plan_name is not None:
            pulumi.set(__self__, "db_instance_plan_name", db_instance_plan_name)
        if plan_configs is not None:
            pulumi.set(__self__, "plan_configs", plan_configs)
        if plan_desc is not None:
            pulumi.set(__self__, "plan_desc", plan_desc)
        if plan_end_date is not None:
            pulumi.set(__self__, "plan_end_date", plan_end_date)
        if plan_id is not None:
            pulumi.set(__self__, "plan_id", plan_id)
        if plan_schedule_type is not None:
            pulumi.set(__self__, "plan_schedule_type", plan_schedule_type)
        if plan_start_date is not None:
            pulumi.set(__self__, "plan_start_date", plan_start_date)
        if plan_type is not None:
            pulumi.set(__self__, "plan_type", plan_type)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the GPDB instance.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter(name="dbInstancePlanName")
    def db_instance_plan_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Plan.
        """
        return pulumi.get(self, "db_instance_plan_name")

    @db_instance_plan_name.setter
    def db_instance_plan_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_plan_name", value)

    @property
    @pulumi.getter(name="planConfigs")
    def plan_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DbInstancePlanPlanConfigArgs']]]]:
        """
        The execution information of the plan. See `plan_config` below.
        """
        return pulumi.get(self, "plan_configs")

    @plan_configs.setter
    def plan_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DbInstancePlanPlanConfigArgs']]]]):
        pulumi.set(self, "plan_configs", value)

    @property
    @pulumi.getter(name="planDesc")
    def plan_desc(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Plan.
        """
        return pulumi.get(self, "plan_desc")

    @plan_desc.setter
    def plan_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_desc", value)

    @property
    @pulumi.getter(name="planEndDate")
    def plan_end_date(self) -> Optional[pulumi.Input[str]]:
        """
        The end time of the Plan.
        """
        return pulumi.get(self, "plan_end_date")

    @plan_end_date.setter
    def plan_end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_end_date", value)

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the plan.
        """
        return pulumi.get(self, "plan_id")

    @plan_id.setter
    def plan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_id", value)

    @property
    @pulumi.getter(name="planScheduleType")
    def plan_schedule_type(self) -> Optional[pulumi.Input[str]]:
        """
        The execution mode of the plan. Valid values: `Postpone`, `Regular`.
        """
        return pulumi.get(self, "plan_schedule_type")

    @plan_schedule_type.setter
    def plan_schedule_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_schedule_type", value)

    @property
    @pulumi.getter(name="planStartDate")
    def plan_start_date(self) -> Optional[pulumi.Input[str]]:
        """
        The start time of the Plan.
        """
        return pulumi.get(self, "plan_start_date")

    @plan_start_date.setter
    def plan_start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_start_date", value)

    @property
    @pulumi.getter(name="planType")
    def plan_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the Plan. Valid values: `PauseResume`, `Resize`.
        """
        return pulumi.get(self, "plan_type")

    @plan_type.setter
    def plan_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The Status of the Plan. Valid values: `active`, `cancel`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class DbInstancePlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 db_instance_plan_name: Optional[pulumi.Input[str]] = None,
                 plan_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DbInstancePlanPlanConfigArgs']]]]] = None,
                 plan_desc: Optional[pulumi.Input[str]] = None,
                 plan_end_date: Optional[pulumi.Input[str]] = None,
                 plan_schedule_type: Optional[pulumi.Input[str]] = None,
                 plan_start_date: Optional[pulumi.Input[str]] = None,
                 plan_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a AnalyticDB for PostgreSQL (GPDB) DB Instance Plan resource.

        For information about AnalyticDB for PostgreSQL (GPDB) DB Instance Plan and how to use it, see [What is DB Instance Plan](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/developer-reference/api-gpdb-2016-05-03-createdbinstanceplan).

        > **NOTE:** Available since v1.189.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.gpdb.get_zones()
        default_get_networks = alicloud.vpc.get_networks(name_regex="^default-NODELETING$")
        default_get_switches = alicloud.vpc.get_switches(vpc_id=default_get_networks.ids[0],
            zone_id=default.ids[0])
        default_instance = alicloud.gpdb.Instance("default",
            db_instance_category="HighAvailability",
            db_instance_class="gpdb.group.segsdx1",
            db_instance_mode="StorageElastic",
            description=name,
            engine="gpdb",
            engine_version="6.0",
            zone_id=default.ids[0],
            instance_network_type="VPC",
            instance_spec="2C16G",
            master_node_num=1,
            payment_type="PayAsYouGo",
            private_ip_address="1.1.1.1",
            seg_storage_type="cloud_essd",
            seg_node_num=4,
            storage_size=50,
            vpc_id=default_get_networks.ids[0],
            vswitch_id=default_get_switches.ids[0],
            ip_whitelists=[alicloud.gpdb.InstanceIpWhitelistArgs(
                security_ip_list="127.0.0.1",
            )])
        default_db_instance_plan = alicloud.gpdb.DbInstancePlan("default",
            db_instance_plan_name=name,
            plan_desc=name,
            plan_type="PauseResume",
            plan_schedule_type="Regular",
            plan_configs=[alicloud.gpdb.DbInstancePlanPlanConfigArgs(
                resume=alicloud.gpdb.DbInstancePlanPlanConfigResumeArgs(
                    plan_cron_time="0 0 0 1/1 * ? ",
                ),
                pause=alicloud.gpdb.DbInstancePlanPlanConfigPauseArgs(
                    plan_cron_time="0 0 10 1/1 * ? ",
                ),
            )],
            db_instance_id=default_instance.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GPDB DB Instance Plan can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:gpdb/dbInstancePlan:DbInstancePlan example <db_instance_id>:<plan_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_instance_id: The ID of the GPDB instance.
        :param pulumi.Input[str] db_instance_plan_name: The name of the Plan.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DbInstancePlanPlanConfigArgs']]]] plan_configs: The execution information of the plan. See `plan_config` below.
        :param pulumi.Input[str] plan_desc: The description of the Plan.
        :param pulumi.Input[str] plan_end_date: The end time of the Plan.
        :param pulumi.Input[str] plan_schedule_type: The execution mode of the plan. Valid values: `Postpone`, `Regular`.
        :param pulumi.Input[str] plan_start_date: The start time of the Plan.
        :param pulumi.Input[str] plan_type: The type of the Plan. Valid values: `PauseResume`, `Resize`.
        :param pulumi.Input[str] status: The Status of the Plan. Valid values: `active`, `cancel`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DbInstancePlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a AnalyticDB for PostgreSQL (GPDB) DB Instance Plan resource.

        For information about AnalyticDB for PostgreSQL (GPDB) DB Instance Plan and how to use it, see [What is DB Instance Plan](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/developer-reference/api-gpdb-2016-05-03-createdbinstanceplan).

        > **NOTE:** Available since v1.189.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.gpdb.get_zones()
        default_get_networks = alicloud.vpc.get_networks(name_regex="^default-NODELETING$")
        default_get_switches = alicloud.vpc.get_switches(vpc_id=default_get_networks.ids[0],
            zone_id=default.ids[0])
        default_instance = alicloud.gpdb.Instance("default",
            db_instance_category="HighAvailability",
            db_instance_class="gpdb.group.segsdx1",
            db_instance_mode="StorageElastic",
            description=name,
            engine="gpdb",
            engine_version="6.0",
            zone_id=default.ids[0],
            instance_network_type="VPC",
            instance_spec="2C16G",
            master_node_num=1,
            payment_type="PayAsYouGo",
            private_ip_address="1.1.1.1",
            seg_storage_type="cloud_essd",
            seg_node_num=4,
            storage_size=50,
            vpc_id=default_get_networks.ids[0],
            vswitch_id=default_get_switches.ids[0],
            ip_whitelists=[alicloud.gpdb.InstanceIpWhitelistArgs(
                security_ip_list="127.0.0.1",
            )])
        default_db_instance_plan = alicloud.gpdb.DbInstancePlan("default",
            db_instance_plan_name=name,
            plan_desc=name,
            plan_type="PauseResume",
            plan_schedule_type="Regular",
            plan_configs=[alicloud.gpdb.DbInstancePlanPlanConfigArgs(
                resume=alicloud.gpdb.DbInstancePlanPlanConfigResumeArgs(
                    plan_cron_time="0 0 0 1/1 * ? ",
                ),
                pause=alicloud.gpdb.DbInstancePlanPlanConfigPauseArgs(
                    plan_cron_time="0 0 10 1/1 * ? ",
                ),
            )],
            db_instance_id=default_instance.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GPDB DB Instance Plan can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:gpdb/dbInstancePlan:DbInstancePlan example <db_instance_id>:<plan_id>
        ```

        :param str resource_name: The name of the resource.
        :param DbInstancePlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DbInstancePlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 db_instance_plan_name: Optional[pulumi.Input[str]] = None,
                 plan_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DbInstancePlanPlanConfigArgs']]]]] = None,
                 plan_desc: Optional[pulumi.Input[str]] = None,
                 plan_end_date: Optional[pulumi.Input[str]] = None,
                 plan_schedule_type: Optional[pulumi.Input[str]] = None,
                 plan_start_date: Optional[pulumi.Input[str]] = None,
                 plan_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DbInstancePlanArgs.__new__(DbInstancePlanArgs)

            if db_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_instance_id'")
            __props__.__dict__["db_instance_id"] = db_instance_id
            if db_instance_plan_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_instance_plan_name'")
            __props__.__dict__["db_instance_plan_name"] = db_instance_plan_name
            if plan_configs is None and not opts.urn:
                raise TypeError("Missing required property 'plan_configs'")
            __props__.__dict__["plan_configs"] = plan_configs
            __props__.__dict__["plan_desc"] = plan_desc
            __props__.__dict__["plan_end_date"] = plan_end_date
            if plan_schedule_type is None and not opts.urn:
                raise TypeError("Missing required property 'plan_schedule_type'")
            __props__.__dict__["plan_schedule_type"] = plan_schedule_type
            __props__.__dict__["plan_start_date"] = plan_start_date
            if plan_type is None and not opts.urn:
                raise TypeError("Missing required property 'plan_type'")
            __props__.__dict__["plan_type"] = plan_type
            __props__.__dict__["status"] = status
            __props__.__dict__["plan_id"] = None
        super(DbInstancePlan, __self__).__init__(
            'alicloud:gpdb/dbInstancePlan:DbInstancePlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            db_instance_id: Optional[pulumi.Input[str]] = None,
            db_instance_plan_name: Optional[pulumi.Input[str]] = None,
            plan_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DbInstancePlanPlanConfigArgs']]]]] = None,
            plan_desc: Optional[pulumi.Input[str]] = None,
            plan_end_date: Optional[pulumi.Input[str]] = None,
            plan_id: Optional[pulumi.Input[str]] = None,
            plan_schedule_type: Optional[pulumi.Input[str]] = None,
            plan_start_date: Optional[pulumi.Input[str]] = None,
            plan_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'DbInstancePlan':
        """
        Get an existing DbInstancePlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_instance_id: The ID of the GPDB instance.
        :param pulumi.Input[str] db_instance_plan_name: The name of the Plan.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DbInstancePlanPlanConfigArgs']]]] plan_configs: The execution information of the plan. See `plan_config` below.
        :param pulumi.Input[str] plan_desc: The description of the Plan.
        :param pulumi.Input[str] plan_end_date: The end time of the Plan.
        :param pulumi.Input[str] plan_id: The ID of the plan.
        :param pulumi.Input[str] plan_schedule_type: The execution mode of the plan. Valid values: `Postpone`, `Regular`.
        :param pulumi.Input[str] plan_start_date: The start time of the Plan.
        :param pulumi.Input[str] plan_type: The type of the Plan. Valid values: `PauseResume`, `Resize`.
        :param pulumi.Input[str] status: The Status of the Plan. Valid values: `active`, `cancel`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DbInstancePlanState.__new__(_DbInstancePlanState)

        __props__.__dict__["db_instance_id"] = db_instance_id
        __props__.__dict__["db_instance_plan_name"] = db_instance_plan_name
        __props__.__dict__["plan_configs"] = plan_configs
        __props__.__dict__["plan_desc"] = plan_desc
        __props__.__dict__["plan_end_date"] = plan_end_date
        __props__.__dict__["plan_id"] = plan_id
        __props__.__dict__["plan_schedule_type"] = plan_schedule_type
        __props__.__dict__["plan_start_date"] = plan_start_date
        __props__.__dict__["plan_type"] = plan_type
        __props__.__dict__["status"] = status
        return DbInstancePlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the GPDB instance.
        """
        return pulumi.get(self, "db_instance_id")

    @property
    @pulumi.getter(name="dbInstancePlanName")
    def db_instance_plan_name(self) -> pulumi.Output[str]:
        """
        The name of the Plan.
        """
        return pulumi.get(self, "db_instance_plan_name")

    @property
    @pulumi.getter(name="planConfigs")
    def plan_configs(self) -> pulumi.Output[Sequence['outputs.DbInstancePlanPlanConfig']]:
        """
        The execution information of the plan. See `plan_config` below.
        """
        return pulumi.get(self, "plan_configs")

    @property
    @pulumi.getter(name="planDesc")
    def plan_desc(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the Plan.
        """
        return pulumi.get(self, "plan_desc")

    @property
    @pulumi.getter(name="planEndDate")
    def plan_end_date(self) -> pulumi.Output[Optional[str]]:
        """
        The end time of the Plan.
        """
        return pulumi.get(self, "plan_end_date")

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> pulumi.Output[str]:
        """
        The ID of the plan.
        """
        return pulumi.get(self, "plan_id")

    @property
    @pulumi.getter(name="planScheduleType")
    def plan_schedule_type(self) -> pulumi.Output[str]:
        """
        The execution mode of the plan. Valid values: `Postpone`, `Regular`.
        """
        return pulumi.get(self, "plan_schedule_type")

    @property
    @pulumi.getter(name="planStartDate")
    def plan_start_date(self) -> pulumi.Output[str]:
        """
        The start time of the Plan.
        """
        return pulumi.get(self, "plan_start_date")

    @property
    @pulumi.getter(name="planType")
    def plan_type(self) -> pulumi.Output[str]:
        """
        The type of the Plan. Valid values: `PauseResume`, `Resize`.
        """
        return pulumi.get(self, "plan_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The Status of the Plan. Valid values: `active`, `cancel`.
        """
        return pulumi.get(self, "status")

