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

__all__ = ['LakeAccountArgs', 'LakeAccount']

@pulumi.input_type
class LakeAccountArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 account_password: pulumi.Input[str],
                 db_cluster_id: pulumi.Input[str],
                 account_description: Optional[pulumi.Input[str]] = None,
                 account_privileges: Optional[pulumi.Input[Sequence[pulumi.Input['LakeAccountAccountPrivilegeArgs']]]] = None,
                 account_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LakeAccount resource.
        :param pulumi.Input[str] account_name: The name of the account.
        :param pulumi.Input[str] account_password: AccountPassword.
        :param pulumi.Input[str] db_cluster_id: The DBCluster ID.
        :param pulumi.Input[str] account_description: The description of the account.
        :param pulumi.Input[Sequence[pulumi.Input['LakeAccountAccountPrivilegeArgs']]] account_privileges: List of permissions granted. See `account_privileges` below.
        :param pulumi.Input[str] account_type: The type of the account.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "account_password", account_password)
        pulumi.set(__self__, "db_cluster_id", db_cluster_id)
        if account_description is not None:
            pulumi.set(__self__, "account_description", account_description)
        if account_privileges is not None:
            pulumi.set(__self__, "account_privileges", account_privileges)
        if account_type is not None:
            pulumi.set(__self__, "account_type", account_type)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        The name of the account.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> pulumi.Input[str]:
        """
        AccountPassword.
        """
        return pulumi.get(self, "account_password")

    @account_password.setter
    def account_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_password", value)

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> pulumi.Input[str]:
        """
        The DBCluster ID.
        """
        return pulumi.get(self, "db_cluster_id")

    @db_cluster_id.setter
    def db_cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_cluster_id", value)

    @property
    @pulumi.getter(name="accountDescription")
    def account_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the account.
        """
        return pulumi.get(self, "account_description")

    @account_description.setter
    def account_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_description", value)

    @property
    @pulumi.getter(name="accountPrivileges")
    def account_privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LakeAccountAccountPrivilegeArgs']]]]:
        """
        List of permissions granted. See `account_privileges` below.
        """
        return pulumi.get(self, "account_privileges")

    @account_privileges.setter
    def account_privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LakeAccountAccountPrivilegeArgs']]]]):
        pulumi.set(self, "account_privileges", value)

    @property
    @pulumi.getter(name="accountType")
    def account_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the account.
        """
        return pulumi.get(self, "account_type")

    @account_type.setter
    def account_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_type", value)


