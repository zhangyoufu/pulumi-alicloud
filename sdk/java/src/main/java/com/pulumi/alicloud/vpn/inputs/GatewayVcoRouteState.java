// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpn.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GatewayVcoRouteState extends com.pulumi.resources.ResourceArgs {

    public static final GatewayVcoRouteState Empty = new GatewayVcoRouteState();

    /**
     * The next hop of the destination route.
     * 
     */
    @Import(name="nextHop")
    private @Nullable Output<String> nextHop;

    /**
     * @return The next hop of the destination route.
     * 
     */
    public Optional<Output<String>> nextHop() {
        return Optional.ofNullable(this.nextHop);
    }

    /**
     * The destination network segment of the destination route.
     * 
     */
    @Import(name="routeDest")
    private @Nullable Output<String> routeDest;

    /**
     * @return The destination network segment of the destination route.
     * 
     */
    public Optional<Output<String>> routeDest() {
        return Optional.ofNullable(this.routeDest);
    }

    /**
     * The status of the vpn route entry.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the vpn route entry.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The id of the vpn attachment.
     * 
     */
    @Import(name="vpnConnectionId")
    private @Nullable Output<String> vpnConnectionId;

    /**
     * @return The id of the vpn attachment.
     * 
     */
    public Optional<Output<String>> vpnConnectionId() {
        return Optional.ofNullable(this.vpnConnectionId);
    }

    /**
     * The weight value of the destination route. Valid values: `0`, `100`.
     * 
     */
    @Import(name="weight")
    private @Nullable Output<Integer> weight;

    /**
     * @return The weight value of the destination route. Valid values: `0`, `100`.
     * 
     */
    public Optional<Output<Integer>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private GatewayVcoRouteState() {}

    private GatewayVcoRouteState(GatewayVcoRouteState $) {
        this.nextHop = $.nextHop;
        this.routeDest = $.routeDest;
        this.status = $.status;
        this.vpnConnectionId = $.vpnConnectionId;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GatewayVcoRouteState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GatewayVcoRouteState $;

        public Builder() {
            $ = new GatewayVcoRouteState();
        }

        public Builder(GatewayVcoRouteState defaults) {
            $ = new GatewayVcoRouteState(Objects.requireNonNull(defaults));
        }

        /**
         * @param nextHop The next hop of the destination route.
         * 
         * @return builder
         * 
         */
        public Builder nextHop(@Nullable Output<String> nextHop) {
            $.nextHop = nextHop;
            return this;
        }

        /**
         * @param nextHop The next hop of the destination route.
         * 
         * @return builder
         * 
         */
        public Builder nextHop(String nextHop) {
            return nextHop(Output.of(nextHop));
        }

        /**
         * @param routeDest The destination network segment of the destination route.
         * 
         * @return builder
         * 
         */
        public Builder routeDest(@Nullable Output<String> routeDest) {
            $.routeDest = routeDest;
            return this;
        }

        /**
         * @param routeDest The destination network segment of the destination route.
         * 
         * @return builder
         * 
         */
        public Builder routeDest(String routeDest) {
            return routeDest(Output.of(routeDest));
        }

        /**
         * @param status The status of the vpn route entry.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the vpn route entry.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param vpnConnectionId The id of the vpn attachment.
         * 
         * @return builder
         * 
         */
        public Builder vpnConnectionId(@Nullable Output<String> vpnConnectionId) {
            $.vpnConnectionId = vpnConnectionId;
            return this;
        }

        /**
         * @param vpnConnectionId The id of the vpn attachment.
         * 
         * @return builder
         * 
         */
        public Builder vpnConnectionId(String vpnConnectionId) {
            return vpnConnectionId(Output.of(vpnConnectionId));
        }

        /**
         * @param weight The weight value of the destination route. Valid values: `0`, `100`.
         * 
         * @return builder
         * 
         */
        public Builder weight(@Nullable Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param weight The weight value of the destination route. Valid values: `0`, `100`.
         * 
         * @return builder
         * 
         */
        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public GatewayVcoRouteState build() {
            return $;
        }
    }

}
