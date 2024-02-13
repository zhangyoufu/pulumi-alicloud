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

__all__ = ['ApplicationScalingRuleArgs', 'ApplicationScalingRule']

@pulumi.input_type
class ApplicationScalingRuleArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[str],
                 scaling_rule_name: pulumi.Input[str],
                 scaling_rule_type: pulumi.Input[str],
                 min_ready_instance_ratio: Optional[pulumi.Input[int]] = None,
                 min_ready_instances: Optional[pulumi.Input[int]] = None,
                 scaling_rule_enable: Optional[pulumi.Input[bool]] = None,
                 scaling_rule_metric: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricArgs']] = None,
                 scaling_rule_timer: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleTimerArgs']] = None):
        """
        The set of arguments for constructing a ApplicationScalingRule resource.
        :param pulumi.Input[str] app_id: Application ID.
        :param pulumi.Input[str] scaling_rule_name: The name of a custom elastic scaling policy. In the application, the policy name cannot be repeated. It must start with a lowercase letter, and can only contain lowercase letters, numbers, and dashes (-), and no more than 32 characters. After the scaling policy is successfully created, the policy name cannot be modified.
        :param pulumi.Input[str] scaling_rule_type: Flexible strategy type. Valid values: `mix`, `timing` and `metric`.
        :param pulumi.Input[int] min_ready_instance_ratio: The min ready instance ratio.
        :param pulumi.Input[int] min_ready_instances: The min ready instances.
        :param pulumi.Input[bool] scaling_rule_enable: True whether the auto scaling policy is enabled. The value description is as follows: true: enabled state. false: disabled status. Valid values: `false`, `true`.
        :param pulumi.Input['ApplicationScalingRuleScalingRuleMetricArgs'] scaling_rule_metric: Monitor the configuration of the indicator elasticity strategy. See `scaling_rule_metric` below.
        :param pulumi.Input['ApplicationScalingRuleScalingRuleTimerArgs'] scaling_rule_timer: Configuration of Timing Resilient Policies. See `scaling_rule_timer` below.
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "scaling_rule_name", scaling_rule_name)
        pulumi.set(__self__, "scaling_rule_type", scaling_rule_type)
        if min_ready_instance_ratio is not None:
            pulumi.set(__self__, "min_ready_instance_ratio", min_ready_instance_ratio)
        if min_ready_instances is not None:
            pulumi.set(__self__, "min_ready_instances", min_ready_instances)
        if scaling_rule_enable is not None:
            pulumi.set(__self__, "scaling_rule_enable", scaling_rule_enable)
        if scaling_rule_metric is not None:
            pulumi.set(__self__, "scaling_rule_metric", scaling_rule_metric)
        if scaling_rule_timer is not None:
            pulumi.set(__self__, "scaling_rule_timer", scaling_rule_timer)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        """
        Application ID.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="scalingRuleName")
    def scaling_rule_name(self) -> pulumi.Input[str]:
        """
        The name of a custom elastic scaling policy. In the application, the policy name cannot be repeated. It must start with a lowercase letter, and can only contain lowercase letters, numbers, and dashes (-), and no more than 32 characters. After the scaling policy is successfully created, the policy name cannot be modified.
        """
        return pulumi.get(self, "scaling_rule_name")

    @scaling_rule_name.setter
    def scaling_rule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "scaling_rule_name", value)

    @property
    @pulumi.getter(name="scalingRuleType")
    def scaling_rule_type(self) -> pulumi.Input[str]:
        """
        Flexible strategy type. Valid values: `mix`, `timing` and `metric`.
        """
        return pulumi.get(self, "scaling_rule_type")

    @scaling_rule_type.setter
    def scaling_rule_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "scaling_rule_type", value)

    @property
    @pulumi.getter(name="minReadyInstanceRatio")
    def min_ready_instance_ratio(self) -> Optional[pulumi.Input[int]]:
        """
        The min ready instance ratio.
        """
        return pulumi.get(self, "min_ready_instance_ratio")

    @min_ready_instance_ratio.setter
    def min_ready_instance_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_ready_instance_ratio", value)

    @property
    @pulumi.getter(name="minReadyInstances")
    def min_ready_instances(self) -> Optional[pulumi.Input[int]]:
        """
        The min ready instances.
        """
        return pulumi.get(self, "min_ready_instances")

    @min_ready_instances.setter
    def min_ready_instances(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_ready_instances", value)

    @property
    @pulumi.getter(name="scalingRuleEnable")
    def scaling_rule_enable(self) -> Optional[pulumi.Input[bool]]:
        """
        True whether the auto scaling policy is enabled. The value description is as follows: true: enabled state. false: disabled status. Valid values: `false`, `true`.
        """
        return pulumi.get(self, "scaling_rule_enable")

    @scaling_rule_enable.setter
    def scaling_rule_enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "scaling_rule_enable", value)

    @property
    @pulumi.getter(name="scalingRuleMetric")
    def scaling_rule_metric(self) -> Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricArgs']]:
        """
        Monitor the configuration of the indicator elasticity strategy. See `scaling_rule_metric` below.
        """
        return pulumi.get(self, "scaling_rule_metric")

    @scaling_rule_metric.setter
    def scaling_rule_metric(self, value: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricArgs']]):
        pulumi.set(self, "scaling_rule_metric", value)

    @property
    @pulumi.getter(name="scalingRuleTimer")
    def scaling_rule_timer(self) -> Optional[pulumi.Input['ApplicationScalingRuleScalingRuleTimerArgs']]:
        """
        Configuration of Timing Resilient Policies. See `scaling_rule_timer` below.
        """
        return pulumi.get(self, "scaling_rule_timer")

    @scaling_rule_timer.setter
    def scaling_rule_timer(self, value: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleTimerArgs']]):
        pulumi.set(self, "scaling_rule_timer", value)


@pulumi.input_type
class _ApplicationScalingRuleState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 min_ready_instance_ratio: Optional[pulumi.Input[int]] = None,
                 min_ready_instances: Optional[pulumi.Input[int]] = None,
                 scaling_rule_enable: Optional[pulumi.Input[bool]] = None,
                 scaling_rule_metric: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricArgs']] = None,
                 scaling_rule_name: Optional[pulumi.Input[str]] = None,
                 scaling_rule_timer: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleTimerArgs']] = None,
                 scaling_rule_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationScalingRule resources.
        :param pulumi.Input[str] app_id: Application ID.
        :param pulumi.Input[int] min_ready_instance_ratio: The min ready instance ratio.
        :param pulumi.Input[int] min_ready_instances: The min ready instances.
        :param pulumi.Input[bool] scaling_rule_enable: True whether the auto scaling policy is enabled. The value description is as follows: true: enabled state. false: disabled status. Valid values: `false`, `true`.
        :param pulumi.Input['ApplicationScalingRuleScalingRuleMetricArgs'] scaling_rule_metric: Monitor the configuration of the indicator elasticity strategy. See `scaling_rule_metric` below.
        :param pulumi.Input[str] scaling_rule_name: The name of a custom elastic scaling policy. In the application, the policy name cannot be repeated. It must start with a lowercase letter, and can only contain lowercase letters, numbers, and dashes (-), and no more than 32 characters. After the scaling policy is successfully created, the policy name cannot be modified.
        :param pulumi.Input['ApplicationScalingRuleScalingRuleTimerArgs'] scaling_rule_timer: Configuration of Timing Resilient Policies. See `scaling_rule_timer` below.
        :param pulumi.Input[str] scaling_rule_type: Flexible strategy type. Valid values: `mix`, `timing` and `metric`.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if min_ready_instance_ratio is not None:
            pulumi.set(__self__, "min_ready_instance_ratio", min_ready_instance_ratio)
        if min_ready_instances is not None:
            pulumi.set(__self__, "min_ready_instances", min_ready_instances)
        if scaling_rule_enable is not None:
            pulumi.set(__self__, "scaling_rule_enable", scaling_rule_enable)
        if scaling_rule_metric is not None:
            pulumi.set(__self__, "scaling_rule_metric", scaling_rule_metric)
        if scaling_rule_name is not None:
            pulumi.set(__self__, "scaling_rule_name", scaling_rule_name)
        if scaling_rule_timer is not None:
            pulumi.set(__self__, "scaling_rule_timer", scaling_rule_timer)
        if scaling_rule_type is not None:
            pulumi.set(__self__, "scaling_rule_type", scaling_rule_type)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        Application ID.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="minReadyInstanceRatio")
    def min_ready_instance_ratio(self) -> Optional[pulumi.Input[int]]:
        """
        The min ready instance ratio.
        """
        return pulumi.get(self, "min_ready_instance_ratio")

    @min_ready_instance_ratio.setter
    def min_ready_instance_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_ready_instance_ratio", value)

    @property
    @pulumi.getter(name="minReadyInstances")
    def min_ready_instances(self) -> Optional[pulumi.Input[int]]:
        """
        The min ready instances.
        """
        return pulumi.get(self, "min_ready_instances")

    @min_ready_instances.setter
    def min_ready_instances(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_ready_instances", value)

    @property
    @pulumi.getter(name="scalingRuleEnable")
    def scaling_rule_enable(self) -> Optional[pulumi.Input[bool]]:
        """
        True whether the auto scaling policy is enabled. The value description is as follows: true: enabled state. false: disabled status. Valid values: `false`, `true`.
        """
        return pulumi.get(self, "scaling_rule_enable")

    @scaling_rule_enable.setter
    def scaling_rule_enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "scaling_rule_enable", value)

    @property
    @pulumi.getter(name="scalingRuleMetric")
    def scaling_rule_metric(self) -> Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricArgs']]:
        """
        Monitor the configuration of the indicator elasticity strategy. See `scaling_rule_metric` below.
        """
        return pulumi.get(self, "scaling_rule_metric")

    @scaling_rule_metric.setter
    def scaling_rule_metric(self, value: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleMetricArgs']]):
        pulumi.set(self, "scaling_rule_metric", value)

    @property
    @pulumi.getter(name="scalingRuleName")
    def scaling_rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a custom elastic scaling policy. In the application, the policy name cannot be repeated. It must start with a lowercase letter, and can only contain lowercase letters, numbers, and dashes (-), and no more than 32 characters. After the scaling policy is successfully created, the policy name cannot be modified.
        """
        return pulumi.get(self, "scaling_rule_name")

    @scaling_rule_name.setter
    def scaling_rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_rule_name", value)

    @property
    @pulumi.getter(name="scalingRuleTimer")
    def scaling_rule_timer(self) -> Optional[pulumi.Input['ApplicationScalingRuleScalingRuleTimerArgs']]:
        """
        Configuration of Timing Resilient Policies. See `scaling_rule_timer` below.
        """
        return pulumi.get(self, "scaling_rule_timer")

    @scaling_rule_timer.setter
    def scaling_rule_timer(self, value: Optional[pulumi.Input['ApplicationScalingRuleScalingRuleTimerArgs']]):
        pulumi.set(self, "scaling_rule_timer", value)

    @property
    @pulumi.getter(name="scalingRuleType")
    def scaling_rule_type(self) -> Optional[pulumi.Input[str]]:
        """
        Flexible strategy type. Valid values: `mix`, `timing` and `metric`.
        """
        return pulumi.get(self, "scaling_rule_type")

    @scaling_rule_type.setter
    def scaling_rule_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_rule_type", value)


class ApplicationScalingRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 min_ready_instance_ratio: Optional[pulumi.Input[int]] = None,
                 min_ready_instances: Optional[pulumi.Input[int]] = None,
                 scaling_rule_enable: Optional[pulumi.Input[bool]] = None,
                 scaling_rule_metric: Optional[pulumi.Input[pulumi.InputType['ApplicationScalingRuleScalingRuleMetricArgs']]] = None,
                 scaling_rule_name: Optional[pulumi.Input[str]] = None,
                 scaling_rule_timer: Optional[pulumi.Input[pulumi.InputType['ApplicationScalingRuleScalingRuleTimerArgs']]] = None,
                 scaling_rule_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Serverless App Engine (SAE) Application Scaling Rule resource.

        For information about Serverless App Engine (SAE) Application Scaling Rule and how to use it, see [What is Application Scaling Rule](https://www.alibabacloud.com/help/en/sae/latest/create-application-scaling-rule).

        > **NOTE:** Available since v1.159.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_regions = alicloud.get_regions(current=True)
        default_random_integer = random.RandomInteger("defaultRandomInteger",
            max=99999,
            min=10000)
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name=name,
            cidr_block="10.4.0.0/24",
            vpc_id=default_network.id,
            zone_id=default_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_namespace = alicloud.sae.Namespace("defaultNamespace",
            namespace_id=default_random_integer.result.apply(lambda result: f"{default_regions.regions[0].id}:example{result}"),
            namespace_name=name,
            namespace_description=name,
            enable_micro_registration=False)
        default_application = alicloud.sae.Application("defaultApplication",
            app_description=name,
            app_name=name,
            namespace_id=default_namespace.id,
            image_url=f"registry-vpc.{default_regions.regions[0].id}.aliyuncs.com/sae-demo-image/consumer:1.0",
            package_type="Image",
            security_group_id=default_security_group.id,
            vpc_id=default_network.id,
            vswitch_id=default_switch.id,
            timezone="Asia/Beijing",
            replicas=5,
            cpu=500,
            memory=2048)
        default_application_scaling_rule = alicloud.sae.ApplicationScalingRule("defaultApplicationScalingRule",
            app_id=default_application.id,
            scaling_rule_name=name,
            scaling_rule_enable=True,
            scaling_rule_type="mix",
            min_ready_instances=3,
            min_ready_instance_ratio=-1,
            scaling_rule_timer=alicloud.sae.ApplicationScalingRuleScalingRuleTimerArgs(
                period="* * *",
                schedules=[
                    alicloud.sae.ApplicationScalingRuleScalingRuleTimerScheduleArgs(
                        at_time="08:00",
                        max_replicas=10,
                        min_replicas=3,
                    ),
                    alicloud.sae.ApplicationScalingRuleScalingRuleTimerScheduleArgs(
                        at_time="20:00",
                        max_replicas=50,
                        min_replicas=3,
                    ),
                ],
            ),
            scaling_rule_metric=alicloud.sae.ApplicationScalingRuleScalingRuleMetricArgs(
                max_replicas=50,
                min_replicas=3,
                metrics=[
                    alicloud.sae.ApplicationScalingRuleScalingRuleMetricMetricArgs(
                        metric_type="CPU",
                        metric_target_average_utilization=20,
                    ),
                    alicloud.sae.ApplicationScalingRuleScalingRuleMetricMetricArgs(
                        metric_type="MEMORY",
                        metric_target_average_utilization=30,
                    ),
                    alicloud.sae.ApplicationScalingRuleScalingRuleMetricMetricArgs(
                        metric_type="tcpActiveConn",
                        metric_target_average_utilization=20,
                    ),
                ],
                scale_up_rules=alicloud.sae.ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs(
                    step=10,
                    disabled=False,
                    stabilization_window_seconds=0,
                ),
                scale_down_rules=alicloud.sae.ApplicationScalingRuleScalingRuleMetricScaleDownRulesArgs(
                    step=10,
                    disabled=False,
                    stabilization_window_seconds=10,
                ),
            ))
        ```

        ## Import

        Serverless App Engine (SAE) Application Scaling Rule can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:sae/applicationScalingRule:ApplicationScalingRule example <app_id>:<scaling_rule_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Application ID.
        :param pulumi.Input[int] min_ready_instance_ratio: The min ready instance ratio.
        :param pulumi.Input[int] min_ready_instances: The min ready instances.
        :param pulumi.Input[bool] scaling_rule_enable: True whether the auto scaling policy is enabled. The value description is as follows: true: enabled state. false: disabled status. Valid values: `false`, `true`.
        :param pulumi.Input[pulumi.InputType['ApplicationScalingRuleScalingRuleMetricArgs']] scaling_rule_metric: Monitor the configuration of the indicator elasticity strategy. See `scaling_rule_metric` below.
        :param pulumi.Input[str] scaling_rule_name: The name of a custom elastic scaling policy. In the application, the policy name cannot be repeated. It must start with a lowercase letter, and can only contain lowercase letters, numbers, and dashes (-), and no more than 32 characters. After the scaling policy is successfully created, the policy name cannot be modified.
        :param pulumi.Input[pulumi.InputType['ApplicationScalingRuleScalingRuleTimerArgs']] scaling_rule_timer: Configuration of Timing Resilient Policies. See `scaling_rule_timer` below.
        :param pulumi.Input[str] scaling_rule_type: Flexible strategy type. Valid values: `mix`, `timing` and `metric`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationScalingRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Serverless App Engine (SAE) Application Scaling Rule resource.

        For information about Serverless App Engine (SAE) Application Scaling Rule and how to use it, see [What is Application Scaling Rule](https://www.alibabacloud.com/help/en/sae/latest/create-application-scaling-rule).

        > **NOTE:** Available since v1.159.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_regions = alicloud.get_regions(current=True)
        default_random_integer = random.RandomInteger("defaultRandomInteger",
            max=99999,
            min=10000)
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name=name,
            cidr_block="10.4.0.0/24",
            vpc_id=default_network.id,
            zone_id=default_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_namespace = alicloud.sae.Namespace("defaultNamespace",
            namespace_id=default_random_integer.result.apply(lambda result: f"{default_regions.regions[0].id}:example{result}"),
            namespace_name=name,
            namespace_description=name,
            enable_micro_registration=False)
        default_application = alicloud.sae.Application("defaultApplication",
            app_description=name,
            app_name=name,
            namespace_id=default_namespace.id,
            image_url=f"registry-vpc.{default_regions.regions[0].id}.aliyuncs.com/sae-demo-image/consumer:1.0",
            package_type="Image",
            security_group_id=default_security_group.id,
            vpc_id=default_network.id,
            vswitch_id=default_switch.id,
            timezone="Asia/Beijing",
            replicas=5,
            cpu=500,
            memory=2048)
        default_application_scaling_rule = alicloud.sae.ApplicationScalingRule("defaultApplicationScalingRule",
            app_id=default_application.id,
            scaling_rule_name=name,
            scaling_rule_enable=True,
            scaling_rule_type="mix",
            min_ready_instances=3,
            min_ready_instance_ratio=-1,
            scaling_rule_timer=alicloud.sae.ApplicationScalingRuleScalingRuleTimerArgs(
                period="* * *",
                schedules=[
                    alicloud.sae.ApplicationScalingRuleScalingRuleTimerScheduleArgs(
                        at_time="08:00",
                        max_replicas=10,
                        min_replicas=3,
                    ),
                    alicloud.sae.ApplicationScalingRuleScalingRuleTimerScheduleArgs(
                        at_time="20:00",
                        max_replicas=50,
                        min_replicas=3,
                    ),
                ],
            ),
            scaling_rule_metric=alicloud.sae.ApplicationScalingRuleScalingRuleMetricArgs(
                max_replicas=50,
                min_replicas=3,
                metrics=[
                    alicloud.sae.ApplicationScalingRuleScalingRuleMetricMetricArgs(
                        metric_type="CPU",
                        metric_target_average_utilization=20,
                    ),
                    alicloud.sae.ApplicationScalingRuleScalingRuleMetricMetricArgs(
                        metric_type="MEMORY",
                        metric_target_average_utilization=30,
                    ),
                    alicloud.sae.ApplicationScalingRuleScalingRuleMetricMetricArgs(
                        metric_type="tcpActiveConn",
                        metric_target_average_utilization=20,
                    ),
                ],
                scale_up_rules=alicloud.sae.ApplicationScalingRuleScalingRuleMetricScaleUpRulesArgs(
                    step=10,
                    disabled=False,
                    stabilization_window_seconds=0,
                ),
                scale_down_rules=alicloud.sae.ApplicationScalingRuleScalingRuleMetricScaleDownRulesArgs(
                    step=10,
                    disabled=False,
                    stabilization_window_seconds=10,
                ),
            ))
        ```

        ## Import

        Serverless App Engine (SAE) Application Scaling Rule can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:sae/applicationScalingRule:ApplicationScalingRule example <app_id>:<scaling_rule_name>
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationScalingRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationScalingRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 min_ready_instance_ratio: Optional[pulumi.Input[int]] = None,
                 min_ready_instances: Optional[pulumi.Input[int]] = None,
                 scaling_rule_enable: Optional[pulumi.Input[bool]] = None,
                 scaling_rule_metric: Optional[pulumi.Input[pulumi.InputType['ApplicationScalingRuleScalingRuleMetricArgs']]] = None,
                 scaling_rule_name: Optional[pulumi.Input[str]] = None,
                 scaling_rule_timer: Optional[pulumi.Input[pulumi.InputType['ApplicationScalingRuleScalingRuleTimerArgs']]] = None,
                 scaling_rule_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationScalingRuleArgs.__new__(ApplicationScalingRuleArgs)

            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            __props__.__dict__["min_ready_instance_ratio"] = min_ready_instance_ratio
            __props__.__dict__["min_ready_instances"] = min_ready_instances
            __props__.__dict__["scaling_rule_enable"] = scaling_rule_enable
            __props__.__dict__["scaling_rule_metric"] = scaling_rule_metric
            if scaling_rule_name is None and not opts.urn:
                raise TypeError("Missing required property 'scaling_rule_name'")
            __props__.__dict__["scaling_rule_name"] = scaling_rule_name
            __props__.__dict__["scaling_rule_timer"] = scaling_rule_timer
            if scaling_rule_type is None and not opts.urn:
                raise TypeError("Missing required property 'scaling_rule_type'")
            __props__.__dict__["scaling_rule_type"] = scaling_rule_type
        super(ApplicationScalingRule, __self__).__init__(
            'alicloud:sae/applicationScalingRule:ApplicationScalingRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            min_ready_instance_ratio: Optional[pulumi.Input[int]] = None,
            min_ready_instances: Optional[pulumi.Input[int]] = None,
            scaling_rule_enable: Optional[pulumi.Input[bool]] = None,
            scaling_rule_metric: Optional[pulumi.Input[pulumi.InputType['ApplicationScalingRuleScalingRuleMetricArgs']]] = None,
            scaling_rule_name: Optional[pulumi.Input[str]] = None,
            scaling_rule_timer: Optional[pulumi.Input[pulumi.InputType['ApplicationScalingRuleScalingRuleTimerArgs']]] = None,
            scaling_rule_type: Optional[pulumi.Input[str]] = None) -> 'ApplicationScalingRule':
        """
        Get an existing ApplicationScalingRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Application ID.
        :param pulumi.Input[int] min_ready_instance_ratio: The min ready instance ratio.
        :param pulumi.Input[int] min_ready_instances: The min ready instances.
        :param pulumi.Input[bool] scaling_rule_enable: True whether the auto scaling policy is enabled. The value description is as follows: true: enabled state. false: disabled status. Valid values: `false`, `true`.
        :param pulumi.Input[pulumi.InputType['ApplicationScalingRuleScalingRuleMetricArgs']] scaling_rule_metric: Monitor the configuration of the indicator elasticity strategy. See `scaling_rule_metric` below.
        :param pulumi.Input[str] scaling_rule_name: The name of a custom elastic scaling policy. In the application, the policy name cannot be repeated. It must start with a lowercase letter, and can only contain lowercase letters, numbers, and dashes (-), and no more than 32 characters. After the scaling policy is successfully created, the policy name cannot be modified.
        :param pulumi.Input[pulumi.InputType['ApplicationScalingRuleScalingRuleTimerArgs']] scaling_rule_timer: Configuration of Timing Resilient Policies. See `scaling_rule_timer` below.
        :param pulumi.Input[str] scaling_rule_type: Flexible strategy type. Valid values: `mix`, `timing` and `metric`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationScalingRuleState.__new__(_ApplicationScalingRuleState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["min_ready_instance_ratio"] = min_ready_instance_ratio
        __props__.__dict__["min_ready_instances"] = min_ready_instances
        __props__.__dict__["scaling_rule_enable"] = scaling_rule_enable
        __props__.__dict__["scaling_rule_metric"] = scaling_rule_metric
        __props__.__dict__["scaling_rule_name"] = scaling_rule_name
        __props__.__dict__["scaling_rule_timer"] = scaling_rule_timer
        __props__.__dict__["scaling_rule_type"] = scaling_rule_type
        return ApplicationScalingRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        Application ID.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="minReadyInstanceRatio")
    def min_ready_instance_ratio(self) -> pulumi.Output[Optional[int]]:
        """
        The min ready instance ratio.
        """
        return pulumi.get(self, "min_ready_instance_ratio")

    @property
    @pulumi.getter(name="minReadyInstances")
    def min_ready_instances(self) -> pulumi.Output[Optional[int]]:
        """
        The min ready instances.
        """
        return pulumi.get(self, "min_ready_instances")

    @property
    @pulumi.getter(name="scalingRuleEnable")
    def scaling_rule_enable(self) -> pulumi.Output[bool]:
        """
        True whether the auto scaling policy is enabled. The value description is as follows: true: enabled state. false: disabled status. Valid values: `false`, `true`.
        """
        return pulumi.get(self, "scaling_rule_enable")

    @property
    @pulumi.getter(name="scalingRuleMetric")
    def scaling_rule_metric(self) -> pulumi.Output[Optional['outputs.ApplicationScalingRuleScalingRuleMetric']]:
        """
        Monitor the configuration of the indicator elasticity strategy. See `scaling_rule_metric` below.
        """
        return pulumi.get(self, "scaling_rule_metric")

    @property
    @pulumi.getter(name="scalingRuleName")
    def scaling_rule_name(self) -> pulumi.Output[str]:
        """
        The name of a custom elastic scaling policy. In the application, the policy name cannot be repeated. It must start with a lowercase letter, and can only contain lowercase letters, numbers, and dashes (-), and no more than 32 characters. After the scaling policy is successfully created, the policy name cannot be modified.
        """
        return pulumi.get(self, "scaling_rule_name")

    @property
    @pulumi.getter(name="scalingRuleTimer")
    def scaling_rule_timer(self) -> pulumi.Output[Optional['outputs.ApplicationScalingRuleScalingRuleTimer']]:
        """
        Configuration of Timing Resilient Policies. See `scaling_rule_timer` below.
        """
        return pulumi.get(self, "scaling_rule_timer")

    @property
    @pulumi.getter(name="scalingRuleType")
    def scaling_rule_type(self) -> pulumi.Output[str]:
        """
        Flexible strategy type. Valid values: `mix`, `timing` and `metric`.
        """
        return pulumi.get(self, "scaling_rule_type")

