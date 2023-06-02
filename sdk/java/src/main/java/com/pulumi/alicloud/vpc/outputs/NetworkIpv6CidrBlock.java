// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class NetworkIpv6CidrBlock {
    /**
     * @return The IPv6 CIDR block of the VPC.
     * 
     */
    private @Nullable String ipv6CidrBlock;
    /**
     * @return The IPv6 address segment type of the VPC. Value:
     * - **BGP** (default): Alibaba Cloud BGP IPv6.
     * - **ChinaMobile**: China Mobile (single line).
     * - **ChinaUnicom**: China Unicom (single line).
     * - **ChinaTelecom**: China Telecom (single line).
     * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
     * 
     */
    private @Nullable String ipv6Isp;

    private NetworkIpv6CidrBlock() {}
    /**
     * @return The IPv6 CIDR block of the VPC.
     * 
     */
    public Optional<String> ipv6CidrBlock() {
        return Optional.ofNullable(this.ipv6CidrBlock);
    }
    /**
     * @return The IPv6 address segment type of the VPC. Value:
     * - **BGP** (default): Alibaba Cloud BGP IPv6.
     * - **ChinaMobile**: China Mobile (single line).
     * - **ChinaUnicom**: China Unicom (single line).
     * - **ChinaTelecom**: China Telecom (single line).
     * &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
     * 
     */
    public Optional<String> ipv6Isp() {
        return Optional.ofNullable(this.ipv6Isp);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(NetworkIpv6CidrBlock defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String ipv6CidrBlock;
        private @Nullable String ipv6Isp;
        public Builder() {}
        public Builder(NetworkIpv6CidrBlock defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipv6CidrBlock = defaults.ipv6CidrBlock;
    	      this.ipv6Isp = defaults.ipv6Isp;
        }

        @CustomType.Setter
        public Builder ipv6CidrBlock(@Nullable String ipv6CidrBlock) {
            this.ipv6CidrBlock = ipv6CidrBlock;
            return this;
        }
        @CustomType.Setter
        public Builder ipv6Isp(@Nullable String ipv6Isp) {
            this.ipv6Isp = ipv6Isp;
            return this;
        }
        public NetworkIpv6CidrBlock build() {
            final var o = new NetworkIpv6CidrBlock();
            o.ipv6CidrBlock = ipv6CidrBlock;
            o.ipv6Isp = ipv6Isp;
            return o;
        }
    }
}
