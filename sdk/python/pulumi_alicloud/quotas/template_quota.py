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

__all__ = ['TemplateQuotaArgs', 'TemplateQuota']

@pulumi.input_type
class TemplateQuotaArgs:
    def __init__(__self__, *,
                 desire_value: pulumi.Input[float],
                 product_code: pulumi.Input[str],
                 quota_action_code: pulumi.Input[str],
                 dimensions: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateQuotaDimensionArgs']]]] = None,
                 effective_time: Optional[pulumi.Input[str]] = None,
                 env_language: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 notice_type: Optional[pulumi.Input[int]] = None,
                 quota_category: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TemplateQuota resource.
        :param pulumi.Input[float] desire_value: Quota application value.
        :param pulumi.Input[str] product_code: The abbreviation of the cloud service name.
        :param pulumi.Input[str] quota_action_code: The quota ID.
        :param pulumi.Input[Sequence[pulumi.Input['TemplateQuotaDimensionArgs']]] dimensions: The Quota Dimensions. See the following `Block Dimensions`.
        :param pulumi.Input[str] effective_time: The UTC time when the quota takes effect.
        :param pulumi.Input[str] env_language: The language of the quota alert notification. Value:
               - zh: Chinese.
               - en: English.
        :param pulumi.Input[str] expire_time: The UTC time when the quota expires.
        :param pulumi.Input[int] notice_type: Whether to notify the result of quota promotion application. Value:
               - 0: No.
               - 3: Yes.
        :param pulumi.Input[str] quota_category: Type of quota. Value:
               - CommonQuota : Generic quota.
               - WhiteListLabel: Equity quota.
               - FlowControl:API rate quota.
        """
        pulumi.set(__self__, "desire_value", desire_value)
        pulumi.set(__self__, "product_code", product_code)
        pulumi.set(__self__, "quota_action_code", quota_action_code)
        if dimensions is not None:
            pulumi.set(__self__, "dimensions", dimensions)
        if effective_time is not None:
            pulumi.set(__self__, "effective_time", effective_time)
        if env_language is not None:
            pulumi.set(__self__, "env_language", env_language)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if notice_type is not None:
            pulumi.set(__self__, "notice_type", notice_type)
        if quota_category is not None:
            pulumi.set(__self__, "quota_category", quota_category)

    @property
    @pulumi.getter(name="desireValue")
    def desire_value(self) -> pulumi.Input[float]:
        """
        Quota application value.
        """
        return pulumi.get(self, "desire_value")

    @desire_value.setter
    def desire_value(self, value: pulumi.Input[float]):
        pulumi.set(self, "desire_value", value)

    @property
    @pulumi.getter(name="productCode")
    def product_code(self) -> pulumi.Input[str]:
        """
        The abbreviation of the cloud service name.
        """
        return pulumi.get(self, "product_code")

    @product_code.setter
    def product_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "product_code", value)

    @property
    @pulumi.getter(name="quotaActionCode")
    def quota_action_code(self) -> pulumi.Input[str]:
        """
        The quota ID.
        """
        return pulumi.get(self, "quota_action_code")

    @quota_action_code.setter
    def quota_action_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "quota_action_code", value)

    @property
    @pulumi.getter
    def dimensions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TemplateQuotaDimensionArgs']]]]:
        """
        The Quota Dimensions. See the following `Block Dimensions`.
        """
        return pulumi.get(self, "dimensions")

    @dimensions.setter
    def dimensions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateQuotaDimensionArgs']]]]):
        pulumi.set(self, "dimensions", value)

    @property
    @pulumi.getter(name="effectiveTime")
    def effective_time(self) -> Optional[pulumi.Input[str]]:
        """
        The UTC time when the quota takes effect.
        """
        return pulumi.get(self, "effective_time")

    @effective_time.setter
    def effective_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "effective_time", value)

    @property
    @pulumi.getter(name="envLanguage")
    def env_language(self) -> Optional[pulumi.Input[str]]:
        """
        The language of the quota alert notification. Value:
        - zh: Chinese.
        - en: English.
        """
        return pulumi.get(self, "env_language")

    @env_language.setter
    def env_language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "env_language", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        The UTC time when the quota expires.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter(name="noticeType")
    def notice_type(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to notify the result of quota promotion application. Value:
        - 0: No.
        - 3: Yes.
        """
        return pulumi.get(self, "notice_type")

    @notice_type.setter
    def notice_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "notice_type", value)

    @property
    @pulumi.getter(name="quotaCategory")
    def quota_category(self) -> Optional[pulumi.Input[str]]:
        """
        Type of quota. Value:
        - CommonQuota : Generic quota.
        - WhiteListLabel: Equity quota.
        - FlowControl:API rate quota.
        """
        return pulumi.get(self, "quota_category")

    @quota_category.setter
    def quota_category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quota_category", value)


