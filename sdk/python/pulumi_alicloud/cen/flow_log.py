# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class FlowLog(pulumi.CustomResource):
    cen_id: pulumi.Output[str]
    """
    The ID of the CEN Instance.
    """
    description: pulumi.Output[str]
    """
    The description of flowlog.
    """
    flow_log_name: pulumi.Output[str]
    """
    The name of flowlog.
    """
    log_store_name: pulumi.Output[str]
    """
    The name of the log store which is in the  `project_name` SLS project.
    """
    project_name: pulumi.Output[str]
    """
    The name of the SLS project.
    """
    status: pulumi.Output[str]
    """
    The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
    """
    def __init__(__self__, resource_name, opts=None, cen_id=None, description=None, flow_log_name=None, log_store_name=None, project_name=None, status=None, __props__=None, __name__=None, __opts__=None):
        """
        This resource used to create a flow log function in Cloud Enterprise Network (CEN). 
        By using the flow log function, you can capture the traffic data of the network instances in different regions of a CEN. 
        You can also use the data aggregated in flow logs to analyze cross-region traffic flows, minimize traffic costs, and troubleshoot network faults.

        For information about CEN flow log and how to use it, see [Manage CEN flowlog](https://www.alibabacloud.com/help/doc-detail/123006.htm).

        > **NOTE:** Available in 1.73.0+



        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cen_flowlog.html.markdown.

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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, cen_id=None, description=None, flow_log_name=None, log_store_name=None, project_name=None, status=None):
        """
        Get an existing FlowLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

