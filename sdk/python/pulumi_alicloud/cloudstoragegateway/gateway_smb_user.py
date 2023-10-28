# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GatewaySmbUserArgs', 'GatewaySmbUser']

@pulumi.input_type
class GatewaySmbUserArgs:
    def __init__(__self__, *,
                 gateway_id: pulumi.Input[str],
                 password: pulumi.Input[str],
                 username: pulumi.Input[str]):
        """
        The set of arguments for constructing a GatewaySmbUser resource.
        :param pulumi.Input[str] gateway_id: The Gateway ID of the Gateway SMB User.
        :param pulumi.Input[str] password: The password of the Gateway SMB User.
        :param pulumi.Input[str] username: The username of the Gateway SMB User.
        """
        pulumi.set(__self__, "gateway_id", gateway_id)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Input[str]:
        """
        The Gateway ID of the Gateway SMB User.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        The password of the Gateway SMB User.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The username of the Gateway SMB User.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _GatewaySmbUserState:
    def __init__(__self__, *,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GatewaySmbUser resources.
        :param pulumi.Input[str] gateway_id: The Gateway ID of the Gateway SMB User.
        :param pulumi.Input[str] password: The password of the Gateway SMB User.
        :param pulumi.Input[str] username: The username of the Gateway SMB User.
        """
        if gateway_id is not None:
            pulumi.set(__self__, "gateway_id", gateway_id)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Gateway ID of the Gateway SMB User.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the Gateway SMB User.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username of the Gateway SMB User.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class GatewaySmbUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cloud Storage Gateway Gateway SMB User resource.

        For information about Cloud Storage Gateway Gateway SMB User and how to use it, see [What is Gateway SMB User](https://www.alibabacloud.com/help/en/cloud-storage-gateway/latest/creategatewaysmbuser).

        > **NOTE:** Available since v1.142.0.

        ## Import

        Cloud Storage Gateway Gateway SMB User can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cloudstoragegateway/gatewaySmbUser:GatewaySmbUser example <gateway_id>:<username>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_id: The Gateway ID of the Gateway SMB User.
        :param pulumi.Input[str] password: The password of the Gateway SMB User.
        :param pulumi.Input[str] username: The username of the Gateway SMB User.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewaySmbUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Storage Gateway Gateway SMB User resource.

        For information about Cloud Storage Gateway Gateway SMB User and how to use it, see [What is Gateway SMB User](https://www.alibabacloud.com/help/en/cloud-storage-gateway/latest/creategatewaysmbuser).

        > **NOTE:** Available since v1.142.0.

        ## Import

        Cloud Storage Gateway Gateway SMB User can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cloudstoragegateway/gatewaySmbUser:GatewaySmbUser example <gateway_id>:<username>
        ```

        :param str resource_name: The name of the resource.
        :param GatewaySmbUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewaySmbUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GatewaySmbUserArgs.__new__(GatewaySmbUserArgs)

            if gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_id'")
            __props__.__dict__["gateway_id"] = gateway_id
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(GatewaySmbUser, __self__).__init__(
            'alicloud:cloudstoragegateway/gatewaySmbUser:GatewaySmbUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            gateway_id: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'GatewaySmbUser':
        """
        Get an existing GatewaySmbUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_id: The Gateway ID of the Gateway SMB User.
        :param pulumi.Input[str] password: The password of the Gateway SMB User.
        :param pulumi.Input[str] username: The username of the Gateway SMB User.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewaySmbUserState.__new__(_GatewaySmbUserState)

        __props__.__dict__["gateway_id"] = gateway_id
        __props__.__dict__["password"] = password
        __props__.__dict__["username"] = username
        return GatewaySmbUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Output[str]:
        """
        The Gateway ID of the Gateway SMB User.
        """
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The password of the Gateway SMB User.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The username of the Gateway SMB User.
        """
        return pulumi.get(self, "username")

