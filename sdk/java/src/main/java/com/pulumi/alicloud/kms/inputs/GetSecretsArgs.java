// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kms.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSecretsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecretsArgs Empty = new GetSecretsArgs();

    /**
     * Default to `false`. Set it to true can output more details.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Output<Boolean> enableDetails;

    /**
     * @return Default to `false`. Set it to true can output more details.
     * 
     */
    public Optional<Output<Boolean>> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * Whether to include the predetermined resource tag in the return value. Default to `false`.
     * 
     */
    @Import(name="fetchTags")
    private @Nullable Output<Boolean> fetchTags;

    /**
     * @return Whether to include the predetermined resource tag in the return value. Default to `false`.
     * 
     */
    public Optional<Output<Boolean>> fetchTags() {
        return Optional.ofNullable(this.fetchTags);
    }

    /**
     * The secret filter. The filter consists of one or more key-value pairs.
     * More details see API [ListSecrets](https://www.alibabacloud.com/help/en/key-management-service/latest/listsecrets).
     * 
     */
    @Import(name="filters")
    private @Nullable Output<String> filters;

    /**
     * @return The secret filter. The filter consists of one or more key-value pairs.
     * More details see API [ListSecrets](https://www.alibabacloud.com/help/en/key-management-service/latest/listsecrets).
     * 
     */
    public Optional<Output<String>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * A list of KMS Secret ids. The value is same as KMS secret_name.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of KMS Secret ids. The value is same as KMS secret_name.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter the results by the KMS secret_name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter the results by the KMS secret_name.
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

    private GetSecretsArgs() {}

    private GetSecretsArgs(GetSecretsArgs $) {
        this.enableDetails = $.enableDetails;
        this.fetchTags = $.fetchTags;
        this.filters = $.filters;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSecretsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecretsArgs $;

        public Builder() {
            $ = new GetSecretsArgs();
        }

        public Builder(GetSecretsArgs defaults) {
            $ = new GetSecretsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableDetails Default to `false`. Set it to true can output more details.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Output<Boolean> enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param enableDetails Default to `false`. Set it to true can output more details.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(Boolean enableDetails) {
            return enableDetails(Output.of(enableDetails));
        }

        /**
         * @param fetchTags Whether to include the predetermined resource tag in the return value. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder fetchTags(@Nullable Output<Boolean> fetchTags) {
            $.fetchTags = fetchTags;
            return this;
        }

        /**
         * @param fetchTags Whether to include the predetermined resource tag in the return value. Default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder fetchTags(Boolean fetchTags) {
            return fetchTags(Output.of(fetchTags));
        }

        /**
         * @param filters The secret filter. The filter consists of one or more key-value pairs.
         * More details see API [ListSecrets](https://www.alibabacloud.com/help/en/key-management-service/latest/listsecrets).
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable Output<String> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters The secret filter. The filter consists of one or more key-value pairs.
         * More details see API [ListSecrets](https://www.alibabacloud.com/help/en/key-management-service/latest/listsecrets).
         * 
         * @return builder
         * 
         */
        public Builder filters(String filters) {
            return filters(Output.of(filters));
        }

        /**
         * @param ids A list of KMS Secret ids. The value is same as KMS secret_name.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of KMS Secret ids. The value is same as KMS secret_name.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of KMS Secret ids. The value is same as KMS secret_name.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter the results by the KMS secret_name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter the results by the KMS secret_name.
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

        public GetSecretsArgs build() {
            return $;
        }
    }

}
