# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EcsSnapshotGroupArgs', 'EcsSnapshotGroup']

@pulumi.input_type
class EcsSnapshotGroupArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exclude_disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instant_access: Optional[pulumi.Input[bool]] = None,
                 instant_access_retention_days: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 snapshot_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a EcsSnapshotGroup resource.
        :param pulumi.Input[str] description: The description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disk_ids: The ID of disk for which to create snapshots. You can specify multiple disk IDs across instances with the same zone.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exclude_disk_ids: The ID of disk N for which you do not need to create snapshots. After this parameter is specified, the created snapshot-consistent group does not contain snapshots of the disk.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[bool] instant_access: Specifies whether to enable the instant access feature.
        :param pulumi.Input[int] instant_access_retention_days: Specify the number of days for which the instant access feature is available. Unit: days. Valid values: `1` to `65535`.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group to which the snapshot consistency group belongs.
        :param pulumi.Input[str] snapshot_group_name: The name of the snapshot-consistent group. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). It must start with a letter or a digit and cannot start with `http://` or `https://`.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the snapshot group.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disk_ids is not None:
            pulumi.set(__self__, "disk_ids", disk_ids)
        if exclude_disk_ids is not None:
            pulumi.set(__self__, "exclude_disk_ids", exclude_disk_ids)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instant_access is not None:
            pulumi.set(__self__, "instant_access", instant_access)
        if instant_access_retention_days is not None:
            pulumi.set(__self__, "instant_access_retention_days", instant_access_retention_days)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if snapshot_group_name is not None:
            pulumi.set(__self__, "snapshot_group_name", snapshot_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="diskIds")
    def disk_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ID of disk for which to create snapshots. You can specify multiple disk IDs across instances with the same zone.
        """
        return pulumi.get(self, "disk_ids")

    @disk_ids.setter
    def disk_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "disk_ids", value)

    @property
    @pulumi.getter(name="excludeDiskIds")
    def exclude_disk_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ID of disk N for which you do not need to create snapshots. After this parameter is specified, the created snapshot-consistent group does not contain snapshots of the disk.
        """
        return pulumi.get(self, "exclude_disk_ids")

    @exclude_disk_ids.setter
    def exclude_disk_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exclude_disk_ids", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instantAccess")
    def instant_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the instant access feature.
        """
        return pulumi.get(self, "instant_access")

    @instant_access.setter
    def instant_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "instant_access", value)

    @property
    @pulumi.getter(name="instantAccessRetentionDays")
    def instant_access_retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specify the number of days for which the instant access feature is available. Unit: days. Valid values: `1` to `65535`.
        """
        return pulumi.get(self, "instant_access_retention_days")

    @instant_access_retention_days.setter
    def instant_access_retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instant_access_retention_days", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group to which the snapshot consistency group belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="snapshotGroupName")
    def snapshot_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the snapshot-consistent group. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). It must start with a letter or a digit and cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "snapshot_group_name")

    @snapshot_group_name.setter
    def snapshot_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_group_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the snapshot group.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _EcsSnapshotGroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exclude_disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instant_access: Optional[pulumi.Input[bool]] = None,
                 instant_access_retention_days: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 snapshot_group_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering EcsSnapshotGroup resources.
        :param pulumi.Input[str] description: The description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disk_ids: The ID of disk for which to create snapshots. You can specify multiple disk IDs across instances with the same zone.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exclude_disk_ids: The ID of disk N for which you do not need to create snapshots. After this parameter is specified, the created snapshot-consistent group does not contain snapshots of the disk.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[bool] instant_access: Specifies whether to enable the instant access feature.
        :param pulumi.Input[int] instant_access_retention_days: Specify the number of days for which the instant access feature is available. Unit: days. Valid values: `1` to `65535`.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group to which the snapshot consistency group belongs.
        :param pulumi.Input[str] snapshot_group_name: The name of the snapshot-consistent group. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). It must start with a letter or a digit and cannot start with `http://` or `https://`.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the snapshot group.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disk_ids is not None:
            pulumi.set(__self__, "disk_ids", disk_ids)
        if exclude_disk_ids is not None:
            pulumi.set(__self__, "exclude_disk_ids", exclude_disk_ids)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instant_access is not None:
            pulumi.set(__self__, "instant_access", instant_access)
        if instant_access_retention_days is not None:
            pulumi.set(__self__, "instant_access_retention_days", instant_access_retention_days)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if snapshot_group_name is not None:
            pulumi.set(__self__, "snapshot_group_name", snapshot_group_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="diskIds")
    def disk_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ID of disk for which to create snapshots. You can specify multiple disk IDs across instances with the same zone.
        """
        return pulumi.get(self, "disk_ids")

    @disk_ids.setter
    def disk_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "disk_ids", value)

    @property
    @pulumi.getter(name="excludeDiskIds")
    def exclude_disk_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ID of disk N for which you do not need to create snapshots. After this parameter is specified, the created snapshot-consistent group does not contain snapshots of the disk.
        """
        return pulumi.get(self, "exclude_disk_ids")

    @exclude_disk_ids.setter
    def exclude_disk_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exclude_disk_ids", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instantAccess")
    def instant_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the instant access feature.
        """
        return pulumi.get(self, "instant_access")

    @instant_access.setter
    def instant_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "instant_access", value)

    @property
    @pulumi.getter(name="instantAccessRetentionDays")
    def instant_access_retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specify the number of days for which the instant access feature is available. Unit: days. Valid values: `1` to `65535`.
        """
        return pulumi.get(self, "instant_access_retention_days")

    @instant_access_retention_days.setter
    def instant_access_retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instant_access_retention_days", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group to which the snapshot consistency group belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="snapshotGroupName")
    def snapshot_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the snapshot-consistent group. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). It must start with a letter or a digit and cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "snapshot_group_name")

    @snapshot_group_name.setter
    def snapshot_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_group_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the snapshot group.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class EcsSnapshotGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exclude_disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instant_access: Optional[pulumi.Input[bool]] = None,
                 instant_access_retention_days: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 snapshot_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provides a ECS Snapshot Group resource.

        For information about ECS Snapshot Group and how to use it, see [What is Snapshot Group](https://www.alibabacloud.com/help/en/doc-detail/210939.html).

        > **NOTE:** Available in v1.160.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.get_zones(available_resource_creation="Instance",
            available_disk_category="cloud_essd")
        default_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=default.zones[0].id,
            system_disk_category="cloud_essd")
        default_get_images = alicloud.ecs.get_images(owners="system")
        default_network = alicloud.vpc.Network("default",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("default",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=default.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("default",
            name="terraform-example",
            vpc_id=default_network.id)
        default_instance = alicloud.ecs.Instance("default",
            availability_zone=default.zones[0].id,
            instance_name="terraform-example",
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id,
            instance_type=default_get_instance_types.instance_types[0].id,
            image_id=default_get_images.images[0].id,
            internet_max_bandwidth_out=10)
        default_ecs_disk = alicloud.ecs.EcsDisk("default",
            zone_id=default.zones[0].id,
            disk_name="terraform-example",
            description="terraform-example",
            category="cloud_essd",
            size=30)
        default_disk_attachment = alicloud.ecs.DiskAttachment("default",
            disk_id=default_ecs_disk.id,
            instance_id=default_instance.id)
        default_ecs_snapshot_group = alicloud.ecs.EcsSnapshotGroup("default",
            description="terraform-example",
            disk_ids=[default_ecs_disk.id],
            snapshot_group_name="terraform-example",
            instance_id=default_instance.id,
            instant_access=True,
            instant_access_retention_days=1,
            tags={
                "Created": "TF",
                "For": "Acceptance",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ECS Snapshot Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsSnapshotGroup:EcsSnapshotGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disk_ids: The ID of disk for which to create snapshots. You can specify multiple disk IDs across instances with the same zone.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exclude_disk_ids: The ID of disk N for which you do not need to create snapshots. After this parameter is specified, the created snapshot-consistent group does not contain snapshots of the disk.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[bool] instant_access: Specifies whether to enable the instant access feature.
        :param pulumi.Input[int] instant_access_retention_days: Specify the number of days for which the instant access feature is available. Unit: days. Valid values: `1` to `65535`.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group to which the snapshot consistency group belongs.
        :param pulumi.Input[str] snapshot_group_name: The name of the snapshot-consistent group. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). It must start with a letter or a digit and cannot start with `http://` or `https://`.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the snapshot group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EcsSnapshotGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ECS Snapshot Group resource.

        For information about ECS Snapshot Group and how to use it, see [What is Snapshot Group](https://www.alibabacloud.com/help/en/doc-detail/210939.html).

        > **NOTE:** Available in v1.160.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.get_zones(available_resource_creation="Instance",
            available_disk_category="cloud_essd")
        default_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=default.zones[0].id,
            system_disk_category="cloud_essd")
        default_get_images = alicloud.ecs.get_images(owners="system")
        default_network = alicloud.vpc.Network("default",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("default",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=default.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("default",
            name="terraform-example",
            vpc_id=default_network.id)
        default_instance = alicloud.ecs.Instance("default",
            availability_zone=default.zones[0].id,
            instance_name="terraform-example",
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id,
            instance_type=default_get_instance_types.instance_types[0].id,
            image_id=default_get_images.images[0].id,
            internet_max_bandwidth_out=10)
        default_ecs_disk = alicloud.ecs.EcsDisk("default",
            zone_id=default.zones[0].id,
            disk_name="terraform-example",
            description="terraform-example",
            category="cloud_essd",
            size=30)
        default_disk_attachment = alicloud.ecs.DiskAttachment("default",
            disk_id=default_ecs_disk.id,
            instance_id=default_instance.id)
        default_ecs_snapshot_group = alicloud.ecs.EcsSnapshotGroup("default",
            description="terraform-example",
            disk_ids=[default_ecs_disk.id],
            snapshot_group_name="terraform-example",
            instance_id=default_instance.id,
            instant_access=True,
            instant_access_retention_days=1,
            tags={
                "Created": "TF",
                "For": "Acceptance",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ECS Snapshot Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsSnapshotGroup:EcsSnapshotGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param EcsSnapshotGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EcsSnapshotGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exclude_disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instant_access: Optional[pulumi.Input[bool]] = None,
                 instant_access_retention_days: Optional[pulumi.Input[int]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 snapshot_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EcsSnapshotGroupArgs.__new__(EcsSnapshotGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["disk_ids"] = disk_ids
            __props__.__dict__["exclude_disk_ids"] = exclude_disk_ids
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["instant_access"] = instant_access
            __props__.__dict__["instant_access_retention_days"] = instant_access_retention_days
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["snapshot_group_name"] = snapshot_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["status"] = None
        super(EcsSnapshotGroup, __self__).__init__(
            'alicloud:ecs/ecsSnapshotGroup:EcsSnapshotGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            exclude_disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instant_access: Optional[pulumi.Input[bool]] = None,
            instant_access_retention_days: Optional[pulumi.Input[int]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            snapshot_group_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'EcsSnapshotGroup':
        """
        Get an existing EcsSnapshotGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disk_ids: The ID of disk for which to create snapshots. You can specify multiple disk IDs across instances with the same zone.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exclude_disk_ids: The ID of disk N for which you do not need to create snapshots. After this parameter is specified, the created snapshot-consistent group does not contain snapshots of the disk.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[bool] instant_access: Specifies whether to enable the instant access feature.
        :param pulumi.Input[int] instant_access_retention_days: Specify the number of days for which the instant access feature is available. Unit: days. Valid values: `1` to `65535`.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group to which the snapshot consistency group belongs.
        :param pulumi.Input[str] snapshot_group_name: The name of the snapshot-consistent group. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). It must start with a letter or a digit and cannot start with `http://` or `https://`.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the snapshot group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EcsSnapshotGroupState.__new__(_EcsSnapshotGroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["disk_ids"] = disk_ids
        __props__.__dict__["exclude_disk_ids"] = exclude_disk_ids
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["instant_access"] = instant_access
        __props__.__dict__["instant_access_retention_days"] = instant_access_retention_days
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["snapshot_group_name"] = snapshot_group_name
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        return EcsSnapshotGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskIds")
    def disk_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The ID of disk for which to create snapshots. You can specify multiple disk IDs across instances with the same zone.
        """
        return pulumi.get(self, "disk_ids")

    @property
    @pulumi.getter(name="excludeDiskIds")
    def exclude_disk_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The ID of disk N for which you do not need to create snapshots. After this parameter is specified, the created snapshot-consistent group does not contain snapshots of the disk.
        """
        return pulumi.get(self, "exclude_disk_ids")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instantAccess")
    def instant_access(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to enable the instant access feature.
        """
        return pulumi.get(self, "instant_access")

    @property
    @pulumi.getter(name="instantAccessRetentionDays")
    def instant_access_retention_days(self) -> pulumi.Output[Optional[int]]:
        """
        Specify the number of days for which the instant access feature is available. Unit: days. Valid values: `1` to `65535`.
        """
        return pulumi.get(self, "instant_access_retention_days")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource group to which the snapshot consistency group belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="snapshotGroupName")
    def snapshot_group_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the snapshot-consistent group. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). It must start with a letter or a digit and cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "snapshot_group_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the snapshot group.
        """
        return pulumi.get(self, "tags")

