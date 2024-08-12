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

__all__ = ['RdsParameterGroupArgs', 'RdsParameterGroup']

@pulumi.input_type
class RdsParameterGroupArgs:
    def __init__(__self__, *,
                 engine: pulumi.Input[str],
                 engine_version: pulumi.Input[str],
                 param_details: pulumi.Input[Sequence[pulumi.Input['RdsParameterGroupParamDetailArgs']]],
                 parameter_group_name: pulumi.Input[str],
                 parameter_group_desc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RdsParameterGroup resource.
        :param pulumi.Input[str] engine: The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
        :param pulumi.Input[str] engine_version: The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
        :param pulumi.Input[Sequence[pulumi.Input['RdsParameterGroupParamDetailArgs']]] param_details: Parameter list. See `param_detail` below.
        :param pulumi.Input[str] parameter_group_name: The name of the parameter template.
        :param pulumi.Input[str] parameter_group_desc: The description of the parameter template.
        """
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "engine_version", engine_version)
        pulumi.set(__self__, "param_details", param_details)
        pulumi.set(__self__, "parameter_group_name", parameter_group_name)
        if parameter_group_desc is not None:
            pulumi.set(__self__, "parameter_group_desc", parameter_group_desc)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Input[str]:
        """
        The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Input[str]:
        """
        The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
        """
        return pulumi.get(self, "engine_version")

    @engine_version.setter
    def engine_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine_version", value)

    @property
    @pulumi.getter(name="paramDetails")
    def param_details(self) -> pulumi.Input[Sequence[pulumi.Input['RdsParameterGroupParamDetailArgs']]]:
        """
        Parameter list. See `param_detail` below.
        """
        return pulumi.get(self, "param_details")

    @param_details.setter
    def param_details(self, value: pulumi.Input[Sequence[pulumi.Input['RdsParameterGroupParamDetailArgs']]]):
        pulumi.set(self, "param_details", value)

    @property
    @pulumi.getter(name="parameterGroupName")
    def parameter_group_name(self) -> pulumi.Input[str]:
        """
        The name of the parameter template.
        """
        return pulumi.get(self, "parameter_group_name")

    @parameter_group_name.setter
    def parameter_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_group_name", value)

    @property
    @pulumi.getter(name="parameterGroupDesc")
    def parameter_group_desc(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the parameter template.
        """
        return pulumi.get(self, "parameter_group_desc")

    @parameter_group_desc.setter
    def parameter_group_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_group_desc", value)


@pulumi.input_type
class _RdsParameterGroupState:
    def __init__(__self__, *,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 param_details: Optional[pulumi.Input[Sequence[pulumi.Input['RdsParameterGroupParamDetailArgs']]]] = None,
                 parameter_group_desc: Optional[pulumi.Input[str]] = None,
                 parameter_group_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RdsParameterGroup resources.
        :param pulumi.Input[str] engine: The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
        :param pulumi.Input[str] engine_version: The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
        :param pulumi.Input[Sequence[pulumi.Input['RdsParameterGroupParamDetailArgs']]] param_details: Parameter list. See `param_detail` below.
        :param pulumi.Input[str] parameter_group_desc: The description of the parameter template.
        :param pulumi.Input[str] parameter_group_name: The name of the parameter template.
        """
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if engine_version is not None:
            pulumi.set(__self__, "engine_version", engine_version)
        if param_details is not None:
            pulumi.set(__self__, "param_details", param_details)
        if parameter_group_desc is not None:
            pulumi.set(__self__, "parameter_group_desc", parameter_group_desc)
        if parameter_group_name is not None:
            pulumi.set(__self__, "parameter_group_name", parameter_group_name)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
        """
        return pulumi.get(self, "engine_version")

    @engine_version.setter
    def engine_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_version", value)

    @property
    @pulumi.getter(name="paramDetails")
    def param_details(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RdsParameterGroupParamDetailArgs']]]]:
        """
        Parameter list. See `param_detail` below.
        """
        return pulumi.get(self, "param_details")

    @param_details.setter
    def param_details(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RdsParameterGroupParamDetailArgs']]]]):
        pulumi.set(self, "param_details", value)

    @property
    @pulumi.getter(name="parameterGroupDesc")
    def parameter_group_desc(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the parameter template.
        """
        return pulumi.get(self, "parameter_group_desc")

    @parameter_group_desc.setter
    def parameter_group_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_group_desc", value)

    @property
    @pulumi.getter(name="parameterGroupName")
    def parameter_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the parameter template.
        """
        return pulumi.get(self, "parameter_group_name")

    @parameter_group_name.setter
    def parameter_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_group_name", value)


class RdsParameterGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 param_details: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RdsParameterGroupParamDetailArgs', 'RdsParameterGroupParamDetailArgsDict']]]]] = None,
                 parameter_group_desc: Optional[pulumi.Input[str]] = None,
                 parameter_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a RDS Parameter Group resource.

        For information about RDS Parameter Group and how to use it, see [What is Parameter Group](https://www.alibabacloud.com/help/en/doc-detail/144839.htm).

        > **NOTE:** Available since v1.119.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf_example"
        default = alicloud.rds.RdsParameterGroup("default",
            engine="mysql",
            engine_version="5.7",
            param_details=[
                {
                    "param_name": "back_log",
                    "param_value": "4000",
                },
                {
                    "param_name": "wait_timeout",
                    "param_value": "86460",
                },
            ],
            parameter_group_desc=name,
            parameter_group_name=name)
        ```

        ## Import

        RDS Parameter Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:rds/rdsParameterGroup:RdsParameterGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] engine: The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
        :param pulumi.Input[str] engine_version: The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RdsParameterGroupParamDetailArgs', 'RdsParameterGroupParamDetailArgsDict']]]] param_details: Parameter list. See `param_detail` below.
        :param pulumi.Input[str] parameter_group_desc: The description of the parameter template.
        :param pulumi.Input[str] parameter_group_name: The name of the parameter template.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RdsParameterGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a RDS Parameter Group resource.

        For information about RDS Parameter Group and how to use it, see [What is Parameter Group](https://www.alibabacloud.com/help/en/doc-detail/144839.htm).

        > **NOTE:** Available since v1.119.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf_example"
        default = alicloud.rds.RdsParameterGroup("default",
            engine="mysql",
            engine_version="5.7",
            param_details=[
                {
                    "param_name": "back_log",
                    "param_value": "4000",
                },
                {
                    "param_name": "wait_timeout",
                    "param_value": "86460",
                },
            ],
            parameter_group_desc=name,
            parameter_group_name=name)
        ```

        ## Import

        RDS Parameter Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:rds/rdsParameterGroup:RdsParameterGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param RdsParameterGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RdsParameterGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 param_details: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RdsParameterGroupParamDetailArgs', 'RdsParameterGroupParamDetailArgsDict']]]]] = None,
                 parameter_group_desc: Optional[pulumi.Input[str]] = None,
                 parameter_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RdsParameterGroupArgs.__new__(RdsParameterGroupArgs)

            if engine is None and not opts.urn:
                raise TypeError("Missing required property 'engine'")
            __props__.__dict__["engine"] = engine
            if engine_version is None and not opts.urn:
                raise TypeError("Missing required property 'engine_version'")
            __props__.__dict__["engine_version"] = engine_version
            if param_details is None and not opts.urn:
                raise TypeError("Missing required property 'param_details'")
            __props__.__dict__["param_details"] = param_details
            __props__.__dict__["parameter_group_desc"] = parameter_group_desc
            if parameter_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'parameter_group_name'")
            __props__.__dict__["parameter_group_name"] = parameter_group_name
        super(RdsParameterGroup, __self__).__init__(
            'alicloud:rds/rdsParameterGroup:RdsParameterGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            engine: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            param_details: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RdsParameterGroupParamDetailArgs', 'RdsParameterGroupParamDetailArgsDict']]]]] = None,
            parameter_group_desc: Optional[pulumi.Input[str]] = None,
            parameter_group_name: Optional[pulumi.Input[str]] = None) -> 'RdsParameterGroup':
        """
        Get an existing RdsParameterGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] engine: The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
        :param pulumi.Input[str] engine_version: The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RdsParameterGroupParamDetailArgs', 'RdsParameterGroupParamDetailArgsDict']]]] param_details: Parameter list. See `param_detail` below.
        :param pulumi.Input[str] parameter_group_desc: The description of the parameter template.
        :param pulumi.Input[str] parameter_group_name: The name of the parameter template.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RdsParameterGroupState.__new__(_RdsParameterGroupState)

        __props__.__dict__["engine"] = engine
        __props__.__dict__["engine_version"] = engine_version
        __props__.__dict__["param_details"] = param_details
        __props__.__dict__["parameter_group_desc"] = parameter_group_desc
        __props__.__dict__["parameter_group_name"] = parameter_group_name
        return RdsParameterGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        """
        The database engine. Valid values: `mysql`, `mariadb`, `PostgreSQL`.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Output[str]:
        """
        The version of the database engine. Valid values: mysql: `5.1`, `5.5`, `5.6`, `5.7`, `8.0`; mariadb: `10.3`; PostgreSQL: `10.0`, `11.0`, `12.0`, `13.0`, `14.0`, `15.0`.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="paramDetails")
    def param_details(self) -> pulumi.Output[Sequence['outputs.RdsParameterGroupParamDetail']]:
        """
        Parameter list. See `param_detail` below.
        """
        return pulumi.get(self, "param_details")

    @property
    @pulumi.getter(name="parameterGroupDesc")
    def parameter_group_desc(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the parameter template.
        """
        return pulumi.get(self, "parameter_group_desc")

    @property
    @pulumi.getter(name="parameterGroupName")
    def parameter_group_name(self) -> pulumi.Output[str]:
        """
        The name of the parameter template.
        """
        return pulumi.get(self, "parameter_group_name")

