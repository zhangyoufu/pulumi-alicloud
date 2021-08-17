# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'EcdPolicyGroupAuthorizeAccessPolicyRule',
    'EcdPolicyGroupAuthorizeSecurityPolicyRule',
    'GetPolicyGroupsGroupResult',
    'GetPolicyGroupsGroupAuthorizeAccessPolicyRuleResult',
    'GetPolicyGroupsGroupAuthorizeSecurityPolicyRuleResult',
]

@pulumi.output_type
class EcdPolicyGroupAuthorizeAccessPolicyRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cidrIp":
            suggest = "cidr_ip"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EcdPolicyGroupAuthorizeAccessPolicyRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EcdPolicyGroupAuthorizeAccessPolicyRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EcdPolicyGroupAuthorizeAccessPolicyRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cidr_ip: Optional[str] = None,
                 description: Optional[str] = None):
        """
        :param str cidr_ip: The cidrip of authorize access rule.
        :param str description: The description of authorize access rule.
        """
        if cidr_ip is not None:
            pulumi.set(__self__, "cidr_ip", cidr_ip)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> Optional[str]:
        """
        The cidrip of authorize access rule.
        """
        return pulumi.get(self, "cidr_ip")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of authorize access rule.
        """
        return pulumi.get(self, "description")


@pulumi.output_type
class EcdPolicyGroupAuthorizeSecurityPolicyRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cidrIp":
            suggest = "cidr_ip"
        elif key == "ipProtocol":
            suggest = "ip_protocol"
        elif key == "portRange":
            suggest = "port_range"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EcdPolicyGroupAuthorizeSecurityPolicyRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EcdPolicyGroupAuthorizeSecurityPolicyRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EcdPolicyGroupAuthorizeSecurityPolicyRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cidr_ip: Optional[str] = None,
                 description: Optional[str] = None,
                 ip_protocol: Optional[str] = None,
                 policy: Optional[str] = None,
                 port_range: Optional[str] = None,
                 priority: Optional[str] = None,
                 type: Optional[str] = None):
        """
        :param str cidr_ip: The cidrip of authorize access rule.
        :param str description: The description of authorize access rule.
        :param str ip_protocol: The ip protocol of security rules.
        :param str policy: The policy of security rules.
        :param str port_range: The port range of security rules.
        :param str priority: The priority of security rules.
        :param str type: The type of security rules.
        """
        if cidr_ip is not None:
            pulumi.set(__self__, "cidr_ip", cidr_ip)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_protocol is not None:
            pulumi.set(__self__, "ip_protocol", ip_protocol)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if port_range is not None:
            pulumi.set(__self__, "port_range", port_range)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> Optional[str]:
        """
        The cidrip of authorize access rule.
        """
        return pulumi.get(self, "cidr_ip")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of authorize access rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> Optional[str]:
        """
        The ip protocol of security rules.
        """
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter
    def policy(self) -> Optional[str]:
        """
        The policy of security rules.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="portRange")
    def port_range(self) -> Optional[str]:
        """
        The port range of security rules.
        """
        return pulumi.get(self, "port_range")

    @property
    @pulumi.getter
    def priority(self) -> Optional[str]:
        """
        The priority of security rules.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of security rules.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetPolicyGroupsGroupResult(dict):
    def __init__(__self__, *,
                 authorize_access_policy_rules: Sequence['outputs.GetPolicyGroupsGroupAuthorizeAccessPolicyRuleResult'],
                 authorize_security_policy_rules: Sequence['outputs.GetPolicyGroupsGroupAuthorizeSecurityPolicyRuleResult'],
                 clipboard: str,
                 domain_list: str,
                 eds_count: int,
                 html_access: str,
                 html_file_transfer: str,
                 id: str,
                 local_drive: str,
                 policy_group_id: str,
                 policy_group_name: str,
                 policy_group_type: str,
                 status: str,
                 usb_redirect: str,
                 visual_quality: str,
                 watermark: str,
                 watermark_transparency: str,
                 watermark_type: str):
        """
        :param Sequence['GetPolicyGroupsGroupAuthorizeAccessPolicyRuleArgs'] authorize_access_policy_rules: The rule of authorize access rule.
        :param Sequence['GetPolicyGroupsGroupAuthorizeSecurityPolicyRuleArgs'] authorize_security_policy_rules: The policy rule.
        :param str clipboard: The clipboard policy.
        :param str domain_list: The list of domain.
        :param int eds_count: The count of eds.
        :param str html_access: The access of html5.
        :param str html_file_transfer: The html5 file transfer.
        :param str id: The ID of the Policy Group.
        :param str local_drive: Local drive redirect policy.
        :param str policy_group_id: The policy group id.
        :param str policy_group_name: The name of policy group.
        :param str policy_group_type: The type of policy group.
        :param str status: The status of policy.
        :param str usb_redirect: The usb redirect policy.
        :param str visual_quality: The quality of visual.sae_ecdsae_nameecd_po
        :param str watermark: The watermark policy.
        :param str watermark_transparency: The watermark transparency.
        :param str watermark_type: The type of watemark.
        """
        pulumi.set(__self__, "authorize_access_policy_rules", authorize_access_policy_rules)
        pulumi.set(__self__, "authorize_security_policy_rules", authorize_security_policy_rules)
        pulumi.set(__self__, "clipboard", clipboard)
        pulumi.set(__self__, "domain_list", domain_list)
        pulumi.set(__self__, "eds_count", eds_count)
        pulumi.set(__self__, "html_access", html_access)
        pulumi.set(__self__, "html_file_transfer", html_file_transfer)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "local_drive", local_drive)
        pulumi.set(__self__, "policy_group_id", policy_group_id)
        pulumi.set(__self__, "policy_group_name", policy_group_name)
        pulumi.set(__self__, "policy_group_type", policy_group_type)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "usb_redirect", usb_redirect)
        pulumi.set(__self__, "visual_quality", visual_quality)
        pulumi.set(__self__, "watermark", watermark)
        pulumi.set(__self__, "watermark_transparency", watermark_transparency)
        pulumi.set(__self__, "watermark_type", watermark_type)

    @property
    @pulumi.getter(name="authorizeAccessPolicyRules")
    def authorize_access_policy_rules(self) -> Sequence['outputs.GetPolicyGroupsGroupAuthorizeAccessPolicyRuleResult']:
        """
        The rule of authorize access rule.
        """
        return pulumi.get(self, "authorize_access_policy_rules")

    @property
    @pulumi.getter(name="authorizeSecurityPolicyRules")
    def authorize_security_policy_rules(self) -> Sequence['outputs.GetPolicyGroupsGroupAuthorizeSecurityPolicyRuleResult']:
        """
        The policy rule.
        """
        return pulumi.get(self, "authorize_security_policy_rules")

    @property
    @pulumi.getter
    def clipboard(self) -> str:
        """
        The clipboard policy.
        """
        return pulumi.get(self, "clipboard")

    @property
    @pulumi.getter(name="domainList")
    def domain_list(self) -> str:
        """
        The list of domain.
        """
        return pulumi.get(self, "domain_list")

    @property
    @pulumi.getter(name="edsCount")
    def eds_count(self) -> int:
        """
        The count of eds.
        """
        return pulumi.get(self, "eds_count")

    @property
    @pulumi.getter(name="htmlAccess")
    def html_access(self) -> str:
        """
        The access of html5.
        """
        return pulumi.get(self, "html_access")

    @property
    @pulumi.getter(name="htmlFileTransfer")
    def html_file_transfer(self) -> str:
        """
        The html5 file transfer.
        """
        return pulumi.get(self, "html_file_transfer")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Policy Group.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="localDrive")
    def local_drive(self) -> str:
        """
        Local drive redirect policy.
        """
        return pulumi.get(self, "local_drive")

    @property
    @pulumi.getter(name="policyGroupId")
    def policy_group_id(self) -> str:
        """
        The policy group id.
        """
        return pulumi.get(self, "policy_group_id")

    @property
    @pulumi.getter(name="policyGroupName")
    def policy_group_name(self) -> str:
        """
        The name of policy group.
        """
        return pulumi.get(self, "policy_group_name")

    @property
    @pulumi.getter(name="policyGroupType")
    def policy_group_type(self) -> str:
        """
        The type of policy group.
        """
        return pulumi.get(self, "policy_group_type")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of policy.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="usbRedirect")
    def usb_redirect(self) -> str:
        """
        The usb redirect policy.
        """
        return pulumi.get(self, "usb_redirect")

    @property
    @pulumi.getter(name="visualQuality")
    def visual_quality(self) -> str:
        """
        The quality of visual.sae_ecdsae_nameecd_po
        """
        return pulumi.get(self, "visual_quality")

    @property
    @pulumi.getter
    def watermark(self) -> str:
        """
        The watermark policy.
        """
        return pulumi.get(self, "watermark")

    @property
    @pulumi.getter(name="watermarkTransparency")
    def watermark_transparency(self) -> str:
        """
        The watermark transparency.
        """
        return pulumi.get(self, "watermark_transparency")

    @property
    @pulumi.getter(name="watermarkType")
    def watermark_type(self) -> str:
        """
        The type of watemark.
        """
        return pulumi.get(self, "watermark_type")


@pulumi.output_type
class GetPolicyGroupsGroupAuthorizeAccessPolicyRuleResult(dict):
    def __init__(__self__, *,
                 cidr_ip: str,
                 description: str):
        """
        :param str cidr_ip: The cidrip of security rules.
        :param str description: The description of security rules.
        """
        pulumi.set(__self__, "cidr_ip", cidr_ip)
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> str:
        """
        The cidrip of security rules.
        """
        return pulumi.get(self, "cidr_ip")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of security rules.
        """
        return pulumi.get(self, "description")