@pulumi.input_type
class _TemplateQuotaState:
    def __init__(__self__, *,
                 desire_value: Optional[pulumi.Input[float]] = None,
                 dimensions: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateQuotaDimensionArgs']]]] = None,
                 effective_time: Optional[pulumi.Input[str]] = None,
                 env_language: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 notice_type: Optional[pulumi.Input[int]] = None,
                 product_code: Optional[pulumi.Input[str]] = None,
                 quota_action_code: Optional[pulumi.Input[str]] = None,
                 quota_category: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TemplateQuota resources.
        :param pulumi.Input[float] desire_value: Quota application value.
        :param pulumi.Input[Sequence[pulumi.Input['TemplateQuotaDimensionArgs']]] dimensions: The Quota Dimensions. See the following `Block Dimensions`.
        :param pulumi.Input[str] effective_time: The UTC time when the quota takes effect.
        :param pulumi.Input[str] env_language: The language of the quota alert notification. Value:
               - zh: Chinese.
               - en: English.
        :param pulumi.Input[str] expire_time: The UTC time when the quota expires.
        :param pulumi.Input[int] notice_type: Whether to notify the result of quota promotion application. Value:
               - 0: No.
               - 3: Yes.
        :param pulumi.Input[str] product_code: The abbreviation of the cloud service name.
        :param pulumi.Input[str] quota_action_code: The quota ID.
        :param pulumi.Input[str] quota_category: Type of quota. Value:
               - CommonQuota : Generic quota.
               - WhiteListLabel: Equity quota.
               - FlowControl:API rate quota.
        """
        if desire_value is not None:
            pulumi.set(__self__, "desire_value", desire_value)
        if dimensions is not None:
            pulumi.set(__self__, "dimensions", dimensions)
        if effective_time is not None:
            pulumi.set(__self__, "effective_time", effective_time)
        if env_language is not None:
            pulumi.set(__self__, "env_language", env_language)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if notice_type is not None:
            pulumi.set(__self__, "notice_type", notice_type)
        if product_code is not None:
            pulumi.set(__self__, "product_code", product_code)
        if quota_action_code is not None:
            pulumi.set(__self__, "quota_action_code", quota_action_code)
        if quota_category is not None:
            pulumi.set(__self__, "quota_category", quota_category)

    @property
    @pulumi.getter(name="desireValue")
    def desire_value(self) -> Optional[pulumi.Input[float]]:
        """
        Quota application value.
        """
        return pulumi.get(self, "desire_value")

    @desire_value.setter
    def desire_value(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "desire_value", value)

    @property
    @pulumi.getter
    def dimensions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TemplateQuotaDimensionArgs']]]]:
        """
        The Quota Dimensions. See the following `Block Dimensions`.
        """
        return pulumi.get(self, "dimensions")

    @dimensions.setter
    def dimensions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateQuotaDimensionArgs']]]]):
        pulumi.set(self, "dimensions", value)

    @property
    @pulumi.getter(name="effectiveTime")
    def effective_time(self) -> Optional[pulumi.Input[str]]:
        """
        The UTC time when the quota takes effect.
        """
        return pulumi.get(self, "effective_time")

    @effective_time.setter
    def effective_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "effective_time", value)

    @property
    @pulumi.getter(name="envLanguage")
    def env_language(self) -> Optional[pulumi.Input[str]]:
        """
        The language of the quota alert notification. Value:
        - zh: Chinese.
        - en: English.
        """
        return pulumi.get(self, "env_language")

    @env_language.setter
    def env_language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "env_language", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        The UTC time when the quota expires.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter(name="noticeType")
    def notice_type(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to notify the result of quota promotion application. Value:
        - 0: No.
        - 3: Yes.
        """
        return pulumi.get(self, "notice_type")

    @notice_type.setter
    def notice_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "notice_type", value)

    @property
    @pulumi.getter(name="productCode")
    def product_code(self) -> Optional[pulumi.Input[str]]:
        """
        The abbreviation of the cloud service name.
        """
        return pulumi.get(self, "product_code")

    @product_code.setter
    def product_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_code", value)

    @property
    @pulumi.getter(name="quotaActionCode")
    def quota_action_code(self) -> Optional[pulumi.Input[str]]:
        """
        The quota ID.
        """
        return pulumi.get(self, "quota_action_code")

    @quota_action_code.setter
    def quota_action_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quota_action_code", value)

    @property
    @pulumi.getter(name="quotaCategory")
    def quota_category(self) -> Optional[pulumi.Input[str]]:
        """
        Type of quota. Value:
        - CommonQuota : Generic quota.
        - WhiteListLabel: Equity quota.
        - FlowControl:API rate quota.
        """
        return pulumi.get(self, "quota_category")

    @quota_category.setter
    def quota_category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quota_category", value)


