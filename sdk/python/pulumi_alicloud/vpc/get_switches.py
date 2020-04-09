# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSwitchesResult:
    """
    A collection of values returned by getSwitches.
    """
    def __init__(__self__, cidr_block=None, id=None, ids=None, is_default=None, name_regex=None, names=None, output_file=None, resource_group_id=None, tags=None, vpc_id=None, vswitches=None, zone_id=None):
        if cidr_block and not isinstance(cidr_block, str):
            raise TypeError("Expected argument 'cidr_block' to be a str")
        __self__.cidr_block = cidr_block
        """
        CIDR block of the VSwitch.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of VSwitch IDs.
        """
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        __self__.is_default = is_default
        """
        Whether the VSwitch is the default one in the region.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of VSwitch names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        __self__.resource_group_id = resource_group_id
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        __self__.vpc_id = vpc_id
        """
        ID of the VPC that owns the VSwitch.
        """
        if vswitches and not isinstance(vswitches, list):
            raise TypeError("Expected argument 'vswitches' to be a list")
        __self__.vswitches = vswitches
        """
        A list of VSwitches. Each element contains the following attributes:
        """
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        __self__.zone_id = zone_id
        """
        ID of the availability zone where the VSwitch is located.
        """
class AwaitableGetSwitchesResult(GetSwitchesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSwitchesResult(
            cidr_block=self.cidr_block,
            id=self.id,
            ids=self.ids,
            is_default=self.is_default,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            resource_group_id=self.resource_group_id,
            tags=self.tags,
            vpc_id=self.vpc_id,
            vswitches=self.vswitches,
            zone_id=self.zone_id)

def get_switches(cidr_block=None,ids=None,is_default=None,name_regex=None,output_file=None,resource_group_id=None,tags=None,vpc_id=None,zone_id=None,opts=None):
    """
    This data source provides a list of VSwitches owned by an Alibaba Cloud account.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/vswitches.html.markdown.


    :param str cidr_block: Filter results by a specific CIDR block. For example: "172.16.0.0/12".
    :param list ids: A list of VSwitch IDs.
    :param bool is_default: Indicate whether the VSwitch is created by the system.
    :param str name_regex: A regex string to filter results by name.
    :param str resource_group_id: The Id of resource group which VSWitch belongs.
    :param dict tags: A mapping of tags to assign to the resource.
    :param str vpc_id: ID of the VPC that owns the VSwitch.
    :param str zone_id: The availability zone of the VSwitch.
    """
    __args__ = dict()


    __args__['cidrBlock'] = cidr_block
    __args__['ids'] = ids
    __args__['isDefault'] = is_default
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['resourceGroupId'] = resource_group_id
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:vpc/getSwitches:getSwitches', __args__, opts=opts).value

    return AwaitableGetSwitchesResult(
        cidr_block=__ret__.get('cidrBlock'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        is_default=__ret__.get('isDefault'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        resource_group_id=__ret__.get('resourceGroupId'),
        tags=__ret__.get('tags'),
        vpc_id=__ret__.get('vpcId'),
        vswitches=__ret__.get('vswitches'),
        zone_id=__ret__.get('zoneId'))
