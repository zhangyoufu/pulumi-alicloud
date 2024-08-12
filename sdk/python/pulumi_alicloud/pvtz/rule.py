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

__all__ = ['RuleArgs', 'Rule']

@pulumi.input_type
class RuleArgs:
    def __init__(__self__, *,
                 endpoint_id: pulumi.Input[str],
                 forward_ips: pulumi.Input[Sequence[pulumi.Input['RuleForwardIpArgs']]],
                 rule_name: pulumi.Input[str],
                 zone_name: pulumi.Input[str],
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Rule resource.
        :param pulumi.Input[str] endpoint_id: The ID of the Endpoint.
        :param pulumi.Input[Sequence[pulumi.Input['RuleForwardIpArgs']]] forward_ips: Forwarding target. See `forward_ips` below.
        :param pulumi.Input[str] rule_name: The name of the resource.
        :param pulumi.Input[str] zone_name: The name of the forwarding zone.
        :param pulumi.Input[str] type: The type of the rule. Valid values: `OUTBOUND`.
        """
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        pulumi.set(__self__, "forward_ips", forward_ips)
        pulumi.set(__self__, "rule_name", rule_name)
        pulumi.set(__self__, "zone_name", zone_name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Input[str]:
        """
        The ID of the Endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="forwardIps")
    def forward_ips(self) -> pulumi.Input[Sequence[pulumi.Input['RuleForwardIpArgs']]]:
        """
        Forwarding target. See `forward_ips` below.
        """
        return pulumi.get(self, "forward_ips")

    @forward_ips.setter
    def forward_ips(self, value: pulumi.Input[Sequence[pulumi.Input['RuleForwardIpArgs']]]):
        pulumi.set(self, "forward_ips", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Input[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> pulumi.Input[str]:
        """
        The name of the forwarding zone.
        """
        return pulumi.get(self, "zone_name")

    @zone_name.setter
    def zone_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the rule. Valid values: `OUTBOUND`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _RuleState:
    def __init__(__self__, *,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 forward_ips: Optional[pulumi.Input[Sequence[pulumi.Input['RuleForwardIpArgs']]]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Rule resources.
        :param pulumi.Input[str] endpoint_id: The ID of the Endpoint.
        :param pulumi.Input[Sequence[pulumi.Input['RuleForwardIpArgs']]] forward_ips: Forwarding target. See `forward_ips` below.
        :param pulumi.Input[str] rule_name: The name of the resource.
        :param pulumi.Input[str] type: The type of the rule. Valid values: `OUTBOUND`.
        :param pulumi.Input[str] zone_name: The name of the forwarding zone.
        """
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if forward_ips is not None:
            pulumi.set(__self__, "forward_ips", forward_ips)
        if rule_name is not None:
            pulumi.set(__self__, "rule_name", rule_name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if zone_name is not None:
            pulumi.set(__self__, "zone_name", zone_name)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="forwardIps")
    def forward_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleForwardIpArgs']]]]:
        """
        Forwarding target. See `forward_ips` below.
        """
        return pulumi.get(self, "forward_ips")

    @forward_ips.setter
    def forward_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleForwardIpArgs']]]]):
        pulumi.set(self, "forward_ips", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the rule. Valid values: `OUTBOUND`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the forwarding zone.
        """
        return pulumi.get(self, "zone_name")

    @zone_name.setter
    def zone_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_name", value)


class Rule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 forward_ips: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RuleForwardIpArgs', 'RuleForwardIpArgsDict']]]]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Private Zone Rule resource.

        For information about Private Zone Rule and how to use it, see [What is Rule](https://www.alibabacloud.com/help/en/privatezone/latest/add-forwarding-rule).

        > **NOTE:** Available since v1.143.0.

        ## Import

        Private Zone Rule can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:pvtz/rule:Rule example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_id: The ID of the Endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RuleForwardIpArgs', 'RuleForwardIpArgsDict']]]] forward_ips: Forwarding target. See `forward_ips` below.
        :param pulumi.Input[str] rule_name: The name of the resource.
        :param pulumi.Input[str] type: The type of the rule. Valid values: `OUTBOUND`.
        :param pulumi.Input[str] zone_name: The name of the forwarding zone.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Private Zone Rule resource.

        For information about Private Zone Rule and how to use it, see [What is Rule](https://www.alibabacloud.com/help/en/privatezone/latest/add-forwarding-rule).

        > **NOTE:** Available since v1.143.0.

        ## Import

        Private Zone Rule can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:pvtz/rule:Rule example <id>
        ```

        :param str resource_name: The name of the resource.
        :param RuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 forward_ips: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RuleForwardIpArgs', 'RuleForwardIpArgsDict']]]]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RuleArgs.__new__(RuleArgs)

            if endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_id'")
            __props__.__dict__["endpoint_id"] = endpoint_id
            if forward_ips is None and not opts.urn:
                raise TypeError("Missing required property 'forward_ips'")
            __props__.__dict__["forward_ips"] = forward_ips
            if rule_name is None and not opts.urn:
                raise TypeError("Missing required property 'rule_name'")
            __props__.__dict__["rule_name"] = rule_name
            __props__.__dict__["type"] = type
            if zone_name is None and not opts.urn:
                raise TypeError("Missing required property 'zone_name'")
            __props__.__dict__["zone_name"] = zone_name
        super(Rule, __self__).__init__(
            'alicloud:pvtz/rule:Rule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint_id: Optional[pulumi.Input[str]] = None,
            forward_ips: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RuleForwardIpArgs', 'RuleForwardIpArgsDict']]]]] = None,
            rule_name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            zone_name: Optional[pulumi.Input[str]] = None) -> 'Rule':
        """
        Get an existing Rule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_id: The ID of the Endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RuleForwardIpArgs', 'RuleForwardIpArgsDict']]]] forward_ips: Forwarding target. See `forward_ips` below.
        :param pulumi.Input[str] rule_name: The name of the resource.
        :param pulumi.Input[str] type: The type of the rule. Valid values: `OUTBOUND`.
        :param pulumi.Input[str] zone_name: The name of the forwarding zone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RuleState.__new__(_RuleState)

        __props__.__dict__["endpoint_id"] = endpoint_id
        __props__.__dict__["forward_ips"] = forward_ips
        __props__.__dict__["rule_name"] = rule_name
        __props__.__dict__["type"] = type
        __props__.__dict__["zone_name"] = zone_name
        return Rule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Output[str]:
        """
        The ID of the Endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="forwardIps")
    def forward_ips(self) -> pulumi.Output[Sequence['outputs.RuleForwardIp']]:
        """
        Forwarding target. See `forward_ips` below.
        """
        return pulumi.get(self, "forward_ips")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the rule. Valid values: `OUTBOUND`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> pulumi.Output[str]:
        """
        The name of the forwarding zone.
        """
        return pulumi.get(self, "zone_name")

