// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGlobalDatabaseNetworksPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGlobalDatabaseNetworksPlainArgs Empty = new GetGlobalDatabaseNetworksPlainArgs();

    /**
     * The ID of the PolarDB cluster.
     * 
     */
    @Import(name="dbClusterId")
    private @Nullable String dbClusterId;

    /**
     * @return The ID of the PolarDB cluster.
     * 
     */
    public Optional<String> dbClusterId() {
        return Optional.ofNullable(this.dbClusterId);
    }

    /**
     * The description of the Global Database Network.
     * 
     */
    @Import(name="description")
    private @Nullable String description;

    /**
     * @return The description of the Global Database Network.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the Global Database Network.
     * 
     */
    @Import(name="gdnId")
    private @Nullable String gdnId;

    /**
     * @return The ID of the Global Database Network.
     * 
     */
    public Optional<String> gdnId() {
        return Optional.ofNullable(this.gdnId);
    }

    /**
     * A list of Global Database Network IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Global Database Network IDs.
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

    @Import(name="pageNumber")
    private @Nullable Integer pageNumber;

    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Integer pageSize;

    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    /**
     * The status of the Global Database Network.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the Global Database Network.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetGlobalDatabaseNetworksPlainArgs() {}

    private GetGlobalDatabaseNetworksPlainArgs(GetGlobalDatabaseNetworksPlainArgs $) {
        this.dbClusterId = $.dbClusterId;
        this.description = $.description;
        this.gdnId = $.gdnId;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGlobalDatabaseNetworksPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGlobalDatabaseNetworksPlainArgs $;

        public Builder() {
            $ = new GetGlobalDatabaseNetworksPlainArgs();
        }

        public Builder(GetGlobalDatabaseNetworksPlainArgs defaults) {
            $ = new GetGlobalDatabaseNetworksPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbClusterId The ID of the PolarDB cluster.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(@Nullable String dbClusterId) {
            $.dbClusterId = dbClusterId;
            return this;
        }

        /**
         * @param description The description of the Global Database Network.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        /**
         * @param gdnId The ID of the Global Database Network.
         * 
         * @return builder
         * 
         */
        public Builder gdnId(@Nullable String gdnId) {
            $.gdnId = gdnId;
            return this;
        }

        /**
         * @param ids A list of Global Database Network IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Global Database Network IDs.
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

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        /**
         * @param status The status of the Global Database Network.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetGlobalDatabaseNetworksPlainArgs build() {
            return $;
        }
    }

}
