// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSecretParametersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecretParametersArgs Empty = new GetSecretParametersArgs();

    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Output<Boolean> enableDetails;

    /**
     * @return Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    public Optional<Output<Boolean>> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of Secret Parameter IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Secret Parameter IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Secret Parameter name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Secret Parameter name.
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

    /**
     * The ID of the Resource Group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the Resource Group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The name of the secret parameter.
     * 
     */
    @Import(name="secretParameterName")
    private @Nullable Output<String> secretParameterName;

    /**
     * @return The name of the secret parameter.
     * 
     */
    public Optional<Output<String>> secretParameterName() {
        return Optional.ofNullable(this.secretParameterName);
    }

    @Import(name="sortField")
    private @Nullable Output<String> sortField;

    public Optional<Output<String>> sortField() {
        return Optional.ofNullable(this.sortField);
    }

    @Import(name="sortOrder")
    private @Nullable Output<String> sortOrder;

    public Optional<Output<String>> sortOrder() {
        return Optional.ofNullable(this.sortOrder);
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

    private GetSecretParametersArgs() {}

    private GetSecretParametersArgs(GetSecretParametersArgs $) {
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.resourceGroupId = $.resourceGroupId;
        this.secretParameterName = $.secretParameterName;
        this.sortField = $.sortField;
        this.sortOrder = $.sortOrder;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSecretParametersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecretParametersArgs $;

        public Builder() {
            $ = new GetSecretParametersArgs();
        }

        public Builder(GetSecretParametersArgs defaults) {
            $ = new GetSecretParametersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableDetails Default to `false`. Set it to `true` can output more details about resource attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Output<Boolean> enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param enableDetails Default to `false`. Set it to `true` can output more details about resource attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(Boolean enableDetails) {
            return enableDetails(Output.of(enableDetails));
        }

        /**
         * @param ids A list of Secret Parameter IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Secret Parameter IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Secret Parameter IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Secret Parameter name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Secret Parameter name.
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

        /**
         * @param resourceGroupId The ID of the Resource Group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the Resource Group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param secretParameterName The name of the secret parameter.
         * 
         * @return builder
         * 
         */
        public Builder secretParameterName(@Nullable Output<String> secretParameterName) {
            $.secretParameterName = secretParameterName;
            return this;
        }

        /**
         * @param secretParameterName The name of the secret parameter.
         * 
         * @return builder
         * 
         */
        public Builder secretParameterName(String secretParameterName) {
            return secretParameterName(Output.of(secretParameterName));
        }

        public Builder sortField(@Nullable Output<String> sortField) {
            $.sortField = sortField;
            return this;
        }

        public Builder sortField(String sortField) {
            return sortField(Output.of(sortField));
        }

        public Builder sortOrder(@Nullable Output<String> sortOrder) {
            $.sortOrder = sortOrder;
            return this;
        }

        public Builder sortOrder(String sortOrder) {
            return sortOrder(Output.of(sortOrder));
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

        public GetSecretParametersArgs build() {
            return $;
        }
    }

}
