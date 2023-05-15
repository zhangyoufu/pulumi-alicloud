// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudstoragegateway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGatewaySmbUsersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGatewaySmbUsersArgs Empty = new GetGatewaySmbUsersArgs();

    /**
     * The Gateway ID.
     * 
     */
    @Import(name="gatewayId", required=true)
    private Output<String> gatewayId;

    /**
     * @return The Gateway ID.
     * 
     */
    public Output<String> gatewayId() {
        return this.gatewayId;
    }

    /**
     * A list of Gateway SMB User IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Gateway SMB User IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Gateway SMB username.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Gateway SMB username.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
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

    private GetGatewaySmbUsersArgs() {}

    private GetGatewaySmbUsersArgs(GetGatewaySmbUsersArgs $) {
        this.gatewayId = $.gatewayId;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGatewaySmbUsersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGatewaySmbUsersArgs $;

        public Builder() {
            $ = new GetGatewaySmbUsersArgs();
        }

        public Builder(GetGatewaySmbUsersArgs defaults) {
            $ = new GetGatewaySmbUsersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param gatewayId The Gateway ID.
         * 
         * @return builder
         * 
         */
        public Builder gatewayId(Output<String> gatewayId) {
            $.gatewayId = gatewayId;
            return this;
        }

        /**
         * @param gatewayId The Gateway ID.
         * 
         * @return builder
         * 
         */
        public Builder gatewayId(String gatewayId) {
            return gatewayId(Output.of(gatewayId));
        }

        /**
         * @param ids A list of Gateway SMB User IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Gateway SMB User IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Gateway SMB User IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Gateway SMB username.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Gateway SMB username.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
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

        public GetGatewaySmbUsersArgs build() {
            $.gatewayId = Objects.requireNonNull($.gatewayId, "expected parameter 'gatewayId' to be non-null");
            return $;
        }
    }

}
