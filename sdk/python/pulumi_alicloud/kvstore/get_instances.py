# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetInstancesResult',
    'AwaitableGetInstancesResult',
    'get_instances',
]

@pulumi.output_type
class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, id=None, ids=None, instance_class=None, instance_type=None, instances=None, name_regex=None, names=None, output_file=None, status=None, tags=None, vpc_id=None, vswitch_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_class and not isinstance(instance_class, str):
            raise TypeError("Expected argument 'instance_class' to be a str")
        pulumi.set(__self__, "instance_class", instance_class)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_id and not isinstance(vswitch_id, str):
            raise TypeError("Expected argument 'vswitch_id' to be a str")
        pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> List[str]:
        """
        A list of RKV instance IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> Optional[str]:
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        """
        (Optional) Database type. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
        * `instance_class`- (Optional) Type of the applied ApsaraDB for Redis instance.
        For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def instances(self) -> List['outputs.GetInstancesInstanceResult']:
        """
        A list of RKV instances. Its every element contains the following attributes:
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> List[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Status of the instance.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        VPC ID the instance belongs to.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[str]:
        """
        VSwitch ID the instance belongs to.
        """
        return pulumi.get(self, "vswitch_id")


class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            id=self.id,
            ids=self.ids,
            instance_class=self.instance_class,
            instance_type=self.instance_type,
            instances=self.instances,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            status=self.status,
            tags=self.tags,
            vpc_id=self.vpc_id,
            vswitch_id=self.vswitch_id)


def get_instances(ids: Optional[List[str]] = None,
                  instance_class: Optional[str] = None,
                  instance_type: Optional[str] = None,
                  name_regex: Optional[str] = None,
                  output_file: Optional[str] = None,
                  status: Optional[str] = None,
                  tags: Optional[Mapping[str, Any]] = None,
                  vpc_id: Optional[str] = None,
                  vswitch_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancesResult:
    """
    The `kvstore.getInstances` data source provides a collection of kvstore instances available in Alicloud account.
    Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.


    :param List[str] ids: A list of RKV instance IDs.
    :param str instance_class: Type of the applied ApsaraDB for Redis instance.
           For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
    :param str instance_type: Database type. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
    :param str name_regex: A regex string to apply to the instance name.
    :param str status: Status of the instance.
    :param Mapping[str, Any] tags: Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
    :param str vpc_id: Used to retrieve instances belong to specified VPC.
    :param str vswitch_id: Used to retrieve instances belong to specified `vswitch` resources.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceClass'] = instance_class
    __args__['instanceType'] = instance_type
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    __args__['vswitchId'] = vswitch_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:kvstore/getInstances:getInstances', __args__, opts=opts, typ=GetInstancesResult).value

    return AwaitableGetInstancesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        instance_class=__ret__.instance_class,
        instance_type=__ret__.instance_type,
        instances=__ret__.instances,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        status=__ret__.status,
        tags=__ret__.tags,
        vpc_id=__ret__.vpc_id,
        vswitch_id=__ret__.vswitch_id)
