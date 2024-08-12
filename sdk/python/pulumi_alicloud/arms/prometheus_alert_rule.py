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

__all__ = ['PrometheusAlertRuleArgs', 'PrometheusAlertRule']

@pulumi.input_type
class PrometheusAlertRuleArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 duration: pulumi.Input[str],
                 expression: pulumi.Input[str],
                 message: pulumi.Input[str],
                 prometheus_alert_rule_name: pulumi.Input[str],
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleAnnotationArgs']]]] = None,
                 dispatch_rule_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleLabelArgs']]]] = None,
                 notify_type: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrometheusAlertRule resource.
        :param pulumi.Input[str] cluster_id: The ID of the cluster.
        :param pulumi.Input[str] duration: The duration of the alert.
        :param pulumi.Input[str] expression: The alert rule expression that follows the PromQL syntax.
        :param pulumi.Input[str] message: The message of the alert notification.
        :param pulumi.Input[str] prometheus_alert_rule_name: The name of the resource.
        :param pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleAnnotationArgs']]] annotations: The annotations of the alert rule. See `annotations` below.
        :param pulumi.Input[str] dispatch_rule_id: The ID of the notification policy. This parameter is required when the `notify_type` parameter is set to `DISPATCH_RULE`.
        :param pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleLabelArgs']]] labels: The labels of the resource. See `labels` below.
        :param pulumi.Input[str] notify_type: The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
        :param pulumi.Input[str] type: The type of the alert rule.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "message", message)
        pulumi.set(__self__, "prometheus_alert_rule_name", prometheus_alert_rule_name)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if dispatch_rule_id is not None:
            pulumi.set(__self__, "dispatch_rule_id", dispatch_rule_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if notify_type is not None:
            pulumi.set(__self__, "notify_type", notify_type)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The ID of the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Input[str]:
        """
        The duration of the alert.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: pulumi.Input[str]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        The alert rule expression that follows the PromQL syntax.
        """
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def message(self) -> pulumi.Input[str]:
        """
        The message of the alert notification.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: pulumi.Input[str]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter(name="prometheusAlertRuleName")
    def prometheus_alert_rule_name(self) -> pulumi.Input[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "prometheus_alert_rule_name")

    @prometheus_alert_rule_name.setter
    def prometheus_alert_rule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "prometheus_alert_rule_name", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleAnnotationArgs']]]]:
        """
        The annotations of the alert rule. See `annotations` below.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleAnnotationArgs']]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="dispatchRuleId")
    def dispatch_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the notification policy. This parameter is required when the `notify_type` parameter is set to `DISPATCH_RULE`.
        """
        return pulumi.get(self, "dispatch_rule_id")

    @dispatch_rule_id.setter
    def dispatch_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dispatch_rule_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleLabelArgs']]]]:
        """
        The labels of the resource. See `labels` below.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleLabelArgs']]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="notifyType")
    def notify_type(self) -> Optional[pulumi.Input[str]]:
        """
        The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
        """
        return pulumi.get(self, "notify_type")

    @notify_type.setter
    def notify_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_type", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the alert rule.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _PrometheusAlertRuleState:
    def __init__(__self__, *,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleAnnotationArgs']]]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 dispatch_rule_id: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[str]] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleLabelArgs']]]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 notify_type: Optional[pulumi.Input[str]] = None,
                 prometheus_alert_rule_id: Optional[pulumi.Input[int]] = None,
                 prometheus_alert_rule_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrometheusAlertRule resources.
        :param pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleAnnotationArgs']]] annotations: The annotations of the alert rule. See `annotations` below.
        :param pulumi.Input[str] cluster_id: The ID of the cluster.
        :param pulumi.Input[str] dispatch_rule_id: The ID of the notification policy. This parameter is required when the `notify_type` parameter is set to `DISPATCH_RULE`.
        :param pulumi.Input[str] duration: The duration of the alert.
        :param pulumi.Input[str] expression: The alert rule expression that follows the PromQL syntax.
        :param pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleLabelArgs']]] labels: The labels of the resource. See `labels` below.
        :param pulumi.Input[str] message: The message of the alert notification.
        :param pulumi.Input[str] notify_type: The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
        :param pulumi.Input[int] prometheus_alert_rule_id: The first ID of the resource.
        :param pulumi.Input[str] prometheus_alert_rule_name: The name of the resource.
        :param pulumi.Input[int] status: The status of the resource. Valid values: `0`, `1`.
        :param pulumi.Input[str] type: The type of the alert rule.
        """
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if dispatch_rule_id is not None:
            pulumi.set(__self__, "dispatch_rule_id", dispatch_rule_id)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if expression is not None:
            pulumi.set(__self__, "expression", expression)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if notify_type is not None:
            pulumi.set(__self__, "notify_type", notify_type)
        if prometheus_alert_rule_id is not None:
            pulumi.set(__self__, "prometheus_alert_rule_id", prometheus_alert_rule_id)
        if prometheus_alert_rule_name is not None:
            pulumi.set(__self__, "prometheus_alert_rule_name", prometheus_alert_rule_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleAnnotationArgs']]]]:
        """
        The annotations of the alert rule. See `annotations` below.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleAnnotationArgs']]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="dispatchRuleId")
    def dispatch_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the notification policy. This parameter is required when the `notify_type` parameter is set to `DISPATCH_RULE`.
        """
        return pulumi.get(self, "dispatch_rule_id")

    @dispatch_rule_id.setter
    def dispatch_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dispatch_rule_id", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[str]]:
        """
        The duration of the alert.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def expression(self) -> Optional[pulumi.Input[str]]:
        """
        The alert rule expression that follows the PromQL syntax.
        """
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleLabelArgs']]]]:
        """
        The labels of the resource. See `labels` below.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusAlertRuleLabelArgs']]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        The message of the alert notification.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter(name="notifyType")
    def notify_type(self) -> Optional[pulumi.Input[str]]:
        """
        The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
        """
        return pulumi.get(self, "notify_type")

    @notify_type.setter
    def notify_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notify_type", value)

    @property
    @pulumi.getter(name="prometheusAlertRuleId")
    def prometheus_alert_rule_id(self) -> Optional[pulumi.Input[int]]:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "prometheus_alert_rule_id")

    @prometheus_alert_rule_id.setter
    def prometheus_alert_rule_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prometheus_alert_rule_id", value)

    @property
    @pulumi.getter(name="prometheusAlertRuleName")
    def prometheus_alert_rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "prometheus_alert_rule_name")

    @prometheus_alert_rule_name.setter
    def prometheus_alert_rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prometheus_alert_rule_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of the resource. Valid values: `0`, `1`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the alert rule.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class PrometheusAlertRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrometheusAlertRuleAnnotationArgs', 'PrometheusAlertRuleAnnotationArgsDict']]]]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 dispatch_rule_id: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[str]] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrometheusAlertRuleLabelArgs', 'PrometheusAlertRuleLabelArgsDict']]]]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 notify_type: Optional[pulumi.Input[str]] = None,
                 prometheus_alert_rule_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule resource.

        For information about Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule and how to use it, see [What is Prometheus Alert Rule](https://www.alibabacloud.com/help/en/doc-detail/212056.htm).

        > **NOTE:** Available since v1.136.0.

        ## Import

        Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:arms/prometheusAlertRule:PrometheusAlertRule example <cluster_id>:<prometheus_alert_rule_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PrometheusAlertRuleAnnotationArgs', 'PrometheusAlertRuleAnnotationArgsDict']]]] annotations: The annotations of the alert rule. See `annotations` below.
        :param pulumi.Input[str] cluster_id: The ID of the cluster.
        :param pulumi.Input[str] dispatch_rule_id: The ID of the notification policy. This parameter is required when the `notify_type` parameter is set to `DISPATCH_RULE`.
        :param pulumi.Input[str] duration: The duration of the alert.
        :param pulumi.Input[str] expression: The alert rule expression that follows the PromQL syntax.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PrometheusAlertRuleLabelArgs', 'PrometheusAlertRuleLabelArgsDict']]]] labels: The labels of the resource. See `labels` below.
        :param pulumi.Input[str] message: The message of the alert notification.
        :param pulumi.Input[str] notify_type: The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
        :param pulumi.Input[str] prometheus_alert_rule_name: The name of the resource.
        :param pulumi.Input[str] type: The type of the alert rule.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrometheusAlertRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule resource.

        For information about Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule and how to use it, see [What is Prometheus Alert Rule](https://www.alibabacloud.com/help/en/doc-detail/212056.htm).

        > **NOTE:** Available since v1.136.0.

        ## Import

        Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:arms/prometheusAlertRule:PrometheusAlertRule example <cluster_id>:<prometheus_alert_rule_id>
        ```

        :param str resource_name: The name of the resource.
        :param PrometheusAlertRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrometheusAlertRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrometheusAlertRuleAnnotationArgs', 'PrometheusAlertRuleAnnotationArgsDict']]]]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 dispatch_rule_id: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[str]] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrometheusAlertRuleLabelArgs', 'PrometheusAlertRuleLabelArgsDict']]]]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 notify_type: Optional[pulumi.Input[str]] = None,
                 prometheus_alert_rule_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrometheusAlertRuleArgs.__new__(PrometheusAlertRuleArgs)

            __props__.__dict__["annotations"] = annotations
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["dispatch_rule_id"] = dispatch_rule_id
            if duration is None and not opts.urn:
                raise TypeError("Missing required property 'duration'")
            __props__.__dict__["duration"] = duration
            if expression is None and not opts.urn:
                raise TypeError("Missing required property 'expression'")
            __props__.__dict__["expression"] = expression
            __props__.__dict__["labels"] = labels
            if message is None and not opts.urn:
                raise TypeError("Missing required property 'message'")
            __props__.__dict__["message"] = message
            __props__.__dict__["notify_type"] = notify_type
            if prometheus_alert_rule_name is None and not opts.urn:
                raise TypeError("Missing required property 'prometheus_alert_rule_name'")
            __props__.__dict__["prometheus_alert_rule_name"] = prometheus_alert_rule_name
            __props__.__dict__["type"] = type
            __props__.__dict__["prometheus_alert_rule_id"] = None
            __props__.__dict__["status"] = None
        super(PrometheusAlertRule, __self__).__init__(
            'alicloud:arms/prometheusAlertRule:PrometheusAlertRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            annotations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrometheusAlertRuleAnnotationArgs', 'PrometheusAlertRuleAnnotationArgsDict']]]]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            dispatch_rule_id: Optional[pulumi.Input[str]] = None,
            duration: Optional[pulumi.Input[str]] = None,
            expression: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrometheusAlertRuleLabelArgs', 'PrometheusAlertRuleLabelArgsDict']]]]] = None,
            message: Optional[pulumi.Input[str]] = None,
            notify_type: Optional[pulumi.Input[str]] = None,
            prometheus_alert_rule_id: Optional[pulumi.Input[int]] = None,
            prometheus_alert_rule_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'PrometheusAlertRule':
        """
        Get an existing PrometheusAlertRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PrometheusAlertRuleAnnotationArgs', 'PrometheusAlertRuleAnnotationArgsDict']]]] annotations: The annotations of the alert rule. See `annotations` below.
        :param pulumi.Input[str] cluster_id: The ID of the cluster.
        :param pulumi.Input[str] dispatch_rule_id: The ID of the notification policy. This parameter is required when the `notify_type` parameter is set to `DISPATCH_RULE`.
        :param pulumi.Input[str] duration: The duration of the alert.
        :param pulumi.Input[str] expression: The alert rule expression that follows the PromQL syntax.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PrometheusAlertRuleLabelArgs', 'PrometheusAlertRuleLabelArgsDict']]]] labels: The labels of the resource. See `labels` below.
        :param pulumi.Input[str] message: The message of the alert notification.
        :param pulumi.Input[str] notify_type: The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
        :param pulumi.Input[int] prometheus_alert_rule_id: The first ID of the resource.
        :param pulumi.Input[str] prometheus_alert_rule_name: The name of the resource.
        :param pulumi.Input[int] status: The status of the resource. Valid values: `0`, `1`.
        :param pulumi.Input[str] type: The type of the alert rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrometheusAlertRuleState.__new__(_PrometheusAlertRuleState)

        __props__.__dict__["annotations"] = annotations
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["dispatch_rule_id"] = dispatch_rule_id
        __props__.__dict__["duration"] = duration
        __props__.__dict__["expression"] = expression
        __props__.__dict__["labels"] = labels
        __props__.__dict__["message"] = message
        __props__.__dict__["notify_type"] = notify_type
        __props__.__dict__["prometheus_alert_rule_id"] = prometheus_alert_rule_id
        __props__.__dict__["prometheus_alert_rule_name"] = prometheus_alert_rule_name
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        return PrometheusAlertRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Sequence['outputs.PrometheusAlertRuleAnnotation']]]:
        """
        The annotations of the alert rule. See `annotations` below.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The ID of the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="dispatchRuleId")
    def dispatch_rule_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the notification policy. This parameter is required when the `notify_type` parameter is set to `DISPATCH_RULE`.
        """
        return pulumi.get(self, "dispatch_rule_id")

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Output[str]:
        """
        The duration of the alert.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Output[str]:
        """
        The alert rule expression that follows the PromQL syntax.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Sequence['outputs.PrometheusAlertRuleLabel']]]:
        """
        The labels of the resource. See `labels` below.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def message(self) -> pulumi.Output[str]:
        """
        The message of the alert notification.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter(name="notifyType")
    def notify_type(self) -> pulumi.Output[Optional[str]]:
        """
        The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
        """
        return pulumi.get(self, "notify_type")

    @property
    @pulumi.getter(name="prometheusAlertRuleId")
    def prometheus_alert_rule_id(self) -> pulumi.Output[int]:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "prometheus_alert_rule_id")

    @property
    @pulumi.getter(name="prometheusAlertRuleName")
    def prometheus_alert_rule_name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "prometheus_alert_rule_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[int]:
        """
        The status of the resource. Valid values: `0`, `1`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the alert rule.
        """
        return pulumi.get(self, "type")

