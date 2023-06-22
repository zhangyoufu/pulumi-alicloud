// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.inputs;

import com.pulumi.alicloud.ess.inputs.ScalingGroupVServerGroupsVserverGroupVserverAttributeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class ScalingGroupVServerGroupsVserverGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final ScalingGroupVServerGroupsVserverGroupArgs Empty = new ScalingGroupVServerGroupsVserverGroupArgs();

    /**
     * Loadbalancer server ID of VServer Group.
     * 
     */
    @Import(name="loadbalancerId", required=true)
    private Output<String> loadbalancerId;

    /**
     * @return Loadbalancer server ID of VServer Group.
     * 
     */
    public Output<String> loadbalancerId() {
        return this.loadbalancerId;
    }

    /**
     * A list of VServer Group attributes. See `vserver_attributes` below.
     * 
     */
    @Import(name="vserverAttributes", required=true)
    private Output<List<ScalingGroupVServerGroupsVserverGroupVserverAttributeArgs>> vserverAttributes;

    /**
     * @return A list of VServer Group attributes. See `vserver_attributes` below.
     * 
     */
    public Output<List<ScalingGroupVServerGroupsVserverGroupVserverAttributeArgs>> vserverAttributes() {
        return this.vserverAttributes;
    }

    private ScalingGroupVServerGroupsVserverGroupArgs() {}

    private ScalingGroupVServerGroupsVserverGroupArgs(ScalingGroupVServerGroupsVserverGroupArgs $) {
        this.loadbalancerId = $.loadbalancerId;
        this.vserverAttributes = $.vserverAttributes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ScalingGroupVServerGroupsVserverGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ScalingGroupVServerGroupsVserverGroupArgs $;

        public Builder() {
            $ = new ScalingGroupVServerGroupsVserverGroupArgs();
        }

        public Builder(ScalingGroupVServerGroupsVserverGroupArgs defaults) {
            $ = new ScalingGroupVServerGroupsVserverGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param loadbalancerId Loadbalancer server ID of VServer Group.
         * 
         * @return builder
         * 
         */
        public Builder loadbalancerId(Output<String> loadbalancerId) {
            $.loadbalancerId = loadbalancerId;
            return this;
        }

        /**
         * @param loadbalancerId Loadbalancer server ID of VServer Group.
         * 
         * @return builder
         * 
         */
        public Builder loadbalancerId(String loadbalancerId) {
            return loadbalancerId(Output.of(loadbalancerId));
        }

        /**
         * @param vserverAttributes A list of VServer Group attributes. See `vserver_attributes` below.
         * 
         * @return builder
         * 
         */
        public Builder vserverAttributes(Output<List<ScalingGroupVServerGroupsVserverGroupVserverAttributeArgs>> vserverAttributes) {
            $.vserverAttributes = vserverAttributes;
            return this;
        }

        /**
         * @param vserverAttributes A list of VServer Group attributes. See `vserver_attributes` below.
         * 
         * @return builder
         * 
         */
        public Builder vserverAttributes(List<ScalingGroupVServerGroupsVserverGroupVserverAttributeArgs> vserverAttributes) {
            return vserverAttributes(Output.of(vserverAttributes));
        }

        /**
         * @param vserverAttributes A list of VServer Group attributes. See `vserver_attributes` below.
         * 
         * @return builder
         * 
         */
        public Builder vserverAttributes(ScalingGroupVServerGroupsVserverGroupVserverAttributeArgs... vserverAttributes) {
            return vserverAttributes(List.of(vserverAttributes));
        }

        public ScalingGroupVServerGroupsVserverGroupArgs build() {
            $.loadbalancerId = Objects.requireNonNull($.loadbalancerId, "expected parameter 'loadbalancerId' to be non-null");
            $.vserverAttributes = Objects.requireNonNull($.vserverAttributes, "expected parameter 'vserverAttributes' to be non-null");
            return $;
        }
    }

}
