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

__all__ = ['AclArgs', 'Acl']

@pulumi.input_type
class AclArgs:
    def __init__(__self__, *,
                 address_ip_version: pulumi.Input[str],
                 acl_entries: Optional[pulumi.Input[Sequence[pulumi.Input['AclAclEntryArgs']]]] = None,
                 acl_name: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Acl resource.
        :param pulumi.Input[str] address_ip_version: The IP version. Valid values: `IPv4` and `IPv6`.
        :param pulumi.Input[Sequence[pulumi.Input['AclAclEntryArgs']]] acl_entries: The entries of the Acl. See the following `Block acl_entries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
        :param pulumi.Input[str] acl_name: The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
        :param pulumi.Input[bool] dry_run: The dry run.
        """
        pulumi.set(__self__, "address_ip_version", address_ip_version)
        if acl_entries is not None:
            warnings.warn("""Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.""", DeprecationWarning)
            pulumi.log.warn("""acl_entries is deprecated: Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.""")
        if acl_entries is not None:
            pulumi.set(__self__, "acl_entries", acl_entries)
        if acl_name is not None:
            pulumi.set(__self__, "acl_name", acl_name)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)

    @property
    @pulumi.getter(name="addressIpVersion")
    def address_ip_version(self) -> pulumi.Input[str]:
        """
        The IP version. Valid values: `IPv4` and `IPv6`.
        """
        return pulumi.get(self, "address_ip_version")

    @address_ip_version.setter
    def address_ip_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "address_ip_version", value)

    @property
    @pulumi.getter(name="aclEntries")
    def acl_entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AclAclEntryArgs']]]]:
        """
        The entries of the Acl. See the following `Block acl_entries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
        """
        return pulumi.get(self, "acl_entries")

    @acl_entries.setter
    def acl_entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AclAclEntryArgs']]]]):
        pulumi.set(self, "acl_entries", value)

    @property
    @pulumi.getter(name="aclName")
    def acl_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
        """
        return pulumi.get(self, "acl_name")

    @acl_name.setter
    def acl_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "acl_name", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        The dry run.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)


@pulumi.input_type
class _AclState:
    def __init__(__self__, *,
                 acl_entries: Optional[pulumi.Input[Sequence[pulumi.Input['AclAclEntryArgs']]]] = None,
                 acl_name: Optional[pulumi.Input[str]] = None,
                 address_ip_version: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Acl resources.
        :param pulumi.Input[Sequence[pulumi.Input['AclAclEntryArgs']]] acl_entries: The entries of the Acl. See the following `Block acl_entries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
        :param pulumi.Input[str] acl_name: The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
        :param pulumi.Input[str] address_ip_version: The IP version. Valid values: `IPv4` and `IPv6`.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[str] status: The status of the resource.
        """
        if acl_entries is not None:
            warnings.warn("""Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.""", DeprecationWarning)
            pulumi.log.warn("""acl_entries is deprecated: Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.""")
        if acl_entries is not None:
            pulumi.set(__self__, "acl_entries", acl_entries)
        if acl_name is not None:
            pulumi.set(__self__, "acl_name", acl_name)
        if address_ip_version is not None:
            pulumi.set(__self__, "address_ip_version", address_ip_version)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="aclEntries")
    def acl_entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AclAclEntryArgs']]]]:
        """
        The entries of the Acl. See the following `Block acl_entries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
        """
        return pulumi.get(self, "acl_entries")

    @acl_entries.setter
    def acl_entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AclAclEntryArgs']]]]):
        pulumi.set(self, "acl_entries", value)

    @property
    @pulumi.getter(name="aclName")
    def acl_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
        """
        return pulumi.get(self, "acl_name")

    @acl_name.setter
    def acl_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "acl_name", value)

    @property
    @pulumi.getter(name="addressIpVersion")
    def address_ip_version(self) -> Optional[pulumi.Input[str]]:
        """
        The IP version. Valid values: `IPv4` and `IPv6`.
        """
        return pulumi.get(self, "address_ip_version")

    @address_ip_version.setter
    def address_ip_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_ip_version", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        The dry run.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

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


