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
    'GetGatewaySmbUsersResult',
    'AwaitableGetGatewaySmbUsersResult',
    'get_gateway_smb_users',
    'get_gateway_smb_users_output',
]

@pulumi.output_type
class GetGatewaySmbUsersResult:
    """
    A collection of values returned by getGatewaySmbUsers.
    """
    def __init__(__self__, gateway_id=None, id=None, ids=None, name_regex=None, output_file=None, users=None):
        if gateway_id and not isinstance(gateway_id, str):
            raise TypeError("Expected argument 'gateway_id' to be a str")
        pulumi.set(__self__, "gateway_id", gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> str:
        return pulumi.get(self, "gateway_id")

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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetGatewaySmbUsersUserResult']:
        return pulumi.get(self, "users")


class AwaitableGetGatewaySmbUsersResult(GetGatewaySmbUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewaySmbUsersResult(
            gateway_id=self.gateway_id,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            users=self.users)


def get_gateway_smb_users(gateway_id: Optional[str] = None,
                          ids: Optional[Sequence[str]] = None,
                          name_regex: Optional[str] = None,
                          output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewaySmbUsersResult:
    """
    This data source provides the Cloud Storage Gateway Gateway SMB Users of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.142.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_networks = alicloud.vpc.get_networks(name_regex="default-NODELETING")
    default_switches = alicloud.vpc.get_switches(vpc_id=default_networks.ids[0])
    example = alicloud.cloudstoragegateway.StorageBundle("example", storage_bundle_name="example_value")
    default_gateway = alicloud.cloudstoragegateway.Gateway("defaultGateway",
        description="tf-acctestDesalone",
        gateway_class="Standard",
        type="File",
        payment_type="PayAsYouGo",
        vswitch_id=default_switches.ids[0],
        release_after_expiration=False,
        public_network_bandwidth=40,
        storage_bundle_id=example.id,
        location="Cloud",
        gateway_name="example_value")
    default_gateway_smb_user = alicloud.cloudstoragegateway.GatewaySmbUser("defaultGatewaySmbUser",
        username="your_username",
        password="password",
        gateway_id=default_gateway.id)
    ids = alicloud.cloudstoragegateway.get_gateway_smb_users_output(gateway_id=default_gateway.id,
        ids=[default_gateway_smb_user.id])
    pulumi.export("cloudStorageGatewayGatewaySmbUserId1", ids.users[0].id)
    ```


    :param str gateway_id: The Gateway ID.
    :param Sequence[str] ids: A list of Gateway SMB User IDs.
    :param str name_regex: A regex string to filter results by Gateway SMB username.
    """
    __args__ = dict()
    __args__['gatewayId'] = gateway_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cloudstoragegateway/getGatewaySmbUsers:getGatewaySmbUsers', __args__, opts=opts, typ=GetGatewaySmbUsersResult).value

    return AwaitableGetGatewaySmbUsersResult(
        gateway_id=__ret__.gateway_id,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        users=__ret__.users)


@_utilities.lift_output_func(get_gateway_smb_users)
def get_gateway_smb_users_output(gateway_id: Optional[pulumi.Input[str]] = None,
                                 ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewaySmbUsersResult]:
    """
    This data source provides the Cloud Storage Gateway Gateway SMB Users of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.142.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_networks = alicloud.vpc.get_networks(name_regex="default-NODELETING")
    default_switches = alicloud.vpc.get_switches(vpc_id=default_networks.ids[0])
    example = alicloud.cloudstoragegateway.StorageBundle("example", storage_bundle_name="example_value")
    default_gateway = alicloud.cloudstoragegateway.Gateway("defaultGateway",
        description="tf-acctestDesalone",
        gateway_class="Standard",
        type="File",
        payment_type="PayAsYouGo",
        vswitch_id=default_switches.ids[0],
        release_after_expiration=False,
        public_network_bandwidth=40,
        storage_bundle_id=example.id,
        location="Cloud",
        gateway_name="example_value")
    default_gateway_smb_user = alicloud.cloudstoragegateway.GatewaySmbUser("defaultGatewaySmbUser",
        username="your_username",
        password="password",
        gateway_id=default_gateway.id)
    ids = alicloud.cloudstoragegateway.get_gateway_smb_users_output(gateway_id=default_gateway.id,
        ids=[default_gateway_smb_user.id])
    pulumi.export("cloudStorageGatewayGatewaySmbUserId1", ids.users[0].id)
    ```


    :param str gateway_id: The Gateway ID.
    :param Sequence[str] ids: A list of Gateway SMB User IDs.
    :param str name_regex: A regex string to filter results by Gateway SMB username.
    """
    ...
