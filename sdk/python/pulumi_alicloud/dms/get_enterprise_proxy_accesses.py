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
    'GetEnterpriseProxyAccessesResult',
    'AwaitableGetEnterpriseProxyAccessesResult',
    'get_enterprise_proxy_accesses',
    'get_enterprise_proxy_accesses_output',
]

@pulumi.output_type
class GetEnterpriseProxyAccessesResult:
    """
    A collection of values returned by getEnterpriseProxyAccesses.
    """
    def __init__(__self__, accesses=None, enable_details=None, id=None, ids=None, output_file=None, proxy_id=None):
        if accesses and not isinstance(accesses, list):
            raise TypeError("Expected argument 'accesses' to be a list")
        pulumi.set(__self__, "accesses", accesses)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if proxy_id and not isinstance(proxy_id, str):
            raise TypeError("Expected argument 'proxy_id' to be a str")
        pulumi.set(__self__, "proxy_id", proxy_id)

    @property
    @pulumi.getter
    def accesses(self) -> Sequence['outputs.GetEnterpriseProxyAccessesAccessResult']:
        """
        A list of Proxy Access Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "accesses")

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
        """
        A list of Proxy Access IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> str:
        """
        The ID of the security agent.
        """
        return pulumi.get(self, "proxy_id")


class AwaitableGetEnterpriseProxyAccessesResult(GetEnterpriseProxyAccessesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnterpriseProxyAccessesResult(
            accesses=self.accesses,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            proxy_id=self.proxy_id)


def get_enterprise_proxy_accesses(enable_details: Optional[bool] = None,
                                  ids: Optional[Sequence[str]] = None,
                                  output_file: Optional[str] = None,
                                  proxy_id: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnterpriseProxyAccessesResult:
    """
    This data source provides DMS Enterprise Proxy Access available to the user.[What is Proxy Access](https://next.api.alibabacloud.com/document/dms-enterprise/2018-11-01/CreateProxyAccess)

    > **NOTE:** Available since v1.195.0.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_enterprise_proxy_access = alicloud.dms.EnterpriseProxyAccess("defaultEnterpriseProxyAccess",
        indep_password="PASSWORD-DEMO",
        proxy_id="1881",
        indep_account="dmstest",
        user_id="104442")
    default_enterprise_proxy_accesses = alicloud.dms.get_enterprise_proxy_accesses_output(ids=[default_enterprise_proxy_access.id],
        proxy_id="1881")
    pulumi.export("alicloudDmsProxyAccesesExampleId", default_enterprise_proxy_accesses.accesses[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Proxy Access IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str proxy_id: The ID of the security agent.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['proxyId'] = proxy_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:dms/getEnterpriseProxyAccesses:getEnterpriseProxyAccesses', __args__, opts=opts, typ=GetEnterpriseProxyAccessesResult).value

    return AwaitableGetEnterpriseProxyAccessesResult(
        accesses=pulumi.get(__ret__, 'accesses'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        output_file=pulumi.get(__ret__, 'output_file'),
        proxy_id=pulumi.get(__ret__, 'proxy_id'))


@_utilities.lift_output_func(get_enterprise_proxy_accesses)
def get_enterprise_proxy_accesses_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                         proxy_id: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnterpriseProxyAccessesResult]:
    """
    This data source provides DMS Enterprise Proxy Access available to the user.[What is Proxy Access](https://next.api.alibabacloud.com/document/dms-enterprise/2018-11-01/CreateProxyAccess)

    > **NOTE:** Available since v1.195.0.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_enterprise_proxy_access = alicloud.dms.EnterpriseProxyAccess("defaultEnterpriseProxyAccess",
        indep_password="PASSWORD-DEMO",
        proxy_id="1881",
        indep_account="dmstest",
        user_id="104442")
    default_enterprise_proxy_accesses = alicloud.dms.get_enterprise_proxy_accesses_output(ids=[default_enterprise_proxy_access.id],
        proxy_id="1881")
    pulumi.export("alicloudDmsProxyAccesesExampleId", default_enterprise_proxy_accesses.accesses[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Proxy Access IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str proxy_id: The ID of the security agent.
    """
    ...
