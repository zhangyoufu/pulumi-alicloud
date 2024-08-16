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
    'ConsumerGroupConsumeRetryPolicy',
    'RocketMQInstanceNetworkInfo',
    'RocketMQInstanceNetworkInfoEndpoint',
    'RocketMQInstanceNetworkInfoInternetInfo',
    'RocketMQInstanceNetworkInfoVpcInfo',
    'RocketMQInstanceProductInfo',
    'RocketMQInstanceSoftware',
    'GetGroupsGroupResult',
    'GetInstancesInstanceResult',
    'GetTopicsTopicResult',
]

@pulumi.output_type
class ConsumerGroupConsumeRetryPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxRetryTimes":
            suggest = "max_retry_times"
        elif key == "retryPolicy":
            suggest = "retry_policy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConsumerGroupConsumeRetryPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConsumerGroupConsumeRetryPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConsumerGroupConsumeRetryPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_retry_times: Optional[int] = None,
                 retry_policy: Optional[str] = None):
        """
        :param int max_retry_times: Maximum number of retries.
        :param str retry_policy: Consume retry policy.
        """
        if max_retry_times is not None:
            pulumi.set(__self__, "max_retry_times", max_retry_times)
        if retry_policy is not None:
            pulumi.set(__self__, "retry_policy", retry_policy)

    @property
    @pulumi.getter(name="maxRetryTimes")
    def max_retry_times(self) -> Optional[int]:
        """
        Maximum number of retries.
        """
        return pulumi.get(self, "max_retry_times")

    @property
    @pulumi.getter(name="retryPolicy")
    def retry_policy(self) -> Optional[str]:
        """
        Consume retry policy.
        """
        return pulumi.get(self, "retry_policy")


@pulumi.output_type
class RocketMQInstanceNetworkInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "internetInfo":
            suggest = "internet_info"
        elif key == "vpcInfo":
            suggest = "vpc_info"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RocketMQInstanceNetworkInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RocketMQInstanceNetworkInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RocketMQInstanceNetworkInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 internet_info: 'outputs.RocketMQInstanceNetworkInfoInternetInfo',
                 vpc_info: 'outputs.RocketMQInstanceNetworkInfoVpcInfo',
                 endpoints: Optional[Sequence['outputs.RocketMQInstanceNetworkInfoEndpoint']] = None):
        """
        :param 'RocketMQInstanceNetworkInfoInternetInfoArgs' internet_info: instance internet info. See `internet_info` below.
        :param 'RocketMQInstanceNetworkInfoVpcInfoArgs' vpc_info: Proprietary network information. See `vpc_info` below.
        :param Sequence['RocketMQInstanceNetworkInfoEndpointArgs'] endpoints: Access point list.
        """
        pulumi.set(__self__, "internet_info", internet_info)
        pulumi.set(__self__, "vpc_info", vpc_info)
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)

    @property
    @pulumi.getter(name="internetInfo")
    def internet_info(self) -> 'outputs.RocketMQInstanceNetworkInfoInternetInfo':
        """
        instance internet info. See `internet_info` below.
        """
        return pulumi.get(self, "internet_info")

    @property
    @pulumi.getter(name="vpcInfo")
    def vpc_info(self) -> 'outputs.RocketMQInstanceNetworkInfoVpcInfo':
        """
        Proprietary network information. See `vpc_info` below.
        """
        return pulumi.get(self, "vpc_info")

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[Sequence['outputs.RocketMQInstanceNetworkInfoEndpoint']]:
        """
        Access point list.
        """
        return pulumi.get(self, "endpoints")


@pulumi.output_type
class RocketMQInstanceNetworkInfoEndpoint(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "endpointType":
            suggest = "endpoint_type"
        elif key == "endpointUrl":
            suggest = "endpoint_url"
        elif key == "ipWhiteLists":
            suggest = "ip_white_lists"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RocketMQInstanceNetworkInfoEndpoint. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RocketMQInstanceNetworkInfoEndpoint.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RocketMQInstanceNetworkInfoEndpoint.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 endpoint_type: Optional[str] = None,
                 endpoint_url: Optional[str] = None,
                 ip_white_lists: Optional[Sequence[str]] = None):
        """
        :param str endpoint_type: Access point type.
        :param str endpoint_url: Access point address.
        :param Sequence[str] ip_white_lists: White list of access addresses.
        """
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)
        if endpoint_url is not None:
            pulumi.set(__self__, "endpoint_url", endpoint_url)
        if ip_white_lists is not None:
            pulumi.set(__self__, "ip_white_lists", ip_white_lists)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[str]:
        """
        Access point type.
        """
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter(name="endpointUrl")
    def endpoint_url(self) -> Optional[str]:
        """
        Access point address.
        """
        return pulumi.get(self, "endpoint_url")

    @property
    @pulumi.getter(name="ipWhiteLists")
    def ip_white_lists(self) -> Optional[Sequence[str]]:
        """
        White list of access addresses.
        """
        return pulumi.get(self, "ip_white_lists")


