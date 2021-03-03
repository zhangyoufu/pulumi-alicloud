# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetQuotaAlarmsResult',
    'AwaitableGetQuotaAlarmsResult',
    'get_quota_alarms',
]

@pulumi.output_type
class GetQuotaAlarmsResult:
    """
    A collection of values returned by getQuotaAlarms.
    """
    def __init__(__self__, alarms=None, enable_details=None, id=None, ids=None, name_regex=None, names=None, output_file=None, product_code=None, quota_action_code=None, quota_alarm_name=None, quota_dimensions=None):
        if alarms and not isinstance(alarms, list):
            raise TypeError("Expected argument 'alarms' to be a list")
        pulumi.set(__self__, "alarms", alarms)
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
        if product_code and not isinstance(product_code, str):
            raise TypeError("Expected argument 'product_code' to be a str")
        pulumi.set(__self__, "product_code", product_code)
        if quota_action_code and not isinstance(quota_action_code, str):
            raise TypeError("Expected argument 'quota_action_code' to be a str")
        pulumi.set(__self__, "quota_action_code", quota_action_code)
        if quota_alarm_name and not isinstance(quota_alarm_name, str):
            raise TypeError("Expected argument 'quota_alarm_name' to be a str")
        pulumi.set(__self__, "quota_alarm_name", quota_alarm_name)
        if quota_dimensions and not isinstance(quota_dimensions, list):
            raise TypeError("Expected argument 'quota_dimensions' to be a list")
        pulumi.set(__self__, "quota_dimensions", quota_dimensions)

    @property
    @pulumi.getter
    def alarms(self) -> Sequence['outputs.GetQuotaAlarmsAlarmResult']:
        return pulumi.get(self, "alarms")

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
    @pulumi.getter(name="productCode")
    def product_code(self) -> Optional[str]:
        return pulumi.get(self, "product_code")

    @property
    @pulumi.getter(name="quotaActionCode")
    def quota_action_code(self) -> Optional[str]:
        return pulumi.get(self, "quota_action_code")

    @property
    @pulumi.getter(name="quotaAlarmName")
    def quota_alarm_name(self) -> Optional[str]:
        return pulumi.get(self, "quota_alarm_name")

    @property
    @pulumi.getter(name="quotaDimensions")
    def quota_dimensions(self) -> Optional[Sequence['outputs.GetQuotaAlarmsQuotaDimensionResult']]:
        return pulumi.get(self, "quota_dimensions")


class AwaitableGetQuotaAlarmsResult(GetQuotaAlarmsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQuotaAlarmsResult(
            alarms=self.alarms,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            product_code=self.product_code,
            quota_action_code=self.quota_action_code,
            quota_alarm_name=self.quota_alarm_name,
            quota_dimensions=self.quota_dimensions)


def get_quota_alarms(enable_details: Optional[bool] = None,
                     ids: Optional[Sequence[str]] = None,
                     name_regex: Optional[str] = None,
                     output_file: Optional[str] = None,
                     product_code: Optional[str] = None,
                     quota_action_code: Optional[str] = None,
                     quota_alarm_name: Optional[str] = None,
                     quota_dimensions: Optional[Sequence[pulumi.InputType['GetQuotaAlarmsQuotaDimensionArgs']]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQuotaAlarmsResult:
    """
    This data source provides the Quotas Quota Alarms of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.116.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.quotas.get_quota_alarms(ids=["5VR90-421F886-81E9-xxx"],
        name_regex="tf-testAcc")
    pulumi.export("firstQuotasQuotaAlarmId", example.alarms[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Quota Alarm IDs.
    :param str name_regex: A regex string to filter results by Quota Alarm name.
    :param str product_code: The Product Code.
    :param str quota_action_code: The Quota Action Code.
    :param str quota_alarm_name: The name of Quota Alarm.
    :param Sequence[pulumi.InputType['GetQuotaAlarmsQuotaDimensionArgs']] quota_dimensions: The Quota Dimensions.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['productCode'] = product_code
    __args__['quotaActionCode'] = quota_action_code
    __args__['quotaAlarmName'] = quota_alarm_name
    __args__['quotaDimensions'] = quota_dimensions
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:quotas/getQuotaAlarms:getQuotaAlarms', __args__, opts=opts, typ=GetQuotaAlarmsResult).value

    return AwaitableGetQuotaAlarmsResult(
        alarms=__ret__.alarms,
        enable_details=__ret__.enable_details,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        product_code=__ret__.product_code,
        quota_action_code=__ret__.quota_action_code,
        quota_alarm_name=__ret__.quota_alarm_name,
        quota_dimensions=__ret__.quota_dimensions)
