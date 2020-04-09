# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class CaCertificate(pulumi.CustomResource):
    ca_certificate: pulumi.Output[str]
    """
    the content of the CA certificate.
    """
    name: pulumi.Output[str]
    """
    Name of the CA Certificate.
    """
    resource_group_id: pulumi.Output[str]
    """
    The Id of resource group which the slb_ca certificate belongs.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, ca_certificate=None, name=None, resource_group_id=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        A Load Balancer CA Certificate is used by the listener of the protocol https.

        For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).

        For information about CA Certificate and how to use it, see [Configure CA Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).




        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/slb_ca_certificate.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ca_certificate: the content of the CA certificate.
        :param pulumi.Input[str] name: Name of the CA Certificate.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the slb_ca certificate belongs.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
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

            if ca_certificate is None:
                raise TypeError("Missing required property 'ca_certificate'")
            __props__['ca_certificate'] = ca_certificate
            __props__['name'] = name
            __props__['resource_group_id'] = resource_group_id
            __props__['tags'] = tags
        super(CaCertificate, __self__).__init__(
            'alicloud:slb/caCertificate:CaCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, ca_certificate=None, name=None, resource_group_id=None, tags=None):
        """
        Get an existing CaCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ca_certificate: the content of the CA certificate.
        :param pulumi.Input[str] name: Name of the CA Certificate.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the slb_ca certificate belongs.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ca_certificate"] = ca_certificate
        __props__["name"] = name
        __props__["resource_group_id"] = resource_group_id
        __props__["tags"] = tags
        return CaCertificate(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

