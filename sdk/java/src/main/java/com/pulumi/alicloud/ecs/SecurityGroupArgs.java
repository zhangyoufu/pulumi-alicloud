// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecurityGroupArgs Empty = new SecurityGroupArgs();

    /**
     * The security group description. Defaults to null.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The security group description. Defaults to null.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Field `inner_access` has been deprecated from provider version 1.55.3. New field `inner_access_policy` instead.
     * 
     * Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
     * 
     * @deprecated
     * Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it.
     * 
     */
    @Deprecated /* Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it. */
    @Import(name="innerAccess")
    private @Nullable Output<Boolean> innerAccess;

    /**
     * @return Field `inner_access` has been deprecated from provider version 1.55.3. New field `inner_access_policy` instead.
     * 
     * Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
     * 
     * @deprecated
     * Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it.
     * 
     */
    @Deprecated /* Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it. */
    public Optional<Output<Boolean>> innerAccess() {
        return Optional.ofNullable(this.innerAccess);
    }

    /**
     * The internal access control policy of the security group. Valid values: `Accept`, `Drop`.
     * 
     */
    @Import(name="innerAccessPolicy")
    private @Nullable Output<String> innerAccessPolicy;

    /**
     * @return The internal access control policy of the security group. Valid values: `Accept`, `Drop`.
     * 
     */
    public Optional<Output<String>> innerAccessPolicy() {
        return Optional.ofNullable(this.innerAccessPolicy);
    }

    /**
     * The name of the security group. Defaults to null.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the security group. Defaults to null.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the resource group to which the security group belongs. **NOTE:** From version 1.115.0, `resource_group_id` can be modified.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which the security group belongs. **NOTE:** From version 1.115.0, `resource_group_id` can be modified.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The type of the security group. Valid values:
     * 
     */
    @Import(name="securityGroupType")
    private @Nullable Output<String> securityGroupType;

    /**
     * @return The type of the security group. Valid values:
     * 
     */
    public Optional<Output<String>> securityGroupType() {
        return Optional.ofNullable(this.securityGroupType);
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
     * The ID of the VPC.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the VPC.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private SecurityGroupArgs() {}

    private SecurityGroupArgs(SecurityGroupArgs $) {
        this.description = $.description;
        this.innerAccess = $.innerAccess;
        this.innerAccessPolicy = $.innerAccessPolicy;
        this.name = $.name;
        this.resourceGroupId = $.resourceGroupId;
        this.securityGroupType = $.securityGroupType;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityGroupArgs $;

        public Builder() {
            $ = new SecurityGroupArgs();
        }

        public Builder(SecurityGroupArgs defaults) {
            $ = new SecurityGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The security group description. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The security group description. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param innerAccess Field `inner_access` has been deprecated from provider version 1.55.3. New field `inner_access_policy` instead.
         * 
         * Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
         * 
         * @return builder
         * 
         * @deprecated
         * Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it.
         * 
         */
        @Deprecated /* Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it. */
        public Builder innerAccess(@Nullable Output<Boolean> innerAccess) {
            $.innerAccess = innerAccess;
            return this;
        }

        /**
         * @param innerAccess Field `inner_access` has been deprecated from provider version 1.55.3. New field `inner_access_policy` instead.
         * 
         * Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
         * 
         * @return builder
         * 
         * @deprecated
         * Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it.
         * 
         */
        @Deprecated /* Field `inner_access` has been deprecated from provider version 1.55.3. Use `inner_access_policy` replaces it. */
        public Builder innerAccess(Boolean innerAccess) {
            return innerAccess(Output.of(innerAccess));
        }

        /**
         * @param innerAccessPolicy The internal access control policy of the security group. Valid values: `Accept`, `Drop`.
         * 
         * @return builder
         * 
         */
        public Builder innerAccessPolicy(@Nullable Output<String> innerAccessPolicy) {
            $.innerAccessPolicy = innerAccessPolicy;
            return this;
        }

        /**
         * @param innerAccessPolicy The internal access control policy of the security group. Valid values: `Accept`, `Drop`.
         * 
         * @return builder
         * 
         */
        public Builder innerAccessPolicy(String innerAccessPolicy) {
            return innerAccessPolicy(Output.of(innerAccessPolicy));
        }

        /**
         * @param name The name of the security group. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the security group. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param resourceGroupId The ID of the resource group to which the security group belongs. **NOTE:** From version 1.115.0, `resource_group_id` can be modified.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group to which the security group belongs. **NOTE:** From version 1.115.0, `resource_group_id` can be modified.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param securityGroupType The type of the security group. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder securityGroupType(@Nullable Output<String> securityGroupType) {
            $.securityGroupType = securityGroupType;
            return this;
        }

        /**
         * @param securityGroupType The type of the security group. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder securityGroupType(String securityGroupType) {
            return securityGroupType(Output.of(securityGroupType));
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
         * @param vpcId The ID of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public SecurityGroupArgs build() {
            return $;
        }
    }

}
