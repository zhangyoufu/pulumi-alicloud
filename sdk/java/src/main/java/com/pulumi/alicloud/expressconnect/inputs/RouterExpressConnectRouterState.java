// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.expressconnect.inputs;

import com.pulumi.alicloud.expressconnect.inputs.RouterExpressConnectRouterRegionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RouterExpressConnectRouterState extends com.pulumi.resources.ResourceArgs {

    public static final RouterExpressConnectRouterState Empty = new RouterExpressConnectRouterState();

    /**
     * ASN representing resources.
     * 
     */
    @Import(name="alibabaSideAsn")
    private @Nullable Output<Integer> alibabaSideAsn;

    /**
     * @return ASN representing resources.
     * 
     */
    public Optional<Output<Integer>> alibabaSideAsn() {
        return Optional.ofNullable(this.alibabaSideAsn);
    }

    /**
     * Represents the creation time of the resource.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Represents the creation time of the resource.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Represents the description of the leased line gateway.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Represents the description of the leased line gateway.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Name of the Gateway representing the leased line.
     * 
     */
    @Import(name="ecrName")
    private @Nullable Output<String> ecrName;

    /**
     * @return Name of the Gateway representing the leased line.
     * 
     */
    public Optional<Output<String>> ecrName() {
        return Optional.ofNullable(this.ecrName);
    }

    /**
     * List of regions representing leased line gateways. See `regions` below.
     * 
     */
    @Import(name="regions")
    private @Nullable Output<List<RouterExpressConnectRouterRegionArgs>> regions;

    /**
     * @return List of regions representing leased line gateways. See `regions` below.
     * 
     */
    public Optional<Output<List<RouterExpressConnectRouterRegionArgs>>> regions() {
        return Optional.ofNullable(this.regions);
    }

    /**
     * The ID of the resource group to which the ECR instance belongs.
     * - A string consisting of letters, numbers, hyphens (-), and underscores (_), and the string length can be 0 to 64 characters.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which the ECR instance belongs.
     * - A string consisting of letters, numbers, hyphens (-), and underscores (_), and the string length can be 0 to 64 characters.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
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
     * The tag of the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return The tag of the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private RouterExpressConnectRouterState() {}

    private RouterExpressConnectRouterState(RouterExpressConnectRouterState $) {
        this.alibabaSideAsn = $.alibabaSideAsn;
        this.createTime = $.createTime;
        this.description = $.description;
        this.ecrName = $.ecrName;
        this.regions = $.regions;
        this.resourceGroupId = $.resourceGroupId;
        this.status = $.status;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouterExpressConnectRouterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouterExpressConnectRouterState $;

        public Builder() {
            $ = new RouterExpressConnectRouterState();
        }

        public Builder(RouterExpressConnectRouterState defaults) {
            $ = new RouterExpressConnectRouterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alibabaSideAsn ASN representing resources.
         * 
         * @return builder
         * 
         */
        public Builder alibabaSideAsn(@Nullable Output<Integer> alibabaSideAsn) {
            $.alibabaSideAsn = alibabaSideAsn;
            return this;
        }

        /**
         * @param alibabaSideAsn ASN representing resources.
         * 
         * @return builder
         * 
         */
        public Builder alibabaSideAsn(Integer alibabaSideAsn) {
            return alibabaSideAsn(Output.of(alibabaSideAsn));
        }

        /**
         * @param createTime Represents the creation time of the resource.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Represents the creation time of the resource.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description Represents the description of the leased line gateway.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Represents the description of the leased line gateway.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ecrName Name of the Gateway representing the leased line.
         * 
         * @return builder
         * 
         */
        public Builder ecrName(@Nullable Output<String> ecrName) {
            $.ecrName = ecrName;
            return this;
        }

        /**
         * @param ecrName Name of the Gateway representing the leased line.
         * 
         * @return builder
         * 
         */
        public Builder ecrName(String ecrName) {
            return ecrName(Output.of(ecrName));
        }

        /**
         * @param regions List of regions representing leased line gateways. See `regions` below.
         * 
         * @return builder
         * 
         */
        public Builder regions(@Nullable Output<List<RouterExpressConnectRouterRegionArgs>> regions) {
            $.regions = regions;
            return this;
        }

        /**
         * @param regions List of regions representing leased line gateways. See `regions` below.
         * 
         * @return builder
         * 
         */
        public Builder regions(List<RouterExpressConnectRouterRegionArgs> regions) {
            return regions(Output.of(regions));
        }

        /**
         * @param regions List of regions representing leased line gateways. See `regions` below.
         * 
         * @return builder
         * 
         */
        public Builder regions(RouterExpressConnectRouterRegionArgs... regions) {
            return regions(List.of(regions));
        }

        /**
         * @param resourceGroupId The ID of the resource group to which the ECR instance belongs.
         * - A string consisting of letters, numbers, hyphens (-), and underscores (_), and the string length can be 0 to 64 characters.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group to which the ECR instance belongs.
         * - A string consisting of letters, numbers, hyphens (-), and underscores (_), and the string length can be 0 to 64 characters.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
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
         * @param tags The tag of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tag of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public RouterExpressConnectRouterState build() {
            return $;
        }
    }

}
