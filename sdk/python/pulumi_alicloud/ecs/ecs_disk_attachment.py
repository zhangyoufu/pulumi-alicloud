# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EcsDiskAttachmentArgs', 'EcsDiskAttachment']

@pulumi.input_type
class EcsDiskAttachmentArgs:
    def __init__(__self__, *,
                 disk_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 bootable: Optional[pulumi.Input[bool]] = None,
                 delete_with_instance: Optional[pulumi.Input[bool]] = None,
                 key_pair_name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EcsDiskAttachment resource.
        :param pulumi.Input[str] disk_id: ID of the Disk to be attached.
        :param pulumi.Input[str] instance_id: ID of the Instance to attach to.
        :param pulumi.Input[bool] bootable: Whether to mount as a system disk. Default to: `false`.
        :param pulumi.Input[bool] delete_with_instance: Indicates whether the disk is released together with the instance. Default to: `false`.
        :param pulumi.Input[str] key_pair_name: The name of key pair
        :param pulumi.Input[str] password: When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
        """
        pulumi.set(__self__, "disk_id", disk_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if bootable is not None:
            pulumi.set(__self__, "bootable", bootable)
        if delete_with_instance is not None:
            pulumi.set(__self__, "delete_with_instance", delete_with_instance)
        if key_pair_name is not None:
            pulumi.set(__self__, "key_pair_name", key_pair_name)
        if password is not None:
            pulumi.set(__self__, "password", password)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Input[str]:
        """
        ID of the Disk to be attached.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of the Instance to attach to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def bootable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to mount as a system disk. Default to: `false`.
        """
        return pulumi.get(self, "bootable")

    @bootable.setter
    def bootable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bootable", value)

    @property
    @pulumi.getter(name="deleteWithInstance")
    def delete_with_instance(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the disk is released together with the instance. Default to: `false`.
        """
        return pulumi.get(self, "delete_with_instance")

    @delete_with_instance.setter
    def delete_with_instance(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_with_instance", value)

    @property
    @pulumi.getter(name="keyPairName")
    def key_pair_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of key pair
        """
        return pulumi.get(self, "key_pair_name")

    @key_pair_name.setter
    def key_pair_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_pair_name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)


@pulumi.input_type
class _EcsDiskAttachmentState:
    def __init__(__self__, *,
                 bootable: Optional[pulumi.Input[bool]] = None,
                 delete_with_instance: Optional[pulumi.Input[bool]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 key_pair_name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EcsDiskAttachment resources.
        :param pulumi.Input[bool] bootable: Whether to mount as a system disk. Default to: `false`.
        :param pulumi.Input[bool] delete_with_instance: Indicates whether the disk is released together with the instance. Default to: `false`.
        :param pulumi.Input[str] device: The name of the cloud disk device.
        :param pulumi.Input[str] disk_id: ID of the Disk to be attached.
        :param pulumi.Input[str] instance_id: ID of the Instance to attach to.
        :param pulumi.Input[str] key_pair_name: The name of key pair
        :param pulumi.Input[str] password: When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
        """
        if bootable is not None:
            pulumi.set(__self__, "bootable", bootable)
        if delete_with_instance is not None:
            pulumi.set(__self__, "delete_with_instance", delete_with_instance)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if disk_id is not None:
            pulumi.set(__self__, "disk_id", disk_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if key_pair_name is not None:
            pulumi.set(__self__, "key_pair_name", key_pair_name)
        if password is not None:
            pulumi.set(__self__, "password", password)

    @property
    @pulumi.getter
    def bootable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to mount as a system disk. Default to: `false`.
        """
        return pulumi.get(self, "bootable")

    @bootable.setter
    def bootable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bootable", value)

    @property
    @pulumi.getter(name="deleteWithInstance")
    def delete_with_instance(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the disk is released together with the instance. Default to: `false`.
        """
        return pulumi.get(self, "delete_with_instance")

    @delete_with_instance.setter
    def delete_with_instance(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_with_instance", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cloud disk device.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Disk to be attached.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Instance to attach to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="keyPairName")
    def key_pair_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of key pair
        """
        return pulumi.get(self, "key_pair_name")

    @key_pair_name.setter
    def key_pair_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_pair_name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)


class EcsDiskAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bootable: Optional[pulumi.Input[bool]] = None,
                 delete_with_instance: Optional[pulumi.Input[bool]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 key_pair_name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Alicloud ECS Disk Attachment as a resource, to attach and detach disks from ECS Instances.

        For information about ECS Disk Attachment and how to use it, see [What is Disk Attachment](https://www.alibabacloud.com/help/en/doc-detail/25515.htm).

        > **NOTE:** Available since v1.122.0+.

        ## Example Usage

        Basic usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.get_zones(available_resource_creation="Instance")
        default_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=default.zones[0].id,
            instance_type_family="ecs.sn1ne")
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vpc_id=default_network.id,
            cidr_block="10.4.0.0/24",
            zone_id=default.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("default",
            name="tf-example",
            description="New security group",
            vpc_id=default_network.id)
        default_get_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            most_recent=True,
            owners="system")
        default_instance = alicloud.ecs.Instance("default",
            availability_zone=default.zones[0].id,
            instance_name=name,
            host_name=name,
            image_id=default_get_images.images[0].id,
            instance_type=default_get_instance_types.instance_types[0].id,
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id)
        disk = alicloud.get_zones(available_resource_creation="VSwitch")
        default_ecs_disk = alicloud.ecs.EcsDisk("default",
            zone_id=disk.zones[0].id,
            category="cloud_efficiency",
            delete_auto_snapshot=True,
            description="Test For Terraform",
            disk_name=name,
            enable_auto_snapshot=True,
            encrypted=True,
            size=500,
            tags={
                "Created": "TF",
                "Environment": "Acceptance-test",
            })
        default_ecs_disk_attachment = alicloud.ecs.EcsDiskAttachment("default",
            disk_id=default_ecs_disk.id,
            instance_id=default_instance.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        The disk attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsDiskAttachment:EcsDiskAttachment example d-abc12345678:i-abc12355
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bootable: Whether to mount as a system disk. Default to: `false`.
        :param pulumi.Input[bool] delete_with_instance: Indicates whether the disk is released together with the instance. Default to: `false`.
        :param pulumi.Input[str] disk_id: ID of the Disk to be attached.
        :param pulumi.Input[str] instance_id: ID of the Instance to attach to.
        :param pulumi.Input[str] key_pair_name: The name of key pair
        :param pulumi.Input[str] password: When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EcsDiskAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Alicloud ECS Disk Attachment as a resource, to attach and detach disks from ECS Instances.

        For information about ECS Disk Attachment and how to use it, see [What is Disk Attachment](https://www.alibabacloud.com/help/en/doc-detail/25515.htm).

        > **NOTE:** Available since v1.122.0+.

        ## Example Usage

        Basic usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.get_zones(available_resource_creation="Instance")
        default_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=default.zones[0].id,
            instance_type_family="ecs.sn1ne")
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vpc_id=default_network.id,
            cidr_block="10.4.0.0/24",
            zone_id=default.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("default",
            name="tf-example",
            description="New security group",
            vpc_id=default_network.id)
        default_get_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            most_recent=True,
            owners="system")
        default_instance = alicloud.ecs.Instance("default",
            availability_zone=default.zones[0].id,
            instance_name=name,
            host_name=name,
            image_id=default_get_images.images[0].id,
            instance_type=default_get_instance_types.instance_types[0].id,
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id)
        disk = alicloud.get_zones(available_resource_creation="VSwitch")
        default_ecs_disk = alicloud.ecs.EcsDisk("default",
            zone_id=disk.zones[0].id,
            category="cloud_efficiency",
            delete_auto_snapshot=True,
            description="Test For Terraform",
            disk_name=name,
            enable_auto_snapshot=True,
            encrypted=True,
            size=500,
            tags={
                "Created": "TF",
                "Environment": "Acceptance-test",
            })
        default_ecs_disk_attachment = alicloud.ecs.EcsDiskAttachment("default",
            disk_id=default_ecs_disk.id,
            instance_id=default_instance.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        The disk attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsDiskAttachment:EcsDiskAttachment example d-abc12345678:i-abc12355
        ```

        :param str resource_name: The name of the resource.
        :param EcsDiskAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EcsDiskAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bootable: Optional[pulumi.Input[bool]] = None,
                 delete_with_instance: Optional[pulumi.Input[bool]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 key_pair_name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EcsDiskAttachmentArgs.__new__(EcsDiskAttachmentArgs)

            __props__.__dict__["bootable"] = bootable
            __props__.__dict__["delete_with_instance"] = delete_with_instance
            if disk_id is None and not opts.urn:
                raise TypeError("Missing required property 'disk_id'")
            __props__.__dict__["disk_id"] = disk_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["key_pair_name"] = key_pair_name
            __props__.__dict__["password"] = password
            __props__.__dict__["device"] = None
        super(EcsDiskAttachment, __self__).__init__(
            'alicloud:ecs/ecsDiskAttachment:EcsDiskAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bootable: Optional[pulumi.Input[bool]] = None,
            delete_with_instance: Optional[pulumi.Input[bool]] = None,
            device: Optional[pulumi.Input[str]] = None,
            disk_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            key_pair_name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None) -> 'EcsDiskAttachment':
        """
        Get an existing EcsDiskAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bootable: Whether to mount as a system disk. Default to: `false`.
        :param pulumi.Input[bool] delete_with_instance: Indicates whether the disk is released together with the instance. Default to: `false`.
        :param pulumi.Input[str] device: The name of the cloud disk device.
        :param pulumi.Input[str] disk_id: ID of the Disk to be attached.
        :param pulumi.Input[str] instance_id: ID of the Instance to attach to.
        :param pulumi.Input[str] key_pair_name: The name of key pair
        :param pulumi.Input[str] password: When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EcsDiskAttachmentState.__new__(_EcsDiskAttachmentState)

        __props__.__dict__["bootable"] = bootable
        __props__.__dict__["delete_with_instance"] = delete_with_instance
        __props__.__dict__["device"] = device
        __props__.__dict__["disk_id"] = disk_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["key_pair_name"] = key_pair_name
        __props__.__dict__["password"] = password
        return EcsDiskAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bootable(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to mount as a system disk. Default to: `false`.
        """
        return pulumi.get(self, "bootable")

    @property
    @pulumi.getter(name="deleteWithInstance")
    def delete_with_instance(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether the disk is released together with the instance. Default to: `false`.
        """
        return pulumi.get(self, "delete_with_instance")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[str]:
        """
        The name of the cloud disk device.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Output[str]:
        """
        ID of the Disk to be attached.
        """
        return pulumi.get(self, "disk_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of the Instance to attach to.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="keyPairName")
    def key_pair_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of key pair
        """
        return pulumi.get(self, "key_pair_name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
        """
        return pulumi.get(self, "password")

