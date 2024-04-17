# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MaliciousFileWhitelistConfigArgs', 'MaliciousFileWhitelistConfig']

@pulumi.input_type
class MaliciousFileWhitelistConfigArgs:
    def __init__(__self__, *,
                 event_name: Optional[pulumi.Input[str]] = None,
                 field: Optional[pulumi.Input[str]] = None,
                 field_value: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 target_value: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MaliciousFileWhitelistConfig resource.
        :param pulumi.Input[str] event_name: The name of the security alert associated with the representative rule.
        :param pulumi.Input[str] field: Represents the alarm associated with the resource and the white field.
        :param pulumi.Input[str] field_value: Represents the whiteout target value in effect for the resource.
        :param pulumi.Input[str] operator: The decision operator in effect on behalf of the resource.
        :param pulumi.Input[str] source: Business Source:
               - agentless: agentless detection.
        :param pulumi.Input[str] target_type: The type of target in effect on behalf of the resource.
        :param pulumi.Input[str] target_value: Represents the specific value of the target type in effect for the resource.
        """
        if event_name is not None:
            pulumi.set(__self__, "event_name", event_name)
        if field is not None:
            pulumi.set(__self__, "field", field)
        if field_value is not None:
            pulumi.set(__self__, "field_value", field_value)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if target_type is not None:
            pulumi.set(__self__, "target_type", target_type)
        if target_value is not None:
            pulumi.set(__self__, "target_value", target_value)

    @property
    @pulumi.getter(name="eventName")
    def event_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the security alert associated with the representative rule.
        """
        return pulumi.get(self, "event_name")

    @event_name.setter
    def event_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_name", value)

    @property
    @pulumi.getter
    def field(self) -> Optional[pulumi.Input[str]]:
        """
        Represents the alarm associated with the resource and the white field.
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field", value)

    @property
    @pulumi.getter(name="fieldValue")
    def field_value(self) -> Optional[pulumi.Input[str]]:
        """
        Represents the whiteout target value in effect for the resource.
        """
        return pulumi.get(self, "field_value")

    @field_value.setter
    def field_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_value", value)

    @property
    @pulumi.getter
    def operator(self) -> Optional[pulumi.Input[str]]:
        """
        The decision operator in effect on behalf of the resource.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Business Source:
        - agentless: agentless detection.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of target in effect on behalf of the resource.
        """
        return pulumi.get(self, "target_type")

    @target_type.setter
    def target_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_type", value)

    @property
    @pulumi.getter(name="targetValue")
    def target_value(self) -> Optional[pulumi.Input[str]]:
        """
        Represents the specific value of the target type in effect for the resource.
        """
        return pulumi.get(self, "target_value")

    @target_value.setter
    def target_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_value", value)


@pulumi.input_type
class _MaliciousFileWhitelistConfigState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 event_name: Optional[pulumi.Input[str]] = None,
                 field: Optional[pulumi.Input[str]] = None,
                 field_value: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 target_value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MaliciousFileWhitelistConfig resources.
        :param pulumi.Input[str] create_time: The creation time of the resource.
        :param pulumi.Input[str] event_name: The name of the security alert associated with the representative rule.
        :param pulumi.Input[str] field: Represents the alarm associated with the resource and the white field.
        :param pulumi.Input[str] field_value: Represents the whiteout target value in effect for the resource.
        :param pulumi.Input[str] operator: The decision operator in effect on behalf of the resource.
        :param pulumi.Input[str] source: Business Source:
               - agentless: agentless detection.
        :param pulumi.Input[str] target_type: The type of target in effect on behalf of the resource.
        :param pulumi.Input[str] target_value: Represents the specific value of the target type in effect for the resource.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if event_name is not None:
            pulumi.set(__self__, "event_name", event_name)
        if field is not None:
            pulumi.set(__self__, "field", field)
        if field_value is not None:
            pulumi.set(__self__, "field_value", field_value)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if target_type is not None:
            pulumi.set(__self__, "target_type", target_type)
        if target_value is not None:
            pulumi.set(__self__, "target_value", target_value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the resource.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="eventName")
    def event_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the security alert associated with the representative rule.
        """
        return pulumi.get(self, "event_name")

    @event_name.setter
    def event_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_name", value)

    @property
    @pulumi.getter
    def field(self) -> Optional[pulumi.Input[str]]:
        """
        Represents the alarm associated with the resource and the white field.
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field", value)

    @property
    @pulumi.getter(name="fieldValue")
    def field_value(self) -> Optional[pulumi.Input[str]]:
        """
        Represents the whiteout target value in effect for the resource.
        """
        return pulumi.get(self, "field_value")

    @field_value.setter
    def field_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_value", value)

    @property
    @pulumi.getter
    def operator(self) -> Optional[pulumi.Input[str]]:
        """
        The decision operator in effect on behalf of the resource.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Business Source:
        - agentless: agentless detection.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of target in effect on behalf of the resource.
        """
        return pulumi.get(self, "target_type")

    @target_type.setter
    def target_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_type", value)

    @property
    @pulumi.getter(name="targetValue")
    def target_value(self) -> Optional[pulumi.Input[str]]:
        """
        Represents the specific value of the target type in effect for the resource.
        """
        return pulumi.get(self, "target_value")

    @target_value.setter
    def target_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_value", value)


class MaliciousFileWhitelistConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 event_name: Optional[pulumi.Input[str]] = None,
                 field: Optional[pulumi.Input[str]] = None,
                 field_value: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 target_value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Threat Detection Malicious File Whitelist Config resource. malicious file add whitelist config.

        For information about Threat Detection Malicious File Whitelist Config and how to use it, see [What is Malicious File Whitelist Config](https://www.alibabacloud.com/help/zh/security-center/developer-reference/api-sas-2018-12-03-createmaliciousfilewhitelistconfig/).

        > **NOTE:** Available since v1.214.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.threatdetection.MaliciousFileWhitelistConfig("default",
            operator="strEquals",
            field="fileMd6",
            target_value="123",
            target_type="SELECTION_KEY",
            event_name="123",
            source="agentless",
            field_value="sadfas")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Threat Detection Malicious File Whitelist Config can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:threatdetection/maliciousFileWhitelistConfig:MaliciousFileWhitelistConfig example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] event_name: The name of the security alert associated with the representative rule.
        :param pulumi.Input[str] field: Represents the alarm associated with the resource and the white field.
        :param pulumi.Input[str] field_value: Represents the whiteout target value in effect for the resource.
        :param pulumi.Input[str] operator: The decision operator in effect on behalf of the resource.
        :param pulumi.Input[str] source: Business Source:
               - agentless: agentless detection.
        :param pulumi.Input[str] target_type: The type of target in effect on behalf of the resource.
        :param pulumi.Input[str] target_value: Represents the specific value of the target type in effect for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MaliciousFileWhitelistConfigArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Threat Detection Malicious File Whitelist Config resource. malicious file add whitelist config.

        For information about Threat Detection Malicious File Whitelist Config and how to use it, see [What is Malicious File Whitelist Config](https://www.alibabacloud.com/help/zh/security-center/developer-reference/api-sas-2018-12-03-createmaliciousfilewhitelistconfig/).

        > **NOTE:** Available since v1.214.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.threatdetection.MaliciousFileWhitelistConfig("default",
            operator="strEquals",
            field="fileMd6",
            target_value="123",
            target_type="SELECTION_KEY",
            event_name="123",
            source="agentless",
            field_value="sadfas")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Threat Detection Malicious File Whitelist Config can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:threatdetection/maliciousFileWhitelistConfig:MaliciousFileWhitelistConfig example <id>
        ```

        :param str resource_name: The name of the resource.
        :param MaliciousFileWhitelistConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MaliciousFileWhitelistConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 event_name: Optional[pulumi.Input[str]] = None,
                 field: Optional[pulumi.Input[str]] = None,
                 field_value: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 target_value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MaliciousFileWhitelistConfigArgs.__new__(MaliciousFileWhitelistConfigArgs)

            __props__.__dict__["event_name"] = event_name
            __props__.__dict__["field"] = field
            __props__.__dict__["field_value"] = field_value
            __props__.__dict__["operator"] = operator
            __props__.__dict__["source"] = source
            __props__.__dict__["target_type"] = target_type
            __props__.__dict__["target_value"] = target_value
            __props__.__dict__["create_time"] = None
        super(MaliciousFileWhitelistConfig, __self__).__init__(
            'alicloud:threatdetection/maliciousFileWhitelistConfig:MaliciousFileWhitelistConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            event_name: Optional[pulumi.Input[str]] = None,
            field: Optional[pulumi.Input[str]] = None,
            field_value: Optional[pulumi.Input[str]] = None,
            operator: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            target_type: Optional[pulumi.Input[str]] = None,
            target_value: Optional[pulumi.Input[str]] = None) -> 'MaliciousFileWhitelistConfig':
        """
        Get an existing MaliciousFileWhitelistConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The creation time of the resource.
        :param pulumi.Input[str] event_name: The name of the security alert associated with the representative rule.
        :param pulumi.Input[str] field: Represents the alarm associated with the resource and the white field.
        :param pulumi.Input[str] field_value: Represents the whiteout target value in effect for the resource.
        :param pulumi.Input[str] operator: The decision operator in effect on behalf of the resource.
        :param pulumi.Input[str] source: Business Source:
               - agentless: agentless detection.
        :param pulumi.Input[str] target_type: The type of target in effect on behalf of the resource.
        :param pulumi.Input[str] target_value: Represents the specific value of the target type in effect for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MaliciousFileWhitelistConfigState.__new__(_MaliciousFileWhitelistConfigState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["event_name"] = event_name
        __props__.__dict__["field"] = field
        __props__.__dict__["field_value"] = field_value
        __props__.__dict__["operator"] = operator
        __props__.__dict__["source"] = source
        __props__.__dict__["target_type"] = target_type
        __props__.__dict__["target_value"] = target_value
        return MaliciousFileWhitelistConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of the resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="eventName")
    def event_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the security alert associated with the representative rule.
        """
        return pulumi.get(self, "event_name")

    @property
    @pulumi.getter
    def field(self) -> pulumi.Output[Optional[str]]:
        """
        Represents the alarm associated with the resource and the white field.
        """
        return pulumi.get(self, "field")

    @property
    @pulumi.getter(name="fieldValue")
    def field_value(self) -> pulumi.Output[Optional[str]]:
        """
        Represents the whiteout target value in effect for the resource.
        """
        return pulumi.get(self, "field_value")

    @property
    @pulumi.getter
    def operator(self) -> pulumi.Output[Optional[str]]:
        """
        The decision operator in effect on behalf of the resource.
        """
        return pulumi.get(self, "operator")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[Optional[str]]:
        """
        Business Source:
        - agentless: agentless detection.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of target in effect on behalf of the resource.
        """
        return pulumi.get(self, "target_type")

    @property
    @pulumi.getter(name="targetValue")
    def target_value(self) -> pulumi.Output[Optional[str]]:
        """
        Represents the specific value of the target type in effect for the resource.
        """
        return pulumi.get(self, "target_value")