class Acl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclAclEntryArgs']]]]] = None,
                 acl_name: Optional[pulumi.Input[str]] = None,
                 address_ip_version: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a Global Accelerator (GA) Acl resource.

        For information about Global Accelerator (GA) Acl and how to use it, see [What is Acl](https://www.alibabacloud.com/help/en/doc-detail/258289.html).

        > **NOTE:** Available in v1.150.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.ga.Acl("default",
            acl_entries=[alicloud.ga.AclAclEntryArgs(
                entry="192.168.1.0/24",
                entry_description="tf-test1",
            )],
            acl_name="tf-testAccAcl",
            address_ip_version="IPv4")
        ```

        ## Import

        Global Accelerator (GA) Acl can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ga/acl:Acl example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclAclEntryArgs']]]] acl_entries: The entries of the Acl. See the following `Block acl_entries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
        :param pulumi.Input[str] acl_name: The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
        :param pulumi.Input[str] address_ip_version: The IP version. Valid values: `IPv4` and `IPv6`.
        :param pulumi.Input[bool] dry_run: The dry run.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Global Accelerator (GA) Acl resource.

        For information about Global Accelerator (GA) Acl and how to use it, see [What is Acl](https://www.alibabacloud.com/help/en/doc-detail/258289.html).

        > **NOTE:** Available in v1.150.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.ga.Acl("default",
            acl_entries=[alicloud.ga.AclAclEntryArgs(
                entry="192.168.1.0/24",
                entry_description="tf-test1",
            )],
            acl_name="tf-testAccAcl",
            address_ip_version="IPv4")
        ```

        ## Import

        Global Accelerator (GA) Acl can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ga/acl:Acl example <id>
        ```

        :param str resource_name: The name of the resource.
        :param AclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclAclEntryArgs']]]]] = None,
                 acl_name: Optional[pulumi.Input[str]] = None,
                 address_ip_version: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AclArgs.__new__(AclArgs)

            if acl_entries is not None and not opts.urn:
                warnings.warn("""Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.""", DeprecationWarning)
                pulumi.log.warn("""acl_entries is deprecated: Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource 'alicloud_ga_acl_entry_attachment'.""")
            __props__.__dict__["acl_entries"] = acl_entries
            __props__.__dict__["acl_name"] = acl_name
            if address_ip_version is None and not opts.urn:
                raise TypeError("Missing required property 'address_ip_version'")
            __props__.__dict__["address_ip_version"] = address_ip_version
            __props__.__dict__["dry_run"] = dry_run
            __props__.__dict__["status"] = None
        super(Acl, __self__).__init__(
            'alicloud:ga/acl:Acl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclAclEntryArgs']]]]] = None,
            acl_name: Optional[pulumi.Input[str]] = None,
            address_ip_version: Optional[pulumi.Input[str]] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Acl':
        """
        Get an existing Acl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AclAclEntryArgs']]]] acl_entries: The entries of the Acl. See the following `Block acl_entries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
        :param pulumi.Input[str] acl_name: The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
        :param pulumi.Input[str] address_ip_version: The IP version. Valid values: `IPv4` and `IPv6`.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[str] status: The status of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AclState.__new__(_AclState)

        __props__.__dict__["acl_entries"] = acl_entries
        __props__.__dict__["acl_name"] = acl_name
        __props__.__dict__["address_ip_version"] = address_ip_version
        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["status"] = status
        return Acl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aclEntries")
    def acl_entries(self) -> pulumi.Output[Sequence['outputs.AclAclEntry']]:
        """
        The entries of the Acl. See the following `Block acl_entries`. **NOTE:** "Field 'acl_entries' has been deprecated from provider version 1.190.0 and it will be removed in the future version. Please use the new resource `ga.AclEntryAttachment`."
        """
        return pulumi.get(self, "acl_entries")

    @property
    @pulumi.getter(name="aclName")
    def acl_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the ACL. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), hyphens (-) and underscores (_). It must start with a letter.
        """
        return pulumi.get(self, "acl_name")

    @property
    @pulumi.getter(name="addressIpVersion")
    def address_ip_version(self) -> pulumi.Output[str]:
        """
        The IP version. Valid values: `IPv4` and `IPv6`.
        """
        return pulumi.get(self, "address_ip_version")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        The dry run.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

