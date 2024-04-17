# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RocketMQTopicArgs', 'RocketMQTopic']

@pulumi.input_type
class RocketMQTopicArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 topic_name: pulumi.Input[str],
                 message_type: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RocketMQTopic resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] topic_name: Topic name and identification.
        :param pulumi.Input[str] message_type: Message type.
        :param pulumi.Input[str] remark: Custom remarks.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "topic_name", topic_name)
        if message_type is not None:
            pulumi.set(__self__, "message_type", message_type)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Input[str]:
        """
        Topic name and identification.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> Optional[pulumi.Input[str]]:
        """
        Message type.
        """
        return pulumi.get(self, "message_type")

    @message_type.setter
    def message_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message_type", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Custom remarks.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)


@pulumi.input_type
class _RocketMQTopicState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 message_type: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RocketMQTopic resources.
        :param pulumi.Input[str] create_time: The creation time of the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] message_type: Message type.
        :param pulumi.Input[str] remark: Custom remarks.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[str] topic_name: Topic name and identification.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if message_type is not None:
            pulumi.set(__self__, "message_type", message_type)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if topic_name is not None:
            pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the resource.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> Optional[pulumi.Input[str]]:
        """
        Message type.
        """
        return pulumi.get(self, "message_type")

    @message_type.setter
    def message_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message_type", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Custom remarks.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        Topic name and identification.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_name", value)


class RocketMQTopic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 message_type: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a RocketMQ Topic resource.

        For information about RocketMQ Topic and how to use it, see [What is Topic](https://www.alibabacloud.com/help/en/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/developer-reference/api-rocketmq-2022-08-01-createtopic).

        > **NOTE:** Available since v1.211.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.get_zones(available_resource_creation="VSwitch")
        create_vpc = alicloud.vpc.Network("createVpc",
            description="example",
            cidr_block="172.16.0.0/12",
            vpc_name=name)
        create_vswitch = alicloud.vpc.Switch("createVswitch",
            description="example",
            vpc_id=create_vpc.id,
            zone_id=default.zones[0].id,
            cidr_block="172.16.0.0/24",
            vswitch_name=name)
        create_instance = alicloud.rocketmq.RocketMQInstance("createInstance",
            auto_renew_period=1,
            product_info=alicloud.rocketmq.RocketMQInstanceProductInfoArgs(
                msg_process_spec="rmq.p2.4xlarge",
                send_receive_ratio=0.3,
                message_retention_time=70,
            ),
            network_info=alicloud.rocketmq.RocketMQInstanceNetworkInfoArgs(
                vpc_info=alicloud.rocketmq.RocketMQInstanceNetworkInfoVpcInfoArgs(
                    vpc_id=create_vpc.id,
                    vswitch_id=create_vswitch.id,
                ),
                internet_info=alicloud.rocketmq.RocketMQInstanceNetworkInfoInternetInfoArgs(
                    internet_spec="enable",
                    flow_out_type="payByBandwidth",
                    flow_out_bandwidth=30,
                ),
            ),
            period=1,
            sub_series_code="cluster_ha",
            remark="example",
            instance_name=name,
            service_code="rmq",
            series_code="professional",
            payment_type="PayAsYouGo",
            period_unit="Month")
        default_rocket_mq_topic = alicloud.rocketmq.RocketMQTopic("default",
            remark="example",
            instance_id=create_instance.id,
            message_type="NORMAL",
            topic_name=name)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        RocketMQ Topic can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:rocketmq/rocketMQTopic:RocketMQTopic example <instance_id>:<topic_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] message_type: Message type.
        :param pulumi.Input[str] remark: Custom remarks.
        :param pulumi.Input[str] topic_name: Topic name and identification.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RocketMQTopicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a RocketMQ Topic resource.

        For information about RocketMQ Topic and how to use it, see [What is Topic](https://www.alibabacloud.com/help/en/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/developer-reference/api-rocketmq-2022-08-01-createtopic).

        > **NOTE:** Available since v1.211.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.get_zones(available_resource_creation="VSwitch")
        create_vpc = alicloud.vpc.Network("createVpc",
            description="example",
            cidr_block="172.16.0.0/12",
            vpc_name=name)
        create_vswitch = alicloud.vpc.Switch("createVswitch",
            description="example",
            vpc_id=create_vpc.id,
            zone_id=default.zones[0].id,
            cidr_block="172.16.0.0/24",
            vswitch_name=name)
        create_instance = alicloud.rocketmq.RocketMQInstance("createInstance",
            auto_renew_period=1,
            product_info=alicloud.rocketmq.RocketMQInstanceProductInfoArgs(
                msg_process_spec="rmq.p2.4xlarge",
                send_receive_ratio=0.3,
                message_retention_time=70,
            ),
            network_info=alicloud.rocketmq.RocketMQInstanceNetworkInfoArgs(
                vpc_info=alicloud.rocketmq.RocketMQInstanceNetworkInfoVpcInfoArgs(
                    vpc_id=create_vpc.id,
                    vswitch_id=create_vswitch.id,
                ),
                internet_info=alicloud.rocketmq.RocketMQInstanceNetworkInfoInternetInfoArgs(
                    internet_spec="enable",
                    flow_out_type="payByBandwidth",
                    flow_out_bandwidth=30,
                ),
            ),
            period=1,
            sub_series_code="cluster_ha",
            remark="example",
            instance_name=name,
            service_code="rmq",
            series_code="professional",
            payment_type="PayAsYouGo",
            period_unit="Month")
        default_rocket_mq_topic = alicloud.rocketmq.RocketMQTopic("default",
            remark="example",
            instance_id=create_instance.id,
            message_type="NORMAL",
            topic_name=name)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        RocketMQ Topic can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:rocketmq/rocketMQTopic:RocketMQTopic example <instance_id>:<topic_name>
        ```

        :param str resource_name: The name of the resource.
        :param RocketMQTopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RocketMQTopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 message_type: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RocketMQTopicArgs.__new__(RocketMQTopicArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["message_type"] = message_type
            __props__.__dict__["remark"] = remark
            if topic_name is None and not opts.urn:
                raise TypeError("Missing required property 'topic_name'")
            __props__.__dict__["topic_name"] = topic_name
            __props__.__dict__["create_time"] = None
            __props__.__dict__["status"] = None
        super(RocketMQTopic, __self__).__init__(
            'alicloud:rocketmq/rocketMQTopic:RocketMQTopic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            message_type: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            topic_name: Optional[pulumi.Input[str]] = None) -> 'RocketMQTopic':
        """
        Get an existing RocketMQTopic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The creation time of the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] message_type: Message type.
        :param pulumi.Input[str] remark: Custom remarks.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[str] topic_name: Topic name and identification.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RocketMQTopicState.__new__(_RocketMQTopicState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["message_type"] = message_type
        __props__.__dict__["remark"] = remark
        __props__.__dict__["status"] = status
        __props__.__dict__["topic_name"] = topic_name
        return RocketMQTopic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of the resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> pulumi.Output[Optional[str]]:
        """
        Message type.
        """
        return pulumi.get(self, "message_type")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Custom remarks.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Output[str]:
        """
        Topic name and identification.
        """
        return pulumi.get(self, "topic_name")

