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

__all__ = ['ScheduledSqlArgs', 'ScheduledSql']

@pulumi.input_type
class ScheduledSqlArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 project: pulumi.Input[str],
                 schedule: pulumi.Input['ScheduledSqlScheduleArgs'],
                 scheduled_sql_configuration: pulumi.Input['ScheduledSqlScheduledSqlConfigurationArgs'],
                 scheduled_sql_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ScheduledSql resource.
        :param pulumi.Input[str] display_name: Task Display Name.
        :param pulumi.Input[str] project: Log project.
        :param pulumi.Input['ScheduledSqlScheduleArgs'] schedule: The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
        :param pulumi.Input['ScheduledSqlScheduledSqlConfigurationArgs'] scheduled_sql_configuration: Task Configuration. See `scheduled_sql_configuration` below.
        :param pulumi.Input[str] scheduled_sql_name: Timed SQL name.
        :param pulumi.Input[str] description: Task Description.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "scheduled_sql_configuration", scheduled_sql_configuration)
        pulumi.set(__self__, "scheduled_sql_name", scheduled_sql_name)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        Task Display Name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        Log project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input['ScheduledSqlScheduleArgs']:
        """
        The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input['ScheduledSqlScheduleArgs']):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="scheduledSqlConfiguration")
    def scheduled_sql_configuration(self) -> pulumi.Input['ScheduledSqlScheduledSqlConfigurationArgs']:
        """
        Task Configuration. See `scheduled_sql_configuration` below.
        """
        return pulumi.get(self, "scheduled_sql_configuration")

    @scheduled_sql_configuration.setter
    def scheduled_sql_configuration(self, value: pulumi.Input['ScheduledSqlScheduledSqlConfigurationArgs']):
        pulumi.set(self, "scheduled_sql_configuration", value)

    @property
    @pulumi.getter(name="scheduledSqlName")
    def scheduled_sql_name(self) -> pulumi.Input[str]:
        """
        Timed SQL name.
        """
        return pulumi.get(self, "scheduled_sql_name")

    @scheduled_sql_name.setter
    def scheduled_sql_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "scheduled_sql_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Task Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _ScheduledSqlState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input['ScheduledSqlScheduleArgs']] = None,
                 scheduled_sql_configuration: Optional[pulumi.Input['ScheduledSqlScheduledSqlConfigurationArgs']] = None,
                 scheduled_sql_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ScheduledSql resources.
        :param pulumi.Input[str] description: Task Description.
        :param pulumi.Input[str] display_name: Task Display Name.
        :param pulumi.Input[str] project: Log project.
        :param pulumi.Input['ScheduledSqlScheduleArgs'] schedule: The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
        :param pulumi.Input['ScheduledSqlScheduledSqlConfigurationArgs'] scheduled_sql_configuration: Task Configuration. See `scheduled_sql_configuration` below.
        :param pulumi.Input[str] scheduled_sql_name: Timed SQL name.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if scheduled_sql_configuration is not None:
            pulumi.set(__self__, "scheduled_sql_configuration", scheduled_sql_configuration)
        if scheduled_sql_name is not None:
            pulumi.set(__self__, "scheduled_sql_name", scheduled_sql_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Task Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Task Display Name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        Log project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input['ScheduledSqlScheduleArgs']]:
        """
        The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input['ScheduledSqlScheduleArgs']]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="scheduledSqlConfiguration")
    def scheduled_sql_configuration(self) -> Optional[pulumi.Input['ScheduledSqlScheduledSqlConfigurationArgs']]:
        """
        Task Configuration. See `scheduled_sql_configuration` below.
        """
        return pulumi.get(self, "scheduled_sql_configuration")

    @scheduled_sql_configuration.setter
    def scheduled_sql_configuration(self, value: Optional[pulumi.Input['ScheduledSqlScheduledSqlConfigurationArgs']]):
        pulumi.set(self, "scheduled_sql_configuration", value)

    @property
    @pulumi.getter(name="scheduledSqlName")
    def scheduled_sql_name(self) -> Optional[pulumi.Input[str]]:
        """
        Timed SQL name.
        """
        return pulumi.get(self, "scheduled_sql_name")

    @scheduled_sql_name.setter
    def scheduled_sql_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scheduled_sql_name", value)


