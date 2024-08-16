// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SubnetArgs extends com.pulumi.resources.ResourceArgs {

    public static final SubnetArgs Empty = new SubnetArgs();

    /**
     * @deprecated
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.119.0. New field &#39;zone_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead. */
    @Import(name="availabilityZone")
    private @Nullable Output<String> availabilityZone;

    /**
     * @deprecated
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.119.0. New field &#39;zone_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead. */
    public Optional<Output<String>> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    @Import(name="cidrBlock", required=true)
    private Output<String> cidrBlock;

    public Output<String> cidrBlock() {
        return this.cidrBlock;
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="enableIpv6")
    private @Nullable Output<Boolean> enableIpv6;

    public Optional<Output<Boolean>> enableIpv6() {
        return Optional.ofNullable(this.enableIpv6);
    }

    @Import(name="ipv6CidrBlockMask")
    private @Nullable Output<Integer> ipv6CidrBlockMask;

    public Optional<Output<Integer>> ipv6CidrBlockMask() {
        return Optional.ofNullable(this.ipv6CidrBlockMask);
    }

    /**
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vswitch_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vswitch_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    @Import(name="vpcId", required=true)
    private Output<String> vpcId;

    public Output<String> vpcId() {
        return this.vpcId;
    }

    @Import(name="vswitchName")
    private @Nullable Output<String> vswitchName;

    public Optional<Output<String>> vswitchName() {
        return Optional.ofNullable(this.vswitchName);
    }

    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private SubnetArgs() {}

    private SubnetArgs(SubnetArgs $) {
        this.availabilityZone = $.availabilityZone;
        this.cidrBlock = $.cidrBlock;
        this.description = $.description;
        this.enableIpv6 = $.enableIpv6;
        this.ipv6CidrBlockMask = $.ipv6CidrBlockMask;
        this.name = $.name;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.vswitchName = $.vswitchName;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubnetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubnetArgs $;

        public Builder() {
            $ = new SubnetArgs();
        }

        public Builder(SubnetArgs defaults) {
            $ = new SubnetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;availability_zone&#39; has been deprecated from provider version 1.119.0. New field &#39;zone_id&#39; instead.
         * 
         */
        @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead. */
        public Builder availabilityZone(@Nullable Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;availability_zone&#39; has been deprecated from provider version 1.119.0. New field &#39;zone_id&#39; instead.
         * 
         */
        @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead. */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        public Builder cidrBlock(Output<String> cidrBlock) {
            $.cidrBlock = cidrBlock;
            return this;
        }

        public Builder cidrBlock(String cidrBlock) {
            return cidrBlock(Output.of(cidrBlock));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder enableIpv6(@Nullable Output<Boolean> enableIpv6) {
            $.enableIpv6 = enableIpv6;
            return this;
        }

        public Builder enableIpv6(Boolean enableIpv6) {
            return enableIpv6(Output.of(enableIpv6));
        }

        public Builder ipv6CidrBlockMask(@Nullable Output<Integer> ipv6CidrBlockMask) {
            $.ipv6CidrBlockMask = ipv6CidrBlockMask;
            return this;
        }

        public Builder ipv6CidrBlockMask(Integer ipv6CidrBlockMask) {
            return ipv6CidrBlockMask(Output.of(ipv6CidrBlockMask));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vswitch_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vswitch_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public Builder vpcId(Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public Builder vswitchName(@Nullable Output<String> vswitchName) {
            $.vswitchName = vswitchName;
            return this;
        }

        public Builder vswitchName(String vswitchName) {
            return vswitchName(Output.of(vswitchName));
        }

        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public SubnetArgs build() {
            if ($.cidrBlock == null) {
                throw new MissingRequiredPropertyException("SubnetArgs", "cidrBlock");
            }
            if ($.vpcId == null) {
                throw new MissingRequiredPropertyException("SubnetArgs", "vpcId");
            }
            return $;
        }
    }

}
