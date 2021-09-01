# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetConfigsConfigResult',
    'GetRulesRuleResult',
]

@pulumi.output_type
class GetConfigsConfigResult(dict):
    def __init__(__self__, *,
                 code: str,
                 config_id: str,
                 default_value: str,
                 description: str,
                 id: str,
                 value: str):
        """
        :param str code: Abnormal Alarm General Configuration Module by Using the Encoding.Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
        :param str config_id: Configure the Number.
        :param str default_value: Default Value.
        :param str description: Abnormal Alarm General Description of the Configuration Item.
        :param str id: The ID of the Config.
        :param str value: The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different.
               * `access_failed_cnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
               * `access_permission_exprie_max_days`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
               * `log_datasize_avg_days`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "default_value", default_value)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        Abnormal Alarm General Configuration Module by Using the Encoding.Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> str:
        """
        Configure the Number.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> str:
        """
        Default Value.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Abnormal Alarm General Description of the Configuration Item.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Config.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different.
        * `access_failed_cnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
        * `access_permission_exprie_max_days`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
        * `log_datasize_avg_days`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GetRulesRuleResult(dict):
    def __init__(__self__, *,
                 category: int,
                 category_name: str,
                 content: str,
                 content_category: str,
                 create_time: str,
                 custom_type: int,
                 description: str,
                 display_name: str,
                 gmt_modified: str,
                 id: str,
                 login_name: str,
                 major_key: str,
                 name: str,
                 product_code: str,
                 product_id: str,
                 risk_level_id: str,
                 risk_level_name: str,
                 rule_id: str,
                 stat_express: str,
                 status: int,
                 target: str,
                 user_id: str,
                 warn_level: int):
        """
        :param int category: Sensitive Data Identification Rules for the Type of.
        :param str category_name: Sensitive Data Identification Rules Belongs Type Name.
        :param str content: Sensitive Data Identification Rules the Content.
        :param str content_category: The Content Classification.
        :param str create_time: Sensitive Data Identification Rules the Creation Time of the Number of Milliseconds.
        :param int custom_type: Sensitive Data Identification Rules of Type. 0: the Built-in 1: The User-Defined.
        :param str description: Sensitive Data Identification a Description of the Rule Information.
        :param str display_name: Sensitive Data Identification Rules, Founder of Account Display Name.
        :param str gmt_modified: Sensitive Data Identification Rules to the Modified Time of the Number of Milliseconds.
        :param str id: The ID of the Rule.
        :param str login_name: Sensitive Data Identification Rules, Founder Of Account Login.
        :param str major_key: The Primary Key.
        :param str name: The name of rule.
        :param str product_code: Product Code.
        :param str product_id: Product ID.
        :param str risk_level_id: Sensitive Data Identification Rules of Risk Level ID. Valid values:1:S1, Weak Risk Level. 2:S2, Medium Risk Level. 3:S3 High Risk Level. 4:S4, the Highest Risk Level.
        :param str risk_level_name: Sensitive Data Identification Rules the Risk Level of. S1: Weak Risk Level S2: Moderate Risk Level S3: High Risk Level S4: the Highest Risk Level.
        :param str rule_id: The first ID of the resource.
        :param str stat_express: Triggered the Alarm Conditions.
        :param int status: Sensitive Data Identification Rules Detection State of.
        :param str target: The Target.
        :param str user_id: The User ID.
        :param int warn_level: The Level of Risk.
        """
        pulumi.set(__self__, "category", category)
        pulumi.set(__self__, "category_name", category_name)
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "content_category", content_category)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "custom_type", custom_type)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "gmt_modified", gmt_modified)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "login_name", login_name)
        pulumi.set(__self__, "major_key", major_key)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "product_code", product_code)
        pulumi.set(__self__, "product_id", product_id)
        pulumi.set(__self__, "risk_level_id", risk_level_id)
        pulumi.set(__self__, "risk_level_name", risk_level_name)
        pulumi.set(__self__, "rule_id", rule_id)
        pulumi.set(__self__, "stat_express", stat_express)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "target", target)
        pulumi.set(__self__, "user_id", user_id)
        pulumi.set(__self__, "warn_level", warn_level)

    @property
    @pulumi.getter
    def category(self) -> int:
        """
        Sensitive Data Identification Rules for the Type of.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="categoryName")
    def category_name(self) -> str:
        """
        Sensitive Data Identification Rules Belongs Type Name.
        """
        return pulumi.get(self, "category_name")

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        Sensitive Data Identification Rules the Content.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentCategory")
    def content_category(self) -> str:
        """
        The Content Classification.
        """
        return pulumi.get(self, "content_category")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Sensitive Data Identification Rules the Creation Time of the Number of Milliseconds.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="customType")
    def custom_type(self) -> int:
        """
        Sensitive Data Identification Rules of Type. 0: the Built-in 1: The User-Defined.
        """
        return pulumi.get(self, "custom_type")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Sensitive Data Identification a Description of the Rule Information.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Sensitive Data Identification Rules, Founder of Account Display Name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="gmtModified")
    def gmt_modified(self) -> str:
        """
        Sensitive Data Identification Rules to the Modified Time of the Number of Milliseconds.
        """
        return pulumi.get(self, "gmt_modified")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Rule.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> str:
        """
        Sensitive Data Identification Rules, Founder Of Account Login.
        """
        return pulumi.get(self, "login_name")

    @property
    @pulumi.getter(name="majorKey")
    def major_key(self) -> str:
        """
        The Primary Key.
        """
        return pulumi.get(self, "major_key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of rule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="productCode")
    def product_code(self) -> str:
        """
        Product Code.
        """
        return pulumi.get(self, "product_code")

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> str:
        """
        Product ID.
        """
        return pulumi.get(self, "product_id")

    @property
    @pulumi.getter(name="riskLevelId")
    def risk_level_id(self) -> str:
        """
        Sensitive Data Identification Rules of Risk Level ID. Valid values:1:S1, Weak Risk Level. 2:S2, Medium Risk Level. 3:S3 High Risk Level. 4:S4, the Highest Risk Level.
        """
        return pulumi.get(self, "risk_level_id")

    @property
    @pulumi.getter(name="riskLevelName")
    def risk_level_name(self) -> str:
        """
        Sensitive Data Identification Rules the Risk Level of. S1: Weak Risk Level S2: Moderate Risk Level S3: High Risk Level S4: the Highest Risk Level.
        """
        return pulumi.get(self, "risk_level_name")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> str:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter(name="statExpress")
    def stat_express(self) -> str:
        """
        Triggered the Alarm Conditions.
        """
        return pulumi.get(self, "stat_express")

    @property
    @pulumi.getter
    def status(self) -> int:
        """
        Sensitive Data Identification Rules Detection State of.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        The Target.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The User ID.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="warnLevel")
    def warn_level(self) -> int:
        """
        The Level of Risk.
        """
        return pulumi.get(self, "warn_level")


