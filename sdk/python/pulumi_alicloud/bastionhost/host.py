# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HostArgs', 'Host']

@pulumi.input_type
class HostArgs:
    def __init__(__self__, *,
                 active_address_type: pulumi.Input[str],
                 host_name: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 os_type: pulumi.Input[str],
                 source: pulumi.Input[str],
                 comment: Optional[pulumi.Input[str]] = None,
                 host_private_address: Optional[pulumi.Input[str]] = None,
                 host_public_address: Optional[pulumi.Input[str]] = None,
                 instance_region_id: Optional[pulumi.Input[str]] = None,
                 source_instance_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Host resource.
        :param pulumi.Input[str] active_address_type: Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
        :param pulumi.Input[str] host_name: Specify the new create a host name of the supports up to 128 characters.
        :param pulumi.Input[str] instance_id: Specify the new create a host where the Bastion host ID of.
        :param pulumi.Input[str] os_type: Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
        :param pulumi.Input[str] source: Specify the new create a host of source. Valid values:
        :param pulumi.Input[str] comment: Specify a host of notes, supports up to 500 characters.
        :param pulumi.Input[str] host_private_address: Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `active_address_type` parameter is set to `Private`.
        :param pulumi.Input[str] host_public_address: Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
        :param pulumi.Input[str] instance_region_id: The instance region id.
        :param pulumi.Input[str] source_instance_id: Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
        """
        pulumi.set(__self__, "active_address_type", active_address_type)
        pulumi.set(__self__, "host_name", host_name)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "os_type", os_type)
        pulumi.set(__self__, "source", source)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if host_private_address is not None:
            pulumi.set(__self__, "host_private_address", host_private_address)
        if host_public_address is not None:
            pulumi.set(__self__, "host_public_address", host_public_address)
        if instance_region_id is not None:
            pulumi.set(__self__, "instance_region_id", instance_region_id)
        if source_instance_id is not None:
            pulumi.set(__self__, "source_instance_id", source_instance_id)

    @property
    @pulumi.getter(name="activeAddressType")
    def active_address_type(self) -> pulumi.Input[str]:
        """
        Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
        """
        return pulumi.get(self, "active_address_type")

    @active_address_type.setter
    def active_address_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "active_address_type", value)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> pulumi.Input[str]:
        """
        Specify the new create a host name of the supports up to 128 characters.
        """
        return pulumi.get(self, "host_name")

    @host_name.setter
    def host_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "host_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specify the new create a host where the Bastion host ID of.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> pulumi.Input[str]:
        """
        Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
        """
        return pulumi.get(self, "os_type")

    @os_type.setter
    def os_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "os_type", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        """
        Specify the new create a host of source. Valid values:
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Specify a host of notes, supports up to 500 characters.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="hostPrivateAddress")
    def host_private_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `active_address_type` parameter is set to `Private`.
        """
        return pulumi.get(self, "host_private_address")

    @host_private_address.setter
    def host_private_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_private_address", value)

    @property
    @pulumi.getter(name="hostPublicAddress")
    def host_public_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
        """
        return pulumi.get(self, "host_public_address")

    @host_public_address.setter
    def host_public_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_public_address", value)

    @property
    @pulumi.getter(name="instanceRegionId")
    def instance_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance region id.
        """
        return pulumi.get(self, "instance_region_id")

    @instance_region_id.setter
    def instance_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_region_id", value)

    @property
    @pulumi.getter(name="sourceInstanceId")
    def source_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
        """
        return pulumi.get(self, "source_instance_id")

    @source_instance_id.setter
    def source_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_instance_id", value)


