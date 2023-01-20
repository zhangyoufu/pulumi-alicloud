// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NatGatewayState extends com.pulumi.resources.ResourceArgs {

    public static final NatGatewayState Empty = new NatGatewayState();

    /**
     * Whether enable the deletion protection or not. Default value: `false`.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     * 
     */
    @Import(name="deletionProtection")
    private @Nullable Output<Boolean> deletionProtection;

    /**
     * @return Whether enable the deletion protection or not. Default value: `false`.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     * 
     */
    public Optional<Output<Boolean>> deletionProtection() {
        return Optional.ofNullable(this.deletionProtection);
    }

    /**
     * Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Specifies whether to only precheck this request. Default value: `false`.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return Specifies whether to only precheck this request. Default value: `false`.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The EIP binding mode of the NAT gateway. Default value: `MULTI_BINDED`. Valid values:
     * 
     */
    @Import(name="eipBindMode")
    private @Nullable Output<String> eipBindMode;

    /**
     * @return The EIP binding mode of the NAT gateway. Default value: `MULTI_BINDED`. Valid values:
     * 
     */
    public Optional<Output<String>> eipBindMode() {
        return Optional.ofNullable(this.eipBindMode);
    }

    /**
     * Specifies whether to forcefully delete the NAT gateway.
     * 
     */
    @Import(name="force")
    private @Nullable Output<Boolean> force;

    /**
     * @return Specifies whether to forcefully delete the NAT gateway.
     * 
     */
    public Optional<Output<Boolean>> force() {
        return Optional.ofNullable(this.force);
    }

    /**
     * The nat gateway will auto create a forward item.
     * 
     */
    @Import(name="forwardTableIds")
    private @Nullable Output<String> forwardTableIds;

    /**
     * @return The nat gateway will auto create a forward item.
     * 
     */
    public Optional<Output<String>> forwardTableIds() {
        return Optional.ofNullable(this.forwardTableIds);
    }

    /**
     * Field `instance_charge_type` has been deprecated from provider version 1.121.0. New field `payment_type` instead.
     * 
     */
    @Import(name="instanceChargeType")
    private @Nullable Output<String> instanceChargeType;

    /**
     * @return Field `instance_charge_type` has been deprecated from provider version 1.121.0. New field `payment_type` instead.
     * 
     */
    public Optional<Output<String>> instanceChargeType() {
        return Optional.ofNullable(this.instanceChargeType);
    }

    /**
     * The internet charge type. Valid values `PayByLcu` and `PayBySpec`. The `PayByLcu` is only support enhanced NAT. **NOTE:** From 1.137.0+, The `PayBySpec` has been deprecated.
     * 
     */
    @Import(name="internetChargeType")
    private @Nullable Output<String> internetChargeType;

    /**
     * @return The internet charge type. Valid values `PayByLcu` and `PayBySpec`. The `PayByLcu` is only support enhanced NAT. **NOTE:** From 1.137.0+, The `PayBySpec` has been deprecated.
     * 
     */
    public Optional<Output<String>> internetChargeType() {
        return Optional.ofNullable(this.internetChargeType);
    }

    /**
     * Field `name` has been deprecated from provider version 1.121.0. New field `nat_gateway_name` instead.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Field `name` has been deprecated from provider version 1.121.0. New field `nat_gateway_name` instead.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
     * 
     */
    @Import(name="natGatewayName")
    private @Nullable Output<String> natGatewayName;

    /**
     * @return Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
     * 
     */
    public Optional<Output<String>> natGatewayName() {
        return Optional.ofNullable(this.natGatewayName);
    }

    /**
     * The type of NAT gateway. Valid values: `Normal` and `Enhanced`. **NOTE:** From 1.137.0+,  The `Normal` has been deprecated.
     * 
     */
    @Import(name="natType")
    private @Nullable Output<String> natType;

    /**
     * @return The type of NAT gateway. Valid values: `Normal` and `Enhanced`. **NOTE:** From 1.137.0+,  The `Normal` has been deprecated.
     * 
     */
    public Optional<Output<String>> natType() {
        return Optional.ofNullable(this.natType);
    }

    /**
     * Indicates the type of the created NAT gateway. Valid values `internet` and `intranet`. `internet`: Internet NAT Gateway. `intranet`: VPC NAT Gateway.
     * 
     */
    @Import(name="networkType")
    private @Nullable Output<String> networkType;

    /**
     * @return Indicates the type of the created NAT gateway. Valid values `internet` and `intranet`. `internet`: Internet NAT Gateway. `intranet`: VPC NAT Gateway.
     * 
     */
    public Optional<Output<String>> networkType() {
        return Optional.ofNullable(this.networkType);
    }

    /**
     * The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
     * 
     */
    @Import(name="paymentType")
    private @Nullable Output<String> paymentType;

    /**
     * @return The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
     * 
     */
    public Optional<Output<String>> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }

    @Import(name="period")
    private @Nullable Output<Integer> period;

    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * The nat gateway will auto create a snat item.
     * 
     */
    @Import(name="snatTableIds")
    private @Nullable Output<String> snatTableIds;

    /**
     * @return The nat gateway will auto create a snat item.
     * 
     */
    public Optional<Output<String>> snatTableIds() {
        return Optional.ofNullable(this.snatTableIds);
    }

    /**
     * The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Effective when `internet_charge_type` is `PayBySpec` and `network_type` is `internet`. Details refer to [Nat Gateway Specification](https://help.aliyun.com/document_detail/203500.html).
     * 
     */
    @Import(name="specification")
    private @Nullable Output<String> specification;

    /**
     * @return The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Effective when `internet_charge_type` is `PayBySpec` and `network_type` is `internet`. Details refer to [Nat Gateway Specification](https://help.aliyun.com/document_detail/203500.html).
     * 
     */
    public Optional<Output<String>> specification() {
        return Optional.ofNullable(this.specification);
    }

    /**
     * (Available in 1.121.0+) The status of NAT gateway.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return (Available in 1.121.0+) The status of NAT gateway.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tags of NAT gateway.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return The tags of NAT gateway.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The VPC ID.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The VPC ID.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * The id of VSwitch.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return The id of VSwitch.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private NatGatewayState() {}

    private NatGatewayState(NatGatewayState $) {
        this.deletionProtection = $.deletionProtection;
        this.description = $.description;
        this.dryRun = $.dryRun;
        this.eipBindMode = $.eipBindMode;
        this.force = $.force;
        this.forwardTableIds = $.forwardTableIds;
        this.instanceChargeType = $.instanceChargeType;
        this.internetChargeType = $.internetChargeType;
        this.name = $.name;
        this.natGatewayName = $.natGatewayName;
        this.natType = $.natType;
        this.networkType = $.networkType;
        this.paymentType = $.paymentType;
        this.period = $.period;
        this.snatTableIds = $.snatTableIds;
        this.specification = $.specification;
        this.status = $.status;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NatGatewayState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NatGatewayState $;

        public Builder() {
            $ = new NatGatewayState();
        }

        public Builder(NatGatewayState defaults) {
            $ = new NatGatewayState(Objects.requireNonNull(defaults));
        }

        /**
         * @param deletionProtection Whether enable the deletion protection or not. Default value: `false`.
         * - true: Enable deletion protection.
         * - false: Disable deletion protection.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(@Nullable Output<Boolean> deletionProtection) {
            $.deletionProtection = deletionProtection;
            return this;
        }

        /**
         * @param deletionProtection Whether enable the deletion protection or not. Default value: `false`.
         * - true: Enable deletion protection.
         * - false: Disable deletion protection.
         * 
         * @return builder
         * 
         */
        public Builder deletionProtection(Boolean deletionProtection) {
            return deletionProtection(Output.of(deletionProtection));
        }

        /**
         * @param description Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dryRun Specifies whether to only precheck this request. Default value: `false`.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun Specifies whether to only precheck this request. Default value: `false`.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param eipBindMode The EIP binding mode of the NAT gateway. Default value: `MULTI_BINDED`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder eipBindMode(@Nullable Output<String> eipBindMode) {
            $.eipBindMode = eipBindMode;
            return this;
        }

        /**
         * @param eipBindMode The EIP binding mode of the NAT gateway. Default value: `MULTI_BINDED`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder eipBindMode(String eipBindMode) {
            return eipBindMode(Output.of(eipBindMode));
        }

        /**
         * @param force Specifies whether to forcefully delete the NAT gateway.
         * 
         * @return builder
         * 
         */
        public Builder force(@Nullable Output<Boolean> force) {
            $.force = force;
            return this;
        }

        /**
         * @param force Specifies whether to forcefully delete the NAT gateway.
         * 
         * @return builder
         * 
         */
        public Builder force(Boolean force) {
            return force(Output.of(force));
        }

        /**
         * @param forwardTableIds The nat gateway will auto create a forward item.
         * 
         * @return builder
         * 
         */
        public Builder forwardTableIds(@Nullable Output<String> forwardTableIds) {
            $.forwardTableIds = forwardTableIds;
            return this;
        }

        /**
         * @param forwardTableIds The nat gateway will auto create a forward item.
         * 
         * @return builder
         * 
         */
        public Builder forwardTableIds(String forwardTableIds) {
            return forwardTableIds(Output.of(forwardTableIds));
        }

        /**
         * @param instanceChargeType Field `instance_charge_type` has been deprecated from provider version 1.121.0. New field `payment_type` instead.
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(@Nullable Output<String> instanceChargeType) {
            $.instanceChargeType = instanceChargeType;
            return this;
        }

        /**
         * @param instanceChargeType Field `instance_charge_type` has been deprecated from provider version 1.121.0. New field `payment_type` instead.
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(String instanceChargeType) {
            return instanceChargeType(Output.of(instanceChargeType));
        }

        /**
         * @param internetChargeType The internet charge type. Valid values `PayByLcu` and `PayBySpec`. The `PayByLcu` is only support enhanced NAT. **NOTE:** From 1.137.0+, The `PayBySpec` has been deprecated.
         * 
         * @return builder
         * 
         */
        public Builder internetChargeType(@Nullable Output<String> internetChargeType) {
            $.internetChargeType = internetChargeType;
            return this;
        }

        /**
         * @param internetChargeType The internet charge type. Valid values `PayByLcu` and `PayBySpec`. The `PayByLcu` is only support enhanced NAT. **NOTE:** From 1.137.0+, The `PayBySpec` has been deprecated.
         * 
         * @return builder
         * 
         */
        public Builder internetChargeType(String internetChargeType) {
            return internetChargeType(Output.of(internetChargeType));
        }

        /**
         * @param name Field `name` has been deprecated from provider version 1.121.0. New field `nat_gateway_name` instead.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Field `name` has been deprecated from provider version 1.121.0. New field `nat_gateway_name` instead.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param natGatewayName Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder natGatewayName(@Nullable Output<String> natGatewayName) {
            $.natGatewayName = natGatewayName;
            return this;
        }

        /**
         * @param natGatewayName Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder natGatewayName(String natGatewayName) {
            return natGatewayName(Output.of(natGatewayName));
        }

        /**
         * @param natType The type of NAT gateway. Valid values: `Normal` and `Enhanced`. **NOTE:** From 1.137.0+,  The `Normal` has been deprecated.
         * 
         * @return builder
         * 
         */
        public Builder natType(@Nullable Output<String> natType) {
            $.natType = natType;
            return this;
        }

        /**
         * @param natType The type of NAT gateway. Valid values: `Normal` and `Enhanced`. **NOTE:** From 1.137.0+,  The `Normal` has been deprecated.
         * 
         * @return builder
         * 
         */
        public Builder natType(String natType) {
            return natType(Output.of(natType));
        }

        /**
         * @param networkType Indicates the type of the created NAT gateway. Valid values `internet` and `intranet`. `internet`: Internet NAT Gateway. `intranet`: VPC NAT Gateway.
         * 
         * @return builder
         * 
         */
        public Builder networkType(@Nullable Output<String> networkType) {
            $.networkType = networkType;
            return this;
        }

        /**
         * @param networkType Indicates the type of the created NAT gateway. Valid values `internet` and `intranet`. `internet`: Internet NAT Gateway. `intranet`: VPC NAT Gateway.
         * 
         * @return builder
         * 
         */
        public Builder networkType(String networkType) {
            return networkType(Output.of(networkType));
        }

        /**
         * @param paymentType The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(@Nullable Output<String> paymentType) {
            $.paymentType = paymentType;
            return this;
        }

        /**
         * @param paymentType The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
         * 
         * @return builder
         * 
         */
        public Builder paymentType(String paymentType) {
            return paymentType(Output.of(paymentType));
        }

        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param snatTableIds The nat gateway will auto create a snat item.
         * 
         * @return builder
         * 
         */
        public Builder snatTableIds(@Nullable Output<String> snatTableIds) {
            $.snatTableIds = snatTableIds;
            return this;
        }

        /**
         * @param snatTableIds The nat gateway will auto create a snat item.
         * 
         * @return builder
         * 
         */
        public Builder snatTableIds(String snatTableIds) {
            return snatTableIds(Output.of(snatTableIds));
        }

        /**
         * @param specification The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Effective when `internet_charge_type` is `PayBySpec` and `network_type` is `internet`. Details refer to [Nat Gateway Specification](https://help.aliyun.com/document_detail/203500.html).
         * 
         * @return builder
         * 
         */
        public Builder specification(@Nullable Output<String> specification) {
            $.specification = specification;
            return this;
        }

        /**
         * @param specification The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Effective when `internet_charge_type` is `PayBySpec` and `network_type` is `internet`. Details refer to [Nat Gateway Specification](https://help.aliyun.com/document_detail/203500.html).
         * 
         * @return builder
         * 
         */
        public Builder specification(String specification) {
            return specification(Output.of(specification));
        }

        /**
         * @param status (Available in 1.121.0+) The status of NAT gateway.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status (Available in 1.121.0+) The status of NAT gateway.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags The tags of NAT gateway.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags of NAT gateway.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcId The VPC ID.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The VPC ID.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vswitchId The id of VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The id of VSwitch.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public NatGatewayState build() {
            return $;
        }
    }

}
