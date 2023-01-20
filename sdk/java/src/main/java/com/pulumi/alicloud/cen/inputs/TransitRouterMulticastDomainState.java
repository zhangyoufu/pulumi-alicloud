// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TransitRouterMulticastDomainState extends com.pulumi.resources.ResourceArgs {

    public static final TransitRouterMulticastDomainState Empty = new TransitRouterMulticastDomainState();

    /**
     * The status of the Transit Router Multicast Domain.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the Transit Router Multicast Domain.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The ID of the transit router.
     * 
     */
    @Import(name="transitRouterId")
    private @Nullable Output<String> transitRouterId;

    /**
     * @return The ID of the transit router.
     * 
     */
    public Optional<Output<String>> transitRouterId() {
        return Optional.ofNullable(this.transitRouterId);
    }

    /**
     * The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
     * 
     */
    @Import(name="transitRouterMulticastDomainDescription")
    private @Nullable Output<String> transitRouterMulticastDomainDescription;

    /**
     * @return The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
     * 
     */
    public Optional<Output<String>> transitRouterMulticastDomainDescription() {
        return Optional.ofNullable(this.transitRouterMulticastDomainDescription);
    }

    /**
     * The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
     * 
     */
    @Import(name="transitRouterMulticastDomainName")
    private @Nullable Output<String> transitRouterMulticastDomainName;

    /**
     * @return The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
     * 
     */
    public Optional<Output<String>> transitRouterMulticastDomainName() {
        return Optional.ofNullable(this.transitRouterMulticastDomainName);
    }

    private TransitRouterMulticastDomainState() {}

    private TransitRouterMulticastDomainState(TransitRouterMulticastDomainState $) {
        this.status = $.status;
        this.tags = $.tags;
        this.transitRouterId = $.transitRouterId;
        this.transitRouterMulticastDomainDescription = $.transitRouterMulticastDomainDescription;
        this.transitRouterMulticastDomainName = $.transitRouterMulticastDomainName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TransitRouterMulticastDomainState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TransitRouterMulticastDomainState $;

        public Builder() {
            $ = new TransitRouterMulticastDomainState();
        }

        public Builder(TransitRouterMulticastDomainState defaults) {
            $ = new TransitRouterMulticastDomainState(Objects.requireNonNull(defaults));
        }

        /**
         * @param status The status of the Transit Router Multicast Domain.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the Transit Router Multicast Domain.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param transitRouterId The ID of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterId(@Nullable Output<String> transitRouterId) {
            $.transitRouterId = transitRouterId;
            return this;
        }

        /**
         * @param transitRouterId The ID of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterId(String transitRouterId) {
            return transitRouterId(Output.of(transitRouterId));
        }

        /**
         * @param transitRouterMulticastDomainDescription The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder transitRouterMulticastDomainDescription(@Nullable Output<String> transitRouterMulticastDomainDescription) {
            $.transitRouterMulticastDomainDescription = transitRouterMulticastDomainDescription;
            return this;
        }

        /**
         * @param transitRouterMulticastDomainDescription The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder transitRouterMulticastDomainDescription(String transitRouterMulticastDomainDescription) {
            return transitRouterMulticastDomainDescription(Output.of(transitRouterMulticastDomainDescription));
        }

        /**
         * @param transitRouterMulticastDomainName The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder transitRouterMulticastDomainName(@Nullable Output<String> transitRouterMulticastDomainName) {
            $.transitRouterMulticastDomainName = transitRouterMulticastDomainName;
            return this;
        }

        /**
         * @param transitRouterMulticastDomainName The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder transitRouterMulticastDomainName(String transitRouterMulticastDomainName) {
            return transitRouterMulticastDomainName(Output.of(transitRouterMulticastDomainName));
        }

        public TransitRouterMulticastDomainState build() {
            return $;
        }
    }

}