class ScheduledSql(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[Union['ScheduledSqlScheduleArgs', 'ScheduledSqlScheduleArgsDict']]] = None,
                 scheduled_sql_configuration: Optional[pulumi.Input[Union['ScheduledSqlScheduledSqlConfigurationArgs', 'ScheduledSqlScheduledSqlConfigurationArgsDict']]] = None,
                 scheduled_sql_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a SLS Scheduled SQL resource. Scheduled SQL task.

        For information about SLS Scheduled SQL and how to use it, see [What is Scheduled SQL](https://www.alibabacloud.com/help/zh/sls/developer-reference/api-sls-2020-12-30-createscheduledsql).

        > **NOTE:** Available since v1.224.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = random.index.Integer("default",
            min=10000,
            max=99999)
        default_k_ie4_kv = alicloud.log.Project("defaultKIe4KV",
            description=f"{name}-{default['result']}",
            project_name=f"{name}-{default['result']}")
        default1_li9we = alicloud.log.Store("default1LI9we",
            hot_ttl=8,
            retention_period=30,
            shard_count=2,
            project_name=default_k_ie4_kv.project_name,
            logstore_name=f"{name}-{default['result']}")
        default_scheduled_sql = alicloud.sls.ScheduledSql("default",
            description="example-tf-scheduled-sql-0006",
            schedule={
                "type": "Cron",
                "time_zone": "+0700",
                "delay": 20,
                "cron_expression": "0 0/1 * * *",
            },
            display_name="example-tf-scheduled-sql-0006",
            scheduled_sql_configuration={
                "script": "* | select * from log",
                "sql_type": "searchQuery",
                "dest_endpoint": "ap-northeast-1.log.aliyuncs.com",
                "dest_project": "job-e2e-project-jj78kur-ap-southeast-1",
                "source_logstore": default1_li9we.logstore_name,
                "dest_logstore": "example-open-api02",
                "role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
                "dest_role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
                "from_time_expr": "@m-1m",
                "to_time_expr": "@m",
                "max_run_time_in_seconds": 1800,
                "resource_pool": "enhanced",
                "max_retries": 5,
                "from_time": 1713196800,
                "to_time": 0,
                "data_format": "log2log",
            },
            scheduled_sql_name=name,
            project=default_k_ie4_kv.project_name)
        ```

        ## Import

        SLS Scheduled SQL can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:sls/scheduledSql:ScheduledSql example <project>:<scheduled_sql_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Task Description.
        :param pulumi.Input[str] display_name: Task Display Name.
        :param pulumi.Input[str] project: Log project.
        :param pulumi.Input[Union['ScheduledSqlScheduleArgs', 'ScheduledSqlScheduleArgsDict']] schedule: The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
        :param pulumi.Input[Union['ScheduledSqlScheduledSqlConfigurationArgs', 'ScheduledSqlScheduledSqlConfigurationArgsDict']] scheduled_sql_configuration: Task Configuration. See `scheduled_sql_configuration` below.
        :param pulumi.Input[str] scheduled_sql_name: Timed SQL name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScheduledSqlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a SLS Scheduled SQL resource. Scheduled SQL task.

        For information about SLS Scheduled SQL and how to use it, see [What is Scheduled SQL](https://www.alibabacloud.com/help/zh/sls/developer-reference/api-sls-2020-12-30-createscheduledsql).

        > **NOTE:** Available since v1.224.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = random.index.Integer("default",
            min=10000,
            max=99999)
        default_k_ie4_kv = alicloud.log.Project("defaultKIe4KV",
            description=f"{name}-{default['result']}",
            project_name=f"{name}-{default['result']}")
        default1_li9we = alicloud.log.Store("default1LI9we",
            hot_ttl=8,
            retention_period=30,
            shard_count=2,
            project_name=default_k_ie4_kv.project_name,
            logstore_name=f"{name}-{default['result']}")
        default_scheduled_sql = alicloud.sls.ScheduledSql("default",
            description="example-tf-scheduled-sql-0006",
            schedule={
                "type": "Cron",
                "time_zone": "+0700",
                "delay": 20,
                "cron_expression": "0 0/1 * * *",
            },
            display_name="example-tf-scheduled-sql-0006",
            scheduled_sql_configuration={
                "script": "* | select * from log",
                "sql_type": "searchQuery",
                "dest_endpoint": "ap-northeast-1.log.aliyuncs.com",
                "dest_project": "job-e2e-project-jj78kur-ap-southeast-1",
                "source_logstore": default1_li9we.logstore_name,
                "dest_logstore": "example-open-api02",
                "role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
                "dest_role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
                "from_time_expr": "@m-1m",
                "to_time_expr": "@m",
                "max_run_time_in_seconds": 1800,
                "resource_pool": "enhanced",
                "max_retries": 5,
                "from_time": 1713196800,
                "to_time": 0,
                "data_format": "log2log",
            },
            scheduled_sql_name=name,
            project=default_k_ie4_kv.project_name)
        ```

        ## Import

        SLS Scheduled SQL can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:sls/scheduledSql:ScheduledSql example <project>:<scheduled_sql_name>
        ```

        :param str resource_name: The name of the resource.
        :param ScheduledSqlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScheduledSqlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[Union['ScheduledSqlScheduleArgs', 'ScheduledSqlScheduleArgsDict']]] = None,
                 scheduled_sql_configuration: Optional[pulumi.Input[Union['ScheduledSqlScheduledSqlConfigurationArgs', 'ScheduledSqlScheduledSqlConfigurationArgsDict']]] = None,
                 scheduled_sql_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScheduledSqlArgs.__new__(ScheduledSqlArgs)

            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
            if scheduled_sql_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'scheduled_sql_configuration'")
            __props__.__dict__["scheduled_sql_configuration"] = scheduled_sql_configuration
            if scheduled_sql_name is None and not opts.urn:
                raise TypeError("Missing required property 'scheduled_sql_name'")
            __props__.__dict__["scheduled_sql_name"] = scheduled_sql_name
        super(ScheduledSql, __self__).__init__(
            'alicloud:sls/scheduledSql:ScheduledSql',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            schedule: Optional[pulumi.Input[Union['ScheduledSqlScheduleArgs', 'ScheduledSqlScheduleArgsDict']]] = None,
            scheduled_sql_configuration: Optional[pulumi.Input[Union['ScheduledSqlScheduledSqlConfigurationArgs', 'ScheduledSqlScheduledSqlConfigurationArgsDict']]] = None,
            scheduled_sql_name: Optional[pulumi.Input[str]] = None) -> 'ScheduledSql':
        """
        Get an existing ScheduledSql resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Task Description.
        :param pulumi.Input[str] display_name: Task Display Name.
        :param pulumi.Input[str] project: Log project.
        :param pulumi.Input[Union['ScheduledSqlScheduleArgs', 'ScheduledSqlScheduleArgsDict']] schedule: The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
        :param pulumi.Input[Union['ScheduledSqlScheduledSqlConfigurationArgs', 'ScheduledSqlScheduledSqlConfigurationArgsDict']] scheduled_sql_configuration: Task Configuration. See `scheduled_sql_configuration` below.
        :param pulumi.Input[str] scheduled_sql_name: Timed SQL name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScheduledSqlState.__new__(_ScheduledSqlState)

        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["project"] = project
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["scheduled_sql_configuration"] = scheduled_sql_configuration
        __props__.__dict__["scheduled_sql_name"] = scheduled_sql_name
        return ScheduledSql(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Task Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Task Display Name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        Log project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output['outputs.ScheduledSqlSchedule']:
        """
        The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="scheduledSqlConfiguration")
    def scheduled_sql_configuration(self) -> pulumi.Output['outputs.ScheduledSqlScheduledSqlConfiguration']:
        """
        Task Configuration. See `scheduled_sql_configuration` below.
        """
        return pulumi.get(self, "scheduled_sql_configuration")

    @property
    @pulumi.getter(name="scheduledSqlName")
    def scheduled_sql_name(self) -> pulumi.Output[str]:
        """
        Timed SQL name.
        """
        return pulumi.get(self, "scheduled_sql_name")

