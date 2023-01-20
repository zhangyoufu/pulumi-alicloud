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
    'ProjectIpWhiteList',
    'ProjectProperties',
    'ProjectPropertiesEncryption',
    'ProjectPropertiesTableLifecycle',
    'ProjectSecurityProperties',
    'ProjectSecurityPropertiesProjectProtection',
    'GetProjectsProjectResult',
    'GetProjectsProjectIpWhiteListResult',
    'GetProjectsProjectPropertiesResult',
    'GetProjectsProjectPropertiesEncryptionResult',
    'GetProjectsProjectPropertiesTableLifecycleResult',
    'GetProjectsProjectSecurityPropertiesResult',
    'GetProjectsProjectSecurityPropertiesProjectProtectionResult',
]

@pulumi.output_type
class ProjectIpWhiteList(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipList":
            suggest = "ip_list"
        elif key == "vpcIpList":
            suggest = "vpc_ip_list"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectIpWhiteList. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectIpWhiteList.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectIpWhiteList.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ip_list: Optional[str] = None,
                 vpc_ip_list: Optional[str] = None):
        """
        :param str ip_list: Classic network IP white list.
        :param str vpc_ip_list: VPC network whitelist.
        """
        if ip_list is not None:
            pulumi.set(__self__, "ip_list", ip_list)
        if vpc_ip_list is not None:
            pulumi.set(__self__, "vpc_ip_list", vpc_ip_list)

    @property
    @pulumi.getter(name="ipList")
    def ip_list(self) -> Optional[str]:
        """
        Classic network IP white list.
        """
        return pulumi.get(self, "ip_list")

    @property
    @pulumi.getter(name="vpcIpList")
    def vpc_ip_list(self) -> Optional[str]:
        """
        VPC network whitelist.
        """
        return pulumi.get(self, "vpc_ip_list")


@pulumi.output_type
class ProjectProperties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowFullScan":
            suggest = "allow_full_scan"
        elif key == "enableDecimal2":
            suggest = "enable_decimal2"
        elif key == "retentionDays":
            suggest = "retention_days"
        elif key == "sqlMeteringMax":
            suggest = "sql_metering_max"
        elif key == "tableLifecycle":
            suggest = "table_lifecycle"
        elif key == "typeSystem":
            suggest = "type_system"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectProperties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectProperties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectProperties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 allow_full_scan: Optional[bool] = None,
                 enable_decimal2: Optional[bool] = None,
                 encryption: Optional['outputs.ProjectPropertiesEncryption'] = None,
                 retention_days: Optional[int] = None,
                 sql_metering_max: Optional[str] = None,
                 table_lifecycle: Optional['outputs.ProjectPropertiesTableLifecycle'] = None,
                 timezone: Optional[str] = None,
                 type_system: Optional[str] = None):
        """
        :param bool allow_full_scan: Whether to allow full table scan.
        :param bool enable_decimal2: Whether to turn on Decimal2.0.
        :param 'ProjectPropertiesEncryptionArgs' encryption: Whether encryption is turned on.See the following `Block Encryption`.
        :param int retention_days: Job default retention time.
        :param str sql_metering_max: SQL charge limit.
        :param 'ProjectPropertiesTableLifecycleArgs' table_lifecycle: Life cycle of tables.See the following `Block TableLifecycle`.
        :param str timezone: Project time zone.
        :param str type_system: Type system.
        """
        if allow_full_scan is not None:
            pulumi.set(__self__, "allow_full_scan", allow_full_scan)
        if enable_decimal2 is not None:
            pulumi.set(__self__, "enable_decimal2", enable_decimal2)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if sql_metering_max is not None:
            pulumi.set(__self__, "sql_metering_max", sql_metering_max)
        if table_lifecycle is not None:
            pulumi.set(__self__, "table_lifecycle", table_lifecycle)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)
        if type_system is not None:
            pulumi.set(__self__, "type_system", type_system)

    @property
    @pulumi.getter(name="allowFullScan")
    def allow_full_scan(self) -> Optional[bool]:
        """
        Whether to allow full table scan.
        """
        return pulumi.get(self, "allow_full_scan")

    @property
    @pulumi.getter(name="enableDecimal2")
    def enable_decimal2(self) -> Optional[bool]:
        """
        Whether to turn on Decimal2.0.
        """
        return pulumi.get(self, "enable_decimal2")

    @property
    @pulumi.getter
    def encryption(self) -> Optional['outputs.ProjectPropertiesEncryption']:
        """
        Whether encryption is turned on.See the following `Block Encryption`.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[int]:
        """
        Job default retention time.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter(name="sqlMeteringMax")
    def sql_metering_max(self) -> Optional[str]:
        """
        SQL charge limit.
        """
        return pulumi.get(self, "sql_metering_max")

    @property
    @pulumi.getter(name="tableLifecycle")
    def table_lifecycle(self) -> Optional['outputs.ProjectPropertiesTableLifecycle']:
        """
        Life cycle of tables.See the following `Block TableLifecycle`.
        """
        return pulumi.get(self, "table_lifecycle")

    @property
    @pulumi.getter
    def timezone(self) -> Optional[str]:
        """
        Project time zone.
        """
        return pulumi.get(self, "timezone")

    @property
    @pulumi.getter(name="typeSystem")
    def type_system(self) -> Optional[str]:
        """
        Type system.
        """
        return pulumi.get(self, "type_system")


