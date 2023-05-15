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
    'GetTransitRouterMulticastDomainSourcesResult',
    'AwaitableGetTransitRouterMulticastDomainSourcesResult',
    'get_transit_router_multicast_domain_sources',
    'get_transit_router_multicast_domain_sources_output',
]

@pulumi.output_type
class GetTransitRouterMulticastDomainSourcesResult:
    """
    A collection of values returned by getTransitRouterMulticastDomainSources.
    """
    def __init__(__self__, id=None, ids=None, output_file=None, sources=None, transit_router_multicast_domain_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if sources and not isinstance(sources, list):
            raise TypeError("Expected argument 'sources' to be a list")
        pulumi.set(__self__, "sources", sources)
        if transit_router_multicast_domain_id and not isinstance(transit_router_multicast_domain_id, str):
            raise TypeError("Expected argument 'transit_router_multicast_domain_id' to be a str")
        pulumi.set(__self__, "transit_router_multicast_domain_id", transit_router_multicast_domain_id)

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
    @pulumi.getter
    def sources(self) -> Sequence['outputs.GetTransitRouterMulticastDomainSourcesSourceResult']:
        """
        A list of Transit Router Multicast Domain Source Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "sources")

    @property
    @pulumi.getter(name="transitRouterMulticastDomainId")
    def transit_router_multicast_domain_id(self) -> str:
        """
        The ID of the multicast domain to which the multicast source belongs.
        """
        return pulumi.get(self, "transit_router_multicast_domain_id")


class AwaitableGetTransitRouterMulticastDomainSourcesResult(GetTransitRouterMulticastDomainSourcesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransitRouterMulticastDomainSourcesResult(
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            sources=self.sources,
            transit_router_multicast_domain_id=self.transit_router_multicast_domain_id)


def get_transit_router_multicast_domain_sources(ids: Optional[Sequence[str]] = None,
                                                output_file: Optional[str] = None,
                                                transit_router_multicast_domain_id: Optional[str] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransitRouterMulticastDomainSourcesResult:
    """
    This data source provides Cen Transit Router Multicast Domain Source available to the user.[What is Transit Router Multicast Domain Source](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-registertransitroutermulticastgroupsources)

    > **NOTE:** Available in 1.195.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.cen.get_transit_router_multicast_domain_sources(transit_router_multicast_domain_id="tr-mcast-domain-xxxxxx")
    pulumi.export("alicloudCenTransitRouterMulticastDomainSourceExampleId", default.sources[0].id)
    ```


    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str transit_router_multicast_domain_id: The ID of the multicast domain to which the multicast source belongs.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['transitRouterMulticastDomainId'] = transit_router_multicast_domain_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getTransitRouterMulticastDomainSources:getTransitRouterMulticastDomainSources', __args__, opts=opts, typ=GetTransitRouterMulticastDomainSourcesResult).value

    return AwaitableGetTransitRouterMulticastDomainSourcesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        sources=__ret__.sources,
        transit_router_multicast_domain_id=__ret__.transit_router_multicast_domain_id)


@_utilities.lift_output_func(get_transit_router_multicast_domain_sources)
def get_transit_router_multicast_domain_sources_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                                       transit_router_multicast_domain_id: Optional[pulumi.Input[str]] = None,
                                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransitRouterMulticastDomainSourcesResult]:
    """
    This data source provides Cen Transit Router Multicast Domain Source available to the user.[What is Transit Router Multicast Domain Source](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-registertransitroutermulticastgroupsources)

    > **NOTE:** Available in 1.195.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.cen.get_transit_router_multicast_domain_sources(transit_router_multicast_domain_id="tr-mcast-domain-xxxxxx")
    pulumi.export("alicloudCenTransitRouterMulticastDomainSourceExampleId", default.sources[0].id)
    ```


    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str transit_router_multicast_domain_id: The ID of the multicast domain to which the multicast source belongs.
    """
    ...
