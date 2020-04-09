# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetServerGroupsResult:
    """
    A collection of values returned by getServerGroups.
    """
    def __init__(__self__, id=None, ids=None, load_balancer_id=None, name_regex=None, names=None, output_file=None, slb_server_groups=None):
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
        A list of SLB VServer groups IDs.
        """
        if load_balancer_id and not isinstance(load_balancer_id, str):
            raise TypeError("Expected argument 'load_balancer_id' to be a str")
        __self__.load_balancer_id = load_balancer_id
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of SLB VServer groups names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if slb_server_groups and not isinstance(slb_server_groups, list):
            raise TypeError("Expected argument 'slb_server_groups' to be a list")
        __self__.slb_server_groups = slb_server_groups
        """
        A list of SLB VServer groups. Each element contains the following attributes:
        """
class AwaitableGetServerGroupsResult(GetServerGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerGroupsResult(
            id=self.id,
            ids=self.ids,
            load_balancer_id=self.load_balancer_id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            slb_server_groups=self.slb_server_groups)

def get_server_groups(ids=None,load_balancer_id=None,name_regex=None,output_file=None,opts=None):
    """
    This data source provides the VServer groups related to a server load balancer.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_server_groups.html.markdown.


    :param list ids: A list of VServer group IDs to filter results.
    :param str load_balancer_id: ID of the SLB.
    :param str name_regex: A regex string to filter results by VServer group name.
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['loadBalancerId'] = load_balancer_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:slb/getServerGroups:getServerGroups', __args__, opts=opts).value

    return AwaitableGetServerGroupsResult(
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        load_balancer_id=__ret__.get('loadBalancerId'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        slb_server_groups=__ret__.get('slbServerGroups'))
