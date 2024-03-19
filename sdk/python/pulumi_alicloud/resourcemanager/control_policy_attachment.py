# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ControlPolicyAttachmentArgs', 'ControlPolicyAttachment']

@pulumi.input_type
class ControlPolicyAttachmentArgs:
    def __init__(__self__, *,
                 policy_id: pulumi.Input[str],
                 target_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ControlPolicyAttachment resource.
        :param pulumi.Input[str] policy_id: The ID of control policy.
        :param pulumi.Input[str] target_id: The ID of target.
        """
        pulumi.set(__self__, "policy_id", policy_id)
        pulumi.set(__self__, "target_id", target_id)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        The ID of control policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Input[str]:
        """
        The ID of target.
        """
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_id", value)


@pulumi.input_type
class _ControlPolicyAttachmentState:
    def __init__(__self__, *,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 target_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ControlPolicyAttachment resources.
        :param pulumi.Input[str] policy_id: The ID of control policy.
        :param pulumi.Input[str] target_id: The ID of target.
        """
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if target_id is not None:
            pulumi.set(__self__, "target_id", target_id)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of control policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of target.
        """
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_id", value)


class ControlPolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Resource Manager Control Policy Attachment resource.

        For information about Resource Manager Control Policy Attachment and how to use it, see [What is Control Policy Attachment](https://www.alibabacloud.com/help/en/resource-management/latest/api-resourcedirectorymaster-2022-04-19-attachcontrolpolicy).

        > **NOTE:** Available since v1.120.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = random.RandomInteger("default",
            min=10000,
            max=99999)
        example_control_policy = alicloud.resourcemanager.ControlPolicy("exampleControlPolicy",
            control_policy_name=name,
            description=name,
            effect_scope="RAM",
            policy_document=\"\"\"  {
            "Version": "1",
            "Statement": [
              {
                "Effect": "Deny",
                "Action": [
                  "ram:UpdateRole",
                  "ram:DeleteRole",
                  "ram:AttachPolicyToRole",
                  "ram:DetachPolicyFromRole"
                ],
                "Resource": "acs:ram:*:*:role/ResourceDirectoryAccountAccessRole"
              }
            ]
          }
        \"\"\")
        example_folder = alicloud.resourcemanager.Folder("exampleFolder", folder_name=default.result.apply(lambda result: f"{name}-{result}"))
        example_control_policy_attachment = alicloud.resourcemanager.ControlPolicyAttachment("exampleControlPolicyAttachment",
            policy_id=example_control_policy.id,
            target_id=example_folder.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Resource Manager Control Policy Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:resourcemanager/controlPolicyAttachment:ControlPolicyAttachment example <policy_id>:<target_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_id: The ID of control policy.
        :param pulumi.Input[str] target_id: The ID of target.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ControlPolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Resource Manager Control Policy Attachment resource.

        For information about Resource Manager Control Policy Attachment and how to use it, see [What is Control Policy Attachment](https://www.alibabacloud.com/help/en/resource-management/latest/api-resourcedirectorymaster-2022-04-19-attachcontrolpolicy).

        > **NOTE:** Available since v1.120.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = random.RandomInteger("default",
            min=10000,
            max=99999)
        example_control_policy = alicloud.resourcemanager.ControlPolicy("exampleControlPolicy",
            control_policy_name=name,
            description=name,
            effect_scope="RAM",
            policy_document=\"\"\"  {
            "Version": "1",
            "Statement": [
              {
                "Effect": "Deny",
                "Action": [
                  "ram:UpdateRole",
                  "ram:DeleteRole",
                  "ram:AttachPolicyToRole",
                  "ram:DetachPolicyFromRole"
                ],
                "Resource": "acs:ram:*:*:role/ResourceDirectoryAccountAccessRole"
              }
            ]
          }
        \"\"\")
        example_folder = alicloud.resourcemanager.Folder("exampleFolder", folder_name=default.result.apply(lambda result: f"{name}-{result}"))
        example_control_policy_attachment = alicloud.resourcemanager.ControlPolicyAttachment("exampleControlPolicyAttachment",
            policy_id=example_control_policy.id,
            target_id=example_folder.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Resource Manager Control Policy Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:resourcemanager/controlPolicyAttachment:ControlPolicyAttachment example <policy_id>:<target_id>
        ```

        :param str resource_name: The name of the resource.
        :param ControlPolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ControlPolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ControlPolicyAttachmentArgs.__new__(ControlPolicyAttachmentArgs)

            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__.__dict__["policy_id"] = policy_id
            if target_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_id'")
            __props__.__dict__["target_id"] = target_id
        super(ControlPolicyAttachment, __self__).__init__(
            'alicloud:resourcemanager/controlPolicyAttachment:ControlPolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            target_id: Optional[pulumi.Input[str]] = None) -> 'ControlPolicyAttachment':
        """
        Get an existing ControlPolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_id: The ID of control policy.
        :param pulumi.Input[str] target_id: The ID of target.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ControlPolicyAttachmentState.__new__(_ControlPolicyAttachmentState)

        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["target_id"] = target_id
        return ControlPolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        The ID of control policy.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Output[str]:
        """
        The ID of target.
        """
        return pulumi.get(self, "target_id")

