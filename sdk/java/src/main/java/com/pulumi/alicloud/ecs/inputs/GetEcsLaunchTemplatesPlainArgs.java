// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEcsLaunchTemplatesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEcsLaunchTemplatesPlainArgs Empty = new GetEcsLaunchTemplatesPlainArgs();

    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    /**
     * @return Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of Launch Template IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Launch Template IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The Launch Template Name.
     * 
     */
    @Import(name="launchTemplateName")
    private @Nullable String launchTemplateName;

    /**
     * @return The Launch Template Name.
     * 
     */
    public Optional<String> launchTemplateName() {
        return Optional.ofNullable(this.launchTemplateName);
    }

    /**
     * A regex string to filter results by Launch Template name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Launch Template name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The template resource group id.
     * 
     */
    @Import(name="templateResourceGroupId")
    private @Nullable String templateResourceGroupId;

    /**
     * @return The template resource group id.
     * 
     */
    public Optional<String> templateResourceGroupId() {
        return Optional.ofNullable(this.templateResourceGroupId);
    }

    @Import(name="templateTags")
    private @Nullable Map<String,String> templateTags;

    public Optional<Map<String,String>> templateTags() {
        return Optional.ofNullable(this.templateTags);
    }

    private GetEcsLaunchTemplatesPlainArgs() {}

    private GetEcsLaunchTemplatesPlainArgs(GetEcsLaunchTemplatesPlainArgs $) {
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.launchTemplateName = $.launchTemplateName;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.templateResourceGroupId = $.templateResourceGroupId;
        this.templateTags = $.templateTags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEcsLaunchTemplatesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEcsLaunchTemplatesPlainArgs $;

        public Builder() {
            $ = new GetEcsLaunchTemplatesPlainArgs();
        }

        public Builder(GetEcsLaunchTemplatesPlainArgs defaults) {
            $ = new GetEcsLaunchTemplatesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableDetails Default to `false`. Set it to `true` can output more details about resource attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids A list of Launch Template IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Launch Template IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param launchTemplateName The Launch Template Name.
         * 
         * @return builder
         * 
         */
        public Builder launchTemplateName(@Nullable String launchTemplateName) {
            $.launchTemplateName = launchTemplateName;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Launch Template name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param templateResourceGroupId The template resource group id.
         * 
         * @return builder
         * 
         */
        public Builder templateResourceGroupId(@Nullable String templateResourceGroupId) {
            $.templateResourceGroupId = templateResourceGroupId;
            return this;
        }

        public Builder templateTags(@Nullable Map<String,String> templateTags) {
            $.templateTags = templateTags;
            return this;
        }

        public GetEcsLaunchTemplatesPlainArgs build() {
            return $;
        }
    }

}
