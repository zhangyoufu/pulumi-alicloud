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

__all__ = ['AccessPointArgs', 'AccessPoint']

@pulumi.input_type
class AccessPointArgs:
    def __init__(__self__, *,
                 access_group: pulumi.Input[str],
                 file_system_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str],
                 access_point_name: Optional[pulumi.Input[str]] = None,
                 enabled_ram: Optional[pulumi.Input[bool]] = None,
                 posix_user: Optional[pulumi.Input['AccessPointPosixUserArgs']] = None,
                 root_path: Optional[pulumi.Input[str]] = None,
                 root_path_permission: Optional[pulumi.Input['AccessPointRootPathPermissionArgs']] = None):
        """
        The set of arguments for constructing a AccessPoint resource.
        :param pulumi.Input[str] access_group: The permission group name.
        :param pulumi.Input[str] file_system_id: The ID of the file system.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        :param pulumi.Input[str] vswitch_id: The vSwitch ID.
        :param pulumi.Input[str] access_point_name: The Access Point Name.
        :param pulumi.Input[bool] enabled_ram: Whether to enable the RAM policy.
        :param pulumi.Input['AccessPointPosixUserArgs'] posix_user: The Posix user. See `posix_user` below.
        :param pulumi.Input[str] root_path: The root directory.
        :param pulumi.Input['AccessPointRootPathPermissionArgs'] root_path_permission: Root permissions. See `root_path_permission` below.
        """
        pulumi.set(__self__, "access_group", access_group)
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        if access_point_name is not None:
            pulumi.set(__self__, "access_point_name", access_point_name)
        if enabled_ram is not None:
            pulumi.set(__self__, "enabled_ram", enabled_ram)
        if posix_user is not None:
            pulumi.set(__self__, "posix_user", posix_user)
        if root_path is not None:
            pulumi.set(__self__, "root_path", root_path)
        if root_path_permission is not None:
            pulumi.set(__self__, "root_path_permission", root_path_permission)

    @property
    @pulumi.getter(name="accessGroup")
    def access_group(self) -> pulumi.Input[str]:
        """
        The permission group name.
        """
        return pulumi.get(self, "access_group")

    @access_group.setter
    def access_group(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_group", value)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Input[str]:
        """
        The ID of the file system.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        The vSwitch ID.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="accessPointName")
    def access_point_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Access Point Name.
        """
        return pulumi.get(self, "access_point_name")

    @access_point_name.setter
    def access_point_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_point_name", value)

    @property
    @pulumi.getter(name="enabledRam")
    def enabled_ram(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the RAM policy.
        """
        return pulumi.get(self, "enabled_ram")

    @enabled_ram.setter
    def enabled_ram(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled_ram", value)

    @property
    @pulumi.getter(name="posixUser")
    def posix_user(self) -> Optional[pulumi.Input['AccessPointPosixUserArgs']]:
        """
        The Posix user. See `posix_user` below.
        """
        return pulumi.get(self, "posix_user")

    @posix_user.setter
    def posix_user(self, value: Optional[pulumi.Input['AccessPointPosixUserArgs']]):
        pulumi.set(self, "posix_user", value)

    @property
    @pulumi.getter(name="rootPath")
    def root_path(self) -> Optional[pulumi.Input[str]]:
        """
        The root directory.
        """
        return pulumi.get(self, "root_path")

    @root_path.setter
    def root_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_path", value)

    @property
    @pulumi.getter(name="rootPathPermission")
    def root_path_permission(self) -> Optional[pulumi.Input['AccessPointRootPathPermissionArgs']]:
        """
        Root permissions. See `root_path_permission` below.
        """
        return pulumi.get(self, "root_path_permission")

    @root_path_permission.setter
    def root_path_permission(self, value: Optional[pulumi.Input['AccessPointRootPathPermissionArgs']]):
        pulumi.set(self, "root_path_permission", value)


@pulumi.input_type
class _AccessPointState:
    def __init__(__self__, *,
                 access_group: Optional[pulumi.Input[str]] = None,
                 access_point_id: Optional[pulumi.Input[str]] = None,
                 access_point_name: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 enabled_ram: Optional[pulumi.Input[bool]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 posix_user: Optional[pulumi.Input['AccessPointPosixUserArgs']] = None,
                 root_path: Optional[pulumi.Input[str]] = None,
                 root_path_permission: Optional[pulumi.Input['AccessPointRootPathPermissionArgs']] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccessPoint resources.
        :param pulumi.Input[str] access_group: The permission group name.
        :param pulumi.Input[str] access_point_id: Access point ID.
        :param pulumi.Input[str] access_point_name: The Access Point Name.
        :param pulumi.Input[str] create_time: Creation time.
        :param pulumi.Input[bool] enabled_ram: Whether to enable the RAM policy.
        :param pulumi.Input[str] file_system_id: The ID of the file system.
        :param pulumi.Input['AccessPointPosixUserArgs'] posix_user: The Posix user. See `posix_user` below.
        :param pulumi.Input[str] root_path: The root directory.
        :param pulumi.Input['AccessPointRootPathPermissionArgs'] root_path_permission: Root permissions. See `root_path_permission` below.
        :param pulumi.Input[str] status: Current access point state.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        :param pulumi.Input[str] vswitch_id: The vSwitch ID.
        """
        if access_group is not None:
            pulumi.set(__self__, "access_group", access_group)
        if access_point_id is not None:
            pulumi.set(__self__, "access_point_id", access_point_id)
        if access_point_name is not None:
            pulumi.set(__self__, "access_point_name", access_point_name)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if enabled_ram is not None:
            pulumi.set(__self__, "enabled_ram", enabled_ram)
        if file_system_id is not None:
            pulumi.set(__self__, "file_system_id", file_system_id)
        if posix_user is not None:
            pulumi.set(__self__, "posix_user", posix_user)
        if root_path is not None:
            pulumi.set(__self__, "root_path", root_path)
        if root_path_permission is not None:
            pulumi.set(__self__, "root_path_permission", root_path_permission)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="accessGroup")
    def access_group(self) -> Optional[pulumi.Input[str]]:
        """
        The permission group name.
        """
        return pulumi.get(self, "access_group")

    @access_group.setter
    def access_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_group", value)

    @property
    @pulumi.getter(name="accessPointId")
    def access_point_id(self) -> Optional[pulumi.Input[str]]:
        """
        Access point ID.
        """
        return pulumi.get(self, "access_point_id")

    @access_point_id.setter
    def access_point_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_point_id", value)

    @property
    @pulumi.getter(name="accessPointName")
    def access_point_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Access Point Name.
        """
        return pulumi.get(self, "access_point_name")

    @access_point_name.setter
    def access_point_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_point_name", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="enabledRam")
    def enabled_ram(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the RAM policy.
        """
        return pulumi.get(self, "enabled_ram")

    @enabled_ram.setter
    def enabled_ram(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled_ram", value)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the file system.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="posixUser")
    def posix_user(self) -> Optional[pulumi.Input['AccessPointPosixUserArgs']]:
        """
        The Posix user. See `posix_user` below.
        """
        return pulumi.get(self, "posix_user")

    @posix_user.setter
    def posix_user(self, value: Optional[pulumi.Input['AccessPointPosixUserArgs']]):
        pulumi.set(self, "posix_user", value)

    @property
    @pulumi.getter(name="rootPath")
    def root_path(self) -> Optional[pulumi.Input[str]]:
        """
        The root directory.
        """
        return pulumi.get(self, "root_path")

    @root_path.setter
    def root_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_path", value)

    @property
    @pulumi.getter(name="rootPathPermission")
    def root_path_permission(self) -> Optional[pulumi.Input['AccessPointRootPathPermissionArgs']]:
        """
        Root permissions. See `root_path_permission` below.
        """
        return pulumi.get(self, "root_path_permission")

    @root_path_permission.setter
    def root_path_permission(self, value: Optional[pulumi.Input['AccessPointRootPathPermissionArgs']]):
        pulumi.set(self, "root_path_permission", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Current access point state.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The vSwitch ID.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class AccessPoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_group: Optional[pulumi.Input[str]] = None,
                 access_point_name: Optional[pulumi.Input[str]] = None,
                 enabled_ram: Optional[pulumi.Input[bool]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 posix_user: Optional[pulumi.Input[Union['AccessPointPosixUserArgs', 'AccessPointPosixUserArgsDict']]] = None,
                 root_path: Optional[pulumi.Input[str]] = None,
                 root_path_permission: Optional[pulumi.Input[Union['AccessPointRootPathPermissionArgs', 'AccessPointRootPathPermissionArgsDict']]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a NAS Access Point resource.

        For information about NAS Access Point and how to use it, see [What is Access Point](https://www.alibabacloud.com/help/zh/nas/developer-reference/api-nas-2017-06-26-createaccesspoint).

        > **NOTE:** Available since v1.224.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        region_id = config.get("regionId")
        if region_id is None:
            region_id = "cn-hangzhou"
        azone = config.get("azone")
        if azone is None:
            azone = "cn-hangzhou-g"
        default = alicloud.get_zones(available_resource_creation="VSwitch")
        defaultky_vc70 = alicloud.vpc.Network("defaultkyVC70",
            cidr_block="172.16.0.0/12",
            description="接入点测试noRootDirectory")
        defaulto_za_pm_o = alicloud.vpc.Switch("defaultoZAPmO",
            vpc_id=defaultky_vc70.id,
            zone_id=default.zones[0].id,
            cidr_block="172.16.0.0/24")
        default_bbc7ev = alicloud.nas.AccessGroup("defaultBbc7ev",
            access_group_type="Vpc",
            access_group_name=name,
            file_system_type="standard")
        default_vt_up_dh = alicloud.nas.FileSystem("defaultVtUpDh",
            storage_type="Performance",
            zone_id=azone,
            encrypt_type=0,
            protocol_type="NFS",
            file_system_type="standard",
            description="AccessPointnoRootDirectory")
        default_access_point = alicloud.nas.AccessPoint("default",
            vpc_id=defaultky_vc70.id,
            access_group=default_bbc7ev.access_group_name,
            vswitch_id=defaulto_za_pm_o.id,
            file_system_id=default_vt_up_dh.id,
            access_point_name=name,
            posix_user={
                "posix_group_id": 123,
                "posix_user_id": 123,
            },
            root_path_permission={
                "owner_group_id": 1,
                "owner_user_id": 1,
                "permission": "0777",
            })
        ```

        ## Import

        NAS Access Point can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:nas/accessPoint:AccessPoint example <file_system_id>:<access_point_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_group: The permission group name.
        :param pulumi.Input[str] access_point_name: The Access Point Name.
        :param pulumi.Input[bool] enabled_ram: Whether to enable the RAM policy.
        :param pulumi.Input[str] file_system_id: The ID of the file system.
        :param pulumi.Input[Union['AccessPointPosixUserArgs', 'AccessPointPosixUserArgsDict']] posix_user: The Posix user. See `posix_user` below.
        :param pulumi.Input[str] root_path: The root directory.
        :param pulumi.Input[Union['AccessPointRootPathPermissionArgs', 'AccessPointRootPathPermissionArgsDict']] root_path_permission: Root permissions. See `root_path_permission` below.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        :param pulumi.Input[str] vswitch_id: The vSwitch ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessPointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a NAS Access Point resource.

        For information about NAS Access Point and how to use it, see [What is Access Point](https://www.alibabacloud.com/help/zh/nas/developer-reference/api-nas-2017-06-26-createaccesspoint).

        > **NOTE:** Available since v1.224.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        region_id = config.get("regionId")
        if region_id is None:
            region_id = "cn-hangzhou"
        azone = config.get("azone")
        if azone is None:
            azone = "cn-hangzhou-g"
        default = alicloud.get_zones(available_resource_creation="VSwitch")
        defaultky_vc70 = alicloud.vpc.Network("defaultkyVC70",
            cidr_block="172.16.0.0/12",
            description="接入点测试noRootDirectory")
        defaulto_za_pm_o = alicloud.vpc.Switch("defaultoZAPmO",
            vpc_id=defaultky_vc70.id,
            zone_id=default.zones[0].id,
            cidr_block="172.16.0.0/24")
        default_bbc7ev = alicloud.nas.AccessGroup("defaultBbc7ev",
            access_group_type="Vpc",
            access_group_name=name,
            file_system_type="standard")
        default_vt_up_dh = alicloud.nas.FileSystem("defaultVtUpDh",
            storage_type="Performance",
            zone_id=azone,
            encrypt_type=0,
            protocol_type="NFS",
            file_system_type="standard",
            description="AccessPointnoRootDirectory")
        default_access_point = alicloud.nas.AccessPoint("default",
            vpc_id=defaultky_vc70.id,
            access_group=default_bbc7ev.access_group_name,
            vswitch_id=defaulto_za_pm_o.id,
            file_system_id=default_vt_up_dh.id,
            access_point_name=name,
            posix_user={
                "posix_group_id": 123,
                "posix_user_id": 123,
            },
            root_path_permission={
                "owner_group_id": 1,
                "owner_user_id": 1,
                "permission": "0777",
            })
        ```

        ## Import

        NAS Access Point can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:nas/accessPoint:AccessPoint example <file_system_id>:<access_point_id>
        ```

        :param str resource_name: The name of the resource.
        :param AccessPointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessPointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_group: Optional[pulumi.Input[str]] = None,
                 access_point_name: Optional[pulumi.Input[str]] = None,
                 enabled_ram: Optional[pulumi.Input[bool]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 posix_user: Optional[pulumi.Input[Union['AccessPointPosixUserArgs', 'AccessPointPosixUserArgsDict']]] = None,
                 root_path: Optional[pulumi.Input[str]] = None,
                 root_path_permission: Optional[pulumi.Input[Union['AccessPointRootPathPermissionArgs', 'AccessPointRootPathPermissionArgsDict']]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessPointArgs.__new__(AccessPointArgs)

            if access_group is None and not opts.urn:
                raise TypeError("Missing required property 'access_group'")
            __props__.__dict__["access_group"] = access_group
            __props__.__dict__["access_point_name"] = access_point_name
            __props__.__dict__["enabled_ram"] = enabled_ram
            if file_system_id is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_id'")
            __props__.__dict__["file_system_id"] = file_system_id
            __props__.__dict__["posix_user"] = posix_user
            __props__.__dict__["root_path"] = root_path
            __props__.__dict__["root_path_permission"] = root_path_permission
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
            __props__.__dict__["access_point_id"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["status"] = None
        super(AccessPoint, __self__).__init__(
            'alicloud:nas/accessPoint:AccessPoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_group: Optional[pulumi.Input[str]] = None,
            access_point_id: Optional[pulumi.Input[str]] = None,
            access_point_name: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            enabled_ram: Optional[pulumi.Input[bool]] = None,
            file_system_id: Optional[pulumi.Input[str]] = None,
            posix_user: Optional[pulumi.Input[Union['AccessPointPosixUserArgs', 'AccessPointPosixUserArgsDict']]] = None,
            root_path: Optional[pulumi.Input[str]] = None,
            root_path_permission: Optional[pulumi.Input[Union['AccessPointRootPathPermissionArgs', 'AccessPointRootPathPermissionArgsDict']]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'AccessPoint':
        """
        Get an existing AccessPoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_group: The permission group name.
        :param pulumi.Input[str] access_point_id: Access point ID.
        :param pulumi.Input[str] access_point_name: The Access Point Name.
        :param pulumi.Input[str] create_time: Creation time.
        :param pulumi.Input[bool] enabled_ram: Whether to enable the RAM policy.
        :param pulumi.Input[str] file_system_id: The ID of the file system.
        :param pulumi.Input[Union['AccessPointPosixUserArgs', 'AccessPointPosixUserArgsDict']] posix_user: The Posix user. See `posix_user` below.
        :param pulumi.Input[str] root_path: The root directory.
        :param pulumi.Input[Union['AccessPointRootPathPermissionArgs', 'AccessPointRootPathPermissionArgsDict']] root_path_permission: Root permissions. See `root_path_permission` below.
        :param pulumi.Input[str] status: Current access point state.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        :param pulumi.Input[str] vswitch_id: The vSwitch ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccessPointState.__new__(_AccessPointState)

        __props__.__dict__["access_group"] = access_group
        __props__.__dict__["access_point_id"] = access_point_id
        __props__.__dict__["access_point_name"] = access_point_name
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["enabled_ram"] = enabled_ram
        __props__.__dict__["file_system_id"] = file_system_id
        __props__.__dict__["posix_user"] = posix_user
        __props__.__dict__["root_path"] = root_path
        __props__.__dict__["root_path_permission"] = root_path_permission
        __props__.__dict__["status"] = status
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vswitch_id"] = vswitch_id
        return AccessPoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessGroup")
    def access_group(self) -> pulumi.Output[str]:
        """
        The permission group name.
        """
        return pulumi.get(self, "access_group")

    @property
    @pulumi.getter(name="accessPointId")
    def access_point_id(self) -> pulumi.Output[str]:
        """
        Access point ID.
        """
        return pulumi.get(self, "access_point_id")

    @property
    @pulumi.getter(name="accessPointName")
    def access_point_name(self) -> pulumi.Output[Optional[str]]:
        """
        The Access Point Name.
        """
        return pulumi.get(self, "access_point_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="enabledRam")
    def enabled_ram(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable the RAM policy.
        """
        return pulumi.get(self, "enabled_ram")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Output[str]:
        """
        The ID of the file system.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="posixUser")
    def posix_user(self) -> pulumi.Output['outputs.AccessPointPosixUser']:
        """
        The Posix user. See `posix_user` below.
        """
        return pulumi.get(self, "posix_user")

    @property
    @pulumi.getter(name="rootPath")
    def root_path(self) -> pulumi.Output[str]:
        """
        The root directory.
        """
        return pulumi.get(self, "root_path")

    @property
    @pulumi.getter(name="rootPathPermission")
    def root_path_permission(self) -> pulumi.Output['outputs.AccessPointRootPathPermission']:
        """
        Root permissions. See `root_path_permission` below.
        """
        return pulumi.get(self, "root_path_permission")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Current access point state.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The vSwitch ID.
        """
        return pulumi.get(self, "vswitch_id")

