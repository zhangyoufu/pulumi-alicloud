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
    def __init__(__self__, category=None, created_by=None, created_date=None, created_date_after=None, has_trigger=None, id=None, ids=None, name_regex=None, names=None, output_file=None, share_type=None, sort_field=None, sort_order=None, tags=None, template_format=None, template_type=None, templates=None):
        if category and not isinstance(category, str):
            raise TypeError("Expected argument 'category' to be a str")
        pulumi.set(__self__, "category", category)
        if created_by and not isinstance(created_by, str):
            raise TypeError("Expected argument 'created_by' to be a str")
        pulumi.set(__self__, "created_by", created_by)
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if created_date_after and not isinstance(created_date_after, str):
            raise TypeError("Expected argument 'created_date_after' to be a str")
        pulumi.set(__self__, "created_date_after", created_date_after)
        if has_trigger and not isinstance(has_trigger, bool):
            raise TypeError("Expected argument 'has_trigger' to be a bool")
        pulumi.set(__self__, "has_trigger", has_trigger)
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
        if sort_field and not isinstance(sort_field, str):
            raise TypeError("Expected argument 'sort_field' to be a str")
        pulumi.set(__self__, "sort_field", sort_field)
        if sort_order and not isinstance(sort_order, str):
            raise TypeError("Expected argument 'sort_order' to be a str")
        pulumi.set(__self__, "sort_order", sort_order)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if template_format and not isinstance(template_format, str):
            raise TypeError("Expected argument 'template_format' to be a str")
        pulumi.set(__self__, "template_format", template_format)
        if template_type and not isinstance(template_type, str):
            raise TypeError("Expected argument 'template_type' to be a str")
        pulumi.set(__self__, "template_type", template_type)
        if templates and not isinstance(templates, list):
            raise TypeError("Expected argument 'templates' to be a list")
        pulumi.set(__self__, "templates", templates)

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[str]:
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> Optional[str]:
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter(name="createdDateAfter")
    def created_date_after(self) -> Optional[str]:
        return pulumi.get(self, "created_date_after")

    @property
    @pulumi.getter(name="hasTrigger")
    def has_trigger(self) -> Optional[bool]:
        return pulumi.get(self, "has_trigger")

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
        """
        A list of OOS Template ids. Each element in the list is same as template_name.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        (Available in v1.114.0+) A list of OOS Template names.
        """
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
    @pulumi.getter(name="sortField")
    def sort_field(self) -> Optional[str]:
        return pulumi.get(self, "sort_field")

    @property
    @pulumi.getter(name="sortOrder")
    def sort_order(self) -> Optional[str]:
        return pulumi.get(self, "sort_order")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateFormat")
    def template_format(self) -> Optional[str]:
        return pulumi.get(self, "template_format")

    @property
    @pulumi.getter(name="templateType")
    def template_type(self) -> Optional[str]:
        return pulumi.get(self, "template_type")

    @property
    @pulumi.getter
    def templates(self) -> Sequence['outputs.GetTemplatesTemplateResult']:
        """
        A list of OOS Templates. Each element contains the following attributes:
        """
        return pulumi.get(self, "templates")


