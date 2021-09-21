# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetDeviceGroupsResult',
    'AwaitableGetDeviceGroupsResult',
    'get_device_groups',
]

@pulumi.output_type
class GetDeviceGroupsResult:
    """
    A collection of values returned by getDeviceGroups.
    """
    def __init__(__self__, enable_details=None, group_name=None, groups=None, id=None, ids=None, iot_instance_id=None, name_regex=None, output_file=None, super_group_id=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        pulumi.set(__self__, "group_name", group_name)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if iot_instance_id and not isinstance(iot_instance_id, str):
            raise TypeError("Expected argument 'iot_instance_id' to be a str")
        pulumi.set(__self__, "iot_instance_id", iot_instance_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if super_group_id and not isinstance(super_group_id, str):
            raise TypeError("Expected argument 'super_group_id' to be a str")
        pulumi.set(__self__, "super_group_id", super_group_id)

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[str]:
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetDeviceGroupsGroupResult']:
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
    @pulumi.getter(name="iotInstanceId")
    def iot_instance_id(self) -> Optional[str]:
        return pulumi.get(self, "iot_instance_id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="superGroupId")
    def super_group_id(self) -> Optional[str]:
        return pulumi.get(self, "super_group_id")


class AwaitableGetDeviceGroupsResult(GetDeviceGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeviceGroupsResult(
            enable_details=self.enable_details,
            group_name=self.group_name,
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            iot_instance_id=self.iot_instance_id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            super_group_id=self.super_group_id)


def get_device_groups(enable_details: Optional[bool] = None,
                      group_name: Optional[str] = None,
                      ids: Optional[Sequence[str]] = None,
                      iot_instance_id: Optional[str] = None,
                      name_regex: Optional[str] = None,
                      output_file: Optional[str] = None,
                      super_group_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeviceGroupsResult:
    """
    This data source provides the Iot Device Groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.134.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.iot.get_device_groups()
    pulumi.export("iotDeviceGroupId1", ids.groups[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param str group_name: The GroupName of the device group.
    :param Sequence[str] ids: A list of device group IDs.
    :param str iot_instance_id: The id of the Iot Instance.
    :param str name_regex: A regex string to filter CEN instances by name.
    :param str super_group_id: The id of the SuperGroup.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['groupName'] = group_name
    __args__['ids'] = ids
    __args__['iotInstanceId'] = iot_instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['superGroupId'] = super_group_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:iot/getDeviceGroups:getDeviceGroups', __args__, opts=opts, typ=GetDeviceGroupsResult).value

    return AwaitableGetDeviceGroupsResult(
        enable_details=__ret__.enable_details,
        group_name=__ret__.group_name,
        groups=__ret__.groups,
        id=__ret__.id,
        ids=__ret__.ids,
        iot_instance_id=__ret__.iot_instance_id,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        super_group_id=__ret__.super_group_id)
