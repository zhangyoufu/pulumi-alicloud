# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetAlarmsResult:
    """
    A collection of values returned by getAlarms.
    """
    def __init__(__self__, alarms=None, id=None, ids=None, metric_type=None, name_regex=None, names=None, output_file=None, scaling_group_id=None):
        if alarms and not isinstance(alarms, list):
            raise TypeError("Expected argument 'alarms' to be a list")
        __self__.alarms = alarms
        """
        A list of alarms. Each element contains the following attributes:
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
        A list of alarm ids.
        """
        if metric_type and not isinstance(metric_type, str):
            raise TypeError("Expected argument 'metric_type' to be a str")
        __self__.metric_type = metric_type
        """
        The type for the alarm's associated metric. 
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of alarm names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if scaling_group_id and not isinstance(scaling_group_id, str):
            raise TypeError("Expected argument 'scaling_group_id' to be a str")
        __self__.scaling_group_id = scaling_group_id
        """
        The scaling group associated with this alarm.
        """
class AwaitableGetAlarmsResult(GetAlarmsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlarmsResult(
            alarms=self.alarms,
            id=self.id,
            ids=self.ids,
            metric_type=self.metric_type,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            scaling_group_id=self.scaling_group_id)

def get_alarms(ids=None,metric_type=None,name_regex=None,output_file=None,scaling_group_id=None,opts=None):
    """
    This data source provides available alarm resources. 

    > **NOTE** Available in 1.72.0+




    :param list ids: A list of alarm IDs.
    :param str metric_type: The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.
    :param str name_regex: A regex string to filter resulting alarms by name.
    :param str scaling_group_id: Scaling group id the alarms belong to.
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['metricType'] = metric_type
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['scalingGroupId'] = scaling_group_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:ess/getAlarms:getAlarms', __args__, opts=opts).value

    return AwaitableGetAlarmsResult(
        alarms=__ret__.get('alarms'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        metric_type=__ret__.get('metricType'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        scaling_group_id=__ret__.get('scalingGroupId'))
