# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SnatEntry(pulumi.CustomResource):
    cidr_block: pulumi.Output[str]
    """
    The destination CIDR block.
    """
    sag_id: pulumi.Output[str]
    """
    The ID of the SAG instance.
    """
    snat_ip: pulumi.Output[str]
    """
    The public IP address.
    """
    def __init__(__self__, resource_name, opts=None, cidr_block=None, sag_id=None, snat_ip=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Sag SnatEntry resource. This topic describes how to add a SNAT entry to enable the SNAT function. The SNAT function can hide internal IP addresses and resolve private IP address conflicts. With this function, on-premises sites can access internal IP addresses, but cannot be accessed by internal IP addresses. If you do not add a SNAT entry, on-premises sites can access each other only when all related IP addresses do not conflict.

        For information about Sag SnatEntry and how to use it, see [What is Sag SnatEntry](https://www.alibabacloud.com/help/doc-detail/124231.htm).

        > **NOTE:** Available in 1.61.0+

        > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]



        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/sag_snat_entry.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_block: The destination CIDR block.
        :param pulumi.Input[str] sag_id: The ID of the SAG instance.
        :param pulumi.Input[str] snat_ip: The public IP address.
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

            if cidr_block is None:
                raise TypeError("Missing required property 'cidr_block'")
            __props__['cidr_block'] = cidr_block
            if sag_id is None:
                raise TypeError("Missing required property 'sag_id'")
            __props__['sag_id'] = sag_id
            if snat_ip is None:
                raise TypeError("Missing required property 'snat_ip'")
            __props__['snat_ip'] = snat_ip
        super(SnatEntry, __self__).__init__(
            'alicloud:rocketmq/snatEntry:SnatEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cidr_block=None, sag_id=None, snat_ip=None):
        """
        Get an existing SnatEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_block: The destination CIDR block.
        :param pulumi.Input[str] sag_id: The ID of the SAG instance.
        :param pulumi.Input[str] snat_ip: The public IP address.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cidr_block"] = cidr_block
        __props__["sag_id"] = sag_id
        __props__["snat_ip"] = snat_ip
        return SnatEntry(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

