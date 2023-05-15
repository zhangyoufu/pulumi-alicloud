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
    'GetPrometheisResult',
    'AwaitableGetPrometheisResult',
    'get_prometheis',
    'get_prometheis_output',
]

@pulumi.output_type
class GetPrometheisResult:
    """
    A collection of values returned by getPrometheis.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, names=None, output_file=None, prometheis=None, resource_group_id=None, tags=None):
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
        if prometheis and not isinstance(prometheis, list):
            raise TypeError("Expected argument 'prometheis' to be a list")
        pulumi.set(__self__, "prometheis", prometheis)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

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
        A list of Prometheus names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def prometheis(self) -> Sequence['outputs.GetPrometheisPrometheiResult']:
        """
        A list of Prometheis. Each element contains the following attributes:
        """
        return pulumi.get(self, "prometheis")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        """
        The tag of the Prometheus.
        """
        return pulumi.get(self, "tags")


class AwaitableGetPrometheisResult(GetPrometheisResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrometheisResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            prometheis=self.prometheis,
            resource_group_id=self.resource_group_id,
            tags=self.tags)


def get_prometheis(ids: Optional[Sequence[str]] = None,
                   name_regex: Optional[str] = None,
                   output_file: Optional[str] = None,
                   resource_group_id: Optional[str] = None,
                   tags: Optional[Mapping[str, Any]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrometheisResult:
    """
    This data source provides the Arms Prometheis of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.203.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.arms.get_prometheis(ids=["example_id"])
    pulumi.export("armsPrometheisId1", ids.prometheis[0].id)
    name_regex = alicloud.arms.get_prometheis(name_regex="tf-example")
    pulumi.export("armsPrometheisId2", name_regex.prometheis[0].id)
    ```


    :param Sequence[str] ids: A list of Prometheus IDs.
    :param str name_regex: A regex string to filter results by Prometheus name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The ID of the resource group.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['resourceGroupId'] = resource_group_id
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:arms/getPrometheis:getPrometheis', __args__, opts=opts, typ=GetPrometheisResult).value

    return AwaitableGetPrometheisResult(
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        prometheis=__ret__.prometheis,
        resource_group_id=__ret__.resource_group_id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_prometheis)
def get_prometheis_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                          output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                          tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrometheisResult]:
    """
    This data source provides the Arms Prometheis of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.203.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.arms.get_prometheis(ids=["example_id"])
    pulumi.export("armsPrometheisId1", ids.prometheis[0].id)
    name_regex = alicloud.arms.get_prometheis(name_regex="tf-example")
    pulumi.export("armsPrometheisId2", name_regex.prometheis[0].id)
    ```


    :param Sequence[str] ids: A list of Prometheus IDs.
    :param str name_regex: A regex string to filter results by Prometheus name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The ID of the resource group.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    """
    ...
