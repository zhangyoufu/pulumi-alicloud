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
                 entry_lists: Optional[pulumi.Input[Sequence[pulumi.Input['AclEntryListArgs']]]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Acl resource.
        :param pulumi.Input[Sequence[pulumi.Input['AclEntryListArgs']]] entry_lists: A list of entry (CIDR blocks) to be added. It contains two sub-fields as `Entry Block` follows. **NOTE:** "Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.",
        :param pulumi.Input[str] ip_version: The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
        :param pulumi.Input[str] name: Name of the access control list.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        if entry_lists is not None:
            warnings.warn("""Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.""", DeprecationWarning)
            pulumi.log.warn("""entry_lists is deprecated: Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.""")
        if entry_lists is not None:
            pulumi.set(__self__, "entry_lists", entry_lists)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="entryLists")
    @_utilities.deprecated("""Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.""")
    def entry_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AclEntryListArgs']]]]:
        """
        A list of entry (CIDR blocks) to be added. It contains two sub-fields as `Entry Block` follows. **NOTE:** "Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.",
        """
        return pulumi.get(self, "entry_lists")

    @entry_lists.setter
    def entry_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AclEntryListArgs']]]]):
        pulumi.set(self, "entry_lists", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[str]]:
        """
        The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the access control list.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _AclState:
    def __init__(__self__, *,
                 entry_lists: Optional[pulumi.Input[Sequence[pulumi.Input['AclEntryListArgs']]]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Acl resources.
        :param pulumi.Input[Sequence[pulumi.Input['AclEntryListArgs']]] entry_lists: A list of entry (CIDR blocks) to be added. It contains two sub-fields as `Entry Block` follows. **NOTE:** "Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.",
        :param pulumi.Input[str] ip_version: The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
        :param pulumi.Input[str] name: Name of the access control list.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        if entry_lists is not None:
            warnings.warn("""Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.""", DeprecationWarning)
            pulumi.log.warn("""entry_lists is deprecated: Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.""")
        if entry_lists is not None:
            pulumi.set(__self__, "entry_lists", entry_lists)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="entryLists")
    @_utilities.deprecated("""Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.""")
    def entry_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AclEntryListArgs']]]]:
        """
        A list of entry (CIDR blocks) to be added. It contains two sub-fields as `Entry Block` follows. **NOTE:** "Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.",
        """
        return pulumi.get(self, "entry_lists")

    @entry_lists.setter
    def entry_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AclEntryListArgs']]]]):
        pulumi.set(self, "entry_lists", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[str]]:
        """
        The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the access control list.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class Acl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entry_lists: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AclEntryListArgs', 'AclEntryListArgsDict']]]]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        An access control list contains multiple IP addresses or CIDR blocks.
        The access control list can help you to define multiple instance listening dimension,
        and to meet the multiple usage for single access control list.

        Server Load Balancer allows you to configure access control for listeners.
        You can configure different whitelists or blacklists for different listeners.

        You can configure access control
        when you create a listener or change access control configuration after a listener is created.

        > **NOTE:** One access control list can be attached to many Listeners in different load balancer as whitelists or blacklists.

        > **NOTE:** The maximum number of access control lists per region  is 50.

        > **NOTE:** The maximum number of IP addresses added each time is 50.

        > **NOTE:** The maximum number of entries per access control list is 300.

        > **NOTE:** The maximum number of listeners that an access control list can be added to is 50.

        For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).

        For information about acl and how to use it, see [Configure an access control list](https://www.alibabacloud.com/help/doc-detail/70015.htm).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        acl = alicloud.slb.Acl("acl",
            name="terraformslbaclconfig",
            ip_version="ipv4")
        ```

        ## Entry Block

        The entry mapping supports the following:

        * `entry` - (Optional, Computed) The CIDR blocks.
        * `comment` - (Optional, Computed) The comment of the entry.

        ## Import

        Server Load balancer access control list can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:slb/acl:Acl example acl-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AclEntryListArgs', 'AclEntryListArgsDict']]]] entry_lists: A list of entry (CIDR blocks) to be added. It contains two sub-fields as `Entry Block` follows. **NOTE:** "Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.",
        :param pulumi.Input[str] ip_version: The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
        :param pulumi.Input[str] name: Name of the access control list.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AclArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An access control list contains multiple IP addresses or CIDR blocks.
        The access control list can help you to define multiple instance listening dimension,
        and to meet the multiple usage for single access control list.

        Server Load Balancer allows you to configure access control for listeners.
        You can configure different whitelists or blacklists for different listeners.

        You can configure access control
        when you create a listener or change access control configuration after a listener is created.

        > **NOTE:** One access control list can be attached to many Listeners in different load balancer as whitelists or blacklists.

        > **NOTE:** The maximum number of access control lists per region  is 50.

        > **NOTE:** The maximum number of IP addresses added each time is 50.

        > **NOTE:** The maximum number of entries per access control list is 300.

        > **NOTE:** The maximum number of listeners that an access control list can be added to is 50.

        For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).

        For information about acl and how to use it, see [Configure an access control list](https://www.alibabacloud.com/help/doc-detail/70015.htm).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        acl = alicloud.slb.Acl("acl",
            name="terraformslbaclconfig",
            ip_version="ipv4")
        ```

        ## Entry Block

        The entry mapping supports the following:

        * `entry` - (Optional, Computed) The CIDR blocks.
        * `comment` - (Optional, Computed) The comment of the entry.

        ## Import

        Server Load balancer access control list can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:slb/acl:Acl example acl-abc123456
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
                 entry_lists: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AclEntryListArgs', 'AclEntryListArgsDict']]]]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AclArgs.__new__(AclArgs)

            __props__.__dict__["entry_lists"] = entry_lists
            __props__.__dict__["ip_version"] = ip_version
            __props__.__dict__["name"] = name
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["tags"] = tags
        super(Acl, __self__).__init__(
            'alicloud:slb/acl:Acl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            entry_lists: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AclEntryListArgs', 'AclEntryListArgsDict']]]]] = None,
            ip_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Acl':
        """
        Get an existing Acl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AclEntryListArgs', 'AclEntryListArgsDict']]]] entry_lists: A list of entry (CIDR blocks) to be added. It contains two sub-fields as `Entry Block` follows. **NOTE:** "Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.",
        :param pulumi.Input[str] ip_version: The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
        :param pulumi.Input[str] name: Name of the access control list.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AclState.__new__(_AclState)

        __props__.__dict__["entry_lists"] = entry_lists
        __props__.__dict__["ip_version"] = ip_version
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["tags"] = tags
        return Acl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="entryLists")
    @_utilities.deprecated("""Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.""")
    def entry_lists(self) -> pulumi.Output[Sequence['outputs.AclEntryList']]:
        """
        A list of entry (CIDR blocks) to be added. It contains two sub-fields as `Entry Block` follows. **NOTE:** "Field 'entry_list' has been deprecated from provider version 1.162.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_acl_entry_attachment'.",
        """
        return pulumi.get(self, "entry_lists")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> pulumi.Output[Optional[str]]:
        """
        The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the access control list.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        Resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

