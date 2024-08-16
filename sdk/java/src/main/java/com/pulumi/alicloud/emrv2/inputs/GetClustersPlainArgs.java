// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emrv2.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetClustersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetClustersPlainArgs Empty = new GetClustersPlainArgs();

    /**
     * The cluster name.
     * 
     */
    @Import(name="clusterName")
    private @Nullable String clusterName;

    /**
     * @return The cluster name.
     * 
     */
    public Optional<String> clusterName() {
        return Optional.ofNullable(this.clusterName);
    }

    /**
     * The cluster states.
     * 
     */
    @Import(name="clusterStates")
    private @Nullable List<String> clusterStates;

    /**
     * @return The cluster states.
     * 
     */
    public Optional<List<String>> clusterStates() {
        return Optional.ofNullable(this.clusterStates);
    }

    /**
     * The cluster types.
     * 
     */
    @Import(name="clusterTypes")
    private @Nullable List<String> clusterTypes;

    /**
     * @return The cluster types.
     * 
     */
    public Optional<List<String>> clusterTypes() {
        return Optional.ofNullable(this.clusterTypes);
    }

    /**
     * A list of Cluster IDS.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Cluster IDS.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The max results is used to list clusters for next page.
     * 
     */
    @Import(name="maxResults")
    private @Nullable Integer maxResults;

    /**
     * @return The max results is used to list clusters for next page.
     * 
     */
    public Optional<Integer> maxResults() {
        return Optional.ofNullable(this.maxResults);
    }

    /**
     * A regex string to filter results by Cluster name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Cluster name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The next token is used to list clusters for next page.
     * 
     */
    @Import(name="nextToken")
    private @Nullable String nextToken;

    /**
     * @return The next token is used to list clusters for next page.
     * 
     */
    public Optional<String> nextToken() {
        return Optional.ofNullable(this.nextToken);
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
     * The cluster payment types.
     * 
     */
    @Import(name="paymentTypes")
    private @Nullable List<String> paymentTypes;

    /**
     * @return The cluster payment types.
     * 
     */
    public Optional<List<String>> paymentTypes() {
        return Optional.ofNullable(this.paymentTypes);
    }

    /**
     * The Resource Group ID.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable String resourceGroupId;

    /**
     * @return The Resource Group ID.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,String> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Map<String,String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GetClustersPlainArgs() {}

    private GetClustersPlainArgs(GetClustersPlainArgs $) {
        this.clusterName = $.clusterName;
        this.clusterStates = $.clusterStates;
        this.clusterTypes = $.clusterTypes;
        this.ids = $.ids;
        this.maxResults = $.maxResults;
        this.nameRegex = $.nameRegex;
        this.nextToken = $.nextToken;
        this.outputFile = $.outputFile;
        this.paymentTypes = $.paymentTypes;
        this.resourceGroupId = $.resourceGroupId;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetClustersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetClustersPlainArgs $;

        public Builder() {
            $ = new GetClustersPlainArgs();
        }

        public Builder(GetClustersPlainArgs defaults) {
            $ = new GetClustersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterName The cluster name.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(@Nullable String clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param clusterStates The cluster states.
         * 
         * @return builder
         * 
         */
        public Builder clusterStates(@Nullable List<String> clusterStates) {
            $.clusterStates = clusterStates;
            return this;
        }

        /**
         * @param clusterStates The cluster states.
         * 
         * @return builder
         * 
         */
        public Builder clusterStates(String... clusterStates) {
            return clusterStates(List.of(clusterStates));
        }

        /**
         * @param clusterTypes The cluster types.
         * 
         * @return builder
         * 
         */
        public Builder clusterTypes(@Nullable List<String> clusterTypes) {
            $.clusterTypes = clusterTypes;
            return this;
        }

        /**
         * @param clusterTypes The cluster types.
         * 
         * @return builder
         * 
         */
        public Builder clusterTypes(String... clusterTypes) {
            return clusterTypes(List.of(clusterTypes));
        }

        /**
         * @param ids A list of Cluster IDS.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Cluster IDS.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param maxResults The max results is used to list clusters for next page.
         * 
         * @return builder
         * 
         */
        public Builder maxResults(@Nullable Integer maxResults) {
            $.maxResults = maxResults;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Cluster name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nextToken The next token is used to list clusters for next page.
         * 
         * @return builder
         * 
         */
        public Builder nextToken(@Nullable String nextToken) {
            $.nextToken = nextToken;
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
         * @param paymentTypes The cluster payment types.
         * 
         * @return builder
         * 
         */
        public Builder paymentTypes(@Nullable List<String> paymentTypes) {
            $.paymentTypes = paymentTypes;
            return this;
        }

        /**
         * @param paymentTypes The cluster payment types.
         * 
         * @return builder
         * 
         */
        public Builder paymentTypes(String... paymentTypes) {
            return paymentTypes(List.of(paymentTypes));
        }

        /**
         * @param resourceGroupId The Resource Group ID.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,String> tags) {
            $.tags = tags;
            return this;
        }

        public GetClustersPlainArgs build() {
            return $;
        }
    }

}
