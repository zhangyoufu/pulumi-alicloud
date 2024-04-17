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
    'GetSecurityGroupsResult',
    'AwaitableGetSecurityGroupsResult',
    'get_security_groups',
    'get_security_groups_output',
]

@pulumi.output_type
class GetSecurityGroupsResult:
    """
    A collection of values returned by getSecurityGroups.
    """
    def __init__(__self__, enable_details=None, groups=None, id=None, ids=None, name_regex=None, names=None, output_file=None, page_number=None, page_size=None, resource_group_id=None, tags=None, total_count=None, vpc_id=None):
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
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetSecurityGroupsGroupResult']:
        """
        A list of Security Groups. Each element contains the following attributes:
        """
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
        """
        A list of Security Group IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of Security Group names.
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

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        """
        The Id of resource group which the security_group belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        """
        A map of tags assigned to the ECS instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The ID of the VPC that owns the security group.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetSecurityGroupsResult(GetSecurityGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityGroupsResult(
            enable_details=self.enable_details,
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size,
            resource_group_id=self.resource_group_id,
            tags=self.tags,
            total_count=self.total_count,
            vpc_id=self.vpc_id)


def get_security_groups(enable_details: Optional[bool] = None,
                        ids: Optional[Sequence[str]] = None,
                        name_regex: Optional[str] = None,
                        output_file: Optional[str] = None,
                        page_number: Optional[int] = None,
                        page_size: Optional[int] = None,
                        resource_group_id: Optional[str] = None,
                        tags: Optional[Mapping[str, Any]] = None,
                        vpc_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityGroupsResult:
    """
    This data source provides a list of Security Groups in an Alibaba Cloud account according to the specified filters.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    # Filter security groups and print the results into a file
    sec_groups_ds = alicloud.ecs.get_security_groups(name_regex="^web-",
        output_file="web_access.json")
    # In conjunction with a VPC
    primary_vpc_ds = alicloud.vpc.Network("primary_vpc_ds")
    primary_sec_groups_ds = alicloud.ecs.get_security_groups_output(vpc_id=primary_vpc_ds.id)
    pulumi.export("firstGroupId", primary_sec_groups_ds.groups[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Security Group IDs.
    :param str name_regex: A regex string to filter the resulting security groups by their names.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The Id of resource group which the security_group belongs.
    :param Mapping[str, Any] tags: A map of tags assigned to the ECS instances. It must be in the format:
           <!--Start PulumiCodeChooser -->
           ```python
           import pulumi
           import pulumi_alicloud as alicloud
           
           tagged_security_groups = alicloud.ecs.get_security_groups(tags={
               "tagKey1": "tagValue1",
               "tagKey2": "tagValue2",
           })
           ```
           <!--End PulumiCodeChooser -->
    :param str vpc_id: Used to retrieve security groups that belong to the specified VPC ID.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    __args__['resourceGroupId'] = resource_group_id
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getSecurityGroups:getSecurityGroups', __args__, opts=opts, typ=GetSecurityGroupsResult).value

    return AwaitableGetSecurityGroupsResult(
        enable_details=pulumi.get(__ret__, 'enable_details'),
        groups=pulumi.get(__ret__, 'groups'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        page_number=pulumi.get(__ret__, 'page_number'),
        page_size=pulumi.get(__ret__, 'page_size'),
        resource_group_id=pulumi.get(__ret__, 'resource_group_id'),
        tags=pulumi.get(__ret__, 'tags'),
        total_count=pulumi.get(__ret__, 'total_count'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))


@_utilities.lift_output_func(get_security_groups)
def get_security_groups_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                               ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                               output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               page_number: Optional[pulumi.Input[Optional[int]]] = None,
                               page_size: Optional[pulumi.Input[Optional[int]]] = None,
                               resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                               tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                               vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityGroupsResult]:
    """
    This data source provides a list of Security Groups in an Alibaba Cloud account according to the specified filters.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    # Filter security groups and print the results into a file
    sec_groups_ds = alicloud.ecs.get_security_groups(name_regex="^web-",
        output_file="web_access.json")
    # In conjunction with a VPC
    primary_vpc_ds = alicloud.vpc.Network("primary_vpc_ds")
    primary_sec_groups_ds = alicloud.ecs.get_security_groups_output(vpc_id=primary_vpc_ds.id)
    pulumi.export("firstGroupId", primary_sec_groups_ds.groups[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Security Group IDs.
    :param str name_regex: A regex string to filter the resulting security groups by their names.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The Id of resource group which the security_group belongs.
    :param Mapping[str, Any] tags: A map of tags assigned to the ECS instances. It must be in the format:
           <!--Start PulumiCodeChooser -->
           ```python
           import pulumi
           import pulumi_alicloud as alicloud
           
           tagged_security_groups = alicloud.ecs.get_security_groups(tags={
               "tagKey1": "tagValue1",
               "tagKey2": "tagValue2",
           })
           ```
           <!--End PulumiCodeChooser -->
    :param str vpc_id: Used to retrieve security groups that belong to the specified VPC ID.
    """
    ...
