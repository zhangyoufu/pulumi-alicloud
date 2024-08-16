# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ChangeSetParameterArgs',
    'StackGroupParameterArgs',
    'StackInstanceParameterOverrideArgs',
    'StackParameterArgs',
    'TemplateScratchPreferenceParameterArgs',
    'TemplateScratchSourceResourceArgs',
    'TemplateScratchSourceResourceGroupArgs',
    'TemplateScratchSourceTagArgs',
]

@pulumi.input_type
class ChangeSetParameterArgs:
    def __init__(__self__, *,
                 parameter_key: pulumi.Input[str],
                 parameter_value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] parameter_key: The parameter key.
        :param pulumi.Input[str] parameter_value: The parameter value.
        """
        pulumi.set(__self__, "parameter_key", parameter_key)
        pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> pulumi.Input[str]:
        """
        The parameter key.
        """
        return pulumi.get(self, "parameter_key")

    @parameter_key.setter
    def parameter_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_key", value)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> pulumi.Input[str]:
        """
        The parameter value.
        """
        return pulumi.get(self, "parameter_value")

    @parameter_value.setter
    def parameter_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_value", value)


@pulumi.input_type
class StackGroupParameterArgs:
    def __init__(__self__, *,
                 parameter_key: Optional[pulumi.Input[str]] = None,
                 parameter_value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] parameter_key: The parameter key.
        :param pulumi.Input[str] parameter_value: The parameter value.
        """
        if parameter_key is not None:
            pulumi.set(__self__, "parameter_key", parameter_key)
        if parameter_value is not None:
            pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> Optional[pulumi.Input[str]]:
        """
        The parameter key.
        """
        return pulumi.get(self, "parameter_key")

    @parameter_key.setter
    def parameter_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_key", value)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> Optional[pulumi.Input[str]]:
        """
        The parameter value.
        """
        return pulumi.get(self, "parameter_value")

    @parameter_value.setter
    def parameter_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_value", value)


@pulumi.input_type
class StackInstanceParameterOverrideArgs:
    def __init__(__self__, *,
                 parameter_key: Optional[pulumi.Input[str]] = None,
                 parameter_value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] parameter_key: The key of override parameter. If you do not specify the key and value of the parameter, ROS uses the key and value that you specified when you created the stack group.
        :param pulumi.Input[str] parameter_value: The value of override parameter. If you do not specify the key and value of the parameter, ROS uses the key and value that you specified when you created the stack group.
        """
        if parameter_key is not None:
            pulumi.set(__self__, "parameter_key", parameter_key)
        if parameter_value is not None:
            pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> Optional[pulumi.Input[str]]:
        """
        The key of override parameter. If you do not specify the key and value of the parameter, ROS uses the key and value that you specified when you created the stack group.
        """
        return pulumi.get(self, "parameter_key")

    @parameter_key.setter
    def parameter_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_key", value)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of override parameter. If you do not specify the key and value of the parameter, ROS uses the key and value that you specified when you created the stack group.
        """
        return pulumi.get(self, "parameter_value")

    @parameter_value.setter
    def parameter_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_value", value)


@pulumi.input_type
class StackParameterArgs:
    def __init__(__self__, *,
                 parameter_value: pulumi.Input[str],
                 parameter_key: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] parameter_value: The parameter value.
        :param pulumi.Input[str] parameter_key: The parameter key.
        """
        pulumi.set(__self__, "parameter_value", parameter_value)
        if parameter_key is not None:
            pulumi.set(__self__, "parameter_key", parameter_key)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> pulumi.Input[str]:
        """
        The parameter value.
        """
        return pulumi.get(self, "parameter_value")

    @parameter_value.setter
    def parameter_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_value", value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> Optional[pulumi.Input[str]]:
        """
        The parameter key.
        """
        return pulumi.get(self, "parameter_key")

    @parameter_key.setter
    def parameter_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_key", value)


@pulumi.input_type
class TemplateScratchPreferenceParameterArgs:
    def __init__(__self__, *,
                 parameter_key: pulumi.Input[str],
                 parameter_value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] parameter_key: Priority parameter key. For more information about values, see [supplementary instructions for request parameters](https://www.alibabacloud.com/help/zh/doc-detail/358846.html#h2-url-4).
        :param pulumi.Input[str] parameter_value: Priority parameter value. For more information about values, see [supplementary instructions for request parameters](https://www.alibabacloud.com/help/zh/doc-detail/358846.html#h2-url-4).
        """
        pulumi.set(__self__, "parameter_key", parameter_key)
        pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="parameterKey")
    def parameter_key(self) -> pulumi.Input[str]:
        """
        Priority parameter key. For more information about values, see [supplementary instructions for request parameters](https://www.alibabacloud.com/help/zh/doc-detail/358846.html#h2-url-4).
        """
        return pulumi.get(self, "parameter_key")

    @parameter_key.setter
    def parameter_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_key", value)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> pulumi.Input[str]:
        """
        Priority parameter value. For more information about values, see [supplementary instructions for request parameters](https://www.alibabacloud.com/help/zh/doc-detail/358846.html#h2-url-4).
        """
        return pulumi.get(self, "parameter_value")

    @parameter_value.setter
    def parameter_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_value", value)


@pulumi.input_type
class TemplateScratchSourceResourceArgs:
    def __init__(__self__, *,
                 resource_id: pulumi.Input[str],
                 resource_type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] resource_id: The ID of the Source Resource.
        :param pulumi.Input[str] resource_type: The type of the Source resource.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the Source Resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        """
        The type of the Source resource.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)


@pulumi.input_type
class TemplateScratchSourceResourceGroupArgs:
    def __init__(__self__, *,
                 resource_group_id: pulumi.Input[str],
                 resource_type_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] resource_group_id: The ID of the Source Resource Group.
        """
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if resource_type_filters is not None:
            pulumi.set(__self__, "resource_type_filters", resource_type_filters)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Input[str]:
        """
        The ID of the Source Resource Group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="resourceTypeFilters")
    def resource_type_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "resource_type_filters")

    @resource_type_filters.setter
    def resource_type_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_type_filters", value)


@pulumi.input_type
class TemplateScratchSourceTagArgs:
    def __init__(__self__, *,
                 resource_tags: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 resource_type_filters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] resource_tags: Source label. **NOTE:** A maximum of 10 source labels can be configured.
        """
        pulumi.set(__self__, "resource_tags", resource_tags)
        if resource_type_filters is not None:
            pulumi.set(__self__, "resource_type_filters", resource_type_filters)

    @property
    @pulumi.getter(name="resourceTags")
    def resource_tags(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        Source label. **NOTE:** A maximum of 10 source labels can be configured.
        """
        return pulumi.get(self, "resource_tags")

    @resource_tags.setter
    def resource_tags(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "resource_tags", value)

    @property
    @pulumi.getter(name="resourceTypeFilters")
    def resource_type_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "resource_type_filters")

    @resource_type_filters.setter
    def resource_type_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_type_filters", value)


