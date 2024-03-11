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
    'GetInstanceTypesResult',
    'AwaitableGetInstanceTypesResult',
    'get_instance_types',
    'get_instance_types_output',
]

@pulumi.output_type
class GetInstanceTypesResult:
    """
    A collection of values returned by getInstanceTypes.
    """
    def __init__(__self__, availability_zone=None, cpu_core_count=None, eni_amount=None, gpu_amount=None, gpu_spec=None, id=None, ids=None, image_id=None, instance_charge_type=None, instance_type_family=None, instance_types=None, is_outdated=None, kubernetes_node_role=None, memory_size=None, minimum_eni_ipv6_address_quantity=None, network_type=None, output_file=None, sorted_by=None, spot_strategy=None, system_disk_category=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if cpu_core_count and not isinstance(cpu_core_count, int):
            raise TypeError("Expected argument 'cpu_core_count' to be a int")
        pulumi.set(__self__, "cpu_core_count", cpu_core_count)
        if eni_amount and not isinstance(eni_amount, int):
            raise TypeError("Expected argument 'eni_amount' to be a int")
        pulumi.set(__self__, "eni_amount", eni_amount)
        if gpu_amount and not isinstance(gpu_amount, int):
            raise TypeError("Expected argument 'gpu_amount' to be a int")
        pulumi.set(__self__, "gpu_amount", gpu_amount)
        if gpu_spec and not isinstance(gpu_spec, str):
            raise TypeError("Expected argument 'gpu_spec' to be a str")
        pulumi.set(__self__, "gpu_spec", gpu_spec)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if instance_charge_type and not isinstance(instance_charge_type, str):
            raise TypeError("Expected argument 'instance_charge_type' to be a str")
        pulumi.set(__self__, "instance_charge_type", instance_charge_type)
        if instance_type_family and not isinstance(instance_type_family, str):
            raise TypeError("Expected argument 'instance_type_family' to be a str")
        pulumi.set(__self__, "instance_type_family", instance_type_family)
        if instance_types and not isinstance(instance_types, list):
            raise TypeError("Expected argument 'instance_types' to be a list")
        pulumi.set(__self__, "instance_types", instance_types)
        if is_outdated and not isinstance(is_outdated, bool):
            raise TypeError("Expected argument 'is_outdated' to be a bool")
        pulumi.set(__self__, "is_outdated", is_outdated)
        if kubernetes_node_role and not isinstance(kubernetes_node_role, str):
            raise TypeError("Expected argument 'kubernetes_node_role' to be a str")
        pulumi.set(__self__, "kubernetes_node_role", kubernetes_node_role)
        if memory_size and not isinstance(memory_size, float):
            raise TypeError("Expected argument 'memory_size' to be a float")
        pulumi.set(__self__, "memory_size", memory_size)
        if minimum_eni_ipv6_address_quantity and not isinstance(minimum_eni_ipv6_address_quantity, int):
            raise TypeError("Expected argument 'minimum_eni_ipv6_address_quantity' to be a int")
        pulumi.set(__self__, "minimum_eni_ipv6_address_quantity", minimum_eni_ipv6_address_quantity)
        if network_type and not isinstance(network_type, str):
            raise TypeError("Expected argument 'network_type' to be a str")
        pulumi.set(__self__, "network_type", network_type)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if sorted_by and not isinstance(sorted_by, str):
            raise TypeError("Expected argument 'sorted_by' to be a str")
        pulumi.set(__self__, "sorted_by", sorted_by)
        if spot_strategy and not isinstance(spot_strategy, str):
            raise TypeError("Expected argument 'spot_strategy' to be a str")
        pulumi.set(__self__, "spot_strategy", spot_strategy)
        if system_disk_category and not isinstance(system_disk_category, str):
            raise TypeError("Expected argument 'system_disk_category' to be a str")
        pulumi.set(__self__, "system_disk_category", system_disk_category)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="cpuCoreCount")
    def cpu_core_count(self) -> Optional[int]:
        """
        Number of CPU cores.
        """
        return pulumi.get(self, "cpu_core_count")

    @property
    @pulumi.getter(name="eniAmount")
    def eni_amount(self) -> Optional[int]:
        """
        The maximum number of network interfaces that an instance type can be attached to.
        """
        return pulumi.get(self, "eni_amount")

    @property
    @pulumi.getter(name="gpuAmount")
    def gpu_amount(self) -> Optional[int]:
        return pulumi.get(self, "gpu_amount")

    @property
    @pulumi.getter(name="gpuSpec")
    def gpu_spec(self) -> Optional[str]:
        return pulumi.get(self, "gpu_spec")

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
        A list of instance type IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[str]:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> Optional[str]:
        return pulumi.get(self, "instance_charge_type")

    @property
    @pulumi.getter(name="instanceTypeFamily")
    def instance_type_family(self) -> Optional[str]:
        return pulumi.get(self, "instance_type_family")

    @property
    @pulumi.getter(name="instanceTypes")
    def instance_types(self) -> Sequence['outputs.GetInstanceTypesInstanceTypeResult']:
        """
        A list of image types. Each element contains the following attributes:
        """
        return pulumi.get(self, "instance_types")

    @property
    @pulumi.getter(name="isOutdated")
    def is_outdated(self) -> Optional[bool]:
        return pulumi.get(self, "is_outdated")

    @property
    @pulumi.getter(name="kubernetesNodeRole")
    def kubernetes_node_role(self) -> Optional[str]:
        return pulumi.get(self, "kubernetes_node_role")

    @property
    @pulumi.getter(name="memorySize")
    def memory_size(self) -> Optional[float]:
        """
        Size of memory, measured in GB.
        """
        return pulumi.get(self, "memory_size")

    @property
    @pulumi.getter(name="minimumEniIpv6AddressQuantity")
    def minimum_eni_ipv6_address_quantity(self) -> Optional[int]:
        return pulumi.get(self, "minimum_eni_ipv6_address_quantity")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[str]:
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="sortedBy")
    def sorted_by(self) -> Optional[str]:
        return pulumi.get(self, "sorted_by")

    @property
    @pulumi.getter(name="spotStrategy")
    def spot_strategy(self) -> Optional[str]:
        return pulumi.get(self, "spot_strategy")

    @property
    @pulumi.getter(name="systemDiskCategory")
    def system_disk_category(self) -> Optional[str]:
        return pulumi.get(self, "system_disk_category")


