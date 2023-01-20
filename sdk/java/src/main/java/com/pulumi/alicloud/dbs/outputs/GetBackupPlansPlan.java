// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dbs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBackupPlansPlan {
    /**
     * @return The ID of the backup gateway.
     * 
     */
    private String backupGatewayId;
    /**
     * @return The Backup method.
     * 
     */
    private String backupMethod;
    /**
     * @return The backup object.
     * 
     */
    private String backupObjects;
    /**
     * @return Full backup cycle.
     * 
     */
    private String backupPeriod;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String backupPlanId;
    /**
     * @return The name of the resource.
     * 
     */
    private String backupPlanName;
    /**
     * @return The retention time of backup data.
     * 
     */
    private Integer backupRetentionPeriod;
    /**
     * @return The start time of full Backup.
     * 
     */
    private String backupStartTime;
    /**
     * @return Built-in storage type.
     * 
     */
    private String backupStorageType;
    /**
     * @return The UID that is backed up across Alibaba cloud accounts.
     * 
     */
    private String crossAliyunId;
    /**
     * @return The name of the RAM role that is backed up across Alibaba cloud accounts.
     * 
     */
    private String crossRoleName;
    /**
     * @return The database type.
     * 
     */
    private String databaseType;
    /**
     * @return The storage time for conversion to archive cold standby is 365 days by default.
     * 
     */
    private Integer duplicationArchivePeriod;
    /**
     * @return The storage time is converted to low-frequency access. The default time is 180 days.
     * 
     */
    private Integer duplicationInfrequentAccessPeriod;
    /**
     * @return Whether to enable incremental log Backup.
     * 
     */
    private Boolean enableBackupLog;
    /**
     * @return The ID of the Backup Plan.
     * 
     */
    private String id;
    /**
     * @return The Instance class.
     * 
     */
    private String instanceClass;
    /**
     * @return The OSS Bucket name.
     * 
     */
    private String ossBucketName;
    /**
     * @return The payment type of the resource.
     * 
     */
    private String paymentType;
    /**
     * @return The ID of the resource group.
     * 
     */
    private String resourceGroupId;
    /**
     * @return The name of the database.
     * 
     */
    private String sourceEndpointDatabaseName;
    /**
     * @return The ID of the database instance.
     * 
     */
    private String sourceEndpointInstanceId;
    /**
     * @return The location of the database.
     * 
     */
    private String sourceEndpointInstanceType;
    /**
     * @return The region of the database.
     * 
     */
    private String sourceEndpointRegion;
    /**
     * @return The Oracle SID name.
     * 
     */
    private String sourceEndpointSid;
    /**
     * @return The source endpoint username.
     * 
     */
    private String sourceEndpointUserName;
    /**
     * @return The status of the resource.
     * 
     */
    private String status;

    private GetBackupPlansPlan() {}
    /**
     * @return The ID of the backup gateway.
     * 
     */
    public String backupGatewayId() {
        return this.backupGatewayId;
    }
    /**
     * @return The Backup method.
     * 
     */
    public String backupMethod() {
        return this.backupMethod;
    }
    /**
     * @return The backup object.
     * 
     */
    public String backupObjects() {
        return this.backupObjects;
    }
    /**
     * @return Full backup cycle.
     * 
     */
    public String backupPeriod() {
        return this.backupPeriod;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String backupPlanId() {
        return this.backupPlanId;
    }
    /**
     * @return The name of the resource.
     * 
     */
    public String backupPlanName() {
        return this.backupPlanName;
    }
    /**
     * @return The retention time of backup data.
     * 
     */
    public Integer backupRetentionPeriod() {
        return this.backupRetentionPeriod;
    }
    /**
     * @return The start time of full Backup.
     * 
     */
    public String backupStartTime() {
        return this.backupStartTime;
    }
    /**
     * @return Built-in storage type.
     * 
     */
    public String backupStorageType() {
        return this.backupStorageType;
    }
    /**
     * @return The UID that is backed up across Alibaba cloud accounts.
     * 
     */
    public String crossAliyunId() {
        return this.crossAliyunId;
    }
    /**
     * @return The name of the RAM role that is backed up across Alibaba cloud accounts.
     * 
     */
    public String crossRoleName() {
        return this.crossRoleName;
    }
    /**
     * @return The database type.
     * 
     */
    public String databaseType() {
        return this.databaseType;
    }
    /**
     * @return The storage time for conversion to archive cold standby is 365 days by default.
     * 
     */
    public Integer duplicationArchivePeriod() {
        return this.duplicationArchivePeriod;
    }
    /**
     * @return The storage time is converted to low-frequency access. The default time is 180 days.
     * 
     */
    public Integer duplicationInfrequentAccessPeriod() {
        return this.duplicationInfrequentAccessPeriod;
    }
    /**
     * @return Whether to enable incremental log Backup.
     * 
     */
    public Boolean enableBackupLog() {
        return this.enableBackupLog;
    }
    /**
     * @return The ID of the Backup Plan.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The Instance class.
     * 
     */
    public String instanceClass() {
        return this.instanceClass;
    }
    /**
     * @return The OSS Bucket name.
     * 
     */
    public String ossBucketName() {
        return this.ossBucketName;
    }
    /**
     * @return The payment type of the resource.
     * 
     */
    public String paymentType() {
        return this.paymentType;
    }
    /**
     * @return The ID of the resource group.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * @return The name of the database.
     * 
     */
    public String sourceEndpointDatabaseName() {
        return this.sourceEndpointDatabaseName;
    }
    /**
     * @return The ID of the database instance.
     * 
     */
    public String sourceEndpointInstanceId() {
        return this.sourceEndpointInstanceId;
    }
    /**
     * @return The location of the database.
     * 
     */
    public String sourceEndpointInstanceType() {
        return this.sourceEndpointInstanceType;
    }
    /**
     * @return The region of the database.
     * 
     */
    public String sourceEndpointRegion() {
        return this.sourceEndpointRegion;
    }
    /**
     * @return The Oracle SID name.
     * 
     */
    public String sourceEndpointSid() {
        return this.sourceEndpointSid;
    }
    /**
     * @return The source endpoint username.
     * 
     */
    public String sourceEndpointUserName() {
        return this.sourceEndpointUserName;
    }
    /**
     * @return The status of the resource.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackupPlansPlan defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String backupGatewayId;
        private String backupMethod;
        private String backupObjects;
        private String backupPeriod;
        private String backupPlanId;
        private String backupPlanName;
        private Integer backupRetentionPeriod;
        private String backupStartTime;
        private String backupStorageType;
        private String crossAliyunId;
        private String crossRoleName;
        private String databaseType;
        private Integer duplicationArchivePeriod;
        private Integer duplicationInfrequentAccessPeriod;
        private Boolean enableBackupLog;
        private String id;
        private String instanceClass;
        private String ossBucketName;
        private String paymentType;
        private String resourceGroupId;
        private String sourceEndpointDatabaseName;
        private String sourceEndpointInstanceId;
        private String sourceEndpointInstanceType;
        private String sourceEndpointRegion;
        private String sourceEndpointSid;
        private String sourceEndpointUserName;
        private String status;
        public Builder() {}
        public Builder(GetBackupPlansPlan defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backupGatewayId = defaults.backupGatewayId;
    	      this.backupMethod = defaults.backupMethod;
    	      this.backupObjects = defaults.backupObjects;
    	      this.backupPeriod = defaults.backupPeriod;
    	      this.backupPlanId = defaults.backupPlanId;
    	      this.backupPlanName = defaults.backupPlanName;
    	      this.backupRetentionPeriod = defaults.backupRetentionPeriod;
    	      this.backupStartTime = defaults.backupStartTime;
    	      this.backupStorageType = defaults.backupStorageType;
    	      this.crossAliyunId = defaults.crossAliyunId;
    	      this.crossRoleName = defaults.crossRoleName;
    	      this.databaseType = defaults.databaseType;
    	      this.duplicationArchivePeriod = defaults.duplicationArchivePeriod;
    	      this.duplicationInfrequentAccessPeriod = defaults.duplicationInfrequentAccessPeriod;
    	      this.enableBackupLog = defaults.enableBackupLog;
    	      this.id = defaults.id;
    	      this.instanceClass = defaults.instanceClass;
    	      this.ossBucketName = defaults.ossBucketName;
    	      this.paymentType = defaults.paymentType;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.sourceEndpointDatabaseName = defaults.sourceEndpointDatabaseName;
    	      this.sourceEndpointInstanceId = defaults.sourceEndpointInstanceId;
    	      this.sourceEndpointInstanceType = defaults.sourceEndpointInstanceType;
    	      this.sourceEndpointRegion = defaults.sourceEndpointRegion;
    	      this.sourceEndpointSid = defaults.sourceEndpointSid;
    	      this.sourceEndpointUserName = defaults.sourceEndpointUserName;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder backupGatewayId(String backupGatewayId) {
            this.backupGatewayId = Objects.requireNonNull(backupGatewayId);
            return this;
        }
        @CustomType.Setter
        public Builder backupMethod(String backupMethod) {
            this.backupMethod = Objects.requireNonNull(backupMethod);
            return this;
        }
        @CustomType.Setter
        public Builder backupObjects(String backupObjects) {
            this.backupObjects = Objects.requireNonNull(backupObjects);
            return this;
        }
        @CustomType.Setter
        public Builder backupPeriod(String backupPeriod) {
            this.backupPeriod = Objects.requireNonNull(backupPeriod);
            return this;
        }
        @CustomType.Setter
        public Builder backupPlanId(String backupPlanId) {
            this.backupPlanId = Objects.requireNonNull(backupPlanId);
            return this;
        }
        @CustomType.Setter
        public Builder backupPlanName(String backupPlanName) {
            this.backupPlanName = Objects.requireNonNull(backupPlanName);
            return this;
        }
        @CustomType.Setter
        public Builder backupRetentionPeriod(Integer backupRetentionPeriod) {
            this.backupRetentionPeriod = Objects.requireNonNull(backupRetentionPeriod);
            return this;
        }
        @CustomType.Setter
        public Builder backupStartTime(String backupStartTime) {
            this.backupStartTime = Objects.requireNonNull(backupStartTime);
            return this;
        }
        @CustomType.Setter
        public Builder backupStorageType(String backupStorageType) {
            this.backupStorageType = Objects.requireNonNull(backupStorageType);
            return this;
        }
        @CustomType.Setter
        public Builder crossAliyunId(String crossAliyunId) {
            this.crossAliyunId = Objects.requireNonNull(crossAliyunId);
            return this;
        }
        @CustomType.Setter
        public Builder crossRoleName(String crossRoleName) {
            this.crossRoleName = Objects.requireNonNull(crossRoleName);
            return this;
        }
        @CustomType.Setter
        public Builder databaseType(String databaseType) {
            this.databaseType = Objects.requireNonNull(databaseType);
            return this;
        }
        @CustomType.Setter
        public Builder duplicationArchivePeriod(Integer duplicationArchivePeriod) {
            this.duplicationArchivePeriod = Objects.requireNonNull(duplicationArchivePeriod);
            return this;
        }
        @CustomType.Setter
        public Builder duplicationInfrequentAccessPeriod(Integer duplicationInfrequentAccessPeriod) {
            this.duplicationInfrequentAccessPeriod = Objects.requireNonNull(duplicationInfrequentAccessPeriod);
            return this;
        }
        @CustomType.Setter
        public Builder enableBackupLog(Boolean enableBackupLog) {
            this.enableBackupLog = Objects.requireNonNull(enableBackupLog);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceClass(String instanceClass) {
            this.instanceClass = Objects.requireNonNull(instanceClass);
            return this;
        }
        @CustomType.Setter
        public Builder ossBucketName(String ossBucketName) {
            this.ossBucketName = Objects.requireNonNull(ossBucketName);
            return this;
        }
        @CustomType.Setter
        public Builder paymentType(String paymentType) {
            this.paymentType = Objects.requireNonNull(paymentType);
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(String resourceGroupId) {
            this.resourceGroupId = Objects.requireNonNull(resourceGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointDatabaseName(String sourceEndpointDatabaseName) {
            this.sourceEndpointDatabaseName = Objects.requireNonNull(sourceEndpointDatabaseName);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointInstanceId(String sourceEndpointInstanceId) {
            this.sourceEndpointInstanceId = Objects.requireNonNull(sourceEndpointInstanceId);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointInstanceType(String sourceEndpointInstanceType) {
            this.sourceEndpointInstanceType = Objects.requireNonNull(sourceEndpointInstanceType);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointRegion(String sourceEndpointRegion) {
            this.sourceEndpointRegion = Objects.requireNonNull(sourceEndpointRegion);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointSid(String sourceEndpointSid) {
            this.sourceEndpointSid = Objects.requireNonNull(sourceEndpointSid);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointUserName(String sourceEndpointUserName) {
            this.sourceEndpointUserName = Objects.requireNonNull(sourceEndpointUserName);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetBackupPlansPlan build() {
            final var o = new GetBackupPlansPlan();
            o.backupGatewayId = backupGatewayId;
            o.backupMethod = backupMethod;
            o.backupObjects = backupObjects;
            o.backupPeriod = backupPeriod;
            o.backupPlanId = backupPlanId;
            o.backupPlanName = backupPlanName;
            o.backupRetentionPeriod = backupRetentionPeriod;
            o.backupStartTime = backupStartTime;
            o.backupStorageType = backupStorageType;
            o.crossAliyunId = crossAliyunId;
            o.crossRoleName = crossRoleName;
            o.databaseType = databaseType;
            o.duplicationArchivePeriod = duplicationArchivePeriod;
            o.duplicationInfrequentAccessPeriod = duplicationInfrequentAccessPeriod;
            o.enableBackupLog = enableBackupLog;
            o.id = id;
            o.instanceClass = instanceClass;
            o.ossBucketName = ossBucketName;
            o.paymentType = paymentType;
            o.resourceGroupId = resourceGroupId;
            o.sourceEndpointDatabaseName = sourceEndpointDatabaseName;
            o.sourceEndpointInstanceId = sourceEndpointInstanceId;
            o.sourceEndpointInstanceType = sourceEndpointInstanceType;
            o.sourceEndpointRegion = sourceEndpointRegion;
            o.sourceEndpointSid = sourceEndpointSid;
            o.sourceEndpointUserName = sourceEndpointUserName;
            o.status = status;
            return o;
        }
    }
}
