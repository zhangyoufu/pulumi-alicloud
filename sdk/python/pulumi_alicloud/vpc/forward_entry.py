# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ForwardEntry(pulumi.CustomResource):
    external_ip: pulumi.Output[str]
    """
    The external ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
    """
    external_port: pulumi.Output[str]
    """
    The external port, valid value is 1~65535|any.
    """
    forward_entry_id: pulumi.Output[str]
    """
    The id of the forward entry on the server.
    """
    forward_table_id: pulumi.Output[str]
    """
    The value can get from `vpc.NatGateway` Attributes "forward_table_ids".
    """
    internal_ip: pulumi.Output[str]
    """
    The internal ip, must a private ip.
    """
    internal_port: pulumi.Output[str]
    """
    The internal port, valid value is 1~65535|any.
    """
    ip_protocol: pulumi.Output[str]
    """
    The ip protocal, valid value is tcp|udp|any.
    """
    name: pulumi.Output[str]
    """
    The name of forward entry.
    """
    def __init__(__self__, resource_name, opts=None, external_ip=None, external_port=None, forward_table_id=None, internal_ip=None, internal_port=None, ip_protocol=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a forward resource.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/forward_entry.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] external_ip: The external ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        :param pulumi.Input[str] external_port: The external port, valid value is 1~65535|any.
        :param pulumi.Input[str] forward_table_id: The value can get from `vpc.NatGateway` Attributes "forward_table_ids".
        :param pulumi.Input[str] internal_ip: The internal ip, must a private ip.
        :param pulumi.Input[str] internal_port: The internal port, valid value is 1~65535|any.
        :param pulumi.Input[str] ip_protocol: The ip protocal, valid value is tcp|udp|any.
        :param pulumi.Input[str] name: The name of forward entry.
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

            if external_ip is None:
                raise TypeError("Missing required property 'external_ip'")
            __props__['external_ip'] = external_ip
            if external_port is None:
                raise TypeError("Missing required property 'external_port'")
            __props__['external_port'] = external_port
            if forward_table_id is None:
                raise TypeError("Missing required property 'forward_table_id'")
            __props__['forward_table_id'] = forward_table_id
            if internal_ip is None:
                raise TypeError("Missing required property 'internal_ip'")
            __props__['internal_ip'] = internal_ip
            if internal_port is None:
                raise TypeError("Missing required property 'internal_port'")
            __props__['internal_port'] = internal_port
            if ip_protocol is None:
                raise TypeError("Missing required property 'ip_protocol'")
            __props__['ip_protocol'] = ip_protocol
            __props__['name'] = name
            __props__['forward_entry_id'] = None
        super(ForwardEntry, __self__).__init__(
            'alicloud:vpc/forwardEntry:ForwardEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, external_ip=None, external_port=None, forward_entry_id=None, forward_table_id=None, internal_ip=None, internal_port=None, ip_protocol=None, name=None):
        """
        Get an existing ForwardEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] external_ip: The external ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidth_packages`.
        :param pulumi.Input[str] external_port: The external port, valid value is 1~65535|any.
        :param pulumi.Input[str] forward_entry_id: The id of the forward entry on the server.
        :param pulumi.Input[str] forward_table_id: The value can get from `vpc.NatGateway` Attributes "forward_table_ids".
        :param pulumi.Input[str] internal_ip: The internal ip, must a private ip.
        :param pulumi.Input[str] internal_port: The internal port, valid value is 1~65535|any.
        :param pulumi.Input[str] ip_protocol: The ip protocal, valid value is tcp|udp|any.
        :param pulumi.Input[str] name: The name of forward entry.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["external_ip"] = external_ip
        __props__["external_port"] = external_port
        __props__["forward_entry_id"] = forward_entry_id
        __props__["forward_table_id"] = forward_table_id
        __props__["internal_ip"] = internal_ip
        __props__["internal_port"] = internal_port
        __props__["ip_protocol"] = ip_protocol
        __props__["name"] = name
        return ForwardEntry(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

