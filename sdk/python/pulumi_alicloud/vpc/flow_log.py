# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FlowLogArgs', 'FlowLog']

@pulumi.input_type
class FlowLogArgs:
    def __init__(__self__, *,
                 log_store_name: pulumi.Input[str],
                 project_name: pulumi.Input[str],
                 resource_id: pulumi.Input[str],
                 resource_type: pulumi.Input[str],
                 traffic_type: pulumi.Input[str],
                 aggregation_interval: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 flow_log_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 traffic_paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a FlowLog resource.
        :param pulumi.Input[str] log_store_name: The name of the logstore.
        :param pulumi.Input[str] project_name: The name of the project.
        :param pulumi.Input[str] resource_id: The ID of the resource.
        :param pulumi.Input[str] resource_type: The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
        :param pulumi.Input[str] traffic_type: The type of traffic collected. Valid values:**All**: All traffic.**Allow**: Access control allowedtraffic.**Drop**: Access control denied traffic.
        :param pulumi.Input[str] aggregation_interval: Data aggregation interval.
        :param pulumi.Input[str] description: The Description of the VPC Flow Log.
        :param pulumi.Input[str] flow_log_name: The Name of the VPC Flow Log.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] status: The status of the VPC Flow Log. Valid values: **Active** and **Inactive**.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the current instance resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] traffic_paths: The collected flow path. Value:**all**: indicates full acquisition.**internetGateway**: indicates public network traffic collection.
        """
        pulumi.set(__self__, "log_store_name", log_store_name)
        pulumi.set(__self__, "project_name", project_name)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "traffic_type", traffic_type)
        if aggregation_interval is not None:
            pulumi.set(__self__, "aggregation_interval", aggregation_interval)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if flow_log_name is not None:
            pulumi.set(__self__, "flow_log_name", flow_log_name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if traffic_paths is not None:
            pulumi.set(__self__, "traffic_paths", traffic_paths)

    @property
    @pulumi.getter(name="logStoreName")
    def log_store_name(self) -> pulumi.Input[str]:
        """
        The name of the logstore.
        """
        return pulumi.get(self, "log_store_name")

    @log_store_name.setter
    def log_store_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_store_name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Input[str]:
        """
        The name of the project.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        """
        The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> pulumi.Input[str]:
        """
        The type of traffic collected. Valid values:**All**: All traffic.**Allow**: Access control allowedtraffic.**Drop**: Access control denied traffic.
        """
        return pulumi.get(self, "traffic_type")

    @traffic_type.setter
    def traffic_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "traffic_type", value)

    @property
    @pulumi.getter(name="aggregationInterval")
    def aggregation_interval(self) -> Optional[pulumi.Input[str]]:
        """
        Data aggregation interval.
        """
        return pulumi.get(self, "aggregation_interval")

    @aggregation_interval.setter
    def aggregation_interval(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aggregation_interval", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Description of the VPC Flow Log.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="flowLogName")
    def flow_log_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the VPC Flow Log.
        """
        return pulumi.get(self, "flow_log_name")

    @flow_log_name.setter
    def flow_log_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flow_log_name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the VPC Flow Log. Valid values: **Active** and **Inactive**.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tag of the current instance resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="trafficPaths")
    def traffic_paths(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The collected flow path. Value:**all**: indicates full acquisition.**internetGateway**: indicates public network traffic collection.
        """
        return pulumi.get(self, "traffic_paths")

    @traffic_paths.setter
    def traffic_paths(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "traffic_paths", value)


@pulumi.input_type
class _FlowLogState:
    def __init__(__self__, *,
                 aggregation_interval: Optional[pulumi.Input[str]] = None,
                 business_status: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 flow_log_id: Optional[pulumi.Input[str]] = None,
                 flow_log_name: Optional[pulumi.Input[str]] = None,
                 log_store_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 traffic_paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FlowLog resources.
        :param pulumi.Input[str] aggregation_interval: Data aggregation interval.
        :param pulumi.Input[str] business_status: Business status.
        :param pulumi.Input[str] create_time: Creation time.
        :param pulumi.Input[str] description: The Description of the VPC Flow Log.
        :param pulumi.Input[str] flow_log_id: The flow log ID.
        :param pulumi.Input[str] flow_log_name: The Name of the VPC Flow Log.
        :param pulumi.Input[str] log_store_name: The name of the logstore.
        :param pulumi.Input[str] project_name: The name of the project.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] resource_id: The ID of the resource.
        :param pulumi.Input[str] resource_type: The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
        :param pulumi.Input[str] status: The status of the VPC Flow Log. Valid values: **Active** and **Inactive**.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the current instance resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] traffic_paths: The collected flow path. Value:**all**: indicates full acquisition.**internetGateway**: indicates public network traffic collection.
        :param pulumi.Input[str] traffic_type: The type of traffic collected. Valid values:**All**: All traffic.**Allow**: Access control allowedtraffic.**Drop**: Access control denied traffic.
        """
        if aggregation_interval is not None:
            pulumi.set(__self__, "aggregation_interval", aggregation_interval)
        if business_status is not None:
            pulumi.set(__self__, "business_status", business_status)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if flow_log_id is not None:
            pulumi.set(__self__, "flow_log_id", flow_log_id)
        if flow_log_name is not None:
            pulumi.set(__self__, "flow_log_name", flow_log_name)
        if log_store_name is not None:
            pulumi.set(__self__, "log_store_name", log_store_name)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if traffic_paths is not None:
            pulumi.set(__self__, "traffic_paths", traffic_paths)
        if traffic_type is not None:
            pulumi.set(__self__, "traffic_type", traffic_type)

    @property
    @pulumi.getter(name="aggregationInterval")
    def aggregation_interval(self) -> Optional[pulumi.Input[str]]:
        """
        Data aggregation interval.
        """
        return pulumi.get(self, "aggregation_interval")

    @aggregation_interval.setter
    def aggregation_interval(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aggregation_interval", value)

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> Optional[pulumi.Input[str]]:
        """
        Business status.
        """
        return pulumi.get(self, "business_status")

    @business_status.setter
    def business_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "business_status", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The Description of the VPC Flow Log.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="flowLogId")
    def flow_log_id(self) -> Optional[pulumi.Input[str]]:
        """
        The flow log ID.
        """
        return pulumi.get(self, "flow_log_id")

    @flow_log_id.setter
    def flow_log_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flow_log_id", value)

    @property
    @pulumi.getter(name="flowLogName")
    def flow_log_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the VPC Flow Log.
        """
        return pulumi.get(self, "flow_log_name")

    @flow_log_name.setter
    def flow_log_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flow_log_name", value)

    @property
    @pulumi.getter(name="logStoreName")
    def log_store_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the logstore.
        """
        return pulumi.get(self, "log_store_name")

    @log_store_name.setter
    def log_store_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_store_name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the project.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the VPC Flow Log. Valid values: **Active** and **Inactive**.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tag of the current instance resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="trafficPaths")
    def traffic_paths(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The collected flow path. Value:**all**: indicates full acquisition.**internetGateway**: indicates public network traffic collection.
        """
        return pulumi.get(self, "traffic_paths")

    @traffic_paths.setter
    def traffic_paths(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "traffic_paths", value)

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of traffic collected. Valid values:**All**: All traffic.**Allow**: Access control allowedtraffic.**Drop**: Access control denied traffic.
        """
        return pulumi.get(self, "traffic_type")

    @traffic_type.setter
    def traffic_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_type", value)


class FlowLog(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aggregation_interval: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 flow_log_name: Optional[pulumi.Input[str]] = None,
                 log_store_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 traffic_paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Vpc Flow Log resource. While it uses vpc.FlowLog to build a vpc flow log resource, it will be active by default.

        For information about Vpc Flow Log and how to use it, see [What is Flow Log](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/flow-logs-overview).

        > **NOTE:** Available in v1.117.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-testacc-example"
        default_rg = alicloud.resourcemanager.ResourceGroup("defaultRg",
            resource_group_name=name,
            display_name="tf-testAcc-rg78")
        default_vpc = alicloud.vpc.Network("defaultVpc",
            vpc_name=f"{name}1",
            cidr_block="10.0.0.0/8")
        modify_rg = alicloud.resourcemanager.ResourceGroup("modifyRG",
            display_name="tf-testAcc-rg405",
            resource_group_name=f"{name}2")
        default_project = alicloud.log.Project("defaultProject")
        default_store = alicloud.log.Store("defaultStore", project=default_project.name)
        default_flow_log = alicloud.vpc.FlowLog("defaultFlowLog",
            flow_log_name=name,
            log_store_name=default_store.name,
            description="tf-testAcc-flowlog",
            traffic_paths=["all"],
            project_name=default_project.name,
            resource_type="VPC",
            resource_group_id=default_rg.id,
            resource_id=default_vpc.id,
            aggregation_interval="1",
            traffic_type="All")
        ```

        ## Import

        Vpc Flow Log can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/flowLog:FlowLog example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aggregation_interval: Data aggregation interval.
        :param pulumi.Input[str] description: The Description of the VPC Flow Log.
        :param pulumi.Input[str] flow_log_name: The Name of the VPC Flow Log.
        :param pulumi.Input[str] log_store_name: The name of the logstore.
        :param pulumi.Input[str] project_name: The name of the project.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] resource_id: The ID of the resource.
        :param pulumi.Input[str] resource_type: The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
        :param pulumi.Input[str] status: The status of the VPC Flow Log. Valid values: **Active** and **Inactive**.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the current instance resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] traffic_paths: The collected flow path. Value:**all**: indicates full acquisition.**internetGateway**: indicates public network traffic collection.
        :param pulumi.Input[str] traffic_type: The type of traffic collected. Valid values:**All**: All traffic.**Allow**: Access control allowedtraffic.**Drop**: Access control denied traffic.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlowLogArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vpc Flow Log resource. While it uses vpc.FlowLog to build a vpc flow log resource, it will be active by default.

        For information about Vpc Flow Log and how to use it, see [What is Flow Log](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/flow-logs-overview).

        > **NOTE:** Available in v1.117.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-testacc-example"
        default_rg = alicloud.resourcemanager.ResourceGroup("defaultRg",
            resource_group_name=name,
            display_name="tf-testAcc-rg78")
        default_vpc = alicloud.vpc.Network("defaultVpc",
            vpc_name=f"{name}1",
            cidr_block="10.0.0.0/8")
        modify_rg = alicloud.resourcemanager.ResourceGroup("modifyRG",
            display_name="tf-testAcc-rg405",
            resource_group_name=f"{name}2")
        default_project = alicloud.log.Project("defaultProject")
        default_store = alicloud.log.Store("defaultStore", project=default_project.name)
        default_flow_log = alicloud.vpc.FlowLog("defaultFlowLog",
            flow_log_name=name,
            log_store_name=default_store.name,
            description="tf-testAcc-flowlog",
            traffic_paths=["all"],
            project_name=default_project.name,
            resource_type="VPC",
            resource_group_id=default_rg.id,
            resource_id=default_vpc.id,
            aggregation_interval="1",
            traffic_type="All")
        ```

        ## Import

        Vpc Flow Log can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/flowLog:FlowLog example <id>
        ```

        :param str resource_name: The name of the resource.
        :param FlowLogArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlowLogArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aggregation_interval: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 flow_log_name: Optional[pulumi.Input[str]] = None,
                 log_store_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 traffic_paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FlowLogArgs.__new__(FlowLogArgs)

            __props__.__dict__["aggregation_interval"] = aggregation_interval
            __props__.__dict__["description"] = description
            __props__.__dict__["flow_log_name"] = flow_log_name
            if log_store_name is None and not opts.urn:
                raise TypeError("Missing required property 'log_store_name'")
            __props__.__dict__["log_store_name"] = log_store_name
            if project_name is None and not opts.urn:
                raise TypeError("Missing required property 'project_name'")
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["resource_group_id"] = resource_group_id
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["traffic_paths"] = traffic_paths
            if traffic_type is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_type'")
            __props__.__dict__["traffic_type"] = traffic_type
            __props__.__dict__["business_status"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["flow_log_id"] = None
        super(FlowLog, __self__).__init__(
            'alicloud:vpc/flowLog:FlowLog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aggregation_interval: Optional[pulumi.Input[str]] = None,
            business_status: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            flow_log_id: Optional[pulumi.Input[str]] = None,
            flow_log_name: Optional[pulumi.Input[str]] = None,
            log_store_name: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            traffic_paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            traffic_type: Optional[pulumi.Input[str]] = None) -> 'FlowLog':
        """
        Get an existing FlowLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aggregation_interval: Data aggregation interval.
        :param pulumi.Input[str] business_status: Business status.
        :param pulumi.Input[str] create_time: Creation time.
        :param pulumi.Input[str] description: The Description of the VPC Flow Log.
        :param pulumi.Input[str] flow_log_id: The flow log ID.
        :param pulumi.Input[str] flow_log_name: The Name of the VPC Flow Log.
        :param pulumi.Input[str] log_store_name: The name of the logstore.
        :param pulumi.Input[str] project_name: The name of the project.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] resource_id: The ID of the resource.
        :param pulumi.Input[str] resource_type: The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
        :param pulumi.Input[str] status: The status of the VPC Flow Log. Valid values: **Active** and **Inactive**.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the current instance resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] traffic_paths: The collected flow path. Value:**all**: indicates full acquisition.**internetGateway**: indicates public network traffic collection.
        :param pulumi.Input[str] traffic_type: The type of traffic collected. Valid values:**All**: All traffic.**Allow**: Access control allowedtraffic.**Drop**: Access control denied traffic.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FlowLogState.__new__(_FlowLogState)

        __props__.__dict__["aggregation_interval"] = aggregation_interval
        __props__.__dict__["business_status"] = business_status
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["flow_log_id"] = flow_log_id
        __props__.__dict__["flow_log_name"] = flow_log_name
        __props__.__dict__["log_store_name"] = log_store_name
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["traffic_paths"] = traffic_paths
        __props__.__dict__["traffic_type"] = traffic_type
        return FlowLog(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aggregationInterval")
    def aggregation_interval(self) -> pulumi.Output[str]:
        """
        Data aggregation interval.
        """
        return pulumi.get(self, "aggregation_interval")

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> pulumi.Output[str]:
        """
        Business status.
        """
        return pulumi.get(self, "business_status")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The Description of the VPC Flow Log.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="flowLogId")
    def flow_log_id(self) -> pulumi.Output[str]:
        """
        The flow log ID.
        """
        return pulumi.get(self, "flow_log_id")

    @property
    @pulumi.getter(name="flowLogName")
    def flow_log_name(self) -> pulumi.Output[Optional[str]]:
        """
        The Name of the VPC Flow Log.
        """
        return pulumi.get(self, "flow_log_name")

    @property
    @pulumi.getter(name="logStoreName")
    def log_store_name(self) -> pulumi.Output[str]:
        """
        The name of the logstore.
        """
        return pulumi.get(self, "log_store_name")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The name of the project.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        """
        The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the VPC Flow Log. Valid values: **Active** and **Inactive**.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The tag of the current instance resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficPaths")
    def traffic_paths(self) -> pulumi.Output[Sequence[str]]:
        """
        The collected flow path. Value:**all**: indicates full acquisition.**internetGateway**: indicates public network traffic collection.
        """
        return pulumi.get(self, "traffic_paths")

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> pulumi.Output[str]:
        """
        The type of traffic collected. Valid values:**All**: All traffic.**Allow**: Access control allowedtraffic.**Drop**: Access control denied traffic.
        """
        return pulumi.get(self, "traffic_type")

