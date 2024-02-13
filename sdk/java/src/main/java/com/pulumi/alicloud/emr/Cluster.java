// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emr;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.emr.ClusterArgs;
import com.pulumi.alicloud.emr.inputs.ClusterState;
import com.pulumi.alicloud.emr.outputs.ClusterBootstrapAction;
import com.pulumi.alicloud.emr.outputs.ClusterConfig;
import com.pulumi.alicloud.emr.outputs.ClusterHostGroup;
import com.pulumi.alicloud.emr.outputs.ClusterMetaStoreConf;
import com.pulumi.alicloud.emr.outputs.ClusterModifyClusterServiceConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a EMR Cluster resource. With this you can create, read, and release  EMR Cluster.
 * 
 * &gt; **DEPRECATED:**  This resource has been deprecated from version `1.204.0`. Please use new resource emrv2_cluster.
 * 
 * &gt; **NOTE:** Available in 1.57.0+.
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * Aliclioud E-MapReduce cluster can be imported using the id e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:emr/cluster:Cluster default C-B47FB8FE96C67XXXX
 * ```
 * 
 */
@ResourceType(type="alicloud:emr/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    /**
     * Boot action parameters.
     * 
     */
    @Export(name="bootstrapActions", refs={List.class,ClusterBootstrapAction.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ClusterBootstrapAction>> bootstrapActions;

    /**
     * @return Boot action parameters.
     * 
     */
    public Output<Optional<List<ClusterBootstrapAction>>> bootstrapActions() {
        return Codegen.optional(this.bootstrapActions);
    }
    /**
     * Charge Type for this cluster. Supported value: PostPaid or PrePaid. Default value: PostPaid.
     * 
     */
    @Export(name="chargeType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> chargeType;

    /**
     * @return Charge Type for this cluster. Supported value: PostPaid or PrePaid. Default value: PostPaid.
     * 
     */
    public Output<Optional<String>> chargeType() {
        return Codegen.optional(this.chargeType);
    }
    /**
     * EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported &#39;GATEWAY&#39; available in 1.61.0+.
     * 
     */
    @Export(name="clusterType", refs={String.class}, tree="[0]")
    private Output<String> clusterType;

    /**
     * @return EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported &#39;GATEWAY&#39; available in 1.61.0+.
     * 
     */
    public Output<String> clusterType() {
        return this.clusterType;
    }
    /**
     * The custom configurations of emr-cluster service.
     * 
     */
    @Export(name="configs", refs={List.class,ClusterConfig.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ClusterConfig>> configs;

    /**
     * @return The custom configurations of emr-cluster service.
     * 
     */
    public Output<Optional<List<ClusterConfig>>> configs() {
        return Codegen.optional(this.configs);
    }
    /**
     * Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
     * 
     */
    @Export(name="depositType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> depositType;

    /**
     * @return Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
     * 
     */
    public Output<Optional<String>> depositType() {
        return Codegen.optional(this.depositType);
    }
    /**
     * High security cluster (true) or not. Default value is false.
     * 
     */
    @Export(name="easEnable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> easEnable;

    /**
     * @return High security cluster (true) or not. Default value is false.
     * 
     */
    public Output<Optional<Boolean>> easEnable() {
        return Codegen.optional(this.easEnable);
    }
    /**
     * EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
     * 
     */
    @Export(name="emrVer", refs={String.class}, tree="[0]")
    private Output<String> emrVer;

    /**
     * @return EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
     * 
     */
    public Output<String> emrVer() {
        return this.emrVer;
    }
    /**
     * High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
     * 
     */
    @Export(name="highAvailabilityEnable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> highAvailabilityEnable;

    /**
     * @return High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
     * 
     */
    public Output<Optional<Boolean>> highAvailabilityEnable() {
        return Codegen.optional(this.highAvailabilityEnable);
    }
    /**
     * Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
     * 
     */
    @Export(name="hostGroups", refs={List.class,ClusterHostGroup.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ClusterHostGroup>> hostGroups;

    /**
     * @return Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
     * 
     */
    public Output<Optional<List<ClusterHostGroup>>> hostGroups() {
        return Codegen.optional(this.hostGroups);
    }
    /**
     * Whether the MASTER node has a public IP address enabled. Default value is false.
     * 
     */
    @Export(name="isOpenPublicIp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isOpenPublicIp;

    /**
     * @return Whether the MASTER node has a public IP address enabled. Default value is false.
     * 
     */
    public Output<Optional<Boolean>> isOpenPublicIp() {
        return Codegen.optional(this.isOpenPublicIp);
    }
    /**
     * Ssh key pair.
     * 
     */
    @Export(name="keyPairName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> keyPairName;

    /**
     * @return Ssh key pair.
     * 
     */
    public Output<Optional<String>> keyPairName() {
        return Codegen.optional(this.keyPairName);
    }
    /**
     * Master ssh password.
     * 
     */
    @Export(name="masterPwd", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> masterPwd;

    /**
     * @return Master ssh password.
     * 
     */
    public Output<Optional<String>> masterPwd() {
        return Codegen.optional(this.masterPwd);
    }
    /**
     * The configuration of emr-cluster service component metadata storage. If meta store type is ’user_rds’, this should be specified.
     * 
     */
    @Export(name="metaStoreConf", refs={ClusterMetaStoreConf.class}, tree="[0]")
    private Output</* @Nullable */ ClusterMetaStoreConf> metaStoreConf;

    /**
     * @return The configuration of emr-cluster service component metadata storage. If meta store type is ’user_rds’, this should be specified.
     * 
     */
    public Output<Optional<ClusterMetaStoreConf>> metaStoreConf() {
        return Codegen.optional(this.metaStoreConf);
    }
    /**
     * The type of emr-cluster service component metadata storage. ’dlf’ or ’local’ or ’user_rds’ .
     * 
     */
    @Export(name="metaStoreType", refs={String.class}, tree="[0]")
    private Output<String> metaStoreType;

    /**
     * @return The type of emr-cluster service component metadata storage. ’dlf’ or ’local’ or ’user_rds’ .
     * 
     */
    public Output<String> metaStoreType() {
        return this.metaStoreType;
    }
    /**
     * The configurations of emr-cluster service modification after cluster created.
     * 
     */
    @Export(name="modifyClusterServiceConfig", refs={ClusterModifyClusterServiceConfig.class}, tree="[0]")
    private Output</* @Nullable */ ClusterModifyClusterServiceConfig> modifyClusterServiceConfig;

    /**
     * @return The configurations of emr-cluster service modification after cluster created.
     * 
     */
    public Output<Optional<ClusterModifyClusterServiceConfig>> modifyClusterServiceConfig() {
        return Codegen.optional(this.modifyClusterServiceConfig);
    }
    /**
     * The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, &#34;-&#34;, &#34;_&#34;.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, &#34;-&#34;, &#34;_&#34;.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Optional software list.
     * 
     */
    @Export(name="optionSoftwareLists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> optionSoftwareLists;

    /**
     * @return Optional software list.
     * 
     */
    public Output<Optional<List<String>>> optionSoftwareLists() {
        return Codegen.optional(this.optionSoftwareLists);
    }
    /**
     * If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * This specify the related cluster id, if this cluster is a Gateway.
     * 
     */
    @Export(name="relatedClusterId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> relatedClusterId;

    /**
     * @return This specify the related cluster id, if this cluster is a Gateway.
     * 
     */
    public Output<Optional<String>> relatedClusterId() {
        return Codegen.optional(this.relatedClusterId);
    }
    /**
     * The Id of resource group which the emr-cluster belongs.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The Id of resource group which the emr-cluster belongs.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * Security Group ID for Cluster, you can also specify this key for each host group.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityGroupId;

    /**
     * @return Security Group ID for Cluster, you can also specify this key for each host group.
     * 
     */
    public Output<Optional<String>> securityGroupId() {
        return Codegen.optional(this.securityGroupId);
    }
    /**
     * If this is set true, we can ssh into cluster. Default value is false.
     * 
     */
    @Export(name="sshEnable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> sshEnable;

    /**
     * @return If this is set true, we can ssh into cluster. Default value is false.
     * 
     */
    public Output<Optional<Boolean>> sshEnable() {
        return Codegen.optional(this.sshEnable);
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Map<String,Object>> tags() {
        return this.tags;
    }
    /**
     * Use local metadb. Default is false.
     * 
     */
    @Export(name="useLocalMetadb", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useLocalMetadb;

    /**
     * @return Use local metadb. Default is false.
     * 
     */
    public Output<Optional<Boolean>> useLocalMetadb() {
        return Codegen.optional(this.useLocalMetadb);
    }
    /**
     * Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
     * 
     */
    @Export(name="userDefinedEmrEcsRole", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userDefinedEmrEcsRole;

    /**
     * @return Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
     * 
     */
    public Output<Optional<String>> userDefinedEmrEcsRole() {
        return Codegen.optional(this.userDefinedEmrEcsRole);
    }
    /**
     * Global vswitch id, you can also specify it in host group.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return Global vswitch id, you can also specify it in host group.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }
    /**
     * Zone ID, e.g. cn-huhehaote-a
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return Zone ID, e.g. cn-huhehaote-a
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Cluster(String name) {
        this(name, ClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cluster(String name, ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:emr/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:emr/cluster:Cluster", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}
