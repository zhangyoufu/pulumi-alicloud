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

__all__ = [
    'GetNatIpsResult',
    'AwaitableGetNatIpsResult',
    'get_nat_ips',
    'get_nat_ips_output',
]

@pulumi.output_type
class GetNatIpsResult:
    """
    A collection of values returned by getNatIps.
    """
    def __init__(__self__, id=None, ids=None, ips=None, name_regex=None, names=None, nat_gateway_id=None, nat_ip_cidr=None, nat_ip_ids=None, nat_ip_names=None, output_file=None, status=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if nat_gateway_id and not isinstance(nat_gateway_id, str):
            raise TypeError("Expected argument 'nat_gateway_id' to be a str")
        pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        if nat_ip_cidr and not isinstance(nat_ip_cidr, str):
            raise TypeError("Expected argument 'nat_ip_cidr' to be a str")
        pulumi.set(__self__, "nat_ip_cidr", nat_ip_cidr)
        if nat_ip_ids and not isinstance(nat_ip_ids, list):
            raise TypeError("Expected argument 'nat_ip_ids' to be a list")
        pulumi.set(__self__, "nat_ip_ids", nat_ip_ids)
        if nat_ip_names and not isinstance(nat_ip_names, list):
            raise TypeError("Expected argument 'nat_ip_names' to be a list")
        pulumi.set(__self__, "nat_ip_names", nat_ip_names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def ips(self) -> Sequence['outputs.GetNatIpsIpResult']:
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> str:
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="natIpCidr")
    def nat_ip_cidr(self) -> Optional[str]:
        return pulumi.get(self, "nat_ip_cidr")

    @property
    @pulumi.getter(name="natIpIds")
    def nat_ip_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "nat_ip_ids")

    @property
    @pulumi.getter(name="natIpNames")
    def nat_ip_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "nat_ip_names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetNatIpsResult(GetNatIpsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNatIpsResult(
            id=self.id,
            ids=self.ids,
            ips=self.ips,
            name_regex=self.name_regex,
            names=self.names,
            nat_gateway_id=self.nat_gateway_id,
            nat_ip_cidr=self.nat_ip_cidr,
            nat_ip_ids=self.nat_ip_ids,
            nat_ip_names=self.nat_ip_names,
            output_file=self.output_file,
            status=self.status)


def get_nat_ips(ids: Optional[Sequence[str]] = None,
                name_regex: Optional[str] = None,
                nat_gateway_id: Optional[str] = None,
                nat_ip_cidr: Optional[str] = None,
                nat_ip_ids: Optional[Sequence[str]] = None,
                nat_ip_names: Optional[Sequence[str]] = None,
                output_file: Optional[str] = None,
                status: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNatIpsResult:
    """
    This data source provides the Vpc Nat Ips of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.136.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("vpcNatIpId1", ids.ips[0].id)
    name_regex = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        name_regex="^my-NatIp")
    pulumi.export("vpcNatIpId2", name_regex.ips[0].id)
    nat_ip_cidr = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        nat_ip_cidr="example_value",
        name_regex="^my-NatIp")
    pulumi.export("vpcNatIpId3", nat_ip_cidr.ips[0].id)
    nat_ip_name = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        ids=["example_value"],
        nat_ip_names=["example_value"])
    pulumi.export("vpcNatIpId4", nat_ip_name.ips[0].id)
    nat_ip_ids = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        ids=["example_value"],
        nat_ip_ids=["example_value"])
    pulumi.export("vpcNatIpId5", nat_ip_ids.ips[0].id)
    status = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        ids=["example_value"],
        status="example_value")
    pulumi.export("vpcNatIpId6", status.ips[0].id)
    ```


    :param Sequence[str] ids: A list of Nat Ip IDs.
    :param str name_regex: A regex string to filter results by Nat Ip name.
    :param str nat_gateway_id: The ID of the Virtual Private Cloud (VPC) NAT gateway to which the NAT IP address belongs.
    :param str nat_ip_cidr: The CIDR block to which the NAT IP address belongs.
    :param Sequence[str] nat_ip_names: The name of the NAT IP address.
    :param str status: The status of the NAT IP address. Valid values: `Available`, `Deleting` and `Creating`.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['natGatewayId'] = nat_gateway_id
    __args__['natIpCidr'] = nat_ip_cidr
    __args__['natIpIds'] = nat_ip_ids
    __args__['natIpNames'] = nat_ip_names
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:vpc/getNatIps:getNatIps', __args__, opts=opts, typ=GetNatIpsResult).value

    return AwaitableGetNatIpsResult(
        id=__ret__.id,
        ids=__ret__.ids,
        ips=__ret__.ips,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        nat_gateway_id=__ret__.nat_gateway_id,
        nat_ip_cidr=__ret__.nat_ip_cidr,
        nat_ip_ids=__ret__.nat_ip_ids,
        nat_ip_names=__ret__.nat_ip_names,
        output_file=__ret__.output_file,
        status=__ret__.status)


@_utilities.lift_output_func(get_nat_ips)
def get_nat_ips_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                       nat_gateway_id: Optional[pulumi.Input[str]] = None,
                       nat_ip_cidr: Optional[pulumi.Input[Optional[str]]] = None,
                       nat_ip_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       nat_ip_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       status: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNatIpsResult]:
    """
    This data source provides the Vpc Nat Ips of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.136.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("vpcNatIpId1", ids.ips[0].id)
    name_regex = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        name_regex="^my-NatIp")
    pulumi.export("vpcNatIpId2", name_regex.ips[0].id)
    nat_ip_cidr = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        nat_ip_cidr="example_value",
        name_regex="^my-NatIp")
    pulumi.export("vpcNatIpId3", nat_ip_cidr.ips[0].id)
    nat_ip_name = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        ids=["example_value"],
        nat_ip_names=["example_value"])
    pulumi.export("vpcNatIpId4", nat_ip_name.ips[0].id)
    nat_ip_ids = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        ids=["example_value"],
        nat_ip_ids=["example_value"])
    pulumi.export("vpcNatIpId5", nat_ip_ids.ips[0].id)
    status = alicloud.vpc.get_nat_ips(nat_gateway_id="example_value",
        ids=["example_value"],
        status="example_value")
    pulumi.export("vpcNatIpId6", status.ips[0].id)
    ```


    :param Sequence[str] ids: A list of Nat Ip IDs.
    :param str name_regex: A regex string to filter results by Nat Ip name.
    :param str nat_gateway_id: The ID of the Virtual Private Cloud (VPC) NAT gateway to which the NAT IP address belongs.
    :param str nat_ip_cidr: The CIDR block to which the NAT IP address belongs.
    :param Sequence[str] nat_ip_names: The name of the NAT IP address.
    :param str status: The status of the NAT IP address. Valid values: `Available`, `Deleting` and `Creating`.
    """
    ...
