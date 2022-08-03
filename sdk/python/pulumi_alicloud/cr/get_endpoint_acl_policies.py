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
    'GetEndpointAclPoliciesResult',
    'AwaitableGetEndpointAclPoliciesResult',
    'get_endpoint_acl_policies',
    'get_endpoint_acl_policies_output',
]

@pulumi.output_type
class GetEndpointAclPoliciesResult:
    """
    A collection of values returned by getEndpointAclPolicies.
    """
    def __init__(__self__, endpoint_type=None, id=None, ids=None, instance_id=None, output_file=None, policies=None):
        if endpoint_type and not isinstance(endpoint_type, str):
            raise TypeError("Expected argument 'endpoint_type' to be a str")
        pulumi.set(__self__, "endpoint_type", endpoint_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> str:
        return pulumi.get(self, "endpoint_type")

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
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def policies(self) -> Sequence['outputs.GetEndpointAclPoliciesPolicyResult']:
        return pulumi.get(self, "policies")


class AwaitableGetEndpointAclPoliciesResult(GetEndpointAclPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointAclPoliciesResult(
            endpoint_type=self.endpoint_type,
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            output_file=self.output_file,
            policies=self.policies)


def get_endpoint_acl_policies(endpoint_type: Optional[str] = None,
                              ids: Optional[Sequence[str]] = None,
                              instance_id: Optional[str] = None,
                              output_file: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEndpointAclPoliciesResult:
    """
    This data source provides the Cr Endpoint Acl Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.139.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cr.get_endpoint_acl_policies(instance_id="example_value",
        endpoint_type="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("crEndpointAclPolicyId1", ids.policies[0].id)
    ```


    :param str endpoint_type: The type of endpoint.
    :param Sequence[str] ids: A list of Endpoint Acl Policy IDs.
    :param str instance_id: The ID of the CR Instance.
    """
    __args__ = dict()
    __args__['endpointType'] = endpoint_type
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cr/getEndpointAclPolicies:getEndpointAclPolicies', __args__, opts=opts, typ=GetEndpointAclPoliciesResult).value

    return AwaitableGetEndpointAclPoliciesResult(
        endpoint_type=__ret__.endpoint_type,
        id=__ret__.id,
        ids=__ret__.ids,
        instance_id=__ret__.instance_id,
        output_file=__ret__.output_file,
        policies=__ret__.policies)


@_utilities.lift_output_func(get_endpoint_acl_policies)
def get_endpoint_acl_policies_output(endpoint_type: Optional[pulumi.Input[str]] = None,
                                     ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                     instance_id: Optional[pulumi.Input[str]] = None,
                                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEndpointAclPoliciesResult]:
    """
    This data source provides the Cr Endpoint Acl Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.139.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cr.get_endpoint_acl_policies(instance_id="example_value",
        endpoint_type="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("crEndpointAclPolicyId1", ids.policies[0].id)
    ```


    :param str endpoint_type: The type of endpoint.
    :param Sequence[str] ids: A list of Endpoint Acl Policy IDs.
    :param str instance_id: The ID of the CR Instance.
    """
    ...
