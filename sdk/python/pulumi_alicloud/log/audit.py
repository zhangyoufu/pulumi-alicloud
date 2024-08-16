# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AuditArgs', 'Audit']

@pulumi.input_type
class AuditArgs:
    def __init__(__self__, *,
                 aliuid: pulumi.Input[str],
                 display_name: pulumi.Input[str],
                 multi_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_directory_type: Optional[pulumi.Input[str]] = None,
                 variable_map: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Audit resource.
        :param pulumi.Input[str] aliuid: Aliuid value of your account.
        :param pulumi.Input[str] display_name: Name of SLS log audit.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] multi_accounts: Multi-account configuration, please fill in multiple aliuid.
        :param pulumi.Input[str] resource_directory_type: Resource Directory type. Optional values are all or custom. If the value is custom, argument multi_account should be provided.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variable_map: Log audit detailed configuration.
        """
        pulumi.set(__self__, "aliuid", aliuid)
        pulumi.set(__self__, "display_name", display_name)
        if multi_accounts is not None:
            pulumi.set(__self__, "multi_accounts", multi_accounts)
        if resource_directory_type is not None:
            pulumi.set(__self__, "resource_directory_type", resource_directory_type)
        if variable_map is not None:
            pulumi.set(__self__, "variable_map", variable_map)

    @property
    @pulumi.getter
    def aliuid(self) -> pulumi.Input[str]:
        """
        Aliuid value of your account.
        """
        return pulumi.get(self, "aliuid")

    @aliuid.setter
    def aliuid(self, value: pulumi.Input[str]):
        pulumi.set(self, "aliuid", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        Name of SLS log audit.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="multiAccounts")
    def multi_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Multi-account configuration, please fill in multiple aliuid.
        """
        return pulumi.get(self, "multi_accounts")

    @multi_accounts.setter
    def multi_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "multi_accounts", value)

    @property
    @pulumi.getter(name="resourceDirectoryType")
    def resource_directory_type(self) -> Optional[pulumi.Input[str]]:
        """
        Resource Directory type. Optional values are all or custom. If the value is custom, argument multi_account should be provided.
        """
        return pulumi.get(self, "resource_directory_type")

    @resource_directory_type.setter
    def resource_directory_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_directory_type", value)

    @property
    @pulumi.getter(name="variableMap")
    def variable_map(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Log audit detailed configuration.
        """
        return pulumi.get(self, "variable_map")

    @variable_map.setter
    def variable_map(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "variable_map", value)


@pulumi.input_type
class _AuditState:
    def __init__(__self__, *,
                 aliuid: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 multi_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_directory_type: Optional[pulumi.Input[str]] = None,
                 variable_map: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Audit resources.
        :param pulumi.Input[str] aliuid: Aliuid value of your account.
        :param pulumi.Input[str] display_name: Name of SLS log audit.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] multi_accounts: Multi-account configuration, please fill in multiple aliuid.
        :param pulumi.Input[str] resource_directory_type: Resource Directory type. Optional values are all or custom. If the value is custom, argument multi_account should be provided.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variable_map: Log audit detailed configuration.
        """
        if aliuid is not None:
            pulumi.set(__self__, "aliuid", aliuid)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if multi_accounts is not None:
            pulumi.set(__self__, "multi_accounts", multi_accounts)
        if resource_directory_type is not None:
            pulumi.set(__self__, "resource_directory_type", resource_directory_type)
        if variable_map is not None:
            pulumi.set(__self__, "variable_map", variable_map)

    @property
    @pulumi.getter
    def aliuid(self) -> Optional[pulumi.Input[str]]:
        """
        Aliuid value of your account.
        """
        return pulumi.get(self, "aliuid")

    @aliuid.setter
    def aliuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aliuid", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of SLS log audit.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="multiAccounts")
    def multi_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Multi-account configuration, please fill in multiple aliuid.
        """
        return pulumi.get(self, "multi_accounts")

    @multi_accounts.setter
    def multi_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "multi_accounts", value)

    @property
    @pulumi.getter(name="resourceDirectoryType")
    def resource_directory_type(self) -> Optional[pulumi.Input[str]]:
        """
        Resource Directory type. Optional values are all or custom. If the value is custom, argument multi_account should be provided.
        """
        return pulumi.get(self, "resource_directory_type")

    @resource_directory_type.setter
    def resource_directory_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_directory_type", value)

    @property
    @pulumi.getter(name="variableMap")
    def variable_map(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Log audit detailed configuration.
        """
        return pulumi.get(self, "variable_map")

    @variable_map.setter
    def variable_map(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "variable_map", value)


class Audit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aliuid: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 multi_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_directory_type: Optional[pulumi.Input[str]] = None,
                 variable_map: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        SLS log audit exists in the form of log service app.

        In addition to inheriting all SLS functions, it also enhances the real-time automatic centralized collection of audit related logs across multi cloud products under multi accounts, and provides support for storage, query and information summary required by audit. It covers actiontrail, OSS, NAS, SLB, API gateway, RDS, WAF, cloud firewall, cloud security center and other products.

        > **NOTE:** Available since v1.81.0

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.get_account()
        example = alicloud.log.Audit("example",
            display_name="tf-audit-example",
            aliuid=default.id,
            variable_map={
                "actiontrail_enabled": "true",
                "actiontrail_ttl": "180",
                "oss_access_enabled": "true",
                "oss_access_ttl": "7",
                "oss_sync_enabled": "true",
                "oss_sync_ttl": "180",
                "oss_metering_enabled": "true",
                "oss_metering_ttl": "180",
                "rds_enabled": "true",
                "rds_audit_collection_policy": "",
                "rds_ttl": "180",
                "rds_slow_enabled": "false",
                "rds_slow_collection_policy": "",
                "rds_slow_ttl": "180",
                "rds_perf_enabled": "false",
                "rds_perf_collection_policy": "",
                "rds_perf_ttl": "180",
                "vpc_flow_enabled": "false",
                "vpc_flow_ttl": "7",
                "vpc_flow_collection_policy": "",
                "vpc_sync_enabled": "true",
                "vpc_sync_ttl": "180",
                "polardb_enabled": "true",
                "polardb_audit_collection_policy": "",
                "polardb_ttl": "180",
                "polardb_slow_enabled": "false",
                "polardb_slow_collection_policy": "",
                "polardb_slow_ttl": "180",
                "polardb_perf_enabled": "false",
                "polardb_perf_collection_policy": "",
                "polardb_perf_ttl": "180",
                "drds_audit_enabled": "true",
                "drds_audit_collection_policy": "",
                "drds_audit_ttl": "7",
                "drds_sync_enabled": "true",
                "drds_sync_ttl": "180",
                "slb_access_enabled": "true",
                "slb_access_collection_policy": "",
                "slb_access_ttl": "7",
                "slb_sync_enabled": "true",
                "slb_sync_ttl": "180",
                "bastion_enabled": "true",
                "bastion_ttl": "180",
                "waf_enabled": "true",
                "waf_ttl": "180",
                "cloudfirewall_enabled": "true",
                "cloudfirewall_ttl": "180",
                "ddos_coo_access_enabled": "false",
                "ddos_coo_access_ttl": "180",
                "ddos_bgp_access_enabled": "false",
                "ddos_bgp_access_ttl": "180",
                "ddos_dip_access_enabled": "false",
                "ddos_dip_access_ttl": "180",
                "sas_crack_enabled": "true",
                "sas_dns_enabled": "true",
                "sas_http_enabled": "true",
                "sas_local_dns_enabled": "true",
                "sas_login_enabled": "true",
                "sas_network_enabled": "true",
                "sas_process_enabled": "true",
                "sas_security_alert_enabled": "true",
                "sas_security_hc_enabled": "true",
                "sas_security_vul_enabled": "true",
                "sas_session_enabled": "true",
                "sas_snapshot_account_enabled": "true",
                "sas_snapshot_port_enabled": "true",
                "sas_snapshot_process_enabled": "true",
                "sas_ttl": "180",
                "apigateway_enabled": "true",
                "apigateway_ttl": "180",
                "nas_enabled": "true",
                "nas_ttl": "180",
                "appconnect_enabled": "false",
                "appconnect_ttl": "180",
                "cps_enabled": "true",
                "cps_ttl": "180",
                "k8s_audit_enabled": "true",
                "k8s_audit_collection_policy": "",
                "k8s_audit_ttl": "180",
                "k8s_event_enabled": "true",
                "k8s_event_collection_policy": "",
                "k8s_event_ttl": "180",
                "k8s_ingress_enabled": "true",
                "k8s_ingress_collection_policy": "",
                "k8s_ingress_ttl": "180",
            })
        ```
        Multiple accounts Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.get_account()
        example = alicloud.log.Audit("example",
            display_name="tf-audit-example",
            aliuid=default.id,
            variable_map={
                "actiontrail_enabled": "true",
                "actiontrail_ttl": "180",
                "oss_access_enabled": "true",
                "oss_access_ttl": "180",
            },
            multi_accounts=[
                "123456789123",
                "12345678912300123",
            ])
        ```
        Resource Directory Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.get_account()
        example = alicloud.log.Audit("example",
            display_name="tf-audit-example",
            aliuid=default.id,
            variable_map={
                "actiontrail_enabled": "true",
                "actiontrail_ttl": "180",
                "oss_access_enabled": "true",
                "oss_access_ttl": "180",
            },
            resource_directory_type="all")
        ```
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.get_account()
        example = alicloud.log.Audit("example",
            display_name="tf-audit-example",
            aliuid=default.id,
            variable_map={
                "actiontrail_enabled": "true",
                "actiontrail_ttl": "180",
                "oss_access_enabled": "true",
                "oss_access_ttl": "180",
            },
            multi_accounts=[],
            resource_directory_type="custom")
        ```

        ## Import

        Log audit can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:log/audit:Audit example tf-audit-example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aliuid: Aliuid value of your account.
        :param pulumi.Input[str] display_name: Name of SLS log audit.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] multi_accounts: Multi-account configuration, please fill in multiple aliuid.
        :param pulumi.Input[str] resource_directory_type: Resource Directory type. Optional values are all or custom. If the value is custom, argument multi_account should be provided.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variable_map: Log audit detailed configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuditArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        SLS log audit exists in the form of log service app.

        In addition to inheriting all SLS functions, it also enhances the real-time automatic centralized collection of audit related logs across multi cloud products under multi accounts, and provides support for storage, query and information summary required by audit. It covers actiontrail, OSS, NAS, SLB, API gateway, RDS, WAF, cloud firewall, cloud security center and other products.

        > **NOTE:** Available since v1.81.0

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.get_account()
        example = alicloud.log.Audit("example",
            display_name="tf-audit-example",
            aliuid=default.id,
            variable_map={
                "actiontrail_enabled": "true",
                "actiontrail_ttl": "180",
                "oss_access_enabled": "true",
                "oss_access_ttl": "7",
                "oss_sync_enabled": "true",
                "oss_sync_ttl": "180",
                "oss_metering_enabled": "true",
                "oss_metering_ttl": "180",
                "rds_enabled": "true",
                "rds_audit_collection_policy": "",
                "rds_ttl": "180",
                "rds_slow_enabled": "false",
                "rds_slow_collection_policy": "",
                "rds_slow_ttl": "180",
                "rds_perf_enabled": "false",
                "rds_perf_collection_policy": "",
                "rds_perf_ttl": "180",
                "vpc_flow_enabled": "false",
                "vpc_flow_ttl": "7",
                "vpc_flow_collection_policy": "",
                "vpc_sync_enabled": "true",
                "vpc_sync_ttl": "180",
                "polardb_enabled": "true",
                "polardb_audit_collection_policy": "",
                "polardb_ttl": "180",
                "polardb_slow_enabled": "false",
                "polardb_slow_collection_policy": "",
                "polardb_slow_ttl": "180",
                "polardb_perf_enabled": "false",
                "polardb_perf_collection_policy": "",
                "polardb_perf_ttl": "180",
                "drds_audit_enabled": "true",
                "drds_audit_collection_policy": "",
                "drds_audit_ttl": "7",
                "drds_sync_enabled": "true",
                "drds_sync_ttl": "180",
                "slb_access_enabled": "true",
                "slb_access_collection_policy": "",
                "slb_access_ttl": "7",
                "slb_sync_enabled": "true",
                "slb_sync_ttl": "180",
                "bastion_enabled": "true",
                "bastion_ttl": "180",
                "waf_enabled": "true",
                "waf_ttl": "180",
                "cloudfirewall_enabled": "true",
                "cloudfirewall_ttl": "180",
                "ddos_coo_access_enabled": "false",
                "ddos_coo_access_ttl": "180",
                "ddos_bgp_access_enabled": "false",
                "ddos_bgp_access_ttl": "180",
                "ddos_dip_access_enabled": "false",
                "ddos_dip_access_ttl": "180",
                "sas_crack_enabled": "true",
                "sas_dns_enabled": "true",
                "sas_http_enabled": "true",
                "sas_local_dns_enabled": "true",
                "sas_login_enabled": "true",
                "sas_network_enabled": "true",
                "sas_process_enabled": "true",
                "sas_security_alert_enabled": "true",
                "sas_security_hc_enabled": "true",
                "sas_security_vul_enabled": "true",
                "sas_session_enabled": "true",
                "sas_snapshot_account_enabled": "true",
                "sas_snapshot_port_enabled": "true",
                "sas_snapshot_process_enabled": "true",
                "sas_ttl": "180",
                "apigateway_enabled": "true",
                "apigateway_ttl": "180",
                "nas_enabled": "true",
                "nas_ttl": "180",
                "appconnect_enabled": "false",
                "appconnect_ttl": "180",
                "cps_enabled": "true",
                "cps_ttl": "180",
                "k8s_audit_enabled": "true",
                "k8s_audit_collection_policy": "",
                "k8s_audit_ttl": "180",
                "k8s_event_enabled": "true",
                "k8s_event_collection_policy": "",
                "k8s_event_ttl": "180",
                "k8s_ingress_enabled": "true",
                "k8s_ingress_collection_policy": "",
                "k8s_ingress_ttl": "180",
            })
        ```
        Multiple accounts Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.get_account()
        example = alicloud.log.Audit("example",
            display_name="tf-audit-example",
            aliuid=default.id,
            variable_map={
                "actiontrail_enabled": "true",
                "actiontrail_ttl": "180",
                "oss_access_enabled": "true",
                "oss_access_ttl": "180",
            },
            multi_accounts=[
                "123456789123",
                "12345678912300123",
            ])
        ```
        Resource Directory Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.get_account()
        example = alicloud.log.Audit("example",
            display_name="tf-audit-example",
            aliuid=default.id,
            variable_map={
                "actiontrail_enabled": "true",
                "actiontrail_ttl": "180",
                "oss_access_enabled": "true",
                "oss_access_ttl": "180",
            },
            resource_directory_type="all")
        ```
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.get_account()
        example = alicloud.log.Audit("example",
            display_name="tf-audit-example",
            aliuid=default.id,
            variable_map={
                "actiontrail_enabled": "true",
                "actiontrail_ttl": "180",
                "oss_access_enabled": "true",
                "oss_access_ttl": "180",
            },
            multi_accounts=[],
            resource_directory_type="custom")
        ```

        ## Import

        Log audit can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:log/audit:Audit example tf-audit-example
        ```

        :param str resource_name: The name of the resource.
        :param AuditArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuditArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aliuid: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 multi_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_directory_type: Optional[pulumi.Input[str]] = None,
                 variable_map: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuditArgs.__new__(AuditArgs)

            if aliuid is None and not opts.urn:
                raise TypeError("Missing required property 'aliuid'")
            __props__.__dict__["aliuid"] = aliuid
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["multi_accounts"] = multi_accounts
            __props__.__dict__["resource_directory_type"] = resource_directory_type
            __props__.__dict__["variable_map"] = variable_map
        super(Audit, __self__).__init__(
            'alicloud:log/audit:Audit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aliuid: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            multi_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_directory_type: Optional[pulumi.Input[str]] = None,
            variable_map: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Audit':
        """
        Get an existing Audit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aliuid: Aliuid value of your account.
        :param pulumi.Input[str] display_name: Name of SLS log audit.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] multi_accounts: Multi-account configuration, please fill in multiple aliuid.
        :param pulumi.Input[str] resource_directory_type: Resource Directory type. Optional values are all or custom. If the value is custom, argument multi_account should be provided.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variable_map: Log audit detailed configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuditState.__new__(_AuditState)

        __props__.__dict__["aliuid"] = aliuid
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["multi_accounts"] = multi_accounts
        __props__.__dict__["resource_directory_type"] = resource_directory_type
        __props__.__dict__["variable_map"] = variable_map
        return Audit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def aliuid(self) -> pulumi.Output[str]:
        """
        Aliuid value of your account.
        """
        return pulumi.get(self, "aliuid")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Name of SLS log audit.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="multiAccounts")
    def multi_accounts(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Multi-account configuration, please fill in multiple aliuid.
        """
        return pulumi.get(self, "multi_accounts")

    @property
    @pulumi.getter(name="resourceDirectoryType")
    def resource_directory_type(self) -> pulumi.Output[Optional[str]]:
        """
        Resource Directory type. Optional values are all or custom. If the value is custom, argument multi_account should be provided.
        """
        return pulumi.get(self, "resource_directory_type")

    @property
    @pulumi.getter(name="variableMap")
    def variable_map(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Log audit detailed configuration.
        """
        return pulumi.get(self, "variable_map")

