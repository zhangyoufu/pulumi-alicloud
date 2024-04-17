# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LogConfigArgs', 'LogConfig']

@pulumi.input_type
class LogConfigArgs:
    def __init__(__self__, *,
                 log_type: pulumi.Input[str],
                 sls_log_store: pulumi.Input[str],
                 sls_project: pulumi.Input[str]):
        """
        The set of arguments for constructing a LogConfig resource.
        :param pulumi.Input[str] log_type: The type the of log. Valid values: `PROVIDER`.
        :param pulumi.Input[str] sls_log_store: The name of the Log Store.
        :param pulumi.Input[str] sls_project: The name of the Project.
        """
        pulumi.set(__self__, "log_type", log_type)
        pulumi.set(__self__, "sls_log_store", sls_log_store)
        pulumi.set(__self__, "sls_project", sls_project)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> pulumi.Input[str]:
        """
        The type the of log. Valid values: `PROVIDER`.
        """
        return pulumi.get(self, "log_type")

    @log_type.setter
    def log_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_type", value)

    @property
    @pulumi.getter(name="slsLogStore")
    def sls_log_store(self) -> pulumi.Input[str]:
        """
        The name of the Log Store.
        """
        return pulumi.get(self, "sls_log_store")

    @sls_log_store.setter
    def sls_log_store(self, value: pulumi.Input[str]):
        pulumi.set(self, "sls_log_store", value)

    @property
    @pulumi.getter(name="slsProject")
    def sls_project(self) -> pulumi.Input[str]:
        """
        The name of the Project.
        """
        return pulumi.get(self, "sls_project")

    @sls_project.setter
    def sls_project(self, value: pulumi.Input[str]):
        pulumi.set(self, "sls_project", value)


@pulumi.input_type
class _LogConfigState:
    def __init__(__self__, *,
                 log_type: Optional[pulumi.Input[str]] = None,
                 sls_log_store: Optional[pulumi.Input[str]] = None,
                 sls_project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogConfig resources.
        :param pulumi.Input[str] log_type: The type the of log. Valid values: `PROVIDER`.
        :param pulumi.Input[str] sls_log_store: The name of the Log Store.
        :param pulumi.Input[str] sls_project: The name of the Project.
        """
        if log_type is not None:
            pulumi.set(__self__, "log_type", log_type)
        if sls_log_store is not None:
            pulumi.set(__self__, "sls_log_store", sls_log_store)
        if sls_project is not None:
            pulumi.set(__self__, "sls_project", sls_project)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type the of log. Valid values: `PROVIDER`.
        """
        return pulumi.get(self, "log_type")

    @log_type.setter
    def log_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_type", value)

    @property
    @pulumi.getter(name="slsLogStore")
    def sls_log_store(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Log Store.
        """
        return pulumi.get(self, "sls_log_store")

    @sls_log_store.setter
    def sls_log_store(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sls_log_store", value)

    @property
    @pulumi.getter(name="slsProject")
    def sls_project(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Project.
        """
        return pulumi.get(self, "sls_project")

    @sls_project.setter
    def sls_project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sls_project", value)


class LogConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 log_type: Optional[pulumi.Input[str]] = None,
                 sls_log_store: Optional[pulumi.Input[str]] = None,
                 sls_project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Api Gateway Log Config resource.

        For information about Api Gateway Log Config and how to use it, see [What is Log Config](https://www.alibabacloud.com/help/en/api-gateway/latest/api-cloudapi-2016-07-14-createlogconfig).

        > **NOTE:** Available since v1.185.0.

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
            max=99999,
            min=10000)
        example = alicloud.log.Project("example",
            project_name=f"{name}-{default['result']}",
            description=name)
        example_store = alicloud.log.Store("example",
            project_name=example.project_name,
            logstore_name=f"{name}-{default['result']}",
            shard_count=3,
            auto_split=True,
            max_split_shard_count=60,
            append_meta=True)
        example_log_config = alicloud.apigateway.LogConfig("example",
            sls_project=example.project_name,
            sls_log_store=example_store.logstore_name,
            log_type="PROVIDER")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Api Gateway Log Config can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:apigateway/logConfig:LogConfig example <log_type>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] log_type: The type the of log. Valid values: `PROVIDER`.
        :param pulumi.Input[str] sls_log_store: The name of the Log Store.
        :param pulumi.Input[str] sls_project: The name of the Project.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Api Gateway Log Config resource.

        For information about Api Gateway Log Config and how to use it, see [What is Log Config](https://www.alibabacloud.com/help/en/api-gateway/latest/api-cloudapi-2016-07-14-createlogconfig).

        > **NOTE:** Available since v1.185.0.

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
            max=99999,
            min=10000)
        example = alicloud.log.Project("example",
            project_name=f"{name}-{default['result']}",
            description=name)
        example_store = alicloud.log.Store("example",
            project_name=example.project_name,
            logstore_name=f"{name}-{default['result']}",
            shard_count=3,
            auto_split=True,
            max_split_shard_count=60,
            append_meta=True)
        example_log_config = alicloud.apigateway.LogConfig("example",
            sls_project=example.project_name,
            sls_log_store=example_store.logstore_name,
            log_type="PROVIDER")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Api Gateway Log Config can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:apigateway/logConfig:LogConfig example <log_type>
        ```

        :param str resource_name: The name of the resource.
        :param LogConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 log_type: Optional[pulumi.Input[str]] = None,
                 sls_log_store: Optional[pulumi.Input[str]] = None,
                 sls_project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogConfigArgs.__new__(LogConfigArgs)

            if log_type is None and not opts.urn:
                raise TypeError("Missing required property 'log_type'")
            __props__.__dict__["log_type"] = log_type
            if sls_log_store is None and not opts.urn:
                raise TypeError("Missing required property 'sls_log_store'")
            __props__.__dict__["sls_log_store"] = sls_log_store
            if sls_project is None and not opts.urn:
                raise TypeError("Missing required property 'sls_project'")
            __props__.__dict__["sls_project"] = sls_project
        super(LogConfig, __self__).__init__(
            'alicloud:apigateway/logConfig:LogConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            log_type: Optional[pulumi.Input[str]] = None,
            sls_log_store: Optional[pulumi.Input[str]] = None,
            sls_project: Optional[pulumi.Input[str]] = None) -> 'LogConfig':
        """
        Get an existing LogConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] log_type: The type the of log. Valid values: `PROVIDER`.
        :param pulumi.Input[str] sls_log_store: The name of the Log Store.
        :param pulumi.Input[str] sls_project: The name of the Project.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogConfigState.__new__(_LogConfigState)

        __props__.__dict__["log_type"] = log_type
        __props__.__dict__["sls_log_store"] = sls_log_store
        __props__.__dict__["sls_project"] = sls_project
        return LogConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> pulumi.Output[str]:
        """
        The type the of log. Valid values: `PROVIDER`.
        """
        return pulumi.get(self, "log_type")

    @property
    @pulumi.getter(name="slsLogStore")
    def sls_log_store(self) -> pulumi.Output[str]:
        """
        The name of the Log Store.
        """
        return pulumi.get(self, "sls_log_store")

    @property
    @pulumi.getter(name="slsProject")
    def sls_project(self) -> pulumi.Output[str]:
        """
        The name of the Project.
        """
        return pulumi.get(self, "sls_project")

