# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AclAclEntryArgs',
    'EndpointGroupEndpointConfigurationArgs',
    'EndpointGroupPortOverridesArgs',
    'ForwardingRuleRuleActionArgs',
    'ForwardingRuleRuleActionForwardGroupConfigArgs',
    'ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs',
    'ForwardingRuleRuleConditionArgs',
    'ForwardingRuleRuleConditionHostConfigArgs',
    'ForwardingRuleRuleConditionPathConfigArgs',
    'ListenerCertificateArgs',
    'ListenerPortRangeArgs',
]

@pulumi.input_type
class AclAclEntryArgs:
    def __init__(__self__, *,
                 entry: Optional[pulumi.Input[str]] = None,
                 entry_description: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] entry: The IP entry that you want to add to the ACL.
        :param pulumi.Input[str] entry_description: The description of the IP entry. The description must be `1` to `256` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_).
        """
        if entry is not None:
            pulumi.set(__self__, "entry", entry)
        if entry_description is not None:
            pulumi.set(__self__, "entry_description", entry_description)

    @property
    @pulumi.getter
    def entry(self) -> Optional[pulumi.Input[str]]:
        """
        The IP entry that you want to add to the ACL.
        """
        return pulumi.get(self, "entry")

    @entry.setter
    def entry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entry", value)

    @property
    @pulumi.getter(name="entryDescription")
    def entry_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the IP entry. The description must be `1` to `256` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_).
        """
        return pulumi.get(self, "entry_description")

    @entry_description.setter
    def entry_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entry_description", value)


@pulumi.input_type
class EndpointGroupEndpointConfigurationArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input[str],
                 type: pulumi.Input[str],
                 weight: pulumi.Input[int],
                 enable_clientip_preservation: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] endpoint: The IP address or domain name of Endpoint N in the endpoint group.
        :param pulumi.Input[str] type: The type of Endpoint N in the endpoint group. Valid values: `Domain`: a custom domain name, `Ip`: a custom IP address, `PublicIp`: an Alibaba Cloud public IP address, `ECS`: an Alibaba Cloud Elastic Compute Service (ECS) instance, `SLB`: an Alibaba Cloud Server Load Balancer (SLB) instance.
        :param pulumi.Input[int] weight: The weight of Endpoint N in the endpoint group. Valid value is 0 to 255.
        :param pulumi.Input[bool] enable_clientip_preservation: Indicates whether client IP addresses are reserved. Valid values: `true`: Client IP addresses are reserved, `false`: Client IP addresses are not reserved. Default value is `false`.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "weight", weight)
        if enable_clientip_preservation is not None:
            pulumi.set(__self__, "enable_clientip_preservation", enable_clientip_preservation)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        """
        The IP address or domain name of Endpoint N in the endpoint group.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of Endpoint N in the endpoint group. Valid values: `Domain`: a custom domain name, `Ip`: a custom IP address, `PublicIp`: an Alibaba Cloud public IP address, `ECS`: an Alibaba Cloud Elastic Compute Service (ECS) instance, `SLB`: an Alibaba Cloud Server Load Balancer (SLB) instance.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Input[int]:
        """
        The weight of Endpoint N in the endpoint group. Valid value is 0 to 255.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: pulumi.Input[int]):
        pulumi.set(self, "weight", value)

    @property
    @pulumi.getter(name="enableClientipPreservation")
    def enable_clientip_preservation(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether client IP addresses are reserved. Valid values: `true`: Client IP addresses are reserved, `false`: Client IP addresses are not reserved. Default value is `false`.
        """
        return pulumi.get(self, "enable_clientip_preservation")

    @enable_clientip_preservation.setter
    def enable_clientip_preservation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_clientip_preservation", value)


@pulumi.input_type
class EndpointGroupPortOverridesArgs:
    def __init__(__self__, *,
                 endpoint_port: Optional[pulumi.Input[int]] = None,
                 listener_port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] endpoint_port: Forwarding port.
        :param pulumi.Input[int] listener_port: Listener port.
        """
        if endpoint_port is not None:
            pulumi.set(__self__, "endpoint_port", endpoint_port)
        if listener_port is not None:
            pulumi.set(__self__, "listener_port", listener_port)

    @property
    @pulumi.getter(name="endpointPort")
    def endpoint_port(self) -> Optional[pulumi.Input[int]]:
        """
        Forwarding port.
        """
        return pulumi.get(self, "endpoint_port")

    @endpoint_port.setter
    def endpoint_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "endpoint_port", value)

    @property
    @pulumi.getter(name="listenerPort")
    def listener_port(self) -> Optional[pulumi.Input[int]]:
        """
        Listener port.
        """
        return pulumi.get(self, "listener_port")

    @listener_port.setter
    def listener_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listener_port", value)


