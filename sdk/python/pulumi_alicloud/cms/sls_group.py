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

__all__ = ['SlsGroupArgs', 'SlsGroup']

@pulumi.input_type
class SlsGroupArgs:
    def __init__(__self__, *,
                 sls_group_configs: pulumi.Input[Sequence[pulumi.Input['SlsGroupSlsGroupConfigArgs']]],
                 sls_group_name: pulumi.Input[str],
                 sls_group_description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SlsGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input['SlsGroupSlsGroupConfigArgs']]] sls_group_configs: The Config of the Sls Group. You can specify up to 25 Config. See `sls_group_config` below.
        :param pulumi.Input[str] sls_group_name: The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
        :param pulumi.Input[str] sls_group_description: The Description of the Sls Group.
        """
        pulumi.set(__self__, "sls_group_configs", sls_group_configs)
        pulumi.set(__self__, "sls_group_name", sls_group_name)
        if sls_group_description is not None:
            pulumi.set(__self__, "sls_group_description", sls_group_description)

    @property
    @pulumi.getter(name="slsGroupConfigs")
    def sls_group_configs(self) -> pulumi.Input[Sequence[pulumi.Input['SlsGroupSlsGroupConfigArgs']]]:
        """
        The Config of the Sls Group. You can specify up to 25 Config. See `sls_group_config` below.
        """
        return pulumi.get(self, "sls_group_configs")

    @sls_group_configs.setter
    def sls_group_configs(self, value: pulumi.Input[Sequence[pulumi.Input['SlsGroupSlsGroupConfigArgs']]]):
        pulumi.set(self, "sls_group_configs", value)

    @property
    @pulumi.getter(name="slsGroupName")
    def sls_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
        """
        return pulumi.get(self, "sls_group_name")

    @sls_group_name.setter
    def sls_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "sls_group_name", value)

    @property
    @pulumi.getter(name="slsGroupDescription")
    def sls_group_description(self) -> Optional[pulumi.Input[str]]:
        """
        The Description of the Sls Group.
        """
        return pulumi.get(self, "sls_group_description")

    @sls_group_description.setter
    def sls_group_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sls_group_description", value)


