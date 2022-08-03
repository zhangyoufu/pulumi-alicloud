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
    'GetBgpPeersResult',
    'AwaitableGetBgpPeersResult',
    'get_bgp_peers',
    'get_bgp_peers_output',
]

@pulumi.output_type
class GetBgpPeersResult:
    """
    A collection of values returned by getBgpPeers.
    """
    def __init__(__self__, bgp_group_id=None, id=None, ids=None, output_file=None, peers=None, router_id=None, status=None):
        if bgp_group_id and not isinstance(bgp_group_id, str):
            raise TypeError("Expected argument 'bgp_group_id' to be a str")
        pulumi.set(__self__, "bgp_group_id", bgp_group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if peers and not isinstance(peers, list):
            raise TypeError("Expected argument 'peers' to be a list")
        pulumi.set(__self__, "peers", peers)
        if router_id and not isinstance(router_id, str):
            raise TypeError("Expected argument 'router_id' to be a str")
        pulumi.set(__self__, "router_id", router_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="bgpGroupId")
    def bgp_group_id(self) -> Optional[str]:
        return pulumi.get(self, "bgp_group_id")

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
    def peers(self) -> Sequence['outputs.GetBgpPeersPeerResult']:
        return pulumi.get(self, "peers")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> Optional[str]:
        return pulumi.get(self, "router_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetBgpPeersResult(GetBgpPeersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBgpPeersResult(
            bgp_group_id=self.bgp_group_id,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            peers=self.peers,
            router_id=self.router_id,
            status=self.status)


def get_bgp_peers(bgp_group_id: Optional[str] = None,
                  ids: Optional[Sequence[str]] = None,
                  output_file: Optional[str] = None,
                  router_id: Optional[str] = None,
                  status: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBgpPeersResult:
    """
    This data source provides the Vpc Bgp Peers of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.153.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.vpc.get_bgp_peers(ids=[
        "example_value-1",
        "example_value-2",
    ])
    pulumi.export("vpcBgpPeerId1", ids.peers[0].id)
    bgp_group_id = alicloud.vpc.get_bgp_peers(bgp_group_id="example_value")
    pulumi.export("vpcBgpPeerId2", bgp_group_id.peers[0].id)
    router_id = alicloud.vpc.get_bgp_peers(router_id="example_value")
    pulumi.export("vpcBgpPeerId3", router_id.peers[0].id)
    status = alicloud.vpc.get_bgp_peers(status="Available")
    pulumi.export("vpcBgpPeerId4", status.peers[0].id)
    ```


    :param str bgp_group_id: The ID of the BGP group.
    :param Sequence[str] ids: A list of Bgp Peer IDs.
    :param str router_id: The ID of the router.
    :param str status: The status of the BGP peer.
    """
    __args__ = dict()
    __args__['bgpGroupId'] = bgp_group_id
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['routerId'] = router_id
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:vpc/getBgpPeers:getBgpPeers', __args__, opts=opts, typ=GetBgpPeersResult).value

    return AwaitableGetBgpPeersResult(
        bgp_group_id=__ret__.bgp_group_id,
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        peers=__ret__.peers,
        router_id=__ret__.router_id,
        status=__ret__.status)


@_utilities.lift_output_func(get_bgp_peers)
def get_bgp_peers_output(bgp_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         router_id: Optional[pulumi.Input[Optional[str]]] = None,
                         status: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBgpPeersResult]:
    """
    This data source provides the Vpc Bgp Peers of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.153.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.vpc.get_bgp_peers(ids=[
        "example_value-1",
        "example_value-2",
    ])
    pulumi.export("vpcBgpPeerId1", ids.peers[0].id)
    bgp_group_id = alicloud.vpc.get_bgp_peers(bgp_group_id="example_value")
    pulumi.export("vpcBgpPeerId2", bgp_group_id.peers[0].id)
    router_id = alicloud.vpc.get_bgp_peers(router_id="example_value")
    pulumi.export("vpcBgpPeerId3", router_id.peers[0].id)
    status = alicloud.vpc.get_bgp_peers(status="Available")
    pulumi.export("vpcBgpPeerId4", status.peers[0].id)
    ```


    :param str bgp_group_id: The ID of the BGP group.
    :param Sequence[str] ids: A list of Bgp Peer IDs.
    :param str router_id: The ID of the router.
    :param str status: The status of the BGP peer.
    """
    ...
