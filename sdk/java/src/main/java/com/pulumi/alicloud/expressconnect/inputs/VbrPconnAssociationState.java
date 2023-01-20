// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.expressconnect.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VbrPconnAssociationState extends com.pulumi.resources.ResourceArgs {

    public static final VbrPconnAssociationState Empty = new VbrPconnAssociationState();

    /**
     * The circuit code provided by the operator for the physical connection.
     * 
     */
    @Import(name="circuitCode")
    private @Nullable Output<String> circuitCode;

    /**
     * @return The circuit code provided by the operator for the physical connection.
     * 
     */
    public Optional<Output<String>> circuitCode() {
        return Optional.ofNullable(this.circuitCode);
    }

    /**
     * Whether IPv6 is enabled. Value:
     * - **true**: on.
     * - **false** (default): Off.
     * 
     */
    @Import(name="enableIpv6")
    private @Nullable Output<Boolean> enableIpv6;

    /**
     * @return Whether IPv6 is enabled. Value:
     * - **true**: on.
     * - **false** (default): Off.
     * 
     */
    public Optional<Output<Boolean>> enableIpv6() {
        return Optional.ofNullable(this.enableIpv6);
    }

    /**
     * The Alibaba cloud IP address of the VBR instance.
     * 
     */
    @Import(name="localGatewayIp")
    private @Nullable Output<String> localGatewayIp;

    /**
     * @return The Alibaba cloud IP address of the VBR instance.
     * 
     */
    public Optional<Output<String>> localGatewayIp() {
        return Optional.ofNullable(this.localGatewayIp);
    }

    /**
     * The IPv6 address on the Alibaba Cloud side of the VBR instance.
     * 
     */
    @Import(name="localIpv6GatewayIp")
    private @Nullable Output<String> localIpv6GatewayIp;

    /**
     * @return The IPv6 address on the Alibaba Cloud side of the VBR instance.
     * 
     */
    public Optional<Output<String>> localIpv6GatewayIp() {
        return Optional.ofNullable(this.localIpv6GatewayIp);
    }

    /**
     * The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
     * 
     */
    @Import(name="peerGatewayIp")
    private @Nullable Output<String> peerGatewayIp;

    /**
     * @return The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
     * 
     */
    public Optional<Output<String>> peerGatewayIp() {
        return Optional.ofNullable(this.peerGatewayIp);
    }

    /**
     * The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
     * 
     */
    @Import(name="peerIpv6GatewayIp")
    private @Nullable Output<String> peerIpv6GatewayIp;

    /**
     * @return The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
     * 
     */
    public Optional<Output<String>> peerIpv6GatewayIp() {
        return Optional.ofNullable(this.peerIpv6GatewayIp);
    }

    /**
     * The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
     * 
     */
    @Import(name="peeringIpv6SubnetMask")
    private @Nullable Output<String> peeringIpv6SubnetMask;

    /**
     * @return The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
     * 
     */
    public Optional<Output<String>> peeringIpv6SubnetMask() {
        return Optional.ofNullable(this.peeringIpv6SubnetMask);
    }

    /**
     * The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
     * 
     */
    @Import(name="peeringSubnetMask")
    private @Nullable Output<String> peeringSubnetMask;

    /**
     * @return The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
     * 
     */
    public Optional<Output<String>> peeringSubnetMask() {
        return Optional.ofNullable(this.peeringSubnetMask);
    }

    /**
     * The ID of the leased line instance.
     * 
     */
    @Import(name="physicalConnectionId")
    private @Nullable Output<String> physicalConnectionId;

    /**
     * @return The ID of the leased line instance.
     * 
     */
    public Optional<Output<String>> physicalConnectionId() {
        return Optional.ofNullable(this.physicalConnectionId);
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
     * The ID of the VBR instance.
     * 
     */
    @Import(name="vbrId")
    private @Nullable Output<String> vbrId;

    /**
     * @return The ID of the VBR instance.
     * 
     */
    public Optional<Output<String>> vbrId() {
        return Optional.ofNullable(this.vbrId);
    }

    /**
     * VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
     * 
     */
    @Import(name="vlanId")
    private @Nullable Output<Integer> vlanId;

    /**
     * @return VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
     * 
     */
    public Optional<Output<Integer>> vlanId() {
        return Optional.ofNullable(this.vlanId);
    }

    private VbrPconnAssociationState() {}

    private VbrPconnAssociationState(VbrPconnAssociationState $) {
        this.circuitCode = $.circuitCode;
        this.enableIpv6 = $.enableIpv6;
        this.localGatewayIp = $.localGatewayIp;
        this.localIpv6GatewayIp = $.localIpv6GatewayIp;
        this.peerGatewayIp = $.peerGatewayIp;
        this.peerIpv6GatewayIp = $.peerIpv6GatewayIp;
        this.peeringIpv6SubnetMask = $.peeringIpv6SubnetMask;
        this.peeringSubnetMask = $.peeringSubnetMask;
        this.physicalConnectionId = $.physicalConnectionId;
        this.status = $.status;
        this.vbrId = $.vbrId;
        this.vlanId = $.vlanId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VbrPconnAssociationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VbrPconnAssociationState $;

        public Builder() {
            $ = new VbrPconnAssociationState();
        }

        public Builder(VbrPconnAssociationState defaults) {
            $ = new VbrPconnAssociationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param circuitCode The circuit code provided by the operator for the physical connection.
         * 
         * @return builder
         * 
         */
        public Builder circuitCode(@Nullable Output<String> circuitCode) {
            $.circuitCode = circuitCode;
            return this;
        }

        /**
         * @param circuitCode The circuit code provided by the operator for the physical connection.
         * 
         * @return builder
         * 
         */
        public Builder circuitCode(String circuitCode) {
            return circuitCode(Output.of(circuitCode));
        }

        /**
         * @param enableIpv6 Whether IPv6 is enabled. Value:
         * - **true**: on.
         * - **false** (default): Off.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(@Nullable Output<Boolean> enableIpv6) {
            $.enableIpv6 = enableIpv6;
            return this;
        }

        /**
         * @param enableIpv6 Whether IPv6 is enabled. Value:
         * - **true**: on.
         * - **false** (default): Off.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(Boolean enableIpv6) {
            return enableIpv6(Output.of(enableIpv6));
        }

        /**
         * @param localGatewayIp The Alibaba cloud IP address of the VBR instance.
         * 
         * @return builder
         * 
         */
        public Builder localGatewayIp(@Nullable Output<String> localGatewayIp) {
            $.localGatewayIp = localGatewayIp;
            return this;
        }

        /**
         * @param localGatewayIp The Alibaba cloud IP address of the VBR instance.
         * 
         * @return builder
         * 
         */
        public Builder localGatewayIp(String localGatewayIp) {
            return localGatewayIp(Output.of(localGatewayIp));
        }

        /**
         * @param localIpv6GatewayIp The IPv6 address on the Alibaba Cloud side of the VBR instance.
         * 
         * @return builder
         * 
         */
        public Builder localIpv6GatewayIp(@Nullable Output<String> localIpv6GatewayIp) {
            $.localIpv6GatewayIp = localIpv6GatewayIp;
            return this;
        }

        /**
         * @param localIpv6GatewayIp The IPv6 address on the Alibaba Cloud side of the VBR instance.
         * 
         * @return builder
         * 
         */
        public Builder localIpv6GatewayIp(String localIpv6GatewayIp) {
            return localIpv6GatewayIp(Output.of(localIpv6GatewayIp));
        }

        /**
         * @param peerGatewayIp The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
         * 
         * @return builder
         * 
         */
        public Builder peerGatewayIp(@Nullable Output<String> peerGatewayIp) {
            $.peerGatewayIp = peerGatewayIp;
            return this;
        }

        /**
         * @param peerGatewayIp The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
         * 
         * @return builder
         * 
         */
        public Builder peerGatewayIp(String peerGatewayIp) {
            return peerGatewayIp(Output.of(peerGatewayIp));
        }

        /**
         * @param peerIpv6GatewayIp The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
         * 
         * @return builder
         * 
         */
        public Builder peerIpv6GatewayIp(@Nullable Output<String> peerIpv6GatewayIp) {
            $.peerIpv6GatewayIp = peerIpv6GatewayIp;
            return this;
        }

        /**
         * @param peerIpv6GatewayIp The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
         * 
         * @return builder
         * 
         */
        public Builder peerIpv6GatewayIp(String peerIpv6GatewayIp) {
            return peerIpv6GatewayIp(Output.of(peerIpv6GatewayIp));
        }

        /**
         * @param peeringIpv6SubnetMask The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
         * 
         * @return builder
         * 
         */
        public Builder peeringIpv6SubnetMask(@Nullable Output<String> peeringIpv6SubnetMask) {
            $.peeringIpv6SubnetMask = peeringIpv6SubnetMask;
            return this;
        }

        /**
         * @param peeringIpv6SubnetMask The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
         * 
         * @return builder
         * 
         */
        public Builder peeringIpv6SubnetMask(String peeringIpv6SubnetMask) {
            return peeringIpv6SubnetMask(Output.of(peeringIpv6SubnetMask));
        }

        /**
         * @param peeringSubnetMask The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
         * 
         * @return builder
         * 
         */
        public Builder peeringSubnetMask(@Nullable Output<String> peeringSubnetMask) {
            $.peeringSubnetMask = peeringSubnetMask;
            return this;
        }

        /**
         * @param peeringSubnetMask The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
         * 
         * @return builder
         * 
         */
        public Builder peeringSubnetMask(String peeringSubnetMask) {
            return peeringSubnetMask(Output.of(peeringSubnetMask));
        }

        /**
         * @param physicalConnectionId The ID of the leased line instance.
         * 
         * @return builder
         * 
         */
        public Builder physicalConnectionId(@Nullable Output<String> physicalConnectionId) {
            $.physicalConnectionId = physicalConnectionId;
            return this;
        }

        /**
         * @param physicalConnectionId The ID of the leased line instance.
         * 
         * @return builder
         * 
         */
        public Builder physicalConnectionId(String physicalConnectionId) {
            return physicalConnectionId(Output.of(physicalConnectionId));
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
         * @param vbrId The ID of the VBR instance.
         * 
         * @return builder
         * 
         */
        public Builder vbrId(@Nullable Output<String> vbrId) {
            $.vbrId = vbrId;
            return this;
        }

        /**
         * @param vbrId The ID of the VBR instance.
         * 
         * @return builder
         * 
         */
        public Builder vbrId(String vbrId) {
            return vbrId(Output.of(vbrId));
        }

        /**
         * @param vlanId VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
         * 
         * @return builder
         * 
         */
        public Builder vlanId(@Nullable Output<Integer> vlanId) {
            $.vlanId = vlanId;
            return this;
        }

        /**
         * @param vlanId VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
         * 
         * @return builder
         * 
         */
        public Builder vlanId(Integer vlanId) {
            return vlanId(Output.of(vlanId));
        }

        public VbrPconnAssociationState build() {
            return $;
        }
    }

}
