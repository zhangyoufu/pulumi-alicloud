# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SharedResourceArgs', 'SharedResource']

@pulumi.input_type
class SharedResourceArgs:
    def __init__(__self__, *,
                 resource_id: pulumi.Input[str],
                 resource_share_id: pulumi.Input[str],
                 resource_type: pulumi.Input[str]):
        """
        The set of arguments for constructing a SharedResource resource.
        :param pulumi.Input[str] resource_id: The resource ID need shared.
        :param pulumi.Input[str] resource_share_id: The resource share ID of resource manager.
        :param pulumi.Input[str] resource_type: The resource type of should shared, valid value `VSwitch`. The following types are added after v1.173.0: `ROSTemplate` and `ServiceCatalogPortfolio`.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_share_id", resource_share_id)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The resource ID need shared.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceShareId")
    def resource_share_id(self) -> pulumi.Input[str]:
        """
        The resource share ID of resource manager.
        """
        return pulumi.get(self, "resource_share_id")

    @resource_share_id.setter
    def resource_share_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_share_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        """
        The resource type of should shared, valid value `VSwitch`. The following types are added after v1.173.0: `ROSTemplate` and `ServiceCatalogPortfolio`.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)


@pulumi.input_type
class _SharedResourceState:
    def __init__(__self__, *,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_share_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SharedResource resources.
        :param pulumi.Input[str] resource_id: The resource ID need shared.
        :param pulumi.Input[str] resource_share_id: The resource share ID of resource manager.
        :param pulumi.Input[str] resource_type: The resource type of should shared, valid value `VSwitch`. The following types are added after v1.173.0: `ROSTemplate` and `ServiceCatalogPortfolio`.
        :param pulumi.Input[str] status: status.
        """
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_share_id is not None:
            pulumi.set(__self__, "resource_share_id", resource_share_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID need shared.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceShareId")
    def resource_share_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource share ID of resource manager.
        """
        return pulumi.get(self, "resource_share_id")

    @resource_share_id.setter
    def resource_share_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_share_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The resource type of should shared, valid value `VSwitch`. The following types are added after v1.173.0: `ROSTemplate` and `ServiceCatalogPortfolio`.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class SharedResource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_share_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Resource Manager Shared Resource resource.

        For information about Resource Manager Shared Resource and how to use it, see [What is Shared Resource](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).

        > **NOTE:** Available in v1.111.0+.

        ## Import

        Resource Manager Shared Resource can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:resourcemanager/sharedResource:SharedResource example <resource_share_id>:<resource_id>:<resource_type>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_id: The resource ID need shared.
        :param pulumi.Input[str] resource_share_id: The resource share ID of resource manager.
        :param pulumi.Input[str] resource_type: The resource type of should shared, valid value `VSwitch`. The following types are added after v1.173.0: `ROSTemplate` and `ServiceCatalogPortfolio`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SharedResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Resource Manager Shared Resource resource.

        For information about Resource Manager Shared Resource and how to use it, see [What is Shared Resource](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).

        > **NOTE:** Available in v1.111.0+.

        ## Import

        Resource Manager Shared Resource can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:resourcemanager/sharedResource:SharedResource example <resource_share_id>:<resource_id>:<resource_type>
        ```

        :param str resource_name: The name of the resource.
        :param SharedResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SharedResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_share_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SharedResourceArgs.__new__(SharedResourceArgs)

            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            if resource_share_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_share_id'")
            __props__.__dict__["resource_share_id"] = resource_share_id
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["status"] = None
        super(SharedResource, __self__).__init__(
            'alicloud:resourcemanager/sharedResource:SharedResource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            resource_share_id: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'SharedResource':
        """
        Get an existing SharedResource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_id: The resource ID need shared.
        :param pulumi.Input[str] resource_share_id: The resource share ID of resource manager.
        :param pulumi.Input[str] resource_type: The resource type of should shared, valid value `VSwitch`. The following types are added after v1.173.0: `ROSTemplate` and `ServiceCatalogPortfolio`.
        :param pulumi.Input[str] status: status.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SharedResourceState.__new__(_SharedResourceState)

        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_share_id"] = resource_share_id
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["status"] = status
        return SharedResource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The resource ID need shared.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceShareId")
    def resource_share_id(self) -> pulumi.Output[str]:
        """
        The resource share ID of resource manager.
        """
        return pulumi.get(self, "resource_share_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        """
        The resource type of should shared, valid value `VSwitch`. The following types are added after v1.173.0: `ROSTemplate` and `ServiceCatalogPortfolio`.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        status.
        """
        return pulumi.get(self, "status")

