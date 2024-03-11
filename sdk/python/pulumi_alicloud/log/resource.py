# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ResourceArgs', 'Resource']

@pulumi.input_type
class ResourceArgs:
    def __init__(__self__, *,
                 schema: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ext_info: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Resource resource.
        :param pulumi.Input[str] schema: The meta store's schema info, which is json string format, used to define table's fields.
        :param pulumi.Input[str] type: The meta store's type, userdefine e.g.
        :param pulumi.Input[str] description: The meta store's description.
        :param pulumi.Input[str] ext_info: The ext info of meta store.
        :param pulumi.Input[str] name: The meta store's name, can be used as table name.
        """
        pulumi.set(__self__, "schema", schema)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ext_info is not None:
            pulumi.set(__self__, "ext_info", ext_info)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Input[str]:
        """
        The meta store's schema info, which is json string format, used to define table's fields.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: pulumi.Input[str]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The meta store's type, userdefine e.g.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The meta store's description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="extInfo")
    def ext_info(self) -> Optional[pulumi.Input[str]]:
        """
        The ext info of meta store.
        """
        return pulumi.get(self, "ext_info")

    @ext_info.setter
    def ext_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ext_info", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The meta store's name, can be used as table name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ResourceState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 ext_info: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Resource resources.
        :param pulumi.Input[str] description: The meta store's description.
        :param pulumi.Input[str] ext_info: The ext info of meta store.
        :param pulumi.Input[str] name: The meta store's name, can be used as table name.
        :param pulumi.Input[str] schema: The meta store's schema info, which is json string format, used to define table's fields.
        :param pulumi.Input[str] type: The meta store's type, userdefine e.g.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ext_info is not None:
            pulumi.set(__self__, "ext_info", ext_info)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The meta store's description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="extInfo")
    def ext_info(self) -> Optional[pulumi.Input[str]]:
        """
        The ext info of meta store.
        """
        return pulumi.get(self, "ext_info")

    @ext_info.setter
    def ext_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ext_info", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The meta store's name, can be used as table name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[str]]:
        """
        The meta store's schema info, which is json string format, used to define table's fields.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The meta store's type, userdefine e.g.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Resource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ext_info: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Log resource is a meta store service provided by log service, resource can be used to define meta store's table structure.

        For information about SLS Resource and how to use it, see [Resource management](https://www.alibabacloud.com/help/en/doc-detail/207732.html)

        > **NOTE:** Available since v1.162.0. log resource region should be set a main region: cn-heyuan.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.log.Resource("example",
            description="user tf resource desc",
            ext_info="{}",
            schema=\"\"\"    {
              "schema": [
                {
                  "column": "col1",
                  "desc": "col1   desc",
                  "ext_info": {
                  },
                  "required": true,
                  "type": "string"
                },
                {
                  "column": "col2",
                  "desc": "col2   desc",
                  "ext_info": "optional",
                  "required": true,
                  "type": "string"
                }
              ]
            }
          
        \"\"\",
            type="userdefine")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Log resource can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:log/resource:Resource example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The meta store's description.
        :param pulumi.Input[str] ext_info: The ext info of meta store.
        :param pulumi.Input[str] name: The meta store's name, can be used as table name.
        :param pulumi.Input[str] schema: The meta store's schema info, which is json string format, used to define table's fields.
        :param pulumi.Input[str] type: The meta store's type, userdefine e.g.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Log resource is a meta store service provided by log service, resource can be used to define meta store's table structure.

        For information about SLS Resource and how to use it, see [Resource management](https://www.alibabacloud.com/help/en/doc-detail/207732.html)

        > **NOTE:** Available since v1.162.0. log resource region should be set a main region: cn-heyuan.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.log.Resource("example",
            description="user tf resource desc",
            ext_info="{}",
            schema=\"\"\"    {
              "schema": [
                {
                  "column": "col1",
                  "desc": "col1   desc",
                  "ext_info": {
                  },
                  "required": true,
                  "type": "string"
                },
                {
                  "column": "col2",
                  "desc": "col2   desc",
                  "ext_info": "optional",
                  "required": true,
                  "type": "string"
                }
              ]
            }
          
        \"\"\",
            type="userdefine")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Log resource can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:log/resource:Resource example <id>
        ```

        :param str resource_name: The name of the resource.
        :param ResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ext_info: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceArgs.__new__(ResourceArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["ext_info"] = ext_info
            __props__.__dict__["name"] = name
            if schema is None and not opts.urn:
                raise TypeError("Missing required property 'schema'")
            __props__.__dict__["schema"] = schema
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(Resource, __self__).__init__(
            'alicloud:log/resource:Resource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            ext_info: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            schema: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Resource':
        """
        Get an existing Resource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The meta store's description.
        :param pulumi.Input[str] ext_info: The ext info of meta store.
        :param pulumi.Input[str] name: The meta store's name, can be used as table name.
        :param pulumi.Input[str] schema: The meta store's schema info, which is json string format, used to define table's fields.
        :param pulumi.Input[str] type: The meta store's type, userdefine e.g.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourceState.__new__(_ResourceState)

        __props__.__dict__["description"] = description
        __props__.__dict__["ext_info"] = ext_info
        __props__.__dict__["name"] = name
        __props__.__dict__["schema"] = schema
        __props__.__dict__["type"] = type
        return Resource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The meta store's description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="extInfo")
    def ext_info(self) -> pulumi.Output[Optional[str]]:
        """
        The ext info of meta store.
        """
        return pulumi.get(self, "ext_info")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The meta store's name, can be used as table name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[str]:
        """
        The meta store's schema info, which is json string format, used to define table's fields.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The meta store's type, userdefine e.g.
        """
        return pulumi.get(self, "type")

