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
    'GetProductAsEndUsersResult',
    'AwaitableGetProductAsEndUsersResult',
    'get_product_as_end_users',
    'get_product_as_end_users_output',
]

@pulumi.output_type
class GetProductAsEndUsersResult:
    """
    A collection of values returned by getProductAsEndUsers.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, output_file=None, sort_by=None, sort_order=None, users=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if sort_by and not isinstance(sort_by, str):
            raise TypeError("Expected argument 'sort_by' to be a str")
        pulumi.set(__self__, "sort_by", sort_by)
        if sort_order and not isinstance(sort_order, str):
            raise TypeError("Expected argument 'sort_order' to be a str")
        pulumi.set(__self__, "sort_order", sort_order)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

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
        A list of Product As End User IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="sortBy")
    def sort_by(self) -> Optional[str]:
        return pulumi.get(self, "sort_by")

    @property
    @pulumi.getter(name="sortOrder")
    def sort_order(self) -> Optional[str]:
        return pulumi.get(self, "sort_order")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetProductAsEndUsersUserResult']:
        """
        A list of Product As End User Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "users")


class AwaitableGetProductAsEndUsersResult(GetProductAsEndUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProductAsEndUsersResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            sort_by=self.sort_by,
            sort_order=self.sort_order,
            users=self.users)


def get_product_as_end_users(ids: Optional[Sequence[str]] = None,
                             name_regex: Optional[str] = None,
                             output_file: Optional[str] = None,
                             sort_by: Optional[str] = None,
                             sort_order: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProductAsEndUsersResult:
    """
    This data source provides Service Catalog Product As End User available to the user.[What is Product As End User](https://www.alibabacloud.com/help/en/servicecatalog/latest/api-doc-servicecatalog-2021-09-01-api-doc-listproductsasenduser)

    > **NOTE:** Available in 1.196.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.servicecatalog.get_product_as_end_users(name_regex="ram模板创建")
    pulumi.export("alicloudServiceCatalogProductAsEndUserExampleId", default.users[0].id)
    ```


    :param Sequence[str] ids: A list of Product As End User IDs.
    :param str name_regex: A regex string to filter results by product name.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['sortBy'] = sort_by
    __args__['sortOrder'] = sort_order
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:servicecatalog/getProductAsEndUsers:getProductAsEndUsers', __args__, opts=opts, typ=GetProductAsEndUsersResult).value

    return AwaitableGetProductAsEndUsersResult(
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        sort_by=__ret__.sort_by,
        sort_order=__ret__.sort_order,
        users=__ret__.users)


@_utilities.lift_output_func(get_product_as_end_users)
def get_product_as_end_users_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    sort_by: Optional[pulumi.Input[Optional[str]]] = None,
                                    sort_order: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProductAsEndUsersResult]:
    """
    This data source provides Service Catalog Product As End User available to the user.[What is Product As End User](https://www.alibabacloud.com/help/en/servicecatalog/latest/api-doc-servicecatalog-2021-09-01-api-doc-listproductsasenduser)

    > **NOTE:** Available in 1.196.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.servicecatalog.get_product_as_end_users(name_regex="ram模板创建")
    pulumi.export("alicloudServiceCatalogProductAsEndUserExampleId", default.users[0].id)
    ```


    :param Sequence[str] ids: A list of Product As End User IDs.
    :param str name_regex: A regex string to filter results by product name.
    """
    ...
