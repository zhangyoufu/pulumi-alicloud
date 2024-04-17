# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RecycleBinArgs', 'RecycleBin']

@pulumi.input_type
class RecycleBinArgs:
    def __init__(__self__, *,
                 file_system_id: pulumi.Input[str],
                 reserved_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a RecycleBin resource.
        :param pulumi.Input[str] file_system_id: The ID of the file system for which you want to enable the recycle bin feature.
        :param pulumi.Input[int] reserved_days: The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
        """
        pulumi.set(__self__, "file_system_id", file_system_id)
        if reserved_days is not None:
            pulumi.set(__self__, "reserved_days", reserved_days)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Input[str]:
        """
        The ID of the file system for which you want to enable the recycle bin feature.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="reservedDays")
    def reserved_days(self) -> Optional[pulumi.Input[int]]:
        """
        The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
        """
        return pulumi.get(self, "reserved_days")

    @reserved_days.setter
    def reserved_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reserved_days", value)


@pulumi.input_type
class _RecycleBinState:
    def __init__(__self__, *,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 reserved_days: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RecycleBin resources.
        :param pulumi.Input[str] file_system_id: The ID of the file system for which you want to enable the recycle bin feature.
        :param pulumi.Input[int] reserved_days: The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
        :param pulumi.Input[str] status: The status of the recycle bin.
        """
        if file_system_id is not None:
            pulumi.set(__self__, "file_system_id", file_system_id)
        if reserved_days is not None:
            pulumi.set(__self__, "reserved_days", reserved_days)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the file system for which you want to enable the recycle bin feature.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="reservedDays")
    def reserved_days(self) -> Optional[pulumi.Input[int]]:
        """
        The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
        """
        return pulumi.get(self, "reserved_days")

    @reserved_days.setter
    def reserved_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reserved_days", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the recycle bin.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class RecycleBin(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 reserved_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a Network Attached Storage (NAS) Recycle Bin resource.

        For information about Network Attached Storage (NAS) Recycle Bin and how to use it, see [What is Recycle Bin](https://www.alibabacloud.com/help/en/doc-detail/264185.html).

        > **NOTE:** Available in v1.155.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.nas.get_zones(file_system_type="standard")
        example_file_system = alicloud.nas.FileSystem("example",
            protocol_type="NFS",
            storage_type="Performance",
            description="terraform-example",
            encrypt_type=1,
            zone_id=example.zones[0].zone_id)
        example_recycle_bin = alicloud.nas.RecycleBin("example",
            file_system_id=example_file_system.id,
            reserved_days=3)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Network Attached Storage (NAS) Recycle Bin can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:nas/recycleBin:RecycleBin example <file_system_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] file_system_id: The ID of the file system for which you want to enable the recycle bin feature.
        :param pulumi.Input[int] reserved_days: The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RecycleBinArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Network Attached Storage (NAS) Recycle Bin resource.

        For information about Network Attached Storage (NAS) Recycle Bin and how to use it, see [What is Recycle Bin](https://www.alibabacloud.com/help/en/doc-detail/264185.html).

        > **NOTE:** Available in v1.155.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.nas.get_zones(file_system_type="standard")
        example_file_system = alicloud.nas.FileSystem("example",
            protocol_type="NFS",
            storage_type="Performance",
            description="terraform-example",
            encrypt_type=1,
            zone_id=example.zones[0].zone_id)
        example_recycle_bin = alicloud.nas.RecycleBin("example",
            file_system_id=example_file_system.id,
            reserved_days=3)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Network Attached Storage (NAS) Recycle Bin can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:nas/recycleBin:RecycleBin example <file_system_id>
        ```

        :param str resource_name: The name of the resource.
        :param RecycleBinArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RecycleBinArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 reserved_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RecycleBinArgs.__new__(RecycleBinArgs)

            if file_system_id is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_id'")
            __props__.__dict__["file_system_id"] = file_system_id
            __props__.__dict__["reserved_days"] = reserved_days
            __props__.__dict__["status"] = None
        super(RecycleBin, __self__).__init__(
            'alicloud:nas/recycleBin:RecycleBin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            file_system_id: Optional[pulumi.Input[str]] = None,
            reserved_days: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'RecycleBin':
        """
        Get an existing RecycleBin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] file_system_id: The ID of the file system for which you want to enable the recycle bin feature.
        :param pulumi.Input[int] reserved_days: The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
        :param pulumi.Input[str] status: The status of the recycle bin.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RecycleBinState.__new__(_RecycleBinState)

        __props__.__dict__["file_system_id"] = file_system_id
        __props__.__dict__["reserved_days"] = reserved_days
        __props__.__dict__["status"] = status
        return RecycleBin(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Output[str]:
        """
        The ID of the file system for which you want to enable the recycle bin feature.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="reservedDays")
    def reserved_days(self) -> pulumi.Output[int]:
        """
        The period for which the files in the recycle bin are retained. Unit: days. Valid values: `1` to `180`.
        """
        return pulumi.get(self, "reserved_days")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the recycle bin.
        """
        return pulumi.get(self, "status")

