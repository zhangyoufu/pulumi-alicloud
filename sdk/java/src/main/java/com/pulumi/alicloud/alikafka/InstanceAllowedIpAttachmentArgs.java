// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alikafka;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class InstanceAllowedIpAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceAllowedIpAttachmentArgs Empty = new InstanceAllowedIpAttachmentArgs();

    /**
     * The allowed ip. It can be a CIDR block.
     * 
     */
    @Import(name="allowedIp", required=true)
    private Output<String> allowedIp;

    /**
     * @return The allowed ip. It can be a CIDR block.
     * 
     */
    public Output<String> allowedIp() {
        return this.allowedIp;
    }

    /**
     * The type of whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowed_type` can be set to `internet`.
     * 
     */
    @Import(name="allowedType", required=true)
    private Output<String> allowedType;

    /**
     * @return The type of whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowed_type` can be set to `internet`.
     * 
     */
    public Output<String> allowedType() {
        return this.allowedType;
    }

    /**
     * The ID of the instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * The Port range.  Valid Value: `9092/9092`, `9093/9093`. **NOTE:** From version 1.179.0, `port_range` can be set to `9093/9093`.
     * - `9092/9092`: port range for a VPC whitelist.
     * - `9093/9093`: port range for an Internet whitelist.
     * 
     */
    @Import(name="portRange", required=true)
    private Output<String> portRange;

    /**
     * @return The Port range.  Valid Value: `9092/9092`, `9093/9093`. **NOTE:** From version 1.179.0, `port_range` can be set to `9093/9093`.
     * - `9092/9092`: port range for a VPC whitelist.
     * - `9093/9093`: port range for an Internet whitelist.
     * 
     */
    public Output<String> portRange() {
        return this.portRange;
    }

    private InstanceAllowedIpAttachmentArgs() {}

    private InstanceAllowedIpAttachmentArgs(InstanceAllowedIpAttachmentArgs $) {
        this.allowedIp = $.allowedIp;
        this.allowedType = $.allowedType;
        this.instanceId = $.instanceId;
        this.portRange = $.portRange;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceAllowedIpAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceAllowedIpAttachmentArgs $;

        public Builder() {
            $ = new InstanceAllowedIpAttachmentArgs();
        }

        public Builder(InstanceAllowedIpAttachmentArgs defaults) {
            $ = new InstanceAllowedIpAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedIp The allowed ip. It can be a CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder allowedIp(Output<String> allowedIp) {
            $.allowedIp = allowedIp;
            return this;
        }

        /**
         * @param allowedIp The allowed ip. It can be a CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder allowedIp(String allowedIp) {
            return allowedIp(Output.of(allowedIp));
        }

        /**
         * @param allowedType The type of whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowed_type` can be set to `internet`.
         * 
         * @return builder
         * 
         */
        public Builder allowedType(Output<String> allowedType) {
            $.allowedType = allowedType;
            return this;
        }

        /**
         * @param allowedType The type of whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowed_type` can be set to `internet`.
         * 
         * @return builder
         * 
         */
        public Builder allowedType(String allowedType) {
            return allowedType(Output.of(allowedType));
        }

        /**
         * @param instanceId The ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param portRange The Port range.  Valid Value: `9092/9092`, `9093/9093`. **NOTE:** From version 1.179.0, `port_range` can be set to `9093/9093`.
         * - `9092/9092`: port range for a VPC whitelist.
         * - `9093/9093`: port range for an Internet whitelist.
         * 
         * @return builder
         * 
         */
        public Builder portRange(Output<String> portRange) {
            $.portRange = portRange;
            return this;
        }

        /**
         * @param portRange The Port range.  Valid Value: `9092/9092`, `9093/9093`. **NOTE:** From version 1.179.0, `port_range` can be set to `9093/9093`.
         * - `9092/9092`: port range for a VPC whitelist.
         * - `9093/9093`: port range for an Internet whitelist.
         * 
         * @return builder
         * 
         */
        public Builder portRange(String portRange) {
            return portRange(Output.of(portRange));
        }

        public InstanceAllowedIpAttachmentArgs build() {
            $.allowedIp = Objects.requireNonNull($.allowedIp, "expected parameter 'allowedIp' to be non-null");
            $.allowedType = Objects.requireNonNull($.allowedType, "expected parameter 'allowedType' to be non-null");
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            $.portRange = Objects.requireNonNull($.portRange, "expected parameter 'portRange' to be non-null");
            return $;
        }
    }

}
