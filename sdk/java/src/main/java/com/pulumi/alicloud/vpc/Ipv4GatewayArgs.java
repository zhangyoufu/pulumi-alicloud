// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class Ipv4GatewayArgs extends com.pulumi.resources.ResourceArgs {

    public static final Ipv4GatewayArgs Empty = new Ipv4GatewayArgs();

    /**
     * The dry run.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The description of the IPv4 gateway. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    @Import(name="ipv4GatewayDescription")
    private @Nullable Output<String> ipv4GatewayDescription;

    /**
     * @return The description of the IPv4 gateway. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    public Optional<Output<String>> ipv4GatewayDescription() {
        return Optional.ofNullable(this.ipv4GatewayDescription);
    }

    /**
     * The name of the IPv4 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
     * 
     */
    @Import(name="ipv4GatewayName")
    private @Nullable Output<String> ipv4GatewayName;

    /**
     * @return The name of the IPv4 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
     * 
     */
    public Optional<Output<String>> ipv4GatewayName() {
        return Optional.ofNullable(this.ipv4GatewayName);
    }

    /**
     * The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
     * 
     */
    @Import(name="vpcId", required=true)
    private Output<String> vpcId;

    /**
     * @return The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    private Ipv4GatewayArgs() {}

    private Ipv4GatewayArgs(Ipv4GatewayArgs $) {
        this.dryRun = $.dryRun;
        this.ipv4GatewayDescription = $.ipv4GatewayDescription;
        this.ipv4GatewayName = $.ipv4GatewayName;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Ipv4GatewayArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Ipv4GatewayArgs $;

        public Builder() {
            $ = new Ipv4GatewayArgs();
        }

        public Builder(Ipv4GatewayArgs defaults) {
            $ = new Ipv4GatewayArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param ipv4GatewayDescription The description of the IPv4 gateway. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder ipv4GatewayDescription(@Nullable Output<String> ipv4GatewayDescription) {
            $.ipv4GatewayDescription = ipv4GatewayDescription;
            return this;
        }

        /**
         * @param ipv4GatewayDescription The description of the IPv4 gateway. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder ipv4GatewayDescription(String ipv4GatewayDescription) {
            return ipv4GatewayDescription(Output.of(ipv4GatewayDescription));
        }

        /**
         * @param ipv4GatewayName The name of the IPv4 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder ipv4GatewayName(@Nullable Output<String> ipv4GatewayName) {
            $.ipv4GatewayName = ipv4GatewayName;
            return this;
        }

        /**
         * @param ipv4GatewayName The name of the IPv4 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder ipv4GatewayName(String ipv4GatewayName) {
            return ipv4GatewayName(Output.of(ipv4GatewayName));
        }

        /**
         * @param vpcId The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public Ipv4GatewayArgs build() {
            $.vpcId = Objects.requireNonNull($.vpcId, "expected parameter 'vpcId' to be non-null");
            return $;
        }
    }

}
