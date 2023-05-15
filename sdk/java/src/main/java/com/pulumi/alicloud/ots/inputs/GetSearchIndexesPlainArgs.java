// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSearchIndexesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSearchIndexesPlainArgs Empty = new GetSearchIndexesPlainArgs();

    /**
     * A list of search index IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of search index IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The name of OTS instance.
     * 
     */
    @Import(name="instanceName", required=true)
    private String instanceName;

    /**
     * @return The name of OTS instance.
     * 
     */
    public String instanceName() {
        return this.instanceName;
    }

    /**
     * A regex string to filter results by search index name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by search index name.
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
     * The name of OTS table.
     * 
     */
    @Import(name="tableName", required=true)
    private String tableName;

    /**
     * @return The name of OTS table.
     * 
     */
    public String tableName() {
        return this.tableName;
    }

    private GetSearchIndexesPlainArgs() {}

    private GetSearchIndexesPlainArgs(GetSearchIndexesPlainArgs $) {
        this.ids = $.ids;
        this.instanceName = $.instanceName;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.tableName = $.tableName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSearchIndexesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSearchIndexesPlainArgs $;

        public Builder() {
            $ = new GetSearchIndexesPlainArgs();
        }

        public Builder(GetSearchIndexesPlainArgs defaults) {
            $ = new GetSearchIndexesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of search index IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of search index IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceName The name of OTS instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by search index name.
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
         * @param tableName The name of OTS table.
         * 
         * @return builder
         * 
         */
        public Builder tableName(String tableName) {
            $.tableName = tableName;
            return this;
        }

        public GetSearchIndexesPlainArgs build() {
            $.instanceName = Objects.requireNonNull($.instanceName, "expected parameter 'instanceName' to be non-null");
            $.tableName = Objects.requireNonNull($.tableName, "expected parameter 'tableName' to be non-null");
            return $;
        }
    }

}
