# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class BastionHostInstance(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    Description of the instance. This name can have a string of 1 to 63 characters.
    """
    license_code: pulumi.Output[str]
    period: pulumi.Output[float]
    """
    Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
    """
    security_group_ids: pulumi.Output[list]
    """
    security group IDs configured to bastionhost
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vswitch_id: pulumi.Output[str]
    """
    vSwtich ID configured to bastionhost
    """
    def __init__(__self__, resource_name, opts=None, description=None, license_code=None, period=None, security_group_ids=None, tags=None, vswitch_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Cloud Bastionhost instance resource ("Yundun_bastionhost" is the short term of this product).

        > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.

        > **NOTE:** Available in 1.63.0+ .

        > **NOTE:** In order to destroy Cloud Bastionhost instance , users are required to apply for white list first



        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/yundun_bastionhost_instance.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[float] period: Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
        :param pulumi.Input[list] security_group_ids: security group IDs configured to bastionhost
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: vSwtich ID configured to bastionhost
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if description is None:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            if license_code is None:
                raise TypeError("Missing required property 'license_code'")
            __props__['license_code'] = license_code
            __props__['period'] = period
            if security_group_ids is None:
                raise TypeError("Missing required property 'security_group_ids'")
            __props__['security_group_ids'] = security_group_ids
            __props__['tags'] = tags
            if vswitch_id is None:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__['vswitch_id'] = vswitch_id
        super(BastionHostInstance, __self__).__init__(
            'alicloud:yundun/bastionHostInstance:BastionHostInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, license_code=None, period=None, security_group_ids=None, tags=None, vswitch_id=None):
        """
        Get an existing BastionHostInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[float] period: Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
        :param pulumi.Input[list] security_group_ids: security group IDs configured to bastionhost
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: vSwtich ID configured to bastionhost
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["license_code"] = license_code
        __props__["period"] = period
        __props__["security_group_ids"] = security_group_ids
        __props__["tags"] = tags
        __props__["vswitch_id"] = vswitch_id
        return BastionHostInstance(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

