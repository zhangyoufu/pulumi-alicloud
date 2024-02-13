# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ChartNamespaceArgs', 'ChartNamespace']

@pulumi.input_type
class ChartNamespaceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 namespace_name: pulumi.Input[str],
                 auto_create_repo: Optional[pulumi.Input[bool]] = None,
                 default_repo_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ChartNamespace resource.
        :param pulumi.Input[str] instance_id: The ID of the Container Registry instance.
        :param pulumi.Input[str] namespace_name: The name of the namespace that you want to create.
        :param pulumi.Input[bool] auto_create_repo: Specifies whether to automatically create repositories in the namespace. Valid values:
        :param pulumi.Input[str] default_repo_type: DefaultRepoType. Valid values: `PRIVATE`, `PUBLIC`.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "namespace_name", namespace_name)
        if auto_create_repo is not None:
            pulumi.set(__self__, "auto_create_repo", auto_create_repo)
        if default_repo_type is not None:
            pulumi.set(__self__, "default_repo_type", default_repo_type)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the Container Registry instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[str]:
        """
        The name of the namespace that you want to create.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="autoCreateRepo")
    def auto_create_repo(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to automatically create repositories in the namespace. Valid values:
        """
        return pulumi.get(self, "auto_create_repo")

    @auto_create_repo.setter
    def auto_create_repo(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_create_repo", value)

    @property
    @pulumi.getter(name="defaultRepoType")
    def default_repo_type(self) -> Optional[pulumi.Input[str]]:
        """
        DefaultRepoType. Valid values: `PRIVATE`, `PUBLIC`.
        """
        return pulumi.get(self, "default_repo_type")

    @default_repo_type.setter
    def default_repo_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_repo_type", value)


@pulumi.input_type
class _ChartNamespaceState:
    def __init__(__self__, *,
                 auto_create_repo: Optional[pulumi.Input[bool]] = None,
                 default_repo_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ChartNamespace resources.
        :param pulumi.Input[bool] auto_create_repo: Specifies whether to automatically create repositories in the namespace. Valid values:
        :param pulumi.Input[str] default_repo_type: DefaultRepoType. Valid values: `PRIVATE`, `PUBLIC`.
        :param pulumi.Input[str] instance_id: The ID of the Container Registry instance.
        :param pulumi.Input[str] namespace_name: The name of the namespace that you want to create.
        """
        if auto_create_repo is not None:
            pulumi.set(__self__, "auto_create_repo", auto_create_repo)
        if default_repo_type is not None:
            pulumi.set(__self__, "default_repo_type", default_repo_type)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if namespace_name is not None:
            pulumi.set(__self__, "namespace_name", namespace_name)

    @property
    @pulumi.getter(name="autoCreateRepo")
    def auto_create_repo(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to automatically create repositories in the namespace. Valid values:
        """
        return pulumi.get(self, "auto_create_repo")

    @auto_create_repo.setter
    def auto_create_repo(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_create_repo", value)

    @property
    @pulumi.getter(name="defaultRepoType")
    def default_repo_type(self) -> Optional[pulumi.Input[str]]:
        """
        DefaultRepoType. Valid values: `PRIVATE`, `PUBLIC`.
        """
        return pulumi.get(self, "default_repo_type")

    @default_repo_type.setter
    def default_repo_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_repo_type", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Container Registry instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the namespace that you want to create.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_name", value)


class ChartNamespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_create_repo: Optional[pulumi.Input[bool]] = None,
                 default_repo_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a CR Chart Namespace resource.

        For information about CR Chart Namespace and how to use it, see [What is Chart Namespace](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createchartnamespace).

        > **NOTE:** Available since v1.149.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "example-name"
        example_registry_enterprise_instance = alicloud.cr.RegistryEnterpriseInstance("exampleRegistryEnterpriseInstance",
            payment_type="Subscription",
            period=1,
            renew_period=0,
            renewal_status="ManualRenewal",
            instance_type="Advanced",
            instance_name=name)
        example_chart_namespace = alicloud.cr.ChartNamespace("exampleChartNamespace",
            instance_id=example_registry_enterprise_instance.id,
            namespace_name=name)
        ```

        ## Import

        CR Chart Namespace can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cr/chartNamespace:ChartNamespace example <instance_id>:<namespace_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_repo: Specifies whether to automatically create repositories in the namespace. Valid values:
        :param pulumi.Input[str] default_repo_type: DefaultRepoType. Valid values: `PRIVATE`, `PUBLIC`.
        :param pulumi.Input[str] instance_id: The ID of the Container Registry instance.
        :param pulumi.Input[str] namespace_name: The name of the namespace that you want to create.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ChartNamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CR Chart Namespace resource.

        For information about CR Chart Namespace and how to use it, see [What is Chart Namespace](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createchartnamespace).

        > **NOTE:** Available since v1.149.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "example-name"
        example_registry_enterprise_instance = alicloud.cr.RegistryEnterpriseInstance("exampleRegistryEnterpriseInstance",
            payment_type="Subscription",
            period=1,
            renew_period=0,
            renewal_status="ManualRenewal",
            instance_type="Advanced",
            instance_name=name)
        example_chart_namespace = alicloud.cr.ChartNamespace("exampleChartNamespace",
            instance_id=example_registry_enterprise_instance.id,
            namespace_name=name)
        ```

        ## Import

        CR Chart Namespace can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cr/chartNamespace:ChartNamespace example <instance_id>:<namespace_name>
        ```

        :param str resource_name: The name of the resource.
        :param ChartNamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ChartNamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_create_repo: Optional[pulumi.Input[bool]] = None,
                 default_repo_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ChartNamespaceArgs.__new__(ChartNamespaceArgs)

            __props__.__dict__["auto_create_repo"] = auto_create_repo
            __props__.__dict__["default_repo_type"] = default_repo_type
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__.__dict__["namespace_name"] = namespace_name
        super(ChartNamespace, __self__).__init__(
            'alicloud:cr/chartNamespace:ChartNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_create_repo: Optional[pulumi.Input[bool]] = None,
            default_repo_type: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None) -> 'ChartNamespace':
        """
        Get an existing ChartNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_repo: Specifies whether to automatically create repositories in the namespace. Valid values:
        :param pulumi.Input[str] default_repo_type: DefaultRepoType. Valid values: `PRIVATE`, `PUBLIC`.
        :param pulumi.Input[str] instance_id: The ID of the Container Registry instance.
        :param pulumi.Input[str] namespace_name: The name of the namespace that you want to create.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ChartNamespaceState.__new__(_ChartNamespaceState)

        __props__.__dict__["auto_create_repo"] = auto_create_repo
        __props__.__dict__["default_repo_type"] = default_repo_type
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["namespace_name"] = namespace_name
        return ChartNamespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoCreateRepo")
    def auto_create_repo(self) -> pulumi.Output[bool]:
        """
        Specifies whether to automatically create repositories in the namespace. Valid values:
        """
        return pulumi.get(self, "auto_create_repo")

    @property
    @pulumi.getter(name="defaultRepoType")
    def default_repo_type(self) -> pulumi.Output[str]:
        """
        DefaultRepoType. Valid values: `PRIVATE`, `PUBLIC`.
        """
        return pulumi.get(self, "default_repo_type")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the Container Registry instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        The name of the namespace that you want to create.
        """
        return pulumi.get(self, "namespace_name")

