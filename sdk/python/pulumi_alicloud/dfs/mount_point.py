# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MountPointArgs', 'MountPoint']

@pulumi.input_type
class MountPointArgs:
    def __init__(__self__, *,
                 access_group_id: pulumi.Input[str],
                 file_system_id: pulumi.Input[str],
                 network_type: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str],
                 alias_prefix: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MountPoint resource.
        :param pulumi.Input[str] access_group_id: The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
        :param pulumi.Input[str] file_system_id: Unique file system identifier, used to retrieve specified file system resources.
        :param pulumi.Input[str] network_type: The network type of the Mount point.  Only VPC (VPC) is supported.
        :param pulumi.Input[str] vpc_id: The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
        :param pulumi.Input[str] vswitch_id: VSwitch ID, which specifies the VSwitch resource used to create the mount point.
        :param pulumi.Input[str] alias_prefix: The mount point alias prefix, which specifies the mount point alias prefix.
        :param pulumi.Input[str] description: The description of the Mount point.  No more than 32 characters in length.
        :param pulumi.Input[str] status: Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
        """
        pulumi.set(__self__, "access_group_id", access_group_id)
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "network_type", network_type)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        if alias_prefix is not None:
            pulumi.set(__self__, "alias_prefix", alias_prefix)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> pulumi.Input[str]:
        """
        The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
        """
        return pulumi.get(self, "access_group_id")

    @access_group_id.setter
    def access_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_group_id", value)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Input[str]:
        """
        Unique file system identifier, used to retrieve specified file system resources.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Input[str]:
        """
        The network type of the Mount point.  Only VPC (VPC) is supported.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        VSwitch ID, which specifies the VSwitch resource used to create the mount point.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="aliasPrefix")
    def alias_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The mount point alias prefix, which specifies the mount point alias prefix.
        """
        return pulumi.get(self, "alias_prefix")

    @alias_prefix.setter
    def alias_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias_prefix", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Mount point.  No more than 32 characters in length.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _MountPointState:
    def __init__(__self__, *,
                 access_group_id: Optional[pulumi.Input[str]] = None,
                 alias_prefix: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 mount_point_id: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MountPoint resources.
        :param pulumi.Input[str] access_group_id: The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
        :param pulumi.Input[str] alias_prefix: The mount point alias prefix, which specifies the mount point alias prefix.
        :param pulumi.Input[str] create_time: The creation time of the Mount point resource.
        :param pulumi.Input[str] description: The description of the Mount point.  No more than 32 characters in length.
        :param pulumi.Input[str] file_system_id: Unique file system identifier, used to retrieve specified file system resources.
        :param pulumi.Input[str] mount_point_id: The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
        :param pulumi.Input[str] network_type: The network type of the Mount point.  Only VPC (VPC) is supported.
        :param pulumi.Input[str] status: Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
        :param pulumi.Input[str] vpc_id: The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
        :param pulumi.Input[str] vswitch_id: VSwitch ID, which specifies the VSwitch resource used to create the mount point.
        """
        if access_group_id is not None:
            pulumi.set(__self__, "access_group_id", access_group_id)
        if alias_prefix is not None:
            pulumi.set(__self__, "alias_prefix", alias_prefix)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if file_system_id is not None:
            pulumi.set(__self__, "file_system_id", file_system_id)
        if mount_point_id is not None:
            pulumi.set(__self__, "mount_point_id", mount_point_id)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
        """
        return pulumi.get(self, "access_group_id")

    @access_group_id.setter
    def access_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_group_id", value)

    @property
    @pulumi.getter(name="aliasPrefix")
    def alias_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The mount point alias prefix, which specifies the mount point alias prefix.
        """
        return pulumi.get(self, "alias_prefix")

    @alias_prefix.setter
    def alias_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias_prefix", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the Mount point resource.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Mount point.  No more than 32 characters in length.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique file system identifier, used to retrieve specified file system resources.
        """
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="mountPointId")
    def mount_point_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
        """
        return pulumi.get(self, "mount_point_id")

    @mount_point_id.setter
    def mount_point_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount_point_id", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        The network type of the Mount point.  Only VPC (VPC) is supported.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        VSwitch ID, which specifies the VSwitch resource used to create the mount point.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class MountPoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_group_id: Optional[pulumi.Input[str]] = None,
                 alias_prefix: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a DFS Mount Point resource.

        For information about DFS Mount Point and how to use it, see [What is Mount Point](https://www.alibabacloud.com/help/en/aibaba-cloud-storage-services/latest/apsara-file-storage-for-hdfs).

        > **NOTE:** Available since v1.140.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.dfs.get_zones()
        default_integer = random.index.Integer("default",
            min=10000,
            max=99999)
        default_vpc = alicloud.vpc.Network("DefaultVPC",
            cidr_block="172.16.0.0/12",
            vpc_name=name)
        default_v_switch = alicloud.vpc.Switch("DefaultVSwitch",
            description="example",
            vpc_id=default_vpc.id,
            cidr_block="172.16.0.0/24",
            vswitch_name=name,
            zone_id=default.zones[0].zone_id)
        default_access_group = alicloud.dfs.AccessGroup("DefaultAccessGroup",
            description="AccessGroup resource manager center example",
            network_type="VPC",
            access_group_name=f"{name}-{default_integer['result']}")
        update_access_group = alicloud.dfs.AccessGroup("UpdateAccessGroup",
            description="Second AccessGroup resource manager center example",
            network_type="VPC",
            access_group_name=f"{name}-update-{default_integer['result']}")
        default_fs = alicloud.dfs.FileSystem("DefaultFs",
            space_capacity=1024,
            description="for mountpoint  example",
            storage_type="STANDARD",
            zone_id=default.zones[0].zone_id,
            protocol_type="HDFS",
            data_redundancy_type="LRS",
            file_system_name=f"{name}-{default_integer['result']}")
        default_mount_point = alicloud.dfs.MountPoint("default",
            vpc_id=default_vpc.id,
            description="mountpoint example",
            network_type="VPC",
            vswitch_id=default_v_switch.id,
            file_system_id=default_fs.id,
            access_group_id=default_access_group.id,
            status="Active")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        DFS Mount Point can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dfs/mountPoint:MountPoint example <file_system_id>:<mount_point_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_group_id: The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
        :param pulumi.Input[str] alias_prefix: The mount point alias prefix, which specifies the mount point alias prefix.
        :param pulumi.Input[str] description: The description of the Mount point.  No more than 32 characters in length.
        :param pulumi.Input[str] file_system_id: Unique file system identifier, used to retrieve specified file system resources.
        :param pulumi.Input[str] network_type: The network type of the Mount point.  Only VPC (VPC) is supported.
        :param pulumi.Input[str] status: Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
        :param pulumi.Input[str] vpc_id: The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
        :param pulumi.Input[str] vswitch_id: VSwitch ID, which specifies the VSwitch resource used to create the mount point.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MountPointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a DFS Mount Point resource.

        For information about DFS Mount Point and how to use it, see [What is Mount Point](https://www.alibabacloud.com/help/en/aibaba-cloud-storage-services/latest/apsara-file-storage-for-hdfs).

        > **NOTE:** Available since v1.140.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.dfs.get_zones()
        default_integer = random.index.Integer("default",
            min=10000,
            max=99999)
        default_vpc = alicloud.vpc.Network("DefaultVPC",
            cidr_block="172.16.0.0/12",
            vpc_name=name)
        default_v_switch = alicloud.vpc.Switch("DefaultVSwitch",
            description="example",
            vpc_id=default_vpc.id,
            cidr_block="172.16.0.0/24",
            vswitch_name=name,
            zone_id=default.zones[0].zone_id)
        default_access_group = alicloud.dfs.AccessGroup("DefaultAccessGroup",
            description="AccessGroup resource manager center example",
            network_type="VPC",
            access_group_name=f"{name}-{default_integer['result']}")
        update_access_group = alicloud.dfs.AccessGroup("UpdateAccessGroup",
            description="Second AccessGroup resource manager center example",
            network_type="VPC",
            access_group_name=f"{name}-update-{default_integer['result']}")
        default_fs = alicloud.dfs.FileSystem("DefaultFs",
            space_capacity=1024,
            description="for mountpoint  example",
            storage_type="STANDARD",
            zone_id=default.zones[0].zone_id,
            protocol_type="HDFS",
            data_redundancy_type="LRS",
            file_system_name=f"{name}-{default_integer['result']}")
        default_mount_point = alicloud.dfs.MountPoint("default",
            vpc_id=default_vpc.id,
            description="mountpoint example",
            network_type="VPC",
            vswitch_id=default_v_switch.id,
            file_system_id=default_fs.id,
            access_group_id=default_access_group.id,
            status="Active")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        DFS Mount Point can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dfs/mountPoint:MountPoint example <file_system_id>:<mount_point_id>
        ```

        :param str resource_name: The name of the resource.
        :param MountPointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MountPointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_group_id: Optional[pulumi.Input[str]] = None,
                 alias_prefix: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 file_system_id: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MountPointArgs.__new__(MountPointArgs)

            if access_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'access_group_id'")
            __props__.__dict__["access_group_id"] = access_group_id
            __props__.__dict__["alias_prefix"] = alias_prefix
            __props__.__dict__["description"] = description
            if file_system_id is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_id'")
            __props__.__dict__["file_system_id"] = file_system_id
            if network_type is None and not opts.urn:
                raise TypeError("Missing required property 'network_type'")
            __props__.__dict__["network_type"] = network_type
            __props__.__dict__["status"] = status
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["mount_point_id"] = None
        super(MountPoint, __self__).__init__(
            'alicloud:dfs/mountPoint:MountPoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_group_id: Optional[pulumi.Input[str]] = None,
            alias_prefix: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            file_system_id: Optional[pulumi.Input[str]] = None,
            mount_point_id: Optional[pulumi.Input[str]] = None,
            network_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'MountPoint':
        """
        Get an existing MountPoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_group_id: The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
        :param pulumi.Input[str] alias_prefix: The mount point alias prefix, which specifies the mount point alias prefix.
        :param pulumi.Input[str] create_time: The creation time of the Mount point resource.
        :param pulumi.Input[str] description: The description of the Mount point.  No more than 32 characters in length.
        :param pulumi.Input[str] file_system_id: Unique file system identifier, used to retrieve specified file system resources.
        :param pulumi.Input[str] mount_point_id: The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
        :param pulumi.Input[str] network_type: The network type of the Mount point.  Only VPC (VPC) is supported.
        :param pulumi.Input[str] status: Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
        :param pulumi.Input[str] vpc_id: The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
        :param pulumi.Input[str] vswitch_id: VSwitch ID, which specifies the VSwitch resource used to create the mount point.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MountPointState.__new__(_MountPointState)

        __props__.__dict__["access_group_id"] = access_group_id
        __props__.__dict__["alias_prefix"] = alias_prefix
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["file_system_id"] = file_system_id
        __props__.__dict__["mount_point_id"] = mount_point_id
        __props__.__dict__["network_type"] = network_type
        __props__.__dict__["status"] = status
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vswitch_id"] = vswitch_id
        return MountPoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> pulumi.Output[str]:
        """
        The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
        """
        return pulumi.get(self, "access_group_id")

    @property
    @pulumi.getter(name="aliasPrefix")
    def alias_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The mount point alias prefix, which specifies the mount point alias prefix.
        """
        return pulumi.get(self, "alias_prefix")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of the Mount point resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the Mount point.  No more than 32 characters in length.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Output[str]:
        """
        Unique file system identifier, used to retrieve specified file system resources.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="mountPointId")
    def mount_point_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
        """
        return pulumi.get(self, "mount_point_id")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Output[str]:
        """
        The network type of the Mount point.  Only VPC (VPC) is supported.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        VSwitch ID, which specifies the VSwitch resource used to create the mount point.
        """
        return pulumi.get(self, "vswitch_id")

