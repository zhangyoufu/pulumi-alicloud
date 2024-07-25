// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.maxcompute.inputs;

import com.pulumi.alicloud.maxcompute.inputs.ProjectSecurityPropertiesProjectProtectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectSecurityPropertiesArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectSecurityPropertiesArgs Empty = new ProjectSecurityPropertiesArgs();

    /**
     * Set whether to enable the [Download permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/download-control), that is, set the ODPS. security.enabledownloadprivilege property.
     * 
     */
    @Import(name="enableDownloadPrivilege")
    private @Nullable Output<Boolean> enableDownloadPrivilege;

    /**
     * @return Set whether to enable the [Download permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/download-control), that is, set the ODPS. security.enabledownloadprivilege property.
     * 
     */
    public Optional<Output<Boolean>> enableDownloadPrivilege() {
        return Optional.ofNullable(this.enableDownloadPrivilege);
    }

    /**
     * Set whether to use the [Label permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/label-based-access-control), that is, set the LabelSecurity attribute, which is not used by default.
     * 
     */
    @Import(name="labelSecurity")
    private @Nullable Output<Boolean> labelSecurity;

    /**
     * @return Set whether to use the [Label permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/label-based-access-control), that is, set the LabelSecurity attribute, which is not used by default.
     * 
     */
    public Optional<Output<Boolean>> labelSecurity() {
        return Optional.ofNullable(this.labelSecurity);
    }

    /**
     * Sets whether to allow the creator of the object to have access to the object, I .e. sets the attribute. The default is the allowed state.
     * 
     */
    @Import(name="objectCreatorHasAccessPermission")
    private @Nullable Output<Boolean> objectCreatorHasAccessPermission;

    /**
     * @return Sets whether to allow the creator of the object to have access to the object, I .e. sets the attribute. The default is the allowed state.
     * 
     */
    public Optional<Output<Boolean>> objectCreatorHasAccessPermission() {
        return Optional.ofNullable(this.objectCreatorHasAccessPermission);
    }

    /**
     * The ObjectCreatorHasGrantPermission attribute is set to allow the object creator to have the authorization permission on the object. The default is the allowed state.
     * 
     */
    @Import(name="objectCreatorHasGrantPermission")
    private @Nullable Output<Boolean> objectCreatorHasGrantPermission;

    /**
     * @return The ObjectCreatorHasGrantPermission attribute is set to allow the object creator to have the authorization permission on the object. The default is the allowed state.
     * 
     */
    public Optional<Output<Boolean>> objectCreatorHasGrantPermission() {
        return Optional.ofNullable(this.objectCreatorHasGrantPermission);
    }

    /**
     * Project protection. See `project_protection` below.
     * 
     */
    @Import(name="projectProtection")
    private @Nullable Output<ProjectSecurityPropertiesProjectProtectionArgs> projectProtection;

    /**
     * @return Project protection. See `project_protection` below.
     * 
     */
    public Optional<Output<ProjectSecurityPropertiesProjectProtectionArgs>> projectProtection() {
        return Optional.ofNullable(this.projectProtection);
    }

    /**
     * Set whether to use the [ACL permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/maxcompute-permissions), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
     * 
     */
    @Import(name="usingAcl")
    private @Nullable Output<Boolean> usingAcl;

    /**
     * @return Set whether to use the [ACL permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/maxcompute-permissions), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
     * 
     */
    public Optional<Output<Boolean>> usingAcl() {
        return Optional.ofNullable(this.usingAcl);
    }

    /**
     * Set whether to use the Policy permission control function (https://www.alibabacloud.com/help/en/maxcompute/user-guide/policy-based-access-control-1), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
     * 
     */
    @Import(name="usingPolicy")
    private @Nullable Output<Boolean> usingPolicy;

    /**
     * @return Set whether to use the Policy permission control function (https://www.alibabacloud.com/help/en/maxcompute/user-guide/policy-based-access-control-1), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
     * 
     */
    public Optional<Output<Boolean>> usingPolicy() {
        return Optional.ofNullable(this.usingPolicy);
    }

    private ProjectSecurityPropertiesArgs() {}

    private ProjectSecurityPropertiesArgs(ProjectSecurityPropertiesArgs $) {
        this.enableDownloadPrivilege = $.enableDownloadPrivilege;
        this.labelSecurity = $.labelSecurity;
        this.objectCreatorHasAccessPermission = $.objectCreatorHasAccessPermission;
        this.objectCreatorHasGrantPermission = $.objectCreatorHasGrantPermission;
        this.projectProtection = $.projectProtection;
        this.usingAcl = $.usingAcl;
        this.usingPolicy = $.usingPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectSecurityPropertiesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectSecurityPropertiesArgs $;

        public Builder() {
            $ = new ProjectSecurityPropertiesArgs();
        }

        public Builder(ProjectSecurityPropertiesArgs defaults) {
            $ = new ProjectSecurityPropertiesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableDownloadPrivilege Set whether to enable the [Download permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/download-control), that is, set the ODPS. security.enabledownloadprivilege property.
         * 
         * @return builder
         * 
         */
        public Builder enableDownloadPrivilege(@Nullable Output<Boolean> enableDownloadPrivilege) {
            $.enableDownloadPrivilege = enableDownloadPrivilege;
            return this;
        }

        /**
         * @param enableDownloadPrivilege Set whether to enable the [Download permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/download-control), that is, set the ODPS. security.enabledownloadprivilege property.
         * 
         * @return builder
         * 
         */
        public Builder enableDownloadPrivilege(Boolean enableDownloadPrivilege) {
            return enableDownloadPrivilege(Output.of(enableDownloadPrivilege));
        }

        /**
         * @param labelSecurity Set whether to use the [Label permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/label-based-access-control), that is, set the LabelSecurity attribute, which is not used by default.
         * 
         * @return builder
         * 
         */
        public Builder labelSecurity(@Nullable Output<Boolean> labelSecurity) {
            $.labelSecurity = labelSecurity;
            return this;
        }

        /**
         * @param labelSecurity Set whether to use the [Label permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/label-based-access-control), that is, set the LabelSecurity attribute, which is not used by default.
         * 
         * @return builder
         * 
         */
        public Builder labelSecurity(Boolean labelSecurity) {
            return labelSecurity(Output.of(labelSecurity));
        }

        /**
         * @param objectCreatorHasAccessPermission Sets whether to allow the creator of the object to have access to the object, I .e. sets the attribute. The default is the allowed state.
         * 
         * @return builder
         * 
         */
        public Builder objectCreatorHasAccessPermission(@Nullable Output<Boolean> objectCreatorHasAccessPermission) {
            $.objectCreatorHasAccessPermission = objectCreatorHasAccessPermission;
            return this;
        }

        /**
         * @param objectCreatorHasAccessPermission Sets whether to allow the creator of the object to have access to the object, I .e. sets the attribute. The default is the allowed state.
         * 
         * @return builder
         * 
         */
        public Builder objectCreatorHasAccessPermission(Boolean objectCreatorHasAccessPermission) {
            return objectCreatorHasAccessPermission(Output.of(objectCreatorHasAccessPermission));
        }

        /**
         * @param objectCreatorHasGrantPermission The ObjectCreatorHasGrantPermission attribute is set to allow the object creator to have the authorization permission on the object. The default is the allowed state.
         * 
         * @return builder
         * 
         */
        public Builder objectCreatorHasGrantPermission(@Nullable Output<Boolean> objectCreatorHasGrantPermission) {
            $.objectCreatorHasGrantPermission = objectCreatorHasGrantPermission;
            return this;
        }

        /**
         * @param objectCreatorHasGrantPermission The ObjectCreatorHasGrantPermission attribute is set to allow the object creator to have the authorization permission on the object. The default is the allowed state.
         * 
         * @return builder
         * 
         */
        public Builder objectCreatorHasGrantPermission(Boolean objectCreatorHasGrantPermission) {
            return objectCreatorHasGrantPermission(Output.of(objectCreatorHasGrantPermission));
        }

        /**
         * @param projectProtection Project protection. See `project_protection` below.
         * 
         * @return builder
         * 
         */
        public Builder projectProtection(@Nullable Output<ProjectSecurityPropertiesProjectProtectionArgs> projectProtection) {
            $.projectProtection = projectProtection;
            return this;
        }

        /**
         * @param projectProtection Project protection. See `project_protection` below.
         * 
         * @return builder
         * 
         */
        public Builder projectProtection(ProjectSecurityPropertiesProjectProtectionArgs projectProtection) {
            return projectProtection(Output.of(projectProtection));
        }

        /**
         * @param usingAcl Set whether to use the [ACL permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/maxcompute-permissions), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
         * 
         * @return builder
         * 
         */
        public Builder usingAcl(@Nullable Output<Boolean> usingAcl) {
            $.usingAcl = usingAcl;
            return this;
        }

        /**
         * @param usingAcl Set whether to use the [ACL permission control function](https://www.alibabacloud.com/help/en/maxcompute/user-guide/maxcompute-permissions), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
         * 
         * @return builder
         * 
         */
        public Builder usingAcl(Boolean usingAcl) {
            return usingAcl(Output.of(usingAcl));
        }

        /**
         * @param usingPolicy Set whether to use the Policy permission control function (https://www.alibabacloud.com/help/en/maxcompute/user-guide/policy-based-access-control-1), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
         * 
         * @return builder
         * 
         */
        public Builder usingPolicy(@Nullable Output<Boolean> usingPolicy) {
            $.usingPolicy = usingPolicy;
            return this;
        }

        /**
         * @param usingPolicy Set whether to use the Policy permission control function (https://www.alibabacloud.com/help/en/maxcompute/user-guide/policy-based-access-control-1), that is, set the CheckPermissionUsingACL attribute, which is in use by default.
         * 
         * @return builder
         * 
         */
        public Builder usingPolicy(Boolean usingPolicy) {
            return usingPolicy(Output.of(usingPolicy));
        }

        public ProjectSecurityPropertiesArgs build() {
            return $;
        }
    }

}
