// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNodeClassesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNodeClassesArgs Empty = new GetNodeClassesArgs();

    /**
     * The PolarDB node cluster series.
     * 
     */
    @Import(name="category")
    private @Nullable Output<String> category;

    /**
     * @return The PolarDB node cluster series.
     * 
     */
    public Optional<Output<String>> category() {
        return Optional.ofNullable(this.category);
    }

    /**
     * The PolarDB node class type by the user.
     * 
     */
    @Import(name="dbNodeClass")
    private @Nullable Output<String> dbNodeClass;

    /**
     * @return The PolarDB node class type by the user.
     * 
     */
    public Optional<Output<String>> dbNodeClass() {
        return Optional.ofNullable(this.dbNodeClass);
    }

    /**
     * Database type. Options are `MySQL`, `PostgreSQL`, `Oracle`. If db_type is set, db_version also needs to be set.
     * 
     */
    @Import(name="dbType")
    private @Nullable Output<String> dbType;

    /**
     * @return Database type. Options are `MySQL`, `PostgreSQL`, `Oracle`. If db_type is set, db_version also needs to be set.
     * 
     */
    public Optional<Output<String>> dbType() {
        return Optional.ofNullable(this.dbType);
    }

    /**
     * Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/98169.htm) `DBVersion`. If db_version is set, db_type also needs to be set.
     * 
     */
    @Import(name="dbVersion")
    private @Nullable Output<String> dbVersion;

    /**
     * @return Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/98169.htm) `DBVersion`. If db_version is set, db_type also needs to be set.
     * 
     */
    public Optional<Output<String>> dbVersion() {
        return Optional.ofNullable(this.dbVersion);
    }

    /**
     * File name where to save data source results (after running `pulumi up`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi up`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`.
     * 
     */
    @Import(name="payType", required=true)
    private Output<String> payType;

    /**
     * @return Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`.
     * 
     */
    public Output<String> payType() {
        return this.payType;
    }

    /**
     * The Region to launch the PolarDB cluster.
     * 
     */
    @Import(name="regionId")
    private @Nullable Output<String> regionId;

    /**
     * @return The Region to launch the PolarDB cluster.
     * 
     */
    public Optional<Output<String>> regionId() {
        return Optional.ofNullable(this.regionId);
    }

    /**
     * The Zone to launch the PolarDB cluster.
     * 
     */
    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    /**
     * @return The Zone to launch the PolarDB cluster.
     * 
     */
    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private GetNodeClassesArgs() {}

    private GetNodeClassesArgs(GetNodeClassesArgs $) {
        this.category = $.category;
        this.dbNodeClass = $.dbNodeClass;
        this.dbType = $.dbType;
        this.dbVersion = $.dbVersion;
        this.outputFile = $.outputFile;
        this.payType = $.payType;
        this.regionId = $.regionId;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNodeClassesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNodeClassesArgs $;

        public Builder() {
            $ = new GetNodeClassesArgs();
        }

        public Builder(GetNodeClassesArgs defaults) {
            $ = new GetNodeClassesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param category The PolarDB node cluster series.
         * 
         * @return builder
         * 
         */
        public Builder category(@Nullable Output<String> category) {
            $.category = category;
            return this;
        }

        /**
         * @param category The PolarDB node cluster series.
         * 
         * @return builder
         * 
         */
        public Builder category(String category) {
            return category(Output.of(category));
        }

        /**
         * @param dbNodeClass The PolarDB node class type by the user.
         * 
         * @return builder
         * 
         */
        public Builder dbNodeClass(@Nullable Output<String> dbNodeClass) {
            $.dbNodeClass = dbNodeClass;
            return this;
        }

        /**
         * @param dbNodeClass The PolarDB node class type by the user.
         * 
         * @return builder
         * 
         */
        public Builder dbNodeClass(String dbNodeClass) {
            return dbNodeClass(Output.of(dbNodeClass));
        }

        /**
         * @param dbType Database type. Options are `MySQL`, `PostgreSQL`, `Oracle`. If db_type is set, db_version also needs to be set.
         * 
         * @return builder
         * 
         */
        public Builder dbType(@Nullable Output<String> dbType) {
            $.dbType = dbType;
            return this;
        }

        /**
         * @param dbType Database type. Options are `MySQL`, `PostgreSQL`, `Oracle`. If db_type is set, db_version also needs to be set.
         * 
         * @return builder
         * 
         */
        public Builder dbType(String dbType) {
            return dbType(Output.of(dbType));
        }

        /**
         * @param dbVersion Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/98169.htm) `DBVersion`. If db_version is set, db_type also needs to be set.
         * 
         * @return builder
         * 
         */
        public Builder dbVersion(@Nullable Output<String> dbVersion) {
            $.dbVersion = dbVersion;
            return this;
        }

        /**
         * @param dbVersion Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/98169.htm) `DBVersion`. If db_version is set, db_type also needs to be set.
         * 
         * @return builder
         * 
         */
        public Builder dbVersion(String dbVersion) {
            return dbVersion(Output.of(dbVersion));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi up`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi up`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param payType Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`.
         * 
         * @return builder
         * 
         */
        public Builder payType(Output<String> payType) {
            $.payType = payType;
            return this;
        }

        /**
         * @param payType Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`.
         * 
         * @return builder
         * 
         */
        public Builder payType(String payType) {
            return payType(Output.of(payType));
        }

        /**
         * @param regionId The Region to launch the PolarDB cluster.
         * 
         * @return builder
         * 
         */
        public Builder regionId(@Nullable Output<String> regionId) {
            $.regionId = regionId;
            return this;
        }

        /**
         * @param regionId The Region to launch the PolarDB cluster.
         * 
         * @return builder
         * 
         */
        public Builder regionId(String regionId) {
            return regionId(Output.of(regionId));
        }

        /**
         * @param zoneId The Zone to launch the PolarDB cluster.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The Zone to launch the PolarDB cluster.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public GetNodeClassesArgs build() {
            $.payType = Objects.requireNonNull($.payType, "expected parameter 'payType' to be non-null");
            return $;
        }
    }

}
