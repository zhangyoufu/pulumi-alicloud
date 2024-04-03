// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class KubernetesPermissionPermission {
    /**
     * @return The ID of the cluster that you want to manage, When `role_type` value is `all-clusters`, the value of `role_type` must be null.
     * 
     */
    private String cluster;
    /**
     * @return Specifies whether to perform a custom authorization. To perform a custom authorization, set `role_name` to a custom cluster role.
     * 
     */
    private @Nullable Boolean isCustom;
    /**
     * @return Specifies whether the permissions are granted to a RAM role. When `uid` is ram role id, the value of `is_ram_role` must be `true`.
     * 
     */
    private @Nullable Boolean isRamRole;
    /**
     * @return The namespace to which the permissions are scoped. This parameter is required only if you set role_type to namespace.
     * 
     */
    private @Nullable String namespace;
    /**
     * @return Specifies the predefined role that you want to assign. Valid values `admin`, `ops`, `dev`, `restricted` and the custom cluster roles.
     * 
     */
    private String roleName;
    /**
     * @return The authorization type. Valid values `cluster`, `namespace` and `all-clusters`.
     * 
     */
    private String roleType;

    private KubernetesPermissionPermission() {}
    /**
     * @return The ID of the cluster that you want to manage, When `role_type` value is `all-clusters`, the value of `role_type` must be null.
     * 
     */
    public String cluster() {
        return this.cluster;
    }
    /**
     * @return Specifies whether to perform a custom authorization. To perform a custom authorization, set `role_name` to a custom cluster role.
     * 
     */
    public Optional<Boolean> isCustom() {
        return Optional.ofNullable(this.isCustom);
    }
    /**
     * @return Specifies whether the permissions are granted to a RAM role. When `uid` is ram role id, the value of `is_ram_role` must be `true`.
     * 
     */
    public Optional<Boolean> isRamRole() {
        return Optional.ofNullable(this.isRamRole);
    }
    /**
     * @return The namespace to which the permissions are scoped. This parameter is required only if you set role_type to namespace.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    /**
     * @return Specifies the predefined role that you want to assign. Valid values `admin`, `ops`, `dev`, `restricted` and the custom cluster roles.
     * 
     */
    public String roleName() {
        return this.roleName;
    }
    /**
     * @return The authorization type. Valid values `cluster`, `namespace` and `all-clusters`.
     * 
     */
    public String roleType() {
        return this.roleType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubernetesPermissionPermission defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cluster;
        private @Nullable Boolean isCustom;
        private @Nullable Boolean isRamRole;
        private @Nullable String namespace;
        private String roleName;
        private String roleType;
        public Builder() {}
        public Builder(KubernetesPermissionPermission defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cluster = defaults.cluster;
    	      this.isCustom = defaults.isCustom;
    	      this.isRamRole = defaults.isRamRole;
    	      this.namespace = defaults.namespace;
    	      this.roleName = defaults.roleName;
    	      this.roleType = defaults.roleType;
        }

        @CustomType.Setter
        public Builder cluster(String cluster) {
            if (cluster == null) {
              throw new MissingRequiredPropertyException("KubernetesPermissionPermission", "cluster");
            }
            this.cluster = cluster;
            return this;
        }
        @CustomType.Setter
        public Builder isCustom(@Nullable Boolean isCustom) {

            this.isCustom = isCustom;
            return this;
        }
        @CustomType.Setter
        public Builder isRamRole(@Nullable Boolean isRamRole) {

            this.isRamRole = isRamRole;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {

            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder roleName(String roleName) {
            if (roleName == null) {
              throw new MissingRequiredPropertyException("KubernetesPermissionPermission", "roleName");
            }
            this.roleName = roleName;
            return this;
        }
        @CustomType.Setter
        public Builder roleType(String roleType) {
            if (roleType == null) {
              throw new MissingRequiredPropertyException("KubernetesPermissionPermission", "roleType");
            }
            this.roleType = roleType;
            return this;
        }
        public KubernetesPermissionPermission build() {
            final var _resultValue = new KubernetesPermissionPermission();
            _resultValue.cluster = cluster;
            _resultValue.isCustom = isCustom;
            _resultValue.isRamRole = isRamRole;
            _resultValue.namespace = namespace;
            _resultValue.roleName = roleName;
            _resultValue.roleType = roleType;
            return _resultValue;
        }
    }
}