@pulumi.input_type
class ForwardingRuleRuleActionArgs:
    def __init__(__self__, *,
                 forward_group_config: pulumi.Input['ForwardingRuleRuleActionForwardGroupConfigArgs'],
                 order: pulumi.Input[int],
                 rule_action_type: pulumi.Input[str]):
        """
        :param pulumi.Input['ForwardingRuleRuleActionForwardGroupConfigArgs'] forward_group_config: Forwarding configuration.
        :param pulumi.Input[int] order: Forwarding priority.
        :param pulumi.Input[str] rule_action_type: Forward action type. Default: forwardgroup.
        """
        pulumi.set(__self__, "forward_group_config", forward_group_config)
        pulumi.set(__self__, "order", order)
        pulumi.set(__self__, "rule_action_type", rule_action_type)

    @property
    @pulumi.getter(name="forwardGroupConfig")
    def forward_group_config(self) -> pulumi.Input['ForwardingRuleRuleActionForwardGroupConfigArgs']:
        """
        Forwarding configuration.
        """
        return pulumi.get(self, "forward_group_config")

    @forward_group_config.setter
    def forward_group_config(self, value: pulumi.Input['ForwardingRuleRuleActionForwardGroupConfigArgs']):
        pulumi.set(self, "forward_group_config", value)

    @property
    @pulumi.getter
    def order(self) -> pulumi.Input[int]:
        """
        Forwarding priority.
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: pulumi.Input[int]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter(name="ruleActionType")
    def rule_action_type(self) -> pulumi.Input[str]:
        """
        Forward action type. Default: forwardgroup.
        """
        return pulumi.get(self, "rule_action_type")

    @rule_action_type.setter
    def rule_action_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_action_type", value)


@pulumi.input_type
class ForwardingRuleRuleActionForwardGroupConfigArgs:
    def __init__(__self__, *,
                 server_group_tuples: pulumi.Input[Sequence[pulumi.Input['ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs']]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs']]] server_group_tuples: Terminal node group configuration.
        """
        pulumi.set(__self__, "server_group_tuples", server_group_tuples)

    @property
    @pulumi.getter(name="serverGroupTuples")
    def server_group_tuples(self) -> pulumi.Input[Sequence[pulumi.Input['ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs']]]:
        """
        Terminal node group configuration.
        """
        return pulumi.get(self, "server_group_tuples")

    @server_group_tuples.setter
    def server_group_tuples(self, value: pulumi.Input[Sequence[pulumi.Input['ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs']]]):
        pulumi.set(self, "server_group_tuples", value)


