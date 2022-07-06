# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['UserPermissionArgs', 'UserPermission']

@pulumi.input_type
class UserPermissionArgs:
    def __init__(__self__, *,
                 sub_account_user_id: pulumi.Input[str],
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['UserPermissionPermissionArgs']]]] = None):
        """
        The set of arguments for constructing a UserPermission resource.
        :param pulumi.Input[str] sub_account_user_id: The configuration of the Load Balancer. See the following `Block load_balancer`.
        :param pulumi.Input[Sequence[pulumi.Input['UserPermissionPermissionArgs']]] permissions: List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
        """
        pulumi.set(__self__, "sub_account_user_id", sub_account_user_id)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="subAccountUserId")
    def sub_account_user_id(self) -> pulumi.Input[str]:
        """
        The configuration of the Load Balancer. See the following `Block load_balancer`.
        """
        return pulumi.get(self, "sub_account_user_id")

    @sub_account_user_id.setter
    def sub_account_user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "sub_account_user_id", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserPermissionPermissionArgs']]]]:
        """
        List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserPermissionPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)


@pulumi.input_type
class _UserPermissionState:
    def __init__(__self__, *,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['UserPermissionPermissionArgs']]]] = None,
                 sub_account_user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserPermission resources.
        :param pulumi.Input[Sequence[pulumi.Input['UserPermissionPermissionArgs']]] permissions: List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
        :param pulumi.Input[str] sub_account_user_id: The configuration of the Load Balancer. See the following `Block load_balancer`.
        """
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if sub_account_user_id is not None:
            pulumi.set(__self__, "sub_account_user_id", sub_account_user_id)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserPermissionPermissionArgs']]]]:
        """
        List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserPermissionPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="subAccountUserId")
    def sub_account_user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The configuration of the Load Balancer. See the following `Block load_balancer`.
        """
        return pulumi.get(self, "sub_account_user_id")

    @sub_account_user_id.setter
    def sub_account_user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_account_user_id", value)


class UserPermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserPermissionPermissionArgs']]]]] = None,
                 sub_account_user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Service Mesh UserPermission resource.

        For information about Service Mesh User Permission and how to use it, see [What is User Permission](https://help.aliyun.com/document_detail/171622.html).

        > **NOTE:** Available in v1.174.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "servicemesh"
        default_versions = alicloud.servicemesh.get_versions(edition="Default")
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_networks = alicloud.vpc.get_networks(name_regex="default-NODELETING")
        default_switches = alicloud.vpc.get_switches(vpc_id=default_networks.ids[0],
            zone_id=default_zones.zones[0].id)
        default_user = alicloud.ram.User("defaultUser")
        default1 = alicloud.servicemesh.ServiceMesh("default1",
            service_mesh_name=name,
            edition="Default",
            version=default_versions.versions[0].version,
            cluster_spec="standard",
            network=alicloud.servicemesh.ServiceMeshNetworkArgs(
                vpc_id=default_networks.ids[0],
                vswitche_list=[default_switches.ids[0]],
            ),
            load_balancer=alicloud.servicemesh.ServiceMeshLoadBalancerArgs(
                pilot_public_eip=False,
                api_server_public_eip=False,
            ))
        example = alicloud.servicemesh.UserPermission("example",
            sub_account_user_id=default_user.id,
            permissions=[alicloud.servicemesh.UserPermissionPermissionArgs(
                role_name="istio-admin",
                service_mesh_id=default1.id,
                role_type="custom",
                is_custom=True,
                is_ram_role=False,
            )])
        ```

        ## Import

        Service Mesh User Permission can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:servicemesh/userPermission:UserPermission example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserPermissionPermissionArgs']]]] permissions: List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
        :param pulumi.Input[str] sub_account_user_id: The configuration of the Load Balancer. See the following `Block load_balancer`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserPermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Service Mesh UserPermission resource.

        For information about Service Mesh User Permission and how to use it, see [What is User Permission](https://help.aliyun.com/document_detail/171622.html).

        > **NOTE:** Available in v1.174.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "servicemesh"
        default_versions = alicloud.servicemesh.get_versions(edition="Default")
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_networks = alicloud.vpc.get_networks(name_regex="default-NODELETING")
        default_switches = alicloud.vpc.get_switches(vpc_id=default_networks.ids[0],
            zone_id=default_zones.zones[0].id)
        default_user = alicloud.ram.User("defaultUser")
        default1 = alicloud.servicemesh.ServiceMesh("default1",
            service_mesh_name=name,
            edition="Default",
            version=default_versions.versions[0].version,
            cluster_spec="standard",
            network=alicloud.servicemesh.ServiceMeshNetworkArgs(
                vpc_id=default_networks.ids[0],
                vswitche_list=[default_switches.ids[0]],
            ),
            load_balancer=alicloud.servicemesh.ServiceMeshLoadBalancerArgs(
                pilot_public_eip=False,
                api_server_public_eip=False,
            ))
        example = alicloud.servicemesh.UserPermission("example",
            sub_account_user_id=default_user.id,
            permissions=[alicloud.servicemesh.UserPermissionPermissionArgs(
                role_name="istio-admin",
                service_mesh_id=default1.id,
                role_type="custom",
                is_custom=True,
                is_ram_role=False,
            )])
        ```

        ## Import

        Service Mesh User Permission can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:servicemesh/userPermission:UserPermission example <id>
        ```

        :param str resource_name: The name of the resource.
        :param UserPermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserPermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserPermissionPermissionArgs']]]]] = None,
                 sub_account_user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserPermissionArgs.__new__(UserPermissionArgs)

            __props__.__dict__["permissions"] = permissions
            if sub_account_user_id is None and not opts.urn:
                raise TypeError("Missing required property 'sub_account_user_id'")
            __props__.__dict__["sub_account_user_id"] = sub_account_user_id
        super(UserPermission, __self__).__init__(
            'alicloud:servicemesh/userPermission:UserPermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserPermissionPermissionArgs']]]]] = None,
            sub_account_user_id: Optional[pulumi.Input[str]] = None) -> 'UserPermission':
        """
        Get an existing UserPermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserPermissionPermissionArgs']]]] permissions: List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
        :param pulumi.Input[str] sub_account_user_id: The configuration of the Load Balancer. See the following `Block load_balancer`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserPermissionState.__new__(_UserPermissionState)

        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["sub_account_user_id"] = sub_account_user_id
        return UserPermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence['outputs.UserPermissionPermission']]:
        """
        List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See the following `Block permissions`.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="subAccountUserId")
    def sub_account_user_id(self) -> pulumi.Output[str]:
        """
        The configuration of the Load Balancer. See the following `Block load_balancer`.
        """
        return pulumi.get(self, "sub_account_user_id")

