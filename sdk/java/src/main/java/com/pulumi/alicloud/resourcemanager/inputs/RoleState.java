// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RoleState extends com.pulumi.resources.ResourceArgs {

    public static final RoleState Empty = new RoleState();

    /**
     * The resource descriptor of the role.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The resource descriptor of the role.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The content of the permissions strategy that plays a role.
     * 
     */
    @Import(name="assumeRolePolicyDocument")
    private @Nullable Output<String> assumeRolePolicyDocument;

    /**
     * @return The content of the permissions strategy that plays a role.
     * 
     */
    public Optional<Output<String>> assumeRolePolicyDocument() {
        return Optional.ofNullable(this.assumeRolePolicyDocument);
    }

    /**
     * The description of the Resource Manager role.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the Resource Manager role.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
     * 
     */
    @Import(name="maxSessionDuration")
    private @Nullable Output<Integer> maxSessionDuration;

    /**
     * @return Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
     * 
     */
    public Optional<Output<Integer>> maxSessionDuration() {
        return Optional.ofNullable(this.maxSessionDuration);
    }

    /**
     * This ID of Resource Manager role. The value is set to `role_name`.
     * 
     */
    @Import(name="roleId")
    private @Nullable Output<String> roleId;

    /**
     * @return This ID of Resource Manager role. The value is set to `role_name`.
     * 
     */
    public Optional<Output<String>> roleId() {
        return Optional.ofNullable(this.roleId);
    }

    /**
     * Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots &#34;.&#34; and dashes &#34;-&#34;.
     * 
     */
    @Import(name="roleName")
    private @Nullable Output<String> roleName;

    /**
     * @return Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots &#34;.&#34; and dashes &#34;-&#34;.
     * 
     */
    public Optional<Output<String>> roleName() {
        return Optional.ofNullable(this.roleName);
    }

    /**
     * Role update time.
     * 
     */
    @Import(name="updateDate")
    private @Nullable Output<String> updateDate;

    /**
     * @return Role update time.
     * 
     */
    public Optional<Output<String>> updateDate() {
        return Optional.ofNullable(this.updateDate);
    }

    private RoleState() {}

    private RoleState(RoleState $) {
        this.arn = $.arn;
        this.assumeRolePolicyDocument = $.assumeRolePolicyDocument;
        this.description = $.description;
        this.maxSessionDuration = $.maxSessionDuration;
        this.roleId = $.roleId;
        this.roleName = $.roleName;
        this.updateDate = $.updateDate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RoleState $;

        public Builder() {
            $ = new RoleState();
        }

        public Builder(RoleState defaults) {
            $ = new RoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The resource descriptor of the role.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The resource descriptor of the role.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param assumeRolePolicyDocument The content of the permissions strategy that plays a role.
         * 
         * @return builder
         * 
         */
        public Builder assumeRolePolicyDocument(@Nullable Output<String> assumeRolePolicyDocument) {
            $.assumeRolePolicyDocument = assumeRolePolicyDocument;
            return this;
        }

        /**
         * @param assumeRolePolicyDocument The content of the permissions strategy that plays a role.
         * 
         * @return builder
         * 
         */
        public Builder assumeRolePolicyDocument(String assumeRolePolicyDocument) {
            return assumeRolePolicyDocument(Output.of(assumeRolePolicyDocument));
        }

        /**
         * @param description The description of the Resource Manager role.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the Resource Manager role.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param maxSessionDuration Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
         * 
         * @return builder
         * 
         */
        public Builder maxSessionDuration(@Nullable Output<Integer> maxSessionDuration) {
            $.maxSessionDuration = maxSessionDuration;
            return this;
        }

        /**
         * @param maxSessionDuration Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
         * 
         * @return builder
         * 
         */
        public Builder maxSessionDuration(Integer maxSessionDuration) {
            return maxSessionDuration(Output.of(maxSessionDuration));
        }

        /**
         * @param roleId This ID of Resource Manager role. The value is set to `role_name`.
         * 
         * @return builder
         * 
         */
        public Builder roleId(@Nullable Output<String> roleId) {
            $.roleId = roleId;
            return this;
        }

        /**
         * @param roleId This ID of Resource Manager role. The value is set to `role_name`.
         * 
         * @return builder
         * 
         */
        public Builder roleId(String roleId) {
            return roleId(Output.of(roleId));
        }

        /**
         * @param roleName Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots &#34;.&#34; and dashes &#34;-&#34;.
         * 
         * @return builder
         * 
         */
        public Builder roleName(@Nullable Output<String> roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param roleName Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots &#34;.&#34; and dashes &#34;-&#34;.
         * 
         * @return builder
         * 
         */
        public Builder roleName(String roleName) {
            return roleName(Output.of(roleName));
        }

        /**
         * @param updateDate Role update time.
         * 
         * @return builder
         * 
         */
        public Builder updateDate(@Nullable Output<String> updateDate) {
            $.updateDate = updateDate;
            return this;
        }

        /**
         * @param updateDate Role update time.
         * 
         * @return builder
         * 
         */
        public Builder updateDate(String updateDate) {
            return updateDate(Output.of(updateDate));
        }

        public RoleState build() {
            return $;
        }
    }

}
