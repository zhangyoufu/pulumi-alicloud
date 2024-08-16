// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TransitRouterState extends com.pulumi.resources.ResourceArgs {

    public static final TransitRouterState Empty = new TransitRouterState();

    /**
     * The ID of the CEN.
     * 
     */
    @Import(name="cenId")
    private @Nullable Output<String> cenId;

    /**
     * @return The ID of the CEN.
     * 
     */
    public Optional<Output<String>> cenId() {
        return Optional.ofNullable(this.cenId);
    }

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
     * Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values: `false`, `true`. Default Value: `false`. The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-listtransitrouteravailableresource) to query the regions that support multicast.
     * 
     */
    @Import(name="supportMulticast")
    private @Nullable Output<Boolean> supportMulticast;

    /**
     * @return Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values: `false`, `true`. Default Value: `false`. The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-listtransitrouteravailableresource) to query the regions that support multicast.
     * 
     */
    public Optional<Output<Boolean>> supportMulticast() {
        return Optional.ofNullable(this.supportMulticast);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The description of the transit router.
     * 
     */
    @Import(name="transitRouterDescription")
    private @Nullable Output<String> transitRouterDescription;

    /**
     * @return The description of the transit router.
     * 
     */
    public Optional<Output<String>> transitRouterDescription() {
        return Optional.ofNullable(this.transitRouterDescription);
    }

    /**
     * The transit router id of the transit router.
     * 
     */
    @Import(name="transitRouterId")
    private @Nullable Output<String> transitRouterId;

    /**
     * @return The transit router id of the transit router.
     * 
     */
    public Optional<Output<String>> transitRouterId() {
        return Optional.ofNullable(this.transitRouterId);
    }

    /**
     * The name of the transit router.
     * 
     */
    @Import(name="transitRouterName")
    private @Nullable Output<String> transitRouterName;

    /**
     * @return The name of the transit router.
     * 
     */
    public Optional<Output<String>> transitRouterName() {
        return Optional.ofNullable(this.transitRouterName);
    }

    /**
     * The Type of the Transit Router. Valid values: `Enterprise`, `Basic`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The Type of the Transit Router. Valid values: `Enterprise`, `Basic`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private TransitRouterState() {}

    private TransitRouterState(TransitRouterState $) {
        this.cenId = $.cenId;
        this.dryRun = $.dryRun;
        this.status = $.status;
        this.supportMulticast = $.supportMulticast;
        this.tags = $.tags;
        this.transitRouterDescription = $.transitRouterDescription;
        this.transitRouterId = $.transitRouterId;
        this.transitRouterName = $.transitRouterName;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TransitRouterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TransitRouterState $;

        public Builder() {
            $ = new TransitRouterState();
        }

        public Builder(TransitRouterState defaults) {
            $ = new TransitRouterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cenId The ID of the CEN.
         * 
         * @return builder
         * 
         */
        public Builder cenId(@Nullable Output<String> cenId) {
            $.cenId = cenId;
            return this;
        }

        /**
         * @param cenId The ID of the CEN.
         * 
         * @return builder
         * 
         */
        public Builder cenId(String cenId) {
            return cenId(Output.of(cenId));
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
         * @param supportMulticast Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values: `false`, `true`. Default Value: `false`. The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-listtransitrouteravailableresource) to query the regions that support multicast.
         * 
         * @return builder
         * 
         */
        public Builder supportMulticast(@Nullable Output<Boolean> supportMulticast) {
            $.supportMulticast = supportMulticast;
            return this;
        }

        /**
         * @param supportMulticast Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values: `false`, `true`. Default Value: `false`. The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-listtransitrouteravailableresource) to query the regions that support multicast.
         * 
         * @return builder
         * 
         */
        public Builder supportMulticast(Boolean supportMulticast) {
            return supportMulticast(Output.of(supportMulticast));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param transitRouterDescription The description of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterDescription(@Nullable Output<String> transitRouterDescription) {
            $.transitRouterDescription = transitRouterDescription;
            return this;
        }

        /**
         * @param transitRouterDescription The description of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterDescription(String transitRouterDescription) {
            return transitRouterDescription(Output.of(transitRouterDescription));
        }

        /**
         * @param transitRouterId The transit router id of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterId(@Nullable Output<String> transitRouterId) {
            $.transitRouterId = transitRouterId;
            return this;
        }

        /**
         * @param transitRouterId The transit router id of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterId(String transitRouterId) {
            return transitRouterId(Output.of(transitRouterId));
        }

        /**
         * @param transitRouterName The name of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterName(@Nullable Output<String> transitRouterName) {
            $.transitRouterName = transitRouterName;
            return this;
        }

        /**
         * @param transitRouterName The name of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterName(String transitRouterName) {
            return transitRouterName(Output.of(transitRouterName));
        }

        /**
         * @param type The Type of the Transit Router. Valid values: `Enterprise`, `Basic`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The Type of the Transit Router. Valid values: `Enterprise`, `Basic`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public TransitRouterState build() {
            return $;
        }
    }

}
