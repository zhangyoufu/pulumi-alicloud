# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HanaInstanceArgs', 'HanaInstance']

@pulumi.input_type
class HanaInstanceArgs:
    def __init__(__self__, *,
                 vault_id: pulumi.Input[str],
                 alert_setting: Optional[pulumi.Input[str]] = None,
                 ecs_instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hana_name: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance_number: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 sid: Optional[pulumi.Input[str]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 validate_certificate: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a HanaInstance resource.
        :param pulumi.Input[str] vault_id: The ID of the backup vault.
        :param pulumi.Input[str] alert_setting: The alert settings. Valid value: `INHERITED`, which indicates that the backup client sends alert notifications in the same way as the backup vault.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ecs_instance_ids: The IDs of ECS instances that host the SAP HANA instance to be registered. HBR installs backup clients on the specified ECS instances.
        :param pulumi.Input[str] hana_name: The name of the SAP HANA instance.
        :param pulumi.Input[str] host: The private or internal IP address of the host where the primary node of the SAP HANA instance resides.
        :param pulumi.Input[int] instance_number: The instance number of the SAP HANA system.
        :param pulumi.Input[str] password: The password that is used to connect with the SAP HANA database.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] sid: The security identifier (SID) of the SAP HANA database.
        :param pulumi.Input[bool] use_ssl: Specifies whether to connect with the SAP HANA database over Secure Sockets Layer (SSL).
        :param pulumi.Input[str] user_name: The username of the SYSTEMDB database.
        :param pulumi.Input[bool] validate_certificate: Specifies whether to verify the SSL certificate of the SAP HANA database.
        """
        pulumi.set(__self__, "vault_id", vault_id)
        if alert_setting is not None:
            pulumi.set(__self__, "alert_setting", alert_setting)
        if ecs_instance_ids is not None:
            pulumi.set(__self__, "ecs_instance_ids", ecs_instance_ids)
        if hana_name is not None:
            pulumi.set(__self__, "hana_name", hana_name)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if instance_number is not None:
            pulumi.set(__self__, "instance_number", instance_number)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if sid is not None:
            pulumi.set(__self__, "sid", sid)
        if use_ssl is not None:
            pulumi.set(__self__, "use_ssl", use_ssl)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)
        if validate_certificate is not None:
            pulumi.set(__self__, "validate_certificate", validate_certificate)

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> pulumi.Input[str]:
        """
        The ID of the backup vault.
        """
        return pulumi.get(self, "vault_id")

    @vault_id.setter
    def vault_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vault_id", value)

    @property
    @pulumi.getter(name="alertSetting")
    def alert_setting(self) -> Optional[pulumi.Input[str]]:
        """
        The alert settings. Valid value: `INHERITED`, which indicates that the backup client sends alert notifications in the same way as the backup vault.
        """
        return pulumi.get(self, "alert_setting")

    @alert_setting.setter
    def alert_setting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_setting", value)

    @property
    @pulumi.getter(name="ecsInstanceIds")
    def ecs_instance_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of ECS instances that host the SAP HANA instance to be registered. HBR installs backup clients on the specified ECS instances.
        """
        return pulumi.get(self, "ecs_instance_ids")

    @ecs_instance_ids.setter
    def ecs_instance_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ecs_instance_ids", value)

    @property
    @pulumi.getter(name="hanaName")
    def hana_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SAP HANA instance.
        """
        return pulumi.get(self, "hana_name")

    @hana_name.setter
    def hana_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hana_name", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The private or internal IP address of the host where the primary node of the SAP HANA instance resides.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="instanceNumber")
    def instance_number(self) -> Optional[pulumi.Input[int]]:
        """
        The instance number of the SAP HANA system.
        """
        return pulumi.get(self, "instance_number")

    @instance_number.setter
    def instance_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_number", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password that is used to connect with the SAP HANA database.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def sid(self) -> Optional[pulumi.Input[str]]:
        """
        The security identifier (SID) of the SAP HANA database.
        """
        return pulumi.get(self, "sid")

    @sid.setter
    def sid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sid", value)

    @property
    @pulumi.getter(name="useSsl")
    def use_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to connect with the SAP HANA database over Secure Sockets Layer (SSL).
        """
        return pulumi.get(self, "use_ssl")

    @use_ssl.setter
    def use_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_ssl", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        The username of the SYSTEMDB database.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="validateCertificate")
    def validate_certificate(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to verify the SSL certificate of the SAP HANA database.
        """
        return pulumi.get(self, "validate_certificate")

    @validate_certificate.setter
    def validate_certificate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validate_certificate", value)


@pulumi.input_type
class _HanaInstanceState:
    def __init__(__self__, *,
                 alert_setting: Optional[pulumi.Input[str]] = None,
                 ecs_instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hana_instance_id: Optional[pulumi.Input[str]] = None,
                 hana_name: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance_number: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 sid: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 validate_certificate: Optional[pulumi.Input[bool]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HanaInstance resources.
        :param pulumi.Input[str] alert_setting: The alert settings. Valid value: `INHERITED`, which indicates that the backup client sends alert notifications in the same way as the backup vault.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ecs_instance_ids: The IDs of ECS instances that host the SAP HANA instance to be registered. HBR installs backup clients on the specified ECS instances.
        :param pulumi.Input[str] hana_instance_id: The id of the Hana Instance.
        :param pulumi.Input[str] hana_name: The name of the SAP HANA instance.
        :param pulumi.Input[str] host: The private or internal IP address of the host where the primary node of the SAP HANA instance resides.
        :param pulumi.Input[int] instance_number: The instance number of the SAP HANA system.
        :param pulumi.Input[str] password: The password that is used to connect with the SAP HANA database.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] sid: The security identifier (SID) of the SAP HANA database.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[bool] use_ssl: Specifies whether to connect with the SAP HANA database over Secure Sockets Layer (SSL).
        :param pulumi.Input[str] user_name: The username of the SYSTEMDB database.
        :param pulumi.Input[bool] validate_certificate: Specifies whether to verify the SSL certificate of the SAP HANA database.
        :param pulumi.Input[str] vault_id: The ID of the backup vault.
        """
        if alert_setting is not None:
            pulumi.set(__self__, "alert_setting", alert_setting)
        if ecs_instance_ids is not None:
            pulumi.set(__self__, "ecs_instance_ids", ecs_instance_ids)
        if hana_instance_id is not None:
            pulumi.set(__self__, "hana_instance_id", hana_instance_id)
        if hana_name is not None:
            pulumi.set(__self__, "hana_name", hana_name)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if instance_number is not None:
            pulumi.set(__self__, "instance_number", instance_number)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if sid is not None:
            pulumi.set(__self__, "sid", sid)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if use_ssl is not None:
            pulumi.set(__self__, "use_ssl", use_ssl)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)
        if validate_certificate is not None:
            pulumi.set(__self__, "validate_certificate", validate_certificate)
        if vault_id is not None:
            pulumi.set(__self__, "vault_id", vault_id)

    @property
    @pulumi.getter(name="alertSetting")
    def alert_setting(self) -> Optional[pulumi.Input[str]]:
        """
        The alert settings. Valid value: `INHERITED`, which indicates that the backup client sends alert notifications in the same way as the backup vault.
        """
        return pulumi.get(self, "alert_setting")

    @alert_setting.setter
    def alert_setting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_setting", value)

    @property
    @pulumi.getter(name="ecsInstanceIds")
    def ecs_instance_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of ECS instances that host the SAP HANA instance to be registered. HBR installs backup clients on the specified ECS instances.
        """
        return pulumi.get(self, "ecs_instance_ids")

    @ecs_instance_ids.setter
    def ecs_instance_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ecs_instance_ids", value)

    @property
    @pulumi.getter(name="hanaInstanceId")
    def hana_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the Hana Instance.
        """
        return pulumi.get(self, "hana_instance_id")

    @hana_instance_id.setter
    def hana_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hana_instance_id", value)

    @property
    @pulumi.getter(name="hanaName")
    def hana_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SAP HANA instance.
        """
        return pulumi.get(self, "hana_name")

    @hana_name.setter
    def hana_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hana_name", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The private or internal IP address of the host where the primary node of the SAP HANA instance resides.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="instanceNumber")
    def instance_number(self) -> Optional[pulumi.Input[int]]:
        """
        The instance number of the SAP HANA system.
        """
        return pulumi.get(self, "instance_number")

    @instance_number.setter
    def instance_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_number", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password that is used to connect with the SAP HANA database.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def sid(self) -> Optional[pulumi.Input[str]]:
        """
        The security identifier (SID) of the SAP HANA database.
        """
        return pulumi.get(self, "sid")

    @sid.setter
    def sid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sid", value)

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
    @pulumi.getter(name="useSsl")
    def use_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to connect with the SAP HANA database over Secure Sockets Layer (SSL).
        """
        return pulumi.get(self, "use_ssl")

    @use_ssl.setter
    def use_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_ssl", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        The username of the SYSTEMDB database.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="validateCertificate")
    def validate_certificate(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to verify the SSL certificate of the SAP HANA database.
        """
        return pulumi.get(self, "validate_certificate")

    @validate_certificate.setter
    def validate_certificate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validate_certificate", value)

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the backup vault.
        """
        return pulumi.get(self, "vault_id")

    @vault_id.setter
    def vault_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_id", value)


class HanaInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_setting: Optional[pulumi.Input[str]] = None,
                 ecs_instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hana_name: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance_number: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 sid: Optional[pulumi.Input[str]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 validate_certificate: Optional[pulumi.Input[bool]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Hybrid Backup Recovery (HBR) Hana Instance resource.

        For information about Hybrid Backup Recovery (HBR) Hana Instance and how to use it, see [What is Hana Instance](https://www.alibabacloud.com/help/en/hybrid-backup-recovery/latest/api-hbr-2017-09-08-createhanainstance).

        > **NOTE:** Available in v1.178.0+.

        > **NOTE:** The `sid` attribute is required when destroying resources.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_resource_groups = alicloud.resourcemanager.get_resource_groups(status="OK")
        example_vault = alicloud.hbr.Vault("exampleVault", vault_name="terraform-example")
        example_hana_instance = alicloud.hbr.HanaInstance("exampleHanaInstance",
            alert_setting="INHERITED",
            hana_name="terraform-example",
            host="1.1.1.1",
            instance_number=1,
            password="YouPassword123",
            resource_group_id=example_resource_groups.groups[0].id,
            sid="HXE",
            use_ssl=False,
            user_name="admin",
            validate_certificate=False,
            vault_id=example_vault.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Hybrid Backup Recovery (HBR) Hana Instance can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:hbr/hanaInstance:HanaInstance example <vault_id>:<hana_instance_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_setting: The alert settings. Valid value: `INHERITED`, which indicates that the backup client sends alert notifications in the same way as the backup vault.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ecs_instance_ids: The IDs of ECS instances that host the SAP HANA instance to be registered. HBR installs backup clients on the specified ECS instances.
        :param pulumi.Input[str] hana_name: The name of the SAP HANA instance.
        :param pulumi.Input[str] host: The private or internal IP address of the host where the primary node of the SAP HANA instance resides.
        :param pulumi.Input[int] instance_number: The instance number of the SAP HANA system.
        :param pulumi.Input[str] password: The password that is used to connect with the SAP HANA database.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] sid: The security identifier (SID) of the SAP HANA database.
        :param pulumi.Input[bool] use_ssl: Specifies whether to connect with the SAP HANA database over Secure Sockets Layer (SSL).
        :param pulumi.Input[str] user_name: The username of the SYSTEMDB database.
        :param pulumi.Input[bool] validate_certificate: Specifies whether to verify the SSL certificate of the SAP HANA database.
        :param pulumi.Input[str] vault_id: The ID of the backup vault.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HanaInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Hybrid Backup Recovery (HBR) Hana Instance resource.

        For information about Hybrid Backup Recovery (HBR) Hana Instance and how to use it, see [What is Hana Instance](https://www.alibabacloud.com/help/en/hybrid-backup-recovery/latest/api-hbr-2017-09-08-createhanainstance).

        > **NOTE:** Available in v1.178.0+.

        > **NOTE:** The `sid` attribute is required when destroying resources.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_resource_groups = alicloud.resourcemanager.get_resource_groups(status="OK")
        example_vault = alicloud.hbr.Vault("exampleVault", vault_name="terraform-example")
        example_hana_instance = alicloud.hbr.HanaInstance("exampleHanaInstance",
            alert_setting="INHERITED",
            hana_name="terraform-example",
            host="1.1.1.1",
            instance_number=1,
            password="YouPassword123",
            resource_group_id=example_resource_groups.groups[0].id,
            sid="HXE",
            use_ssl=False,
            user_name="admin",
            validate_certificate=False,
            vault_id=example_vault.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Hybrid Backup Recovery (HBR) Hana Instance can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:hbr/hanaInstance:HanaInstance example <vault_id>:<hana_instance_id>
        ```

        :param str resource_name: The name of the resource.
        :param HanaInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HanaInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_setting: Optional[pulumi.Input[str]] = None,
                 ecs_instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hana_name: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance_number: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 sid: Optional[pulumi.Input[str]] = None,
                 use_ssl: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 validate_certificate: Optional[pulumi.Input[bool]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HanaInstanceArgs.__new__(HanaInstanceArgs)

            __props__.__dict__["alert_setting"] = alert_setting
            __props__.__dict__["ecs_instance_ids"] = ecs_instance_ids
            __props__.__dict__["hana_name"] = hana_name
            __props__.__dict__["host"] = host
            __props__.__dict__["instance_number"] = instance_number
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["sid"] = sid
            __props__.__dict__["use_ssl"] = use_ssl
            __props__.__dict__["user_name"] = user_name
            __props__.__dict__["validate_certificate"] = validate_certificate
            if vault_id is None and not opts.urn:
                raise TypeError("Missing required property 'vault_id'")
            __props__.__dict__["vault_id"] = vault_id
            __props__.__dict__["hana_instance_id"] = None
            __props__.__dict__["status"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(HanaInstance, __self__).__init__(
            'alicloud:hbr/hanaInstance:HanaInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_setting: Optional[pulumi.Input[str]] = None,
            ecs_instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            hana_instance_id: Optional[pulumi.Input[str]] = None,
            hana_name: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            instance_number: Optional[pulumi.Input[int]] = None,
            password: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            sid: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            use_ssl: Optional[pulumi.Input[bool]] = None,
            user_name: Optional[pulumi.Input[str]] = None,
            validate_certificate: Optional[pulumi.Input[bool]] = None,
            vault_id: Optional[pulumi.Input[str]] = None) -> 'HanaInstance':
        """
        Get an existing HanaInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_setting: The alert settings. Valid value: `INHERITED`, which indicates that the backup client sends alert notifications in the same way as the backup vault.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ecs_instance_ids: The IDs of ECS instances that host the SAP HANA instance to be registered. HBR installs backup clients on the specified ECS instances.
        :param pulumi.Input[str] hana_instance_id: The id of the Hana Instance.
        :param pulumi.Input[str] hana_name: The name of the SAP HANA instance.
        :param pulumi.Input[str] host: The private or internal IP address of the host where the primary node of the SAP HANA instance resides.
        :param pulumi.Input[int] instance_number: The instance number of the SAP HANA system.
        :param pulumi.Input[str] password: The password that is used to connect with the SAP HANA database.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] sid: The security identifier (SID) of the SAP HANA database.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[bool] use_ssl: Specifies whether to connect with the SAP HANA database over Secure Sockets Layer (SSL).
        :param pulumi.Input[str] user_name: The username of the SYSTEMDB database.
        :param pulumi.Input[bool] validate_certificate: Specifies whether to verify the SSL certificate of the SAP HANA database.
        :param pulumi.Input[str] vault_id: The ID of the backup vault.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HanaInstanceState.__new__(_HanaInstanceState)

        __props__.__dict__["alert_setting"] = alert_setting
        __props__.__dict__["ecs_instance_ids"] = ecs_instance_ids
        __props__.__dict__["hana_instance_id"] = hana_instance_id
        __props__.__dict__["hana_name"] = hana_name
        __props__.__dict__["host"] = host
        __props__.__dict__["instance_number"] = instance_number
        __props__.__dict__["password"] = password
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["sid"] = sid
        __props__.__dict__["status"] = status
        __props__.__dict__["use_ssl"] = use_ssl
        __props__.__dict__["user_name"] = user_name
        __props__.__dict__["validate_certificate"] = validate_certificate
        __props__.__dict__["vault_id"] = vault_id
        return HanaInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertSetting")
    def alert_setting(self) -> pulumi.Output[str]:
        """
        The alert settings. Valid value: `INHERITED`, which indicates that the backup client sends alert notifications in the same way as the backup vault.
        """
        return pulumi.get(self, "alert_setting")

    @property
    @pulumi.getter(name="ecsInstanceIds")
    def ecs_instance_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IDs of ECS instances that host the SAP HANA instance to be registered. HBR installs backup clients on the specified ECS instances.
        """
        return pulumi.get(self, "ecs_instance_ids")

    @property
    @pulumi.getter(name="hanaInstanceId")
    def hana_instance_id(self) -> pulumi.Output[str]:
        """
        The id of the Hana Instance.
        """
        return pulumi.get(self, "hana_instance_id")

    @property
    @pulumi.getter(name="hanaName")
    def hana_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the SAP HANA instance.
        """
        return pulumi.get(self, "hana_name")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[str]]:
        """
        The private or internal IP address of the host where the primary node of the SAP HANA instance resides.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="instanceNumber")
    def instance_number(self) -> pulumi.Output[Optional[int]]:
        """
        The instance number of the SAP HANA system.
        """
        return pulumi.get(self, "instance_number")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password that is used to connect with the SAP HANA database.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def sid(self) -> pulumi.Output[Optional[str]]:
        """
        The security identifier (SID) of the SAP HANA database.
        """
        return pulumi.get(self, "sid")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="useSsl")
    def use_ssl(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to connect with the SAP HANA database over Secure Sockets Layer (SSL).
        """
        return pulumi.get(self, "use_ssl")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[Optional[str]]:
        """
        The username of the SYSTEMDB database.
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="validateCertificate")
    def validate_certificate(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to verify the SSL certificate of the SAP HANA database.
        """
        return pulumi.get(self, "validate_certificate")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> pulumi.Output[str]:
        """
        The ID of the backup vault.
        """
        return pulumi.get(self, "vault_id")

