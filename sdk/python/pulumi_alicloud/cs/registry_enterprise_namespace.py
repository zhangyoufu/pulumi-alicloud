# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['RegistryEnterpriseNamespace']


class RegistryEnterpriseNamespace(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_create: Optional[pulumi.Input[bool]] = None,
                 default_visibility: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        This resource will help you to manager Container Registry Enterprise Edition namespaces.

        For information about Container Registry Enterprise Edition namespaces and how to use it, see [Create a Namespace](https://www.alibabacloud.com/help/doc-detail/145483.htm)

        > **NOTE:** Available in v1.86.0+.

        > **NOTE:** You need to set your registry password in Container Registry Enterprise Edition console before use this resource.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        my_namespace = alicloud.cs.RegistryEnterpriseNamespace("my-namespace",
            auto_create=False,
            default_visibility="PUBLIC",
            instance_id="cri-xxx")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create: Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
        :param pulumi.Input[str] default_visibility: `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
        :param pulumi.Input[str] instance_id: ID of Container Registry Enterprise Edition instance.
        :param pulumi.Input[str] name: Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if auto_create is None:
                raise TypeError("Missing required property 'auto_create'")
            __props__['auto_create'] = auto_create
            if default_visibility is None:
                raise TypeError("Missing required property 'default_visibility'")
            __props__['default_visibility'] = default_visibility
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['name'] = name
        super(RegistryEnterpriseNamespace, __self__).__init__(
            'alicloud:cs/registryEnterpriseNamespace:RegistryEnterpriseNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_create: Optional[pulumi.Input[bool]] = None,
            default_visibility: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'RegistryEnterpriseNamespace':
        """
        Get an existing RegistryEnterpriseNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create: Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
        :param pulumi.Input[str] default_visibility: `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
        :param pulumi.Input[str] instance_id: ID of Container Registry Enterprise Edition instance.
        :param pulumi.Input[str] name: Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_create"] = auto_create
        __props__["default_visibility"] = default_visibility
        __props__["instance_id"] = instance_id
        __props__["name"] = name
        return RegistryEnterpriseNamespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoCreate")
    def auto_create(self) -> bool:
        """
        Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
        """
        return pulumi.get(self, "auto_create")

    @property
    @pulumi.getter(name="defaultVisibility")
    def default_visibility(self) -> str:
        """
        `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
        """
        return pulumi.get(self, "default_visibility")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        ID of Container Registry Enterprise Edition instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