@pulumi.input_type
class _SlsGroupState:
    def __init__(__self__, *,
                 sls_group_configs: Optional[pulumi.Input[Sequence[pulumi.Input['SlsGroupSlsGroupConfigArgs']]]] = None,
                 sls_group_description: Optional[pulumi.Input[str]] = None,
                 sls_group_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SlsGroup resources.
        :param pulumi.Input[Sequence[pulumi.Input['SlsGroupSlsGroupConfigArgs']]] sls_group_configs: The Config of the Sls Group. You can specify up to 25 Config. See `sls_group_config` below.
        :param pulumi.Input[str] sls_group_description: The Description of the Sls Group.
        :param pulumi.Input[str] sls_group_name: The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
        """
        if sls_group_configs is not None:
            pulumi.set(__self__, "sls_group_configs", sls_group_configs)
        if sls_group_description is not None:
            pulumi.set(__self__, "sls_group_description", sls_group_description)
        if sls_group_name is not None:
            pulumi.set(__self__, "sls_group_name", sls_group_name)

    @property
    @pulumi.getter(name="slsGroupConfigs")
    def sls_group_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SlsGroupSlsGroupConfigArgs']]]]:
        """
        The Config of the Sls Group. You can specify up to 25 Config. See `sls_group_config` below.
        """
        return pulumi.get(self, "sls_group_configs")

    @sls_group_configs.setter
    def sls_group_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SlsGroupSlsGroupConfigArgs']]]]):
        pulumi.set(self, "sls_group_configs", value)

    @property
    @pulumi.getter(name="slsGroupDescription")
    def sls_group_description(self) -> Optional[pulumi.Input[str]]:
        """
        The Description of the Sls Group.
        """
        return pulumi.get(self, "sls_group_description")

    @sls_group_description.setter
    def sls_group_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sls_group_description", value)

    @property
    @pulumi.getter(name="slsGroupName")
    def sls_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
        """
        return pulumi.get(self, "sls_group_name")

    @sls_group_name.setter
    def sls_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sls_group_name", value)


class SlsGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 sls_group_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SlsGroupSlsGroupConfigArgs']]]]] = None,
                 sls_group_description: Optional[pulumi.Input[str]] = None,
                 sls_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cloud Monitor Service Sls Group resource.

        For information about Cloud Monitor Service Sls Group and how to use it, see [What is Sls Group](https://www.alibabacloud.com/help/doc-detail/28608.htm).

        > **NOTE:** Available since v1.171.0.

        ## Import

        Cloud Monitor Service Sls Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cms/slsGroup:SlsGroup example <sls_group_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SlsGroupSlsGroupConfigArgs']]]] sls_group_configs: The Config of the Sls Group. You can specify up to 25 Config. See `sls_group_config` below.
        :param pulumi.Input[str] sls_group_description: The Description of the Sls Group.
        :param pulumi.Input[str] sls_group_name: The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SlsGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Monitor Service Sls Group resource.

        For information about Cloud Monitor Service Sls Group and how to use it, see [What is Sls Group](https://www.alibabacloud.com/help/doc-detail/28608.htm).

        > **NOTE:** Available since v1.171.0.

        ## Import

        Cloud Monitor Service Sls Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cms/slsGroup:SlsGroup example <sls_group_name>
        ```

        :param str resource_name: The name of the resource.
        :param SlsGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SlsGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 sls_group_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SlsGroupSlsGroupConfigArgs']]]]] = None,
                 sls_group_description: Optional[pulumi.Input[str]] = None,
                 sls_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SlsGroupArgs.__new__(SlsGroupArgs)

            if sls_group_configs is None and not opts.urn:
                raise TypeError("Missing required property 'sls_group_configs'")
            __props__.__dict__["sls_group_configs"] = sls_group_configs
            __props__.__dict__["sls_group_description"] = sls_group_description
            if sls_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'sls_group_name'")
            __props__.__dict__["sls_group_name"] = sls_group_name
        super(SlsGroup, __self__).__init__(
            'alicloud:cms/slsGroup:SlsGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            sls_group_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SlsGroupSlsGroupConfigArgs']]]]] = None,
            sls_group_description: Optional[pulumi.Input[str]] = None,
            sls_group_name: Optional[pulumi.Input[str]] = None) -> 'SlsGroup':
        """
        Get an existing SlsGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SlsGroupSlsGroupConfigArgs']]]] sls_group_configs: The Config of the Sls Group. You can specify up to 25 Config. See `sls_group_config` below.
        :param pulumi.Input[str] sls_group_description: The Description of the Sls Group.
        :param pulumi.Input[str] sls_group_name: The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SlsGroupState.__new__(_SlsGroupState)

        __props__.__dict__["sls_group_configs"] = sls_group_configs
        __props__.__dict__["sls_group_description"] = sls_group_description
        __props__.__dict__["sls_group_name"] = sls_group_name
        return SlsGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="slsGroupConfigs")
    def sls_group_configs(self) -> pulumi.Output[Sequence['outputs.SlsGroupSlsGroupConfig']]:
        """
        The Config of the Sls Group. You can specify up to 25 Config. See `sls_group_config` below.
        """
        return pulumi.get(self, "sls_group_configs")

    @property
    @pulumi.getter(name="slsGroupDescription")
    def sls_group_description(self) -> pulumi.Output[Optional[str]]:
        """
        The Description of the Sls Group.
        """
        return pulumi.get(self, "sls_group_description")

    @property
    @pulumi.getter(name="slsGroupName")
    def sls_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
        """
        return pulumi.get(self, "sls_group_name")

