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
    'GetAggregateConfigRulesResult',
    'AwaitableGetAggregateConfigRulesResult',
    'get_aggregate_config_rules',
    'get_aggregate_config_rules_output',
]

@pulumi.output_type
class GetAggregateConfigRulesResult:
    """
    A collection of values returned by getAggregateConfigRules.
    """
    def __init__(__self__, aggregate_config_rule_name=None, aggregator_id=None, enable_details=None, id=None, ids=None, name_regex=None, names=None, output_file=None, risk_level=None, rules=None, status=None):
        if aggregate_config_rule_name and not isinstance(aggregate_config_rule_name, str):
            raise TypeError("Expected argument 'aggregate_config_rule_name' to be a str")
        pulumi.set(__self__, "aggregate_config_rule_name", aggregate_config_rule_name)
        if aggregator_id and not isinstance(aggregator_id, str):
            raise TypeError("Expected argument 'aggregator_id' to be a str")
        pulumi.set(__self__, "aggregator_id", aggregator_id)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if risk_level and not isinstance(risk_level, int):
            raise TypeError("Expected argument 'risk_level' to be a int")
        pulumi.set(__self__, "risk_level", risk_level)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="aggregateConfigRuleName")
    def aggregate_config_rule_name(self) -> Optional[str]:
        """
        The name of the rule.
        """
        return pulumi.get(self, "aggregate_config_rule_name")

    @property
    @pulumi.getter(name="aggregatorId")
    def aggregator_id(self) -> str:
        """
        The ID of Aggregator.
        """
        return pulumi.get(self, "aggregator_id")

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of Aggregate Config Rule names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="riskLevel")
    def risk_level(self) -> Optional[int]:
        """
        The risk level of the resources that are not compliant with the rule. Valid values: `1`: critical, `2`: warning, `3`: info.
        """
        return pulumi.get(self, "risk_level")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetAggregateConfigRulesRuleResult']:
        """
        A list of Config Aggregate Config Rules. Each element contains the following attributes:
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the rule.
        """
        return pulumi.get(self, "status")


class AwaitableGetAggregateConfigRulesResult(GetAggregateConfigRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAggregateConfigRulesResult(
            aggregate_config_rule_name=self.aggregate_config_rule_name,
            aggregator_id=self.aggregator_id,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            risk_level=self.risk_level,
            rules=self.rules,
            status=self.status)


def get_aggregate_config_rules(aggregate_config_rule_name: Optional[str] = None,
                               aggregator_id: Optional[str] = None,
                               enable_details: Optional[bool] = None,
                               ids: Optional[Sequence[str]] = None,
                               name_regex: Optional[str] = None,
                               output_file: Optional[str] = None,
                               risk_level: Optional[int] = None,
                               status: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAggregateConfigRulesResult:
    """
    This data source provides the Config Aggregate Config Rules of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.124.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cfg.get_aggregate_config_rules(aggregator_id="ca-3a9b626622af001d****",
        ids=["cr-5154626622af0034****"],
        name_regex="the_resource_name")
    pulumi.export("firstConfigAggregateConfigRuleId", example.rules[0].id)
    ```


    :param str aggregate_config_rule_name: The config rule name.
    :param str aggregator_id: The ID of aggregator.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Aggregate Config Rule IDs.
    :param str name_regex: A regex string to filter results by Aggregate Config Rule name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param int risk_level: Optional, ForceNew) The Risk Level. Valid values `1`: critical, `2`: warning, `3`: info.
    :param str status: The state of the config rule, valid values: `ACTIVE`, `DELETING`, `EVALUATING` and `INACTIVE`.
    """
    __args__ = dict()
    __args__['aggregateConfigRuleName'] = aggregate_config_rule_name
    __args__['aggregatorId'] = aggregator_id
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['riskLevel'] = risk_level
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cfg/getAggregateConfigRules:getAggregateConfigRules', __args__, opts=opts, typ=GetAggregateConfigRulesResult).value

    return AwaitableGetAggregateConfigRulesResult(
        aggregate_config_rule_name=pulumi.get(__ret__, 'aggregate_config_rule_name'),
        aggregator_id=pulumi.get(__ret__, 'aggregator_id'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        risk_level=pulumi.get(__ret__, 'risk_level'),
        rules=pulumi.get(__ret__, 'rules'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_aggregate_config_rules)
def get_aggregate_config_rules_output(aggregate_config_rule_name: Optional[pulumi.Input[Optional[str]]] = None,
                                      aggregator_id: Optional[pulumi.Input[str]] = None,
                                      enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                                      ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                      name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                      risk_level: Optional[pulumi.Input[Optional[int]]] = None,
                                      status: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAggregateConfigRulesResult]:
    """
    This data source provides the Config Aggregate Config Rules of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.124.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cfg.get_aggregate_config_rules(aggregator_id="ca-3a9b626622af001d****",
        ids=["cr-5154626622af0034****"],
        name_regex="the_resource_name")
    pulumi.export("firstConfigAggregateConfigRuleId", example.rules[0].id)
    ```


    :param str aggregate_config_rule_name: The config rule name.
    :param str aggregator_id: The ID of aggregator.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Aggregate Config Rule IDs.
    :param str name_regex: A regex string to filter results by Aggregate Config Rule name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param int risk_level: Optional, ForceNew) The Risk Level. Valid values `1`: critical, `2`: warning, `3`: info.
    :param str status: The state of the config rule, valid values: `ACTIVE`, `DELETING`, `EVALUATING` and `INACTIVE`.
    """
    ...
