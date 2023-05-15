// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.clickhouse.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDbClustersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDbClustersPlainArgs Empty = new GetDbClustersPlainArgs();

    /**
     * The DBCluster description.
     * 
     */
    @Import(name="dbClusterDescription")
    private @Nullable String dbClusterDescription;

    /**
     * @return The DBCluster description.
     * 
     */
    public Optional<String> dbClusterDescription() {
        return Optional.ofNullable(this.dbClusterDescription);
    }

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
     * A list of DBCluster IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of DBCluster IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
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
     * The status of the DBCluster. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the DBCluster. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetDbClustersPlainArgs() {}

    private GetDbClustersPlainArgs(GetDbClustersPlainArgs $) {
        this.dbClusterDescription = $.dbClusterDescription;
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDbClustersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDbClustersPlainArgs $;

        public Builder() {
            $ = new GetDbClustersPlainArgs();
        }

        public Builder(GetDbClustersPlainArgs defaults) {
            $ = new GetDbClustersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbClusterDescription The DBCluster description.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterDescription(@Nullable String dbClusterDescription) {
            $.dbClusterDescription = dbClusterDescription;
            return this;
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
         * @param ids A list of DBCluster IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of DBCluster IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
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
         * @param status The status of the DBCluster. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetDbClustersPlainArgs build() {
            return $;
        }
    }

}
