# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Execution']


class Execution(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 loop_mode: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 parent_execution_id: Optional[pulumi.Input[str]] = None,
                 safety_check: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 template_version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a OOS Execution resource. For information about Alicloud OOS Execution and how to use it, see [What is Resource Alicloud OOS Execution](https://www.alibabacloud.com/help/doc-detail/120771.htm).

        > **NOTE:** Available in 1.93.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.oos.Template("default",
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
            template_name="test-name",
            version_name="test",
            tags={
                "Created": "TF",
                "For": "acceptance Test",
            })
        example = alicloud.oos.Execution("example",
            template_name=default.template_name,
            description="From TF Test",
            parameters="				{\"Status\":\"Running\"}\n")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of OOS Execution.
        :param pulumi.Input[str] loop_mode: The loop mode of OOS Execution.
        :param pulumi.Input[str] mode: The mode of OOS Execution. Valid: `Automatic`, `Debug`. Default to `Automatic`.
        :param pulumi.Input[str] parameters: The parameters required by the template. Default to `{}`.
        :param pulumi.Input[str] parent_execution_id: The id of parent execution.
        :param pulumi.Input[str] safety_check: The mode of safety check.
        :param pulumi.Input[str] template_name: The name of execution template.
        :param pulumi.Input[str] template_version: The version of execution template.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['loop_mode'] = loop_mode
            __props__['mode'] = mode
            __props__['parameters'] = parameters
            __props__['parent_execution_id'] = parent_execution_id
            __props__['safety_check'] = safety_check
            if template_name is None:
                raise TypeError("Missing required property 'template_name'")
            __props__['template_name'] = template_name
            __props__['template_version'] = template_version
            __props__['counters'] = None
            __props__['create_date'] = None
            __props__['end_date'] = None
            __props__['executed_by'] = None
            __props__['is_parent'] = None
            __props__['outputs'] = None
            __props__['ram_role'] = None
            __props__['start_date'] = None
            __props__['status'] = None
            __props__['status_message'] = None
            __props__['template_id'] = None
            __props__['update_date'] = None
        super(Execution, __self__).__init__(
            'alicloud:oos/execution:Execution',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            counters: Optional[pulumi.Input[str]] = None,
            create_date: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            end_date: Optional[pulumi.Input[str]] = None,
            executed_by: Optional[pulumi.Input[str]] = None,
            is_parent: Optional[pulumi.Input[bool]] = None,
            loop_mode: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            outputs: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[str]] = None,
            parent_execution_id: Optional[pulumi.Input[str]] = None,
            ram_role: Optional[pulumi.Input[str]] = None,
            safety_check: Optional[pulumi.Input[str]] = None,
            start_date: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            status_message: Optional[pulumi.Input[str]] = None,
            template_id: Optional[pulumi.Input[str]] = None,
            template_name: Optional[pulumi.Input[str]] = None,
            template_version: Optional[pulumi.Input[str]] = None,
            update_date: Optional[pulumi.Input[str]] = None) -> 'Execution':
        """
        Get an existing Execution resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] counters: The counters of OOS Execution.
        :param pulumi.Input[str] create_date: The time when the execution was created.
        :param pulumi.Input[str] description: The description of OOS Execution.
        :param pulumi.Input[str] end_date: The time when the execution was ended.
        :param pulumi.Input[str] executed_by: The user who execute the template.
        :param pulumi.Input[bool] is_parent: Whether to include subtasks.
        :param pulumi.Input[str] loop_mode: The loop mode of OOS Execution.
        :param pulumi.Input[str] mode: The mode of OOS Execution. Valid: `Automatic`, `Debug`. Default to `Automatic`.
        :param pulumi.Input[str] outputs: The outputs of OOS Execution.
        :param pulumi.Input[str] parameters: The parameters required by the template. Default to `{}`.
        :param pulumi.Input[str] parent_execution_id: The id of parent execution.
        :param pulumi.Input[str] ram_role: The role that executes the current template.
        :param pulumi.Input[str] safety_check: The mode of safety check.
        :param pulumi.Input[str] start_date: The time when the execution was started.
        :param pulumi.Input[str] status: The status of OOS Execution.
        :param pulumi.Input[str] status_message: The message of status.
        :param pulumi.Input[str] template_id: The id of template.
        :param pulumi.Input[str] template_name: The name of execution template.
        :param pulumi.Input[str] template_version: The version of execution template.
        :param pulumi.Input[str] update_date: The time when the execution was updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["counters"] = counters
        __props__["create_date"] = create_date
        __props__["description"] = description
        __props__["end_date"] = end_date
        __props__["executed_by"] = executed_by
        __props__["is_parent"] = is_parent
        __props__["loop_mode"] = loop_mode
        __props__["mode"] = mode
        __props__["outputs"] = outputs
        __props__["parameters"] = parameters
        __props__["parent_execution_id"] = parent_execution_id
        __props__["ram_role"] = ram_role
        __props__["safety_check"] = safety_check
        __props__["start_date"] = start_date
        __props__["status"] = status
        __props__["status_message"] = status_message
        __props__["template_id"] = template_id
        __props__["template_name"] = template_name
        __props__["template_version"] = template_version
        __props__["update_date"] = update_date
        return Execution(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def counters(self) -> str:
        """
        The counters of OOS Execution.
        """
        return pulumi.get(self, "counters")

    @property
    @pulumi.getter(name="createDate")
    def create_date(self) -> str:
        """
        The time when the execution was created.
        """
        return pulumi.get(self, "create_date")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of OOS Execution.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> str:
        """
        The time when the execution was ended.
        """
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter(name="executedBy")
    def executed_by(self) -> str:
        """
        The user who execute the template.
        """
        return pulumi.get(self, "executed_by")

    @property
    @pulumi.getter(name="isParent")
    def is_parent(self) -> bool:
        """
        Whether to include subtasks.
        """
        return pulumi.get(self, "is_parent")

    @property
    @pulumi.getter(name="loopMode")
    def loop_mode(self) -> Optional[str]:
        """
        The loop mode of OOS Execution.
        """
        return pulumi.get(self, "loop_mode")

    @property
    @pulumi.getter
    def mode(self) -> Optional[str]:
        """
        The mode of OOS Execution. Valid: `Automatic`, `Debug`. Default to `Automatic`.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def outputs(self) -> str:
        """
        The outputs of OOS Execution.
        """
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[str]:
        """
        The parameters required by the template. Default to `{}`.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="parentExecutionId")
    def parent_execution_id(self) -> Optional[str]:
        """
        The id of parent execution.
        """
        return pulumi.get(self, "parent_execution_id")

    @property
    @pulumi.getter(name="ramRole")
    def ram_role(self) -> str:
        """
        The role that executes the current template.
        """
        return pulumi.get(self, "ram_role")

    @property
    @pulumi.getter(name="safetyCheck")
    def safety_check(self) -> Optional[str]:
        """
        The mode of safety check.
        """
        return pulumi.get(self, "safety_check")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> str:
        """
        The time when the execution was started.
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of OOS Execution.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> str:
        """
        The message of status.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> str:
        """
        The id of template.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> str:
        """
        The name of execution template.
        """
        return pulumi.get(self, "template_name")

    @property
    @pulumi.getter(name="templateVersion")
    def template_version(self) -> str:
        """
        The version of execution template.
        """
        return pulumi.get(self, "template_version")

    @property
    @pulumi.getter(name="updateDate")
    def update_date(self) -> str:
        """
        The time when the execution was updated.
        """
        return pulumi.get(self, "update_date")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

