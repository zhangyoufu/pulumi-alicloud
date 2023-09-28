// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicyAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final PolicyAttachmentState Empty = new PolicyAttachmentState();

    /**
     * The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
     * 
     */
    @Import(name="policyName")
    private @Nullable Output<String> policyName;

    /**
     * @return The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
     * 
     */
    public Optional<Output<String>> policyName() {
        return Optional.ofNullable(this.policyName);
    }

    /**
     * The type of the policy. Valid values: `Custom`, `System`.
     * 
     */
    @Import(name="policyType")
    private @Nullable Output<String> policyType;

    /**
     * @return The type of the policy. Valid values: `Custom`, `System`.
     * 
     */
    public Optional<Output<String>> policyType() {
        return Optional.ofNullable(this.policyType);
    }

    /**
     * The name of the object to which you want to attach the policy.
     * 
     */
    @Import(name="principalName")
    private @Nullable Output<String> principalName;

    /**
     * @return The name of the object to which you want to attach the policy.
     * 
     */
    public Optional<Output<String>> principalName() {
        return Optional.ofNullable(this.principalName);
    }

    /**
     * The type of the object to which you want to attach the policy. Valid values: `IMSUser`: RAM user, `IMSGroup`: RAM user group, `ServiceRole`: RAM role.
     * 
     */
    @Import(name="principalType")
    private @Nullable Output<String> principalType;

    /**
     * @return The type of the object to which you want to attach the policy. Valid values: `IMSUser`: RAM user, `IMSGroup`: RAM user group, `ServiceRole`: RAM role.
     * 
     */
    public Optional<Output<String>> principalType() {
        return Optional.ofNullable(this.principalType);
    }

    /**
     * The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    private PolicyAttachmentState() {}

    private PolicyAttachmentState(PolicyAttachmentState $) {
        this.policyName = $.policyName;
        this.policyType = $.policyType;
        this.principalName = $.principalName;
        this.principalType = $.principalType;
        this.resourceGroupId = $.resourceGroupId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyAttachmentState $;

        public Builder() {
            $ = new PolicyAttachmentState();
        }

        public Builder(PolicyAttachmentState defaults) {
            $ = new PolicyAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param policyName The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder policyName(@Nullable Output<String> policyName) {
            $.policyName = policyName;
            return this;
        }

        /**
         * @param policyName The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder policyName(String policyName) {
            return policyName(Output.of(policyName));
        }

        /**
         * @param policyType The type of the policy. Valid values: `Custom`, `System`.
         * 
         * @return builder
         * 
         */
        public Builder policyType(@Nullable Output<String> policyType) {
            $.policyType = policyType;
            return this;
        }

        /**
         * @param policyType The type of the policy. Valid values: `Custom`, `System`.
         * 
         * @return builder
         * 
         */
        public Builder policyType(String policyType) {
            return policyType(Output.of(policyType));
        }

        /**
         * @param principalName The name of the object to which you want to attach the policy.
         * 
         * @return builder
         * 
         */
        public Builder principalName(@Nullable Output<String> principalName) {
            $.principalName = principalName;
            return this;
        }

        /**
         * @param principalName The name of the object to which you want to attach the policy.
         * 
         * @return builder
         * 
         */
        public Builder principalName(String principalName) {
            return principalName(Output.of(principalName));
        }

        /**
         * @param principalType The type of the object to which you want to attach the policy. Valid values: `IMSUser`: RAM user, `IMSGroup`: RAM user group, `ServiceRole`: RAM role.
         * 
         * @return builder
         * 
         */
        public Builder principalType(@Nullable Output<String> principalType) {
            $.principalType = principalType;
            return this;
        }

        /**
         * @param principalType The type of the object to which you want to attach the policy. Valid values: `IMSUser`: RAM user, `IMSGroup`: RAM user group, `ServiceRole`: RAM role.
         * 
         * @return builder
         * 
         */
        public Builder principalType(String principalType) {
            return principalType(Output.of(principalType));
        }

        /**
         * @param resourceGroupId The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        public PolicyAttachmentState build() {
            return $;
        }
    }

}
