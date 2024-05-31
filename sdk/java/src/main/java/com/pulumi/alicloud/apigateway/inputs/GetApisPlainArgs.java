// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetApisPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetApisPlainArgs Empty = new GetApisPlainArgs();

    /**
     * The ID of the API.
     * 
     */
    @Import(name="apiId")
    private @Nullable String apiId;

    /**
     * @return The ID of the API.
     * 
     */
    public Optional<String> apiId() {
        return Optional.ofNullable(this.apiId);
    }

    /**
     * The ID of the API group.
     * 
     */
    @Import(name="groupId")
    private @Nullable String groupId;

    /**
     * @return The ID of the API group.
     * 
     */
    public Optional<String> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * A list of API IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of API IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by API name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by API name.
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

    private GetApisPlainArgs() {}

    private GetApisPlainArgs(GetApisPlainArgs $) {
        this.apiId = $.apiId;
        this.groupId = $.groupId;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetApisPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetApisPlainArgs $;

        public Builder() {
            $ = new GetApisPlainArgs();
        }

        public Builder(GetApisPlainArgs defaults) {
            $ = new GetApisPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiId The ID of the API.
         * 
         * @return builder
         * 
         */
        public Builder apiId(@Nullable String apiId) {
            $.apiId = apiId;
            return this;
        }

        /**
         * @param groupId The ID of the API group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable String groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param ids A list of API IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of API IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by API name.
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

        public GetApisPlainArgs build() {
            return $;
        }
    }

}
