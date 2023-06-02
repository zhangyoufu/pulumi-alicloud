// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kvstore.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetZonesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetZonesPlainArgs Empty = new GetZonesPlainArgs();

    /**
     * Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
     * * product_type - (Optional, Available in v1.130.0+) The type of the service. Valid values: `Local`, `Tair_rdb`, `Tair_scm`, `Tair_essd`, `OnECS`.
     * 
     */
    @Import(name="engine")
    private @Nullable String engine;

    /**
     * @return Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
     * * product_type - (Optional, Available in v1.130.0+) The type of the service. Valid values: `Local`, `Tair_rdb`, `Tair_scm`, `Tair_essd`, `OnECS`.
     * 
     */
    public Optional<String> engine() {
        return Optional.ofNullable(this.engine);
    }

    /**
     * Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
     * 
     */
    @Import(name="instanceChargeType")
    private @Nullable String instanceChargeType;

    /**
     * @return Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
     * 
     */
    public Optional<String> instanceChargeType() {
        return Optional.ofNullable(this.instanceChargeType);
    }

    /**
     * Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch KVStore instances.
     * 
     */
    @Import(name="multi")
    private @Nullable Boolean multi;

    /**
     * @return Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch KVStore instances.
     * 
     */
    public Optional<Boolean> multi() {
        return Optional.ofNullable(this.multi);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="productType")
    private @Nullable String productType;

    public Optional<String> productType() {
        return Optional.ofNullable(this.productType);
    }

    private GetZonesPlainArgs() {}

    private GetZonesPlainArgs(GetZonesPlainArgs $) {
        this.engine = $.engine;
        this.instanceChargeType = $.instanceChargeType;
        this.multi = $.multi;
        this.outputFile = $.outputFile;
        this.productType = $.productType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetZonesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetZonesPlainArgs $;

        public Builder() {
            $ = new GetZonesPlainArgs();
        }

        public Builder(GetZonesPlainArgs defaults) {
            $ = new GetZonesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param engine Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
         * * product_type - (Optional, Available in v1.130.0+) The type of the service. Valid values: `Local`, `Tair_rdb`, `Tair_scm`, `Tair_essd`, `OnECS`.
         * 
         * @return builder
         * 
         */
        public Builder engine(@Nullable String engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param instanceChargeType Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(@Nullable String instanceChargeType) {
            $.instanceChargeType = instanceChargeType;
            return this;
        }

        /**
         * @param multi Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch KVStore instances.
         * 
         * @return builder
         * 
         */
        public Builder multi(@Nullable Boolean multi) {
            $.multi = multi;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder productType(@Nullable String productType) {
            $.productType = productType;
            return this;
        }

        public GetZonesPlainArgs build() {
            return $;
        }
    }

}
