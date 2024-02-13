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

__all__ = ['AScriptArgs', 'AScript']

@pulumi.input_type
class AScriptArgs:
    def __init__(__self__, *,
                 ascript_name: pulumi.Input[str],
                 enabled: pulumi.Input[bool],
                 listener_id: pulumi.Input[str],
                 position: pulumi.Input[str],
                 script_content: pulumi.Input[str],
                 ext_attribute_enabled: Optional[pulumi.Input[bool]] = None,
                 ext_attributes: Optional[pulumi.Input[Sequence[pulumi.Input['AScriptExtAttributeArgs']]]] = None):
        """
        The set of arguments for constructing a AScript resource.
        :param pulumi.Input[str] ascript_name: The name of AScript.
        :param pulumi.Input[bool] enabled: Whether scripts are enabled.
        :param pulumi.Input[str] listener_id: Listener ID of script attribution
        :param pulumi.Input[str] position: Execution location of AScript.
        :param pulumi.Input[str] script_content: The content of AScript.
        :param pulumi.Input[bool] ext_attribute_enabled: Whether extension parameters are enabled.
        :param pulumi.Input[Sequence[pulumi.Input['AScriptExtAttributeArgs']]] ext_attributes: Extended attribute list. See `ext_attributes` below for details.
        """
        pulumi.set(__self__, "ascript_name", ascript_name)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "listener_id", listener_id)
        pulumi.set(__self__, "position", position)
        pulumi.set(__self__, "script_content", script_content)
        if ext_attribute_enabled is not None:
            pulumi.set(__self__, "ext_attribute_enabled", ext_attribute_enabled)
        if ext_attributes is not None:
            pulumi.set(__self__, "ext_attributes", ext_attributes)

    @property
    @pulumi.getter(name="ascriptName")
    def ascript_name(self) -> pulumi.Input[str]:
        """
        The name of AScript.
        """
        return pulumi.get(self, "ascript_name")

    @ascript_name.setter
    def ascript_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "ascript_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Whether scripts are enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Input[str]:
        """
        Listener ID of script attribution
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter
    def position(self) -> pulumi.Input[str]:
        """
        Execution location of AScript.
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: pulumi.Input[str]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="scriptContent")
    def script_content(self) -> pulumi.Input[str]:
        """
        The content of AScript.
        """
        return pulumi.get(self, "script_content")

    @script_content.setter
    def script_content(self, value: pulumi.Input[str]):
        pulumi.set(self, "script_content", value)

    @property
    @pulumi.getter(name="extAttributeEnabled")
    def ext_attribute_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether extension parameters are enabled.
        """
        return pulumi.get(self, "ext_attribute_enabled")

    @ext_attribute_enabled.setter
    def ext_attribute_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ext_attribute_enabled", value)

    @property
    @pulumi.getter(name="extAttributes")
    def ext_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AScriptExtAttributeArgs']]]]:
        """
        Extended attribute list. See `ext_attributes` below for details.
        """
        return pulumi.get(self, "ext_attributes")

    @ext_attributes.setter
    def ext_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AScriptExtAttributeArgs']]]]):
        pulumi.set(self, "ext_attributes", value)


@pulumi.input_type
class _AScriptState:
    def __init__(__self__, *,
                 ascript_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ext_attribute_enabled: Optional[pulumi.Input[bool]] = None,
                 ext_attributes: Optional[pulumi.Input[Sequence[pulumi.Input['AScriptExtAttributeArgs']]]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 load_balancer_id: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[str]] = None,
                 script_content: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AScript resources.
        :param pulumi.Input[str] ascript_name: The name of AScript.
        :param pulumi.Input[bool] enabled: Whether scripts are enabled.
        :param pulumi.Input[bool] ext_attribute_enabled: Whether extension parameters are enabled.
        :param pulumi.Input[Sequence[pulumi.Input['AScriptExtAttributeArgs']]] ext_attributes: Extended attribute list. See `ext_attributes` below for details.
        :param pulumi.Input[str] listener_id: Listener ID of script attribution
        :param pulumi.Input[str] load_balancer_id: The ID of load balancer instance.
        :param pulumi.Input[str] position: Execution location of AScript.
        :param pulumi.Input[str] script_content: The content of AScript.
        :param pulumi.Input[str] status: The status of AScript.
        """
        if ascript_name is not None:
            pulumi.set(__self__, "ascript_name", ascript_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if ext_attribute_enabled is not None:
            pulumi.set(__self__, "ext_attribute_enabled", ext_attribute_enabled)
        if ext_attributes is not None:
            pulumi.set(__self__, "ext_attributes", ext_attributes)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)
        if load_balancer_id is not None:
            pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if script_content is not None:
            pulumi.set(__self__, "script_content", script_content)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="ascriptName")
    def ascript_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of AScript.
        """
        return pulumi.get(self, "ascript_name")

    @ascript_name.setter
    def ascript_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ascript_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether scripts are enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="extAttributeEnabled")
    def ext_attribute_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether extension parameters are enabled.
        """
        return pulumi.get(self, "ext_attribute_enabled")

    @ext_attribute_enabled.setter
    def ext_attribute_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ext_attribute_enabled", value)

    @property
    @pulumi.getter(name="extAttributes")
    def ext_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AScriptExtAttributeArgs']]]]:
        """
        Extended attribute list. See `ext_attributes` below for details.
        """
        return pulumi.get(self, "ext_attributes")

    @ext_attributes.setter
    def ext_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AScriptExtAttributeArgs']]]]):
        pulumi.set(self, "ext_attributes", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        Listener ID of script attribution
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of load balancer instance.
        """
        return pulumi.get(self, "load_balancer_id")

    @load_balancer_id.setter
    def load_balancer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_balancer_id", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[str]]:
        """
        Execution location of AScript.
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter(name="scriptContent")
    def script_content(self) -> Optional[pulumi.Input[str]]:
        """
        The content of AScript.
        """
        return pulumi.get(self, "script_content")

    @script_content.setter
    def script_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "script_content", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of AScript.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class AScript(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ascript_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ext_attribute_enabled: Optional[pulumi.Input[bool]] = None,
                 ext_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AScriptExtAttributeArgs']]]]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[str]] = None,
                 script_content: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Alb Ascript resource.

        For information about Alb Ascript and how to use it, see [What is AScript](https://www.alibabacloud.com/help/en/slb/application-load-balancer/developer-reference/api-alb-2020-06-16-createascripts).

        > **NOTE:** Available since v1.195.0.

        ## Import

        Alb AScript can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:alb/aScript:AScript example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ascript_name: The name of AScript.
        :param pulumi.Input[bool] enabled: Whether scripts are enabled.
        :param pulumi.Input[bool] ext_attribute_enabled: Whether extension parameters are enabled.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AScriptExtAttributeArgs']]]] ext_attributes: Extended attribute list. See `ext_attributes` below for details.
        :param pulumi.Input[str] listener_id: Listener ID of script attribution
        :param pulumi.Input[str] position: Execution location of AScript.
        :param pulumi.Input[str] script_content: The content of AScript.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AScriptArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Alb Ascript resource.

        For information about Alb Ascript and how to use it, see [What is AScript](https://www.alibabacloud.com/help/en/slb/application-load-balancer/developer-reference/api-alb-2020-06-16-createascripts).

        > **NOTE:** Available since v1.195.0.

        ## Import

        Alb AScript can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:alb/aScript:AScript example <id>
        ```

        :param str resource_name: The name of the resource.
        :param AScriptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AScriptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ascript_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ext_attribute_enabled: Optional[pulumi.Input[bool]] = None,
                 ext_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AScriptExtAttributeArgs']]]]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[str]] = None,
                 script_content: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AScriptArgs.__new__(AScriptArgs)

            if ascript_name is None and not opts.urn:
                raise TypeError("Missing required property 'ascript_name'")
            __props__.__dict__["ascript_name"] = ascript_name
            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["ext_attribute_enabled"] = ext_attribute_enabled
            __props__.__dict__["ext_attributes"] = ext_attributes
            if listener_id is None and not opts.urn:
                raise TypeError("Missing required property 'listener_id'")
            __props__.__dict__["listener_id"] = listener_id
            if position is None and not opts.urn:
                raise TypeError("Missing required property 'position'")
            __props__.__dict__["position"] = position
            if script_content is None and not opts.urn:
                raise TypeError("Missing required property 'script_content'")
            __props__.__dict__["script_content"] = script_content
            __props__.__dict__["load_balancer_id"] = None
            __props__.__dict__["status"] = None
        super(AScript, __self__).__init__(
            'alicloud:alb/aScript:AScript',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ascript_name: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            ext_attribute_enabled: Optional[pulumi.Input[bool]] = None,
            ext_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AScriptExtAttributeArgs']]]]] = None,
            listener_id: Optional[pulumi.Input[str]] = None,
            load_balancer_id: Optional[pulumi.Input[str]] = None,
            position: Optional[pulumi.Input[str]] = None,
            script_content: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'AScript':
        """
        Get an existing AScript resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ascript_name: The name of AScript.
        :param pulumi.Input[bool] enabled: Whether scripts are enabled.
        :param pulumi.Input[bool] ext_attribute_enabled: Whether extension parameters are enabled.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AScriptExtAttributeArgs']]]] ext_attributes: Extended attribute list. See `ext_attributes` below for details.
        :param pulumi.Input[str] listener_id: Listener ID of script attribution
        :param pulumi.Input[str] load_balancer_id: The ID of load balancer instance.
        :param pulumi.Input[str] position: Execution location of AScript.
        :param pulumi.Input[str] script_content: The content of AScript.
        :param pulumi.Input[str] status: The status of AScript.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AScriptState.__new__(_AScriptState)

        __props__.__dict__["ascript_name"] = ascript_name
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["ext_attribute_enabled"] = ext_attribute_enabled
        __props__.__dict__["ext_attributes"] = ext_attributes
        __props__.__dict__["listener_id"] = listener_id
        __props__.__dict__["load_balancer_id"] = load_balancer_id
        __props__.__dict__["position"] = position
        __props__.__dict__["script_content"] = script_content
        __props__.__dict__["status"] = status
        return AScript(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ascriptName")
    def ascript_name(self) -> pulumi.Output[str]:
        """
        The name of AScript.
        """
        return pulumi.get(self, "ascript_name")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether scripts are enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="extAttributeEnabled")
    def ext_attribute_enabled(self) -> pulumi.Output[bool]:
        """
        Whether extension parameters are enabled.
        """
        return pulumi.get(self, "ext_attribute_enabled")

    @property
    @pulumi.getter(name="extAttributes")
    def ext_attributes(self) -> pulumi.Output[Sequence['outputs.AScriptExtAttribute']]:
        """
        Extended attribute list. See `ext_attributes` below for details.
        """
        return pulumi.get(self, "ext_attributes")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Output[str]:
        """
        Listener ID of script attribution
        """
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> pulumi.Output[str]:
        """
        The ID of load balancer instance.
        """
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter
    def position(self) -> pulumi.Output[str]:
        """
        Execution location of AScript.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter(name="scriptContent")
    def script_content(self) -> pulumi.Output[str]:
        """
        The content of AScript.
        """
        return pulumi.get(self, "script_content")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of AScript.
        """
        return pulumi.get(self, "status")

