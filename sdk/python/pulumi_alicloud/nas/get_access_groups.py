# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetAccessGroupsResult:
    """
    A collection of values returned by getAccessGroups.
    """
    def __init__(__self__, description=None, groups=None, id=None, ids=None, name_regex=None, names=None, output_file=None, type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Destription of the AccessGroup.
        """
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        __self__.groups = groups
        """
        A list of AccessGroups. Each element contains the following attributes:
        """
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
        A list of AccessGroup IDs, the value is set to `names` .
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of AccessGroup names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        AccessGroupType of the AccessGroup.
        """
class AwaitableGetAccessGroupsResult(GetAccessGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessGroupsResult(
            description=self.description,
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            type=self.type)

def get_access_groups(description=None,name_regex=None,output_file=None,type=None,opts=None):
    """
    This data source provides user-available access groups. Use when you can create mount points

    > NOTE: Available in 1.35.0+




    :param str description: Filter results by a specific Description.
    :param str name_regex: A regex string to filter AccessGroups by name. 
    :param str type: Filter results by a specific AccessGroupType.
    """
    __args__ = dict()


    __args__['description'] = description
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:nas/getAccessGroups:getAccessGroups', __args__, opts=opts).value

    return AwaitableGetAccessGroupsResult(
        description=__ret__.get('description'),
        groups=__ret__.get('groups'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        type=__ret__.get('type'))
