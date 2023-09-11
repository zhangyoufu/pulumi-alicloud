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

__all__ = ['VirtualNodeArgs', 'VirtualNode']

@pulumi.input_type
class VirtualNodeArgs:
    def __init__(__self__, *,
                 kube_config: pulumi.Input[str],
                 security_group_id: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str],
                 eip_instance_id: Optional[pulumi.Input[str]] = None,
                 enable_public_network: Optional[pulumi.Input[bool]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 taints: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNodeTaintArgs']]]] = None,
                 virtual_node_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VirtualNode resource.
        :param pulumi.Input[str] kube_config: The kube config for the k8s cluster. It needs to be connected after Base64 encoding.
        :param pulumi.Input[str] security_group_id: The security group ID.
        :param pulumi.Input[str] vswitch_id: The vswitch id.
        :param pulumi.Input[str] eip_instance_id: The Id of eip.
        :param pulumi.Input[bool] enable_public_network: Whether to enable public network. **NOTE:** If `eip_instance_id` is not configured and `enable_public_network` is true, the system will create an elastic public network IP.
        :param pulumi.Input[str] resource_group_id: The resource group ID.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input['VirtualNodeTaintArgs']]] taints: The taint. See `taints` below.
        :param pulumi.Input[str] virtual_node_name: The name of the virtual node. The length of the name is limited to `2` to `128` characters. It can contain uppercase and lowercase letters, Chinese characters, numbers, half-width colon (:), underscores (_), or hyphens (-), and must start with letters.
        :param pulumi.Input[str] zone_id: The Zone.
        """
        pulumi.set(__self__, "kube_config", kube_config)
        pulumi.set(__self__, "security_group_id", security_group_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        if eip_instance_id is not None:
            pulumi.set(__self__, "eip_instance_id", eip_instance_id)
        if enable_public_network is not None:
            pulumi.set(__self__, "enable_public_network", enable_public_network)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if taints is not None:
            pulumi.set(__self__, "taints", taints)
        if virtual_node_name is not None:
            pulumi.set(__self__, "virtual_node_name", virtual_node_name)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="kubeConfig")
    def kube_config(self) -> pulumi.Input[str]:
        """
        The kube config for the k8s cluster. It needs to be connected after Base64 encoding.
        """
        return pulumi.get(self, "kube_config")

    @kube_config.setter
    def kube_config(self, value: pulumi.Input[str]):
        pulumi.set(self, "kube_config", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Input[str]:
        """
        The security group ID.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        The vswitch id.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="eipInstanceId")
    def eip_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of eip.
        """
        return pulumi.get(self, "eip_instance_id")

    @eip_instance_id.setter
    def eip_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_instance_id", value)

    @property
    @pulumi.getter(name="enablePublicNetwork")
    def enable_public_network(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable public network. **NOTE:** If `eip_instance_id` is not configured and `enable_public_network` is true, the system will create an elastic public network IP.
        """
        return pulumi.get(self, "enable_public_network")

    @enable_public_network.setter
    def enable_public_network(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_public_network", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def taints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNodeTaintArgs']]]]:
        """
        The taint. See `taints` below.
        """
        return pulumi.get(self, "taints")

    @taints.setter
    def taints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNodeTaintArgs']]]]):
        pulumi.set(self, "taints", value)

    @property
    @pulumi.getter(name="virtualNodeName")
    def virtual_node_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the virtual node. The length of the name is limited to `2` to `128` characters. It can contain uppercase and lowercase letters, Chinese characters, numbers, half-width colon (:), underscores (_), or hyphens (-), and must start with letters.
        """
        return pulumi.get(self, "virtual_node_name")

    @virtual_node_name.setter
    def virtual_node_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_node_name", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Zone.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


@pulumi.input_type
class _VirtualNodeState:
    def __init__(__self__, *,
                 eip_instance_id: Optional[pulumi.Input[str]] = None,
                 enable_public_network: Optional[pulumi.Input[bool]] = None,
                 kube_config: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 taints: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNodeTaintArgs']]]] = None,
                 virtual_node_name: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VirtualNode resources.
        :param pulumi.Input[str] eip_instance_id: The Id of eip.
        :param pulumi.Input[bool] enable_public_network: Whether to enable public network. **NOTE:** If `eip_instance_id` is not configured and `enable_public_network` is true, the system will create an elastic public network IP.
        :param pulumi.Input[str] kube_config: The kube config for the k8s cluster. It needs to be connected after Base64 encoding.
        :param pulumi.Input[str] resource_group_id: The resource group ID.
        :param pulumi.Input[str] security_group_id: The security group ID.
        :param pulumi.Input[str] status: The Status of the virtual node. Valid values: `Cleaned`, `Failed`, `Pending`, `Ready`.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input['VirtualNodeTaintArgs']]] taints: The taint. See `taints` below.
        :param pulumi.Input[str] virtual_node_name: The name of the virtual node. The length of the name is limited to `2` to `128` characters. It can contain uppercase and lowercase letters, Chinese characters, numbers, half-width colon (:), underscores (_), or hyphens (-), and must start with letters.
        :param pulumi.Input[str] vswitch_id: The vswitch id.
        :param pulumi.Input[str] zone_id: The Zone.
        """
        if eip_instance_id is not None:
            pulumi.set(__self__, "eip_instance_id", eip_instance_id)
        if enable_public_network is not None:
            pulumi.set(__self__, "enable_public_network", enable_public_network)
        if kube_config is not None:
            pulumi.set(__self__, "kube_config", kube_config)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if taints is not None:
            pulumi.set(__self__, "taints", taints)
        if virtual_node_name is not None:
            pulumi.set(__self__, "virtual_node_name", virtual_node_name)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="eipInstanceId")
    def eip_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of eip.
        """
        return pulumi.get(self, "eip_instance_id")

    @eip_instance_id.setter
    def eip_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_instance_id", value)

    @property
    @pulumi.getter(name="enablePublicNetwork")
    def enable_public_network(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable public network. **NOTE:** If `eip_instance_id` is not configured and `enable_public_network` is true, the system will create an elastic public network IP.
        """
        return pulumi.get(self, "enable_public_network")

    @enable_public_network.setter
    def enable_public_network(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_public_network", value)

    @property
    @pulumi.getter(name="kubeConfig")
    def kube_config(self) -> Optional[pulumi.Input[str]]:
        """
        The kube config for the k8s cluster. It needs to be connected after Base64 encoding.
        """
        return pulumi.get(self, "kube_config")

    @kube_config.setter
    def kube_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kube_config", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The security group ID.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The Status of the virtual node. Valid values: `Cleaned`, `Failed`, `Pending`, `Ready`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def taints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNodeTaintArgs']]]]:
        """
        The taint. See `taints` below.
        """
        return pulumi.get(self, "taints")

    @taints.setter
    def taints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualNodeTaintArgs']]]]):
        pulumi.set(self, "taints", value)

    @property
    @pulumi.getter(name="virtualNodeName")
    def virtual_node_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the virtual node. The length of the name is limited to `2` to `128` characters. It can contain uppercase and lowercase letters, Chinese characters, numbers, half-width colon (:), underscores (_), or hyphens (-), and must start with letters.
        """
        return pulumi.get(self, "virtual_node_name")

    @virtual_node_name.setter
    def virtual_node_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_node_name", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The vswitch id.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Zone.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class VirtualNode(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eip_instance_id: Optional[pulumi.Input[str]] = None,
                 enable_public_network: Optional[pulumi.Input[bool]] = None,
                 kube_config: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 taints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNodeTaintArgs']]]]] = None,
                 virtual_node_name: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ECI Virtual Node resource.

        For information about ECI Virtual Node and how to use it, see [What is Virtual Node](https://www.alibabacloud.com/help/en/doc-detail/89129.html).

        > **NOTE:** Available since v1.145.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_zones = alicloud.eci.get_zones()
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="10.0.0.0/8")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name=name,
            cidr_block="10.1.0.0/16",
            vpc_id=default_network.id,
            zone_id=default_zones.zones[0].zone_ids[0])
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_eip_address = alicloud.ecs.EipAddress("defaultEipAddress",
            isp="BGP",
            address_name=name,
            netmode="public",
            bandwidth="1",
            security_protection_types=["AntiDDoS_Enhanced"],
            payment_type="PayAsYouGo")
        default_resource_groups = alicloud.resourcemanager.get_resource_groups()
        default_virtual_node = alicloud.eci.VirtualNode("defaultVirtualNode",
            security_group_id=default_security_group.id,
            virtual_node_name=name,
            vswitch_id=default_switch.id,
            enable_public_network=False,
            eip_instance_id=default_eip_address.id,
            resource_group_id=default_resource_groups.groups[0].id,
            kube_config="kube_config",
            tags={
                "Created": "TF",
            },
            taints=[alicloud.eci.VirtualNodeTaintArgs(
                effect="NoSchedule",
                key="TF",
                value="example",
            )])
        ```

        ## Import

        ECI Virtual Node can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:eci/virtualNode:VirtualNode example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eip_instance_id: The Id of eip.
        :param pulumi.Input[bool] enable_public_network: Whether to enable public network. **NOTE:** If `eip_instance_id` is not configured and `enable_public_network` is true, the system will create an elastic public network IP.
        :param pulumi.Input[str] kube_config: The kube config for the k8s cluster. It needs to be connected after Base64 encoding.
        :param pulumi.Input[str] resource_group_id: The resource group ID.
        :param pulumi.Input[str] security_group_id: The security group ID.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNodeTaintArgs']]]] taints: The taint. See `taints` below.
        :param pulumi.Input[str] virtual_node_name: The name of the virtual node. The length of the name is limited to `2` to `128` characters. It can contain uppercase and lowercase letters, Chinese characters, numbers, half-width colon (:), underscores (_), or hyphens (-), and must start with letters.
        :param pulumi.Input[str] vswitch_id: The vswitch id.
        :param pulumi.Input[str] zone_id: The Zone.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualNodeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ECI Virtual Node resource.

        For information about ECI Virtual Node and how to use it, see [What is Virtual Node](https://www.alibabacloud.com/help/en/doc-detail/89129.html).

        > **NOTE:** Available since v1.145.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_zones = alicloud.eci.get_zones()
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="10.0.0.0/8")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name=name,
            cidr_block="10.1.0.0/16",
            vpc_id=default_network.id,
            zone_id=default_zones.zones[0].zone_ids[0])
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_eip_address = alicloud.ecs.EipAddress("defaultEipAddress",
            isp="BGP",
            address_name=name,
            netmode="public",
            bandwidth="1",
            security_protection_types=["AntiDDoS_Enhanced"],
            payment_type="PayAsYouGo")
        default_resource_groups = alicloud.resourcemanager.get_resource_groups()
        default_virtual_node = alicloud.eci.VirtualNode("defaultVirtualNode",
            security_group_id=default_security_group.id,
            virtual_node_name=name,
            vswitch_id=default_switch.id,
            enable_public_network=False,
            eip_instance_id=default_eip_address.id,
            resource_group_id=default_resource_groups.groups[0].id,
            kube_config="kube_config",
            tags={
                "Created": "TF",
            },
            taints=[alicloud.eci.VirtualNodeTaintArgs(
                effect="NoSchedule",
                key="TF",
                value="example",
            )])
        ```

        ## Import

        ECI Virtual Node can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:eci/virtualNode:VirtualNode example <id>
        ```

        :param str resource_name: The name of the resource.
        :param VirtualNodeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualNodeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eip_instance_id: Optional[pulumi.Input[str]] = None,
                 enable_public_network: Optional[pulumi.Input[bool]] = None,
                 kube_config: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 taints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNodeTaintArgs']]]]] = None,
                 virtual_node_name: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualNodeArgs.__new__(VirtualNodeArgs)

            __props__.__dict__["eip_instance_id"] = eip_instance_id
            __props__.__dict__["enable_public_network"] = enable_public_network
            if kube_config is None and not opts.urn:
                raise TypeError("Missing required property 'kube_config'")
            __props__.__dict__["kube_config"] = kube_config
            __props__.__dict__["resource_group_id"] = resource_group_id
            if security_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_id'")
            __props__.__dict__["security_group_id"] = security_group_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["taints"] = taints
            __props__.__dict__["virtual_node_name"] = virtual_node_name
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["status"] = None
        super(VirtualNode, __self__).__init__(
            'alicloud:eci/virtualNode:VirtualNode',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            eip_instance_id: Optional[pulumi.Input[str]] = None,
            enable_public_network: Optional[pulumi.Input[bool]] = None,
            kube_config: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            taints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNodeTaintArgs']]]]] = None,
            virtual_node_name: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'VirtualNode':
        """
        Get an existing VirtualNode resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eip_instance_id: The Id of eip.
        :param pulumi.Input[bool] enable_public_network: Whether to enable public network. **NOTE:** If `eip_instance_id` is not configured and `enable_public_network` is true, the system will create an elastic public network IP.
        :param pulumi.Input[str] kube_config: The kube config for the k8s cluster. It needs to be connected after Base64 encoding.
        :param pulumi.Input[str] resource_group_id: The resource group ID.
        :param pulumi.Input[str] security_group_id: The security group ID.
        :param pulumi.Input[str] status: The Status of the virtual node. Valid values: `Cleaned`, `Failed`, `Pending`, `Ready`.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNodeTaintArgs']]]] taints: The taint. See `taints` below.
        :param pulumi.Input[str] virtual_node_name: The name of the virtual node. The length of the name is limited to `2` to `128` characters. It can contain uppercase and lowercase letters, Chinese characters, numbers, half-width colon (:), underscores (_), or hyphens (-), and must start with letters.
        :param pulumi.Input[str] vswitch_id: The vswitch id.
        :param pulumi.Input[str] zone_id: The Zone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VirtualNodeState.__new__(_VirtualNodeState)

        __props__.__dict__["eip_instance_id"] = eip_instance_id
        __props__.__dict__["enable_public_network"] = enable_public_network
        __props__.__dict__["kube_config"] = kube_config
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["security_group_id"] = security_group_id
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["taints"] = taints
        __props__.__dict__["virtual_node_name"] = virtual_node_name
        __props__.__dict__["vswitch_id"] = vswitch_id
        __props__.__dict__["zone_id"] = zone_id
        return VirtualNode(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eipInstanceId")
    def eip_instance_id(self) -> pulumi.Output[str]:
        """
        The Id of eip.
        """
        return pulumi.get(self, "eip_instance_id")

    @property
    @pulumi.getter(name="enablePublicNetwork")
    def enable_public_network(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable public network. **NOTE:** If `eip_instance_id` is not configured and `enable_public_network` is true, the system will create an elastic public network IP.
        """
        return pulumi.get(self, "enable_public_network")

    @property
    @pulumi.getter(name="kubeConfig")
    def kube_config(self) -> pulumi.Output[str]:
        """
        The kube config for the k8s cluster. It needs to be connected after Base64 encoding.
        """
        return pulumi.get(self, "kube_config")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        The security group ID.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The Status of the virtual node. Valid values: `Cleaned`, `Failed`, `Pending`, `Ready`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def taints(self) -> pulumi.Output[Optional[Sequence['outputs.VirtualNodeTaint']]]:
        """
        The taint. See `taints` below.
        """
        return pulumi.get(self, "taints")

    @property
    @pulumi.getter(name="virtualNodeName")
    def virtual_node_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the virtual node. The length of the name is limited to `2` to `128` characters. It can contain uppercase and lowercase letters, Chinese characters, numbers, half-width colon (:), underscores (_), or hyphens (-), and must start with letters.
        """
        return pulumi.get(self, "virtual_node_name")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The vswitch id.
        """
        return pulumi.get(self, "vswitch_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The Zone.
        """
        return pulumi.get(self, "zone_id")

