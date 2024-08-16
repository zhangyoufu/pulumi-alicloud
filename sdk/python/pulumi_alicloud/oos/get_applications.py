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
    'GetApplicationsResult',
    'AwaitableGetApplicationsResult',
    'get_applications',
    'get_applications_output',
]

@pulumi.output_type
class GetApplicationsResult:
    """
    A collection of values returned by getApplications.
    """
    def __init__(__self__, applications=None, id=None, ids=None, name_regex=None, names=None, output_file=None, tags=None):
        if applications and not isinstance(applications, list):
            raise TypeError("Expected argument 'applications' to be a list")
        pulumi.set(__self__, "applications", applications)
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
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def applications(self) -> Sequence['outputs.GetApplicationsApplicationResult']:
        return pulumi.get(self, "applications")

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
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "tags")


class AwaitableGetApplicationsResult(GetApplicationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationsResult(
            applications=self.applications,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            tags=self.tags)


def get_applications(ids: Optional[Sequence[str]] = None,
                     name_regex: Optional[str] = None,
                     output_file: Optional[str] = None,
                     tags: Optional[Mapping[str, str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationsResult:
    """
    This data source provides the Oos Applications of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.145.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.oos.get_applications()
    pulumi.export("oosApplicationId1", ids.applications[0].id)
    name_regex = alicloud.oos.get_applications(name_regex="^my-Application")
    pulumi.export("oosApplicationId2", name_regex.applications[0].id)
    ```


    :param Sequence[str] ids: A list of Application IDs. Its element value is same as Application Name.
    :param str name_regex: A regex string to filter results by Application name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param Mapping[str, str] tags: A mapping of tags to assign to the resource.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:oos/getApplications:getApplications', __args__, opts=opts, typ=GetApplicationsResult).value

    return AwaitableGetApplicationsResult(
        applications=pulumi.get(__ret__, 'applications'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_applications)
def get_applications_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationsResult]:
    """
    This data source provides the Oos Applications of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.145.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.oos.get_applications()
    pulumi.export("oosApplicationId1", ids.applications[0].id)
    name_regex = alicloud.oos.get_applications(name_regex="^my-Application")
    pulumi.export("oosApplicationId2", name_regex.applications[0].id)
    ```


    :param Sequence[str] ids: A list of Application IDs. Its element value is same as Application Name.
    :param str name_regex: A regex string to filter results by Application name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param Mapping[str, str] tags: A mapping of tags to assign to the resource.
    """
    ...
