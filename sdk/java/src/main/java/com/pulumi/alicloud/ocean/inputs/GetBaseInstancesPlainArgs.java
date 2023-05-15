// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ocean.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBaseInstancesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBaseInstancesPlainArgs Empty = new GetBaseInstancesPlainArgs();

    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of Instance IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Instance IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * OceanBase cluster ID.
     * 
     */
    @Import(name="instanceId")
    private @Nullable String instanceId;

    /**
     * @return OceanBase cluster ID.
     * 
     */
    public Optional<String> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * OceanBase cluster name.
     * 
     */
    @Import(name="instanceName")
    private @Nullable String instanceName;

    /**
     * @return OceanBase cluster name.
     * 
     */
    public Optional<String> instanceName() {
        return Optional.ofNullable(this.instanceName);
    }

    /**
     * A regex string to filter results by Instance name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Instance name.
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
     * The ID of the enterprise resource group to which the instance resides.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable String resourceGroupId;

    /**
     * @return The ID of the enterprise resource group to which the instance resides.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The filter keyword for the query list.
     * 
     */
    @Import(name="searchKey")
    private @Nullable String searchKey;

    /**
     * @return The filter keyword for the query list.
     * 
     */
    public Optional<String> searchKey() {
        return Optional.ofNullable(this.searchKey);
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetBaseInstancesPlainArgs() {}

    private GetBaseInstancesPlainArgs(GetBaseInstancesPlainArgs $) {
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.instanceName = $.instanceName;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.resourceGroupId = $.resourceGroupId;
        this.searchKey = $.searchKey;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBaseInstancesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBaseInstancesPlainArgs $;

        public Builder() {
            $ = new GetBaseInstancesPlainArgs();
        }

        public Builder(GetBaseInstancesPlainArgs defaults) {
            $ = new GetBaseInstancesPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids A list of Instance IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Instance IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId OceanBase cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable String instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceName OceanBase cluster name.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(@Nullable String instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Instance name.
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

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the enterprise resource group to which the instance resides.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param searchKey The filter keyword for the query list.
         * 
         * @return builder
         * 
         */
        public Builder searchKey(@Nullable String searchKey) {
            $.searchKey = searchKey;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetBaseInstancesPlainArgs build() {
            return $;
        }
    }

}
