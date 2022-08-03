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

__all__ = ['AggregatorArgs', 'Aggregator']

@pulumi.input_type
class AggregatorArgs:
    def __init__(__self__, *,
                 aggregator_name: pulumi.Input[str],
                 description: pulumi.Input[str],
                 aggregator_accounts: Optional[pulumi.Input[Sequence[pulumi.Input['AggregatorAggregatorAccountArgs']]]] = None,
                 aggregator_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Aggregator resource.
        :param pulumi.Input[str] aggregator_name: The name of aggregator.
        :param pulumi.Input[str] description: The description of aggregator.
        :param pulumi.Input[Sequence[pulumi.Input['AggregatorAggregatorAccountArgs']]] aggregator_accounts: The information of account in aggregator. If the aggregator_type is RD, it is optional and means add all members in the resource directory to the account group. **NOTE:** the field `aggregator_accounts` is not required from version 1.148.0.
        :param pulumi.Input[str] aggregator_type: The type of aggregator. Valid values: `CUSTOM`, `RD`. The Default value: `CUSTOM`.
        """
        pulumi.set(__self__, "aggregator_name", aggregator_name)
        pulumi.set(__self__, "description", description)
        if aggregator_accounts is not None:
            pulumi.set(__self__, "aggregator_accounts", aggregator_accounts)
        if aggregator_type is not None:
            pulumi.set(__self__, "aggregator_type", aggregator_type)

    @property
    @pulumi.getter(name="aggregatorName")
    def aggregator_name(self) -> pulumi.Input[str]:
        """
        The name of aggregator.
        """
        return pulumi.get(self, "aggregator_name")

    @aggregator_name.setter
    def aggregator_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "aggregator_name", value)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        The description of aggregator.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="aggregatorAccounts")
    def aggregator_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AggregatorAggregatorAccountArgs']]]]:
        """
        The information of account in aggregator. If the aggregator_type is RD, it is optional and means add all members in the resource directory to the account group. **NOTE:** the field `aggregator_accounts` is not required from version 1.148.0.
        """
        return pulumi.get(self, "aggregator_accounts")

    @aggregator_accounts.setter
    def aggregator_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AggregatorAggregatorAccountArgs']]]]):
        pulumi.set(self, "aggregator_accounts", value)

    @property
    @pulumi.getter(name="aggregatorType")
    def aggregator_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of aggregator. Valid values: `CUSTOM`, `RD`. The Default value: `CUSTOM`.
        """
        return pulumi.get(self, "aggregator_type")

    @aggregator_type.setter
    def aggregator_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aggregator_type", value)


@pulumi.input_type
class _AggregatorState:
    def __init__(__self__, *,
                 aggregator_accounts: Optional[pulumi.Input[Sequence[pulumi.Input['AggregatorAggregatorAccountArgs']]]] = None,
                 aggregator_name: Optional[pulumi.Input[str]] = None,
                 aggregator_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Aggregator resources.
        :param pulumi.Input[Sequence[pulumi.Input['AggregatorAggregatorAccountArgs']]] aggregator_accounts: The information of account in aggregator. If the aggregator_type is RD, it is optional and means add all members in the resource directory to the account group. **NOTE:** the field `aggregator_accounts` is not required from version 1.148.0.
        :param pulumi.Input[str] aggregator_name: The name of aggregator.
        :param pulumi.Input[str] aggregator_type: The type of aggregator. Valid values: `CUSTOM`, `RD`. The Default value: `CUSTOM`.
        :param pulumi.Input[str] description: The description of aggregator.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `0`: creating `1`: normal `2`: deleting.
        """
        if aggregator_accounts is not None:
            pulumi.set(__self__, "aggregator_accounts", aggregator_accounts)
        if aggregator_name is not None:
            pulumi.set(__self__, "aggregator_name", aggregator_name)
        if aggregator_type is not None:
            pulumi.set(__self__, "aggregator_type", aggregator_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="aggregatorAccounts")
    def aggregator_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AggregatorAggregatorAccountArgs']]]]:
        """
        The information of account in aggregator. If the aggregator_type is RD, it is optional and means add all members in the resource directory to the account group. **NOTE:** the field `aggregator_accounts` is not required from version 1.148.0.
        """
        return pulumi.get(self, "aggregator_accounts")

    @aggregator_accounts.setter
    def aggregator_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AggregatorAggregatorAccountArgs']]]]):
        pulumi.set(self, "aggregator_accounts", value)

    @property
    @pulumi.getter(name="aggregatorName")
    def aggregator_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of aggregator.
        """
        return pulumi.get(self, "aggregator_name")

    @aggregator_name.setter
    def aggregator_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aggregator_name", value)

    @property
    @pulumi.getter(name="aggregatorType")
    def aggregator_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of aggregator. Valid values: `CUSTOM`, `RD`. The Default value: `CUSTOM`.
        """
        return pulumi.get(self, "aggregator_type")

    @aggregator_type.setter
    def aggregator_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aggregator_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of aggregator.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource. Valid values: `0`: creating `1`: normal `2`: deleting.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Aggregator(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aggregator_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregatorAggregatorAccountArgs']]]]] = None,
                 aggregator_name: Optional[pulumi.Input[str]] = None,
                 aggregator_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cloud Config Aggregator resource.

        For information about Cloud Config Aggregate Config Rule and how to use it, see [What is Aggregator](https://www.alibabacloud.com/help/en/doc-detail/211197.html).

        > **NOTE:** Available in v1.124.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.cfg.Aggregator("example",
            aggregator_accounts=[alicloud.cfg.AggregatorAggregatorAccountArgs(
                account_id="123968452689****",
                account_name="tf-testacc1234",
                account_type="ResourceDirectory",
            )],
            aggregator_name="tf-testaccConfigAggregator1234",
            description="tf-testaccConfigAggregator1234")
        ```

        ## Import

        Cloud Config Aggregator can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cfg/aggregator:Aggregator example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregatorAggregatorAccountArgs']]]] aggregator_accounts: The information of account in aggregator. If the aggregator_type is RD, it is optional and means add all members in the resource directory to the account group. **NOTE:** the field `aggregator_accounts` is not required from version 1.148.0.
        :param pulumi.Input[str] aggregator_name: The name of aggregator.
        :param pulumi.Input[str] aggregator_type: The type of aggregator. Valid values: `CUSTOM`, `RD`. The Default value: `CUSTOM`.
        :param pulumi.Input[str] description: The description of aggregator.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AggregatorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Config Aggregator resource.

        For information about Cloud Config Aggregate Config Rule and how to use it, see [What is Aggregator](https://www.alibabacloud.com/help/en/doc-detail/211197.html).

        > **NOTE:** Available in v1.124.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.cfg.Aggregator("example",
            aggregator_accounts=[alicloud.cfg.AggregatorAggregatorAccountArgs(
                account_id="123968452689****",
                account_name="tf-testacc1234",
                account_type="ResourceDirectory",
            )],
            aggregator_name="tf-testaccConfigAggregator1234",
            description="tf-testaccConfigAggregator1234")
        ```

        ## Import

        Cloud Config Aggregator can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cfg/aggregator:Aggregator example <id>
        ```

        :param str resource_name: The name of the resource.
        :param AggregatorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AggregatorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aggregator_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregatorAggregatorAccountArgs']]]]] = None,
                 aggregator_name: Optional[pulumi.Input[str]] = None,
                 aggregator_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AggregatorArgs.__new__(AggregatorArgs)

            __props__.__dict__["aggregator_accounts"] = aggregator_accounts
            if aggregator_name is None and not opts.urn:
                raise TypeError("Missing required property 'aggregator_name'")
            __props__.__dict__["aggregator_name"] = aggregator_name
            __props__.__dict__["aggregator_type"] = aggregator_type
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["status"] = None
        super(Aggregator, __self__).__init__(
            'alicloud:cfg/aggregator:Aggregator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aggregator_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregatorAggregatorAccountArgs']]]]] = None,
            aggregator_name: Optional[pulumi.Input[str]] = None,
            aggregator_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Aggregator':
        """
        Get an existing Aggregator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AggregatorAggregatorAccountArgs']]]] aggregator_accounts: The information of account in aggregator. If the aggregator_type is RD, it is optional and means add all members in the resource directory to the account group. **NOTE:** the field `aggregator_accounts` is not required from version 1.148.0.
        :param pulumi.Input[str] aggregator_name: The name of aggregator.
        :param pulumi.Input[str] aggregator_type: The type of aggregator. Valid values: `CUSTOM`, `RD`. The Default value: `CUSTOM`.
        :param pulumi.Input[str] description: The description of aggregator.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `0`: creating `1`: normal `2`: deleting.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AggregatorState.__new__(_AggregatorState)

        __props__.__dict__["aggregator_accounts"] = aggregator_accounts
        __props__.__dict__["aggregator_name"] = aggregator_name
        __props__.__dict__["aggregator_type"] = aggregator_type
        __props__.__dict__["description"] = description
        __props__.__dict__["status"] = status
        return Aggregator(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aggregatorAccounts")
    def aggregator_accounts(self) -> pulumi.Output[Sequence['outputs.AggregatorAggregatorAccount']]:
        """
        The information of account in aggregator. If the aggregator_type is RD, it is optional and means add all members in the resource directory to the account group. **NOTE:** the field `aggregator_accounts` is not required from version 1.148.0.
        """
        return pulumi.get(self, "aggregator_accounts")

    @property
    @pulumi.getter(name="aggregatorName")
    def aggregator_name(self) -> pulumi.Output[str]:
        """
        The name of aggregator.
        """
        return pulumi.get(self, "aggregator_name")

    @property
    @pulumi.getter(name="aggregatorType")
    def aggregator_type(self) -> pulumi.Output[str]:
        """
        The type of aggregator. Valid values: `CUSTOM`, `RD`. The Default value: `CUSTOM`.
        """
        return pulumi.get(self, "aggregator_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of aggregator.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource. Valid values: `0`: creating `1`: normal `2`: deleting.
        """
        return pulumi.get(self, "status")

