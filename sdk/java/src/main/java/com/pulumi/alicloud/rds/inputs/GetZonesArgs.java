// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetZonesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetZonesArgs Empty = new GetZonesArgs();

    /**
     * DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`, `serverless_basic`, `cluster`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
     * 
     */
    @Import(name="category")
    private @Nullable Output<String> category;

    /**
     * @return DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`, `serverless_basic`, `cluster`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
     * 
     */
    public Optional<Output<String>> category() {
        return Optional.ofNullable(this.category);
    }

    @Import(name="dbInstanceClass")
    private @Nullable Output<String> dbInstanceClass;

    public Optional<Output<String>> dbInstanceClass() {
        return Optional.ofNullable(this.dbInstanceClass);
    }

    /**
     * The DB instance storage space required by the user. Valid values: &#34;cloud_ssd&#34;, &#34;local_ssd&#34;, &#34;cloud_essd&#34;, &#34;cloud_essd2&#34;, &#34;cloud_essd3&#34;.
     * 
     */
    @Import(name="dbInstanceStorageType")
    private @Nullable Output<String> dbInstanceStorageType;

    /**
     * @return The DB instance storage space required by the user. Valid values: &#34;cloud_ssd&#34;, &#34;local_ssd&#34;, &#34;cloud_essd&#34;, &#34;cloud_essd2&#34;, &#34;cloud_essd3&#34;.
     * 
     */
    public Optional<Output<String>> dbInstanceStorageType() {
        return Optional.ofNullable(this.dbInstanceStorageType);
    }

    /**
     * Database type. Valid values: &#34;MySQL&#34;, &#34;SQLServer&#34;, &#34;PostgreSQL&#34;, &#34;PPAS&#34;, &#34;MariaDB&#34;. If not set, it will match all of engines.
     * 
     */
    @Import(name="engine")
    private @Nullable Output<String> engine;

    /**
     * @return Database type. Valid values: &#34;MySQL&#34;, &#34;SQLServer&#34;, &#34;PostgreSQL&#34;, &#34;PPAS&#34;, &#34;MariaDB&#34;. If not set, it will match all of engines.
     * 
     */
    public Optional<Output<String>> engine() {
        return Optional.ofNullable(this.engine);
    }

    /**
     * Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     * 
     */
    @Import(name="engineVersion")
    private @Nullable Output<String> engineVersion;

    /**
     * @return Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     * 
     */
    public Optional<Output<String>> engineVersion() {
        return Optional.ofNullable(this.engineVersion);
    }

    /**
     * Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid` and `Serverless`. Default to `PostPaid`.
     * 
     */
    @Import(name="instanceChargeType")
    private @Nullable Output<String> instanceChargeType;

    /**
     * @return Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid` and `Serverless`. Default to `PostPaid`.
     * 
     */
    public Optional<Output<String>> instanceChargeType() {
        return Optional.ofNullable(this.instanceChargeType);
    }

    /**
     * It has been deprecated from version 1.137.0 and using `multi_zone` instead.
     * 
     */
    @Import(name="multi")
    private @Nullable Output<Boolean> multi;

    /**
     * @return It has been deprecated from version 1.137.0 and using `multi_zone` instead.
     * 
     */
    public Optional<Output<Boolean>> multi() {
        return Optional.ofNullable(this.multi);
    }

    /**
     * Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
     * 
     */
    @Import(name="multiZone")
    private @Nullable Output<Boolean> multiZone;

    /**
     * @return Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
     * 
     */
    public Optional<Output<Boolean>> multiZone() {
        return Optional.ofNullable(this.multiZone);
    }

    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetZonesArgs() {}

    private GetZonesArgs(GetZonesArgs $) {
        this.category = $.category;
        this.dbInstanceClass = $.dbInstanceClass;
        this.dbInstanceStorageType = $.dbInstanceStorageType;
        this.engine = $.engine;
        this.engineVersion = $.engineVersion;
        this.instanceChargeType = $.instanceChargeType;
        this.multi = $.multi;
        this.multiZone = $.multiZone;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetZonesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetZonesArgs $;

        public Builder() {
            $ = new GetZonesArgs();
        }

        public Builder(GetZonesArgs defaults) {
            $ = new GetZonesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param category DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`, `serverless_basic`, `cluster`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
         * 
         * @return builder
         * 
         */
        public Builder category(@Nullable Output<String> category) {
            $.category = category;
            return this;
        }

        /**
         * @param category DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`, `AlwaysOn`, `serverless_basic`, `cluster`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
         * 
         * @return builder
         * 
         */
        public Builder category(String category) {
            return category(Output.of(category));
        }

        public Builder dbInstanceClass(@Nullable Output<String> dbInstanceClass) {
            $.dbInstanceClass = dbInstanceClass;
            return this;
        }

        public Builder dbInstanceClass(String dbInstanceClass) {
            return dbInstanceClass(Output.of(dbInstanceClass));
        }

        /**
         * @param dbInstanceStorageType The DB instance storage space required by the user. Valid values: &#34;cloud_ssd&#34;, &#34;local_ssd&#34;, &#34;cloud_essd&#34;, &#34;cloud_essd2&#34;, &#34;cloud_essd3&#34;.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceStorageType(@Nullable Output<String> dbInstanceStorageType) {
            $.dbInstanceStorageType = dbInstanceStorageType;
            return this;
        }

        /**
         * @param dbInstanceStorageType The DB instance storage space required by the user. Valid values: &#34;cloud_ssd&#34;, &#34;local_ssd&#34;, &#34;cloud_essd&#34;, &#34;cloud_essd2&#34;, &#34;cloud_essd3&#34;.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceStorageType(String dbInstanceStorageType) {
            return dbInstanceStorageType(Output.of(dbInstanceStorageType));
        }

        /**
         * @param engine Database type. Valid values: &#34;MySQL&#34;, &#34;SQLServer&#34;, &#34;PostgreSQL&#34;, &#34;PPAS&#34;, &#34;MariaDB&#34;. If not set, it will match all of engines.
         * 
         * @return builder
         * 
         */
        public Builder engine(@Nullable Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine Database type. Valid values: &#34;MySQL&#34;, &#34;SQLServer&#34;, &#34;PostgreSQL&#34;, &#34;PPAS&#34;, &#34;MariaDB&#34;. If not set, it will match all of engines.
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param engineVersion Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(@Nullable Output<String> engineVersion) {
            $.engineVersion = engineVersion;
            return this;
        }

        /**
         * @param engineVersion Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
         * 
         * @return builder
         * 
         */
        public Builder engineVersion(String engineVersion) {
            return engineVersion(Output.of(engineVersion));
        }

        /**
         * @param instanceChargeType Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid` and `Serverless`. Default to `PostPaid`.
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(@Nullable Output<String> instanceChargeType) {
            $.instanceChargeType = instanceChargeType;
            return this;
        }

        /**
         * @param instanceChargeType Filter the results by a specific instance charge type. Valid values: `PrePaid` and `PostPaid` and `Serverless`. Default to `PostPaid`.
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(String instanceChargeType) {
            return instanceChargeType(Output.of(instanceChargeType));
        }

        /**
         * @param multi It has been deprecated from version 1.137.0 and using `multi_zone` instead.
         * 
         * @return builder
         * 
         */
        public Builder multi(@Nullable Output<Boolean> multi) {
            $.multi = multi;
            return this;
        }

        /**
         * @param multi It has been deprecated from version 1.137.0 and using `multi_zone` instead.
         * 
         * @return builder
         * 
         */
        public Builder multi(Boolean multi) {
            return multi(Output.of(multi));
        }

        /**
         * @param multiZone Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
         * 
         * @return builder
         * 
         */
        public Builder multiZone(@Nullable Output<Boolean> multiZone) {
            $.multiZone = multiZone;
            return this;
        }

        /**
         * @param multiZone Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
         * 
         * @return builder
         * 
         */
        public Builder multiZone(Boolean multiZone) {
            return multiZone(Output.of(multiZone));
        }

        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        public GetZonesArgs build() {
            return $;
        }
    }

}