@pulumi.input_type
class _HostState:
    def __init__(__self__, *,
                 active_address_type: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 host_id: Optional[pulumi.Input[str]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 host_private_address: Optional[pulumi.Input[str]] = None,
                 host_public_address: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_region_id: Optional[pulumi.Input[str]] = None,
                 os_type: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Host resources.
        :param pulumi.Input[str] active_address_type: Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
        :param pulumi.Input[str] comment: Specify a host of notes, supports up to 500 characters.
        :param pulumi.Input[str] host_id: The host ID.
        :param pulumi.Input[str] host_name: Specify the new create a host name of the supports up to 128 characters.
        :param pulumi.Input[str] host_private_address: Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `active_address_type` parameter is set to `Private`.
        :param pulumi.Input[str] host_public_address: Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
        :param pulumi.Input[str] instance_id: Specify the new create a host where the Bastion host ID of.
        :param pulumi.Input[str] instance_region_id: The instance region id.
        :param pulumi.Input[str] os_type: Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
        :param pulumi.Input[str] source: Specify the new create a host of source. Valid values:
        :param pulumi.Input[str] source_instance_id: Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
        """
        if active_address_type is not None:
            pulumi.set(__self__, "active_address_type", active_address_type)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if host_id is not None:
            pulumi.set(__self__, "host_id", host_id)
        if host_name is not None:
            pulumi.set(__self__, "host_name", host_name)
        if host_private_address is not None:
            pulumi.set(__self__, "host_private_address", host_private_address)
        if host_public_address is not None:
            pulumi.set(__self__, "host_public_address", host_public_address)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_region_id is not None:
            pulumi.set(__self__, "instance_region_id", instance_region_id)
        if os_type is not None:
            pulumi.set(__self__, "os_type", os_type)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if source_instance_id is not None:
            pulumi.set(__self__, "source_instance_id", source_instance_id)

    @property
    @pulumi.getter(name="activeAddressType")
    def active_address_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
        """
        return pulumi.get(self, "active_address_type")

    @active_address_type.setter
    def active_address_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "active_address_type", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Specify a host of notes, supports up to 500 characters.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> Optional[pulumi.Input[str]]:
        """
        The host ID.
        """
        return pulumi.get(self, "host_id")

    @host_id.setter
    def host_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_id", value)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the new create a host name of the supports up to 128 characters.
        """
        return pulumi.get(self, "host_name")

    @host_name.setter
    def host_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_name", value)

    @property
    @pulumi.getter(name="hostPrivateAddress")
    def host_private_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `active_address_type` parameter is set to `Private`.
        """
        return pulumi.get(self, "host_private_address")

    @host_private_address.setter
    def host_private_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_private_address", value)

    @property
    @pulumi.getter(name="hostPublicAddress")
    def host_public_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
        """
        return pulumi.get(self, "host_public_address")

    @host_public_address.setter
    def host_public_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_public_address", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the new create a host where the Bastion host ID of.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceRegionId")
    def instance_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance region id.
        """
        return pulumi.get(self, "instance_region_id")

    @instance_region_id.setter
    def instance_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_region_id", value)

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
        """
        return pulumi.get(self, "os_type")

    @os_type.setter
    def os_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "os_type", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the new create a host of source. Valid values:
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="sourceInstanceId")
    def source_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
        """
        return pulumi.get(self, "source_instance_id")

    @source_instance_id.setter
    def source_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_instance_id", value)


class Host(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_address_type: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 host_private_address: Optional[pulumi.Input[str]] = None,
                 host_public_address: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_region_id: Optional[pulumi.Input[str]] = None,
                 os_type: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Bastion Host Host resource.

        For information about Bastion Host Host and how to use it, see [What is Host](https://www.alibabacloud.com/help/en/doc-detail/201330.htm).

        > **NOTE:** Available since v1.135.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf_example"
        default = alicloud.get_zones(available_resource_creation="VSwitch")
        default_get_networks = alicloud.vpc.get_networks(name_regex="^default-NODELETING$",
            cidr_block="10.4.0.0/16")
        default_get_switches = alicloud.vpc.get_switches(cidr_block="10.4.0.0/24",
            vpc_id=default_get_networks.ids[0],
            zone_id=default.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("default", vpc_id=default_get_networks.ids[0])
        default_instance = alicloud.bastionhost.Instance("default",
            description=name,
            license_code="bhah_ent_50_asset",
            plan_code="cloudbastion",
            storage="5",
            bandwidth="5",
            period=1,
            vswitch_id=default_get_switches.ids[0],
            security_group_ids=[default_security_group.id])
        default_host = alicloud.bastionhost.Host("default",
            instance_id=default_instance.id,
            host_name=name,
            active_address_type="Private",
            host_private_address="172.16.0.10",
            os_type="Linux",
            source="Local")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Bastion Host Host can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:bastionhost/host:Host example <instance_id>:<host_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] active_address_type: Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
        :param pulumi.Input[str] comment: Specify a host of notes, supports up to 500 characters.
        :param pulumi.Input[str] host_name: Specify the new create a host name of the supports up to 128 characters.
        :param pulumi.Input[str] host_private_address: Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `active_address_type` parameter is set to `Private`.
        :param pulumi.Input[str] host_public_address: Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
        :param pulumi.Input[str] instance_id: Specify the new create a host where the Bastion host ID of.
        :param pulumi.Input[str] instance_region_id: The instance region id.
        :param pulumi.Input[str] os_type: Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
        :param pulumi.Input[str] source: Specify the new create a host of source. Valid values:
        :param pulumi.Input[str] source_instance_id: Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HostArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Bastion Host Host resource.

        For information about Bastion Host Host and how to use it, see [What is Host](https://www.alibabacloud.com/help/en/doc-detail/201330.htm).

        > **NOTE:** Available since v1.135.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf_example"
        default = alicloud.get_zones(available_resource_creation="VSwitch")
        default_get_networks = alicloud.vpc.get_networks(name_regex="^default-NODELETING$",
            cidr_block="10.4.0.0/16")
        default_get_switches = alicloud.vpc.get_switches(cidr_block="10.4.0.0/24",
            vpc_id=default_get_networks.ids[0],
            zone_id=default.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("default", vpc_id=default_get_networks.ids[0])
        default_instance = alicloud.bastionhost.Instance("default",
            description=name,
            license_code="bhah_ent_50_asset",
            plan_code="cloudbastion",
            storage="5",
            bandwidth="5",
            period=1,
            vswitch_id=default_get_switches.ids[0],
            security_group_ids=[default_security_group.id])
        default_host = alicloud.bastionhost.Host("default",
            instance_id=default_instance.id,
            host_name=name,
            active_address_type="Private",
            host_private_address="172.16.0.10",
            os_type="Linux",
            source="Local")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Bastion Host Host can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:bastionhost/host:Host example <instance_id>:<host_id>
        ```

        :param str resource_name: The name of the resource.
        :param HostArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HostArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_address_type: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 host_private_address: Optional[pulumi.Input[str]] = None,
                 host_public_address: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_region_id: Optional[pulumi.Input[str]] = None,
                 os_type: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HostArgs.__new__(HostArgs)

            if active_address_type is None and not opts.urn:
                raise TypeError("Missing required property 'active_address_type'")
            __props__.__dict__["active_address_type"] = active_address_type
            __props__.__dict__["comment"] = comment
            if host_name is None and not opts.urn:
                raise TypeError("Missing required property 'host_name'")
            __props__.__dict__["host_name"] = host_name
            __props__.__dict__["host_private_address"] = host_private_address
            __props__.__dict__["host_public_address"] = host_public_address
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["instance_region_id"] = instance_region_id
            if os_type is None and not opts.urn:
                raise TypeError("Missing required property 'os_type'")
            __props__.__dict__["os_type"] = os_type
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            __props__.__dict__["source_instance_id"] = source_instance_id
            __props__.__dict__["host_id"] = None
        super(Host, __self__).__init__(
            'alicloud:bastionhost/host:Host',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active_address_type: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            host_id: Optional[pulumi.Input[str]] = None,
            host_name: Optional[pulumi.Input[str]] = None,
            host_private_address: Optional[pulumi.Input[str]] = None,
            host_public_address: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instance_region_id: Optional[pulumi.Input[str]] = None,
            os_type: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            source_instance_id: Optional[pulumi.Input[str]] = None) -> 'Host':
        """
        Get an existing Host resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] active_address_type: Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
        :param pulumi.Input[str] comment: Specify a host of notes, supports up to 500 characters.
        :param pulumi.Input[str] host_id: The host ID.
        :param pulumi.Input[str] host_name: Specify the new create a host name of the supports up to 128 characters.
        :param pulumi.Input[str] host_private_address: Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `active_address_type` parameter is set to `Private`.
        :param pulumi.Input[str] host_public_address: Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
        :param pulumi.Input[str] instance_id: Specify the new create a host where the Bastion host ID of.
        :param pulumi.Input[str] instance_region_id: The instance region id.
        :param pulumi.Input[str] os_type: Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
        :param pulumi.Input[str] source: Specify the new create a host of source. Valid values:
        :param pulumi.Input[str] source_instance_id: Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HostState.__new__(_HostState)

        __props__.__dict__["active_address_type"] = active_address_type
        __props__.__dict__["comment"] = comment
        __props__.__dict__["host_id"] = host_id
        __props__.__dict__["host_name"] = host_name
        __props__.__dict__["host_private_address"] = host_private_address
        __props__.__dict__["host_public_address"] = host_public_address
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["instance_region_id"] = instance_region_id
        __props__.__dict__["os_type"] = os_type
        __props__.__dict__["source"] = source
        __props__.__dict__["source_instance_id"] = source_instance_id
        return Host(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeAddressType")
    def active_address_type(self) -> pulumi.Output[str]:
        """
        Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
        """
        return pulumi.get(self, "active_address_type")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Specify a host of notes, supports up to 500 characters.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> pulumi.Output[str]:
        """
        The host ID.
        """
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> pulumi.Output[str]:
        """
        Specify the new create a host name of the supports up to 128 characters.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter(name="hostPrivateAddress")
    def host_private_address(self) -> pulumi.Output[Optional[str]]:
        """
        Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `active_address_type` parameter is set to `Private`.
        """
        return pulumi.get(self, "host_private_address")

    @property
    @pulumi.getter(name="hostPublicAddress")
    def host_public_address(self) -> pulumi.Output[Optional[str]]:
        """
        Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
        """
        return pulumi.get(self, "host_public_address")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specify the new create a host where the Bastion host ID of.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceRegionId")
    def instance_region_id(self) -> pulumi.Output[Optional[str]]:
        """
        The instance region id.
        """
        return pulumi.get(self, "instance_region_id")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> pulumi.Output[str]:
        """
        Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        Specify the new create a host of source. Valid values:
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourceInstanceId")
    def source_instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
        """
        return pulumi.get(self, "source_instance_id")

