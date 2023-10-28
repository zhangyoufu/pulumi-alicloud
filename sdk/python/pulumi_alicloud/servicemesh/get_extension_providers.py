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
    'GetExtensionProvidersResult',
    'AwaitableGetExtensionProvidersResult',
    'get_extension_providers',
    'get_extension_providers_output',
]

@pulumi.output_type
class GetExtensionProvidersResult:
    """
    A collection of values returned by getExtensionProviders.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, names=None, output_file=None, providers=None, service_mesh_id=None, type=None):
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
        if providers and not isinstance(providers, list):
            raise TypeError("Expected argument 'providers' to be a list")
        pulumi.set(__self__, "providers", providers)
        if service_mesh_id and not isinstance(service_mesh_id, str):
            raise TypeError("Expected argument 'service_mesh_id' to be a str")
        pulumi.set(__self__, "service_mesh_id", service_mesh_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

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
        A list of Extension Provider names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def providers(self) -> Sequence['outputs.GetExtensionProvidersProviderResult']:
        """
        A list of Service Mesh Extension Providers. Each element contains the following attributes:
        """
        return pulumi.get(self, "providers")

    @property
    @pulumi.getter(name="serviceMeshId")
    def service_mesh_id(self) -> str:
        """
        The ID of the Service Mesh.
        """
        return pulumi.get(self, "service_mesh_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the Service Mesh Extension Provider.
        """
        return pulumi.get(self, "type")


class AwaitableGetExtensionProvidersResult(GetExtensionProvidersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExtensionProvidersResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            providers=self.providers,
            service_mesh_id=self.service_mesh_id,
            type=self.type)


def get_extension_providers(ids: Optional[Sequence[str]] = None,
                            name_regex: Optional[str] = None,
                            output_file: Optional[str] = None,
                            service_mesh_id: Optional[str] = None,
                            type: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExtensionProvidersResult:
    """
    This data source provides the Service Mesh Extension Providers of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.191.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.servicemesh.get_extension_providers(ids=["example_id"],
        service_mesh_id="example_service_mesh_id",
        type="httpextauth")
    pulumi.export("serviceMeshExtensionProvidersId1", ids.providers[0].id)
    name_regex = alicloud.servicemesh.get_extension_providers(name_regex="^my-ServiceMeshExtensionProvider",
        service_mesh_id="example_service_mesh_id",
        type="httpextauth")
    pulumi.export("serviceMeshExtensionProvidersId2", name_regex.providers[0].id)
    ```


    :param Sequence[str] ids: A list of Service Mesh Extension Provider IDs.
    :param str name_regex: A regex string to filter results by Service Mesh Extension Provider name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str service_mesh_id: The ID of the Service Mesh.
    :param str type: The type of the Service Mesh Extension Provider. Valid values: `httpextauth`, `grpcextauth`.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['serviceMeshId'] = service_mesh_id
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:servicemesh/getExtensionProviders:getExtensionProviders', __args__, opts=opts, typ=GetExtensionProvidersResult).value

    return AwaitableGetExtensionProvidersResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        providers=pulumi.get(__ret__, 'providers'),
        service_mesh_id=pulumi.get(__ret__, 'service_mesh_id'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_extension_providers)
def get_extension_providers_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                   name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                   output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   service_mesh_id: Optional[pulumi.Input[str]] = None,
                                   type: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetExtensionProvidersResult]:
    """
    This data source provides the Service Mesh Extension Providers of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.191.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.servicemesh.get_extension_providers(ids=["example_id"],
        service_mesh_id="example_service_mesh_id",
        type="httpextauth")
    pulumi.export("serviceMeshExtensionProvidersId1", ids.providers[0].id)
    name_regex = alicloud.servicemesh.get_extension_providers(name_regex="^my-ServiceMeshExtensionProvider",
        service_mesh_id="example_service_mesh_id",
        type="httpextauth")
    pulumi.export("serviceMeshExtensionProvidersId2", name_regex.providers[0].id)
    ```


    :param Sequence[str] ids: A list of Service Mesh Extension Provider IDs.
    :param str name_regex: A regex string to filter results by Service Mesh Extension Provider name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str service_mesh_id: The ID of the Service Mesh.
    :param str type: The type of the Service Mesh Extension Provider. Valid values: `httpextauth`, `grpcextauth`.
    """
    ...
