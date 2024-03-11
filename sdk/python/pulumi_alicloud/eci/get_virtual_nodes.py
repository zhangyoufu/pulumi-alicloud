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
    'GetVirtualNodesResult',
    'AwaitableGetVirtualNodesResult',
    'get_virtual_nodes',
    'get_virtual_nodes_output',
]

@pulumi.output_type
class GetVirtualNodesResult:
    """
    A collection of values returned by getVirtualNodes.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, names=None, nodes=None, output_file=None, resource_group_id=None, security_group_id=None, status=None, tags=None, virtual_node_name=None, vswitch_id=None):
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
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        pulumi.set(__self__, "nodes", nodes)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if virtual_node_name and not isinstance(virtual_node_name, str):
            raise TypeError("Expected argument 'virtual_node_name' to be a str")
        pulumi.set(__self__, "virtual_node_name", virtual_node_name)
        if vswitch_id and not isinstance(vswitch_id, str):
            raise TypeError("Expected argument 'vswitch_id' to be a str")
        pulumi.set(__self__, "vswitch_id", vswitch_id)

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
    @pulumi.getter
    def nodes(self) -> Sequence['outputs.GetVirtualNodesNodeResult']:
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[str]:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualNodeName")
    def virtual_node_name(self) -> Optional[str]:
        return pulumi.get(self, "virtual_node_name")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[str]:
        return pulumi.get(self, "vswitch_id")


class AwaitableGetVirtualNodesResult(GetVirtualNodesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualNodesResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            nodes=self.nodes,
            output_file=self.output_file,
            resource_group_id=self.resource_group_id,
            security_group_id=self.security_group_id,
            status=self.status,
            tags=self.tags,
            virtual_node_name=self.virtual_node_name,
            vswitch_id=self.vswitch_id)


def get_virtual_nodes(ids: Optional[Sequence[str]] = None,
                      name_regex: Optional[str] = None,
                      output_file: Optional[str] = None,
                      resource_group_id: Optional[str] = None,
                      security_group_id: Optional[str] = None,
                      status: Optional[str] = None,
                      tags: Optional[Mapping[str, Any]] = None,
                      virtual_node_name: Optional[str] = None,
                      vswitch_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualNodesResult:
    """
    This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.145.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.eci.get_virtual_nodes(ids=[
        "example_value-1",
        "example_value-2",
    ])
    pulumi.export("eciVirtualNodeId1", ids.nodes[0].id)
    name_regex = alicloud.eci.get_virtual_nodes(name_regex="^my-VirtualNode")
    pulumi.export("eciVirtualNodeId2", name_regex.nodes[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Virtual Node IDs.
    :param str name_regex: A regex string to filter results by Virtual Node name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The resource group ID.
    :param str security_group_id: The security group ID.
    :param str status: The Status of the virtual node.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    :param str virtual_node_name: The name of the virtual node.
    :param str vswitch_id: The vswitch id.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['resourceGroupId'] = resource_group_id
    __args__['securityGroupId'] = security_group_id
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['virtualNodeName'] = virtual_node_name
    __args__['vswitchId'] = vswitch_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:eci/getVirtualNodes:getVirtualNodes', __args__, opts=opts, typ=GetVirtualNodesResult).value

    return AwaitableGetVirtualNodesResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        nodes=pulumi.get(__ret__, 'nodes'),
        output_file=pulumi.get(__ret__, 'output_file'),
        resource_group_id=pulumi.get(__ret__, 'resource_group_id'),
        security_group_id=pulumi.get(__ret__, 'security_group_id'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        virtual_node_name=pulumi.get(__ret__, 'virtual_node_name'),
        vswitch_id=pulumi.get(__ret__, 'vswitch_id'))


@_utilities.lift_output_func(get_virtual_nodes)
def get_virtual_nodes_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                             output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                             security_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                             status: Optional[pulumi.Input[Optional[str]]] = None,
                             tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                             virtual_node_name: Optional[pulumi.Input[Optional[str]]] = None,
                             vswitch_id: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVirtualNodesResult]:
    """
    This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.145.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.eci.get_virtual_nodes(ids=[
        "example_value-1",
        "example_value-2",
    ])
    pulumi.export("eciVirtualNodeId1", ids.nodes[0].id)
    name_regex = alicloud.eci.get_virtual_nodes(name_regex="^my-VirtualNode")
    pulumi.export("eciVirtualNodeId2", name_regex.nodes[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Virtual Node IDs.
    :param str name_regex: A regex string to filter results by Virtual Node name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The resource group ID.
    :param str security_group_id: The security group ID.
    :param str status: The Status of the virtual node.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    :param str virtual_node_name: The name of the virtual node.
    :param str vswitch_id: The vswitch id.
    """
    ...
