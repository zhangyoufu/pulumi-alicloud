# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Alias(pulumi.CustomResource):
    alias_name: pulumi.Output[str]
    """
    The alias of CMK. `Encrypt`、`GenerateDataKey`、`DescribeKey` can be called using aliases. Length of characters other than prefixes: minimum length of 1 character and maximum length of 255 characters. Must contain prefix `alias/`.
    """
    key_id: pulumi.Output[str]
    """
    The id of the key.
    """
    def __init__(__self__, resource_name, opts=None, alias_name=None, key_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Create an alias for the master key (CMK).

        > **NOTE:** Available in v1.77.0+.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias_name: The alias of CMK. `Encrypt`、`GenerateDataKey`、`DescribeKey` can be called using aliases. Length of characters other than prefixes: minimum length of 1 character and maximum length of 255 characters. Must contain prefix `alias/`.
        :param pulumi.Input[str] key_id: The id of the key.
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

            if alias_name is None:
                raise TypeError("Missing required property 'alias_name'")
            __props__['alias_name'] = alias_name
            if key_id is None:
                raise TypeError("Missing required property 'key_id'")
            __props__['key_id'] = key_id
        super(Alias, __self__).__init__(
            'alicloud:kms/alias:Alias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, alias_name=None, key_id=None):
        """
        Get an existing Alias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias_name: The alias of CMK. `Encrypt`、`GenerateDataKey`、`DescribeKey` can be called using aliases. Length of characters other than prefixes: minimum length of 1 character and maximum length of 255 characters. Must contain prefix `alias/`.
        :param pulumi.Input[str] key_id: The id of the key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alias_name"] = alias_name
        __props__["key_id"] = key_id
        return Alias(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

