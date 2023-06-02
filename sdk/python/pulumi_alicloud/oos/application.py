# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 application_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Application resource.
        :param pulumi.Input[str] application_name: The name of the application.
        :param pulumi.Input[str] description: Application group description information.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the resource.
        """
        pulumi.set(__self__, "application_name", application_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Input[str]:
        """
        The name of the application.
        """
        return pulumi.get(self, "application_name")

    @application_name.setter
    def application_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Application group description information.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tag of the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ApplicationState:
    def __init__(__self__, *,
                 application_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Application resources.
        :param pulumi.Input[str] application_name: The name of the application.
        :param pulumi.Input[str] description: Application group description information.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the resource.
        """
        if application_name is not None:
            pulumi.set(__self__, "application_name", application_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the application.
        """
        return pulumi.get(self, "application_name")

    @application_name.setter
    def application_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Application group description information.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tag of the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class Application(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provides a OOS Application resource.

        For information about OOS Application and how to use it, see [What is Application](https://www.alibabacloud.com/help/en/doc-detail/120556.html).

        > **NOTE:** Available in v1.145.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_resource_groups = alicloud.resourcemanager.get_resource_groups()
        default_application = alicloud.oos.Application("defaultApplication",
            resource_group_id=default_resource_groups.groups[0].id,
            application_name="terraform-example",
            description="terraform-example",
            tags={
                "Created": "TF",
            })
        ```

        ## Import

        OOS Application can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:oos/application:Application example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_name: The name of the application.
        :param pulumi.Input[str] description: Application group description information.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a OOS Application resource.

        For information about OOS Application and how to use it, see [What is Application](https://www.alibabacloud.com/help/en/doc-detail/120556.html).

        > **NOTE:** Available in v1.145.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_resource_groups = alicloud.resourcemanager.get_resource_groups()
        default_application = alicloud.oos.Application("defaultApplication",
            resource_group_id=default_resource_groups.groups[0].id,
            application_name="terraform-example",
            description="terraform-example",
            tags={
                "Created": "TF",
            })
        ```

        ## Import

        OOS Application can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:oos/application:Application example <id>
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationArgs.__new__(ApplicationArgs)

            if application_name is None and not opts.urn:
                raise TypeError("Missing required property 'application_name'")
            __props__.__dict__["application_name"] = application_name
            __props__.__dict__["description"] = description
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["tags"] = tags
        super(Application, __self__).__init__(
            'alicloud:oos/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_name: The name of the application.
        :param pulumi.Input[str] description: Application group description information.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationState.__new__(_ApplicationState)

        __props__.__dict__["application_name"] = application_name
        __props__.__dict__["description"] = description
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["tags"] = tags
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Output[str]:
        """
        The name of the application.
        """
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Application group description information.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The tag of the resource.
        """
        return pulumi.get(self, "tags")

