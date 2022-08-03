# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AppAttachmentArgs', 'AppAttachment']

@pulumi.input_type
class AppAttachmentArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 app_id: pulumi.Input[str],
                 group_id: pulumi.Input[str],
                 stage_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a AppAttachment resource.
        :param pulumi.Input[str] api_id: The api_id that app apply to access.
        :param pulumi.Input[str] app_id: The app that apply to the authorization.
        :param pulumi.Input[str] group_id: The group that the api belongs to.
        :param pulumi.Input[str] stage_name: Stage that the app apply to access.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "stage_name", stage_name)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        """
        The api_id that app apply to access.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        """
        The app that apply to the authorization.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        The group that the api belongs to.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> pulumi.Input[str]:
        """
        Stage that the app apply to access.
        """
        return pulumi.get(self, "stage_name")

    @stage_name.setter
    def stage_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "stage_name", value)


@pulumi.input_type
class _AppAttachmentState:
    def __init__(__self__, *,
                 api_id: Optional[pulumi.Input[str]] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 stage_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppAttachment resources.
        :param pulumi.Input[str] api_id: The api_id that app apply to access.
        :param pulumi.Input[str] app_id: The app that apply to the authorization.
        :param pulumi.Input[str] group_id: The group that the api belongs to.
        :param pulumi.Input[str] stage_name: Stage that the app apply to access.
        """
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if stage_name is not None:
            pulumi.set(__self__, "stage_name", stage_name)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        The api_id that app apply to access.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        The app that apply to the authorization.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The group that the api belongs to.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> Optional[pulumi.Input[str]]:
        """
        Stage that the app apply to access.
        """
        return pulumi.get(self, "stage_name")

    @stage_name.setter
    def stage_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stage_name", value)


class AppAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 stage_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        foo = alicloud.apigateway.AppAttachment("foo",
            api_id="d29d25b9cfdf4742b1a3f6537299a749",
            app_id="20898181",
            group_id="aaef8cdbb404420f9398a74ed1db7fff",
            stage_name="PRE")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The api_id that app apply to access.
        :param pulumi.Input[str] app_id: The app that apply to the authorization.
        :param pulumi.Input[str] group_id: The group that the api belongs to.
        :param pulumi.Input[str] stage_name: Stage that the app apply to access.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        foo = alicloud.apigateway.AppAttachment("foo",
            api_id="d29d25b9cfdf4742b1a3f6537299a749",
            app_id="20898181",
            group_id="aaef8cdbb404420f9398a74ed1db7fff",
            stage_name="PRE")
        ```

        :param str resource_name: The name of the resource.
        :param AppAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 stage_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppAttachmentArgs.__new__(AppAttachmentArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            if stage_name is None and not opts.urn:
                raise TypeError("Missing required property 'stage_name'")
            __props__.__dict__["stage_name"] = stage_name
        super(AppAttachment, __self__).__init__(
            'alicloud:apigateway/appAttachment:AppAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            stage_name: Optional[pulumi.Input[str]] = None) -> 'AppAttachment':
        """
        Get an existing AppAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The api_id that app apply to access.
        :param pulumi.Input[str] app_id: The app that apply to the authorization.
        :param pulumi.Input[str] group_id: The group that the api belongs to.
        :param pulumi.Input[str] stage_name: Stage that the app apply to access.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppAttachmentState.__new__(_AppAttachmentState)

        __props__.__dict__["api_id"] = api_id
        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["stage_name"] = stage_name
        return AppAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        """
        The api_id that app apply to access.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        The app that apply to the authorization.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        The group that the api belongs to.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> pulumi.Output[str]:
        """
        Stage that the app apply to access.
        """
        return pulumi.get(self, "stage_name")

