# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, group_name=None, id=None, ids=None, name_regex=None, names=None, output_file=None, policy_name=None, policy_type=None, users=None):
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        __self__.group_name = group_name
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
        A list of ram user IDs. 
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of ram user names. 
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if policy_name and not isinstance(policy_name, str):
            raise TypeError("Expected argument 'policy_name' to be a str")
        __self__.policy_name = policy_name
        if policy_type and not isinstance(policy_type, str):
            raise TypeError("Expected argument 'policy_type' to be a str")
        __self__.policy_type = policy_type
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        __self__.users = users
        """
        A list of users. Each element contains the following attributes:
        """
class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            group_name=self.group_name,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            policy_name=self.policy_name,
            policy_type=self.policy_type,
            users=self.users)

def get_users(group_name=None,ids=None,name_regex=None,output_file=None,policy_name=None,policy_type=None,opts=None):
    """
    This data source provides a list of RAM users in an Alibaba Cloud account according to the specified filters.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ram_users.html.markdown.


    :param str group_name: Filter results by a specific group name. Returned users are in the specified group. 
    :param list ids: A list of ram user IDs. 
    :param str name_regex: A regex string to filter resulting users by their names.
           * `ids` (Optional, Available 1.53.0+) - A list of ram user IDs.
    :param str policy_name: Filter results by a specific policy name. If you set this parameter without setting `policy_type`, the later will be automatically set to `System`. Returned users are attached to the specified policy.
    :param str policy_type: Filter results by a specific policy type. Valid values are `Custom` and `System`. If you set this parameter, you must set `policy_name` as well.
    """
    __args__ = dict()


    __args__['groupName'] = group_name
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['policyName'] = policy_name
    __args__['policyType'] = policy_type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ram/getUsers:getUsers', __args__, opts=opts).value

    return AwaitableGetUsersResult(
        group_name=__ret__.get('groupName'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        policy_name=__ret__.get('policyName'),
        policy_type=__ret__.get('policyType'),
        users=__ret__.get('users'))
