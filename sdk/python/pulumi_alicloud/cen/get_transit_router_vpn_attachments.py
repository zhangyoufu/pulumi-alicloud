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
    'GetTransitRouterVpnAttachmentsResult',
    'AwaitableGetTransitRouterVpnAttachmentsResult',
    'get_transit_router_vpn_attachments',
    'get_transit_router_vpn_attachments_output',
]

@pulumi.output_type
class GetTransitRouterVpnAttachmentsResult:
    """
    A collection of values returned by getTransitRouterVpnAttachments.
    """
    def __init__(__self__, attachments=None, cen_id=None, id=None, ids=None, name_regex=None, names=None, output_file=None, status=None, transit_router_id=None):
        if attachments and not isinstance(attachments, list):
            raise TypeError("Expected argument 'attachments' to be a list")
        pulumi.set(__self__, "attachments", attachments)
        if cen_id and not isinstance(cen_id, str):
            raise TypeError("Expected argument 'cen_id' to be a str")
        pulumi.set(__self__, "cen_id", cen_id)
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
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if transit_router_id and not isinstance(transit_router_id, str):
            raise TypeError("Expected argument 'transit_router_id' to be a str")
        pulumi.set(__self__, "transit_router_id", transit_router_id)

    @property
    @pulumi.getter
    def attachments(self) -> Sequence['outputs.GetTransitRouterVpnAttachmentsAttachmentResult']:
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> str:
        return pulumi.get(self, "cen_id")

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
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> Optional[str]:
        return pulumi.get(self, "transit_router_id")


class AwaitableGetTransitRouterVpnAttachmentsResult(GetTransitRouterVpnAttachmentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransitRouterVpnAttachmentsResult(
            attachments=self.attachments,
            cen_id=self.cen_id,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            status=self.status,
            transit_router_id=self.transit_router_id)


def get_transit_router_vpn_attachments(cen_id: Optional[str] = None,
                                       ids: Optional[Sequence[str]] = None,
                                       name_regex: Optional[str] = None,
                                       output_file: Optional[str] = None,
                                       status: Optional[str] = None,
                                       transit_router_id: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransitRouterVpnAttachmentsResult:
    """
    This data source provides the Cen Transit Router Vpn Attachments of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.183.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cen.get_transit_router_vpn_attachments(cen_id="example_value")
    pulumi.export("cenTransitRouterVpnAttachmentId1", ids.attachments[0].id)
    ```


    :param str cen_id: The id of the cen.
    :param Sequence[str] ids: A list of Transit Router Vpn Attachment IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the transit router attachment.
    :param str transit_router_id: The ID of the forwarding router instance.
    """
    __args__ = dict()
    __args__['cenId'] = cen_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['transitRouterId'] = transit_router_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getTransitRouterVpnAttachments:getTransitRouterVpnAttachments', __args__, opts=opts, typ=GetTransitRouterVpnAttachmentsResult).value

    return AwaitableGetTransitRouterVpnAttachmentsResult(
        attachments=__ret__.attachments,
        cen_id=__ret__.cen_id,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        status=__ret__.status,
        transit_router_id=__ret__.transit_router_id)


@_utilities.lift_output_func(get_transit_router_vpn_attachments)
def get_transit_router_vpn_attachments_output(cen_id: Optional[pulumi.Input[str]] = None,
                                              ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                              name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                              status: Optional[pulumi.Input[Optional[str]]] = None,
                                              transit_router_id: Optional[pulumi.Input[Optional[str]]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransitRouterVpnAttachmentsResult]:
    """
    This data source provides the Cen Transit Router Vpn Attachments of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.183.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cen.get_transit_router_vpn_attachments(cen_id="example_value")
    pulumi.export("cenTransitRouterVpnAttachmentId1", ids.attachments[0].id)
    ```


    :param str cen_id: The id of the cen.
    :param Sequence[str] ids: A list of Transit Router Vpn Attachment IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the transit router attachment.
    :param str transit_router_id: The ID of the forwarding router instance.
    """
    ...