@pulumi.output_type
class GetPolicyGroupsGroupAuthorizeSecurityPolicyRuleResult(dict):
    def __init__(__self__, *,
                 cidr_ip: str,
                 description: str,
                 ip_protocol: str,
                 policy: str,
                 port_range: str,
                 priority: str,
                 type: str):
        """
        :param str cidr_ip: The cidrip of security rules.
        :param str description: The description of security rules.
        :param str ip_protocol: The ip protocol of security rules.
        :param str policy: The policy of security rules.
        :param str port_range: The port range of security rules.
        :param str priority: The priority of security rules.
        :param str type: The type of security rules.
        """
        pulumi.set(__self__, "cidr_ip", cidr_ip)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "ip_protocol", ip_protocol)
        pulumi.set(__self__, "policy", policy)
        pulumi.set(__self__, "port_range", port_range)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> str:
        """
        The cidrip of security rules.
        """
        return pulumi.get(self, "cidr_ip")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of security rules.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> str:
        """
        The ip protocol of security rules.
        """
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter
    def policy(self) -> str:
        """
        The policy of security rules.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="portRange")
    def port_range(self) -> str:
        """
        The port range of security rules.
        """
        return pulumi.get(self, "port_range")

    @property
    @pulumi.getter
    def priority(self) -> str:
        """
        The priority of security rules.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of security rules.
        """
        return pulumi.get(self, "type")


