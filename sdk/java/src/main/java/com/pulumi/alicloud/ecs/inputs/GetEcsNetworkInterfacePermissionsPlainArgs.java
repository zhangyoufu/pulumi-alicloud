// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEcsNetworkInterfacePermissionsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEcsNetworkInterfacePermissionsPlainArgs Empty = new GetEcsNetworkInterfacePermissionsPlainArgs();

    /**
     * A list of Network Interface Permission IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Network Interface Permission IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The ID of the network interface.
     * 
     */
    @Import(name="networkInterfaceId", required=true)
    private String networkInterfaceId;

    /**
     * @return The ID of the network interface.
     * 
     */
    public String networkInterfaceId() {
        return this.networkInterfaceId;
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
     * The Status of the Network Interface Permissions.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The Status of the Network Interface Permissions.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetEcsNetworkInterfacePermissionsPlainArgs() {}

    private GetEcsNetworkInterfacePermissionsPlainArgs(GetEcsNetworkInterfacePermissionsPlainArgs $) {
        this.ids = $.ids;
        this.networkInterfaceId = $.networkInterfaceId;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEcsNetworkInterfacePermissionsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEcsNetworkInterfacePermissionsPlainArgs $;

        public Builder() {
            $ = new GetEcsNetworkInterfacePermissionsPlainArgs();
        }

        public Builder(GetEcsNetworkInterfacePermissionsPlainArgs defaults) {
            $ = new GetEcsNetworkInterfacePermissionsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Network Interface Permission IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Network Interface Permission IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param networkInterfaceId The ID of the network interface.
         * 
         * @return builder
         * 
         */
        public Builder networkInterfaceId(String networkInterfaceId) {
            $.networkInterfaceId = networkInterfaceId;
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
         * @param status The Status of the Network Interface Permissions.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetEcsNetworkInterfacePermissionsPlainArgs build() {
            $.networkInterfaceId = Objects.requireNonNull($.networkInterfaceId, "expected parameter 'networkInterfaceId' to be non-null");
            return $;
        }
    }

}
