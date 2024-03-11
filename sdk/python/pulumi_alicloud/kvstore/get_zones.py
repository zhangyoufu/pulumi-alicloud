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
    'GetZonesResult',
    'AwaitableGetZonesResult',
    'get_zones',
    'get_zones_output',
]

@pulumi.output_type
class GetZonesResult:
    """
    A collection of values returned by getZones.
    """
    def __init__(__self__, engine=None, id=None, ids=None, instance_charge_type=None, multi=None, output_file=None, product_type=None, zones=None):
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_charge_type and not isinstance(instance_charge_type, str):
            raise TypeError("Expected argument 'instance_charge_type' to be a str")
        pulumi.set(__self__, "instance_charge_type", instance_charge_type)
        if multi and not isinstance(multi, bool):
            raise TypeError("Expected argument 'multi' to be a bool")
        pulumi.set(__self__, "multi", multi)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if product_type and not isinstance(product_type, str):
            raise TypeError("Expected argument 'product_type' to be a str")
        pulumi.set(__self__, "product_type", product_type)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter
    def engine(self) -> Optional[str]:
        return pulumi.get(self, "engine")

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
        A list of zone IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> Optional[str]:
        return pulumi.get(self, "instance_charge_type")

    @property
    @pulumi.getter
    def multi(self) -> Optional[bool]:
        return pulumi.get(self, "multi")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> Optional[str]:
        return pulumi.get(self, "product_type")

    @property
    @pulumi.getter
    def zones(self) -> Sequence['outputs.GetZonesZoneResult']:
        """
        A list of availability zones. Each element contains the following attributes:
        """
        return pulumi.get(self, "zones")


class AwaitableGetZonesResult(GetZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZonesResult(
            engine=self.engine,
            id=self.id,
            ids=self.ids,
            instance_charge_type=self.instance_charge_type,
            multi=self.multi,
            output_file=self.output_file,
            product_type=self.product_type,
            zones=self.zones)


def get_zones(engine: Optional[str] = None,
              instance_charge_type: Optional[str] = None,
              multi: Optional[bool] = None,
              output_file: Optional[str] = None,
              product_type: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZonesResult:
    """
    This data source provides availability zones for KVStore that can be accessed by an Alibaba Cloud account within the region configured in the provider.

    > **NOTE:** Available since v1.73.0.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    zones_ids = alicloud.kvstore.get_zones(instance_charge_type="PostPaid")
    ```
    <!--End PulumiCodeChooser -->


    :param str engine: Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
           * product_type - (Optional, Available since v1.130.0+) The type of the service. Valid values: `Local`, `Tair_rdb`, `Tair_scm`, `Tair_essd`, `OnECS`.
    :param str instance_charge_type: Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
    :param bool multi: Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch KVStore instances.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['engine'] = engine
    __args__['instanceChargeType'] = instance_charge_type
    __args__['multi'] = multi
    __args__['outputFile'] = output_file
    __args__['productType'] = product_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:kvstore/getZones:getZones', __args__, opts=opts, typ=GetZonesResult).value

    return AwaitableGetZonesResult(
        engine=pulumi.get(__ret__, 'engine'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        instance_charge_type=pulumi.get(__ret__, 'instance_charge_type'),
        multi=pulumi.get(__ret__, 'multi'),
        output_file=pulumi.get(__ret__, 'output_file'),
        product_type=pulumi.get(__ret__, 'product_type'),
        zones=pulumi.get(__ret__, 'zones'))


@_utilities.lift_output_func(get_zones)
def get_zones_output(engine: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_charge_type: Optional[pulumi.Input[Optional[str]]] = None,
                     multi: Optional[pulumi.Input[Optional[bool]]] = None,
                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     product_type: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZonesResult]:
    """
    This data source provides availability zones for KVStore that can be accessed by an Alibaba Cloud account within the region configured in the provider.

    > **NOTE:** Available since v1.73.0.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    zones_ids = alicloud.kvstore.get_zones(instance_charge_type="PostPaid")
    ```
    <!--End PulumiCodeChooser -->


    :param str engine: Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
           * product_type - (Optional, Available since v1.130.0+) The type of the service. Valid values: `Local`, `Tair_rdb`, `Tair_scm`, `Tair_essd`, `OnECS`.
    :param str instance_charge_type: Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
    :param bool multi: Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch KVStore instances.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
