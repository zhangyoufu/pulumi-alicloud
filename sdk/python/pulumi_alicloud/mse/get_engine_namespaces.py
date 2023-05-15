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
    'GetEngineNamespacesResult',
    'AwaitableGetEngineNamespacesResult',
    'get_engine_namespaces',
    'get_engine_namespaces_output',
]

@pulumi.output_type
class GetEngineNamespacesResult:
    """
    A collection of values returned by getEngineNamespaces.
    """
    def __init__(__self__, accept_language=None, cluster_id=None, id=None, ids=None, namespaces=None, output_file=None):
        if accept_language and not isinstance(accept_language, str):
            raise TypeError("Expected argument 'accept_language' to be a str")
        pulumi.set(__self__, "accept_language", accept_language)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if namespaces and not isinstance(namespaces, list):
            raise TypeError("Expected argument 'namespaces' to be a list")
        pulumi.set(__self__, "namespaces", namespaces)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> Optional[str]:
        return pulumi.get(self, "accept_language")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

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
    def namespaces(self) -> Sequence['outputs.GetEngineNamespacesNamespaceResult']:
        return pulumi.get(self, "namespaces")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetEngineNamespacesResult(GetEngineNamespacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEngineNamespacesResult(
            accept_language=self.accept_language,
            cluster_id=self.cluster_id,
            id=self.id,
            ids=self.ids,
            namespaces=self.namespaces,
            output_file=self.output_file)


def get_engine_namespaces(accept_language: Optional[str] = None,
                          cluster_id: Optional[str] = None,
                          ids: Optional[Sequence[str]] = None,
                          output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEngineNamespacesResult:
    """
    This data source provides the Mse Engine Namespaces of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.166.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.mse.get_engine_namespaces(cluster_id="example_value",
        ids=["example_value"])
    pulumi.export("mseEngineNamespaceId1", ids.namespaces[0].id)
    ```


    :param str accept_language: The language type of the returned information. Valid values: `zh`, `en`.
    :param str cluster_id: The id of the cluster.
    :param Sequence[str] ids: A list of Engine Namespace IDs. It is formatted to `<cluster_id>:<namespace_id>`.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['acceptLanguage'] = accept_language
    __args__['clusterId'] = cluster_id
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:mse/getEngineNamespaces:getEngineNamespaces', __args__, opts=opts, typ=GetEngineNamespacesResult).value

    return AwaitableGetEngineNamespacesResult(
        accept_language=__ret__.accept_language,
        cluster_id=__ret__.cluster_id,
        id=__ret__.id,
        ids=__ret__.ids,
        namespaces=__ret__.namespaces,
        output_file=__ret__.output_file)


@_utilities.lift_output_func(get_engine_namespaces)
def get_engine_namespaces_output(accept_language: Optional[pulumi.Input[Optional[str]]] = None,
                                 cluster_id: Optional[pulumi.Input[str]] = None,
                                 ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEngineNamespacesResult]:
    """
    This data source provides the Mse Engine Namespaces of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.166.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.mse.get_engine_namespaces(cluster_id="example_value",
        ids=["example_value"])
    pulumi.export("mseEngineNamespaceId1", ids.namespaces[0].id)
    ```


    :param str accept_language: The language type of the returned information. Valid values: `zh`, `en`.
    :param str cluster_id: The id of the cluster.
    :param Sequence[str] ids: A list of Engine Namespace IDs. It is formatted to `<cluster_id>:<namespace_id>`.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
