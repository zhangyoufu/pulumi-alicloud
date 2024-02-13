# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AutoSnapshotPolicyArgs', 'AutoSnapshotPolicy']

@pulumi.input_type
class AutoSnapshotPolicyArgs:
    def __init__(__self__, *,
                 repeat_weekdays: pulumi.Input[Sequence[pulumi.Input[str]]],
                 time_points: pulumi.Input[Sequence[pulumi.Input[str]]],
                 auto_snapshot_policy_name: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a AutoSnapshotPolicy resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: The day on which an auto snapshot is created.
               - A maximum of 7 time points can be selected.
               - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The point in time at which an auto snapshot is created.
               - A maximum of 24 time points can be selected.
               - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
        :param pulumi.Input[str] auto_snapshot_policy_name: The name of the automatic snapshot policy. Limits:
               - The name must be `2` to `128` characters in length,
               - The name must start with a letter.
               - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
               - The value of this parameter is empty by default.
        :param pulumi.Input[int] retention_days: The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
               - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
        """
        pulumi.set(__self__, "repeat_weekdays", repeat_weekdays)
        pulumi.set(__self__, "time_points", time_points)
        if auto_snapshot_policy_name is not None:
            pulumi.set(__self__, "auto_snapshot_policy_name", auto_snapshot_policy_name)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The day on which an auto snapshot is created.
        - A maximum of 7 time points can be selected.
        - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
        """
        return pulumi.get(self, "repeat_weekdays")

    @repeat_weekdays.setter
    def repeat_weekdays(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "repeat_weekdays", value)

    @property
    @pulumi.getter(name="timePoints")
    def time_points(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The point in time at which an auto snapshot is created.
        - A maximum of 24 time points can be selected.
        - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
        """
        return pulumi.get(self, "time_points")

    @time_points.setter
    def time_points(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "time_points", value)

    @property
    @pulumi.getter(name="autoSnapshotPolicyName")
    def auto_snapshot_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the automatic snapshot policy. Limits:
        - The name must be `2` to `128` characters in length,
        - The name must start with a letter.
        - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
        - The value of this parameter is empty by default.
        """
        return pulumi.get(self, "auto_snapshot_policy_name")

    @auto_snapshot_policy_name.setter
    def auto_snapshot_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_snapshot_policy_name", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
        - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)


@pulumi.input_type
class _AutoSnapshotPolicyState:
    def __init__(__self__, *,
                 auto_snapshot_policy_name: Optional[pulumi.Input[str]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering AutoSnapshotPolicy resources.
        :param pulumi.Input[str] auto_snapshot_policy_name: The name of the automatic snapshot policy. Limits:
               - The name must be `2` to `128` characters in length,
               - The name must start with a letter.
               - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
               - The value of this parameter is empty by default.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: The day on which an auto snapshot is created.
               - A maximum of 7 time points can be selected.
               - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
        :param pulumi.Input[int] retention_days: The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
               - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
        :param pulumi.Input[str] status: The status of the automatic snapshot policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The point in time at which an auto snapshot is created.
               - A maximum of 24 time points can be selected.
               - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
        """
        if auto_snapshot_policy_name is not None:
            pulumi.set(__self__, "auto_snapshot_policy_name", auto_snapshot_policy_name)
        if repeat_weekdays is not None:
            pulumi.set(__self__, "repeat_weekdays", repeat_weekdays)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if time_points is not None:
            pulumi.set(__self__, "time_points", time_points)

    @property
    @pulumi.getter(name="autoSnapshotPolicyName")
    def auto_snapshot_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the automatic snapshot policy. Limits:
        - The name must be `2` to `128` characters in length,
        - The name must start with a letter.
        - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
        - The value of this parameter is empty by default.
        """
        return pulumi.get(self, "auto_snapshot_policy_name")

    @auto_snapshot_policy_name.setter
    def auto_snapshot_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_snapshot_policy_name", value)

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The day on which an auto snapshot is created.
        - A maximum of 7 time points can be selected.
        - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
        """
        return pulumi.get(self, "repeat_weekdays")

    @repeat_weekdays.setter
    def repeat_weekdays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "repeat_weekdays", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
        - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the automatic snapshot policy.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="timePoints")
    def time_points(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The point in time at which an auto snapshot is created.
        - A maximum of 24 time points can be selected.
        - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
        """
        return pulumi.get(self, "time_points")

    @time_points.setter
    def time_points(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "time_points", value)


class AutoSnapshotPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_snapshot_policy_name: Optional[pulumi.Input[str]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a Network Attached Storage (NAS) Auto Snapshot Policy resource.

        For information about Network Attached Storage (NAS) Auto Snapshot Policy and how to use it, see [What is Auto Snapshot Policy](https://www.alibabacloud.com/help/en/doc-detail/135662.html).

        > **NOTE:** Available in v1.153.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.nas.AutoSnapshotPolicy("example",
            auto_snapshot_policy_name="example_value",
            repeat_weekdays=[
                "3",
                "4",
                "5",
            ],
            retention_days=30,
            time_points=[
                "3",
                "4",
                "5",
            ])
        ```

        ## Import

        Network Attached Storage (NAS) Auto Snapshot Policy can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:nas/autoSnapshotPolicy:AutoSnapshotPolicy example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_snapshot_policy_name: The name of the automatic snapshot policy. Limits:
               - The name must be `2` to `128` characters in length,
               - The name must start with a letter.
               - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
               - The value of this parameter is empty by default.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: The day on which an auto snapshot is created.
               - A maximum of 7 time points can be selected.
               - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
        :param pulumi.Input[int] retention_days: The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
               - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The point in time at which an auto snapshot is created.
               - A maximum of 24 time points can be selected.
               - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AutoSnapshotPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Network Attached Storage (NAS) Auto Snapshot Policy resource.

        For information about Network Attached Storage (NAS) Auto Snapshot Policy and how to use it, see [What is Auto Snapshot Policy](https://www.alibabacloud.com/help/en/doc-detail/135662.html).

        > **NOTE:** Available in v1.153.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.nas.AutoSnapshotPolicy("example",
            auto_snapshot_policy_name="example_value",
            repeat_weekdays=[
                "3",
                "4",
                "5",
            ],
            retention_days=30,
            time_points=[
                "3",
                "4",
                "5",
            ])
        ```

        ## Import

        Network Attached Storage (NAS) Auto Snapshot Policy can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:nas/autoSnapshotPolicy:AutoSnapshotPolicy example <id>
        ```

        :param str resource_name: The name of the resource.
        :param AutoSnapshotPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutoSnapshotPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_snapshot_policy_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = AutoSnapshotPolicyArgs.__new__(AutoSnapshotPolicyArgs)

            __props__.__dict__["auto_snapshot_policy_name"] = auto_snapshot_policy_name
            if repeat_weekdays is None and not opts.urn:
                raise TypeError("Missing required property 'repeat_weekdays'")
            __props__.__dict__["repeat_weekdays"] = repeat_weekdays
            __props__.__dict__["retention_days"] = retention_days
            if time_points is None and not opts.urn:
                raise TypeError("Missing required property 'time_points'")
            __props__.__dict__["time_points"] = time_points
            __props__.__dict__["status"] = None
        super(AutoSnapshotPolicy, __self__).__init__(
            'alicloud:nas/autoSnapshotPolicy:AutoSnapshotPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_snapshot_policy_name: Optional[pulumi.Input[str]] = None,
            repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            retention_days: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'AutoSnapshotPolicy':
        """
        Get an existing AutoSnapshotPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_snapshot_policy_name: The name of the automatic snapshot policy. Limits:
               - The name must be `2` to `128` characters in length,
               - The name must start with a letter.
               - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
               - The value of this parameter is empty by default.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: The day on which an auto snapshot is created.
               - A maximum of 7 time points can be selected.
               - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
        :param pulumi.Input[int] retention_days: The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
               - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
        :param pulumi.Input[str] status: The status of the automatic snapshot policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The point in time at which an auto snapshot is created.
               - A maximum of 24 time points can be selected.
               - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AutoSnapshotPolicyState.__new__(_AutoSnapshotPolicyState)

        __props__.__dict__["auto_snapshot_policy_name"] = auto_snapshot_policy_name
        __props__.__dict__["repeat_weekdays"] = repeat_weekdays
        __props__.__dict__["retention_days"] = retention_days
        __props__.__dict__["status"] = status
        __props__.__dict__["time_points"] = time_points
        return AutoSnapshotPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoSnapshotPolicyName")
    def auto_snapshot_policy_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the automatic snapshot policy. Limits:
        - The name must be `2` to `128` characters in length,
        - The name must start with a letter.
        - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
        - The value of this parameter is empty by default.
        """
        return pulumi.get(self, "auto_snapshot_policy_name")

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> pulumi.Output[Sequence[str]]:
        """
        The day on which an auto snapshot is created.
        - A maximum of 7 time points can be selected.
        - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
        """
        return pulumi.get(self, "repeat_weekdays")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[int]:
        """
        The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
        - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the automatic snapshot policy.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="timePoints")
    def time_points(self) -> pulumi.Output[Sequence[str]]:
        """
        The point in time at which an auto snapshot is created.
        - A maximum of 24 time points can be selected.
        - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
        """
        return pulumi.get(self, "time_points")

