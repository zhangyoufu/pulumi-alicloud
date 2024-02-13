# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EcsNetworkInterfaceAttachmentArgs', 'EcsNetworkInterfaceAttachment']

@pulumi.input_type
class EcsNetworkInterfaceAttachmentArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 network_interface_id: pulumi.Input[str],
                 trunk_network_instance_id: Optional[pulumi.Input[str]] = None,
                 wait_for_network_configuration_ready: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a EcsNetworkInterfaceAttachment resource.
        :param pulumi.Input[str] instance_id: The instance id.
        :param pulumi.Input[str] network_interface_id: The network interface id.
        :param pulumi.Input[str] trunk_network_instance_id: The trunk network instance id.
        :param pulumi.Input[bool] wait_for_network_configuration_ready: The wait for network configuration ready.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        if trunk_network_instance_id is not None:
            pulumi.set(__self__, "trunk_network_instance_id", trunk_network_instance_id)
        if wait_for_network_configuration_ready is not None:
            pulumi.set(__self__, "wait_for_network_configuration_ready", wait_for_network_configuration_ready)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Input[str]:
        """
        The network interface id.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="trunkNetworkInstanceId")
    def trunk_network_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The trunk network instance id.
        """
        return pulumi.get(self, "trunk_network_instance_id")

    @trunk_network_instance_id.setter
    def trunk_network_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trunk_network_instance_id", value)

    @property
    @pulumi.getter(name="waitForNetworkConfigurationReady")
    def wait_for_network_configuration_ready(self) -> Optional[pulumi.Input[bool]]:
        """
        The wait for network configuration ready.
        """
        return pulumi.get(self, "wait_for_network_configuration_ready")

    @wait_for_network_configuration_ready.setter
    def wait_for_network_configuration_ready(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_for_network_configuration_ready", value)


@pulumi.input_type
class _EcsNetworkInterfaceAttachmentState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 trunk_network_instance_id: Optional[pulumi.Input[str]] = None,
                 wait_for_network_configuration_ready: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering EcsNetworkInterfaceAttachment resources.
        :param pulumi.Input[str] instance_id: The instance id.
        :param pulumi.Input[str] network_interface_id: The network interface id.
        :param pulumi.Input[str] trunk_network_instance_id: The trunk network instance id.
        :param pulumi.Input[bool] wait_for_network_configuration_ready: The wait for network configuration ready.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)
        if trunk_network_instance_id is not None:
            pulumi.set(__self__, "trunk_network_instance_id", trunk_network_instance_id)
        if wait_for_network_configuration_ready is not None:
            pulumi.set(__self__, "wait_for_network_configuration_ready", wait_for_network_configuration_ready)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        The network interface id.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="trunkNetworkInstanceId")
    def trunk_network_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The trunk network instance id.
        """
        return pulumi.get(self, "trunk_network_instance_id")

    @trunk_network_instance_id.setter
    def trunk_network_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trunk_network_instance_id", value)

    @property
    @pulumi.getter(name="waitForNetworkConfigurationReady")
    def wait_for_network_configuration_ready(self) -> Optional[pulumi.Input[bool]]:
        """
        The wait for network configuration ready.
        """
        return pulumi.get(self, "wait_for_network_configuration_ready")

    @wait_for_network_configuration_ready.setter
    def wait_for_network_configuration_ready(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_for_network_configuration_ready", value)


class EcsNetworkInterfaceAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 trunk_network_instance_id: Optional[pulumi.Input[str]] = None,
                 wait_for_network_configuration_ready: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a ECS Network Interface Attachment resource.

        For information about ECS Network Interface Attachment and how to use it, see [What is Network Interface Attachment](https://www.alibabacloud.com/help/en/doc-detail/58515.htm).

        > **NOTE:** Available since v1.123.1+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_zones = alicloud.get_zones(available_resource_creation="Instance")
        default_instance_types = alicloud.ecs.get_instance_types(availability_zone=default_zones.zones[0].id,
            eni_amount=3)
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="192.168.0.0/24")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name=name,
            cidr_block="192.168.0.0/24",
            zone_id=default_zones.zones[0].id,
            vpc_id=default_network.id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup",
            description="New security group",
            vpc_id=default_network.id)
        default_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            most_recent=True,
            owners="system")
        default_instance = alicloud.ecs.Instance("defaultInstance",
            availability_zone=default_zones.zones[0].id,
            instance_name=name,
            host_name="tf-example",
            image_id=default_images.images[0].id,
            instance_type=default_instance_types.instance_types[0].id,
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id)
        default_resource_groups = alicloud.resourcemanager.get_resource_groups(status="OK")
        default_ecs_network_interface = alicloud.ecs.EcsNetworkInterface("defaultEcsNetworkInterface",
            network_interface_name=name,
            vswitch_id=default_switch.id,
            security_group_ids=[default_security_group.id],
            description="Basic example",
            primary_ip_address="192.168.0.2",
            tags={
                "Created": "TF",
                "For": "example",
            },
            resource_group_id=default_resource_groups.ids[0])
        default_ecs_network_interface_attachment = alicloud.ecs.EcsNetworkInterfaceAttachment("defaultEcsNetworkInterfaceAttachment",
            network_interface_id=default_ecs_network_interface.id,
            instance_id=default_instance.id)
        ```

        ## Import

        ECS Network Interface Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsNetworkInterfaceAttachment:EcsNetworkInterfaceAttachment example eni-abcd1234:i-abcd1234
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The instance id.
        :param pulumi.Input[str] network_interface_id: The network interface id.
        :param pulumi.Input[str] trunk_network_instance_id: The trunk network instance id.
        :param pulumi.Input[bool] wait_for_network_configuration_ready: The wait for network configuration ready.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EcsNetworkInterfaceAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ECS Network Interface Attachment resource.

        For information about ECS Network Interface Attachment and how to use it, see [What is Network Interface Attachment](https://www.alibabacloud.com/help/en/doc-detail/58515.htm).

        > **NOTE:** Available since v1.123.1+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_zones = alicloud.get_zones(available_resource_creation="Instance")
        default_instance_types = alicloud.ecs.get_instance_types(availability_zone=default_zones.zones[0].id,
            eni_amount=3)
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="192.168.0.0/24")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name=name,
            cidr_block="192.168.0.0/24",
            zone_id=default_zones.zones[0].id,
            vpc_id=default_network.id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup",
            description="New security group",
            vpc_id=default_network.id)
        default_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            most_recent=True,
            owners="system")
        default_instance = alicloud.ecs.Instance("defaultInstance",
            availability_zone=default_zones.zones[0].id,
            instance_name=name,
            host_name="tf-example",
            image_id=default_images.images[0].id,
            instance_type=default_instance_types.instance_types[0].id,
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id)
        default_resource_groups = alicloud.resourcemanager.get_resource_groups(status="OK")
        default_ecs_network_interface = alicloud.ecs.EcsNetworkInterface("defaultEcsNetworkInterface",
            network_interface_name=name,
            vswitch_id=default_switch.id,
            security_group_ids=[default_security_group.id],
            description="Basic example",
            primary_ip_address="192.168.0.2",
            tags={
                "Created": "TF",
                "For": "example",
            },
            resource_group_id=default_resource_groups.ids[0])
        default_ecs_network_interface_attachment = alicloud.ecs.EcsNetworkInterfaceAttachment("defaultEcsNetworkInterfaceAttachment",
            network_interface_id=default_ecs_network_interface.id,
            instance_id=default_instance.id)
        ```

        ## Import

        ECS Network Interface Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsNetworkInterfaceAttachment:EcsNetworkInterfaceAttachment example eni-abcd1234:i-abcd1234
        ```

        :param str resource_name: The name of the resource.
        :param EcsNetworkInterfaceAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EcsNetworkInterfaceAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 trunk_network_instance_id: Optional[pulumi.Input[str]] = None,
                 wait_for_network_configuration_ready: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EcsNetworkInterfaceAttachmentArgs.__new__(EcsNetworkInterfaceAttachmentArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if network_interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_interface_id'")
            __props__.__dict__["network_interface_id"] = network_interface_id
            __props__.__dict__["trunk_network_instance_id"] = trunk_network_instance_id
            __props__.__dict__["wait_for_network_configuration_ready"] = wait_for_network_configuration_ready
        super(EcsNetworkInterfaceAttachment, __self__).__init__(
            'alicloud:ecs/ecsNetworkInterfaceAttachment:EcsNetworkInterfaceAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None,
            trunk_network_instance_id: Optional[pulumi.Input[str]] = None,
            wait_for_network_configuration_ready: Optional[pulumi.Input[bool]] = None) -> 'EcsNetworkInterfaceAttachment':
        """
        Get an existing EcsNetworkInterfaceAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The instance id.
        :param pulumi.Input[str] network_interface_id: The network interface id.
        :param pulumi.Input[str] trunk_network_instance_id: The trunk network instance id.
        :param pulumi.Input[bool] wait_for_network_configuration_ready: The wait for network configuration ready.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EcsNetworkInterfaceAttachmentState.__new__(_EcsNetworkInterfaceAttachmentState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["network_interface_id"] = network_interface_id
        __props__.__dict__["trunk_network_instance_id"] = trunk_network_instance_id
        __props__.__dict__["wait_for_network_configuration_ready"] = wait_for_network_configuration_ready
        return EcsNetworkInterfaceAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        The network interface id.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="trunkNetworkInstanceId")
    def trunk_network_instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        The trunk network instance id.
        """
        return pulumi.get(self, "trunk_network_instance_id")

    @property
    @pulumi.getter(name="waitForNetworkConfigurationReady")
    def wait_for_network_configuration_ready(self) -> pulumi.Output[Optional[bool]]:
        """
        The wait for network configuration ready.
        """
        return pulumi.get(self, "wait_for_network_configuration_ready")

