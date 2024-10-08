// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIpv6GatewaysPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIpv6GatewaysPlainArgs Empty = new GetIpv6GatewaysPlainArgs();

    /**
     * A list of Ipv6 Gateway IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Ipv6 Gateway IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The name of the IPv6 gateway.
     * 
     */
    @Import(name="ipv6GatewayName")
    private @Nullable String ipv6GatewayName;

    /**
     * @return The name of the IPv6 gateway.
     * 
     */
    public Optional<String> ipv6GatewayName() {
        return Optional.ofNullable(this.ipv6GatewayName);
    }

    /**
     * A regex string to filter results by Ipv6 Gateway name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Ipv6 Gateway name.
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
     * The status of the resource. Valid values: `Available`, `Deleting`, `Pending`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the resource. Valid values: `Available`, `Deleting`, `Pending`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.
     * 
     */
    @Import(name="vpcId")
    private @Nullable String vpcId;

    /**
     * @return The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private GetIpv6GatewaysPlainArgs() {}

    private GetIpv6GatewaysPlainArgs(GetIpv6GatewaysPlainArgs $) {
        this.ids = $.ids;
        this.ipv6GatewayName = $.ipv6GatewayName;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.status = $.status;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpv6GatewaysPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpv6GatewaysPlainArgs $;

        public Builder() {
            $ = new GetIpv6GatewaysPlainArgs();
        }

        public Builder(GetIpv6GatewaysPlainArgs defaults) {
            $ = new GetIpv6GatewaysPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Ipv6 Gateway IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Ipv6 Gateway IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param ipv6GatewayName The name of the IPv6 gateway.
         * 
         * @return builder
         * 
         */
        public Builder ipv6GatewayName(@Nullable String ipv6GatewayName) {
            $.ipv6GatewayName = ipv6GatewayName;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Ipv6 Gateway name.
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
         * @param status The status of the resource. Valid values: `Available`, `Deleting`, `Pending`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param vpcId The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable String vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        public GetIpv6GatewaysPlainArgs build() {
            return $;
        }
    }

}
