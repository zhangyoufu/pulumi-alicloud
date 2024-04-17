# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EcsInvocationArgs', 'EcsInvocation']

@pulumi.input_type
class EcsInvocationArgs:
    def __init__(__self__, *,
                 command_id: pulumi.Input[str],
                 instance_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 frequency: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 repeat_mode: Optional[pulumi.Input[str]] = None,
                 timed: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 windows_password_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EcsInvocation resource.
        :param pulumi.Input[str] command_id: The ID of the command.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: The list of instances to execute the command. You can specify up to 50 instance IDs.
        :param pulumi.Input[str] frequency: The schedule on which the recurring execution of the command takes place. Take note of the following items:
               * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
               * When you set Timed to true, you must specify Frequency.
               * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
        :param pulumi.Input[Mapping[str, Any]] parameters: The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
        :param pulumi.Input[str] repeat_mode: Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeat_mode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeat_mode` is specified.
        :param pulumi.Input[bool] timed: Specifies whether to periodically run the command. Default value: `false`.
        :param pulumi.Input[str] username: The username that is used to run the command on the ECS instance. 
               * For Linux instances, the root username is used.
               * For Windows instances, the System username is used.
               * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
        :param pulumi.Input[str] windows_password_name: The name of the password used to run the command on a Windows instance.
        """
        pulumi.set(__self__, "command_id", command_id)
        pulumi.set(__self__, "instance_ids", instance_ids)
        if frequency is not None:
            pulumi.set(__self__, "frequency", frequency)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if repeat_mode is not None:
            pulumi.set(__self__, "repeat_mode", repeat_mode)
        if timed is not None:
            pulumi.set(__self__, "timed", timed)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if windows_password_name is not None:
            pulumi.set(__self__, "windows_password_name", windows_password_name)

    @property
    @pulumi.getter(name="commandId")
    def command_id(self) -> pulumi.Input[str]:
        """
        The ID of the command.
        """
        return pulumi.get(self, "command_id")

    @command_id.setter
    def command_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "command_id", value)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The list of instances to execute the command. You can specify up to 50 instance IDs.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "instance_ids", value)

    @property
    @pulumi.getter
    def frequency(self) -> Optional[pulumi.Input[str]]:
        """
        The schedule on which the recurring execution of the command takes place. Take note of the following items:
        * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
        * When you set Timed to true, you must specify Frequency.
        * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
        """
        return pulumi.get(self, "frequency")

    @frequency.setter
    def frequency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "frequency", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="repeatMode")
    def repeat_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeat_mode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeat_mode` is specified.
        """
        return pulumi.get(self, "repeat_mode")

    @repeat_mode.setter
    def repeat_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repeat_mode", value)

    @property
    @pulumi.getter
    def timed(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to periodically run the command. Default value: `false`.
        """
        return pulumi.get(self, "timed")

    @timed.setter
    def timed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "timed", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username that is used to run the command on the ECS instance. 
        * For Linux instances, the root username is used.
        * For Windows instances, the System username is used.
        * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="windowsPasswordName")
    def windows_password_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the password used to run the command on a Windows instance.
        """
        return pulumi.get(self, "windows_password_name")

    @windows_password_name.setter
    def windows_password_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "windows_password_name", value)


@pulumi.input_type
class _EcsInvocationState:
    def __init__(__self__, *,
                 command_id: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 repeat_mode: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 timed: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 windows_password_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EcsInvocation resources.
        :param pulumi.Input[str] command_id: The ID of the command.
        :param pulumi.Input[str] frequency: The schedule on which the recurring execution of the command takes place. Take note of the following items:
               * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
               * When you set Timed to true, you must specify Frequency.
               * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: The list of instances to execute the command. You can specify up to 50 instance IDs.
        :param pulumi.Input[Mapping[str, Any]] parameters: The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
        :param pulumi.Input[str] repeat_mode: Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeat_mode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeat_mode` is specified.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[bool] timed: Specifies whether to periodically run the command. Default value: `false`.
        :param pulumi.Input[str] username: The username that is used to run the command on the ECS instance. 
               * For Linux instances, the root username is used.
               * For Windows instances, the System username is used.
               * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
        :param pulumi.Input[str] windows_password_name: The name of the password used to run the command on a Windows instance.
        """
        if command_id is not None:
            pulumi.set(__self__, "command_id", command_id)
        if frequency is not None:
            pulumi.set(__self__, "frequency", frequency)
        if instance_ids is not None:
            pulumi.set(__self__, "instance_ids", instance_ids)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if repeat_mode is not None:
            pulumi.set(__self__, "repeat_mode", repeat_mode)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if timed is not None:
            pulumi.set(__self__, "timed", timed)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if windows_password_name is not None:
            pulumi.set(__self__, "windows_password_name", windows_password_name)

    @property
    @pulumi.getter(name="commandId")
    def command_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the command.
        """
        return pulumi.get(self, "command_id")

    @command_id.setter
    def command_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command_id", value)

    @property
    @pulumi.getter
    def frequency(self) -> Optional[pulumi.Input[str]]:
        """
        The schedule on which the recurring execution of the command takes place. Take note of the following items:
        * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
        * When you set Timed to true, you must specify Frequency.
        * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
        """
        return pulumi.get(self, "frequency")

    @frequency.setter
    def frequency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "frequency", value)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of instances to execute the command. You can specify up to 50 instance IDs.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "instance_ids", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="repeatMode")
    def repeat_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeat_mode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeat_mode` is specified.
        """
        return pulumi.get(self, "repeat_mode")

    @repeat_mode.setter
    def repeat_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repeat_mode", value)

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
    @pulumi.getter
    def timed(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to periodically run the command. Default value: `false`.
        """
        return pulumi.get(self, "timed")

    @timed.setter
    def timed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "timed", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username that is used to run the command on the ECS instance. 
        * For Linux instances, the root username is used.
        * For Windows instances, the System username is used.
        * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="windowsPasswordName")
    def windows_password_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the password used to run the command on a Windows instance.
        """
        return pulumi.get(self, "windows_password_name")

    @windows_password_name.setter
    def windows_password_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "windows_password_name", value)


