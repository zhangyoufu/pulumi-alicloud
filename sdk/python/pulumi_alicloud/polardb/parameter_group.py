# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ParameterGroupArgs', 'ParameterGroup']

@pulumi.input_type
class ParameterGroupArgs:
    def __init__(__self__, *,
                 db_type: pulumi.Input[str],
                 db_version: pulumi.Input[str],
                 parameters: pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ParameterGroup resource.
        :param pulumi.Input[str] db_type: The type of the database engine. Only `MySQL` is supported.
        :param pulumi.Input[str] db_version: The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
        :param pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]] parameters: The parameter template. See the following `Block parameters`.
        :param pulumi.Input[str] description: The description of the parameter template. It must be 0 to 200 characters in length.
        :param pulumi.Input[str] name: The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
        """
        pulumi.set(__self__, "db_type", db_type)
        pulumi.set(__self__, "db_version", db_version)
        pulumi.set(__self__, "parameters", parameters)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="dbType")
    def db_type(self) -> pulumi.Input[str]:
        """
        The type of the database engine. Only `MySQL` is supported.
        """
        return pulumi.get(self, "db_type")

    @db_type.setter
    def db_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_type", value)

    @property
    @pulumi.getter(name="dbVersion")
    def db_version(self) -> pulumi.Input[str]:
        """
        The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
        """
        return pulumi.get(self, "db_version")

    @db_version.setter
    def db_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_version", value)

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]]:
        """
        The parameter template. See the following `Block parameters`.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the parameter template. It must be 0 to 200 characters in length.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ParameterGroupState:
    def __init__(__self__, *,
                 db_type: Optional[pulumi.Input[str]] = None,
                 db_version: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]]] = None):
        """
        Input properties used for looking up and filtering ParameterGroup resources.
        :param pulumi.Input[str] db_type: The type of the database engine. Only `MySQL` is supported.
        :param pulumi.Input[str] db_version: The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
        :param pulumi.Input[str] description: The description of the parameter template. It must be 0 to 200 characters in length.
        :param pulumi.Input[str] name: The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
        :param pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]] parameters: The parameter template. See the following `Block parameters`.
        """
        if db_type is not None:
            pulumi.set(__self__, "db_type", db_type)
        if db_version is not None:
            pulumi.set(__self__, "db_version", db_version)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="dbType")
    def db_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the database engine. Only `MySQL` is supported.
        """
        return pulumi.get(self, "db_type")

    @db_type.setter
    def db_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_type", value)

    @property
    @pulumi.getter(name="dbVersion")
    def db_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
        """
        return pulumi.get(self, "db_version")

    @db_version.setter
    def db_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_version", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the parameter template. It must be 0 to 200 characters in length.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]]]:
        """
        The parameter template. See the following `Block parameters`.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]]]):
        pulumi.set(self, "parameters", value)


class ParameterGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_type: Optional[pulumi.Input[str]] = None,
                 db_version: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParameterArgs']]]]] = None,
                 __props__=None):
        """
        Provides a PolarDB Parameter Group resource.

        For information about PolarDB Parameter Group and how to use it, see [What is Parameter Group](https://www.alibabacloud.com/help/en/polardb/polardb-for-mysql/user-guide/apply-a-parameter-template).

        > **NOTE:** Available in v1.183.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.polardb.ParameterGroup("example",
            db_type="MySQL",
            db_version="8.0",
            description="example_value",
            parameters=[alicloud.polardb.ParameterGroupParameterArgs(
                param_name="wait_timeout",
                param_value="86400",
            )])
        ```

        ## Import

        PolarDB Parameter Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:polardb/parameterGroup:ParameterGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_type: The type of the database engine. Only `MySQL` is supported.
        :param pulumi.Input[str] db_version: The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
        :param pulumi.Input[str] description: The description of the parameter template. It must be 0 to 200 characters in length.
        :param pulumi.Input[str] name: The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParameterArgs']]]] parameters: The parameter template. See the following `Block parameters`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ParameterGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a PolarDB Parameter Group resource.

        For information about PolarDB Parameter Group and how to use it, see [What is Parameter Group](https://www.alibabacloud.com/help/en/polardb/polardb-for-mysql/user-guide/apply-a-parameter-template).

        > **NOTE:** Available in v1.183.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.polardb.ParameterGroup("example",
            db_type="MySQL",
            db_version="8.0",
            description="example_value",
            parameters=[alicloud.polardb.ParameterGroupParameterArgs(
                param_name="wait_timeout",
                param_value="86400",
            )])
        ```

        ## Import

        PolarDB Parameter Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:polardb/parameterGroup:ParameterGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param ParameterGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ParameterGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_type: Optional[pulumi.Input[str]] = None,
                 db_version: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParameterArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ParameterGroupArgs.__new__(ParameterGroupArgs)

            if db_type is None and not opts.urn:
                raise TypeError("Missing required property 'db_type'")
            __props__.__dict__["db_type"] = db_type
            if db_version is None and not opts.urn:
                raise TypeError("Missing required property 'db_version'")
            __props__.__dict__["db_version"] = db_version
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if parameters is None and not opts.urn:
                raise TypeError("Missing required property 'parameters'")
            __props__.__dict__["parameters"] = parameters
        super(ParameterGroup, __self__).__init__(
            'alicloud:polardb/parameterGroup:ParameterGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            db_type: Optional[pulumi.Input[str]] = None,
            db_version: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParameterArgs']]]]] = None) -> 'ParameterGroup':
        """
        Get an existing ParameterGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_type: The type of the database engine. Only `MySQL` is supported.
        :param pulumi.Input[str] db_version: The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
        :param pulumi.Input[str] description: The description of the parameter template. It must be 0 to 200 characters in length.
        :param pulumi.Input[str] name: The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParameterArgs']]]] parameters: The parameter template. See the following `Block parameters`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ParameterGroupState.__new__(_ParameterGroupState)

        __props__.__dict__["db_type"] = db_type
        __props__.__dict__["db_version"] = db_version
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["parameters"] = parameters
        return ParameterGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dbType")
    def db_type(self) -> pulumi.Output[str]:
        """
        The type of the database engine. Only `MySQL` is supported.
        """
        return pulumi.get(self, "db_type")

    @property
    @pulumi.getter(name="dbVersion")
    def db_version(self) -> pulumi.Output[str]:
        """
        The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
        """
        return pulumi.get(self, "db_version")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the parameter template. It must be 0 to 200 characters in length.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Sequence['outputs.ParameterGroupParameter']]:
        """
        The parameter template. See the following `Block parameters`.
        """
        return pulumi.get(self, "parameters")

