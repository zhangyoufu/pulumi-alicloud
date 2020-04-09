# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Project(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the maxcompute project. 
    """
    order_type: pulumi.Output[str]
    """
    The type of payment, only `PayAsYouGo` supported currently.
    """
    specification_type: pulumi.Output[str]
    """
    The type of resource Specification, only `OdpsStandard` supported currently.
    """
    def __init__(__self__, resource_name, opts=None, name=None, order_type=None, specification_type=None, __props__=None, __name__=None, __opts__=None):
        """
        The project is the basic unit of operation in maxcompute. It is similar to the concept of Database or Schema in traditional databases, and sets the boundary for maxcompute multi-user isolation and access control. [Refer to details](https://www.alibabacloud.com/help/doc-detail/27818.html).

        ->**NOTE:** Available in 1.77.0+.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/maxcompute_project.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the maxcompute project. 
        :param pulumi.Input[str] order_type: The type of payment, only `PayAsYouGo` supported currently.
        :param pulumi.Input[str] specification_type: The type of resource Specification, only `OdpsStandard` supported currently.
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

            __props__['name'] = name
            if order_type is None:
                raise TypeError("Missing required property 'order_type'")
            __props__['order_type'] = order_type
            if specification_type is None:
                raise TypeError("Missing required property 'specification_type'")
            __props__['specification_type'] = specification_type
        super(Project, __self__).__init__(
            'alicloud:maxcompute/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, order_type=None, specification_type=None):
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the maxcompute project. 
        :param pulumi.Input[str] order_type: The type of payment, only `PayAsYouGo` supported currently.
        :param pulumi.Input[str] specification_type: The type of resource Specification, only `OdpsStandard` supported currently.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["order_type"] = order_type
        __props__["specification_type"] = specification_type
        return Project(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

