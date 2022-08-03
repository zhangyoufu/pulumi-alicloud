# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceSlbConnAddr',
    'InstanceUiProxyConnAddr',
    'InstanceZkConnAddr',
    'GetInstanceTypesCoreInstanceTypeResult',
    'GetInstanceTypesMasterInstanceTypeResult',
    'GetInstanceTypesTypeResult',
    'GetInstancesInstanceResult',
    'GetZonesZoneResult',
]

@pulumi.output_type
class InstanceSlbConnAddr(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "connAddr":
            suggest = "conn_addr"
        elif key == "connAddrPort":
            suggest = "conn_addr_port"
        elif key == "netType":
            suggest = "net_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceSlbConnAddr. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceSlbConnAddr.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceSlbConnAddr.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 conn_addr: Optional[str] = None,
                 conn_addr_port: Optional[str] = None,
                 net_type: Optional[str] = None):
        if conn_addr is not None:
            pulumi.set(__self__, "conn_addr", conn_addr)
        if conn_addr_port is not None:
            pulumi.set(__self__, "conn_addr_port", conn_addr_port)
        if net_type is not None:
            pulumi.set(__self__, "net_type", net_type)

    @property
    @pulumi.getter(name="connAddr")
    def conn_addr(self) -> Optional[str]:
        return pulumi.get(self, "conn_addr")

    @property
    @pulumi.getter(name="connAddrPort")
    def conn_addr_port(self) -> Optional[str]:
        return pulumi.get(self, "conn_addr_port")

    @property
    @pulumi.getter(name="netType")
    def net_type(self) -> Optional[str]:
        return pulumi.get(self, "net_type")


@pulumi.output_type
class InstanceUiProxyConnAddr(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "connAddr":
            suggest = "conn_addr"
        elif key == "connAddrPort":
            suggest = "conn_addr_port"
        elif key == "netType":
            suggest = "net_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceUiProxyConnAddr. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceUiProxyConnAddr.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceUiProxyConnAddr.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 conn_addr: Optional[str] = None,
                 conn_addr_port: Optional[str] = None,
                 net_type: Optional[str] = None):
        if conn_addr is not None:
            pulumi.set(__self__, "conn_addr", conn_addr)
        if conn_addr_port is not None:
            pulumi.set(__self__, "conn_addr_port", conn_addr_port)
        if net_type is not None:
            pulumi.set(__self__, "net_type", net_type)

    @property
    @pulumi.getter(name="connAddr")
    def conn_addr(self) -> Optional[str]:
        return pulumi.get(self, "conn_addr")

    @property
    @pulumi.getter(name="connAddrPort")
    def conn_addr_port(self) -> Optional[str]:
        return pulumi.get(self, "conn_addr_port")

    @property
    @pulumi.getter(name="netType")
    def net_type(self) -> Optional[str]:
        return pulumi.get(self, "net_type")


@pulumi.output_type
class InstanceZkConnAddr(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "connAddr":
            suggest = "conn_addr"
        elif key == "connAddrPort":
            suggest = "conn_addr_port"
        elif key == "netType":
            suggest = "net_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceZkConnAddr. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceZkConnAddr.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceZkConnAddr.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 conn_addr: Optional[str] = None,
                 conn_addr_port: Optional[str] = None,
                 net_type: Optional[str] = None):
        if conn_addr is not None:
            pulumi.set(__self__, "conn_addr", conn_addr)
        if conn_addr_port is not None:
            pulumi.set(__self__, "conn_addr_port", conn_addr_port)
        if net_type is not None:
            pulumi.set(__self__, "net_type", net_type)

    @property
    @pulumi.getter(name="connAddr")
    def conn_addr(self) -> Optional[str]:
        return pulumi.get(self, "conn_addr")

    @property
    @pulumi.getter(name="connAddrPort")
    def conn_addr_port(self) -> Optional[str]:
        return pulumi.get(self, "conn_addr_port")

    @property
    @pulumi.getter(name="netType")
    def net_type(self) -> Optional[str]:
        return pulumi.get(self, "net_type")


