# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 account_password: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 account_description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[str] account_name: The name of the account. Valid values: `root`.
        :param pulumi.Input[str] account_password: The Password of the Account.
               * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%!^(MISSING)&*()_+-=`.
               * The password must be `8` to `32` characters in length.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[str] account_description: The description of the account.
               * The description must start with a letter, and cannot start with `http://` or `https://`.
               * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "account_password", account_password)
        pulumi.set(__self__, "instance_id", instance_id)
        if account_description is not None:
            pulumi.set(__self__, "account_description", account_description)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        The name of the account. Valid values: `root`.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> pulumi.Input[str]:
        """
        The Password of the Account.
        * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%!^(MISSING)&*()_+-=`.
        * The password must be `8` to `32` characters in length.
        """
        return pulumi.get(self, "account_password")

    @account_password.setter
    def account_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_password", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="accountDescription")
    def account_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the account.
        * The description must start with a letter, and cannot start with `http://` or `https://`.
        * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "account_description")

    @account_description.setter
    def account_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_description", value)


@pulumi.input_type
class _AccountState:
    def __init__(__self__, *,
                 account_description: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Account resources.
        :param pulumi.Input[str] account_description: The description of the account.
               * The description must start with a letter, and cannot start with `http://` or `https://`.
               * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
        :param pulumi.Input[str] account_name: The name of the account. Valid values: `root`.
        :param pulumi.Input[str] account_password: The Password of the Account.
               * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%!^(MISSING)&*()_+-=`.
               * The password must be `8` to `32` characters in length.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[str] status: The status of the account. Valid values: `Unavailable`, `Available`.
        """
        if account_description is not None:
            pulumi.set(__self__, "account_description", account_description)
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if account_password is not None:
            pulumi.set(__self__, "account_password", account_password)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accountDescription")
    def account_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the account.
        * The description must start with a letter, and cannot start with `http://` or `https://`.
        * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "account_description")

    @account_description.setter
    def account_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_description", value)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the account. Valid values: `root`.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> Optional[pulumi.Input[str]]:
        """
        The Password of the Account.
        * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%!^(MISSING)&*()_+-=`.
        * The password must be `8` to `32` characters in length.
        """
        return pulumi.get(self, "account_password")

    @account_password.setter
    def account_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_password", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the account. Valid values: `Unavailable`, `Available`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_description: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a MongoDB Account resource.

        For information about MongoDB Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/doc-detail/62154.html).

        > **NOTE:** Available since v1.148.0.

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
        default_zones = alicloud.mongodb.get_zones()
        index = len(default_zones.zones) - 1
        zone_id = default_zones.zones[index].id
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name=name,
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=zone_id)
        default_instance = alicloud.mongodb.Instance("defaultInstance",
            engine_version="4.2",
            db_instance_class="dds.mongo.mid",
            db_instance_storage=10,
            vswitch_id=default_switch.id,
            security_ip_lists=[
                "10.168.1.12",
                "100.69.7.112",
            ],
            tags={
                "Created": "TF",
                "For": "example",
            })
        default_account = alicloud.mongodb.Account("defaultAccount",
            account_name="root",
            account_password="Example_123",
            instance_id=default_instance.id,
            account_description=name)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        MongoDB Account can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:mongodb/account:Account example <instance_id>:<account_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_description: The description of the account.
               * The description must start with a letter, and cannot start with `http://` or `https://`.
               * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
        :param pulumi.Input[str] account_name: The name of the account. Valid values: `root`.
        :param pulumi.Input[str] account_password: The Password of the Account.
               * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%!^(MISSING)&*()_+-=`.
               * The password must be `8` to `32` characters in length.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a MongoDB Account resource.

        For information about MongoDB Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/doc-detail/62154.html).

        > **NOTE:** Available since v1.148.0.

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
        default_zones = alicloud.mongodb.get_zones()
        index = len(default_zones.zones) - 1
        zone_id = default_zones.zones[index].id
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name=name,
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=zone_id)
        default_instance = alicloud.mongodb.Instance("defaultInstance",
            engine_version="4.2",
            db_instance_class="dds.mongo.mid",
            db_instance_storage=10,
            vswitch_id=default_switch.id,
            security_ip_lists=[
                "10.168.1.12",
                "100.69.7.112",
            ],
            tags={
                "Created": "TF",
                "For": "example",
            })
        default_account = alicloud.mongodb.Account("defaultAccount",
            account_name="root",
            account_password="Example_123",
            instance_id=default_instance.id,
            account_description=name)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        MongoDB Account can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:mongodb/account:Account example <instance_id>:<account_name>
        ```

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_description: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountArgs.__new__(AccountArgs)

            __props__.__dict__["account_description"] = account_description
            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            if account_password is None and not opts.urn:
                raise TypeError("Missing required property 'account_password'")
            __props__.__dict__["account_password"] = None if account_password is None else pulumi.Output.secret(account_password)
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["status"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accountPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Account, __self__).__init__(
            'alicloud:mongodb/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_description: Optional[pulumi.Input[str]] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            account_password: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_description: The description of the account.
               * The description must start with a letter, and cannot start with `http://` or `https://`.
               * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
        :param pulumi.Input[str] account_name: The name of the account. Valid values: `root`.
        :param pulumi.Input[str] account_password: The Password of the Account.
               * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%!^(MISSING)&*()_+-=`.
               * The password must be `8` to `32` characters in length.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[str] status: The status of the account. Valid values: `Unavailable`, `Available`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountState.__new__(_AccountState)

        __props__.__dict__["account_description"] = account_description
        __props__.__dict__["account_name"] = account_name
        __props__.__dict__["account_password"] = account_password
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["status"] = status
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountDescription")
    def account_description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the account.
        * The description must start with a letter, and cannot start with `http://` or `https://`.
        * It must be `2` to `256` characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "account_description")

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        The name of the account. Valid values: `root`.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> pulumi.Output[str]:
        """
        The Password of the Account.
        * The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!#$%!^(MISSING)&*()_+-=`.
        * The password must be `8` to `32` characters in length.
        """
        return pulumi.get(self, "account_password")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the account. Valid values: `Unavailable`, `Available`.
        """
        return pulumi.get(self, "status")

