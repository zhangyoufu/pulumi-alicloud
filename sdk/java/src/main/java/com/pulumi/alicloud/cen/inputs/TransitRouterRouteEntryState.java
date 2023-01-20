// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TransitRouterRouteEntryState extends com.pulumi.resources.ResourceArgs {

    public static final TransitRouterRouteEntryState Empty = new TransitRouterRouteEntryState();

    /**
     * The dry run.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The associating status of the Transit Router.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The associating status of the Transit Router.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The description of the transit router route entry.
     * 
     */
    @Import(name="transitRouterRouteEntryDescription")
    private @Nullable Output<String> transitRouterRouteEntryDescription;

    /**
     * @return The description of the transit router route entry.
     * 
     */
    public Optional<Output<String>> transitRouterRouteEntryDescription() {
        return Optional.ofNullable(this.transitRouterRouteEntryDescription);
    }

    /**
     * The CIDR of the transit router route entry.
     * 
     */
    @Import(name="transitRouterRouteEntryDestinationCidrBlock")
    private @Nullable Output<String> transitRouterRouteEntryDestinationCidrBlock;

    /**
     * @return The CIDR of the transit router route entry.
     * 
     */
    public Optional<Output<String>> transitRouterRouteEntryDestinationCidrBlock() {
        return Optional.ofNullable(this.transitRouterRouteEntryDestinationCidrBlock);
    }

    /**
     * The ID of the route entry.
     * 
     */
    @Import(name="transitRouterRouteEntryId")
    private @Nullable Output<String> transitRouterRouteEntryId;

    /**
     * @return The ID of the route entry.
     * 
     */
    public Optional<Output<String>> transitRouterRouteEntryId() {
        return Optional.ofNullable(this.transitRouterRouteEntryId);
    }

    /**
     * The name of the transit router route entry.
     * 
     */
    @Import(name="transitRouterRouteEntryName")
    private @Nullable Output<String> transitRouterRouteEntryName;

    /**
     * @return The name of the transit router route entry.
     * 
     */
    public Optional<Output<String>> transitRouterRouteEntryName() {
        return Optional.ofNullable(this.transitRouterRouteEntryName);
    }

    /**
     * The ID of the transit router route entry next hop.
     * 
     */
    @Import(name="transitRouterRouteEntryNextHopId")
    private @Nullable Output<String> transitRouterRouteEntryNextHopId;

    /**
     * @return The ID of the transit router route entry next hop.
     * 
     */
    public Optional<Output<String>> transitRouterRouteEntryNextHopId() {
        return Optional.ofNullable(this.transitRouterRouteEntryNextHopId);
    }

    /**
     * The Type of the transit router route entry next hop,Valid values `Attachment` and `BlackHole`.
     * 
     */
    @Import(name="transitRouterRouteEntryNextHopType")
    private @Nullable Output<String> transitRouterRouteEntryNextHopType;

    /**
     * @return The Type of the transit router route entry next hop,Valid values `Attachment` and `BlackHole`.
     * 
     */
    public Optional<Output<String>> transitRouterRouteEntryNextHopType() {
        return Optional.ofNullable(this.transitRouterRouteEntryNextHopType);
    }

    /**
     * The ID of the transit router route table.
     * 
     */
    @Import(name="transitRouterRouteTableId")
    private @Nullable Output<String> transitRouterRouteTableId;

    /**
     * @return The ID of the transit router route table.
     * 
     */
    public Optional<Output<String>> transitRouterRouteTableId() {
        return Optional.ofNullable(this.transitRouterRouteTableId);
    }

    private TransitRouterRouteEntryState() {}

    private TransitRouterRouteEntryState(TransitRouterRouteEntryState $) {
        this.dryRun = $.dryRun;
        this.status = $.status;
        this.transitRouterRouteEntryDescription = $.transitRouterRouteEntryDescription;
        this.transitRouterRouteEntryDestinationCidrBlock = $.transitRouterRouteEntryDestinationCidrBlock;
        this.transitRouterRouteEntryId = $.transitRouterRouteEntryId;
        this.transitRouterRouteEntryName = $.transitRouterRouteEntryName;
        this.transitRouterRouteEntryNextHopId = $.transitRouterRouteEntryNextHopId;
        this.transitRouterRouteEntryNextHopType = $.transitRouterRouteEntryNextHopType;
        this.transitRouterRouteTableId = $.transitRouterRouteTableId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TransitRouterRouteEntryState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TransitRouterRouteEntryState $;

        public Builder() {
            $ = new TransitRouterRouteEntryState();
        }

        public Builder(TransitRouterRouteEntryState defaults) {
            $ = new TransitRouterRouteEntryState(Objects.requireNonNull(defaults));
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param status The associating status of the Transit Router.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The associating status of the Transit Router.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param transitRouterRouteEntryDescription The description of the transit router route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryDescription(@Nullable Output<String> transitRouterRouteEntryDescription) {
            $.transitRouterRouteEntryDescription = transitRouterRouteEntryDescription;
            return this;
        }

        /**
         * @param transitRouterRouteEntryDescription The description of the transit router route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryDescription(String transitRouterRouteEntryDescription) {
            return transitRouterRouteEntryDescription(Output.of(transitRouterRouteEntryDescription));
        }

        /**
         * @param transitRouterRouteEntryDestinationCidrBlock The CIDR of the transit router route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryDestinationCidrBlock(@Nullable Output<String> transitRouterRouteEntryDestinationCidrBlock) {
            $.transitRouterRouteEntryDestinationCidrBlock = transitRouterRouteEntryDestinationCidrBlock;
            return this;
        }

        /**
         * @param transitRouterRouteEntryDestinationCidrBlock The CIDR of the transit router route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryDestinationCidrBlock(String transitRouterRouteEntryDestinationCidrBlock) {
            return transitRouterRouteEntryDestinationCidrBlock(Output.of(transitRouterRouteEntryDestinationCidrBlock));
        }

        /**
         * @param transitRouterRouteEntryId The ID of the route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryId(@Nullable Output<String> transitRouterRouteEntryId) {
            $.transitRouterRouteEntryId = transitRouterRouteEntryId;
            return this;
        }

        /**
         * @param transitRouterRouteEntryId The ID of the route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryId(String transitRouterRouteEntryId) {
            return transitRouterRouteEntryId(Output.of(transitRouterRouteEntryId));
        }

        /**
         * @param transitRouterRouteEntryName The name of the transit router route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryName(@Nullable Output<String> transitRouterRouteEntryName) {
            $.transitRouterRouteEntryName = transitRouterRouteEntryName;
            return this;
        }

        /**
         * @param transitRouterRouteEntryName The name of the transit router route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryName(String transitRouterRouteEntryName) {
            return transitRouterRouteEntryName(Output.of(transitRouterRouteEntryName));
        }

        /**
         * @param transitRouterRouteEntryNextHopId The ID of the transit router route entry next hop.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryNextHopId(@Nullable Output<String> transitRouterRouteEntryNextHopId) {
            $.transitRouterRouteEntryNextHopId = transitRouterRouteEntryNextHopId;
            return this;
        }

        /**
         * @param transitRouterRouteEntryNextHopId The ID of the transit router route entry next hop.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryNextHopId(String transitRouterRouteEntryNextHopId) {
            return transitRouterRouteEntryNextHopId(Output.of(transitRouterRouteEntryNextHopId));
        }

        /**
         * @param transitRouterRouteEntryNextHopType The Type of the transit router route entry next hop,Valid values `Attachment` and `BlackHole`.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryNextHopType(@Nullable Output<String> transitRouterRouteEntryNextHopType) {
            $.transitRouterRouteEntryNextHopType = transitRouterRouteEntryNextHopType;
            return this;
        }

        /**
         * @param transitRouterRouteEntryNextHopType The Type of the transit router route entry next hop,Valid values `Attachment` and `BlackHole`.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryNextHopType(String transitRouterRouteEntryNextHopType) {
            return transitRouterRouteEntryNextHopType(Output.of(transitRouterRouteEntryNextHopType));
        }

        /**
         * @param transitRouterRouteTableId The ID of the transit router route table.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteTableId(@Nullable Output<String> transitRouterRouteTableId) {
            $.transitRouterRouteTableId = transitRouterRouteTableId;
            return this;
        }

        /**
         * @param transitRouterRouteTableId The ID of the transit router route table.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteTableId(String transitRouterRouteTableId) {
            return transitRouterRouteTableId(Output.of(transitRouterRouteTableId));
        }

        public TransitRouterRouteEntryState build() {
            return $;
        }
    }

}
