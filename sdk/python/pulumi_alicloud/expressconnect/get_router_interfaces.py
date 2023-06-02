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
from ._inputs import *

__all__ = [
    'GetRouterInterfacesResult',
    'AwaitableGetRouterInterfacesResult',
    'get_router_interfaces',
    'get_router_interfaces_output',
]

@pulumi.output_type
class GetRouterInterfacesResult:
    """
    A collection of values returned by getRouterInterfaces.
    """
    def __init__(__self__, filters=None, id=None, ids=None, include_reservation_data=None, interfaces=None, name_regex=None, names=None, output_file=None, page_number=None, page_size=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if include_reservation_data and not isinstance(include_reservation_data, str):
            raise TypeError("Expected argument 'include_reservation_data' to be a str")
        pulumi.set(__self__, "include_reservation_data", include_reservation_data)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if page_number and not isinstance(page_number, int):
            raise TypeError("Expected argument 'page_number' to be a int")
        pulumi.set(__self__, "page_number", page_number)
        if page_size and not isinstance(page_size, int):
            raise TypeError("Expected argument 'page_size' to be a int")
        pulumi.set(__self__, "page_size", page_size)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetRouterInterfacesFilterResult']]:
        return pulumi.get(self, "filters")

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
        A list of Router Interface IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="includeReservationData")
    def include_reservation_data(self) -> Optional[str]:
        return pulumi.get(self, "include_reservation_data")

    @property
    @pulumi.getter
    def interfaces(self) -> Sequence['outputs.GetRouterInterfacesInterfaceResult']:
        """
        A list of Router Interface Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of name of Router Interfaces.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="pageNumber")
    def page_number(self) -> Optional[int]:
        return pulumi.get(self, "page_number")

    @property
    @pulumi.getter(name="pageSize")
    def page_size(self) -> Optional[int]:
        return pulumi.get(self, "page_size")


class AwaitableGetRouterInterfacesResult(GetRouterInterfacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterInterfacesResult(
            filters=self.filters,
            id=self.id,
            ids=self.ids,
            include_reservation_data=self.include_reservation_data,
            interfaces=self.interfaces,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size)


def get_router_interfaces(filters: Optional[Sequence[pulumi.InputType['GetRouterInterfacesFilterArgs']]] = None,
                          ids: Optional[Sequence[str]] = None,
                          include_reservation_data: Optional[str] = None,
                          name_regex: Optional[str] = None,
                          output_file: Optional[str] = None,
                          page_number: Optional[int] = None,
                          page_size: Optional[int] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterInterfacesResult:
    """
    This data source provides Router Interface available to the user.[What is Router Interface](https://www.alibabacloud.com/help/doc-detail/52412.htm)

    > **NOTE:** Available in 1.199.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.expressconnect.get_router_interfaces(ids=[alicloud_router_interface["default"]["id"]],
        name_regex=%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
    pulumi.export("alicloudRouterInterfaceExampleId", default.interfaces[0].id)
    ```


    :param Sequence[str] ids: A list of Router Interface IDs.
    :param str include_reservation_data: Does it contain renewal data. Valid values: `true`, `false`.
    :param str name_regex: A regex string to filter results by Group Metric Rule name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['ids'] = ids
    __args__['includeReservationData'] = include_reservation_data
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:expressconnect/getRouterInterfaces:getRouterInterfaces', __args__, opts=opts, typ=GetRouterInterfacesResult).value

    return AwaitableGetRouterInterfacesResult(
        filters=__ret__.filters,
        id=__ret__.id,
        ids=__ret__.ids,
        include_reservation_data=__ret__.include_reservation_data,
        interfaces=__ret__.interfaces,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        page_number=__ret__.page_number,
        page_size=__ret__.page_size)


@_utilities.lift_output_func(get_router_interfaces)
def get_router_interfaces_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetRouterInterfacesFilterArgs']]]]] = None,
                                 ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 include_reservation_data: Optional[pulumi.Input[Optional[str]]] = None,
                                 name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 page_number: Optional[pulumi.Input[Optional[int]]] = None,
                                 page_size: Optional[pulumi.Input[Optional[int]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterInterfacesResult]:
    """
    This data source provides Router Interface available to the user.[What is Router Interface](https://www.alibabacloud.com/help/doc-detail/52412.htm)

    > **NOTE:** Available in 1.199.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.expressconnect.get_router_interfaces(ids=[alicloud_router_interface["default"]["id"]],
        name_regex=%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
    pulumi.export("alicloudRouterInterfaceExampleId", default.interfaces[0].id)
    ```


    :param Sequence[str] ids: A list of Router Interface IDs.
    :param str include_reservation_data: Does it contain renewal data. Valid values: `true`, `false`.
    :param str name_regex: A regex string to filter results by Group Metric Rule name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