@pulumi.input_type
class _LakeAccountState:
    def __init__(__self__, *,
                 account_description: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 account_privileges: Optional[pulumi.Input[Sequence[pulumi.Input['LakeAccountAccountPrivilegeArgs']]]] = None,
                 account_type: Optional[pulumi.Input[str]] = None,
                 db_cluster_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LakeAccount resources.
        :param pulumi.Input[str] account_description: The description of the account.
        :param pulumi.Input[str] account_name: The name of the account.
        :param pulumi.Input[str] account_password: AccountPassword.
        :param pulumi.Input[Sequence[pulumi.Input['LakeAccountAccountPrivilegeArgs']]] account_privileges: List of permissions granted. See `account_privileges` below.
        :param pulumi.Input[str] account_type: The type of the account.
        :param pulumi.Input[str] db_cluster_id: The DBCluster ID.
        :param pulumi.Input[str] status: The status of the resource.
        """
        if account_description is not None:
            pulumi.set(__self__, "account_description", account_description)
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if account_password is not None:
            pulumi.set(__self__, "account_password", account_password)
        if account_privileges is not None:
            pulumi.set(__self__, "account_privileges", account_privileges)
        if account_type is not None:
            pulumi.set(__self__, "account_type", account_type)
        if db_cluster_id is not None:
            pulumi.set(__self__, "db_cluster_id", db_cluster_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accountDescription")
    def account_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the account.
        """
        return pulumi.get(self, "account_description")

    @account_description.setter
    def account_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_description", value)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the account.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> Optional[pulumi.Input[str]]:
        """
        AccountPassword.
        """
        return pulumi.get(self, "account_password")

    @account_password.setter
    def account_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_password", value)

    @property
    @pulumi.getter(name="accountPrivileges")
    def account_privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LakeAccountAccountPrivilegeArgs']]]]:
        """
        List of permissions granted. See `account_privileges` below.
        """
        return pulumi.get(self, "account_privileges")

    @account_privileges.setter
    def account_privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LakeAccountAccountPrivilegeArgs']]]]):
        pulumi.set(self, "account_privileges", value)

    @property
    @pulumi.getter(name="accountType")
    def account_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the account.
        """
        return pulumi.get(self, "account_type")

    @account_type.setter
    def account_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_type", value)

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The DBCluster ID.
        """
        return pulumi.get(self, "db_cluster_id")

    @db_cluster_id.setter
    def db_cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_cluster_id", value)

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


class LakeAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_description: Optional[pulumi.Input[str]] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 account_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LakeAccountAccountPrivilegeArgs']]]]] = None,
                 account_type: Optional[pulumi.Input[str]] = None,
                 db_cluster_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ADB Lake Account resource. Account of the DBClusterLakeVesion.

        For information about ADB Lake Account and how to use it, see [What is Lake Account](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2021-12-01-createaccount).
        For information about ADB Lake Account Privileges and how to use it, see [What are Lake Account Privileges](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2021-12-01-modifyaccountprivileges/).

        > **NOTE:** Available since v1.214.0.

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
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        v_pcid = alicloud.vpc.Network("vPCID",
            vpc_name=name,
            cidr_block="172.16.0.0/12")
        v_switchid = alicloud.vpc.Switch("vSWITCHID",
            vpc_id=v_pcid.id,
            zone_id="cn-hangzhou-k",
            vswitch_name=name,
            cidr_block="172.16.0.0/24")
        create_instance = alicloud.adb.DBClusterLakeVersion("createInstance",
            storage_resource="0ACU",
            zone_id="cn-hangzhou-k",
            vpc_id=v_pcid.id,
            vswitch_id=v_switchid.id,
            db_cluster_description=name,
            compute_resource="16ACU",
            db_cluster_version="5.0",
            payment_type="PayAsYouGo",
            security_ips="127.0.0.1")
        default_lake_account = alicloud.adb.LakeAccount("defaultLakeAccount",
            db_cluster_id=create_instance.id,
            account_type="Super",
            account_name="tfnormal",
            account_password="normal@2023",
            account_privileges=[
                alicloud.adb.LakeAccountAccountPrivilegeArgs(
                    privilege_type="Database",
                    privilege_object=alicloud.adb.LakeAccountAccountPrivilegePrivilegeObjectArgs(
                        database="MYSQL",
                    ),
                    privileges=[
                        "select",
                        "update",
                    ],
                ),
                alicloud.adb.LakeAccountAccountPrivilegeArgs(
                    privilege_type="Table",
                    privilege_object=alicloud.adb.LakeAccountAccountPrivilegePrivilegeObjectArgs(
                        database="INFORMATION_SCHEMA",
                        table="ENGINES",
                    ),
                    privileges=["update"],
                ),
                alicloud.adb.LakeAccountAccountPrivilegeArgs(
                    privilege_type="Column",
                    privilege_object=alicloud.adb.LakeAccountAccountPrivilegePrivilegeObjectArgs(
                        table="COLUMNS",
                        column="PRIVILEGES",
                        database="INFORMATION_SCHEMA",
                    ),
                    privileges=["update"],
                ),
            ],
            account_description=name)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ADB Lake Account can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:adb/lakeAccount:LakeAccount example <db_cluster_id>:<account_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_description: The description of the account.
        :param pulumi.Input[str] account_name: The name of the account.
        :param pulumi.Input[str] account_password: AccountPassword.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LakeAccountAccountPrivilegeArgs']]]] account_privileges: List of permissions granted. See `account_privileges` below.
        :param pulumi.Input[str] account_type: The type of the account.
        :param pulumi.Input[str] db_cluster_id: The DBCluster ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LakeAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ADB Lake Account resource. Account of the DBClusterLakeVesion.

        For information about ADB Lake Account and how to use it, see [What is Lake Account](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2021-12-01-createaccount).
        For information about ADB Lake Account Privileges and how to use it, see [What are Lake Account Privileges](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2021-12-01-modifyaccountprivileges/).

        > **NOTE:** Available since v1.214.0.

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
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        v_pcid = alicloud.vpc.Network("vPCID",
            vpc_name=name,
            cidr_block="172.16.0.0/12")
        v_switchid = alicloud.vpc.Switch("vSWITCHID",
            vpc_id=v_pcid.id,
            zone_id="cn-hangzhou-k",
            vswitch_name=name,
            cidr_block="172.16.0.0/24")
        create_instance = alicloud.adb.DBClusterLakeVersion("createInstance",
            storage_resource="0ACU",
            zone_id="cn-hangzhou-k",
            vpc_id=v_pcid.id,
            vswitch_id=v_switchid.id,
            db_cluster_description=name,
            compute_resource="16ACU",
            db_cluster_version="5.0",
            payment_type="PayAsYouGo",
            security_ips="127.0.0.1")
        default_lake_account = alicloud.adb.LakeAccount("defaultLakeAccount",
            db_cluster_id=create_instance.id,
            account_type="Super",
            account_name="tfnormal",
            account_password="normal@2023",
            account_privileges=[
                alicloud.adb.LakeAccountAccountPrivilegeArgs(
                    privilege_type="Database",
                    privilege_object=alicloud.adb.LakeAccountAccountPrivilegePrivilegeObjectArgs(
                        database="MYSQL",
                    ),
                    privileges=[
                        "select",
                        "update",
                    ],
                ),
                alicloud.adb.LakeAccountAccountPrivilegeArgs(
                    privilege_type="Table",
                    privilege_object=alicloud.adb.LakeAccountAccountPrivilegePrivilegeObjectArgs(
                        database="INFORMATION_SCHEMA",
                        table="ENGINES",
                    ),
                    privileges=["update"],
                ),
                alicloud.adb.LakeAccountAccountPrivilegeArgs(
                    privilege_type="Column",
                    privilege_object=alicloud.adb.LakeAccountAccountPrivilegePrivilegeObjectArgs(
                        table="COLUMNS",
                        column="PRIVILEGES",
                        database="INFORMATION_SCHEMA",
                    ),
                    privileges=["update"],
                ),
            ],
            account_description=name)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ADB Lake Account can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:adb/lakeAccount:LakeAccount example <db_cluster_id>:<account_name>
        ```

        :param str resource_name: The name of the resource.
        :param LakeAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LakeAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
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
                 account_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LakeAccountAccountPrivilegeArgs']]]]] = None,
                 account_type: Optional[pulumi.Input[str]] = None,
                 db_cluster_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LakeAccountArgs.__new__(LakeAccountArgs)

            __props__.__dict__["account_description"] = account_description
            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            if account_password is None and not opts.urn:
                raise TypeError("Missing required property 'account_password'")
            __props__.__dict__["account_password"] = None if account_password is None else pulumi.Output.secret(account_password)
            __props__.__dict__["account_privileges"] = account_privileges
            __props__.__dict__["account_type"] = account_type
            if db_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_cluster_id'")
            __props__.__dict__["db_cluster_id"] = db_cluster_id
            __props__.__dict__["status"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accountPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(LakeAccount, __self__).__init__(
            'alicloud:adb/lakeAccount:LakeAccount',
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
            account_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LakeAccountAccountPrivilegeArgs']]]]] = None,
            account_type: Optional[pulumi.Input[str]] = None,
            db_cluster_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'LakeAccount':
        """
        Get an existing LakeAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_description: The description of the account.
        :param pulumi.Input[str] account_name: The name of the account.
        :param pulumi.Input[str] account_password: AccountPassword.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LakeAccountAccountPrivilegeArgs']]]] account_privileges: List of permissions granted. See `account_privileges` below.
        :param pulumi.Input[str] account_type: The type of the account.
        :param pulumi.Input[str] db_cluster_id: The DBCluster ID.
        :param pulumi.Input[str] status: The status of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LakeAccountState.__new__(_LakeAccountState)

        __props__.__dict__["account_description"] = account_description
        __props__.__dict__["account_name"] = account_name
        __props__.__dict__["account_password"] = account_password
        __props__.__dict__["account_privileges"] = account_privileges
        __props__.__dict__["account_type"] = account_type
        __props__.__dict__["db_cluster_id"] = db_cluster_id
        __props__.__dict__["status"] = status
        return LakeAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountDescription")
    def account_description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the account.
        """
        return pulumi.get(self, "account_description")

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        The name of the account.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> pulumi.Output[str]:
        """
        AccountPassword.
        """
        return pulumi.get(self, "account_password")

    @property
    @pulumi.getter(name="accountPrivileges")
    def account_privileges(self) -> pulumi.Output[Sequence['outputs.LakeAccountAccountPrivilege']]:
        """
        List of permissions granted. See `account_privileges` below.
        """
        return pulumi.get(self, "account_privileges")

    @property
    @pulumi.getter(name="accountType")
    def account_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the account.
        """
        return pulumi.get(self, "account_type")

    @property
    @pulumi.getter(name="dbClusterId")
    def db_cluster_id(self) -> pulumi.Output[str]:
        """
        The DBCluster ID.
        """
        return pulumi.get(self, "db_cluster_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

