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
    'GetInstancesResult',
    'AwaitableGetInstancesResult',
    'get_instances',
    'get_instances_output',
]

@pulumi.output_type
class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, domain_type=None, id=None, ids=None, instances=None, lang=None, output_file=None, user_client_ip=None):
        if domain_type and not isinstance(domain_type, str):
            raise TypeError("Expected argument 'domain_type' to be a str")
        pulumi.set(__self__, "domain_type", domain_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if lang and not isinstance(lang, str):
            raise TypeError("Expected argument 'lang' to be a str")
        pulumi.set(__self__, "lang", lang)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if user_client_ip and not isinstance(user_client_ip, str):
            raise TypeError("Expected argument 'user_client_ip' to be a str")
        pulumi.set(__self__, "user_client_ip", user_client_ip)

    @property
    @pulumi.getter(name="domainType")
    def domain_type(self) -> Optional[str]:
        return pulumi.get(self, "domain_type")

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
        A list of instance IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetInstancesInstanceResult']:
        """
        A list of instances. Each element contains the following attributes:
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def lang(self) -> Optional[str]:
        return pulumi.get(self, "lang")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="userClientIp")
    def user_client_ip(self) -> Optional[str]:
        return pulumi.get(self, "user_client_ip")


class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            domain_type=self.domain_type,
            id=self.id,
            ids=self.ids,
            instances=self.instances,
            lang=self.lang,
            output_file=self.output_file,
            user_client_ip=self.user_client_ip)


def get_instances(domain_type: Optional[str] = None,
                  ids: Optional[Sequence[str]] = None,
                  lang: Optional[str] = None,
                  output_file: Optional[str] = None,
                  user_client_ip: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancesResult:
    """
    > **DEPRECATED:**  This resource has been renamed to dns_get_alidns_instances from version 1.95.0.

    This data source provides a list of DNS instances in an Alibaba Cloud account according to the specified filters.

    > **NOTE:**  Available in 1.84.0+.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.dns.get_instances(ids=["dns-cn-oew1npk****"])
    pulumi.export("firstInstanceId", example.instances[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of instance IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['domainType'] = domain_type
    __args__['ids'] = ids
    __args__['lang'] = lang
    __args__['outputFile'] = output_file
    __args__['userClientIp'] = user_client_ip
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:dns/getInstances:getInstances', __args__, opts=opts, typ=GetInstancesResult).value

    return AwaitableGetInstancesResult(
        domain_type=pulumi.get(__ret__, 'domain_type'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        instances=pulumi.get(__ret__, 'instances'),
        lang=pulumi.get(__ret__, 'lang'),
        output_file=pulumi.get(__ret__, 'output_file'),
        user_client_ip=pulumi.get(__ret__, 'user_client_ip'))


@_utilities.lift_output_func(get_instances)
def get_instances_output(domain_type: Optional[pulumi.Input[Optional[str]]] = None,
                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         lang: Optional[pulumi.Input[Optional[str]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         user_client_ip: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstancesResult]:
    """
    > **DEPRECATED:**  This resource has been renamed to dns_get_alidns_instances from version 1.95.0.

    This data source provides a list of DNS instances in an Alibaba Cloud account according to the specified filters.

    > **NOTE:**  Available in 1.84.0+.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.dns.get_instances(ids=["dns-cn-oew1npk****"])
    pulumi.export("firstInstanceId", example.instances[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of instance IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