class AwaitableGetTemplatesResult(GetTemplatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTemplatesResult(
            category=self.category,
            created_by=self.created_by,
            created_date=self.created_date,
            created_date_after=self.created_date_after,
            has_trigger=self.has_trigger,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            share_type=self.share_type,
            sort_field=self.sort_field,
            sort_order=self.sort_order,
            tags=self.tags,
            template_format=self.template_format,
            template_type=self.template_type,
            templates=self.templates)


def get_templates(category: Optional[str] = None,
                  created_by: Optional[str] = None,
                  created_date: Optional[str] = None,
                  created_date_after: Optional[str] = None,
                  has_trigger: Optional[bool] = None,
                  ids: Optional[Sequence[str]] = None,
                  name_regex: Optional[str] = None,
                  output_file: Optional[str] = None,
                  share_type: Optional[str] = None,
                  sort_field: Optional[str] = None,
                  sort_order: Optional[str] = None,
                  tags: Optional[Mapping[str, Any]] = None,
                  template_format: Optional[str] = None,
                  template_type: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTemplatesResult:
    """
    This data source provides a list of OOS Templates in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in v1.92.0+.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.oos.get_templates(has_trigger=False,
        name_regex="test",
        share_type="Private",
        tags={
            "Created": "TF",
            "For": "template Test",
        })
    pulumi.export("firstTemplateName", example.templates[0].template_name)
    ```
    <!--End PulumiCodeChooser -->


    :param str category: The category of template.
    :param str created_by: The creator of the template.
    :param str created_date: The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.
    :param str created_date_after: Create a template whose time is greater than or equal to the specified time. The format is: YYYY-MM-DDThh:mm:ssZ.
    :param bool has_trigger: Is it triggered successfully.
    :param Sequence[str] ids: A list of OOS Template ids. Each element in the list is same as template_name.
    :param str name_regex: A regex string to filter the results by the template_name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str share_type: The sharing type of the template. Valid values: `Private`, `Public`.
    :param str sort_field: Sort field. Valid values: `TotalExecutionCount`, `Popularity`, `TemplateName` and `CreatedDate`. Default to `TotalExecutionCount`.
    :param str sort_order: Sort order. Valid values: `Ascending`, `Descending`. Default to `Descending`
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    :param str template_format: The format of the template. Valid values: `JSON`, `YAML`.
    :param str template_type: The type of OOS Template.
    """
    __args__ = dict()
    __args__['category'] = category
    __args__['createdBy'] = created_by
    __args__['createdDate'] = created_date
    __args__['createdDateAfter'] = created_date_after
    __args__['hasTrigger'] = has_trigger
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['shareType'] = share_type
    __args__['sortField'] = sort_field
    __args__['sortOrder'] = sort_order
    __args__['tags'] = tags
    __args__['templateFormat'] = template_format
    __args__['templateType'] = template_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:oos/getTemplates:getTemplates', __args__, opts=opts, typ=GetTemplatesResult).value

    return AwaitableGetTemplatesResult(
        category=pulumi.get(__ret__, 'category'),
        created_by=pulumi.get(__ret__, 'created_by'),
        created_date=pulumi.get(__ret__, 'created_date'),
        created_date_after=pulumi.get(__ret__, 'created_date_after'),
        has_trigger=pulumi.get(__ret__, 'has_trigger'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        share_type=pulumi.get(__ret__, 'share_type'),
        sort_field=pulumi.get(__ret__, 'sort_field'),
        sort_order=pulumi.get(__ret__, 'sort_order'),
        tags=pulumi.get(__ret__, 'tags'),
        template_format=pulumi.get(__ret__, 'template_format'),
        template_type=pulumi.get(__ret__, 'template_type'),
        templates=pulumi.get(__ret__, 'templates'))


@_utilities.lift_output_func(get_templates)
def get_templates_output(category: Optional[pulumi.Input[Optional[str]]] = None,
                         created_by: Optional[pulumi.Input[Optional[str]]] = None,
                         created_date: Optional[pulumi.Input[Optional[str]]] = None,
                         created_date_after: Optional[pulumi.Input[Optional[str]]] = None,
                         has_trigger: Optional[pulumi.Input[Optional[bool]]] = None,
                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         share_type: Optional[pulumi.Input[Optional[str]]] = None,
                         sort_field: Optional[pulumi.Input[Optional[str]]] = None,
                         sort_order: Optional[pulumi.Input[Optional[str]]] = None,
                         tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                         template_format: Optional[pulumi.Input[Optional[str]]] = None,
                         template_type: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTemplatesResult]:
    """
    This data source provides a list of OOS Templates in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in v1.92.0+.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.oos.get_templates(has_trigger=False,
        name_regex="test",
        share_type="Private",
        tags={
            "Created": "TF",
            "For": "template Test",
        })
    pulumi.export("firstTemplateName", example.templates[0].template_name)
    ```
    <!--End PulumiCodeChooser -->


    :param str category: The category of template.
    :param str created_by: The creator of the template.
    :param str created_date: The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.
    :param str created_date_after: Create a template whose time is greater than or equal to the specified time. The format is: YYYY-MM-DDThh:mm:ssZ.
    :param bool has_trigger: Is it triggered successfully.
    :param Sequence[str] ids: A list of OOS Template ids. Each element in the list is same as template_name.
    :param str name_regex: A regex string to filter the results by the template_name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str share_type: The sharing type of the template. Valid values: `Private`, `Public`.
    :param str sort_field: Sort field. Valid values: `TotalExecutionCount`, `Popularity`, `TemplateName` and `CreatedDate`. Default to `TotalExecutionCount`.
    :param str sort_order: Sort order. Valid values: `Ascending`, `Descending`. Default to `Descending`
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    :param str template_format: The format of the template. Valid values: `JSON`, `YAML`.
    :param str template_type: The type of OOS Template.
    """
    ...
