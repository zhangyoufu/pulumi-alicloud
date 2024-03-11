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
    'GetInterRegionTrafficQosPoliciesResult',
    'AwaitableGetInterRegionTrafficQosPoliciesResult',
    'get_inter_region_traffic_qos_policies',
    'get_inter_region_traffic_qos_policies_output',
]

@pulumi.output_type
class GetInterRegionTrafficQosPoliciesResult:
    """
    A collection of values returned by getInterRegionTrafficQosPolicies.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, names=None, output_file=None, policies=None, status=None, traffic_qos_policy_description=None, traffic_qos_policy_id=None, traffic_qos_policy_name=None, transit_router_attachment_id=None, transit_router_id=None):
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
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if traffic_qos_policy_description and not isinstance(traffic_qos_policy_description, str):
            raise TypeError("Expected argument 'traffic_qos_policy_description' to be a str")
        pulumi.set(__self__, "traffic_qos_policy_description", traffic_qos_policy_description)
        if traffic_qos_policy_id and not isinstance(traffic_qos_policy_id, str):
            raise TypeError("Expected argument 'traffic_qos_policy_id' to be a str")
        pulumi.set(__self__, "traffic_qos_policy_id", traffic_qos_policy_id)
        if traffic_qos_policy_name and not isinstance(traffic_qos_policy_name, str):
            raise TypeError("Expected argument 'traffic_qos_policy_name' to be a str")
        pulumi.set(__self__, "traffic_qos_policy_name", traffic_qos_policy_name)
        if transit_router_attachment_id and not isinstance(transit_router_attachment_id, str):
            raise TypeError("Expected argument 'transit_router_attachment_id' to be a str")
        pulumi.set(__self__, "transit_router_attachment_id", transit_router_attachment_id)
        if transit_router_id and not isinstance(transit_router_id, str):
            raise TypeError("Expected argument 'transit_router_id' to be a str")
        pulumi.set(__self__, "transit_router_id", transit_router_id)

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
        """
        A list of Inter Region Traffic Qos Policy names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def policies(self) -> Sequence['outputs.GetInterRegionTrafficQosPoliciesPolicyResult']:
        """
        A list of Cen Inter Region Traffic Qos Policies. Each element contains the following attributes:
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the Inter Region Traffic Qos Policy.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="trafficQosPolicyDescription")
    def traffic_qos_policy_description(self) -> Optional[str]:
        return pulumi.get(self, "traffic_qos_policy_description")

    @property
    @pulumi.getter(name="trafficQosPolicyId")
    def traffic_qos_policy_id(self) -> Optional[str]:
        return pulumi.get(self, "traffic_qos_policy_id")

    @property
    @pulumi.getter(name="trafficQosPolicyName")
    def traffic_qos_policy_name(self) -> Optional[str]:
        return pulumi.get(self, "traffic_qos_policy_name")

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> str:
        """
        The ID of the inter-region connection.
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> str:
        """
        The ID of the transit router.
        """
        return pulumi.get(self, "transit_router_id")


class AwaitableGetInterRegionTrafficQosPoliciesResult(GetInterRegionTrafficQosPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInterRegionTrafficQosPoliciesResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            policies=self.policies,
            status=self.status,
            traffic_qos_policy_description=self.traffic_qos_policy_description,
            traffic_qos_policy_id=self.traffic_qos_policy_id,
            traffic_qos_policy_name=self.traffic_qos_policy_name,
            transit_router_attachment_id=self.transit_router_attachment_id,
            transit_router_id=self.transit_router_id)