@pulumi.output_type
class GetInstanceTypesCoreInstanceTypeResult(dict):
    def __init__(__self__, *,
                 category: str,
                 cpu_size: int,
                 engine: str,
                 instance_type: str,
                 max_core_count: int,
                 mem_size: int,
                 storage_type: str,
                 version: str,
                 zone: str):
        """
        :param str category: Name of the category, single or cluster.
        :param int cpu_size: Cpu size of the instance type.
        :param str engine: The engine name, `singlehbase`, `hbase`, `hbaseue`, `bds`.
        :param str instance_type: The hbase instance type of create hbase cluster instance.
        :param int max_core_count: Max count of the core instance nodes.
        :param int mem_size: Mem size of the instance type.
        :param str storage_type: Name of the storage type.
        :param str version: The engine version, singlehbase/hbase=1.1/2.0, bds=1.0.
        :param str zone: Name of zone id.
        """
        pulumi.set(__self__, "category", category)
        pulumi.set(__self__, "cpu_size", cpu_size)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "max_core_count", max_core_count)
        pulumi.set(__self__, "mem_size", mem_size)
        pulumi.set(__self__, "storage_type", storage_type)
        pulumi.set(__self__, "version", version)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def category(self) -> str:
        """
        Name of the category, single or cluster.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="cpuSize")
    def cpu_size(self) -> int:
        """
        Cpu size of the instance type.
        """
        return pulumi.get(self, "cpu_size")

    @property
    @pulumi.getter
    def engine(self) -> str:
        """
        The engine name, `singlehbase`, `hbase`, `hbaseue`, `bds`.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        The hbase instance type of create hbase cluster instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="maxCoreCount")
    def max_core_count(self) -> int:
        """
        Max count of the core instance nodes.
        """
        return pulumi.get(self, "max_core_count")

    @property
    @pulumi.getter(name="memSize")
    def mem_size(self) -> int:
        """
        Mem size of the instance type.
        """
        return pulumi.get(self, "mem_size")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> str:
        """
        Name of the storage type.
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The engine version, singlehbase/hbase=1.1/2.0, bds=1.0.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        Name of zone id.
        """
        return pulumi.get(self, "zone")


@pulumi.output_type
class GetInstanceTypesMasterInstanceTypeResult(dict):
    def __init__(__self__, *,
                 cpu_size: int,
                 instance_type: str,
                 mem_size: int):
        """
        :param int cpu_size: Cpu size of the instance type.
        :param str instance_type: The hbase instance type of create hbase cluster instance.
        :param int mem_size: Mem size of the instance type.
        """
        pulumi.set(__self__, "cpu_size", cpu_size)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "mem_size", mem_size)

    @property
    @pulumi.getter(name="cpuSize")
    def cpu_size(self) -> int:
        """
        Cpu size of the instance type.
        """
        return pulumi.get(self, "cpu_size")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        The hbase instance type of create hbase cluster instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="memSize")
    def mem_size(self) -> int:
        """
        Mem size of the instance type.
        """
        return pulumi.get(self, "mem_size")


