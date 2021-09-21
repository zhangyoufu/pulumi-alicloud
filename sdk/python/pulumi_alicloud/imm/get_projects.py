# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetProjectsResult',
    'AwaitableGetProjectsResult',
    'get_projects',
]

@pulumi.output_type
class GetProjectsResult:
    """
    A collection of values returned by getProjects.
    """
    def __init__(__self__, id=None, ids=None, output_file=None, projects=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if projects and not isinstance(projects, list):
            raise TypeError("Expected argument 'projects' to be a list")
        pulumi.set(__self__, "projects", projects)

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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def projects(self) -> Sequence['outputs.GetProjectsProjectResult']:
        return pulumi.get(self, "projects")


class AwaitableGetProjectsResult(GetProjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectsResult(
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            projects=self.projects)


def get_projects(ids: Optional[Sequence[str]] = None,
                 output_file: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectsResult:
    """
    This data source provides the Intelligent Media Management Projects of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.134.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.imm.get_projects(ids=["example_id"])
    pulumi.export("immProjectId1", ids.projects[0].id)
    ```


    :param Sequence[str] ids: A list of Project IDs.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:imm/getProjects:getProjects', __args__, opts=opts, typ=GetProjectsResult).value

    return AwaitableGetProjectsResult(
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        projects=__ret__.projects)
