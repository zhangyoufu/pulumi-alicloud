# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AutoSnapShotPolicyArgs', 'AutoSnapShotPolicy']

@pulumi.input_type
class AutoSnapShotPolicyArgs:
    def __init__(__self__, *,
                 policy_name: pulumi.Input[str],
                 repeat_weekdays: pulumi.Input[Sequence[pulumi.Input[str]]],
                 retention_days: pulumi.Input[int],
                 time_points: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a AutoSnapShotPolicy resource.
        :param pulumi.Input[str] policy_name: Automatic snapshot policy name
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
        :param pulumi.Input[int] retention_days: Automatic snapshot retention days.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
        """
        pulumi.set(__self__, "policy_name", policy_name)
        pulumi.set(__self__, "repeat_weekdays", repeat_weekdays)
        pulumi.set(__self__, "retention_days", retention_days)
        pulumi.set(__self__, "time_points", time_points)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Input[str]:
        """
        Automatic snapshot policy name
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
        """
        return pulumi.get(self, "repeat_weekdays")

    @repeat_weekdays.setter
    def repeat_weekdays(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "repeat_weekdays", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Input[int]:
        """
        Automatic snapshot retention days.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: pulumi.Input[int]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter(name="timePoints")
    def time_points(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
        """
        return pulumi.get(self, "time_points")

    @time_points.setter
    def time_points(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "time_points", value)


@pulumi.input_type
class _AutoSnapShotPolicyState:
    def __init__(__self__, *,
                 applied_dbfs_number: Optional[pulumi.Input[int]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 last_modified: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 status_detail: Optional[pulumi.Input[str]] = None,
                 time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering AutoSnapShotPolicy resources.
        :param pulumi.Input[int] applied_dbfs_number: The number of database file systems set by the automatic snapshot policy.
        :param pulumi.Input[str] create_time: The creation time of the resource
        :param pulumi.Input[str] last_modified: Last modification time of automatic snapshot policy
        :param pulumi.Input[str] policy_id: Automatic snapshot policy ID
        :param pulumi.Input[str] policy_name: Automatic snapshot policy name
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
        :param pulumi.Input[int] retention_days: Automatic snapshot retention days.
        :param pulumi.Input[str] status: Automatic snapshot policy status
        :param pulumi.Input[str] status_detail: Automatic snapshot policy status details
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
        """
        if applied_dbfs_number is not None:
            pulumi.set(__self__, "applied_dbfs_number", applied_dbfs_number)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if last_modified is not None:
            pulumi.set(__self__, "last_modified", last_modified)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)
        if repeat_weekdays is not None:
            pulumi.set(__self__, "repeat_weekdays", repeat_weekdays)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if status_detail is not None:
            pulumi.set(__self__, "status_detail", status_detail)
        if time_points is not None:
            pulumi.set(__self__, "time_points", time_points)

    @property
    @pulumi.getter(name="appliedDbfsNumber")
    def applied_dbfs_number(self) -> Optional[pulumi.Input[int]]:
        """
        The number of database file systems set by the automatic snapshot policy.
        """
        return pulumi.get(self, "applied_dbfs_number")

    @applied_dbfs_number.setter
    def applied_dbfs_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "applied_dbfs_number", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the resource
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> Optional[pulumi.Input[str]]:
        """
        Last modification time of automatic snapshot policy
        """
        return pulumi.get(self, "last_modified")

    @last_modified.setter
    def last_modified(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_modified", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Automatic snapshot policy ID
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        Automatic snapshot policy name
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
        """
        return pulumi.get(self, "repeat_weekdays")

    @repeat_weekdays.setter
    def repeat_weekdays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "repeat_weekdays", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Automatic snapshot retention days.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Automatic snapshot policy status
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="statusDetail")
    def status_detail(self) -> Optional[pulumi.Input[str]]:
        """
        Automatic snapshot policy status details
        """
        return pulumi.get(self, "status_detail")

    @status_detail.setter
    def status_detail(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status_detail", value)

    @property
    @pulumi.getter(name="timePoints")
    def time_points(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
        """
        return pulumi.get(self, "time_points")

    @time_points.setter
    def time_points(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "time_points", value)


class AutoSnapShotPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a Dbfs Auto Snap Shot Policy resource.

        For information about Dbfs Auto Snap Shot Policy and how to use it.

        > **NOTE:** Available since v1.202.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.databasefilesystem.AutoSnapShotPolicy("default",
            policy_name="tf-example",
            repeat_weekdays=["2"],
            retention_days=1,
            time_points=["01"])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Dbfs Auto Snap Shot Policy can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:databasefilesystem/autoSnapShotPolicy:AutoSnapShotPolicy example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_name: Automatic snapshot policy name
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
        :param pulumi.Input[int] retention_days: Automatic snapshot retention days.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AutoSnapShotPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Dbfs Auto Snap Shot Policy resource.

        For information about Dbfs Auto Snap Shot Policy and how to use it.

        > **NOTE:** Available since v1.202.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.databasefilesystem.AutoSnapShotPolicy("default",
            policy_name="tf-example",
            repeat_weekdays=["2"],
            retention_days=1,
            time_points=["01"])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Dbfs Auto Snap Shot Policy can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:databasefilesystem/autoSnapShotPolicy:AutoSnapShotPolicy example <id>
        ```

        :param str resource_name: The name of the resource.
        :param AutoSnapShotPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutoSnapShotPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AutoSnapShotPolicyArgs.__new__(AutoSnapShotPolicyArgs)

            if policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'policy_name'")
            __props__.__dict__["policy_name"] = policy_name
            if repeat_weekdays is None and not opts.urn:
                raise TypeError("Missing required property 'repeat_weekdays'")
            __props__.__dict__["repeat_weekdays"] = repeat_weekdays
            if retention_days is None and not opts.urn:
                raise TypeError("Missing required property 'retention_days'")
            __props__.__dict__["retention_days"] = retention_days
            if time_points is None and not opts.urn:
                raise TypeError("Missing required property 'time_points'")
            __props__.__dict__["time_points"] = time_points
            __props__.__dict__["applied_dbfs_number"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["last_modified"] = None
            __props__.__dict__["policy_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["status_detail"] = None
        super(AutoSnapShotPolicy, __self__).__init__(
            'alicloud:databasefilesystem/autoSnapShotPolicy:AutoSnapShotPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            applied_dbfs_number: Optional[pulumi.Input[int]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            last_modified: Optional[pulumi.Input[str]] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            policy_name: Optional[pulumi.Input[str]] = None,
            repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            retention_days: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            status_detail: Optional[pulumi.Input[str]] = None,
            time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'AutoSnapShotPolicy':
        """
        Get an existing AutoSnapShotPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] applied_dbfs_number: The number of database file systems set by the automatic snapshot policy.
        :param pulumi.Input[str] create_time: The creation time of the resource
        :param pulumi.Input[str] last_modified: Last modification time of automatic snapshot policy
        :param pulumi.Input[str] policy_id: Automatic snapshot policy ID
        :param pulumi.Input[str] policy_name: Automatic snapshot policy name
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
        :param pulumi.Input[int] retention_days: Automatic snapshot retention days.
        :param pulumi.Input[str] status: Automatic snapshot policy status
        :param pulumi.Input[str] status_detail: Automatic snapshot policy status details
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AutoSnapShotPolicyState.__new__(_AutoSnapShotPolicyState)

        __props__.__dict__["applied_dbfs_number"] = applied_dbfs_number
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["last_modified"] = last_modified
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["policy_name"] = policy_name
        __props__.__dict__["repeat_weekdays"] = repeat_weekdays
        __props__.__dict__["retention_days"] = retention_days
        __props__.__dict__["status"] = status
        __props__.__dict__["status_detail"] = status_detail
        __props__.__dict__["time_points"] = time_points
        return AutoSnapShotPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appliedDbfsNumber")
    def applied_dbfs_number(self) -> pulumi.Output[int]:
        """
        The number of database file systems set by the automatic snapshot policy.
        """
        return pulumi.get(self, "applied_dbfs_number")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of the resource
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> pulumi.Output[str]:
        """
        Last modification time of automatic snapshot policy
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        Automatic snapshot policy ID
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[str]:
        """
        Automatic snapshot policy name
        """
        return pulumi.get(self, "policy_name")

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> pulumi.Output[Sequence[str]]:
        """
        A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
        """
        return pulumi.get(self, "repeat_weekdays")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[int]:
        """
        Automatic snapshot retention days.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Automatic snapshot policy status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusDetail")
    def status_detail(self) -> pulumi.Output[str]:
        """
        Automatic snapshot policy status details
        """
        return pulumi.get(self, "status_detail")

    @property
    @pulumi.getter(name="timePoints")
    def time_points(self) -> pulumi.Output[Sequence[str]]:
        """
        The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
        """
        return pulumi.get(self, "time_points")

