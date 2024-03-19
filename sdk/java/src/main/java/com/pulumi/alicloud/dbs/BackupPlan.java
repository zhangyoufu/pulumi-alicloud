// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dbs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dbs.BackupPlanArgs;
import com.pulumi.alicloud.dbs.inputs.BackupPlanState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DBS Backup Plan resource.
 * 
 * For information about DBS Backup Plan and how to use it, see [What is Backup Plan](https://www.alibabacloud.com/help/en/dbs/developer-reference/api-dbs-2019-03-06-createandstartbackupplan).
 * 
 * &gt; **NOTE:** Available since v1.185.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.rds.RdsFunctions;
 * import com.pulumi.alicloud.rds.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.rds.inputs.GetInstanceClassesArgs;
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.vpc.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.rds.Instance;
 * import com.pulumi.alicloud.rds.InstanceArgs;
 * import com.pulumi.alicloud.rds.Database;
 * import com.pulumi.alicloud.rds.DatabaseArgs;
 * import com.pulumi.alicloud.rds.RdsAccount;
 * import com.pulumi.alicloud.rds.RdsAccountArgs;
 * import com.pulumi.alicloud.rds.AccountPrivilege;
 * import com.pulumi.alicloud.rds.AccountPrivilegeArgs;
 * import com.pulumi.alicloud.dbs.BackupPlan;
 * import com.pulumi.alicloud.dbs.BackupPlanArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups(GetResourceGroupsArgs.builder()
 *             .status(&#34;OK&#34;)
 *             .build());
 * 
 *         final var defaultZones = RdsFunctions.getZones(GetZonesArgs.builder()
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;8.0&#34;)
 *             .instanceChargeType(&#34;PostPaid&#34;)
 *             .category(&#34;HighAvailability&#34;)
 *             .dbInstanceStorageType(&#34;cloud_essd&#34;)
 *             .build());
 * 
 *         final var defaultInstanceClasses = RdsFunctions.getInstanceClasses(GetInstanceClassesArgs.builder()
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;8.0&#34;)
 *             .category(&#34;HighAvailability&#34;)
 *             .dbInstanceStorageType(&#34;cloud_essd&#34;)
 *             .instanceChargeType(&#34;PostPaid&#34;)
 *             .build());
 * 
 *         final var defaultNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex(&#34;^default-NODELETING&#34;)
 *             .build());
 * 
 *         final var defaultSwitches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         final var vswitchId = defaultSwitches.applyValue(getSwitchesResult -&gt; getSwitchesResult.ids()[0]);
 * 
 *         final var zoneId = defaultZones.applyValue(getZonesResult -&gt; getZonesResult.ids()[0]);
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;8.0&#34;)
 *             .dbInstanceStorageType(&#34;cloud_essd&#34;)
 *             .instanceType(defaultInstanceClasses.applyValue(getInstanceClassesResult -&gt; getInstanceClassesResult.instanceClasses()[0].instanceClass()))
 *             .instanceStorage(defaultInstanceClasses.applyValue(getInstanceClassesResult -&gt; getInstanceClassesResult.instanceClasses()[0].storageRange().min()))
 *             .vswitchId(vswitchId)
 *             .instanceName(name)
 *             .build());
 * 
 *         var defaultDatabase = new Database(&#34;defaultDatabase&#34;, DatabaseArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .build());
 * 
 *         var defaultRdsAccount = new RdsAccount(&#34;defaultRdsAccount&#34;, RdsAccountArgs.builder()        
 *             .dbInstanceId(defaultInstance.id())
 *             .accountName(&#34;tfnormal000&#34;)
 *             .accountPassword(&#34;Test12345&#34;)
 *             .build());
 * 
 *         var defaultAccountPrivilege = new AccountPrivilege(&#34;defaultAccountPrivilege&#34;, AccountPrivilegeArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .accountName(defaultRdsAccount.accountName())
 *             .privilege(&#34;ReadWrite&#34;)
 *             .dbNames(defaultDatabase.name())
 *             .build());
 * 
 *         var defaultBackupPlan = new BackupPlan(&#34;defaultBackupPlan&#34;, BackupPlanArgs.builder()        
 *             .backupPlanName(name)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .instanceClass(&#34;xlarge&#34;)
 *             .backupMethod(&#34;logical&#34;)
 *             .databaseType(&#34;MySQL&#34;)
 *             .databaseRegion(&#34;cn-hangzhou&#34;)
 *             .storageRegion(&#34;cn-hangzhou&#34;)
 *             .instanceType(&#34;RDS&#34;)
 *             .sourceEndpointInstanceType(&#34;RDS&#34;)
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .sourceEndpointRegion(&#34;cn-hangzhou&#34;)
 *             .sourceEndpointInstanceId(defaultInstance.id())
 *             .sourceEndpointUserName(defaultAccountPrivilege.accountName())
 *             .sourceEndpointPassword(defaultRdsAccount.accountPassword())
 *             .backupObjects(defaultDatabase.name().applyValue(name -&gt; String.format(&#34;[{{\&#34;DBName\&#34;:\&#34;%s\&#34;}}]&#34;, name)))
 *             .backupPeriod(&#34;Monday&#34;)
 *             .backupStartTime(&#34;14:22&#34;)
 *             .backupStorageType(&#34;system&#34;)
 *             .backupRetentionPeriod(740)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * DBS Backup Plan can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dbs/backupPlan:BackupPlan example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dbs/backupPlan:BackupPlan")
public class BackupPlan extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the backup gateway. This parameter is required when the `source_endpoint_instance_type` is `Agent`.
     * 
     */
    @Export(name="backupGatewayId", refs={String.class}, tree="[0]")
    private Output<String> backupGatewayId;

    /**
     * @return The ID of the backup gateway. This parameter is required when the `source_endpoint_instance_type` is `Agent`.
     * 
     */
    public Output<String> backupGatewayId() {
        return this.backupGatewayId;
    }
    /**
     * The backup log interval seconds.
     * 
     */
    @Export(name="backupLogIntervalSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> backupLogIntervalSeconds;

    /**
     * @return The backup log interval seconds.
     * 
     */
    public Output<Optional<Integer>> backupLogIntervalSeconds() {
        return Codegen.optional(this.backupLogIntervalSeconds);
    }
    /**
     * Backup method. Valid values: `duplication`, `logical`, `physical`.
     * 
     */
    @Export(name="backupMethod", refs={String.class}, tree="[0]")
    private Output<String> backupMethod;

    /**
     * @return Backup method. Valid values: `duplication`, `logical`, `physical`.
     * 
     */
    public Output<String> backupMethod() {
        return this.backupMethod;
    }
    /**
     * The backup object.
     * 
     */
    @Export(name="backupObjects", refs={String.class}, tree="[0]")
    private Output<String> backupObjects;

    /**
     * @return The backup object.
     * 
     */
    public Output<String> backupObjects() {
        return this.backupObjects;
    }
    /**
     * Full backup cycle, Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. supports the selection of multiple fetch values, separated by English commas (,).
     * 
     */
    @Export(name="backupPeriod", refs={String.class}, tree="[0]")
    private Output<String> backupPeriod;

    /**
     * @return Full backup cycle, Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. supports the selection of multiple fetch values, separated by English commas (,).
     * 
     */
    public Output<String> backupPeriod() {
        return this.backupPeriod;
    }
    /**
     * The name of the resource.
     * 
     */
    @Export(name="backupPlanName", refs={String.class}, tree="[0]")
    private Output<String> backupPlanName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> backupPlanName() {
        return this.backupPlanName;
    }
    /**
     * The backup rate limit.
     * 
     */
    @Export(name="backupRateLimit", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backupRateLimit;

    /**
     * @return The backup rate limit.
     * 
     */
    public Output<Optional<String>> backupRateLimit() {
        return Codegen.optional(this.backupRateLimit);
    }
    /**
     * The retention time of backup data. Valid values: 0 to 1825. Default value: 730 days.
     * 
     */
    @Export(name="backupRetentionPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> backupRetentionPeriod;

    /**
     * @return The retention time of backup data. Valid values: 0 to 1825. Default value: 730 days.
     * 
     */
    public Output<Integer> backupRetentionPeriod() {
        return this.backupRetentionPeriod;
    }
    /**
     * The backup speed limit.
     * 
     */
    @Export(name="backupSpeedLimit", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backupSpeedLimit;

    /**
     * @return The backup speed limit.
     * 
     */
    public Output<Optional<String>> backupSpeedLimit() {
        return Codegen.optional(this.backupSpeedLimit);
    }
    /**
     * The start time of full Backup. The format is `&lt;I&gt; HH:mm&lt;/I&gt;` Z(UTC time).
     * 
     */
    @Export(name="backupStartTime", refs={String.class}, tree="[0]")
    private Output<String> backupStartTime;

    /**
     * @return The start time of full Backup. The format is `&lt;I&gt; HH:mm&lt;/I&gt;` Z(UTC time).
     * 
     */
    public Output<String> backupStartTime() {
        return this.backupStartTime;
    }
    /**
     * Built-in storage type, Valid values: `system`.
     * 
     */
    @Export(name="backupStorageType", refs={String.class}, tree="[0]")
    private Output<String> backupStorageType;

    /**
     * @return Built-in storage type, Valid values: `system`.
     * 
     */
    public Output<String> backupStorageType() {
        return this.backupStorageType;
    }
    /**
     * The backup strategy type. Valid values: `simple`, `manual`.
     * 
     */
    @Export(name="backupStrategyType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backupStrategyType;

    /**
     * @return The backup strategy type. Valid values: `simple`, `manual`.
     * 
     */
    public Output<Optional<String>> backupStrategyType() {
        return Codegen.optional(this.backupStrategyType);
    }
    /**
     * The UID that is backed up across Alibaba cloud accounts.
     * 
     */
    @Export(name="crossAliyunId", refs={String.class}, tree="[0]")
    private Output<String> crossAliyunId;

    /**
     * @return The UID that is backed up across Alibaba cloud accounts.
     * 
     */
    public Output<String> crossAliyunId() {
        return this.crossAliyunId;
    }
    /**
     * The name of the RAM role that is backed up across Alibaba cloud accounts.
     * 
     */
    @Export(name="crossRoleName", refs={String.class}, tree="[0]")
    private Output<String> crossRoleName;

    /**
     * @return The name of the RAM role that is backed up across Alibaba cloud accounts.
     * 
     */
    public Output<String> crossRoleName() {
        return this.crossRoleName;
    }
    /**
     * The database region.
     * 
     */
    @Export(name="databaseRegion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> databaseRegion;

    /**
     * @return The database region.
     * 
     */
    public Output<Optional<String>> databaseRegion() {
        return Codegen.optional(this.databaseRegion);
    }
    /**
     * Database type. Valid values: `DRDS`, `FIle`, `MSSQL`, `MariaDB`, `MongoDB`, `MySQL`, `Oracle`, `PPAS`, `PostgreSQL`, `Redis`.
     * 
     */
    @Export(name="databaseType", refs={String.class}, tree="[0]")
    private Output<String> databaseType;

    /**
     * @return Database type. Valid values: `DRDS`, `FIle`, `MSSQL`, `MariaDB`, `MongoDB`, `MySQL`, `Oracle`, `PPAS`, `PostgreSQL`, `Redis`.
     * 
     */
    public Output<String> databaseType() {
        return this.databaseType;
    }
    /**
     * The storage time for conversion to archive cold standby is 365 days by default.
     * 
     */
    @Export(name="duplicationArchivePeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> duplicationArchivePeriod;

    /**
     * @return The storage time for conversion to archive cold standby is 365 days by default.
     * 
     */
    public Output<Integer> duplicationArchivePeriod() {
        return this.duplicationArchivePeriod;
    }
    /**
     * The storage time is converted to low-frequency access. The default time is 180 days.
     * 
     */
    @Export(name="duplicationInfrequentAccessPeriod", refs={Integer.class}, tree="[0]")
    private Output<Integer> duplicationInfrequentAccessPeriod;

    /**
     * @return The storage time is converted to low-frequency access. The default time is 180 days.
     * 
     */
    public Output<Integer> duplicationInfrequentAccessPeriod() {
        return this.duplicationInfrequentAccessPeriod;
    }
    /**
     * Whether to enable incremental log Backup.
     * 
     */
    @Export(name="enableBackupLog", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enableBackupLog;

    /**
     * @return Whether to enable incremental log Backup.
     * 
     */
    public Output<Boolean> enableBackupLog() {
        return this.enableBackupLog;
    }
    /**
     * The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`.
     * 
     */
    @Export(name="instanceClass", refs={String.class}, tree="[0]")
    private Output<String> instanceClass;

    /**
     * @return The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`.
     * 
     */
    public Output<String> instanceClass() {
        return this.instanceClass;
    }
    /**
     * The instance type. Valid values: `RDS`, `PolarDB`, `DDS`, `Kvstore`, `Other`.
     * 
     */
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instanceType;

    /**
     * @return The instance type. Valid values: `RDS`, `PolarDB`, `DDS`, `Kvstore`, `Other`.
     * 
     */
    public Output<Optional<String>> instanceType() {
        return Codegen.optional(this.instanceType);
    }
    /**
     * The OSS Bucket name. The system automatically generates a new name by default.
     * 
     */
    @Export(name="ossBucketName", refs={String.class}, tree="[0]")
    private Output<String> ossBucketName;

    /**
     * @return The OSS Bucket name. The system automatically generates a new name by default.
     * 
     */
    public Output<String> ossBucketName() {
        return this.ossBucketName;
    }
    /**
     * The payment type of the resource. Valid values: `PayAsYouGo`, `Subscription`.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output<String> paymentType;

    /**
     * @return The payment type of the resource. Valid values: `PayAsYouGo`, `Subscription`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * Specify that the prepaid instance is of the package year or monthly type. Valid values: `Month`, `Year`.
     * 
     */
    @Export(name="period", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> period;

    /**
     * @return Specify that the prepaid instance is of the package year or monthly type. Valid values: `Month`, `Year`.
     * 
     */
    public Output<Optional<String>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The name of the database. This parameter is required when the `database_type` is `PostgreSQL` or `MongoDB`.
     * 
     */
    @Export(name="sourceEndpointDatabaseName", refs={String.class}, tree="[0]")
    private Output<String> sourceEndpointDatabaseName;

    /**
     * @return The name of the database. This parameter is required when the `database_type` is `PostgreSQL` or `MongoDB`.
     * 
     */
    public Output<String> sourceEndpointDatabaseName() {
        return this.sourceEndpointDatabaseName;
    }
    /**
     * The ID of the database instance. This parameter is required when the `source_endpoint_instance_type` is `RDS`, `ECS`, `DDS`, or `Express`.
     * 
     */
    @Export(name="sourceEndpointInstanceId", refs={String.class}, tree="[0]")
    private Output<String> sourceEndpointInstanceId;

    /**
     * @return The ID of the database instance. This parameter is required when the `source_endpoint_instance_type` is `RDS`, `ECS`, `DDS`, or `Express`.
     * 
     */
    public Output<String> sourceEndpointInstanceId() {
        return this.sourceEndpointInstanceId;
    }
    /**
     * The location of the database. Valid values: `RDS`, `ECS`, `Express`, `Agent`, `DDS`, `Other`.
     * 
     */
    @Export(name="sourceEndpointInstanceType", refs={String.class}, tree="[0]")
    private Output<String> sourceEndpointInstanceType;

    /**
     * @return The location of the database. Valid values: `RDS`, `ECS`, `Express`, `Agent`, `DDS`, `Other`.
     * 
     */
    public Output<String> sourceEndpointInstanceType() {
        return this.sourceEndpointInstanceType;
    }
    /**
     * The source endpoint ip.
     * 
     */
    @Export(name="sourceEndpointIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceEndpointIp;

    /**
     * @return The source endpoint ip.
     * 
     */
    public Output<Optional<String>> sourceEndpointIp() {
        return Codegen.optional(this.sourceEndpointIp);
    }
    /**
     * Oracle SID name. This parameter is required when the `database_type` is `Oracle`.
     * 
     */
    @Export(name="sourceEndpointOracleSid", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceEndpointOracleSid;

    /**
     * @return Oracle SID name. This parameter is required when the `database_type` is `Oracle`.
     * 
     */
    public Output<Optional<String>> sourceEndpointOracleSid() {
        return Codegen.optional(this.sourceEndpointOracleSid);
    }
    /**
     * The source endpoint password.  This parameter is not required when the `database_type` is `Redis`, or when the `source_endpoint_instance_type` is `Agent` and the `database_type` is `MSSQL`. This parameter is required in other scenarios.
     * 
     */
    @Export(name="sourceEndpointPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceEndpointPassword;

    /**
     * @return The source endpoint password.  This parameter is not required when the `database_type` is `Redis`, or when the `source_endpoint_instance_type` is `Agent` and the `database_type` is `MSSQL`. This parameter is required in other scenarios.
     * 
     */
    public Output<Optional<String>> sourceEndpointPassword() {
        return Codegen.optional(this.sourceEndpointPassword);
    }
    /**
     * The source endpoint port.
     * 
     */
    @Export(name="sourceEndpointPort", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sourceEndpointPort;

    /**
     * @return The source endpoint port.
     * 
     */
    public Output<Optional<Integer>> sourceEndpointPort() {
        return Codegen.optional(this.sourceEndpointPort);
    }
    /**
     * The region of the database. This parameter is required when the `source_endpoint_instance_type` is `RDS`, `ECS`, `DDS`, `Express`, or `Agent`.
     * 
     */
    @Export(name="sourceEndpointRegion", refs={String.class}, tree="[0]")
    private Output<String> sourceEndpointRegion;

    /**
     * @return The region of the database. This parameter is required when the `source_endpoint_instance_type` is `RDS`, `ECS`, `DDS`, `Express`, or `Agent`.
     * 
     */
    public Output<String> sourceEndpointRegion() {
        return this.sourceEndpointRegion;
    }
    /**
     * Oracle SID name. This parameter is required when the `database_type` is `Oracle`.
     * 
     */
    @Export(name="sourceEndpointSid", refs={String.class}, tree="[0]")
    private Output<String> sourceEndpointSid;

    /**
     * @return Oracle SID name. This parameter is required when the `database_type` is `Oracle`.
     * 
     */
    public Output<String> sourceEndpointSid() {
        return this.sourceEndpointSid;
    }
    /**
     * The source endpoint username. This parameter is not required when the `database_type` is `Redis`, or when the `source_endpoint_instance_type` is `Agent` and the `database_type` is `MSSQL`. This parameter is required in other scenarios.
     * 
     */
    @Export(name="sourceEndpointUserName", refs={String.class}, tree="[0]")
    private Output<String> sourceEndpointUserName;

    /**
     * @return The source endpoint username. This parameter is not required when the `database_type` is `Redis`, or when the `source_endpoint_instance_type` is `Agent` and the `database_type` is `MSSQL`. This parameter is required in other scenarios.
     * 
     */
    public Output<String> sourceEndpointUserName() {
        return this.sourceEndpointUserName;
    }
    /**
     * The status of the resource. Valid values: `pause`, `running`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource. Valid values: `pause`, `running`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The storage region.
     * 
     */
    @Export(name="storageRegion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> storageRegion;

    /**
     * @return The storage region.
     * 
     */
    public Output<Optional<String>> storageRegion() {
        return Codegen.optional(this.storageRegion);
    }
    /**
     * Specify purchase duration. When the parameter `period` is `Year`, the `used_time` value is 1 to 9. When the parameter `period` is `Month`, the `used_time` value is 1 to 11.
     * 
     */
    @Export(name="usedTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> usedTime;

    /**
     * @return Specify purchase duration. When the parameter `period` is `Year`, the `used_time` value is 1 to 9. When the parameter `period` is `Month`, the `used_time` value is 1 to 11.
     * 
     */
    public Output<Optional<Integer>> usedTime() {
        return Codegen.optional(this.usedTime);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BackupPlan(String name) {
        this(name, BackupPlanArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BackupPlan(String name, BackupPlanArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BackupPlan(String name, BackupPlanArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dbs/backupPlan:BackupPlan", name, args == null ? BackupPlanArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BackupPlan(String name, Output<String> id, @Nullable BackupPlanState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dbs/backupPlan:BackupPlan", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "sourceEndpointPassword"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static BackupPlan get(String name, Output<String> id, @Nullable BackupPlanState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BackupPlan(name, id, state, options);
    }
}
