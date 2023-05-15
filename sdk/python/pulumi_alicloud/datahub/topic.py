# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TopicArgs', 'Topic']

@pulumi.input_type
class TopicArgs:
    def __init__(__self__, *,
                 project_name: pulumi.Input[str],
                 comment: Optional[pulumi.Input[str]] = None,
                 life_cycle: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 record_schema: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 record_type: Optional[pulumi.Input[str]] = None,
                 shard_count: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Topic resource.
        :param pulumi.Input[str] project_name: The name of the datahub project that this topic belongs to. It is case-insensitive.
        :param pulumi.Input[str] comment: Comment of the datahub topic. It cannot be longer than 255 characters.
               
               **Notes:** Currently `life_cycle` can not be modified and it will be supported in the next future.
        :param pulumi.Input[int] life_cycle: How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
        :param pulumi.Input[str] name: The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        :param pulumi.Input[Mapping[str, Any]] record_schema: Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
               - BIGINT
               - STRING
               - BOOLEAN
               - DOUBLE
               - TIMESTAMP
        :param pulumi.Input[str] record_type: The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.
        :param pulumi.Input[int] shard_count: The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
        """
        pulumi.set(__self__, "project_name", project_name)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if life_cycle is not None:
            pulumi.set(__self__, "life_cycle", life_cycle)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if record_schema is not None:
            pulumi.set(__self__, "record_schema", record_schema)
        if record_type is not None:
            pulumi.set(__self__, "record_type", record_type)
        if shard_count is not None:
            pulumi.set(__self__, "shard_count", shard_count)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Input[str]:
        """
        The name of the datahub project that this topic belongs to. It is case-insensitive.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment of the datahub topic. It cannot be longer than 255 characters.

        **Notes:** Currently `life_cycle` can not be modified and it will be supported in the next future.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="lifeCycle")
    def life_cycle(self) -> Optional[pulumi.Input[int]]:
        """
        How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
        """
        return pulumi.get(self, "life_cycle")

    @life_cycle.setter
    def life_cycle(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "life_cycle", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="recordSchema")
    def record_schema(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
        - BIGINT
        - STRING
        - BOOLEAN
        - DOUBLE
        - TIMESTAMP
        """
        return pulumi.get(self, "record_schema")

    @record_schema.setter
    def record_schema(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "record_schema", value)

    @property
    @pulumi.getter(name="recordType")
    def record_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.
        """
        return pulumi.get(self, "record_type")

    @record_type.setter
    def record_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "record_type", value)

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
        """
        return pulumi.get(self, "shard_count")

    @shard_count.setter
    def shard_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "shard_count", value)


@pulumi.input_type
class _TopicState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 last_modify_time: Optional[pulumi.Input[str]] = None,
                 life_cycle: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 record_schema: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 record_type: Optional[pulumi.Input[str]] = None,
                 shard_count: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Topic resources.
        :param pulumi.Input[str] comment: Comment of the datahub topic. It cannot be longer than 255 characters.
               
               **Notes:** Currently `life_cycle` can not be modified and it will be supported in the next future.
        :param pulumi.Input[str] create_time: Create time of the datahub topic. It is a human-readable string rather than 64-bits UTC.
        :param pulumi.Input[str] last_modify_time: Last modify time of the datahub topic. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
        :param pulumi.Input[int] life_cycle: How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
        :param pulumi.Input[str] name: The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        :param pulumi.Input[str] project_name: The name of the datahub project that this topic belongs to. It is case-insensitive.
        :param pulumi.Input[Mapping[str, Any]] record_schema: Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
               - BIGINT
               - STRING
               - BOOLEAN
               - DOUBLE
               - TIMESTAMP
        :param pulumi.Input[str] record_type: The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.
        :param pulumi.Input[int] shard_count: The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if last_modify_time is not None:
            pulumi.set(__self__, "last_modify_time", last_modify_time)
        if life_cycle is not None:
            pulumi.set(__self__, "life_cycle", life_cycle)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if record_schema is not None:
            pulumi.set(__self__, "record_schema", record_schema)
        if record_type is not None:
            pulumi.set(__self__, "record_type", record_type)
        if shard_count is not None:
            pulumi.set(__self__, "shard_count", shard_count)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment of the datahub topic. It cannot be longer than 255 characters.

        **Notes:** Currently `life_cycle` can not be modified and it will be supported in the next future.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time of the datahub topic. It is a human-readable string rather than 64-bits UTC.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="lastModifyTime")
    def last_modify_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last modify time of the datahub topic. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
        """
        return pulumi.get(self, "last_modify_time")

    @last_modify_time.setter
    def last_modify_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_modify_time", value)

    @property
    @pulumi.getter(name="lifeCycle")
    def life_cycle(self) -> Optional[pulumi.Input[int]]:
        """
        How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
        """
        return pulumi.get(self, "life_cycle")

    @life_cycle.setter
    def life_cycle(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "life_cycle", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the datahub project that this topic belongs to. It is case-insensitive.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="recordSchema")
    def record_schema(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
        - BIGINT
        - STRING
        - BOOLEAN
        - DOUBLE
        - TIMESTAMP
        """
        return pulumi.get(self, "record_schema")

    @record_schema.setter
    def record_schema(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "record_schema", value)

    @property
    @pulumi.getter(name="recordType")
    def record_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.
        """
        return pulumi.get(self, "record_type")

    @record_type.setter
    def record_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "record_type", value)

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
        """
        return pulumi.get(self, "shard_count")

    @shard_count.setter
    def shard_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "shard_count", value)


class Topic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 life_cycle: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 record_schema: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 record_type: Optional[pulumi.Input[str]] = None,
                 shard_count: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The topic is the basic unit of Datahub data source and is used to define one kind of data or stream. It contains a set of subscriptions. You can manage the datahub source of an application by using topics. [Refer to details](https://help.aliyun.com/document_detail/47440.html).

        ## Example Usage

        Basic Usage

        - BLob Topic

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.datahub.Topic("example",
            comment="created by terraform",
            life_cycle=7,
            project_name="tf_datahub_project",
            record_type="BLOB",
            shard_count=3)
        ```
        - Tuple Topic

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.datahub.Topic("example",
            comment="created by terraform",
            life_cycle=7,
            project_name="tf_datahub_project",
            record_schema={
                "bigint_field": "BIGINT",
                "boolean_field": "BOOLEAN",
                "double_field": "DOUBLE",
                "string_field": "STRING",
                "timestamp_field": "TIMESTAMP",
            },
            record_type="TUPLE",
            shard_count=3)
        ```

        ## Import

        Datahub topic can be imported using the ID, e.g.

        ```sh
         $ pulumi import alicloud:datahub/topic:Topic example tf_datahub_project:tf_datahub_topic
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment of the datahub topic. It cannot be longer than 255 characters.
               
               **Notes:** Currently `life_cycle` can not be modified and it will be supported in the next future.
        :param pulumi.Input[int] life_cycle: How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
        :param pulumi.Input[str] name: The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        :param pulumi.Input[str] project_name: The name of the datahub project that this topic belongs to. It is case-insensitive.
        :param pulumi.Input[Mapping[str, Any]] record_schema: Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
               - BIGINT
               - STRING
               - BOOLEAN
               - DOUBLE
               - TIMESTAMP
        :param pulumi.Input[str] record_type: The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.
        :param pulumi.Input[int] shard_count: The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TopicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The topic is the basic unit of Datahub data source and is used to define one kind of data or stream. It contains a set of subscriptions. You can manage the datahub source of an application by using topics. [Refer to details](https://help.aliyun.com/document_detail/47440.html).

        ## Example Usage

        Basic Usage

        - BLob Topic

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.datahub.Topic("example",
            comment="created by terraform",
            life_cycle=7,
            project_name="tf_datahub_project",
            record_type="BLOB",
            shard_count=3)
        ```
        - Tuple Topic

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.datahub.Topic("example",
            comment="created by terraform",
            life_cycle=7,
            project_name="tf_datahub_project",
            record_schema={
                "bigint_field": "BIGINT",
                "boolean_field": "BOOLEAN",
                "double_field": "DOUBLE",
                "string_field": "STRING",
                "timestamp_field": "TIMESTAMP",
            },
            record_type="TUPLE",
            shard_count=3)
        ```

        ## Import

        Datahub topic can be imported using the ID, e.g.

        ```sh
         $ pulumi import alicloud:datahub/topic:Topic example tf_datahub_project:tf_datahub_topic
        ```

        :param str resource_name: The name of the resource.
        :param TopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 life_cycle: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 record_schema: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 record_type: Optional[pulumi.Input[str]] = None,
                 shard_count: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TopicArgs.__new__(TopicArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["life_cycle"] = life_cycle
            __props__.__dict__["name"] = name
            if project_name is None and not opts.urn:
                raise TypeError("Missing required property 'project_name'")
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["record_schema"] = record_schema
            __props__.__dict__["record_type"] = record_type
            __props__.__dict__["shard_count"] = shard_count
            __props__.__dict__["create_time"] = None
            __props__.__dict__["last_modify_time"] = None
        super(Topic, __self__).__init__(
            'alicloud:datahub/topic:Topic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            last_modify_time: Optional[pulumi.Input[str]] = None,
            life_cycle: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            record_schema: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            record_type: Optional[pulumi.Input[str]] = None,
            shard_count: Optional[pulumi.Input[int]] = None) -> 'Topic':
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment of the datahub topic. It cannot be longer than 255 characters.
               
               **Notes:** Currently `life_cycle` can not be modified and it will be supported in the next future.
        :param pulumi.Input[str] create_time: Create time of the datahub topic. It is a human-readable string rather than 64-bits UTC.
        :param pulumi.Input[str] last_modify_time: Last modify time of the datahub topic. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
        :param pulumi.Input[int] life_cycle: How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
        :param pulumi.Input[str] name: The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        :param pulumi.Input[str] project_name: The name of the datahub project that this topic belongs to. It is case-insensitive.
        :param pulumi.Input[Mapping[str, Any]] record_schema: Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
               - BIGINT
               - STRING
               - BOOLEAN
               - DOUBLE
               - TIMESTAMP
        :param pulumi.Input[str] record_type: The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.
        :param pulumi.Input[int] shard_count: The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TopicState.__new__(_TopicState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["last_modify_time"] = last_modify_time
        __props__.__dict__["life_cycle"] = life_cycle
        __props__.__dict__["name"] = name
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["record_schema"] = record_schema
        __props__.__dict__["record_type"] = record_type
        __props__.__dict__["shard_count"] = shard_count
        return Topic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment of the datahub topic. It cannot be longer than 255 characters.

        **Notes:** Currently `life_cycle` can not be modified and it will be supported in the next future.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the datahub topic. It is a human-readable string rather than 64-bits UTC.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="lastModifyTime")
    def last_modify_time(self) -> pulumi.Output[str]:
        """
        Last modify time of the datahub topic. It is the same as *create_time* at the beginning. It is also a human-readable string rather than 64-bits UTC.
        """
        return pulumi.get(self, "last_modify_time")

    @property
    @pulumi.getter(name="lifeCycle")
    def life_cycle(self) -> pulumi.Output[Optional[int]]:
        """
        How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.
        """
        return pulumi.get(self, "life_cycle")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The name of the datahub project that this topic belongs to. It is case-insensitive.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="recordSchema")
    def record_schema(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are:
        - BIGINT
        - STRING
        - BOOLEAN
        - DOUBLE
        - TIMESTAMP
        """
        return pulumi.get(self, "record_schema")

    @property
    @pulumi.getter(name="recordType")
    def record_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.
        """
        return pulumi.get(self, "record_type")

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> pulumi.Output[Optional[int]]:
        """
        The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.
        """
        return pulumi.get(self, "shard_count")

