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
    'GetUsersResult',
    'AwaitableGetUsersResult',
    'get_users',
    'get_users_output',
]

@pulumi.output_type
class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, directory_id=None, enable_details=None, id=None, ids=None, name_regex=None, names=None, output_file=None, provision_type=None, status=None, users=None):
        if directory_id and not isinstance(directory_id, str):
            raise TypeError("Expected argument 'directory_id' to be a str")
        pulumi.set(__self__, "directory_id", directory_id)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
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
        if provision_type and not isinstance(provision_type, str):
            raise TypeError("Expected argument 'provision_type' to be a str")
        pulumi.set(__self__, "provision_type", provision_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> str:
        return pulumi.get(self, "directory_id")

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

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
    @pulumi.getter(name="provisionType")
    def provision_type(self) -> Optional[str]:
        return pulumi.get(self, "provision_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetUsersUserResult']:
        return pulumi.get(self, "users")


class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            directory_id=self.directory_id,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            provision_type=self.provision_type,
            status=self.status,
            users=self.users)


def get_users(directory_id: Optional[str] = None,
              enable_details: Optional[bool] = None,
              ids: Optional[Sequence[str]] = None,
              name_regex: Optional[str] = None,
              output_file: Optional[str] = None,
              provision_type: Optional[str] = None,
              status: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUsersResult:
    """
    This data source provides the Cloud Sso Users of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.140.0+.

    > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cloudsso.get_users(directory_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("cloudSsoUserId1", ids.users[0].id)
    name_regex = alicloud.cloudsso.get_users(directory_id="example_value",
        name_regex="^my-User")
    pulumi.export("cloudSsoUserId2", name_regex.users[0].id)
    provision_type = alicloud.cloudsso.get_users(directory_id="example_value",
        ids=["example_value-1"],
        provision_type="Manual")
    pulumi.export("cloudSsoUserId3", provision_type.users[0].id)
    status = alicloud.cloudsso.get_users(directory_id="example_value",
        ids=["example_value-1"],
        status="Enabled")
    pulumi.export("cloudSsoUserId4", status.users[0].id)
    ```


    :param str directory_id: The ID of the Directory.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of User IDs.
    :param str name_regex: A regex string to filter results by User name.
    :param str provision_type: ProvisionType.
    :param str status: User status. Valid values: `Enabled` and `Disabled`.
    """
    __args__ = dict()
    __args__['directoryId'] = directory_id
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['provisionType'] = provision_type
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cloudsso/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult).value

    return AwaitableGetUsersResult(
        directory_id=__ret__.directory_id,
        enable_details=__ret__.enable_details,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        provision_type=__ret__.provision_type,
        status=__ret__.status,
        users=__ret__.users)


@_utilities.lift_output_func(get_users)
def get_users_output(directory_id: Optional[pulumi.Input[str]] = None,
                     enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                     ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     provision_type: Optional[pulumi.Input[Optional[str]]] = None,
                     status: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUsersResult]:
    """
    This data source provides the Cloud Sso Users of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.140.0+.

    > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cloudsso.get_users(directory_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("cloudSsoUserId1", ids.users[0].id)
    name_regex = alicloud.cloudsso.get_users(directory_id="example_value",
        name_regex="^my-User")
    pulumi.export("cloudSsoUserId2", name_regex.users[0].id)
    provision_type = alicloud.cloudsso.get_users(directory_id="example_value",
        ids=["example_value-1"],
        provision_type="Manual")
    pulumi.export("cloudSsoUserId3", provision_type.users[0].id)
    status = alicloud.cloudsso.get_users(directory_id="example_value",
        ids=["example_value-1"],
        status="Enabled")
    pulumi.export("cloudSsoUserId4", status.users[0].id)
    ```


    :param str directory_id: The ID of the Directory.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of User IDs.
    :param str name_regex: A regex string to filter results by User name.
    :param str provision_type: ProvisionType.
    :param str status: User status. Valid values: `Enabled` and `Disabled`.
    """
    ...