@pulumi.output_type
class RocketMQInstanceNetworkInfoInternetInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "flowOutType":
            suggest = "flow_out_type"
        elif key == "internetSpec":
            suggest = "internet_spec"
        elif key == "flowOutBandwidth":
            suggest = "flow_out_bandwidth"
        elif key == "ipWhitelists":
            suggest = "ip_whitelists"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RocketMQInstanceNetworkInfoInternetInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RocketMQInstanceNetworkInfoInternetInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RocketMQInstanceNetworkInfoInternetInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 flow_out_type: str,
                 internet_spec: str,
                 flow_out_bandwidth: Optional[int] = None,
                 ip_whitelists: Optional[Sequence[str]] = None):
        """
        :param str flow_out_type: Public network billing type. The parameter values are as follows:
               - payByBandwidth: Fixed bandwidth billing. Set this value when enabling public network access.
               - uninvolved: Not involved. Set this value when disabling public network access.
        :param str internet_spec: Whether to enable public network access. Instances by default support VPC access. If public network access is enabled, Alibaba Cloud Message Queue RocketMQ version will incur charges for public network outbound bandwidth. For specific billing information, please refer to [Public Network Access Fees](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/internet-access-fee). The parameter values are as follows:
               - enable: Enable public network access
               - disable: Disable public network access
        :param int flow_out_bandwidth: Public network bandwidth specification. Unit: Mb/s.This field should only be filled when the public network billing type is set to payByBandwidth.The value range is [1 - 1000].
        :param Sequence[str] ip_whitelists: internet ip whitelist.
        """
        pulumi.set(__self__, "flow_out_type", flow_out_type)
        pulumi.set(__self__, "internet_spec", internet_spec)
        if flow_out_bandwidth is not None:
            pulumi.set(__self__, "flow_out_bandwidth", flow_out_bandwidth)
        if ip_whitelists is not None:
            pulumi.set(__self__, "ip_whitelists", ip_whitelists)

    @property
    @pulumi.getter(name="flowOutType")
    def flow_out_type(self) -> str:
        """
        Public network billing type. The parameter values are as follows:
        - payByBandwidth: Fixed bandwidth billing. Set this value when enabling public network access.
        - uninvolved: Not involved. Set this value when disabling public network access.
        """
        return pulumi.get(self, "flow_out_type")

    @property
    @pulumi.getter(name="internetSpec")
    def internet_spec(self) -> str:
        """
        Whether to enable public network access. Instances by default support VPC access. If public network access is enabled, Alibaba Cloud Message Queue RocketMQ version will incur charges for public network outbound bandwidth. For specific billing information, please refer to [Public Network Access Fees](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/internet-access-fee). The parameter values are as follows:
        - enable: Enable public network access
        - disable: Disable public network access
        """
        return pulumi.get(self, "internet_spec")

    @property
    @pulumi.getter(name="flowOutBandwidth")
    def flow_out_bandwidth(self) -> Optional[int]:
        """
        Public network bandwidth specification. Unit: Mb/s.This field should only be filled when the public network billing type is set to payByBandwidth.The value range is [1 - 1000].
        """
        return pulumi.get(self, "flow_out_bandwidth")

    @property
    @pulumi.getter(name="ipWhitelists")
    def ip_whitelists(self) -> Optional[Sequence[str]]:
        """
        internet ip whitelist.
        """
        return pulumi.get(self, "ip_whitelists")


