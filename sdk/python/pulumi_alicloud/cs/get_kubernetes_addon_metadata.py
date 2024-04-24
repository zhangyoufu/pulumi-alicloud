# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetKubernetesAddonMetadataResult',
    'AwaitableGetKubernetesAddonMetadataResult',
    'get_kubernetes_addon_metadata',
    'get_kubernetes_addon_metadata_output',
]

@pulumi.output_type
class GetKubernetesAddonMetadataResult:
    """
    A collection of values returned by getKubernetesAddonMetadata.
    """
    def __init__(__self__, cluster_id=None, config_schema=None, id=None, name=None, version=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if config_schema and not isinstance(config_schema, str):
            raise TypeError("Expected argument 'config_schema' to be a str")
        pulumi.set(__self__, "config_schema", config_schema)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="configSchema")
    def config_schema(self) -> str:
        """
        The addon configuration that can be customized. The returned format is the standard json schema. If return empty, it means that the addon does not support custom configuration yet.
        """
        return pulumi.get(self, "config_schema")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> str:
        return pulumi.get(self, "version")


class AwaitableGetKubernetesAddonMetadataResult(GetKubernetesAddonMetadataResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubernetesAddonMetadataResult(
            cluster_id=self.cluster_id,
            config_schema=self.config_schema,
            id=self.id,
            name=self.name,
            version=self.version)


def get_kubernetes_addon_metadata(cluster_id: Optional[str] = None,
                                  name: Optional[str] = None,
                                  version: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubernetesAddonMetadataResult:
    """
    This data source provides metadata of kubernetes cluster addons.

    > **NOTE:** Available in 1.166.0+.


    :param str cluster_id: The id of kubernetes cluster.
    :param str name: The name of the cluster addon. You can get a list of available addons that the cluster can install by using data source `cs_get_kubernetes_addons`.
    :param str version: The version of the cluster addon.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    __args__['version'] = version
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cs/getKubernetesAddonMetadata:getKubernetesAddonMetadata', __args__, opts=opts, typ=GetKubernetesAddonMetadataResult).value

    return AwaitableGetKubernetesAddonMetadataResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        config_schema=pulumi.get(__ret__, 'config_schema'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        version=pulumi.get(__ret__, 'version'))


@_utilities.lift_output_func(get_kubernetes_addon_metadata)
def get_kubernetes_addon_metadata_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                         name: Optional[pulumi.Input[str]] = None,
                                         version: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKubernetesAddonMetadataResult]:
    """
    This data source provides metadata of kubernetes cluster addons.

    > **NOTE:** Available in 1.166.0+.


    :param str cluster_id: The id of kubernetes cluster.
    :param str name: The name of the cluster addon. You can get a list of available addons that the cluster can install by using data source `cs_get_kubernetes_addons`.
    :param str version: The version of the cluster addon.
    """
    ...
