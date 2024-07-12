// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ens;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EipInstanceAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final EipInstanceAttachmentArgs Empty = new EipInstanceAttachmentArgs();

    /**
     * The first ID of the resource
     * 
     */
    @Import(name="allocationId", required=true)
    private Output<String> allocationId;

    /**
     * @return The first ID of the resource
     * 
     */
    public Output<String> allocationId() {
        return this.allocationId;
    }

    /**
     * Instance ID
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return Instance ID
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * The type of the EIP instance. Value:
     * - `Nat`:NAT gateway.
     * - `SlbInstance`: Server Load Balancer (ELB).
     * - `NetworkInterface`: Secondary ENI.
     * - `EnsInstance` (default): The ENS instance.
     * 
     */
    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    /**
     * @return The type of the EIP instance. Value:
     * - `Nat`:NAT gateway.
     * - `SlbInstance`: Server Load Balancer (ELB).
     * - `NetworkInterface`: Secondary ENI.
     * - `EnsInstance` (default): The ENS instance.
     * 
     */
    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    /**
     * Indicates whether the EIP is a backup EIP. Value:
     * - true: Spare.
     * - false: not standby.
     * 
     */
    @Import(name="standby")
    private @Nullable Output<Boolean> standby;

    /**
     * @return Indicates whether the EIP is a backup EIP. Value:
     * - true: Spare.
     * - false: not standby.
     * 
     */
    public Optional<Output<Boolean>> standby() {
        return Optional.ofNullable(this.standby);
    }

    private EipInstanceAttachmentArgs() {}

    private EipInstanceAttachmentArgs(EipInstanceAttachmentArgs $) {
        this.allocationId = $.allocationId;
        this.instanceId = $.instanceId;
        this.instanceType = $.instanceType;
        this.standby = $.standby;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EipInstanceAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EipInstanceAttachmentArgs $;

        public Builder() {
            $ = new EipInstanceAttachmentArgs();
        }

        public Builder(EipInstanceAttachmentArgs defaults) {
            $ = new EipInstanceAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allocationId The first ID of the resource
         * 
         * @return builder
         * 
         */
        public Builder allocationId(Output<String> allocationId) {
            $.allocationId = allocationId;
            return this;
        }

        /**
         * @param allocationId The first ID of the resource
         * 
         * @return builder
         * 
         */
        public Builder allocationId(String allocationId) {
            return allocationId(Output.of(allocationId));
        }

        /**
         * @param instanceId Instance ID
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Instance ID
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param instanceType The type of the EIP instance. Value:
         * - `Nat`:NAT gateway.
         * - `SlbInstance`: Server Load Balancer (ELB).
         * - `NetworkInterface`: Secondary ENI.
         * - `EnsInstance` (default): The ENS instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType The type of the EIP instance. Value:
         * - `Nat`:NAT gateway.
         * - `SlbInstance`: Server Load Balancer (ELB).
         * - `NetworkInterface`: Secondary ENI.
         * - `EnsInstance` (default): The ENS instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param standby Indicates whether the EIP is a backup EIP. Value:
         * - true: Spare.
         * - false: not standby.
         * 
         * @return builder
         * 
         */
        public Builder standby(@Nullable Output<Boolean> standby) {
            $.standby = standby;
            return this;
        }

        /**
         * @param standby Indicates whether the EIP is a backup EIP. Value:
         * - true: Spare.
         * - false: not standby.
         * 
         * @return builder
         * 
         */
        public Builder standby(Boolean standby) {
            return standby(Output.of(standby));
        }

        public EipInstanceAttachmentArgs build() {
            if ($.allocationId == null) {
                throw new MissingRequiredPropertyException("EipInstanceAttachmentArgs", "allocationId");
            }
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("EipInstanceAttachmentArgs", "instanceId");
            }
            return $;
        }
    }

}
