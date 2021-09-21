# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetMscSubSubscriptionsResult',
    'AwaitableGetMscSubSubscriptionsResult',
    'get_msc_sub_subscriptions',
]

@pulumi.output_type
class GetMscSubSubscriptionsResult:
    """
    A collection of values returned by getMscSubSubscriptions.
    """
    def __init__(__self__, id=None, output_file=None, subscriptions=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if subscriptions and not isinstance(subscriptions, list):
            raise TypeError("Expected argument 'subscriptions' to be a list")
        pulumi.set(__self__, "subscriptions", subscriptions)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def subscriptions(self) -> Sequence['outputs.GetMscSubSubscriptionsSubscriptionResult']:
        return pulumi.get(self, "subscriptions")


class AwaitableGetMscSubSubscriptionsResult(GetMscSubSubscriptionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMscSubSubscriptionsResult(
            id=self.id,
            output_file=self.output_file,
            subscriptions=self.subscriptions)


def get_msc_sub_subscriptions(output_file: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMscSubSubscriptionsResult:
    """
    This data source provides the Message Center Subscriptions of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.135.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.get_msc_sub_subscriptions()
    pulumi.export("mscSubSubscriptionId1", default.subscriptions[0].id)
    ```
    """
    __args__ = dict()
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:index/getMscSubSubscriptions:getMscSubSubscriptions', __args__, opts=opts, typ=GetMscSubSubscriptionsResult).value

    return AwaitableGetMscSubSubscriptionsResult(
        id=__ret__.id,
        output_file=__ret__.output_file,
        subscriptions=__ret__.subscriptions)
