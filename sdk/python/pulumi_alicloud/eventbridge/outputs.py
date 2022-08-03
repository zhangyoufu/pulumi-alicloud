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
    'RuleTarget',
    'RuleTargetParamList',
    'GetEventBusesBusResult',
    'GetEventSourcesSourceResult',
    'GetRulesRuleResult',
    'GetRulesRuleTargetResult',
]

@pulumi.output_type
class RuleTarget(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "paramLists":
            suggest = "param_lists"
        elif key == "targetId":
            suggest = "target_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RuleTarget. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RuleTarget.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RuleTarget.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 endpoint: str,
                 param_lists: Sequence['outputs.RuleTargetParamList'],
                 target_id: str,
                 type: str):
        """
        :param str endpoint: The endpoint of target.
        :param Sequence['RuleTargetParamListArgs'] param_lists: A list of param.
        :param str target_id: The ID of target.
        :param str type: The type of target. Valid values: `acs.fc.function`, `acs.mns.topic`, `acs.mns.queue`,`http`,`acs.sms`,`acs.mail`,`acs.dingtalk`,`https`, `acs.eventbridge`,`acs.rabbitmq` and `acs.rocketmq`.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "param_lists", param_lists)
        pulumi.set(__self__, "target_id", target_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The endpoint of target.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="paramLists")
    def param_lists(self) -> Sequence['outputs.RuleTargetParamList']:
        """
        A list of param.
        """
        return pulumi.get(self, "param_lists")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> str:
        """
        The ID of target.
        """
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of target. Valid values: `acs.fc.function`, `acs.mns.topic`, `acs.mns.queue`,`http`,`acs.sms`,`acs.mail`,`acs.dingtalk`,`https`, `acs.eventbridge`,`acs.rabbitmq` and `acs.rocketmq`.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class RuleTargetParamList(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceKey":
            suggest = "resource_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RuleTargetParamList. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RuleTargetParamList.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RuleTargetParamList.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 form: str,
                 resource_key: str,
                 template: Optional[str] = None,
                 value: Optional[str] = None):
        """
        :param str form: The format of param.  Valid values: `ORIGINAL`, `TEMPLATE`, `JSONPATH`, `CONSTANT`.
        :param str resource_key: The resource key of param.  For more information, see [Event target parameters](https://help.aliyun.com/document_detail/185887.htm)
        :param str template: The template of param.
        :param str value: The value of param.
        """
        pulumi.set(__self__, "form", form)
        pulumi.set(__self__, "resource_key", resource_key)
        if template is not None:
            pulumi.set(__self__, "template", template)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def form(self) -> str:
        """
        The format of param.  Valid values: `ORIGINAL`, `TEMPLATE`, `JSONPATH`, `CONSTANT`.
        """
        return pulumi.get(self, "form")

    @property
    @pulumi.getter(name="resourceKey")
    def resource_key(self) -> str:
        """
        The resource key of param.  For more information, see [Event target parameters](https://help.aliyun.com/document_detail/185887.htm)
        """
        return pulumi.get(self, "resource_key")

    @property
    @pulumi.getter
    def template(self) -> Optional[str]:
        """
        The template of param.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The value of param.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GetEventBusesBusResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 description: str,
                 event_bus_name: str,
                 id: str):
        """
        :param str create_time: The time of this bus was created.
        :param str description: The description of event bus.
        :param str event_bus_name: The name of event bus.
        :param str id: The ID of the Event Bus. Its value is same as Queue Name.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "event_bus_name", event_bus_name)
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time of this bus was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of event bus.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> str:
        """
        The name of event bus.
        """
        return pulumi.get(self, "event_bus_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Event Bus. Its value is same as Queue Name.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class GetEventSourcesSourceResult(dict):
    def __init__(__self__, *,
                 description: str,
                 event_source_name: str,
                 external_source_config: Mapping[str, Any],
                 external_source_type: str,
                 id: str,
                 linked_external_source: bool,
                 type: str):
        """
        :param str description: The detail describe of event source.
        :param str event_source_name: The code name of event source.
        :param Mapping[str, Any] external_source_config: The config of external data source.
        :param str external_source_type: The type of external data source.
        :param str id: The ID of the Event Source.
        :param bool linked_external_source: Whether to connect to an external data source.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "event_source_name", event_source_name)
        pulumi.set(__self__, "external_source_config", external_source_config)
        pulumi.set(__self__, "external_source_type", external_source_type)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "linked_external_source", linked_external_source)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The detail describe of event source.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eventSourceName")
    def event_source_name(self) -> str:
        """
        The code name of event source.
        """
        return pulumi.get(self, "event_source_name")

    @property
    @pulumi.getter(name="externalSourceConfig")
    def external_source_config(self) -> Mapping[str, Any]:
        """
        The config of external data source.
        """
        return pulumi.get(self, "external_source_config")

    @property
    @pulumi.getter(name="externalSourceType")
    def external_source_type(self) -> str:
        """
        The type of external data source.
        """
        return pulumi.get(self, "external_source_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Event Source.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="linkedExternalSource")
    def linked_external_source(self) -> bool:
        """
        Whether to connect to an external data source.
        """
        return pulumi.get(self, "linked_external_source")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


@pulumi.output_type
class GetRulesRuleResult(dict):
    def __init__(__self__, *,
                 description: str,
                 event_bus_name: str,
                 filter_pattern: str,
                 id: str,
                 rule_name: str,
                 status: str,
                 targets: Sequence['outputs.GetRulesRuleTargetResult']):
        """
        :param str description: The description of rule.
        :param str event_bus_name: The name of event bus.
        :param str filter_pattern: The pattern to match interested events.
        :param str id: The ID of the Rule.
        :param str rule_name: The name of rule.
        :param str status: Rule status, either Enable or Disable.
        :param Sequence['GetRulesRuleTargetArgs'] targets: The target for rule.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "event_bus_name", event_bus_name)
        pulumi.set(__self__, "filter_pattern", filter_pattern)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "rule_name", rule_name)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "targets", targets)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> str:
        """
        The name of event bus.
        """
        return pulumi.get(self, "event_bus_name")

    @property
    @pulumi.getter(name="filterPattern")
    def filter_pattern(self) -> str:
        """
        The pattern to match interested events.
        """
        return pulumi.get(self, "filter_pattern")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Rule.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> str:
        """
        The name of rule.
        """
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Rule status, either Enable or Disable.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def targets(self) -> Sequence['outputs.GetRulesRuleTargetResult']:
        """
        The target for rule.
        """
        return pulumi.get(self, "targets")


@pulumi.output_type
class GetRulesRuleTargetResult(dict):
    def __init__(__self__, *,
                 endpoint: str,
                 target_id: str,
                 type: str):
        """
        :param str endpoint: The endpoint.
        :param str target_id: The id of target.
        :param str type: The type of target.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "target_id", target_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The endpoint.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> str:
        """
        The id of target.
        """
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of target.
        """
        return pulumi.get(self, "type")


