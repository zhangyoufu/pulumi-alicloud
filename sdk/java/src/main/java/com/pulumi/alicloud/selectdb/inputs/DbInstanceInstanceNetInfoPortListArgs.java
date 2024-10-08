// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.selectdb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DbInstanceInstanceNetInfoPortListArgs extends com.pulumi.resources.ResourceArgs {

    public static final DbInstanceInstanceNetInfoPortListArgs Empty = new DbInstanceInstanceNetInfoPortListArgs();

    /**
     * The port that is used to connect.
     * 
     */
    @Import(name="port")
    private @Nullable Output<String> port;

    /**
     * @return The port that is used to connect.
     * 
     */
    public Optional<Output<String>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The protocol of the port.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return The protocol of the port.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    private DbInstanceInstanceNetInfoPortListArgs() {}

    private DbInstanceInstanceNetInfoPortListArgs(DbInstanceInstanceNetInfoPortListArgs $) {
        this.port = $.port;
        this.protocol = $.protocol;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DbInstanceInstanceNetInfoPortListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DbInstanceInstanceNetInfoPortListArgs $;

        public Builder() {
            $ = new DbInstanceInstanceNetInfoPortListArgs();
        }

        public Builder(DbInstanceInstanceNetInfoPortListArgs defaults) {
            $ = new DbInstanceInstanceNetInfoPortListArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param port The port that is used to connect.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<String> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port that is used to connect.
         * 
         * @return builder
         * 
         */
        public Builder port(String port) {
            return port(Output.of(port));
        }

        /**
         * @param protocol The protocol of the port.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The protocol of the port.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        public DbInstanceInstanceNetInfoPortListArgs build() {
            return $;
        }
    }

}
