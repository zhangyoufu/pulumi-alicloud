# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Zone(pulumi.CustomResource):
    creation_time: pulumi.Output[str]
    is_ptr: pulumi.Output[bool]
    lang: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    The name of the Private Zone.
    """
    proxy_pattern: pulumi.Output[str]
    record_count: pulumi.Output[float]
    """
    The count of the Private Zone Record.
    """
    remark: pulumi.Output[str]
    """
    The remark of the Private Zone.
    * `proxy_pattern - (Optional, Available in 1.69.0+) The recursive DNS proxy. Valid values:
    - ZONE: indicates that the recursive DNS proxy is disabled.
    - RECORD: indicates that the recursive DNS proxy is enabled.
    """
    update_time: pulumi.Output[str]
    user_client_ip: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, lang=None, name=None, proxy_pattern=None, remark=None, user_client_ip=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Zone resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Private Zone.
        :param pulumi.Input[str] remark: The remark of the Private Zone.
               * `proxy_pattern - (Optional, Available in 1.69.0+) The recursive DNS proxy. Valid values:
               - ZONE: indicates that the recursive DNS proxy is disabled.
               - RECORD: indicates that the recursive DNS proxy is enabled.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/pvtz_zone.html.markdown.
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

            __props__['lang'] = lang
            __props__['name'] = name
            __props__['proxy_pattern'] = proxy_pattern
            __props__['remark'] = remark
            __props__['user_client_ip'] = user_client_ip
            __props__['creation_time'] = None
            __props__['is_ptr'] = None
            __props__['record_count'] = None
            __props__['update_time'] = None
        super(Zone, __self__).__init__(
            'alicloud:pvtz/zone:Zone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, creation_time=None, is_ptr=None, lang=None, name=None, proxy_pattern=None, record_count=None, remark=None, update_time=None, user_client_ip=None):
        """
        Get an existing Zone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Private Zone.
        :param pulumi.Input[float] record_count: The count of the Private Zone Record.
        :param pulumi.Input[str] remark: The remark of the Private Zone.
               * `proxy_pattern - (Optional, Available in 1.69.0+) The recursive DNS proxy. Valid values:
               - ZONE: indicates that the recursive DNS proxy is disabled.
               - RECORD: indicates that the recursive DNS proxy is enabled.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/pvtz_zone.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["creation_time"] = creation_time
        __props__["is_ptr"] = is_ptr
        __props__["lang"] = lang
        __props__["name"] = name
        __props__["proxy_pattern"] = proxy_pattern
        __props__["record_count"] = record_count
        __props__["remark"] = remark
        __props__["update_time"] = update_time
        __props__["user_client_ip"] = user_client_ip
        return Zone(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

