# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EcsAutoSnapshotPolicyAttachmentArgs', 'EcsAutoSnapshotPolicyAttachment']

@pulumi.input_type
class EcsAutoSnapshotPolicyAttachmentArgs:
    def __init__(__self__, *,
                 auto_snapshot_policy_id: pulumi.Input[str],
                 disk_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a EcsAutoSnapshotPolicyAttachment resource.
        :param pulumi.Input[str] auto_snapshot_policy_id: The auto snapshot policy id.
        :param pulumi.Input[str] disk_id: The disk id.
        """
        pulumi.set(__self__, "auto_snapshot_policy_id", auto_snapshot_policy_id)
        pulumi.set(__self__, "disk_id", disk_id)

    @property
    @pulumi.getter(name="autoSnapshotPolicyId")
    def auto_snapshot_policy_id(self) -> pulumi.Input[str]:
        """
        The auto snapshot policy id.
        """
        return pulumi.get(self, "auto_snapshot_policy_id")

    @auto_snapshot_policy_id.setter
    def auto_snapshot_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "auto_snapshot_policy_id", value)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Input[str]:
        """
        The disk id.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk_id", value)


@pulumi.input_type
class _EcsAutoSnapshotPolicyAttachmentState:
    def __init__(__self__, *,
                 auto_snapshot_policy_id: Optional[pulumi.Input[str]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EcsAutoSnapshotPolicyAttachment resources.
        :param pulumi.Input[str] auto_snapshot_policy_id: The auto snapshot policy id.
        :param pulumi.Input[str] disk_id: The disk id.
        """
        if auto_snapshot_policy_id is not None:
            pulumi.set(__self__, "auto_snapshot_policy_id", auto_snapshot_policy_id)
        if disk_id is not None:
            pulumi.set(__self__, "disk_id", disk_id)

    @property
    @pulumi.getter(name="autoSnapshotPolicyId")
    def auto_snapshot_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The auto snapshot policy id.
        """
        return pulumi.get(self, "auto_snapshot_policy_id")

    @auto_snapshot_policy_id.setter
    def auto_snapshot_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_snapshot_policy_id", value)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> Optional[pulumi.Input[str]]:
        """
        The disk id.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_id", value)


class EcsAutoSnapshotPolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_snapshot_policy_id: Optional[pulumi.Input[str]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ECS Auto Snapshot Policy Attachment resource.

        For information about ECS Auto Snapshot Policy Attachment and how to use it, see [What is Auto Snapshot Policy Attachment](https://www.alibabacloud.com/help/en/doc-detail/25531.htm).

        > **NOTE:** Available in v1.122.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.ecs.EcsAutoSnapshotPolicyAttachment("example",
            auto_snapshot_policy_id="s-ge465xxxx",
            disk_id="d-gw835xxxx")
        ```

        ## Import

        ECS Auto Snapshot Policy Attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ecs/ecsAutoSnapshotPolicyAttachment:EcsAutoSnapshotPolicyAttachment example s-abcd12345:d-abcd12345
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_snapshot_policy_id: The auto snapshot policy id.
        :param pulumi.Input[str] disk_id: The disk id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EcsAutoSnapshotPolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ECS Auto Snapshot Policy Attachment resource.

        For information about ECS Auto Snapshot Policy Attachment and how to use it, see [What is Auto Snapshot Policy Attachment](https://www.alibabacloud.com/help/en/doc-detail/25531.htm).

        > **NOTE:** Available in v1.122.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.ecs.EcsAutoSnapshotPolicyAttachment("example",
            auto_snapshot_policy_id="s-ge465xxxx",
            disk_id="d-gw835xxxx")
        ```

        ## Import

        ECS Auto Snapshot Policy Attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ecs/ecsAutoSnapshotPolicyAttachment:EcsAutoSnapshotPolicyAttachment example s-abcd12345:d-abcd12345
        ```

        :param str resource_name: The name of the resource.
        :param EcsAutoSnapshotPolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EcsAutoSnapshotPolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_snapshot_policy_id: Optional[pulumi.Input[str]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EcsAutoSnapshotPolicyAttachmentArgs.__new__(EcsAutoSnapshotPolicyAttachmentArgs)

            if auto_snapshot_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'auto_snapshot_policy_id'")
            __props__.__dict__["auto_snapshot_policy_id"] = auto_snapshot_policy_id
            if disk_id is None and not opts.urn:
                raise TypeError("Missing required property 'disk_id'")
            __props__.__dict__["disk_id"] = disk_id
        super(EcsAutoSnapshotPolicyAttachment, __self__).__init__(
            'alicloud:ecs/ecsAutoSnapshotPolicyAttachment:EcsAutoSnapshotPolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_snapshot_policy_id: Optional[pulumi.Input[str]] = None,
            disk_id: Optional[pulumi.Input[str]] = None) -> 'EcsAutoSnapshotPolicyAttachment':
        """
        Get an existing EcsAutoSnapshotPolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_snapshot_policy_id: The auto snapshot policy id.
        :param pulumi.Input[str] disk_id: The disk id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EcsAutoSnapshotPolicyAttachmentState.__new__(_EcsAutoSnapshotPolicyAttachmentState)

        __props__.__dict__["auto_snapshot_policy_id"] = auto_snapshot_policy_id
        __props__.__dict__["disk_id"] = disk_id
        return EcsAutoSnapshotPolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoSnapshotPolicyId")
    def auto_snapshot_policy_id(self) -> pulumi.Output[str]:
        """
        The auto snapshot policy id.
        """
        return pulumi.get(self, "auto_snapshot_policy_id")

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Output[str]:
        """
        The disk id.
        """
        return pulumi.get(self, "disk_id")

