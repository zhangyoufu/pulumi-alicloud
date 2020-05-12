# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Topic(pulumi.CustomResource):
    instance_id: pulumi.Output[str]
    """
    ID of the ONS Instance that owns the topics.
    """
    message_type: pulumi.Output[float]
    """
    The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
    """
    perm: pulumi.Output[float]
    """
    This attribute is used to set the read-write mode for the topic. Read [Request parameters](https://www.alibabacloud.com/help/doc-detail/56880.html) for further details.
    """
    remark: pulumi.Output[str]
    """
    This attribute is a concise description of topic. The length cannot exceed 128.
    """
    topic: pulumi.Output[str]
    """
    Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
    """
    def __init__(__self__, resource_name, opts=None, instance_id=None, message_type=None, perm=None, remark=None, topic=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an ONS topic resource.

        For more information about how to use it, see [RocketMQ Topic Management API](https://www.alibabacloud.com/help/doc-detail/29591.html). 

        > **NOTE:** Available in 1.53.0+

        ## Example Usage



        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "onsInstanceName"
        topic = config.get("topic")
        if topic is None:
            topic = "onsTopicName"
        default_instance = alicloud.rocketmq.Instance("defaultInstance", remark="default_ons_instance_remark")
        default_topic = alicloud.rocketmq.Topic("defaultTopic",
            instance_id=default_instance.id,
            message_type=0,
            remark="dafault_ons_topic_remark",
            topic=topic)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: ID of the ONS Instance that owns the topics.
        :param pulumi.Input[float] message_type: The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
        :param pulumi.Input[float] perm: This attribute is used to set the read-write mode for the topic. Read [Request parameters](https://www.alibabacloud.com/help/doc-detail/56880.html) for further details.
        :param pulumi.Input[str] remark: This attribute is a concise description of topic. The length cannot exceed 128.
        :param pulumi.Input[str] topic: Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
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

            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            if message_type is None:
                raise TypeError("Missing required property 'message_type'")
            __props__['message_type'] = message_type
            __props__['perm'] = perm
            __props__['remark'] = remark
            if topic is None:
                raise TypeError("Missing required property 'topic'")
            __props__['topic'] = topic
        super(Topic, __self__).__init__(
            'alicloud:rocketmq/topic:Topic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, instance_id=None, message_type=None, perm=None, remark=None, topic=None):
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: ID of the ONS Instance that owns the topics.
        :param pulumi.Input[float] message_type: The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
        :param pulumi.Input[float] perm: This attribute is used to set the read-write mode for the topic. Read [Request parameters](https://www.alibabacloud.com/help/doc-detail/56880.html) for further details.
        :param pulumi.Input[str] remark: This attribute is a concise description of topic. The length cannot exceed 128.
        :param pulumi.Input[str] topic: Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["instance_id"] = instance_id
        __props__["message_type"] = message_type
        __props__["perm"] = perm
        __props__["remark"] = remark
        __props__["topic"] = topic
        return Topic(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

