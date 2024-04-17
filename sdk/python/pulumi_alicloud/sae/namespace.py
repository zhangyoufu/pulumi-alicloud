# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NamespaceArgs', 'Namespace']

@pulumi.input_type
class NamespaceArgs:
    def __init__(__self__, *,
                 namespace_name: pulumi.Input[str],
                 enable_micro_registration: Optional[pulumi.Input[bool]] = None,
                 namespace_description: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 namespace_short_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Namespace resource.
        :param pulumi.Input[str] namespace_name: The Name of Namespace.
        :param pulumi.Input[bool] enable_micro_registration: Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
        :param pulumi.Input[str] namespace_description: The Description of Namespace.
        :param pulumi.Input[str] namespace_id: The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
        :param pulumi.Input[str] namespace_short_id: The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
        """
        pulumi.set(__self__, "namespace_name", namespace_name)
        if enable_micro_registration is not None:
            pulumi.set(__self__, "enable_micro_registration", enable_micro_registration)
        if namespace_description is not None:
            pulumi.set(__self__, "namespace_description", namespace_description)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if namespace_short_id is not None:
            pulumi.set(__self__, "namespace_short_id", namespace_short_id)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[str]:
        """
        The Name of Namespace.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="enableMicroRegistration")
    def enable_micro_registration(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
        """
        return pulumi.get(self, "enable_micro_registration")

    @enable_micro_registration.setter
    def enable_micro_registration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_micro_registration", value)

    @property
    @pulumi.getter(name="namespaceDescription")
    def namespace_description(self) -> Optional[pulumi.Input[str]]:
        """
        The Description of Namespace.
        """
        return pulumi.get(self, "namespace_description")

    @namespace_description.setter
    def namespace_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_description", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="namespaceShortId")
    def namespace_short_id(self) -> Optional[pulumi.Input[str]]:
        """
        The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
        """
        return pulumi.get(self, "namespace_short_id")

    @namespace_short_id.setter
    def namespace_short_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_short_id", value)


@pulumi.input_type
class _NamespaceState:
    def __init__(__self__, *,
                 enable_micro_registration: Optional[pulumi.Input[bool]] = None,
                 namespace_description: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 namespace_short_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Namespace resources.
        :param pulumi.Input[bool] enable_micro_registration: Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
        :param pulumi.Input[str] namespace_description: The Description of Namespace.
        :param pulumi.Input[str] namespace_id: The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
        :param pulumi.Input[str] namespace_name: The Name of Namespace.
        :param pulumi.Input[str] namespace_short_id: The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
        """
        if enable_micro_registration is not None:
            pulumi.set(__self__, "enable_micro_registration", enable_micro_registration)
        if namespace_description is not None:
            pulumi.set(__self__, "namespace_description", namespace_description)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if namespace_name is not None:
            pulumi.set(__self__, "namespace_name", namespace_name)
        if namespace_short_id is not None:
            pulumi.set(__self__, "namespace_short_id", namespace_short_id)

    @property
    @pulumi.getter(name="enableMicroRegistration")
    def enable_micro_registration(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
        """
        return pulumi.get(self, "enable_micro_registration")

    @enable_micro_registration.setter
    def enable_micro_registration(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_micro_registration", value)

    @property
    @pulumi.getter(name="namespaceDescription")
    def namespace_description(self) -> Optional[pulumi.Input[str]]:
        """
        The Description of Namespace.
        """
        return pulumi.get(self, "namespace_description")

    @namespace_description.setter
    def namespace_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_description", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of Namespace.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="namespaceShortId")
    def namespace_short_id(self) -> Optional[pulumi.Input[str]]:
        """
        The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
        """
        return pulumi.get(self, "namespace_short_id")

    @namespace_short_id.setter
    def namespace_short_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_short_id", value)


class Namespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_micro_registration: Optional[pulumi.Input[bool]] = None,
                 namespace_description: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 namespace_short_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Serverless App Engine (SAE) Namespace resource.

        For information about SAE Namespace and how to use it, see [What is Namespace](https://www.alibabacloud.com/help/en/sae/latest/createnamespace).

        > **NOTE:** Available since v1.129.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.get_regions(current=True)
        default_integer = random.index.Integer("default",
            max=99999,
            min=10000)
        example = alicloud.sae.Namespace("example",
            namespace_id=f"{default.regions[0].id}:example{default_integer['result']}",
            namespace_name=name,
            namespace_description=name,
            enable_micro_registration=False)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Serverless App Engine (SAE) Namespace can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:sae/namespace:Namespace example <namespace_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_micro_registration: Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
        :param pulumi.Input[str] namespace_description: The Description of Namespace.
        :param pulumi.Input[str] namespace_id: The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
        :param pulumi.Input[str] namespace_name: The Name of Namespace.
        :param pulumi.Input[str] namespace_short_id: The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Serverless App Engine (SAE) Namespace resource.

        For information about SAE Namespace and how to use it, see [What is Namespace](https://www.alibabacloud.com/help/en/sae/latest/createnamespace).

        > **NOTE:** Available since v1.129.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.get_regions(current=True)
        default_integer = random.index.Integer("default",
            max=99999,
            min=10000)
        example = alicloud.sae.Namespace("example",
            namespace_id=f"{default.regions[0].id}:example{default_integer['result']}",
            namespace_name=name,
            namespace_description=name,
            enable_micro_registration=False)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Serverless App Engine (SAE) Namespace can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:sae/namespace:Namespace example <namespace_id>
        ```

        :param str resource_name: The name of the resource.
        :param NamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_micro_registration: Optional[pulumi.Input[bool]] = None,
                 namespace_description: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 namespace_short_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NamespaceArgs.__new__(NamespaceArgs)

            __props__.__dict__["enable_micro_registration"] = enable_micro_registration
            __props__.__dict__["namespace_description"] = namespace_description
            __props__.__dict__["namespace_id"] = namespace_id
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__.__dict__["namespace_name"] = namespace_name
            __props__.__dict__["namespace_short_id"] = namespace_short_id
        super(Namespace, __self__).__init__(
            'alicloud:sae/namespace:Namespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enable_micro_registration: Optional[pulumi.Input[bool]] = None,
            namespace_description: Optional[pulumi.Input[str]] = None,
            namespace_id: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            namespace_short_id: Optional[pulumi.Input[str]] = None) -> 'Namespace':
        """
        Get an existing Namespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_micro_registration: Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
        :param pulumi.Input[str] namespace_description: The Description of Namespace.
        :param pulumi.Input[str] namespace_id: The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
        :param pulumi.Input[str] namespace_name: The Name of Namespace.
        :param pulumi.Input[str] namespace_short_id: The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NamespaceState.__new__(_NamespaceState)

        __props__.__dict__["enable_micro_registration"] = enable_micro_registration
        __props__.__dict__["namespace_description"] = namespace_description
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["namespace_name"] = namespace_name
        __props__.__dict__["namespace_short_id"] = namespace_short_id
        return Namespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enableMicroRegistration")
    def enable_micro_registration(self) -> pulumi.Output[bool]:
        """
        Specifies whether to enable the SAE built-in registry. If you do not use the built-in registry, you can set `enable_micro_registration` to `false` to accelerate the creation of the namespace. Default value: `true`. Valid values:
        """
        return pulumi.get(self, "enable_micro_registration")

    @property
    @pulumi.getter(name="namespaceDescription")
    def namespace_description(self) -> pulumi.Output[Optional[str]]:
        """
        The Description of Namespace.
        """
        return pulumi.get(self, "namespace_description")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[str]:
        """
        The ID of the Namespace. It can contain 2 to 32 lowercase characters. The value is in format `{RegionId}:{namespace}`.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        The Name of Namespace.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="namespaceShortId")
    def namespace_short_id(self) -> pulumi.Output[str]:
        """
        The short ID of the Namespace. You do not need to specify a region ID. The value of `namespace_short_id` can be up to 20 characters in length and can contain only lowercase letters and digits.
        """
        return pulumi.get(self, "namespace_short_id")

