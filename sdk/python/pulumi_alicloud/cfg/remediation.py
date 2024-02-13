# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RemediationArgs', 'Remediation']

@pulumi.input_type
class RemediationArgs:
    def __init__(__self__, *,
                 config_rule_id: pulumi.Input[str],
                 invoke_type: pulumi.Input[str],
                 params: pulumi.Input[str],
                 remediation_template_id: pulumi.Input[str],
                 remediation_type: pulumi.Input[str],
                 remediation_source_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Remediation resource.
        :param pulumi.Input[str] config_rule_id: Rule ID.
        :param pulumi.Input[str] invoke_type: Execution type, valid values: `Manual`, `Automatic`.
        :param pulumi.Input[str] params: Remediation parameter.
        :param pulumi.Input[str] remediation_template_id: Remediation template ID.
        :param pulumi.Input[str] remediation_type: Remediation type, valid values: `OOS`, `FC`.
               
               The following arguments will be discarded. Please use new fields as soon as possible:
        :param pulumi.Input[str] remediation_source_type: Remediation resource type, valid values: `ALIYUN` , `CUSTOMER`.
        """
        pulumi.set(__self__, "config_rule_id", config_rule_id)
        pulumi.set(__self__, "invoke_type", invoke_type)
        pulumi.set(__self__, "params", params)
        pulumi.set(__self__, "remediation_template_id", remediation_template_id)
        pulumi.set(__self__, "remediation_type", remediation_type)
        if remediation_source_type is not None:
            pulumi.set(__self__, "remediation_source_type", remediation_source_type)

    @property
    @pulumi.getter(name="configRuleId")
    def config_rule_id(self) -> pulumi.Input[str]:
        """
        Rule ID.
        """
        return pulumi.get(self, "config_rule_id")

    @config_rule_id.setter
    def config_rule_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_rule_id", value)

    @property
    @pulumi.getter(name="invokeType")
    def invoke_type(self) -> pulumi.Input[str]:
        """
        Execution type, valid values: `Manual`, `Automatic`.
        """
        return pulumi.get(self, "invoke_type")

    @invoke_type.setter
    def invoke_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "invoke_type", value)

    @property
    @pulumi.getter
    def params(self) -> pulumi.Input[str]:
        """
        Remediation parameter.
        """
        return pulumi.get(self, "params")

    @params.setter
    def params(self, value: pulumi.Input[str]):
        pulumi.set(self, "params", value)

    @property
    @pulumi.getter(name="remediationTemplateId")
    def remediation_template_id(self) -> pulumi.Input[str]:
        """
        Remediation template ID.
        """
        return pulumi.get(self, "remediation_template_id")

    @remediation_template_id.setter
    def remediation_template_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "remediation_template_id", value)

    @property
    @pulumi.getter(name="remediationType")
    def remediation_type(self) -> pulumi.Input[str]:
        """
        Remediation type, valid values: `OOS`, `FC`.

        The following arguments will be discarded. Please use new fields as soon as possible:
        """
        return pulumi.get(self, "remediation_type")

    @remediation_type.setter
    def remediation_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "remediation_type", value)

    @property
    @pulumi.getter(name="remediationSourceType")
    def remediation_source_type(self) -> Optional[pulumi.Input[str]]:
        """
        Remediation resource type, valid values: `ALIYUN` , `CUSTOMER`.
        """
        return pulumi.get(self, "remediation_source_type")

    @remediation_source_type.setter
    def remediation_source_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remediation_source_type", value)


@pulumi.input_type
class _RemediationState:
    def __init__(__self__, *,
                 config_rule_id: Optional[pulumi.Input[str]] = None,
                 invoke_type: Optional[pulumi.Input[str]] = None,
                 params: Optional[pulumi.Input[str]] = None,
                 remediation_id: Optional[pulumi.Input[str]] = None,
                 remediation_source_type: Optional[pulumi.Input[str]] = None,
                 remediation_template_id: Optional[pulumi.Input[str]] = None,
                 remediation_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Remediation resources.
        :param pulumi.Input[str] config_rule_id: Rule ID.
        :param pulumi.Input[str] invoke_type: Execution type, valid values: `Manual`, `Automatic`.
        :param pulumi.Input[str] params: Remediation parameter.
        :param pulumi.Input[str] remediation_id: Remediation ID.
        :param pulumi.Input[str] remediation_source_type: Remediation resource type, valid values: `ALIYUN` , `CUSTOMER`.
        :param pulumi.Input[str] remediation_template_id: Remediation template ID.
        :param pulumi.Input[str] remediation_type: Remediation type, valid values: `OOS`, `FC`.
               
               The following arguments will be discarded. Please use new fields as soon as possible:
        """
        if config_rule_id is not None:
            pulumi.set(__self__, "config_rule_id", config_rule_id)
        if invoke_type is not None:
            pulumi.set(__self__, "invoke_type", invoke_type)
        if params is not None:
            pulumi.set(__self__, "params", params)
        if remediation_id is not None:
            pulumi.set(__self__, "remediation_id", remediation_id)
        if remediation_source_type is not None:
            pulumi.set(__self__, "remediation_source_type", remediation_source_type)
        if remediation_template_id is not None:
            pulumi.set(__self__, "remediation_template_id", remediation_template_id)
        if remediation_type is not None:
            pulumi.set(__self__, "remediation_type", remediation_type)

    @property
    @pulumi.getter(name="configRuleId")
    def config_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        Rule ID.
        """
        return pulumi.get(self, "config_rule_id")

    @config_rule_id.setter
    def config_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_rule_id", value)

    @property
    @pulumi.getter(name="invokeType")
    def invoke_type(self) -> Optional[pulumi.Input[str]]:
        """
        Execution type, valid values: `Manual`, `Automatic`.
        """
        return pulumi.get(self, "invoke_type")

    @invoke_type.setter
    def invoke_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invoke_type", value)

    @property
    @pulumi.getter
    def params(self) -> Optional[pulumi.Input[str]]:
        """
        Remediation parameter.
        """
        return pulumi.get(self, "params")

    @params.setter
    def params(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "params", value)

    @property
    @pulumi.getter(name="remediationId")
    def remediation_id(self) -> Optional[pulumi.Input[str]]:
        """
        Remediation ID.
        """
        return pulumi.get(self, "remediation_id")

    @remediation_id.setter
    def remediation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remediation_id", value)

    @property
    @pulumi.getter(name="remediationSourceType")
    def remediation_source_type(self) -> Optional[pulumi.Input[str]]:
        """
        Remediation resource type, valid values: `ALIYUN` , `CUSTOMER`.
        """
        return pulumi.get(self, "remediation_source_type")

    @remediation_source_type.setter
    def remediation_source_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remediation_source_type", value)

    @property
    @pulumi.getter(name="remediationTemplateId")
    def remediation_template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Remediation template ID.
        """
        return pulumi.get(self, "remediation_template_id")

    @remediation_template_id.setter
    def remediation_template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remediation_template_id", value)

    @property
    @pulumi.getter(name="remediationType")
    def remediation_type(self) -> Optional[pulumi.Input[str]]:
        """
        Remediation type, valid values: `OOS`, `FC`.

        The following arguments will be discarded. Please use new fields as soon as possible:
        """
        return pulumi.get(self, "remediation_type")

    @remediation_type.setter
    def remediation_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remediation_type", value)


class Remediation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_rule_id: Optional[pulumi.Input[str]] = None,
                 invoke_type: Optional[pulumi.Input[str]] = None,
                 params: Optional[pulumi.Input[str]] = None,
                 remediation_source_type: Optional[pulumi.Input[str]] = None,
                 remediation_template_id: Optional[pulumi.Input[str]] = None,
                 remediation_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Config Remediation resource.

        For information about Config Remediation and how to use it, see [What is Remediation](https://www.alibabacloud.com/help/en/cloud-config/latest/api-config-2020-09-07-createremediation).

        > **NOTE:** Available since v1.204.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example-oss"
        default_regions = alicloud.get_regions(current=True)
        default_bucket = alicloud.oss.Bucket("defaultBucket",
            bucket=name,
            acl="public-read",
            tags={
                "For": "example",
            })
        default_rule = alicloud.cfg.Rule("defaultRule",
            description="If the ACL policy of the OSS bucket denies read access from the Internet, the configuration is considered compliant.",
            source_owner="ALIYUN",
            source_identifier="oss-bucket-public-read-prohibited",
            risk_level=1,
            tag_key_scope="For",
            tag_value_scope="example",
            region_ids_scope=default_regions.regions[0].id,
            config_rule_trigger_types="ConfigurationItemChangeNotification",
            resource_types_scopes=["ACS::OSS::Bucket"],
            rule_name="oss-bucket-public-read-prohibited")
        default_remediation = alicloud.cfg.Remediation("defaultRemediation",
            config_rule_id=default_rule.config_rule_id,
            remediation_template_id="ACS-OSS-PutBucketAcl",
            remediation_source_type="ALIYUN",
            invoke_type="MANUAL_EXECUTION",
            params=default_bucket.bucket.apply(lambda bucket: f"{{\\"bucketName\\": \\"{bucket}\\", \\"regionId\\": \\"{default_regions.regions[0].id}\\", \\"permissionName\\": \\"private\\"}}"),
            remediation_type="OOS")
        ```

        ## Import

        Config Remediation can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cfg/remediation:Remediation example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_rule_id: Rule ID.
        :param pulumi.Input[str] invoke_type: Execution type, valid values: `Manual`, `Automatic`.
        :param pulumi.Input[str] params: Remediation parameter.
        :param pulumi.Input[str] remediation_source_type: Remediation resource type, valid values: `ALIYUN` , `CUSTOMER`.
        :param pulumi.Input[str] remediation_template_id: Remediation template ID.
        :param pulumi.Input[str] remediation_type: Remediation type, valid values: `OOS`, `FC`.
               
               The following arguments will be discarded. Please use new fields as soon as possible:
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RemediationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Config Remediation resource.

        For information about Config Remediation and how to use it, see [What is Remediation](https://www.alibabacloud.com/help/en/cloud-config/latest/api-config-2020-09-07-createremediation).

        > **NOTE:** Available since v1.204.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example-oss"
        default_regions = alicloud.get_regions(current=True)
        default_bucket = alicloud.oss.Bucket("defaultBucket",
            bucket=name,
            acl="public-read",
            tags={
                "For": "example",
            })
        default_rule = alicloud.cfg.Rule("defaultRule",
            description="If the ACL policy of the OSS bucket denies read access from the Internet, the configuration is considered compliant.",
            source_owner="ALIYUN",
            source_identifier="oss-bucket-public-read-prohibited",
            risk_level=1,
            tag_key_scope="For",
            tag_value_scope="example",
            region_ids_scope=default_regions.regions[0].id,
            config_rule_trigger_types="ConfigurationItemChangeNotification",
            resource_types_scopes=["ACS::OSS::Bucket"],
            rule_name="oss-bucket-public-read-prohibited")
        default_remediation = alicloud.cfg.Remediation("defaultRemediation",
            config_rule_id=default_rule.config_rule_id,
            remediation_template_id="ACS-OSS-PutBucketAcl",
            remediation_source_type="ALIYUN",
            invoke_type="MANUAL_EXECUTION",
            params=default_bucket.bucket.apply(lambda bucket: f"{{\\"bucketName\\": \\"{bucket}\\", \\"regionId\\": \\"{default_regions.regions[0].id}\\", \\"permissionName\\": \\"private\\"}}"),
            remediation_type="OOS")
        ```

        ## Import

        Config Remediation can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cfg/remediation:Remediation example <id>
        ```

        :param str resource_name: The name of the resource.
        :param RemediationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RemediationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_rule_id: Optional[pulumi.Input[str]] = None,
                 invoke_type: Optional[pulumi.Input[str]] = None,
                 params: Optional[pulumi.Input[str]] = None,
                 remediation_source_type: Optional[pulumi.Input[str]] = None,
                 remediation_template_id: Optional[pulumi.Input[str]] = None,
                 remediation_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RemediationArgs.__new__(RemediationArgs)

            if config_rule_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_rule_id'")
            __props__.__dict__["config_rule_id"] = config_rule_id
            if invoke_type is None and not opts.urn:
                raise TypeError("Missing required property 'invoke_type'")
            __props__.__dict__["invoke_type"] = invoke_type
            if params is None and not opts.urn:
                raise TypeError("Missing required property 'params'")
            __props__.__dict__["params"] = params
            __props__.__dict__["remediation_source_type"] = remediation_source_type
            if remediation_template_id is None and not opts.urn:
                raise TypeError("Missing required property 'remediation_template_id'")
            __props__.__dict__["remediation_template_id"] = remediation_template_id
            if remediation_type is None and not opts.urn:
                raise TypeError("Missing required property 'remediation_type'")
            __props__.__dict__["remediation_type"] = remediation_type
            __props__.__dict__["remediation_id"] = None
        super(Remediation, __self__).__init__(
            'alicloud:cfg/remediation:Remediation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_rule_id: Optional[pulumi.Input[str]] = None,
            invoke_type: Optional[pulumi.Input[str]] = None,
            params: Optional[pulumi.Input[str]] = None,
            remediation_id: Optional[pulumi.Input[str]] = None,
            remediation_source_type: Optional[pulumi.Input[str]] = None,
            remediation_template_id: Optional[pulumi.Input[str]] = None,
            remediation_type: Optional[pulumi.Input[str]] = None) -> 'Remediation':
        """
        Get an existing Remediation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_rule_id: Rule ID.
        :param pulumi.Input[str] invoke_type: Execution type, valid values: `Manual`, `Automatic`.
        :param pulumi.Input[str] params: Remediation parameter.
        :param pulumi.Input[str] remediation_id: Remediation ID.
        :param pulumi.Input[str] remediation_source_type: Remediation resource type, valid values: `ALIYUN` , `CUSTOMER`.
        :param pulumi.Input[str] remediation_template_id: Remediation template ID.
        :param pulumi.Input[str] remediation_type: Remediation type, valid values: `OOS`, `FC`.
               
               The following arguments will be discarded. Please use new fields as soon as possible:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RemediationState.__new__(_RemediationState)

        __props__.__dict__["config_rule_id"] = config_rule_id
        __props__.__dict__["invoke_type"] = invoke_type
        __props__.__dict__["params"] = params
        __props__.__dict__["remediation_id"] = remediation_id
        __props__.__dict__["remediation_source_type"] = remediation_source_type
        __props__.__dict__["remediation_template_id"] = remediation_template_id
        __props__.__dict__["remediation_type"] = remediation_type
        return Remediation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configRuleId")
    def config_rule_id(self) -> pulumi.Output[str]:
        """
        Rule ID.
        """
        return pulumi.get(self, "config_rule_id")

    @property
    @pulumi.getter(name="invokeType")
    def invoke_type(self) -> pulumi.Output[str]:
        """
        Execution type, valid values: `Manual`, `Automatic`.
        """
        return pulumi.get(self, "invoke_type")

    @property
    @pulumi.getter
    def params(self) -> pulumi.Output[str]:
        """
        Remediation parameter.
        """
        return pulumi.get(self, "params")

    @property
    @pulumi.getter(name="remediationId")
    def remediation_id(self) -> pulumi.Output[str]:
        """
        Remediation ID.
        """
        return pulumi.get(self, "remediation_id")

    @property
    @pulumi.getter(name="remediationSourceType")
    def remediation_source_type(self) -> pulumi.Output[str]:
        """
        Remediation resource type, valid values: `ALIYUN` , `CUSTOMER`.
        """
        return pulumi.get(self, "remediation_source_type")

    @property
    @pulumi.getter(name="remediationTemplateId")
    def remediation_template_id(self) -> pulumi.Output[str]:
        """
        Remediation template ID.
        """
        return pulumi.get(self, "remediation_template_id")

    @property
    @pulumi.getter(name="remediationType")
    def remediation_type(self) -> pulumi.Output[str]:
        """
        Remediation type, valid values: `OOS`, `FC`.

        The following arguments will be discarded. Please use new fields as soon as possible:
        """
        return pulumi.get(self, "remediation_type")

