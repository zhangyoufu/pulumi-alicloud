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
    'GetResourceGroupsResult',
    'AwaitableGetResourceGroupsResult',
    'get_resource_groups',
    'get_resource_groups_output',
]

@pulumi.output_type
class GetResourceGroupsResult:
    """
    A collection of values returned by getResourceGroups.
    """
    def __init__(__self__, enable_details=None, groups=None, id=None, ids=None, name_regex=None, names=None, output_file=None, status=None):
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
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetResourceGroupsGroupResult']:
        """
        A list of resource groups. Each element contains the following attributes:
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
        A list of resource group IDs.
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
        A list of resource group names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the regional resource group.
        """
        return pulumi.get(self, "status")


class AwaitableGetResourceGroupsResult(GetResourceGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceGroupsResult(
            enable_details=self.enable_details,
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            status=self.status)


def get_resource_groups(enable_details: Optional[bool] = None,
                        ids: Optional[Sequence[str]] = None,
                        name_regex: Optional[str] = None,
                        output_file: Optional[str] = None,
                        status: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceGroupsResult:
    """
    This data source provides resource groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.84.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.resourcemanager.get_resource_groups(name_regex="tftest")
    pulumi.export("firstResourceGroupId", example.groups[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to true can output more details.
    :param Sequence[str] ids: A list of resource group IDs.
    :param str name_regex: A regex string to filter results by resource group name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource group. Possible values:`Creating`,`Deleted`,`Deleting`(Available 1.114.0+) `OK` and `PendingDelete`.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:resourcemanager/getResourceGroups:getResourceGroups', __args__, opts=opts, typ=GetResourceGroupsResult).value

    return AwaitableGetResourceGroupsResult(
        enable_details=__ret__.enable_details,
        groups=__ret__.groups,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        status=__ret__.status)


@_utilities.lift_output_func(get_resource_groups)
def get_resource_groups_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                               ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                               output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               status: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResourceGroupsResult]:
    """
    This data source provides resource groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.84.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.resourcemanager.get_resource_groups(name_regex="tftest")
    pulumi.export("firstResourceGroupId", example.groups[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to true can output more details.
    :param Sequence[str] ids: A list of resource group IDs.
    :param str name_regex: A regex string to filter results by resource group name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource group. Possible values:`Creating`,`Deleted`,`Deleting`(Available 1.114.0+) `OK` and `PendingDelete`.
    """
    ...
