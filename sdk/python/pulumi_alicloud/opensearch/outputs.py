# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'AppGroupOrder',
    'AppGroupQuota',
    'GetAppGroupsGroupResult',
    'GetAppGroupsGroupQuotaResult',
]

@pulumi.output_type
class AppGroupOrder(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "autoRenew":
            suggest = "auto_renew"
        elif key == "pricingCycle":
            suggest = "pricing_cycle"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AppGroupOrder. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AppGroupOrder.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AppGroupOrder.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 auto_renew: Optional[bool] = None,
                 duration: Optional[int] = None,
                 pricing_cycle: Optional[str] = None):
        """
        :param bool auto_renew: Whether to renew automatically. It only takes effect when the parameter payment_type takes the value `Subscription`.
        :param int duration: Order cycle. The minimum value is not less than 0.
        :param str pricing_cycle: Order cycle unit. Valid values: `Year` and `Month`.
        """
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if pricing_cycle is not None:
            pulumi.set(__self__, "pricing_cycle", pricing_cycle)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[bool]:
        """
        Whether to renew automatically. It only takes effect when the parameter payment_type takes the value `Subscription`.
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter
    def duration(self) -> Optional[int]:
        """
        Order cycle. The minimum value is not less than 0.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="pricingCycle")
    def pricing_cycle(self) -> Optional[str]:
        """
        Order cycle unit. Valid values: `Year` and `Month`.
        """
        return pulumi.get(self, "pricing_cycle")


