# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EnvFeatureArgs', 'EnvFeature']

@pulumi.input_type
class EnvFeatureArgs:
    def __init__(__self__, *,
                 env_feature_name: pulumi.Input[str],
                 environment_id: pulumi.Input[str],
                 feature_version: pulumi.Input[str]):
        """
        The set of arguments for constructing a EnvFeature resource.
        :param pulumi.Input[str] env_feature_name: The name of the resource.
        :param pulumi.Input[str] environment_id: The first ID of the resource.
        :param pulumi.Input[str] feature_version: Version information of the Feature. You can query Feature information by using ListEnvironmentFeatures.
        """
        pulumi.set(__self__, "env_feature_name", env_feature_name)
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "feature_version", feature_version)

    @property
    @pulumi.getter(name="envFeatureName")
    def env_feature_name(self) -> pulumi.Input[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "env_feature_name")

    @env_feature_name.setter
    def env_feature_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "env_feature_name", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[str]:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="featureVersion")
    def feature_version(self) -> pulumi.Input[str]:
        """
        Version information of the Feature. You can query Feature information by using ListEnvironmentFeatures.
        """
        return pulumi.get(self, "feature_version")

    @feature_version.setter
    def feature_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "feature_version", value)


@pulumi.input_type
class _EnvFeatureState:
    def __init__(__self__, *,
                 env_feature_name: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 feature_version: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EnvFeature resources.
        :param pulumi.Input[str] env_feature_name: The name of the resource.
        :param pulumi.Input[str] environment_id: The first ID of the resource.
        :param pulumi.Input[str] feature_version: Version information of the Feature. You can query Feature information by using ListEnvironmentFeatures.
        :param pulumi.Input[str] namespace: Namespace.
        :param pulumi.Input[str] status: Status.
        """
        if env_feature_name is not None:
            pulumi.set(__self__, "env_feature_name", env_feature_name)
        if environment_id is not None:
            pulumi.set(__self__, "environment_id", environment_id)
        if feature_version is not None:
            pulumi.set(__self__, "feature_version", feature_version)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="envFeatureName")
    def env_feature_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "env_feature_name")

    @env_feature_name.setter
    def env_feature_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "env_feature_name", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="featureVersion")
    def feature_version(self) -> Optional[pulumi.Input[str]]:
        """
        Version information of the Feature. You can query Feature information by using ListEnvironmentFeatures.
        """
        return pulumi.get(self, "feature_version")

    @feature_version.setter
    def feature_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "feature_version", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Namespace.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class EnvFeature(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 env_feature_name: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 feature_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ARMS Env Feature resource. Feature of the arms environment.

        For information about ARMS Env Feature and how to use it, see [What is Env Feature](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-installenvironmentfeature).

        > **NOTE:** Available since v1.212.0.

        ## Import

        ARMS Env Feature can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:arms/envFeature:EnvFeature example <environment_id>:<env_feature_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] env_feature_name: The name of the resource.
        :param pulumi.Input[str] environment_id: The first ID of the resource.
        :param pulumi.Input[str] feature_version: Version information of the Feature. You can query Feature information by using ListEnvironmentFeatures.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvFeatureArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ARMS Env Feature resource. Feature of the arms environment.

        For information about ARMS Env Feature and how to use it, see [What is Env Feature](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-installenvironmentfeature).

        > **NOTE:** Available since v1.212.0.

        ## Import

        ARMS Env Feature can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:arms/envFeature:EnvFeature example <environment_id>:<env_feature_name>
        ```

        :param str resource_name: The name of the resource.
        :param EnvFeatureArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvFeatureArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 env_feature_name: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 feature_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvFeatureArgs.__new__(EnvFeatureArgs)

            if env_feature_name is None and not opts.urn:
                raise TypeError("Missing required property 'env_feature_name'")
            __props__.__dict__["env_feature_name"] = env_feature_name
            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            if feature_version is None and not opts.urn:
                raise TypeError("Missing required property 'feature_version'")
            __props__.__dict__["feature_version"] = feature_version
            __props__.__dict__["namespace"] = None
            __props__.__dict__["status"] = None
        super(EnvFeature, __self__).__init__(
            'alicloud:arms/envFeature:EnvFeature',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            env_feature_name: Optional[pulumi.Input[str]] = None,
            environment_id: Optional[pulumi.Input[str]] = None,
            feature_version: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'EnvFeature':
        """
        Get an existing EnvFeature resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] env_feature_name: The name of the resource.
        :param pulumi.Input[str] environment_id: The first ID of the resource.
        :param pulumi.Input[str] feature_version: Version information of the Feature. You can query Feature information by using ListEnvironmentFeatures.
        :param pulumi.Input[str] namespace: Namespace.
        :param pulumi.Input[str] status: Status.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnvFeatureState.__new__(_EnvFeatureState)

        __props__.__dict__["env_feature_name"] = env_feature_name
        __props__.__dict__["environment_id"] = environment_id
        __props__.__dict__["feature_version"] = feature_version
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["status"] = status
        return EnvFeature(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="envFeatureName")
    def env_feature_name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "env_feature_name")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[str]:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="featureVersion")
    def feature_version(self) -> pulumi.Output[str]:
        """
        Version information of the Feature. You can query Feature information by using ListEnvironmentFeatures.
        """
        return pulumi.get(self, "feature_version")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[str]:
        """
        Namespace.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status.
        """
        return pulumi.get(self, "status")

