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
    'GetClustersResult',
    'AwaitableGetClustersResult',
    'get_clusters',
    'get_clusters_output',
]

@pulumi.output_type
class GetClustersResult:
    """
    A collection of values returned by getClusters.
    """
    def __init__(__self__, cluster_alias_name=None, clusters=None, enable_details=None, id=None, ids=None, name_regex=None, names=None, output_file=None, request_pars=None, status=None):
        if cluster_alias_name and not isinstance(cluster_alias_name, str):
            raise TypeError("Expected argument 'cluster_alias_name' to be a str")
        pulumi.set(__self__, "cluster_alias_name", cluster_alias_name)
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
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
        if request_pars and not isinstance(request_pars, str):
            raise TypeError("Expected argument 'request_pars' to be a str")
        pulumi.set(__self__, "request_pars", request_pars)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="clusterAliasName")
    def cluster_alias_name(self) -> Optional[str]:
        return pulumi.get(self, "cluster_alias_name")

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.GetClustersClusterResult']:
        """
        A list of MSE Clusters. Each element contains the following attributes:
        """
        return pulumi.get(self, "clusters")

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
        """
        A list of MSE Cluster ids.
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
        A list of MSE Cluster names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="requestPars")
    def request_pars(self) -> Optional[str]:
        return pulumi.get(self, "request_pars")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of MSE Cluster.
        """
        return pulumi.get(self, "status")


class AwaitableGetClustersResult(GetClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClustersResult(
            cluster_alias_name=self.cluster_alias_name,
            clusters=self.clusters,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            request_pars=self.request_pars,
            status=self.status)


def get_clusters(cluster_alias_name: Optional[str] = None,
                 enable_details: Optional[bool] = None,
                 ids: Optional[Sequence[str]] = None,
                 name_regex: Optional[str] = None,
                 output_file: Optional[str] = None,
                 request_pars: Optional[str] = None,
                 status: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClustersResult:
    """
    This data source provides a list of MSE Clusters in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in v1.94.0+.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.mse.get_clusters(ids=["mse-cn-0d9xxxx"],
        status="INIT_SUCCESS")
    pulumi.export("clusterId", example.clusters[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_alias_name: The alias name of MSE Cluster.
    :param Sequence[str] ids: A list of MSE Cluster ids.
    :param str name_regex: A regex string to filter the results by the cluster alias name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of MSE Cluster. Valid: `DESTROY_FAILED`, `DESTROY_ING`, `DESTROY_SUCCESS`, `INIT_FAILED`, `INIT_ING`, `INIT_SUCCESS`, `INIT_TIME_OUT`, `RESTART_FAILED`, `RESTART_ING`, `RESTART_SUCCESS`, `SCALE_FAILED`, `SCALE_ING`, `SCALE_SUCCESS`
    """
    __args__ = dict()
    __args__['clusterAliasName'] = cluster_alias_name
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['requestPars'] = request_pars
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:mse/getClusters:getClusters', __args__, opts=opts, typ=GetClustersResult).value

    return AwaitableGetClustersResult(
        cluster_alias_name=pulumi.get(__ret__, 'cluster_alias_name'),
        clusters=pulumi.get(__ret__, 'clusters'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        request_pars=pulumi.get(__ret__, 'request_pars'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_clusters)
def get_clusters_output(cluster_alias_name: Optional[pulumi.Input[Optional[str]]] = None,
                        enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                        ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        request_pars: Optional[pulumi.Input[Optional[str]]] = None,
                        status: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClustersResult]:
    """
    This data source provides a list of MSE Clusters in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in v1.94.0+.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.mse.get_clusters(ids=["mse-cn-0d9xxxx"],
        status="INIT_SUCCESS")
    pulumi.export("clusterId", example.clusters[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_alias_name: The alias name of MSE Cluster.
    :param Sequence[str] ids: A list of MSE Cluster ids.
    :param str name_regex: A regex string to filter the results by the cluster alias name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of MSE Cluster. Valid: `DESTROY_FAILED`, `DESTROY_ING`, `DESTROY_SUCCESS`, `INIT_FAILED`, `INIT_ING`, `INIT_SUCCESS`, `INIT_TIME_OUT`, `RESTART_FAILED`, `RESTART_ING`, `RESTART_SUCCESS`, `SCALE_FAILED`, `SCALE_ING`, `SCALE_SUCCESS`
    """
    ...
