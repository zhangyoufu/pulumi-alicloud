// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class DbInstanceEndpointAddressArgs extends com.pulumi.resources.ResourceArgs {

    public static final DbInstanceEndpointAddressArgs Empty = new DbInstanceEndpointAddressArgs();

    /**
     * The prefix of the public endpoint.
     * 
     */
    @Import(name="connectionStringPrefix", required=true)
    private Output<String> connectionStringPrefix;

    /**
     * @return The prefix of the public endpoint.
     * 
     */
    public Output<String> connectionStringPrefix() {
        return this.connectionStringPrefix;
    }

    /**
     * The Endpoint ID of the instance.
     * 
     */
    @Import(name="dbInstanceEndpointId", required=true)
    private Output<String> dbInstanceEndpointId;

    /**
     * @return The Endpoint ID of the instance.
     * 
     */
    public Output<String> dbInstanceEndpointId() {
        return this.dbInstanceEndpointId;
    }

    /**
     * The ID of the instance.
     * 
     */
    @Import(name="dbInstanceId", required=true)
    private Output<String> dbInstanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> dbInstanceId() {
        return this.dbInstanceId;
    }

    /**
     * The port number of the public endpoint.
     * 
     */
    @Import(name="port", required=true)
    private Output<String> port;

    /**
     * @return The port number of the public endpoint.
     * 
     */
    public Output<String> port() {
        return this.port;
    }

    private DbInstanceEndpointAddressArgs() {}

    private DbInstanceEndpointAddressArgs(DbInstanceEndpointAddressArgs $) {
        this.connectionStringPrefix = $.connectionStringPrefix;
        this.dbInstanceEndpointId = $.dbInstanceEndpointId;
        this.dbInstanceId = $.dbInstanceId;
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DbInstanceEndpointAddressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DbInstanceEndpointAddressArgs $;

        public Builder() {
            $ = new DbInstanceEndpointAddressArgs();
        }

        public Builder(DbInstanceEndpointAddressArgs defaults) {
            $ = new DbInstanceEndpointAddressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectionStringPrefix The prefix of the public endpoint.
         * 
         * @return builder
         * 
         */
        public Builder connectionStringPrefix(Output<String> connectionStringPrefix) {
            $.connectionStringPrefix = connectionStringPrefix;
            return this;
        }

        /**
         * @param connectionStringPrefix The prefix of the public endpoint.
         * 
         * @return builder
         * 
         */
        public Builder connectionStringPrefix(String connectionStringPrefix) {
            return connectionStringPrefix(Output.of(connectionStringPrefix));
        }

        /**
         * @param dbInstanceEndpointId The Endpoint ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceEndpointId(Output<String> dbInstanceEndpointId) {
            $.dbInstanceEndpointId = dbInstanceEndpointId;
            return this;
        }

        /**
         * @param dbInstanceEndpointId The Endpoint ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceEndpointId(String dbInstanceEndpointId) {
            return dbInstanceEndpointId(Output.of(dbInstanceEndpointId));
        }

        /**
         * @param dbInstanceId The ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(Output<String> dbInstanceId) {
            $.dbInstanceId = dbInstanceId;
            return this;
        }

        /**
         * @param dbInstanceId The ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(String dbInstanceId) {
            return dbInstanceId(Output.of(dbInstanceId));
        }

        /**
         * @param port The port number of the public endpoint.
         * 
         * @return builder
         * 
         */
        public Builder port(Output<String> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port number of the public endpoint.
         * 
         * @return builder
         * 
         */
        public Builder port(String port) {
            return port(Output.of(port));
        }

        public DbInstanceEndpointAddressArgs build() {
            $.connectionStringPrefix = Objects.requireNonNull($.connectionStringPrefix, "expected parameter 'connectionStringPrefix' to be non-null");
            $.dbInstanceEndpointId = Objects.requireNonNull($.dbInstanceEndpointId, "expected parameter 'dbInstanceEndpointId' to be non-null");
            $.dbInstanceId = Objects.requireNonNull($.dbInstanceId, "expected parameter 'dbInstanceId' to be non-null");
            $.port = Objects.requireNonNull($.port, "expected parameter 'port' to be non-null");
            return $;
        }
    }

}
