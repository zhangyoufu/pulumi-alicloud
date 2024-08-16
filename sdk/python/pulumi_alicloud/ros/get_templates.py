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

__all__ = [
    'GetTemplatesResult',
    'AwaitableGetTemplatesResult',
    'get_templates',
    'get_templates_output',
]

@pulumi.output_type
class GetTemplatesResult:
    """
    A collection of values returned by getTemplates.
    """
    def __init__(__self__, enable_details=None, id=None, ids=None, name_regex=None, names=None, output_file=None, share_type=None, tags=None, template_name=None, templates=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if share_type and not isinstance(share_type, str):
            raise TypeError("Expected argument 'share_type' to be a str")
        pulumi.set(__self__, "share_type", share_type)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if template_name and not isinstance(template_name, str):
            raise TypeError("Expected argument 'template_name' to be a str")
        pulumi.set(__self__, "template_name", template_name)
        if templates and not isinstance(templates, list):
            raise TypeError("Expected argument 'templates' to be a list")
        pulumi.set(__self__, "templates", templates)

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="shareType")
    def share_type(self) -> Optional[str]:
        return pulumi.get(self, "share_type")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> Optional[str]:
        return pulumi.get(self, "template_name")

    @property
    @pulumi.getter
    def templates(self) -> Sequence['outputs.GetTemplatesTemplateResult']:
        return pulumi.get(self, "templates")


class AwaitableGetTemplatesResult(GetTemplatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTemplatesResult(
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            share_type=self.share_type,
            tags=self.tags,
            template_name=self.template_name,
            templates=self.templates)


def get_templates(enable_details: Optional[bool] = None,
                  ids: Optional[Sequence[str]] = None,
                  name_regex: Optional[str] = None,
                  output_file: Optional[str] = None,
                  share_type: Optional[str] = None,
                  tags: Optional[Mapping[str, str]] = None,
                  template_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTemplatesResult:
    """
    This data source provides the Ros Templates of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.108.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.ros.get_templates(ids=["example_value"],
        name_regex="the_resource_name")
    pulumi.export("firstRosTemplateId", example.templates[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Template IDs.
    :param str name_regex: A regex string to filter results by Template name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str share_type: Share Type. Valid Values: `Private`, `Shared`
    :param Mapping[str, str] tags: Query the resource bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
    :param str template_name: The name of the template.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['shareType'] = share_type
    __args__['tags'] = tags
    __args__['templateName'] = template_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ros/getTemplates:getTemplates', __args__, opts=opts, typ=GetTemplatesResult).value

    return AwaitableGetTemplatesResult(
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        share_type=pulumi.get(__ret__, 'share_type'),
        tags=pulumi.get(__ret__, 'tags'),
        template_name=pulumi.get(__ret__, 'template_name'),
        templates=pulumi.get(__ret__, 'templates'))


@_utilities.lift_output_func(get_templates)
def get_templates_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         share_type: Optional[pulumi.Input[Optional[str]]] = None,
                         tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                         template_name: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTemplatesResult]:
    """
    This data source provides the Ros Templates of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.108.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.ros.get_templates(ids=["example_value"],
        name_regex="the_resource_name")
    pulumi.export("firstRosTemplateId", example.templates[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Template IDs.
    :param str name_regex: A regex string to filter results by Template name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str share_type: Share Type. Valid Values: `Private`, `Shared`
    :param Mapping[str, str] tags: Query the resource bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
    :param str template_name: The name of the template.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
    """
    ...