@pulumi.output_type
class ProjectPropertiesEncryption(dict):
    def __init__(__self__, *,
                 algorithm: Optional[str] = None,
                 enable: Optional[bool] = None,
                 key: Optional[str] = None):
        """
        :param str algorithm: Algorithm.
        :param bool enable: Whether to open.
        :param str key: Encryption algorithm key.
        """
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if key is not None:
            pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[str]:
        """
        Algorithm.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def enable(self) -> Optional[bool]:
        """
        Whether to open.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        """
        Encryption algorithm key.
        """
        return pulumi.get(self, "key")


@pulumi.output_type
class ProjectPropertiesTableLifecycle(dict):
    def __init__(__self__, *,
                 type: Optional[str] = None,
                 value: Optional[str] = None):
        """
        :param str type: Life cycle type.
        :param str value: The value of the life cycle.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Life cycle type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The value of the life cycle.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class ProjectSecurityProperties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "enableDownloadPrivilege":
            suggest = "enable_download_privilege"
        elif key == "labelSecurity":
            suggest = "label_security"
        elif key == "objectCreatorHasAccessPermission":
            suggest = "object_creator_has_access_permission"
        elif key == "objectCreatorHasGrantPermission":
            suggest = "object_creator_has_grant_permission"
        elif key == "projectProtection":
            suggest = "project_protection"
        elif key == "usingAcl":
            suggest = "using_acl"
        elif key == "usingPolicy":
            suggest = "using_policy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectSecurityProperties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectSecurityProperties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectSecurityProperties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enable_download_privilege: Optional[bool] = None,
                 label_security: Optional[bool] = None,
                 object_creator_has_access_permission: Optional[bool] = None,
                 object_creator_has_grant_permission: Optional[bool] = None,
                 project_protection: Optional['outputs.ProjectSecurityPropertiesProjectProtection'] = None,
                 using_acl: Optional[bool] = None,
                 using_policy: Optional[bool] = None):
        """
        :param bool enable_download_privilege: Whether to enable download permission check.
        :param bool label_security: Label authorization.
        :param bool object_creator_has_access_permission: Project creator permissions.
        :param bool object_creator_has_grant_permission: Does the project creator have authorization rights.
        :param 'ProjectSecurityPropertiesProjectProtectionArgs' project_protection: Project protection.See the following `Block ProjectProtection`.
        :param bool using_acl: Whether to turn on ACL.
        :param bool using_policy: Whether to enable Policy.
        """
        if enable_download_privilege is not None:
            pulumi.set(__self__, "enable_download_privilege", enable_download_privilege)
        if label_security is not None:
            pulumi.set(__self__, "label_security", label_security)
        if object_creator_has_access_permission is not None:
            pulumi.set(__self__, "object_creator_has_access_permission", object_creator_has_access_permission)
        if object_creator_has_grant_permission is not None:
            pulumi.set(__self__, "object_creator_has_grant_permission", object_creator_has_grant_permission)
        if project_protection is not None:
            pulumi.set(__self__, "project_protection", project_protection)
        if using_acl is not None:
            pulumi.set(__self__, "using_acl", using_acl)
        if using_policy is not None:
            pulumi.set(__self__, "using_policy", using_policy)

    @property
    @pulumi.getter(name="enableDownloadPrivilege")
    def enable_download_privilege(self) -> Optional[bool]:
        """
        Whether to enable download permission check.
        """
        return pulumi.get(self, "enable_download_privilege")

    @property
    @pulumi.getter(name="labelSecurity")
    def label_security(self) -> Optional[bool]:
        """
        Label authorization.
        """
        return pulumi.get(self, "label_security")

    @property
    @pulumi.getter(name="objectCreatorHasAccessPermission")
    def object_creator_has_access_permission(self) -> Optional[bool]:
        """
        Project creator permissions.
        """
        return pulumi.get(self, "object_creator_has_access_permission")

    @property
    @pulumi.getter(name="objectCreatorHasGrantPermission")
    def object_creator_has_grant_permission(self) -> Optional[bool]:
        """
        Does the project creator have authorization rights.
        """
        return pulumi.get(self, "object_creator_has_grant_permission")

    @property
    @pulumi.getter(name="projectProtection")
    def project_protection(self) -> Optional['outputs.ProjectSecurityPropertiesProjectProtection']:
        """
        Project protection.See the following `Block ProjectProtection`.
        """
        return pulumi.get(self, "project_protection")

    @property
    @pulumi.getter(name="usingAcl")
    def using_acl(self) -> Optional[bool]:
        """
        Whether to turn on ACL.
        """
        return pulumi.get(self, "using_acl")

    @property
    @pulumi.getter(name="usingPolicy")
    def using_policy(self) -> Optional[bool]:
        """
        Whether to enable Policy.
        """
        return pulumi.get(self, "using_policy")


@pulumi.output_type
class ProjectSecurityPropertiesProjectProtection(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "exceptionPolicy":
            suggest = "exception_policy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectSecurityPropertiesProjectProtection. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectSecurityPropertiesProjectProtection.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectSecurityPropertiesProjectProtection.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 exception_policy: Optional[str] = None,
                 protected: Optional[bool] = None):
        """
        :param str exception_policy: Exclusion policy.
        :param bool protected: Is it turned on.
        """
        if exception_policy is not None:
            pulumi.set(__self__, "exception_policy", exception_policy)
        if protected is not None:
            pulumi.set(__self__, "protected", protected)

    @property
    @pulumi.getter(name="exceptionPolicy")
    def exception_policy(self) -> Optional[str]:
        """
        Exclusion policy.
        """
        return pulumi.get(self, "exception_policy")

    @property
    @pulumi.getter
    def protected(self) -> Optional[bool]:
        """
        Is it turned on.
        """
        return pulumi.get(self, "protected")


@pulumi.output_type
class GetProjectsProjectResult(dict):
    def __init__(__self__, *,
                 comment: str,
                 default_quota: str,
                 id: str,
                 ip_white_list: 'outputs.GetProjectsProjectIpWhiteListResult',
                 owner: str,
                 project_name: str,
                 properties: 'outputs.GetProjectsProjectPropertiesResult',
                 security_properties: 'outputs.GetProjectsProjectSecurityPropertiesResult',
                 status: str,
                 type: str):
        """
        :param str default_quota: Default Computing Resource Group
        :param str id: Project ID. The value is the same as `project_name`.
        :param 'GetProjectsProjectIpWhiteListArgs' ip_white_list: IP whitelist
        :param str owner: Project owner
        :param str project_name: The name of the resource
        :param 'GetProjectsProjectPropertiesArgs' properties: Project base attributes
        :param 'GetProjectsProjectSecurityPropertiesArgs' security_properties: Security-related attributes
        :param str status: The status of the resource
        :param str type: Project type
        """
        pulumi.set(__self__, "comment", comment)
        pulumi.set(__self__, "default_quota", default_quota)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip_white_list", ip_white_list)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "project_name", project_name)
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "security_properties", security_properties)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def comment(self) -> str:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="defaultQuota")
    def default_quota(self) -> str:
        """
        Default Computing Resource Group
        """
        return pulumi.get(self, "default_quota")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Project ID. The value is the same as `project_name`.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipWhiteList")
    def ip_white_list(self) -> 'outputs.GetProjectsProjectIpWhiteListResult':
        """
        IP whitelist
        """
        return pulumi.get(self, "ip_white_list")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        Project owner
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.GetProjectsProjectPropertiesResult':
        """
        Project base attributes
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="securityProperties")
    def security_properties(self) -> 'outputs.GetProjectsProjectSecurityPropertiesResult':
        """
        Security-related attributes
        """
        return pulumi.get(self, "security_properties")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the resource
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Project type
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetProjectsProjectIpWhiteListResult(dict):
    def __init__(__self__, *,
                 ip_list: str,
                 vpc_ip_list: str):
        """
        :param str ip_list: Classic network IP white list.
        :param str vpc_ip_list: VPC network whitelist.
        """
        pulumi.set(__self__, "ip_list", ip_list)
        pulumi.set(__self__, "vpc_ip_list", vpc_ip_list)

    @property
    @pulumi.getter(name="ipList")
    def ip_list(self) -> str:
        """
        Classic network IP white list.
        """
        return pulumi.get(self, "ip_list")

    @property
    @pulumi.getter(name="vpcIpList")
    def vpc_ip_list(self) -> str:
        """
        VPC network whitelist.
        """
        return pulumi.get(self, "vpc_ip_list")


