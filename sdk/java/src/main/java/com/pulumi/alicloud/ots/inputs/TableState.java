// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots.inputs;

import com.pulumi.alicloud.ots.inputs.TableDefinedColumnArgs;
import com.pulumi.alicloud.ots.inputs.TablePrimaryKeyArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TableState extends com.pulumi.resources.ResourceArgs {

    public static final TableState Empty = new TableState();

    /**
     * The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `defined_column` should not be more than 32. See `defined_column` below.
     * 
     */
    @Import(name="definedColumns")
    private @Nullable Output<List<TableDefinedColumnArgs>> definedColumns;

    /**
     * @return The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `defined_column` should not be more than 32. See `defined_column` below.
     * 
     */
    public Optional<Output<List<TableDefinedColumnArgs>>> definedColumns() {
        return Optional.ofNullable(this.definedColumns);
    }

    /**
     * The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
     * 
     */
    @Import(name="deviationCellVersionInSec")
    private @Nullable Output<String> deviationCellVersionInSec;

    /**
     * @return The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
     * 
     */
    public Optional<Output<String>> deviationCellVersionInSec() {
        return Optional.ofNullable(this.deviationCellVersionInSec);
    }

    /**
     * Whether enable OTS server side encryption. Default value is false.
     * 
     */
    @Import(name="enableSse")
    private @Nullable Output<Boolean> enableSse;

    /**
     * @return Whether enable OTS server side encryption. Default value is false.
     * 
     */
    public Optional<Output<Boolean>> enableSse() {
        return Optional.ofNullable(this.enableSse);
    }

    /**
     * The name of the OTS instance in which table will located.
     * 
     */
    @Import(name="instanceName")
    private @Nullable Output<String> instanceName;

    /**
     * @return The name of the OTS instance in which table will located.
     * 
     */
    public Optional<Output<String>> instanceName() {
        return Optional.ofNullable(this.instanceName);
    }

    /**
     * The maximum number of versions stored in this table. The valid value is 1-2147483647.
     * 
     */
    @Import(name="maxVersion")
    private @Nullable Output<Integer> maxVersion;

    /**
     * @return The maximum number of versions stored in this table. The valid value is 1-2147483647.
     * 
     */
    public Optional<Output<Integer>> maxVersion() {
        return Optional.ofNullable(this.maxVersion);
    }

    /**
     * The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primary_key` should not be less than one and not be more than four. See `primary_key` below.
     * 
     */
    @Import(name="primaryKeys")
    private @Nullable Output<List<TablePrimaryKeyArgs>> primaryKeys;

    /**
     * @return The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primary_key` should not be less than one and not be more than four. See `primary_key` below.
     * 
     */
    public Optional<Output<List<TablePrimaryKeyArgs>>> primaryKeys() {
        return Optional.ofNullable(this.primaryKeys);
    }

    /**
     * The key type of OTS server side encryption. Only `SSE_KMS_SERVICE` is allowed.
     * 
     */
    @Import(name="sseKeyType")
    private @Nullable Output<String> sseKeyType;

    /**
     * @return The key type of OTS server side encryption. Only `SSE_KMS_SERVICE` is allowed.
     * 
     */
    public Optional<Output<String>> sseKeyType() {
        return Optional.ofNullable(this.sseKeyType);
    }

    /**
     * The table name of the OTS instance. If changed, a new table would be created.
     * 
     */
    @Import(name="tableName")
    private @Nullable Output<String> tableName;

    /**
     * @return The table name of the OTS instance. If changed, a new table would be created.
     * 
     */
    public Optional<Output<String>> tableName() {
        return Optional.ofNullable(this.tableName);
    }

    /**
     * The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
     * 
     */
    @Import(name="timeToLive")
    private @Nullable Output<Integer> timeToLive;

    /**
     * @return The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
     * 
     */
    public Optional<Output<Integer>> timeToLive() {
        return Optional.ofNullable(this.timeToLive);
    }

    private TableState() {}

    private TableState(TableState $) {
        this.definedColumns = $.definedColumns;
        this.deviationCellVersionInSec = $.deviationCellVersionInSec;
        this.enableSse = $.enableSse;
        this.instanceName = $.instanceName;
        this.maxVersion = $.maxVersion;
        this.primaryKeys = $.primaryKeys;
        this.sseKeyType = $.sseKeyType;
        this.tableName = $.tableName;
        this.timeToLive = $.timeToLive;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TableState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TableState $;

        public Builder() {
            $ = new TableState();
        }

        public Builder(TableState defaults) {
            $ = new TableState(Objects.requireNonNull(defaults));
        }

        /**
         * @param definedColumns The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `defined_column` should not be more than 32. See `defined_column` below.
         * 
         * @return builder
         * 
         */
        public Builder definedColumns(@Nullable Output<List<TableDefinedColumnArgs>> definedColumns) {
            $.definedColumns = definedColumns;
            return this;
        }

        /**
         * @param definedColumns The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `defined_column` should not be more than 32. See `defined_column` below.
         * 
         * @return builder
         * 
         */
        public Builder definedColumns(List<TableDefinedColumnArgs> definedColumns) {
            return definedColumns(Output.of(definedColumns));
        }

        /**
         * @param definedColumns The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `defined_column` should not be more than 32. See `defined_column` below.
         * 
         * @return builder
         * 
         */
        public Builder definedColumns(TableDefinedColumnArgs... definedColumns) {
            return definedColumns(List.of(definedColumns));
        }

        /**
         * @param deviationCellVersionInSec The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
         * 
         * @return builder
         * 
         */
        public Builder deviationCellVersionInSec(@Nullable Output<String> deviationCellVersionInSec) {
            $.deviationCellVersionInSec = deviationCellVersionInSec;
            return this;
        }

        /**
         * @param deviationCellVersionInSec The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
         * 
         * @return builder
         * 
         */
        public Builder deviationCellVersionInSec(String deviationCellVersionInSec) {
            return deviationCellVersionInSec(Output.of(deviationCellVersionInSec));
        }

        /**
         * @param enableSse Whether enable OTS server side encryption. Default value is false.
         * 
         * @return builder
         * 
         */
        public Builder enableSse(@Nullable Output<Boolean> enableSse) {
            $.enableSse = enableSse;
            return this;
        }

        /**
         * @param enableSse Whether enable OTS server side encryption. Default value is false.
         * 
         * @return builder
         * 
         */
        public Builder enableSse(Boolean enableSse) {
            return enableSse(Output.of(enableSse));
        }

        /**
         * @param instanceName The name of the OTS instance in which table will located.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(@Nullable Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName The name of the OTS instance in which table will located.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        /**
         * @param maxVersion The maximum number of versions stored in this table. The valid value is 1-2147483647.
         * 
         * @return builder
         * 
         */
        public Builder maxVersion(@Nullable Output<Integer> maxVersion) {
            $.maxVersion = maxVersion;
            return this;
        }

        /**
         * @param maxVersion The maximum number of versions stored in this table. The valid value is 1-2147483647.
         * 
         * @return builder
         * 
         */
        public Builder maxVersion(Integer maxVersion) {
            return maxVersion(Output.of(maxVersion));
        }

        /**
         * @param primaryKeys The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primary_key` should not be less than one and not be more than four. See `primary_key` below.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeys(@Nullable Output<List<TablePrimaryKeyArgs>> primaryKeys) {
            $.primaryKeys = primaryKeys;
            return this;
        }

        /**
         * @param primaryKeys The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primary_key` should not be less than one and not be more than four. See `primary_key` below.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeys(List<TablePrimaryKeyArgs> primaryKeys) {
            return primaryKeys(Output.of(primaryKeys));
        }

        /**
         * @param primaryKeys The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primary_key` should not be less than one and not be more than four. See `primary_key` below.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeys(TablePrimaryKeyArgs... primaryKeys) {
            return primaryKeys(List.of(primaryKeys));
        }

        /**
         * @param sseKeyType The key type of OTS server side encryption. Only `SSE_KMS_SERVICE` is allowed.
         * 
         * @return builder
         * 
         */
        public Builder sseKeyType(@Nullable Output<String> sseKeyType) {
            $.sseKeyType = sseKeyType;
            return this;
        }

        /**
         * @param sseKeyType The key type of OTS server side encryption. Only `SSE_KMS_SERVICE` is allowed.
         * 
         * @return builder
         * 
         */
        public Builder sseKeyType(String sseKeyType) {
            return sseKeyType(Output.of(sseKeyType));
        }

        /**
         * @param tableName The table name of the OTS instance. If changed, a new table would be created.
         * 
         * @return builder
         * 
         */
        public Builder tableName(@Nullable Output<String> tableName) {
            $.tableName = tableName;
            return this;
        }

        /**
         * @param tableName The table name of the OTS instance. If changed, a new table would be created.
         * 
         * @return builder
         * 
         */
        public Builder tableName(String tableName) {
            return tableName(Output.of(tableName));
        }

        /**
         * @param timeToLive The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
         * 
         * @return builder
         * 
         */
        public Builder timeToLive(@Nullable Output<Integer> timeToLive) {
            $.timeToLive = timeToLive;
            return this;
        }

        /**
         * @param timeToLive The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
         * 
         * @return builder
         * 
         */
        public Builder timeToLive(Integer timeToLive) {
            return timeToLive(Output.of(timeToLive));
        }

        public TableState build() {
            return $;
        }
    }

}
