# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EipAssociationArgs', 'EipAssociation']

@pulumi.input_type
class EipAssociationArgs:
    def __init__(__self__, *,
                 allocation_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 force: Optional[pulumi.Input[bool]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EipAssociation resource.
        :param pulumi.Input[str] allocation_id: The ID of the EIP that you want to associate with an instance.
        :param pulumi.Input[str] instance_id: The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
        :param pulumi.Input[bool] force: When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
        :param pulumi.Input[str] instance_type: The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
        :param pulumi.Input[str] mode: The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instance_type` is set to `NetworkInterface`.
        :param pulumi.Input[str] private_ip_address: The IP address in the CIDR block of the vSwitch.
        :param pulumi.Input[str] vpc_id: The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instance_type` is set to `IpAddress`.
        """
        pulumi.set(__self__, "allocation_id", allocation_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if force is not None:
            pulumi.set(__self__, "force", force)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if private_ip_address is not None:
            pulumi.set(__self__, "private_ip_address", private_ip_address)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> pulumi.Input[str]:
        """
        The ID of the EIP that you want to associate with an instance.
        """
        return pulumi.get(self, "allocation_id")

    @allocation_id.setter
    def allocation_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "allocation_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def force(self) -> Optional[pulumi.Input[bool]]:
        """
        When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
        """
        return pulumi.get(self, "force")

    @force.setter
    def force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instance_type` is set to `NetworkInterface`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address in the CIDR block of the vSwitch.
        """
        return pulumi.get(self, "private_ip_address")

    @private_ip_address.setter
    def private_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip_address", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instance_type` is set to `IpAddress`.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class _EipAssociationState:
    def __init__(__self__, *,
                 allocation_id: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EipAssociation resources.
        :param pulumi.Input[str] allocation_id: The ID of the EIP that you want to associate with an instance.
        :param pulumi.Input[bool] force: When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
        :param pulumi.Input[str] instance_id: The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
        :param pulumi.Input[str] instance_type: The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
        :param pulumi.Input[str] mode: The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instance_type` is set to `NetworkInterface`.
        :param pulumi.Input[str] private_ip_address: The IP address in the CIDR block of the vSwitch.
        :param pulumi.Input[str] vpc_id: The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instance_type` is set to `IpAddress`.
        """
        if allocation_id is not None:
            pulumi.set(__self__, "allocation_id", allocation_id)
        if force is not None:
            pulumi.set(__self__, "force", force)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if private_ip_address is not None:
            pulumi.set(__self__, "private_ip_address", private_ip_address)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the EIP that you want to associate with an instance.
        """
        return pulumi.get(self, "allocation_id")

    @allocation_id.setter
    def allocation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allocation_id", value)

    @property
    @pulumi.getter
    def force(self) -> Optional[pulumi.Input[bool]]:
        """
        When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
        """
        return pulumi.get(self, "force")

    @force.setter
    def force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instance_type` is set to `NetworkInterface`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address in the CIDR block of the vSwitch.
        """
        return pulumi.get(self, "private_ip_address")

    @private_ip_address.setter
    def private_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip_address", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instance_type` is set to `IpAddress`.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class EipAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_id: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Alicloud EIP Association resource for associating Elastic IP to ECS Instance, SLB Instance or Nat Gateway.

        > **NOTE:** `ecs.EipAssociation` is useful in scenarios where EIPs are either
         pre-existing or distributed to customers or users and therefore cannot be changed.

        > **NOTE:** From version 1.7.1, the resource support to associate EIP to SLB Instance or Nat Gateway.

        > **NOTE:** One EIP can only be associated with ECS or SLB instance which in the VPC.

        > **NOTE:** Available since v1.117.0.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        example_zones = alicloud.get_zones(available_resource_creation="Instance")
        example_instance_types = alicloud.ecs.get_instance_types(availability_zone=example_zones.zones[0].id,
            cpu_core_count=1,
            memory_size=2)
        example_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            owners="system")
        example_network = alicloud.vpc.Network("exampleNetwork",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        example_switch = alicloud.vpc.Switch("exampleSwitch",
            vswitch_name=name,
            cidr_block="10.4.0.0/24",
            vpc_id=example_network.id,
            zone_id=example_zones.zones[0].id)
        example_security_group = alicloud.ecs.SecurityGroup("exampleSecurityGroup", vpc_id=example_network.id)
        example_instance = alicloud.ecs.Instance("exampleInstance",
            availability_zone=example_zones.zones[0].id,
            instance_name=name,
            image_id=example_images.images[0].id,
            instance_type=example_instance_types.instance_types[0].id,
            security_groups=[example_security_group.id],
            vswitch_id=example_switch.id,
            tags={
                "Created": "TF",
                "For": "example",
            })
        example_eip_address = alicloud.ecs.EipAddress("exampleEipAddress", address_name=name)
        example_eip_association = alicloud.ecs.EipAssociation("exampleEipAssociation",
            allocation_id=example_eip_address.id,
            instance_id=example_instance.id)
        ```
        ## Module Support

        You can use the existing eip module
        to create several EIP instances and associate them with other resources one-click, like ECS instances, SLB, Nat Gateway and so on.

        ## Import

        Elastic IP address association can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/eipAssociation:EipAssociation example <allocation_id>:<instance_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_id: The ID of the EIP that you want to associate with an instance.
        :param pulumi.Input[bool] force: When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
        :param pulumi.Input[str] instance_id: The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
        :param pulumi.Input[str] instance_type: The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
        :param pulumi.Input[str] mode: The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instance_type` is set to `NetworkInterface`.
        :param pulumi.Input[str] private_ip_address: The IP address in the CIDR block of the vSwitch.
        :param pulumi.Input[str] vpc_id: The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instance_type` is set to `IpAddress`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EipAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Alicloud EIP Association resource for associating Elastic IP to ECS Instance, SLB Instance or Nat Gateway.

        > **NOTE:** `ecs.EipAssociation` is useful in scenarios where EIPs are either
         pre-existing or distributed to customers or users and therefore cannot be changed.

        > **NOTE:** From version 1.7.1, the resource support to associate EIP to SLB Instance or Nat Gateway.

        > **NOTE:** One EIP can only be associated with ECS or SLB instance which in the VPC.

        > **NOTE:** Available since v1.117.0.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        example_zones = alicloud.get_zones(available_resource_creation="Instance")
        example_instance_types = alicloud.ecs.get_instance_types(availability_zone=example_zones.zones[0].id,
            cpu_core_count=1,
            memory_size=2)
        example_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            owners="system")
        example_network = alicloud.vpc.Network("exampleNetwork",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        example_switch = alicloud.vpc.Switch("exampleSwitch",
            vswitch_name=name,
            cidr_block="10.4.0.0/24",
            vpc_id=example_network.id,
            zone_id=example_zones.zones[0].id)
        example_security_group = alicloud.ecs.SecurityGroup("exampleSecurityGroup", vpc_id=example_network.id)
        example_instance = alicloud.ecs.Instance("exampleInstance",
            availability_zone=example_zones.zones[0].id,
            instance_name=name,
            image_id=example_images.images[0].id,
            instance_type=example_instance_types.instance_types[0].id,
            security_groups=[example_security_group.id],
            vswitch_id=example_switch.id,
            tags={
                "Created": "TF",
                "For": "example",
            })
        example_eip_address = alicloud.ecs.EipAddress("exampleEipAddress", address_name=name)
        example_eip_association = alicloud.ecs.EipAssociation("exampleEipAssociation",
            allocation_id=example_eip_address.id,
            instance_id=example_instance.id)
        ```
        ## Module Support

        You can use the existing eip module
        to create several EIP instances and associate them with other resources one-click, like ECS instances, SLB, Nat Gateway and so on.

        ## Import

        Elastic IP address association can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/eipAssociation:EipAssociation example <allocation_id>:<instance_id>
        ```

        :param str resource_name: The name of the resource.
        :param EipAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EipAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_id: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EipAssociationArgs.__new__(EipAssociationArgs)

            if allocation_id is None and not opts.urn:
                raise TypeError("Missing required property 'allocation_id'")
            __props__.__dict__["allocation_id"] = allocation_id
            __props__.__dict__["force"] = force
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["mode"] = mode
            __props__.__dict__["private_ip_address"] = private_ip_address
            __props__.__dict__["vpc_id"] = vpc_id
        super(EipAssociation, __self__).__init__(
            'alicloud:ecs/eipAssociation:EipAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allocation_id: Optional[pulumi.Input[str]] = None,
            force: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            private_ip_address: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'EipAssociation':
        """
        Get an existing EipAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_id: The ID of the EIP that you want to associate with an instance.
        :param pulumi.Input[bool] force: When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
        :param pulumi.Input[str] instance_id: The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
        :param pulumi.Input[str] instance_type: The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
        :param pulumi.Input[str] mode: The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instance_type` is set to `NetworkInterface`.
        :param pulumi.Input[str] private_ip_address: The IP address in the CIDR block of the vSwitch.
        :param pulumi.Input[str] vpc_id: The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instance_type` is set to `IpAddress`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EipAssociationState.__new__(_EipAssociationState)

        __props__.__dict__["allocation_id"] = allocation_id
        __props__.__dict__["force"] = force
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["instance_type"] = instance_type
        __props__.__dict__["mode"] = mode
        __props__.__dict__["private_ip_address"] = private_ip_address
        __props__.__dict__["vpc_id"] = vpc_id
        return EipAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> pulumi.Output[str]:
        """
        The ID of the EIP that you want to associate with an instance.
        """
        return pulumi.get(self, "allocation_id")

    @property
    @pulumi.getter
    def force(self) -> pulumi.Output[Optional[bool]]:
        """
        When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
        """
        return pulumi.get(self, "force")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instance_type` is set to `NetworkInterface`.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> pulumi.Output[Optional[str]]:
        """
        The IP address in the CIDR block of the vSwitch.
        """
        return pulumi.get(self, "private_ip_address")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instance_type` is set to `IpAddress`.
        """
        return pulumi.get(self, "vpc_id")

