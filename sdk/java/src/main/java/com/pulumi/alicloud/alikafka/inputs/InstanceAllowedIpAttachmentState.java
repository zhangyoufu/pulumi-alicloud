// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alikafka.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceAllowedIpAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final InstanceAllowedIpAttachmentState Empty = new InstanceAllowedIpAttachmentState();

    /**
     * The IP address whitelist. It can be a CIDR block.
     * 
     */
    @Import(name="allowedIp")
    private @Nullable Output<String> allowedIp;

    /**
     * @return The IP address whitelist. It can be a CIDR block.
     * 
     */
    public Optional<Output<String>> allowedIp() {
        return Optional.ofNullable(this.allowedIp);
    }

    /**
     * The type of the whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowed_type` can be set to `internet`.
     * 
     */
    @Import(name="allowedType")
    private @Nullable Output<String> allowedType;

    /**
     * @return The type of the whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowed_type` can be set to `internet`.
     * 
     */
    public Optional<Output<String>> allowedType() {
        return Optional.ofNullable(this.allowedType);
    }

    /**
     * The ID of the instance.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The Port range. Valid Value: `9092/9092`, `9093/9093`, `9094/9094`, `9095/9095`. **NOTE:** From version 1.179.0, `port_range` can be set to `9093/9093`. From version 1.219.0, `port_range` can be set to `9094/9094`, `9095/9095`.
     * - `9092/9092`: The port range for access from virtual private clouds (VPCs) by using the default endpoint.
     * - `9093/9093`: The port range for access from the Internet.
     * - `9094/9094`: The port range for access from VPCs by using the Simple Authentication and Security Layer (SASL) endpoint.
     * - `9095/9095`: The port range for access from VPCs by using the Secure Sockets Layer (SSL) endpoint.
     * 
     */
    @Import(name="portRange")
    private @Nullable Output<String> portRange;

    /**
     * @return The Port range. Valid Value: `9092/9092`, `9093/9093`, `9094/9094`, `9095/9095`. **NOTE:** From version 1.179.0, `port_range` can be set to `9093/9093`. From version 1.219.0, `port_range` can be set to `9094/9094`, `9095/9095`.
     * - `9092/9092`: The port range for access from virtual private clouds (VPCs) by using the default endpoint.
     * - `9093/9093`: The port range for access from the Internet.
     * - `9094/9094`: The port range for access from VPCs by using the Simple Authentication and Security Layer (SASL) endpoint.
     * - `9095/9095`: The port range for access from VPCs by using the Secure Sockets Layer (SSL) endpoint.
     * 
     */
    public Optional<Output<String>> portRange() {
        return Optional.ofNullable(this.portRange);
    }

    private InstanceAllowedIpAttachmentState() {}

    private InstanceAllowedIpAttachmentState(InstanceAllowedIpAttachmentState $) {
        this.allowedIp = $.allowedIp;
        this.allowedType = $.allowedType;
        this.instanceId = $.instanceId;
        this.portRange = $.portRange;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceAllowedIpAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceAllowedIpAttachmentState $;

        public Builder() {
            $ = new InstanceAllowedIpAttachmentState();
        }

        public Builder(InstanceAllowedIpAttachmentState defaults) {
            $ = new InstanceAllowedIpAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedIp The IP address whitelist. It can be a CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder allowedIp(@Nullable Output<String> allowedIp) {
            $.allowedIp = allowedIp;
            return this;
        }

        /**
         * @param allowedIp The IP address whitelist. It can be a CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder allowedIp(String allowedIp) {
            return allowedIp(Output.of(allowedIp));
        }

        /**
         * @param allowedType The type of the whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowed_type` can be set to `internet`.
         * 
         * @return builder
         * 
         */
        public Builder allowedType(@Nullable Output<String> allowedType) {
            $.allowedType = allowedType;
            return this;
        }

        /**
         * @param allowedType The type of the whitelist. Valid Value: `vpc`, `internet`. **NOTE:** From version 1.179.0, `allowed_type` can be set to `internet`.
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
        public Builder instanceId(@Nullable Output<String> instanceId) {
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
         * @param portRange The Port range. Valid Value: `9092/9092`, `9093/9093`, `9094/9094`, `9095/9095`. **NOTE:** From version 1.179.0, `port_range` can be set to `9093/9093`. From version 1.219.0, `port_range` can be set to `9094/9094`, `9095/9095`.
         * - `9092/9092`: The port range for access from virtual private clouds (VPCs) by using the default endpoint.
         * - `9093/9093`: The port range for access from the Internet.
         * - `9094/9094`: The port range for access from VPCs by using the Simple Authentication and Security Layer (SASL) endpoint.
         * - `9095/9095`: The port range for access from VPCs by using the Secure Sockets Layer (SSL) endpoint.
         * 
         * @return builder
         * 
         */
        public Builder portRange(@Nullable Output<String> portRange) {
            $.portRange = portRange;
            return this;
        }

        /**
         * @param portRange The Port range. Valid Value: `9092/9092`, `9093/9093`, `9094/9094`, `9095/9095`. **NOTE:** From version 1.179.0, `port_range` can be set to `9093/9093`. From version 1.219.0, `port_range` can be set to `9094/9094`, `9095/9095`.
         * - `9092/9092`: The port range for access from virtual private clouds (VPCs) by using the default endpoint.
         * - `9093/9093`: The port range for access from the Internet.
         * - `9094/9094`: The port range for access from VPCs by using the Simple Authentication and Security Layer (SASL) endpoint.
         * - `9095/9095`: The port range for access from VPCs by using the Secure Sockets Layer (SSL) endpoint.
         * 
         * @return builder
         * 
         */
        public Builder portRange(String portRange) {
            return portRange(Output.of(portRange));
        }

        public InstanceAllowedIpAttachmentState build() {
            return $;
        }
    }

}
