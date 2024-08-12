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

__all__ = ['EventRuleArgs', 'EventRule']

@pulumi.input_type
class EventRuleArgs:
    def __init__(__self__, *,
                 event_pattern: pulumi.Input['EventRuleEventPatternArgs'],
                 rule_name: pulumi.Input[str],
                 contact_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleContactParameterArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fc_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleFcParameterArgs']]]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 mns_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleMnsParameterArgs']]]] = None,
                 open_api_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleOpenApiParameterArgs']]]] = None,
                 silence_time: Optional[pulumi.Input[int]] = None,
                 sls_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleSlsParameterArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 webhook_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleWebhookParameterArgs']]]] = None):
        """
        The set of arguments for constructing a EventRule resource.
        :param pulumi.Input['EventRuleEventPatternArgs'] event_pattern: Event mode, used to describe the trigger conditions for this event. See `event_pattern` below.
        :param pulumi.Input[str] rule_name: The name of the event-triggered alert rule.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleContactParameterArgs']]] contact_parameters: The information about the alert contact groups that receive alert notifications. See `contact_parameters` below.
        :param pulumi.Input[str] description: The description of the event-triggered alert rule.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleFcParameterArgs']]] fc_parameters: The information about the recipients in Function Compute. See `fc_parameters` below.
        :param pulumi.Input[str] group_id: The ID of the application group to which the event-triggered alert rule belongs.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleMnsParameterArgs']]] mns_parameters: The information about the recipients in Message Service (MNS). See `mns_parameters` below.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleOpenApiParameterArgs']]] open_api_parameters: The parameters of API callback notification. See `open_api_parameters` below.
        :param pulumi.Input[int] silence_time: The silence time.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleSlsParameterArgs']]] sls_parameters: The information about the recipients in Simple Log Service. See `sls_parameters` below.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `ENABLED`, `DISABLED`.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleWebhookParameterArgs']]] webhook_parameters: The information about the callback URLs that are used to receive alert notifications. See `webhook_parameters` below.
        """
        pulumi.set(__self__, "event_pattern", event_pattern)
        pulumi.set(__self__, "rule_name", rule_name)
        if contact_parameters is not None:
            pulumi.set(__self__, "contact_parameters", contact_parameters)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if fc_parameters is not None:
            pulumi.set(__self__, "fc_parameters", fc_parameters)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if mns_parameters is not None:
            pulumi.set(__self__, "mns_parameters", mns_parameters)
        if open_api_parameters is not None:
            pulumi.set(__self__, "open_api_parameters", open_api_parameters)
        if silence_time is not None:
            pulumi.set(__self__, "silence_time", silence_time)
        if sls_parameters is not None:
            pulumi.set(__self__, "sls_parameters", sls_parameters)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if webhook_parameters is not None:
            pulumi.set(__self__, "webhook_parameters", webhook_parameters)

    @property
    @pulumi.getter(name="eventPattern")
    def event_pattern(self) -> pulumi.Input['EventRuleEventPatternArgs']:
        """
        Event mode, used to describe the trigger conditions for this event. See `event_pattern` below.
        """
        return pulumi.get(self, "event_pattern")

    @event_pattern.setter
    def event_pattern(self, value: pulumi.Input['EventRuleEventPatternArgs']):
        pulumi.set(self, "event_pattern", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Input[str]:
        """
        The name of the event-triggered alert rule.
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter(name="contactParameters")
    def contact_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleContactParameterArgs']]]]:
        """
        The information about the alert contact groups that receive alert notifications. See `contact_parameters` below.
        """
        return pulumi.get(self, "contact_parameters")

    @contact_parameters.setter
    def contact_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleContactParameterArgs']]]]):
        pulumi.set(self, "contact_parameters", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the event-triggered alert rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="fcParameters")
    def fc_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleFcParameterArgs']]]]:
        """
        The information about the recipients in Function Compute. See `fc_parameters` below.
        """
        return pulumi.get(self, "fc_parameters")

    @fc_parameters.setter
    def fc_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleFcParameterArgs']]]]):
        pulumi.set(self, "fc_parameters", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the application group to which the event-triggered alert rule belongs.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="mnsParameters")
    def mns_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleMnsParameterArgs']]]]:
        """
        The information about the recipients in Message Service (MNS). See `mns_parameters` below.
        """
        return pulumi.get(self, "mns_parameters")

    @mns_parameters.setter
    def mns_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleMnsParameterArgs']]]]):
        pulumi.set(self, "mns_parameters", value)

    @property
    @pulumi.getter(name="openApiParameters")
    def open_api_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleOpenApiParameterArgs']]]]:
        """
        The parameters of API callback notification. See `open_api_parameters` below.
        """
        return pulumi.get(self, "open_api_parameters")

    @open_api_parameters.setter
    def open_api_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleOpenApiParameterArgs']]]]):
        pulumi.set(self, "open_api_parameters", value)

    @property
    @pulumi.getter(name="silenceTime")
    def silence_time(self) -> Optional[pulumi.Input[int]]:
        """
        The silence time.
        """
        return pulumi.get(self, "silence_time")

    @silence_time.setter
    def silence_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "silence_time", value)

    @property
    @pulumi.getter(name="slsParameters")
    def sls_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleSlsParameterArgs']]]]:
        """
        The information about the recipients in Simple Log Service. See `sls_parameters` below.
        """
        return pulumi.get(self, "sls_parameters")

    @sls_parameters.setter
    def sls_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleSlsParameterArgs']]]]):
        pulumi.set(self, "sls_parameters", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource. Valid values: `ENABLED`, `DISABLED`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="webhookParameters")
    def webhook_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleWebhookParameterArgs']]]]:
        """
        The information about the callback URLs that are used to receive alert notifications. See `webhook_parameters` below.
        """
        return pulumi.get(self, "webhook_parameters")

    @webhook_parameters.setter
    def webhook_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleWebhookParameterArgs']]]]):
        pulumi.set(self, "webhook_parameters", value)


@pulumi.input_type
class _EventRuleState:
    def __init__(__self__, *,
                 contact_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleContactParameterArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_pattern: Optional[pulumi.Input['EventRuleEventPatternArgs']] = None,
                 fc_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleFcParameterArgs']]]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 mns_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleMnsParameterArgs']]]] = None,
                 open_api_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleOpenApiParameterArgs']]]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 silence_time: Optional[pulumi.Input[int]] = None,
                 sls_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleSlsParameterArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 webhook_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleWebhookParameterArgs']]]] = None):
        """
        Input properties used for looking up and filtering EventRule resources.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleContactParameterArgs']]] contact_parameters: The information about the alert contact groups that receive alert notifications. See `contact_parameters` below.
        :param pulumi.Input[str] description: The description of the event-triggered alert rule.
        :param pulumi.Input['EventRuleEventPatternArgs'] event_pattern: Event mode, used to describe the trigger conditions for this event. See `event_pattern` below.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleFcParameterArgs']]] fc_parameters: The information about the recipients in Function Compute. See `fc_parameters` below.
        :param pulumi.Input[str] group_id: The ID of the application group to which the event-triggered alert rule belongs.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleMnsParameterArgs']]] mns_parameters: The information about the recipients in Message Service (MNS). See `mns_parameters` below.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleOpenApiParameterArgs']]] open_api_parameters: The parameters of API callback notification. See `open_api_parameters` below.
        :param pulumi.Input[str] rule_name: The name of the event-triggered alert rule.
        :param pulumi.Input[int] silence_time: The silence time.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleSlsParameterArgs']]] sls_parameters: The information about the recipients in Simple Log Service. See `sls_parameters` below.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `ENABLED`, `DISABLED`.
        :param pulumi.Input[Sequence[pulumi.Input['EventRuleWebhookParameterArgs']]] webhook_parameters: The information about the callback URLs that are used to receive alert notifications. See `webhook_parameters` below.
        """
        if contact_parameters is not None:
            pulumi.set(__self__, "contact_parameters", contact_parameters)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if event_pattern is not None:
            pulumi.set(__self__, "event_pattern", event_pattern)
        if fc_parameters is not None:
            pulumi.set(__self__, "fc_parameters", fc_parameters)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if mns_parameters is not None:
            pulumi.set(__self__, "mns_parameters", mns_parameters)
        if open_api_parameters is not None:
            pulumi.set(__self__, "open_api_parameters", open_api_parameters)
        if rule_name is not None:
            pulumi.set(__self__, "rule_name", rule_name)
        if silence_time is not None:
            pulumi.set(__self__, "silence_time", silence_time)
        if sls_parameters is not None:
            pulumi.set(__self__, "sls_parameters", sls_parameters)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if webhook_parameters is not None:
            pulumi.set(__self__, "webhook_parameters", webhook_parameters)

    @property
    @pulumi.getter(name="contactParameters")
    def contact_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleContactParameterArgs']]]]:
        """
        The information about the alert contact groups that receive alert notifications. See `contact_parameters` below.
        """
        return pulumi.get(self, "contact_parameters")

    @contact_parameters.setter
    def contact_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleContactParameterArgs']]]]):
        pulumi.set(self, "contact_parameters", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the event-triggered alert rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="eventPattern")
    def event_pattern(self) -> Optional[pulumi.Input['EventRuleEventPatternArgs']]:
        """
        Event mode, used to describe the trigger conditions for this event. See `event_pattern` below.
        """
        return pulumi.get(self, "event_pattern")

    @event_pattern.setter
    def event_pattern(self, value: Optional[pulumi.Input['EventRuleEventPatternArgs']]):
        pulumi.set(self, "event_pattern", value)

    @property
    @pulumi.getter(name="fcParameters")
    def fc_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleFcParameterArgs']]]]:
        """
        The information about the recipients in Function Compute. See `fc_parameters` below.
        """
        return pulumi.get(self, "fc_parameters")

    @fc_parameters.setter
    def fc_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleFcParameterArgs']]]]):
        pulumi.set(self, "fc_parameters", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the application group to which the event-triggered alert rule belongs.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="mnsParameters")
    def mns_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleMnsParameterArgs']]]]:
        """
        The information about the recipients in Message Service (MNS). See `mns_parameters` below.
        """
        return pulumi.get(self, "mns_parameters")

    @mns_parameters.setter
    def mns_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleMnsParameterArgs']]]]):
        pulumi.set(self, "mns_parameters", value)

    @property
    @pulumi.getter(name="openApiParameters")
    def open_api_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleOpenApiParameterArgs']]]]:
        """
        The parameters of API callback notification. See `open_api_parameters` below.
        """
        return pulumi.get(self, "open_api_parameters")

    @open_api_parameters.setter
    def open_api_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleOpenApiParameterArgs']]]]):
        pulumi.set(self, "open_api_parameters", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the event-triggered alert rule.
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter(name="silenceTime")
    def silence_time(self) -> Optional[pulumi.Input[int]]:
        """
        The silence time.
        """
        return pulumi.get(self, "silence_time")

    @silence_time.setter
    def silence_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "silence_time", value)

    @property
    @pulumi.getter(name="slsParameters")
    def sls_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleSlsParameterArgs']]]]:
        """
        The information about the recipients in Simple Log Service. See `sls_parameters` below.
        """
        return pulumi.get(self, "sls_parameters")

    @sls_parameters.setter
    def sls_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleSlsParameterArgs']]]]):
        pulumi.set(self, "sls_parameters", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource. Valid values: `ENABLED`, `DISABLED`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="webhookParameters")
    def webhook_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleWebhookParameterArgs']]]]:
        """
        The information about the callback URLs that are used to receive alert notifications. See `webhook_parameters` below.
        """
        return pulumi.get(self, "webhook_parameters")

    @webhook_parameters.setter
    def webhook_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventRuleWebhookParameterArgs']]]]):
        pulumi.set(self, "webhook_parameters", value)


class EventRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contact_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleContactParameterArgs', 'EventRuleContactParameterArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_pattern: Optional[pulumi.Input[Union['EventRuleEventPatternArgs', 'EventRuleEventPatternArgsDict']]] = None,
                 fc_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleFcParameterArgs', 'EventRuleFcParameterArgsDict']]]]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 mns_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleMnsParameterArgs', 'EventRuleMnsParameterArgsDict']]]]] = None,
                 open_api_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleOpenApiParameterArgs', 'EventRuleOpenApiParameterArgsDict']]]]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 silence_time: Optional[pulumi.Input[int]] = None,
                 sls_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleSlsParameterArgs', 'EventRuleSlsParameterArgsDict']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 webhook_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleWebhookParameterArgs', 'EventRuleWebhookParameterArgsDict']]]]] = None,
                 __props__=None):
        """
        Provides a Cloud Monitor Service Event Rule resource.

        For information about Cloud Monitor Service Event Rule and how to use it, see [What is Event Rule](https://www.alibabacloud.com/help/en/cloudmonitor/latest/puteventrule).

        > **NOTE:** Available since v1.182.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.cms.MonitorGroup("default", monitor_group_name=name)
        example = alicloud.cms.EventRule("example",
            rule_name=name,
            group_id=default.id,
            silence_time=100,
            description=name,
            status="ENABLED",
            event_pattern={
                "product": "ecs",
                "sql_filter": "example_value",
                "name_lists": ["example_value"],
                "level_lists": ["CRITICAL"],
                "event_type_lists": ["StatusNotification"],
            })
        ```

        ## Import

        Cloud Monitor Service Event Rule can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cms/eventRule:EventRule example <rule_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleContactParameterArgs', 'EventRuleContactParameterArgsDict']]]] contact_parameters: The information about the alert contact groups that receive alert notifications. See `contact_parameters` below.
        :param pulumi.Input[str] description: The description of the event-triggered alert rule.
        :param pulumi.Input[Union['EventRuleEventPatternArgs', 'EventRuleEventPatternArgsDict']] event_pattern: Event mode, used to describe the trigger conditions for this event. See `event_pattern` below.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleFcParameterArgs', 'EventRuleFcParameterArgsDict']]]] fc_parameters: The information about the recipients in Function Compute. See `fc_parameters` below.
        :param pulumi.Input[str] group_id: The ID of the application group to which the event-triggered alert rule belongs.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleMnsParameterArgs', 'EventRuleMnsParameterArgsDict']]]] mns_parameters: The information about the recipients in Message Service (MNS). See `mns_parameters` below.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleOpenApiParameterArgs', 'EventRuleOpenApiParameterArgsDict']]]] open_api_parameters: The parameters of API callback notification. See `open_api_parameters` below.
        :param pulumi.Input[str] rule_name: The name of the event-triggered alert rule.
        :param pulumi.Input[int] silence_time: The silence time.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleSlsParameterArgs', 'EventRuleSlsParameterArgsDict']]]] sls_parameters: The information about the recipients in Simple Log Service. See `sls_parameters` below.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `ENABLED`, `DISABLED`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleWebhookParameterArgs', 'EventRuleWebhookParameterArgsDict']]]] webhook_parameters: The information about the callback URLs that are used to receive alert notifications. See `webhook_parameters` below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Monitor Service Event Rule resource.

        For information about Cloud Monitor Service Event Rule and how to use it, see [What is Event Rule](https://www.alibabacloud.com/help/en/cloudmonitor/latest/puteventrule).

        > **NOTE:** Available since v1.182.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.cms.MonitorGroup("default", monitor_group_name=name)
        example = alicloud.cms.EventRule("example",
            rule_name=name,
            group_id=default.id,
            silence_time=100,
            description=name,
            status="ENABLED",
            event_pattern={
                "product": "ecs",
                "sql_filter": "example_value",
                "name_lists": ["example_value"],
                "level_lists": ["CRITICAL"],
                "event_type_lists": ["StatusNotification"],
            })
        ```

        ## Import

        Cloud Monitor Service Event Rule can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cms/eventRule:EventRule example <rule_name>
        ```

        :param str resource_name: The name of the resource.
        :param EventRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contact_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleContactParameterArgs', 'EventRuleContactParameterArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_pattern: Optional[pulumi.Input[Union['EventRuleEventPatternArgs', 'EventRuleEventPatternArgsDict']]] = None,
                 fc_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleFcParameterArgs', 'EventRuleFcParameterArgsDict']]]]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 mns_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleMnsParameterArgs', 'EventRuleMnsParameterArgsDict']]]]] = None,
                 open_api_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleOpenApiParameterArgs', 'EventRuleOpenApiParameterArgsDict']]]]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 silence_time: Optional[pulumi.Input[int]] = None,
                 sls_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleSlsParameterArgs', 'EventRuleSlsParameterArgsDict']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 webhook_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleWebhookParameterArgs', 'EventRuleWebhookParameterArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventRuleArgs.__new__(EventRuleArgs)

            __props__.__dict__["contact_parameters"] = contact_parameters
            __props__.__dict__["description"] = description
            if event_pattern is None and not opts.urn:
                raise TypeError("Missing required property 'event_pattern'")
            __props__.__dict__["event_pattern"] = event_pattern
            __props__.__dict__["fc_parameters"] = fc_parameters
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["mns_parameters"] = mns_parameters
            __props__.__dict__["open_api_parameters"] = open_api_parameters
            if rule_name is None and not opts.urn:
                raise TypeError("Missing required property 'rule_name'")
            __props__.__dict__["rule_name"] = rule_name
            __props__.__dict__["silence_time"] = silence_time
            __props__.__dict__["sls_parameters"] = sls_parameters
            __props__.__dict__["status"] = status
            __props__.__dict__["webhook_parameters"] = webhook_parameters
        super(EventRule, __self__).__init__(
            'alicloud:cms/eventRule:EventRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            contact_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleContactParameterArgs', 'EventRuleContactParameterArgsDict']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            event_pattern: Optional[pulumi.Input[Union['EventRuleEventPatternArgs', 'EventRuleEventPatternArgsDict']]] = None,
            fc_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleFcParameterArgs', 'EventRuleFcParameterArgsDict']]]]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            mns_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleMnsParameterArgs', 'EventRuleMnsParameterArgsDict']]]]] = None,
            open_api_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleOpenApiParameterArgs', 'EventRuleOpenApiParameterArgsDict']]]]] = None,
            rule_name: Optional[pulumi.Input[str]] = None,
            silence_time: Optional[pulumi.Input[int]] = None,
            sls_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleSlsParameterArgs', 'EventRuleSlsParameterArgsDict']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            webhook_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventRuleWebhookParameterArgs', 'EventRuleWebhookParameterArgsDict']]]]] = None) -> 'EventRule':
        """
        Get an existing EventRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleContactParameterArgs', 'EventRuleContactParameterArgsDict']]]] contact_parameters: The information about the alert contact groups that receive alert notifications. See `contact_parameters` below.
        :param pulumi.Input[str] description: The description of the event-triggered alert rule.
        :param pulumi.Input[Union['EventRuleEventPatternArgs', 'EventRuleEventPatternArgsDict']] event_pattern: Event mode, used to describe the trigger conditions for this event. See `event_pattern` below.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleFcParameterArgs', 'EventRuleFcParameterArgsDict']]]] fc_parameters: The information about the recipients in Function Compute. See `fc_parameters` below.
        :param pulumi.Input[str] group_id: The ID of the application group to which the event-triggered alert rule belongs.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleMnsParameterArgs', 'EventRuleMnsParameterArgsDict']]]] mns_parameters: The information about the recipients in Message Service (MNS). See `mns_parameters` below.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleOpenApiParameterArgs', 'EventRuleOpenApiParameterArgsDict']]]] open_api_parameters: The parameters of API callback notification. See `open_api_parameters` below.
        :param pulumi.Input[str] rule_name: The name of the event-triggered alert rule.
        :param pulumi.Input[int] silence_time: The silence time.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleSlsParameterArgs', 'EventRuleSlsParameterArgsDict']]]] sls_parameters: The information about the recipients in Simple Log Service. See `sls_parameters` below.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `ENABLED`, `DISABLED`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventRuleWebhookParameterArgs', 'EventRuleWebhookParameterArgsDict']]]] webhook_parameters: The information about the callback URLs that are used to receive alert notifications. See `webhook_parameters` below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventRuleState.__new__(_EventRuleState)

        __props__.__dict__["contact_parameters"] = contact_parameters
        __props__.__dict__["description"] = description
        __props__.__dict__["event_pattern"] = event_pattern
        __props__.__dict__["fc_parameters"] = fc_parameters
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["mns_parameters"] = mns_parameters
        __props__.__dict__["open_api_parameters"] = open_api_parameters
        __props__.__dict__["rule_name"] = rule_name
        __props__.__dict__["silence_time"] = silence_time
        __props__.__dict__["sls_parameters"] = sls_parameters
        __props__.__dict__["status"] = status
        __props__.__dict__["webhook_parameters"] = webhook_parameters
        return EventRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contactParameters")
    def contact_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.EventRuleContactParameter']]]:
        """
        The information about the alert contact groups that receive alert notifications. See `contact_parameters` below.
        """
        return pulumi.get(self, "contact_parameters")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the event-triggered alert rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eventPattern")
    def event_pattern(self) -> pulumi.Output['outputs.EventRuleEventPattern']:
        """
        Event mode, used to describe the trigger conditions for this event. See `event_pattern` below.
        """
        return pulumi.get(self, "event_pattern")

    @property
    @pulumi.getter(name="fcParameters")
    def fc_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.EventRuleFcParameter']]]:
        """
        The information about the recipients in Function Compute. See `fc_parameters` below.
        """
        return pulumi.get(self, "fc_parameters")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the application group to which the event-triggered alert rule belongs.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="mnsParameters")
    def mns_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.EventRuleMnsParameter']]]:
        """
        The information about the recipients in Message Service (MNS). See `mns_parameters` below.
        """
        return pulumi.get(self, "mns_parameters")

    @property
    @pulumi.getter(name="openApiParameters")
    def open_api_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.EventRuleOpenApiParameter']]]:
        """
        The parameters of API callback notification. See `open_api_parameters` below.
        """
        return pulumi.get(self, "open_api_parameters")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Output[str]:
        """
        The name of the event-triggered alert rule.
        """
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter(name="silenceTime")
    def silence_time(self) -> pulumi.Output[Optional[int]]:
        """
        The silence time.
        """
        return pulumi.get(self, "silence_time")

    @property
    @pulumi.getter(name="slsParameters")
    def sls_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.EventRuleSlsParameter']]]:
        """
        The information about the recipients in Simple Log Service. See `sls_parameters` below.
        """
        return pulumi.get(self, "sls_parameters")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource. Valid values: `ENABLED`, `DISABLED`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="webhookParameters")
    def webhook_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.EventRuleWebhookParameter']]]:
        """
        The information about the callback URLs that are used to receive alert notifications. See `webhook_parameters` below.
        """
        return pulumi.get(self, "webhook_parameters")