class TemplateQuota(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 desire_value: Optional[pulumi.Input[float]] = None,
                 dimensions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateQuotaDimensionArgs']]]]] = None,
                 effective_time: Optional[pulumi.Input[str]] = None,
                 env_language: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 notice_type: Optional[pulumi.Input[int]] = None,
                 product_code: Optional[pulumi.Input[str]] = None,
                 quota_action_code: Optional[pulumi.Input[str]] = None,
                 quota_category: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Quotas Template Quota resource.

        For information about Quotas Template Quota and how to use it, see [What is Template Quota](https://help.aliyun.com/document_detail/450615.html).

        > **NOTE:** Available in v1.206.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.quotas.TemplateQuota("default",
            desire_value=1001,
            dimensions=[alicloud.quotas.TemplateQuotaDimensionArgs(
                key="regionId",
                value="cn-hangzhou",
            )],
            env_language="zh",
            notice_type=3,
            product_code="gws",
            quota_action_code="q_desktop-count",
            quota_category="CommonQuota")
        ```

        ## Import

        Quotas Template Quota can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:quotas/templateQuota:TemplateQuota example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] desire_value: Quota application value.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateQuotaDimensionArgs']]]] dimensions: The Quota Dimensions. See the following `Block Dimensions`.
        :param pulumi.Input[str] effective_time: The UTC time when the quota takes effect.
        :param pulumi.Input[str] env_language: The language of the quota alert notification. Value:
               - zh: Chinese.
               - en: English.
        :param pulumi.Input[str] expire_time: The UTC time when the quota expires.
        :param pulumi.Input[int] notice_type: Whether to notify the result of quota promotion application. Value:
               - 0: No.
               - 3: Yes.
        :param pulumi.Input[str] product_code: The abbreviation of the cloud service name.
        :param pulumi.Input[str] quota_action_code: The quota ID.
        :param pulumi.Input[str] quota_category: Type of quota. Value:
               - CommonQuota : Generic quota.
               - WhiteListLabel: Equity quota.
               - FlowControl:API rate quota.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TemplateQuotaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Quotas Template Quota resource.

        For information about Quotas Template Quota and how to use it, see [What is Template Quota](https://help.aliyun.com/document_detail/450615.html).

        > **NOTE:** Available in v1.206.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.quotas.TemplateQuota("default",
            desire_value=1001,
            dimensions=[alicloud.quotas.TemplateQuotaDimensionArgs(
                key="regionId",
                value="cn-hangzhou",
            )],
            env_language="zh",
            notice_type=3,
            product_code="gws",
            quota_action_code="q_desktop-count",
            quota_category="CommonQuota")
        ```

        ## Import

        Quotas Template Quota can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:quotas/templateQuota:TemplateQuota example <id>
        ```

        :param str resource_name: The name of the resource.
        :param TemplateQuotaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TemplateQuotaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 desire_value: Optional[pulumi.Input[float]] = None,
                 dimensions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateQuotaDimensionArgs']]]]] = None,
                 effective_time: Optional[pulumi.Input[str]] = None,
                 env_language: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 notice_type: Optional[pulumi.Input[int]] = None,
                 product_code: Optional[pulumi.Input[str]] = None,
                 quota_action_code: Optional[pulumi.Input[str]] = None,
                 quota_category: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TemplateQuotaArgs.__new__(TemplateQuotaArgs)

            if desire_value is None and not opts.urn:
                raise TypeError("Missing required property 'desire_value'")
            __props__.__dict__["desire_value"] = desire_value
            __props__.__dict__["dimensions"] = dimensions
            __props__.__dict__["effective_time"] = effective_time
            __props__.__dict__["env_language"] = env_language
            __props__.__dict__["expire_time"] = expire_time
            __props__.__dict__["notice_type"] = notice_type
            if product_code is None and not opts.urn:
                raise TypeError("Missing required property 'product_code'")
            __props__.__dict__["product_code"] = product_code
            if quota_action_code is None and not opts.urn:
                raise TypeError("Missing required property 'quota_action_code'")
            __props__.__dict__["quota_action_code"] = quota_action_code
            __props__.__dict__["quota_category"] = quota_category
        super(TemplateQuota, __self__).__init__(
            'alicloud:quotas/templateQuota:TemplateQuota',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            desire_value: Optional[pulumi.Input[float]] = None,
            dimensions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateQuotaDimensionArgs']]]]] = None,
            effective_time: Optional[pulumi.Input[str]] = None,
            env_language: Optional[pulumi.Input[str]] = None,
            expire_time: Optional[pulumi.Input[str]] = None,
            notice_type: Optional[pulumi.Input[int]] = None,
            product_code: Optional[pulumi.Input[str]] = None,
            quota_action_code: Optional[pulumi.Input[str]] = None,
            quota_category: Optional[pulumi.Input[str]] = None) -> 'TemplateQuota':
        """
        Get an existing TemplateQuota resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] desire_value: Quota application value.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateQuotaDimensionArgs']]]] dimensions: The Quota Dimensions. See the following `Block Dimensions`.
        :param pulumi.Input[str] effective_time: The UTC time when the quota takes effect.
        :param pulumi.Input[str] env_language: The language of the quota alert notification. Value:
               - zh: Chinese.
               - en: English.
        :param pulumi.Input[str] expire_time: The UTC time when the quota expires.
        :param pulumi.Input[int] notice_type: Whether to notify the result of quota promotion application. Value:
               - 0: No.
               - 3: Yes.
        :param pulumi.Input[str] product_code: The abbreviation of the cloud service name.
        :param pulumi.Input[str] quota_action_code: The quota ID.
        :param pulumi.Input[str] quota_category: Type of quota. Value:
               - CommonQuota : Generic quota.
               - WhiteListLabel: Equity quota.
               - FlowControl:API rate quota.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TemplateQuotaState.__new__(_TemplateQuotaState)

        __props__.__dict__["desire_value"] = desire_value
        __props__.__dict__["dimensions"] = dimensions
        __props__.__dict__["effective_time"] = effective_time
        __props__.__dict__["env_language"] = env_language
        __props__.__dict__["expire_time"] = expire_time
        __props__.__dict__["notice_type"] = notice_type
        __props__.__dict__["product_code"] = product_code
        __props__.__dict__["quota_action_code"] = quota_action_code
        __props__.__dict__["quota_category"] = quota_category
        return TemplateQuota(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="desireValue")
    def desire_value(self) -> pulumi.Output[float]:
        """
        Quota application value.
        """
        return pulumi.get(self, "desire_value")

    @property
    @pulumi.getter
    def dimensions(self) -> pulumi.Output[Optional[Sequence['outputs.TemplateQuotaDimension']]]:
        """
        The Quota Dimensions. See the following `Block Dimensions`.
        """
        return pulumi.get(self, "dimensions")

    @property
    @pulumi.getter(name="effectiveTime")
    def effective_time(self) -> pulumi.Output[Optional[str]]:
        """
        The UTC time when the quota takes effect.
        """
        return pulumi.get(self, "effective_time")

    @property
    @pulumi.getter(name="envLanguage")
    def env_language(self) -> pulumi.Output[str]:
        """
        The language of the quota alert notification. Value:
        - zh: Chinese.
        - en: English.
        """
        return pulumi.get(self, "env_language")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[Optional[str]]:
        """
        The UTC time when the quota expires.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="noticeType")
    def notice_type(self) -> pulumi.Output[int]:
        """
        Whether to notify the result of quota promotion application. Value:
        - 0: No.
        - 3: Yes.
        """
        return pulumi.get(self, "notice_type")

    @property
    @pulumi.getter(name="productCode")
    def product_code(self) -> pulumi.Output[str]:
        """
        The abbreviation of the cloud service name.
        """
        return pulumi.get(self, "product_code")

    @property
    @pulumi.getter(name="quotaActionCode")
    def quota_action_code(self) -> pulumi.Output[str]:
        """
        The quota ID.
        """
        return pulumi.get(self, "quota_action_code")

    @property
    @pulumi.getter(name="quotaCategory")
    def quota_category(self) -> pulumi.Output[Optional[str]]:
        """
        Type of quota. Value:
        - CommonQuota : Generic quota.
        - WhiteListLabel: Equity quota.
        - FlowControl:API rate quota.
        """
        return pulumi.get(self, "quota_category")

