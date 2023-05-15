// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBgpPeersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBgpPeersPlainArgs Empty = new GetBgpPeersPlainArgs();

    /**
     * The ID of the BGP group.
     * 
     */
    @Import(name="bgpGroupId")
    private @Nullable String bgpGroupId;

    /**
     * @return The ID of the BGP group.
     * 
     */
    public Optional<String> bgpGroupId() {
        return Optional.ofNullable(this.bgpGroupId);
    }

    /**
     * A list of Bgp Peer IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Bgp Peer IDs.
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
     * The ID of the router.
     * 
     */
    @Import(name="routerId")
    private @Nullable String routerId;

    /**
     * @return The ID of the router.
     * 
     */
    public Optional<String> routerId() {
        return Optional.ofNullable(this.routerId);
    }

    /**
     * The status of the BGP peer.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the BGP peer.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetBgpPeersPlainArgs() {}

    private GetBgpPeersPlainArgs(GetBgpPeersPlainArgs $) {
        this.bgpGroupId = $.bgpGroupId;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.routerId = $.routerId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBgpPeersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBgpPeersPlainArgs $;

        public Builder() {
            $ = new GetBgpPeersPlainArgs();
        }

        public Builder(GetBgpPeersPlainArgs defaults) {
            $ = new GetBgpPeersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bgpGroupId The ID of the BGP group.
         * 
         * @return builder
         * 
         */
        public Builder bgpGroupId(@Nullable String bgpGroupId) {
            $.bgpGroupId = bgpGroupId;
            return this;
        }

        /**
         * @param ids A list of Bgp Peer IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Bgp Peer IDs.
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
         * @param routerId The ID of the router.
         * 
         * @return builder
         * 
         */
        public Builder routerId(@Nullable String routerId) {
            $.routerId = routerId;
            return this;
        }

        /**
         * @param status The status of the BGP peer.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetBgpPeersPlainArgs build() {
            return $;
        }
    }

}