@pulumi.input_type
class ForwardingRuleRuleActionForwardGroupConfigServerGroupTupleArgs:
    def __init__(__self__, *,
                 endpoint_group_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] endpoint_group_id: Terminal node group ID.
        """
        pulumi.set(__self__, "endpoint_group_id", endpoint_group_id)

    @property
    @pulumi.getter(name="endpointGroupId")
    def endpoint_group_id(self) -> pulumi.Input[str]:
        """
        Terminal node group ID.
        """
        return pulumi.get(self, "endpoint_group_id")

    @endpoint_group_id.setter
    def endpoint_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_group_id", value)


@pulumi.input_type
class ForwardingRuleRuleConditionArgs:
    def __init__(__self__, *,
                 rule_condition_type: pulumi.Input[str],
                 host_configs: Optional[pulumi.Input[Sequence[pulumi.Input['ForwardingRuleRuleConditionHostConfigArgs']]]] = None,
                 path_config: Optional[pulumi.Input['ForwardingRuleRuleConditionPathConfigArgs']] = None):
        """
        :param pulumi.Input[str] rule_condition_type: Forwarding condition type. Valid value: `Host`, `Path`.
        :param pulumi.Input[Sequence[pulumi.Input['ForwardingRuleRuleConditionHostConfigArgs']]] host_configs: Domain name configuration information.
        :param pulumi.Input['ForwardingRuleRuleConditionPathConfigArgs'] path_config: Path configuration information.
        """
        pulumi.set(__self__, "rule_condition_type", rule_condition_type)
        if host_configs is not None:
            pulumi.set(__self__, "host_configs", host_configs)
        if path_config is not None:
            pulumi.set(__self__, "path_config", path_config)

    @property
    @pulumi.getter(name="ruleConditionType")
    def rule_condition_type(self) -> pulumi.Input[str]:
        """
        Forwarding condition type. Valid value: `Host`, `Path`.
        """
        return pulumi.get(self, "rule_condition_type")

    @rule_condition_type.setter
    def rule_condition_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_condition_type", value)

    @property
    @pulumi.getter(name="hostConfigs")
    def host_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ForwardingRuleRuleConditionHostConfigArgs']]]]:
        """
        Domain name configuration information.
        """
        return pulumi.get(self, "host_configs")

    @host_configs.setter
    def host_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ForwardingRuleRuleConditionHostConfigArgs']]]]):
        pulumi.set(self, "host_configs", value)

    @property
    @pulumi.getter(name="pathConfig")
    def path_config(self) -> Optional[pulumi.Input['ForwardingRuleRuleConditionPathConfigArgs']]:
        """
        Path configuration information.
        """
        return pulumi.get(self, "path_config")

    @path_config.setter
    def path_config(self, value: Optional[pulumi.Input['ForwardingRuleRuleConditionPathConfigArgs']]):
        pulumi.set(self, "path_config", value)


@pulumi.input_type
class ForwardingRuleRuleConditionHostConfigArgs:
    def __init__(__self__, *,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: The length of the path is 1-128 characters. It must start with a forward slash (/), and can only contain letters, numbers, dollar sign ($), dash (-), and underscores (_) , half width full stop (.), plus sign (+), forward slash (/), and (&), wavy line (~), at (@), half width colon (:), apostrophe ('). It supports asterisk (*) and half width question mark (?) as wildcards.
        """
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The length of the path is 1-128 characters. It must start with a forward slash (/), and can only contain letters, numbers, dollar sign ($), dash (-), and underscores (_) , half width full stop (.), plus sign (+), forward slash (/), and (&), wavy line (~), at (@), half width colon (:), apostrophe ('). It supports asterisk (*) and half width question mark (?) as wildcards.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class ForwardingRuleRuleConditionPathConfigArgs:
    def __init__(__self__, *,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: The length of the path is 1-128 characters. It must start with a forward slash (/), and can only contain letters, numbers, dollar sign ($), dash (-), and underscores (_) , half width full stop (.), plus sign (+), forward slash (/), and (&), wavy line (~), at (@), half width colon (:), apostrophe ('). It supports asterisk (*) and half width question mark (?) as wildcards.
        """
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The length of the path is 1-128 characters. It must start with a forward slash (/), and can only contain letters, numbers, dollar sign ($), dash (-), and underscores (_) , half width full stop (.), plus sign (+), forward slash (/), and (&), wavy line (~), at (@), half width colon (:), apostrophe ('). It supports asterisk (*) and half width question mark (?) as wildcards.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class ListenerCertificateArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The id of the certificate.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the certificate.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class ListenerPortRangeArgs:
    def __init__(__self__, *,
                 from_port: pulumi.Input[int],
                 to_port: pulumi.Input[int]):
        """
        :param pulumi.Input[int] from_port: The initial listening port used to receive requests and forward them to terminal nodes.
        :param pulumi.Input[int] to_port: The end listening port used to receive requests and forward them to terminal nodes.
        """
        pulumi.set(__self__, "from_port", from_port)
        pulumi.set(__self__, "to_port", to_port)

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> pulumi.Input[int]:
        """
        The initial listening port used to receive requests and forward them to terminal nodes.
        """
        return pulumi.get(self, "from_port")

    @from_port.setter
    def from_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "from_port", value)

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> pulumi.Input[int]:
        """
        The end listening port used to receive requests and forward them to terminal nodes.
        """
        return pulumi.get(self, "to_port")

    @to_port.setter
    def to_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "to_port", value)


