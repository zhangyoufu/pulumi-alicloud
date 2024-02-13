# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TemplateArgs', 'Template']

@pulumi.input_type
class TemplateArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 template_name: pulumi.Input[str],
                 auto_delete_executions: Optional[pulumi.Input[bool]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 version_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Template resource.
        :param pulumi.Input[str] content: The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
        :param pulumi.Input[str] template_name: The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
        :param pulumi.Input[bool] auto_delete_executions: When deleting a template, whether to delete its related executions. Default to `false`.
        :param pulumi.Input[str] resource_group_id: The ID of resource group which the template belongs.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] version_name: The name of template version.
        """
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "template_name", template_name)
        if auto_delete_executions is not None:
            pulumi.set(__self__, "auto_delete_executions", auto_delete_executions)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if version_name is not None:
            pulumi.set(__self__, "version_name", version_name)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Input[str]:
        """
        The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
        """
        return pulumi.get(self, "template_name")

    @template_name.setter
    def template_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_name", value)

    @property
    @pulumi.getter(name="autoDeleteExecutions")
    def auto_delete_executions(self) -> Optional[pulumi.Input[bool]]:
        """
        When deleting a template, whether to delete its related executions. Default to `false`.
        """
        return pulumi.get(self, "auto_delete_executions")

    @auto_delete_executions.setter
    def auto_delete_executions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_delete_executions", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of resource group which the template belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="versionName")
    def version_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of template version.
        """
        return pulumi.get(self, "version_name")

    @version_name.setter
    def version_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_name", value)