@pulumi.output_type
class AppGroupQuota(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "computeResource":
            suggest = "compute_resource"
        elif key == "docSize":
            suggest = "doc_size"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AppGroupQuota. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AppGroupQuota.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AppGroupQuota.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 compute_resource: int,
                 doc_size: int,
                 spec: str,
                 qps: Optional[int] = None):
        """
        :param int compute_resource: Computing resources. Unit: LCU.
        :param int doc_size: Storage Size. Unit: GB.
        :param str spec: Specification. Valid values: 
               * `opensearch.share.junior`: Entry-level.
               * `opensearch.share.common`: Shared universal.
               * `opensearch.share.compute`: Shared computing.
               * `opensearch.share.storage`: Shared storage type.
               * `opensearch.private.common`: Exclusive universal type.
               * `opensearch.private.compute`: Exclusive computing type.
               * `opensearch.private.storage`: Exclusive storage type
        :param int qps: Search request. Unit: times/second.
        """
        pulumi.set(__self__, "compute_resource", compute_resource)
        pulumi.set(__self__, "doc_size", doc_size)
        pulumi.set(__self__, "spec", spec)
        if qps is not None:
            pulumi.set(__self__, "qps", qps)

    @property
    @pulumi.getter(name="computeResource")
    def compute_resource(self) -> int:
        """
        Computing resources. Unit: LCU.
        """
        return pulumi.get(self, "compute_resource")

    @property
    @pulumi.getter(name="docSize")
    def doc_size(self) -> int:
        """
        Storage Size. Unit: GB.
        """
        return pulumi.get(self, "doc_size")

    @property
    @pulumi.getter
    def spec(self) -> str:
        """
        Specification. Valid values: 
        * `opensearch.share.junior`: Entry-level.
        * `opensearch.share.common`: Shared universal.
        * `opensearch.share.compute`: Shared computing.
        * `opensearch.share.storage`: Shared storage type.
        * `opensearch.private.common`: Exclusive universal type.
        * `opensearch.private.compute`: Exclusive computing type.
        * `opensearch.private.storage`: Exclusive storage type
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def qps(self) -> Optional[int]:
        """
        Search request. Unit: times/second.
        """
        return pulumi.get(self, "qps")


@pulumi.output_type
class GetAppGroupsGroupResult(dict):
    def __init__(__self__, *,
                 app_group_id: str,
                 app_group_name: str,
                 charge_way: int,
                 commodity_code: str,
                 create_time: int,
                 current_version: str,
                 description: str,
                 domain: str,
                 expire_on: str,
                 first_rank_algo_deployment_id: int,
                 has_pending_quota_review_task: int,
                 id: str,
                 instance_id: str,
                 lock_mode: str,
                 locked_by_expiration: int,
                 payment_type: str,
                 pending_second_rank_algo_deployment_id: int,
                 processing_order_id: str,
                 produced: int,
                 project_id: str,
                 quotas: Sequence['outputs.GetAppGroupsGroupQuotaResult'],
                 resource_group_id: str,
                 second_rank_algo_deployment_id: int,
                 status: str,
                 switched_time: int,
                 type: str):
        """
        :param str app_group_id: The ID of the App Group.
        :param str app_group_name: Application Group Name.
        :param int charge_way: Billing model. Valid values:`compute_resource` and `qps`.
        :param str commodity_code: The commodity code.
        :param int create_time: The time of creation.
        :param str current_version: The version of Application Group Name.
        :param str description: The description of the resource.
        :param str domain: Domain name.
        :param str expire_on: Expiration Time.
        :param int first_rank_algo_deployment_id: Coarse deployment ID.
        :param int has_pending_quota_review_task: Whether the quota status is under approval. Valid status:
               * `0`: normal
               * `1`: Approving.
        :param str instance_id: The Instance ID.
        :param str lock_mode: Locked state. Valid status: `Unlock`,`LockByExpiration`,`ManualLock`.
        :param int locked_by_expiration: Instance is automatically locked after expiration.
        :param str payment_type: The billing method of the resource. Valid values: `Subscription` and `PayAsYouGo`.
        :param int pending_second_rank_algo_deployment_id: Refine deployment ID in deployment.
        :param str processing_order_id: Unfinished order number.
        :param int produced: Whether the production is completed. Valid values:
               * `0`: producing.
               * `1`: completed.
        :param str project_id: The Project ID.
        :param Sequence['GetAppGroupsGroupQuotaArgs'] quotas: Quota information.
        :param str resource_group_id: The Resource Group ID.
        :param int second_rank_algo_deployment_id: Refine deployment ID.
        :param str status: The status of the resource. Valid values: `producing`,`review_pending`,`config_pending`,`normal`,`frozen`.
        :param int switched_time: The Switched time.
        :param str type: Application type. Valid Values: `standard`, `enhanced`.
        """
        pulumi.set(__self__, "app_group_id", app_group_id)
        pulumi.set(__self__, "app_group_name", app_group_name)
        pulumi.set(__self__, "charge_way", charge_way)
        pulumi.set(__self__, "commodity_code", commodity_code)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "current_version", current_version)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "expire_on", expire_on)
        pulumi.set(__self__, "first_rank_algo_deployment_id", first_rank_algo_deployment_id)
        pulumi.set(__self__, "has_pending_quota_review_task", has_pending_quota_review_task)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "lock_mode", lock_mode)
        pulumi.set(__self__, "locked_by_expiration", locked_by_expiration)
        pulumi.set(__self__, "payment_type", payment_type)
        pulumi.set(__self__, "pending_second_rank_algo_deployment_id", pending_second_rank_algo_deployment_id)
        pulumi.set(__self__, "processing_order_id", processing_order_id)
        pulumi.set(__self__, "produced", produced)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "quotas", quotas)
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        pulumi.set(__self__, "second_rank_algo_deployment_id", second_rank_algo_deployment_id)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "switched_time", switched_time)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="appGroupId")
    def app_group_id(self) -> str:
        """
        The ID of the App Group.
        """
        return pulumi.get(self, "app_group_id")

    @property
    @pulumi.getter(name="appGroupName")
    def app_group_name(self) -> str:
        """
        Application Group Name.
        """
        return pulumi.get(self, "app_group_name")

    @property
    @pulumi.getter(name="chargeWay")
    def charge_way(self) -> int:
        """
        Billing model. Valid values:`compute_resource` and `qps`.
        """
        return pulumi.get(self, "charge_way")

    @property
    @pulumi.getter(name="commodityCode")
    def commodity_code(self) -> str:
        """
        The commodity code.
        """
        return pulumi.get(self, "commodity_code")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> int:
        """
        The time of creation.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="currentVersion")
    def current_version(self) -> str:
        """
        The version of Application Group Name.
        """
        return pulumi.get(self, "current_version")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        Domain name.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="expireOn")
    def expire_on(self) -> str:
        """
        Expiration Time.
        """
        return pulumi.get(self, "expire_on")

    @property
    @pulumi.getter(name="firstRankAlgoDeploymentId")
    def first_rank_algo_deployment_id(self) -> int:
        """
        Coarse deployment ID.
        """
        return pulumi.get(self, "first_rank_algo_deployment_id")

    @property
    @pulumi.getter(name="hasPendingQuotaReviewTask")
    def has_pending_quota_review_task(self) -> int:
        """
        Whether the quota status is under approval. Valid status:
        * `0`: normal
        * `1`: Approving.
        """
        return pulumi.get(self, "has_pending_quota_review_task")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="lockMode")
    def lock_mode(self) -> str:
        """
        Locked state. Valid status: `Unlock`,`LockByExpiration`,`ManualLock`.
        """
        return pulumi.get(self, "lock_mode")

    @property
    @pulumi.getter(name="lockedByExpiration")
    def locked_by_expiration(self) -> int:
        """
        Instance is automatically locked after expiration.
        """
        return pulumi.get(self, "locked_by_expiration")

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> str:
        """
        The billing method of the resource. Valid values: `Subscription` and `PayAsYouGo`.
        """
        return pulumi.get(self, "payment_type")

    @property
    @pulumi.getter(name="pendingSecondRankAlgoDeploymentId")
    def pending_second_rank_algo_deployment_id(self) -> int:
        """
        Refine deployment ID in deployment.
        """
        return pulumi.get(self, "pending_second_rank_algo_deployment_id")

    @property
    @pulumi.getter(name="processingOrderId")
    def processing_order_id(self) -> str:
        """
        Unfinished order number.
        """
        return pulumi.get(self, "processing_order_id")

    @property
    @pulumi.getter
    def produced(self) -> int:
        """
        Whether the production is completed. Valid values:
        * `0`: producing.
        * `1`: completed.
        """
        return pulumi.get(self, "produced")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The Project ID.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def quotas(self) -> Sequence['outputs.GetAppGroupsGroupQuotaResult']:
        """
        Quota information.
        """
        return pulumi.get(self, "quotas")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> str:
        """
        The Resource Group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="secondRankAlgoDeploymentId")
    def second_rank_algo_deployment_id(self) -> int:
        """
        Refine deployment ID.
        """
        return pulumi.get(self, "second_rank_algo_deployment_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the resource. Valid values: `producing`,`review_pending`,`config_pending`,`normal`,`frozen`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="switchedTime")
    def switched_time(self) -> int:
        """
        The Switched time.
        """
        return pulumi.get(self, "switched_time")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Application type. Valid Values: `standard`, `enhanced`.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetAppGroupsGroupQuotaResult(dict):
    def __init__(__self__, *,
                 compute_resource: str,
                 doc_size: str,
                 spec: str):
        """
        :param str compute_resource: Computing resources. Unit: LCU.
        :param str doc_size: Storage Size. Unit: GB.
        :param str spec: Specification. Valid values:
               * `opensearch.share.junior`: Entry-level.
               * `opensearch.share.common`: Shared universal.
               * `opensearch.share.compute`: Shared computing.
               * `opensearch.share.storage`: Shared storage type.
               * `opensearch.private.common`: Exclusive universal type.
               * `opensearch.private.compute`: Exclusive computing type.
               * `opensearch.private.storage`: Exclusive storage type
        """
        pulumi.set(__self__, "compute_resource", compute_resource)
        pulumi.set(__self__, "doc_size", doc_size)
        pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter(name="computeResource")
    def compute_resource(self) -> str:
        """
        Computing resources. Unit: LCU.
        """
        return pulumi.get(self, "compute_resource")

    @property
    @pulumi.getter(name="docSize")
    def doc_size(self) -> str:
        """
        Storage Size. Unit: GB.
        """
        return pulumi.get(self, "doc_size")

    @property
    @pulumi.getter
    def spec(self) -> str:
        """
        Specification. Valid values:
        * `opensearch.share.junior`: Entry-level.
        * `opensearch.share.common`: Shared universal.
        * `opensearch.share.compute`: Shared computing.
        * `opensearch.share.storage`: Shared storage type.
        * `opensearch.private.common`: Exclusive universal type.
        * `opensearch.private.compute`: Exclusive computing type.
        * `opensearch.private.storage`: Exclusive storage type
        """
        return pulumi.get(self, "spec")


