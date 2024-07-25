// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.maxcompute.outputs;

import com.pulumi.alicloud.maxcompute.outputs.ProjectPropertiesEncryption;
import com.pulumi.alicloud.maxcompute.outputs.ProjectPropertiesTableLifecycle;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ProjectProperties {
    /**
     * @return Whether to allow full table scan. Default: false.
     * 
     */
    private @Nullable Boolean allowFullScan;
    /**
     * @return Whether to turn on Decimal2.0.
     * 
     */
    private @Nullable Boolean enableDecimal2;
    /**
     * @return Storage encryption. For details, see [Storage Encryption](https://www.alibabacloud.com/help/en/maxcompute/security-and-compliance/storage-encryption)
     * &gt; **NOTE :**:  To enable storage encryption, you need to modify the parameters of the basic attributes of the MaxCompute project. This operation permission is authenticated by RAM, and you need to have the Super_Administrator role permission of the corresponding project.  To configure the permissions and IP whitelist parameters of the MaxCompute project, you must have the management permissions (Admin) of the corresponding project, including Super_Administrator, Admin, or custom management permissions. For more information, see the project management permissions list.  You can turn on storage encryption only for projects that have not turned on storage encryption. For projects that have turned on storage encryption, you cannot turn off storage encryption or change the encryption algorithm. See `encryption` below.
     * 
     */
    private @Nullable ProjectPropertiesEncryption encryption;
    /**
     * @return Set the number of days to retain backup data. During this time, you can restore the current version to any backup version. The value range of days is [0,30], and the default value is 1. 0 means backup is turned off. The effective policy after adjusting the backup cycle is: Extend the backup cycle: The new backup cycle takes effect on the same day. Shorten the backup cycle: The system will automatically delete backup data that has exceeded the retention cycle.
     * 
     */
    private @Nullable Integer retentionDays;
    /**
     * @return Set the maximum threshold of single SQL consumption, that is, set the ODPS. SQL. metering.value.max attribute. For details, see [Consumption Monitoring Alarm](https://www.alibabacloud.com/help/en/maxcompute/product-overview/consumption-control). Unit: scan volume (GB)* complexity. .
     * 
     */
    private @Nullable String sqlMeteringMax;
    /**
     * @return Set whether the lifecycle of the table in the project needs to be configured, that is, set the ODPS. table.lifecycle property,. See `table_lifecycle` below.
     * 
     */
    private @Nullable ProjectPropertiesTableLifecycle tableLifecycle;
    /**
     * @return Project time zone, example value: Asia/Shanghai.
     * 
     */
    private @Nullable String timezone;
    /**
     * @return Data type version. Value:(1/2/hive) 1: The original MaxCompute type system. 2: New type system introduced by MaxCompute 2.0. hive: the type system of the Hive compatibility mode introduced by MaxCompute 2.0.
     * 
     */
    private @Nullable String typeSystem;

    private ProjectProperties() {}
    /**
     * @return Whether to allow full table scan. Default: false.
     * 
     */
    public Optional<Boolean> allowFullScan() {
        return Optional.ofNullable(this.allowFullScan);
    }
    /**
     * @return Whether to turn on Decimal2.0.
     * 
     */
    public Optional<Boolean> enableDecimal2() {
        return Optional.ofNullable(this.enableDecimal2);
    }
    /**
     * @return Storage encryption. For details, see [Storage Encryption](https://www.alibabacloud.com/help/en/maxcompute/security-and-compliance/storage-encryption)
     * &gt; **NOTE :**:  To enable storage encryption, you need to modify the parameters of the basic attributes of the MaxCompute project. This operation permission is authenticated by RAM, and you need to have the Super_Administrator role permission of the corresponding project.  To configure the permissions and IP whitelist parameters of the MaxCompute project, you must have the management permissions (Admin) of the corresponding project, including Super_Administrator, Admin, or custom management permissions. For more information, see the project management permissions list.  You can turn on storage encryption only for projects that have not turned on storage encryption. For projects that have turned on storage encryption, you cannot turn off storage encryption or change the encryption algorithm. See `encryption` below.
     * 
     */
    public Optional<ProjectPropertiesEncryption> encryption() {
        return Optional.ofNullable(this.encryption);
    }
    /**
     * @return Set the number of days to retain backup data. During this time, you can restore the current version to any backup version. The value range of days is [0,30], and the default value is 1. 0 means backup is turned off. The effective policy after adjusting the backup cycle is: Extend the backup cycle: The new backup cycle takes effect on the same day. Shorten the backup cycle: The system will automatically delete backup data that has exceeded the retention cycle.
     * 
     */
    public Optional<Integer> retentionDays() {
        return Optional.ofNullable(this.retentionDays);
    }
    /**
     * @return Set the maximum threshold of single SQL consumption, that is, set the ODPS. SQL. metering.value.max attribute. For details, see [Consumption Monitoring Alarm](https://www.alibabacloud.com/help/en/maxcompute/product-overview/consumption-control). Unit: scan volume (GB)* complexity. .
     * 
     */
    public Optional<String> sqlMeteringMax() {
        return Optional.ofNullable(this.sqlMeteringMax);
    }
    /**
     * @return Set whether the lifecycle of the table in the project needs to be configured, that is, set the ODPS. table.lifecycle property,. See `table_lifecycle` below.
     * 
     */
    public Optional<ProjectPropertiesTableLifecycle> tableLifecycle() {
        return Optional.ofNullable(this.tableLifecycle);
    }
    /**
     * @return Project time zone, example value: Asia/Shanghai.
     * 
     */
    public Optional<String> timezone() {
        return Optional.ofNullable(this.timezone);
    }
    /**
     * @return Data type version. Value:(1/2/hive) 1: The original MaxCompute type system. 2: New type system introduced by MaxCompute 2.0. hive: the type system of the Hive compatibility mode introduced by MaxCompute 2.0.
     * 
     */
    public Optional<String> typeSystem() {
        return Optional.ofNullable(this.typeSystem);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ProjectProperties defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean allowFullScan;
        private @Nullable Boolean enableDecimal2;
        private @Nullable ProjectPropertiesEncryption encryption;
        private @Nullable Integer retentionDays;
        private @Nullable String sqlMeteringMax;
        private @Nullable ProjectPropertiesTableLifecycle tableLifecycle;
        private @Nullable String timezone;
        private @Nullable String typeSystem;
        public Builder() {}
        public Builder(ProjectProperties defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowFullScan = defaults.allowFullScan;
    	      this.enableDecimal2 = defaults.enableDecimal2;
    	      this.encryption = defaults.encryption;
    	      this.retentionDays = defaults.retentionDays;
    	      this.sqlMeteringMax = defaults.sqlMeteringMax;
    	      this.tableLifecycle = defaults.tableLifecycle;
    	      this.timezone = defaults.timezone;
    	      this.typeSystem = defaults.typeSystem;
        }

        @CustomType.Setter
        public Builder allowFullScan(@Nullable Boolean allowFullScan) {

            this.allowFullScan = allowFullScan;
            return this;
        }
        @CustomType.Setter
        public Builder enableDecimal2(@Nullable Boolean enableDecimal2) {

            this.enableDecimal2 = enableDecimal2;
            return this;
        }
        @CustomType.Setter
        public Builder encryption(@Nullable ProjectPropertiesEncryption encryption) {

            this.encryption = encryption;
            return this;
        }
        @CustomType.Setter
        public Builder retentionDays(@Nullable Integer retentionDays) {

            this.retentionDays = retentionDays;
            return this;
        }
        @CustomType.Setter
        public Builder sqlMeteringMax(@Nullable String sqlMeteringMax) {

            this.sqlMeteringMax = sqlMeteringMax;
            return this;
        }
        @CustomType.Setter
        public Builder tableLifecycle(@Nullable ProjectPropertiesTableLifecycle tableLifecycle) {

            this.tableLifecycle = tableLifecycle;
            return this;
        }
        @CustomType.Setter
        public Builder timezone(@Nullable String timezone) {

            this.timezone = timezone;
            return this;
        }
        @CustomType.Setter
        public Builder typeSystem(@Nullable String typeSystem) {

            this.typeSystem = typeSystem;
            return this;
        }
        public ProjectProperties build() {
            final var _resultValue = new ProjectProperties();
            _resultValue.allowFullScan = allowFullScan;
            _resultValue.enableDecimal2 = enableDecimal2;
            _resultValue.encryption = encryption;
            _resultValue.retentionDays = retentionDays;
            _resultValue.sqlMeteringMax = sqlMeteringMax;
            _resultValue.tableLifecycle = tableLifecycle;
            _resultValue.timezone = timezone;
            _resultValue.typeSystem = typeSystem;
            return _resultValue;
        }
    }
}
