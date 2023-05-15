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
    'GetNetworkPackagesResult',
    'AwaitableGetNetworkPackagesResult',
    'get_network_packages',
    'get_network_packages_output',
]

@pulumi.output_type
class GetNetworkPackagesResult:
    """
    A collection of values returned by getNetworkPackages.
    """
    def __init__(__self__, id=None, ids=None, output_file=None, packages=None, status=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if packages and not isinstance(packages, list):
            raise TypeError("Expected argument 'packages' to be a list")
        pulumi.set(__self__, "packages", packages)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

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
    def packages(self) -> Sequence['outputs.GetNetworkPackagesPackageResult']:
        return pulumi.get(self, "packages")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetNetworkPackagesResult(GetNetworkPackagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkPackagesResult(
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            packages=self.packages,
            status=self.status)


def get_network_packages(ids: Optional[Sequence[str]] = None,
                         output_file: Optional[str] = None,
                         status: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkPackagesResult:
    """
    This data source provides the Ecd Network Packages of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.142.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_simple_office_site = alicloud.eds.SimpleOfficeSite("defaultSimpleOfficeSite",
        cidr_block="172.16.0.0/12",
        desktop_access_type="Internet",
        office_site_name="example_value")
    default_network_package = alicloud.eds.NetworkPackage("defaultNetworkPackage",
        bandwidth=10,
        office_site_id=default_simple_office_site.id)
    default_network_packages = alicloud.eds.get_network_packages_output(ids=[default_network_package.id])
    pulumi.export("ecdNetworkPackageId1", default_network_packages.packages[0].id)
    ```


    :param Sequence[str] ids: A list of Network Package IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:eds/getNetworkPackages:getNetworkPackages', __args__, opts=opts, typ=GetNetworkPackagesResult).value

    return AwaitableGetNetworkPackagesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        packages=__ret__.packages,
        status=__ret__.status)


@_utilities.lift_output_func(get_network_packages)
def get_network_packages_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                status: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetworkPackagesResult]:
    """
    This data source provides the Ecd Network Packages of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.142.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_simple_office_site = alicloud.eds.SimpleOfficeSite("defaultSimpleOfficeSite",
        cidr_block="172.16.0.0/12",
        desktop_access_type="Internet",
        office_site_name="example_value")
    default_network_package = alicloud.eds.NetworkPackage("defaultNetworkPackage",
        bandwidth=10,
        office_site_id=default_simple_office_site.id)
    default_network_packages = alicloud.eds.get_network_packages_output(ids=[default_network_package.id])
    pulumi.export("ecdNetworkPackageId1", default_network_packages.packages[0].id)
    ```


    :param Sequence[str] ids: A list of Network Package IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
    """
    ...
