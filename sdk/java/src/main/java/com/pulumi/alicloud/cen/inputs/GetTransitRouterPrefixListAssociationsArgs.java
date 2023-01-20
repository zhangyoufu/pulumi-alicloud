// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTransitRouterPrefixListAssociationsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTransitRouterPrefixListAssociationsArgs Empty = new GetTransitRouterPrefixListAssociationsArgs();

    /**
     * A list of Cen Transit Router Prefix List Association IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Cen Transit Router Prefix List Association IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The ID of the Alibaba Cloud account to which the prefix list belongs.
     * 
     */
    @Import(name="ownerUid")
    private @Nullable Output<Integer> ownerUid;

    /**
     * @return The ID of the Alibaba Cloud account to which the prefix list belongs.
     * 
     */
    public Optional<Output<Integer>> ownerUid() {
        return Optional.ofNullable(this.ownerUid);
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
     * The ID of the prefix list.
     * 
     */
    @Import(name="prefixListId")
    private @Nullable Output<String> prefixListId;

    /**
     * @return The ID of the prefix list.
     * 
     */
    public Optional<Output<String>> prefixListId() {
        return Optional.ofNullable(this.prefixListId);
    }

    /**
     * The status of the prefix list.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the prefix list.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The ID of the transit router.
     * 
     */
    @Import(name="transitRouterId", required=true)
    private Output<String> transitRouterId;

    /**
     * @return The ID of the transit router.
     * 
     */
    public Output<String> transitRouterId() {
        return this.transitRouterId;
    }

    /**
     * The ID of the route table of the transit router.
     * 
     */
    @Import(name="transitRouterTableId", required=true)
    private Output<String> transitRouterTableId;

    /**
     * @return The ID of the route table of the transit router.
     * 
     */
    public Output<String> transitRouterTableId() {
        return this.transitRouterTableId;
    }

    private GetTransitRouterPrefixListAssociationsArgs() {}

    private GetTransitRouterPrefixListAssociationsArgs(GetTransitRouterPrefixListAssociationsArgs $) {
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.ownerUid = $.ownerUid;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.prefixListId = $.prefixListId;
        this.status = $.status;
        this.transitRouterId = $.transitRouterId;
        this.transitRouterTableId = $.transitRouterTableId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTransitRouterPrefixListAssociationsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTransitRouterPrefixListAssociationsArgs $;

        public Builder() {
            $ = new GetTransitRouterPrefixListAssociationsArgs();
        }

        public Builder(GetTransitRouterPrefixListAssociationsArgs defaults) {
            $ = new GetTransitRouterPrefixListAssociationsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Cen Transit Router Prefix List Association IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Cen Transit Router Prefix List Association IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Cen Transit Router Prefix List Association IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param ownerUid The ID of the Alibaba Cloud account to which the prefix list belongs.
         * 
         * @return builder
         * 
         */
        public Builder ownerUid(@Nullable Output<Integer> ownerUid) {
            $.ownerUid = ownerUid;
            return this;
        }

        /**
         * @param ownerUid The ID of the Alibaba Cloud account to which the prefix list belongs.
         * 
         * @return builder
         * 
         */
        public Builder ownerUid(Integer ownerUid) {
            return ownerUid(Output.of(ownerUid));
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
         * @param prefixListId The ID of the prefix list.
         * 
         * @return builder
         * 
         */
        public Builder prefixListId(@Nullable Output<String> prefixListId) {
            $.prefixListId = prefixListId;
            return this;
        }

        /**
         * @param prefixListId The ID of the prefix list.
         * 
         * @return builder
         * 
         */
        public Builder prefixListId(String prefixListId) {
            return prefixListId(Output.of(prefixListId));
        }

        /**
         * @param status The status of the prefix list.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the prefix list.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param transitRouterId The ID of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterId(Output<String> transitRouterId) {
            $.transitRouterId = transitRouterId;
            return this;
        }

        /**
         * @param transitRouterId The ID of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterId(String transitRouterId) {
            return transitRouterId(Output.of(transitRouterId));
        }

        /**
         * @param transitRouterTableId The ID of the route table of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterTableId(Output<String> transitRouterTableId) {
            $.transitRouterTableId = transitRouterTableId;
            return this;
        }

        /**
         * @param transitRouterTableId The ID of the route table of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterTableId(String transitRouterTableId) {
            return transitRouterTableId(Output.of(transitRouterTableId));
        }

        public GetTransitRouterPrefixListAssociationsArgs build() {
            $.transitRouterId = Objects.requireNonNull($.transitRouterId, "expected parameter 'transitRouterId' to be non-null");
            $.transitRouterTableId = Objects.requireNonNull($.transitRouterTableId, "expected parameter 'transitRouterTableId' to be non-null");
            return $;
        }
    }

}
