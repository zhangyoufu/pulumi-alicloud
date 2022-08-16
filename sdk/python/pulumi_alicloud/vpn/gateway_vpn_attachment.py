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

__all__ = ['GatewayVpnAttachmentArgs', 'GatewayVpnAttachment']

@pulumi.input_type
class GatewayVpnAttachmentArgs:
    def __init__(__self__, *,
                 customer_gateway_id: pulumi.Input[str],
                 local_subnet: pulumi.Input[str],
                 remote_subnet: pulumi.Input[str],
                 bgp_config: Optional[pulumi.Input['GatewayVpnAttachmentBgpConfigArgs']] = None,
                 effect_immediately: Optional[pulumi.Input[bool]] = None,
                 enable_dpd: Optional[pulumi.Input[bool]] = None,
                 enable_nat_traversal: Optional[pulumi.Input[bool]] = None,
                 health_check_config: Optional[pulumi.Input['GatewayVpnAttachmentHealthCheckConfigArgs']] = None,
                 ike_config: Optional[pulumi.Input['GatewayVpnAttachmentIkeConfigArgs']] = None,
                 ipsec_config: Optional[pulumi.Input['GatewayVpnAttachmentIpsecConfigArgs']] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 vpn_attachment_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GatewayVpnAttachment resource.
        :param pulumi.Input[str] customer_gateway_id: The ID of the customer gateway.
        :param pulumi.Input[str] local_subnet: The CIDR block of the virtual private cloud (VPC).
        :param pulumi.Input[str] remote_subnet: The CIDR block of the on-premises data center.
        :param pulumi.Input['GatewayVpnAttachmentBgpConfigArgs'] bgp_config: Bgp configuration information. See the following `Block bgp_config`.
        :param pulumi.Input[bool] effect_immediately: Indicates whether IPsec-VPN negotiations are initiated immediately. Valid values.
        :param pulumi.Input[bool] enable_dpd: Whether to enable the DPD (peer survival detection) function.
        :param pulumi.Input[bool] enable_nat_traversal: Allow NAT penetration.
        :param pulumi.Input['GatewayVpnAttachmentHealthCheckConfigArgs'] health_check_config: Health check configuration information. See the following `Block health_check_config`.
        :param pulumi.Input['GatewayVpnAttachmentIkeConfigArgs'] ike_config: Configuration negotiated in the second stage. See the following `Block ike_config`.
        :param pulumi.Input['GatewayVpnAttachmentIpsecConfigArgs'] ipsec_config: Configuration negotiated in the second stage. See the following `Block ipsec_config`.
        :param pulumi.Input[str] network_type: The network type of the IPsec connection. Valid values: `public`, `private`.
        :param pulumi.Input[str] vpn_attachment_name: The name of the vpn attachment.
        """
        pulumi.set(__self__, "customer_gateway_id", customer_gateway_id)
        pulumi.set(__self__, "local_subnet", local_subnet)
        pulumi.set(__self__, "remote_subnet", remote_subnet)
        if bgp_config is not None:
            pulumi.set(__self__, "bgp_config", bgp_config)
        if effect_immediately is not None:
            pulumi.set(__self__, "effect_immediately", effect_immediately)
        if enable_dpd is not None:
            pulumi.set(__self__, "enable_dpd", enable_dpd)
        if enable_nat_traversal is not None:
            pulumi.set(__self__, "enable_nat_traversal", enable_nat_traversal)
        if health_check_config is not None:
            pulumi.set(__self__, "health_check_config", health_check_config)
        if ike_config is not None:
            pulumi.set(__self__, "ike_config", ike_config)
        if ipsec_config is not None:
            pulumi.set(__self__, "ipsec_config", ipsec_config)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if vpn_attachment_name is not None:
            pulumi.set(__self__, "vpn_attachment_name", vpn_attachment_name)

    @property
    @pulumi.getter(name="customerGatewayId")
    def customer_gateway_id(self) -> pulumi.Input[str]:
        """
        The ID of the customer gateway.
        """
        return pulumi.get(self, "customer_gateway_id")

    @customer_gateway_id.setter
    def customer_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "customer_gateway_id", value)

    @property
    @pulumi.getter(name="localSubnet")
    def local_subnet(self) -> pulumi.Input[str]:
        """
        The CIDR block of the virtual private cloud (VPC).
        """
        return pulumi.get(self, "local_subnet")

    @local_subnet.setter
    def local_subnet(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_subnet", value)

    @property
    @pulumi.getter(name="remoteSubnet")
    def remote_subnet(self) -> pulumi.Input[str]:
        """
        The CIDR block of the on-premises data center.
        """
        return pulumi.get(self, "remote_subnet")

    @remote_subnet.setter
    def remote_subnet(self, value: pulumi.Input[str]):
        pulumi.set(self, "remote_subnet", value)

    @property
    @pulumi.getter(name="bgpConfig")
    def bgp_config(self) -> Optional[pulumi.Input['GatewayVpnAttachmentBgpConfigArgs']]:
        """
        Bgp configuration information. See the following `Block bgp_config`.
        """
        return pulumi.get(self, "bgp_config")

    @bgp_config.setter
    def bgp_config(self, value: Optional[pulumi.Input['GatewayVpnAttachmentBgpConfigArgs']]):
        pulumi.set(self, "bgp_config", value)

    @property
    @pulumi.getter(name="effectImmediately")
    def effect_immediately(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether IPsec-VPN negotiations are initiated immediately. Valid values.
        """
        return pulumi.get(self, "effect_immediately")

    @effect_immediately.setter
    def effect_immediately(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "effect_immediately", value)

    @property
    @pulumi.getter(name="enableDpd")
    def enable_dpd(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the DPD (peer survival detection) function.
        """
        return pulumi.get(self, "enable_dpd")

    @enable_dpd.setter
    def enable_dpd(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_dpd", value)

    @property
    @pulumi.getter(name="enableNatTraversal")
    def enable_nat_traversal(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow NAT penetration.
        """
        return pulumi.get(self, "enable_nat_traversal")

    @enable_nat_traversal.setter
    def enable_nat_traversal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_nat_traversal", value)

    @property
    @pulumi.getter(name="healthCheckConfig")
    def health_check_config(self) -> Optional[pulumi.Input['GatewayVpnAttachmentHealthCheckConfigArgs']]:
        """
        Health check configuration information. See the following `Block health_check_config`.
        """
        return pulumi.get(self, "health_check_config")

    @health_check_config.setter
    def health_check_config(self, value: Optional[pulumi.Input['GatewayVpnAttachmentHealthCheckConfigArgs']]):
        pulumi.set(self, "health_check_config", value)

    @property
    @pulumi.getter(name="ikeConfig")
    def ike_config(self) -> Optional[pulumi.Input['GatewayVpnAttachmentIkeConfigArgs']]:
        """
        Configuration negotiated in the second stage. See the following `Block ike_config`.
        """
        return pulumi.get(self, "ike_config")

    @ike_config.setter
    def ike_config(self, value: Optional[pulumi.Input['GatewayVpnAttachmentIkeConfigArgs']]):
        pulumi.set(self, "ike_config", value)

    @property
    @pulumi.getter(name="ipsecConfig")
    def ipsec_config(self) -> Optional[pulumi.Input['GatewayVpnAttachmentIpsecConfigArgs']]:
        """
        Configuration negotiated in the second stage. See the following `Block ipsec_config`.
        """
        return pulumi.get(self, "ipsec_config")

    @ipsec_config.setter
    def ipsec_config(self, value: Optional[pulumi.Input['GatewayVpnAttachmentIpsecConfigArgs']]):
        pulumi.set(self, "ipsec_config", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        The network type of the IPsec connection. Valid values: `public`, `private`.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter(name="vpnAttachmentName")
    def vpn_attachment_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the vpn attachment.
        """
        return pulumi.get(self, "vpn_attachment_name")

    @vpn_attachment_name.setter
    def vpn_attachment_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_attachment_name", value)


@pulumi.input_type
class _GatewayVpnAttachmentState:
    def __init__(__self__, *,
                 bgp_config: Optional[pulumi.Input['GatewayVpnAttachmentBgpConfigArgs']] = None,
                 customer_gateway_id: Optional[pulumi.Input[str]] = None,
                 effect_immediately: Optional[pulumi.Input[bool]] = None,
                 enable_dpd: Optional[pulumi.Input[bool]] = None,
                 enable_nat_traversal: Optional[pulumi.Input[bool]] = None,
                 health_check_config: Optional[pulumi.Input['GatewayVpnAttachmentHealthCheckConfigArgs']] = None,
                 ike_config: Optional[pulumi.Input['GatewayVpnAttachmentIkeConfigArgs']] = None,
                 ipsec_config: Optional[pulumi.Input['GatewayVpnAttachmentIpsecConfigArgs']] = None,
                 local_subnet: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 remote_subnet: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpn_attachment_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GatewayVpnAttachment resources.
        :param pulumi.Input['GatewayVpnAttachmentBgpConfigArgs'] bgp_config: Bgp configuration information. See the following `Block bgp_config`.
        :param pulumi.Input[str] customer_gateway_id: The ID of the customer gateway.
        :param pulumi.Input[bool] effect_immediately: Indicates whether IPsec-VPN negotiations are initiated immediately. Valid values.
        :param pulumi.Input[bool] enable_dpd: Whether to enable the DPD (peer survival detection) function.
        :param pulumi.Input[bool] enable_nat_traversal: Allow NAT penetration.
        :param pulumi.Input['GatewayVpnAttachmentHealthCheckConfigArgs'] health_check_config: Health check configuration information. See the following `Block health_check_config`.
        :param pulumi.Input['GatewayVpnAttachmentIkeConfigArgs'] ike_config: Configuration negotiated in the second stage. See the following `Block ike_config`.
        :param pulumi.Input['GatewayVpnAttachmentIpsecConfigArgs'] ipsec_config: Configuration negotiated in the second stage. See the following `Block ipsec_config`.
        :param pulumi.Input[str] local_subnet: The CIDR block of the virtual private cloud (VPC).
        :param pulumi.Input[str] network_type: The network type of the IPsec connection. Valid values: `public`, `private`.
        :param pulumi.Input[str] remote_subnet: The CIDR block of the on-premises data center.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[str] vpn_attachment_name: The name of the vpn attachment.
        """
        if bgp_config is not None:
            pulumi.set(__self__, "bgp_config", bgp_config)
        if customer_gateway_id is not None:
            pulumi.set(__self__, "customer_gateway_id", customer_gateway_id)
        if effect_immediately is not None:
            pulumi.set(__self__, "effect_immediately", effect_immediately)
        if enable_dpd is not None:
            pulumi.set(__self__, "enable_dpd", enable_dpd)
        if enable_nat_traversal is not None:
            pulumi.set(__self__, "enable_nat_traversal", enable_nat_traversal)
        if health_check_config is not None:
            pulumi.set(__self__, "health_check_config", health_check_config)
        if ike_config is not None:
            pulumi.set(__self__, "ike_config", ike_config)
        if ipsec_config is not None:
            pulumi.set(__self__, "ipsec_config", ipsec_config)
        if local_subnet is not None:
            pulumi.set(__self__, "local_subnet", local_subnet)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if remote_subnet is not None:
            pulumi.set(__self__, "remote_subnet", remote_subnet)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpn_attachment_name is not None:
            pulumi.set(__self__, "vpn_attachment_name", vpn_attachment_name)

    @property
    @pulumi.getter(name="bgpConfig")
    def bgp_config(self) -> Optional[pulumi.Input['GatewayVpnAttachmentBgpConfigArgs']]:
        """
        Bgp configuration information. See the following `Block bgp_config`.
        """
        return pulumi.get(self, "bgp_config")

    @bgp_config.setter
    def bgp_config(self, value: Optional[pulumi.Input['GatewayVpnAttachmentBgpConfigArgs']]):
        pulumi.set(self, "bgp_config", value)

    @property
    @pulumi.getter(name="customerGatewayId")
    def customer_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the customer gateway.
        """
        return pulumi.get(self, "customer_gateway_id")

    @customer_gateway_id.setter
    def customer_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_gateway_id", value)

    @property
    @pulumi.getter(name="effectImmediately")
    def effect_immediately(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether IPsec-VPN negotiations are initiated immediately. Valid values.
        """
        return pulumi.get(self, "effect_immediately")

    @effect_immediately.setter
    def effect_immediately(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "effect_immediately", value)

    @property
    @pulumi.getter(name="enableDpd")
    def enable_dpd(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the DPD (peer survival detection) function.
        """
        return pulumi.get(self, "enable_dpd")

    @enable_dpd.setter
    def enable_dpd(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_dpd", value)

    @property
    @pulumi.getter(name="enableNatTraversal")
    def enable_nat_traversal(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow NAT penetration.
        """
        return pulumi.get(self, "enable_nat_traversal")

    @enable_nat_traversal.setter
    def enable_nat_traversal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_nat_traversal", value)

    @property
    @pulumi.getter(name="healthCheckConfig")
    def health_check_config(self) -> Optional[pulumi.Input['GatewayVpnAttachmentHealthCheckConfigArgs']]:
        """
        Health check configuration information. See the following `Block health_check_config`.
        """
        return pulumi.get(self, "health_check_config")

    @health_check_config.setter
    def health_check_config(self, value: Optional[pulumi.Input['GatewayVpnAttachmentHealthCheckConfigArgs']]):
        pulumi.set(self, "health_check_config", value)

    @property
    @pulumi.getter(name="ikeConfig")
    def ike_config(self) -> Optional[pulumi.Input['GatewayVpnAttachmentIkeConfigArgs']]:
        """
        Configuration negotiated in the second stage. See the following `Block ike_config`.
        """
        return pulumi.get(self, "ike_config")

    @ike_config.setter
    def ike_config(self, value: Optional[pulumi.Input['GatewayVpnAttachmentIkeConfigArgs']]):
        pulumi.set(self, "ike_config", value)

    @property
    @pulumi.getter(name="ipsecConfig")
    def ipsec_config(self) -> Optional[pulumi.Input['GatewayVpnAttachmentIpsecConfigArgs']]:
        """
        Configuration negotiated in the second stage. See the following `Block ipsec_config`.
        """
        return pulumi.get(self, "ipsec_config")

    @ipsec_config.setter
    def ipsec_config(self, value: Optional[pulumi.Input['GatewayVpnAttachmentIpsecConfigArgs']]):
        pulumi.set(self, "ipsec_config", value)

    @property
    @pulumi.getter(name="localSubnet")
    def local_subnet(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block of the virtual private cloud (VPC).
        """
        return pulumi.get(self, "local_subnet")

    @local_subnet.setter
    def local_subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_subnet", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        The network type of the IPsec connection. Valid values: `public`, `private`.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter(name="remoteSubnet")
    def remote_subnet(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block of the on-premises data center.
        """
        return pulumi.get(self, "remote_subnet")

    @remote_subnet.setter
    def remote_subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_subnet", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpnAttachmentName")
    def vpn_attachment_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the vpn attachment.
        """
        return pulumi.get(self, "vpn_attachment_name")

    @vpn_attachment_name.setter
    def vpn_attachment_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_attachment_name", value)


class GatewayVpnAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bgp_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentBgpConfigArgs']]] = None,
                 customer_gateway_id: Optional[pulumi.Input[str]] = None,
                 effect_immediately: Optional[pulumi.Input[bool]] = None,
                 enable_dpd: Optional[pulumi.Input[bool]] = None,
                 enable_nat_traversal: Optional[pulumi.Input[bool]] = None,
                 health_check_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentHealthCheckConfigArgs']]] = None,
                 ike_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentIkeConfigArgs']]] = None,
                 ipsec_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentIpsecConfigArgs']]] = None,
                 local_subnet: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 remote_subnet: Optional[pulumi.Input[str]] = None,
                 vpn_attachment_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a VPN Gateway Vpn Attachment resource.

        For information about VPN Gateway Vpn Attachment and how to use it, see [What is Vpn Attachment](https://www.alibabacloud.com/help/zh/virtual-private-cloud/latest/createvpnattachment).

        > **NOTE:** Available in v1.181.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_customer_gateway = alicloud.vpn.CustomerGateway("defaultCustomerGateway",
            ip_address="42.104.22.210",
            asn="45014",
            description="testAccVpnConnectionDesc")
        default_gateway_vpn_attachment = alicloud.vpn.GatewayVpnAttachment("defaultGatewayVpnAttachment",
            customer_gateway_id=default_customer_gateway.id,
            network_type="public",
            local_subnet="0.0.0.0/0",
            remote_subnet="0.0.0.0/0",
            effect_immediately=False,
            ike_config=alicloud.vpn.GatewayVpnAttachmentIkeConfigArgs(
                ike_auth_alg="md5",
                ike_enc_alg="des",
                ike_version="ikev2",
                ike_mode="main",
                ike_lifetime=86400,
                psk="tf-testvpn2",
                ike_pfs="group1",
                remote_id="testbob2",
                local_id="testalice2",
            ),
            ipsec_config=alicloud.vpn.GatewayVpnAttachmentIpsecConfigArgs(
                ipsec_pfs="group5",
                ipsec_enc_alg="des",
                ipsec_auth_alg="md5",
                ipsec_lifetime=86400,
            ),
            bgp_config=alicloud.vpn.GatewayVpnAttachmentBgpConfigArgs(
                enable=True,
                local_asn=45014,
                tunnel_cidr="169.254.11.0/30",
                local_bgp_ip="169.254.11.1",
            ),
            health_check_config=alicloud.vpn.GatewayVpnAttachmentHealthCheckConfigArgs(
                enable=True,
                sip="192.168.1.1",
                dip="10.0.0.1",
                interval=10,
                retry=10,
                policy="revoke_route",
            ),
            enable_dpd=True,
            enable_nat_traversal=True,
            vpn_attachment_name=var["name"])
        ```

        ## Import

        VPN Gateway Vpn Attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpn/gatewayVpnAttachment:GatewayVpnAttachment example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GatewayVpnAttachmentBgpConfigArgs']] bgp_config: Bgp configuration information. See the following `Block bgp_config`.
        :param pulumi.Input[str] customer_gateway_id: The ID of the customer gateway.
        :param pulumi.Input[bool] effect_immediately: Indicates whether IPsec-VPN negotiations are initiated immediately. Valid values.
        :param pulumi.Input[bool] enable_dpd: Whether to enable the DPD (peer survival detection) function.
        :param pulumi.Input[bool] enable_nat_traversal: Allow NAT penetration.
        :param pulumi.Input[pulumi.InputType['GatewayVpnAttachmentHealthCheckConfigArgs']] health_check_config: Health check configuration information. See the following `Block health_check_config`.
        :param pulumi.Input[pulumi.InputType['GatewayVpnAttachmentIkeConfigArgs']] ike_config: Configuration negotiated in the second stage. See the following `Block ike_config`.
        :param pulumi.Input[pulumi.InputType['GatewayVpnAttachmentIpsecConfigArgs']] ipsec_config: Configuration negotiated in the second stage. See the following `Block ipsec_config`.
        :param pulumi.Input[str] local_subnet: The CIDR block of the virtual private cloud (VPC).
        :param pulumi.Input[str] network_type: The network type of the IPsec connection. Valid values: `public`, `private`.
        :param pulumi.Input[str] remote_subnet: The CIDR block of the on-premises data center.
        :param pulumi.Input[str] vpn_attachment_name: The name of the vpn attachment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewayVpnAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPN Gateway Vpn Attachment resource.

        For information about VPN Gateway Vpn Attachment and how to use it, see [What is Vpn Attachment](https://www.alibabacloud.com/help/zh/virtual-private-cloud/latest/createvpnattachment).

        > **NOTE:** Available in v1.181.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_customer_gateway = alicloud.vpn.CustomerGateway("defaultCustomerGateway",
            ip_address="42.104.22.210",
            asn="45014",
            description="testAccVpnConnectionDesc")
        default_gateway_vpn_attachment = alicloud.vpn.GatewayVpnAttachment("defaultGatewayVpnAttachment",
            customer_gateway_id=default_customer_gateway.id,
            network_type="public",
            local_subnet="0.0.0.0/0",
            remote_subnet="0.0.0.0/0",
            effect_immediately=False,
            ike_config=alicloud.vpn.GatewayVpnAttachmentIkeConfigArgs(
                ike_auth_alg="md5",
                ike_enc_alg="des",
                ike_version="ikev2",
                ike_mode="main",
                ike_lifetime=86400,
                psk="tf-testvpn2",
                ike_pfs="group1",
                remote_id="testbob2",
                local_id="testalice2",
            ),
            ipsec_config=alicloud.vpn.GatewayVpnAttachmentIpsecConfigArgs(
                ipsec_pfs="group5",
                ipsec_enc_alg="des",
                ipsec_auth_alg="md5",
                ipsec_lifetime=86400,
            ),
            bgp_config=alicloud.vpn.GatewayVpnAttachmentBgpConfigArgs(
                enable=True,
                local_asn=45014,
                tunnel_cidr="169.254.11.0/30",
                local_bgp_ip="169.254.11.1",
            ),
            health_check_config=alicloud.vpn.GatewayVpnAttachmentHealthCheckConfigArgs(
                enable=True,
                sip="192.168.1.1",
                dip="10.0.0.1",
                interval=10,
                retry=10,
                policy="revoke_route",
            ),
            enable_dpd=True,
            enable_nat_traversal=True,
            vpn_attachment_name=var["name"])
        ```

        ## Import

        VPN Gateway Vpn Attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpn/gatewayVpnAttachment:GatewayVpnAttachment example <id>
        ```

        :param str resource_name: The name of the resource.
        :param GatewayVpnAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewayVpnAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bgp_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentBgpConfigArgs']]] = None,
                 customer_gateway_id: Optional[pulumi.Input[str]] = None,
                 effect_immediately: Optional[pulumi.Input[bool]] = None,
                 enable_dpd: Optional[pulumi.Input[bool]] = None,
                 enable_nat_traversal: Optional[pulumi.Input[bool]] = None,
                 health_check_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentHealthCheckConfigArgs']]] = None,
                 ike_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentIkeConfigArgs']]] = None,
                 ipsec_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentIpsecConfigArgs']]] = None,
                 local_subnet: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 remote_subnet: Optional[pulumi.Input[str]] = None,
                 vpn_attachment_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GatewayVpnAttachmentArgs.__new__(GatewayVpnAttachmentArgs)

            __props__.__dict__["bgp_config"] = bgp_config
            if customer_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'customer_gateway_id'")
            __props__.__dict__["customer_gateway_id"] = customer_gateway_id
            __props__.__dict__["effect_immediately"] = effect_immediately
            __props__.__dict__["enable_dpd"] = enable_dpd
            __props__.__dict__["enable_nat_traversal"] = enable_nat_traversal
            __props__.__dict__["health_check_config"] = health_check_config
            __props__.__dict__["ike_config"] = ike_config
            __props__.__dict__["ipsec_config"] = ipsec_config
            if local_subnet is None and not opts.urn:
                raise TypeError("Missing required property 'local_subnet'")
            __props__.__dict__["local_subnet"] = local_subnet
            __props__.__dict__["network_type"] = network_type
            if remote_subnet is None and not opts.urn:
                raise TypeError("Missing required property 'remote_subnet'")
            __props__.__dict__["remote_subnet"] = remote_subnet
            __props__.__dict__["vpn_attachment_name"] = vpn_attachment_name
            __props__.__dict__["status"] = None
        super(GatewayVpnAttachment, __self__).__init__(
            'alicloud:vpn/gatewayVpnAttachment:GatewayVpnAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bgp_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentBgpConfigArgs']]] = None,
            customer_gateway_id: Optional[pulumi.Input[str]] = None,
            effect_immediately: Optional[pulumi.Input[bool]] = None,
            enable_dpd: Optional[pulumi.Input[bool]] = None,
            enable_nat_traversal: Optional[pulumi.Input[bool]] = None,
            health_check_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentHealthCheckConfigArgs']]] = None,
            ike_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentIkeConfigArgs']]] = None,
            ipsec_config: Optional[pulumi.Input[pulumi.InputType['GatewayVpnAttachmentIpsecConfigArgs']]] = None,
            local_subnet: Optional[pulumi.Input[str]] = None,
            network_type: Optional[pulumi.Input[str]] = None,
            remote_subnet: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vpn_attachment_name: Optional[pulumi.Input[str]] = None) -> 'GatewayVpnAttachment':
        """
        Get an existing GatewayVpnAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GatewayVpnAttachmentBgpConfigArgs']] bgp_config: Bgp configuration information. See the following `Block bgp_config`.
        :param pulumi.Input[str] customer_gateway_id: The ID of the customer gateway.
        :param pulumi.Input[bool] effect_immediately: Indicates whether IPsec-VPN negotiations are initiated immediately. Valid values.
        :param pulumi.Input[bool] enable_dpd: Whether to enable the DPD (peer survival detection) function.
        :param pulumi.Input[bool] enable_nat_traversal: Allow NAT penetration.
        :param pulumi.Input[pulumi.InputType['GatewayVpnAttachmentHealthCheckConfigArgs']] health_check_config: Health check configuration information. See the following `Block health_check_config`.
        :param pulumi.Input[pulumi.InputType['GatewayVpnAttachmentIkeConfigArgs']] ike_config: Configuration negotiated in the second stage. See the following `Block ike_config`.
        :param pulumi.Input[pulumi.InputType['GatewayVpnAttachmentIpsecConfigArgs']] ipsec_config: Configuration negotiated in the second stage. See the following `Block ipsec_config`.
        :param pulumi.Input[str] local_subnet: The CIDR block of the virtual private cloud (VPC).
        :param pulumi.Input[str] network_type: The network type of the IPsec connection. Valid values: `public`, `private`.
        :param pulumi.Input[str] remote_subnet: The CIDR block of the on-premises data center.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[str] vpn_attachment_name: The name of the vpn attachment.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewayVpnAttachmentState.__new__(_GatewayVpnAttachmentState)

        __props__.__dict__["bgp_config"] = bgp_config
        __props__.__dict__["customer_gateway_id"] = customer_gateway_id
        __props__.__dict__["effect_immediately"] = effect_immediately
        __props__.__dict__["enable_dpd"] = enable_dpd
        __props__.__dict__["enable_nat_traversal"] = enable_nat_traversal
        __props__.__dict__["health_check_config"] = health_check_config
        __props__.__dict__["ike_config"] = ike_config
        __props__.__dict__["ipsec_config"] = ipsec_config
        __props__.__dict__["local_subnet"] = local_subnet
        __props__.__dict__["network_type"] = network_type
        __props__.__dict__["remote_subnet"] = remote_subnet
        __props__.__dict__["status"] = status
        __props__.__dict__["vpn_attachment_name"] = vpn_attachment_name
        return GatewayVpnAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bgpConfig")
    def bgp_config(self) -> pulumi.Output['outputs.GatewayVpnAttachmentBgpConfig']:
        """
        Bgp configuration information. See the following `Block bgp_config`.
        """
        return pulumi.get(self, "bgp_config")

    @property
    @pulumi.getter(name="customerGatewayId")
    def customer_gateway_id(self) -> pulumi.Output[str]:
        """
        The ID of the customer gateway.
        """
        return pulumi.get(self, "customer_gateway_id")

    @property
    @pulumi.getter(name="effectImmediately")
    def effect_immediately(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether IPsec-VPN negotiations are initiated immediately. Valid values.
        """
        return pulumi.get(self, "effect_immediately")

    @property
    @pulumi.getter(name="enableDpd")
    def enable_dpd(self) -> pulumi.Output[bool]:
        """
        Whether to enable the DPD (peer survival detection) function.
        """
        return pulumi.get(self, "enable_dpd")

    @property
    @pulumi.getter(name="enableNatTraversal")
    def enable_nat_traversal(self) -> pulumi.Output[bool]:
        """
        Allow NAT penetration.
        """
        return pulumi.get(self, "enable_nat_traversal")

    @property
    @pulumi.getter(name="healthCheckConfig")
    def health_check_config(self) -> pulumi.Output['outputs.GatewayVpnAttachmentHealthCheckConfig']:
        """
        Health check configuration information. See the following `Block health_check_config`.
        """
        return pulumi.get(self, "health_check_config")

    @property
    @pulumi.getter(name="ikeConfig")
    def ike_config(self) -> pulumi.Output['outputs.GatewayVpnAttachmentIkeConfig']:
        """
        Configuration negotiated in the second stage. See the following `Block ike_config`.
        """
        return pulumi.get(self, "ike_config")

    @property
    @pulumi.getter(name="ipsecConfig")
    def ipsec_config(self) -> pulumi.Output['outputs.GatewayVpnAttachmentIpsecConfig']:
        """
        Configuration negotiated in the second stage. See the following `Block ipsec_config`.
        """
        return pulumi.get(self, "ipsec_config")

    @property
    @pulumi.getter(name="localSubnet")
    def local_subnet(self) -> pulumi.Output[str]:
        """
        The CIDR block of the virtual private cloud (VPC).
        """
        return pulumi.get(self, "local_subnet")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Output[str]:
        """
        The network type of the IPsec connection. Valid values: `public`, `private`.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="remoteSubnet")
    def remote_subnet(self) -> pulumi.Output[str]:
        """
        The CIDR block of the on-premises data center.
        """
        return pulumi.get(self, "remote_subnet")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpnAttachmentName")
    def vpn_attachment_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the vpn attachment.
        """
        return pulumi.get(self, "vpn_attachment_name")

