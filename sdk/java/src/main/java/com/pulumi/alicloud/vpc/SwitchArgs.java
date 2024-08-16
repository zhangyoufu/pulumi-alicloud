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


public final class SwitchArgs extends com.pulumi.resources.ResourceArgs {

    public static final SwitchArgs Empty = new SwitchArgs();

    /**
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.119.0. New field &#39;zone_id&#39; instead.
     * 
     * @deprecated
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.119.0. New field &#39;zone_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead. */
    @Import(name="availabilityZone")
    private @Nullable Output<String> availabilityZone;

    /**
     * @return Field &#39;availability_zone&#39; has been deprecated from provider version 1.119.0. New field &#39;zone_id&#39; instead.
     * 
     * @deprecated
     * Field &#39;availability_zone&#39; has been deprecated from provider version 1.119.0. New field &#39;zone_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead. */
    public Optional<Output<String>> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * The IPv4 CIDR block of the VSwitch.
     * 
     */
    @Import(name="cidrBlock", required=true)
    private Output<String> cidrBlock;

    /**
     * @return The IPv4 CIDR block of the VSwitch.
     * 
     */
    public Output<String> cidrBlock() {
        return this.cidrBlock;
    }

    /**
     * The description of VSwitch.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of VSwitch.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Whether the IPv6 function is enabled in the switch. Value:
     * - **true**: enables IPv6.
     * - **false** (default): IPv6 is not enabled.
     * 
     */
    @Import(name="enableIpv6")
    private @Nullable Output<Boolean> enableIpv6;

    /**
     * @return Whether the IPv6 function is enabled in the switch. Value:
     * - **true**: enables IPv6.
     * - **false** (default): IPv6 is not enabled.
     * 
     */
    public Optional<Output<Boolean>> enableIpv6() {
        return Optional.ofNullable(this.enableIpv6);
    }

    /**
     * The IPv6 CIDR block of the VSwitch.
     * 
     */
    @Import(name="ipv6CidrBlockMask")
    private @Nullable Output<Integer> ipv6CidrBlockMask;

    /**
     * @return The IPv6 CIDR block of the VSwitch.
     * 
     */
    public Optional<Output<Integer>> ipv6CidrBlockMask() {
        return Optional.ofNullable(this.ipv6CidrBlockMask);
    }

    /**
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vswitch_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vswitch_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vswitch_name&#39; instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vswitch_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The tags of VSwitch.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return The tags of VSwitch.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The VPC ID.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    @Import(name="vpcId", required=true)
    private Output<String> vpcId;

    /**
     * @return The VPC ID.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     * The name of the VSwitch.
     * 
     */
    @Import(name="vswitchName")
    private @Nullable Output<String> vswitchName;

    /**
     * @return The name of the VSwitch.
     * 
     */
    public Optional<Output<String>> vswitchName() {
        return Optional.ofNullable(this.vswitchName);
    }

    /**
     * The AZ for the VSwitch. **Note:** Required for a VPC VSwitch.
     * 
     */
    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    /**
     * @return The AZ for the VSwitch. **Note:** Required for a VPC VSwitch.
     * 
     */
    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private SwitchArgs() {}

    private SwitchArgs(SwitchArgs $) {
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
    public static Builder builder(SwitchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SwitchArgs $;

        public Builder() {
            $ = new SwitchArgs();
        }

        public Builder(SwitchArgs defaults) {
            $ = new SwitchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param availabilityZone Field &#39;availability_zone&#39; has been deprecated from provider version 1.119.0. New field &#39;zone_id&#39; instead.
         * 
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
         * @param availabilityZone Field &#39;availability_zone&#39; has been deprecated from provider version 1.119.0. New field &#39;zone_id&#39; instead.
         * 
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

        /**
         * @param cidrBlock The IPv4 CIDR block of the VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(Output<String> cidrBlock) {
            $.cidrBlock = cidrBlock;
            return this;
        }

        /**
         * @param cidrBlock The IPv4 CIDR block of the VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(String cidrBlock) {
            return cidrBlock(Output.of(cidrBlock));
        }

        /**
         * @param description The description of VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param enableIpv6 Whether the IPv6 function is enabled in the switch. Value:
         * - **true**: enables IPv6.
         * - **false** (default): IPv6 is not enabled.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(@Nullable Output<Boolean> enableIpv6) {
            $.enableIpv6 = enableIpv6;
            return this;
        }

        /**
         * @param enableIpv6 Whether the IPv6 function is enabled in the switch. Value:
         * - **true**: enables IPv6.
         * - **false** (default): IPv6 is not enabled.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(Boolean enableIpv6) {
            return enableIpv6(Output.of(enableIpv6));
        }

        /**
         * @param ipv6CidrBlockMask The IPv6 CIDR block of the VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder ipv6CidrBlockMask(@Nullable Output<Integer> ipv6CidrBlockMask) {
            $.ipv6CidrBlockMask = ipv6CidrBlockMask;
            return this;
        }

        /**
         * @param ipv6CidrBlockMask The IPv6 CIDR block of the VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder ipv6CidrBlockMask(Integer ipv6CidrBlockMask) {
            return ipv6CidrBlockMask(Output.of(ipv6CidrBlockMask));
        }

        /**
         * @param name Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vswitch_name&#39; instead.
         * 
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
         * @param name Field &#39;name&#39; has been deprecated from provider version 1.119.0. New field &#39;vswitch_name&#39; instead.
         * 
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

        /**
         * @param tags The tags of VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags of VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcId The VPC ID.
         * 
         * The following arguments will be discarded. Please use new fields as soon as possible:
         * 
         * @return builder
         * 
         */
        public Builder vpcId(Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The VPC ID.
         * 
         * The following arguments will be discarded. Please use new fields as soon as possible:
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vswitchName The name of the VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder vswitchName(@Nullable Output<String> vswitchName) {
            $.vswitchName = vswitchName;
            return this;
        }

        /**
         * @param vswitchName The name of the VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder vswitchName(String vswitchName) {
            return vswitchName(Output.of(vswitchName));
        }

        /**
         * @param zoneId The AZ for the VSwitch. **Note:** Required for a VPC VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The AZ for the VSwitch. **Note:** Required for a VPC VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public SwitchArgs build() {
            if ($.cidrBlock == null) {
                throw new MissingRequiredPropertyException("SwitchArgs", "cidrBlock");
            }
            if ($.vpcId == null) {
                throw new MissingRequiredPropertyException("SwitchArgs", "vpcId");
            }
            return $;
        }
    }

}
