# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConfigArgs', 'Config']

@pulumi.input_type
class ConfigArgs:
    def __init__(__self__, *,
                 code: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Config resource.
        :param pulumi.Input[str] code: Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
        :param pulumi.Input[str] description: Abnormal Alarm General Description of the Configuration Item.
        :param pulumi.Input[str] lang: The language of the request and response. Valid values: `zh`,`en`.
               * `zh`: Chinese.
               * `en`: English.
        :param pulumi.Input[str] value: The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
               * `access_failed_cnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
               * `access_permission_exprie_max_days`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
               * `log_datasize_avg_days`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input[str]]:
        """
        Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Abnormal Alarm General Description of the Configuration Item.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        The language of the request and response. Valid values: `zh`,`en`.
        * `zh`: Chinese.
        * `en`: English.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
        * `access_failed_cnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
        * `access_permission_exprie_max_days`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
        * `log_datasize_avg_days`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class _ConfigState:
    def __init__(__self__, *,
                 code: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Config resources.
        :param pulumi.Input[str] code: Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
        :param pulumi.Input[str] description: Abnormal Alarm General Description of the Configuration Item.
        :param pulumi.Input[str] lang: The language of the request and response. Valid values: `zh`,`en`.
               * `zh`: Chinese.
               * `en`: English.
        :param pulumi.Input[str] value: The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
               * `access_failed_cnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
               * `access_permission_exprie_max_days`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
               * `log_datasize_avg_days`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input[str]]:
        """
        Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Abnormal Alarm General Description of the Configuration Item.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        The language of the request and response. Valid values: `zh`,`en`.
        * `zh`: Chinese.
        * `en`: English.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
        * `access_failed_cnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
        * `access_permission_exprie_max_days`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
        * `log_datasize_avg_days`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class Config(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 code: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Data Security Center Config resource.

        For information about Data Security Center Config and how to use it, see [What is Config](https://help.aliyun.com/product/88674.html).

        > **NOTE:** Available in v1.133.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.sddp.Config("default",
            code="access_failed_cnt",
            value="10")
        ```

        ## Import

        Data Security Center Config can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:sddp/config:Config example <code>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] code: Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
        :param pulumi.Input[str] description: Abnormal Alarm General Description of the Configuration Item.
        :param pulumi.Input[str] lang: The language of the request and response. Valid values: `zh`,`en`.
               * `zh`: Chinese.
               * `en`: English.
        :param pulumi.Input[str] value: The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
               * `access_failed_cnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
               * `access_permission_exprie_max_days`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
               * `log_datasize_avg_days`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ConfigArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Data Security Center Config resource.

        For information about Data Security Center Config and how to use it, see [What is Config](https://help.aliyun.com/product/88674.html).

        > **NOTE:** Available in v1.133.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.sddp.Config("default",
            code="access_failed_cnt",
            value="10")
        ```

        ## Import

        Data Security Center Config can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:sddp/config:Config example <code>
        ```

        :param str resource_name: The name of the resource.
        :param ConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 code: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigArgs.__new__(ConfigArgs)

            __props__.__dict__["code"] = code
            __props__.__dict__["description"] = description
            __props__.__dict__["lang"] = lang
            __props__.__dict__["value"] = value
        super(Config, __self__).__init__(
            'alicloud:sddp/config:Config',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            code: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            lang: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'Config':
        """
        Get an existing Config resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] code: Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
        :param pulumi.Input[str] description: Abnormal Alarm General Description of the Configuration Item.
        :param pulumi.Input[str] lang: The language of the request and response. Valid values: `zh`,`en`.
               * `zh`: Chinese.
               * `en`: English.
        :param pulumi.Input[str] value: The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
               * `access_failed_cnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
               * `access_permission_exprie_max_days`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
               * `log_datasize_avg_days`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigState.__new__(_ConfigState)

        __props__.__dict__["code"] = code
        __props__.__dict__["description"] = description
        __props__.__dict__["lang"] = lang
        __props__.__dict__["value"] = value
        return Config(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def code(self) -> pulumi.Output[Optional[str]]:
        """
        Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Abnormal Alarm General Description of the Configuration Item.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def lang(self) -> pulumi.Output[Optional[str]]:
        """
        The language of the request and response. Valid values: `zh`,`en`.
        * `zh`: Chinese.
        * `en`: English.
        """
        return pulumi.get(self, "lang")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[Optional[str]]:
        """
        The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
        * `access_failed_cnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
        * `access_permission_exprie_max_days`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
        * `log_datasize_avg_days`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
        """
        return pulumi.get(self, "value")

