// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SharedTargetState extends com.pulumi.resources.ResourceArgs {

    public static final SharedTargetState Empty = new SharedTargetState();

    /**
     * The resource share ID of resource manager.
     * 
     */
    @Import(name="resourceShareId")
    private @Nullable Output<String> resourceShareId;

    /**
     * @return The resource share ID of resource manager.
     * 
     */
    public Optional<Output<String>> resourceShareId() {
        return Optional.ofNullable(this.resourceShareId);
    }

    /**
     * The status of shared target.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of shared target.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The member account ID in resource directory.
     * 
     */
    @Import(name="targetId")
    private @Nullable Output<String> targetId;

    /**
     * @return The member account ID in resource directory.
     * 
     */
    public Optional<Output<String>> targetId() {
        return Optional.ofNullable(this.targetId);
    }

    private SharedTargetState() {}

    private SharedTargetState(SharedTargetState $) {
        this.resourceShareId = $.resourceShareId;
        this.status = $.status;
        this.targetId = $.targetId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SharedTargetState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SharedTargetState $;

        public Builder() {
            $ = new SharedTargetState();
        }

        public Builder(SharedTargetState defaults) {
            $ = new SharedTargetState(Objects.requireNonNull(defaults));
        }

        /**
         * @param resourceShareId The resource share ID of resource manager.
         * 
         * @return builder
         * 
         */
        public Builder resourceShareId(@Nullable Output<String> resourceShareId) {
            $.resourceShareId = resourceShareId;
            return this;
        }

        /**
         * @param resourceShareId The resource share ID of resource manager.
         * 
         * @return builder
         * 
         */
        public Builder resourceShareId(String resourceShareId) {
            return resourceShareId(Output.of(resourceShareId));
        }

        /**
         * @param status The status of shared target.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of shared target.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param targetId The member account ID in resource directory.
         * 
         * @return builder
         * 
         */
        public Builder targetId(@Nullable Output<String> targetId) {
            $.targetId = targetId;
            return this;
        }

        /**
         * @param targetId The member account ID in resource directory.
         * 
         * @return builder
         * 
         */
        public Builder targetId(String targetId) {
            return targetId(Output.of(targetId));
        }

        public SharedTargetState build() {
            return $;
        }
    }

}
