# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RamDirectoryArgs', 'RamDirectory']

@pulumi.input_type
class RamDirectoryArgs:
    def __init__(__self__, *,
                 ram_directory_name: pulumi.Input[str],
                 vswitch_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 desktop_access_type: Optional[pulumi.Input[str]] = None,
                 enable_admin_access: Optional[pulumi.Input[bool]] = None,
                 enable_internet_access: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a RamDirectory resource.
        :param pulumi.Input[str] ram_directory_name: The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: List of VSwitch IDs in the directory.
        :param pulumi.Input[str] desktop_access_type: The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        :param pulumi.Input[bool] enable_admin_access: Whether to enable public network access.
        :param pulumi.Input[bool] enable_internet_access: Whether to grant local administrator rights to users who use cloud desktops.
        """
        pulumi.set(__self__, "ram_directory_name", ram_directory_name)
        pulumi.set(__self__, "vswitch_ids", vswitch_ids)
        if desktop_access_type is not None:
            pulumi.set(__self__, "desktop_access_type", desktop_access_type)
        if enable_admin_access is not None:
            pulumi.set(__self__, "enable_admin_access", enable_admin_access)
        if enable_internet_access is not None:
            pulumi.set(__self__, "enable_internet_access", enable_internet_access)

    @property
    @pulumi.getter(name="ramDirectoryName")
    def ram_directory_name(self) -> pulumi.Input[str]:
        """
        The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "ram_directory_name")

    @ram_directory_name.setter
    def ram_directory_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "ram_directory_name", value)

    @property
    @pulumi.getter(name="vswitchIds")
    def vswitch_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of VSwitch IDs in the directory.
        """
        return pulumi.get(self, "vswitch_ids")

    @vswitch_ids.setter
    def vswitch_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "vswitch_ids", value)

    @property
    @pulumi.getter(name="desktopAccessType")
    def desktop_access_type(self) -> Optional[pulumi.Input[str]]:
        """
        The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        """
        return pulumi.get(self, "desktop_access_type")

    @desktop_access_type.setter
    def desktop_access_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "desktop_access_type", value)

    @property
    @pulumi.getter(name="enableAdminAccess")
    def enable_admin_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable public network access.
        """
        return pulumi.get(self, "enable_admin_access")

    @enable_admin_access.setter
    def enable_admin_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_admin_access", value)

    @property
    @pulumi.getter(name="enableInternetAccess")
    def enable_internet_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to grant local administrator rights to users who use cloud desktops.
        """
        return pulumi.get(self, "enable_internet_access")

    @enable_internet_access.setter
    def enable_internet_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_internet_access", value)


@pulumi.input_type
class _RamDirectoryState:
    def __init__(__self__, *,
                 desktop_access_type: Optional[pulumi.Input[str]] = None,
                 enable_admin_access: Optional[pulumi.Input[bool]] = None,
                 enable_internet_access: Optional[pulumi.Input[bool]] = None,
                 ram_directory_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering RamDirectory resources.
        :param pulumi.Input[str] desktop_access_type: The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        :param pulumi.Input[bool] enable_admin_access: Whether to enable public network access.
        :param pulumi.Input[bool] enable_internet_access: Whether to grant local administrator rights to users who use cloud desktops.
        :param pulumi.Input[str] ram_directory_name: The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        :param pulumi.Input[str] status: The status of directory.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: List of VSwitch IDs in the directory.
        """
        if desktop_access_type is not None:
            pulumi.set(__self__, "desktop_access_type", desktop_access_type)
        if enable_admin_access is not None:
            pulumi.set(__self__, "enable_admin_access", enable_admin_access)
        if enable_internet_access is not None:
            pulumi.set(__self__, "enable_internet_access", enable_internet_access)
        if ram_directory_name is not None:
            pulumi.set(__self__, "ram_directory_name", ram_directory_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vswitch_ids is not None:
            pulumi.set(__self__, "vswitch_ids", vswitch_ids)

    @property
    @pulumi.getter(name="desktopAccessType")
    def desktop_access_type(self) -> Optional[pulumi.Input[str]]:
        """
        The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        """
        return pulumi.get(self, "desktop_access_type")

    @desktop_access_type.setter
    def desktop_access_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "desktop_access_type", value)

    @property
    @pulumi.getter(name="enableAdminAccess")
    def enable_admin_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable public network access.
        """
        return pulumi.get(self, "enable_admin_access")

    @enable_admin_access.setter
    def enable_admin_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_admin_access", value)

    @property
    @pulumi.getter(name="enableInternetAccess")
    def enable_internet_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to grant local administrator rights to users who use cloud desktops.
        """
        return pulumi.get(self, "enable_internet_access")

    @enable_internet_access.setter
    def enable_internet_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_internet_access", value)

    @property
    @pulumi.getter(name="ramDirectoryName")
    def ram_directory_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "ram_directory_name")

    @ram_directory_name.setter
    def ram_directory_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ram_directory_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of directory.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vswitchIds")
    def vswitch_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of VSwitch IDs in the directory.
        """
        return pulumi.get(self, "vswitch_ids")

    @vswitch_ids.setter
    def vswitch_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vswitch_ids", value)


class RamDirectory(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 desktop_access_type: Optional[pulumi.Input[str]] = None,
                 enable_admin_access: Optional[pulumi.Input[bool]] = None,
                 enable_internet_access: Optional[pulumi.Input[bool]] = None,
                 ram_directory_name: Optional[pulumi.Input[str]] = None,
                 vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a ECD Ram Directory resource.

        For information about ECD Ram Directory and how to use it, see [What is Ram Directory](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createramdirectory).

        > **NOTE:** Available since v1.174.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_zones = alicloud.eds.get_zones()
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            zone_id=default_zones.ids[0],
            vswitch_name=name)
        default_ram_directory = alicloud.eds.RamDirectory("defaultRamDirectory",
            desktop_access_type="INTERNET",
            enable_admin_access=True,
            enable_internet_access=True,
            ram_directory_name=name,
            vswitch_ids=[default_switch.id])
        ```

        ## Import

        ECD Ram Directory can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:eds/ramDirectory:RamDirectory example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] desktop_access_type: The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        :param pulumi.Input[bool] enable_admin_access: Whether to enable public network access.
        :param pulumi.Input[bool] enable_internet_access: Whether to grant local administrator rights to users who use cloud desktops.
        :param pulumi.Input[str] ram_directory_name: The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: List of VSwitch IDs in the directory.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RamDirectoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ECD Ram Directory resource.

        For information about ECD Ram Directory and how to use it, see [What is Ram Directory](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createramdirectory).

        > **NOTE:** Available since v1.174.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_zones = alicloud.eds.get_zones()
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            zone_id=default_zones.ids[0],
            vswitch_name=name)
        default_ram_directory = alicloud.eds.RamDirectory("defaultRamDirectory",
            desktop_access_type="INTERNET",
            enable_admin_access=True,
            enable_internet_access=True,
            ram_directory_name=name,
            vswitch_ids=[default_switch.id])
        ```

        ## Import

        ECD Ram Directory can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:eds/ramDirectory:RamDirectory example <id>
        ```

        :param str resource_name: The name of the resource.
        :param RamDirectoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RamDirectoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 desktop_access_type: Optional[pulumi.Input[str]] = None,
                 enable_admin_access: Optional[pulumi.Input[bool]] = None,
                 enable_internet_access: Optional[pulumi.Input[bool]] = None,
                 ram_directory_name: Optional[pulumi.Input[str]] = None,
                 vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RamDirectoryArgs.__new__(RamDirectoryArgs)

            __props__.__dict__["desktop_access_type"] = desktop_access_type
            __props__.__dict__["enable_admin_access"] = enable_admin_access
            __props__.__dict__["enable_internet_access"] = enable_internet_access
            if ram_directory_name is None and not opts.urn:
                raise TypeError("Missing required property 'ram_directory_name'")
            __props__.__dict__["ram_directory_name"] = ram_directory_name
            if vswitch_ids is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_ids'")
            __props__.__dict__["vswitch_ids"] = vswitch_ids
            __props__.__dict__["status"] = None
        super(RamDirectory, __self__).__init__(
            'alicloud:eds/ramDirectory:RamDirectory',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            desktop_access_type: Optional[pulumi.Input[str]] = None,
            enable_admin_access: Optional[pulumi.Input[bool]] = None,
            enable_internet_access: Optional[pulumi.Input[bool]] = None,
            ram_directory_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'RamDirectory':
        """
        Get an existing RamDirectory resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] desktop_access_type: The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        :param pulumi.Input[bool] enable_admin_access: Whether to enable public network access.
        :param pulumi.Input[bool] enable_internet_access: Whether to grant local administrator rights to users who use cloud desktops.
        :param pulumi.Input[str] ram_directory_name: The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        :param pulumi.Input[str] status: The status of directory.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: List of VSwitch IDs in the directory.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RamDirectoryState.__new__(_RamDirectoryState)

        __props__.__dict__["desktop_access_type"] = desktop_access_type
        __props__.__dict__["enable_admin_access"] = enable_admin_access
        __props__.__dict__["enable_internet_access"] = enable_internet_access
        __props__.__dict__["ram_directory_name"] = ram_directory_name
        __props__.__dict__["status"] = status
        __props__.__dict__["vswitch_ids"] = vswitch_ids
        return RamDirectory(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="desktopAccessType")
    def desktop_access_type(self) -> pulumi.Output[str]:
        """
        The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        """
        return pulumi.get(self, "desktop_access_type")

    @property
    @pulumi.getter(name="enableAdminAccess")
    def enable_admin_access(self) -> pulumi.Output[bool]:
        """
        Whether to enable public network access.
        """
        return pulumi.get(self, "enable_admin_access")

    @property
    @pulumi.getter(name="enableInternetAccess")
    def enable_internet_access(self) -> pulumi.Output[bool]:
        """
        Whether to grant local administrator rights to users who use cloud desktops.
        """
        return pulumi.get(self, "enable_internet_access")

    @property
    @pulumi.getter(name="ramDirectoryName")
    def ram_directory_name(self) -> pulumi.Output[str]:
        """
        The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "ram_directory_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of directory.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vswitchIds")
    def vswitch_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        List of VSwitch IDs in the directory.
        """
        return pulumi.get(self, "vswitch_ids")

