# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VbrPconnAssociationArgs', 'VbrPconnAssociation']

@pulumi.input_type
class VbrPconnAssociationArgs:
    def __init__(__self__, *,
                 physical_connection_id: pulumi.Input[str],
                 vbr_id: pulumi.Input[str],
                 vlan_id: pulumi.Input[int],
                 enable_ipv6: Optional[pulumi.Input[bool]] = None,
                 local_gateway_ip: Optional[pulumi.Input[str]] = None,
                 local_ipv6_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peer_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peer_ipv6_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peering_ipv6_subnet_mask: Optional[pulumi.Input[str]] = None,
                 peering_subnet_mask: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VbrPconnAssociation resource.
        :param pulumi.Input[str] physical_connection_id: The ID of the leased line instance.
        :param pulumi.Input[str] vbr_id: The ID of the VBR instance.
        :param pulumi.Input[int] vlan_id: VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
        :param pulumi.Input[bool] enable_ipv6: Whether IPv6 is enabled. Value:
               - **true**: on.
               - **false** (default): Off.
        :param pulumi.Input[str] local_gateway_ip: The Alibaba cloud IP address of the VBR instance.
        :param pulumi.Input[str] local_ipv6_gateway_ip: The IPv6 address on the Alibaba Cloud side of the VBR instance.
        :param pulumi.Input[str] peer_gateway_ip: The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        :param pulumi.Input[str] peer_ipv6_gateway_ip: The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        :param pulumi.Input[str] peering_ipv6_subnet_mask: The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
        :param pulumi.Input[str] peering_subnet_mask: The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
        """
        pulumi.set(__self__, "physical_connection_id", physical_connection_id)
        pulumi.set(__self__, "vbr_id", vbr_id)
        pulumi.set(__self__, "vlan_id", vlan_id)
        if enable_ipv6 is not None:
            pulumi.set(__self__, "enable_ipv6", enable_ipv6)
        if local_gateway_ip is not None:
            pulumi.set(__self__, "local_gateway_ip", local_gateway_ip)
        if local_ipv6_gateway_ip is not None:
            pulumi.set(__self__, "local_ipv6_gateway_ip", local_ipv6_gateway_ip)
        if peer_gateway_ip is not None:
            pulumi.set(__self__, "peer_gateway_ip", peer_gateway_ip)
        if peer_ipv6_gateway_ip is not None:
            pulumi.set(__self__, "peer_ipv6_gateway_ip", peer_ipv6_gateway_ip)
        if peering_ipv6_subnet_mask is not None:
            pulumi.set(__self__, "peering_ipv6_subnet_mask", peering_ipv6_subnet_mask)
        if peering_subnet_mask is not None:
            pulumi.set(__self__, "peering_subnet_mask", peering_subnet_mask)

    @property
    @pulumi.getter(name="physicalConnectionId")
    def physical_connection_id(self) -> pulumi.Input[str]:
        """
        The ID of the leased line instance.
        """
        return pulumi.get(self, "physical_connection_id")

    @physical_connection_id.setter
    def physical_connection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "physical_connection_id", value)

    @property
    @pulumi.getter(name="vbrId")
    def vbr_id(self) -> pulumi.Input[str]:
        """
        The ID of the VBR instance.
        """
        return pulumi.get(self, "vbr_id")

    @vbr_id.setter
    def vbr_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vbr_id", value)

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> pulumi.Input[int]:
        """
        VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
        """
        return pulumi.get(self, "vlan_id")

    @vlan_id.setter
    def vlan_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "vlan_id", value)

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether IPv6 is enabled. Value:
        - **true**: on.
        - **false** (default): Off.
        """
        return pulumi.get(self, "enable_ipv6")

    @enable_ipv6.setter
    def enable_ipv6(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ipv6", value)

    @property
    @pulumi.getter(name="localGatewayIp")
    def local_gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The Alibaba cloud IP address of the VBR instance.
        """
        return pulumi.get(self, "local_gateway_ip")

    @local_gateway_ip.setter
    def local_gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_gateway_ip", value)

    @property
    @pulumi.getter(name="localIpv6GatewayIp")
    def local_ipv6_gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 address on the Alibaba Cloud side of the VBR instance.
        """
        return pulumi.get(self, "local_ipv6_gateway_ip")

    @local_ipv6_gateway_ip.setter
    def local_ipv6_gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_ipv6_gateway_ip", value)

    @property
    @pulumi.getter(name="peerGatewayIp")
    def peer_gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        """
        return pulumi.get(self, "peer_gateway_ip")

    @peer_gateway_ip.setter
    def peer_gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_gateway_ip", value)

    @property
    @pulumi.getter(name="peerIpv6GatewayIp")
    def peer_ipv6_gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        """
        return pulumi.get(self, "peer_ipv6_gateway_ip")

    @peer_ipv6_gateway_ip.setter
    def peer_ipv6_gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_ipv6_gateway_ip", value)

    @property
    @pulumi.getter(name="peeringIpv6SubnetMask")
    def peering_ipv6_subnet_mask(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
        """
        return pulumi.get(self, "peering_ipv6_subnet_mask")

    @peering_ipv6_subnet_mask.setter
    def peering_ipv6_subnet_mask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peering_ipv6_subnet_mask", value)

    @property
    @pulumi.getter(name="peeringSubnetMask")
    def peering_subnet_mask(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
        """
        return pulumi.get(self, "peering_subnet_mask")

    @peering_subnet_mask.setter
    def peering_subnet_mask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peering_subnet_mask", value)


@pulumi.input_type
class _VbrPconnAssociationState:
    def __init__(__self__, *,
                 circuit_code: Optional[pulumi.Input[str]] = None,
                 enable_ipv6: Optional[pulumi.Input[bool]] = None,
                 local_gateway_ip: Optional[pulumi.Input[str]] = None,
                 local_ipv6_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peer_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peer_ipv6_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peering_ipv6_subnet_mask: Optional[pulumi.Input[str]] = None,
                 peering_subnet_mask: Optional[pulumi.Input[str]] = None,
                 physical_connection_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vbr_id: Optional[pulumi.Input[str]] = None,
                 vlan_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering VbrPconnAssociation resources.
        :param pulumi.Input[str] circuit_code: The circuit code provided by the operator for the physical connection.
        :param pulumi.Input[bool] enable_ipv6: Whether IPv6 is enabled. Value:
               - **true**: on.
               - **false** (default): Off.
        :param pulumi.Input[str] local_gateway_ip: The Alibaba cloud IP address of the VBR instance.
        :param pulumi.Input[str] local_ipv6_gateway_ip: The IPv6 address on the Alibaba Cloud side of the VBR instance.
        :param pulumi.Input[str] peer_gateway_ip: The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        :param pulumi.Input[str] peer_ipv6_gateway_ip: The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        :param pulumi.Input[str] peering_ipv6_subnet_mask: The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
        :param pulumi.Input[str] peering_subnet_mask: The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
        :param pulumi.Input[str] physical_connection_id: The ID of the leased line instance.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[str] vbr_id: The ID of the VBR instance.
        :param pulumi.Input[int] vlan_id: VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
        """
        if circuit_code is not None:
            pulumi.set(__self__, "circuit_code", circuit_code)
        if enable_ipv6 is not None:
            pulumi.set(__self__, "enable_ipv6", enable_ipv6)
        if local_gateway_ip is not None:
            pulumi.set(__self__, "local_gateway_ip", local_gateway_ip)
        if local_ipv6_gateway_ip is not None:
            pulumi.set(__self__, "local_ipv6_gateway_ip", local_ipv6_gateway_ip)
        if peer_gateway_ip is not None:
            pulumi.set(__self__, "peer_gateway_ip", peer_gateway_ip)
        if peer_ipv6_gateway_ip is not None:
            pulumi.set(__self__, "peer_ipv6_gateway_ip", peer_ipv6_gateway_ip)
        if peering_ipv6_subnet_mask is not None:
            pulumi.set(__self__, "peering_ipv6_subnet_mask", peering_ipv6_subnet_mask)
        if peering_subnet_mask is not None:
            pulumi.set(__self__, "peering_subnet_mask", peering_subnet_mask)
        if physical_connection_id is not None:
            pulumi.set(__self__, "physical_connection_id", physical_connection_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vbr_id is not None:
            pulumi.set(__self__, "vbr_id", vbr_id)
        if vlan_id is not None:
            pulumi.set(__self__, "vlan_id", vlan_id)

    @property
    @pulumi.getter(name="circuitCode")
    def circuit_code(self) -> Optional[pulumi.Input[str]]:
        """
        The circuit code provided by the operator for the physical connection.
        """
        return pulumi.get(self, "circuit_code")

    @circuit_code.setter
    def circuit_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "circuit_code", value)

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether IPv6 is enabled. Value:
        - **true**: on.
        - **false** (default): Off.
        """
        return pulumi.get(self, "enable_ipv6")

    @enable_ipv6.setter
    def enable_ipv6(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ipv6", value)

    @property
    @pulumi.getter(name="localGatewayIp")
    def local_gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The Alibaba cloud IP address of the VBR instance.
        """
        return pulumi.get(self, "local_gateway_ip")

    @local_gateway_ip.setter
    def local_gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_gateway_ip", value)

    @property
    @pulumi.getter(name="localIpv6GatewayIp")
    def local_ipv6_gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 address on the Alibaba Cloud side of the VBR instance.
        """
        return pulumi.get(self, "local_ipv6_gateway_ip")

    @local_ipv6_gateway_ip.setter
    def local_ipv6_gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_ipv6_gateway_ip", value)

    @property
    @pulumi.getter(name="peerGatewayIp")
    def peer_gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        """
        return pulumi.get(self, "peer_gateway_ip")

    @peer_gateway_ip.setter
    def peer_gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_gateway_ip", value)

    @property
    @pulumi.getter(name="peerIpv6GatewayIp")
    def peer_ipv6_gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        """
        return pulumi.get(self, "peer_ipv6_gateway_ip")

    @peer_ipv6_gateway_ip.setter
    def peer_ipv6_gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_ipv6_gateway_ip", value)

    @property
    @pulumi.getter(name="peeringIpv6SubnetMask")
    def peering_ipv6_subnet_mask(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
        """
        return pulumi.get(self, "peering_ipv6_subnet_mask")

    @peering_ipv6_subnet_mask.setter
    def peering_ipv6_subnet_mask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peering_ipv6_subnet_mask", value)

    @property
    @pulumi.getter(name="peeringSubnetMask")
    def peering_subnet_mask(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
        """
        return pulumi.get(self, "peering_subnet_mask")

    @peering_subnet_mask.setter
    def peering_subnet_mask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peering_subnet_mask", value)

    @property
    @pulumi.getter(name="physicalConnectionId")
    def physical_connection_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the leased line instance.
        """
        return pulumi.get(self, "physical_connection_id")

    @physical_connection_id.setter
    def physical_connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "physical_connection_id", value)

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
    @pulumi.getter(name="vbrId")
    def vbr_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VBR instance.
        """
        return pulumi.get(self, "vbr_id")

    @vbr_id.setter
    def vbr_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vbr_id", value)

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> Optional[pulumi.Input[int]]:
        """
        VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
        """
        return pulumi.get(self, "vlan_id")

    @vlan_id.setter
    def vlan_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlan_id", value)


class VbrPconnAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_ipv6: Optional[pulumi.Input[bool]] = None,
                 local_gateway_ip: Optional[pulumi.Input[str]] = None,
                 local_ipv6_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peer_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peer_ipv6_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peering_ipv6_subnet_mask: Optional[pulumi.Input[str]] = None,
                 peering_subnet_mask: Optional[pulumi.Input[str]] = None,
                 physical_connection_id: Optional[pulumi.Input[str]] = None,
                 vbr_id: Optional[pulumi.Input[str]] = None,
                 vlan_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a Express Connect Vbr Pconn Association resource.

        For information about Express Connect Vbr Pconn Association and how to use it, see [What is Vbr Pconn Association](https://www.alibabacloud.com/help/en/express-connect/latest/associatephysicalconnectiontovirtualborderrouter#doc-api-Vpc-AssociatePhysicalConnectionToVirtualBorderRouter).

        > **NOTE:** Available in v1.196.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        name_regex = alicloud.expressconnect.get_physical_connections(name_regex="^preserved-NODELETING")
        default_virtual_border_router = alicloud.expressconnect.VirtualBorderRouter("defaultVirtualBorderRouter",
            local_gateway_ip="10.0.0.1",
            peer_gateway_ip="10.0.0.2",
            peering_subnet_mask="255.255.255.252",
            physical_connection_id=name_regex.connections[0].id,
            virtual_border_router_name=var["name"],
            vlan_id=100,
            min_rx_interval=1000,
            min_tx_interval=1000,
            detect_multiplier=10)
        default_vbr_pconn_association = alicloud.expressconnect.VbrPconnAssociation("defaultVbrPconnAssociation",
            peer_gateway_ip="10.0.0.6",
            local_gateway_ip="10.0.0.5",
            physical_connection_id=name_regex.connections[1].id,
            vbr_id=default_virtual_border_router.id,
            peering_subnet_mask="255.255.255.252",
            vlan_id=1122,
            enable_ipv6=False)
        ```

        ## Import

        Express Connect Vbr Pconn Association can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:expressconnect/vbrPconnAssociation:VbrPconnAssociation example <VbrId>:<PhysicalConnectionId>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_ipv6: Whether IPv6 is enabled. Value:
               - **true**: on.
               - **false** (default): Off.
        :param pulumi.Input[str] local_gateway_ip: The Alibaba cloud IP address of the VBR instance.
        :param pulumi.Input[str] local_ipv6_gateway_ip: The IPv6 address on the Alibaba Cloud side of the VBR instance.
        :param pulumi.Input[str] peer_gateway_ip: The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        :param pulumi.Input[str] peer_ipv6_gateway_ip: The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        :param pulumi.Input[str] peering_ipv6_subnet_mask: The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
        :param pulumi.Input[str] peering_subnet_mask: The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
        :param pulumi.Input[str] physical_connection_id: The ID of the leased line instance.
        :param pulumi.Input[str] vbr_id: The ID of the VBR instance.
        :param pulumi.Input[int] vlan_id: VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VbrPconnAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Express Connect Vbr Pconn Association resource.

        For information about Express Connect Vbr Pconn Association and how to use it, see [What is Vbr Pconn Association](https://www.alibabacloud.com/help/en/express-connect/latest/associatephysicalconnectiontovirtualborderrouter#doc-api-Vpc-AssociatePhysicalConnectionToVirtualBorderRouter).

        > **NOTE:** Available in v1.196.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        name_regex = alicloud.expressconnect.get_physical_connections(name_regex="^preserved-NODELETING")
        default_virtual_border_router = alicloud.expressconnect.VirtualBorderRouter("defaultVirtualBorderRouter",
            local_gateway_ip="10.0.0.1",
            peer_gateway_ip="10.0.0.2",
            peering_subnet_mask="255.255.255.252",
            physical_connection_id=name_regex.connections[0].id,
            virtual_border_router_name=var["name"],
            vlan_id=100,
            min_rx_interval=1000,
            min_tx_interval=1000,
            detect_multiplier=10)
        default_vbr_pconn_association = alicloud.expressconnect.VbrPconnAssociation("defaultVbrPconnAssociation",
            peer_gateway_ip="10.0.0.6",
            local_gateway_ip="10.0.0.5",
            physical_connection_id=name_regex.connections[1].id,
            vbr_id=default_virtual_border_router.id,
            peering_subnet_mask="255.255.255.252",
            vlan_id=1122,
            enable_ipv6=False)
        ```

        ## Import

        Express Connect Vbr Pconn Association can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:expressconnect/vbrPconnAssociation:VbrPconnAssociation example <VbrId>:<PhysicalConnectionId>
        ```

        :param str resource_name: The name of the resource.
        :param VbrPconnAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VbrPconnAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_ipv6: Optional[pulumi.Input[bool]] = None,
                 local_gateway_ip: Optional[pulumi.Input[str]] = None,
                 local_ipv6_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peer_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peer_ipv6_gateway_ip: Optional[pulumi.Input[str]] = None,
                 peering_ipv6_subnet_mask: Optional[pulumi.Input[str]] = None,
                 peering_subnet_mask: Optional[pulumi.Input[str]] = None,
                 physical_connection_id: Optional[pulumi.Input[str]] = None,
                 vbr_id: Optional[pulumi.Input[str]] = None,
                 vlan_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VbrPconnAssociationArgs.__new__(VbrPconnAssociationArgs)

            __props__.__dict__["enable_ipv6"] = enable_ipv6
            __props__.__dict__["local_gateway_ip"] = local_gateway_ip
            __props__.__dict__["local_ipv6_gateway_ip"] = local_ipv6_gateway_ip
            __props__.__dict__["peer_gateway_ip"] = peer_gateway_ip
            __props__.__dict__["peer_ipv6_gateway_ip"] = peer_ipv6_gateway_ip
            __props__.__dict__["peering_ipv6_subnet_mask"] = peering_ipv6_subnet_mask
            __props__.__dict__["peering_subnet_mask"] = peering_subnet_mask
            if physical_connection_id is None and not opts.urn:
                raise TypeError("Missing required property 'physical_connection_id'")
            __props__.__dict__["physical_connection_id"] = physical_connection_id
            if vbr_id is None and not opts.urn:
                raise TypeError("Missing required property 'vbr_id'")
            __props__.__dict__["vbr_id"] = vbr_id
            if vlan_id is None and not opts.urn:
                raise TypeError("Missing required property 'vlan_id'")
            __props__.__dict__["vlan_id"] = vlan_id
            __props__.__dict__["circuit_code"] = None
            __props__.__dict__["status"] = None
        super(VbrPconnAssociation, __self__).__init__(
            'alicloud:expressconnect/vbrPconnAssociation:VbrPconnAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            circuit_code: Optional[pulumi.Input[str]] = None,
            enable_ipv6: Optional[pulumi.Input[bool]] = None,
            local_gateway_ip: Optional[pulumi.Input[str]] = None,
            local_ipv6_gateway_ip: Optional[pulumi.Input[str]] = None,
            peer_gateway_ip: Optional[pulumi.Input[str]] = None,
            peer_ipv6_gateway_ip: Optional[pulumi.Input[str]] = None,
            peering_ipv6_subnet_mask: Optional[pulumi.Input[str]] = None,
            peering_subnet_mask: Optional[pulumi.Input[str]] = None,
            physical_connection_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vbr_id: Optional[pulumi.Input[str]] = None,
            vlan_id: Optional[pulumi.Input[int]] = None) -> 'VbrPconnAssociation':
        """
        Get an existing VbrPconnAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] circuit_code: The circuit code provided by the operator for the physical connection.
        :param pulumi.Input[bool] enable_ipv6: Whether IPv6 is enabled. Value:
               - **true**: on.
               - **false** (default): Off.
        :param pulumi.Input[str] local_gateway_ip: The Alibaba cloud IP address of the VBR instance.
        :param pulumi.Input[str] local_ipv6_gateway_ip: The IPv6 address on the Alibaba Cloud side of the VBR instance.
        :param pulumi.Input[str] peer_gateway_ip: The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        :param pulumi.Input[str] peer_ipv6_gateway_ip: The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        :param pulumi.Input[str] peering_ipv6_subnet_mask: The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
        :param pulumi.Input[str] peering_subnet_mask: The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
        :param pulumi.Input[str] physical_connection_id: The ID of the leased line instance.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[str] vbr_id: The ID of the VBR instance.
        :param pulumi.Input[int] vlan_id: VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VbrPconnAssociationState.__new__(_VbrPconnAssociationState)

        __props__.__dict__["circuit_code"] = circuit_code
        __props__.__dict__["enable_ipv6"] = enable_ipv6
        __props__.__dict__["local_gateway_ip"] = local_gateway_ip
        __props__.__dict__["local_ipv6_gateway_ip"] = local_ipv6_gateway_ip
        __props__.__dict__["peer_gateway_ip"] = peer_gateway_ip
        __props__.__dict__["peer_ipv6_gateway_ip"] = peer_ipv6_gateway_ip
        __props__.__dict__["peering_ipv6_subnet_mask"] = peering_ipv6_subnet_mask
        __props__.__dict__["peering_subnet_mask"] = peering_subnet_mask
        __props__.__dict__["physical_connection_id"] = physical_connection_id
        __props__.__dict__["status"] = status
        __props__.__dict__["vbr_id"] = vbr_id
        __props__.__dict__["vlan_id"] = vlan_id
        return VbrPconnAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="circuitCode")
    def circuit_code(self) -> pulumi.Output[str]:
        """
        The circuit code provided by the operator for the physical connection.
        """
        return pulumi.get(self, "circuit_code")

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> pulumi.Output[bool]:
        """
        Whether IPv6 is enabled. Value:
        - **true**: on.
        - **false** (default): Off.
        """
        return pulumi.get(self, "enable_ipv6")

    @property
    @pulumi.getter(name="localGatewayIp")
    def local_gateway_ip(self) -> pulumi.Output[Optional[str]]:
        """
        The Alibaba cloud IP address of the VBR instance.
        """
        return pulumi.get(self, "local_gateway_ip")

    @property
    @pulumi.getter(name="localIpv6GatewayIp")
    def local_ipv6_gateway_ip(self) -> pulumi.Output[Optional[str]]:
        """
        The IPv6 address on the Alibaba Cloud side of the VBR instance.
        """
        return pulumi.get(self, "local_ipv6_gateway_ip")

    @property
    @pulumi.getter(name="peerGatewayIp")
    def peer_gateway_ip(self) -> pulumi.Output[Optional[str]]:
        """
        The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        """
        return pulumi.get(self, "peer_gateway_ip")

    @property
    @pulumi.getter(name="peerIpv6GatewayIp")
    def peer_ipv6_gateway_ip(self) -> pulumi.Output[Optional[str]]:
        """
        The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
        """
        return pulumi.get(self, "peer_ipv6_gateway_ip")

    @property
    @pulumi.getter(name="peeringIpv6SubnetMask")
    def peering_ipv6_subnet_mask(self) -> pulumi.Output[Optional[str]]:
        """
        The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
        """
        return pulumi.get(self, "peering_ipv6_subnet_mask")

    @property
    @pulumi.getter(name="peeringSubnetMask")
    def peering_subnet_mask(self) -> pulumi.Output[Optional[str]]:
        """
        The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
        """
        return pulumi.get(self, "peering_subnet_mask")

    @property
    @pulumi.getter(name="physicalConnectionId")
    def physical_connection_id(self) -> pulumi.Output[str]:
        """
        The ID of the leased line instance.
        """
        return pulumi.get(self, "physical_connection_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vbrId")
    def vbr_id(self) -> pulumi.Output[str]:
        """
        The ID of the VBR instance.
        """
        return pulumi.get(self, "vbr_id")

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> pulumi.Output[int]:
        """
        VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
        """
        return pulumi.get(self, "vlan_id")

