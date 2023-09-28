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


public final class HAVipAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final HAVipAttachmentArgs Empty = new HAVipAttachmentArgs();

    /**
     * Whether to force the ECS instance or Eni instance bound to AVIP to be unbound. The value is:
     * - **True**: Force unbinding.
     * - **False** (default): unbinding is not forced.
     * &gt; **NOTE:**  If the value of this parameter is **False**, the Master instance bound to HaVip cannot be unbound.
     * 
     */
    @Import(name="force")
    private @Nullable Output<Boolean> force;

    /**
     * @return Whether to force the ECS instance or Eni instance bound to AVIP to be unbound. The value is:
     * - **True**: Force unbinding.
     * - **False** (default): unbinding is not forced.
     * &gt; **NOTE:**  If the value of this parameter is **False**, the Master instance bound to HaVip cannot be unbound.
     * 
     */
    public Optional<Output<Boolean>> force() {
        return Optional.ofNullable(this.force);
    }

    /**
     * The ID of the HaVip instance.
     * 
     */
    @Import(name="haVipId")
    private @Nullable Output<String> haVipId;

    /**
     * @return The ID of the HaVip instance.
     * 
     */
    public Optional<Output<String>> haVipId() {
        return Optional.ofNullable(this.haVipId);
    }

    /**
     * . Field &#39;havip_id&#39; has been deprecated from provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
     * 
     * @deprecated
     * Field &#39;havip_id&#39; has been deprecated since provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'havip_id' has been deprecated since provider version 1.211.0. New field 'ha_vip_id' instead. */
    @Import(name="havipId")
    private @Nullable Output<String> havipId;

    /**
     * @return . Field &#39;havip_id&#39; has been deprecated from provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
     * 
     * @deprecated
     * Field &#39;havip_id&#39; has been deprecated since provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'havip_id' has been deprecated since provider version 1.211.0. New field 'ha_vip_id' instead. */
    public Optional<Output<String>> havipId() {
        return Optional.ofNullable(this.havipId);
    }

    /**
     * The ID of the ECS instance bound to the HaVip instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The ID of the ECS instance bound to the HaVip instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * The type of the instance associated with the VIIP.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    /**
     * @return The type of the instance associated with the VIIP.
     * 
     * The following arguments will be discarded. Please use new fields as soon as possible:
     * 
     */
    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    private HAVipAttachmentArgs() {}

    private HAVipAttachmentArgs(HAVipAttachmentArgs $) {
        this.force = $.force;
        this.haVipId = $.haVipId;
        this.havipId = $.havipId;
        this.instanceId = $.instanceId;
        this.instanceType = $.instanceType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HAVipAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HAVipAttachmentArgs $;

        public Builder() {
            $ = new HAVipAttachmentArgs();
        }

        public Builder(HAVipAttachmentArgs defaults) {
            $ = new HAVipAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param force Whether to force the ECS instance or Eni instance bound to AVIP to be unbound. The value is:
         * - **True**: Force unbinding.
         * - **False** (default): unbinding is not forced.
         * &gt; **NOTE:**  If the value of this parameter is **False**, the Master instance bound to HaVip cannot be unbound.
         * 
         * @return builder
         * 
         */
        public Builder force(@Nullable Output<Boolean> force) {
            $.force = force;
            return this;
        }

        /**
         * @param force Whether to force the ECS instance or Eni instance bound to AVIP to be unbound. The value is:
         * - **True**: Force unbinding.
         * - **False** (default): unbinding is not forced.
         * &gt; **NOTE:**  If the value of this parameter is **False**, the Master instance bound to HaVip cannot be unbound.
         * 
         * @return builder
         * 
         */
        public Builder force(Boolean force) {
            return force(Output.of(force));
        }

        /**
         * @param haVipId The ID of the HaVip instance.
         * 
         * @return builder
         * 
         */
        public Builder haVipId(@Nullable Output<String> haVipId) {
            $.haVipId = haVipId;
            return this;
        }

        /**
         * @param haVipId The ID of the HaVip instance.
         * 
         * @return builder
         * 
         */
        public Builder haVipId(String haVipId) {
            return haVipId(Output.of(haVipId));
        }

        /**
         * @param havipId . Field &#39;havip_id&#39; has been deprecated from provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;havip_id&#39; has been deprecated since provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
         * 
         */
        @Deprecated /* Field 'havip_id' has been deprecated since provider version 1.211.0. New field 'ha_vip_id' instead. */
        public Builder havipId(@Nullable Output<String> havipId) {
            $.havipId = havipId;
            return this;
        }

        /**
         * @param havipId . Field &#39;havip_id&#39; has been deprecated from provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;havip_id&#39; has been deprecated since provider version 1.211.0. New field &#39;ha_vip_id&#39; instead.
         * 
         */
        @Deprecated /* Field 'havip_id' has been deprecated since provider version 1.211.0. New field 'ha_vip_id' instead. */
        public Builder havipId(String havipId) {
            return havipId(Output.of(havipId));
        }

        /**
         * @param instanceId The ID of the ECS instance bound to the HaVip instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the ECS instance bound to the HaVip instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param instanceType The type of the instance associated with the VIIP.
         * 
         * The following arguments will be discarded. Please use new fields as soon as possible:
         * 
         * @return builder
         * 
         */
        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType The type of the instance associated with the VIIP.
         * 
         * The following arguments will be discarded. Please use new fields as soon as possible:
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        public HAVipAttachmentArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
