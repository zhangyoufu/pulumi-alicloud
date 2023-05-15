# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetInstancesResult',
    'AwaitableGetInstancesResult',
    'get_instances',
    'get_instances_output',
]

@pulumi.output_type
class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, enable_details=None, id=None, ids=None, instances=None, name_regex=None, names=None, output_file=None, query_str=None, status=None, support_engine=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
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
        if query_str and not isinstance(query_str, str):
            raise TypeError("Expected argument 'query_str' to be a str")
        pulumi.set(__self__, "query_str", query_str)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if support_engine and not isinstance(support_engine, int):
            raise TypeError("Expected argument 'support_engine' to be a int")
        pulumi.set(__self__, "support_engine", support_engine)

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetInstancesInstanceResult']:
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="queryStr")
    def query_str(self) -> Optional[str]:
        return pulumi.get(self, "query_str")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="supportEngine")
    def support_engine(self) -> Optional[int]:
        return pulumi.get(self, "support_engine")


class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            instances=self.instances,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            query_str=self.query_str,
            status=self.status,
            support_engine=self.support_engine)


def get_instances(enable_details: Optional[bool] = None,
                  ids: Optional[Sequence[str]] = None,
                  name_regex: Optional[str] = None,
                  output_file: Optional[str] = None,
                  query_str: Optional[str] = None,
                  status: Optional[str] = None,
                  support_engine: Optional[int] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancesResult:
    """
    This data source provides the Lindorm Instances of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.132.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.lindorm.get_instances()
    pulumi.export("lindormInstanceId1", ids.instances[0].id)
    name_regex = alicloud.lindorm.get_instances(name_regex="^my-Instance")
    pulumi.export("lindormInstanceId2", name_regex.instances[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Instance IDs.
    :param str name_regex: A regex string to filter results by Instance name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str query_str: The query str, which can use `instance_name` keyword for fuzzy search.
    :param str status: The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
    :param int support_engine: The support engine. Valid values: `1` to `7`.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['queryStr'] = query_str
    __args__['status'] = status
    __args__['supportEngine'] = support_engine
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:lindorm/getInstances:getInstances', __args__, opts=opts, typ=GetInstancesResult).value

    return AwaitableGetInstancesResult(
        enable_details=__ret__.enable_details,
        id=__ret__.id,
        ids=__ret__.ids,
        instances=__ret__.instances,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        query_str=__ret__.query_str,
        status=__ret__.status,
        support_engine=__ret__.support_engine)


@_utilities.lift_output_func(get_instances)
def get_instances_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         query_str: Optional[pulumi.Input[Optional[str]]] = None,
                         status: Optional[pulumi.Input[Optional[str]]] = None,
                         support_engine: Optional[pulumi.Input[Optional[int]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstancesResult]:
    """
    This data source provides the Lindorm Instances of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.132.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.lindorm.get_instances()
    pulumi.export("lindormInstanceId1", ids.instances[0].id)
    name_regex = alicloud.lindorm.get_instances(name_regex="^my-Instance")
    pulumi.export("lindormInstanceId2", name_regex.instances[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Instance IDs.
    :param str name_regex: A regex string to filter results by Instance name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str query_str: The query str, which can use `instance_name` keyword for fuzzy search.
    :param str status: The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
    :param int support_engine: The support engine. Valid values: `1` to `7`.
    """
    ...
