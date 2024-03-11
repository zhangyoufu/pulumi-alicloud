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
    'GetDbClustersResult',
    'AwaitableGetDbClustersResult',
    'get_db_clusters',
    'get_db_clusters_output',
]

@pulumi.output_type
class GetDbClustersResult:
    """
    A collection of values returned by getDbClusters.
    """
    def __init__(__self__, clusters=None, db_cluster_description=None, enable_details=None, id=None, ids=None, output_file=None, status=None):
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if db_cluster_description and not isinstance(db_cluster_description, str):
            raise TypeError("Expected argument 'db_cluster_description' to be a str")
        pulumi.set(__self__, "db_cluster_description", db_cluster_description)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.GetDbClustersClusterResult']:
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter(name="dbClusterDescription")
    def db_cluster_description(self) -> Optional[str]:
        return pulumi.get(self, "db_cluster_description")

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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetDbClustersResult(GetDbClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDbClustersResult(
            clusters=self.clusters,
            db_cluster_description=self.db_cluster_description,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            status=self.status)


def get_db_clusters(db_cluster_description: Optional[str] = None,
                    enable_details: Optional[bool] = None,
                    ids: Optional[Sequence[str]] = None,
                    output_file: Optional[str] = None,
                    status: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDbClustersResult:
    """
    This data source provides the Click House DBCluster of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.134.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_db_cluster = alicloud.clickhouse.DbCluster("defaultDbCluster",
        db_cluster_version="20.3.10.75",
        category="Basic",
        db_cluster_class="S8",
        db_cluster_network_type="vpc",
        db_node_group_count=1,
        payment_type="PayAsYouGo",
        db_node_storage="500",
        storage_type="cloud_essd",
        vswitch_id="your_vswitch_id")
    default_db_clusters = alicloud.clickhouse.get_db_clusters_output(ids=[default_db_cluster.id])
    pulumi.export("dbCluster", default_db_clusters.ids[0])
    ```
    <!--End PulumiCodeChooser -->


    :param str db_cluster_description: The DBCluster description.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of DBCluster IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the DBCluster. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`.
    """
    __args__ = dict()
    __args__['dbClusterDescription'] = db_cluster_description
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:clickhouse/getDbClusters:getDbClusters', __args__, opts=opts, typ=GetDbClustersResult).value

    return AwaitableGetDbClustersResult(
        clusters=pulumi.get(__ret__, 'clusters'),
        db_cluster_description=pulumi.get(__ret__, 'db_cluster_description'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        output_file=pulumi.get(__ret__, 'output_file'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_db_clusters)
def get_db_clusters_output(db_cluster_description: Optional[pulumi.Input[Optional[str]]] = None,
                           enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                           ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           status: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDbClustersResult]:
    """
    This data source provides the Click House DBCluster of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.134.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_db_cluster = alicloud.clickhouse.DbCluster("defaultDbCluster",
        db_cluster_version="20.3.10.75",
        category="Basic",
        db_cluster_class="S8",
        db_cluster_network_type="vpc",
        db_node_group_count=1,
        payment_type="PayAsYouGo",
        db_node_storage="500",
        storage_type="cloud_essd",
        vswitch_id="your_vswitch_id")
    default_db_clusters = alicloud.clickhouse.get_db_clusters_output(ids=[default_db_cluster.id])
    pulumi.export("dbCluster", default_db_clusters.ids[0])
    ```
    <!--End PulumiCodeChooser -->


    :param str db_cluster_description: The DBCluster description.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of DBCluster IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the DBCluster. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`.
    """
    ...
