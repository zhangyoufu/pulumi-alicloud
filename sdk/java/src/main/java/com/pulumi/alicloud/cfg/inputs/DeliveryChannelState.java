// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DeliveryChannelState extends com.pulumi.resources.ResourceArgs {

    public static final DeliveryChannelState Empty = new DeliveryChannelState();

    /**
     * The Alibaba Cloud Resource Name (ARN) of the role to be assumed by the delivery method.
     * 
     */
    @Import(name="deliveryChannelAssumeRoleArn")
    private @Nullable Output<String> deliveryChannelAssumeRoleArn;

    /**
     * @return The Alibaba Cloud Resource Name (ARN) of the role to be assumed by the delivery method.
     * 
     */
    public Optional<Output<String>> deliveryChannelAssumeRoleArn() {
        return Optional.ofNullable(this.deliveryChannelAssumeRoleArn);
    }

    /**
     * The rule attached to the delivery method. This parameter is applicable only to delivery methods of the MNS type. Please refer to api [PutDeliveryChannel](https://www.alibabacloud.com/help/en/doc-detail/174253.htm) for example format.
     * 
     */
    @Import(name="deliveryChannelCondition")
    private @Nullable Output<String> deliveryChannelCondition;

    /**
     * @return The rule attached to the delivery method. This parameter is applicable only to delivery methods of the MNS type. Please refer to api [PutDeliveryChannel](https://www.alibabacloud.com/help/en/doc-detail/174253.htm) for example format.
     * 
     */
    public Optional<Output<String>> deliveryChannelCondition() {
        return Optional.ofNullable(this.deliveryChannelCondition);
    }

    /**
     * The name of the delivery channel.
     * 
     */
    @Import(name="deliveryChannelName")
    private @Nullable Output<String> deliveryChannelName;

    /**
     * @return The name of the delivery channel.
     * 
     */
    public Optional<Output<String>> deliveryChannelName() {
        return Optional.ofNullable(this.deliveryChannelName);
    }

    /**
     * The ARN of the delivery destination. This parameter is required when you create a delivery method. The value must be in one of the following formats:
     * - `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
     * - `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
     * - `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
     * 
     */
    @Import(name="deliveryChannelTargetArn")
    private @Nullable Output<String> deliveryChannelTargetArn;

    /**
     * @return The ARN of the delivery destination. This parameter is required when you create a delivery method. The value must be in one of the following formats:
     * - `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
     * - `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
     * - `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
     * 
     */
    public Optional<Output<String>> deliveryChannelTargetArn() {
        return Optional.ofNullable(this.deliveryChannelTargetArn);
    }

    /**
     * The type of the delivery method. This parameter is required when you create a delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
     * 
     */
    @Import(name="deliveryChannelType")
    private @Nullable Output<String> deliveryChannelType;

    /**
     * @return The type of the delivery method. This parameter is required when you create a delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
     * 
     */
    public Optional<Output<String>> deliveryChannelType() {
        return Optional.ofNullable(this.deliveryChannelType);
    }

    /**
     * The description of the delivery method.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the delivery method.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The status of the delivery method. Valid values: `0`: The delivery method is disabled., `1`: The delivery destination is enabled. This is the default value.
     * 
     */
    @Import(name="status")
    private @Nullable Output<Integer> status;

    /**
     * @return The status of the delivery method. Valid values: `0`: The delivery method is disabled., `1`: The delivery destination is enabled. This is the default value.
     * 
     */
    public Optional<Output<Integer>> status() {
        return Optional.ofNullable(this.status);
    }

    private DeliveryChannelState() {}

    private DeliveryChannelState(DeliveryChannelState $) {
        this.deliveryChannelAssumeRoleArn = $.deliveryChannelAssumeRoleArn;
        this.deliveryChannelCondition = $.deliveryChannelCondition;
        this.deliveryChannelName = $.deliveryChannelName;
        this.deliveryChannelTargetArn = $.deliveryChannelTargetArn;
        this.deliveryChannelType = $.deliveryChannelType;
        this.description = $.description;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DeliveryChannelState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DeliveryChannelState $;

        public Builder() {
            $ = new DeliveryChannelState();
        }

        public Builder(DeliveryChannelState defaults) {
            $ = new DeliveryChannelState(Objects.requireNonNull(defaults));
        }

        /**
         * @param deliveryChannelAssumeRoleArn The Alibaba Cloud Resource Name (ARN) of the role to be assumed by the delivery method.
         * 
         * @return builder
         * 
         */
        public Builder deliveryChannelAssumeRoleArn(@Nullable Output<String> deliveryChannelAssumeRoleArn) {
            $.deliveryChannelAssumeRoleArn = deliveryChannelAssumeRoleArn;
            return this;
        }

        /**
         * @param deliveryChannelAssumeRoleArn The Alibaba Cloud Resource Name (ARN) of the role to be assumed by the delivery method.
         * 
         * @return builder
         * 
         */
        public Builder deliveryChannelAssumeRoleArn(String deliveryChannelAssumeRoleArn) {
            return deliveryChannelAssumeRoleArn(Output.of(deliveryChannelAssumeRoleArn));
        }

        /**
         * @param deliveryChannelCondition The rule attached to the delivery method. This parameter is applicable only to delivery methods of the MNS type. Please refer to api [PutDeliveryChannel](https://www.alibabacloud.com/help/en/doc-detail/174253.htm) for example format.
         * 
         * @return builder
         * 
         */
        public Builder deliveryChannelCondition(@Nullable Output<String> deliveryChannelCondition) {
            $.deliveryChannelCondition = deliveryChannelCondition;
            return this;
        }

        /**
         * @param deliveryChannelCondition The rule attached to the delivery method. This parameter is applicable only to delivery methods of the MNS type. Please refer to api [PutDeliveryChannel](https://www.alibabacloud.com/help/en/doc-detail/174253.htm) for example format.
         * 
         * @return builder
         * 
         */
        public Builder deliveryChannelCondition(String deliveryChannelCondition) {
            return deliveryChannelCondition(Output.of(deliveryChannelCondition));
        }

        /**
         * @param deliveryChannelName The name of the delivery channel.
         * 
         * @return builder
         * 
         */
        public Builder deliveryChannelName(@Nullable Output<String> deliveryChannelName) {
            $.deliveryChannelName = deliveryChannelName;
            return this;
        }

        /**
         * @param deliveryChannelName The name of the delivery channel.
         * 
         * @return builder
         * 
         */
        public Builder deliveryChannelName(String deliveryChannelName) {
            return deliveryChannelName(Output.of(deliveryChannelName));
        }

        /**
         * @param deliveryChannelTargetArn The ARN of the delivery destination. This parameter is required when you create a delivery method. The value must be in one of the following formats:
         * - `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
         * - `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
         * - `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
         * 
         * @return builder
         * 
         */
        public Builder deliveryChannelTargetArn(@Nullable Output<String> deliveryChannelTargetArn) {
            $.deliveryChannelTargetArn = deliveryChannelTargetArn;
            return this;
        }

        /**
         * @param deliveryChannelTargetArn The ARN of the delivery destination. This parameter is required when you create a delivery method. The value must be in one of the following formats:
         * - `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
         * - `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
         * - `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
         * 
         * @return builder
         * 
         */
        public Builder deliveryChannelTargetArn(String deliveryChannelTargetArn) {
            return deliveryChannelTargetArn(Output.of(deliveryChannelTargetArn));
        }

        /**
         * @param deliveryChannelType The type of the delivery method. This parameter is required when you create a delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
         * 
         * @return builder
         * 
         */
        public Builder deliveryChannelType(@Nullable Output<String> deliveryChannelType) {
            $.deliveryChannelType = deliveryChannelType;
            return this;
        }

        /**
         * @param deliveryChannelType The type of the delivery method. This parameter is required when you create a delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
         * 
         * @return builder
         * 
         */
        public Builder deliveryChannelType(String deliveryChannelType) {
            return deliveryChannelType(Output.of(deliveryChannelType));
        }

        /**
         * @param description The description of the delivery method.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the delivery method.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param status The status of the delivery method. Valid values: `0`: The delivery method is disabled., `1`: The delivery destination is enabled. This is the default value.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<Integer> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the delivery method. Valid values: `0`: The delivery method is disabled., `1`: The delivery destination is enabled. This is the default value.
         * 
         * @return builder
         * 
         */
        public Builder status(Integer status) {
            return status(Output.of(status));
        }

        public DeliveryChannelState build() {
            return $;
        }
    }

}
