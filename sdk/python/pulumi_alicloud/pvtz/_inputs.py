# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'EndpointIpConfigArgs',
    'RuleAttachmentVpcArgs',
    'RuleForwardIpArgs',
    'ZoneAttachmentVpcArgs',
    'ZoneUserInfoArgs',
]

@pulumi.input_type
class EndpointIpConfigArgs:
    def __init__(__self__, *,
                 cidr_block: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 ip: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] cidr_block: The Subnet mask.
        :param pulumi.Input[str] vswitch_id: The Vswitch id.
        :param pulumi.Input[str] zone_id: The Zone ID.
        :param pulumi.Input[str] ip: The IP address within the parameter range of the subnet mask.  It is recommended to use the IP address assigned by the system.
        """
        pulumi.set(__self__, "cidr_block", cidr_block)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        pulumi.set(__self__, "zone_id", zone_id)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Input[str]:
        """
        The Subnet mask.
        """
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        The Vswitch id.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        The Zone ID.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address within the parameter range of the subnet mask.  It is recommended to use the IP address assigned by the system.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)


@pulumi.input_type
class RuleAttachmentVpcArgs:
    def __init__(__self__, *,
                 region_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] region_id: The region of the vpc. If not set, the current region will instead of.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.  **NOTE:** The VPC that can be associated with the forwarding rule must belong to the same region as the Endpoint.
        """
        pulumi.set(__self__, "region_id", region_id)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> pulumi.Input[str]:
        """
        The region of the vpc. If not set, the current region will instead of.
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPC.  **NOTE:** The VPC that can be associated with the forwarding rule must belong to the same region as the Endpoint.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class RuleForwardIpArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 port: pulumi.Input[int]):
        """
        :param pulumi.Input[str] ip: The ip of the forwarding destination.
        :param pulumi.Input[int] port: The port of the forwarding destination.
        """
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        The ip of the forwarding destination.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        The port of the forwarding destination.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class ZoneAttachmentVpcArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 region_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] vpc_id: The Id of the vpc.
        :param pulumi.Input[str] region_id: The region of the vpc. If not set, the current region will instead of.
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        if region_id is not None:
            pulumi.set(__self__, "region_id", region_id)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The Id of the vpc.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the vpc. If not set, the current region will instead of.
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_id", value)


@pulumi.input_type
class ZoneUserInfoArgs:
    def __init__(__self__, *,
                 region_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] region_ids: The list of the region IDs.
        :param pulumi.Input[str] user_id: The user ID belonging to the region is used for cross-account synchronization scenarios.
        """
        if region_ids is not None:
            pulumi.set(__self__, "region_ids", region_ids)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="regionIds")
    def region_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of the region IDs.
        """
        return pulumi.get(self, "region_ids")

    @region_ids.setter
    def region_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "region_ids", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The user ID belonging to the region is used for cross-account synchronization scenarios.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


