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

__all__ = ['IpsecServerArgs', 'IpsecServer']

@pulumi.input_type
class IpsecServerArgs:
    def __init__(__self__, *,
                 client_ip_pool: pulumi.Input[str],
                 local_subnet: pulumi.Input[str],
                 vpn_gateway_id: pulumi.Input[str],
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 effect_immediately: Optional[pulumi.Input[bool]] = None,
                 ike_configs: Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIkeConfigArgs']]]] = None,
                 ipsec_configs: Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIpsecConfigArgs']]]] = None,
                 ipsec_server_name: Optional[pulumi.Input[str]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 psk_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a IpsecServer resource.
        :param pulumi.Input[str] client_ip_pool: The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
        :param pulumi.Input[str] local_subnet: The local CIDR block. It refers to the CIDR block of the virtual private cloud (VPC) that is used to connect with the client. Separate multiple CIDR blocks with commas (,). Example: `192.168.1.0/24,192.168.2.0/24`.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN gateway.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[bool] effect_immediately: Specifies whether you want the configuration to immediately take effect.
        :param pulumi.Input[Sequence[pulumi.Input['IpsecServerIkeConfigArgs']]] ike_configs: The configuration of Phase 1 negotiations. See `ike_config` below.
        :param pulumi.Input[Sequence[pulumi.Input['IpsecServerIpsecConfigArgs']]] ipsec_configs: The configuration of Phase 2 negotiations. See `ipsec_config` below.
        :param pulumi.Input[str] ipsec_server_name: The name of the IPsec server. The name must be `2` to `128` characters in length, and can contain digits, hyphens (-), and underscores (_). It must start with a letter.
        :param pulumi.Input[str] psk: The pre-shared key. The pre-shared key is used to authenticate the VPN gateway and the client. By default, the system generates a random string that is 16 bits in length. You can also specify the pre-shared key. It can contain at most 100 characters.
        :param pulumi.Input[bool] psk_enabled: Whether to enable the pre-shared key authentication method. The value is only `true`, which indicates that the pre-shared key authentication method is enabled.
        """
        pulumi.set(__self__, "client_ip_pool", client_ip_pool)
        pulumi.set(__self__, "local_subnet", local_subnet)
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if effect_immediately is not None:
            pulumi.set(__self__, "effect_immediately", effect_immediately)
        if ike_configs is not None:
            pulumi.set(__self__, "ike_configs", ike_configs)
        if ipsec_configs is not None:
            pulumi.set(__self__, "ipsec_configs", ipsec_configs)
        if ipsec_server_name is not None:
            pulumi.set(__self__, "ipsec_server_name", ipsec_server_name)
        if psk is not None:
            pulumi.set(__self__, "psk", psk)
        if psk_enabled is not None:
            pulumi.set(__self__, "psk_enabled", psk_enabled)

    @property
    @pulumi.getter(name="clientIpPool")
    def client_ip_pool(self) -> pulumi.Input[str]:
        """
        The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
        """
        return pulumi.get(self, "client_ip_pool")

    @client_ip_pool.setter
    def client_ip_pool(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_ip_pool", value)

    @property
    @pulumi.getter(name="localSubnet")
    def local_subnet(self) -> pulumi.Input[str]:
        """
        The local CIDR block. It refers to the CIDR block of the virtual private cloud (VPC) that is used to connect with the client. Separate multiple CIDR blocks with commas (,). Example: `192.168.1.0/24,192.168.2.0/24`.
        """
        return pulumi.get(self, "local_subnet")

    @local_subnet.setter
    def local_subnet(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_subnet", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPN gateway.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpn_gateway_id", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        The dry run.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="effectImmediately")
    def effect_immediately(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether you want the configuration to immediately take effect.
        """
        return pulumi.get(self, "effect_immediately")

    @effect_immediately.setter
    def effect_immediately(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "effect_immediately", value)

    @property
    @pulumi.getter(name="ikeConfigs")
    def ike_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIkeConfigArgs']]]]:
        """
        The configuration of Phase 1 negotiations. See `ike_config` below.
        """
        return pulumi.get(self, "ike_configs")

    @ike_configs.setter
    def ike_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIkeConfigArgs']]]]):
        pulumi.set(self, "ike_configs", value)

    @property
    @pulumi.getter(name="ipsecConfigs")
    def ipsec_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIpsecConfigArgs']]]]:
        """
        The configuration of Phase 2 negotiations. See `ipsec_config` below.
        """
        return pulumi.get(self, "ipsec_configs")

    @ipsec_configs.setter
    def ipsec_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIpsecConfigArgs']]]]):
        pulumi.set(self, "ipsec_configs", value)

    @property
    @pulumi.getter(name="ipsecServerName")
    def ipsec_server_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IPsec server. The name must be `2` to `128` characters in length, and can contain digits, hyphens (-), and underscores (_). It must start with a letter.
        """
        return pulumi.get(self, "ipsec_server_name")

    @ipsec_server_name.setter
    def ipsec_server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipsec_server_name", value)

    @property
    @pulumi.getter
    def psk(self) -> Optional[pulumi.Input[str]]:
        """
        The pre-shared key. The pre-shared key is used to authenticate the VPN gateway and the client. By default, the system generates a random string that is 16 bits in length. You can also specify the pre-shared key. It can contain at most 100 characters.
        """
        return pulumi.get(self, "psk")

    @psk.setter
    def psk(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psk", value)

    @property
    @pulumi.getter(name="pskEnabled")
    def psk_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the pre-shared key authentication method. The value is only `true`, which indicates that the pre-shared key authentication method is enabled.
        """
        return pulumi.get(self, "psk_enabled")

    @psk_enabled.setter
    def psk_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "psk_enabled", value)


@pulumi.input_type
class _IpsecServerState:
    def __init__(__self__, *,
                 client_ip_pool: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 effect_immediately: Optional[pulumi.Input[bool]] = None,
                 ike_configs: Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIkeConfigArgs']]]] = None,
                 ipsec_configs: Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIpsecConfigArgs']]]] = None,
                 ipsec_server_name: Optional[pulumi.Input[str]] = None,
                 local_subnet: Optional[pulumi.Input[str]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 psk_enabled: Optional[pulumi.Input[bool]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpsecServer resources.
        :param pulumi.Input[str] client_ip_pool: The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[bool] effect_immediately: Specifies whether you want the configuration to immediately take effect.
        :param pulumi.Input[Sequence[pulumi.Input['IpsecServerIkeConfigArgs']]] ike_configs: The configuration of Phase 1 negotiations. See `ike_config` below.
        :param pulumi.Input[Sequence[pulumi.Input['IpsecServerIpsecConfigArgs']]] ipsec_configs: The configuration of Phase 2 negotiations. See `ipsec_config` below.
        :param pulumi.Input[str] ipsec_server_name: The name of the IPsec server. The name must be `2` to `128` characters in length, and can contain digits, hyphens (-), and underscores (_). It must start with a letter.
        :param pulumi.Input[str] local_subnet: The local CIDR block. It refers to the CIDR block of the virtual private cloud (VPC) that is used to connect with the client. Separate multiple CIDR blocks with commas (,). Example: `192.168.1.0/24,192.168.2.0/24`.
        :param pulumi.Input[str] psk: The pre-shared key. The pre-shared key is used to authenticate the VPN gateway and the client. By default, the system generates a random string that is 16 bits in length. You can also specify the pre-shared key. It can contain at most 100 characters.
        :param pulumi.Input[bool] psk_enabled: Whether to enable the pre-shared key authentication method. The value is only `true`, which indicates that the pre-shared key authentication method is enabled.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN gateway.
        """
        if client_ip_pool is not None:
            pulumi.set(__self__, "client_ip_pool", client_ip_pool)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if effect_immediately is not None:
            pulumi.set(__self__, "effect_immediately", effect_immediately)
        if ike_configs is not None:
            pulumi.set(__self__, "ike_configs", ike_configs)
        if ipsec_configs is not None:
            pulumi.set(__self__, "ipsec_configs", ipsec_configs)
        if ipsec_server_name is not None:
            pulumi.set(__self__, "ipsec_server_name", ipsec_server_name)
        if local_subnet is not None:
            pulumi.set(__self__, "local_subnet", local_subnet)
        if psk is not None:
            pulumi.set(__self__, "psk", psk)
        if psk_enabled is not None:
            pulumi.set(__self__, "psk_enabled", psk_enabled)
        if vpn_gateway_id is not None:
            pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter(name="clientIpPool")
    def client_ip_pool(self) -> Optional[pulumi.Input[str]]:
        """
        The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
        """
        return pulumi.get(self, "client_ip_pool")

    @client_ip_pool.setter
    def client_ip_pool(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_ip_pool", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        The dry run.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="effectImmediately")
    def effect_immediately(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether you want the configuration to immediately take effect.
        """
        return pulumi.get(self, "effect_immediately")

    @effect_immediately.setter
    def effect_immediately(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "effect_immediately", value)

    @property
    @pulumi.getter(name="ikeConfigs")
    def ike_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIkeConfigArgs']]]]:
        """
        The configuration of Phase 1 negotiations. See `ike_config` below.
        """
        return pulumi.get(self, "ike_configs")

    @ike_configs.setter
    def ike_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIkeConfigArgs']]]]):
        pulumi.set(self, "ike_configs", value)

    @property
    @pulumi.getter(name="ipsecConfigs")
    def ipsec_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIpsecConfigArgs']]]]:
        """
        The configuration of Phase 2 negotiations. See `ipsec_config` below.
        """
        return pulumi.get(self, "ipsec_configs")

    @ipsec_configs.setter
    def ipsec_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpsecServerIpsecConfigArgs']]]]):
        pulumi.set(self, "ipsec_configs", value)

    @property
    @pulumi.getter(name="ipsecServerName")
    def ipsec_server_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IPsec server. The name must be `2` to `128` characters in length, and can contain digits, hyphens (-), and underscores (_). It must start with a letter.
        """
        return pulumi.get(self, "ipsec_server_name")

    @ipsec_server_name.setter
    def ipsec_server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipsec_server_name", value)

    @property
    @pulumi.getter(name="localSubnet")
    def local_subnet(self) -> Optional[pulumi.Input[str]]:
        """
        The local CIDR block. It refers to the CIDR block of the virtual private cloud (VPC) that is used to connect with the client. Separate multiple CIDR blocks with commas (,). Example: `192.168.1.0/24,192.168.2.0/24`.
        """
        return pulumi.get(self, "local_subnet")

    @local_subnet.setter
    def local_subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_subnet", value)

    @property
    @pulumi.getter
    def psk(self) -> Optional[pulumi.Input[str]]:
        """
        The pre-shared key. The pre-shared key is used to authenticate the VPN gateway and the client. By default, the system generates a random string that is 16 bits in length. You can also specify the pre-shared key. It can contain at most 100 characters.
        """
        return pulumi.get(self, "psk")

    @psk.setter
    def psk(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psk", value)

    @property
    @pulumi.getter(name="pskEnabled")
    def psk_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the pre-shared key authentication method. The value is only `true`, which indicates that the pre-shared key authentication method is enabled.
        """
        return pulumi.get(self, "psk_enabled")

    @psk_enabled.setter
    def psk_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "psk_enabled", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPN gateway.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_gateway_id", value)


class IpsecServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_ip_pool: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 effect_immediately: Optional[pulumi.Input[bool]] = None,
                 ike_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpsecServerIkeConfigArgs']]]]] = None,
                 ipsec_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpsecServerIpsecConfigArgs']]]]] = None,
                 ipsec_server_name: Optional[pulumi.Input[str]] = None,
                 local_subnet: Optional[pulumi.Input[str]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 psk_enabled: Optional[pulumi.Input[bool]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a VPN Ipsec Server resource.

        For information about VPN Ipsec Server and how to use it, see [What is Ipsec Server](https://www.alibabacloud.com/help/en/doc-detail/205454.html).

        > **NOTE:** Available since v1.161.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.get_zones(available_resource_creation="VSwitch")
        default_get_networks = alicloud.vpc.get_networks(name_regex="^default-NODELETING$")
        default_get_switches = alicloud.vpc.get_switches(vpc_id=default_get_networks.ids[0],
            zone_id=default.ids[0])
        default_gateway = alicloud.vpn.Gateway("default",
            name=name,
            vpc_id=default_get_networks.ids[0],
            bandwidth=10,
            enable_ssl=True,
            description=name,
            instance_charge_type="PrePaid",
            vswitch_id=default_get_switches.ids[0])
        foo = alicloud.vpn.IpsecServer("foo",
            client_ip_pool="10.0.0.0/24",
            ipsec_server_name=name,
            local_subnet="192.168.0.0/24",
            vpn_gateway_id=default_gateway.id,
            psk_enabled=True)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        VPN Ipsec Server can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:vpn/ipsecServer:IpsecServer example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_ip_pool: The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[bool] effect_immediately: Specifies whether you want the configuration to immediately take effect.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpsecServerIkeConfigArgs']]]] ike_configs: The configuration of Phase 1 negotiations. See `ike_config` below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpsecServerIpsecConfigArgs']]]] ipsec_configs: The configuration of Phase 2 negotiations. See `ipsec_config` below.
        :param pulumi.Input[str] ipsec_server_name: The name of the IPsec server. The name must be `2` to `128` characters in length, and can contain digits, hyphens (-), and underscores (_). It must start with a letter.
        :param pulumi.Input[str] local_subnet: The local CIDR block. It refers to the CIDR block of the virtual private cloud (VPC) that is used to connect with the client. Separate multiple CIDR blocks with commas (,). Example: `192.168.1.0/24,192.168.2.0/24`.
        :param pulumi.Input[str] psk: The pre-shared key. The pre-shared key is used to authenticate the VPN gateway and the client. By default, the system generates a random string that is 16 bits in length. You can also specify the pre-shared key. It can contain at most 100 characters.
        :param pulumi.Input[bool] psk_enabled: Whether to enable the pre-shared key authentication method. The value is only `true`, which indicates that the pre-shared key authentication method is enabled.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN gateway.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpsecServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPN Ipsec Server resource.

        For information about VPN Ipsec Server and how to use it, see [What is Ipsec Server](https://www.alibabacloud.com/help/en/doc-detail/205454.html).

        > **NOTE:** Available since v1.161.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.get_zones(available_resource_creation="VSwitch")
        default_get_networks = alicloud.vpc.get_networks(name_regex="^default-NODELETING$")
        default_get_switches = alicloud.vpc.get_switches(vpc_id=default_get_networks.ids[0],
            zone_id=default.ids[0])
        default_gateway = alicloud.vpn.Gateway("default",
            name=name,
            vpc_id=default_get_networks.ids[0],
            bandwidth=10,
            enable_ssl=True,
            description=name,
            instance_charge_type="PrePaid",
            vswitch_id=default_get_switches.ids[0])
        foo = alicloud.vpn.IpsecServer("foo",
            client_ip_pool="10.0.0.0/24",
            ipsec_server_name=name,
            local_subnet="192.168.0.0/24",
            vpn_gateway_id=default_gateway.id,
            psk_enabled=True)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        VPN Ipsec Server can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:vpn/ipsecServer:IpsecServer example <id>
        ```

        :param str resource_name: The name of the resource.
        :param IpsecServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpsecServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_ip_pool: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 effect_immediately: Optional[pulumi.Input[bool]] = None,
                 ike_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpsecServerIkeConfigArgs']]]]] = None,
                 ipsec_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpsecServerIpsecConfigArgs']]]]] = None,
                 ipsec_server_name: Optional[pulumi.Input[str]] = None,
                 local_subnet: Optional[pulumi.Input[str]] = None,
                 psk: Optional[pulumi.Input[str]] = None,
                 psk_enabled: Optional[pulumi.Input[bool]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpsecServerArgs.__new__(IpsecServerArgs)

            if client_ip_pool is None and not opts.urn:
                raise TypeError("Missing required property 'client_ip_pool'")
            __props__.__dict__["client_ip_pool"] = client_ip_pool
            __props__.__dict__["dry_run"] = dry_run
            __props__.__dict__["effect_immediately"] = effect_immediately
            __props__.__dict__["ike_configs"] = ike_configs
            __props__.__dict__["ipsec_configs"] = ipsec_configs
            __props__.__dict__["ipsec_server_name"] = ipsec_server_name
            if local_subnet is None and not opts.urn:
                raise TypeError("Missing required property 'local_subnet'")
            __props__.__dict__["local_subnet"] = local_subnet
            __props__.__dict__["psk"] = psk
            __props__.__dict__["psk_enabled"] = psk_enabled
            if vpn_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpn_gateway_id'")
            __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
        super(IpsecServer, __self__).__init__(
            'alicloud:vpn/ipsecServer:IpsecServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_ip_pool: Optional[pulumi.Input[str]] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            effect_immediately: Optional[pulumi.Input[bool]] = None,
            ike_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpsecServerIkeConfigArgs']]]]] = None,
            ipsec_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpsecServerIpsecConfigArgs']]]]] = None,
            ipsec_server_name: Optional[pulumi.Input[str]] = None,
            local_subnet: Optional[pulumi.Input[str]] = None,
            psk: Optional[pulumi.Input[str]] = None,
            psk_enabled: Optional[pulumi.Input[bool]] = None,
            vpn_gateway_id: Optional[pulumi.Input[str]] = None) -> 'IpsecServer':
        """
        Get an existing IpsecServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_ip_pool: The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[bool] effect_immediately: Specifies whether you want the configuration to immediately take effect.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpsecServerIkeConfigArgs']]]] ike_configs: The configuration of Phase 1 negotiations. See `ike_config` below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpsecServerIpsecConfigArgs']]]] ipsec_configs: The configuration of Phase 2 negotiations. See `ipsec_config` below.
        :param pulumi.Input[str] ipsec_server_name: The name of the IPsec server. The name must be `2` to `128` characters in length, and can contain digits, hyphens (-), and underscores (_). It must start with a letter.
        :param pulumi.Input[str] local_subnet: The local CIDR block. It refers to the CIDR block of the virtual private cloud (VPC) that is used to connect with the client. Separate multiple CIDR blocks with commas (,). Example: `192.168.1.0/24,192.168.2.0/24`.
        :param pulumi.Input[str] psk: The pre-shared key. The pre-shared key is used to authenticate the VPN gateway and the client. By default, the system generates a random string that is 16 bits in length. You can also specify the pre-shared key. It can contain at most 100 characters.
        :param pulumi.Input[bool] psk_enabled: Whether to enable the pre-shared key authentication method. The value is only `true`, which indicates that the pre-shared key authentication method is enabled.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN gateway.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpsecServerState.__new__(_IpsecServerState)

        __props__.__dict__["client_ip_pool"] = client_ip_pool
        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["effect_immediately"] = effect_immediately
        __props__.__dict__["ike_configs"] = ike_configs
        __props__.__dict__["ipsec_configs"] = ipsec_configs
        __props__.__dict__["ipsec_server_name"] = ipsec_server_name
        __props__.__dict__["local_subnet"] = local_subnet
        __props__.__dict__["psk"] = psk
        __props__.__dict__["psk_enabled"] = psk_enabled
        __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
        return IpsecServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientIpPool")
    def client_ip_pool(self) -> pulumi.Output[str]:
        """
        The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
        """
        return pulumi.get(self, "client_ip_pool")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        The dry run.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter(name="effectImmediately")
    def effect_immediately(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether you want the configuration to immediately take effect.
        """
        return pulumi.get(self, "effect_immediately")

    @property
    @pulumi.getter(name="ikeConfigs")
    def ike_configs(self) -> pulumi.Output[Sequence['outputs.IpsecServerIkeConfig']]:
        """
        The configuration of Phase 1 negotiations. See `ike_config` below.
        """
        return pulumi.get(self, "ike_configs")

    @property
    @pulumi.getter(name="ipsecConfigs")
    def ipsec_configs(self) -> pulumi.Output[Sequence['outputs.IpsecServerIpsecConfig']]:
        """
        The configuration of Phase 2 negotiations. See `ipsec_config` below.
        """
        return pulumi.get(self, "ipsec_configs")

    @property
    @pulumi.getter(name="ipsecServerName")
    def ipsec_server_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the IPsec server. The name must be `2` to `128` characters in length, and can contain digits, hyphens (-), and underscores (_). It must start with a letter.
        """
        return pulumi.get(self, "ipsec_server_name")

    @property
    @pulumi.getter(name="localSubnet")
    def local_subnet(self) -> pulumi.Output[str]:
        """
        The local CIDR block. It refers to the CIDR block of the virtual private cloud (VPC) that is used to connect with the client. Separate multiple CIDR blocks with commas (,). Example: `192.168.1.0/24,192.168.2.0/24`.
        """
        return pulumi.get(self, "local_subnet")

    @property
    @pulumi.getter
    def psk(self) -> pulumi.Output[str]:
        """
        The pre-shared key. The pre-shared key is used to authenticate the VPN gateway and the client. By default, the system generates a random string that is 16 bits in length. You can also specify the pre-shared key. It can contain at most 100 characters.
        """
        return pulumi.get(self, "psk")

    @property
    @pulumi.getter(name="pskEnabled")
    def psk_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable the pre-shared key authentication method. The value is only `true`, which indicates that the pre-shared key authentication method is enabled.
        """
        return pulumi.get(self, "psk_enabled")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPN gateway.
        """
        return pulumi.get(self, "vpn_gateway_id")

