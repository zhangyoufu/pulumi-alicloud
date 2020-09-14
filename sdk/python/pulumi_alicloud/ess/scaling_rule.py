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

__all__ = ['ScalingRule']


class ScalingRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adjustment_type: Optional[pulumi.Input[str]] = None,
                 adjustment_value: Optional[pulumi.Input[float]] = None,
                 cooldown: Optional[pulumi.Input[float]] = None,
                 disable_scale_in: Optional[pulumi.Input[bool]] = None,
                 estimated_instance_warmup: Optional[pulumi.Input[float]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 scaling_rule_name: Optional[pulumi.Input[str]] = None,
                 scaling_rule_type: Optional[pulumi.Input[str]] = None,
                 step_adjustments: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ScalingRuleStepAdjustmentArgs']]]]] = None,
                 target_value: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a ScalingRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adjustment_type: Adjustment mode of a scaling rule. Optional values:
               - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
               - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
               - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
        :param pulumi.Input[float] adjustment_value: The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
               - QuantityChangeInCapacity：(0, 500] U (-500, 0]
               - PercentChangeInCapacity：[0, 10000] U [-100, 0]
               - TotalCapacity：[0, 1000]
        :param pulumi.Input[float] cooldown: The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
        :param pulumi.Input[bool] disable_scale_in: Indicates whether scale in by the target tracking policy is disabled. Default to false.
        :param pulumi.Input[float] estimated_instance_warmup: The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
        :param pulumi.Input[str] metric_name: A CloudMonitor metric name.
        :param pulumi.Input[str] scaling_group_id: ID of the scaling group of a scaling rule.
        :param pulumi.Input[str] scaling_rule_name: Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
        :param pulumi.Input[str] scaling_rule_type: The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ScalingRuleStepAdjustmentArgs']]]] step_adjustments: Steps for StepScalingRule. See Block stepAdjustment below for details.
        :param pulumi.Input[float] target_value: The target value for the metric.
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

            __props__['adjustment_type'] = adjustment_type
            __props__['adjustment_value'] = adjustment_value
            __props__['cooldown'] = cooldown
            __props__['disable_scale_in'] = disable_scale_in
            __props__['estimated_instance_warmup'] = estimated_instance_warmup
            __props__['metric_name'] = metric_name
            if scaling_group_id is None:
                raise TypeError("Missing required property 'scaling_group_id'")
            __props__['scaling_group_id'] = scaling_group_id
            __props__['scaling_rule_name'] = scaling_rule_name
            __props__['scaling_rule_type'] = scaling_rule_type
            __props__['step_adjustments'] = step_adjustments
            __props__['target_value'] = target_value
            __props__['ari'] = None
        super(ScalingRule, __self__).__init__(
            'alicloud:ess/scalingRule:ScalingRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adjustment_type: Optional[pulumi.Input[str]] = None,
            adjustment_value: Optional[pulumi.Input[float]] = None,
            ari: Optional[pulumi.Input[str]] = None,
            cooldown: Optional[pulumi.Input[float]] = None,
            disable_scale_in: Optional[pulumi.Input[bool]] = None,
            estimated_instance_warmup: Optional[pulumi.Input[float]] = None,
            metric_name: Optional[pulumi.Input[str]] = None,
            scaling_group_id: Optional[pulumi.Input[str]] = None,
            scaling_rule_name: Optional[pulumi.Input[str]] = None,
            scaling_rule_type: Optional[pulumi.Input[str]] = None,
            step_adjustments: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ScalingRuleStepAdjustmentArgs']]]]] = None,
            target_value: Optional[pulumi.Input[float]] = None) -> 'ScalingRule':
        """
        Get an existing ScalingRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adjustment_type: Adjustment mode of a scaling rule. Optional values:
               - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
               - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
               - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
        :param pulumi.Input[float] adjustment_value: The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
               - QuantityChangeInCapacity：(0, 500] U (-500, 0]
               - PercentChangeInCapacity：[0, 10000] U [-100, 0]
               - TotalCapacity：[0, 1000]
        :param pulumi.Input[float] cooldown: The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
        :param pulumi.Input[bool] disable_scale_in: Indicates whether scale in by the target tracking policy is disabled. Default to false.
        :param pulumi.Input[float] estimated_instance_warmup: The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
        :param pulumi.Input[str] metric_name: A CloudMonitor metric name.
        :param pulumi.Input[str] scaling_group_id: ID of the scaling group of a scaling rule.
        :param pulumi.Input[str] scaling_rule_name: Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
        :param pulumi.Input[str] scaling_rule_type: The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ScalingRuleStepAdjustmentArgs']]]] step_adjustments: Steps for StepScalingRule. See Block stepAdjustment below for details.
        :param pulumi.Input[float] target_value: The target value for the metric.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["adjustment_type"] = adjustment_type
        __props__["adjustment_value"] = adjustment_value
        __props__["ari"] = ari
        __props__["cooldown"] = cooldown
        __props__["disable_scale_in"] = disable_scale_in
        __props__["estimated_instance_warmup"] = estimated_instance_warmup
        __props__["metric_name"] = metric_name
        __props__["scaling_group_id"] = scaling_group_id
        __props__["scaling_rule_name"] = scaling_rule_name
        __props__["scaling_rule_type"] = scaling_rule_type
        __props__["step_adjustments"] = step_adjustments
        __props__["target_value"] = target_value
        return ScalingRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adjustmentType")
    def adjustment_type(self) -> pulumi.Output[Optional[str]]:
        """
        Adjustment mode of a scaling rule. Optional values:
        - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
        - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
        - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
        """
        return pulumi.get(self, "adjustment_type")

    @property
    @pulumi.getter(name="adjustmentValue")
    def adjustment_value(self) -> pulumi.Output[Optional[float]]:
        """
        The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
        - QuantityChangeInCapacity：(0, 500] U (-500, 0]
        - PercentChangeInCapacity：[0, 10000] U [-100, 0]
        - TotalCapacity：[0, 1000]
        """
        return pulumi.get(self, "adjustment_value")

    @property
    @pulumi.getter
    def ari(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ari")

    @property
    @pulumi.getter
    def cooldown(self) -> pulumi.Output[Optional[float]]:
        """
        The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
        """
        return pulumi.get(self, "cooldown")

    @property
    @pulumi.getter(name="disableScaleIn")
    def disable_scale_in(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether scale in by the target tracking policy is disabled. Default to false.
        """
        return pulumi.get(self, "disable_scale_in")

    @property
    @pulumi.getter(name="estimatedInstanceWarmup")
    def estimated_instance_warmup(self) -> pulumi.Output[float]:
        """
        The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
        """
        return pulumi.get(self, "estimated_instance_warmup")

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Output[Optional[str]]:
        """
        A CloudMonitor metric name.
        """
        return pulumi.get(self, "metric_name")

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Output[str]:
        """
        ID of the scaling group of a scaling rule.
        """
        return pulumi.get(self, "scaling_group_id")

    @property
    @pulumi.getter(name="scalingRuleName")
    def scaling_rule_name(self) -> pulumi.Output[str]:
        """
        Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
        """
        return pulumi.get(self, "scaling_rule_name")

    @property
    @pulumi.getter(name="scalingRuleType")
    def scaling_rule_type(self) -> pulumi.Output[Optional[str]]:
        """
        The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
        """
        return pulumi.get(self, "scaling_rule_type")

    @property
    @pulumi.getter(name="stepAdjustments")
    def step_adjustments(self) -> pulumi.Output[Optional[List['outputs.ScalingRuleStepAdjustment']]]:
        """
        Steps for StepScalingRule. See Block stepAdjustment below for details.
        """
        return pulumi.get(self, "step_adjustments")

    @property
    @pulumi.getter(name="targetValue")
    def target_value(self) -> pulumi.Output[Optional[float]]:
        """
        The target value for the metric.
        """
        return pulumi.get(self, "target_value")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

