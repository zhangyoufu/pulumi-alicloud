// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PublicIpAddressPoolCidrBlockState extends com.pulumi.resources.ResourceArgs {

    public static final PublicIpAddressPoolCidrBlockState Empty = new PublicIpAddressPoolCidrBlockState();

    /**
     * The CIDR block.
     * 
     */
    @Import(name="cidrBlock")
    private @Nullable Output<String> cidrBlock;

    /**
     * @return The CIDR block.
     * 
     */
    public Optional<Output<String>> cidrBlock() {
        return Optional.ofNullable(this.cidrBlock);
    }

    /**
     * The ID of the VPC Public IP address pool.
     * 
     */
    @Import(name="publicIpAddressPoolId")
    private @Nullable Output<String> publicIpAddressPoolId;

    /**
     * @return The ID of the VPC Public IP address pool.
     * 
     */
    public Optional<Output<String>> publicIpAddressPoolId() {
        return Optional.ofNullable(this.publicIpAddressPoolId);
    }

    /**
     * The status of the VPC Public Ip Address Pool Cidr Block.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the VPC Public Ip Address Pool Cidr Block.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private PublicIpAddressPoolCidrBlockState() {}

    private PublicIpAddressPoolCidrBlockState(PublicIpAddressPoolCidrBlockState $) {
        this.cidrBlock = $.cidrBlock;
        this.publicIpAddressPoolId = $.publicIpAddressPoolId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PublicIpAddressPoolCidrBlockState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PublicIpAddressPoolCidrBlockState $;

        public Builder() {
            $ = new PublicIpAddressPoolCidrBlockState();
        }

        public Builder(PublicIpAddressPoolCidrBlockState defaults) {
            $ = new PublicIpAddressPoolCidrBlockState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidrBlock The CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(@Nullable Output<String> cidrBlock) {
            $.cidrBlock = cidrBlock;
            return this;
        }

        /**
         * @param cidrBlock The CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(String cidrBlock) {
            return cidrBlock(Output.of(cidrBlock));
        }

        /**
         * @param publicIpAddressPoolId The ID of the VPC Public IP address pool.
         * 
         * @return builder
         * 
         */
        public Builder publicIpAddressPoolId(@Nullable Output<String> publicIpAddressPoolId) {
            $.publicIpAddressPoolId = publicIpAddressPoolId;
            return this;
        }

        /**
         * @param publicIpAddressPoolId The ID of the VPC Public IP address pool.
         * 
         * @return builder
         * 
         */
        public Builder publicIpAddressPoolId(String publicIpAddressPoolId) {
            return publicIpAddressPoolId(Output.of(publicIpAddressPoolId));
        }

        /**
         * @param status The status of the VPC Public Ip Address Pool Cidr Block.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the VPC Public Ip Address Pool Cidr Block.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public PublicIpAddressPoolCidrBlockState build() {
            return $;
        }
    }

}
