# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GroupMembershipArgs', 'GroupMembership']

@pulumi.input_type
class GroupMembershipArgs:
    def __init__(__self__, *,
                 group_name: pulumi.Input[str],
                 user_names: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a GroupMembership resource.
        :param pulumi.Input[str] group_name: Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_names: Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
        """
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "user_names", user_names)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[str]:
        """
        Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="userNames")
    def user_names(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
        """
        return pulumi.get(self, "user_names")

    @user_names.setter
    def user_names(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "user_names", value)


@pulumi.input_type
class _GroupMembershipState:
    def __init__(__self__, *,
                 group_name: Optional[pulumi.Input[str]] = None,
                 user_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering GroupMembership resources.
        :param pulumi.Input[str] group_name: Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_names: Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
        """
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if user_names is not None:
            pulumi.set(__self__, "user_names", user_names)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="userNames")
    def user_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
        """
        return pulumi.get(self, "user_names")

    @user_names.setter
    def user_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "user_names", value)


class GroupMembership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 user_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a RAM Group membership resource.

        > **NOTE:** Available since v1.0.0+.

        ## Import

        RAM Group membership can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ram/groupMembership:GroupMembership example my-group
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_name: Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_names: Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupMembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a RAM Group membership resource.

        > **NOTE:** Available since v1.0.0+.

        ## Import

        RAM Group membership can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ram/groupMembership:GroupMembership example my-group
        ```

        :param str resource_name: The name of the resource.
        :param GroupMembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupMembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 user_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupMembershipArgs.__new__(GroupMembershipArgs)

            if group_name is None and not opts.urn:
                raise TypeError("Missing required property 'group_name'")
            __props__.__dict__["group_name"] = group_name
            if user_names is None and not opts.urn:
                raise TypeError("Missing required property 'user_names'")
            __props__.__dict__["user_names"] = user_names
        super(GroupMembership, __self__).__init__(
            'alicloud:ram/groupMembership:GroupMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_name: Optional[pulumi.Input[str]] = None,
            user_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'GroupMembership':
        """
        Get an existing GroupMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_name: Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_names: Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupMembershipState.__new__(_GroupMembershipState)

        __props__.__dict__["group_name"] = group_name
        __props__.__dict__["user_names"] = user_names
        return GroupMembership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="userNames")
    def user_names(self) -> pulumi.Output[Sequence[str]]:
        """
        Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
        """
        return pulumi.get(self, "user_names")

