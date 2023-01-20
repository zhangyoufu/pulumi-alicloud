// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NasBackupPlanState extends com.pulumi.resources.ResourceArgs {

    public static final NasBackupPlanState Empty = new NasBackupPlanState();

    /**
     * Backup type. Valid values: `COMPLETE`.
     * 
     */
    @Import(name="backupType")
    private @Nullable Output<String> backupType;

    /**
     * @return Backup type. Valid values: `COMPLETE`.
     * 
     */
    public Optional<Output<String>> backupType() {
        return Optional.ofNullable(this.backupType);
    }

    /**
     * This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
     * 
     * @deprecated
     * Field &#39;create_time&#39; has been deprecated from provider version 1.153.0.
     * 
     */
    @Deprecated /* Field 'create_time' has been deprecated from provider version 1.153.0. */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
     * 
     * @deprecated
     * Field &#39;create_time&#39; has been deprecated from provider version 1.153.0.
     * 
     */
    @Deprecated /* Field 'create_time' has been deprecated from provider version 1.153.0. */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The role name created in the original account RAM backup by the cross account managed by the current account.
     * 
     */
    @Import(name="crossAccountRoleName")
    private @Nullable Output<String> crossAccountRoleName;

    /**
     * @return The role name created in the original account RAM backup by the cross account managed by the current account.
     * 
     */
    public Optional<Output<String>> crossAccountRoleName() {
        return Optional.ofNullable(this.crossAccountRoleName);
    }

    /**
     * The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
     * 
     */
    @Import(name="crossAccountType")
    private @Nullable Output<String> crossAccountType;

    /**
     * @return The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
     * 
     */
    public Optional<Output<String>> crossAccountType() {
        return Optional.ofNullable(this.crossAccountType);
    }

    /**
     * The original account ID of the cross account backup managed by the current account.
     * 
     */
    @Import(name="crossAccountUserId")
    private @Nullable Output<Integer> crossAccountUserId;

    /**
     * @return The original account ID of the cross account backup managed by the current account.
     * 
     */
    public Optional<Output<Integer>> crossAccountUserId() {
        return Optional.ofNullable(this.crossAccountUserId);
    }

    /**
     * Whether to disable the backup task. Valid values: `true`, `false`.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Whether to disable the backup task. Valid values: `true`, `false`.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * The File System ID of Nas.
     * 
     */
    @Import(name="fileSystemId")
    private @Nullable Output<String> fileSystemId;

    /**
     * @return The File System ID of Nas.
     * 
     */
    public Optional<Output<String>> fileSystemId() {
        return Optional.ofNullable(this.fileSystemId);
    }

    /**
     * The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
     * 
     */
    @Import(name="nasBackupPlanName")
    private @Nullable Output<String> nasBackupPlanName;

    /**
     * @return The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
     * 
     */
    public Optional<Output<String>> nasBackupPlanName() {
        return Optional.ofNullable(this.nasBackupPlanName);
    }

    /**
     * This parameter specifies whether to use Windows VSS to define a backup path.
     * 
     */
    @Import(name="options")
    private @Nullable Output<String> options;

    /**
     * @return This parameter specifies whether to use Windows VSS to define a backup path.
     * 
     */
    public Optional<Output<String>> options() {
        return Optional.ofNullable(this.options);
    }

    /**
     * List of backup path. Up to 65536 characters. e.g.`[&#34;/home&#34;, &#34;/var&#34;]`. **Note** You should at least specify a backup path, empty array not allowed here.
     * 
     */
    @Import(name="paths")
    private @Nullable Output<List<String>> paths;

    /**
     * @return List of backup path. Up to 65536 characters. e.g.`[&#34;/home&#34;, &#34;/var&#34;]`. **Note** You should at least specify a backup path, empty array not allowed here.
     * 
     */
    public Optional<Output<List<String>>> paths() {
        return Optional.ofNullable(this.paths);
    }

    /**
     * Backup retention days, the minimum is 1.
     * 
     */
    @Import(name="retention")
    private @Nullable Output<String> retention;

    /**
     * @return Backup retention days, the minimum is 1.
     * 
     */
    public Optional<Output<String>> retention() {
        return Optional.ofNullable(this.retention);
    }

    /**
     * Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
     * 
     */
    @Import(name="schedule")
    private @Nullable Output<String> schedule;

    /**
     * @return Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
     * 
     */
    public Optional<Output<String>> schedule() {
        return Optional.ofNullable(this.schedule);
    }

    /**
     * The ID of Backup vault.
     * 
     */
    @Import(name="vaultId")
    private @Nullable Output<String> vaultId;

    /**
     * @return The ID of Backup vault.
     * 
     */
    public Optional<Output<String>> vaultId() {
        return Optional.ofNullable(this.vaultId);
    }

    private NasBackupPlanState() {}

    private NasBackupPlanState(NasBackupPlanState $) {
        this.backupType = $.backupType;
        this.createTime = $.createTime;
        this.crossAccountRoleName = $.crossAccountRoleName;
        this.crossAccountType = $.crossAccountType;
        this.crossAccountUserId = $.crossAccountUserId;
        this.disabled = $.disabled;
        this.fileSystemId = $.fileSystemId;
        this.nasBackupPlanName = $.nasBackupPlanName;
        this.options = $.options;
        this.paths = $.paths;
        this.retention = $.retention;
        this.schedule = $.schedule;
        this.vaultId = $.vaultId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NasBackupPlanState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NasBackupPlanState $;

        public Builder() {
            $ = new NasBackupPlanState();
        }

        public Builder(NasBackupPlanState defaults) {
            $ = new NasBackupPlanState(Objects.requireNonNull(defaults));
        }

        /**
         * @param backupType Backup type. Valid values: `COMPLETE`.
         * 
         * @return builder
         * 
         */
        public Builder backupType(@Nullable Output<String> backupType) {
            $.backupType = backupType;
            return this;
        }

        /**
         * @param backupType Backup type. Valid values: `COMPLETE`.
         * 
         * @return builder
         * 
         */
        public Builder backupType(String backupType) {
            return backupType(Output.of(backupType));
        }

        /**
         * @param createTime This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;create_time&#39; has been deprecated from provider version 1.153.0.
         * 
         */
        @Deprecated /* Field 'create_time' has been deprecated from provider version 1.153.0. */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;create_time&#39; has been deprecated from provider version 1.153.0.
         * 
         */
        @Deprecated /* Field 'create_time' has been deprecated from provider version 1.153.0. */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param crossAccountRoleName The role name created in the original account RAM backup by the cross account managed by the current account.
         * 
         * @return builder
         * 
         */
        public Builder crossAccountRoleName(@Nullable Output<String> crossAccountRoleName) {
            $.crossAccountRoleName = crossAccountRoleName;
            return this;
        }

        /**
         * @param crossAccountRoleName The role name created in the original account RAM backup by the cross account managed by the current account.
         * 
         * @return builder
         * 
         */
        public Builder crossAccountRoleName(String crossAccountRoleName) {
            return crossAccountRoleName(Output.of(crossAccountRoleName));
        }

        /**
         * @param crossAccountType The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
         * 
         * @return builder
         * 
         */
        public Builder crossAccountType(@Nullable Output<String> crossAccountType) {
            $.crossAccountType = crossAccountType;
            return this;
        }

        /**
         * @param crossAccountType The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
         * 
         * @return builder
         * 
         */
        public Builder crossAccountType(String crossAccountType) {
            return crossAccountType(Output.of(crossAccountType));
        }

        /**
         * @param crossAccountUserId The original account ID of the cross account backup managed by the current account.
         * 
         * @return builder
         * 
         */
        public Builder crossAccountUserId(@Nullable Output<Integer> crossAccountUserId) {
            $.crossAccountUserId = crossAccountUserId;
            return this;
        }

        /**
         * @param crossAccountUserId The original account ID of the cross account backup managed by the current account.
         * 
         * @return builder
         * 
         */
        public Builder crossAccountUserId(Integer crossAccountUserId) {
            return crossAccountUserId(Output.of(crossAccountUserId));
        }

        /**
         * @param disabled Whether to disable the backup task. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Whether to disable the backup task. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param fileSystemId The File System ID of Nas.
         * 
         * @return builder
         * 
         */
        public Builder fileSystemId(@Nullable Output<String> fileSystemId) {
            $.fileSystemId = fileSystemId;
            return this;
        }

        /**
         * @param fileSystemId The File System ID of Nas.
         * 
         * @return builder
         * 
         */
        public Builder fileSystemId(String fileSystemId) {
            return fileSystemId(Output.of(fileSystemId));
        }

        /**
         * @param nasBackupPlanName The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
         * 
         * @return builder
         * 
         */
        public Builder nasBackupPlanName(@Nullable Output<String> nasBackupPlanName) {
            $.nasBackupPlanName = nasBackupPlanName;
            return this;
        }

        /**
         * @param nasBackupPlanName The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
         * 
         * @return builder
         * 
         */
        public Builder nasBackupPlanName(String nasBackupPlanName) {
            return nasBackupPlanName(Output.of(nasBackupPlanName));
        }

        /**
         * @param options This parameter specifies whether to use Windows VSS to define a backup path.
         * 
         * @return builder
         * 
         */
        public Builder options(@Nullable Output<String> options) {
            $.options = options;
            return this;
        }

        /**
         * @param options This parameter specifies whether to use Windows VSS to define a backup path.
         * 
         * @return builder
         * 
         */
        public Builder options(String options) {
            return options(Output.of(options));
        }

        /**
         * @param paths List of backup path. Up to 65536 characters. e.g.`[&#34;/home&#34;, &#34;/var&#34;]`. **Note** You should at least specify a backup path, empty array not allowed here.
         * 
         * @return builder
         * 
         */
        public Builder paths(@Nullable Output<List<String>> paths) {
            $.paths = paths;
            return this;
        }

        /**
         * @param paths List of backup path. Up to 65536 characters. e.g.`[&#34;/home&#34;, &#34;/var&#34;]`. **Note** You should at least specify a backup path, empty array not allowed here.
         * 
         * @return builder
         * 
         */
        public Builder paths(List<String> paths) {
            return paths(Output.of(paths));
        }

        /**
         * @param paths List of backup path. Up to 65536 characters. e.g.`[&#34;/home&#34;, &#34;/var&#34;]`. **Note** You should at least specify a backup path, empty array not allowed here.
         * 
         * @return builder
         * 
         */
        public Builder paths(String... paths) {
            return paths(List.of(paths));
        }

        /**
         * @param retention Backup retention days, the minimum is 1.
         * 
         * @return builder
         * 
         */
        public Builder retention(@Nullable Output<String> retention) {
            $.retention = retention;
            return this;
        }

        /**
         * @param retention Backup retention days, the minimum is 1.
         * 
         * @return builder
         * 
         */
        public Builder retention(String retention) {
            return retention(Output.of(retention));
        }

        /**
         * @param schedule Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
         * 
         * @return builder
         * 
         */
        public Builder schedule(@Nullable Output<String> schedule) {
            $.schedule = schedule;
            return this;
        }

        /**
         * @param schedule Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
         * 
         * @return builder
         * 
         */
        public Builder schedule(String schedule) {
            return schedule(Output.of(schedule));
        }

        /**
         * @param vaultId The ID of Backup vault.
         * 
         * @return builder
         * 
         */
        public Builder vaultId(@Nullable Output<String> vaultId) {
            $.vaultId = vaultId;
            return this;
        }

        /**
         * @param vaultId The ID of Backup vault.
         * 
         * @return builder
         * 
         */
        public Builder vaultId(String vaultId) {
            return vaultId(Output.of(vaultId));
        }

        public NasBackupPlanState build() {
            return $;
        }
    }

}
