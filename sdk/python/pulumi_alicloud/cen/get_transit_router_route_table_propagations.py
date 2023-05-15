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
    'GetTransitRouterRouteTablePropagationsResult',
    'AwaitableGetTransitRouterRouteTablePropagationsResult',
    'get_transit_router_route_table_propagations',
    'get_transit_router_route_table_propagations_output',
]

@pulumi.output_type
class GetTransitRouterRouteTablePropagationsResult:
    """
    A collection of values returned by getTransitRouterRouteTablePropagations.
    """
    def __init__(__self__, id=None, ids=None, output_file=None, propagations=None, status=None, transit_router_route_table_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if propagations and not isinstance(propagations, list):
            raise TypeError("Expected argument 'propagations' to be a list")
        pulumi.set(__self__, "propagations", propagations)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if transit_router_route_table_id and not isinstance(transit_router_route_table_id, str):
            raise TypeError("Expected argument 'transit_router_route_table_id' to be a str")
        pulumi.set(__self__, "transit_router_route_table_id", transit_router_route_table_id)

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
        A list of CEN Transit Router Route Table Association IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def propagations(self) -> Sequence['outputs.GetTransitRouterRouteTablePropagationsPropagationResult']:
        """
        A list of CEN Transit Router Route Table Propagations. Each element contains the following attributes:
        """
        return pulumi.get(self, "propagations")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the route table.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transitRouterRouteTableId")
    def transit_router_route_table_id(self) -> str:
        """
        ID of the transit router route table.
        """
        return pulumi.get(self, "transit_router_route_table_id")


class AwaitableGetTransitRouterRouteTablePropagationsResult(GetTransitRouterRouteTablePropagationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransitRouterRouteTablePropagationsResult(
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            propagations=self.propagations,
            status=self.status,
            transit_router_route_table_id=self.transit_router_route_table_id)


def get_transit_router_route_table_propagations(ids: Optional[Sequence[str]] = None,
                                                output_file: Optional[str] = None,
                                                status: Optional[str] = None,
                                                transit_router_route_table_id: Optional[str] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransitRouterRouteTablePropagationsResult:
    """
    This data source provides CEN Transit Router Route Table Propagations available to the user.[What is Cen Transit Router Route Table Propagations](https://help.aliyun.com/document_detail/261245.html)

    > **NOTE:** Available in 1.126.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.cen.get_transit_router_route_table_propagations(transit_router_route_table_id="rtb-id1")
    pulumi.export("firstTransitRouterPeerAttachmentsTransitRouterAttachmentResourceType", default.propagations[0].resource_type)
    ```


    :param Sequence[str] ids: A list of CEN Transit Router Route Table Association IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the route table, including `Active`, `Enabling`, `Disabling`, `Deleted`.
    :param str transit_router_route_table_id: ID of the route table of the VPC or VBR.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['transitRouterRouteTableId'] = transit_router_route_table_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getTransitRouterRouteTablePropagations:getTransitRouterRouteTablePropagations', __args__, opts=opts, typ=GetTransitRouterRouteTablePropagationsResult).value

    return AwaitableGetTransitRouterRouteTablePropagationsResult(
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        propagations=__ret__.propagations,
        status=__ret__.status,
        transit_router_route_table_id=__ret__.transit_router_route_table_id)


@_utilities.lift_output_func(get_transit_router_route_table_propagations)
def get_transit_router_route_table_propagations_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                                       status: Optional[pulumi.Input[Optional[str]]] = None,
                                                       transit_router_route_table_id: Optional[pulumi.Input[str]] = None,
                                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransitRouterRouteTablePropagationsResult]:
    """
    This data source provides CEN Transit Router Route Table Propagations available to the user.[What is Cen Transit Router Route Table Propagations](https://help.aliyun.com/document_detail/261245.html)

    > **NOTE:** Available in 1.126.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.cen.get_transit_router_route_table_propagations(transit_router_route_table_id="rtb-id1")
    pulumi.export("firstTransitRouterPeerAttachmentsTransitRouterAttachmentResourceType", default.propagations[0].resource_type)
    ```


    :param Sequence[str] ids: A list of CEN Transit Router Route Table Association IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the route table, including `Active`, `Enabling`, `Disabling`, `Deleted`.
    :param str transit_router_route_table_id: ID of the route table of the VPC or VBR.
    """
    ...
