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
    'GetTrafficMarkingPoliciesResult',
    'AwaitableGetTrafficMarkingPoliciesResult',
    'get_traffic_marking_policies',
    'get_traffic_marking_policies_output',
]

@pulumi.output_type
class GetTrafficMarkingPoliciesResult:
    """
    A collection of values returned by getTrafficMarkingPolicies.
    """
    def __init__(__self__, description=None, id=None, ids=None, name_regex=None, names=None, output_file=None, policies=None, status=None, transit_router_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
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
        if transit_router_id and not isinstance(transit_router_id, str):
            raise TypeError("Expected argument 'transit_router_id' to be a str")
        pulumi.set(__self__, "transit_router_id", transit_router_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

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
    def policies(self) -> Sequence['outputs.GetTrafficMarkingPoliciesPolicyResult']:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> str:
        return pulumi.get(self, "transit_router_id")


class AwaitableGetTrafficMarkingPoliciesResult(GetTrafficMarkingPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrafficMarkingPoliciesResult(
            description=self.description,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            policies=self.policies,
            status=self.status,
            transit_router_id=self.transit_router_id)


def get_traffic_marking_policies(description: Optional[str] = None,
                                 ids: Optional[Sequence[str]] = None,
                                 name_regex: Optional[str] = None,
                                 output_file: Optional[str] = None,
                                 status: Optional[str] = None,
                                 transit_router_id: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrafficMarkingPoliciesResult:
    """
    This data source provides the Cen Traffic Marking Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.173.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cen.get_traffic_marking_policies(transit_router_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("cenTrafficMarkingPolicyId1", ids.policies[0].id)
    name_regex = alicloud.cen.get_traffic_marking_policies(transit_router_id="example_value",
        name_regex="^my-TrafficMarkingPolicy")
    pulumi.export("cenTrafficMarkingPolicyId2", name_regex.policies[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str description: The description of the Traffic Marking Policy.
    :param Sequence[str] ids: A list of Traffic Marking Policy IDs.
    :param str name_regex: A regex string to filter results by Traffic Marking Policy name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource.
    :param str transit_router_id: The ID of the transit router.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['transitRouterId'] = transit_router_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getTrafficMarkingPolicies:getTrafficMarkingPolicies', __args__, opts=opts, typ=GetTrafficMarkingPoliciesResult).value

    return AwaitableGetTrafficMarkingPoliciesResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        policies=pulumi.get(__ret__, 'policies'),
        status=pulumi.get(__ret__, 'status'),
        transit_router_id=pulumi.get(__ret__, 'transit_router_id'))


@_utilities.lift_output_func(get_traffic_marking_policies)
def get_traffic_marking_policies_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                                        ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                        name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                        status: Optional[pulumi.Input[Optional[str]]] = None,
                                        transit_router_id: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTrafficMarkingPoliciesResult]:
    """
    This data source provides the Cen Traffic Marking Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.173.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cen.get_traffic_marking_policies(transit_router_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("cenTrafficMarkingPolicyId1", ids.policies[0].id)
    name_regex = alicloud.cen.get_traffic_marking_policies(transit_router_id="example_value",
        name_regex="^my-TrafficMarkingPolicy")
    pulumi.export("cenTrafficMarkingPolicyId2", name_regex.policies[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str description: The description of the Traffic Marking Policy.
    :param Sequence[str] ids: A list of Traffic Marking Policy IDs.
    :param str name_regex: A regex string to filter results by Traffic Marking Policy name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource.
    :param str transit_router_id: The ID of the transit router.
    """
    ...
