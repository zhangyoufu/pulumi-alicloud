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
    'GetIpv6EgressRulesResult',
    'AwaitableGetIpv6EgressRulesResult',
    'get_ipv6_egress_rules',
    'get_ipv6_egress_rules_output',
]

@pulumi.output_type
class GetIpv6EgressRulesResult:
    """
    A collection of values returned by getIpv6EgressRules.
    """
    def __init__(__self__, id=None, ids=None, instance_id=None, ipv6_egress_rule_name=None, ipv6_gateway_id=None, name_regex=None, names=None, output_file=None, rules=None, status=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if ipv6_egress_rule_name and not isinstance(ipv6_egress_rule_name, str):
            raise TypeError("Expected argument 'ipv6_egress_rule_name' to be a str")
        pulumi.set(__self__, "ipv6_egress_rule_name", ipv6_egress_rule_name)
        if ipv6_gateway_id and not isinstance(ipv6_gateway_id, str):
            raise TypeError("Expected argument 'ipv6_gateway_id' to be a str")
        pulumi.set(__self__, "ipv6_gateway_id", ipv6_gateway_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

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
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="ipv6EgressRuleName")
    def ipv6_egress_rule_name(self) -> Optional[str]:
        return pulumi.get(self, "ipv6_egress_rule_name")

    @property
    @pulumi.getter(name="ipv6GatewayId")
    def ipv6_gateway_id(self) -> str:
        return pulumi.get(self, "ipv6_gateway_id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetIpv6EgressRulesRuleResult']:
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetIpv6EgressRulesResult(GetIpv6EgressRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpv6EgressRulesResult(
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            ipv6_egress_rule_name=self.ipv6_egress_rule_name,
            ipv6_gateway_id=self.ipv6_gateway_id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            rules=self.rules,
            status=self.status)


def get_ipv6_egress_rules(ids: Optional[Sequence[str]] = None,
                          instance_id: Optional[str] = None,
                          ipv6_egress_rule_name: Optional[str] = None,
                          ipv6_gateway_id: Optional[str] = None,
                          name_regex: Optional[str] = None,
                          output_file: Optional[str] = None,
                          status: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpv6EgressRulesResult:
    """
    This data source provides the Vpc Ipv6 Egress Rules of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.142.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.vpc.get_ipv6_egress_rules(ipv6_gateway_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("vpcIpv6EgressRuleId1", ids.rules[0].id)
    name_regex = alicloud.vpc.get_ipv6_egress_rules(ipv6_gateway_id="example_value",
        name_regex="^my-Ipv6EgressRule")
    pulumi.export("vpcIpv6EgressRuleId2", name_regex.rules[0].id)
    status = alicloud.vpc.get_ipv6_egress_rules(ipv6_gateway_id="example_value",
        status="Available")
    pulumi.export("vpcIpv6EgressRuleId3", status.rules[0].id)
    ipv6_egress_rule_name = alicloud.vpc.get_ipv6_egress_rules(ipv6_gateway_id="example_value",
        ipv6_egress_rule_name="example_value")
    pulumi.export("vpcIpv6EgressRuleId4", ipv6_egress_rule_name.rules[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Ipv6 Egress Rule IDs.
    :param str instance_id: The ID of the instance to which the egress-only rule is applied.
    :param str ipv6_egress_rule_name: The name of the resource.
    :param str ipv6_gateway_id: The ID of the IPv6 gateway.
    :param str name_regex: A regex string to filter results by Ipv6 Egress Rule name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource. Valid values: `Available`, `Pending` and `Deleting`.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['ipv6EgressRuleName'] = ipv6_egress_rule_name
    __args__['ipv6GatewayId'] = ipv6_gateway_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:vpc/getIpv6EgressRules:getIpv6EgressRules', __args__, opts=opts, typ=GetIpv6EgressRulesResult).value

    return AwaitableGetIpv6EgressRulesResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        ipv6_egress_rule_name=pulumi.get(__ret__, 'ipv6_egress_rule_name'),
        ipv6_gateway_id=pulumi.get(__ret__, 'ipv6_gateway_id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        rules=pulumi.get(__ret__, 'rules'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_ipv6_egress_rules)
def get_ipv6_egress_rules_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 ipv6_egress_rule_name: Optional[pulumi.Input[Optional[str]]] = None,
                                 ipv6_gateway_id: Optional[pulumi.Input[str]] = None,
                                 name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 status: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpv6EgressRulesResult]:
    """
    This data source provides the Vpc Ipv6 Egress Rules of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.142.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.vpc.get_ipv6_egress_rules(ipv6_gateway_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("vpcIpv6EgressRuleId1", ids.rules[0].id)
    name_regex = alicloud.vpc.get_ipv6_egress_rules(ipv6_gateway_id="example_value",
        name_regex="^my-Ipv6EgressRule")
    pulumi.export("vpcIpv6EgressRuleId2", name_regex.rules[0].id)
    status = alicloud.vpc.get_ipv6_egress_rules(ipv6_gateway_id="example_value",
        status="Available")
    pulumi.export("vpcIpv6EgressRuleId3", status.rules[0].id)
    ipv6_egress_rule_name = alicloud.vpc.get_ipv6_egress_rules(ipv6_gateway_id="example_value",
        ipv6_egress_rule_name="example_value")
    pulumi.export("vpcIpv6EgressRuleId4", ipv6_egress_rule_name.rules[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Ipv6 Egress Rule IDs.
    :param str instance_id: The ID of the instance to which the egress-only rule is applied.
    :param str ipv6_egress_rule_name: The name of the resource.
    :param str ipv6_gateway_id: The ID of the IPv6 gateway.
    :param str name_regex: A regex string to filter results by Ipv6 Egress Rule name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource. Valid values: `Available`, `Pending` and `Deleting`.
    """
    ...
