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
    'GetWebLockConfigsResult',
    'AwaitableGetWebLockConfigsResult',
    'get_web_lock_configs',
    'get_web_lock_configs_output',
]

@pulumi.output_type
class GetWebLockConfigsResult:
    """
    A collection of values returned by getWebLockConfigs.
    """
    def __init__(__self__, configs=None, id=None, ids=None, lang=None, output_file=None, page_number=None, page_size=None, remark=None, source_ip=None, status=None):
        if configs and not isinstance(configs, list):
            raise TypeError("Expected argument 'configs' to be a list")
        pulumi.set(__self__, "configs", configs)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if lang and not isinstance(lang, str):
            raise TypeError("Expected argument 'lang' to be a str")
        pulumi.set(__self__, "lang", lang)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if page_number and not isinstance(page_number, int):
            raise TypeError("Expected argument 'page_number' to be a int")
        pulumi.set(__self__, "page_number", page_number)
        if page_size and not isinstance(page_size, int):
            raise TypeError("Expected argument 'page_size' to be a int")
        pulumi.set(__self__, "page_size", page_size)
        if remark and not isinstance(remark, str):
            raise TypeError("Expected argument 'remark' to be a str")
        pulumi.set(__self__, "remark", remark)
        if source_ip and not isinstance(source_ip, str):
            raise TypeError("Expected argument 'source_ip' to be a str")
        pulumi.set(__self__, "source_ip", source_ip)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def configs(self) -> Sequence['outputs.GetWebLockConfigsConfigResult']:
        """
        A list of Web Lock Config Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "configs")

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
        A list of Web Lock Config IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def lang(self) -> Optional[str]:
        return pulumi.get(self, "lang")

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
    @pulumi.getter
    def remark(self) -> Optional[str]:
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[str]:
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetWebLockConfigsResult(GetWebLockConfigsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebLockConfigsResult(
            configs=self.configs,
            id=self.id,
            ids=self.ids,
            lang=self.lang,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size,
            remark=self.remark,
            source_ip=self.source_ip,
            status=self.status)


def get_web_lock_configs(ids: Optional[Sequence[str]] = None,
                         lang: Optional[str] = None,
                         output_file: Optional[str] = None,
                         page_number: Optional[int] = None,
                         page_size: Optional[int] = None,
                         remark: Optional[str] = None,
                         source_ip: Optional[str] = None,
                         status: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebLockConfigsResult:
    """
    This data source provides Threat Detection Web Lock Config available to the user.[What is Web Lock Config](https://www.alibabacloud.com/help/en/security-center/latest/api-sas-2018-12-03-describeweblockconfiglist)

    > **NOTE:** Available since v1.195.0.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.threatdetection.get_web_lock_configs(ids=[alicloud_threat_detection_web_lock_config["default"]["id"]])
    pulumi.export("alicloudThreatDetectionWebLockConfigExampleId", default.configs[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Web Lock Config IDs.
    :param str lang: The language of the content within the request and the response. Valid values: `zh`, `en`.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str remark: The string that allows you to search for servers in fuzzy match mode. You can enter a server name or IP address.
    :param str source_ip: The source IP address of the request.
    :param str status: The protection status of the server that you want to query. Valid values: `on`, `off`.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['lang'] = lang
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    __args__['remark'] = remark
    __args__['sourceIp'] = source_ip
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:threatdetection/getWebLockConfigs:getWebLockConfigs', __args__, opts=opts, typ=GetWebLockConfigsResult).value

    return AwaitableGetWebLockConfigsResult(
        configs=pulumi.get(__ret__, 'configs'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        lang=pulumi.get(__ret__, 'lang'),
        output_file=pulumi.get(__ret__, 'output_file'),
        page_number=pulumi.get(__ret__, 'page_number'),
        page_size=pulumi.get(__ret__, 'page_size'),
        remark=pulumi.get(__ret__, 'remark'),
        source_ip=pulumi.get(__ret__, 'source_ip'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_web_lock_configs)
def get_web_lock_configs_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                lang: Optional[pulumi.Input[Optional[str]]] = None,
                                output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                page_number: Optional[pulumi.Input[Optional[int]]] = None,
                                page_size: Optional[pulumi.Input[Optional[int]]] = None,
                                remark: Optional[pulumi.Input[Optional[str]]] = None,
                                source_ip: Optional[pulumi.Input[Optional[str]]] = None,
                                status: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWebLockConfigsResult]:
    """
    This data source provides Threat Detection Web Lock Config available to the user.[What is Web Lock Config](https://www.alibabacloud.com/help/en/security-center/latest/api-sas-2018-12-03-describeweblockconfiglist)

    > **NOTE:** Available since v1.195.0.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.threatdetection.get_web_lock_configs(ids=[alicloud_threat_detection_web_lock_config["default"]["id"]])
    pulumi.export("alicloudThreatDetectionWebLockConfigExampleId", default.configs[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Web Lock Config IDs.
    :param str lang: The language of the content within the request and the response. Valid values: `zh`, `en`.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str remark: The string that allows you to search for servers in fuzzy match mode. You can enter a server name or IP address.
    :param str source_ip: The source IP address of the request.
    :param str status: The protection status of the server that you want to query. Valid values: `on`, `off`.
    """
    ...
