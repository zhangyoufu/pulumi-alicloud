# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetTablesResult:
    """
    A collection of values returned by getTables.
    """
    def __init__(__self__, id=None, ids=None, instance_name=None, name_regex=None, names=None, output_file=None, tables=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of table IDs.
        """
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        __self__.instance_name = instance_name
        """
        The OTS instance name.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of table names.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if tables and not isinstance(tables, list):
            raise TypeError("Expected argument 'tables' to be a list")
        __self__.tables = tables
        """
        A list of tables. Each element contains the following attributes:
        """
class AwaitableGetTablesResult(GetTablesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTablesResult(
            id=self.id,
            ids=self.ids,
            instance_name=self.instance_name,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            tables=self.tables)

def get_tables(ids=None,instance_name=None,name_regex=None,output_file=None,opts=None):
    """
    This data source provides the ots tables of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.40.0+.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ots_tables.html.markdown.


    :param list ids: A list of table IDs.
    :param str instance_name: The name of OTS instance.
    :param str name_regex: A regex string to filter results by table name.
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['instanceName'] = instance_name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:oss/getTables:getTables', __args__, opts=opts).value

    return AwaitableGetTablesResult(
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        instance_name=__ret__.get('instanceName'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        tables=__ret__.get('tables'))
