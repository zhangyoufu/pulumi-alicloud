# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    instance_status: pulumi.Output[float]
    """
    The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
    """
    instance_type: pulumi.Output[float]
    """
    The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.
    """
    name: pulumi.Output[str]
    """
    Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
    """
    release_time: pulumi.Output[str]
    """
    Platinum edition instance expiration time.
    """
    remark: pulumi.Output[str]
    """
    This attribute is a concise description of instance. The length cannot exceed 128.
    """
    def __init__(__self__, resource_name, opts=None, name=None, remark=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an ONS instance resource.

        For more information about how to use it, see [RocketMQ Instance Management API](https://www.alibabacloud.com/help/doc-detail/106354.html). 

        > **NOTE:** The number of instances in the same region cannot exceed 8. At present, the resource does not support region "mq-internet-access" and "ap-southeast-5". 

        > **NOTE:** Available in 1.51.0+

        ## Example Usage



        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.rocketmq.Instance("example", remark="tf-example-ons-instance-remark")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
        :param pulumi.Input[str] remark: This attribute is a concise description of instance. The length cannot exceed 128.
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
            __props__['remark'] = remark
            __props__['instance_status'] = None
            __props__['instance_type'] = None
            __props__['release_time'] = None
        super(Instance, __self__).__init__(
            'alicloud:rocketmq/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, instance_status=None, instance_type=None, name=None, release_time=None, remark=None):
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] instance_status: The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
        :param pulumi.Input[float] instance_type: The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.
        :param pulumi.Input[str] name: Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
        :param pulumi.Input[str] release_time: Platinum edition instance expiration time.
        :param pulumi.Input[str] remark: This attribute is a concise description of instance. The length cannot exceed 128.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["instance_status"] = instance_status
        __props__["instance_type"] = instance_type
        __props__["name"] = name
        __props__["release_time"] = release_time
        __props__["remark"] = remark
        return Instance(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

