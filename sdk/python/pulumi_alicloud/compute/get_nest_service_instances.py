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
from ._inputs import *

__all__ = [
    'GetNestServiceInstancesResult',
    'AwaitableGetNestServiceInstancesResult',
    'get_nest_service_instances',
    'get_nest_service_instances_output',
]

@pulumi.output_type
class GetNestServiceInstancesResult:
    """
    A collection of values returned by getNestServiceInstances.
    """
    def __init__(__self__, filters=None, id=None, ids=None, name_regex=None, names=None, output_file=None, service_instances=None, status=None, tags=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if service_instances and not isinstance(service_instances, list):
            raise TypeError("Expected argument 'service_instances' to be a list")
        pulumi.set(__self__, "service_instances", service_instances)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetNestServiceInstancesFilterResult']]:
        return pulumi.get(self, "filters")

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
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of Service Instance names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="serviceInstances")
    def service_instances(self) -> Sequence['outputs.GetNestServiceInstancesServiceInstanceResult']:
        """
        A list of Service Instances. Each element contains the following attributes:
        """
        return pulumi.get(self, "service_instances")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the Service Instance.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        """
        The tag of the Service Instance.
        """
        return pulumi.get(self, "tags")


class AwaitableGetNestServiceInstancesResult(GetNestServiceInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNestServiceInstancesResult(
            filters=self.filters,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            service_instances=self.service_instances,
            status=self.status,
            tags=self.tags)


def get_nest_service_instances(filters: Optional[Sequence[Union['GetNestServiceInstancesFilterArgs', 'GetNestServiceInstancesFilterArgsDict']]] = None,
                               ids: Optional[Sequence[str]] = None,
                               name_regex: Optional[str] = None,
                               output_file: Optional[str] = None,
                               status: Optional[str] = None,
                               tags: Optional[Mapping[str, Any]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNestServiceInstancesResult:
    """
    This data source provides the Compute Nest Service Instances of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.205.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.compute.get_nest_service_instances(ids=["example_id"])
    pulumi.export("armsPrometheisId1", ids.service_instances[0].id)
    name_regex = alicloud.compute.get_nest_service_instances(name_regex="tf-example")
    pulumi.export("armsPrometheisId2", name_regex.service_instances[0].id)
    ```


    :param Sequence[Union['GetNestServiceInstancesFilterArgs', 'GetNestServiceInstancesFilterArgsDict']] filters: The conditions that are used to filter. See the following `Block filter`.
    :param Sequence[str] ids: A list of Service Instance IDs.
    :param str name_regex: A regex string to filter results by Service Instance name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the Service Instance. Valid Values: `Created`, `Deploying`, `DeployedFailed`, `Deployed`, `Upgrading`, `Deleting`, `Deleted`, `DeletedFailed`.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:compute/getNestServiceInstances:getNestServiceInstances', __args__, opts=opts, typ=GetNestServiceInstancesResult).value

    return AwaitableGetNestServiceInstancesResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        service_instances=pulumi.get(__ret__, 'service_instances'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_nest_service_instances)
def get_nest_service_instances_output(filters: Optional[pulumi.Input[Optional[Sequence[Union['GetNestServiceInstancesFilterArgs', 'GetNestServiceInstancesFilterArgsDict']]]]] = None,
                                      ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                      name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                      status: Optional[pulumi.Input[Optional[str]]] = None,
                                      tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNestServiceInstancesResult]:
    """
    This data source provides the Compute Nest Service Instances of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.205.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.compute.get_nest_service_instances(ids=["example_id"])
    pulumi.export("armsPrometheisId1", ids.service_instances[0].id)
    name_regex = alicloud.compute.get_nest_service_instances(name_regex="tf-example")
    pulumi.export("armsPrometheisId2", name_regex.service_instances[0].id)
    ```


    :param Sequence[Union['GetNestServiceInstancesFilterArgs', 'GetNestServiceInstancesFilterArgsDict']] filters: The conditions that are used to filter. See the following `Block filter`.
    :param Sequence[str] ids: A list of Service Instance IDs.
    :param str name_regex: A regex string to filter results by Service Instance name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the Service Instance. Valid Values: `Created`, `Deploying`, `DeployedFailed`, `Deployed`, `Upgrading`, `Deleting`, `Deleted`, `DeletedFailed`.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    """
    ...
