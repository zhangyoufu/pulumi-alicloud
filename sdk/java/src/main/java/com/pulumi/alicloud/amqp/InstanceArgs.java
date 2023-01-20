// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.amqp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceArgs Empty = new InstanceArgs();

    /**
     * The instance name.
     * 
     */
    @Import(name="instanceName")
    private @Nullable Output<String> instanceName;

    /**
     * @return The instance name.
     * 
     */
    public Optional<Output<String>> instanceName() {
        return Optional.ofNullable(this.instanceName);
    }

    /**
     * The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
     * 
     */
    @Import(name="instanceType", required=true)
    private Output<String> instanceType;

    /**
     * @return The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }

    @Import(name="logistics")
    private @Nullable Output<String> logistics;

    public Optional<Output<String>> logistics() {
        return Optional.ofNullable(this.logistics);
    }

    /**
     * The max eip tps. It is valid when `support_eip` is true. The valid value is [128, 45000] with the step size 128.
     * 
     */
    @Import(name="maxEipTps")
    private @Nullable Output<String> maxEipTps;

    /**
     * @return The max eip tps. It is valid when `support_eip` is true. The valid value is [128, 45000] with the step size 128.
     * 
     */
    public Optional<Output<String>> maxEipTps() {
        return Optional.ofNullable(this.maxEipTps);
    }

    /**
     * The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
     * 
     */
    @Import(name="maxTps", required=true)
    private Output<String> maxTps;

    /**
     * @return The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
     * 
     */
    public Output<String> maxTps() {
        return this.maxTps;
    }

    /**
     * The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
     * 
     */
    @Import(name="modifyType")
    private @Nullable Output<String> modifyType;

    /**
     * @return The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
     * 
     */
    public Optional<Output<String>> modifyType() {
        return Optional.ofNullable(this.modifyType);
    }

    /**
     * The payment type. Valid values: `Subscription`.
     * 
     */
    @Import(name="paymentType", required=true)
    private Output<String> paymentType;

    /**
     * @return The payment type. Valid values: `Subscription`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }

    /**
     * The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The queue capacity. The smallest value is 50 and the step size 5.
     * 
     */
    @Import(name="queueCapacity", required=true)
    private Output<String> queueCapacity;

    /**
     * @return The queue capacity. The smallest value is 50 and the step size 5.
     * 
     */
    public Output<String> queueCapacity() {
        return this.queueCapacity;
    }

    /**
     * RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
     * 
     */
    @Import(name="renewalDuration")
    private @Nullable Output<Integer> renewalDuration;

    /**
     * @return RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
     * 
     */
    public Optional<Output<Integer>> renewalDuration() {
        return Optional.ofNullable(this.renewalDuration);
    }

    /**
     * Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
     * 
     */
    @Import(name="renewalDurationUnit")
    private @Nullable Output<String> renewalDurationUnit;

    /**
     * @return Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
     * 
     */
    public Optional<Output<String>> renewalDurationUnit() {
        return Optional.ofNullable(this.renewalDurationUnit);
    }

    /**
     * Whether to renew an instance automatically or not. Default to &#34;ManualRenewal&#34;.
     * 
     */
    @Import(name="renewalStatus")
    private @Nullable Output<String> renewalStatus;

    /**
     * @return Whether to renew an instance automatically or not. Default to &#34;ManualRenewal&#34;.
     * 
     */
    public Optional<Output<String>> renewalStatus() {
        return Optional.ofNullable(this.renewalStatus);
    }

    /**
     * The storage size. It is valid when `instance_type` is vip.
     * 
     */
    @Import(name="storageSize")
    private @Nullable Output<String> storageSize;

    /**
     * @return The storage size. It is valid when `instance_type` is vip.
     * 
     */
    public Optional<Output<String>> storageSize() {
        return Optional.ofNullable(this.storageSize);
    }

    /**
     * Whether to support EIP.
     * 
     */
    @Import(name="supportEip", required=true)
    private Output<Boolean> supportEip;

    /**
     * @return Whether to support EIP.
     * 
     */
    public Output<Boolean> supportEip() {
        return this.supportEip;
    }

    private InstanceArgs() {}

    private InstanceArgs(InstanceArgs $) {
        this.instanceName = $.instanceName;
        this.instanceType = $.instanceType;
        this.logistics = $.logistics;
        this.maxEipTps = $.maxEipTps;
        this.maxTps = $.maxTps;
        this.modifyType = $.modifyType;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.queueCapacity = $.queueCapacity;
        this.renewalDuration = $.renewalDuration;
        this.renewalDurationUnit = $.renewalDurationUnit;
        this.renewalStatus = $.renewalStatus;
        this.storageSize = $.storageSize;
        this.supportEip = $.supportEip;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceArgs $;

        public Builder() {
            $ = new InstanceArgs();
        }

        public Builder(InstanceArgs defaults) {
            $ = new InstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceName The instance name.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(@Nullable Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName The instance name.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        /**
         * @param instanceType The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        public Builder logistics(@Nullable Output<String> logistics) {
            $.logistics = logistics;
            return this;
        }

        public Builder logistics(String logistics) {
            return logistics(Output.of(logistics));
        }

        /**
         * @param maxEipTps The max eip tps. It is valid when `support_eip` is true. The valid value is [128, 45000] with the step size 128.
         * 
         * @return builder
         * 
         */
        public Builder maxEipTps(@Nullable Output<String> maxEipTps) {
            $.maxEipTps = maxEipTps;
            return this;
        }

        /**
         * @param maxEipTps The max eip tps. It is valid when `support_eip` is true. The valid value is [128, 45000] with the step size 128.
         * 
         * @return builder
         * 
         */
        public Builder maxEipTps(String maxEipTps) {
            return maxEipTps(Output.of(maxEipTps));
        }

        /**
         * @param maxTps The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
         * 
         * @return builder
         * 
         */
        public Builder maxTps(Output<String> maxTps) {
            $.maxTps = maxTps;
            return this;
        }

        /**
         * @param maxTps The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
         * 
         * @return builder
         * 
         */
        public Builder maxTps(String maxTps) {
            return maxTps(Output.of(maxTps));
        }

        /**
         * @param modifyType The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
         * 
         * @return builder
         * 
         */
        public Builder modifyType(@Nullable Output<String> modifyType) {
            $.modifyType = modifyType;
            return this;
        }

        /**
         * @param modifyType The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
         * 
         * @return builder
         * 
         */
        public Builder modifyType(String modifyType) {
            return modifyType(Output.of(modifyType));
        }

        /**
         * @param paymentType The payment type. Valid values: `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The payment type. Valid values: `Subscription`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param period The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param queueCapacity The queue capacity. The smallest value is 50 and the step size 5.
         * 
         * @return builder
         * 
         */
        public Builder queueCapacity(Output<String> queueCapacity) {
            $.queueCapacity = queueCapacity;
            return this;
        }

        /**
         * @param queueCapacity The queue capacity. The smallest value is 50 and the step size 5.
         * 
         * @return builder
         * 
         */
        public Builder queueCapacity(String queueCapacity) {
            return queueCapacity(Output.of(queueCapacity));
        }

        /**
         * @param renewalDuration RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
         * 
         * @return builder
         * 
         */
        public Builder renewalDuration(@Nullable Output<Integer> renewalDuration) {
            $.renewalDuration = renewalDuration;
            return this;
        }

        /**
         * @param renewalDuration RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
         * 
         * @return builder
         * 
         */
        public Builder renewalDuration(Integer renewalDuration) {
            return renewalDuration(Output.of(renewalDuration));
        }

        /**
         * @param renewalDurationUnit Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
         * 
         * @return builder
         * 
         */
        public Builder renewalDurationUnit(@Nullable Output<String> renewalDurationUnit) {
            $.renewalDurationUnit = renewalDurationUnit;
            return this;
        }

        /**
         * @param renewalDurationUnit Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
         * 
         * @return builder
         * 
         */
        public Builder renewalDurationUnit(String renewalDurationUnit) {
            return renewalDurationUnit(Output.of(renewalDurationUnit));
        }

        /**
         * @param renewalStatus Whether to renew an instance automatically or not. Default to &#34;ManualRenewal&#34;.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(@Nullable Output<String> renewalStatus) {
            $.renewalStatus = renewalStatus;
            return this;
        }

        /**
         * @param renewalStatus Whether to renew an instance automatically or not. Default to &#34;ManualRenewal&#34;.
         * 
         * @return builder
         * 
         */
        public Builder renewalStatus(String renewalStatus) {
            return renewalStatus(Output.of(renewalStatus));
        }

        /**
         * @param storageSize The storage size. It is valid when `instance_type` is vip.
         * 
         * @return builder
         * 
         */
        public Builder storageSize(@Nullable Output<String> storageSize) {
            $.storageSize = storageSize;
            return this;
        }

        /**
         * @param storageSize The storage size. It is valid when `instance_type` is vip.
         * 
         * @return builder
         * 
         */
        public Builder storageSize(String storageSize) {
            return storageSize(Output.of(storageSize));
        }

        /**
         * @param supportEip Whether to support EIP.
         * 
         * @return builder
         * 
         */
        public Builder supportEip(Output<Boolean> supportEip) {
            $.supportEip = supportEip;
            return this;
        }

        /**
         * @param supportEip Whether to support EIP.
         * 
         * @return builder
         * 
         */
        public Builder supportEip(Boolean supportEip) {
            return supportEip(Output.of(supportEip));
        }

        public InstanceArgs build() {
            $.instanceType = Objects.requireNonNull($.instanceType, "expected parameter 'instanceType' to be non-null");
            $.maxTps = Objects.requireNonNull($.maxTps, "expected parameter 'maxTps' to be non-null");
            $.paymentType = Objects.requireNonNull($.paymentType, "expected parameter 'paymentType' to be non-null");
            $.queueCapacity = Objects.requireNonNull($.queueCapacity, "expected parameter 'queueCapacity' to be non-null");
            $.supportEip = Objects.requireNonNull($.supportEip, "expected parameter 'supportEip' to be non-null");
            return $;
        }
    }

}
