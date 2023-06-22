// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class Ipv6InternetBandwidthState extends com.pulumi.resources.ResourceArgs {

    public static final Ipv6InternetBandwidthState Empty = new Ipv6InternetBandwidthState();

    /**
     * The amount of Internet bandwidth resources of the IPv6 address, Unit: `Mbit/s`. Valid values: `1` to `5000`. **NOTE:** If `internet_charge_type` is set to `PayByTraffic`, the amount of Internet bandwidth resources of the IPv6 address is limited by the specification of the IPv6 gateway. `Small` (default): specifies the Free edition and the Internet bandwidth is from `1` to `500` Mbit/s. `Medium`: specifies the Medium edition and the Internet bandwidth is from `1` to `1000` Mbit/s. `Large`: specifies the Large edition and the Internet bandwidth is from `1` to `2000` Mbit/s.
     * 
     */
    @Import(name="bandwidth")
    private @Nullable Output<Integer> bandwidth;

    /**
     * @return The amount of Internet bandwidth resources of the IPv6 address, Unit: `Mbit/s`. Valid values: `1` to `5000`. **NOTE:** If `internet_charge_type` is set to `PayByTraffic`, the amount of Internet bandwidth resources of the IPv6 address is limited by the specification of the IPv6 gateway. `Small` (default): specifies the Free edition and the Internet bandwidth is from `1` to `500` Mbit/s. `Medium`: specifies the Medium edition and the Internet bandwidth is from `1` to `1000` Mbit/s. `Large`: specifies the Large edition and the Internet bandwidth is from `1` to `2000` Mbit/s.
     * 
     */
    public Optional<Output<Integer>> bandwidth() {
        return Optional.ofNullable(this.bandwidth);
    }

    /**
     * The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: `PayByBandwidth`, `PayByTraffic`.
     * 
     */
    @Import(name="internetChargeType")
    private @Nullable Output<String> internetChargeType;

    /**
     * @return The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: `PayByBandwidth`, `PayByTraffic`.
     * 
     */
    public Optional<Output<String>> internetChargeType() {
        return Optional.ofNullable(this.internetChargeType);
    }

    /**
     * The ID of the IPv6 address instance.
     * 
     */
    @Import(name="ipv6AddressId")
    private @Nullable Output<String> ipv6AddressId;

    /**
     * @return The ID of the IPv6 address instance.
     * 
     */
    public Optional<Output<String>> ipv6AddressId() {
        return Optional.ofNullable(this.ipv6AddressId);
    }

    /**
     * The ID of the IPv6 gateway to which the IPv6 address belongs.
     * 
     */
    @Import(name="ipv6GatewayId")
    private @Nullable Output<String> ipv6GatewayId;

    /**
     * @return The ID of the IPv6 gateway to which the IPv6 address belongs.
     * 
     */
    public Optional<Output<String>> ipv6GatewayId() {
        return Optional.ofNullable(this.ipv6GatewayId);
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

    private Ipv6InternetBandwidthState() {}

    private Ipv6InternetBandwidthState(Ipv6InternetBandwidthState $) {
        this.bandwidth = $.bandwidth;
        this.internetChargeType = $.internetChargeType;
        this.ipv6AddressId = $.ipv6AddressId;
        this.ipv6GatewayId = $.ipv6GatewayId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Ipv6InternetBandwidthState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Ipv6InternetBandwidthState $;

        public Builder() {
            $ = new Ipv6InternetBandwidthState();
        }

        public Builder(Ipv6InternetBandwidthState defaults) {
            $ = new Ipv6InternetBandwidthState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bandwidth The amount of Internet bandwidth resources of the IPv6 address, Unit: `Mbit/s`. Valid values: `1` to `5000`. **NOTE:** If `internet_charge_type` is set to `PayByTraffic`, the amount of Internet bandwidth resources of the IPv6 address is limited by the specification of the IPv6 gateway. `Small` (default): specifies the Free edition and the Internet bandwidth is from `1` to `500` Mbit/s. `Medium`: specifies the Medium edition and the Internet bandwidth is from `1` to `1000` Mbit/s. `Large`: specifies the Large edition and the Internet bandwidth is from `1` to `2000` Mbit/s.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(@Nullable Output<Integer> bandwidth) {
            $.bandwidth = bandwidth;
            return this;
        }

        /**
         * @param bandwidth The amount of Internet bandwidth resources of the IPv6 address, Unit: `Mbit/s`. Valid values: `1` to `5000`. **NOTE:** If `internet_charge_type` is set to `PayByTraffic`, the amount of Internet bandwidth resources of the IPv6 address is limited by the specification of the IPv6 gateway. `Small` (default): specifies the Free edition and the Internet bandwidth is from `1` to `500` Mbit/s. `Medium`: specifies the Medium edition and the Internet bandwidth is from `1` to `1000` Mbit/s. `Large`: specifies the Large edition and the Internet bandwidth is from `1` to `2000` Mbit/s.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(Integer bandwidth) {
            return bandwidth(Output.of(bandwidth));
        }

        /**
         * @param internetChargeType The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: `PayByBandwidth`, `PayByTraffic`.
         * 
         * @return builder
         * 
         */
        public Builder internetChargeType(@Nullable Output<String> internetChargeType) {
            $.internetChargeType = internetChargeType;
            return this;
        }

        /**
         * @param internetChargeType The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: `PayByBandwidth`, `PayByTraffic`.
         * 
         * @return builder
         * 
         */
        public Builder internetChargeType(String internetChargeType) {
            return internetChargeType(Output.of(internetChargeType));
        }

        /**
         * @param ipv6AddressId The ID of the IPv6 address instance.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressId(@Nullable Output<String> ipv6AddressId) {
            $.ipv6AddressId = ipv6AddressId;
            return this;
        }

        /**
         * @param ipv6AddressId The ID of the IPv6 address instance.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressId(String ipv6AddressId) {
            return ipv6AddressId(Output.of(ipv6AddressId));
        }

        /**
         * @param ipv6GatewayId The ID of the IPv6 gateway to which the IPv6 address belongs.
         * 
         * @return builder
         * 
         */
        public Builder ipv6GatewayId(@Nullable Output<String> ipv6GatewayId) {
            $.ipv6GatewayId = ipv6GatewayId;
            return this;
        }

        /**
         * @param ipv6GatewayId The ID of the IPv6 gateway to which the IPv6 address belongs.
         * 
         * @return builder
         * 
         */
        public Builder ipv6GatewayId(String ipv6GatewayId) {
            return ipv6GatewayId(Output.of(ipv6GatewayId));
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

        public Ipv6InternetBandwidthState build() {
            return $;
        }
    }

}
