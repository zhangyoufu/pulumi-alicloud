// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.outputs;

import com.pulumi.alicloud.hbr.outputs.PolicyRuleRetentionRule;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PolicyRule {
    /**
     * @return This parameter is required only when the value of `RuleType` is **TRANSITION. The minimum value is 30, and the Retention-ArchiveDays needs to be greater than or equal to 60.
     * 
     */
    private @Nullable Integer archiveDays;
    /**
     * @return This parameter is required only when the `RuleType` value is **BACKUP. Backup Type.
     * 
     */
    private @Nullable String backupType;
    /**
     * @return This parameter is required only when `RuleType` is set to `BACKUP`.
     * 
     */
    private @Nullable Integer keepLatestSnapshots;
    /**
     * @return Only when the `RuleType` value is.
     * 
     */
    private @Nullable String replicationRegionId;
    /**
     * @return Retention time, in days.
     * 
     */
    private @Nullable Integer retention;
    /**
     * @return This parameter is required only when the value of `RuleType` is `TRANSITION`. See `retention_rules` below.
     * 
     */
    private @Nullable List<PolicyRuleRetentionRule> retentionRules;
    /**
     * @return Rule ID.
     * 
     */
    private @Nullable String ruleId;
    /**
     * @return Rule Type.
     * 
     */
    private String ruleType;
    /**
     * @return This parameter is required only if you set the `RuleType` parameter to `BACKUP`. This parameter specifies the backup schedule settings. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is complete. For example, `I|1631685600|P1D` specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.  *   startTime: the time at which the system starts to run a backup job. The time must follow the UNIX time format. Unit: seconds. *   interval: the interval at which the system runs a backup job. The interval must follow the ISO 8601 standard. For example, PT1H specifies an interval of one hour. P1D specifies an interval of one day.
     * 
     */
    private @Nullable String schedule;
    /**
     * @return Vault ID.
     * 
     */
    private @Nullable String vaultId;

    private PolicyRule() {}
    /**
     * @return This parameter is required only when the value of `RuleType` is **TRANSITION. The minimum value is 30, and the Retention-ArchiveDays needs to be greater than or equal to 60.
     * 
     */
    public Optional<Integer> archiveDays() {
        return Optional.ofNullable(this.archiveDays);
    }
    /**
     * @return This parameter is required only when the `RuleType` value is **BACKUP. Backup Type.
     * 
     */
    public Optional<String> backupType() {
        return Optional.ofNullable(this.backupType);
    }
    /**
     * @return This parameter is required only when `RuleType` is set to `BACKUP`.
     * 
     */
    public Optional<Integer> keepLatestSnapshots() {
        return Optional.ofNullable(this.keepLatestSnapshots);
    }
    /**
     * @return Only when the `RuleType` value is.
     * 
     */
    public Optional<String> replicationRegionId() {
        return Optional.ofNullable(this.replicationRegionId);
    }
    /**
     * @return Retention time, in days.
     * 
     */
    public Optional<Integer> retention() {
        return Optional.ofNullable(this.retention);
    }
    /**
     * @return This parameter is required only when the value of `RuleType` is `TRANSITION`. See `retention_rules` below.
     * 
     */
    public List<PolicyRuleRetentionRule> retentionRules() {
        return this.retentionRules == null ? List.of() : this.retentionRules;
    }
    /**
     * @return Rule ID.
     * 
     */
    public Optional<String> ruleId() {
        return Optional.ofNullable(this.ruleId);
    }
    /**
     * @return Rule Type.
     * 
     */
    public String ruleType() {
        return this.ruleType;
    }
    /**
     * @return This parameter is required only if you set the `RuleType` parameter to `BACKUP`. This parameter specifies the backup schedule settings. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is complete. For example, `I|1631685600|P1D` specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.  *   startTime: the time at which the system starts to run a backup job. The time must follow the UNIX time format. Unit: seconds. *   interval: the interval at which the system runs a backup job. The interval must follow the ISO 8601 standard. For example, PT1H specifies an interval of one hour. P1D specifies an interval of one day.
     * 
     */
    public Optional<String> schedule() {
        return Optional.ofNullable(this.schedule);
    }
    /**
     * @return Vault ID.
     * 
     */
    public Optional<String> vaultId() {
        return Optional.ofNullable(this.vaultId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PolicyRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer archiveDays;
        private @Nullable String backupType;
        private @Nullable Integer keepLatestSnapshots;
        private @Nullable String replicationRegionId;
        private @Nullable Integer retention;
        private @Nullable List<PolicyRuleRetentionRule> retentionRules;
        private @Nullable String ruleId;
        private String ruleType;
        private @Nullable String schedule;
        private @Nullable String vaultId;
        public Builder() {}
        public Builder(PolicyRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.archiveDays = defaults.archiveDays;
    	      this.backupType = defaults.backupType;
    	      this.keepLatestSnapshots = defaults.keepLatestSnapshots;
    	      this.replicationRegionId = defaults.replicationRegionId;
    	      this.retention = defaults.retention;
    	      this.retentionRules = defaults.retentionRules;
    	      this.ruleId = defaults.ruleId;
    	      this.ruleType = defaults.ruleType;
    	      this.schedule = defaults.schedule;
    	      this.vaultId = defaults.vaultId;
        }

        @CustomType.Setter
        public Builder archiveDays(@Nullable Integer archiveDays) {

            this.archiveDays = archiveDays;
            return this;
        }
        @CustomType.Setter
        public Builder backupType(@Nullable String backupType) {

            this.backupType = backupType;
            return this;
        }
        @CustomType.Setter
        public Builder keepLatestSnapshots(@Nullable Integer keepLatestSnapshots) {

            this.keepLatestSnapshots = keepLatestSnapshots;
            return this;
        }
        @CustomType.Setter
        public Builder replicationRegionId(@Nullable String replicationRegionId) {

            this.replicationRegionId = replicationRegionId;
            return this;
        }
        @CustomType.Setter
        public Builder retention(@Nullable Integer retention) {

            this.retention = retention;
            return this;
        }
        @CustomType.Setter
        public Builder retentionRules(@Nullable List<PolicyRuleRetentionRule> retentionRules) {

            this.retentionRules = retentionRules;
            return this;
        }
        public Builder retentionRules(PolicyRuleRetentionRule... retentionRules) {
            return retentionRules(List.of(retentionRules));
        }
        @CustomType.Setter
        public Builder ruleId(@Nullable String ruleId) {

            this.ruleId = ruleId;
            return this;
        }
        @CustomType.Setter
        public Builder ruleType(String ruleType) {
            if (ruleType == null) {
              throw new MissingRequiredPropertyException("PolicyRule", "ruleType");
            }
            this.ruleType = ruleType;
            return this;
        }
        @CustomType.Setter
        public Builder schedule(@Nullable String schedule) {

            this.schedule = schedule;
            return this;
        }
        @CustomType.Setter
        public Builder vaultId(@Nullable String vaultId) {

            this.vaultId = vaultId;
            return this;
        }
        public PolicyRule build() {
            final var _resultValue = new PolicyRule();
            _resultValue.archiveDays = archiveDays;
            _resultValue.backupType = backupType;
            _resultValue.keepLatestSnapshots = keepLatestSnapshots;
            _resultValue.replicationRegionId = replicationRegionId;
            _resultValue.retention = retention;
            _resultValue.retentionRules = retentionRules;
            _resultValue.ruleId = ruleId;
            _resultValue.ruleType = ruleType;
            _resultValue.schedule = schedule;
            _resultValue.vaultId = vaultId;
            return _resultValue;
        }
    }
}