class AwaitableGetInstanceTypesResult(GetInstanceTypesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceTypesResult(
            availability_zone=self.availability_zone,
            cpu_core_count=self.cpu_core_count,
            eni_amount=self.eni_amount,
            gpu_amount=self.gpu_amount,
            gpu_spec=self.gpu_spec,
            id=self.id,
            ids=self.ids,
            image_id=self.image_id,
            instance_charge_type=self.instance_charge_type,
            instance_type_family=self.instance_type_family,
            instance_types=self.instance_types,
            is_outdated=self.is_outdated,
            kubernetes_node_role=self.kubernetes_node_role,
            memory_size=self.memory_size,
            minimum_eni_ipv6_address_quantity=self.minimum_eni_ipv6_address_quantity,
            network_type=self.network_type,
            output_file=self.output_file,
            sorted_by=self.sorted_by,
            spot_strategy=self.spot_strategy,
            system_disk_category=self.system_disk_category)


def get_instance_types(availability_zone: Optional[str] = None,
                       cpu_core_count: Optional[int] = None,
                       eni_amount: Optional[int] = None,
                       gpu_amount: Optional[int] = None,
                       gpu_spec: Optional[str] = None,
                       image_id: Optional[str] = None,
                       instance_charge_type: Optional[str] = None,
                       instance_type_family: Optional[str] = None,
                       is_outdated: Optional[bool] = None,
                       kubernetes_node_role: Optional[str] = None,
                       memory_size: Optional[float] = None,
                       minimum_eni_ipv6_address_quantity: Optional[int] = None,
                       network_type: Optional[str] = None,
                       output_file: Optional[str] = None,
                       sorted_by: Optional[str] = None,
                       spot_strategy: Optional[str] = None,
                       system_disk_category: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceTypesResult:
    """
    This data source provides the ECS instance types of Alibaba Cloud.

    > **NOTE:** By default, only the upgraded instance types are returned. If you want to get outdated instance types, you must set `is_outdated` to true.

    > **NOTE:** If one instance type is sold out, it will not be exported.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    types_ds = alicloud.ecs.get_instance_types(cpu_core_count=1,
        memory_size=2)
    instance = alicloud.ecs.Instance("instance", instance_type=types_ds.instance_types[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str availability_zone: The zone where instance types are supported.
    :param int cpu_core_count: Filter the results to a specific number of cpu cores.
    :param int eni_amount: Filter the result whose network interface number is no more than `eni_amount`.
    :param int gpu_amount: The GPU amount of an instance type.
    :param str gpu_spec: The GPU spec of an instance type.
    :param str image_id: The ID of the image.
    :param str instance_charge_type: Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
    :param str instance_type_family: Filter the results based on their family name. For example: 'ecs.n4'.
    :param bool is_outdated: If true, outdated instance types are included in the results. Default to false.
    :param str kubernetes_node_role: Filter the result which is used to create a kubernetes cluster
           and managed kubernetes cluster. Optional Values: `Master` and `Worker`.
    :param float memory_size: Filter the results to a specific memory size in GB.
    :param int minimum_eni_ipv6_address_quantity: The minimum number of IPv6 addresses per ENI. **Note:** If an instance type supports fewer IPv6 addresses per ENI than the specified value, information about the instance type is not queried.
    :param str network_type: Filter the results by network type. Valid values: `Classic` and `Vpc`.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str spot_strategy: Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
    :param str system_disk_category: Filter the results by system disk category. Valid values: `cloud`, `ephemeral_ssd`, `cloud_essd`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd_entry`. 
           **NOTE**: Its default value `cloud_efficiency` has been removed from the version v1.150.0.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['cpuCoreCount'] = cpu_core_count
    __args__['eniAmount'] = eni_amount
    __args__['gpuAmount'] = gpu_amount
    __args__['gpuSpec'] = gpu_spec
    __args__['imageId'] = image_id
    __args__['instanceChargeType'] = instance_charge_type
    __args__['instanceTypeFamily'] = instance_type_family
    __args__['isOutdated'] = is_outdated
    __args__['kubernetesNodeRole'] = kubernetes_node_role
    __args__['memorySize'] = memory_size
    __args__['minimumEniIpv6AddressQuantity'] = minimum_eni_ipv6_address_quantity
    __args__['networkType'] = network_type
    __args__['outputFile'] = output_file
    __args__['sortedBy'] = sorted_by
    __args__['spotStrategy'] = spot_strategy
    __args__['systemDiskCategory'] = system_disk_category
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getInstanceTypes:getInstanceTypes', __args__, opts=opts, typ=GetInstanceTypesResult).value

    return AwaitableGetInstanceTypesResult(
        availability_zone=pulumi.get(__ret__, 'availability_zone'),
        cpu_core_count=pulumi.get(__ret__, 'cpu_core_count'),
        eni_amount=pulumi.get(__ret__, 'eni_amount'),
        gpu_amount=pulumi.get(__ret__, 'gpu_amount'),
        gpu_spec=pulumi.get(__ret__, 'gpu_spec'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        image_id=pulumi.get(__ret__, 'image_id'),
        instance_charge_type=pulumi.get(__ret__, 'instance_charge_type'),
        instance_type_family=pulumi.get(__ret__, 'instance_type_family'),
        instance_types=pulumi.get(__ret__, 'instance_types'),
        is_outdated=pulumi.get(__ret__, 'is_outdated'),
        kubernetes_node_role=pulumi.get(__ret__, 'kubernetes_node_role'),
        memory_size=pulumi.get(__ret__, 'memory_size'),
        minimum_eni_ipv6_address_quantity=pulumi.get(__ret__, 'minimum_eni_ipv6_address_quantity'),
        network_type=pulumi.get(__ret__, 'network_type'),
        output_file=pulumi.get(__ret__, 'output_file'),
        sorted_by=pulumi.get(__ret__, 'sorted_by'),
        spot_strategy=pulumi.get(__ret__, 'spot_strategy'),
        system_disk_category=pulumi.get(__ret__, 'system_disk_category'))


@_utilities.lift_output_func(get_instance_types)
def get_instance_types_output(availability_zone: Optional[pulumi.Input[Optional[str]]] = None,
                              cpu_core_count: Optional[pulumi.Input[Optional[int]]] = None,
                              eni_amount: Optional[pulumi.Input[Optional[int]]] = None,
                              gpu_amount: Optional[pulumi.Input[Optional[int]]] = None,
                              gpu_spec: Optional[pulumi.Input[Optional[str]]] = None,
                              image_id: Optional[pulumi.Input[Optional[str]]] = None,
                              instance_charge_type: Optional[pulumi.Input[Optional[str]]] = None,
                              instance_type_family: Optional[pulumi.Input[Optional[str]]] = None,
                              is_outdated: Optional[pulumi.Input[Optional[bool]]] = None,
                              kubernetes_node_role: Optional[pulumi.Input[Optional[str]]] = None,
                              memory_size: Optional[pulumi.Input[Optional[float]]] = None,
                              minimum_eni_ipv6_address_quantity: Optional[pulumi.Input[Optional[int]]] = None,
                              network_type: Optional[pulumi.Input[Optional[str]]] = None,
                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              sorted_by: Optional[pulumi.Input[Optional[str]]] = None,
                              spot_strategy: Optional[pulumi.Input[Optional[str]]] = None,
                              system_disk_category: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceTypesResult]:
    """
    This data source provides the ECS instance types of Alibaba Cloud.

    > **NOTE:** By default, only the upgraded instance types are returned. If you want to get outdated instance types, you must set `is_outdated` to true.

    > **NOTE:** If one instance type is sold out, it will not be exported.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    types_ds = alicloud.ecs.get_instance_types(cpu_core_count=1,
        memory_size=2)
    instance = alicloud.ecs.Instance("instance", instance_type=types_ds.instance_types[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str availability_zone: The zone where instance types are supported.
    :param int cpu_core_count: Filter the results to a specific number of cpu cores.
    :param int eni_amount: Filter the result whose network interface number is no more than `eni_amount`.
    :param int gpu_amount: The GPU amount of an instance type.
    :param str gpu_spec: The GPU spec of an instance type.
    :param str image_id: The ID of the image.
    :param str instance_charge_type: Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
    :param str instance_type_family: Filter the results based on their family name. For example: 'ecs.n4'.
    :param bool is_outdated: If true, outdated instance types are included in the results. Default to false.
    :param str kubernetes_node_role: Filter the result which is used to create a kubernetes cluster
           and managed kubernetes cluster. Optional Values: `Master` and `Worker`.
    :param float memory_size: Filter the results to a specific memory size in GB.
    :param int minimum_eni_ipv6_address_quantity: The minimum number of IPv6 addresses per ENI. **Note:** If an instance type supports fewer IPv6 addresses per ENI than the specified value, information about the instance type is not queried.
    :param str network_type: Filter the results by network type. Valid values: `Classic` and `Vpc`.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str spot_strategy: Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
    :param str system_disk_category: Filter the results by system disk category. Valid values: `cloud`, `ephemeral_ssd`, `cloud_essd`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd_entry`. 
           **NOTE**: Its default value `cloud_efficiency` has been removed from the version v1.150.0.
    """
    ...
