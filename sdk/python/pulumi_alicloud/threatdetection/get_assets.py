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
    'GetAssetsResult',
    'AwaitableGetAssetsResult',
    'get_assets',
    'get_assets_output',
]

@pulumi.output_type
class GetAssetsResult:
    """
    A collection of values returned by getAssets.
    """
    def __init__(__self__, assets=None, criteria=None, id=None, ids=None, importance=None, logical_exp=None, machine_types=None, no_group_trace=None, output_file=None, page_number=None, page_size=None):
        if assets and not isinstance(assets, list):
            raise TypeError("Expected argument 'assets' to be a list")
        pulumi.set(__self__, "assets", assets)
        if criteria and not isinstance(criteria, str):
            raise TypeError("Expected argument 'criteria' to be a str")
        pulumi.set(__self__, "criteria", criteria)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if importance and not isinstance(importance, int):
            raise TypeError("Expected argument 'importance' to be a int")
        pulumi.set(__self__, "importance", importance)
        if logical_exp and not isinstance(logical_exp, str):
            raise TypeError("Expected argument 'logical_exp' to be a str")
        pulumi.set(__self__, "logical_exp", logical_exp)
        if machine_types and not isinstance(machine_types, str):
            raise TypeError("Expected argument 'machine_types' to be a str")
        pulumi.set(__self__, "machine_types", machine_types)
        if no_group_trace and not isinstance(no_group_trace, bool):
            raise TypeError("Expected argument 'no_group_trace' to be a bool")
        pulumi.set(__self__, "no_group_trace", no_group_trace)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if page_number and not isinstance(page_number, int):
            raise TypeError("Expected argument 'page_number' to be a int")
        pulumi.set(__self__, "page_number", page_number)
        if page_size and not isinstance(page_size, int):
            raise TypeError("Expected argument 'page_size' to be a int")
        pulumi.set(__self__, "page_size", page_size)

    @property
    @pulumi.getter
    def assets(self) -> Sequence['outputs.GetAssetsAssetResult']:
        """
        A list of Asset Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "assets")

    @property
    @pulumi.getter
    def criteria(self) -> Optional[str]:
        return pulumi.get(self, "criteria")

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
        A list of Asset IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def importance(self) -> Optional[int]:
        return pulumi.get(self, "importance")

    @property
    @pulumi.getter(name="logicalExp")
    def logical_exp(self) -> Optional[str]:
        return pulumi.get(self, "logical_exp")

    @property
    @pulumi.getter(name="machineTypes")
    def machine_types(self) -> Optional[str]:
        return pulumi.get(self, "machine_types")

    @property
    @pulumi.getter(name="noGroupTrace")
    def no_group_trace(self) -> Optional[bool]:
        return pulumi.get(self, "no_group_trace")

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


class AwaitableGetAssetsResult(GetAssetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAssetsResult(
            assets=self.assets,
            criteria=self.criteria,
            id=self.id,
            ids=self.ids,
            importance=self.importance,
            logical_exp=self.logical_exp,
            machine_types=self.machine_types,
            no_group_trace=self.no_group_trace,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size)


def get_assets(criteria: Optional[str] = None,
               ids: Optional[Sequence[str]] = None,
               importance: Optional[int] = None,
               logical_exp: Optional[str] = None,
               machine_types: Optional[str] = None,
               no_group_trace: Optional[bool] = None,
               output_file: Optional[str] = None,
               page_number: Optional[int] = None,
               page_size: Optional[int] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAssetsResult:
    """
    This data source provides Threat Detection Asset available to the user.[What is Asset](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describecloudcenterinstances)

    > **NOTE:** Available since v1.195.0.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.threatdetection.get_assets()
    pulumi.export("alicloudThreatDetectionAssetExampleId", default.assets[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str criteria: Set the conditions for searching assets. This parameter is in JSON format. Note the case when you enter the parameter. **NOTE:** You can search for assets by using conditions such as the instance ID, instance name, VPC ID, region, and public IP address of the asset.
    :param Sequence[str] ids: A list of Asset IDs.
    :param int importance: Set asset importance. Value:
           - **2**: Significant assets
           - **1**: General assets
           - **0**: Test asset
    :param str logical_exp: Set the logical relationship between multiple search conditions. The default value is **OR**. Valid values:
           - **OR**: indicates that the relationship between multiple search conditions is **OR**.
           - **AND**: indicates that the relationship between multiple search conditions is **AND**.
    :param str machine_types: The type of asset to query. Value:
           - **ecs**: server.
           - **cloud_product**: Cloud product.
    :param bool no_group_trace: Specifies whether to internationalize the name of the default group. Default value: false
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['criteria'] = criteria
    __args__['ids'] = ids
    __args__['importance'] = importance
    __args__['logicalExp'] = logical_exp
    __args__['machineTypes'] = machine_types
    __args__['noGroupTrace'] = no_group_trace
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:threatdetection/getAssets:getAssets', __args__, opts=opts, typ=GetAssetsResult).value

    return AwaitableGetAssetsResult(
        assets=pulumi.get(__ret__, 'assets'),
        criteria=pulumi.get(__ret__, 'criteria'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        importance=pulumi.get(__ret__, 'importance'),
        logical_exp=pulumi.get(__ret__, 'logical_exp'),
        machine_types=pulumi.get(__ret__, 'machine_types'),
        no_group_trace=pulumi.get(__ret__, 'no_group_trace'),
        output_file=pulumi.get(__ret__, 'output_file'),
        page_number=pulumi.get(__ret__, 'page_number'),
        page_size=pulumi.get(__ret__, 'page_size'))


@_utilities.lift_output_func(get_assets)
def get_assets_output(criteria: Optional[pulumi.Input[Optional[str]]] = None,
                      ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                      importance: Optional[pulumi.Input[Optional[int]]] = None,
                      logical_exp: Optional[pulumi.Input[Optional[str]]] = None,
                      machine_types: Optional[pulumi.Input[Optional[str]]] = None,
                      no_group_trace: Optional[pulumi.Input[Optional[bool]]] = None,
                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      page_number: Optional[pulumi.Input[Optional[int]]] = None,
                      page_size: Optional[pulumi.Input[Optional[int]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAssetsResult]:
    """
    This data source provides Threat Detection Asset available to the user.[What is Asset](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describecloudcenterinstances)

    > **NOTE:** Available since v1.195.0.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.threatdetection.get_assets()
    pulumi.export("alicloudThreatDetectionAssetExampleId", default.assets[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str criteria: Set the conditions for searching assets. This parameter is in JSON format. Note the case when you enter the parameter. **NOTE:** You can search for assets by using conditions such as the instance ID, instance name, VPC ID, region, and public IP address of the asset.
    :param Sequence[str] ids: A list of Asset IDs.
    :param int importance: Set asset importance. Value:
           - **2**: Significant assets
           - **1**: General assets
           - **0**: Test asset
    :param str logical_exp: Set the logical relationship between multiple search conditions. The default value is **OR**. Valid values:
           - **OR**: indicates that the relationship between multiple search conditions is **OR**.
           - **AND**: indicates that the relationship between multiple search conditions is **AND**.
    :param str machine_types: The type of asset to query. Value:
           - **ecs**: server.
           - **cloud_product**: Cloud product.
    :param bool no_group_trace: Specifies whether to internationalize the name of the default group. Default value: false
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
