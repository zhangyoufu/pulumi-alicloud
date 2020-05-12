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
    def __init__(__self__, id=None, ids=None, multi=None, output_file=None, zones=None):
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
        if multi and not isinstance(multi, bool):
            raise TypeError("Expected argument 'multi' to be a bool")
        __self__.multi = multi
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        __self__.zones = zones
        """
        A list of availability zones. Each element contains the following attributes:
        """
class AwaitableGetZonesResult(GetZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZonesResult(
            id=self.id,
            ids=self.ids,
            multi=self.multi,
            output_file=self.output_file,
            zones=self.zones)

def get_zones(multi=None,output_file=None,opts=None):
    """
    This data source provides availability zones for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.

    > **NOTE:** Available in v1.73.0+.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    zones_ids = alicloud.hbase.get_zones()
    # Create an HBase instance with the first matched zone
    hbase = alicloud.hbase.Instance("hbase", zone_id=zones_ids.zones[0]["id"])
    # Other properties...
    ```



    :param bool multi: Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch HBase instances.
    """
    __args__ = dict()


    __args__['multi'] = multi
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:hbase/getZones:getZones', __args__, opts=opts).value

    return AwaitableGetZonesResult(
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        multi=__ret__.get('multi'),
        output_file=__ret__.get('outputFile'),
        zones=__ret__.get('zones'))
