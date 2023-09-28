// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RouteTableAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final RouteTableAttachmentState Empty = new RouteTableAttachmentState();

    /**
     * The ID of the route table to be bound to the switch.
     * 
     */
    @Import(name="routeTableId")
    private @Nullable Output<String> routeTableId;

    /**
     * @return The ID of the route table to be bound to the switch.
     * 
     */
    public Optional<Output<String>> routeTableId() {
        return Optional.ofNullable(this.routeTableId);
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
     * The ID of the switch to bind the route table.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return The ID of the switch to bind the route table.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private RouteTableAttachmentState() {}

    private RouteTableAttachmentState(RouteTableAttachmentState $) {
        this.routeTableId = $.routeTableId;
        this.status = $.status;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouteTableAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouteTableAttachmentState $;

        public Builder() {
            $ = new RouteTableAttachmentState();
        }

        public Builder(RouteTableAttachmentState defaults) {
            $ = new RouteTableAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param routeTableId The ID of the route table to be bound to the switch.
         * 
         * @return builder
         * 
         */
        public Builder routeTableId(@Nullable Output<String> routeTableId) {
            $.routeTableId = routeTableId;
            return this;
        }

        /**
         * @param routeTableId The ID of the route table to be bound to the switch.
         * 
         * @return builder
         * 
         */
        public Builder routeTableId(String routeTableId) {
            return routeTableId(Output.of(routeTableId));
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
         * @param vswitchId The ID of the switch to bind the route table.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The ID of the switch to bind the route table.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public RouteTableAttachmentState build() {
            return $;
        }
    }

}
