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
    'GetAuditPoliciesResult',
    'AwaitableGetAuditPoliciesResult',
    'get_audit_policies',
    'get_audit_policies_output',
]

@pulumi.output_type
class GetAuditPoliciesResult:
    """
    A collection of values returned by getAuditPolicies.
    """
    def __init__(__self__, db_instance_id=None, id=None, output_file=None, policies=None):
        if db_instance_id and not isinstance(db_instance_id, str):
            raise TypeError("Expected argument 'db_instance_id' to be a str")
        pulumi.set(__self__, "db_instance_id", db_instance_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> str:
        return pulumi.get(self, "db_instance_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def policies(self) -> Sequence['outputs.GetAuditPoliciesPolicyResult']:
        return pulumi.get(self, "policies")


class AwaitableGetAuditPoliciesResult(GetAuditPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuditPoliciesResult(
            db_instance_id=self.db_instance_id,
            id=self.id,
            output_file=self.output_file,
            policies=self.policies)


def get_audit_policies(db_instance_id: Optional[str] = None,
                       output_file: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuditPoliciesResult:
    """
    This data source provides the Mongodb Audit Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.148.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.mongodb.get_audit_policies(db_instance_id="example_value")
    pulumi.export("mongodbAuditPolicyId1", example.policies[0].id)
    ```


    :param str db_instance_id: The ID of the instance.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['dbInstanceId'] = db_instance_id
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:mongodb/getAuditPolicies:getAuditPolicies', __args__, opts=opts, typ=GetAuditPoliciesResult).value

    return AwaitableGetAuditPoliciesResult(
        db_instance_id=__ret__.db_instance_id,
        id=__ret__.id,
        output_file=__ret__.output_file,
        policies=__ret__.policies)


@_utilities.lift_output_func(get_audit_policies)
def get_audit_policies_output(db_instance_id: Optional[pulumi.Input[str]] = None,
                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAuditPoliciesResult]:
    """
    This data source provides the Mongodb Audit Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.148.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.mongodb.get_audit_policies(db_instance_id="example_value")
    pulumi.export("mongodbAuditPolicyId1", example.policies[0].id)
    ```


    :param str db_instance_id: The ID of the instance.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
