// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceZkConnAddrArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceZkConnAddrArgs Empty = new InstanceZkConnAddrArgs();

    /**
     * The Phoenix address.
     * 
     */
    @Import(name="connAddr")
    private @Nullable Output<String> connAddr;

    /**
     * @return The Phoenix address.
     * 
     */
    public Optional<Output<String>> connAddr() {
        return Optional.ofNullable(this.connAddr);
    }

    /**
     * The number of the port over which Phoenix connects to the instance.
     * 
     */
    @Import(name="connAddrPort")
    private @Nullable Output<String> connAddrPort;

    /**
     * @return The number of the port over which Phoenix connects to the instance.
     * 
     */
    public Optional<Output<String>> connAddrPort() {
        return Optional.ofNullable(this.connAddrPort);
    }

    /**
     * The type of the network. Valid values:
     * 
     */
    @Import(name="netType")
    private @Nullable Output<String> netType;

    /**
     * @return The type of the network. Valid values:
     * 
     */
    public Optional<Output<String>> netType() {
        return Optional.ofNullable(this.netType);
    }

    private InstanceZkConnAddrArgs() {}

    private InstanceZkConnAddrArgs(InstanceZkConnAddrArgs $) {
        this.connAddr = $.connAddr;
        this.connAddrPort = $.connAddrPort;
        this.netType = $.netType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceZkConnAddrArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceZkConnAddrArgs $;

        public Builder() {
            $ = new InstanceZkConnAddrArgs();
        }

        public Builder(InstanceZkConnAddrArgs defaults) {
            $ = new InstanceZkConnAddrArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connAddr The Phoenix address.
         * 
         * @return builder
         * 
         */
        public Builder connAddr(@Nullable Output<String> connAddr) {
            $.connAddr = connAddr;
            return this;
        }

        /**
         * @param connAddr The Phoenix address.
         * 
         * @return builder
         * 
         */
        public Builder connAddr(String connAddr) {
            return connAddr(Output.of(connAddr));
        }

        /**
         * @param connAddrPort The number of the port over which Phoenix connects to the instance.
         * 
         * @return builder
         * 
         */
        public Builder connAddrPort(@Nullable Output<String> connAddrPort) {
            $.connAddrPort = connAddrPort;
            return this;
        }

        /**
         * @param connAddrPort The number of the port over which Phoenix connects to the instance.
         * 
         * @return builder
         * 
         */
        public Builder connAddrPort(String connAddrPort) {
            return connAddrPort(Output.of(connAddrPort));
        }

        /**
         * @param netType The type of the network. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder netType(@Nullable Output<String> netType) {
            $.netType = netType;
            return this;
        }

        /**
         * @param netType The type of the network. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder netType(String netType) {
            return netType(Output.of(netType));
        }

        public InstanceZkConnAddrArgs build() {
            return $;
        }
    }

}
