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
    'GetAddressPoolsResult',
    'AwaitableGetAddressPoolsResult',
    'get_address_pools',
    'get_address_pools_output',
]

@pulumi.output_type
class GetAddressPoolsResult:
    """
    A collection of values returned by getAddressPools.
    """
    def __init__(__self__, enable_details=None, id=None, ids=None, instance_id=None, name_regex=None, names=None, output_file=None, pools=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if pools and not isinstance(pools, list):
            raise TypeError("Expected argument 'pools' to be a list")
        pulumi.set(__self__, "pools", pools)

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

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
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def pools(self) -> Sequence['outputs.GetAddressPoolsPoolResult']:
        return pulumi.get(self, "pools")


class AwaitableGetAddressPoolsResult(GetAddressPoolsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAddressPoolsResult(
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            pools=self.pools)


def get_address_pools(enable_details: Optional[bool] = None,
                      ids: Optional[Sequence[str]] = None,
                      instance_id: Optional[str] = None,
                      name_regex: Optional[str] = None,
                      output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAddressPoolsResult:
    """
    This data source provides the Alidns Address Pools of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.152.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.dns.get_address_pools(instance_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("alidnsAddressPoolId1", ids.pools[0].id)
    name_regex = alicloud.dns.get_address_pools(instance_id="example_value",
        name_regex="^my-AddressPool")
    pulumi.export("alidnsAddressPoolId2", name_regex.pools[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Address Pool IDs.
    :param str instance_id: The id of the instance.
    :param str name_regex: A regex string to filter results by Address Pool name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:dns/getAddressPools:getAddressPools', __args__, opts=opts, typ=GetAddressPoolsResult).value

    return AwaitableGetAddressPoolsResult(
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        pools=pulumi.get(__ret__, 'pools'))


@_utilities.lift_output_func(get_address_pools)
def get_address_pools_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                             ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             instance_id: Optional[pulumi.Input[str]] = None,
                             name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                             output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAddressPoolsResult]:
    """
    This data source provides the Alidns Address Pools of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.152.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.dns.get_address_pools(instance_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("alidnsAddressPoolId1", ids.pools[0].id)
    name_regex = alicloud.dns.get_address_pools(instance_id="example_value",
        name_regex="^my-AddressPool")
    pulumi.export("alidnsAddressPoolId2", name_regex.pools[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Address Pool IDs.
    :param str instance_id: The id of the instance.
    :param str name_regex: A regex string to filter results by Address Pool name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