@pulumi.output_type
class GetInstanceTypesTypeResult(dict):
    def __init__(__self__, *,
                 cpu_size: int,
                 mem_size: int,
                 value: str):
        """
        :param int cpu_size: Cpu size of the instance type.
        :param int mem_size: Mem size of the instance type.
        :param str value: Name of the instance type.
        """
        pulumi.set(__self__, "cpu_size", cpu_size)
        pulumi.set(__self__, "mem_size", mem_size)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="cpuSize")
    def cpu_size(self) -> int:
        """
        Cpu size of the instance type.
        """
        return pulumi.get(self, "cpu_size")

    @property
    @pulumi.getter(name="memSize")
    def mem_size(self) -> int:
        """
        Mem size of the instance type.
        """
        return pulumi.get(self, "mem_size")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Name of the instance type.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GetInstancesInstanceResult(dict):
    def __init__(__self__, *,
                 backup_status: str,
                 core_disk_size: int,
                 core_disk_type: str,
                 core_instance_type: str,
                 core_node_count: int,
                 created_time: str,
                 deletion_protection: bool,
                 engine: str,
                 engine_version: str,
                 expire_time: str,
                 id: str,
                 master_instance_type: str,
                 master_node_count: int,
                 name: str,
                 network_type: str,
                 pay_type: str,
                 region_id: str,
                 status: str,
                 vpc_id: str,
                 vswitch_id: str,
                 zone_id: str,
                 tags: Optional[Mapping[str, Any]] = None):
        """
        :param str backup_status: The Backup Status of the instance.
        :param int core_disk_size: Core node disk size, unit:GB.
        :param str core_disk_type: Cloud_ssd or cloud_efficiency
        :param str core_instance_type: Like hbase.sn1.large, hbase.sn1.large, hbase.n1.2xlarge and so on.
        :param int core_node_count: Same with "core_instance_quantity"
        :param str created_time: The created time of the instance.
        :param bool deletion_protection: The switch of delete protection.
        :param str engine: The engine of the instance.
        :param str engine_version: The engine_version of the instance.
        :param str expire_time: The expire time of the instance.
        :param str id: The ID of the HBase instance.
        :param str master_instance_type: Like hbase.sn1.large, hbase.sn1.large, hbase.n1.2xlarge and so on.
        :param int master_node_count: The node count of master
        :param str name: The name of the HBase instance.
        :param str network_type: Classic network or VPC.
        :param str pay_type: Billing method. Value options are `PostPaid` for  Pay-As-You-Go and `PrePaid` for yearly or monthly subscription.
        :param str region_id: Region ID the instance belongs to.
        :param str status: Status of the instance.
        :param str vpc_id: VPC ID the instance belongs to.
        :param str vswitch_id: VSwitch ID the instance belongs to.
        :param str zone_id: Zone ID the instance belongs to.
        :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "backup_status", backup_status)
        pulumi.set(__self__, "core_disk_size", core_disk_size)
        pulumi.set(__self__, "core_disk_type", core_disk_type)
        pulumi.set(__self__, "core_instance_type", core_instance_type)
        pulumi.set(__self__, "core_node_count", core_node_count)
        pulumi.set(__self__, "created_time", created_time)
        pulumi.set(__self__, "deletion_protection", deletion_protection)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "engine_version", engine_version)
        pulumi.set(__self__, "expire_time", expire_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "master_instance_type", master_instance_type)
        pulumi.set(__self__, "master_node_count", master_node_count)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "network_type", network_type)
        pulumi.set(__self__, "pay_type", pay_type)
        pulumi.set(__self__, "region_id", region_id)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        pulumi.set(__self__, "zone_id", zone_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="backupStatus")
    def backup_status(self) -> str:
        """
        The Backup Status of the instance.
        """
        return pulumi.get(self, "backup_status")

    @property
    @pulumi.getter(name="coreDiskSize")
    def core_disk_size(self) -> int:
        """
        Core node disk size, unit:GB.
        """
        return pulumi.get(self, "core_disk_size")

    @property
    @pulumi.getter(name="coreDiskType")
    def core_disk_type(self) -> str:
        """
        Cloud_ssd or cloud_efficiency
        """
        return pulumi.get(self, "core_disk_type")

    @property
    @pulumi.getter(name="coreInstanceType")
    def core_instance_type(self) -> str:
        """
        Like hbase.sn1.large, hbase.sn1.large, hbase.n1.2xlarge and so on.
        """
        return pulumi.get(self, "core_instance_type")

    @property
    @pulumi.getter(name="coreNodeCount")
    def core_node_count(self) -> int:
        """
        Same with "core_instance_quantity"
        """
        return pulumi.get(self, "core_node_count")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        The created time of the instance.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> bool:
        """
        The switch of delete protection.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def engine(self) -> str:
        """
        The engine of the instance.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> str:
        """
        The engine_version of the instance.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        The expire time of the instance.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the HBase instance.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="masterInstanceType")
    def master_instance_type(self) -> str:
        """
        Like hbase.sn1.large, hbase.sn1.large, hbase.n1.2xlarge and so on.
        """
        return pulumi.get(self, "master_instance_type")

    @property
    @pulumi.getter(name="masterNodeCount")
    def master_node_count(self) -> int:
        """
        The node count of master
        """
        return pulumi.get(self, "master_node_count")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the HBase instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> str:
        """
        Classic network or VPC.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="payType")
    def pay_type(self) -> str:
        """
        Billing method. Value options are `PostPaid` for  Pay-As-You-Go and `PrePaid` for yearly or monthly subscription.
        """
        return pulumi.get(self, "pay_type")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> str:
        """
        Region ID the instance belongs to.
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the instance.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        VPC ID the instance belongs to.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> str:
        """
        VSwitch ID the instance belongs to.
        """
        return pulumi.get(self, "vswitch_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        """
        Zone ID the instance belongs to.
        """
        return pulumi.get(self, "zone_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class GetZonesZoneResult(dict):
    def __init__(__self__, *,
                 id: str,
                 multi_zone_ids: Sequence[str]):
        """
        :param str id: ID of the zone.
        :param Sequence[str] multi_zone_ids: A list of zone ids in which the multi zone. Removed from v1.99.0.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "multi_zone_ids", multi_zone_ids)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the zone.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="multiZoneIds")
    def multi_zone_ids(self) -> Sequence[str]:
        """
        A list of zone ids in which the multi zone. Removed from v1.99.0.
        """
        return pulumi.get(self, "multi_zone_ids")


