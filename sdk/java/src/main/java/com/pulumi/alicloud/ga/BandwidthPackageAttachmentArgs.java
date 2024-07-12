// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class BandwidthPackageAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final BandwidthPackageAttachmentArgs Empty = new BandwidthPackageAttachmentArgs();

    /**
     * The ID of the Global Accelerator instance.
     * 
     */
    @Import(name="acceleratorId", required=true)
    private Output<String> acceleratorId;

    /**
     * @return The ID of the Global Accelerator instance.
     * 
     */
    public Output<String> acceleratorId() {
        return this.acceleratorId;
    }

    /**
     * The ID of the Bandwidth Package. **NOTE:** From version 1.192.0, `bandwidth_package_id` can be modified.
     * 
     */
    @Import(name="bandwidthPackageId", required=true)
    private Output<String> bandwidthPackageId;

    /**
     * @return The ID of the Bandwidth Package. **NOTE:** From version 1.192.0, `bandwidth_package_id` can be modified.
     * 
     */
    public Output<String> bandwidthPackageId() {
        return this.bandwidthPackageId;
    }

    private BandwidthPackageAttachmentArgs() {}

    private BandwidthPackageAttachmentArgs(BandwidthPackageAttachmentArgs $) {
        this.acceleratorId = $.acceleratorId;
        this.bandwidthPackageId = $.bandwidthPackageId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BandwidthPackageAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BandwidthPackageAttachmentArgs $;

        public Builder() {
            $ = new BandwidthPackageAttachmentArgs();
        }

        public Builder(BandwidthPackageAttachmentArgs defaults) {
            $ = new BandwidthPackageAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acceleratorId The ID of the Global Accelerator instance.
         * 
         * @return builder
         * 
         */
        public Builder acceleratorId(Output<String> acceleratorId) {
            $.acceleratorId = acceleratorId;
            return this;
        }

        /**
         * @param acceleratorId The ID of the Global Accelerator instance.
         * 
         * @return builder
         * 
         */
        public Builder acceleratorId(String acceleratorId) {
            return acceleratorId(Output.of(acceleratorId));
        }

        /**
         * @param bandwidthPackageId The ID of the Bandwidth Package. **NOTE:** From version 1.192.0, `bandwidth_package_id` can be modified.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthPackageId(Output<String> bandwidthPackageId) {
            $.bandwidthPackageId = bandwidthPackageId;
            return this;
        }

        /**
         * @param bandwidthPackageId The ID of the Bandwidth Package. **NOTE:** From version 1.192.0, `bandwidth_package_id` can be modified.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthPackageId(String bandwidthPackageId) {
            return bandwidthPackageId(Output.of(bandwidthPackageId));
        }

        public BandwidthPackageAttachmentArgs build() {
            if ($.acceleratorId == null) {
                throw new MissingRequiredPropertyException("BandwidthPackageAttachmentArgs", "acceleratorId");
            }
            if ($.bandwidthPackageId == null) {
                throw new MissingRequiredPropertyException("BandwidthPackageAttachmentArgs", "bandwidthPackageId");
            }
            return $;
        }
    }

}