@pulumi.input_type
class _TemplateState:
    def __init__(__self__, *,
                 auto_delete_executions: Optional[pulumi.Input[bool]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 created_by: Optional[pulumi.Input[str]] = None,
                 created_date: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 has_trigger: Optional[pulumi.Input[bool]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 share_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 template_format: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 template_type: Optional[pulumi.Input[str]] = None,
                 template_version: Optional[pulumi.Input[str]] = None,
                 updated_by: Optional[pulumi.Input[str]] = None,
                 updated_date: Optional[pulumi.Input[str]] = None,
                 version_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Template resources.
        :param pulumi.Input[bool] auto_delete_executions: When deleting a template, whether to delete its related executions. Default to `false`.
        :param pulumi.Input[str] content: The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
        :param pulumi.Input[str] created_by: The creator of the template.
        :param pulumi.Input[str] created_date: The time when the template is created.
        :param pulumi.Input[str] description: The description of the template.
        :param pulumi.Input[bool] has_trigger: Is it triggered successfully.
        :param pulumi.Input[str] resource_group_id: The ID of resource group which the template belongs.
        :param pulumi.Input[str] share_type: The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] template_format: The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
        :param pulumi.Input[str] template_id: The id of OOS Template.
        :param pulumi.Input[str] template_name: The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
        :param pulumi.Input[str] template_type: The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
        :param pulumi.Input[str] template_version: The version of OOS Template.
        :param pulumi.Input[str] updated_by: The user who updated the template.
        :param pulumi.Input[str] updated_date: The time when the template was updated.
        :param pulumi.Input[str] version_name: The name of template version.
        """
        if auto_delete_executions is not None:
            pulumi.set(__self__, "auto_delete_executions", auto_delete_executions)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if created_date is not None:
            pulumi.set(__self__, "created_date", created_date)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if has_trigger is not None:
            pulumi.set(__self__, "has_trigger", has_trigger)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if share_type is not None:
            pulumi.set(__self__, "share_type", share_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if template_format is not None:
            pulumi.set(__self__, "template_format", template_format)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)
        if template_name is not None:
            pulumi.set(__self__, "template_name", template_name)
        if template_type is not None:
            pulumi.set(__self__, "template_type", template_type)
        if template_version is not None:
            pulumi.set(__self__, "template_version", template_version)
        if updated_by is not None:
            pulumi.set(__self__, "updated_by", updated_by)
        if updated_date is not None:
            pulumi.set(__self__, "updated_date", updated_date)
        if version_name is not None:
            pulumi.set(__self__, "version_name", version_name)

    @property
    @pulumi.getter(name="autoDeleteExecutions")
    def auto_delete_executions(self) -> Optional[pulumi.Input[bool]]:
        """
        When deleting a template, whether to delete its related executions. Default to `false`.
        """
        return pulumi.get(self, "auto_delete_executions")

    @auto_delete_executions.setter
    def auto_delete_executions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_delete_executions", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[pulumi.Input[str]]:
        """
        The creator of the template.
        """
        return pulumi.get(self, "created_by")

    @created_by.setter
    def created_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_by", value)

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the template is created.
        """
        return pulumi.get(self, "created_date")

    @created_date.setter
    def created_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_date", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the template.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="hasTrigger")
    def has_trigger(self) -> Optional[pulumi.Input[bool]]:
        """
        Is it triggered successfully.
        """
        return pulumi.get(self, "has_trigger")

    @has_trigger.setter
    def has_trigger(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "has_trigger", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of resource group which the template belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="shareType")
    def share_type(self) -> Optional[pulumi.Input[str]]:
        """
        The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
        """
        return pulumi.get(self, "share_type")

    @share_type.setter
    def share_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "share_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="templateFormat")
    def template_format(self) -> Optional[pulumi.Input[str]]:
        """
        The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
        """
        return pulumi.get(self, "template_format")

    @template_format.setter
    def template_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_format", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of OOS Template.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
        """
        return pulumi.get(self, "template_name")

    @template_name.setter
    def template_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_name", value)

    @property
    @pulumi.getter(name="templateType")
    def template_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
        """
        return pulumi.get(self, "template_type")

    @template_type.setter
    def template_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_type", value)

    @property
    @pulumi.getter(name="templateVersion")
    def template_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of OOS Template.
        """
        return pulumi.get(self, "template_version")

    @template_version.setter
    def template_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_version", value)

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> Optional[pulumi.Input[str]]:
        """
        The user who updated the template.
        """
        return pulumi.get(self, "updated_by")

    @updated_by.setter
    def updated_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_by", value)

    @property
    @pulumi.getter(name="updatedDate")
    def updated_date(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the template was updated.
        """
        return pulumi.get(self, "updated_date")

    @updated_date.setter
    def updated_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_date", value)

    @property
    @pulumi.getter(name="versionName")
    def version_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of template version.
        """
        return pulumi.get(self, "version_name")

    @version_name.setter
    def version_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_name", value)


class Template(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_delete_executions: Optional[pulumi.Input[bool]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 version_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a OOS Template resource. For information about Alicloud OOS Template and how to use it, see [What is Resource Alicloud OOS Template](https://www.alibabacloud.com/help/doc-detail/120761.htm).

        > **NOTE:** Available in 1.92.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.oos.Template("example",
            content=\"\"\"  {
            "FormatVersion": "OOS-2019-06-01",
            "Description": "Update Describe instances of given status",
            "Parameters":{
              "Status":{
                "Type": "String",
                "Description": "(Required) The status of the Ecs instance."
              }
            },
            "Tasks": [
              {
                "Properties" :{
                  "Parameters":{
                    "Status": "{{ Status }}"
                  },
                  "API": "DescribeInstances",
                  "Service": "Ecs"
                },
                "Name": "foo",
                "Action": "ACS::ExecuteApi"
              }]
          }
          
        \"\"\",
            tags={
                "Created": "TF",
                "For": "acceptance Test",
            },
            template_name="test-name",
            version_name="test")
        ```

        ## Import

        OOS Template can be imported using the id or template_name, e.g.

        ```sh
        $ pulumi import alicloud:oos/template:Template example template_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_delete_executions: When deleting a template, whether to delete its related executions. Default to `false`.
        :param pulumi.Input[str] content: The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
        :param pulumi.Input[str] resource_group_id: The ID of resource group which the template belongs.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] template_name: The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
        :param pulumi.Input[str] version_name: The name of template version.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a OOS Template resource. For information about Alicloud OOS Template and how to use it, see [What is Resource Alicloud OOS Template](https://www.alibabacloud.com/help/doc-detail/120761.htm).

        > **NOTE:** Available in 1.92.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.oos.Template("example",
            content=\"\"\"  {
            "FormatVersion": "OOS-2019-06-01",
            "Description": "Update Describe instances of given status",
            "Parameters":{
              "Status":{
                "Type": "String",
                "Description": "(Required) The status of the Ecs instance."
              }
            },
            "Tasks": [
              {
                "Properties" :{
                  "Parameters":{
                    "Status": "{{ Status }}"
                  },
                  "API": "DescribeInstances",
                  "Service": "Ecs"
                },
                "Name": "foo",
                "Action": "ACS::ExecuteApi"
              }]
          }
          
        \"\"\",
            tags={
                "Created": "TF",
                "For": "acceptance Test",
            },
            template_name="test-name",
            version_name="test")
        ```

        ## Import

        OOS Template can be imported using the id or template_name, e.g.

        ```sh
        $ pulumi import alicloud:oos/template:Template example template_name
        ```

        :param str resource_name: The name of the resource.
        :param TemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_delete_executions: Optional[pulumi.Input[bool]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 version_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TemplateArgs.__new__(TemplateArgs)

            __props__.__dict__["auto_delete_executions"] = auto_delete_executions
            if content is None and not opts.urn:
                raise TypeError("Missing required property 'content'")
            __props__.__dict__["content"] = content
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["tags"] = tags
            if template_name is None and not opts.urn:
                raise TypeError("Missing required property 'template_name'")
            __props__.__dict__["template_name"] = template_name
            __props__.__dict__["version_name"] = version_name
            __props__.__dict__["created_by"] = None
            __props__.__dict__["created_date"] = None
            __props__.__dict__["description"] = None
            __props__.__dict__["has_trigger"] = None
            __props__.__dict__["share_type"] = None
            __props__.__dict__["template_format"] = None
            __props__.__dict__["template_id"] = None
            __props__.__dict__["template_type"] = None
            __props__.__dict__["template_version"] = None
            __props__.__dict__["updated_by"] = None
            __props__.__dict__["updated_date"] = None
        super(Template, __self__).__init__(
            'alicloud:oos/template:Template',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_delete_executions: Optional[pulumi.Input[bool]] = None,
            content: Optional[pulumi.Input[str]] = None,
            created_by: Optional[pulumi.Input[str]] = None,
            created_date: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            has_trigger: Optional[pulumi.Input[bool]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            share_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            template_format: Optional[pulumi.Input[str]] = None,
            template_id: Optional[pulumi.Input[str]] = None,
            template_name: Optional[pulumi.Input[str]] = None,
            template_type: Optional[pulumi.Input[str]] = None,
            template_version: Optional[pulumi.Input[str]] = None,
            updated_by: Optional[pulumi.Input[str]] = None,
            updated_date: Optional[pulumi.Input[str]] = None,
            version_name: Optional[pulumi.Input[str]] = None) -> 'Template':
        """
        Get an existing Template resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_delete_executions: When deleting a template, whether to delete its related executions. Default to `false`.
        :param pulumi.Input[str] content: The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
        :param pulumi.Input[str] created_by: The creator of the template.
        :param pulumi.Input[str] created_date: The time when the template is created.
        :param pulumi.Input[str] description: The description of the template.
        :param pulumi.Input[bool] has_trigger: Is it triggered successfully.
        :param pulumi.Input[str] resource_group_id: The ID of resource group which the template belongs.
        :param pulumi.Input[str] share_type: The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] template_format: The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
        :param pulumi.Input[str] template_id: The id of OOS Template.
        :param pulumi.Input[str] template_name: The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
        :param pulumi.Input[str] template_type: The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
        :param pulumi.Input[str] template_version: The version of OOS Template.
        :param pulumi.Input[str] updated_by: The user who updated the template.
        :param pulumi.Input[str] updated_date: The time when the template was updated.
        :param pulumi.Input[str] version_name: The name of template version.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TemplateState.__new__(_TemplateState)

        __props__.__dict__["auto_delete_executions"] = auto_delete_executions
        __props__.__dict__["content"] = content
        __props__.__dict__["created_by"] = created_by
        __props__.__dict__["created_date"] = created_date
        __props__.__dict__["description"] = description
        __props__.__dict__["has_trigger"] = has_trigger
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["share_type"] = share_type
        __props__.__dict__["tags"] = tags
        __props__.__dict__["template_format"] = template_format
        __props__.__dict__["template_id"] = template_id
        __props__.__dict__["template_name"] = template_name
        __props__.__dict__["template_type"] = template_type
        __props__.__dict__["template_version"] = template_version
        __props__.__dict__["updated_by"] = updated_by
        __props__.__dict__["updated_date"] = updated_date
        __props__.__dict__["version_name"] = version_name
        return Template(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoDeleteExecutions")
    def auto_delete_executions(self) -> pulumi.Output[Optional[bool]]:
        """
        When deleting a template, whether to delete its related executions. Default to `false`.
        """
        return pulumi.get(self, "auto_delete_executions")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> pulumi.Output[str]:
        """
        The creator of the template.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> pulumi.Output[str]:
        """
        The time when the template is created.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the template.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="hasTrigger")
    def has_trigger(self) -> pulumi.Output[bool]:
        """
        Is it triggered successfully.
        """
        return pulumi.get(self, "has_trigger")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The ID of resource group which the template belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="shareType")
    def share_type(self) -> pulumi.Output[str]:
        """
        The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
        """
        return pulumi.get(self, "share_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateFormat")
    def template_format(self) -> pulumi.Output[str]:
        """
        The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
        """
        return pulumi.get(self, "template_format")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[str]:
        """
        The id of OOS Template.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Output[str]:
        """
        The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
        """
        return pulumi.get(self, "template_name")

    @property
    @pulumi.getter(name="templateType")
    def template_type(self) -> pulumi.Output[str]:
        """
        The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
        """
        return pulumi.get(self, "template_type")

    @property
    @pulumi.getter(name="templateVersion")
    def template_version(self) -> pulumi.Output[str]:
        """
        The version of OOS Template.
        """
        return pulumi.get(self, "template_version")

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> pulumi.Output[str]:
        """
        The user who updated the template.
        """
        return pulumi.get(self, "updated_by")

    @property
    @pulumi.getter(name="updatedDate")
    def updated_date(self) -> pulumi.Output[str]:
        """
        The time when the template was updated.
        """
        return pulumi.get(self, "updated_date")

    @property
    @pulumi.getter(name="versionName")
    def version_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of template version.
        """
        return pulumi.get(self, "version_name")

