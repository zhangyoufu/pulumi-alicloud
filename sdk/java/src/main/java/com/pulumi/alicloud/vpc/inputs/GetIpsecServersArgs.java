// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIpsecServersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIpsecServersArgs Empty = new GetIpsecServersArgs();

    /**
     * A list of Ipsec Server IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Ipsec Server IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The name of the IPsec server.
     * 
     */
    @Import(name="ipsecServerName")
    private @Nullable Output<String> ipsecServerName;

    /**
     * @return The name of the IPsec server.
     * 
     */
    public Optional<Output<String>> ipsecServerName() {
        return Optional.ofNullable(this.ipsecServerName);
    }

    /**
     * A regex string to filter results by Ipsec Server name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Ipsec Server name.
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

    /**
     * The ID of the VPN gateway.
     * 
     */
    @Import(name="vpnGatewayId")
    private @Nullable Output<String> vpnGatewayId;

    /**
     * @return The ID of the VPN gateway.
     * 
     */
    public Optional<Output<String>> vpnGatewayId() {
        return Optional.ofNullable(this.vpnGatewayId);
    }

    private GetIpsecServersArgs() {}

    private GetIpsecServersArgs(GetIpsecServersArgs $) {
        this.ids = $.ids;
        this.ipsecServerName = $.ipsecServerName;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.vpnGatewayId = $.vpnGatewayId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpsecServersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpsecServersArgs $;

        public Builder() {
            $ = new GetIpsecServersArgs();
        }

        public Builder(GetIpsecServersArgs defaults) {
            $ = new GetIpsecServersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Ipsec Server IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Ipsec Server IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Ipsec Server IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param ipsecServerName The name of the IPsec server.
         * 
         * @return builder
         * 
         */
        public Builder ipsecServerName(@Nullable Output<String> ipsecServerName) {
            $.ipsecServerName = ipsecServerName;
            return this;
        }

        /**
         * @param ipsecServerName The name of the IPsec server.
         * 
         * @return builder
         * 
         */
        public Builder ipsecServerName(String ipsecServerName) {
            return ipsecServerName(Output.of(ipsecServerName));
        }

        /**
         * @param nameRegex A regex string to filter results by Ipsec Server name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Ipsec Server name.
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

        /**
         * @param vpnGatewayId The ID of the VPN gateway.
         * 
         * @return builder
         * 
         */
        public Builder vpnGatewayId(@Nullable Output<String> vpnGatewayId) {
            $.vpnGatewayId = vpnGatewayId;
            return this;
        }

        /**
         * @param vpnGatewayId The ID of the VPN gateway.
         * 
         * @return builder
         * 
         */
        public Builder vpnGatewayId(String vpnGatewayId) {
            return vpnGatewayId(Output.of(vpnGatewayId));
        }

        public GetIpsecServersArgs build() {
            return $;
        }
    }

}
