# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AggregateCompliancePackConfigRuleArgs',
    'AggregateCompliancePackConfigRuleConfigRuleParameterArgs',
    'AggregatorAggregatorAccountArgs',
    'CompliancePackConfigRuleArgs',
    'CompliancePackConfigRuleConfigRuleParameterArgs',
]

@pulumi.input_type
class AggregateCompliancePackConfigRuleArgs:
    def __init__(__self__, *,
                 config_rule_parameters: pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleConfigRuleParameterArgs']]],
                 managed_rule_identifier: pulumi.Input[str]):
        """
        :param pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleConfigRuleParameterArgs']]] config_rule_parameters: A list of parameter rules.
        :param pulumi.Input[str] managed_rule_identifier: The Managed Rule Identifier.
        """
        pulumi.set(__self__, "config_rule_parameters", config_rule_parameters)
        pulumi.set(__self__, "managed_rule_identifier", managed_rule_identifier)

    @property
    @pulumi.getter(name="configRuleParameters")
    def config_rule_parameters(self) -> pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleConfigRuleParameterArgs']]]:
        """
        A list of parameter rules.
        """
        return pulumi.get(self, "config_rule_parameters")

    @config_rule_parameters.setter
    def config_rule_parameters(self, value: pulumi.Input[Sequence[pulumi.Input['AggregateCompliancePackConfigRuleConfigRuleParameterArgs']]]):
        pulumi.set(self, "config_rule_parameters", value)

    @property
    @pulumi.getter(name="managedRuleIdentifier")
    def managed_rule_identifier(self) -> pulumi.Input[str]:
        """
        The Managed Rule Identifier.
        """
        return pulumi.get(self, "managed_rule_identifier")

    @managed_rule_identifier.setter
    def managed_rule_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "managed_rule_identifier", value)


@pulumi.input_type
class AggregateCompliancePackConfigRuleConfigRuleParameterArgs:
    def __init__(__self__, *,
                 parameter_name: pulumi.Input[str],
                 parameter_value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] parameter_name: The Parameter Name.
        :param pulumi.Input[str] parameter_value: The Parameter Value.
        """
        pulumi.set(__self__, "parameter_name", parameter_name)
        pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterName")
    def parameter_name(self) -> pulumi.Input[str]:
        """
        The Parameter Name.
        """
        return pulumi.get(self, "parameter_name")

    @parameter_name.setter
    def parameter_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_name", value)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> pulumi.Input[str]:
        """
        The Parameter Value.
        """
        return pulumi.get(self, "parameter_value")

    @parameter_value.setter
    def parameter_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_value", value)


@pulumi.input_type
class AggregatorAggregatorAccountArgs:
    def __init__(__self__, *,
                 account_id: pulumi.Input[str],
                 account_name: pulumi.Input[str],
                 account_type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] account_id: Aggregator account Uid.
        :param pulumi.Input[str] account_name: Aggregator account name.
        :param pulumi.Input[str] account_type: Aggregator account source type. Valid values: `ResourceDirectory`.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "account_type", account_type)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Input[str]:
        """
        Aggregator account Uid.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        Aggregator account name.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="accountType")
    def account_type(self) -> pulumi.Input[str]:
        """
        Aggregator account source type. Valid values: `ResourceDirectory`.
        """
        return pulumi.get(self, "account_type")

    @account_type.setter
    def account_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_type", value)


@pulumi.input_type
class CompliancePackConfigRuleArgs:
    def __init__(__self__, *,
                 config_rule_parameters: pulumi.Input[Sequence[pulumi.Input['CompliancePackConfigRuleConfigRuleParameterArgs']]],
                 managed_rule_identifier: pulumi.Input[str]):
        """
        :param pulumi.Input[Sequence[pulumi.Input['CompliancePackConfigRuleConfigRuleParameterArgs']]] config_rule_parameters: A list of Config Rule Parameters.
        :param pulumi.Input[str] managed_rule_identifier: The Managed Rule Identifier.
        """
        pulumi.set(__self__, "config_rule_parameters", config_rule_parameters)
        pulumi.set(__self__, "managed_rule_identifier", managed_rule_identifier)

    @property
    @pulumi.getter(name="configRuleParameters")
    def config_rule_parameters(self) -> pulumi.Input[Sequence[pulumi.Input['CompliancePackConfigRuleConfigRuleParameterArgs']]]:
        """
        A list of Config Rule Parameters.
        """
        return pulumi.get(self, "config_rule_parameters")

    @config_rule_parameters.setter
    def config_rule_parameters(self, value: pulumi.Input[Sequence[pulumi.Input['CompliancePackConfigRuleConfigRuleParameterArgs']]]):
        pulumi.set(self, "config_rule_parameters", value)

    @property
    @pulumi.getter(name="managedRuleIdentifier")
    def managed_rule_identifier(self) -> pulumi.Input[str]:
        """
        The Managed Rule Identifier.
        """
        return pulumi.get(self, "managed_rule_identifier")

    @managed_rule_identifier.setter
    def managed_rule_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "managed_rule_identifier", value)


@pulumi.input_type
class CompliancePackConfigRuleConfigRuleParameterArgs:
    def __init__(__self__, *,
                 parameter_name: pulumi.Input[str],
                 parameter_value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] parameter_name: The parameter name.
        :param pulumi.Input[str] parameter_value: The parameter value.
        """
        pulumi.set(__self__, "parameter_name", parameter_name)
        if parameter_value is not None:
            pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterName")
    def parameter_name(self) -> pulumi.Input[str]:
        """
        The parameter name.
        """
        return pulumi.get(self, "parameter_name")

    @parameter_name.setter
    def parameter_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_name", value)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> Optional[pulumi.Input[str]]:
        """
        The parameter value.
        """
        return pulumi.get(self, "parameter_value")

    @parameter_value.setter
    def parameter_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_value", value)