class EcsInvocation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command_id: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 repeat_mode: Optional[pulumi.Input[str]] = None,
                 timed: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 windows_password_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ECS Invocation resource.

        For information about ECS Invocation and how to use it, see [What is Invocation](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/invokecommand#t9958.html).

        > **NOTE:** Available since v1.168.0+.

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
        default = alicloud.get_zones(available_disk_category="cloud_efficiency",
            available_resource_creation="VSwitch")
        default_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=default.zones[0].id)
        default_get_images = alicloud.ecs.get_images(name_regex="^ubuntu",
            most_recent=True,
            owners="system")
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            zone_id=default.zones[0].id,
            vswitch_name=name)
        default_security_group = alicloud.ecs.SecurityGroup("default",
            name=name,
            vpc_id=default_network.id)
        default_security_group_rule = alicloud.ecs.SecurityGroupRule("default",
            type="ingress",
            ip_protocol="tcp",
            nic_type="intranet",
            policy="accept",
            port_range="22/22",
            priority=1,
            security_group_id=default_security_group.id,
            cidr_ip="172.16.0.0/24")
        default_instance = alicloud.ecs.Instance("default",
            vswitch_id=default_switch.id,
            image_id=default_get_images.images[0].id,
            instance_type=default_get_instance_types.instance_types[0].id,
            system_disk_category="cloud_efficiency",
            internet_charge_type="PayByTraffic",
            internet_max_bandwidth_out=5,
            security_groups=[default_security_group.id],
            instance_name=name)
        default_command = alicloud.ecs.Command("default",
            name=name,
            command_content="ZWNobyBoZWxsbyx7e25hbWV9fQ==",
            description="For Terraform Test",
            type="RunShellScript",
            working_dir="/root",
            enable_parameter=True)
        default_ecs_invocation = alicloud.ecs.EcsInvocation("default",
            command_id=default_command.id,
            instance_ids=[default_instance.id])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ECS Invocation can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsInvocation:EcsInvocation example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] command_id: The ID of the command.
        :param pulumi.Input[str] frequency: The schedule on which the recurring execution of the command takes place. Take note of the following items:
               * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
               * When you set Timed to true, you must specify Frequency.
               * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: The list of instances to execute the command. You can specify up to 50 instance IDs.
        :param pulumi.Input[Mapping[str, Any]] parameters: The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
        :param pulumi.Input[str] repeat_mode: Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeat_mode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeat_mode` is specified.
        :param pulumi.Input[bool] timed: Specifies whether to periodically run the command. Default value: `false`.
        :param pulumi.Input[str] username: The username that is used to run the command on the ECS instance. 
               * For Linux instances, the root username is used.
               * For Windows instances, the System username is used.
               * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
        :param pulumi.Input[str] windows_password_name: The name of the password used to run the command on a Windows instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EcsInvocationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ECS Invocation resource.

        For information about ECS Invocation and how to use it, see [What is Invocation](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/invokecommand#t9958.html).

        > **NOTE:** Available since v1.168.0+.

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
        default = alicloud.get_zones(available_disk_category="cloud_efficiency",
            available_resource_creation="VSwitch")
        default_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=default.zones[0].id)
        default_get_images = alicloud.ecs.get_images(name_regex="^ubuntu",
            most_recent=True,
            owners="system")
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            zone_id=default.zones[0].id,
            vswitch_name=name)
        default_security_group = alicloud.ecs.SecurityGroup("default",
            name=name,
            vpc_id=default_network.id)
        default_security_group_rule = alicloud.ecs.SecurityGroupRule("default",
            type="ingress",
            ip_protocol="tcp",
            nic_type="intranet",
            policy="accept",
            port_range="22/22",
            priority=1,
            security_group_id=default_security_group.id,
            cidr_ip="172.16.0.0/24")
        default_instance = alicloud.ecs.Instance("default",
            vswitch_id=default_switch.id,
            image_id=default_get_images.images[0].id,
            instance_type=default_get_instance_types.instance_types[0].id,
            system_disk_category="cloud_efficiency",
            internet_charge_type="PayByTraffic",
            internet_max_bandwidth_out=5,
            security_groups=[default_security_group.id],
            instance_name=name)
        default_command = alicloud.ecs.Command("default",
            name=name,
            command_content="ZWNobyBoZWxsbyx7e25hbWV9fQ==",
            description="For Terraform Test",
            type="RunShellScript",
            working_dir="/root",
            enable_parameter=True)
        default_ecs_invocation = alicloud.ecs.EcsInvocation("default",
            command_id=default_command.id,
            instance_ids=[default_instance.id])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ECS Invocation can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsInvocation:EcsInvocation example <id>
        ```

        :param str resource_name: The name of the resource.
        :param EcsInvocationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EcsInvocationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command_id: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 repeat_mode: Optional[pulumi.Input[str]] = None,
                 timed: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 windows_password_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EcsInvocationArgs.__new__(EcsInvocationArgs)

            if command_id is None and not opts.urn:
                raise TypeError("Missing required property 'command_id'")
            __props__.__dict__["command_id"] = command_id
            __props__.__dict__["frequency"] = frequency
            if instance_ids is None and not opts.urn:
                raise TypeError("Missing required property 'instance_ids'")
            __props__.__dict__["instance_ids"] = instance_ids
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["repeat_mode"] = repeat_mode
            __props__.__dict__["timed"] = timed
            __props__.__dict__["username"] = username
            __props__.__dict__["windows_password_name"] = windows_password_name
            __props__.__dict__["status"] = None
        super(EcsInvocation, __self__).__init__(
            'alicloud:ecs/ecsInvocation:EcsInvocation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            command_id: Optional[pulumi.Input[str]] = None,
            frequency: Optional[pulumi.Input[str]] = None,
            instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            repeat_mode: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            timed: Optional[pulumi.Input[bool]] = None,
            username: Optional[pulumi.Input[str]] = None,
            windows_password_name: Optional[pulumi.Input[str]] = None) -> 'EcsInvocation':
        """
        Get an existing EcsInvocation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] command_id: The ID of the command.
        :param pulumi.Input[str] frequency: The schedule on which the recurring execution of the command takes place. Take note of the following items:
               * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
               * When you set Timed to true, you must specify Frequency.
               * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: The list of instances to execute the command. You can specify up to 50 instance IDs.
        :param pulumi.Input[Mapping[str, Any]] parameters: The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
        :param pulumi.Input[str] repeat_mode: Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeat_mode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeat_mode` is specified.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[bool] timed: Specifies whether to periodically run the command. Default value: `false`.
        :param pulumi.Input[str] username: The username that is used to run the command on the ECS instance. 
               * For Linux instances, the root username is used.
               * For Windows instances, the System username is used.
               * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
        :param pulumi.Input[str] windows_password_name: The name of the password used to run the command on a Windows instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EcsInvocationState.__new__(_EcsInvocationState)

        __props__.__dict__["command_id"] = command_id
        __props__.__dict__["frequency"] = frequency
        __props__.__dict__["instance_ids"] = instance_ids
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["repeat_mode"] = repeat_mode
        __props__.__dict__["status"] = status
        __props__.__dict__["timed"] = timed
        __props__.__dict__["username"] = username
        __props__.__dict__["windows_password_name"] = windows_password_name
        return EcsInvocation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="commandId")
    def command_id(self) -> pulumi.Output[str]:
        """
        The ID of the command.
        """
        return pulumi.get(self, "command_id")

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Output[Optional[str]]:
        """
        The schedule on which the recurring execution of the command takes place. Take note of the following items:
        * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
        * When you set Timed to true, you must specify Frequency.
        * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of instances to execute the command. You can specify up to 50 instance IDs.
        """
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="repeatMode")
    def repeat_mode(self) -> pulumi.Output[str]:
        """
        Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeat_mode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeat_mode` is specified.
        """
        return pulumi.get(self, "repeat_mode")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def timed(self) -> pulumi.Output[bool]:
        """
        Specifies whether to periodically run the command. Default value: `false`.
        """
        return pulumi.get(self, "timed")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The username that is used to run the command on the ECS instance. 
        * For Linux instances, the root username is used.
        * For Windows instances, the System username is used.
        * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="windowsPasswordName")
    def windows_password_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the password used to run the command on a Windows instance.
        """
        return pulumi.get(self, "windows_password_name")

