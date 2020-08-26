# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetDatabasesResult',
    'AwaitableGetDatabasesResult',
    'get_databases',
]

@pulumi.output_type
class GetDatabasesResult:
    """
    A collection of values returned by getDatabases.
    """
    def __init__(__self__, databases=None, db_cluster_id=None, id=None, name_regex=None, names=None):
        if databases and not isinstance(databases, list):
            raise TypeError("Expected argument 'databases' to be a list")
        pulumi.set(__self__, "databases", databases)
        if db_cluster_id and not isinstance(db_cluster_id, str):
            raise TypeError("Expected argument 'db_cluster_id' to be a str")
        pulumi.set(__self__, "db_cluster_id", db_cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)

    @property
    @pulumi.getter
    def databases(self) -> List['outputs.GetDatabasesDatabaseResult']:
        """
        A list of PolarDB cluster databases. Each element contains the following attributes:
        """
        return pulumi.get(self, "databases")

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> str:
        return pulumi.get(self, "db_cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> List[str]:
        """
        database name of the cluster.
        """
        return pulumi.get(self, "names")


class AwaitableGetDatabasesResult(GetDatabasesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasesResult(
            databases=self.databases,
            db_cluster_id=self.db_cluster_id,
            id=self.id,
            name_regex=self.name_regex,
            names=self.names)


def get_databases(db_cluster_id: Optional[str] = None,
                  name_regex: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasesResult:
    """
    The `polardb.getDatabases` data source provides a collection of PolarDB cluster database available in Alibaba Cloud account.
    Filters support regular expression for the database name, searches by clusterId.

    > **NOTE:** Available in v1.70.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    polardb_clusters_ds = alicloud.polardb.get_clusters(description_regex="pc-\\w+",
        status="Running")
    default = alicloud.polardb.get_databases(db_cluster_id=polardb_clusters_ds.clusters[0].id)
    pulumi.export("ends", default.databases[0].db_name)
    ```


    :param str db_cluster_id: The polarDB cluster ID.
    :param str name_regex: A regex string to filter results by database name.
    """
    __args__ = dict()
    __args__['dbClusterId'] = db_cluster_id
    __args__['nameRegex'] = name_regex
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:polardb/getDatabases:getDatabases', __args__, opts=opts, typ=GetDatabasesResult).value

    return AwaitableGetDatabasesResult(
        databases=__ret__.databases,
        db_cluster_id=__ret__.db_cluster_id,
        id=__ret__.id,
        name_regex=__ret__.name_regex,
        names=__ret__.names)
