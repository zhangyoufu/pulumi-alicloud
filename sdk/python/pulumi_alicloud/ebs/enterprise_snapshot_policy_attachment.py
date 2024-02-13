# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EnterpriseSnapshotPolicyAttachmentArgs', 'EnterpriseSnapshotPolicyAttachment']

@pulumi.input_type
class EnterpriseSnapshotPolicyAttachmentArgs:
    def __init__(__self__, *,
                 policy_id: pulumi.Input[str],
                 disk_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EnterpriseSnapshotPolicyAttachment resource.
        :param pulumi.Input[str] policy_id: the enterprise snapshot policy id.
        :param pulumi.Input[str] disk_id: Cloud Disk ID.
        """
        pulumi.set(__self__, "policy_id", policy_id)
        if disk_id is not None:
            pulumi.set(__self__, "disk_id", disk_id)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        the enterprise snapshot policy id.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud Disk ID.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_id", value)


@pulumi.input_type
class _EnterpriseSnapshotPolicyAttachmentState:
    def __init__(__self__, *,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EnterpriseSnapshotPolicyAttachment resources.
        :param pulumi.Input[str] disk_id: Cloud Disk ID.
        :param pulumi.Input[str] policy_id: the enterprise snapshot policy id.
        """
        if disk_id is not None:
            pulumi.set(__self__, "disk_id", disk_id)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud Disk ID.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_id", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        the enterprise snapshot policy id.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)


class EnterpriseSnapshotPolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a EBS Enterprise Snapshot Policy Attachment resource. Enterprise-level snapshot policy cloud disk binding relationship.

        For information about EBS Enterprise Snapshot Policy Attachment and how to use it, see [What is Enterprise Snapshot Policy Attachment](https://www.alibabacloud.com/help/en/).

        > **NOTE:** Available since v1.215.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_jk_w46o = alicloud.ecs.EcsDisk("defaultJkW46o",
            category="cloud_essd",
            description="esp-attachment-test",
            zone_id="cn-hangzhou-i",
            performance_level="PL1",
            size=20,
            disk_name=name)
        default_pe3jj_r = alicloud.ebs.EnterpriseSnapshotPolicy("defaultPE3jjR",
            status="DISABLED",
            desc="DESC",
            schedule=alicloud.ebs.EnterpriseSnapshotPolicyScheduleArgs(
                cron_expression="0 0 0 1 * ?",
            ),
            enterprise_snapshot_policy_name=name,
            target_type="DISK",
            retain_rule=alicloud.ebs.EnterpriseSnapshotPolicyRetainRuleArgs(
                time_interval=120,
                time_unit="DAYS",
            ))
        default = alicloud.ebs.EnterpriseSnapshotPolicyAttachment("default",
            policy_id=default_pe3jj_r.id,
            disk_id=default_jk_w46o.id)
        ```

        ## Import

        EBS Enterprise Snapshot Policy Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ebs/enterpriseSnapshotPolicyAttachment:EnterpriseSnapshotPolicyAttachment example <policy_id>:<disk_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_id: Cloud Disk ID.
        :param pulumi.Input[str] policy_id: the enterprise snapshot policy id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnterpriseSnapshotPolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a EBS Enterprise Snapshot Policy Attachment resource. Enterprise-level snapshot policy cloud disk binding relationship.

        For information about EBS Enterprise Snapshot Policy Attachment and how to use it, see [What is Enterprise Snapshot Policy Attachment](https://www.alibabacloud.com/help/en/).

        > **NOTE:** Available since v1.215.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_jk_w46o = alicloud.ecs.EcsDisk("defaultJkW46o",
            category="cloud_essd",
            description="esp-attachment-test",
            zone_id="cn-hangzhou-i",
            performance_level="PL1",
            size=20,
            disk_name=name)
        default_pe3jj_r = alicloud.ebs.EnterpriseSnapshotPolicy("defaultPE3jjR",
            status="DISABLED",
            desc="DESC",
            schedule=alicloud.ebs.EnterpriseSnapshotPolicyScheduleArgs(
                cron_expression="0 0 0 1 * ?",
            ),
            enterprise_snapshot_policy_name=name,
            target_type="DISK",
            retain_rule=alicloud.ebs.EnterpriseSnapshotPolicyRetainRuleArgs(
                time_interval=120,
                time_unit="DAYS",
            ))
        default = alicloud.ebs.EnterpriseSnapshotPolicyAttachment("default",
            policy_id=default_pe3jj_r.id,
            disk_id=default_jk_w46o.id)
        ```

        ## Import

        EBS Enterprise Snapshot Policy Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ebs/enterpriseSnapshotPolicyAttachment:EnterpriseSnapshotPolicyAttachment example <policy_id>:<disk_id>
        ```

        :param str resource_name: The name of the resource.
        :param EnterpriseSnapshotPolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnterpriseSnapshotPolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnterpriseSnapshotPolicyAttachmentArgs.__new__(EnterpriseSnapshotPolicyAttachmentArgs)

            __props__.__dict__["disk_id"] = disk_id
            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__.__dict__["policy_id"] = policy_id
        super(EnterpriseSnapshotPolicyAttachment, __self__).__init__(
            'alicloud:ebs/enterpriseSnapshotPolicyAttachment:EnterpriseSnapshotPolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disk_id: Optional[pulumi.Input[str]] = None,
            policy_id: Optional[pulumi.Input[str]] = None) -> 'EnterpriseSnapshotPolicyAttachment':
        """
        Get an existing EnterpriseSnapshotPolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_id: Cloud Disk ID.
        :param pulumi.Input[str] policy_id: the enterprise snapshot policy id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnterpriseSnapshotPolicyAttachmentState.__new__(_EnterpriseSnapshotPolicyAttachmentState)

        __props__.__dict__["disk_id"] = disk_id
        __props__.__dict__["policy_id"] = policy_id
        return EnterpriseSnapshotPolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Output[str]:
        """
        Cloud Disk ID.
        """
        return pulumi.get(self, "disk_id")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        the enterprise snapshot policy id.
        """
        return pulumi.get(self, "policy_id")

