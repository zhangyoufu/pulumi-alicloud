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
    'GetDataCentersResult',
    'AwaitableGetDataCentersResult',
    'get_data_centers',
    'get_data_centers_output',
]

@pulumi.output_type
class GetDataCentersResult:
    """
    A collection of values returned by getDataCenters.
    """
    def __init__(__self__, centers=None, cluster_id=None, id=None, ids=None, name_regex=None, names=None, output_file=None):
        if centers and not isinstance(centers, list):
            raise TypeError("Expected argument 'centers' to be a list")
        pulumi.set(__self__, "centers", centers)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
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

    @property
    @pulumi.getter
    def centers(self) -> Sequence['outputs.GetDataCentersCenterResult']:
        """
        A list of Cassandra data centers. Its every element contains the following attributes:
        """
        return pulumi.get(self, "centers")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The ID of the Cassandra cluster.
        """
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
        """
        The list of Cassandra data center ids.
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
        The name list of Cassandra data centers.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetDataCentersResult(GetDataCentersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataCentersResult(
            centers=self.centers,
            cluster_id=self.cluster_id,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)


def get_data_centers(cluster_id: Optional[str] = None,
                     ids: Optional[Sequence[str]] = None,
                     name_regex: Optional[str] = None,
                     output_file: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDataCentersResult:
    """
    The `cassandra_get_data_centers` data source provides a collection of Cassandra Data Centers available in Alicloud account.
    Filters support regular expression for the cluster name or ids.

    > **NOTE:**  Available in 1.88.0+.

    > **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    cassandra = alicloud.cassandra.get_data_centers(name_regex="tf_testAccCassandra_dc",
        cluster_id="cds-xxxxx")
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_id: The cluster id of dataCenters belongs to.
    :param Sequence[str] ids: The list of Cassandra data center ids.
    :param str name_regex: A regex string to apply to the cluster name.
    :param str output_file: The name of file that can save the collection of data centers after running `pulumi preview`.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cassandra/getDataCenters:getDataCenters', __args__, opts=opts, typ=GetDataCentersResult).value

    return AwaitableGetDataCentersResult(
        centers=pulumi.get(__ret__, 'centers'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'))


@_utilities.lift_output_func(get_data_centers)
def get_data_centers_output(cluster_id: Optional[pulumi.Input[str]] = None,
                            ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDataCentersResult]:
    """
    The `cassandra_get_data_centers` data source provides a collection of Cassandra Data Centers available in Alicloud account.
    Filters support regular expression for the cluster name or ids.

    > **NOTE:**  Available in 1.88.0+.

    > **DEPRECATED:**  This data source has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    cassandra = alicloud.cassandra.get_data_centers(name_regex="tf_testAccCassandra_dc",
        cluster_id="cds-xxxxx")
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_id: The cluster id of dataCenters belongs to.
    :param Sequence[str] ids: The list of Cassandra data center ids.
    :param str name_regex: A regex string to apply to the cluster name.
    :param str output_file: The name of file that can save the collection of data centers after running `pulumi preview`.
    """
    ...
