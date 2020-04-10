# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetLifecycleHooksResult:
    """
    A collection of values returned by getLifecycleHooks.
    """
    def __init__(__self__, hooks=None, id=None, ids=None, name_regex=None, names=None, output_file=None, scaling_group_id=None):
        if hooks and not isinstance(hooks, list):
            raise TypeError("Expected argument 'hooks' to be a list")
        __self__.hooks = hooks
        """
        A list of lifecycle hooks. Each element contains the following attributes:
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
        A list of lifecycle hook ids.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of lifecycle hook names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if scaling_group_id and not isinstance(scaling_group_id, str):
            raise TypeError("Expected argument 'scaling_group_id' to be a str")
        __self__.scaling_group_id = scaling_group_id
        """
        ID of the scaling group.
        """
class AwaitableGetLifecycleHooksResult(GetLifecycleHooksResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLifecycleHooksResult(
            hooks=self.hooks,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            scaling_group_id=self.scaling_group_id)

def get_lifecycle_hooks(ids=None,name_regex=None,output_file=None,scaling_group_id=None,opts=None):
    """
    This data source provides available lifecycle hook resources. 

    > **NOTE:** Available in 1.72.0+




    :param list ids: A list of lifecycle hook IDs.
    :param str name_regex: A regex string to filter resulting lifecycle hook by name.
    :param str scaling_group_id: Scaling group id the lifecycle hooks belong to.
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['scalingGroupId'] = scaling_group_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ess/getLifecycleHooks:getLifecycleHooks', __args__, opts=opts).value

    return AwaitableGetLifecycleHooksResult(
        hooks=__ret__.get('hooks'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        scaling_group_id=__ret__.get('scalingGroupId'))
