# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['JobMonitorRuleArgs', 'JobMonitorRule']

@pulumi.input_type
class JobMonitorRuleArgs:
    def __init__(__self__, *,
                 dts_job_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 delay_rule_time: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a JobMonitorRule resource.
        :param pulumi.Input[str] dts_job_id: Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
        :param pulumi.Input[str] type: Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
        :param pulumi.Input[str] delay_rule_time: Trigger delay alarm threshold, which is measured in seconds.
        :param pulumi.Input[str] phone: The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
        :param pulumi.Input[str] state: Whether to enable monitoring rules, valid values: `Y`, `N`.
        """
        pulumi.set(__self__, "dts_job_id", dts_job_id)
        pulumi.set(__self__, "type", type)
        if delay_rule_time is not None:
            pulumi.set(__self__, "delay_rule_time", delay_rule_time)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="dtsJobId")
    def dts_job_id(self) -> pulumi.Input[str]:
        """
        Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
        """
        return pulumi.get(self, "dts_job_id")

    @dts_job_id.setter
    def dts_job_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dts_job_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="delayRuleTime")
    def delay_rule_time(self) -> Optional[pulumi.Input[str]]:
        """
        Trigger delay alarm threshold, which is measured in seconds.
        """
        return pulumi.get(self, "delay_rule_time")

    @delay_rule_time.setter
    def delay_rule_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delay_rule_time", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[str]]:
        """
        The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Whether to enable monitoring rules, valid values: `Y`, `N`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


@pulumi.input_type
class _JobMonitorRuleState:
    def __init__(__self__, *,
                 delay_rule_time: Optional[pulumi.Input[str]] = None,
                 dts_job_id: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering JobMonitorRule resources.
        :param pulumi.Input[str] delay_rule_time: Trigger delay alarm threshold, which is measured in seconds.
        :param pulumi.Input[str] dts_job_id: Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
        :param pulumi.Input[str] phone: The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
        :param pulumi.Input[str] state: Whether to enable monitoring rules, valid values: `Y`, `N`.
        :param pulumi.Input[str] type: Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
        """
        if delay_rule_time is not None:
            pulumi.set(__self__, "delay_rule_time", delay_rule_time)
        if dts_job_id is not None:
            pulumi.set(__self__, "dts_job_id", dts_job_id)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="delayRuleTime")
    def delay_rule_time(self) -> Optional[pulumi.Input[str]]:
        """
        Trigger delay alarm threshold, which is measured in seconds.
        """
        return pulumi.get(self, "delay_rule_time")

    @delay_rule_time.setter
    def delay_rule_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delay_rule_time", value)

    @property
    @pulumi.getter(name="dtsJobId")
    def dts_job_id(self) -> Optional[pulumi.Input[str]]:
        """
        Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
        """
        return pulumi.get(self, "dts_job_id")

    @dts_job_id.setter
    def dts_job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dts_job_id", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[str]]:
        """
        The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Whether to enable monitoring rules, valid values: `Y`, `N`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class JobMonitorRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delay_rule_time: Optional[pulumi.Input[str]] = None,
                 dts_job_id: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a DTS Job Monitor Rule resource.

        For information about DTS Job Monitor Rule and how to use it, see [What is Job Monitor Rule](https://www.aliyun.com/product/dts).

        > **NOTE:** Available since v1.134.0.

        ## Import

        DTS Job Monitor Rule can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dts/jobMonitorRule:JobMonitorRule example <dts_job_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] delay_rule_time: Trigger delay alarm threshold, which is measured in seconds.
        :param pulumi.Input[str] dts_job_id: Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
        :param pulumi.Input[str] phone: The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
        :param pulumi.Input[str] state: Whether to enable monitoring rules, valid values: `Y`, `N`.
        :param pulumi.Input[str] type: Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobMonitorRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a DTS Job Monitor Rule resource.

        For information about DTS Job Monitor Rule and how to use it, see [What is Job Monitor Rule](https://www.aliyun.com/product/dts).

        > **NOTE:** Available since v1.134.0.

        ## Import

        DTS Job Monitor Rule can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dts/jobMonitorRule:JobMonitorRule example <dts_job_id>
        ```

        :param str resource_name: The name of the resource.
        :param JobMonitorRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobMonitorRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delay_rule_time: Optional[pulumi.Input[str]] = None,
                 dts_job_id: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = JobMonitorRuleArgs.__new__(JobMonitorRuleArgs)

            __props__.__dict__["delay_rule_time"] = delay_rule_time
            if dts_job_id is None and not opts.urn:
                raise TypeError("Missing required property 'dts_job_id'")
            __props__.__dict__["dts_job_id"] = dts_job_id
            __props__.__dict__["phone"] = phone
            __props__.__dict__["state"] = state
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(JobMonitorRule, __self__).__init__(
            'alicloud:dts/jobMonitorRule:JobMonitorRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            delay_rule_time: Optional[pulumi.Input[str]] = None,
            dts_job_id: Optional[pulumi.Input[str]] = None,
            phone: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'JobMonitorRule':
        """
        Get an existing JobMonitorRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] delay_rule_time: Trigger delay alarm threshold, which is measured in seconds.
        :param pulumi.Input[str] dts_job_id: Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
        :param pulumi.Input[str] phone: The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
        :param pulumi.Input[str] state: Whether to enable monitoring rules, valid values: `Y`, `N`.
        :param pulumi.Input[str] type: Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _JobMonitorRuleState.__new__(_JobMonitorRuleState)

        __props__.__dict__["delay_rule_time"] = delay_rule_time
        __props__.__dict__["dts_job_id"] = dts_job_id
        __props__.__dict__["phone"] = phone
        __props__.__dict__["state"] = state
        __props__.__dict__["type"] = type
        return JobMonitorRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="delayRuleTime")
    def delay_rule_time(self) -> pulumi.Output[str]:
        """
        Trigger delay alarm threshold, which is measured in seconds.
        """
        return pulumi.get(self, "delay_rule_time")

    @property
    @pulumi.getter(name="dtsJobId")
    def dts_job_id(self) -> pulumi.Output[str]:
        """
        Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
        """
        return pulumi.get(self, "dts_job_id")

    @property
    @pulumi.getter
    def phone(self) -> pulumi.Output[Optional[str]]:
        """
        The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
        """
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Whether to enable monitoring rules, valid values: `Y`, `N`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
        """
        return pulumi.get(self, "type")

