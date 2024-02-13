# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HanaBackupPlanArgs', 'HanaBackupPlan']

@pulumi.input_type
class HanaBackupPlanArgs:
    def __init__(__self__, *,
                 backup_type: pulumi.Input[str],
                 cluster_id: pulumi.Input[str],
                 database_name: pulumi.Input[str],
                 plan_name: pulumi.Input[str],
                 schedule: pulumi.Input[str],
                 vault_id: pulumi.Input[str],
                 backup_prefix: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HanaBackupPlan resource.
        :param pulumi.Input[str] backup_type: The backup type. Valid values:
        :param pulumi.Input[str] cluster_id: The ID of the SAP HANA instance.
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] plan_name: The name of the backup plan.
        :param pulumi.Input[str] schedule: The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
        :param pulumi.Input[str] vault_id: The ID of the backup vault.
        :param pulumi.Input[str] backup_prefix: The backup prefix.
        :param pulumi.Input[str] resource_group_id: The resource attribute field that represents the resource group ID.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `Enabled`, `Disabled`.
        """
        pulumi.set(__self__, "backup_type", backup_type)
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "plan_name", plan_name)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "vault_id", vault_id)
        if backup_prefix is not None:
            pulumi.set(__self__, "backup_prefix", backup_prefix)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> pulumi.Input[str]:
        """
        The backup type. Valid values:
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The ID of the SAP HANA instance.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        The name of the database.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="planName")
    def plan_name(self) -> pulumi.Input[str]:
        """
        The name of the backup plan.
        """
        return pulumi.get(self, "plan_name")

    @plan_name.setter
    def plan_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "plan_name", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input[str]:
        """
        The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule", value)

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
    @pulumi.getter(name="backupPrefix")
    def backup_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The backup prefix.
        """
        return pulumi.get(self, "backup_prefix")

    @backup_prefix.setter
    def backup_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_prefix", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource attribute field that represents the resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource. Valid values: `Enabled`, `Disabled`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _HanaBackupPlanState:
    def __init__(__self__, *,
                 backup_prefix: Optional[pulumi.Input[str]] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 plan_id: Optional[pulumi.Input[str]] = None,
                 plan_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HanaBackupPlan resources.
        :param pulumi.Input[str] backup_prefix: The backup prefix.
        :param pulumi.Input[str] backup_type: The backup type. Valid values:
        :param pulumi.Input[str] cluster_id: The ID of the SAP HANA instance.
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] plan_id: The id of the plan.
        :param pulumi.Input[str] plan_name: The name of the backup plan.
        :param pulumi.Input[str] resource_group_id: The resource attribute field that represents the resource group ID.
        :param pulumi.Input[str] schedule: The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `Enabled`, `Disabled`.
        :param pulumi.Input[str] vault_id: The ID of the backup vault.
        """
        if backup_prefix is not None:
            pulumi.set(__self__, "backup_prefix", backup_prefix)
        if backup_type is not None:
            pulumi.set(__self__, "backup_type", backup_type)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if plan_id is not None:
            pulumi.set(__self__, "plan_id", plan_id)
        if plan_name is not None:
            pulumi.set(__self__, "plan_name", plan_name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vault_id is not None:
            pulumi.set(__self__, "vault_id", vault_id)

    @property
    @pulumi.getter(name="backupPrefix")
    def backup_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The backup prefix.
        """
        return pulumi.get(self, "backup_prefix")

    @backup_prefix.setter
    def backup_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_prefix", value)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> Optional[pulumi.Input[str]]:
        """
        The backup type. Valid values:
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the SAP HANA instance.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the database.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the plan.
        """
        return pulumi.get(self, "plan_id")

    @plan_id.setter
    def plan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_id", value)

    @property
    @pulumi.getter(name="planName")
    def plan_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the backup plan.
        """
        return pulumi.get(self, "plan_name")

    @plan_name.setter
    def plan_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource attribute field that represents the resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource. Valid values: `Enabled`, `Disabled`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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


class HanaBackupPlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_prefix: Optional[pulumi.Input[str]] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 plan_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Hybrid Backup Recovery (HBR) Hana Backup Plan resource.

        For information about Hybrid Backup Recovery (HBR) Hana Backup Plan and how to use it, see [What is Hana Backup Plan](https://www.alibabacloud.com/help/en/hybrid-backup-recovery/latest/api-hbr-2017-09-08-createhanabackupplan).

        > **NOTE:** Available in v1.179.0+.

        ## Example Usage

        Basic Usage

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
        example_hana_backup_plan = alicloud.hbr.HanaBackupPlan("exampleHanaBackupPlan",
            backup_prefix="DIFF_DATA_BACKUP",
            backup_type="COMPLETE",
            cluster_id=example_hana_instance.hana_instance_id,
            database_name="SYSTEMDB",
            plan_name="terraform-example",
            resource_group_id=example_resource_groups.groups[0].id,
            schedule="I|1602673264|P1D",
            vault_id=example_hana_instance.vault_id)
        ```

        ## Import

        Hybrid Backup Recovery (HBR) Hana Backup Plan can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:hbr/hanaBackupPlan:HanaBackupPlan example <plan_id>:<vault_id>:<cluster_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_prefix: The backup prefix.
        :param pulumi.Input[str] backup_type: The backup type. Valid values:
        :param pulumi.Input[str] cluster_id: The ID of the SAP HANA instance.
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] plan_name: The name of the backup plan.
        :param pulumi.Input[str] resource_group_id: The resource attribute field that represents the resource group ID.
        :param pulumi.Input[str] schedule: The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `Enabled`, `Disabled`.
        :param pulumi.Input[str] vault_id: The ID of the backup vault.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HanaBackupPlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Hybrid Backup Recovery (HBR) Hana Backup Plan resource.

        For information about Hybrid Backup Recovery (HBR) Hana Backup Plan and how to use it, see [What is Hana Backup Plan](https://www.alibabacloud.com/help/en/hybrid-backup-recovery/latest/api-hbr-2017-09-08-createhanabackupplan).

        > **NOTE:** Available in v1.179.0+.

        ## Example Usage

        Basic Usage

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
        example_hana_backup_plan = alicloud.hbr.HanaBackupPlan("exampleHanaBackupPlan",
            backup_prefix="DIFF_DATA_BACKUP",
            backup_type="COMPLETE",
            cluster_id=example_hana_instance.hana_instance_id,
            database_name="SYSTEMDB",
            plan_name="terraform-example",
            resource_group_id=example_resource_groups.groups[0].id,
            schedule="I|1602673264|P1D",
            vault_id=example_hana_instance.vault_id)
        ```

        ## Import

        Hybrid Backup Recovery (HBR) Hana Backup Plan can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:hbr/hanaBackupPlan:HanaBackupPlan example <plan_id>:<vault_id>:<cluster_id>
        ```

        :param str resource_name: The name of the resource.
        :param HanaBackupPlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HanaBackupPlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_prefix: Optional[pulumi.Input[str]] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 plan_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vault_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HanaBackupPlanArgs.__new__(HanaBackupPlanArgs)

            __props__.__dict__["backup_prefix"] = backup_prefix
            if backup_type is None and not opts.urn:
                raise TypeError("Missing required property 'backup_type'")
            __props__.__dict__["backup_type"] = backup_type
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            if plan_name is None and not opts.urn:
                raise TypeError("Missing required property 'plan_name'")
            __props__.__dict__["plan_name"] = plan_name
            __props__.__dict__["resource_group_id"] = resource_group_id
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["status"] = status
            if vault_id is None and not opts.urn:
                raise TypeError("Missing required property 'vault_id'")
            __props__.__dict__["vault_id"] = vault_id
            __props__.__dict__["plan_id"] = None
        super(HanaBackupPlan, __self__).__init__(
            'alicloud:hbr/hanaBackupPlan:HanaBackupPlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_prefix: Optional[pulumi.Input[str]] = None,
            backup_type: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            plan_id: Optional[pulumi.Input[str]] = None,
            plan_name: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vault_id: Optional[pulumi.Input[str]] = None) -> 'HanaBackupPlan':
        """
        Get an existing HanaBackupPlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_prefix: The backup prefix.
        :param pulumi.Input[str] backup_type: The backup type. Valid values:
        :param pulumi.Input[str] cluster_id: The ID of the SAP HANA instance.
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] plan_id: The id of the plan.
        :param pulumi.Input[str] plan_name: The name of the backup plan.
        :param pulumi.Input[str] resource_group_id: The resource attribute field that represents the resource group ID.
        :param pulumi.Input[str] schedule: The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `Enabled`, `Disabled`.
        :param pulumi.Input[str] vault_id: The ID of the backup vault.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HanaBackupPlanState.__new__(_HanaBackupPlanState)

        __props__.__dict__["backup_prefix"] = backup_prefix
        __props__.__dict__["backup_type"] = backup_type
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["plan_id"] = plan_id
        __props__.__dict__["plan_name"] = plan_name
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["status"] = status
        __props__.__dict__["vault_id"] = vault_id
        return HanaBackupPlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupPrefix")
    def backup_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The backup prefix.
        """
        return pulumi.get(self, "backup_prefix")

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> pulumi.Output[str]:
        """
        The backup type. Valid values:
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The ID of the SAP HANA instance.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        The name of the database.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> pulumi.Output[str]:
        """
        The id of the plan.
        """
        return pulumi.get(self, "plan_id")

    @property
    @pulumi.getter(name="planName")
    def plan_name(self) -> pulumi.Output[str]:
        """
        The name of the backup plan.
        """
        return pulumi.get(self, "plan_name")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The resource attribute field that represents the resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[str]:
        """
        The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource. Valid values: `Enabled`, `Disabled`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> pulumi.Output[str]:
        """
        The ID of the backup vault.
        """
        return pulumi.get(self, "vault_id")

