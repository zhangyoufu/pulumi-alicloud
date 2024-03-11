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
    'GetAggregatorsResult',
    'AwaitableGetAggregatorsResult',
    'get_aggregators',
    'get_aggregators_output',
]

@pulumi.output_type
class GetAggregatorsResult:
    """
    A collection of values returned by getAggregators.
    """
    def __init__(__self__, aggregators=None, enable_details=None, id=None, ids=None, name_regex=None, names=None, output_file=None, status=None):
        if aggregators and not isinstance(aggregators, list):
            raise TypeError("Expected argument 'aggregators' to be a list")
        pulumi.set(__self__, "aggregators", aggregators)
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
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def aggregators(self) -> Sequence['outputs.GetAggregatorsAggregatorResult']:
        """
        A list of config aggregators. Each element contains the following attributes:
        """
        return pulumi.get(self, "aggregators")

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
        """
        A list of Aggregator names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")


class AwaitableGetAggregatorsResult(GetAggregatorsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAggregatorsResult(
            aggregators=self.aggregators,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            status=self.status)


def get_aggregators(enable_details: Optional[bool] = None,
                    ids: Optional[Sequence[str]] = None,
                    name_regex: Optional[str] = None,
                    output_file: Optional[str] = None,
                    status: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAggregatorsResult:
    """
    This data source provides the Config Aggregators of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.124.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cfg.get_aggregators(ids=["ca-3ce2626622af0005****"],
        name_regex="the_resource_name")
    pulumi.export("firstConfigAggregatorId", example.aggregators[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of aggregator ids.
    :param str name_regex: A regex string to filter results by aggregator name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource. Valid Values:  `0`: creating `1`: normal `2`: deleting.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cfg/getAggregators:getAggregators', __args__, opts=opts, typ=GetAggregatorsResult).value

    return AwaitableGetAggregatorsResult(
        aggregators=pulumi.get(__ret__, 'aggregators'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_aggregators)
def get_aggregators_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                           ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           status: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAggregatorsResult]:
    """
    This data source provides the Config Aggregators of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.124.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cfg.get_aggregators(ids=["ca-3ce2626622af0005****"],
        name_regex="the_resource_name")
    pulumi.export("firstConfigAggregatorId", example.aggregators[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of aggregator ids.
    :param str name_regex: A regex string to filter results by aggregator name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource. Valid Values:  `0`: creating `1`: normal `2`: deleting.
    """
    ...
