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
    'GetRegistryEnterpriseNamespacesResult',
    'AwaitableGetRegistryEnterpriseNamespacesResult',
    'get_registry_enterprise_namespaces',
    'get_registry_enterprise_namespaces_output',
]

@pulumi.output_type
class GetRegistryEnterpriseNamespacesResult:
    """
    A collection of values returned by getRegistryEnterpriseNamespaces.
    """
    def __init__(__self__, id=None, ids=None, instance_id=None, name_regex=None, names=None, namespaces=None, output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
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
        A list of matched Container Registry Enterprise Edition namespaces. Each item formats as `<instance_id>:<namespace_name>`. Before 1.161.0, its element is a namespace uuid.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        ID of Container Registry Enterprise Edition instance.
        """
        return pulumi.get(self, "instance_id")

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
    def namespaces(self) -> Sequence['outputs.GetRegistryEnterpriseNamespacesNamespaceResult']:
        """
        A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:
        """
        return pulumi.get(self, "namespaces")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetRegistryEnterpriseNamespacesResult(GetRegistryEnterpriseNamespacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistryEnterpriseNamespacesResult(
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            names=self.names,
            namespaces=self.namespaces,
            output_file=self.output_file)


def get_registry_enterprise_namespaces(ids: Optional[Sequence[str]] = None,
                                       instance_id: Optional[str] = None,
                                       name_regex: Optional[str] = None,
                                       output_file: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegistryEnterpriseNamespacesResult:
    """
    This data source provides a list Container Registry Enterprise Edition namespaces on Alibaba Cloud.

    > **NOTE:** Available in v1.86.0+

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    # Declare the data source
    my_namespaces = alicloud.cs.get_registry_enterprise_namespaces(instance_id="cri-xxx",
        name_regex="my-namespace",
        output_file="my-namespace-json")
    pulumi.export("output", my_namespaces.namespaces)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of ids to filter results by namespace id. Each item formats as `<instance_id>:<namespace_name>`.
    :param str instance_id: ID of Container Registry Enterprise Edition instance.
    :param str name_regex: A regex string to filter results by namespace name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cs/getRegistryEnterpriseNamespaces:getRegistryEnterpriseNamespaces', __args__, opts=opts, typ=GetRegistryEnterpriseNamespacesResult).value

    return AwaitableGetRegistryEnterpriseNamespacesResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        namespaces=pulumi.get(__ret__, 'namespaces'),
        output_file=pulumi.get(__ret__, 'output_file'))


@_utilities.lift_output_func(get_registry_enterprise_namespaces)
def get_registry_enterprise_namespaces_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                              instance_id: Optional[pulumi.Input[str]] = None,
                                              name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegistryEnterpriseNamespacesResult]:
    """
    This data source provides a list Container Registry Enterprise Edition namespaces on Alibaba Cloud.

    > **NOTE:** Available in v1.86.0+

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    # Declare the data source
    my_namespaces = alicloud.cs.get_registry_enterprise_namespaces(instance_id="cri-xxx",
        name_regex="my-namespace",
        output_file="my-namespace-json")
    pulumi.export("output", my_namespaces.namespaces)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of ids to filter results by namespace id. Each item formats as `<instance_id>:<namespace_name>`.
    :param str instance_id: ID of Container Registry Enterprise Edition instance.
    :param str name_regex: A regex string to filter results by namespace name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
