# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetZonesResult:
    """
    A collection of values returned by getZones.
    """
    def __init__(__self__, id=None, ids=None, keyword=None, names=None, output_file=None, zones=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of zone IDs. 
        """
        if keyword and not isinstance(keyword, str):
            raise TypeError("Expected argument 'keyword' to be a str")
        __self__.keyword = keyword
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of zone names. 
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        __self__.zones = zones
        """
        A list of zones. Each element contains the following attributes:
        """
class AwaitableGetZonesResult(GetZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZonesResult(
            id=self.id,
            ids=self.ids,
            keyword=self.keyword,
            names=self.names,
            output_file=self.output_file,
            zones=self.zones)

def get_zones(ids=None,keyword=None,output_file=None,opts=None):
    """
    This data source lists a number of Private Zones resource information owned by an Alibaba Cloud account.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    pvtz_zones_ds = alicloud.pvtz.get_zones(keyword=alicloud_pvtz_zone["basic"]["zone_name"])
    pulumi.export("firstZoneId", pvtz_zones_ds.zones[0]["id"])
    ```



    :param list ids: A list of zone IDs. 
    :param str keyword: keyword for zone name.
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['keyword'] = keyword
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:pvtz/getZones:getZones', __args__, opts=opts).value

    return AwaitableGetZonesResult(
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        keyword=__ret__.get('keyword'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        zones=__ret__.get('zones'))
