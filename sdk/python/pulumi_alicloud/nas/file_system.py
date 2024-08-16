# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FileSystemArgs', 'FileSystem']

@pulumi.input_type
class FileSystemArgs:
    def __init__(__self__, *,
                 protocol_type: pulumi.Input[str],
                 storage_type: pulumi.Input[str],
                 capacity: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encrypt_type: Optional[pulumi.Input[int]] = None,
                 file_system_type: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FileSystem resource.
        :param pulumi.Input[str] protocol_type: The protocol type of the file system.
               Valid values:
               `NFS`,
               `SMB` (Available when the `file_system_type` is `standard`),
               `cpfs` (Available when the `file_system_type` is `cpfs`).
        :param pulumi.Input[str] storage_type: The storage type of the file System. 
               * Valid values:
               * `Performance` (Available when the `file_system_type` is `standard`)
               * `Capacity` (Available when the `file_system_type` is `standard`)
        :param pulumi.Input[int] capacity: The capacity of the file system. The `capacity` is required when the `file_system_type` is `extreme`.
               Unit: gib; **Note**: The minimum value is 100.
        :param pulumi.Input[str] description: The File System description.
        :param pulumi.Input[int] encrypt_type: Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt. 
               * Valid values:
        :param pulumi.Input[str] file_system_type: the type of the file system. 
               Valid values:
               `standard` (Default),
               `extreme`,
               `cpfs`.
        :param pulumi.Input[str] kms_key_id: The id of the KMS key. The `kms_key_id` is required when the `encrypt_type` is `2`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The id of the VPC. The `vpc_id` is required when the `file_system_type` is `cpfs`.
        :param pulumi.Input[str] vswitch_id: The id of the vSwitch. The `vswitch_id` is required when the `file_system_type` is `cpfs`.
        :param pulumi.Input[str] zone_id: The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocol_type` and `storage_type` configuration.
        """
        pulumi.set(__self__, "protocol_type", protocol_type)
        pulumi.set(__self__, "storage_type", storage_type)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encrypt_type is not None:
            pulumi.set(__self__, "encrypt_type", encrypt_type)
        if file_system_type is not None:
            pulumi.set(__self__, "file_system_type", file_system_type)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> pulumi.Input[str]:
        """
        The protocol type of the file system.
        Valid values:
        `NFS`,
        `SMB` (Available when the `file_system_type` is `standard`),
        `cpfs` (Available when the `file_system_type` is `cpfs`).
        """
        return pulumi.get(self, "protocol_type")

    @protocol_type.setter
    def protocol_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol_type", value)

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Input[str]:
        """
        The storage type of the file System. 
        * Valid values:
        * `Performance` (Available when the `file_system_type` is `standard`)
        * `Capacity` (Available when the `file_system_type` is `standard`)
        """
        return pulumi.get(self, "storage_type")

    @storage_type.setter
    def storage_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_type", value)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The capacity of the file system. The `capacity` is required when the `file_system_type` is `extreme`.
        Unit: gib; **Note**: The minimum value is 100.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The File System description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptType")
    def encrypt_type(self) -> Optional[pulumi.Input[int]]:
        """
        Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt. 
        * Valid values:
        """
        return pulumi.get(self, "encrypt_type")

    @encrypt_type.setter
    def encrypt_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "encrypt_type", value)

    @property
    @pulumi.getter(name="fileSystemType")
    def file_system_type(self) -> Optional[pulumi.Input[str]]:
        """
        the type of the file system. 
        Valid values:
        `standard` (Default),
        `extreme`,
        `cpfs`.
        """
        return pulumi.get(self, "file_system_type")

    @file_system_type.setter
    def file_system_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_type", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the KMS key. The `kms_key_id` is required when the `encrypt_type` is `2`.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the VPC. The `vpc_id` is required when the `file_system_type` is `cpfs`.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the vSwitch. The `vswitch_id` is required when the `file_system_type` is `cpfs`.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocol_type` and `storage_type` configuration.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


@pulumi.input_type
class _FileSystemState:
    def __init__(__self__, *,
                 capacity: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encrypt_type: Optional[pulumi.Input[int]] = None,
                 file_system_type: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FileSystem resources.
        :param pulumi.Input[int] capacity: The capacity of the file system. The `capacity` is required when the `file_system_type` is `extreme`.
               Unit: gib; **Note**: The minimum value is 100.
        :param pulumi.Input[str] description: The File System description.
        :param pulumi.Input[int] encrypt_type: Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt. 
               * Valid values:
        :param pulumi.Input[str] file_system_type: the type of the file system. 
               Valid values:
               `standard` (Default),
               `extreme`,
               `cpfs`.
        :param pulumi.Input[str] kms_key_id: The id of the KMS key. The `kms_key_id` is required when the `encrypt_type` is `2`.
        :param pulumi.Input[str] protocol_type: The protocol type of the file system.
               Valid values:
               `NFS`,
               `SMB` (Available when the `file_system_type` is `standard`),
               `cpfs` (Available when the `file_system_type` is `cpfs`).
        :param pulumi.Input[str] storage_type: The storage type of the file System. 
               * Valid values:
               * `Performance` (Available when the `file_system_type` is `standard`)
               * `Capacity` (Available when the `file_system_type` is `standard`)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The id of the VPC. The `vpc_id` is required when the `file_system_type` is `cpfs`.
        :param pulumi.Input[str] vswitch_id: The id of the vSwitch. The `vswitch_id` is required when the `file_system_type` is `cpfs`.
        :param pulumi.Input[str] zone_id: The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocol_type` and `storage_type` configuration.
        """
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encrypt_type is not None:
            pulumi.set(__self__, "encrypt_type", encrypt_type)
        if file_system_type is not None:
            pulumi.set(__self__, "file_system_type", file_system_type)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if protocol_type is not None:
            pulumi.set(__self__, "protocol_type", protocol_type)
        if storage_type is not None:
            pulumi.set(__self__, "storage_type", storage_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The capacity of the file system. The `capacity` is required when the `file_system_type` is `extreme`.
        Unit: gib; **Note**: The minimum value is 100.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The File System description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptType")
    def encrypt_type(self) -> Optional[pulumi.Input[int]]:
        """
        Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt. 
        * Valid values:
        """
        return pulumi.get(self, "encrypt_type")

    @encrypt_type.setter
    def encrypt_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "encrypt_type", value)

    @property
    @pulumi.getter(name="fileSystemType")
    def file_system_type(self) -> Optional[pulumi.Input[str]]:
        """
        the type of the file system. 
        Valid values:
        `standard` (Default),
        `extreme`,
        `cpfs`.
        """
        return pulumi.get(self, "file_system_type")

    @file_system_type.setter
    def file_system_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_type", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the KMS key. The `kms_key_id` is required when the `encrypt_type` is `2`.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol type of the file system.
        Valid values:
        `NFS`,
        `SMB` (Available when the `file_system_type` is `standard`),
        `cpfs` (Available when the `file_system_type` is `cpfs`).
        """
        return pulumi.get(self, "protocol_type")

    @protocol_type.setter
    def protocol_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol_type", value)

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> Optional[pulumi.Input[str]]:
        """
        The storage type of the file System. 
        * Valid values:
        * `Performance` (Available when the `file_system_type` is `standard`)
        * `Capacity` (Available when the `file_system_type` is `standard`)
        """
        return pulumi.get(self, "storage_type")

    @storage_type.setter
    def storage_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the VPC. The `vpc_id` is required when the `file_system_type` is `cpfs`.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the vSwitch. The `vswitch_id` is required when the `file_system_type` is `cpfs`.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocol_type` and `storage_type` configuration.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class FileSystem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encrypt_type: Optional[pulumi.Input[int]] = None,
                 file_system_type: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Nas File System resource.

        After activating NAS, you can create a file system and purchase a storage package for it in the NAS console. The NAS console also enables you to view the file system details and remove unnecessary file systems.

        For information about NAS file system and how to use it, see [Manage file systems](https://www.alibabacloud.com/help/doc-detail/27530.htm)

        > **NOTE:** Available in v1.33.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.nas.get_zones(file_system_type="standard")
        foo = alicloud.nas.FileSystem("foo",
            protocol_type="NFS",
            storage_type="Performance",
            description="terraform-example",
            encrypt_type=1,
            zone_id=example.zones[0].zone_id)
        ```

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.nas.get_zones(file_system_type="extreme")
        foo = alicloud.nas.FileSystem("foo",
            file_system_type="extreme",
            protocol_type="NFS",
            zone_id=example.zones[0].zone_id,
            storage_type="standard",
            capacity=100)
        ```

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.nas.get_zones(file_system_type="cpfs")
        example_network = alicloud.vpc.Network("example",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        example_switch = alicloud.vpc.Switch("example",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=example_network.id,
            zone_id=example.zones[1].zone_id)
        example_file_system = alicloud.nas.FileSystem("example",
            protocol_type="cpfs",
            storage_type="advance_200",
            file_system_type="cpfs",
            capacity=3600,
            zone_id=example.zones[1].zone_id,
            vpc_id=example_network.id,
            vswitch_id=example_switch.id)
        ```

        ## Import

        Nas File System can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:nas/fileSystem:FileSystem foo 1337849c59
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] capacity: The capacity of the file system. The `capacity` is required when the `file_system_type` is `extreme`.
               Unit: gib; **Note**: The minimum value is 100.
        :param pulumi.Input[str] description: The File System description.
        :param pulumi.Input[int] encrypt_type: Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt. 
               * Valid values:
        :param pulumi.Input[str] file_system_type: the type of the file system. 
               Valid values:
               `standard` (Default),
               `extreme`,
               `cpfs`.
        :param pulumi.Input[str] kms_key_id: The id of the KMS key. The `kms_key_id` is required when the `encrypt_type` is `2`.
        :param pulumi.Input[str] protocol_type: The protocol type of the file system.
               Valid values:
               `NFS`,
               `SMB` (Available when the `file_system_type` is `standard`),
               `cpfs` (Available when the `file_system_type` is `cpfs`).
        :param pulumi.Input[str] storage_type: The storage type of the file System. 
               * Valid values:
               * `Performance` (Available when the `file_system_type` is `standard`)
               * `Capacity` (Available when the `file_system_type` is `standard`)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The id of the VPC. The `vpc_id` is required when the `file_system_type` is `cpfs`.
        :param pulumi.Input[str] vswitch_id: The id of the vSwitch. The `vswitch_id` is required when the `file_system_type` is `cpfs`.
        :param pulumi.Input[str] zone_id: The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocol_type` and `storage_type` configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FileSystemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Nas File System resource.

        After activating NAS, you can create a file system and purchase a storage package for it in the NAS console. The NAS console also enables you to view the file system details and remove unnecessary file systems.

        For information about NAS file system and how to use it, see [Manage file systems](https://www.alibabacloud.com/help/doc-detail/27530.htm)

        > **NOTE:** Available in v1.33.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.nas.get_zones(file_system_type="standard")
        foo = alicloud.nas.FileSystem("foo",
            protocol_type="NFS",
            storage_type="Performance",
            description="terraform-example",
            encrypt_type=1,
            zone_id=example.zones[0].zone_id)
        ```

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.nas.get_zones(file_system_type="extreme")
        foo = alicloud.nas.FileSystem("foo",
            file_system_type="extreme",
            protocol_type="NFS",
            zone_id=example.zones[0].zone_id,
            storage_type="standard",
            capacity=100)
        ```

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.nas.get_zones(file_system_type="cpfs")
        example_network = alicloud.vpc.Network("example",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        example_switch = alicloud.vpc.Switch("example",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=example_network.id,
            zone_id=example.zones[1].zone_id)
        example_file_system = alicloud.nas.FileSystem("example",
            protocol_type="cpfs",
            storage_type="advance_200",
            file_system_type="cpfs",
            capacity=3600,
            zone_id=example.zones[1].zone_id,
            vpc_id=example_network.id,
            vswitch_id=example_switch.id)
        ```

        ## Import

        Nas File System can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:nas/fileSystem:FileSystem foo 1337849c59
        ```

        :param str resource_name: The name of the resource.
        :param FileSystemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FileSystemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encrypt_type: Optional[pulumi.Input[int]] = None,
                 file_system_type: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FileSystemArgs.__new__(FileSystemArgs)

            __props__.__dict__["capacity"] = capacity
            __props__.__dict__["description"] = description
            __props__.__dict__["encrypt_type"] = encrypt_type
            __props__.__dict__["file_system_type"] = file_system_type
            __props__.__dict__["kms_key_id"] = kms_key_id
            if protocol_type is None and not opts.urn:
                raise TypeError("Missing required property 'protocol_type'")
            __props__.__dict__["protocol_type"] = protocol_type
            if storage_type is None and not opts.urn:
                raise TypeError("Missing required property 'storage_type'")
            __props__.__dict__["storage_type"] = storage_type
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["vswitch_id"] = vswitch_id
            __props__.__dict__["zone_id"] = zone_id
        super(FileSystem, __self__).__init__(
            'alicloud:nas/fileSystem:FileSystem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            capacity: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            encrypt_type: Optional[pulumi.Input[int]] = None,
            file_system_type: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            protocol_type: Optional[pulumi.Input[str]] = None,
            storage_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'FileSystem':
        """
        Get an existing FileSystem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] capacity: The capacity of the file system. The `capacity` is required when the `file_system_type` is `extreme`.
               Unit: gib; **Note**: The minimum value is 100.
        :param pulumi.Input[str] description: The File System description.
        :param pulumi.Input[int] encrypt_type: Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt. 
               * Valid values:
        :param pulumi.Input[str] file_system_type: the type of the file system. 
               Valid values:
               `standard` (Default),
               `extreme`,
               `cpfs`.
        :param pulumi.Input[str] kms_key_id: The id of the KMS key. The `kms_key_id` is required when the `encrypt_type` is `2`.
        :param pulumi.Input[str] protocol_type: The protocol type of the file system.
               Valid values:
               `NFS`,
               `SMB` (Available when the `file_system_type` is `standard`),
               `cpfs` (Available when the `file_system_type` is `cpfs`).
        :param pulumi.Input[str] storage_type: The storage type of the file System. 
               * Valid values:
               * `Performance` (Available when the `file_system_type` is `standard`)
               * `Capacity` (Available when the `file_system_type` is `standard`)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The id of the VPC. The `vpc_id` is required when the `file_system_type` is `cpfs`.
        :param pulumi.Input[str] vswitch_id: The id of the vSwitch. The `vswitch_id` is required when the `file_system_type` is `cpfs`.
        :param pulumi.Input[str] zone_id: The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocol_type` and `storage_type` configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FileSystemState.__new__(_FileSystemState)

        __props__.__dict__["capacity"] = capacity
        __props__.__dict__["description"] = description
        __props__.__dict__["encrypt_type"] = encrypt_type
        __props__.__dict__["file_system_type"] = file_system_type
        __props__.__dict__["kms_key_id"] = kms_key_id
        __props__.__dict__["protocol_type"] = protocol_type
        __props__.__dict__["storage_type"] = storage_type
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vswitch_id"] = vswitch_id
        __props__.__dict__["zone_id"] = zone_id
        return FileSystem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Output[int]:
        """
        The capacity of the file system. The `capacity` is required when the `file_system_type` is `extreme`.
        Unit: gib; **Note**: The minimum value is 100.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The File System description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encryptType")
    def encrypt_type(self) -> pulumi.Output[Optional[int]]:
        """
        Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt. 
        * Valid values:
        """
        return pulumi.get(self, "encrypt_type")

    @property
    @pulumi.getter(name="fileSystemType")
    def file_system_type(self) -> pulumi.Output[Optional[str]]:
        """
        the type of the file system. 
        Valid values:
        `standard` (Default),
        `extreme`,
        `cpfs`.
        """
        return pulumi.get(self, "file_system_type")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[str]:
        """
        The id of the KMS key. The `kms_key_id` is required when the `encrypt_type` is `2`.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> pulumi.Output[str]:
        """
        The protocol type of the file system.
        Valid values:
        `NFS`,
        `SMB` (Available when the `file_system_type` is `standard`),
        `cpfs` (Available when the `file_system_type` is `cpfs`).
        """
        return pulumi.get(self, "protocol_type")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Output[str]:
        """
        The storage type of the file System. 
        * Valid values:
        * `Performance` (Available when the `file_system_type` is `standard`)
        * `Capacity` (Available when the `file_system_type` is `standard`)
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[Optional[str]]:
        """
        The id of the VPC. The `vpc_id` is required when the `file_system_type` is `cpfs`.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[Optional[str]]:
        """
        The id of the vSwitch. The `vswitch_id` is required when the `file_system_type` is `cpfs`.
        """
        return pulumi.get(self, "vswitch_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The available zones information that supports nas.When FileSystemType=standard, this parameter is not required. **Note:** By default, a qualified availability zone is randomly selected according to the `protocol_type` and `storage_type` configuration.
        """
        return pulumi.get(self, "zone_id")

