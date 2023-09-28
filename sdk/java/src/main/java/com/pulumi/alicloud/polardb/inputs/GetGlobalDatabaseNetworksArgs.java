// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGlobalDatabaseNetworksArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGlobalDatabaseNetworksArgs Empty = new GetGlobalDatabaseNetworksArgs();

    /**
     * The ID of the cluster.
     * 
     */
    @Import(name="dbClusterId")
    private @Nullable Output<String> dbClusterId;

    /**
     * @return The ID of the cluster.
     * 
     */
    public Optional<Output<String>> dbClusterId() {
        return Optional.ofNullable(this.dbClusterId);
    }

    /**
     * The description of the Global Database Network.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the Global Database Network.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the Global Database Network.
     * 
     */
    @Import(name="gdnId")
    private @Nullable Output<String> gdnId;

    /**
     * @return The ID of the Global Database Network.
     * 
     */
    public Optional<Output<String>> gdnId() {
        return Optional.ofNullable(this.gdnId);
    }

    /**
     * A list of Global Database Network IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Global Database Network IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
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

    @Import(name="pageNumber")
    private @Nullable Output<Integer> pageNumber;

    public Optional<Output<Integer>> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Output<Integer> pageSize;

    public Optional<Output<Integer>> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    /**
     * The status of the Global Database Network. Valid values:
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the Global Database Network. Valid values:
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private GetGlobalDatabaseNetworksArgs() {}

    private GetGlobalDatabaseNetworksArgs(GetGlobalDatabaseNetworksArgs $) {
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
    public static Builder builder(GetGlobalDatabaseNetworksArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGlobalDatabaseNetworksArgs $;

        public Builder() {
            $ = new GetGlobalDatabaseNetworksArgs();
        }

        public Builder(GetGlobalDatabaseNetworksArgs defaults) {
            $ = new GetGlobalDatabaseNetworksArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbClusterId The ID of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(@Nullable Output<String> dbClusterId) {
            $.dbClusterId = dbClusterId;
            return this;
        }

        /**
         * @param dbClusterId The ID of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(String dbClusterId) {
            return dbClusterId(Output.of(dbClusterId));
        }

        /**
         * @param description The description of the Global Database Network.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the Global Database Network.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param gdnId The ID of the Global Database Network.
         * 
         * @return builder
         * 
         */
        public Builder gdnId(@Nullable Output<String> gdnId) {
            $.gdnId = gdnId;
            return this;
        }

        /**
         * @param gdnId The ID of the Global Database Network.
         * 
         * @return builder
         * 
         */
        public Builder gdnId(String gdnId) {
            return gdnId(Output.of(gdnId));
        }

        /**
         * @param ids A list of Global Database Network IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Global Database Network IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
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

        public Builder pageNumber(@Nullable Output<Integer> pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageNumber(Integer pageNumber) {
            return pageNumber(Output.of(pageNumber));
        }

        public Builder pageSize(@Nullable Output<Integer> pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        public Builder pageSize(Integer pageSize) {
            return pageSize(Output.of(pageSize));
        }

        /**
         * @param status The status of the Global Database Network. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the Global Database Network. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public GetGlobalDatabaseNetworksArgs build() {
            return $;
        }
    }

}
