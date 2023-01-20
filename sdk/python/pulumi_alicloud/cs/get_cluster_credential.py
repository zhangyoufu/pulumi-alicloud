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
    'GetClusterCredentialResult',
    'AwaitableGetClusterCredentialResult',
    'get_cluster_credential',
    'get_cluster_credential_output',
]

@pulumi.output_type
class GetClusterCredentialResult:
    """
    A collection of values returned by getClusterCredential.
    """
    def __init__(__self__, certificate_authority=None, cluster_id=None, cluster_name=None, expiration=None, id=None, kube_config=None, output_file=None, temporary_duration_minutes=None):
        if certificate_authority and not isinstance(certificate_authority, dict):
            raise TypeError("Expected argument 'certificate_authority' to be a dict")
        pulumi.set(__self__, "certificate_authority", certificate_authority)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_name and not isinstance(cluster_name, str):
            raise TypeError("Expected argument 'cluster_name' to be a str")
        pulumi.set(__self__, "cluster_name", cluster_name)
        if expiration and not isinstance(expiration, str):
            raise TypeError("Expected argument 'expiration' to be a str")
        pulumi.set(__self__, "expiration", expiration)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kube_config and not isinstance(kube_config, str):
            raise TypeError("Expected argument 'kube_config' to be a str")
        pulumi.set(__self__, "kube_config", kube_config)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if temporary_duration_minutes and not isinstance(temporary_duration_minutes, int):
            raise TypeError("Expected argument 'temporary_duration_minutes' to be a int")
        pulumi.set(__self__, "temporary_duration_minutes", temporary_duration_minutes)

    @property
    @pulumi.getter(name="certificateAuthority")
    def certificate_authority(self) -> 'outputs.GetClusterCredentialCertificateAuthorityResult':
        """
        (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
        """
        return pulumi.get(self, "certificate_authority")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The id of target cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        """
        The name of target cluster.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter
    def expiration(self) -> str:
        """
        Expiration time of kube config. Format: UTC time in rfc3339.
        """
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kubeConfig")
    def kube_config(self) -> str:
        """
        (Sensitive) The kube config to use to authenticate with the cluster.
        """
        return pulumi.get(self, "kube_config")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="temporaryDurationMinutes")
    def temporary_duration_minutes(self) -> Optional[int]:
        return pulumi.get(self, "temporary_duration_minutes")


class AwaitableGetClusterCredentialResult(GetClusterCredentialResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterCredentialResult(
            certificate_authority=self.certificate_authority,
            cluster_id=self.cluster_id,
            cluster_name=self.cluster_name,
            expiration=self.expiration,
            id=self.id,
            kube_config=self.kube_config,
            output_file=self.output_file,
            temporary_duration_minutes=self.temporary_duration_minutes)


def get_cluster_credential(cluster_id: Optional[str] = None,
                           output_file: Optional[str] = None,
                           temporary_duration_minutes: Optional[int] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterCredentialResult:
    """
    This data source provides Container Service cluster credential on Alibaba Cloud.

    > **NOTE:** Available in v1.187.0+

    > **NOTE:** This datasource can be used on all kinds of ACK clusters, including managed clusters, imported kubernetes clusters, serverless clusters and edge clusters. Please make sure that the target cluster is not in the failed state before using this datasource, since the api server of clusters in the failed state cannot be accessed.


    :param str cluster_id: The id of target cluster.
    :param int temporary_duration_minutes: Automatic expiration time of the returned credential. The valid value between `15` and `4320`, in minutes. When this field is omitted, the expiration time will be determined by the system automatically and the result will be in the attributed field `expiration`.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['outputFile'] = output_file
    __args__['temporaryDurationMinutes'] = temporary_duration_minutes
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cs/getClusterCredential:getClusterCredential', __args__, opts=opts, typ=GetClusterCredentialResult).value

    return AwaitableGetClusterCredentialResult(
        certificate_authority=__ret__.certificate_authority,
        cluster_id=__ret__.cluster_id,
        cluster_name=__ret__.cluster_name,
        expiration=__ret__.expiration,
        id=__ret__.id,
        kube_config=__ret__.kube_config,
        output_file=__ret__.output_file,
        temporary_duration_minutes=__ret__.temporary_duration_minutes)


@_utilities.lift_output_func(get_cluster_credential)
def get_cluster_credential_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  temporary_duration_minutes: Optional[pulumi.Input[Optional[int]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterCredentialResult]:
    """
    This data source provides Container Service cluster credential on Alibaba Cloud.

    > **NOTE:** Available in v1.187.0+

    > **NOTE:** This datasource can be used on all kinds of ACK clusters, including managed clusters, imported kubernetes clusters, serverless clusters and edge clusters. Please make sure that the target cluster is not in the failed state before using this datasource, since the api server of clusters in the failed state cannot be accessed.


    :param str cluster_id: The id of target cluster.
    :param int temporary_duration_minutes: Automatic expiration time of the returned credential. The valid value between `15` and `4320`, in minutes. When this field is omitted, the expiration time will be determined by the system automatically and the result will be in the attributed field `expiration`.
    """
    ...
