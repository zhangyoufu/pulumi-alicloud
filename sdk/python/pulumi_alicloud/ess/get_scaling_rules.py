# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetScalingRulesResult:
    """
    A collection of values returned by getScalingRules.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, names=None, output_file=None, rules=None, scaling_group_id=None, type=None):
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
        A list of scaling rule ids.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of scaling rule names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        __self__.rules = rules
        """
        A list of scaling rules. Each element contains the following attributes:
        """
        if scaling_group_id and not isinstance(scaling_group_id, str):
            raise TypeError("Expected argument 'scaling_group_id' to be a str")
        __self__.scaling_group_id = scaling_group_id
        """
        ID of the scaling group.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Type of the scaling rule.
        """
class AwaitableGetScalingRulesResult(GetScalingRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScalingRulesResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            rules=self.rules,
            scaling_group_id=self.scaling_group_id,
            type=self.type)

def get_scaling_rules(ids=None,name_regex=None,output_file=None,scaling_group_id=None,type=None,opts=None):
    """
    This data source provides available scaling rule resources. 




    :param list ids: A list of scaling rule IDs.
    :param str name_regex: A regex string to filter resulting scaling rules by name.
    :param str scaling_group_id: Scaling group id the scaling rules belong to.
    :param str type: Type of scaling rule.
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['scalingGroupId'] = scaling_group_id
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ess/getScalingRules:getScalingRules', __args__, opts=opts).value

    return AwaitableGetScalingRulesResult(
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        rules=__ret__.get('rules'),
        scaling_group_id=__ret__.get('scalingGroupId'),
        type=__ret__.get('type'))
