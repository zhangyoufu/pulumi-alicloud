# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KvNamespaceArgs', 'KvNamespace']

@pulumi.input_type
class KvNamespaceArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 namespace: pulumi.Input[str]):
        """
        The set of arguments for constructing a KvNamespace resource.
        :param pulumi.Input[str] description: Namespace description information
        :param pulumi.Input[str] namespace: Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        Namespace description information
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class _KvNamespaceState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering KvNamespace resources.
        :param pulumi.Input[str] description: Namespace description information
        :param pulumi.Input[str] namespace: Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
        :param pulumi.Input[str] status: The status of the resource
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Namespace description information
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class KvNamespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Dcdn Kv Namespace resource.

        For information about Dcdn Kv Namespace and how to use it, see [What is Kv Namespace](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-putdcdnkvnamespace).

        > **NOTE:** Available since v1.198.0.

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
            name = "terraform-example"
        default = random.index.Integer("default",
            min=10000,
            max=99999)
        default_kv_namespace = alicloud.dcdn.KvNamespace("default",
            description=name,
            namespace=f"{name}-{default['result']}")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Dcdn Kv Namespace can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dcdn/kvNamespace:KvNamespace example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Namespace description information
        :param pulumi.Input[str] namespace: Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KvNamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Dcdn Kv Namespace resource.

        For information about Dcdn Kv Namespace and how to use it, see [What is Kv Namespace](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-putdcdnkvnamespace).

        > **NOTE:** Available since v1.198.0.

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
            name = "terraform-example"
        default = random.index.Integer("default",
            min=10000,
            max=99999)
        default_kv_namespace = alicloud.dcdn.KvNamespace("default",
            description=name,
            namespace=f"{name}-{default['result']}")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Dcdn Kv Namespace can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dcdn/kvNamespace:KvNamespace example
        ```

        :param str resource_name: The name of the resource.
        :param KvNamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KvNamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KvNamespaceArgs.__new__(KvNamespaceArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if namespace is None and not opts.urn:
                raise TypeError("Missing required property 'namespace'")
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["status"] = None
        super(KvNamespace, __self__).__init__(
            'alicloud:dcdn/kvNamespace:KvNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'KvNamespace':
        """
        Get an existing KvNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Namespace description information
        :param pulumi.Input[str] namespace: Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
        :param pulumi.Input[str] status: The status of the resource
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KvNamespaceState.__new__(_KvNamespaceState)

        __props__.__dict__["description"] = description
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["status"] = status
        return KvNamespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Namespace description information
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[str]:
        """
        Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource
        """
        return pulumi.get(self, "status")

