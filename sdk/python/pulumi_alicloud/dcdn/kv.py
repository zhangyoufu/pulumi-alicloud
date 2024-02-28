# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KvArgs', 'Kv']

@pulumi.input_type
class KvArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 namespace: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        The set of arguments for constructing a Kv resource.
        :param pulumi.Input[str] key: The name of the key to Put, the longest 512, cannot contain spaces.
        :param pulumi.Input[str] namespace: The name specified when the customer calls PutDcdnKvNamespace.
        :param pulumi.Input[str] value: The content of key, up to 2M(2*1000*1000).
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The name of the key to Put, the longest 512, cannot contain spaces.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        The name specified when the customer calls PutDcdnKvNamespace.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The content of key, up to 2M(2*1000*1000).
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class _KvState:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Kv resources.
        :param pulumi.Input[str] key: The name of the key to Put, the longest 512, cannot contain spaces.
        :param pulumi.Input[str] namespace: The name specified when the customer calls PutDcdnKvNamespace.
        :param pulumi.Input[str] value: The content of key, up to 2M(2*1000*1000).
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the key to Put, the longest 512, cannot contain spaces.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The name specified when the customer calls PutDcdnKvNamespace.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The content of key, up to 2M(2*1000*1000).
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class Kv(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Dcdn Kv resource.

        For information about Dcdn Kv and how to use it, see [What is Kv](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-putdcdnkv).

        > **NOTE:** Available since v1.198.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_random_integer = random.RandomInteger("defaultRandomInteger",
            min=10000,
            max=99999)
        default_kv_namespace = alicloud.dcdn.KvNamespace("defaultKvNamespace",
            description=name,
            namespace=default_random_integer.result.apply(lambda result: f"{name}-{result}"))
        default_kv = alicloud.dcdn.Kv("defaultKv",
            value="example-value",
            key=default_random_integer.result.apply(lambda result: f"{name}-{result}"),
            namespace=default_kv_namespace.namespace)
        ```

        ## Import

        Dcdn Kv can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dcdn/kv:Kv example <namespace>:<key>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: The name of the key to Put, the longest 512, cannot contain spaces.
        :param pulumi.Input[str] namespace: The name specified when the customer calls PutDcdnKvNamespace.
        :param pulumi.Input[str] value: The content of key, up to 2M(2*1000*1000).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KvArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Dcdn Kv resource.

        For information about Dcdn Kv and how to use it, see [What is Kv](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-putdcdnkv).

        > **NOTE:** Available since v1.198.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_random_integer = random.RandomInteger("defaultRandomInteger",
            min=10000,
            max=99999)
        default_kv_namespace = alicloud.dcdn.KvNamespace("defaultKvNamespace",
            description=name,
            namespace=default_random_integer.result.apply(lambda result: f"{name}-{result}"))
        default_kv = alicloud.dcdn.Kv("defaultKv",
            value="example-value",
            key=default_random_integer.result.apply(lambda result: f"{name}-{result}"),
            namespace=default_kv_namespace.namespace)
        ```

        ## Import

        Dcdn Kv can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dcdn/kv:Kv example <namespace>:<key>
        ```

        :param str resource_name: The name of the resource.
        :param KvArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KvArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KvArgs.__new__(KvArgs)

            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            if namespace is None and not opts.urn:
                raise TypeError("Missing required property 'namespace'")
            __props__.__dict__["namespace"] = namespace
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
        super(Kv, __self__).__init__(
            'alicloud:dcdn/kv:Kv',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'Kv':
        """
        Get an existing Kv resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: The name of the key to Put, the longest 512, cannot contain spaces.
        :param pulumi.Input[str] namespace: The name specified when the customer calls PutDcdnKvNamespace.
        :param pulumi.Input[str] value: The content of key, up to 2M(2*1000*1000).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KvState.__new__(_KvState)

        __props__.__dict__["key"] = key
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["value"] = value
        return Kv(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The name of the key to Put, the longest 512, cannot contain spaces.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[str]:
        """
        The name specified when the customer calls PutDcdnKvNamespace.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The content of key, up to 2M(2*1000*1000).
        """
        return pulumi.get(self, "value")

