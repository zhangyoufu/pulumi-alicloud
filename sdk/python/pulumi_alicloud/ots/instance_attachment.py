# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class InstanceAttachment(pulumi.CustomResource):
    instance_name: pulumi.Output[str]
    """
    The name of the OTS instance.
    """
    vpc_id: pulumi.Output[str]
    """
    The ID of attaching VPC to instance.
    """
    vpc_name: pulumi.Output[str]
    """
    The name of attaching VPC to instance.
    """
    vswitch_id: pulumi.Output[str]
    """
    The ID of attaching VSwitch to instance.
    """
    def __init__(__self__, resource_name, opts=None, instance_name=None, vpc_name=None, vswitch_id=None, __props__=None, __name__=None, __opts__=None):
        """
        This resource will help you to bind a VPC to an OTS instance.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Create an OTS instance
        foo_instance = alicloud.ots.Instance("fooInstance",
            accessed_by="Vpc",
            description="for table",
            tags={
                "Created": "TF",
                "For": "Building table",
            })
        foo_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        foo_network = alicloud.vpc.Network("fooNetwork", cidr_block="172.16.0.0/16")
        foo_switch = alicloud.vpc.Switch("fooSwitch",
            availability_zone=foo_zones.zones[0]["id"],
            cidr_block="172.16.1.0/24",
            vpc_id=foo_network.id)
        foo_instance_attachment = alicloud.ots.InstanceAttachment("fooInstanceAttachment",
            instance_name=foo_instance.name,
            vpc_name="attachment1",
            vswitch_id=foo_switch.id)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_name: The name of the OTS instance.
        :param pulumi.Input[str] vpc_name: The name of attaching VPC to instance.
        :param pulumi.Input[str] vswitch_id: The ID of attaching VSwitch to instance.
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

            if instance_name is None:
                raise TypeError("Missing required property 'instance_name'")
            __props__['instance_name'] = instance_name
            if vpc_name is None:
                raise TypeError("Missing required property 'vpc_name'")
            __props__['vpc_name'] = vpc_name
            if vswitch_id is None:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__['vswitch_id'] = vswitch_id
            __props__['vpc_id'] = None
        super(InstanceAttachment, __self__).__init__(
            'alicloud:ots/instanceAttachment:InstanceAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, instance_name=None, vpc_id=None, vpc_name=None, vswitch_id=None):
        """
        Get an existing InstanceAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_name: The name of the OTS instance.
        :param pulumi.Input[str] vpc_id: The ID of attaching VPC to instance.
        :param pulumi.Input[str] vpc_name: The name of attaching VPC to instance.
        :param pulumi.Input[str] vswitch_id: The ID of attaching VSwitch to instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["instance_name"] = instance_name
        __props__["vpc_id"] = vpc_id
        __props__["vpc_name"] = vpc_name
        __props__["vswitch_id"] = vswitch_id
        return InstanceAttachment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