@pulumi.output_type
class RocketMQInstanceNetworkInfoVpcInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "vpcId":
            suggest = "vpc_id"
        elif key == "vswitchId":
            suggest = "vswitch_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RocketMQInstanceNetworkInfoVpcInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RocketMQInstanceNetworkInfoVpcInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RocketMQInstanceNetworkInfoVpcInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 vpc_id: str,
                 vswitch_id: str):
        """
        :param str vpc_id: Proprietary Network.
        :param str vswitch_id: VPC network switch.
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        Proprietary Network.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> str:
        """
        VPC network switch.
        """
        return pulumi.get(self, "vswitch_id")


@pulumi.output_type
class RocketMQInstanceProductInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "msgProcessSpec":
            suggest = "msg_process_spec"
        elif key == "autoScaling":
            suggest = "auto_scaling"
        elif key == "messageRetentionTime":
            suggest = "message_retention_time"
        elif key == "sendReceiveRatio":
            suggest = "send_receive_ratio"
        elif key == "supportAutoScaling":
            suggest = "support_auto_scaling"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RocketMQInstanceProductInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RocketMQInstanceProductInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RocketMQInstanceProductInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 msg_process_spec: str,
                 auto_scaling: Optional[bool] = None,
                 message_retention_time: Optional[int] = None,
                 send_receive_ratio: Optional[float] = None,
                 support_auto_scaling: Optional[bool] = None):
        """
        :param str msg_process_spec: Message sending and receiving calculation specifications. For details about the upper limit for sending and receiving messages, see [Instance Specifications](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/instance-specifications).
        :param bool auto_scaling: is open auto scaling.
        :param int message_retention_time: Duration of message retention. Unit: hours.For the range of values, please refer to [Usage Limits](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/usage-limits)>Resource Quotas>Limitations on Message Retention.The message storage in AlibabaCloud RocketMQ is fully implemented in a serverless and elastic manner, with charges based on the actual storage space. You can control the storage capacity of messages by adjusting the duration of message retention. For more information, please see [Storage Fees](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/storage-fees).
        :param float send_receive_ratio: message send receive ratio.Value range: [0.2, 0.5].
        :param bool support_auto_scaling: is support auto scaling.
        """
        pulumi.set(__self__, "msg_process_spec", msg_process_spec)
        if auto_scaling is not None:
            pulumi.set(__self__, "auto_scaling", auto_scaling)
        if message_retention_time is not None:
            pulumi.set(__self__, "message_retention_time", message_retention_time)
        if send_receive_ratio is not None:
            pulumi.set(__self__, "send_receive_ratio", send_receive_ratio)
        if support_auto_scaling is not None:
            pulumi.set(__self__, "support_auto_scaling", support_auto_scaling)

    @property
    @pulumi.getter(name="msgProcessSpec")
    def msg_process_spec(self) -> str:
        """
        Message sending and receiving calculation specifications. For details about the upper limit for sending and receiving messages, see [Instance Specifications](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/instance-specifications).
        """
        return pulumi.get(self, "msg_process_spec")

    @property
    @pulumi.getter(name="autoScaling")
    def auto_scaling(self) -> Optional[bool]:
        """
        is open auto scaling.
        """
        return pulumi.get(self, "auto_scaling")

    @property
    @pulumi.getter(name="messageRetentionTime")
    def message_retention_time(self) -> Optional[int]:
        """
        Duration of message retention. Unit: hours.For the range of values, please refer to [Usage Limits](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/usage-limits)>Resource Quotas>Limitations on Message Retention.The message storage in AlibabaCloud RocketMQ is fully implemented in a serverless and elastic manner, with charges based on the actual storage space. You can control the storage capacity of messages by adjusting the duration of message retention. For more information, please see [Storage Fees](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/storage-fees).
        """
        return pulumi.get(self, "message_retention_time")

    @property
    @pulumi.getter(name="sendReceiveRatio")
    def send_receive_ratio(self) -> Optional[float]:
        """
        message send receive ratio.Value range: [0.2, 0.5].
        """
        return pulumi.get(self, "send_receive_ratio")

    @property
    @pulumi.getter(name="supportAutoScaling")
    def support_auto_scaling(self) -> Optional[bool]:
        """
        is support auto scaling.
        """
        return pulumi.get(self, "support_auto_scaling")


@pulumi.output_type
class RocketMQInstanceSoftware(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maintainTime":
            suggest = "maintain_time"
        elif key == "softwareVersion":
            suggest = "software_version"
        elif key == "upgradeMethod":
            suggest = "upgrade_method"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RocketMQInstanceSoftware. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RocketMQInstanceSoftware.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RocketMQInstanceSoftware.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 maintain_time: Optional[str] = None,
                 software_version: Optional[str] = None,
                 upgrade_method: Optional[str] = None):
        """
        :param str maintain_time: Upgrade time period.
        :param str software_version: Software version.
        :param str upgrade_method: Upgrade method.
        """
        if maintain_time is not None:
            pulumi.set(__self__, "maintain_time", maintain_time)
        if software_version is not None:
            pulumi.set(__self__, "software_version", software_version)
        if upgrade_method is not None:
            pulumi.set(__self__, "upgrade_method", upgrade_method)

    @property
    @pulumi.getter(name="maintainTime")
    def maintain_time(self) -> Optional[str]:
        """
        Upgrade time period.
        """
        return pulumi.get(self, "maintain_time")

    @property
    @pulumi.getter(name="softwareVersion")
    def software_version(self) -> Optional[str]:
        """
        Software version.
        """
        return pulumi.get(self, "software_version")

    @property
    @pulumi.getter(name="upgradeMethod")
    def upgrade_method(self) -> Optional[str]:
        """
        Upgrade method.
        """
        return pulumi.get(self, "upgrade_method")


@pulumi.output_type
class GetGroupsGroupResult(dict):
    def __init__(__self__, *,
                 group_name: str,
                 group_type: str,
                 id: str,
                 independent_naming: bool,
                 instance_id: str,
                 owner: str,
                 remark: str,
                 tags: Mapping[str, str]):
        """
        :param str group_name: The name of the group.
        :param str group_type: Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
        :param str id: The name of the group.
        :param bool independent_naming: Indicates whether namespaces are available. Read [Fields in SubscribeInfoDo](https://www.alibabacloud.com/help/doc-detail/29619.html) for further details.
        :param str instance_id: ID of the ONS Instance that owns the groups.
        :param str owner: The ID of the group owner, which is the Alibaba Cloud UID.
        :param str remark: Remark of the group.
        :param Mapping[str, str] tags: A map of tags assigned to the Ons instance.
        """
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "group_type", group_type)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "independent_naming", independent_naming)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "remark", remark)
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> str:
        """
        The name of the group.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="groupType")
    def group_type(self) -> str:
        """
        Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
        """
        return pulumi.get(self, "group_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The name of the group.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="independentNaming")
    def independent_naming(self) -> bool:
        """
        Indicates whether namespaces are available. Read [Fields in SubscribeInfoDo](https://www.alibabacloud.com/help/doc-detail/29619.html) for further details.
        """
        return pulumi.get(self, "independent_naming")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        ID of the ONS Instance that owns the groups.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The ID of the group owner, which is the Alibaba Cloud UID.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def remark(self) -> str:
        """
        Remark of the group.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags assigned to the Ons instance.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class GetInstancesInstanceResult(dict):
    def __init__(__self__, *,
                 http_internal_endpoint: str,
                 http_internet_endpoint: str,
                 http_internet_secure_endpoint: str,
                 id: str,
                 independent_naming: bool,
                 instance_id: str,
                 instance_name: str,
                 instance_status: int,
                 instance_type: int,
                 release_time: str,
                 remark: str,
                 status: int,
                 tags: Mapping[str, str],
                 tcp_endpoint: str):
        """
        :param str http_internal_endpoint: The internal HTTP endpoint for the Message Queue for Apache RocketMQ instance.
        :param str http_internet_endpoint: The public HTTP endpoint for the Message Queue for Apache RocketMQ instance.
        :param str http_internet_secure_endpoint: The public HTTPS endpoint for the Message Queue for Apache RocketMQ instance.
        :param str id: ID of the instance.
        :param bool independent_naming: Indicates whether any namespace is configured for the Message Queue for Apache RocketMQ instance.
        :param str instance_id: ID of the instance.
        :param str instance_name: Name of the instance.
        :param int instance_status: The status of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        :param int instance_type: The type of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        :param str release_time: The automatic release time of an Enterprise Platinum Edition instance.
        :param str remark: This attribute is a concise description of instance.
        :param int status: The status of Ons instance. Valid values: `0` deploying, `2` arrears, `5` running, `7` upgrading.
        :param Mapping[str, str] tags: A map of tags assigned to the Ons instance.
        :param str tcp_endpoint: The TCP endpoint for the Message Queue for Apache RocketMQ instance.
        """
        pulumi.set(__self__, "http_internal_endpoint", http_internal_endpoint)
        pulumi.set(__self__, "http_internet_endpoint", http_internet_endpoint)
        pulumi.set(__self__, "http_internet_secure_endpoint", http_internet_secure_endpoint)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "independent_naming", independent_naming)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "instance_status", instance_status)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "release_time", release_time)
        pulumi.set(__self__, "remark", remark)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "tcp_endpoint", tcp_endpoint)

    @property
    @pulumi.getter(name="httpInternalEndpoint")
    def http_internal_endpoint(self) -> str:
        """
        The internal HTTP endpoint for the Message Queue for Apache RocketMQ instance.
        """
        return pulumi.get(self, "http_internal_endpoint")

    @property
    @pulumi.getter(name="httpInternetEndpoint")
    def http_internet_endpoint(self) -> str:
        """
        The public HTTP endpoint for the Message Queue for Apache RocketMQ instance.
        """
        return pulumi.get(self, "http_internet_endpoint")

    @property
    @pulumi.getter(name="httpInternetSecureEndpoint")
    def http_internet_secure_endpoint(self) -> str:
        """
        The public HTTPS endpoint for the Message Queue for Apache RocketMQ instance.
        """
        return pulumi.get(self, "http_internet_secure_endpoint")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the instance.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="independentNaming")
    def independent_naming(self) -> bool:
        """
        Indicates whether any namespace is configured for the Message Queue for Apache RocketMQ instance.
        """
        return pulumi.get(self, "independent_naming")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        """
        Name of the instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="instanceStatus")
    def instance_status(self) -> int:
        """
        The status of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        """
        return pulumi.get(self, "instance_status")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> int:
        """
        The type of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="releaseTime")
    def release_time(self) -> str:
        """
        The automatic release time of an Enterprise Platinum Edition instance.
        """
        return pulumi.get(self, "release_time")

    @property
    @pulumi.getter
    def remark(self) -> str:
        """
        This attribute is a concise description of instance.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter
    def status(self) -> int:
        """
        The status of Ons instance. Valid values: `0` deploying, `2` arrears, `5` running, `7` upgrading.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags assigned to the Ons instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tcpEndpoint")
    def tcp_endpoint(self) -> str:
        """
        The TCP endpoint for the Message Queue for Apache RocketMQ instance.
        """
        return pulumi.get(self, "tcp_endpoint")


@pulumi.output_type
class GetTopicsTopicResult(dict):
    def __init__(__self__, *,
                 id: str,
                 independent_naming: bool,
                 instance_id: str,
                 message_type: int,
                 owner: str,
                 perm: int,
                 relation: int,
                 relation_name: str,
                 remark: str,
                 tags: Mapping[str, str],
                 topic: str,
                 topic_name: str):
        """
        :param str id: The id of the topic.
        :param bool independent_naming: Indicates whether namespaces are available. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        :param str instance_id: ID of the ONS Instance that owns the topics.
        :param int message_type: The type of the message. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        :param str owner: The ID of the topic owner, which is the Alibaba Cloud UID.
        :param int perm: This attribute is used to set the read-write mode for the topic.
        :param int relation: The relation ID. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        :param str relation_name: The name of the relation, for example, owner, publishable, subscribable, and publishable and subscribable.
        :param str remark: Remark of the topic.
        :param Mapping[str, str] tags: A map of tags assigned to the Ons instance.
        :param str topic: The name of the topic.
        :param str topic_name: The name of the topic.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "independent_naming", independent_naming)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "message_type", message_type)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "perm", perm)
        pulumi.set(__self__, "relation", relation)
        pulumi.set(__self__, "relation_name", relation_name)
        pulumi.set(__self__, "remark", remark)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "topic", topic)
        pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the topic.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="independentNaming")
    def independent_naming(self) -> bool:
        """
        Indicates whether namespaces are available. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        """
        return pulumi.get(self, "independent_naming")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        ID of the ONS Instance that owns the topics.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> int:
        """
        The type of the message. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        """
        return pulumi.get(self, "message_type")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The ID of the topic owner, which is the Alibaba Cloud UID.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def perm(self) -> int:
        """
        This attribute is used to set the read-write mode for the topic.
        """
        return pulumi.get(self, "perm")

    @property
    @pulumi.getter
    def relation(self) -> int:
        """
        The relation ID. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
        """
        return pulumi.get(self, "relation")

    @property
    @pulumi.getter(name="relationName")
    def relation_name(self) -> str:
        """
        The name of the relation, for example, owner, publishable, subscribable, and publishable and subscribable.
        """
        return pulumi.get(self, "relation_name")

    @property
    @pulumi.getter
    def remark(self) -> str:
        """
        Remark of the topic.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags assigned to the Ons instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def topic(self) -> str:
        """
        The name of the topic.
        """
        return pulumi.get(self, "topic")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> str:
        """
        The name of the topic.
        """
        return pulumi.get(self, "topic_name")


