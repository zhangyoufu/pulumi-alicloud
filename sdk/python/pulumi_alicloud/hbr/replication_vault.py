# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReplicationVaultArgs', 'ReplicationVault']

@pulumi.input_type
class ReplicationVaultArgs:
    def __init__(__self__, *,
                 replication_source_region_id: pulumi.Input[str],
                 replication_source_vault_id: pulumi.Input[str],
                 vault_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 vault_storage_class: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReplicationVault resource.
        :param pulumi.Input[str] replication_source_region_id: The ID of the region where the source vault resides.
        :param pulumi.Input[str] replication_source_vault_id: The ID of the source vault.
        :param pulumi.Input[str] vault_name: The name of the backup vault. The name must be 1 to 64 characters in length.
        :param pulumi.Input[str] description: The description of the backup vault. The description must be 0 to 255 characters in length.
        :param pulumi.Input[str] vault_storage_class: The storage type of the backup vault. Valid values: `STANDARD`.
        """
        pulumi.set(__self__, "replication_source_region_id", replication_source_region_id)
        pulumi.set(__self__, "replication_source_vault_id", replication_source_vault_id)
        pulumi.set(__self__, "vault_name", vault_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if vault_storage_class is not None:
            pulumi.set(__self__, "vault_storage_class", vault_storage_class)

    @property
    @pulumi.getter(name="replicationSourceRegionId")
    def replication_source_region_id(self) -> pulumi.Input[str]:
        """
        The ID of the region where the source vault resides.
        """
        return pulumi.get(self, "replication_source_region_id")

    @replication_source_region_id.setter
    def replication_source_region_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "replication_source_region_id", value)

    @property
    @pulumi.getter(name="replicationSourceVaultId")
    def replication_source_vault_id(self) -> pulumi.Input[str]:
        """
        The ID of the source vault.
        """
        return pulumi.get(self, "replication_source_vault_id")

    @replication_source_vault_id.setter
    def replication_source_vault_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "replication_source_vault_id", value)

    @property
    @pulumi.getter(name="vaultName")
    def vault_name(self) -> pulumi.Input[str]:
        """
        The name of the backup vault. The name must be 1 to 64 characters in length.
        """
        return pulumi.get(self, "vault_name")

    @vault_name.setter
    def vault_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vault_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the backup vault. The description must be 0 to 255 characters in length.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="vaultStorageClass")
    def vault_storage_class(self) -> Optional[pulumi.Input[str]]:
        """
        The storage type of the backup vault. Valid values: `STANDARD`.
        """
        return pulumi.get(self, "vault_storage_class")

    @vault_storage_class.setter
    def vault_storage_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_storage_class", value)


@pulumi.input_type
class _ReplicationVaultState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 replication_source_region_id: Optional[pulumi.Input[str]] = None,
                 replication_source_vault_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vault_name: Optional[pulumi.Input[str]] = None,
                 vault_storage_class: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReplicationVault resources.
        :param pulumi.Input[str] description: The description of the backup vault. The description must be 0 to 255 characters in length.
        :param pulumi.Input[str] replication_source_region_id: The ID of the region where the source vault resides.
        :param pulumi.Input[str] replication_source_vault_id: The ID of the source vault.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[str] vault_name: The name of the backup vault. The name must be 1 to 64 characters in length.
        :param pulumi.Input[str] vault_storage_class: The storage type of the backup vault. Valid values: `STANDARD`.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if replication_source_region_id is not None:
            pulumi.set(__self__, "replication_source_region_id", replication_source_region_id)
        if replication_source_vault_id is not None:
            pulumi.set(__self__, "replication_source_vault_id", replication_source_vault_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vault_name is not None:
            pulumi.set(__self__, "vault_name", vault_name)
        if vault_storage_class is not None:
            pulumi.set(__self__, "vault_storage_class", vault_storage_class)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the backup vault. The description must be 0 to 255 characters in length.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="replicationSourceRegionId")
    def replication_source_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the region where the source vault resides.
        """
        return pulumi.get(self, "replication_source_region_id")

    @replication_source_region_id.setter
    def replication_source_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_source_region_id", value)

    @property
    @pulumi.getter(name="replicationSourceVaultId")
    def replication_source_vault_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the source vault.
        """
        return pulumi.get(self, "replication_source_vault_id")

    @replication_source_vault_id.setter
    def replication_source_vault_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_source_vault_id", value)

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
    @pulumi.getter(name="vaultName")
    def vault_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the backup vault. The name must be 1 to 64 characters in length.
        """
        return pulumi.get(self, "vault_name")

    @vault_name.setter
    def vault_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_name", value)

    @property
    @pulumi.getter(name="vaultStorageClass")
    def vault_storage_class(self) -> Optional[pulumi.Input[str]]:
        """
        The storage type of the backup vault. Valid values: `STANDARD`.
        """
        return pulumi.get(self, "vault_storage_class")

    @vault_storage_class.setter
    def vault_storage_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_storage_class", value)


class ReplicationVault(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 replication_source_region_id: Optional[pulumi.Input[str]] = None,
                 replication_source_vault_id: Optional[pulumi.Input[str]] = None,
                 vault_name: Optional[pulumi.Input[str]] = None,
                 vault_storage_class: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Hybrid Backup Recovery (HBR) Replication Vault resource.

        For information about Hybrid Backup Recovery (HBR) Replication Vault and how to use it, see [What is Replication Vault](https://www.alibabacloud.com/help/en/doc-detail/345603.html).

        > **NOTE:** Available in v1.152.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_pulumi as pulumi

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-testAccReplicationVault"
        region_source = config.get("regionSource")
        if region_source is None:
            region_source = "you Replication value source region"
        source = pulumi.providers.Alicloud("source", region=region_source)
        default_vault = alicloud.hbr.Vault("defaultVault", vault_name=name,
        opts=pulumi.ResourceOptions(provider=alicloud["source"]))
        default_replication_vault_regions = alicloud.hbr.get_replication_vault_regions()
        region_replication = default_replication_vault_regions.regions[0].replication_region_id
        replication = pulumi.providers.Alicloud("replication", region=region_replication)
        default_replication_vault = alicloud.hbr.ReplicationVault("defaultReplicationVault",
            replication_source_region_id=region_replication,
            replication_source_vault_id=default_vault.id,
            vault_name=name,
            vault_storage_class="STANDARD",
            description=name,
            opts=pulumi.ResourceOptions(provider=alicloud["replication"]))
        ```

        ## Import

        Hybrid Backup Recovery (HBR) Replication Vault can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:hbr/replicationVault:ReplicationVault example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the backup vault. The description must be 0 to 255 characters in length.
        :param pulumi.Input[str] replication_source_region_id: The ID of the region where the source vault resides.
        :param pulumi.Input[str] replication_source_vault_id: The ID of the source vault.
        :param pulumi.Input[str] vault_name: The name of the backup vault. The name must be 1 to 64 characters in length.
        :param pulumi.Input[str] vault_storage_class: The storage type of the backup vault. Valid values: `STANDARD`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicationVaultArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Hybrid Backup Recovery (HBR) Replication Vault resource.

        For information about Hybrid Backup Recovery (HBR) Replication Vault and how to use it, see [What is Replication Vault](https://www.alibabacloud.com/help/en/doc-detail/345603.html).

        > **NOTE:** Available in v1.152.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_pulumi as pulumi

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-testAccReplicationVault"
        region_source = config.get("regionSource")
        if region_source is None:
            region_source = "you Replication value source region"
        source = pulumi.providers.Alicloud("source", region=region_source)
        default_vault = alicloud.hbr.Vault("defaultVault", vault_name=name,
        opts=pulumi.ResourceOptions(provider=alicloud["source"]))
        default_replication_vault_regions = alicloud.hbr.get_replication_vault_regions()
        region_replication = default_replication_vault_regions.regions[0].replication_region_id
        replication = pulumi.providers.Alicloud("replication", region=region_replication)
        default_replication_vault = alicloud.hbr.ReplicationVault("defaultReplicationVault",
            replication_source_region_id=region_replication,
            replication_source_vault_id=default_vault.id,
            vault_name=name,
            vault_storage_class="STANDARD",
            description=name,
            opts=pulumi.ResourceOptions(provider=alicloud["replication"]))
        ```

        ## Import

        Hybrid Backup Recovery (HBR) Replication Vault can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:hbr/replicationVault:ReplicationVault example <id>
        ```

        :param str resource_name: The name of the resource.
        :param ReplicationVaultArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicationVaultArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 replication_source_region_id: Optional[pulumi.Input[str]] = None,
                 replication_source_vault_id: Optional[pulumi.Input[str]] = None,
                 vault_name: Optional[pulumi.Input[str]] = None,
                 vault_storage_class: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReplicationVaultArgs.__new__(ReplicationVaultArgs)

            __props__.__dict__["description"] = description
            if replication_source_region_id is None and not opts.urn:
                raise TypeError("Missing required property 'replication_source_region_id'")
            __props__.__dict__["replication_source_region_id"] = replication_source_region_id
            if replication_source_vault_id is None and not opts.urn:
                raise TypeError("Missing required property 'replication_source_vault_id'")
            __props__.__dict__["replication_source_vault_id"] = replication_source_vault_id
            if vault_name is None and not opts.urn:
                raise TypeError("Missing required property 'vault_name'")
            __props__.__dict__["vault_name"] = vault_name
            __props__.__dict__["vault_storage_class"] = vault_storage_class
            __props__.__dict__["status"] = None
        super(ReplicationVault, __self__).__init__(
            'alicloud:hbr/replicationVault:ReplicationVault',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            replication_source_region_id: Optional[pulumi.Input[str]] = None,
            replication_source_vault_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vault_name: Optional[pulumi.Input[str]] = None,
            vault_storage_class: Optional[pulumi.Input[str]] = None) -> 'ReplicationVault':
        """
        Get an existing ReplicationVault resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the backup vault. The description must be 0 to 255 characters in length.
        :param pulumi.Input[str] replication_source_region_id: The ID of the region where the source vault resides.
        :param pulumi.Input[str] replication_source_vault_id: The ID of the source vault.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[str] vault_name: The name of the backup vault. The name must be 1 to 64 characters in length.
        :param pulumi.Input[str] vault_storage_class: The storage type of the backup vault. Valid values: `STANDARD`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReplicationVaultState.__new__(_ReplicationVaultState)

        __props__.__dict__["description"] = description
        __props__.__dict__["replication_source_region_id"] = replication_source_region_id
        __props__.__dict__["replication_source_vault_id"] = replication_source_vault_id
        __props__.__dict__["status"] = status
        __props__.__dict__["vault_name"] = vault_name
        __props__.__dict__["vault_storage_class"] = vault_storage_class
        return ReplicationVault(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the backup vault. The description must be 0 to 255 characters in length.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="replicationSourceRegionId")
    def replication_source_region_id(self) -> pulumi.Output[str]:
        """
        The ID of the region where the source vault resides.
        """
        return pulumi.get(self, "replication_source_region_id")

    @property
    @pulumi.getter(name="replicationSourceVaultId")
    def replication_source_vault_id(self) -> pulumi.Output[str]:
        """
        The ID of the source vault.
        """
        return pulumi.get(self, "replication_source_vault_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vaultName")
    def vault_name(self) -> pulumi.Output[str]:
        """
        The name of the backup vault. The name must be 1 to 64 characters in length.
        """
        return pulumi.get(self, "vault_name")

    @property
    @pulumi.getter(name="vaultStorageClass")
    def vault_storage_class(self) -> pulumi.Output[str]:
        """
        The storage type of the backup vault. Valid values: `STANDARD`.
        """
        return pulumi.get(self, "vault_storage_class")

