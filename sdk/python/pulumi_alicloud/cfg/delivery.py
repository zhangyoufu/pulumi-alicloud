# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DeliveryArgs', 'Delivery']

@pulumi.input_type
class DeliveryArgs:
    def __init__(__self__, *,
                 delivery_channel_target_arn: pulumi.Input[str],
                 delivery_channel_type: pulumi.Input[str],
                 configuration_item_change_notification: Optional[pulumi.Input[bool]] = None,
                 configuration_snapshot: Optional[pulumi.Input[bool]] = None,
                 delivery_channel_condition: Optional[pulumi.Input[str]] = None,
                 delivery_channel_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 non_compliant_notification: Optional[pulumi.Input[bool]] = None,
                 oversized_data_oss_target_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Delivery resource.
        :param pulumi.Input[str] delivery_channel_target_arn: The ARN of the delivery destination. The value must be in one of the following formats:
               * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
               * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
               * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        :param pulumi.Input[str] delivery_channel_type: The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        :param pulumi.Input[bool] configuration_item_change_notification: Open or close delivery configuration change history. true: open, false: close.
        :param pulumi.Input[bool] configuration_snapshot: Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
        :param pulumi.Input[str] delivery_channel_condition: The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
        :param pulumi.Input[str] delivery_channel_name: The name of the delivery method.
        :param pulumi.Input[str] description: The description of the delivery method.
        :param pulumi.Input[bool] non_compliant_notification: Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
        :param pulumi.Input[str] oversized_data_oss_target_arn: The oss ARN of the delivery channel when the value data oversized limit. 
               * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
               * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
        :param pulumi.Input[int] status: The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
        """
        pulumi.set(__self__, "delivery_channel_target_arn", delivery_channel_target_arn)
        pulumi.set(__self__, "delivery_channel_type", delivery_channel_type)
        if configuration_item_change_notification is not None:
            pulumi.set(__self__, "configuration_item_change_notification", configuration_item_change_notification)
        if configuration_snapshot is not None:
            pulumi.set(__self__, "configuration_snapshot", configuration_snapshot)
        if delivery_channel_condition is not None:
            pulumi.set(__self__, "delivery_channel_condition", delivery_channel_condition)
        if delivery_channel_name is not None:
            pulumi.set(__self__, "delivery_channel_name", delivery_channel_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if non_compliant_notification is not None:
            pulumi.set(__self__, "non_compliant_notification", non_compliant_notification)
        if oversized_data_oss_target_arn is not None:
            pulumi.set(__self__, "oversized_data_oss_target_arn", oversized_data_oss_target_arn)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="deliveryChannelTargetArn")
    def delivery_channel_target_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the delivery destination. The value must be in one of the following formats:
        * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
        * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
        * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        """
        return pulumi.get(self, "delivery_channel_target_arn")

    @delivery_channel_target_arn.setter
    def delivery_channel_target_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "delivery_channel_target_arn", value)

    @property
    @pulumi.getter(name="deliveryChannelType")
    def delivery_channel_type(self) -> pulumi.Input[str]:
        """
        The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        """
        return pulumi.get(self, "delivery_channel_type")

    @delivery_channel_type.setter
    def delivery_channel_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "delivery_channel_type", value)

    @property
    @pulumi.getter(name="configurationItemChangeNotification")
    def configuration_item_change_notification(self) -> Optional[pulumi.Input[bool]]:
        """
        Open or close delivery configuration change history. true: open, false: close.
        """
        return pulumi.get(self, "configuration_item_change_notification")

    @configuration_item_change_notification.setter
    def configuration_item_change_notification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "configuration_item_change_notification", value)

    @property
    @pulumi.getter(name="configurationSnapshot")
    def configuration_snapshot(self) -> Optional[pulumi.Input[bool]]:
        """
        Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
        """
        return pulumi.get(self, "configuration_snapshot")

    @configuration_snapshot.setter
    def configuration_snapshot(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "configuration_snapshot", value)

    @property
    @pulumi.getter(name="deliveryChannelCondition")
    def delivery_channel_condition(self) -> Optional[pulumi.Input[str]]:
        """
        The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
        """
        return pulumi.get(self, "delivery_channel_condition")

    @delivery_channel_condition.setter
    def delivery_channel_condition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delivery_channel_condition", value)

    @property
    @pulumi.getter(name="deliveryChannelName")
    def delivery_channel_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the delivery method.
        """
        return pulumi.get(self, "delivery_channel_name")

    @delivery_channel_name.setter
    def delivery_channel_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delivery_channel_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the delivery method.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="nonCompliantNotification")
    def non_compliant_notification(self) -> Optional[pulumi.Input[bool]]:
        """
        Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
        """
        return pulumi.get(self, "non_compliant_notification")

    @non_compliant_notification.setter
    def non_compliant_notification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "non_compliant_notification", value)

    @property
    @pulumi.getter(name="oversizedDataOssTargetArn")
    def oversized_data_oss_target_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The oss ARN of the delivery channel when the value data oversized limit. 
        * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
        * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
        """
        return pulumi.get(self, "oversized_data_oss_target_arn")

    @oversized_data_oss_target_arn.setter
    def oversized_data_oss_target_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oversized_data_oss_target_arn", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _DeliveryState:
    def __init__(__self__, *,
                 configuration_item_change_notification: Optional[pulumi.Input[bool]] = None,
                 configuration_snapshot: Optional[pulumi.Input[bool]] = None,
                 delivery_channel_condition: Optional[pulumi.Input[str]] = None,
                 delivery_channel_name: Optional[pulumi.Input[str]] = None,
                 delivery_channel_target_arn: Optional[pulumi.Input[str]] = None,
                 delivery_channel_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 non_compliant_notification: Optional[pulumi.Input[bool]] = None,
                 oversized_data_oss_target_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Delivery resources.
        :param pulumi.Input[bool] configuration_item_change_notification: Open or close delivery configuration change history. true: open, false: close.
        :param pulumi.Input[bool] configuration_snapshot: Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
        :param pulumi.Input[str] delivery_channel_condition: The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
        :param pulumi.Input[str] delivery_channel_name: The name of the delivery method.
        :param pulumi.Input[str] delivery_channel_target_arn: The ARN of the delivery destination. The value must be in one of the following formats:
               * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
               * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
               * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        :param pulumi.Input[str] delivery_channel_type: The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        :param pulumi.Input[str] description: The description of the delivery method.
        :param pulumi.Input[bool] non_compliant_notification: Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
        :param pulumi.Input[str] oversized_data_oss_target_arn: The oss ARN of the delivery channel when the value data oversized limit. 
               * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
               * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
        :param pulumi.Input[int] status: The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
        """
        if configuration_item_change_notification is not None:
            pulumi.set(__self__, "configuration_item_change_notification", configuration_item_change_notification)
        if configuration_snapshot is not None:
            pulumi.set(__self__, "configuration_snapshot", configuration_snapshot)
        if delivery_channel_condition is not None:
            pulumi.set(__self__, "delivery_channel_condition", delivery_channel_condition)
        if delivery_channel_name is not None:
            pulumi.set(__self__, "delivery_channel_name", delivery_channel_name)
        if delivery_channel_target_arn is not None:
            pulumi.set(__self__, "delivery_channel_target_arn", delivery_channel_target_arn)
        if delivery_channel_type is not None:
            pulumi.set(__self__, "delivery_channel_type", delivery_channel_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if non_compliant_notification is not None:
            pulumi.set(__self__, "non_compliant_notification", non_compliant_notification)
        if oversized_data_oss_target_arn is not None:
            pulumi.set(__self__, "oversized_data_oss_target_arn", oversized_data_oss_target_arn)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="configurationItemChangeNotification")
    def configuration_item_change_notification(self) -> Optional[pulumi.Input[bool]]:
        """
        Open or close delivery configuration change history. true: open, false: close.
        """
        return pulumi.get(self, "configuration_item_change_notification")

    @configuration_item_change_notification.setter
    def configuration_item_change_notification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "configuration_item_change_notification", value)

    @property
    @pulumi.getter(name="configurationSnapshot")
    def configuration_snapshot(self) -> Optional[pulumi.Input[bool]]:
        """
        Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
        """
        return pulumi.get(self, "configuration_snapshot")

    @configuration_snapshot.setter
    def configuration_snapshot(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "configuration_snapshot", value)

    @property
    @pulumi.getter(name="deliveryChannelCondition")
    def delivery_channel_condition(self) -> Optional[pulumi.Input[str]]:
        """
        The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
        """
        return pulumi.get(self, "delivery_channel_condition")

    @delivery_channel_condition.setter
    def delivery_channel_condition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delivery_channel_condition", value)

    @property
    @pulumi.getter(name="deliveryChannelName")
    def delivery_channel_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the delivery method.
        """
        return pulumi.get(self, "delivery_channel_name")

    @delivery_channel_name.setter
    def delivery_channel_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delivery_channel_name", value)

    @property
    @pulumi.getter(name="deliveryChannelTargetArn")
    def delivery_channel_target_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the delivery destination. The value must be in one of the following formats:
        * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
        * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
        * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        """
        return pulumi.get(self, "delivery_channel_target_arn")

    @delivery_channel_target_arn.setter
    def delivery_channel_target_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delivery_channel_target_arn", value)

    @property
    @pulumi.getter(name="deliveryChannelType")
    def delivery_channel_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        """
        return pulumi.get(self, "delivery_channel_type")

    @delivery_channel_type.setter
    def delivery_channel_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delivery_channel_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the delivery method.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="nonCompliantNotification")
    def non_compliant_notification(self) -> Optional[pulumi.Input[bool]]:
        """
        Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
        """
        return pulumi.get(self, "non_compliant_notification")

    @non_compliant_notification.setter
    def non_compliant_notification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "non_compliant_notification", value)

    @property
    @pulumi.getter(name="oversizedDataOssTargetArn")
    def oversized_data_oss_target_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The oss ARN of the delivery channel when the value data oversized limit. 
        * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
        * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
        """
        return pulumi.get(self, "oversized_data_oss_target_arn")

    @oversized_data_oss_target_arn.setter
    def oversized_data_oss_target_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oversized_data_oss_target_arn", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)


class Delivery(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_item_change_notification: Optional[pulumi.Input[bool]] = None,
                 configuration_snapshot: Optional[pulumi.Input[bool]] = None,
                 delivery_channel_condition: Optional[pulumi.Input[str]] = None,
                 delivery_channel_name: Optional[pulumi.Input[str]] = None,
                 delivery_channel_target_arn: Optional[pulumi.Input[str]] = None,
                 delivery_channel_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 non_compliant_notification: Optional[pulumi.Input[bool]] = None,
                 oversized_data_oss_target_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a Cloud Config Delivery resource.

        For information about Cloud Config Delivery and how to use it, see [What is Delivery](https://www.alibabacloud.com/help/en/cloud-config/latest/api-doc-config-2020-09-07-api-doc-createconfigdeliverychannel).

        > **NOTE:** Available since v1.171.0+.

        ## Import

        Cloud Config Delivery can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cfg/delivery:Delivery example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] configuration_item_change_notification: Open or close delivery configuration change history. true: open, false: close.
        :param pulumi.Input[bool] configuration_snapshot: Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
        :param pulumi.Input[str] delivery_channel_condition: The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
        :param pulumi.Input[str] delivery_channel_name: The name of the delivery method.
        :param pulumi.Input[str] delivery_channel_target_arn: The ARN of the delivery destination. The value must be in one of the following formats:
               * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
               * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
               * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        :param pulumi.Input[str] delivery_channel_type: The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        :param pulumi.Input[str] description: The description of the delivery method.
        :param pulumi.Input[bool] non_compliant_notification: Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
        :param pulumi.Input[str] oversized_data_oss_target_arn: The oss ARN of the delivery channel when the value data oversized limit. 
               * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
               * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
        :param pulumi.Input[int] status: The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeliveryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Config Delivery resource.

        For information about Cloud Config Delivery and how to use it, see [What is Delivery](https://www.alibabacloud.com/help/en/cloud-config/latest/api-doc-config-2020-09-07-api-doc-createconfigdeliverychannel).

        > **NOTE:** Available since v1.171.0+.

        ## Import

        Cloud Config Delivery can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:cfg/delivery:Delivery example <id>
        ```

        :param str resource_name: The name of the resource.
        :param DeliveryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeliveryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_item_change_notification: Optional[pulumi.Input[bool]] = None,
                 configuration_snapshot: Optional[pulumi.Input[bool]] = None,
                 delivery_channel_condition: Optional[pulumi.Input[str]] = None,
                 delivery_channel_name: Optional[pulumi.Input[str]] = None,
                 delivery_channel_target_arn: Optional[pulumi.Input[str]] = None,
                 delivery_channel_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 non_compliant_notification: Optional[pulumi.Input[bool]] = None,
                 oversized_data_oss_target_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeliveryArgs.__new__(DeliveryArgs)

            __props__.__dict__["configuration_item_change_notification"] = configuration_item_change_notification
            __props__.__dict__["configuration_snapshot"] = configuration_snapshot
            __props__.__dict__["delivery_channel_condition"] = delivery_channel_condition
            __props__.__dict__["delivery_channel_name"] = delivery_channel_name
            if delivery_channel_target_arn is None and not opts.urn:
                raise TypeError("Missing required property 'delivery_channel_target_arn'")
            __props__.__dict__["delivery_channel_target_arn"] = delivery_channel_target_arn
            if delivery_channel_type is None and not opts.urn:
                raise TypeError("Missing required property 'delivery_channel_type'")
            __props__.__dict__["delivery_channel_type"] = delivery_channel_type
            __props__.__dict__["description"] = description
            __props__.__dict__["non_compliant_notification"] = non_compliant_notification
            __props__.__dict__["oversized_data_oss_target_arn"] = oversized_data_oss_target_arn
            __props__.__dict__["status"] = status
        super(Delivery, __self__).__init__(
            'alicloud:cfg/delivery:Delivery',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            configuration_item_change_notification: Optional[pulumi.Input[bool]] = None,
            configuration_snapshot: Optional[pulumi.Input[bool]] = None,
            delivery_channel_condition: Optional[pulumi.Input[str]] = None,
            delivery_channel_name: Optional[pulumi.Input[str]] = None,
            delivery_channel_target_arn: Optional[pulumi.Input[str]] = None,
            delivery_channel_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            non_compliant_notification: Optional[pulumi.Input[bool]] = None,
            oversized_data_oss_target_arn: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[int]] = None) -> 'Delivery':
        """
        Get an existing Delivery resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] configuration_item_change_notification: Open or close delivery configuration change history. true: open, false: close.
        :param pulumi.Input[bool] configuration_snapshot: Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
        :param pulumi.Input[str] delivery_channel_condition: The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
        :param pulumi.Input[str] delivery_channel_name: The name of the delivery method.
        :param pulumi.Input[str] delivery_channel_target_arn: The ARN of the delivery destination. The value must be in one of the following formats:
               * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
               * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
               * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        :param pulumi.Input[str] delivery_channel_type: The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        :param pulumi.Input[str] description: The description of the delivery method.
        :param pulumi.Input[bool] non_compliant_notification: Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
        :param pulumi.Input[str] oversized_data_oss_target_arn: The oss ARN of the delivery channel when the value data oversized limit. 
               * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
               * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
        :param pulumi.Input[int] status: The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeliveryState.__new__(_DeliveryState)

        __props__.__dict__["configuration_item_change_notification"] = configuration_item_change_notification
        __props__.__dict__["configuration_snapshot"] = configuration_snapshot
        __props__.__dict__["delivery_channel_condition"] = delivery_channel_condition
        __props__.__dict__["delivery_channel_name"] = delivery_channel_name
        __props__.__dict__["delivery_channel_target_arn"] = delivery_channel_target_arn
        __props__.__dict__["delivery_channel_type"] = delivery_channel_type
        __props__.__dict__["description"] = description
        __props__.__dict__["non_compliant_notification"] = non_compliant_notification
        __props__.__dict__["oversized_data_oss_target_arn"] = oversized_data_oss_target_arn
        __props__.__dict__["status"] = status
        return Delivery(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configurationItemChangeNotification")
    def configuration_item_change_notification(self) -> pulumi.Output[bool]:
        """
        Open or close delivery configuration change history. true: open, false: close.
        """
        return pulumi.get(self, "configuration_item_change_notification")

    @property
    @pulumi.getter(name="configurationSnapshot")
    def configuration_snapshot(self) -> pulumi.Output[bool]:
        """
        Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
        """
        return pulumi.get(self, "configuration_snapshot")

    @property
    @pulumi.getter(name="deliveryChannelCondition")
    def delivery_channel_condition(self) -> pulumi.Output[Optional[str]]:
        """
        The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
        """
        return pulumi.get(self, "delivery_channel_condition")

    @property
    @pulumi.getter(name="deliveryChannelName")
    def delivery_channel_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the delivery method.
        """
        return pulumi.get(self, "delivery_channel_name")

    @property
    @pulumi.getter(name="deliveryChannelTargetArn")
    def delivery_channel_target_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the delivery destination. The value must be in one of the following formats:
        * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
        * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
        * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        """
        return pulumi.get(self, "delivery_channel_target_arn")

    @property
    @pulumi.getter(name="deliveryChannelType")
    def delivery_channel_type(self) -> pulumi.Output[str]:
        """
        The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        """
        return pulumi.get(self, "delivery_channel_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the delivery method.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="nonCompliantNotification")
    def non_compliant_notification(self) -> pulumi.Output[bool]:
        """
        Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
        """
        return pulumi.get(self, "non_compliant_notification")

    @property
    @pulumi.getter(name="oversizedDataOssTargetArn")
    def oversized_data_oss_target_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The oss ARN of the delivery channel when the value data oversized limit. 
        * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
        * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
        """
        return pulumi.get(self, "oversized_data_oss_target_arn")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[int]:
        """
        The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
        """
        return pulumi.get(self, "status")

