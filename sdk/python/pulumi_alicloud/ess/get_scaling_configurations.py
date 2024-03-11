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
    'GetScalingConfigurationsResult',
    'AwaitableGetScalingConfigurationsResult',
    'get_scaling_configurations',
    'get_scaling_configurations_output',
]

@pulumi.output_type
class GetScalingConfigurationsResult:
    """
    A collection of values returned by getScalingConfigurations.
    """
    def __init__(__self__, configurations=None, id=None, ids=None, name_regex=None, names=None, output_file=None, scaling_group_id=None):
        if configurations and not isinstance(configurations, list):
            raise TypeError("Expected argument 'configurations' to be a list")
        pulumi.set(__self__, "configurations", configurations)
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
        if scaling_group_id and not isinstance(scaling_group_id, str):
            raise TypeError("Expected argument 'scaling_group_id' to be a str")
        pulumi.set(__self__, "scaling_group_id", scaling_group_id)

    @property
    @pulumi.getter
    def configurations(self) -> Sequence['outputs.GetScalingConfigurationsConfigurationResult']:
        """
        A list of scaling rules. Each element contains the following attributes:
        """
        return pulumi.get(self, "configurations")

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
        A list of scaling configuration ids.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of scaling configuration names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> Optional[str]:
        """
        ID of the scaling group.
        """
        return pulumi.get(self, "scaling_group_id")


class AwaitableGetScalingConfigurationsResult(GetScalingConfigurationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScalingConfigurationsResult(
            configurations=self.configurations,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            scaling_group_id=self.scaling_group_id)


def get_scaling_configurations(ids: Optional[Sequence[str]] = None,
                               name_regex: Optional[str] = None,
                               output_file: Optional[str] = None,
                               scaling_group_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScalingConfigurationsResult:
    """
    This data source provides available scaling configuration resources.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    scalingconfigurations_ds = alicloud.ess.get_scaling_configurations(ids=[
            "scaling_configuration_id1",
            "scaling_configuration_id2",
        ],
        name_regex="scaling_configuration_name",
        scaling_group_id="scaling_group_id")
    pulumi.export("firstScalingRule", scalingconfigurations_ds.configurations[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of scaling configuration IDs.
    :param str name_regex: A regex string to filter resulting scaling configurations by name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str scaling_group_id: Scaling group id the scaling configurations belong to.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['scalingGroupId'] = scaling_group_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ess/getScalingConfigurations:getScalingConfigurations', __args__, opts=opts, typ=GetScalingConfigurationsResult).value

    return AwaitableGetScalingConfigurationsResult(
        configurations=pulumi.get(__ret__, 'configurations'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        scaling_group_id=pulumi.get(__ret__, 'scaling_group_id'))


@_utilities.lift_output_func(get_scaling_configurations)
def get_scaling_configurations_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                      name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                      scaling_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetScalingConfigurationsResult]:
    """
    This data source provides available scaling configuration resources.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    scalingconfigurations_ds = alicloud.ess.get_scaling_configurations(ids=[
            "scaling_configuration_id1",
            "scaling_configuration_id2",
        ],
        name_regex="scaling_configuration_name",
        scaling_group_id="scaling_group_id")
    pulumi.export("firstScalingRule", scalingconfigurations_ds.configurations[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of scaling configuration IDs.
    :param str name_regex: A regex string to filter resulting scaling configurations by name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str scaling_group_id: Scaling group id the scaling configurations belong to.
    """
    ...
