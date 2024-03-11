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
from ._inputs import *

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 key_num: pulumi.Input[int],
                 secret_num: pulumi.Input[int],
                 spec: pulumi.Input[int],
                 vpc_id: pulumi.Input[str],
                 vpc_num: pulumi.Input[int],
                 vswitch_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 zone_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 bind_vpcs: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceBindVpcArgs']]]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_storage: Optional[pulumi.Input[int]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 product_version: Optional[pulumi.Input[str]] = None,
                 renew_period: Optional[pulumi.Input[int]] = None,
                 renew_status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[int] key_num: Maximum number of stored keys.
        :param pulumi.Input[int] secret_num: Maximum number of Secrets.
        :param pulumi.Input[int] spec: The computation performance level of the KMS instance.
        :param pulumi.Input[str] vpc_id: Instance VPC id.
        :param pulumi.Input[int] vpc_num: The number of managed accesses. The maximum number of VPCs that can access this KMS instance.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: Instance bind vswitches.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_ids: zone id.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceBindVpcArgs']]] bind_vpcs: Aucillary VPCs used to access this KMS instance. See `bind_vpcs` below.
        :param pulumi.Input[str] log: Instance Audit Log Switch.
        :param pulumi.Input[int] log_storage: Instance log capacity.
        :param pulumi.Input[int] period: Purchase cycle, in months.
        :param pulumi.Input[str] product_version: KMS Instance commodity type (software/hardware).
        :param pulumi.Input[int] renew_period: Automatic renewal period, in months.
        :param pulumi.Input[str] renew_status: Renewal options (manual renewal, automatic renewal, no renewal).
        """
        pulumi.set(__self__, "key_num", key_num)
        pulumi.set(__self__, "secret_num", secret_num)
        pulumi.set(__self__, "spec", spec)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vpc_num", vpc_num)
        pulumi.set(__self__, "vswitch_ids", vswitch_ids)
        pulumi.set(__self__, "zone_ids", zone_ids)
        if bind_vpcs is not None:
            pulumi.set(__self__, "bind_vpcs", bind_vpcs)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if log_storage is not None:
            pulumi.set(__self__, "log_storage", log_storage)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if product_version is not None:
            pulumi.set(__self__, "product_version", product_version)
        if renew_period is not None:
            pulumi.set(__self__, "renew_period", renew_period)
        if renew_status is not None:
            pulumi.set(__self__, "renew_status", renew_status)

    @property
    @pulumi.getter(name="keyNum")
    def key_num(self) -> pulumi.Input[int]:
        """
        Maximum number of stored keys.
        """
        return pulumi.get(self, "key_num")

    @key_num.setter
    def key_num(self, value: pulumi.Input[int]):
        pulumi.set(self, "key_num", value)

    @property
    @pulumi.getter(name="secretNum")
    def secret_num(self) -> pulumi.Input[int]:
        """
        Maximum number of Secrets.
        """
        return pulumi.get(self, "secret_num")

    @secret_num.setter
    def secret_num(self, value: pulumi.Input[int]):
        pulumi.set(self, "secret_num", value)

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Input[int]:
        """
        The computation performance level of the KMS instance.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: pulumi.Input[int]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        Instance VPC id.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcNum")
    def vpc_num(self) -> pulumi.Input[int]:
        """
        The number of managed accesses. The maximum number of VPCs that can access this KMS instance.
        """
        return pulumi.get(self, "vpc_num")

    @vpc_num.setter
    def vpc_num(self, value: pulumi.Input[int]):
        pulumi.set(self, "vpc_num", value)

    @property
    @pulumi.getter(name="vswitchIds")
    def vswitch_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Instance bind vswitches.
        """
        return pulumi.get(self, "vswitch_ids")

    @vswitch_ids.setter
    def vswitch_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "vswitch_ids", value)

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        zone id.
        """
        return pulumi.get(self, "zone_ids")

    @zone_ids.setter
    def zone_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "zone_ids", value)

    @property
    @pulumi.getter(name="bindVpcs")
    def bind_vpcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceBindVpcArgs']]]]:
        """
        Aucillary VPCs used to access this KMS instance. See `bind_vpcs` below.
        """
        return pulumi.get(self, "bind_vpcs")

    @bind_vpcs.setter
    def bind_vpcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceBindVpcArgs']]]]):
        pulumi.set(self, "bind_vpcs", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Instance Audit Log Switch.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter(name="logStorage")
    def log_storage(self) -> Optional[pulumi.Input[int]]:
        """
        Instance log capacity.
        """
        return pulumi.get(self, "log_storage")

    @log_storage.setter
    def log_storage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "log_storage", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Purchase cycle, in months.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="productVersion")
    def product_version(self) -> Optional[pulumi.Input[str]]:
        """
        KMS Instance commodity type (software/hardware).
        """
        return pulumi.get(self, "product_version")

    @product_version.setter
    def product_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_version", value)

    @property
    @pulumi.getter(name="renewPeriod")
    def renew_period(self) -> Optional[pulumi.Input[int]]:
        """
        Automatic renewal period, in months.
        """
        return pulumi.get(self, "renew_period")

    @renew_period.setter
    def renew_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "renew_period", value)

    @property
    @pulumi.getter(name="renewStatus")
    def renew_status(self) -> Optional[pulumi.Input[str]]:
        """
        Renewal options (manual renewal, automatic renewal, no renewal).
        """
        return pulumi.get(self, "renew_status")

    @renew_status.setter
    def renew_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "renew_status", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 bind_vpcs: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceBindVpcArgs']]]] = None,
                 ca_certificate_chain_pem: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 key_num: Optional[pulumi.Input[int]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_storage: Optional[pulumi.Input[int]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 product_version: Optional[pulumi.Input[str]] = None,
                 renew_period: Optional[pulumi.Input[int]] = None,
                 renew_status: Optional[pulumi.Input[str]] = None,
                 secret_num: Optional[pulumi.Input[int]] = None,
                 spec: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_num: Optional[pulumi.Input[int]] = None,
                 vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceBindVpcArgs']]] bind_vpcs: Aucillary VPCs used to access this KMS instance. See `bind_vpcs` below.
        :param pulumi.Input[str] ca_certificate_chain_pem: KMS instance certificate chain in PEM format.
        :param pulumi.Input[str] create_time: The creation time of the resource.
        :param pulumi.Input[str] instance_name: The name of the resource.
        :param pulumi.Input[int] key_num: Maximum number of stored keys.
        :param pulumi.Input[str] log: Instance Audit Log Switch.
        :param pulumi.Input[int] log_storage: Instance log capacity.
        :param pulumi.Input[int] period: Purchase cycle, in months.
        :param pulumi.Input[str] product_version: KMS Instance commodity type (software/hardware).
        :param pulumi.Input[int] renew_period: Automatic renewal period, in months.
        :param pulumi.Input[str] renew_status: Renewal options (manual renewal, automatic renewal, no renewal).
        :param pulumi.Input[int] secret_num: Maximum number of Secrets.
        :param pulumi.Input[int] spec: The computation performance level of the KMS instance.
        :param pulumi.Input[str] status: Instance status.
        :param pulumi.Input[str] vpc_id: Instance VPC id.
        :param pulumi.Input[int] vpc_num: The number of managed accesses. The maximum number of VPCs that can access this KMS instance.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: Instance bind vswitches.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_ids: zone id.
        """
        if bind_vpcs is not None:
            pulumi.set(__self__, "bind_vpcs", bind_vpcs)
        if ca_certificate_chain_pem is not None:
            pulumi.set(__self__, "ca_certificate_chain_pem", ca_certificate_chain_pem)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if key_num is not None:
            pulumi.set(__self__, "key_num", key_num)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if log_storage is not None:
            pulumi.set(__self__, "log_storage", log_storage)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if product_version is not None:
            pulumi.set(__self__, "product_version", product_version)
        if renew_period is not None:
            pulumi.set(__self__, "renew_period", renew_period)
        if renew_status is not None:
            pulumi.set(__self__, "renew_status", renew_status)
        if secret_num is not None:
            pulumi.set(__self__, "secret_num", secret_num)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_num is not None:
            pulumi.set(__self__, "vpc_num", vpc_num)
        if vswitch_ids is not None:
            pulumi.set(__self__, "vswitch_ids", vswitch_ids)
        if zone_ids is not None:
            pulumi.set(__self__, "zone_ids", zone_ids)

    @property
    @pulumi.getter(name="bindVpcs")
    def bind_vpcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceBindVpcArgs']]]]:
        """
        Aucillary VPCs used to access this KMS instance. See `bind_vpcs` below.
        """
        return pulumi.get(self, "bind_vpcs")

    @bind_vpcs.setter
    def bind_vpcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceBindVpcArgs']]]]):
        pulumi.set(self, "bind_vpcs", value)

    @property
    @pulumi.getter(name="caCertificateChainPem")
    def ca_certificate_chain_pem(self) -> Optional[pulumi.Input[str]]:
        """
        KMS instance certificate chain in PEM format.
        """
        return pulumi.get(self, "ca_certificate_chain_pem")

    @ca_certificate_chain_pem.setter
    def ca_certificate_chain_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_certificate_chain_pem", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the resource.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="keyNum")
    def key_num(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of stored keys.
        """
        return pulumi.get(self, "key_num")

    @key_num.setter
    def key_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "key_num", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Instance Audit Log Switch.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter(name="logStorage")
    def log_storage(self) -> Optional[pulumi.Input[int]]:
        """
        Instance log capacity.
        """
        return pulumi.get(self, "log_storage")

    @log_storage.setter
    def log_storage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "log_storage", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Purchase cycle, in months.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="productVersion")
    def product_version(self) -> Optional[pulumi.Input[str]]:
        """
        KMS Instance commodity type (software/hardware).
        """
        return pulumi.get(self, "product_version")

    @product_version.setter
    def product_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_version", value)

    @property
    @pulumi.getter(name="renewPeriod")
    def renew_period(self) -> Optional[pulumi.Input[int]]:
        """
        Automatic renewal period, in months.
        """
        return pulumi.get(self, "renew_period")

    @renew_period.setter
    def renew_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "renew_period", value)

    @property
    @pulumi.getter(name="renewStatus")
    def renew_status(self) -> Optional[pulumi.Input[str]]:
        """
        Renewal options (manual renewal, automatic renewal, no renewal).
        """
        return pulumi.get(self, "renew_status")

    @renew_status.setter
    def renew_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "renew_status", value)

    @property
    @pulumi.getter(name="secretNum")
    def secret_num(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of Secrets.
        """
        return pulumi.get(self, "secret_num")

    @secret_num.setter
    def secret_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "secret_num", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input[int]]:
        """
        The computation performance level of the KMS instance.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Instance status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance VPC id.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcNum")
    def vpc_num(self) -> Optional[pulumi.Input[int]]:
        """
        The number of managed accesses. The maximum number of VPCs that can access this KMS instance.
        """
        return pulumi.get(self, "vpc_num")

    @vpc_num.setter
    def vpc_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vpc_num", value)

    @property
    @pulumi.getter(name="vswitchIds")
    def vswitch_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Instance bind vswitches.
        """
        return pulumi.get(self, "vswitch_ids")

    @vswitch_ids.setter
    def vswitch_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vswitch_ids", value)

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        zone id.
        """
        return pulumi.get(self, "zone_ids")

    @zone_ids.setter
    def zone_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "zone_ids", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bind_vpcs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceBindVpcArgs']]]]] = None,
                 key_num: Optional[pulumi.Input[int]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_storage: Optional[pulumi.Input[int]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 product_version: Optional[pulumi.Input[str]] = None,
                 renew_period: Optional[pulumi.Input[int]] = None,
                 renew_status: Optional[pulumi.Input[str]] = None,
                 secret_num: Optional[pulumi.Input[int]] = None,
                 spec: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_num: Optional[pulumi.Input[int]] = None,
                 vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a KMS Instance resource.

        For information about KMS Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/zh/key-management-service/latest/kms-instance-management).

        > **NOTE:** Available since v1.210.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_networks = alicloud.vpc.get_networks(name_regex="^default-NODELETING$",
            cidr_block="172.16.0.0/16")
        default_switches = alicloud.vpc.get_switches(vpc_id=default_networks.ids[0],
            zone_id="cn-hangzhou-k")
        default_instance = alicloud.kms.Instance("defaultInstance",
            product_version="3",
            vpc_id=default_networks.ids[0],
            zone_ids=[
                "cn-hangzhou-k",
                "cn-hangzhou-j",
            ],
            vswitch_ids=[default_switches.ids[0]],
            vpc_num=1,
            key_num=1000,
            secret_num=0,
            spec=1000)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        KMS Instance can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:kms/instance:Instance example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceBindVpcArgs']]]] bind_vpcs: Aucillary VPCs used to access this KMS instance. See `bind_vpcs` below.
        :param pulumi.Input[int] key_num: Maximum number of stored keys.
        :param pulumi.Input[str] log: Instance Audit Log Switch.
        :param pulumi.Input[int] log_storage: Instance log capacity.
        :param pulumi.Input[int] period: Purchase cycle, in months.
        :param pulumi.Input[str] product_version: KMS Instance commodity type (software/hardware).
        :param pulumi.Input[int] renew_period: Automatic renewal period, in months.
        :param pulumi.Input[str] renew_status: Renewal options (manual renewal, automatic renewal, no renewal).
        :param pulumi.Input[int] secret_num: Maximum number of Secrets.
        :param pulumi.Input[int] spec: The computation performance level of the KMS instance.
        :param pulumi.Input[str] vpc_id: Instance VPC id.
        :param pulumi.Input[int] vpc_num: The number of managed accesses. The maximum number of VPCs that can access this KMS instance.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: Instance bind vswitches.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_ids: zone id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a KMS Instance resource.

        For information about KMS Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/zh/key-management-service/latest/kms-instance-management).

        > **NOTE:** Available since v1.210.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_networks = alicloud.vpc.get_networks(name_regex="^default-NODELETING$",
            cidr_block="172.16.0.0/16")
        default_switches = alicloud.vpc.get_switches(vpc_id=default_networks.ids[0],
            zone_id="cn-hangzhou-k")
        default_instance = alicloud.kms.Instance("defaultInstance",
            product_version="3",
            vpc_id=default_networks.ids[0],
            zone_ids=[
                "cn-hangzhou-k",
                "cn-hangzhou-j",
            ],
            vswitch_ids=[default_switches.ids[0]],
            vpc_num=1,
            key_num=1000,
            secret_num=0,
            spec=1000)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        KMS Instance can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:kms/instance:Instance example <id>
        ```

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bind_vpcs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceBindVpcArgs']]]]] = None,
                 key_num: Optional[pulumi.Input[int]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 log_storage: Optional[pulumi.Input[int]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 product_version: Optional[pulumi.Input[str]] = None,
                 renew_period: Optional[pulumi.Input[int]] = None,
                 renew_status: Optional[pulumi.Input[str]] = None,
                 secret_num: Optional[pulumi.Input[int]] = None,
                 spec: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_num: Optional[pulumi.Input[int]] = None,
                 vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["bind_vpcs"] = bind_vpcs
            if key_num is None and not opts.urn:
                raise TypeError("Missing required property 'key_num'")
            __props__.__dict__["key_num"] = key_num
            __props__.__dict__["log"] = log
            __props__.__dict__["log_storage"] = log_storage
            __props__.__dict__["period"] = period
            __props__.__dict__["product_version"] = product_version
            __props__.__dict__["renew_period"] = renew_period
            __props__.__dict__["renew_status"] = renew_status
            if secret_num is None and not opts.urn:
                raise TypeError("Missing required property 'secret_num'")
            __props__.__dict__["secret_num"] = secret_num
            if spec is None and not opts.urn:
                raise TypeError("Missing required property 'spec'")
            __props__.__dict__["spec"] = spec
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            if vpc_num is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_num'")
            __props__.__dict__["vpc_num"] = vpc_num
            if vswitch_ids is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_ids'")
            __props__.__dict__["vswitch_ids"] = vswitch_ids
            if zone_ids is None and not opts.urn:
                raise TypeError("Missing required property 'zone_ids'")
            __props__.__dict__["zone_ids"] = zone_ids
            __props__.__dict__["ca_certificate_chain_pem"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["instance_name"] = None
            __props__.__dict__["status"] = None
        super(Instance, __self__).__init__(
            'alicloud:kms/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bind_vpcs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceBindVpcArgs']]]]] = None,
            ca_certificate_chain_pem: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            key_num: Optional[pulumi.Input[int]] = None,
            log: Optional[pulumi.Input[str]] = None,
            log_storage: Optional[pulumi.Input[int]] = None,
            period: Optional[pulumi.Input[int]] = None,
            product_version: Optional[pulumi.Input[str]] = None,
            renew_period: Optional[pulumi.Input[int]] = None,
            renew_status: Optional[pulumi.Input[str]] = None,
            secret_num: Optional[pulumi.Input[int]] = None,
            spec: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vpc_num: Optional[pulumi.Input[int]] = None,
            vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceBindVpcArgs']]]] bind_vpcs: Aucillary VPCs used to access this KMS instance. See `bind_vpcs` below.
        :param pulumi.Input[str] ca_certificate_chain_pem: KMS instance certificate chain in PEM format.
        :param pulumi.Input[str] create_time: The creation time of the resource.
        :param pulumi.Input[str] instance_name: The name of the resource.
        :param pulumi.Input[int] key_num: Maximum number of stored keys.
        :param pulumi.Input[str] log: Instance Audit Log Switch.
        :param pulumi.Input[int] log_storage: Instance log capacity.
        :param pulumi.Input[int] period: Purchase cycle, in months.
        :param pulumi.Input[str] product_version: KMS Instance commodity type (software/hardware).
        :param pulumi.Input[int] renew_period: Automatic renewal period, in months.
        :param pulumi.Input[str] renew_status: Renewal options (manual renewal, automatic renewal, no renewal).
        :param pulumi.Input[int] secret_num: Maximum number of Secrets.
        :param pulumi.Input[int] spec: The computation performance level of the KMS instance.
        :param pulumi.Input[str] status: Instance status.
        :param pulumi.Input[str] vpc_id: Instance VPC id.
        :param pulumi.Input[int] vpc_num: The number of managed accesses. The maximum number of VPCs that can access this KMS instance.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: Instance bind vswitches.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_ids: zone id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["bind_vpcs"] = bind_vpcs
        __props__.__dict__["ca_certificate_chain_pem"] = ca_certificate_chain_pem
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["key_num"] = key_num
        __props__.__dict__["log"] = log
        __props__.__dict__["log_storage"] = log_storage
        __props__.__dict__["period"] = period
        __props__.__dict__["product_version"] = product_version
        __props__.__dict__["renew_period"] = renew_period
        __props__.__dict__["renew_status"] = renew_status
        __props__.__dict__["secret_num"] = secret_num
        __props__.__dict__["spec"] = spec
        __props__.__dict__["status"] = status
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vpc_num"] = vpc_num
        __props__.__dict__["vswitch_ids"] = vswitch_ids
        __props__.__dict__["zone_ids"] = zone_ids
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bindVpcs")
    def bind_vpcs(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceBindVpc']]]:
        """
        Aucillary VPCs used to access this KMS instance. See `bind_vpcs` below.
        """
        return pulumi.get(self, "bind_vpcs")

    @property
    @pulumi.getter(name="caCertificateChainPem")
    def ca_certificate_chain_pem(self) -> pulumi.Output[str]:
        """
        KMS instance certificate chain in PEM format.
        """
        return pulumi.get(self, "ca_certificate_chain_pem")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of the resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="keyNum")
    def key_num(self) -> pulumi.Output[int]:
        """
        Maximum number of stored keys.
        """
        return pulumi.get(self, "key_num")

    @property
    @pulumi.getter
    def log(self) -> pulumi.Output[str]:
        """
        Instance Audit Log Switch.
        """
        return pulumi.get(self, "log")

    @property
    @pulumi.getter(name="logStorage")
    def log_storage(self) -> pulumi.Output[int]:
        """
        Instance log capacity.
        """
        return pulumi.get(self, "log_storage")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        """
        Purchase cycle, in months.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="productVersion")
    def product_version(self) -> pulumi.Output[Optional[str]]:
        """
        KMS Instance commodity type (software/hardware).
        """
        return pulumi.get(self, "product_version")

    @property
    @pulumi.getter(name="renewPeriod")
    def renew_period(self) -> pulumi.Output[Optional[int]]:
        """
        Automatic renewal period, in months.
        """
        return pulumi.get(self, "renew_period")

    @property
    @pulumi.getter(name="renewStatus")
    def renew_status(self) -> pulumi.Output[Optional[str]]:
        """
        Renewal options (manual renewal, automatic renewal, no renewal).
        """
        return pulumi.get(self, "renew_status")

    @property
    @pulumi.getter(name="secretNum")
    def secret_num(self) -> pulumi.Output[int]:
        """
        Maximum number of Secrets.
        """
        return pulumi.get(self, "secret_num")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output[int]:
        """
        The computation performance level of the KMS instance.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Instance status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        Instance VPC id.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcNum")
    def vpc_num(self) -> pulumi.Output[int]:
        """
        The number of managed accesses. The maximum number of VPCs that can access this KMS instance.
        """
        return pulumi.get(self, "vpc_num")

    @property
    @pulumi.getter(name="vswitchIds")
    def vswitch_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Instance bind vswitches.
        """
        return pulumi.get(self, "vswitch_ids")

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        zone id.
        """
        return pulumi.get(self, "zone_ids")

