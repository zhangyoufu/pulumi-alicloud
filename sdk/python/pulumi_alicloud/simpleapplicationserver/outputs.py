# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetImagesImageResult',
    'GetInstancesInstanceResult',
    'GetServerPlansPlanResult',
]

@pulumi.output_type
class GetImagesImageResult(dict):
    def __init__(__self__, *,
                 description: str,
                 id: str,
                 image_id: str,
                 image_name: str,
                 image_type: str):
        """
        :param str description: The description of the image.
        :param str id: The ID of the Instance Image.
        :param str image_id: The ID of the image.
        :param str image_name: The name of the resource.
        :param str image_type: The type of the image. Valid values: `app`, `custom`, `system`.
               * `system`: operating system (OS) image.
               * `app`: application image.
               * `custom`: custom image.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "image_id", image_id)
        pulumi.set(__self__, "image_name", image_name)
        pulumi.set(__self__, "image_type", image_type)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the image.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Instance Image.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        """
        The ID of the image.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="imageType")
    def image_type(self) -> str:
        """
        The type of the image. Valid values: `app`, `custom`, `system`.
        * `system`: operating system (OS) image.
        * `app`: application image.
        * `custom`: custom image.
        """
        return pulumi.get(self, "image_type")


@pulumi.output_type
class GetInstancesInstanceResult(dict):
    def __init__(__self__, *,
                 business_status: str,
                 create_time: str,
                 ddos_status: str,
                 expired_time: str,
                 id: str,
                 image_id: str,
                 inner_ip_address: str,
                 instance_id: str,
                 instance_name: str,
                 payment_type: str,
                 plan_id: str,
                 public_ip_address: str,
                 status: str):
        """
        :param str business_status: The billing status of the simple application server. Valid values: `Normal`, `Expired` and `Overdue`.
        :param str create_time: The time when the simple application server was created.
        :param str ddos_status: The DDoS protection status. Valid values: `Normal`, `BlackHole`, and `Defense`.
        :param str expired_time: The time when the simple application server expires.
        :param str id: The ID of the Instance.
        :param str image_id: The ID of the simple application server Image.
        :param str inner_ip_address: The internal IP address of the simple application server.
        :param str instance_id: The ID of the simple application server.
        :param str instance_name: The name of the resource.
        :param str payment_type: The billing method of the simple application server.
        :param str plan_id: The ID of the simple application server plan.
        :param str public_ip_address: The public IP address of the simple application server.
        :param str status: The status of the resource.
        """
        pulumi.set(__self__, "business_status", business_status)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "ddos_status", ddos_status)
        pulumi.set(__self__, "expired_time", expired_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "image_id", image_id)
        pulumi.set(__self__, "inner_ip_address", inner_ip_address)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "payment_type", payment_type)
        pulumi.set(__self__, "plan_id", plan_id)
        pulumi.set(__self__, "public_ip_address", public_ip_address)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> str:
        """
        The billing status of the simple application server. Valid values: `Normal`, `Expired` and `Overdue`.
        """
        return pulumi.get(self, "business_status")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the simple application server was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="ddosStatus")
    def ddos_status(self) -> str:
        """
        The DDoS protection status. Valid values: `Normal`, `BlackHole`, and `Defense`.
        """
        return pulumi.get(self, "ddos_status")

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> str:
        """
        The time when the simple application server expires.
        """
        return pulumi.get(self, "expired_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Instance.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        """
        The ID of the simple application server Image.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="innerIpAddress")
    def inner_ip_address(self) -> str:
        """
        The internal IP address of the simple application server.
        """
        return pulumi.get(self, "inner_ip_address")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The ID of the simple application server.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> str:
        """
        The billing method of the simple application server.
        """
        return pulumi.get(self, "payment_type")

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> str:
        """
        The ID of the simple application server plan.
        """
        return pulumi.get(self, "plan_id")

    @property
    @pulumi.getter(name="publicIpAddress")
    def public_ip_address(self) -> str:
        """
        The public IP address of the simple application server.
        """
        return pulumi.get(self, "public_ip_address")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetServerPlansPlanResult(dict):
    def __init__(__self__, *,
                 bandwidth: int,
                 core: int,
                 disk_size: int,
                 flow: int,
                 id: str,
                 memory: int,
                 plan_id: str):
        """
        :param int bandwidth: The peak bandwidth. Unit: Mbit/s.
        :param int core: The number of CPU cores.
        :param int disk_size: The size of the enhanced SSD (ESSD). Unit: GB.
        :param int flow: The monthly data transfer quota. Unit: GB.
        :param str id: The ID of the Instance Plan.
        :param int memory: The memory size. Unit: GB.
        :param str plan_id: The ID of the Instance Plan.
        """
        pulumi.set(__self__, "bandwidth", bandwidth)
        pulumi.set(__self__, "core", core)
        pulumi.set(__self__, "disk_size", disk_size)
        pulumi.set(__self__, "flow", flow)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "memory", memory)
        pulumi.set(__self__, "plan_id", plan_id)

    @property
    @pulumi.getter
    def bandwidth(self) -> int:
        """
        The peak bandwidth. Unit: Mbit/s.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def core(self) -> int:
        """
        The number of CPU cores.
        """
        return pulumi.get(self, "core")

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> int:
        """
        The size of the enhanced SSD (ESSD). Unit: GB.
        """
        return pulumi.get(self, "disk_size")

    @property
    @pulumi.getter
    def flow(self) -> int:
        """
        The monthly data transfer quota. Unit: GB.
        """
        return pulumi.get(self, "flow")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Instance Plan.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def memory(self) -> int:
        """
        The memory size. Unit: GB.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> str:
        """
        The ID of the Instance Plan.
        """
        return pulumi.get(self, "plan_id")


