// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NodePoolManagementAutoUpgradePolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final NodePoolManagementAutoUpgradePolicyArgs Empty = new NodePoolManagementAutoUpgradePolicyArgs();

    /**
     * Specifies whether  to automatically update the kubelet. Valid values: `true`: yes; `false`: no.
     * 
     */
    @Import(name="autoUpgradeKubelet")
    private @Nullable Output<Boolean> autoUpgradeKubelet;

    /**
     * @return Specifies whether  to automatically update the kubelet. Valid values: `true`: yes; `false`: no.
     * 
     */
    public Optional<Output<Boolean>> autoUpgradeKubelet() {
        return Optional.ofNullable(this.autoUpgradeKubelet);
    }

    private NodePoolManagementAutoUpgradePolicyArgs() {}

    private NodePoolManagementAutoUpgradePolicyArgs(NodePoolManagementAutoUpgradePolicyArgs $) {
        this.autoUpgradeKubelet = $.autoUpgradeKubelet;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NodePoolManagementAutoUpgradePolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NodePoolManagementAutoUpgradePolicyArgs $;

        public Builder() {
            $ = new NodePoolManagementAutoUpgradePolicyArgs();
        }

        public Builder(NodePoolManagementAutoUpgradePolicyArgs defaults) {
            $ = new NodePoolManagementAutoUpgradePolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoUpgradeKubelet Specifies whether  to automatically update the kubelet. Valid values: `true`: yes; `false`: no.
         * 
         * @return builder
         * 
         */
        public Builder autoUpgradeKubelet(@Nullable Output<Boolean> autoUpgradeKubelet) {
            $.autoUpgradeKubelet = autoUpgradeKubelet;
            return this;
        }

        /**
         * @param autoUpgradeKubelet Specifies whether  to automatically update the kubelet. Valid values: `true`: yes; `false`: no.
         * 
         * @return builder
         * 
         */
        public Builder autoUpgradeKubelet(Boolean autoUpgradeKubelet) {
            return autoUpgradeKubelet(Output.of(autoUpgradeKubelet));
        }

        public NodePoolManagementAutoUpgradePolicyArgs build() {
            return $;
        }
    }

}
