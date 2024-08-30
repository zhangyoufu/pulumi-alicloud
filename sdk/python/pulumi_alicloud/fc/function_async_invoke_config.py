# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FunctionAsyncInvokeConfigArgs', 'FunctionAsyncInvokeConfig']

@pulumi.input_type
class FunctionAsyncInvokeConfigArgs:
    def __init__(__self__, *,
                 function_name: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 destination_config: Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigArgs']] = None,
                 maximum_event_age_in_seconds: Optional[pulumi.Input[int]] = None,
                 maximum_retry_attempts: Optional[pulumi.Input[int]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 stateful_invocation: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a FunctionAsyncInvokeConfig resource.
        :param pulumi.Input[str] function_name: Name of the Function Compute Function.
        :param pulumi.Input[str] service_name: Name of the Function Compute Function, omitting any version or alias qualifier.
        :param pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigArgs'] destination_config: Configuration block with destination configuration. See `destination_config` below.
        :param pulumi.Input[int] maximum_event_age_in_seconds: Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
        :param pulumi.Input[int] maximum_retry_attempts: Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
        :param pulumi.Input[str] qualifier: Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
        :param pulumi.Input[bool] stateful_invocation: Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
        """
        pulumi.set(__self__, "function_name", function_name)
        pulumi.set(__self__, "service_name", service_name)
        if destination_config is not None:
            pulumi.set(__self__, "destination_config", destination_config)
        if maximum_event_age_in_seconds is not None:
            pulumi.set(__self__, "maximum_event_age_in_seconds", maximum_event_age_in_seconds)
        if maximum_retry_attempts is not None:
            pulumi.set(__self__, "maximum_retry_attempts", maximum_retry_attempts)
        if qualifier is not None:
            pulumi.set(__self__, "qualifier", qualifier)
        if stateful_invocation is not None:
            pulumi.set(__self__, "stateful_invocation", stateful_invocation)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Input[str]:
        """
        Name of the Function Compute Function.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        Name of the Function Compute Function, omitting any version or alias qualifier.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="destinationConfig")
    def destination_config(self) -> Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigArgs']]:
        """
        Configuration block with destination configuration. See `destination_config` below.
        """
        return pulumi.get(self, "destination_config")

    @destination_config.setter
    def destination_config(self, value: Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigArgs']]):
        pulumi.set(self, "destination_config", value)

    @property
    @pulumi.getter(name="maximumEventAgeInSeconds")
    def maximum_event_age_in_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
        """
        return pulumi.get(self, "maximum_event_age_in_seconds")

    @maximum_event_age_in_seconds.setter
    def maximum_event_age_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_event_age_in_seconds", value)

    @property
    @pulumi.getter(name="maximumRetryAttempts")
    def maximum_retry_attempts(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
        """
        return pulumi.get(self, "maximum_retry_attempts")

    @maximum_retry_attempts.setter
    def maximum_retry_attempts(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_retry_attempts", value)

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[pulumi.Input[str]]:
        """
        Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
        """
        return pulumi.get(self, "qualifier")

    @qualifier.setter
    def qualifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "qualifier", value)

    @property
    @pulumi.getter(name="statefulInvocation")
    def stateful_invocation(self) -> Optional[pulumi.Input[bool]]:
        """
        Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
        """
        return pulumi.get(self, "stateful_invocation")

    @stateful_invocation.setter
    def stateful_invocation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "stateful_invocation", value)


@pulumi.input_type
class _FunctionAsyncInvokeConfigState:
    def __init__(__self__, *,
                 created_time: Optional[pulumi.Input[str]] = None,
                 destination_config: Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigArgs']] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 last_modified_time: Optional[pulumi.Input[str]] = None,
                 maximum_event_age_in_seconds: Optional[pulumi.Input[int]] = None,
                 maximum_retry_attempts: Optional[pulumi.Input[int]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 stateful_invocation: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering FunctionAsyncInvokeConfig resources.
        :param pulumi.Input[str] created_time: The date this resource was created.
        :param pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigArgs'] destination_config: Configuration block with destination configuration. See `destination_config` below.
        :param pulumi.Input[str] function_name: Name of the Function Compute Function.
        :param pulumi.Input[str] last_modified_time: The date this resource was last modified.
        :param pulumi.Input[int] maximum_event_age_in_seconds: Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
        :param pulumi.Input[int] maximum_retry_attempts: Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
        :param pulumi.Input[str] qualifier: Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
        :param pulumi.Input[str] service_name: Name of the Function Compute Function, omitting any version or alias qualifier.
        :param pulumi.Input[bool] stateful_invocation: Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
        """
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if destination_config is not None:
            pulumi.set(__self__, "destination_config", destination_config)
        if function_name is not None:
            pulumi.set(__self__, "function_name", function_name)
        if last_modified_time is not None:
            pulumi.set(__self__, "last_modified_time", last_modified_time)
        if maximum_event_age_in_seconds is not None:
            pulumi.set(__self__, "maximum_event_age_in_seconds", maximum_event_age_in_seconds)
        if maximum_retry_attempts is not None:
            pulumi.set(__self__, "maximum_retry_attempts", maximum_retry_attempts)
        if qualifier is not None:
            pulumi.set(__self__, "qualifier", qualifier)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if stateful_invocation is not None:
            pulumi.set(__self__, "stateful_invocation", stateful_invocation)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[pulumi.Input[str]]:
        """
        The date this resource was created.
        """
        return pulumi.get(self, "created_time")

    @created_time.setter
    def created_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_time", value)

    @property
    @pulumi.getter(name="destinationConfig")
    def destination_config(self) -> Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigArgs']]:
        """
        Configuration block with destination configuration. See `destination_config` below.
        """
        return pulumi.get(self, "destination_config")

    @destination_config.setter
    def destination_config(self, value: Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigArgs']]):
        pulumi.set(self, "destination_config", value)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Function Compute Function.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[pulumi.Input[str]]:
        """
        The date this resource was last modified.
        """
        return pulumi.get(self, "last_modified_time")

    @last_modified_time.setter
    def last_modified_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_modified_time", value)

    @property
    @pulumi.getter(name="maximumEventAgeInSeconds")
    def maximum_event_age_in_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
        """
        return pulumi.get(self, "maximum_event_age_in_seconds")

    @maximum_event_age_in_seconds.setter
    def maximum_event_age_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_event_age_in_seconds", value)

    @property
    @pulumi.getter(name="maximumRetryAttempts")
    def maximum_retry_attempts(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
        """
        return pulumi.get(self, "maximum_retry_attempts")

    @maximum_retry_attempts.setter
    def maximum_retry_attempts(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "maximum_retry_attempts", value)

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[pulumi.Input[str]]:
        """
        Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
        """
        return pulumi.get(self, "qualifier")

    @qualifier.setter
    def qualifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "qualifier", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Function Compute Function, omitting any version or alias qualifier.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="statefulInvocation")
    def stateful_invocation(self) -> Optional[pulumi.Input[bool]]:
        """
        Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
        """
        return pulumi.get(self, "stateful_invocation")

    @stateful_invocation.setter
    def stateful_invocation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "stateful_invocation", value)


class FunctionAsyncInvokeConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_config: Optional[pulumi.Input[Union['FunctionAsyncInvokeConfigDestinationConfigArgs', 'FunctionAsyncInvokeConfigDestinationConfigArgsDict']]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 maximum_event_age_in_seconds: Optional[pulumi.Input[int]] = None,
                 maximum_retry_attempts: Optional[pulumi.Input[int]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 stateful_invocation: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages an asynchronous invocation configuration for a FC Function or Alias.\\
         For the detailed information, please refer to the [developer guide](https://www.alibabacloud.com/help/en/fc/developer-reference/api-fc-open-2021-04-06-putfunctionasyncinvokeconfig).

        > **NOTE:** Available since v1.100.0.

        ## Example Usage

        ### Destination Configuration

        > **NOTE** Ensure the FC Function RAM Role has necessary permissions for the destination, such as `mns:SendMessage`, `mns:PublishMessage` or `fc:InvokeFunction`, otherwise the API will return a generic error.

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        default = alicloud.get_account()
        default_get_regions = alicloud.get_regions(current=True)
        default_integer = random.index.Integer("default",
            max=99999,
            min=10000)
        default_role = alicloud.ram.Role("default",
            name=f"examplerole{default_integer['result']}",
            document=\"\"\"\\x09{
        \\x09\\x09"Statement": [
        \\x09\\x09  {
        \\x09\\x09\\x09"Action": "sts:AssumeRole",
        \\x09\\x09\\x09"Effect": "Allow",
        \\x09\\x09\\x09"Principal": {
        \\x09\\x09\\x09  "Service": [
        \\x09\\x09\\x09\\x09"fc.aliyuncs.com"
        \\x09\\x09\\x09  ]
        \\x09\\x09\\x09}
        \\x09\\x09  }
        \\x09\\x09],
        \\x09\\x09"Version": "1"
        \\x09}
        \"\"\",
            description="this is a example",
            force=True)
        default_policy = alicloud.ram.Policy("default",
            policy_name=f"examplepolicy{default_integer['result']}",
            policy_document=\"\"\"\\x09{
        \\x09\\x09"Version": "1",
        \\x09\\x09"Statement": [
        \\x09\\x09  {
        \\x09\\x09\\x09"Action": "mns:*",
        \\x09\\x09\\x09"Resource": "*",
        \\x09\\x09\\x09"Effect": "Allow"
        \\x09\\x09  }
        \\x09\\x09]
        \\x09  }
        \"\"\")
        default_role_policy_attachment = alicloud.ram.RolePolicyAttachment("default",
            role_name=default_role.name,
            policy_name=default_policy.policy_name,
            policy_type="Custom")
        default_service = alicloud.fc.Service("default",
            name=f"example-value-{default_integer['result']}",
            description="example-value",
            role=default_role.arn,
            internet_access=False)
        default_bucket = alicloud.oss.Bucket("default", bucket=f"terraform-example-{default_integer['result']}")
        # If you upload the function by OSS Bucket, you need to specify path can't upload by content.
        default_bucket_object = alicloud.oss.BucketObject("default",
            bucket=default_bucket.id,
            key="index.py",
            content=\"\"\"import logging 
        def handler(event, context): 
        logger = logging.getLogger() 
        logger.info('hello world') 
        return 'hello world'\"\"\")
        default_function = alicloud.fc.Function("default",
            service=default_service.name,
            name=f"terraform-example-{default_integer['result']}",
            description="example",
            oss_bucket=default_bucket.id,
            oss_key=default_bucket_object.key,
            memory_size=512,
            runtime="python3.10",
            handler="hello.handler")
        default_queue = alicloud.mns.Queue("default", name=f"terraform-example-{default_integer['result']}")
        default_topic = alicloud.mns.Topic("default", name=f"terraform-example-{default_integer['result']}")
        default_function_async_invoke_config = alicloud.fc.FunctionAsyncInvokeConfig("default",
            service_name=default_service.name,
            function_name=default_function.name,
            destination_config={
                "on_failure": {
                    "destination": default_queue.name.apply(lambda name: f"acs:mns:{default_get_regions.regions[0].id}:{default.id}:/queues/{name}/messages"),
                },
                "on_success": {
                    "destination": default_topic.name.apply(lambda name: f"acs:mns:{default_get_regions.regions[0].id}:{default.id}:/topics/{name}/messages"),
                },
            },
            maximum_event_age_in_seconds=60,
            maximum_retry_attempts=0,
            stateful_invocation=True,
            qualifier="LATEST")
        ```

        ## Import

        Function Compute Function Async Invoke Configs can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:fc/functionAsyncInvokeConfig:FunctionAsyncInvokeConfig example my_function
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['FunctionAsyncInvokeConfigDestinationConfigArgs', 'FunctionAsyncInvokeConfigDestinationConfigArgsDict']] destination_config: Configuration block with destination configuration. See `destination_config` below.
        :param pulumi.Input[str] function_name: Name of the Function Compute Function.
        :param pulumi.Input[int] maximum_event_age_in_seconds: Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
        :param pulumi.Input[int] maximum_retry_attempts: Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
        :param pulumi.Input[str] qualifier: Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
        :param pulumi.Input[str] service_name: Name of the Function Compute Function, omitting any version or alias qualifier.
        :param pulumi.Input[bool] stateful_invocation: Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionAsyncInvokeConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an asynchronous invocation configuration for a FC Function or Alias.\\
         For the detailed information, please refer to the [developer guide](https://www.alibabacloud.com/help/en/fc/developer-reference/api-fc-open-2021-04-06-putfunctionasyncinvokeconfig).

        > **NOTE:** Available since v1.100.0.

        ## Example Usage

        ### Destination Configuration

        > **NOTE** Ensure the FC Function RAM Role has necessary permissions for the destination, such as `mns:SendMessage`, `mns:PublishMessage` or `fc:InvokeFunction`, otherwise the API will return a generic error.

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        default = alicloud.get_account()
        default_get_regions = alicloud.get_regions(current=True)
        default_integer = random.index.Integer("default",
            max=99999,
            min=10000)
        default_role = alicloud.ram.Role("default",
            name=f"examplerole{default_integer['result']}",
            document=\"\"\"\\x09{
        \\x09\\x09"Statement": [
        \\x09\\x09  {
        \\x09\\x09\\x09"Action": "sts:AssumeRole",
        \\x09\\x09\\x09"Effect": "Allow",
        \\x09\\x09\\x09"Principal": {
        \\x09\\x09\\x09  "Service": [
        \\x09\\x09\\x09\\x09"fc.aliyuncs.com"
        \\x09\\x09\\x09  ]
        \\x09\\x09\\x09}
        \\x09\\x09  }
        \\x09\\x09],
        \\x09\\x09"Version": "1"
        \\x09}
        \"\"\",
            description="this is a example",
            force=True)
        default_policy = alicloud.ram.Policy("default",
            policy_name=f"examplepolicy{default_integer['result']}",
            policy_document=\"\"\"\\x09{
        \\x09\\x09"Version": "1",
        \\x09\\x09"Statement": [
        \\x09\\x09  {
        \\x09\\x09\\x09"Action": "mns:*",
        \\x09\\x09\\x09"Resource": "*",
        \\x09\\x09\\x09"Effect": "Allow"
        \\x09\\x09  }
        \\x09\\x09]
        \\x09  }
        \"\"\")
        default_role_policy_attachment = alicloud.ram.RolePolicyAttachment("default",
            role_name=default_role.name,
            policy_name=default_policy.policy_name,
            policy_type="Custom")
        default_service = alicloud.fc.Service("default",
            name=f"example-value-{default_integer['result']}",
            description="example-value",
            role=default_role.arn,
            internet_access=False)
        default_bucket = alicloud.oss.Bucket("default", bucket=f"terraform-example-{default_integer['result']}")
        # If you upload the function by OSS Bucket, you need to specify path can't upload by content.
        default_bucket_object = alicloud.oss.BucketObject("default",
            bucket=default_bucket.id,
            key="index.py",
            content=\"\"\"import logging 
        def handler(event, context): 
        logger = logging.getLogger() 
        logger.info('hello world') 
        return 'hello world'\"\"\")
        default_function = alicloud.fc.Function("default",
            service=default_service.name,
            name=f"terraform-example-{default_integer['result']}",
            description="example",
            oss_bucket=default_bucket.id,
            oss_key=default_bucket_object.key,
            memory_size=512,
            runtime="python3.10",
            handler="hello.handler")
        default_queue = alicloud.mns.Queue("default", name=f"terraform-example-{default_integer['result']}")
        default_topic = alicloud.mns.Topic("default", name=f"terraform-example-{default_integer['result']}")
        default_function_async_invoke_config = alicloud.fc.FunctionAsyncInvokeConfig("default",
            service_name=default_service.name,
            function_name=default_function.name,
            destination_config={
                "on_failure": {
                    "destination": default_queue.name.apply(lambda name: f"acs:mns:{default_get_regions.regions[0].id}:{default.id}:/queues/{name}/messages"),
                },
                "on_success": {
                    "destination": default_topic.name.apply(lambda name: f"acs:mns:{default_get_regions.regions[0].id}:{default.id}:/topics/{name}/messages"),
                },
            },
            maximum_event_age_in_seconds=60,
            maximum_retry_attempts=0,
            stateful_invocation=True,
            qualifier="LATEST")
        ```

        ## Import

        Function Compute Function Async Invoke Configs can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:fc/functionAsyncInvokeConfig:FunctionAsyncInvokeConfig example my_function
        ```

        :param str resource_name: The name of the resource.
        :param FunctionAsyncInvokeConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionAsyncInvokeConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_config: Optional[pulumi.Input[Union['FunctionAsyncInvokeConfigDestinationConfigArgs', 'FunctionAsyncInvokeConfigDestinationConfigArgsDict']]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 maximum_event_age_in_seconds: Optional[pulumi.Input[int]] = None,
                 maximum_retry_attempts: Optional[pulumi.Input[int]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 stateful_invocation: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionAsyncInvokeConfigArgs.__new__(FunctionAsyncInvokeConfigArgs)

            __props__.__dict__["destination_config"] = destination_config
            if function_name is None and not opts.urn:
                raise TypeError("Missing required property 'function_name'")
            __props__.__dict__["function_name"] = function_name
            __props__.__dict__["maximum_event_age_in_seconds"] = maximum_event_age_in_seconds
            __props__.__dict__["maximum_retry_attempts"] = maximum_retry_attempts
            __props__.__dict__["qualifier"] = qualifier
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["stateful_invocation"] = stateful_invocation
            __props__.__dict__["created_time"] = None
            __props__.__dict__["last_modified_time"] = None
        super(FunctionAsyncInvokeConfig, __self__).__init__(
            'alicloud:fc/functionAsyncInvokeConfig:FunctionAsyncInvokeConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_time: Optional[pulumi.Input[str]] = None,
            destination_config: Optional[pulumi.Input[Union['FunctionAsyncInvokeConfigDestinationConfigArgs', 'FunctionAsyncInvokeConfigDestinationConfigArgsDict']]] = None,
            function_name: Optional[pulumi.Input[str]] = None,
            last_modified_time: Optional[pulumi.Input[str]] = None,
            maximum_event_age_in_seconds: Optional[pulumi.Input[int]] = None,
            maximum_retry_attempts: Optional[pulumi.Input[int]] = None,
            qualifier: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            stateful_invocation: Optional[pulumi.Input[bool]] = None) -> 'FunctionAsyncInvokeConfig':
        """
        Get an existing FunctionAsyncInvokeConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_time: The date this resource was created.
        :param pulumi.Input[Union['FunctionAsyncInvokeConfigDestinationConfigArgs', 'FunctionAsyncInvokeConfigDestinationConfigArgsDict']] destination_config: Configuration block with destination configuration. See `destination_config` below.
        :param pulumi.Input[str] function_name: Name of the Function Compute Function.
        :param pulumi.Input[str] last_modified_time: The date this resource was last modified.
        :param pulumi.Input[int] maximum_event_age_in_seconds: Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
        :param pulumi.Input[int] maximum_retry_attempts: Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
        :param pulumi.Input[str] qualifier: Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
        :param pulumi.Input[str] service_name: Name of the Function Compute Function, omitting any version or alias qualifier.
        :param pulumi.Input[bool] stateful_invocation: Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionAsyncInvokeConfigState.__new__(_FunctionAsyncInvokeConfigState)

        __props__.__dict__["created_time"] = created_time
        __props__.__dict__["destination_config"] = destination_config
        __props__.__dict__["function_name"] = function_name
        __props__.__dict__["last_modified_time"] = last_modified_time
        __props__.__dict__["maximum_event_age_in_seconds"] = maximum_event_age_in_seconds
        __props__.__dict__["maximum_retry_attempts"] = maximum_retry_attempts
        __props__.__dict__["qualifier"] = qualifier
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["stateful_invocation"] = stateful_invocation
        return FunctionAsyncInvokeConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        The date this resource was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="destinationConfig")
    def destination_config(self) -> pulumi.Output[Optional['outputs.FunctionAsyncInvokeConfigDestinationConfig']]:
        """
        Configuration block with destination configuration. See `destination_config` below.
        """
        return pulumi.get(self, "destination_config")

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Output[str]:
        """
        Name of the Function Compute Function.
        """
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[str]:
        """
        The date this resource was last modified.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="maximumEventAgeInSeconds")
    def maximum_event_age_in_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum age of a request that Function Compute sends to a function for processing in seconds. Valid values between 1 and 2592000 (between 60 and 21600 before v1.167.0).
        """
        return pulumi.get(self, "maximum_event_age_in_seconds")

    @property
    @pulumi.getter(name="maximumRetryAttempts")
    def maximum_retry_attempts(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum number of times to retry when the function returns an error. Valid values between 0 and 8 (between 0 and 2 before v1.167.0). Defaults to 2.
        """
        return pulumi.get(self, "maximum_retry_attempts")

    @property
    @pulumi.getter
    def qualifier(self) -> pulumi.Output[Optional[str]]:
        """
        Function Compute Function published version, `LATEST`, or Function Compute Alias name. The default value is `LATEST`.
        """
        return pulumi.get(self, "qualifier")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        Name of the Function Compute Function, omitting any version or alias qualifier.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="statefulInvocation")
    def stateful_invocation(self) -> pulumi.Output[Optional[bool]]:
        """
        Function Compute async job configuration(also known as Task Mode). valid values true or false, default `false`
        """
        return pulumi.get(self, "stateful_invocation")

