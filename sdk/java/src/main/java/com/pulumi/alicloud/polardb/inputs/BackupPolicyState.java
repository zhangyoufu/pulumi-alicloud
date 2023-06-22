// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BackupPolicyState extends com.pulumi.resources.ResourceArgs {

    public static final BackupPolicyState Empty = new BackupPolicyState();

    /**
     * The backup frequency. Valid values are `Normal`, `2/24H`, `3/24H`, `4/24H`.Default to `Normal`.
     * 
     */
    @Import(name="backupFrequency")
    private @Nullable Output<String> backupFrequency;

    /**
     * @return The backup frequency. Valid values are `Normal`, `2/24H`, `3/24H`, `4/24H`.Default to `Normal`.
     * 
     */
    public Optional<Output<String>> backupFrequency() {
        return Optional.ofNullable(this.backupFrequency);
    }

    /**
     * Cluster backup retention days, Fixed for 7 days, not modified.
     * 
     */
    @Import(name="backupRetentionPeriod")
    private @Nullable Output<String> backupRetentionPeriod;

    /**
     * @return Cluster backup retention days, Fixed for 7 days, not modified.
     * 
     */
    public Optional<Output<String>> backupRetentionPeriod() {
        return Optional.ofNullable(this.backupRetentionPeriod);
    }

    /**
     * Specifies whether to retain backups when you delete a cluster. Valid values are `ALL`, `LATEST`, `NONE`. Default to `NONE`. Value options can refer to the latest docs [ModifyBackupPolicy](https://www.alibabacloud.com/help/en/polardb/latest/modifybackuppolicy)
     * 
     */
    @Import(name="backupRetentionPolicyOnClusterDeletion")
    private @Nullable Output<String> backupRetentionPolicyOnClusterDeletion;

    /**
     * @return Specifies whether to retain backups when you delete a cluster. Valid values are `ALL`, `LATEST`, `NONE`. Default to `NONE`. Value options can refer to the latest docs [ModifyBackupPolicy](https://www.alibabacloud.com/help/en/polardb/latest/modifybackuppolicy)
     * 
     */
    public Optional<Output<String>> backupRetentionPolicyOnClusterDeletion() {
        return Optional.ofNullable(this.backupRetentionPolicyOnClusterDeletion);
    }

    /**
     * The Id of cluster that can run database.The backup frequency. Valid values are `Normal`, `2/24H`, `3/24H`, `4/24H`.Default to `Normal`.
     * 
     */
    @Import(name="dataLevel1BackupFrequency")
    private @Nullable Output<String> dataLevel1BackupFrequency;

    /**
     * @return The Id of cluster that can run database.The backup frequency. Valid values are `Normal`, `2/24H`, `3/24H`, `4/24H`.Default to `Normal`.
     * 
     */
    public Optional<Output<String>> dataLevel1BackupFrequency() {
        return Optional.ofNullable(this.dataLevel1BackupFrequency);
    }

    /**
     * PolarDB Cluster of level-1 backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;].
     * &gt; **NOTE:** Note Select at least two values. Separate multiple values with commas (,).
     * 
     */
    @Import(name="dataLevel1BackupPeriods")
    private @Nullable Output<List<String>> dataLevel1BackupPeriods;

    /**
     * @return PolarDB Cluster of level-1 backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;].
     * &gt; **NOTE:** Note Select at least two values. Separate multiple values with commas (,).
     * 
     */
    public Optional<Output<List<String>>> dataLevel1BackupPeriods() {
        return Optional.ofNullable(this.dataLevel1BackupPeriods);
    }

    /**
     * The retention period of level-1 backups. Valid values: 3 to 14. Unit: days.
     * 
     */
    @Import(name="dataLevel1BackupRetentionPeriod")
    private @Nullable Output<Integer> dataLevel1BackupRetentionPeriod;

    /**
     * @return The retention period of level-1 backups. Valid values: 3 to 14. Unit: days.
     * 
     */
    public Optional<Output<Integer>> dataLevel1BackupRetentionPeriod() {
        return Optional.ofNullable(this.dataLevel1BackupRetentionPeriod);
    }

    /**
     * The time period during which automatic backup is performed. The format is HH: MMZ HH: MMZ (UTC time), and the entered value must be an hour apart, such as 14:00z-15:00z.
     * 
     */
    @Import(name="dataLevel1BackupTime")
    private @Nullable Output<String> dataLevel1BackupTime;

    /**
     * @return The time period during which automatic backup is performed. The format is HH: MMZ HH: MMZ (UTC time), and the entered value must be an hour apart, such as 14:00z-15:00z.
     * 
     */
    public Optional<Output<String>> dataLevel1BackupTime() {
        return Optional.ofNullable(this.dataLevel1BackupTime);
    }

    /**
     * PolarDB Cluster of level-2 backup is a cross regional backup area.
     * 
     */
    @Import(name="dataLevel2BackupAnotherRegionRegion")
    private @Nullable Output<String> dataLevel2BackupAnotherRegionRegion;

    /**
     * @return PolarDB Cluster of level-2 backup is a cross regional backup area.
     * 
     */
    public Optional<Output<String>> dataLevel2BackupAnotherRegionRegion() {
        return Optional.ofNullable(this.dataLevel2BackupAnotherRegionRegion);
    }

    /**
     * PolarDB Cluster of level-2 backup cross region backup retention period. Valid values are `0`, `30 to 7300`, `-1`. Default to `0`.
     * 
     */
    @Import(name="dataLevel2BackupAnotherRegionRetentionPeriod")
    private @Nullable Output<Integer> dataLevel2BackupAnotherRegionRetentionPeriod;

    /**
     * @return PolarDB Cluster of level-2 backup cross region backup retention period. Valid values are `0`, `30 to 7300`, `-1`. Default to `0`.
     * 
     */
    public Optional<Output<Integer>> dataLevel2BackupAnotherRegionRetentionPeriod() {
        return Optional.ofNullable(this.dataLevel2BackupAnotherRegionRetentionPeriod);
    }

    /**
     * PolarDB Cluster of level-2 backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;].
     * &gt; **NOTE:** Note Select at least two values. Separate multiple values with commas (,).
     * 
     */
    @Import(name="dataLevel2BackupPeriods")
    private @Nullable Output<List<String>> dataLevel2BackupPeriods;

    /**
     * @return PolarDB Cluster of level-2 backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;].
     * &gt; **NOTE:** Note Select at least two values. Separate multiple values with commas (,).
     * 
     */
    public Optional<Output<List<String>>> dataLevel2BackupPeriods() {
        return Optional.ofNullable(this.dataLevel2BackupPeriods);
    }

    /**
     * The retention period of level-2 backups. Valid values are `0`, `30 to 7300`, `-1`. Default to `0`.
     * 
     */
    @Import(name="dataLevel2BackupRetentionPeriod")
    private @Nullable Output<Integer> dataLevel2BackupRetentionPeriod;

    /**
     * @return The retention period of level-2 backups. Valid values are `0`, `30 to 7300`, `-1`. Default to `0`.
     * 
     */
    public Optional<Output<Integer>> dataLevel2BackupRetentionPeriod() {
        return Optional.ofNullable(this.dataLevel2BackupRetentionPeriod);
    }

    /**
     * The Id of cluster that can run database.
     * 
     */
    @Import(name="dbClusterId")
    private @Nullable Output<String> dbClusterId;

    /**
     * @return The Id of cluster that can run database.
     * 
     */
    public Optional<Output<String>> dbClusterId() {
        return Optional.ofNullable(this.dbClusterId);
    }

    /**
     * Indicates whether the log backup feature was enabled. Valid values are `0`, `1`. `1` By default, the log backup feature is enabled and cannot be disabled.
     * 
     */
    @Import(name="enableBackupLog")
    private @Nullable Output<Integer> enableBackupLog;

    /**
     * @return Indicates whether the log backup feature was enabled. Valid values are `0`, `1`. `1` By default, the log backup feature is enabled and cannot be disabled.
     * 
     */
    public Optional<Output<Integer>> enableBackupLog() {
        return Optional.ofNullable(this.enableBackupLog);
    }

    /**
     * The region in which you want to store cross-region log backups. For information about regions that support the cross-region backup feature, see [Overview.](https://www.alibabacloud.com/help/en/polardb/latest/backup-and-restoration-overview)
     * 
     */
    @Import(name="logBackupAnotherRegionRegion")
    private @Nullable Output<String> logBackupAnotherRegionRegion;

    /**
     * @return The region in which you want to store cross-region log backups. For information about regions that support the cross-region backup feature, see [Overview.](https://www.alibabacloud.com/help/en/polardb/latest/backup-and-restoration-overview)
     * 
     */
    public Optional<Output<String>> logBackupAnotherRegionRegion() {
        return Optional.ofNullable(this.logBackupAnotherRegionRegion);
    }

    /**
     * The retention period of cross-region log backups. Default value: OFF. Valid values are `0`, `30 to 7300`, `-1`.
     * &gt; **NOTE:** Note When you create a cluster, the default value of this parameter is 0.
     * 
     */
    @Import(name="logBackupAnotherRegionRetentionPeriod")
    private @Nullable Output<Integer> logBackupAnotherRegionRetentionPeriod;

    /**
     * @return The retention period of cross-region log backups. Default value: OFF. Valid values are `0`, `30 to 7300`, `-1`.
     * &gt; **NOTE:** Note When you create a cluster, the default value of this parameter is 0.
     * 
     */
    public Optional<Output<Integer>> logBackupAnotherRegionRetentionPeriod() {
        return Optional.ofNullable(this.logBackupAnotherRegionRetentionPeriod);
    }

    /**
     * The retention period of the log backups. Valid values are `3 to 7300`, `-1`.
     * 
     */
    @Import(name="logBackupRetentionPeriod")
    private @Nullable Output<Integer> logBackupRetentionPeriod;

    /**
     * @return The retention period of the log backups. Valid values are `3 to 7300`, `-1`.
     * 
     */
    public Optional<Output<Integer>> logBackupRetentionPeriod() {
        return Optional.ofNullable(this.logBackupRetentionPeriod);
    }

    /**
     * PolarDB Cluster backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;]. Default to [&#34;Tuesday&#34;, &#34;Thursday&#34;, &#34;Saturday&#34;].
     * 
     */
    @Import(name="preferredBackupPeriods")
    private @Nullable Output<List<String>> preferredBackupPeriods;

    /**
     * @return PolarDB Cluster backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;]. Default to [&#34;Tuesday&#34;, &#34;Thursday&#34;, &#34;Saturday&#34;].
     * 
     */
    public Optional<Output<List<String>>> preferredBackupPeriods() {
        return Optional.ofNullable(this.preferredBackupPeriods);
    }

    /**
     * PolarDB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to &#34;02:00Z-03:00Z&#34;. China time is 8 hours behind it.
     * 
     */
    @Import(name="preferredBackupTime")
    private @Nullable Output<String> preferredBackupTime;

    /**
     * @return PolarDB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to &#34;02:00Z-03:00Z&#34;. China time is 8 hours behind it.
     * 
     */
    public Optional<Output<String>> preferredBackupTime() {
        return Optional.ofNullable(this.preferredBackupTime);
    }

    private BackupPolicyState() {}

    private BackupPolicyState(BackupPolicyState $) {
        this.backupFrequency = $.backupFrequency;
        this.backupRetentionPeriod = $.backupRetentionPeriod;
        this.backupRetentionPolicyOnClusterDeletion = $.backupRetentionPolicyOnClusterDeletion;
        this.dataLevel1BackupFrequency = $.dataLevel1BackupFrequency;
        this.dataLevel1BackupPeriods = $.dataLevel1BackupPeriods;
        this.dataLevel1BackupRetentionPeriod = $.dataLevel1BackupRetentionPeriod;
        this.dataLevel1BackupTime = $.dataLevel1BackupTime;
        this.dataLevel2BackupAnotherRegionRegion = $.dataLevel2BackupAnotherRegionRegion;
        this.dataLevel2BackupAnotherRegionRetentionPeriod = $.dataLevel2BackupAnotherRegionRetentionPeriod;
        this.dataLevel2BackupPeriods = $.dataLevel2BackupPeriods;
        this.dataLevel2BackupRetentionPeriod = $.dataLevel2BackupRetentionPeriod;
        this.dbClusterId = $.dbClusterId;
        this.enableBackupLog = $.enableBackupLog;
        this.logBackupAnotherRegionRegion = $.logBackupAnotherRegionRegion;
        this.logBackupAnotherRegionRetentionPeriod = $.logBackupAnotherRegionRetentionPeriod;
        this.logBackupRetentionPeriod = $.logBackupRetentionPeriod;
        this.preferredBackupPeriods = $.preferredBackupPeriods;
        this.preferredBackupTime = $.preferredBackupTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackupPolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackupPolicyState $;

        public Builder() {
            $ = new BackupPolicyState();
        }

        public Builder(BackupPolicyState defaults) {
            $ = new BackupPolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param backupFrequency The backup frequency. Valid values are `Normal`, `2/24H`, `3/24H`, `4/24H`.Default to `Normal`.
         * 
         * @return builder
         * 
         */
        public Builder backupFrequency(@Nullable Output<String> backupFrequency) {
            $.backupFrequency = backupFrequency;
            return this;
        }

        /**
         * @param backupFrequency The backup frequency. Valid values are `Normal`, `2/24H`, `3/24H`, `4/24H`.Default to `Normal`.
         * 
         * @return builder
         * 
         */
        public Builder backupFrequency(String backupFrequency) {
            return backupFrequency(Output.of(backupFrequency));
        }

        /**
         * @param backupRetentionPeriod Cluster backup retention days, Fixed for 7 days, not modified.
         * 
         * @return builder
         * 
         */
        public Builder backupRetentionPeriod(@Nullable Output<String> backupRetentionPeriod) {
            $.backupRetentionPeriod = backupRetentionPeriod;
            return this;
        }

        /**
         * @param backupRetentionPeriod Cluster backup retention days, Fixed for 7 days, not modified.
         * 
         * @return builder
         * 
         */
        public Builder backupRetentionPeriod(String backupRetentionPeriod) {
            return backupRetentionPeriod(Output.of(backupRetentionPeriod));
        }

        /**
         * @param backupRetentionPolicyOnClusterDeletion Specifies whether to retain backups when you delete a cluster. Valid values are `ALL`, `LATEST`, `NONE`. Default to `NONE`. Value options can refer to the latest docs [ModifyBackupPolicy](https://www.alibabacloud.com/help/en/polardb/latest/modifybackuppolicy)
         * 
         * @return builder
         * 
         */
        public Builder backupRetentionPolicyOnClusterDeletion(@Nullable Output<String> backupRetentionPolicyOnClusterDeletion) {
            $.backupRetentionPolicyOnClusterDeletion = backupRetentionPolicyOnClusterDeletion;
            return this;
        }

        /**
         * @param backupRetentionPolicyOnClusterDeletion Specifies whether to retain backups when you delete a cluster. Valid values are `ALL`, `LATEST`, `NONE`. Default to `NONE`. Value options can refer to the latest docs [ModifyBackupPolicy](https://www.alibabacloud.com/help/en/polardb/latest/modifybackuppolicy)
         * 
         * @return builder
         * 
         */
        public Builder backupRetentionPolicyOnClusterDeletion(String backupRetentionPolicyOnClusterDeletion) {
            return backupRetentionPolicyOnClusterDeletion(Output.of(backupRetentionPolicyOnClusterDeletion));
        }

        /**
         * @param dataLevel1BackupFrequency The Id of cluster that can run database.The backup frequency. Valid values are `Normal`, `2/24H`, `3/24H`, `4/24H`.Default to `Normal`.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel1BackupFrequency(@Nullable Output<String> dataLevel1BackupFrequency) {
            $.dataLevel1BackupFrequency = dataLevel1BackupFrequency;
            return this;
        }

        /**
         * @param dataLevel1BackupFrequency The Id of cluster that can run database.The backup frequency. Valid values are `Normal`, `2/24H`, `3/24H`, `4/24H`.Default to `Normal`.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel1BackupFrequency(String dataLevel1BackupFrequency) {
            return dataLevel1BackupFrequency(Output.of(dataLevel1BackupFrequency));
        }

        /**
         * @param dataLevel1BackupPeriods PolarDB Cluster of level-1 backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;].
         * &gt; **NOTE:** Note Select at least two values. Separate multiple values with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder dataLevel1BackupPeriods(@Nullable Output<List<String>> dataLevel1BackupPeriods) {
            $.dataLevel1BackupPeriods = dataLevel1BackupPeriods;
            return this;
        }

        /**
         * @param dataLevel1BackupPeriods PolarDB Cluster of level-1 backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;].
         * &gt; **NOTE:** Note Select at least two values. Separate multiple values with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder dataLevel1BackupPeriods(List<String> dataLevel1BackupPeriods) {
            return dataLevel1BackupPeriods(Output.of(dataLevel1BackupPeriods));
        }

        /**
         * @param dataLevel1BackupPeriods PolarDB Cluster of level-1 backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;].
         * &gt; **NOTE:** Note Select at least two values. Separate multiple values with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder dataLevel1BackupPeriods(String... dataLevel1BackupPeriods) {
            return dataLevel1BackupPeriods(List.of(dataLevel1BackupPeriods));
        }

        /**
         * @param dataLevel1BackupRetentionPeriod The retention period of level-1 backups. Valid values: 3 to 14. Unit: days.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel1BackupRetentionPeriod(@Nullable Output<Integer> dataLevel1BackupRetentionPeriod) {
            $.dataLevel1BackupRetentionPeriod = dataLevel1BackupRetentionPeriod;
            return this;
        }

        /**
         * @param dataLevel1BackupRetentionPeriod The retention period of level-1 backups. Valid values: 3 to 14. Unit: days.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel1BackupRetentionPeriod(Integer dataLevel1BackupRetentionPeriod) {
            return dataLevel1BackupRetentionPeriod(Output.of(dataLevel1BackupRetentionPeriod));
        }

        /**
         * @param dataLevel1BackupTime The time period during which automatic backup is performed. The format is HH: MMZ HH: MMZ (UTC time), and the entered value must be an hour apart, such as 14:00z-15:00z.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel1BackupTime(@Nullable Output<String> dataLevel1BackupTime) {
            $.dataLevel1BackupTime = dataLevel1BackupTime;
            return this;
        }

        /**
         * @param dataLevel1BackupTime The time period during which automatic backup is performed. The format is HH: MMZ HH: MMZ (UTC time), and the entered value must be an hour apart, such as 14:00z-15:00z.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel1BackupTime(String dataLevel1BackupTime) {
            return dataLevel1BackupTime(Output.of(dataLevel1BackupTime));
        }

        /**
         * @param dataLevel2BackupAnotherRegionRegion PolarDB Cluster of level-2 backup is a cross regional backup area.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel2BackupAnotherRegionRegion(@Nullable Output<String> dataLevel2BackupAnotherRegionRegion) {
            $.dataLevel2BackupAnotherRegionRegion = dataLevel2BackupAnotherRegionRegion;
            return this;
        }

        /**
         * @param dataLevel2BackupAnotherRegionRegion PolarDB Cluster of level-2 backup is a cross regional backup area.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel2BackupAnotherRegionRegion(String dataLevel2BackupAnotherRegionRegion) {
            return dataLevel2BackupAnotherRegionRegion(Output.of(dataLevel2BackupAnotherRegionRegion));
        }

        /**
         * @param dataLevel2BackupAnotherRegionRetentionPeriod PolarDB Cluster of level-2 backup cross region backup retention period. Valid values are `0`, `30 to 7300`, `-1`. Default to `0`.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel2BackupAnotherRegionRetentionPeriod(@Nullable Output<Integer> dataLevel2BackupAnotherRegionRetentionPeriod) {
            $.dataLevel2BackupAnotherRegionRetentionPeriod = dataLevel2BackupAnotherRegionRetentionPeriod;
            return this;
        }

        /**
         * @param dataLevel2BackupAnotherRegionRetentionPeriod PolarDB Cluster of level-2 backup cross region backup retention period. Valid values are `0`, `30 to 7300`, `-1`. Default to `0`.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel2BackupAnotherRegionRetentionPeriod(Integer dataLevel2BackupAnotherRegionRetentionPeriod) {
            return dataLevel2BackupAnotherRegionRetentionPeriod(Output.of(dataLevel2BackupAnotherRegionRetentionPeriod));
        }

        /**
         * @param dataLevel2BackupPeriods PolarDB Cluster of level-2 backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;].
         * &gt; **NOTE:** Note Select at least two values. Separate multiple values with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder dataLevel2BackupPeriods(@Nullable Output<List<String>> dataLevel2BackupPeriods) {
            $.dataLevel2BackupPeriods = dataLevel2BackupPeriods;
            return this;
        }

        /**
         * @param dataLevel2BackupPeriods PolarDB Cluster of level-2 backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;].
         * &gt; **NOTE:** Note Select at least two values. Separate multiple values with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder dataLevel2BackupPeriods(List<String> dataLevel2BackupPeriods) {
            return dataLevel2BackupPeriods(Output.of(dataLevel2BackupPeriods));
        }

        /**
         * @param dataLevel2BackupPeriods PolarDB Cluster of level-2 backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;].
         * &gt; **NOTE:** Note Select at least two values. Separate multiple values with commas (,).
         * 
         * @return builder
         * 
         */
        public Builder dataLevel2BackupPeriods(String... dataLevel2BackupPeriods) {
            return dataLevel2BackupPeriods(List.of(dataLevel2BackupPeriods));
        }

        /**
         * @param dataLevel2BackupRetentionPeriod The retention period of level-2 backups. Valid values are `0`, `30 to 7300`, `-1`. Default to `0`.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel2BackupRetentionPeriod(@Nullable Output<Integer> dataLevel2BackupRetentionPeriod) {
            $.dataLevel2BackupRetentionPeriod = dataLevel2BackupRetentionPeriod;
            return this;
        }

        /**
         * @param dataLevel2BackupRetentionPeriod The retention period of level-2 backups. Valid values are `0`, `30 to 7300`, `-1`. Default to `0`.
         * 
         * @return builder
         * 
         */
        public Builder dataLevel2BackupRetentionPeriod(Integer dataLevel2BackupRetentionPeriod) {
            return dataLevel2BackupRetentionPeriod(Output.of(dataLevel2BackupRetentionPeriod));
        }

        /**
         * @param dbClusterId The Id of cluster that can run database.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(@Nullable Output<String> dbClusterId) {
            $.dbClusterId = dbClusterId;
            return this;
        }

        /**
         * @param dbClusterId The Id of cluster that can run database.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(String dbClusterId) {
            return dbClusterId(Output.of(dbClusterId));
        }

        /**
         * @param enableBackupLog Indicates whether the log backup feature was enabled. Valid values are `0`, `1`. `1` By default, the log backup feature is enabled and cannot be disabled.
         * 
         * @return builder
         * 
         */
        public Builder enableBackupLog(@Nullable Output<Integer> enableBackupLog) {
            $.enableBackupLog = enableBackupLog;
            return this;
        }

        /**
         * @param enableBackupLog Indicates whether the log backup feature was enabled. Valid values are `0`, `1`. `1` By default, the log backup feature is enabled and cannot be disabled.
         * 
         * @return builder
         * 
         */
        public Builder enableBackupLog(Integer enableBackupLog) {
            return enableBackupLog(Output.of(enableBackupLog));
        }

        /**
         * @param logBackupAnotherRegionRegion The region in which you want to store cross-region log backups. For information about regions that support the cross-region backup feature, see [Overview.](https://www.alibabacloud.com/help/en/polardb/latest/backup-and-restoration-overview)
         * 
         * @return builder
         * 
         */
        public Builder logBackupAnotherRegionRegion(@Nullable Output<String> logBackupAnotherRegionRegion) {
            $.logBackupAnotherRegionRegion = logBackupAnotherRegionRegion;
            return this;
        }

        /**
         * @param logBackupAnotherRegionRegion The region in which you want to store cross-region log backups. For information about regions that support the cross-region backup feature, see [Overview.](https://www.alibabacloud.com/help/en/polardb/latest/backup-and-restoration-overview)
         * 
         * @return builder
         * 
         */
        public Builder logBackupAnotherRegionRegion(String logBackupAnotherRegionRegion) {
            return logBackupAnotherRegionRegion(Output.of(logBackupAnotherRegionRegion));
        }

        /**
         * @param logBackupAnotherRegionRetentionPeriod The retention period of cross-region log backups. Default value: OFF. Valid values are `0`, `30 to 7300`, `-1`.
         * &gt; **NOTE:** Note When you create a cluster, the default value of this parameter is 0.
         * 
         * @return builder
         * 
         */
        public Builder logBackupAnotherRegionRetentionPeriod(@Nullable Output<Integer> logBackupAnotherRegionRetentionPeriod) {
            $.logBackupAnotherRegionRetentionPeriod = logBackupAnotherRegionRetentionPeriod;
            return this;
        }

        /**
         * @param logBackupAnotherRegionRetentionPeriod The retention period of cross-region log backups. Default value: OFF. Valid values are `0`, `30 to 7300`, `-1`.
         * &gt; **NOTE:** Note When you create a cluster, the default value of this parameter is 0.
         * 
         * @return builder
         * 
         */
        public Builder logBackupAnotherRegionRetentionPeriod(Integer logBackupAnotherRegionRetentionPeriod) {
            return logBackupAnotherRegionRetentionPeriod(Output.of(logBackupAnotherRegionRetentionPeriod));
        }

        /**
         * @param logBackupRetentionPeriod The retention period of the log backups. Valid values are `3 to 7300`, `-1`.
         * 
         * @return builder
         * 
         */
        public Builder logBackupRetentionPeriod(@Nullable Output<Integer> logBackupRetentionPeriod) {
            $.logBackupRetentionPeriod = logBackupRetentionPeriod;
            return this;
        }

        /**
         * @param logBackupRetentionPeriod The retention period of the log backups. Valid values are `3 to 7300`, `-1`.
         * 
         * @return builder
         * 
         */
        public Builder logBackupRetentionPeriod(Integer logBackupRetentionPeriod) {
            return logBackupRetentionPeriod(Output.of(logBackupRetentionPeriod));
        }

        /**
         * @param preferredBackupPeriods PolarDB Cluster backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;]. Default to [&#34;Tuesday&#34;, &#34;Thursday&#34;, &#34;Saturday&#34;].
         * 
         * @return builder
         * 
         */
        public Builder preferredBackupPeriods(@Nullable Output<List<String>> preferredBackupPeriods) {
            $.preferredBackupPeriods = preferredBackupPeriods;
            return this;
        }

        /**
         * @param preferredBackupPeriods PolarDB Cluster backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;]. Default to [&#34;Tuesday&#34;, &#34;Thursday&#34;, &#34;Saturday&#34;].
         * 
         * @return builder
         * 
         */
        public Builder preferredBackupPeriods(List<String> preferredBackupPeriods) {
            return preferredBackupPeriods(Output.of(preferredBackupPeriods));
        }

        /**
         * @param preferredBackupPeriods PolarDB Cluster backup period. Valid values: [&#34;Monday&#34;, &#34;Tuesday&#34;, &#34;Wednesday&#34;, &#34;Thursday&#34;, &#34;Friday&#34;, &#34;Saturday&#34;, &#34;Sunday&#34;]. Default to [&#34;Tuesday&#34;, &#34;Thursday&#34;, &#34;Saturday&#34;].
         * 
         * @return builder
         * 
         */
        public Builder preferredBackupPeriods(String... preferredBackupPeriods) {
            return preferredBackupPeriods(List.of(preferredBackupPeriods));
        }

        /**
         * @param preferredBackupTime PolarDB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to &#34;02:00Z-03:00Z&#34;. China time is 8 hours behind it.
         * 
         * @return builder
         * 
         */
        public Builder preferredBackupTime(@Nullable Output<String> preferredBackupTime) {
            $.preferredBackupTime = preferredBackupTime;
            return this;
        }

        /**
         * @param preferredBackupTime PolarDB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to &#34;02:00Z-03:00Z&#34;. China time is 8 hours behind it.
         * 
         * @return builder
         * 
         */
        public Builder preferredBackupTime(String preferredBackupTime) {
            return preferredBackupTime(Output.of(preferredBackupTime));
        }

        public BackupPolicyState build() {
            return $;
        }
    }

}
