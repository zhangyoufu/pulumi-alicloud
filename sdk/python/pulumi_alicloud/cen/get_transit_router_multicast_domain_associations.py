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
    'GetTransitRouterMulticastDomainAssociationsResult',
    'AwaitableGetTransitRouterMulticastDomainAssociationsResult',
    'get_transit_router_multicast_domain_associations',
    'get_transit_router_multicast_domain_associations_output',
]

@pulumi.output_type
class GetTransitRouterMulticastDomainAssociationsResult:
    """
    A collection of values returned by getTransitRouterMulticastDomainAssociations.
    """
    def __init__(__self__, associations=None, id=None, ids=None, output_file=None, resource_id=None, resource_type=None, status=None, transit_router_attachment_id=None, transit_router_multicast_domain_id=None, vswitch_id=None):
        if associations and not isinstance(associations, list):
            raise TypeError("Expected argument 'associations' to be a list")
        pulumi.set(__self__, "associations", associations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)
        if resource_type and not isinstance(resource_type, str):
            raise TypeError("Expected argument 'resource_type' to be a str")
        pulumi.set(__self__, "resource_type", resource_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if transit_router_attachment_id and not isinstance(transit_router_attachment_id, str):
            raise TypeError("Expected argument 'transit_router_attachment_id' to be a str")
        pulumi.set(__self__, "transit_router_attachment_id", transit_router_attachment_id)
        if transit_router_multicast_domain_id and not isinstance(transit_router_multicast_domain_id, str):
            raise TypeError("Expected argument 'transit_router_multicast_domain_id' to be a str")
        pulumi.set(__self__, "transit_router_multicast_domain_id", transit_router_multicast_domain_id)
        if vswitch_id and not isinstance(vswitch_id, str):
            raise TypeError("Expected argument 'vswitch_id' to be a str")
        pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter
    def associations(self) -> Sequence['outputs.GetTransitRouterMulticastDomainAssociationsAssociationResult']:
        """
        A list of Cen Transit Router Multicast Domain Associations. Each element contains the following attributes:
        """
        return pulumi.get(self, "associations")

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
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[str]:
        """
        The ID of the resource associated with the multicast domain.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        """
        The type of resource associated with the multicast domain.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the associated resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> Optional[str]:
        """
        The ID of the network instance connection.
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @property
    @pulumi.getter(name="transitRouterMulticastDomainId")
    def transit_router_multicast_domain_id(self) -> str:
        """
        The ID of the multicast domain.
        """
        return pulumi.get(self, "transit_router_multicast_domain_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[str]:
        """
        The ID of the vSwitch.
        """
        return pulumi.get(self, "vswitch_id")


class AwaitableGetTransitRouterMulticastDomainAssociationsResult(GetTransitRouterMulticastDomainAssociationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransitRouterMulticastDomainAssociationsResult(
            associations=self.associations,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            resource_id=self.resource_id,
            resource_type=self.resource_type,
            status=self.status,
            transit_router_attachment_id=self.transit_router_attachment_id,
            transit_router_multicast_domain_id=self.transit_router_multicast_domain_id,
            vswitch_id=self.vswitch_id)


def get_transit_router_multicast_domain_associations(ids: Optional[Sequence[str]] = None,
                                                     output_file: Optional[str] = None,
                                                     resource_id: Optional[str] = None,
                                                     resource_type: Optional[str] = None,
                                                     status: Optional[str] = None,
                                                     transit_router_attachment_id: Optional[str] = None,
                                                     transit_router_multicast_domain_id: Optional[str] = None,
                                                     vswitch_id: Optional[str] = None,
                                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransitRouterMulticastDomainAssociationsResult:
    """
    This data source provides the Cen Transit Router Multicast Domain Associations of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.195.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cen.get_transit_router_multicast_domain_associations(ids=["example_id"],
        transit_router_multicast_domain_id="your_transit_router_multicast_domain_id")
    pulumi.export("cenTransitRouterMulticastDomainId0", ids.associations[0].id)
    ```


    :param Sequence[str] ids: A list of Transit Router Multicast Domain Association IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_id: The ID of the resource associated with the multicast domain.
    :param str resource_type: The type of resource associated with the multicast domain. Valid Value: `VPC`.
    :param str status: The status of the associated resource. Valid Value: `Associated`, `Associating`, `Dissociating`.
    :param str transit_router_attachment_id: The ID of the network instance connection.
    :param str transit_router_multicast_domain_id: The ID of the multicast domain.
    :param str vswitch_id: The ID of the vSwitch.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['resourceId'] = resource_id
    __args__['resourceType'] = resource_type
    __args__['status'] = status
    __args__['transitRouterAttachmentId'] = transit_router_attachment_id
    __args__['transitRouterMulticastDomainId'] = transit_router_multicast_domain_id
    __args__['vswitchId'] = vswitch_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getTransitRouterMulticastDomainAssociations:getTransitRouterMulticastDomainAssociations', __args__, opts=opts, typ=GetTransitRouterMulticastDomainAssociationsResult).value

    return AwaitableGetTransitRouterMulticastDomainAssociationsResult(
        associations=__ret__.associations,
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        resource_id=__ret__.resource_id,
        resource_type=__ret__.resource_type,
        status=__ret__.status,
        transit_router_attachment_id=__ret__.transit_router_attachment_id,
        transit_router_multicast_domain_id=__ret__.transit_router_multicast_domain_id,
        vswitch_id=__ret__.vswitch_id)


@_utilities.lift_output_func(get_transit_router_multicast_domain_associations)
def get_transit_router_multicast_domain_associations_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                                            resource_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                            resource_type: Optional[pulumi.Input[Optional[str]]] = None,
                                                            status: Optional[pulumi.Input[Optional[str]]] = None,
                                                            transit_router_attachment_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                            transit_router_multicast_domain_id: Optional[pulumi.Input[str]] = None,
                                                            vswitch_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransitRouterMulticastDomainAssociationsResult]:
    """
    This data source provides the Cen Transit Router Multicast Domain Associations of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.195.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cen.get_transit_router_multicast_domain_associations(ids=["example_id"],
        transit_router_multicast_domain_id="your_transit_router_multicast_domain_id")
    pulumi.export("cenTransitRouterMulticastDomainId0", ids.associations[0].id)
    ```


    :param Sequence[str] ids: A list of Transit Router Multicast Domain Association IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_id: The ID of the resource associated with the multicast domain.
    :param str resource_type: The type of resource associated with the multicast domain. Valid Value: `VPC`.
    :param str status: The status of the associated resource. Valid Value: `Associated`, `Associating`, `Dissociating`.
    :param str transit_router_attachment_id: The ID of the network instance connection.
    :param str transit_router_multicast_domain_id: The ID of the multicast domain.
    :param str vswitch_id: The ID of the vSwitch.
    """
    ...
