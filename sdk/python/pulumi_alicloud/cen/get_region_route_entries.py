# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetRegionRouteEntriesResult:
    """
    A collection of values returned by getRegionRouteEntries.
    """
    def __init__(__self__, entries=None, id=None, instance_id=None, output_file=None, region_id=None):
        if entries and not isinstance(entries, list):
            raise TypeError("Expected argument 'entries' to be a list")
        __self__.entries = entries
        """
        A list of CEN Route Entries. Each element contains the following attributes:
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        __self__.instance_id = instance_id
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if region_id and not isinstance(region_id, str):
            raise TypeError("Expected argument 'region_id' to be a str")
        __self__.region_id = region_id
class AwaitableGetRegionRouteEntriesResult(GetRegionRouteEntriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionRouteEntriesResult(
            entries=self.entries,
            id=self.id,
            instance_id=self.instance_id,
            output_file=self.output_file,
            region_id=self.region_id)

def get_region_route_entries(instance_id=None,output_file=None,region_id=None,opts=None):
    """
    This data source provides CEN Regional Route Entries available to the user.




    :param str instance_id: ID of the CEN instance.
    :param str region_id: ID of the region.
    """
    __args__ = dict()


    __args__['instanceId'] = instance_id
    __args__['outputFile'] = output_file
    __args__['regionId'] = region_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getRegionRouteEntries:getRegionRouteEntries', __args__, opts=opts).value

    return AwaitableGetRegionRouteEntriesResult(
        entries=__ret__.get('entries'),
        id=__ret__.get('id'),
        instance_id=__ret__.get('instanceId'),
        output_file=__ret__.get('outputFile'),
        region_id=__ret__.get('regionId'))
