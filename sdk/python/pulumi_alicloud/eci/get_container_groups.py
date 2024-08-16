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
    'GetContainerGroupsResult',
    'AwaitableGetContainerGroupsResult',
    'get_container_groups',
    'get_container_groups_output',
]

@pulumi.output_type
class GetContainerGroupsResult:
    """
    A collection of values returned by getContainerGroups.
    """
    def __init__(__self__, container_group_name=None, enable_details=None, groups=None, id=None, ids=None, limit=None, name_regex=None, names=None, output_file=None, resource_group_id=None, status=None, tags=None, vswitch_id=None, with_event=None, zone_id=None):
        if container_group_name and not isinstance(container_group_name, str):
            raise TypeError("Expected argument 'container_group_name' to be a str")
        pulumi.set(__self__, "container_group_name", container_group_name)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if limit and not isinstance(limit, int):
            raise TypeError("Expected argument 'limit' to be a int")
        pulumi.set(__self__, "limit", limit)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vswitch_id and not isinstance(vswitch_id, str):
            raise TypeError("Expected argument 'vswitch_id' to be a str")
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        if with_event and not isinstance(with_event, bool):
            raise TypeError("Expected argument 'with_event' to be a bool")
        pulumi.set(__self__, "with_event", with_event)
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="containerGroupName")
    def container_group_name(self) -> Optional[str]:
        return pulumi.get(self, "container_group_name")

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetContainerGroupsGroupResult']:
        return pulumi.get(self, "groups")

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
    @pulumi.getter
    def limit(self) -> Optional[int]:
        return pulumi.get(self, "limit")

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
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[str]:
        return pulumi.get(self, "vswitch_id")

    @property
    @pulumi.getter(name="withEvent")
    def with_event(self) -> Optional[bool]:
        return pulumi.get(self, "with_event")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[str]:
        return pulumi.get(self, "zone_id")


class AwaitableGetContainerGroupsResult(GetContainerGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContainerGroupsResult(
            container_group_name=self.container_group_name,
            enable_details=self.enable_details,
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            limit=self.limit,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            resource_group_id=self.resource_group_id,
            status=self.status,
            tags=self.tags,
            vswitch_id=self.vswitch_id,
            with_event=self.with_event,
            zone_id=self.zone_id)


def get_container_groups(container_group_name: Optional[str] = None,
                         enable_details: Optional[bool] = None,
                         ids: Optional[Sequence[str]] = None,
                         limit: Optional[int] = None,
                         name_regex: Optional[str] = None,
                         output_file: Optional[str] = None,
                         resource_group_id: Optional[str] = None,
                         status: Optional[str] = None,
                         tags: Optional[Mapping[str, str]] = None,
                         vswitch_id: Optional[str] = None,
                         with_event: Optional[bool] = None,
                         zone_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContainerGroupsResult:
    """
    This data source provides the Eci Container Groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.111.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.eci.get_container_groups(ids=["example_value"])
    pulumi.export("firstEciContainerGroupId", example.groups[0].id)
    ```


    :param str container_group_name: The name of ContainerGroup.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Container Group IDs.
    :param int limit: The maximum number of resources returned in the response. Default value is `20`. Maximum value: `20`. The number of returned results is no greater than the specified number.
    :param str name_regex: A regex string to filter results by Container Group name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The ID of the resource group to which the container group belongs. If you have not specified a resource group for the container group, it is added to the default resource group.
    :param str status: The status list. For more information, see the description of ContainerGroup arrays.
    :param str vswitch_id: The ID of the vSwitch. Currently, container groups can only be deployed in VPC networks.
    :param str zone_id: The ID of the zone where you want to deploy the container group. If no value is specified, the system assigns a zone to the container group. By default, no value is specified.
    """
    __args__ = dict()
    __args__['containerGroupName'] = container_group_name
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['limit'] = limit
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['resourceGroupId'] = resource_group_id
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['vswitchId'] = vswitch_id
    __args__['withEvent'] = with_event
    __args__['zoneId'] = zone_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:eci/getContainerGroups:getContainerGroups', __args__, opts=opts, typ=GetContainerGroupsResult).value

    return AwaitableGetContainerGroupsResult(
        container_group_name=pulumi.get(__ret__, 'container_group_name'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        groups=pulumi.get(__ret__, 'groups'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        limit=pulumi.get(__ret__, 'limit'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        resource_group_id=pulumi.get(__ret__, 'resource_group_id'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        vswitch_id=pulumi.get(__ret__, 'vswitch_id'),
        with_event=pulumi.get(__ret__, 'with_event'),
        zone_id=pulumi.get(__ret__, 'zone_id'))


@_utilities.lift_output_func(get_container_groups)
def get_container_groups_output(container_group_name: Optional[pulumi.Input[Optional[str]]] = None,
                                enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                                ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                limit: Optional[pulumi.Input[Optional[int]]] = None,
                                name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                status: Optional[pulumi.Input[Optional[str]]] = None,
                                tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                                vswitch_id: Optional[pulumi.Input[Optional[str]]] = None,
                                with_event: Optional[pulumi.Input[Optional[bool]]] = None,
                                zone_id: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetContainerGroupsResult]:
    """
    This data source provides the Eci Container Groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.111.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.eci.get_container_groups(ids=["example_value"])
    pulumi.export("firstEciContainerGroupId", example.groups[0].id)
    ```


    :param str container_group_name: The name of ContainerGroup.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Container Group IDs.
    :param int limit: The maximum number of resources returned in the response. Default value is `20`. Maximum value: `20`. The number of returned results is no greater than the specified number.
    :param str name_regex: A regex string to filter results by Container Group name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The ID of the resource group to which the container group belongs. If you have not specified a resource group for the container group, it is added to the default resource group.
    :param str status: The status list. For more information, see the description of ContainerGroup arrays.
    :param str vswitch_id: The ID of the vSwitch. Currently, container groups can only be deployed in VPC networks.
    :param str zone_id: The ID of the zone where you want to deploy the container group. If no value is specified, the system assigns a zone to the container group. By default, no value is specified.
    """
    ...
