# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Instance']


class Instance(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 auto_renew_period: Optional[pulumi.Input[float]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 backup_id: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 kms_encrypted_password: Optional[pulumi.Input[str]] = None,
                 kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 maintain_end_time: Optional[pulumi.Input[str]] = None,
                 maintain_start_time: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['InstanceParameterArgs']]]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[float]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 security_ips: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_auth_mode: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an ApsaraDB Redis / Memcache instance resource. A DB instance is an isolated database environment in the cloud. It can be associated with IP whitelists and backup configuration which are separate resource providers.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        creation = config.get("creation")
        if creation is None:
            creation = "KVStore"
        name = config.get("name")
        if name is None:
            name = "kvstoreinstancevpc"
        default_zones = alicloud.get_zones(available_resource_creation=creation)
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            availability_zone=default_zones.zones[0].id,
            cidr_block="172.16.0.0/24",
            vpc_id=default_network.id)
        default_instance = alicloud.kvstore.Instance("defaultInstance",
            engine_version="4.0",
            instance_class="redis.master.small.default",
            instance_name=name,
            instance_type="Redis",
            private_ip="172.16.0.10",
            security_ips=["10.0.0.1"],
            vswitch_id=default_switch.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: Whether to renewal a DB instance automatically or not. It is valid when instance_charge_type is `PrePaid`. Default to `false`.
        :param pulumi.Input[float] auto_renew_period: Auto-renewal period of an instance, in the unit of the month. It is valid when instance_charge_type is `PrePaid`. Valid value:[1~12], Default to 1.
        :param pulumi.Input[str] availability_zone: The Zone to launch the DB instance.
        :param pulumi.Input[str] backup_id: If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
        :param pulumi.Input[str] engine_version: Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
        :param pulumi.Input[str] instance_charge_type: Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
        :param pulumi.Input[str] instance_name: The name of DB instance. It a string of 2 to 256 characters.
        :param pulumi.Input[str] instance_type: The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
        :param pulumi.Input[Mapping[str, Any]] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        :param pulumi.Input[str] maintain_end_time: The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        :param pulumi.Input[str] maintain_start_time: The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['InstanceParameterArgs']]]] parameters: Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
        :param pulumi.Input[str] password: The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
        :param pulumi.Input[float] period: The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        :param pulumi.Input[str] private_ip: Set the instance's private IP.
        :param pulumi.Input[str] resource_group_id: The ID of resource group which the resource belongs.
        :param pulumi.Input[str] security_group_id: The Security Group ID of ECS.
        :param pulumi.Input[List[pulumi.Input[str]]] security_ips: Set the instance's IP whitelist of the default security group.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_auth_mode: Only meaningful if instance_type is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
        :param pulumi.Input[str] vswitch_id: The ID of VSwitch.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['auto_renew'] = auto_renew
            __props__['auto_renew_period'] = auto_renew_period
            __props__['availability_zone'] = availability_zone
            __props__['backup_id'] = backup_id
            __props__['engine_version'] = engine_version
            __props__['instance_charge_type'] = instance_charge_type
            if instance_class is None:
                raise TypeError("Missing required property 'instance_class'")
            __props__['instance_class'] = instance_class
            __props__['instance_name'] = instance_name
            __props__['instance_type'] = instance_type
            __props__['kms_encrypted_password'] = kms_encrypted_password
            __props__['kms_encryption_context'] = kms_encryption_context
            __props__['maintain_end_time'] = maintain_end_time
            __props__['maintain_start_time'] = maintain_start_time
            __props__['parameters'] = parameters
            __props__['password'] = password
            __props__['period'] = period
            __props__['private_ip'] = private_ip
            __props__['resource_group_id'] = resource_group_id
            __props__['security_group_id'] = security_group_id
            __props__['security_ips'] = security_ips
            __props__['tags'] = tags
            __props__['vpc_auth_mode'] = vpc_auth_mode
            __props__['vswitch_id'] = vswitch_id
            __props__['connection_domain'] = None
        super(Instance, __self__).__init__(
            'alicloud:kvstore/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_renew: Optional[pulumi.Input[bool]] = None,
            auto_renew_period: Optional[pulumi.Input[float]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            backup_id: Optional[pulumi.Input[str]] = None,
            connection_domain: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            instance_charge_type: Optional[pulumi.Input[str]] = None,
            instance_class: Optional[pulumi.Input[str]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            kms_encrypted_password: Optional[pulumi.Input[str]] = None,
            kms_encryption_context: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            maintain_end_time: Optional[pulumi.Input[str]] = None,
            maintain_start_time: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['InstanceParameterArgs']]]]] = None,
            password: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[float]] = None,
            private_ip: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            security_ips: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc_auth_mode: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: Whether to renewal a DB instance automatically or not. It is valid when instance_charge_type is `PrePaid`. Default to `false`.
        :param pulumi.Input[float] auto_renew_period: Auto-renewal period of an instance, in the unit of the month. It is valid when instance_charge_type is `PrePaid`. Valid value:[1~12], Default to 1.
        :param pulumi.Input[str] availability_zone: The Zone to launch the DB instance.
        :param pulumi.Input[str] backup_id: If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
        :param pulumi.Input[str] connection_domain: Instance connection domain (only Intranet access supported).
        :param pulumi.Input[str] engine_version: Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
        :param pulumi.Input[str] instance_charge_type: Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
        :param pulumi.Input[str] instance_name: The name of DB instance. It a string of 2 to 256 characters.
        :param pulumi.Input[str] instance_type: The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
        :param pulumi.Input[Mapping[str, Any]] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        :param pulumi.Input[str] maintain_end_time: The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        :param pulumi.Input[str] maintain_start_time: The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['InstanceParameterArgs']]]] parameters: Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
        :param pulumi.Input[str] password: The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
        :param pulumi.Input[float] period: The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        :param pulumi.Input[str] private_ip: Set the instance's private IP.
        :param pulumi.Input[str] resource_group_id: The ID of resource group which the resource belongs.
        :param pulumi.Input[str] security_group_id: The Security Group ID of ECS.
        :param pulumi.Input[List[pulumi.Input[str]]] security_ips: Set the instance's IP whitelist of the default security group.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_auth_mode: Only meaningful if instance_type is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
        :param pulumi.Input[str] vswitch_id: The ID of VSwitch.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_renew"] = auto_renew
        __props__["auto_renew_period"] = auto_renew_period
        __props__["availability_zone"] = availability_zone
        __props__["backup_id"] = backup_id
        __props__["connection_domain"] = connection_domain
        __props__["engine_version"] = engine_version
        __props__["instance_charge_type"] = instance_charge_type
        __props__["instance_class"] = instance_class
        __props__["instance_name"] = instance_name
        __props__["instance_type"] = instance_type
        __props__["kms_encrypted_password"] = kms_encrypted_password
        __props__["kms_encryption_context"] = kms_encryption_context
        __props__["maintain_end_time"] = maintain_end_time
        __props__["maintain_start_time"] = maintain_start_time
        __props__["parameters"] = parameters
        __props__["password"] = password
        __props__["period"] = period
        __props__["private_ip"] = private_ip
        __props__["resource_group_id"] = resource_group_id
        __props__["security_group_id"] = security_group_id
        __props__["security_ips"] = security_ips
        __props__["tags"] = tags
        __props__["vpc_auth_mode"] = vpc_auth_mode
        __props__["vswitch_id"] = vswitch_id
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[bool]:
        """
        Whether to renewal a DB instance automatically or not. It is valid when instance_charge_type is `PrePaid`. Default to `false`.
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter(name="autoRenewPeriod")
    def auto_renew_period(self) -> Optional[float]:
        """
        Auto-renewal period of an instance, in the unit of the month. It is valid when instance_charge_type is `PrePaid`. Valid value:[1~12], Default to 1.
        """
        return pulumi.get(self, "auto_renew_period")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The Zone to launch the DB instance.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> Optional[str]:
        """
        If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
        """
        return pulumi.get(self, "backup_id")

    @property
    @pulumi.getter(name="connectionDomain")
    def connection_domain(self) -> str:
        """
        Instance connection domain (only Intranet access supported).
        """
        return pulumi.get(self, "connection_domain")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[str]:
        """
        Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> Optional[str]:
        """
        Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
        """
        return pulumi.get(self, "instance_charge_type")

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> str:
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[str]:
        """
        The name of DB instance. It a string of 2 to 256 characters.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        """
        The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="kmsEncryptedPassword")
    def kms_encrypted_password(self) -> Optional[str]:
        """
        An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
        """
        return pulumi.get(self, "kms_encrypted_password")

    @property
    @pulumi.getter(name="kmsEncryptionContext")
    def kms_encryption_context(self) -> Optional[Mapping[str, Any]]:
        """
        An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        return pulumi.get(self, "kms_encryption_context")

    @property
    @pulumi.getter(name="maintainEndTime")
    def maintain_end_time(self) -> str:
        """
        The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        """
        return pulumi.get(self, "maintain_end_time")

    @property
    @pulumi.getter(name="maintainStartTime")
    def maintain_start_time(self) -> str:
        """
        The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        """
        return pulumi.get(self, "maintain_start_time")

    @property
    @pulumi.getter
    def parameters(self) -> List['outputs.InstanceParameter']:
        """
        Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        """
        The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def period(self) -> Optional[float]:
        """
        The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        Set the instance's private IP.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        """
        The ID of resource group which the resource belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        """
        The Security Group ID of ECS.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="securityIps")
    def security_ips(self) -> List[str]:
        """
        Set the instance's IP whitelist of the default security group.
        """
        return pulumi.get(self, "security_ips")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcAuthMode")
    def vpc_auth_mode(self) -> str:
        """
        Only meaningful if instance_type is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
        """
        return pulumi.get(self, "vpc_auth_mode")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[str]:
        """
        The ID of VSwitch.
        """
        return pulumi.get(self, "vswitch_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

