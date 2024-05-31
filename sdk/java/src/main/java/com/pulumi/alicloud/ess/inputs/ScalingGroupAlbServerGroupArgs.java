// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ScalingGroupAlbServerGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final ScalingGroupAlbServerGroupArgs Empty = new ScalingGroupAlbServerGroupArgs();

    /**
     * The ID of ALB server group.
     * 
     */
    @Import(name="albServerGroupId")
    private @Nullable Output<String> albServerGroupId;

    /**
     * @return The ID of ALB server group.
     * 
     */
    public Optional<Output<String>> albServerGroupId() {
        return Optional.ofNullable(this.albServerGroupId);
    }

    /**
     * The port number used by an ECS instance after Auto Scaling adds the ECS instance to ALB server group.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The port number used by an ECS instance after Auto Scaling adds the ECS instance to ALB server group.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The weight of the ECS instance as a backend server after Auto Scaling adds the ECS instance to ALB server group.
     * 
     */
    @Import(name="weight")
    private @Nullable Output<Integer> weight;

    /**
     * @return The weight of the ECS instance as a backend server after Auto Scaling adds the ECS instance to ALB server group.
     * 
     */
    public Optional<Output<Integer>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private ScalingGroupAlbServerGroupArgs() {}

    private ScalingGroupAlbServerGroupArgs(ScalingGroupAlbServerGroupArgs $) {
        this.albServerGroupId = $.albServerGroupId;
        this.port = $.port;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ScalingGroupAlbServerGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ScalingGroupAlbServerGroupArgs $;

        public Builder() {
            $ = new ScalingGroupAlbServerGroupArgs();
        }

        public Builder(ScalingGroupAlbServerGroupArgs defaults) {
            $ = new ScalingGroupAlbServerGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param albServerGroupId The ID of ALB server group.
         * 
         * @return builder
         * 
         */
        public Builder albServerGroupId(@Nullable Output<String> albServerGroupId) {
            $.albServerGroupId = albServerGroupId;
            return this;
        }

        /**
         * @param albServerGroupId The ID of ALB server group.
         * 
         * @return builder
         * 
         */
        public Builder albServerGroupId(String albServerGroupId) {
            return albServerGroupId(Output.of(albServerGroupId));
        }

        /**
         * @param port The port number used by an ECS instance after Auto Scaling adds the ECS instance to ALB server group.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port number used by an ECS instance after Auto Scaling adds the ECS instance to ALB server group.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param weight The weight of the ECS instance as a backend server after Auto Scaling adds the ECS instance to ALB server group.
         * 
         * @return builder
         * 
         */
        public Builder weight(@Nullable Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param weight The weight of the ECS instance as a backend server after Auto Scaling adds the ECS instance to ALB server group.
         * 
         * @return builder
         * 
         */
        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public ScalingGroupAlbServerGroupArgs build() {
            return $;
        }
    }

}
