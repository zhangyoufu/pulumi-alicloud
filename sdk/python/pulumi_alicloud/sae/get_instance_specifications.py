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
    'GetInstanceSpecificationsResult',
    'AwaitableGetInstanceSpecificationsResult',
    'get_instance_specifications',
    'get_instance_specifications_output',
]

@pulumi.output_type
class GetInstanceSpecificationsResult:
    """
    A collection of values returned by getInstanceSpecifications.
    """
    def __init__(__self__, id=None, ids=None, output_file=None, specifications=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if specifications and not isinstance(specifications, list):
            raise TypeError("Expected argument 'specifications' to be a list")
        pulumi.set(__self__, "specifications", specifications)

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
    def specifications(self) -> Sequence['outputs.GetInstanceSpecificationsSpecificationResult']:
        return pulumi.get(self, "specifications")


class AwaitableGetInstanceSpecificationsResult(GetInstanceSpecificationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceSpecificationsResult(
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            specifications=self.specifications)


def get_instance_specifications(ids: Optional[Sequence[str]] = None,
                                output_file: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceSpecificationsResult:
    """
    This data source provides the Sae Instance Specifications of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.139.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.sae.get_instance_specifications()
    pulumi.export("saeInstanceSpecificationId1", ids.specifications[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Instance Specification IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:sae/getInstanceSpecifications:getInstanceSpecifications', __args__, opts=opts, typ=GetInstanceSpecificationsResult).value

    return AwaitableGetInstanceSpecificationsResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        output_file=pulumi.get(__ret__, 'output_file'),
        specifications=pulumi.get(__ret__, 'specifications'))


@_utilities.lift_output_func(get_instance_specifications)
def get_instance_specifications_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceSpecificationsResult]:
    """
    This data source provides the Sae Instance Specifications of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.139.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.sae.get_instance_specifications()
    pulumi.export("saeInstanceSpecificationId1", ids.specifications[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Instance Specification IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
