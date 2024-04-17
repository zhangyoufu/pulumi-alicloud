# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnapshotArgs', 'Snapshot']

@pulumi.input_type
class SnapshotArgs:
    def __init__(__self__, *,
                 disk_id: pulumi.Input[str],
                 snapshot_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a Snapshot resource.
        :param pulumi.Input[str] disk_id: The ID of the disk.
        :param pulumi.Input[str] snapshot_name: The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
        """
        pulumi.set(__self__, "disk_id", disk_id)
        pulumi.set(__self__, "snapshot_name", snapshot_name)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Input[str]:
        """
        The ID of the disk.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk_id", value)

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> pulumi.Input[str]:
        """
        The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
        """
        return pulumi.get(self, "snapshot_name")

    @snapshot_name.setter
    def snapshot_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "snapshot_name", value)


@pulumi.input_type
class _SnapshotState:
    def __init__(__self__, *,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Snapshot resources.
        :param pulumi.Input[str] disk_id: The ID of the disk.
        :param pulumi.Input[str] snapshot_name: The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
        :param pulumi.Input[str] status: The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
        """
        if disk_id is not None:
            pulumi.set(__self__, "disk_id", disk_id)
        if snapshot_name is not None:
            pulumi.set(__self__, "snapshot_name", snapshot_name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the disk.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_id", value)

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
        """
        return pulumi.get(self, "snapshot_name")

    @snapshot_name.setter
    def snapshot_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Snapshot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Simple Application Server Snapshot resource.

        For information about Simple Application Server Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/doc-detail/190452.htm).

        > **NOTE:** Available since v1.143.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf_example"
        default = alicloud.simpleapplicationserver.get_images()
        default_get_server_plans = alicloud.simpleapplicationserver.get_server_plans()
        default_instance = alicloud.simpleapplicationserver.Instance("default",
            payment_type="Subscription",
            plan_id=default_get_server_plans.plans[0].id,
            instance_name=name,
            image_id=default.images[0].id,
            period=1,
            data_disk_size=100)
        default_get_server_disks = alicloud.simpleapplicationserver.get_server_disks_output(instance_id=default_instance.id)
        default_snapshot = alicloud.simpleapplicationserver.Snapshot("default",
            disk_id=default_get_server_disks.ids[0],
            snapshot_name=name)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Simple Application Server Snapshot can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:simpleapplicationserver/snapshot:Snapshot example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_id: The ID of the disk.
        :param pulumi.Input[str] snapshot_name: The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Simple Application Server Snapshot resource.

        For information about Simple Application Server Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/doc-detail/190452.htm).

        > **NOTE:** Available since v1.143.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf_example"
        default = alicloud.simpleapplicationserver.get_images()
        default_get_server_plans = alicloud.simpleapplicationserver.get_server_plans()
        default_instance = alicloud.simpleapplicationserver.Instance("default",
            payment_type="Subscription",
            plan_id=default_get_server_plans.plans[0].id,
            instance_name=name,
            image_id=default.images[0].id,
            period=1,
            data_disk_size=100)
        default_get_server_disks = alicloud.simpleapplicationserver.get_server_disks_output(instance_id=default_instance.id)
        default_snapshot = alicloud.simpleapplicationserver.Snapshot("default",
            disk_id=default_get_server_disks.ids[0],
            snapshot_name=name)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Simple Application Server Snapshot can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:simpleapplicationserver/snapshot:Snapshot example <id>
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnapshotArgs.__new__(SnapshotArgs)

            if disk_id is None and not opts.urn:
                raise TypeError("Missing required property 'disk_id'")
            __props__.__dict__["disk_id"] = disk_id
            if snapshot_name is None and not opts.urn:
                raise TypeError("Missing required property 'snapshot_name'")
            __props__.__dict__["snapshot_name"] = snapshot_name
            __props__.__dict__["status"] = None
        super(Snapshot, __self__).__init__(
            'alicloud:simpleapplicationserver/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disk_id: Optional[pulumi.Input[str]] = None,
            snapshot_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Snapshot':
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_id: The ID of the disk.
        :param pulumi.Input[str] snapshot_name: The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
        :param pulumi.Input[str] status: The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotState.__new__(_SnapshotState)

        __props__.__dict__["disk_id"] = disk_id
        __props__.__dict__["snapshot_name"] = snapshot_name
        __props__.__dict__["status"] = status
        return Snapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Output[str]:
        """
        The ID of the disk.
        """
        return pulumi.get(self, "disk_id")

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> pulumi.Output[str]:
        """
        The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
        """
        return pulumi.get(self, "snapshot_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
        """
        return pulumi.get(self, "status")

