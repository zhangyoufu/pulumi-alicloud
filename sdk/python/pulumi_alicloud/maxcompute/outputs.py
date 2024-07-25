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
        :param str ip_list: Set the IP address whitelist in the classic network. Only devices in the whitelist are allowed to access the project.> **NOTE:** If you only configure a classic network IP address whitelist, access to the classic network is restricted and all access to the VPC is prohibited.
        :param str vpc_ip_list: Set the IP address whitelist in the VPC network to allow only devices in the whitelist to access the project space.> **NOTE:** If you only configure a VPC network IP address whitelist, access to the VPC network is restricted and access to the classic network is prohibited.
        """
        if ip_list is not None:
            pulumi.set(__self__, "ip_list", ip_list)
        if vpc_ip_list is not None:
            pulumi.set(__self__, "vpc_ip_list", vpc_ip_list)

    @property
    @pulumi.getter(name="ipList")
    def ip_list(self) -> Optional[str]:
        """
        Set the IP address whitelist in the classic network. Only devices in the whitelist are allowed to access the project.> **NOTE:** If you only configure a classic network IP address whitelist, access to the classic network is restricted and all access to the VPC is prohibited.
        """
        return pulumi.get(self, "ip_list")

    @property
    @pulumi.getter(name="vpcIpList")
    def vpc_ip_list(self) -> Optional[str]:
        """
        Set the IP address whitelist in the VPC network to allow only devices in the whitelist to access the project space.> **NOTE:** If you only configure a VPC network IP address whitelist, access to the VPC network is restricted and access to the classic network is prohibited.
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
        :param bool allow_full_scan: Whether to allow full table scan. Default: false.
        :param bool enable_decimal2: Whether to turn on Decimal2.0.
        :param 'ProjectPropertiesEncryptionArgs' encryption: Storage encryption. For details, see [Storage Encryption](https://www.alibabacloud.com/help/en/maxcompute/security-and-compliance/storage-encryption)
               > **NOTE :**:  To enable storage encryption, you need to modify the parameters of the basic attributes of the MaxCompute project. This operation permission is authenticated by RAM, and you need to have the Super_Administrator role permission of the corresponding project.  To configure the permissions and IP whitelist parameters of the MaxCompute project, you must have the management permissions (Admin) of the corresponding project, including Super_Administrator, Admin, or custom management permissions. For more information, see the project management permissions list.  You can turn on storage encryption only for projects that have not turned on storage encryption. For projects that have turned on storage encryption, you cannot turn off storage encryption or change the encryption algorithm. See `encryption` below.
        :param int retention_days: Set the number of days to retain backup data. During this time, you can restore the current version to any backup version. The value range of days is [0,30], and the default value is 1. 0 means backup is turned off. The effective policy after adjusting the backup cycle is: Extend the backup cycle: The new backup cycle takes effect on the same day. Shorten the backup cycle: The system will automatically delete backup data that has exceeded the retention cycle.
        :param str sql_metering_max: Set the maximum threshold of single SQL consumption, that is, set the ODPS. SQL. metering.value.max attribute. For details, see [Consumption Monitoring Alarm](https://www.alibabacloud.com/help/en/maxcompute/product-overview/consumption-control). Unit: scan volume (GB)* complexity. .
        :param 'ProjectPropertiesTableLifecycleArgs' table_lifecycle: Set whether the lifecycle of the table in the project needs to be configured, that is, set the ODPS. table.lifecycle property,. See `table_lifecycle` below.
        :param str timezone: Project time zone, example value: Asia/Shanghai.
        :param str type_system: Data type version. Value:(1/2/hive) 1: The original MaxCompute type system. 2: New type system introduced by MaxCompute 2.0. hive: the type system of the Hive compatibility mode introduced by MaxCompute 2.0.
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
        Whether to allow full table scan. Default: false.
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
        Storage encryption. For details, see [Storage Encryption](https://www.alibabacloud.com/help/en/maxcompute/security-and-compliance/storage-encryption)
        > **NOTE :**:  To enable storage encryption, you need to modify the parameters of the basic attributes of the MaxCompute project. This operation permission is authenticated by RAM, and you need to have the Super_Administrator role permission of the corresponding project.  To configure the permissions and IP whitelist parameters of the MaxCompute project, you must have the management permissions (Admin) of the corresponding project, including Super_Administrator, Admin, or custom management permissions. For more information, see the project management permissions list.  You can turn on storage encryption only for projects that have not turned on storage encryption. For projects that have turned on storage encryption, you cannot turn off storage encryption or change the encryption algorithm. See `encryption` below.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[int]:
        """
        Set the number of days to retain backup data. During this time, you can restore the current version to any backup version. The value range of days is [0,30], and the default value is 1. 0 means backup is turned off. The effective policy after adjusting the backup cycle is: Extend the backup cycle: The new backup cycle takes effect on the same day. Shorten the backup cycle: The system will automatically delete backup data that has exceeded the retention cycle.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter(name="sqlMeteringMax")
    def sql_metering_max(self) -> Optional[str]:
        """
        Set the maximum threshold of single SQL consumption, that is, set the ODPS. SQL. metering.value.max attribute. For details, see [Consumption Monitoring Alarm](https://www.alibabacloud.com/help/en/maxcompute/product-overview/consumption-control). Unit: scan volume (GB)* complexity. .
        """
        return pulumi.get(self, "sql_metering_max")

    @property
    @pulumi.getter(name="tableLifecycle")
    def table_lifecycle(self) -> Optional['outputs.ProjectPropertiesTableLifecycle']:
        """
        Set whether the lifecycle of the table in the project needs to be configured, that is, set the ODPS. table.lifecycle property,. See `table_lifecycle` below.
        """
        return pulumi.get(self, "table_lifecycle")

    @property
    @pulumi.getter
    def timezone(self) -> Optional[str]:
        """
        Project time zone, example value: Asia/Shanghai.
        """
        return pulumi.get(self, "timezone")

    @property
    @pulumi.getter(name="typeSystem")
    def type_system(self) -> Optional[str]:
        """
        Data type version. Value:(1/2/hive) 1: The original MaxCompute type system. 2: New type system introduced by MaxCompute 2.0. hive: the type system of the Hive compatibility mode introduced by MaxCompute 2.0.
        """
        return pulumi.get(self, "type_system")


@pulumi.output_type
class ProjectPropertiesEncryption(dict):
    def __init__(__self__, *,
                 algorithm: Optional[str] = None,
                 enable: Optional[bool] = None,
                 key: Optional[str] = None):
        """
        :param str algorithm: The encryption algorithm supported by the key, including AES256, AESCTR, and RC4.
        :param bool enable: Only enable function is supported. Value: (true).
        :param str key: The encryption algorithm Key, the Key type used by the project, including the Default Key (MaxCompute Default Key) and the self-contained Key (BYOK). The MaxCompute Default Key is the Default Key created inside MaxCompute.
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
        The encryption algorithm supported by the key, including AES256, AESCTR, and RC4.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def enable(self) -> Optional[bool]:
        """
        Only enable function is supported. Value: (true).
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        """
        The encryption algorithm Key, the Key type used by the project, including the Default Key (MaxCompute Default Key) and the self-contained Key (BYOK). The MaxCompute Default Key is the Default Key created inside MaxCompute.
        """
        return pulumi.get(self, "key")


@pulumi.output_type
class ProjectPropertiesTableLifecycle(dict):
    def __init__(__self__, *,
                 type: Optional[str] = None,
                 value: Optional[str] = None):
        """
        :param str type: Project type
        :param str value: The value of the life cycle, in days. The value range is 1~37231, and the default value is 37231.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Project type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The value of the life cycle, in days. The value range is 1~37231, and the default value is 37231.
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
        :param bool enable_download_privilege: Set whether to enable the [Download permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/download-control), that is, set the ODPS. security.enabledownloadprivilege property.
        :param bool label_security: Set whether to use the [Label permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/label-based-access-control), that is, set the LabelSecurity attribute, which is not used by default.
        :param bool object_creator_has_access_permission: Sets whether to allow the creator of the object to have access to the object, I .e. sets the attribute. The default is the allowed state.
        :param bool object_creator_has_grant_permission: The ObjectCreatorHasGrantPermission attribute is set to allow the object creator to have the authorization permission on the object. The default is the allowed state.
        :param 'ProjectSecurityPropertiesProjectProtectionArgs' project_protection: Project protection. See `project_protection` below.
        :param bool using_acl: Set whether to use the [ACL permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/maxcompute-permissions), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
        :param bool using_policy: Set whether to use the Policy permission control function (https://www.alibabacloud.com/help/en/maxcompute/user-guide/policy-based-access-control-1), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
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
        Set whether to enable the [Download permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/download-control), that is, set the ODPS. security.enabledownloadprivilege property.
        """
        return pulumi.get(self, "enable_download_privilege")

    @property
    @pulumi.getter(name="labelSecurity")
    def label_security(self) -> Optional[bool]:
        """
        Set whether to use the [Label permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/label-based-access-control), that is, set the LabelSecurity attribute, which is not used by default.
        """
        return pulumi.get(self, "label_security")

    @property
    @pulumi.getter(name="objectCreatorHasAccessPermission")
    def object_creator_has_access_permission(self) -> Optional[bool]:
        """
        Sets whether to allow the creator of the object to have access to the object, I .e. sets the attribute. The default is the allowed state.
        """
        return pulumi.get(self, "object_creator_has_access_permission")

    @property
    @pulumi.getter(name="objectCreatorHasGrantPermission")
    def object_creator_has_grant_permission(self) -> Optional[bool]:
        """
        The ObjectCreatorHasGrantPermission attribute is set to allow the object creator to have the authorization permission on the object. The default is the allowed state.
        """
        return pulumi.get(self, "object_creator_has_grant_permission")

    @property
    @pulumi.getter(name="projectProtection")
    def project_protection(self) -> Optional['outputs.ProjectSecurityPropertiesProjectProtection']:
        """
        Project protection. See `project_protection` below.
        """
        return pulumi.get(self, "project_protection")

    @property
    @pulumi.getter(name="usingAcl")
    def using_acl(self) -> Optional[bool]:
        """
        Set whether to use the [ACL permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/maxcompute-permissions), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
        """
        return pulumi.get(self, "using_acl")

    @property
    @pulumi.getter(name="usingPolicy")
    def using_policy(self) -> Optional[bool]:
        """
        Set whether to use the Policy permission control function (https://www.alibabacloud.com/help/en/maxcompute/user-guide/policy-based-access-control-1), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
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
        :param str exception_policy: Set [Exceptions or Trusted Items](https://www.alibabacloud.com/help/en/maxcompute/security-and-compliance/project-data-protection).
        :param bool protected: Whether enabled, value:(true/false).
        """
        if exception_policy is not None:
            pulumi.set(__self__, "exception_policy", exception_policy)
        if protected is not None:
            pulumi.set(__self__, "protected", protected)

    @property
    @pulumi.getter(name="exceptionPolicy")
    def exception_policy(self) -> Optional[str]:
        """
        Set [Exceptions or Trusted Items](https://www.alibabacloud.com/help/en/maxcompute/security-and-compliance/project-data-protection).
        """
        return pulumi.get(self, "exception_policy")

    @property
    @pulumi.getter
    def protected(self) -> Optional[bool]:
        """
        Whether enabled, value:(true/false).
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


