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
    'GetEcsDedicatedHostClustersResult',
    'AwaitableGetEcsDedicatedHostClustersResult',
    'get_ecs_dedicated_host_clusters',
    'get_ecs_dedicated_host_clusters_output',
]

@pulumi.output_type
class GetEcsDedicatedHostClustersResult:
    """
    A collection of values returned by getEcsDedicatedHostClusters.
    """
    def __init__(__self__, clusters=None, dedicated_host_cluster_ids=None, dedicated_host_cluster_name=None, id=None, ids=None, name_regex=None, names=None, output_file=None, tags=None, zone_id=None):
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if dedicated_host_cluster_ids and not isinstance(dedicated_host_cluster_ids, list):
            raise TypeError("Expected argument 'dedicated_host_cluster_ids' to be a list")
        pulumi.set(__self__, "dedicated_host_cluster_ids", dedicated_host_cluster_ids)
        if dedicated_host_cluster_name and not isinstance(dedicated_host_cluster_name, str):
            raise TypeError("Expected argument 'dedicated_host_cluster_name' to be a str")
        pulumi.set(__self__, "dedicated_host_cluster_name", dedicated_host_cluster_name)
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
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.GetEcsDedicatedHostClustersClusterResult']:
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter(name="dedicatedHostClusterIds")
    def dedicated_host_cluster_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "dedicated_host_cluster_ids")

    @property
    @pulumi.getter(name="dedicatedHostClusterName")
    def dedicated_host_cluster_name(self) -> Optional[str]:
        return pulumi.get(self, "dedicated_host_cluster_name")

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
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[str]:
        return pulumi.get(self, "zone_id")


class AwaitableGetEcsDedicatedHostClustersResult(GetEcsDedicatedHostClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEcsDedicatedHostClustersResult(
            clusters=self.clusters,
            dedicated_host_cluster_ids=self.dedicated_host_cluster_ids,
            dedicated_host_cluster_name=self.dedicated_host_cluster_name,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            tags=self.tags,
            zone_id=self.zone_id)


def get_ecs_dedicated_host_clusters(dedicated_host_cluster_ids: Optional[Sequence[str]] = None,
                                    dedicated_host_cluster_name: Optional[str] = None,
                                    ids: Optional[Sequence[str]] = None,
                                    name_regex: Optional[str] = None,
                                    output_file: Optional[str] = None,
                                    tags: Optional[Mapping[str, Any]] = None,
                                    zone_id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEcsDedicatedHostClustersResult:
    """
    This data source provides the Ecs Dedicated Host Clusters of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.146.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ecs.get_ecs_dedicated_host_clusters(ids=["example_id"])
    pulumi.export("ecsDedicatedHostClusterId1", ids.clusters[0].id)
    name_regex = alicloud.ecs.get_ecs_dedicated_host_clusters(name_regex="^my-DedicatedHostCluster")
    pulumi.export("ecsDedicatedHostClusterId2", name_regex.clusters[0].id)
    zone_id = alicloud.ecs.get_ecs_dedicated_host_clusters(zone_id="example_value")
    pulumi.export("ecsDedicatedHostClusterId3", zone_id.clusters[0].id)
    cluster_name = alicloud.ecs.get_ecs_dedicated_host_clusters(dedicated_host_cluster_name="example_value")
    pulumi.export("ecsDedicatedHostClusterId4", cluster_name.clusters[0].id)
    cluster_ids = alicloud.ecs.get_ecs_dedicated_host_clusters(dedicated_host_cluster_ids=["example_id"])
    pulumi.export("ecsDedicatedHostClusterId5", cluster_ids.clusters[0].id)
    ```


    :param Sequence[str] dedicated_host_cluster_ids: The IDs of dedicated host clusters.
    :param str dedicated_host_cluster_name: The name of the dedicated host cluster.
    :param Sequence[str] ids: A list of Dedicated Host Cluster IDs.
    :param str name_regex: A regex string to filter results by Dedicated Host Cluster name.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    :param str zone_id: The zone ID of the dedicated host cluster.
    """
    __args__ = dict()
    __args__['dedicatedHostClusterIds'] = dedicated_host_cluster_ids
    __args__['dedicatedHostClusterName'] = dedicated_host_cluster_name
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['tags'] = tags
    __args__['zoneId'] = zone_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getEcsDedicatedHostClusters:getEcsDedicatedHostClusters', __args__, opts=opts, typ=GetEcsDedicatedHostClustersResult).value

    return AwaitableGetEcsDedicatedHostClustersResult(
        clusters=__ret__.clusters,
        dedicated_host_cluster_ids=__ret__.dedicated_host_cluster_ids,
        dedicated_host_cluster_name=__ret__.dedicated_host_cluster_name,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        tags=__ret__.tags,
        zone_id=__ret__.zone_id)


@_utilities.lift_output_func(get_ecs_dedicated_host_clusters)
def get_ecs_dedicated_host_clusters_output(dedicated_host_cluster_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                           dedicated_host_cluster_name: Optional[pulumi.Input[Optional[str]]] = None,
                                           ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                           name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                           tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                                           zone_id: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEcsDedicatedHostClustersResult]:
    """
    This data source provides the Ecs Dedicated Host Clusters of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.146.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ecs.get_ecs_dedicated_host_clusters(ids=["example_id"])
    pulumi.export("ecsDedicatedHostClusterId1", ids.clusters[0].id)
    name_regex = alicloud.ecs.get_ecs_dedicated_host_clusters(name_regex="^my-DedicatedHostCluster")
    pulumi.export("ecsDedicatedHostClusterId2", name_regex.clusters[0].id)
    zone_id = alicloud.ecs.get_ecs_dedicated_host_clusters(zone_id="example_value")
    pulumi.export("ecsDedicatedHostClusterId3", zone_id.clusters[0].id)
    cluster_name = alicloud.ecs.get_ecs_dedicated_host_clusters(dedicated_host_cluster_name="example_value")
    pulumi.export("ecsDedicatedHostClusterId4", cluster_name.clusters[0].id)
    cluster_ids = alicloud.ecs.get_ecs_dedicated_host_clusters(dedicated_host_cluster_ids=["example_id"])
    pulumi.export("ecsDedicatedHostClusterId5", cluster_ids.clusters[0].id)
    ```


    :param Sequence[str] dedicated_host_cluster_ids: The IDs of dedicated host clusters.
    :param str dedicated_host_cluster_name: The name of the dedicated host cluster.
    :param Sequence[str] ids: A list of Dedicated Host Cluster IDs.
    :param str name_regex: A regex string to filter results by Dedicated Host Cluster name.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    :param str zone_id: The zone ID of the dedicated host cluster.
    """
    ...
