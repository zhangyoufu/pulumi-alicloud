// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class Ipv6EgressRuleState extends com.pulumi.resources.ResourceArgs {

    public static final Ipv6EgressRuleState Empty = new Ipv6EgressRuleState();

    /**
     * The description of the egress-only rule. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the egress-only rule. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the IPv6 address to which you want to apply the egress-only rule.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The ID of the IPv6 address to which you want to apply the egress-only rule.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The type of instance to which you want to apply the egress-only rule. Valid values: `Ipv6Address`. `Ipv6Address` (default): an IPv6 address.
     * 
     */
    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    /**
     * @return The type of instance to which you want to apply the egress-only rule. Valid values: `Ipv6Address`. `Ipv6Address` (default): an IPv6 address.
     * 
     */
    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    /**
     * The name of the egress-only rule. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    @Import(name="ipv6EgressRuleName")
    private @Nullable Output<String> ipv6EgressRuleName;

    /**
     * @return The name of the egress-only rule. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    public Optional<Output<String>> ipv6EgressRuleName() {
        return Optional.ofNullable(this.ipv6EgressRuleName);
    }

    /**
     * The ID of the IPv6 gateway.
     * 
     */
    @Import(name="ipv6GatewayId")
    private @Nullable Output<String> ipv6GatewayId;

    /**
     * @return The ID of the IPv6 gateway.
     * 
     */
    public Optional<Output<String>> ipv6GatewayId() {
        return Optional.ofNullable(this.ipv6GatewayId);
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private Ipv6EgressRuleState() {}

    private Ipv6EgressRuleState(Ipv6EgressRuleState $) {
        this.description = $.description;
        this.instanceId = $.instanceId;
        this.instanceType = $.instanceType;
        this.ipv6EgressRuleName = $.ipv6EgressRuleName;
        this.ipv6GatewayId = $.ipv6GatewayId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Ipv6EgressRuleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Ipv6EgressRuleState $;

        public Builder() {
            $ = new Ipv6EgressRuleState();
        }

        public Builder(Ipv6EgressRuleState defaults) {
            $ = new Ipv6EgressRuleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the egress-only rule. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the egress-only rule. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param instanceId The ID of the IPv6 address to which you want to apply the egress-only rule.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the IPv6 address to which you want to apply the egress-only rule.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param instanceType The type of instance to which you want to apply the egress-only rule. Valid values: `Ipv6Address`. `Ipv6Address` (default): an IPv6 address.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType The type of instance to which you want to apply the egress-only rule. Valid values: `Ipv6Address`. `Ipv6Address` (default): an IPv6 address.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param ipv6EgressRuleName The name of the egress-only rule. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6EgressRuleName(@Nullable Output<String> ipv6EgressRuleName) {
            $.ipv6EgressRuleName = ipv6EgressRuleName;
            return this;
        }

        /**
         * @param ipv6EgressRuleName The name of the egress-only rule. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6EgressRuleName(String ipv6EgressRuleName) {
            return ipv6EgressRuleName(Output.of(ipv6EgressRuleName));
        }

        /**
         * @param ipv6GatewayId The ID of the IPv6 gateway.
         * 
         * @return builder
         * 
         */
        public Builder ipv6GatewayId(@Nullable Output<String> ipv6GatewayId) {
            $.ipv6GatewayId = ipv6GatewayId;
            return this;
        }

        /**
         * @param ipv6GatewayId The ID of the IPv6 gateway.
         * 
         * @return builder
         * 
         */
        public Builder ipv6GatewayId(String ipv6GatewayId) {
            return ipv6GatewayId(Output.of(ipv6GatewayId));
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public Ipv6EgressRuleState build() {
            return $;
        }
    }

}
