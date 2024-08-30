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
    'GetIndustrialPidLoopsResult',
    'AwaitableGetIndustrialPidLoopsResult',
    'get_industrial_pid_loops',
    'get_industrial_pid_loops_output',
]

@pulumi.output_type
class GetIndustrialPidLoopsResult:
    """
    A collection of values returned by getIndustrialPidLoops.
    """
    def __init__(__self__, enable_details=None, id=None, ids=None, loops=None, name_regex=None, names=None, output_file=None, pid_loop_name=None, pid_project_id=None, status=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if loops and not isinstance(loops, list):
            raise TypeError("Expected argument 'loops' to be a list")
        pulumi.set(__self__, "loops", loops)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if pid_loop_name and not isinstance(pid_loop_name, str):
            raise TypeError("Expected argument 'pid_loop_name' to be a str")
        pulumi.set(__self__, "pid_loop_name", pid_loop_name)
        if pid_project_id and not isinstance(pid_project_id, str):
            raise TypeError("Expected argument 'pid_project_id' to be a str")
        pulumi.set(__self__, "pid_project_id", pid_project_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

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
    @pulumi.getter
    def loops(self) -> Sequence['outputs.GetIndustrialPidLoopsLoopResult']:
        """
        A list of Brain Industrial Pid Loops. Each element contains the following attributes:
        """
        return pulumi.get(self, "loops")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of Pid Loop names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="pidLoopName")
    def pid_loop_name(self) -> Optional[str]:
        """
        The name of Pid Loop.
        """
        return pulumi.get(self, "pid_loop_name")

    @property
    @pulumi.getter(name="pidProjectId")
    def pid_project_id(self) -> str:
        return pulumi.get(self, "pid_project_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of Pid Loop.
        """
        return pulumi.get(self, "status")


class AwaitableGetIndustrialPidLoopsResult(GetIndustrialPidLoopsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIndustrialPidLoopsResult(
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            loops=self.loops,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            pid_loop_name=self.pid_loop_name,
            pid_project_id=self.pid_project_id,
            status=self.status)


def get_industrial_pid_loops(enable_details: Optional[bool] = None,
                             ids: Optional[Sequence[str]] = None,
                             name_regex: Optional[str] = None,
                             output_file: Optional[str] = None,
                             pid_loop_name: Optional[str] = None,
                             pid_project_id: Optional[str] = None,
                             status: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIndustrialPidLoopsResult:
    """
    This data source provides the Brain Industrial Pid Loops of the current Alibaba Cloud user.

    > **NOTE:** Available since v1.117.0.

    > **DEPRECATED:**  This data source has been deprecated from version `1.229.1`.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.brain.get_industrial_pid_loops(pid_project_id="856c6b8f-ca63-40a4-xxxx-xxxx",
        ids=["742a3d4e-d8b0-47c8-xxxx-xxxx"],
        name_regex="tf-testACC")
    pulumi.export("firstBrainIndustrialPidLoopId", example.loops[0].id)
    ```


    :param Sequence[str] ids: A list of Pid Loop IDs.
    :param str name_regex: A regex string to filter results by Pid Loop name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str pid_loop_name: The name of Pid Loop.
    :param str pid_project_id: The pid project id.
    :param str status: The status of Pid Loop.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['pidLoopName'] = pid_loop_name
    __args__['pidProjectId'] = pid_project_id
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:brain/getIndustrialPidLoops:getIndustrialPidLoops', __args__, opts=opts, typ=GetIndustrialPidLoopsResult).value

    return AwaitableGetIndustrialPidLoopsResult(
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        loops=pulumi.get(__ret__, 'loops'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        pid_loop_name=pulumi.get(__ret__, 'pid_loop_name'),
        pid_project_id=pulumi.get(__ret__, 'pid_project_id'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_industrial_pid_loops)
def get_industrial_pid_loops_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                                    ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    pid_loop_name: Optional[pulumi.Input[Optional[str]]] = None,
                                    pid_project_id: Optional[pulumi.Input[str]] = None,
                                    status: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIndustrialPidLoopsResult]:
    """
    This data source provides the Brain Industrial Pid Loops of the current Alibaba Cloud user.

    > **NOTE:** Available since v1.117.0.

    > **DEPRECATED:**  This data source has been deprecated from version `1.229.1`.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.brain.get_industrial_pid_loops(pid_project_id="856c6b8f-ca63-40a4-xxxx-xxxx",
        ids=["742a3d4e-d8b0-47c8-xxxx-xxxx"],
        name_regex="tf-testACC")
    pulumi.export("firstBrainIndustrialPidLoopId", example.loops[0].id)
    ```


    :param Sequence[str] ids: A list of Pid Loop IDs.
    :param str name_regex: A regex string to filter results by Pid Loop name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str pid_loop_name: The name of Pid Loop.
    :param str pid_project_id: The pid project id.
    :param str status: The status of Pid Loop.
    """
    ...
