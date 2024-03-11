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
    'GetSecondaryIndexesResult',
    'AwaitableGetSecondaryIndexesResult',
    'get_secondary_indexes',
    'get_secondary_indexes_output',
]

@pulumi.output_type
class GetSecondaryIndexesResult:
    """
    A collection of values returned by getSecondaryIndexes.
    """
    def __init__(__self__, id=None, ids=None, indexes=None, instance_name=None, name_regex=None, names=None, output_file=None, table_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if indexes and not isinstance(indexes, list):
            raise TypeError("Expected argument 'indexes' to be a list")
        pulumi.set(__self__, "indexes", indexes)
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        pulumi.set(__self__, "instance_name", instance_name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if table_name and not isinstance(table_name, str):
            raise TypeError("Expected argument 'table_name' to be a str")
        pulumi.set(__self__, "table_name", table_name)

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
        A list of secondary index IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def indexes(self) -> Sequence['outputs.GetSecondaryIndexesIndexResult']:
        """
        A list of indexes. Each element contains the following attributes:
        """
        return pulumi.get(self, "indexes")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        """
        The OTS instance name.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of secondary index  names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> str:
        """
        The table name of the OTS which could not be changed.
        """
        return pulumi.get(self, "table_name")


class AwaitableGetSecondaryIndexesResult(GetSecondaryIndexesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecondaryIndexesResult(
            id=self.id,
            ids=self.ids,
            indexes=self.indexes,
            instance_name=self.instance_name,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            table_name=self.table_name)


def get_secondary_indexes(ids: Optional[Sequence[str]] = None,
                          instance_name: Optional[str] = None,
                          name_regex: Optional[str] = None,
                          output_file: Optional[str] = None,
                          table_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecondaryIndexesResult:
    """
    This data source provides the ots secondary index of the current Alibaba Cloud user.

    For information about OTS secondary index and how to use it, see [Secondary index overview](https://www.alibabacloud.com/help/en/tablestore/latest/secondary-index-overview).

    > **NOTE:** Available in v1.187.0+.

    ## Example Usage


    :param Sequence[str] ids: A list of secondary index IDs.
    :param str instance_name: The name of OTS instance.
    :param str name_regex: A regex string to filter results by secondary index name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str table_name: The name of OTS table.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceName'] = instance_name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['tableName'] = table_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ots/getSecondaryIndexes:getSecondaryIndexes', __args__, opts=opts, typ=GetSecondaryIndexesResult).value

    return AwaitableGetSecondaryIndexesResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        indexes=pulumi.get(__ret__, 'indexes'),
        instance_name=pulumi.get(__ret__, 'instance_name'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        table_name=pulumi.get(__ret__, 'table_name'))


@_utilities.lift_output_func(get_secondary_indexes)
def get_secondary_indexes_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 instance_name: Optional[pulumi.Input[str]] = None,
                                 name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 table_name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecondaryIndexesResult]:
    """
    This data source provides the ots secondary index of the current Alibaba Cloud user.

    For information about OTS secondary index and how to use it, see [Secondary index overview](https://www.alibabacloud.com/help/en/tablestore/latest/secondary-index-overview).

    > **NOTE:** Available in v1.187.0+.

    ## Example Usage


    :param Sequence[str] ids: A list of secondary index IDs.
    :param str instance_name: The name of OTS instance.
    :param str name_regex: A regex string to filter results by secondary index name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str table_name: The name of OTS table.
    """
    ...
