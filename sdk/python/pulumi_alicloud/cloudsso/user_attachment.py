# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UserAttachmentArgs', 'UserAttachment']

@pulumi.input_type
class UserAttachmentArgs:
    def __init__(__self__, *,
                 directory_id: pulumi.Input[str],
                 group_id: pulumi.Input[str],
                 user_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a UserAttachment resource.
        :param pulumi.Input[str] directory_id: The ID of the Directory.
        :param pulumi.Input[str] group_id: The Group ID.
        :param pulumi.Input[str] user_id: The User ID.
        """
        pulumi.set(__self__, "directory_id", directory_id)
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> pulumi.Input[str]:
        """
        The ID of the Directory.
        """
        return pulumi.get(self, "directory_id")

    @directory_id.setter
    def directory_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "directory_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        The Group ID.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        The User ID.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _UserAttachmentState:
    def __init__(__self__, *,
                 directory_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserAttachment resources.
        :param pulumi.Input[str] directory_id: The ID of the Directory.
        :param pulumi.Input[str] group_id: The Group ID.
        :param pulumi.Input[str] user_id: The User ID.
        """
        if directory_id is not None:
            pulumi.set(__self__, "directory_id", directory_id)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Directory.
        """
        return pulumi.get(self, "directory_id")

    @directory_id.setter
    def directory_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Group ID.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The User ID.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class UserAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 directory_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cloud SSO User Attachment resource.

        For information about Cloud SSO User Attachment and how to use it, see [What is User Attachment](https://www.alibabacloud.com/help/en/doc-detail/264683.htm).

        > **NOTE:** Available in v1.141.0+.

        > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region

        ## Import

        Cloud SSO User Attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cloudsso/userAttachment:UserAttachment example <directory_id>:<group_id>:<user_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] directory_id: The ID of the Directory.
        :param pulumi.Input[str] group_id: The Group ID.
        :param pulumi.Input[str] user_id: The User ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud SSO User Attachment resource.

        For information about Cloud SSO User Attachment and how to use it, see [What is User Attachment](https://www.alibabacloud.com/help/en/doc-detail/264683.htm).

        > **NOTE:** Available in v1.141.0+.

        > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region

        ## Import

        Cloud SSO User Attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cloudsso/userAttachment:UserAttachment example <directory_id>:<group_id>:<user_id>
        ```

        :param str resource_name: The name of the resource.
        :param UserAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 directory_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserAttachmentArgs.__new__(UserAttachmentArgs)

            if directory_id is None and not opts.urn:
                raise TypeError("Missing required property 'directory_id'")
            __props__.__dict__["directory_id"] = directory_id
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(UserAttachment, __self__).__init__(
            'alicloud:cloudsso/userAttachment:UserAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            directory_id: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'UserAttachment':
        """
        Get an existing UserAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] directory_id: The ID of the Directory.
        :param pulumi.Input[str] group_id: The Group ID.
        :param pulumi.Input[str] user_id: The User ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserAttachmentState.__new__(_UserAttachmentState)

        __props__.__dict__["directory_id"] = directory_id
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["user_id"] = user_id
        return UserAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> pulumi.Output[str]:
        """
        The ID of the Directory.
        """
        return pulumi.get(self, "directory_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        The Group ID.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        The User ID.
        """
        return pulumi.get(self, "user_id")

