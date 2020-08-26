# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['FlowLog']


class FlowLog(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 flow_log_name: Optional[pulumi.Input[str]] = None,
                 log_store_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        This resource used to create a flow log function in Cloud Enterprise Network (CEN).
        By using the flow log function, you can capture the traffic data of the network instances in different regions of a CEN.
        You can also use the data aggregated in flow logs to analyze cross-region traffic flows, minimize traffic costs, and troubleshoot network faults.

        For information about CEN flow log and how to use it, see [Manage CEN flowlog](https://www.alibabacloud.com/help/doc-detail/123006.htm).

        > **NOTE:** Available in 1.73.0+

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Create a cen flowlog resource and use it to publish a route entry pointing to an ECS.
        default_instance = alicloud.cen.Instance("defaultInstance")
        default_project = alicloud.log.Project("defaultProject", description="create by terraform")
        default_store = alicloud.log.Store("defaultStore",
            project=default_project.name,
            retention_period=3650,
            shard_count=3,
            auto_split=True,
            max_split_shard_count=60,
            append_meta=True)
        default_flow_log = alicloud.cen.FlowLog("defaultFlowLog",
            flow_log_name="my-flowlog",
            cen_id=default_instance.id,
            project_name=default_project.name,
            log_store_name=default_store.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_id: The ID of the CEN Instance.
        :param pulumi.Input[str] description: The description of flowlog.
        :param pulumi.Input[str] flow_log_name: The name of flowlog.
        :param pulumi.Input[str] log_store_name: The name of the log store which is in the  `project_name` SLS project.
        :param pulumi.Input[str] project_name: The name of the SLS project.
        :param pulumi.Input[str] status: The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
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

            if cen_id is None:
                raise TypeError("Missing required property 'cen_id'")
            __props__['cen_id'] = cen_id
            __props__['description'] = description
            __props__['flow_log_name'] = flow_log_name
            if log_store_name is None:
                raise TypeError("Missing required property 'log_store_name'")
            __props__['log_store_name'] = log_store_name
            if project_name is None:
                raise TypeError("Missing required property 'project_name'")
            __props__['project_name'] = project_name
            __props__['status'] = status
        super(FlowLog, __self__).__init__(
            'alicloud:cen/flowLog:FlowLog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cen_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            flow_log_name: Optional[pulumi.Input[str]] = None,
            log_store_name: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'FlowLog':
        """
        Get an existing FlowLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_id: The ID of the CEN Instance.
        :param pulumi.Input[str] description: The description of flowlog.
        :param pulumi.Input[str] flow_log_name: The name of flowlog.
        :param pulumi.Input[str] log_store_name: The name of the log store which is in the  `project_name` SLS project.
        :param pulumi.Input[str] project_name: The name of the SLS project.
        :param pulumi.Input[str] status: The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cen_id"] = cen_id
        __props__["description"] = description
        __props__["flow_log_name"] = flow_log_name
        __props__["log_store_name"] = log_store_name
        __props__["project_name"] = project_name
        __props__["status"] = status
        return FlowLog(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> str:
        """
        The ID of the CEN Instance.
        """
        return pulumi.get(self, "cen_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of flowlog.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="flowLogName")
    def flow_log_name(self) -> Optional[str]:
        """
        The name of flowlog.
        """
        return pulumi.get(self, "flow_log_name")

    @property
    @pulumi.getter(name="logStoreName")
    def log_store_name(self) -> str:
        """
        The name of the log store which is in the  `project_name` SLS project.
        """
        return pulumi.get(self, "log_store_name")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> str:
        """
        The name of the SLS project.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
        """
        return pulumi.get(self, "status")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

