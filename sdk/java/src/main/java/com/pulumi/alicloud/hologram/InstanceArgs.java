// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hologram;

import com.pulumi.alicloud.hologram.inputs.InstanceEndpointArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceArgs Empty = new InstanceArgs();

    /**
     * Whether to pay automatically. The default value is true. Value:
     * - true: automatic payment
     * - false: only generate orders, not pay
     * &gt; **NOTE:**  The default value is true. If the balance of your payment method is insufficient, you can set the parameter AutoPay to false, and an unpaid order will be generated. You can log in to the user Center to pay by yourself.
     * 
     */
    @Import(name="autoPay")
    private @Nullable Output<Boolean> autoPay;

    /**
     * @return Whether to pay automatically. The default value is true. Value:
     * - true: automatic payment
     * - false: only generate orders, not pay
     * &gt; **NOTE:**  The default value is true. If the balance of your payment method is insufficient, you can set the parameter AutoPay to false, and an unpaid order will be generated. You can log in to the user Center to pay by yourself.
     * 
     */
    public Optional<Output<Boolean>> autoPay() {
        return Optional.ofNullable(this.autoPay);
    }

    /**
     * Instance low-frequency storage space. Unit: GB.
     * &gt; **NOTE:**  PayAsYouGo (PostPaid) instances ignore this parameter.
     * 
     */
    @Import(name="coldStorageSize")
    private @Nullable Output<Integer> coldStorageSize;

    /**
     * @return Instance low-frequency storage space. Unit: GB.
     * &gt; **NOTE:**  PayAsYouGo (PostPaid) instances ignore this parameter.
     * 
     */
    public Optional<Output<Integer>> coldStorageSize() {
        return Optional.ofNullable(this.coldStorageSize);
    }

    /**
     * Instance specifications. Value:
     * - 8 cores 32 GB (number of compute nodes: 1)
     * - 16 cores 64 GB (number of compute nodes: 1)
     * - 32 core 128 GB (number of compute nodes: 2)
     * - 64 core 256 GB (number of compute nodes: 4)
     * - 96 core 384 GB (number of computing nodes: 6)
     * - 128 core 512 GB (number of compute nodes: 8)
     * &gt; **NOTE:** Just fill in the audit number. Please submit a work order application for purchasing 1024 or above specifications. Shared instance types do not need to specify specifications. The specification of - 8 core 32GB (number of computing nodes: 1) is only for experience use and cannot be used for production.
     * 
     */
    @Import(name="cpu")
    private @Nullable Output<Integer> cpu;

    /**
     * @return Instance specifications. Value:
     * - 8 cores 32 GB (number of compute nodes: 1)
     * - 16 cores 64 GB (number of compute nodes: 1)
     * - 32 core 128 GB (number of compute nodes: 2)
     * - 64 core 256 GB (number of compute nodes: 4)
     * - 96 core 384 GB (number of computing nodes: 6)
     * - 128 core 512 GB (number of compute nodes: 8)
     * &gt; **NOTE:** Just fill in the audit number. Please submit a work order application for purchasing 1024 or above specifications. Shared instance types do not need to specify specifications. The specification of - 8 core 32GB (number of computing nodes: 1) is only for experience use and cannot be used for production.
     * 
     */
    public Optional<Output<Integer>> cpu() {
        return Optional.ofNullable(this.cpu);
    }

    /**
     * The buying cycle. Buy for 2 months. If the Payment type is PayAsYouGo (PostPaid), you do not need to specify it.
     * 
     */
    @Import(name="duration")
    private @Nullable Output<Integer> duration;

    /**
     * @return The buying cycle. Buy for 2 months. If the Payment type is PayAsYouGo (PostPaid), you do not need to specify it.
     * 
     */
    public Optional<Output<Integer>> duration() {
        return Optional.ofNullable(this.duration);
    }

    /**
     * List of domain names. See `endpoints` below.
     * 
     */
    @Import(name="endpoints")
    private @Nullable Output<List<InstanceEndpointArgs>> endpoints;

    /**
     * @return List of domain names. See `endpoints` below.
     * 
     */
    public Optional<Output<List<InstanceEndpointArgs>>> endpoints() {
        return Optional.ofNullable(this.endpoints);
    }

    /**
     * Number of gateway nodes.
     * 
     */
    @Import(name="gatewayCount")
    private @Nullable Output<Integer> gatewayCount;

    /**
     * @return Number of gateway nodes.
     * 
     */
    public Optional<Output<Integer>> gatewayCount() {
        return Optional.ofNullable(this.gatewayCount);
    }

    /**
     * Initialize the database and split multiple database names &#34;,&#34;.
     * 
     */
    @Import(name="initialDatabases")
    private @Nullable Output<String> initialDatabases;

    /**
     * @return Initialize the database and split multiple database names &#34;,&#34;.
     * 
     */
    public Optional<Output<String>> initialDatabases() {
        return Optional.ofNullable(this.initialDatabases);
    }

    /**
     * The name of the resource.
     * 
     */
    @Import(name="instanceName", required=true)
    private Output<String> instanceName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }

    /**
     * The instance type. Value:
     * - Standard: Universal.
     * - Follower: Read-only slave instance.
     * - Warehouse: calculation group type.
     * - Shared: Shared.
     * 
     */
    @Import(name="instanceType", required=true)
    private Output<String> instanceType;

    /**
     * @return The instance type. Value:
     * - Standard: Universal.
     * - Follower: Read-only slave instance.
     * - Warehouse: calculation group type.
     * - Shared: Shared.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }

    /**
     * The ID of the primary instance.
     * 
     */
    @Import(name="leaderInstanceId")
    private @Nullable Output<String> leaderInstanceId;

    /**
     * @return The ID of the primary instance.
     * 
     */
    public Optional<Output<String>> leaderInstanceId() {
        return Optional.ofNullable(this.leaderInstanceId);
    }

    /**
     * The payment type of the resource.
     * 
     */
    @Import(name="paymentType", required=true)
    private Output<String> paymentType;

    /**
     * @return The payment type of the resource.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }

    /**
     * Billing cycle. Value:
     * - Month: monthly billing
     * - Hour: hourly billing
     * &gt; **NOTE:**  Subscription instances (PrePaid) only supports Month. PayAsYouGo instances (PostPaid) only supports Hour. The Shared type is automatically set to Hour without specifying it.
     * 
     */
    @Import(name="pricingCycle")
    private @Nullable Output<String> pricingCycle;

    /**
     * @return Billing cycle. Value:
     * - Month: monthly billing
     * - Hour: hourly billing
     * &gt; **NOTE:**  Subscription instances (PrePaid) only supports Month. PayAsYouGo instances (PostPaid) only supports Hour. The Shared type is automatically set to Hour without specifying it.
     * 
     */
    public Optional<Output<String>> pricingCycle() {
        return Optional.ofNullable(this.pricingCycle);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * Change matching type. Value:
     * - UPGRADE: UPGRADE
     * - DOWNGRADE: Downgrading
     * &gt; **NOTE:** The upgrade specification cannot be less than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is larger than the original specification. The downgrading specification cannot be greater than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is smaller than the original specification.
     * 
     */
    @Import(name="scaleType")
    private @Nullable Output<String> scaleType;

    /**
     * @return Change matching type. Value:
     * - UPGRADE: UPGRADE
     * - DOWNGRADE: Downgrading
     * &gt; **NOTE:** The upgrade specification cannot be less than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is larger than the original specification. The downgrading specification cannot be greater than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is smaller than the original specification.
     * 
     */
    public Optional<Output<String>> scaleType() {
        return Optional.ofNullable(this.scaleType);
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The standard storage space of the instance. Unit: GB.
     * &gt; **NOTE:**  PayAsYouGo instances (PostPaid) ignore this parameter.
     * 
     */
    @Import(name="storageSize")
    private @Nullable Output<Integer> storageSize;

    /**
     * @return The standard storage space of the instance. Unit: GB.
     * &gt; **NOTE:**  PayAsYouGo instances (PostPaid) ignore this parameter.
     * 
     */
    public Optional<Output<Integer>> storageSize() {
        return Optional.ofNullable(this.storageSize);
    }

    /**
     * Instance tag.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Instance tag.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The zone Id. Refer to &#34;Instructions for Use&#34;.
     * 
     */
    @Import(name="zoneId", required=true)
    private Output<String> zoneId;

    /**
     * @return The zone Id. Refer to &#34;Instructions for Use&#34;.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    private InstanceArgs() {}

    private InstanceArgs(InstanceArgs $) {
        this.autoPay = $.autoPay;
        this.coldStorageSize = $.coldStorageSize;
        this.cpu = $.cpu;
        this.duration = $.duration;
        this.endpoints = $.endpoints;
        this.gatewayCount = $.gatewayCount;
        this.initialDatabases = $.initialDatabases;
        this.instanceName = $.instanceName;
        this.instanceType = $.instanceType;
        this.leaderInstanceId = $.leaderInstanceId;
        this.paymentType = $.paymentType;
        this.pricingCycle = $.pricingCycle;
        this.resourceGroupId = $.resourceGroupId;
        this.scaleType = $.scaleType;
        this.status = $.status;
        this.storageSize = $.storageSize;
        this.tags = $.tags;
        this.zoneId = $.zoneId;
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
         * @param autoPay Whether to pay automatically. The default value is true. Value:
         * - true: automatic payment
         * - false: only generate orders, not pay
         * &gt; **NOTE:**  The default value is true. If the balance of your payment method is insufficient, you can set the parameter AutoPay to false, and an unpaid order will be generated. You can log in to the user Center to pay by yourself.
         * 
         * @return builder
         * 
         */
        public Builder autoPay(@Nullable Output<Boolean> autoPay) {
            $.autoPay = autoPay;
            return this;
        }

        /**
         * @param autoPay Whether to pay automatically. The default value is true. Value:
         * - true: automatic payment
         * - false: only generate orders, not pay
         * &gt; **NOTE:**  The default value is true. If the balance of your payment method is insufficient, you can set the parameter AutoPay to false, and an unpaid order will be generated. You can log in to the user Center to pay by yourself.
         * 
         * @return builder
         * 
         */
        public Builder autoPay(Boolean autoPay) {
            return autoPay(Output.of(autoPay));
        }

        /**
         * @param coldStorageSize Instance low-frequency storage space. Unit: GB.
         * &gt; **NOTE:**  PayAsYouGo (PostPaid) instances ignore this parameter.
         * 
         * @return builder
         * 
         */
        public Builder coldStorageSize(@Nullable Output<Integer> coldStorageSize) {
            $.coldStorageSize = coldStorageSize;
            return this;
        }

        /**
         * @param coldStorageSize Instance low-frequency storage space. Unit: GB.
         * &gt; **NOTE:**  PayAsYouGo (PostPaid) instances ignore this parameter.
         * 
         * @return builder
         * 
         */
        public Builder coldStorageSize(Integer coldStorageSize) {
            return coldStorageSize(Output.of(coldStorageSize));
        }

        /**
         * @param cpu Instance specifications. Value:
         * - 8 cores 32 GB (number of compute nodes: 1)
         * - 16 cores 64 GB (number of compute nodes: 1)
         * - 32 core 128 GB (number of compute nodes: 2)
         * - 64 core 256 GB (number of compute nodes: 4)
         * - 96 core 384 GB (number of computing nodes: 6)
         * - 128 core 512 GB (number of compute nodes: 8)
         * &gt; **NOTE:** Just fill in the audit number. Please submit a work order application for purchasing 1024 or above specifications. Shared instance types do not need to specify specifications. The specification of - 8 core 32GB (number of computing nodes: 1) is only for experience use and cannot be used for production.
         * 
         * @return builder
         * 
         */
        public Builder cpu(@Nullable Output<Integer> cpu) {
            $.cpu = cpu;
            return this;
        }

        /**
         * @param cpu Instance specifications. Value:
         * - 8 cores 32 GB (number of compute nodes: 1)
         * - 16 cores 64 GB (number of compute nodes: 1)
         * - 32 core 128 GB (number of compute nodes: 2)
         * - 64 core 256 GB (number of compute nodes: 4)
         * - 96 core 384 GB (number of computing nodes: 6)
         * - 128 core 512 GB (number of compute nodes: 8)
         * &gt; **NOTE:** Just fill in the audit number. Please submit a work order application for purchasing 1024 or above specifications. Shared instance types do not need to specify specifications. The specification of - 8 core 32GB (number of computing nodes: 1) is only for experience use and cannot be used for production.
         * 
         * @return builder
         * 
         */
        public Builder cpu(Integer cpu) {
            return cpu(Output.of(cpu));
        }

        /**
         * @param duration The buying cycle. Buy for 2 months. If the Payment type is PayAsYouGo (PostPaid), you do not need to specify it.
         * 
         * @return builder
         * 
         */
        public Builder duration(@Nullable Output<Integer> duration) {
            $.duration = duration;
            return this;
        }

        /**
         * @param duration The buying cycle. Buy for 2 months. If the Payment type is PayAsYouGo (PostPaid), you do not need to specify it.
         * 
         * @return builder
         * 
         */
        public Builder duration(Integer duration) {
            return duration(Output.of(duration));
        }

        /**
         * @param endpoints List of domain names. See `endpoints` below.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(@Nullable Output<List<InstanceEndpointArgs>> endpoints) {
            $.endpoints = endpoints;
            return this;
        }

        /**
         * @param endpoints List of domain names. See `endpoints` below.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(List<InstanceEndpointArgs> endpoints) {
            return endpoints(Output.of(endpoints));
        }

        /**
         * @param endpoints List of domain names. See `endpoints` below.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(InstanceEndpointArgs... endpoints) {
            return endpoints(List.of(endpoints));
        }

        /**
         * @param gatewayCount Number of gateway nodes.
         * 
         * @return builder
         * 
         */
        public Builder gatewayCount(@Nullable Output<Integer> gatewayCount) {
            $.gatewayCount = gatewayCount;
            return this;
        }

        /**
         * @param gatewayCount Number of gateway nodes.
         * 
         * @return builder
         * 
         */
        public Builder gatewayCount(Integer gatewayCount) {
            return gatewayCount(Output.of(gatewayCount));
        }

        /**
         * @param initialDatabases Initialize the database and split multiple database names &#34;,&#34;.
         * 
         * @return builder
         * 
         */
        public Builder initialDatabases(@Nullable Output<String> initialDatabases) {
            $.initialDatabases = initialDatabases;
            return this;
        }

        /**
         * @param initialDatabases Initialize the database and split multiple database names &#34;,&#34;.
         * 
         * @return builder
         * 
         */
        public Builder initialDatabases(String initialDatabases) {
            return initialDatabases(Output.of(initialDatabases));
        }

        /**
         * @param instanceName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        /**
         * @param instanceType The instance type. Value:
         * - Standard: Universal.
         * - Follower: Read-only slave instance.
         * - Warehouse: calculation group type.
         * - Shared: Shared.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType The instance type. Value:
         * - Standard: Universal.
         * - Follower: Read-only slave instance.
         * - Warehouse: calculation group type.
         * - Shared: Shared.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param leaderInstanceId The ID of the primary instance.
         * 
         * @return builder
         * 
         */
        public Builder leaderInstanceId(@Nullable Output<String> leaderInstanceId) {
            $.leaderInstanceId = leaderInstanceId;
            return this;
        }

        /**
         * @param leaderInstanceId The ID of the primary instance.
         * 
         * @return builder
         * 
         */
        public Builder leaderInstanceId(String leaderInstanceId) {
            return leaderInstanceId(Output.of(leaderInstanceId));
        }

        /**
         * @param paymentType The payment type of the resource.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The payment type of the resource.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        /**
         * @param pricingCycle Billing cycle. Value:
         * - Month: monthly billing
         * - Hour: hourly billing
         * &gt; **NOTE:**  Subscription instances (PrePaid) only supports Month. PayAsYouGo instances (PostPaid) only supports Hour. The Shared type is automatically set to Hour without specifying it.
         * 
         * @return builder
         * 
         */
        public Builder pricingCycle(@Nullable Output<String> pricingCycle) {
            $.pricingCycle = pricingCycle;
            return this;
        }

        /**
         * @param pricingCycle Billing cycle. Value:
         * - Month: monthly billing
         * - Hour: hourly billing
         * &gt; **NOTE:**  Subscription instances (PrePaid) only supports Month. PayAsYouGo instances (PostPaid) only supports Hour. The Shared type is automatically set to Hour without specifying it.
         * 
         * @return builder
         * 
         */
        public Builder pricingCycle(String pricingCycle) {
            return pricingCycle(Output.of(pricingCycle));
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param scaleType Change matching type. Value:
         * - UPGRADE: UPGRADE
         * - DOWNGRADE: Downgrading
         * &gt; **NOTE:** The upgrade specification cannot be less than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is larger than the original specification. The downgrading specification cannot be greater than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is smaller than the original specification.
         * 
         * @return builder
         * 
         */
        public Builder scaleType(@Nullable Output<String> scaleType) {
            $.scaleType = scaleType;
            return this;
        }

        /**
         * @param scaleType Change matching type. Value:
         * - UPGRADE: UPGRADE
         * - DOWNGRADE: Downgrading
         * &gt; **NOTE:** The upgrade specification cannot be less than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is larger than the original specification. The downgrading specification cannot be greater than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is smaller than the original specification.
         * 
         * @return builder
         * 
         */
        public Builder scaleType(String scaleType) {
            return scaleType(Output.of(scaleType));
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param storageSize The standard storage space of the instance. Unit: GB.
         * &gt; **NOTE:**  PayAsYouGo instances (PostPaid) ignore this parameter.
         * 
         * @return builder
         * 
         */
        public Builder storageSize(@Nullable Output<Integer> storageSize) {
            $.storageSize = storageSize;
            return this;
        }

        /**
         * @param storageSize The standard storage space of the instance. Unit: GB.
         * &gt; **NOTE:**  PayAsYouGo instances (PostPaid) ignore this parameter.
         * 
         * @return builder
         * 
         */
        public Builder storageSize(Integer storageSize) {
            return storageSize(Output.of(storageSize));
        }

        /**
         * @param tags Instance tag.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Instance tag.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param zoneId The zone Id. Refer to &#34;Instructions for Use&#34;.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The zone Id. Refer to &#34;Instructions for Use&#34;.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public InstanceArgs build() {
            if ($.instanceName == null) {
                throw new MissingRequiredPropertyException("InstanceArgs", "instanceName");
            }
            if ($.instanceType == null) {
                throw new MissingRequiredPropertyException("InstanceArgs", "instanceType");
            }
            if ($.paymentType == null) {
                throw new MissingRequiredPropertyException("InstanceArgs", "paymentType");
            }
            if ($.zoneId == null) {
                throw new MissingRequiredPropertyException("InstanceArgs", "zoneId");
            }
            return $;
        }
    }

}