def get_inter_region_traffic_qos_policies(ids: Optional[Sequence[str]] = None,
                                          name_regex: Optional[str] = None,
                                          output_file: Optional[str] = None,
                                          status: Optional[str] = None,
                                          traffic_qos_policy_description: Optional[str] = None,
                                          traffic_qos_policy_id: Optional[str] = None,
                                          traffic_qos_policy_name: Optional[str] = None,
                                          transit_router_attachment_id: Optional[str] = None,
                                          transit_router_id: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInterRegionTrafficQosPoliciesResult:
    """
    This data source provides the Cen Inter Region Traffic Qos Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.195.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cen.get_inter_region_traffic_qos_policies(ids=["example_id"],
        transit_router_id="your_transit_router_id",
        transit_router_attachment_id="your_transit_router_attachment_id")
    pulumi.export("cenInterRegionTrafficQosPolicyId0", ids.policies[0].id)
    name_regex = alicloud.cen.get_inter_region_traffic_qos_policies(name_regex="^my-name",
        transit_router_id="your_transit_router_id",
        transit_router_attachment_id="your_transit_router_attachment_id")
    pulumi.export("cenInterRegionTrafficQosPolicyId1", name_regex.policies[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Inter Region Traffic Qos Policy IDs.
    :param str name_regex: A regex string to filter results by Inter Region Traffic Qos Policy name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the traffic scheduling policy. Valid Value: `Creating`, `Active`, `Modifying`, `Deleting`, `Deleted`.
    :param str traffic_qos_policy_description: The description of the QoS policy.
    :param str traffic_qos_policy_id: The ID of the QoS policy.
    :param str traffic_qos_policy_name: The name of the QoS policy.
    :param str transit_router_attachment_id: The ID of the inter-region connection.
    :param str transit_router_id: The ID of the transit router.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['trafficQosPolicyDescription'] = traffic_qos_policy_description
    __args__['trafficQosPolicyId'] = traffic_qos_policy_id
    __args__['trafficQosPolicyName'] = traffic_qos_policy_name
    __args__['transitRouterAttachmentId'] = transit_router_attachment_id
    __args__['transitRouterId'] = transit_router_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getInterRegionTrafficQosPolicies:getInterRegionTrafficQosPolicies', __args__, opts=opts, typ=GetInterRegionTrafficQosPoliciesResult).value

    return AwaitableGetInterRegionTrafficQosPoliciesResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        policies=pulumi.get(__ret__, 'policies'),
        status=pulumi.get(__ret__, 'status'),
        traffic_qos_policy_description=pulumi.get(__ret__, 'traffic_qos_policy_description'),
        traffic_qos_policy_id=pulumi.get(__ret__, 'traffic_qos_policy_id'),
        traffic_qos_policy_name=pulumi.get(__ret__, 'traffic_qos_policy_name'),
        transit_router_attachment_id=pulumi.get(__ret__, 'transit_router_attachment_id'),
        transit_router_id=pulumi.get(__ret__, 'transit_router_id'))


@_utilities.lift_output_func(get_inter_region_traffic_qos_policies)
def get_inter_region_traffic_qos_policies_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                                 name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                                 status: Optional[pulumi.Input[Optional[str]]] = None,
                                                 traffic_qos_policy_description: Optional[pulumi.Input[Optional[str]]] = None,
                                                 traffic_qos_policy_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                 traffic_qos_policy_name: Optional[pulumi.Input[Optional[str]]] = None,
                                                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
                                                 transit_router_id: Optional[pulumi.Input[str]] = None,
                                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInterRegionTrafficQosPoliciesResult]:
    """
    This data source provides the Cen Inter Region Traffic Qos Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.195.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cen.get_inter_region_traffic_qos_policies(ids=["example_id"],
        transit_router_id="your_transit_router_id",
        transit_router_attachment_id="your_transit_router_attachment_id")
    pulumi.export("cenInterRegionTrafficQosPolicyId0", ids.policies[0].id)
    name_regex = alicloud.cen.get_inter_region_traffic_qos_policies(name_regex="^my-name",
        transit_router_id="your_transit_router_id",
        transit_router_attachment_id="your_transit_router_attachment_id")
    pulumi.export("cenInterRegionTrafficQosPolicyId1", name_regex.policies[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Inter Region Traffic Qos Policy IDs.
    :param str name_regex: A regex string to filter results by Inter Region Traffic Qos Policy name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the traffic scheduling policy. Valid Value: `Creating`, `Active`, `Modifying`, `Deleting`, `Deleted`.
    :param str traffic_qos_policy_description: The description of the QoS policy.
    :param str traffic_qos_policy_id: The ID of the QoS policy.
    :param str traffic_qos_policy_name: The name of the QoS policy.
    :param str transit_router_attachment_id: The ID of the inter-region connection.
    :param str transit_router_id: The ID of the transit router.
    """
    ...