@pulumi.output_type
class GetProjectsProjectPropertiesResult(dict):
    def __init__(__self__, *,
                 allow_full_scan: bool,
                 enable_decimal2: bool,
                 encryption: 'outputs.GetProjectsProjectPropertiesEncryptionResult',
                 retention_days: str,
                 sql_metering_max: str,
                 table_lifecycle: 'outputs.GetProjectsProjectPropertiesTableLifecycleResult',
                 timezone: str,
                 type_system: str):
        """
        :param bool allow_full_scan: Whether to allow full table scan.
        :param bool enable_decimal2: Whether to turn on Decimal2.0.
        :param 'GetProjectsProjectPropertiesEncryptionArgs' encryption: Whether encryption is turned on.
        :param str retention_days: Job default retention time.
        :param str sql_metering_max: SQL charge limit.
        :param 'GetProjectsProjectPropertiesTableLifecycleArgs' table_lifecycle: Life cycle of tables.
        :param str timezone: Project time zone.
        :param str type_system: Type system.
        """
        pulumi.set(__self__, "allow_full_scan", allow_full_scan)
        pulumi.set(__self__, "enable_decimal2", enable_decimal2)
        pulumi.set(__self__, "encryption", encryption)
        pulumi.set(__self__, "retention_days", retention_days)
        pulumi.set(__self__, "sql_metering_max", sql_metering_max)
        pulumi.set(__self__, "table_lifecycle", table_lifecycle)
        pulumi.set(__self__, "timezone", timezone)
        pulumi.set(__self__, "type_system", type_system)

    @property
    @pulumi.getter(name="allowFullScan")
    def allow_full_scan(self) -> bool:
        """
        Whether to allow full table scan.
        """
        return pulumi.get(self, "allow_full_scan")

    @property
    @pulumi.getter(name="enableDecimal2")
    def enable_decimal2(self) -> bool:
        """
        Whether to turn on Decimal2.0.
        """
        return pulumi.get(self, "enable_decimal2")

    @property
    @pulumi.getter
    def encryption(self) -> 'outputs.GetProjectsProjectPropertiesEncryptionResult':
        """
        Whether encryption is turned on.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> str:
        """
        Job default retention time.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter(name="sqlMeteringMax")
    def sql_metering_max(self) -> str:
        """
        SQL charge limit.
        """
        return pulumi.get(self, "sql_metering_max")

    @property
    @pulumi.getter(name="tableLifecycle")
    def table_lifecycle(self) -> 'outputs.GetProjectsProjectPropertiesTableLifecycleResult':
        """
        Life cycle of tables.
        """
        return pulumi.get(self, "table_lifecycle")

    @property
    @pulumi.getter
    def timezone(self) -> str:
        """
        Project time zone.
        """
        return pulumi.get(self, "timezone")

    @property
    @pulumi.getter(name="typeSystem")
    def type_system(self) -> str:
        """
        Type system.
        """
        return pulumi.get(self, "type_system")


