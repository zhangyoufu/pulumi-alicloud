# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetRouteEntriesResult',
    'AwaitableGetRouteEntriesResult',
    'get_route_entries',
]

@pulumi.output_type
class GetRouteEntriesResult:
    """
    A collection of values returned by getRouteEntries.
    """
    def __init__(__self__, cidr_block=None, entries=None, id=None, instance_id=None, output_file=None, route_table_id=None):
        if cidr_block and not isinstance(cidr_block, str):
            raise TypeError("Expected argument 'cidr_block' to be a str")
        pulumi.set(__self__, "cidr_block", cidr_block)
        if entries and not isinstance(entries, list):
            raise TypeError("Expected argument 'entries' to be a list")
        pulumi.set(__self__, "entries", entries)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if route_table_id and not isinstance(route_table_id, str):
            raise TypeError("Expected argument 'route_table_id' to be a str")
        pulumi.set(__self__, "route_table_id", route_table_id)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> Optional[str]:
        """
        The destination CIDR block of the conflicted route entry.
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter
    def entries(self) -> List['outputs.GetRouteEntriesEntryResult']:
        """
        A list of CEN Route Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "entries")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        ID of the CEN child instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> str:
        """
        ID of the route table.
        """
        return pulumi.get(self, "route_table_id")


class AwaitableGetRouteEntriesResult(GetRouteEntriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteEntriesResult(
            cidr_block=self.cidr_block,
            entries=self.entries,
            id=self.id,
            instance_id=self.instance_id,
            output_file=self.output_file,
            route_table_id=self.route_table_id)


def get_route_entries(cidr_block: Optional[str] = None,
                      instance_id: Optional[str] = None,
                      output_file: Optional[str] = None,
                      route_table_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteEntriesResult:
    """
    This data source provides CEN Route Entries available to the user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    entry = alicloud.cen.get_route_entries(instance_id="cen-id1",
        route_table_id="vtb-id1")
    pulumi.export("firstRouteEntriesRouteEntryCidrBlock", entry.entries[0].cidr_block)
    ```


    :param str cidr_block: The destination CIDR block of the route entry to query.
    :param str instance_id: ID of the CEN instance.
    :param str route_table_id: ID of the route table of the VPC or VBR.
    """
    __args__ = dict()
    __args__['cidrBlock'] = cidr_block
    __args__['instanceId'] = instance_id
    __args__['outputFile'] = output_file
    __args__['routeTableId'] = route_table_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getRouteEntries:getRouteEntries', __args__, opts=opts, typ=GetRouteEntriesResult).value

    return AwaitableGetRouteEntriesResult(
        cidr_block=__ret__.cidr_block,
        entries=__ret__.entries,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        output_file=__ret__.output_file,
        route_table_id=__ret__.route_table_id)
