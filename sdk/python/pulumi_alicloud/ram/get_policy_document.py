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
from ._inputs import *

__all__ = [
    'GetPolicyDocumentResult',
    'AwaitableGetPolicyDocumentResult',
    'get_policy_document',
    'get_policy_document_output',
]

@pulumi.output_type
class GetPolicyDocumentResult:
    """
    A collection of values returned by getPolicyDocument.
    """
    def __init__(__self__, document=None, id=None, output_file=None, statements=None, version=None):
        if document and not isinstance(document, str):
            raise TypeError("Expected argument 'document' to be a str")
        pulumi.set(__self__, "document", document)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if statements and not isinstance(statements, list):
            raise TypeError("Expected argument 'statements' to be a list")
        pulumi.set(__self__, "statements", statements)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def document(self) -> str:
        """
        Standard policy document rendered based on the arguments above.
        """
        return pulumi.get(self, "document")

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
    def statements(self) -> Optional[Sequence['outputs.GetPolicyDocumentStatementResult']]:
        return pulumi.get(self, "statements")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


class AwaitableGetPolicyDocumentResult(GetPolicyDocumentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyDocumentResult(
            document=self.document,
            id=self.id,
            output_file=self.output_file,
            statements=self.statements,
            version=self.version)


def get_policy_document(output_file: Optional[str] = None,
                        statements: Optional[Sequence[Union['GetPolicyDocumentStatementArgs', 'GetPolicyDocumentStatementArgsDict']]] = None,
                        version: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPolicyDocumentResult:
    """
    This data source Generates a RAM policy document of the current Alibaba Cloud user.

    > **NOTE:** Available since v1.184.0+.

    ## Example Usage

    ### Basic Example

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    basic_example = alicloud.ram.get_policy_document(version="1",
        statements=[{
            "effect": "Allow",
            "actions": ["oss:*"],
            "resources": [
                "acs:oss:*:*:myphotos",
                "acs:oss:*:*:myphotos/*",
            ],
        }])
    default = alicloud.ram.Policy("default",
        policy_name="tf-example",
        policy_document=basic_example.document,
        force=True)
    ```

    `data.alicloud_ram_policy_document.basic_example.document` will evaluate to:

    ### Example Multiple Condition Keys and Values

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    multiple_condition = alicloud.ram.get_policy_document(version="1",
        statements=[
            {
                "effect": "Allow",
                "actions": [
                    "oss:ListBuckets",
                    "oss:GetBucketStat",
                    "oss:GetBucketInfo",
                    "oss:GetBucketTagging",
                    "oss:GetBucketAcl",
                ],
                "resources": ["acs:oss:*:*:*"],
            },
            {
                "effect": "Allow",
                "actions": [
                    "oss:GetObject",
                    "oss:GetObjectAcl",
                ],
                "resources": ["acs:oss:*:*:myphotos/hangzhou/2015/*"],
            },
            {
                "effect": "Allow",
                "actions": ["oss:ListObjects"],
                "resources": ["acs:oss:*:*:myphotos"],
                "conditions": [
                    {
                        "operator": "StringLike",
                        "variable": "oss:Delimiter",
                        "values": ["/"],
                    },
                    {
                        "operator": "StringLike",
                        "variable": "oss:Prefix",
                        "values": [
                            "",
                            "hangzhou/",
                            "hangzhou/2015/*",
                        ],
                    },
                ],
            },
        ])
    policy = alicloud.ram.Policy("policy",
        policy_name="tf-example-condition",
        policy_document=multiple_condition.document,
        force=True)
    ```

    `data.alicloud_ram_policy_document.multiple_condition.document` will evaluate to:

    ### Example Assume-Role Policy with RAM Principal

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ram_example = alicloud.ram.get_policy_document(statements=[{
        "effect": "Allow",
        "actions": ["sts:AssumeRole"],
        "principals": [{
            "entity": "RAM",
            "identifiers": ["acs:ram::123456789012****:root"],
        }],
    }])
    role = alicloud.ram.Role("role",
        name="tf-example-role-ram",
        document=ram_example.document,
        force=True)
    ```

    `data.alicloud_ram_policy_document.ram_example.document` will evaluate to:

    ### Example Assume-Role Policy with Service Principal

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    service_example = alicloud.ram.get_policy_document(statements=[{
        "effect": "Allow",
        "actions": ["sts:AssumeRole"],
        "principals": [{
            "entity": "Service",
            "identifiers": ["ecs.aliyuncs.com"],
        }],
    }])
    role = alicloud.ram.Role("role",
        name="tf-example-role-service",
        document=service_example.document,
        force=True)
    ```

    `data.alicloud_ram_policy_document.service_example.document` will evaluate to:

    ### Example Assume-Role Policy with Federated Principal

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    federated_example = alicloud.ram.get_policy_document(statements=[{
        "effect": "Allow",
        "actions": ["sts:AssumeRole"],
        "principals": [{
            "entity": "Federated",
            "identifiers": ["acs:ram::123456789012****:saml-provider/testprovider"],
        }],
        "conditions": [{
            "operator": "StringEquals",
            "variable": "saml:recipient",
            "values": ["https://signin.aliyun.com/saml-role/sso"],
        }],
    }])
    role = alicloud.ram.Role("role",
        name="tf-example-role-federated",
        document=federated_example.document,
        force=True)
    ```

    `data.alicloud_ram_policy_document.federated_example.document` will evaluate to:


    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param Sequence[Union['GetPolicyDocumentStatementArgs', 'GetPolicyDocumentStatementArgsDict']] statements: Statement of the RAM policy document. See the following `Block statement`. See `statement` below.
    :param str version: Version of the RAM policy document. Valid value is `1`. Default value is `1`.
    """
    __args__ = dict()
    __args__['outputFile'] = output_file
    __args__['statements'] = statements
    __args__['version'] = version
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ram/getPolicyDocument:getPolicyDocument', __args__, opts=opts, typ=GetPolicyDocumentResult).value

    return AwaitableGetPolicyDocumentResult(
        document=pulumi.get(__ret__, 'document'),
        id=pulumi.get(__ret__, 'id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        statements=pulumi.get(__ret__, 'statements'),
        version=pulumi.get(__ret__, 'version'))


@_utilities.lift_output_func(get_policy_document)
def get_policy_document_output(output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               statements: Optional[pulumi.Input[Optional[Sequence[Union['GetPolicyDocumentStatementArgs', 'GetPolicyDocumentStatementArgsDict']]]]] = None,
                               version: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPolicyDocumentResult]:
    """
    This data source Generates a RAM policy document of the current Alibaba Cloud user.

    > **NOTE:** Available since v1.184.0+.

    ## Example Usage

    ### Basic Example

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    basic_example = alicloud.ram.get_policy_document(version="1",
        statements=[{
            "effect": "Allow",
            "actions": ["oss:*"],
            "resources": [
                "acs:oss:*:*:myphotos",
                "acs:oss:*:*:myphotos/*",
            ],
        }])
    default = alicloud.ram.Policy("default",
        policy_name="tf-example",
        policy_document=basic_example.document,
        force=True)
    ```

    `data.alicloud_ram_policy_document.basic_example.document` will evaluate to:

    ### Example Multiple Condition Keys and Values

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    multiple_condition = alicloud.ram.get_policy_document(version="1",
        statements=[
            {
                "effect": "Allow",
                "actions": [
                    "oss:ListBuckets",
                    "oss:GetBucketStat",
                    "oss:GetBucketInfo",
                    "oss:GetBucketTagging",
                    "oss:GetBucketAcl",
                ],
                "resources": ["acs:oss:*:*:*"],
            },
            {
                "effect": "Allow",
                "actions": [
                    "oss:GetObject",
                    "oss:GetObjectAcl",
                ],
                "resources": ["acs:oss:*:*:myphotos/hangzhou/2015/*"],
            },
            {
                "effect": "Allow",
                "actions": ["oss:ListObjects"],
                "resources": ["acs:oss:*:*:myphotos"],
                "conditions": [
                    {
                        "operator": "StringLike",
                        "variable": "oss:Delimiter",
                        "values": ["/"],
                    },
                    {
                        "operator": "StringLike",
                        "variable": "oss:Prefix",
                        "values": [
                            "",
                            "hangzhou/",
                            "hangzhou/2015/*",
                        ],
                    },
                ],
            },
        ])
    policy = alicloud.ram.Policy("policy",
        policy_name="tf-example-condition",
        policy_document=multiple_condition.document,
        force=True)
    ```

    `data.alicloud_ram_policy_document.multiple_condition.document` will evaluate to:

    ### Example Assume-Role Policy with RAM Principal

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ram_example = alicloud.ram.get_policy_document(statements=[{
        "effect": "Allow",
        "actions": ["sts:AssumeRole"],
        "principals": [{
            "entity": "RAM",
            "identifiers": ["acs:ram::123456789012****:root"],
        }],
    }])
    role = alicloud.ram.Role("role",
        name="tf-example-role-ram",
        document=ram_example.document,
        force=True)
    ```

    `data.alicloud_ram_policy_document.ram_example.document` will evaluate to:

    ### Example Assume-Role Policy with Service Principal

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    service_example = alicloud.ram.get_policy_document(statements=[{
        "effect": "Allow",
        "actions": ["sts:AssumeRole"],
        "principals": [{
            "entity": "Service",
            "identifiers": ["ecs.aliyuncs.com"],
        }],
    }])
    role = alicloud.ram.Role("role",
        name="tf-example-role-service",
        document=service_example.document,
        force=True)
    ```

    `data.alicloud_ram_policy_document.service_example.document` will evaluate to:

    ### Example Assume-Role Policy with Federated Principal

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    federated_example = alicloud.ram.get_policy_document(statements=[{
        "effect": "Allow",
        "actions": ["sts:AssumeRole"],
        "principals": [{
            "entity": "Federated",
            "identifiers": ["acs:ram::123456789012****:saml-provider/testprovider"],
        }],
        "conditions": [{
            "operator": "StringEquals",
            "variable": "saml:recipient",
            "values": ["https://signin.aliyun.com/saml-role/sso"],
        }],
    }])
    role = alicloud.ram.Role("role",
        name="tf-example-role-federated",
        document=federated_example.document,
        force=True)
    ```

    `data.alicloud_ram_policy_document.federated_example.document` will evaluate to:


    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param Sequence[Union['GetPolicyDocumentStatementArgs', 'GetPolicyDocumentStatementArgsDict']] statements: Statement of the RAM policy document. See the following `Block statement`. See `statement` below.
    :param str version: Version of the RAM policy document. Valid value is `1`. Default value is `1`.
    """
    ...
