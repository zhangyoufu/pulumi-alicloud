// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetHealthCheckTemplatesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetHealthCheckTemplatesArgs Empty = new GetHealthCheckTemplatesArgs();

    /**
     * The health check template ids.
     * 
     */
    @Import(name="healthCheckTemplateIds")
    private @Nullable Output<List<String>> healthCheckTemplateIds;

    /**
     * @return The health check template ids.
     * 
     */
    public Optional<Output<List<String>>> healthCheckTemplateIds() {
        return Optional.ofNullable(this.healthCheckTemplateIds);
    }

    /**
     * The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    @Import(name="healthCheckTemplateName")
    private @Nullable Output<String> healthCheckTemplateName;

    /**
     * @return The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    public Optional<Output<String>> healthCheckTemplateName() {
        return Optional.ofNullable(this.healthCheckTemplateName);
    }

    /**
     * A list of Health Check Template IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Health Check Template IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Health Check Template name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Health Check Template name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetHealthCheckTemplatesArgs() {}

    private GetHealthCheckTemplatesArgs(GetHealthCheckTemplatesArgs $) {
        this.healthCheckTemplateIds = $.healthCheckTemplateIds;
        this.healthCheckTemplateName = $.healthCheckTemplateName;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetHealthCheckTemplatesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetHealthCheckTemplatesArgs $;

        public Builder() {
            $ = new GetHealthCheckTemplatesArgs();
        }

        public Builder(GetHealthCheckTemplatesArgs defaults) {
            $ = new GetHealthCheckTemplatesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param healthCheckTemplateIds The health check template ids.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckTemplateIds(@Nullable Output<List<String>> healthCheckTemplateIds) {
            $.healthCheckTemplateIds = healthCheckTemplateIds;
            return this;
        }

        /**
         * @param healthCheckTemplateIds The health check template ids.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckTemplateIds(List<String> healthCheckTemplateIds) {
            return healthCheckTemplateIds(Output.of(healthCheckTemplateIds));
        }

        /**
         * @param healthCheckTemplateIds The health check template ids.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckTemplateIds(String... healthCheckTemplateIds) {
            return healthCheckTemplateIds(List.of(healthCheckTemplateIds));
        }

        /**
         * @param healthCheckTemplateName The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckTemplateName(@Nullable Output<String> healthCheckTemplateName) {
            $.healthCheckTemplateName = healthCheckTemplateName;
            return this;
        }

        /**
         * @param healthCheckTemplateName The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckTemplateName(String healthCheckTemplateName) {
            return healthCheckTemplateName(Output.of(healthCheckTemplateName));
        }

        /**
         * @param ids A list of Health Check Template IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Health Check Template IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Health Check Template IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Health Check Template name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Health Check Template name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        public GetHealthCheckTemplatesArgs build() {
            return $;
        }
    }

}
