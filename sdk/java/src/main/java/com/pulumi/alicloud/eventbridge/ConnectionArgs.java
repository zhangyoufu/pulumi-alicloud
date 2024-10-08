// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge;

import com.pulumi.alicloud.eventbridge.inputs.ConnectionAuthParametersArgs;
import com.pulumi.alicloud.eventbridge.inputs.ConnectionNetworkParametersArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionArgs Empty = new ConnectionArgs();

    /**
     * The parameters that are configured for authentication. See `auth_parameters` below.
     * 
     */
    @Import(name="authParameters")
    private @Nullable Output<ConnectionAuthParametersArgs> authParameters;

    /**
     * @return The parameters that are configured for authentication. See `auth_parameters` below.
     * 
     */
    public Optional<Output<ConnectionAuthParametersArgs>> authParameters() {
        return Optional.ofNullable(this.authParameters);
    }

    /**
     * The name of the connection.
     * 
     */
    @Import(name="connectionName", required=true)
    private Output<String> connectionName;

    /**
     * @return The name of the connection.
     * 
     */
    public Output<String> connectionName() {
        return this.connectionName;
    }

    /**
     * The description of the connection.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the connection.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The parameters that are configured for the network. See `network_parameters` below.
     * 
     */
    @Import(name="networkParameters", required=true)
    private Output<ConnectionNetworkParametersArgs> networkParameters;

    /**
     * @return The parameters that are configured for the network. See `network_parameters` below.
     * 
     */
    public Output<ConnectionNetworkParametersArgs> networkParameters() {
        return this.networkParameters;
    }

    private ConnectionArgs() {}

    private ConnectionArgs(ConnectionArgs $) {
        this.authParameters = $.authParameters;
        this.connectionName = $.connectionName;
        this.description = $.description;
        this.networkParameters = $.networkParameters;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionArgs $;

        public Builder() {
            $ = new ConnectionArgs();
        }

        public Builder(ConnectionArgs defaults) {
            $ = new ConnectionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authParameters The parameters that are configured for authentication. See `auth_parameters` below.
         * 
         * @return builder
         * 
         */
        public Builder authParameters(@Nullable Output<ConnectionAuthParametersArgs> authParameters) {
            $.authParameters = authParameters;
            return this;
        }

        /**
         * @param authParameters The parameters that are configured for authentication. See `auth_parameters` below.
         * 
         * @return builder
         * 
         */
        public Builder authParameters(ConnectionAuthParametersArgs authParameters) {
            return authParameters(Output.of(authParameters));
        }

        /**
         * @param connectionName The name of the connection.
         * 
         * @return builder
         * 
         */
        public Builder connectionName(Output<String> connectionName) {
            $.connectionName = connectionName;
            return this;
        }

        /**
         * @param connectionName The name of the connection.
         * 
         * @return builder
         * 
         */
        public Builder connectionName(String connectionName) {
            return connectionName(Output.of(connectionName));
        }

        /**
         * @param description The description of the connection.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the connection.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param networkParameters The parameters that are configured for the network. See `network_parameters` below.
         * 
         * @return builder
         * 
         */
        public Builder networkParameters(Output<ConnectionNetworkParametersArgs> networkParameters) {
            $.networkParameters = networkParameters;
            return this;
        }

        /**
         * @param networkParameters The parameters that are configured for the network. See `network_parameters` below.
         * 
         * @return builder
         * 
         */
        public Builder networkParameters(ConnectionNetworkParametersArgs networkParameters) {
            return networkParameters(Output.of(networkParameters));
        }

        public ConnectionArgs build() {
            if ($.connectionName == null) {
                throw new MissingRequiredPropertyException("ConnectionArgs", "connectionName");
            }
            if ($.networkParameters == null) {
                throw new MissingRequiredPropertyException("ConnectionArgs", "networkParameters");
            }
            return $;
        }
    }

}
