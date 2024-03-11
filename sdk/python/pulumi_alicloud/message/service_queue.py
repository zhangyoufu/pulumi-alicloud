# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServiceQueueArgs', 'ServiceQueue']

@pulumi.input_type
class ServiceQueueArgs:
    def __init__(__self__, *,
                 queue_name: pulumi.Input[str],
                 delay_seconds: Optional[pulumi.Input[int]] = None,
                 logging_enabled: Optional[pulumi.Input[bool]] = None,
                 maximum_message_size: Optional[pulumi.Input[int]] = None,
                 message_retention_period: Optional[pulumi.Input[int]] = None,
                 polling_wait_seconds: Optional[pulumi.Input[int]] = None,
                 visibility_timeout: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ServiceQueue resource.
        :param pulumi.Input[str] queue_name: Two queues on a single account in the same region cannot have the same name. A queue name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 120 characters.
        :param pulumi.Input[int] delay_seconds: The delay period after which a message sent to the queue can be consumed. Unit: seconds. Valid values: 0-604800 seconds. Default value: 0.
        :param pulumi.Input[bool] logging_enabled: Specifies whether to enable the log management feature. Default value: false. Valid values:
        :param pulumi.Input[int] maximum_message_size: The maximum size of a message body that can be sent to the queue. Unit: bytes. Valid value range: 1024-65536. Default value: 65536.
        :param pulumi.Input[int] message_retention_period: The maximum period for which a message can be retained in the queue. After the specified period, the message is deleted no matter whether the message is consumed. Unit: seconds. Valid values: 60-604800. Default value: 345600.
        :param pulumi.Input[int] polling_wait_seconds: The maximum period for which a ReceiveMessage request waits if no message is available in the queue. Unit: seconds. Valid values: 0-30. Default value: 0.
        :param pulumi.Input[int] visibility_timeout: The invisibility period for which the received message remains the Inactive state. Unit: seconds. Valid values: 1-43200. Default value: 30.
        """
        pulumi.set(__self__, "queue_name", queue_name)
        if delay_seconds is not None:
            pulumi.set(__self__, "delay_seconds", delay_seconds)
        if logging_enabled is not None:
            pulumi.set(__self__, "logging_enabled", logging_enabled)
        if maximum_message_size is not None:
            pulumi.set(__self__, "maximum_message_size", maximum_message_size)
        if message_retention_period is not None:
            pulumi.set(__self__, "message_retention_period", message_retention_period)
        if polling_wait_seconds is not None:
            pulumi.set(__self__, "polling_wait_seconds", polling_wait_seconds)
        if visibility_timeout is not None:
            pulumi.set(__self__, "visibility_timeout", visibility_timeout)

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> pulumi.Input[str]:
        """
        Two queues on a single account in the same region cannot have the same name. A queue name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 120 characters.
        """
        return pulumi.get(self, "queue_name")

    @queue_name.setter
    def queue_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "queue_name", value)

    @property
    @pulumi.getter(name="delaySeconds")
    def delay_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The delay period after which a message sent to the queue can be consumed. Unit: seconds. Valid values: 0-604800 seconds. Default value: 0.
        """
        return pulumi.get(self, "delay_seconds")

    @delay_seconds.setter
    def delay_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delay_seconds", value)

    @property
    @pulumi.getter(name="loggingEnabled")
    def logging_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the log management feature. Default value: false. Valid values:
        """
        return pulumi.get(self, "logging_enabled")

    @logging_enabled.setter
    def logging_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "logging_enabled", value)

    @property
    @pulumi.getter(name="maximumMessageSize")
    def maximum_message_size(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum size of a message body that can be sent to the queue. Unit: bytes. Valid value range: 1024-65536. Default value: 65536.
        """
        return pulumi.get(self, "maximum_message_size")

    @maximum_message_size.setter
    def maximum_message_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_message_size", value)

    @property
    @pulumi.getter(name="messageRetentionPeriod")
    def message_retention_period(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum period for which a message can be retained in the queue. After the specified period, the message is deleted no matter whether the message is consumed. Unit: seconds. Valid values: 60-604800. Default value: 345600.
        """
        return pulumi.get(self, "message_retention_period")

    @message_retention_period.setter
    def message_retention_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_retention_period", value)

    @property
    @pulumi.getter(name="pollingWaitSeconds")
    def polling_wait_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum period for which a ReceiveMessage request waits if no message is available in the queue. Unit: seconds. Valid values: 0-30. Default value: 0.
        """
        return pulumi.get(self, "polling_wait_seconds")

    @polling_wait_seconds.setter
    def polling_wait_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "polling_wait_seconds", value)

    @property
    @pulumi.getter(name="visibilityTimeout")
    def visibility_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The invisibility period for which the received message remains the Inactive state. Unit: seconds. Valid values: 1-43200. Default value: 30.
        """
        return pulumi.get(self, "visibility_timeout")

    @visibility_timeout.setter
    def visibility_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "visibility_timeout", value)


@pulumi.input_type
class _ServiceQueueState:
    def __init__(__self__, *,
                 delay_seconds: Optional[pulumi.Input[int]] = None,
                 logging_enabled: Optional[pulumi.Input[bool]] = None,
                 maximum_message_size: Optional[pulumi.Input[int]] = None,
                 message_retention_period: Optional[pulumi.Input[int]] = None,
                 polling_wait_seconds: Optional[pulumi.Input[int]] = None,
                 queue_name: Optional[pulumi.Input[str]] = None,
                 visibility_timeout: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ServiceQueue resources.
        :param pulumi.Input[int] delay_seconds: The delay period after which a message sent to the queue can be consumed. Unit: seconds. Valid values: 0-604800 seconds. Default value: 0.
        :param pulumi.Input[bool] logging_enabled: Specifies whether to enable the log management feature. Default value: false. Valid values:
        :param pulumi.Input[int] maximum_message_size: The maximum size of a message body that can be sent to the queue. Unit: bytes. Valid value range: 1024-65536. Default value: 65536.
        :param pulumi.Input[int] message_retention_period: The maximum period for which a message can be retained in the queue. After the specified period, the message is deleted no matter whether the message is consumed. Unit: seconds. Valid values: 60-604800. Default value: 345600.
        :param pulumi.Input[int] polling_wait_seconds: The maximum period for which a ReceiveMessage request waits if no message is available in the queue. Unit: seconds. Valid values: 0-30. Default value: 0.
        :param pulumi.Input[str] queue_name: Two queues on a single account in the same region cannot have the same name. A queue name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 120 characters.
        :param pulumi.Input[int] visibility_timeout: The invisibility period for which the received message remains the Inactive state. Unit: seconds. Valid values: 1-43200. Default value: 30.
        """
        if delay_seconds is not None:
            pulumi.set(__self__, "delay_seconds", delay_seconds)
        if logging_enabled is not None:
            pulumi.set(__self__, "logging_enabled", logging_enabled)
        if maximum_message_size is not None:
            pulumi.set(__self__, "maximum_message_size", maximum_message_size)
        if message_retention_period is not None:
            pulumi.set(__self__, "message_retention_period", message_retention_period)
        if polling_wait_seconds is not None:
            pulumi.set(__self__, "polling_wait_seconds", polling_wait_seconds)
        if queue_name is not None:
            pulumi.set(__self__, "queue_name", queue_name)
        if visibility_timeout is not None:
            pulumi.set(__self__, "visibility_timeout", visibility_timeout)

    @property
    @pulumi.getter(name="delaySeconds")
    def delay_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The delay period after which a message sent to the queue can be consumed. Unit: seconds. Valid values: 0-604800 seconds. Default value: 0.
        """
        return pulumi.get(self, "delay_seconds")

    @delay_seconds.setter
    def delay_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delay_seconds", value)

    @property
    @pulumi.getter(name="loggingEnabled")
    def logging_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the log management feature. Default value: false. Valid values:
        """
        return pulumi.get(self, "logging_enabled")

    @logging_enabled.setter
    def logging_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "logging_enabled", value)

    @property
    @pulumi.getter(name="maximumMessageSize")
    def maximum_message_size(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum size of a message body that can be sent to the queue. Unit: bytes. Valid value range: 1024-65536. Default value: 65536.
        """
        return pulumi.get(self, "maximum_message_size")

    @maximum_message_size.setter
    def maximum_message_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_message_size", value)

    @property
    @pulumi.getter(name="messageRetentionPeriod")
    def message_retention_period(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum period for which a message can be retained in the queue. After the specified period, the message is deleted no matter whether the message is consumed. Unit: seconds. Valid values: 60-604800. Default value: 345600.
        """
        return pulumi.get(self, "message_retention_period")

    @message_retention_period.setter
    def message_retention_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_retention_period", value)

    @property
    @pulumi.getter(name="pollingWaitSeconds")
    def polling_wait_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum period for which a ReceiveMessage request waits if no message is available in the queue. Unit: seconds. Valid values: 0-30. Default value: 0.
        """
        return pulumi.get(self, "polling_wait_seconds")

    @polling_wait_seconds.setter
    def polling_wait_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "polling_wait_seconds", value)

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> Optional[pulumi.Input[str]]:
        """
        Two queues on a single account in the same region cannot have the same name. A queue name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 120 characters.
        """
        return pulumi.get(self, "queue_name")

    @queue_name.setter
    def queue_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue_name", value)

    @property
    @pulumi.getter(name="visibilityTimeout")
    def visibility_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The invisibility period for which the received message remains the Inactive state. Unit: seconds. Valid values: 1-43200. Default value: 30.
        """
        return pulumi.get(self, "visibility_timeout")

    @visibility_timeout.setter
    def visibility_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "visibility_timeout", value)


