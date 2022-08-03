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
    'GetNamespacesResult',
    'AwaitableGetNamespacesResult',
    'get_namespaces',
    'get_namespaces_output',
]

@pulumi.output_type
class GetNamespacesResult:
    """
    A collection of values returned by getNamespaces.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, names=None, namespaces=None, output_file=None):
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
        if namespaces and not isinstance(namespaces, list):
            raise TypeError("Expected argument 'namespaces' to be a list")
        pulumi.set(__self__, "namespaces", namespaces)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)

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
        """
        A list of matched Container Registry namespaces. Its element is a namespace name.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of namespace names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def namespaces(self) -> Sequence['outputs.GetNamespacesNamespaceResult']:
        """
        A list of matched Container Registry namespaces. Each element contains the following attributes:
        """
        return pulumi.get(self, "namespaces")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetNamespacesResult(GetNamespacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNamespacesResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            namespaces=self.namespaces,
            output_file=self.output_file)


def get_namespaces(name_regex: Optional[str] = None,
                   output_file: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNamespacesResult:
    """
    This data source provides a list Container Registry namespaces on Alibaba Cloud.

    > **NOTE:** Available in v1.35.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    my_namespaces = alicloud.cr.get_namespaces(name_regex="my-namespace",
        output_file="my-namespace-json")
    pulumi.export("output", my_namespaces.namespaces)
    ```


    :param str name_regex: A regex string to filter results by namespace name.
    """
    __args__ = dict()
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cr/getNamespaces:getNamespaces', __args__, opts=opts, typ=GetNamespacesResult).value

    return AwaitableGetNamespacesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        namespaces=__ret__.namespaces,
        output_file=__ret__.output_file)


@_utilities.lift_output_func(get_namespaces)
def get_namespaces_output(name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                          output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNamespacesResult]:
    """
    This data source provides a list Container Registry namespaces on Alibaba Cloud.

    > **NOTE:** Available in v1.35.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    my_namespaces = alicloud.cr.get_namespaces(name_regex="my-namespace",
        output_file="my-namespace-json")
    pulumi.export("output", my_namespaces.namespaces)
    ```


    :param str name_regex: A regex string to filter results by namespace name.
    """
    ...
