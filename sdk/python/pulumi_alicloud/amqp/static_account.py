# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StaticAccountArgs', 'StaticAccount']

@pulumi.input_type
class StaticAccountArgs:
    def __init__(__self__, *,
                 access_key: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 secret_key: pulumi.Input[str]):
        """
        The set of arguments for constructing a StaticAccount resource.
        :param pulumi.Input[str] access_key: Access key.
        :param pulumi.Input[str] instance_id: Amqp instance ID.
        :param pulumi.Input[str] secret_key: Secret key.
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "secret_key", secret_key)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Input[str]:
        """
        Access key.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Amqp instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Input[str]:
        """
        Secret key.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_key", value)


@pulumi.input_type
class _StaticAccountState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 master_uid: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StaticAccount resources.
        :param pulumi.Input[str] access_key: Access key.
        :param pulumi.Input[int] create_time: Create time stamp. Unix timestamp, to millisecond level.
        :param pulumi.Input[str] instance_id: Amqp instance ID.
        :param pulumi.Input[str] master_uid: The ID of the user's primary account.
        :param pulumi.Input[str] password: Static password.
        :param pulumi.Input[str] secret_key: Secret key.
        :param pulumi.Input[str] user_name: Static user name.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if master_uid is not None:
            pulumi.set(__self__, "master_uid", master_uid)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Access key.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[int]]:
        """
        Create time stamp. Unix timestamp, to millisecond level.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Amqp instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="masterUid")
    def master_uid(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the user's primary account.
        """
        return pulumi.get(self, "master_uid")

    @master_uid.setter
    def master_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_uid", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Static password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        Secret key.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        Static user name.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class StaticAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Amqp Static Account resource.

        For information about Amqp Static Account and how to use it, see [What is Static Account](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/create-a-pair-of-static-username-and-password).

        > **NOTE:** Available since v1.195.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        access_key = config.get("accessKey")
        if access_key is None:
            access_key = "access_key"
        secret_key = config.get("secretKey")
        if secret_key is None:
            secret_key = "secret_key"
        default_instance = alicloud.amqp.Instance("defaultInstance",
            instance_type="enterprise",
            max_tps="3000",
            queue_capacity="200",
            storage_size="700",
            support_eip=False,
            max_eip_tps="128",
            payment_type="Subscription",
            period=1)
        default_static_account = alicloud.amqp.StaticAccount("defaultStaticAccount",
            instance_id=default_instance.id,
            access_key=access_key,
            secret_key=secret_key)
        ```

        ## Import

        Amqp Static Account can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:amqp/staticAccount:StaticAccount example <instance_id>:<access_key>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: Access key.
        :param pulumi.Input[str] instance_id: Amqp instance ID.
        :param pulumi.Input[str] secret_key: Secret key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StaticAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Amqp Static Account resource.

        For information about Amqp Static Account and how to use it, see [What is Static Account](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/create-a-pair-of-static-username-and-password).

        > **NOTE:** Available since v1.195.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        access_key = config.get("accessKey")
        if access_key is None:
            access_key = "access_key"
        secret_key = config.get("secretKey")
        if secret_key is None:
            secret_key = "secret_key"
        default_instance = alicloud.amqp.Instance("defaultInstance",
            instance_type="enterprise",
            max_tps="3000",
            queue_capacity="200",
            storage_size="700",
            support_eip=False,
            max_eip_tps="128",
            payment_type="Subscription",
            period=1)
        default_static_account = alicloud.amqp.StaticAccount("defaultStaticAccount",
            instance_id=default_instance.id,
            access_key=access_key,
            secret_key=secret_key)
        ```

        ## Import

        Amqp Static Account can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:amqp/staticAccount:StaticAccount example <instance_id>:<access_key>
        ```

        :param str resource_name: The name of the resource.
        :param StaticAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StaticAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StaticAccountArgs.__new__(StaticAccountArgs)

            if access_key is None and not opts.urn:
                raise TypeError("Missing required property 'access_key'")
            __props__.__dict__["access_key"] = access_key
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_key'")
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            __props__.__dict__["create_time"] = None
            __props__.__dict__["master_uid"] = None
            __props__.__dict__["password"] = None
            __props__.__dict__["user_name"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(StaticAccount, __self__).__init__(
            'alicloud:amqp/staticAccount:StaticAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[int]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            master_uid: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'StaticAccount':
        """
        Get an existing StaticAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: Access key.
        :param pulumi.Input[int] create_time: Create time stamp. Unix timestamp, to millisecond level.
        :param pulumi.Input[str] instance_id: Amqp instance ID.
        :param pulumi.Input[str] master_uid: The ID of the user's primary account.
        :param pulumi.Input[str] password: Static password.
        :param pulumi.Input[str] secret_key: Secret key.
        :param pulumi.Input[str] user_name: Static user name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StaticAccountState.__new__(_StaticAccountState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["master_uid"] = master_uid
        __props__.__dict__["password"] = password
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["user_name"] = user_name
        return StaticAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        Access key.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[int]:
        """
        Create time stamp. Unix timestamp, to millisecond level.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Amqp instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="masterUid")
    def master_uid(self) -> pulumi.Output[str]:
        """
        The ID of the user's primary account.
        """
        return pulumi.get(self, "master_uid")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Static password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        Secret key.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        Static user name.
        """
        return pulumi.get(self, "user_name")