class ServiceQueue(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delay_seconds: Optional[pulumi.Input[int]] = None,
                 logging_enabled: Optional[pulumi.Input[bool]] = None,
                 maximum_message_size: Optional[pulumi.Input[int]] = None,
                 message_retention_period: Optional[pulumi.Input[int]] = None,
                 polling_wait_seconds: Optional[pulumi.Input[int]] = None,
                 queue_name: Optional[pulumi.Input[str]] = None,
                 visibility_timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a Message Notification Service Queue resource.

        For information about Message Notification Service Queue and how to use it, see [What is Queue](https://www.alibabacloud.com/help/en/message-service/latest/createqueue).

        > **NOTE:** Available since v1.188.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        queue = alicloud.message.ServiceQueue("queue",
            queue_name=name,
            delay_seconds=60478,
            maximum_message_size=12357,
            message_retention_period=256000,
            visibility_timeout=30,
            polling_wait_seconds=3,
            logging_enabled=True)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Message Notification Service Queue can be imported using the id or queue_name, e.g.

        ```sh
        $ pulumi import alicloud:message/serviceQueue:ServiceQueue example <queue_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] delay_seconds: The delay period after which a message sent to the queue can be consumed. Unit: seconds. Valid values: 0-604800 seconds. Default value: 0.
        :param pulumi.Input[bool] logging_enabled: Specifies whether to enable the log management feature. Default value: false. Valid values:
        :param pulumi.Input[int] maximum_message_size: The maximum size of a message body that can be sent to the queue. Unit: bytes. Valid value range: 1024-65536. Default value: 65536.
        :param pulumi.Input[int] message_retention_period: The maximum period for which a message can be retained in the queue. After the specified period, the message is deleted no matter whether the message is consumed. Unit: seconds. Valid values: 60-604800. Default value: 345600.
        :param pulumi.Input[int] polling_wait_seconds: The maximum period for which a ReceiveMessage request waits if no message is available in the queue. Unit: seconds. Valid values: 0-30. Default value: 0.
        :param pulumi.Input[str] queue_name: Two queues on a single account in the same region cannot have the same name. A queue name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 120 characters.
        :param pulumi.Input[int] visibility_timeout: The invisibility period for which the received message remains the Inactive state. Unit: seconds. Valid values: 1-43200. Default value: 30.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceQueueArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Message Notification Service Queue resource.

        For information about Message Notification Service Queue and how to use it, see [What is Queue](https://www.alibabacloud.com/help/en/message-service/latest/createqueue).

        > **NOTE:** Available since v1.188.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        queue = alicloud.message.ServiceQueue("queue",
            queue_name=name,
            delay_seconds=60478,
            maximum_message_size=12357,
            message_retention_period=256000,
            visibility_timeout=30,
            polling_wait_seconds=3,
            logging_enabled=True)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Message Notification Service Queue can be imported using the id or queue_name, e.g.

        ```sh
        $ pulumi import alicloud:message/serviceQueue:ServiceQueue example <queue_name>
        ```

        :param str resource_name: The name of the resource.
        :param ServiceQueueArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceQueueArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delay_seconds: Optional[pulumi.Input[int]] = None,
                 logging_enabled: Optional[pulumi.Input[bool]] = None,
                 maximum_message_size: Optional[pulumi.Input[int]] = None,
                 message_retention_period: Optional[pulumi.Input[int]] = None,
                 polling_wait_seconds: Optional[pulumi.Input[int]] = None,
                 queue_name: Optional[pulumi.Input[str]] = None,
                 visibility_timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceQueueArgs.__new__(ServiceQueueArgs)

            __props__.__dict__["delay_seconds"] = delay_seconds
            __props__.__dict__["logging_enabled"] = logging_enabled
            __props__.__dict__["maximum_message_size"] = maximum_message_size
            __props__.__dict__["message_retention_period"] = message_retention_period
            __props__.__dict__["polling_wait_seconds"] = polling_wait_seconds
            if queue_name is None and not opts.urn:
                raise TypeError("Missing required property 'queue_name'")
            __props__.__dict__["queue_name"] = queue_name
            __props__.__dict__["visibility_timeout"] = visibility_timeout
        super(ServiceQueue, __self__).__init__(
            'alicloud:message/serviceQueue:ServiceQueue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            delay_seconds: Optional[pulumi.Input[int]] = None,
            logging_enabled: Optional[pulumi.Input[bool]] = None,
            maximum_message_size: Optional[pulumi.Input[int]] = None,
            message_retention_period: Optional[pulumi.Input[int]] = None,
            polling_wait_seconds: Optional[pulumi.Input[int]] = None,
            queue_name: Optional[pulumi.Input[str]] = None,
            visibility_timeout: Optional[pulumi.Input[int]] = None) -> 'ServiceQueue':
        """
        Get an existing ServiceQueue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] delay_seconds: The delay period after which a message sent to the queue can be consumed. Unit: seconds. Valid values: 0-604800 seconds. Default value: 0.
        :param pulumi.Input[bool] logging_enabled: Specifies whether to enable the log management feature. Default value: false. Valid values:
        :param pulumi.Input[int] maximum_message_size: The maximum size of a message body that can be sent to the queue. Unit: bytes. Valid value range: 1024-65536. Default value: 65536.
        :param pulumi.Input[int] message_retention_period: The maximum period for which a message can be retained in the queue. After the specified period, the message is deleted no matter whether the message is consumed. Unit: seconds. Valid values: 60-604800. Default value: 345600.
        :param pulumi.Input[int] polling_wait_seconds: The maximum period for which a ReceiveMessage request waits if no message is available in the queue. Unit: seconds. Valid values: 0-30. Default value: 0.
        :param pulumi.Input[str] queue_name: Two queues on a single account in the same region cannot have the same name. A queue name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 120 characters.
        :param pulumi.Input[int] visibility_timeout: The invisibility period for which the received message remains the Inactive state. Unit: seconds. Valid values: 1-43200. Default value: 30.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceQueueState.__new__(_ServiceQueueState)

        __props__.__dict__["delay_seconds"] = delay_seconds
        __props__.__dict__["logging_enabled"] = logging_enabled
        __props__.__dict__["maximum_message_size"] = maximum_message_size
        __props__.__dict__["message_retention_period"] = message_retention_period
        __props__.__dict__["polling_wait_seconds"] = polling_wait_seconds
        __props__.__dict__["queue_name"] = queue_name
        __props__.__dict__["visibility_timeout"] = visibility_timeout
        return ServiceQueue(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="delaySeconds")
    def delay_seconds(self) -> pulumi.Output[int]:
        """
        The delay period after which a message sent to the queue can be consumed. Unit: seconds. Valid values: 0-604800 seconds. Default value: 0.
        """
        return pulumi.get(self, "delay_seconds")

    @property
    @pulumi.getter(name="loggingEnabled")
    def logging_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to enable the log management feature. Default value: false. Valid values:
        """
        return pulumi.get(self, "logging_enabled")

    @property
    @pulumi.getter(name="maximumMessageSize")
    def maximum_message_size(self) -> pulumi.Output[int]:
        """
        The maximum size of a message body that can be sent to the queue. Unit: bytes. Valid value range: 1024-65536. Default value: 65536.
        """
        return pulumi.get(self, "maximum_message_size")

    @property
    @pulumi.getter(name="messageRetentionPeriod")
    def message_retention_period(self) -> pulumi.Output[int]:
        """
        The maximum period for which a message can be retained in the queue. After the specified period, the message is deleted no matter whether the message is consumed. Unit: seconds. Valid values: 60-604800. Default value: 345600.
        """
        return pulumi.get(self, "message_retention_period")

    @property
    @pulumi.getter(name="pollingWaitSeconds")
    def polling_wait_seconds(self) -> pulumi.Output[int]:
        """
        The maximum period for which a ReceiveMessage request waits if no message is available in the queue. Unit: seconds. Valid values: 0-30. Default value: 0.
        """
        return pulumi.get(self, "polling_wait_seconds")

    @property
    @pulumi.getter(name="queueName")
    def queue_name(self) -> pulumi.Output[str]:
        """
        Two queues on a single account in the same region cannot have the same name. A queue name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 120 characters.
        """
        return pulumi.get(self, "queue_name")

    @property
    @pulumi.getter(name="visibilityTimeout")
    def visibility_timeout(self) -> pulumi.Output[int]:
        """
        The invisibility period for which the received message remains the Inactive state. Unit: seconds. Valid values: 1-43200. Default value: 30.
        """
        return pulumi.get(self, "visibility_timeout")

