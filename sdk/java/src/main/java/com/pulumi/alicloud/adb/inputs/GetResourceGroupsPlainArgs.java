// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.adb.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetResourceGroupsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetResourceGroupsPlainArgs Empty = new GetResourceGroupsPlainArgs();

    /**
     * DBClusterId
     * 
     */
    @Import(name="dbClusterId", required=true)
    private String dbClusterId;

    /**
     * @return DBClusterId
     * 
     */
    public String dbClusterId() {
        return this.dbClusterId;
    }

    /**
     * The name of the resource pool, which cannot exceed 64 bytes in length.
     * 
     */
    @Import(name="groupName")
    private @Nullable String groupName;

    /**
     * @return The name of the resource pool, which cannot exceed 64 bytes in length.
     * 
     */
    public Optional<String> groupName() {
        return Optional.ofNullable(this.groupName);
    }

    /**
     * A list of AnalyticDB for MySQL (ADB) Resource Group IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of AnalyticDB for MySQL (ADB) Resource Group IDs.
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

    private GetResourceGroupsPlainArgs() {}

    private GetResourceGroupsPlainArgs(GetResourceGroupsPlainArgs $) {
        this.dbClusterId = $.dbClusterId;
        this.groupName = $.groupName;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetResourceGroupsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetResourceGroupsPlainArgs $;

        public Builder() {
            $ = new GetResourceGroupsPlainArgs();
        }

        public Builder(GetResourceGroupsPlainArgs defaults) {
            $ = new GetResourceGroupsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbClusterId DBClusterId
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(String dbClusterId) {
            $.dbClusterId = dbClusterId;
            return this;
        }

        /**
         * @param groupName The name of the resource pool, which cannot exceed 64 bytes in length.
         * 
         * @return builder
         * 
         */
        public Builder groupName(@Nullable String groupName) {
            $.groupName = groupName;
            return this;
        }

        /**
         * @param ids A list of AnalyticDB for MySQL (ADB) Resource Group IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of AnalyticDB for MySQL (ADB) Resource Group IDs.
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

        public GetResourceGroupsPlainArgs build() {
            $.dbClusterId = Objects.requireNonNull($.dbClusterId, "expected parameter 'dbClusterId' to be non-null");
            return $;
        }
    }

}
