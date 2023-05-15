// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final AttachmentArgs Empty = new AttachmentArgs();

    /**
     * Whether to remove forcibly &#34;AutoCreated&#34; ECS instances in order to release scaling group capacity &#34;MaxSize&#34; for attaching ECS instances. Default to false.
     * 
     * &gt; **NOTE:** &#34;AutoCreated&#34; ECS instance will be deleted after it is removed from scaling group, but &#34;Attached&#34; will be not.
     * 
     * &gt; **NOTE:** Restrictions on attaching ECS instances:
     * 
     * - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
     * - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
     * - The attached ECS instances must in the running state.
     * - The attached ECS instances has not been attached to other scaling groups.
     * - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.
     * 
     */
    @Import(name="force")
    private @Nullable Output<Boolean> force;

    /**
     * @return Whether to remove forcibly &#34;AutoCreated&#34; ECS instances in order to release scaling group capacity &#34;MaxSize&#34; for attaching ECS instances. Default to false.
     * 
     * &gt; **NOTE:** &#34;AutoCreated&#34; ECS instance will be deleted after it is removed from scaling group, but &#34;Attached&#34; will be not.
     * 
     * &gt; **NOTE:** Restrictions on attaching ECS instances:
     * 
     * - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
     * - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
     * - The attached ECS instances must in the running state.
     * - The attached ECS instances has not been attached to other scaling groups.
     * - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.
     * 
     */
    public Optional<Output<Boolean>> force() {
        return Optional.ofNullable(this.force);
    }

    /**
     * ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
     * 
     */
    @Import(name="instanceIds", required=true)
    private Output<List<String>> instanceIds;

    /**
     * @return ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
     * 
     */
    public Output<List<String>> instanceIds() {
        return this.instanceIds;
    }

    /**
     * ID of the scaling group of a scaling configuration.
     * 
     */
    @Import(name="scalingGroupId", required=true)
    private Output<String> scalingGroupId;

    /**
     * @return ID of the scaling group of a scaling configuration.
     * 
     */
    public Output<String> scalingGroupId() {
        return this.scalingGroupId;
    }

    private AttachmentArgs() {}

    private AttachmentArgs(AttachmentArgs $) {
        this.force = $.force;
        this.instanceIds = $.instanceIds;
        this.scalingGroupId = $.scalingGroupId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AttachmentArgs $;

        public Builder() {
            $ = new AttachmentArgs();
        }

        public Builder(AttachmentArgs defaults) {
            $ = new AttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param force Whether to remove forcibly &#34;AutoCreated&#34; ECS instances in order to release scaling group capacity &#34;MaxSize&#34; for attaching ECS instances. Default to false.
         * 
         * &gt; **NOTE:** &#34;AutoCreated&#34; ECS instance will be deleted after it is removed from scaling group, but &#34;Attached&#34; will be not.
         * 
         * &gt; **NOTE:** Restrictions on attaching ECS instances:
         * 
         * - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
         * - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
         * - The attached ECS instances must in the running state.
         * - The attached ECS instances has not been attached to other scaling groups.
         * - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.
         * 
         * @return builder
         * 
         */
        public Builder force(@Nullable Output<Boolean> force) {
            $.force = force;
            return this;
        }

        /**
         * @param force Whether to remove forcibly &#34;AutoCreated&#34; ECS instances in order to release scaling group capacity &#34;MaxSize&#34; for attaching ECS instances. Default to false.
         * 
         * &gt; **NOTE:** &#34;AutoCreated&#34; ECS instance will be deleted after it is removed from scaling group, but &#34;Attached&#34; will be not.
         * 
         * &gt; **NOTE:** Restrictions on attaching ECS instances:
         * 
         * - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
         * - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
         * - The attached ECS instances must in the running state.
         * - The attached ECS instances has not been attached to other scaling groups.
         * - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.
         * 
         * @return builder
         * 
         */
        public Builder force(Boolean force) {
            return force(Output.of(force));
        }

        /**
         * @param instanceIds ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
         * 
         * @return builder
         * 
         */
        public Builder instanceIds(Output<List<String>> instanceIds) {
            $.instanceIds = instanceIds;
            return this;
        }

        /**
         * @param instanceIds ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
         * 
         * @return builder
         * 
         */
        public Builder instanceIds(List<String> instanceIds) {
            return instanceIds(Output.of(instanceIds));
        }

        /**
         * @param instanceIds ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
         * 
         * @return builder
         * 
         */
        public Builder instanceIds(String... instanceIds) {
            return instanceIds(List.of(instanceIds));
        }

        /**
         * @param scalingGroupId ID of the scaling group of a scaling configuration.
         * 
         * @return builder
         * 
         */
        public Builder scalingGroupId(Output<String> scalingGroupId) {
            $.scalingGroupId = scalingGroupId;
            return this;
        }

        /**
         * @param scalingGroupId ID of the scaling group of a scaling configuration.
         * 
         * @return builder
         * 
         */
        public Builder scalingGroupId(String scalingGroupId) {
            return scalingGroupId(Output.of(scalingGroupId));
        }

        public AttachmentArgs build() {
            $.instanceIds = Objects.requireNonNull($.instanceIds, "expected parameter 'instanceIds' to be non-null");
            $.scalingGroupId = Objects.requireNonNull($.scalingGroupId, "expected parameter 'scalingGroupId' to be non-null");
            return $;
        }
    }

}