@pulumi.output_type
class GetProjectsProjectPropertiesEncryptionResult(dict):
    def __init__(__self__, *,
                 algorithm: str,
                 enable: bool,
                 key: str):
        """
        :param str algorithm: Algorithm.
        :param bool enable: Whether to open.
        :param str key: Encryption algorithm key.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "enable", enable)
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        Algorithm.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def enable(self) -> bool:
        """
        Whether to open.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Encryption algorithm key.
        """
        return pulumi.get(self, "key")


@pulumi.output_type
class GetProjectsProjectPropertiesTableLifecycleResult(dict):
    def __init__(__self__, *,
                 type: str,
                 value: str):
        """
        :param str type: Project type
        :param str value: The value of the life cycle.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Project type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the life cycle.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GetProjectsProjectSecurityPropertiesResult(dict):
    def __init__(__self__, *,
                 enable_download_privilege: bool,
                 label_security: bool,
                 object_creator_has_access_permission: bool,
                 object_creator_has_grant_permission: bool,
                 project_protection: 'outputs.GetProjectsProjectSecurityPropertiesProjectProtectionResult',
                 using_acl: bool,
                 using_policy: bool):
        """
        :param bool enable_download_privilege: Whether to enable download permission check.
        :param bool label_security: Label authorization.
        :param bool object_creator_has_access_permission: Project creator permissions.
        :param bool object_creator_has_grant_permission: Does the project creator have authorization rights.
        :param 'GetProjectsProjectSecurityPropertiesProjectProtectionArgs' project_protection: Project protection.
        :param bool using_acl: Whether to turn on ACL.
        :param bool using_policy: Whether to enable Policy.
        """
        pulumi.set(__self__, "enable_download_privilege", enable_download_privilege)
        pulumi.set(__self__, "label_security", label_security)
        pulumi.set(__self__, "object_creator_has_access_permission", object_creator_has_access_permission)
        pulumi.set(__self__, "object_creator_has_grant_permission", object_creator_has_grant_permission)
        pulumi.set(__self__, "project_protection", project_protection)
        pulumi.set(__self__, "using_acl", using_acl)
        pulumi.set(__self__, "using_policy", using_policy)

    @property
    @pulumi.getter(name="enableDownloadPrivilege")
    def enable_download_privilege(self) -> bool:
        """
        Whether to enable download permission check.
        """
        return pulumi.get(self, "enable_download_privilege")

    @property
    @pulumi.getter(name="labelSecurity")
    def label_security(self) -> bool:
        """
        Label authorization.
        """
        return pulumi.get(self, "label_security")

    @property
    @pulumi.getter(name="objectCreatorHasAccessPermission")
    def object_creator_has_access_permission(self) -> bool:
        """
        Project creator permissions.
        """
        return pulumi.get(self, "object_creator_has_access_permission")

    @property
    @pulumi.getter(name="objectCreatorHasGrantPermission")
    def object_creator_has_grant_permission(self) -> bool:
        """
        Does the project creator have authorization rights.
        """
        return pulumi.get(self, "object_creator_has_grant_permission")

    @property
    @pulumi.getter(name="projectProtection")
    def project_protection(self) -> 'outputs.GetProjectsProjectSecurityPropertiesProjectProtectionResult':
        """
        Project protection.
        """
        return pulumi.get(self, "project_protection")

    @property
    @pulumi.getter(name="usingAcl")
    def using_acl(self) -> bool:
        """
        Whether to turn on ACL.
        """
        return pulumi.get(self, "using_acl")

    @property
    @pulumi.getter(name="usingPolicy")
    def using_policy(self) -> bool:
        """
        Whether to enable Policy.
        """
        return pulumi.get(self, "using_policy")


@pulumi.output_type
class GetProjectsProjectSecurityPropertiesProjectProtectionResult(dict):
    def __init__(__self__, *,
                 exception_policy: str,
                 protected: bool):
        """
        :param str exception_policy: Exclusion policy.
        :param bool protected: Is it turned on.
        """
        pulumi.set(__self__, "exception_policy", exception_policy)
        pulumi.set(__self__, "protected", protected)

    @property
    @pulumi.getter(name="exceptionPolicy")
    def exception_policy(self) -> str:
        """
        Exclusion policy.
        """
        return pulumi.get(self, "exception_policy")

    @property
    @pulumi.getter
    def protected(self) -> bool:
        """
        Is it turned on.
        """
        return pulumi.get(self, "protected")


