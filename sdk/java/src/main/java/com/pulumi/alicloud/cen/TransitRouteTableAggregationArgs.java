// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TransitRouteTableAggregationArgs extends com.pulumi.resources.ResourceArgs {

    public static final TransitRouteTableAggregationArgs Empty = new TransitRouteTableAggregationArgs();

    /**
     * The destination CIDR block of the aggregate route. CIDR blocks that start with `0` or `100.64`. Multicast CIDR blocks, including `224.0.0.1` to `239.255.255.254`.
     * 
     */
    @Import(name="transitRouteTableAggregationCidr", required=true)
    private Output<String> transitRouteTableAggregationCidr;

    /**
     * @return The destination CIDR block of the aggregate route. CIDR blocks that start with `0` or `100.64`. Multicast CIDR blocks, including `224.0.0.1` to `239.255.255.254`.
     * 
     */
    public Output<String> transitRouteTableAggregationCidr() {
        return this.transitRouteTableAggregationCidr;
    }

    /**
     * The description of the aggregate route.
     * 
     */
    @Import(name="transitRouteTableAggregationDescription")
    private @Nullable Output<String> transitRouteTableAggregationDescription;

    /**
     * @return The description of the aggregate route.
     * 
     */
    public Optional<Output<String>> transitRouteTableAggregationDescription() {
        return Optional.ofNullable(this.transitRouteTableAggregationDescription);
    }

    /**
     * The name of the aggregate route.
     * 
     */
    @Import(name="transitRouteTableAggregationName")
    private @Nullable Output<String> transitRouteTableAggregationName;

    /**
     * @return The name of the aggregate route.
     * 
     */
    public Optional<Output<String>> transitRouteTableAggregationName() {
        return Optional.ofNullable(this.transitRouteTableAggregationName);
    }

    /**
     * The scope of networks that you want to advertise the aggregate route. Valid Value: `VPC`.
     * 
     */
    @Import(name="transitRouteTableAggregationScope", required=true)
    private Output<String> transitRouteTableAggregationScope;

    /**
     * @return The scope of networks that you want to advertise the aggregate route. Valid Value: `VPC`.
     * 
     */
    public Output<String> transitRouteTableAggregationScope() {
        return this.transitRouteTableAggregationScope;
    }

    /**
     * The ID of the route table of the Enterprise Edition transit router.
     * 
     */
    @Import(name="transitRouteTableId", required=true)
    private Output<String> transitRouteTableId;

    /**
     * @return The ID of the route table of the Enterprise Edition transit router.
     * 
     */
    public Output<String> transitRouteTableId() {
        return this.transitRouteTableId;
    }

    private TransitRouteTableAggregationArgs() {}

    private TransitRouteTableAggregationArgs(TransitRouteTableAggregationArgs $) {
        this.transitRouteTableAggregationCidr = $.transitRouteTableAggregationCidr;
        this.transitRouteTableAggregationDescription = $.transitRouteTableAggregationDescription;
        this.transitRouteTableAggregationName = $.transitRouteTableAggregationName;
        this.transitRouteTableAggregationScope = $.transitRouteTableAggregationScope;
        this.transitRouteTableId = $.transitRouteTableId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TransitRouteTableAggregationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TransitRouteTableAggregationArgs $;

        public Builder() {
            $ = new TransitRouteTableAggregationArgs();
        }

        public Builder(TransitRouteTableAggregationArgs defaults) {
            $ = new TransitRouteTableAggregationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param transitRouteTableAggregationCidr The destination CIDR block of the aggregate route. CIDR blocks that start with `0` or `100.64`. Multicast CIDR blocks, including `224.0.0.1` to `239.255.255.254`.
         * 
         * @return builder
         * 
         */
        public Builder transitRouteTableAggregationCidr(Output<String> transitRouteTableAggregationCidr) {
            $.transitRouteTableAggregationCidr = transitRouteTableAggregationCidr;
            return this;
        }

        /**
         * @param transitRouteTableAggregationCidr The destination CIDR block of the aggregate route. CIDR blocks that start with `0` or `100.64`. Multicast CIDR blocks, including `224.0.0.1` to `239.255.255.254`.
         * 
         * @return builder
         * 
         */
        public Builder transitRouteTableAggregationCidr(String transitRouteTableAggregationCidr) {
            return transitRouteTableAggregationCidr(Output.of(transitRouteTableAggregationCidr));
        }

        /**
         * @param transitRouteTableAggregationDescription The description of the aggregate route.
         * 
         * @return builder
         * 
         */
        public Builder transitRouteTableAggregationDescription(@Nullable Output<String> transitRouteTableAggregationDescription) {
            $.transitRouteTableAggregationDescription = transitRouteTableAggregationDescription;
            return this;
        }

        /**
         * @param transitRouteTableAggregationDescription The description of the aggregate route.
         * 
         * @return builder
         * 
         */
        public Builder transitRouteTableAggregationDescription(String transitRouteTableAggregationDescription) {
            return transitRouteTableAggregationDescription(Output.of(transitRouteTableAggregationDescription));
        }

        /**
         * @param transitRouteTableAggregationName The name of the aggregate route.
         * 
         * @return builder
         * 
         */
        public Builder transitRouteTableAggregationName(@Nullable Output<String> transitRouteTableAggregationName) {
            $.transitRouteTableAggregationName = transitRouteTableAggregationName;
            return this;
        }

        /**
         * @param transitRouteTableAggregationName The name of the aggregate route.
         * 
         * @return builder
         * 
         */
        public Builder transitRouteTableAggregationName(String transitRouteTableAggregationName) {
            return transitRouteTableAggregationName(Output.of(transitRouteTableAggregationName));
        }

        /**
         * @param transitRouteTableAggregationScope The scope of networks that you want to advertise the aggregate route. Valid Value: `VPC`.
         * 
         * @return builder
         * 
         */
        public Builder transitRouteTableAggregationScope(Output<String> transitRouteTableAggregationScope) {
            $.transitRouteTableAggregationScope = transitRouteTableAggregationScope;
            return this;
        }

        /**
         * @param transitRouteTableAggregationScope The scope of networks that you want to advertise the aggregate route. Valid Value: `VPC`.
         * 
         * @return builder
         * 
         */
        public Builder transitRouteTableAggregationScope(String transitRouteTableAggregationScope) {
            return transitRouteTableAggregationScope(Output.of(transitRouteTableAggregationScope));
        }

        /**
         * @param transitRouteTableId The ID of the route table of the Enterprise Edition transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouteTableId(Output<String> transitRouteTableId) {
            $.transitRouteTableId = transitRouteTableId;
            return this;
        }

        /**
         * @param transitRouteTableId The ID of the route table of the Enterprise Edition transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouteTableId(String transitRouteTableId) {
            return transitRouteTableId(Output.of(transitRouteTableId));
        }

        public TransitRouteTableAggregationArgs build() {
            $.transitRouteTableAggregationCidr = Objects.requireNonNull($.transitRouteTableAggregationCidr, "expected parameter 'transitRouteTableAggregationCidr' to be non-null");
            $.transitRouteTableAggregationScope = Objects.requireNonNull($.transitRouteTableAggregationScope, "expected parameter 'transitRouteTableAggregationScope' to be non-null");
            $.transitRouteTableId = Objects.requireNonNull($.transitRouteTableId, "expected parameter 'transitRouteTableId' to be non-null");
            return $;
        }
    }

}
