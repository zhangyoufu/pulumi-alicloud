# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['JobTemplateArgs', 'JobTemplate']

@pulumi.input_type
class JobTemplateArgs:
    def __init__(__self__, *,
                 command_line: pulumi.Input[str],
                 job_template_name: pulumi.Input[str],
                 array_request: Optional[pulumi.Input[str]] = None,
                 clock_time: Optional[pulumi.Input[str]] = None,
                 gpu: Optional[pulumi.Input[int]] = None,
                 mem: Optional[pulumi.Input[str]] = None,
                 node: Optional[pulumi.Input[int]] = None,
                 package_path: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 re_runable: Optional[pulumi.Input[bool]] = None,
                 runas_user: Optional[pulumi.Input[str]] = None,
                 stderr_redirect_path: Optional[pulumi.Input[str]] = None,
                 stdout_redirect_path: Optional[pulumi.Input[str]] = None,
                 task: Optional[pulumi.Input[int]] = None,
                 thread: Optional[pulumi.Input[int]] = None,
                 variables: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a JobTemplate resource.
        :param pulumi.Input[str] command_line: Job Commands.
        :param pulumi.Input[str] job_template_name: A Job Template Name.
        :param pulumi.Input[str] array_request: Queue Jobs, Is of the Form: 1-10:2.
        :param pulumi.Input[str] clock_time: Job Maximum Run Time.
        :param pulumi.Input[int] gpu: A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
        :param pulumi.Input[str] mem: A Single Compute Node Maximum Memory.
        :param pulumi.Input[int] node: Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
        :param pulumi.Input[str] package_path: Job Commands the Directory.
        :param pulumi.Input[int] priority: The Job Priority.
        :param pulumi.Input[str] queue: The Job Queue.
        :param pulumi.Input[bool] re_runable: If the Job Is Support for the Re-Run.
        :param pulumi.Input[str] runas_user: The name of the user who performed the job.
        :param pulumi.Input[str] stderr_redirect_path: Error Output Path.
        :param pulumi.Input[str] stdout_redirect_path: Standard Output Path and.
        :param pulumi.Input[int] task: A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
        :param pulumi.Input[int] thread: A Single Task and the Number of Required Threads.
        :param pulumi.Input[str] variables: The Job of the Environment Variable.
        """
        pulumi.set(__self__, "command_line", command_line)
        pulumi.set(__self__, "job_template_name", job_template_name)
        if array_request is not None:
            pulumi.set(__self__, "array_request", array_request)
        if clock_time is not None:
            pulumi.set(__self__, "clock_time", clock_time)
        if gpu is not None:
            pulumi.set(__self__, "gpu", gpu)
        if mem is not None:
            pulumi.set(__self__, "mem", mem)
        if node is not None:
            pulumi.set(__self__, "node", node)
        if package_path is not None:
            pulumi.set(__self__, "package_path", package_path)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if re_runable is not None:
            pulumi.set(__self__, "re_runable", re_runable)
        if runas_user is not None:
            pulumi.set(__self__, "runas_user", runas_user)
        if stderr_redirect_path is not None:
            pulumi.set(__self__, "stderr_redirect_path", stderr_redirect_path)
        if stdout_redirect_path is not None:
            pulumi.set(__self__, "stdout_redirect_path", stdout_redirect_path)
        if task is not None:
            pulumi.set(__self__, "task", task)
        if thread is not None:
            pulumi.set(__self__, "thread", thread)
        if variables is not None:
            pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter(name="commandLine")
    def command_line(self) -> pulumi.Input[str]:
        """
        Job Commands.
        """
        return pulumi.get(self, "command_line")

    @command_line.setter
    def command_line(self, value: pulumi.Input[str]):
        pulumi.set(self, "command_line", value)

    @property
    @pulumi.getter(name="jobTemplateName")
    def job_template_name(self) -> pulumi.Input[str]:
        """
        A Job Template Name.
        """
        return pulumi.get(self, "job_template_name")

    @job_template_name.setter
    def job_template_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "job_template_name", value)

    @property
    @pulumi.getter(name="arrayRequest")
    def array_request(self) -> Optional[pulumi.Input[str]]:
        """
        Queue Jobs, Is of the Form: 1-10:2.
        """
        return pulumi.get(self, "array_request")

    @array_request.setter
    def array_request(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "array_request", value)

    @property
    @pulumi.getter(name="clockTime")
    def clock_time(self) -> Optional[pulumi.Input[str]]:
        """
        Job Maximum Run Time.
        """
        return pulumi.get(self, "clock_time")

    @clock_time.setter
    def clock_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "clock_time", value)

    @property
    @pulumi.getter
    def gpu(self) -> Optional[pulumi.Input[int]]:
        """
        A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
        """
        return pulumi.get(self, "gpu")

    @gpu.setter
    def gpu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "gpu", value)

    @property
    @pulumi.getter
    def mem(self) -> Optional[pulumi.Input[str]]:
        """
        A Single Compute Node Maximum Memory.
        """
        return pulumi.get(self, "mem")

    @mem.setter
    def mem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mem", value)

    @property
    @pulumi.getter
    def node(self) -> Optional[pulumi.Input[int]]:
        """
        Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
        """
        return pulumi.get(self, "node")

    @node.setter
    def node(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node", value)

    @property
    @pulumi.getter(name="packagePath")
    def package_path(self) -> Optional[pulumi.Input[str]]:
        """
        Job Commands the Directory.
        """
        return pulumi.get(self, "package_path")

    @package_path.setter
    def package_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_path", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The Job Priority.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def queue(self) -> Optional[pulumi.Input[str]]:
        """
        The Job Queue.
        """
        return pulumi.get(self, "queue")

    @queue.setter
    def queue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue", value)

    @property
    @pulumi.getter(name="reRunable")
    def re_runable(self) -> Optional[pulumi.Input[bool]]:
        """
        If the Job Is Support for the Re-Run.
        """
        return pulumi.get(self, "re_runable")

    @re_runable.setter
    def re_runable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "re_runable", value)

    @property
    @pulumi.getter(name="runasUser")
    def runas_user(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user who performed the job.
        """
        return pulumi.get(self, "runas_user")

    @runas_user.setter
    def runas_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "runas_user", value)

    @property
    @pulumi.getter(name="stderrRedirectPath")
    def stderr_redirect_path(self) -> Optional[pulumi.Input[str]]:
        """
        Error Output Path.
        """
        return pulumi.get(self, "stderr_redirect_path")

    @stderr_redirect_path.setter
    def stderr_redirect_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stderr_redirect_path", value)

    @property
    @pulumi.getter(name="stdoutRedirectPath")
    def stdout_redirect_path(self) -> Optional[pulumi.Input[str]]:
        """
        Standard Output Path and.
        """
        return pulumi.get(self, "stdout_redirect_path")

    @stdout_redirect_path.setter
    def stdout_redirect_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stdout_redirect_path", value)

    @property
    @pulumi.getter
    def task(self) -> Optional[pulumi.Input[int]]:
        """
        A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
        """
        return pulumi.get(self, "task")

    @task.setter
    def task(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "task", value)

    @property
    @pulumi.getter
    def thread(self) -> Optional[pulumi.Input[int]]:
        """
        A Single Task and the Number of Required Threads.
        """
        return pulumi.get(self, "thread")

    @thread.setter
    def thread(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "thread", value)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[str]]:
        """
        The Job of the Environment Variable.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variables", value)


@pulumi.input_type
class _JobTemplateState:
    def __init__(__self__, *,
                 array_request: Optional[pulumi.Input[str]] = None,
                 clock_time: Optional[pulumi.Input[str]] = None,
                 command_line: Optional[pulumi.Input[str]] = None,
                 gpu: Optional[pulumi.Input[int]] = None,
                 job_template_name: Optional[pulumi.Input[str]] = None,
                 mem: Optional[pulumi.Input[str]] = None,
                 node: Optional[pulumi.Input[int]] = None,
                 package_path: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 re_runable: Optional[pulumi.Input[bool]] = None,
                 runas_user: Optional[pulumi.Input[str]] = None,
                 stderr_redirect_path: Optional[pulumi.Input[str]] = None,
                 stdout_redirect_path: Optional[pulumi.Input[str]] = None,
                 task: Optional[pulumi.Input[int]] = None,
                 thread: Optional[pulumi.Input[int]] = None,
                 variables: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering JobTemplate resources.
        :param pulumi.Input[str] array_request: Queue Jobs, Is of the Form: 1-10:2.
        :param pulumi.Input[str] clock_time: Job Maximum Run Time.
        :param pulumi.Input[str] command_line: Job Commands.
        :param pulumi.Input[int] gpu: A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
        :param pulumi.Input[str] job_template_name: A Job Template Name.
        :param pulumi.Input[str] mem: A Single Compute Node Maximum Memory.
        :param pulumi.Input[int] node: Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
        :param pulumi.Input[str] package_path: Job Commands the Directory.
        :param pulumi.Input[int] priority: The Job Priority.
        :param pulumi.Input[str] queue: The Job Queue.
        :param pulumi.Input[bool] re_runable: If the Job Is Support for the Re-Run.
        :param pulumi.Input[str] runas_user: The name of the user who performed the job.
        :param pulumi.Input[str] stderr_redirect_path: Error Output Path.
        :param pulumi.Input[str] stdout_redirect_path: Standard Output Path and.
        :param pulumi.Input[int] task: A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
        :param pulumi.Input[int] thread: A Single Task and the Number of Required Threads.
        :param pulumi.Input[str] variables: The Job of the Environment Variable.
        """
        if array_request is not None:
            pulumi.set(__self__, "array_request", array_request)
        if clock_time is not None:
            pulumi.set(__self__, "clock_time", clock_time)
        if command_line is not None:
            pulumi.set(__self__, "command_line", command_line)
        if gpu is not None:
            pulumi.set(__self__, "gpu", gpu)
        if job_template_name is not None:
            pulumi.set(__self__, "job_template_name", job_template_name)
        if mem is not None:
            pulumi.set(__self__, "mem", mem)
        if node is not None:
            pulumi.set(__self__, "node", node)
        if package_path is not None:
            pulumi.set(__self__, "package_path", package_path)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if re_runable is not None:
            pulumi.set(__self__, "re_runable", re_runable)
        if runas_user is not None:
            pulumi.set(__self__, "runas_user", runas_user)
        if stderr_redirect_path is not None:
            pulumi.set(__self__, "stderr_redirect_path", stderr_redirect_path)
        if stdout_redirect_path is not None:
            pulumi.set(__self__, "stdout_redirect_path", stdout_redirect_path)
        if task is not None:
            pulumi.set(__self__, "task", task)
        if thread is not None:
            pulumi.set(__self__, "thread", thread)
        if variables is not None:
            pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter(name="arrayRequest")
    def array_request(self) -> Optional[pulumi.Input[str]]:
        """
        Queue Jobs, Is of the Form: 1-10:2.
        """
        return pulumi.get(self, "array_request")

    @array_request.setter
    def array_request(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "array_request", value)

    @property
    @pulumi.getter(name="clockTime")
    def clock_time(self) -> Optional[pulumi.Input[str]]:
        """
        Job Maximum Run Time.
        """
        return pulumi.get(self, "clock_time")

    @clock_time.setter
    def clock_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "clock_time", value)

    @property
    @pulumi.getter(name="commandLine")
    def command_line(self) -> Optional[pulumi.Input[str]]:
        """
        Job Commands.
        """
        return pulumi.get(self, "command_line")

    @command_line.setter
    def command_line(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command_line", value)

    @property
    @pulumi.getter
    def gpu(self) -> Optional[pulumi.Input[int]]:
        """
        A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
        """
        return pulumi.get(self, "gpu")

    @gpu.setter
    def gpu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "gpu", value)

    @property
    @pulumi.getter(name="jobTemplateName")
    def job_template_name(self) -> Optional[pulumi.Input[str]]:
        """
        A Job Template Name.
        """
        return pulumi.get(self, "job_template_name")

    @job_template_name.setter
    def job_template_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_template_name", value)

    @property
    @pulumi.getter
    def mem(self) -> Optional[pulumi.Input[str]]:
        """
        A Single Compute Node Maximum Memory.
        """
        return pulumi.get(self, "mem")

    @mem.setter
    def mem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mem", value)

    @property
    @pulumi.getter
    def node(self) -> Optional[pulumi.Input[int]]:
        """
        Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
        """
        return pulumi.get(self, "node")

    @node.setter
    def node(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node", value)

    @property
    @pulumi.getter(name="packagePath")
    def package_path(self) -> Optional[pulumi.Input[str]]:
        """
        Job Commands the Directory.
        """
        return pulumi.get(self, "package_path")

    @package_path.setter
    def package_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_path", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The Job Priority.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def queue(self) -> Optional[pulumi.Input[str]]:
        """
        The Job Queue.
        """
        return pulumi.get(self, "queue")

    @queue.setter
    def queue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue", value)

    @property
    @pulumi.getter(name="reRunable")
    def re_runable(self) -> Optional[pulumi.Input[bool]]:
        """
        If the Job Is Support for the Re-Run.
        """
        return pulumi.get(self, "re_runable")

    @re_runable.setter
    def re_runable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "re_runable", value)

    @property
    @pulumi.getter(name="runasUser")
    def runas_user(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user who performed the job.
        """
        return pulumi.get(self, "runas_user")

    @runas_user.setter
    def runas_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "runas_user", value)

    @property
    @pulumi.getter(name="stderrRedirectPath")
    def stderr_redirect_path(self) -> Optional[pulumi.Input[str]]:
        """
        Error Output Path.
        """
        return pulumi.get(self, "stderr_redirect_path")

    @stderr_redirect_path.setter
    def stderr_redirect_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stderr_redirect_path", value)

    @property
    @pulumi.getter(name="stdoutRedirectPath")
    def stdout_redirect_path(self) -> Optional[pulumi.Input[str]]:
        """
        Standard Output Path and.
        """
        return pulumi.get(self, "stdout_redirect_path")

    @stdout_redirect_path.setter
    def stdout_redirect_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stdout_redirect_path", value)

    @property
    @pulumi.getter
    def task(self) -> Optional[pulumi.Input[int]]:
        """
        A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
        """
        return pulumi.get(self, "task")

    @task.setter
    def task(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "task", value)

    @property
    @pulumi.getter
    def thread(self) -> Optional[pulumi.Input[int]]:
        """
        A Single Task and the Number of Required Threads.
        """
        return pulumi.get(self, "thread")

    @thread.setter
    def thread(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "thread", value)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[str]]:
        """
        The Job of the Environment Variable.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variables", value)


class JobTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 array_request: Optional[pulumi.Input[str]] = None,
                 clock_time: Optional[pulumi.Input[str]] = None,
                 command_line: Optional[pulumi.Input[str]] = None,
                 gpu: Optional[pulumi.Input[int]] = None,
                 job_template_name: Optional[pulumi.Input[str]] = None,
                 mem: Optional[pulumi.Input[str]] = None,
                 node: Optional[pulumi.Input[int]] = None,
                 package_path: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 re_runable: Optional[pulumi.Input[bool]] = None,
                 runas_user: Optional[pulumi.Input[str]] = None,
                 stderr_redirect_path: Optional[pulumi.Input[str]] = None,
                 stdout_redirect_path: Optional[pulumi.Input[str]] = None,
                 task: Optional[pulumi.Input[int]] = None,
                 thread: Optional[pulumi.Input[int]] = None,
                 variables: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Ehpc Job Template resource.

        For information about Ehpc Job Template and how to use it, see [What is Job Template](https://www.alibabacloud.com/help/product/57664.html).

        > **NOTE:** Available in v1.133.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.ehpc.JobTemplate("default",
            command_line="./LammpsTest/lammps.pbs",
            job_template_name="example_value")
        ```

        ## Import

        Ehpc Job Template can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ehpc/jobTemplate:JobTemplate example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] array_request: Queue Jobs, Is of the Form: 1-10:2.
        :param pulumi.Input[str] clock_time: Job Maximum Run Time.
        :param pulumi.Input[str] command_line: Job Commands.
        :param pulumi.Input[int] gpu: A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
        :param pulumi.Input[str] job_template_name: A Job Template Name.
        :param pulumi.Input[str] mem: A Single Compute Node Maximum Memory.
        :param pulumi.Input[int] node: Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
        :param pulumi.Input[str] package_path: Job Commands the Directory.
        :param pulumi.Input[int] priority: The Job Priority.
        :param pulumi.Input[str] queue: The Job Queue.
        :param pulumi.Input[bool] re_runable: If the Job Is Support for the Re-Run.
        :param pulumi.Input[str] runas_user: The name of the user who performed the job.
        :param pulumi.Input[str] stderr_redirect_path: Error Output Path.
        :param pulumi.Input[str] stdout_redirect_path: Standard Output Path and.
        :param pulumi.Input[int] task: A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
        :param pulumi.Input[int] thread: A Single Task and the Number of Required Threads.
        :param pulumi.Input[str] variables: The Job of the Environment Variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Ehpc Job Template resource.

        For information about Ehpc Job Template and how to use it, see [What is Job Template](https://www.alibabacloud.com/help/product/57664.html).

        > **NOTE:** Available in v1.133.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.ehpc.JobTemplate("default",
            command_line="./LammpsTest/lammps.pbs",
            job_template_name="example_value")
        ```

        ## Import

        Ehpc Job Template can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:ehpc/jobTemplate:JobTemplate example <id>
        ```

        :param str resource_name: The name of the resource.
        :param JobTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 array_request: Optional[pulumi.Input[str]] = None,
                 clock_time: Optional[pulumi.Input[str]] = None,
                 command_line: Optional[pulumi.Input[str]] = None,
                 gpu: Optional[pulumi.Input[int]] = None,
                 job_template_name: Optional[pulumi.Input[str]] = None,
                 mem: Optional[pulumi.Input[str]] = None,
                 node: Optional[pulumi.Input[int]] = None,
                 package_path: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 re_runable: Optional[pulumi.Input[bool]] = None,
                 runas_user: Optional[pulumi.Input[str]] = None,
                 stderr_redirect_path: Optional[pulumi.Input[str]] = None,
                 stdout_redirect_path: Optional[pulumi.Input[str]] = None,
                 task: Optional[pulumi.Input[int]] = None,
                 thread: Optional[pulumi.Input[int]] = None,
                 variables: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = JobTemplateArgs.__new__(JobTemplateArgs)

            __props__.__dict__["array_request"] = array_request
            __props__.__dict__["clock_time"] = clock_time
            if command_line is None and not opts.urn:
                raise TypeError("Missing required property 'command_line'")
            __props__.__dict__["command_line"] = command_line
            __props__.__dict__["gpu"] = gpu
            if job_template_name is None and not opts.urn:
                raise TypeError("Missing required property 'job_template_name'")
            __props__.__dict__["job_template_name"] = job_template_name
            __props__.__dict__["mem"] = mem
            __props__.__dict__["node"] = node
            __props__.__dict__["package_path"] = package_path
            __props__.__dict__["priority"] = priority
            __props__.__dict__["queue"] = queue
            __props__.__dict__["re_runable"] = re_runable
            __props__.__dict__["runas_user"] = runas_user
            __props__.__dict__["stderr_redirect_path"] = stderr_redirect_path
            __props__.__dict__["stdout_redirect_path"] = stdout_redirect_path
            __props__.__dict__["task"] = task
            __props__.__dict__["thread"] = thread
            __props__.__dict__["variables"] = variables
        super(JobTemplate, __self__).__init__(
            'alicloud:ehpc/jobTemplate:JobTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            array_request: Optional[pulumi.Input[str]] = None,
            clock_time: Optional[pulumi.Input[str]] = None,
            command_line: Optional[pulumi.Input[str]] = None,
            gpu: Optional[pulumi.Input[int]] = None,
            job_template_name: Optional[pulumi.Input[str]] = None,
            mem: Optional[pulumi.Input[str]] = None,
            node: Optional[pulumi.Input[int]] = None,
            package_path: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            queue: Optional[pulumi.Input[str]] = None,
            re_runable: Optional[pulumi.Input[bool]] = None,
            runas_user: Optional[pulumi.Input[str]] = None,
            stderr_redirect_path: Optional[pulumi.Input[str]] = None,
            stdout_redirect_path: Optional[pulumi.Input[str]] = None,
            task: Optional[pulumi.Input[int]] = None,
            thread: Optional[pulumi.Input[int]] = None,
            variables: Optional[pulumi.Input[str]] = None) -> 'JobTemplate':
        """
        Get an existing JobTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] array_request: Queue Jobs, Is of the Form: 1-10:2.
        :param pulumi.Input[str] clock_time: Job Maximum Run Time.
        :param pulumi.Input[str] command_line: Job Commands.
        :param pulumi.Input[int] gpu: A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
        :param pulumi.Input[str] job_template_name: A Job Template Name.
        :param pulumi.Input[str] mem: A Single Compute Node Maximum Memory.
        :param pulumi.Input[int] node: Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
        :param pulumi.Input[str] package_path: Job Commands the Directory.
        :param pulumi.Input[int] priority: The Job Priority.
        :param pulumi.Input[str] queue: The Job Queue.
        :param pulumi.Input[bool] re_runable: If the Job Is Support for the Re-Run.
        :param pulumi.Input[str] runas_user: The name of the user who performed the job.
        :param pulumi.Input[str] stderr_redirect_path: Error Output Path.
        :param pulumi.Input[str] stdout_redirect_path: Standard Output Path and.
        :param pulumi.Input[int] task: A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
        :param pulumi.Input[int] thread: A Single Task and the Number of Required Threads.
        :param pulumi.Input[str] variables: The Job of the Environment Variable.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _JobTemplateState.__new__(_JobTemplateState)

        __props__.__dict__["array_request"] = array_request
        __props__.__dict__["clock_time"] = clock_time
        __props__.__dict__["command_line"] = command_line
        __props__.__dict__["gpu"] = gpu
        __props__.__dict__["job_template_name"] = job_template_name
        __props__.__dict__["mem"] = mem
        __props__.__dict__["node"] = node
        __props__.__dict__["package_path"] = package_path
        __props__.__dict__["priority"] = priority
        __props__.__dict__["queue"] = queue
        __props__.__dict__["re_runable"] = re_runable
        __props__.__dict__["runas_user"] = runas_user
        __props__.__dict__["stderr_redirect_path"] = stderr_redirect_path
        __props__.__dict__["stdout_redirect_path"] = stdout_redirect_path
        __props__.__dict__["task"] = task
        __props__.__dict__["thread"] = thread
        __props__.__dict__["variables"] = variables
        return JobTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="arrayRequest")
    def array_request(self) -> pulumi.Output[Optional[str]]:
        """
        Queue Jobs, Is of the Form: 1-10:2.
        """
        return pulumi.get(self, "array_request")

    @property
    @pulumi.getter(name="clockTime")
    def clock_time(self) -> pulumi.Output[Optional[str]]:
        """
        Job Maximum Run Time.
        """
        return pulumi.get(self, "clock_time")

    @property
    @pulumi.getter(name="commandLine")
    def command_line(self) -> pulumi.Output[str]:
        """
        Job Commands.
        """
        return pulumi.get(self, "command_line")

    @property
    @pulumi.getter
    def gpu(self) -> pulumi.Output[Optional[int]]:
        """
        A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
        """
        return pulumi.get(self, "gpu")

    @property
    @pulumi.getter(name="jobTemplateName")
    def job_template_name(self) -> pulumi.Output[str]:
        """
        A Job Template Name.
        """
        return pulumi.get(self, "job_template_name")

    @property
    @pulumi.getter
    def mem(self) -> pulumi.Output[Optional[str]]:
        """
        A Single Compute Node Maximum Memory.
        """
        return pulumi.get(self, "mem")

    @property
    @pulumi.getter
    def node(self) -> pulumi.Output[Optional[int]]:
        """
        Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
        """
        return pulumi.get(self, "node")

    @property
    @pulumi.getter(name="packagePath")
    def package_path(self) -> pulumi.Output[Optional[str]]:
        """
        Job Commands the Directory.
        """
        return pulumi.get(self, "package_path")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[int]]:
        """
        The Job Priority.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def queue(self) -> pulumi.Output[Optional[str]]:
        """
        The Job Queue.
        """
        return pulumi.get(self, "queue")

    @property
    @pulumi.getter(name="reRunable")
    def re_runable(self) -> pulumi.Output[bool]:
        """
        If the Job Is Support for the Re-Run.
        """
        return pulumi.get(self, "re_runable")

    @property
    @pulumi.getter(name="runasUser")
    def runas_user(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the user who performed the job.
        """
        return pulumi.get(self, "runas_user")

    @property
    @pulumi.getter(name="stderrRedirectPath")
    def stderr_redirect_path(self) -> pulumi.Output[Optional[str]]:
        """
        Error Output Path.
        """
        return pulumi.get(self, "stderr_redirect_path")

    @property
    @pulumi.getter(name="stdoutRedirectPath")
    def stdout_redirect_path(self) -> pulumi.Output[Optional[str]]:
        """
        Standard Output Path and.
        """
        return pulumi.get(self, "stdout_redirect_path")

    @property
    @pulumi.getter
    def task(self) -> pulumi.Output[Optional[int]]:
        """
        A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
        """
        return pulumi.get(self, "task")

    @property
    @pulumi.getter
    def thread(self) -> pulumi.Output[Optional[int]]:
        """
        A Single Task and the Number of Required Threads.
        """
        return pulumi.get(self, "thread")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Optional[str]]:
        """
        The Job of the Environment Variable.
        """
        return pulumi.get(self, "variables")

