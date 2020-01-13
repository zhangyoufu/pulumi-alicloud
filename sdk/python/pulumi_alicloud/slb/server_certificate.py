# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ServerCertificate(pulumi.CustomResource):
    alicloud_certifacte_id: pulumi.Output[str]
    alicloud_certifacte_name: pulumi.Output[str]
    alicloud_certificate_id: pulumi.Output[str]
    """
    an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
    """
    alicloud_certificate_name: pulumi.Output[str]
    """
    the name of the certificate specified by `alicloud_certificate_id`.but it is not supported on the international site of alibaba cloud now.
    """
    alicloud_certificate_region_id: pulumi.Output[str]
    """
    the region of the certificate specified by `alicloud_certificate_id`. but it is not supported on the international site of alibaba cloud now.
    """
    name: pulumi.Output[str]
    """
    Name of the Server Certificate.
    """
    private_key: pulumi.Output[str]
    """
    the content of privat key of the ssl certificate specified by `server_certificate`. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
    """
    resource_group_id: pulumi.Output[str]
    """
    The Id of resource group which the slb server certificate belongs.
    """
    server_certificate: pulumi.Output[str]
    """
    the content of the ssl certificate. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, alicloud_certifacte_id=None, alicloud_certifacte_name=None, alicloud_certificate_id=None, alicloud_certificate_name=None, alicloud_certificate_region_id=None, name=None, private_key=None, resource_group_id=None, server_certificate=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        A Load Balancer Server Certificate is an ssl Certificate used by the listener of the protocol https.
        
        For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).
        
        For information about Server Certificate and how to use it, see [Configure Server Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alicloud_certificate_id: an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
        :param pulumi.Input[str] alicloud_certificate_name: the name of the certificate specified by `alicloud_certificate_id`.but it is not supported on the international site of alibaba cloud now.
        :param pulumi.Input[str] alicloud_certificate_region_id: the region of the certificate specified by `alicloud_certificate_id`. but it is not supported on the international site of alibaba cloud now.
        :param pulumi.Input[str] name: Name of the Server Certificate.
        :param pulumi.Input[str] private_key: the content of privat key of the ssl certificate specified by `server_certificate`. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the slb server certificate belongs.
        :param pulumi.Input[str] server_certificate: the content of the ssl certificate. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/slb_server_certificate.html.markdown.
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

            __props__['alicloud_certifacte_id'] = alicloud_certifacte_id
            __props__['alicloud_certifacte_name'] = alicloud_certifacte_name
            __props__['alicloud_certificate_id'] = alicloud_certificate_id
            __props__['alicloud_certificate_name'] = alicloud_certificate_name
            __props__['alicloud_certificate_region_id'] = alicloud_certificate_region_id
            __props__['name'] = name
            __props__['private_key'] = private_key
            __props__['resource_group_id'] = resource_group_id
            __props__['server_certificate'] = server_certificate
            __props__['tags'] = tags
        super(ServerCertificate, __self__).__init__(
            'alicloud:slb/serverCertificate:ServerCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, alicloud_certifacte_id=None, alicloud_certifacte_name=None, alicloud_certificate_id=None, alicloud_certificate_name=None, alicloud_certificate_region_id=None, name=None, private_key=None, resource_group_id=None, server_certificate=None, tags=None):
        """
        Get an existing ServerCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alicloud_certificate_id: an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
        :param pulumi.Input[str] alicloud_certificate_name: the name of the certificate specified by `alicloud_certificate_id`.but it is not supported on the international site of alibaba cloud now.
        :param pulumi.Input[str] alicloud_certificate_region_id: the region of the certificate specified by `alicloud_certificate_id`. but it is not supported on the international site of alibaba cloud now.
        :param pulumi.Input[str] name: Name of the Server Certificate.
        :param pulumi.Input[str] private_key: the content of privat key of the ssl certificate specified by `server_certificate`. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the slb server certificate belongs.
        :param pulumi.Input[str] server_certificate: the content of the ssl certificate. where `alicloud_certificate_id` is null, it is required, otherwise it is ignored.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/slb_server_certificate.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["alicloud_certifacte_id"] = alicloud_certifacte_id
        __props__["alicloud_certifacte_name"] = alicloud_certifacte_name
        __props__["alicloud_certificate_id"] = alicloud_certificate_id
        __props__["alicloud_certificate_name"] = alicloud_certificate_name
        __props__["alicloud_certificate_region_id"] = alicloud_certificate_region_id
        __props__["name"] = name
        __props__["private_key"] = private_key
        __props__["resource_group_id"] = resource_group_id
        __props__["server_certificate"] = server_certificate
        __props__["tags"] = tags
        return ServerCertificate(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

