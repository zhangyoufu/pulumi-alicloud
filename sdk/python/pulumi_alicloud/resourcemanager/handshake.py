# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HandshakeArgs', 'Handshake']

@pulumi.input_type
class HandshakeArgs:
    def __init__(__self__, *,
                 target_entity: pulumi.Input[str],
                 target_type: pulumi.Input[str],
                 note: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Handshake resource.
        :param pulumi.Input[str] target_entity: Invited account ID or login email.
        :param pulumi.Input[str] target_type: Type of account being invited. Valid values: `Account`, `Email`.
        :param pulumi.Input[str] note: Remarks. The maximum length is 1024 characters.
        """
        pulumi.set(__self__, "target_entity", target_entity)
        pulumi.set(__self__, "target_type", target_type)
        if note is not None:
            pulumi.set(__self__, "note", note)

    @property
    @pulumi.getter(name="targetEntity")
    def target_entity(self) -> pulumi.Input[str]:
        """
        Invited account ID or login email.
        """
        return pulumi.get(self, "target_entity")

    @target_entity.setter
    def target_entity(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_entity", value)

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> pulumi.Input[str]:
        """
        Type of account being invited. Valid values: `Account`, `Email`.
        """
        return pulumi.get(self, "target_type")

    @target_type.setter
    def target_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_type", value)

    @property
    @pulumi.getter
    def note(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks. The maximum length is 1024 characters.
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "note", value)


@pulumi.input_type
class _HandshakeState:
    def __init__(__self__, *,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 master_account_id: Optional[pulumi.Input[str]] = None,
                 master_account_name: Optional[pulumi.Input[str]] = None,
                 modify_time: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 resource_directory_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 target_entity: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Handshake resources.
        :param pulumi.Input[str] expire_time: The expiration time of the invitation.
        :param pulumi.Input[str] master_account_id: Resource account master account ID.
        :param pulumi.Input[str] master_account_name: The name of the main account of the resource directory.
        :param pulumi.Input[str] modify_time: The modification time of the invitation.
        :param pulumi.Input[str] note: Remarks. The maximum length is 1024 characters.
        :param pulumi.Input[str] resource_directory_id: Resource directory ID.
        :param pulumi.Input[str] status: Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
        :param pulumi.Input[str] target_entity: Invited account ID or login email.
        :param pulumi.Input[str] target_type: Type of account being invited. Valid values: `Account`, `Email`.
        """
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if master_account_id is not None:
            pulumi.set(__self__, "master_account_id", master_account_id)
        if master_account_name is not None:
            pulumi.set(__self__, "master_account_name", master_account_name)
        if modify_time is not None:
            pulumi.set(__self__, "modify_time", modify_time)
        if note is not None:
            pulumi.set(__self__, "note", note)
        if resource_directory_id is not None:
            pulumi.set(__self__, "resource_directory_id", resource_directory_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if target_entity is not None:
            pulumi.set(__self__, "target_entity", target_entity)
        if target_type is not None:
            pulumi.set(__self__, "target_type", target_type)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        The expiration time of the invitation.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter(name="masterAccountId")
    def master_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource account master account ID.
        """
        return pulumi.get(self, "master_account_id")

    @master_account_id.setter
    def master_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_account_id", value)

    @property
    @pulumi.getter(name="masterAccountName")
    def master_account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the main account of the resource directory.
        """
        return pulumi.get(self, "master_account_name")

    @master_account_name.setter
    def master_account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_account_name", value)

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> Optional[pulumi.Input[str]]:
        """
        The modification time of the invitation.
        """
        return pulumi.get(self, "modify_time")

    @modify_time.setter
    def modify_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modify_time", value)

    @property
    @pulumi.getter
    def note(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks. The maximum length is 1024 characters.
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "note", value)

    @property
    @pulumi.getter(name="resourceDirectoryId")
    def resource_directory_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource directory ID.
        """
        return pulumi.get(self, "resource_directory_id")

    @resource_directory_id.setter
    def resource_directory_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_directory_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="targetEntity")
    def target_entity(self) -> Optional[pulumi.Input[str]]:
        """
        Invited account ID or login email.
        """
        return pulumi.get(self, "target_entity")

    @target_entity.setter
    def target_entity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_entity", value)

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of account being invited. Valid values: `Account`, `Email`.
        """
        return pulumi.get(self, "target_type")

    @target_type.setter
    def target_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_type", value)


class Handshake(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 target_entity: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Resource Manager handshake resource. You can invite accounts to join a resource directory for unified management.
        For information about Resource Manager handshake and how to use it, see [What is Resource Manager handshake](https://www.alibabacloud.com/help/en/doc-detail/135287.htm).

        > **NOTE:** Available in v1.82.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Add a Resource Manager handshake.
        example = alicloud.resourcemanager.Handshake("example",
            note="test resource manager handshake",
            target_entity="1182775234******",
            target_type="Account")
        ```

        ## Import

        Resource Manager handshake can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:resourcemanager/handshake:Handshake example h-QmdexeFm1kE*****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] note: Remarks. The maximum length is 1024 characters.
        :param pulumi.Input[str] target_entity: Invited account ID or login email.
        :param pulumi.Input[str] target_type: Type of account being invited. Valid values: `Account`, `Email`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HandshakeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Resource Manager handshake resource. You can invite accounts to join a resource directory for unified management.
        For information about Resource Manager handshake and how to use it, see [What is Resource Manager handshake](https://www.alibabacloud.com/help/en/doc-detail/135287.htm).

        > **NOTE:** Available in v1.82.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Add a Resource Manager handshake.
        example = alicloud.resourcemanager.Handshake("example",
            note="test resource manager handshake",
            target_entity="1182775234******",
            target_type="Account")
        ```

        ## Import

        Resource Manager handshake can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:resourcemanager/handshake:Handshake example h-QmdexeFm1kE*****
        ```

        :param str resource_name: The name of the resource.
        :param HandshakeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HandshakeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 target_entity: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HandshakeArgs.__new__(HandshakeArgs)

            __props__.__dict__["note"] = note
            if target_entity is None and not opts.urn:
                raise TypeError("Missing required property 'target_entity'")
            __props__.__dict__["target_entity"] = target_entity
            if target_type is None and not opts.urn:
                raise TypeError("Missing required property 'target_type'")
            __props__.__dict__["target_type"] = target_type
            __props__.__dict__["expire_time"] = None
            __props__.__dict__["master_account_id"] = None
            __props__.__dict__["master_account_name"] = None
            __props__.__dict__["modify_time"] = None
            __props__.__dict__["resource_directory_id"] = None
            __props__.__dict__["status"] = None
        super(Handshake, __self__).__init__(
            'alicloud:resourcemanager/handshake:Handshake',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expire_time: Optional[pulumi.Input[str]] = None,
            master_account_id: Optional[pulumi.Input[str]] = None,
            master_account_name: Optional[pulumi.Input[str]] = None,
            modify_time: Optional[pulumi.Input[str]] = None,
            note: Optional[pulumi.Input[str]] = None,
            resource_directory_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            target_entity: Optional[pulumi.Input[str]] = None,
            target_type: Optional[pulumi.Input[str]] = None) -> 'Handshake':
        """
        Get an existing Handshake resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expire_time: The expiration time of the invitation.
        :param pulumi.Input[str] master_account_id: Resource account master account ID.
        :param pulumi.Input[str] master_account_name: The name of the main account of the resource directory.
        :param pulumi.Input[str] modify_time: The modification time of the invitation.
        :param pulumi.Input[str] note: Remarks. The maximum length is 1024 characters.
        :param pulumi.Input[str] resource_directory_id: Resource directory ID.
        :param pulumi.Input[str] status: Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
        :param pulumi.Input[str] target_entity: Invited account ID or login email.
        :param pulumi.Input[str] target_type: Type of account being invited. Valid values: `Account`, `Email`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HandshakeState.__new__(_HandshakeState)

        __props__.__dict__["expire_time"] = expire_time
        __props__.__dict__["master_account_id"] = master_account_id
        __props__.__dict__["master_account_name"] = master_account_name
        __props__.__dict__["modify_time"] = modify_time
        __props__.__dict__["note"] = note
        __props__.__dict__["resource_directory_id"] = resource_directory_id
        __props__.__dict__["status"] = status
        __props__.__dict__["target_entity"] = target_entity
        __props__.__dict__["target_type"] = target_type
        return Handshake(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        The expiration time of the invitation.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="masterAccountId")
    def master_account_id(self) -> pulumi.Output[str]:
        """
        Resource account master account ID.
        """
        return pulumi.get(self, "master_account_id")

    @property
    @pulumi.getter(name="masterAccountName")
    def master_account_name(self) -> pulumi.Output[str]:
        """
        The name of the main account of the resource directory.
        """
        return pulumi.get(self, "master_account_name")

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> pulumi.Output[str]:
        """
        The modification time of the invitation.
        """
        return pulumi.get(self, "modify_time")

    @property
    @pulumi.getter
    def note(self) -> pulumi.Output[Optional[str]]:
        """
        Remarks. The maximum length is 1024 characters.
        """
        return pulumi.get(self, "note")

    @property
    @pulumi.getter(name="resourceDirectoryId")
    def resource_directory_id(self) -> pulumi.Output[str]:
        """
        Resource directory ID.
        """
        return pulumi.get(self, "resource_directory_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="targetEntity")
    def target_entity(self) -> pulumi.Output[str]:
        """
        Invited account ID or login email.
        """
        return pulumi.get(self, "target_entity")

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> pulumi.Output[str]:
        """
        Type of account being invited. Valid values: `Account`, `Email`.
        """
        return pulumi.get(self, "target_type")

