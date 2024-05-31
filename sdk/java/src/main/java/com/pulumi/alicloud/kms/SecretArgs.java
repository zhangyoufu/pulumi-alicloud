// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretArgs Empty = new SecretArgs();

    /**
     * The description of the secret.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the secret.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the KMS instance.
     * 
     */
    @Import(name="dkmsInstanceId")
    private @Nullable Output<String> dkmsInstanceId;

    /**
     * @return The ID of the KMS instance.
     * 
     */
    public Optional<Output<String>> dkmsInstanceId() {
        return Optional.ofNullable(this.dkmsInstanceId);
    }

    /**
     * Specifies whether to enable automatic rotation. Default value: `false`. Valid values: `true`, `false`.
     * 
     */
    @Import(name="enableAutomaticRotation")
    private @Nullable Output<Boolean> enableAutomaticRotation;

    /**
     * @return Specifies whether to enable automatic rotation. Default value: `false`. Valid values: `true`, `false`.
     * 
     */
    public Optional<Output<Boolean>> enableAutomaticRotation() {
        return Optional.ofNullable(this.enableAutomaticRotation);
    }

    /**
     * The ID of the KMS key.
     * 
     */
    @Import(name="encryptionKeyId")
    private @Nullable Output<String> encryptionKeyId;

    /**
     * @return The ID of the KMS key.
     * 
     */
    public Optional<Output<String>> encryptionKeyId() {
        return Optional.ofNullable(this.encryptionKeyId);
    }

    /**
     * The extended configuration of the secret. For more information, see [How to use it](https://www.alibabacloud.com/help/en/key-management-service/latest/kms-createsecret).
     * 
     */
    @Import(name="extendedConfig")
    private @Nullable Output<String> extendedConfig;

    /**
     * @return The extended configuration of the secret. For more information, see [How to use it](https://www.alibabacloud.com/help/en/key-management-service/latest/kms-createsecret).
     * 
     */
    public Optional<Output<String>> extendedConfig() {
        return Optional.ofNullable(this.extendedConfig);
    }

    /**
     * Specifies whether to immediately delete a secret. Default value: `false`. Valid values: `true`, `false`.
     * 
     */
    @Import(name="forceDeleteWithoutRecovery")
    private @Nullable Output<Boolean> forceDeleteWithoutRecovery;

    /**
     * @return Specifies whether to immediately delete a secret. Default value: `false`. Valid values: `true`, `false`.
     * 
     */
    public Optional<Output<Boolean>> forceDeleteWithoutRecovery() {
        return Optional.ofNullable(this.forceDeleteWithoutRecovery);
    }

    /**
     * The content of the secret policy. The value is in the JSON format. The value can be up to 32,768 bytes in length. For more information, see [How to use it](https://www.alibabacloud.com/help/en/kms/developer-reference/api-setsecretpolicy).
     * 
     */
    @Import(name="policy")
    private @Nullable Output<String> policy;

    /**
     * @return The content of the secret policy. The value is in the JSON format. The value can be up to 32,768 bytes in length. For more information, see [How to use it](https://www.alibabacloud.com/help/en/kms/developer-reference/api-setsecretpolicy).
     * 
     */
    public Optional<Output<String>> policy() {
        return Optional.ofNullable(this.policy);
    }

    /**
     * Specifies the recovery period of the secret if you do not forcibly delete it. Default value: `30`. **NOTE:**  If `force_delete_without_recovery` is set to `true`, `recovery_window_in_days` will be ignored.
     * 
     */
    @Import(name="recoveryWindowInDays")
    private @Nullable Output<Integer> recoveryWindowInDays;

    /**
     * @return Specifies the recovery period of the secret if you do not forcibly delete it. Default value: `30`. **NOTE:**  If `force_delete_without_recovery` is set to `true`, `recovery_window_in_days` will be ignored.
     * 
     */
    public Optional<Output<Integer>> recoveryWindowInDays() {
        return Optional.ofNullable(this.recoveryWindowInDays);
    }

    /**
     * The interval for automatic rotation.
     * 
     */
    @Import(name="rotationInterval")
    private @Nullable Output<String> rotationInterval;

    /**
     * @return The interval for automatic rotation.
     * 
     */
    public Optional<Output<String>> rotationInterval() {
        return Optional.ofNullable(this.rotationInterval);
    }

    /**
     * The data of the secret. **NOTE:** From version 1.204.1, attribute `secret_data` updating diff will be ignored when `secret_type` is not Generic.
     * 
     */
    @Import(name="secretData", required=true)
    private Output<String> secretData;

    /**
     * @return The data of the secret. **NOTE:** From version 1.204.1, attribute `secret_data` updating diff will be ignored when `secret_type` is not Generic.
     * 
     */
    public Output<String> secretData() {
        return this.secretData;
    }

    /**
     * The type of the secret value. Default value: `text`. Valid values: `text`, `binary`.
     * 
     */
    @Import(name="secretDataType")
    private @Nullable Output<String> secretDataType;

    /**
     * @return The type of the secret value. Default value: `text`. Valid values: `text`, `binary`.
     * 
     */
    public Optional<Output<String>> secretDataType() {
        return Optional.ofNullable(this.secretDataType);
    }

    /**
     * The name of the secret.
     * 
     */
    @Import(name="secretName", required=true)
    private Output<String> secretName;

    /**
     * @return The name of the secret.
     * 
     */
    public Output<String> secretName() {
        return this.secretName;
    }

    /**
     * The type of the secret. Valid values:
     * - `Generic`: Generic secret.
     * - `Rds`: ApsaraDB RDS secret.
     * - `RAMCredentials`: RAM secret.
     * - `ECS`: ECS secret.
     * 
     */
    @Import(name="secretType")
    private @Nullable Output<String> secretType;

    /**
     * @return The type of the secret. Valid values:
     * - `Generic`: Generic secret.
     * - `Rds`: ApsaraDB RDS secret.
     * - `RAMCredentials`: RAM secret.
     * - `ECS`: ECS secret.
     * 
     */
    public Optional<Output<String>> secretType() {
        return Optional.ofNullable(this.secretType);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The version number of the initial version.
     * 
     */
    @Import(name="versionId", required=true)
    private Output<String> versionId;

    /**
     * @return The version number of the initial version.
     * 
     */
    public Output<String> versionId() {
        return this.versionId;
    }

    /**
     * The stage label that is used to mark the new version.
     * 
     */
    @Import(name="versionStages")
    private @Nullable Output<List<String>> versionStages;

    /**
     * @return The stage label that is used to mark the new version.
     * 
     */
    public Optional<Output<List<String>>> versionStages() {
        return Optional.ofNullable(this.versionStages);
    }

    private SecretArgs() {}

    private SecretArgs(SecretArgs $) {
        this.description = $.description;
        this.dkmsInstanceId = $.dkmsInstanceId;
        this.enableAutomaticRotation = $.enableAutomaticRotation;
        this.encryptionKeyId = $.encryptionKeyId;
        this.extendedConfig = $.extendedConfig;
        this.forceDeleteWithoutRecovery = $.forceDeleteWithoutRecovery;
        this.policy = $.policy;
        this.recoveryWindowInDays = $.recoveryWindowInDays;
        this.rotationInterval = $.rotationInterval;
        this.secretData = $.secretData;
        this.secretDataType = $.secretDataType;
        this.secretName = $.secretName;
        this.secretType = $.secretType;
        this.tags = $.tags;
        this.versionId = $.versionId;
        this.versionStages = $.versionStages;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretArgs $;

        public Builder() {
            $ = new SecretArgs();
        }

        public Builder(SecretArgs defaults) {
            $ = new SecretArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the secret.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the secret.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dkmsInstanceId The ID of the KMS instance.
         * 
         * @return builder
         * 
         */
        public Builder dkmsInstanceId(@Nullable Output<String> dkmsInstanceId) {
            $.dkmsInstanceId = dkmsInstanceId;
            return this;
        }

        /**
         * @param dkmsInstanceId The ID of the KMS instance.
         * 
         * @return builder
         * 
         */
        public Builder dkmsInstanceId(String dkmsInstanceId) {
            return dkmsInstanceId(Output.of(dkmsInstanceId));
        }

        /**
         * @param enableAutomaticRotation Specifies whether to enable automatic rotation. Default value: `false`. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableAutomaticRotation(@Nullable Output<Boolean> enableAutomaticRotation) {
            $.enableAutomaticRotation = enableAutomaticRotation;
            return this;
        }

        /**
         * @param enableAutomaticRotation Specifies whether to enable automatic rotation. Default value: `false`. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableAutomaticRotation(Boolean enableAutomaticRotation) {
            return enableAutomaticRotation(Output.of(enableAutomaticRotation));
        }

        /**
         * @param encryptionKeyId The ID of the KMS key.
         * 
         * @return builder
         * 
         */
        public Builder encryptionKeyId(@Nullable Output<String> encryptionKeyId) {
            $.encryptionKeyId = encryptionKeyId;
            return this;
        }

        /**
         * @param encryptionKeyId The ID of the KMS key.
         * 
         * @return builder
         * 
         */
        public Builder encryptionKeyId(String encryptionKeyId) {
            return encryptionKeyId(Output.of(encryptionKeyId));
        }

        /**
         * @param extendedConfig The extended configuration of the secret. For more information, see [How to use it](https://www.alibabacloud.com/help/en/key-management-service/latest/kms-createsecret).
         * 
         * @return builder
         * 
         */
        public Builder extendedConfig(@Nullable Output<String> extendedConfig) {
            $.extendedConfig = extendedConfig;
            return this;
        }

        /**
         * @param extendedConfig The extended configuration of the secret. For more information, see [How to use it](https://www.alibabacloud.com/help/en/key-management-service/latest/kms-createsecret).
         * 
         * @return builder
         * 
         */
        public Builder extendedConfig(String extendedConfig) {
            return extendedConfig(Output.of(extendedConfig));
        }

        /**
         * @param forceDeleteWithoutRecovery Specifies whether to immediately delete a secret. Default value: `false`. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder forceDeleteWithoutRecovery(@Nullable Output<Boolean> forceDeleteWithoutRecovery) {
            $.forceDeleteWithoutRecovery = forceDeleteWithoutRecovery;
            return this;
        }

        /**
         * @param forceDeleteWithoutRecovery Specifies whether to immediately delete a secret. Default value: `false`. Valid values: `true`, `false`.
         * 
         * @return builder
         * 
         */
        public Builder forceDeleteWithoutRecovery(Boolean forceDeleteWithoutRecovery) {
            return forceDeleteWithoutRecovery(Output.of(forceDeleteWithoutRecovery));
        }

        /**
         * @param policy The content of the secret policy. The value is in the JSON format. The value can be up to 32,768 bytes in length. For more information, see [How to use it](https://www.alibabacloud.com/help/en/kms/developer-reference/api-setsecretpolicy).
         * 
         * @return builder
         * 
         */
        public Builder policy(@Nullable Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy The content of the secret policy. The value is in the JSON format. The value can be up to 32,768 bytes in length. For more information, see [How to use it](https://www.alibabacloud.com/help/en/kms/developer-reference/api-setsecretpolicy).
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        /**
         * @param recoveryWindowInDays Specifies the recovery period of the secret if you do not forcibly delete it. Default value: `30`. **NOTE:**  If `force_delete_without_recovery` is set to `true`, `recovery_window_in_days` will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder recoveryWindowInDays(@Nullable Output<Integer> recoveryWindowInDays) {
            $.recoveryWindowInDays = recoveryWindowInDays;
            return this;
        }

        /**
         * @param recoveryWindowInDays Specifies the recovery period of the secret if you do not forcibly delete it. Default value: `30`. **NOTE:**  If `force_delete_without_recovery` is set to `true`, `recovery_window_in_days` will be ignored.
         * 
         * @return builder
         * 
         */
        public Builder recoveryWindowInDays(Integer recoveryWindowInDays) {
            return recoveryWindowInDays(Output.of(recoveryWindowInDays));
        }

        /**
         * @param rotationInterval The interval for automatic rotation.
         * 
         * @return builder
         * 
         */
        public Builder rotationInterval(@Nullable Output<String> rotationInterval) {
            $.rotationInterval = rotationInterval;
            return this;
        }

        /**
         * @param rotationInterval The interval for automatic rotation.
         * 
         * @return builder
         * 
         */
        public Builder rotationInterval(String rotationInterval) {
            return rotationInterval(Output.of(rotationInterval));
        }

        /**
         * @param secretData The data of the secret. **NOTE:** From version 1.204.1, attribute `secret_data` updating diff will be ignored when `secret_type` is not Generic.
         * 
         * @return builder
         * 
         */
        public Builder secretData(Output<String> secretData) {
            $.secretData = secretData;
            return this;
        }

        /**
         * @param secretData The data of the secret. **NOTE:** From version 1.204.1, attribute `secret_data` updating diff will be ignored when `secret_type` is not Generic.
         * 
         * @return builder
         * 
         */
        public Builder secretData(String secretData) {
            return secretData(Output.of(secretData));
        }

        /**
         * @param secretDataType The type of the secret value. Default value: `text`. Valid values: `text`, `binary`.
         * 
         * @return builder
         * 
         */
        public Builder secretDataType(@Nullable Output<String> secretDataType) {
            $.secretDataType = secretDataType;
            return this;
        }

        /**
         * @param secretDataType The type of the secret value. Default value: `text`. Valid values: `text`, `binary`.
         * 
         * @return builder
         * 
         */
        public Builder secretDataType(String secretDataType) {
            return secretDataType(Output.of(secretDataType));
        }

        /**
         * @param secretName The name of the secret.
         * 
         * @return builder
         * 
         */
        public Builder secretName(Output<String> secretName) {
            $.secretName = secretName;
            return this;
        }

        /**
         * @param secretName The name of the secret.
         * 
         * @return builder
         * 
         */
        public Builder secretName(String secretName) {
            return secretName(Output.of(secretName));
        }

        /**
         * @param secretType The type of the secret. Valid values:
         * - `Generic`: Generic secret.
         * - `Rds`: ApsaraDB RDS secret.
         * - `RAMCredentials`: RAM secret.
         * - `ECS`: ECS secret.
         * 
         * @return builder
         * 
         */
        public Builder secretType(@Nullable Output<String> secretType) {
            $.secretType = secretType;
            return this;
        }

        /**
         * @param secretType The type of the secret. Valid values:
         * - `Generic`: Generic secret.
         * - `Rds`: ApsaraDB RDS secret.
         * - `RAMCredentials`: RAM secret.
         * - `ECS`: ECS secret.
         * 
         * @return builder
         * 
         */
        public Builder secretType(String secretType) {
            return secretType(Output.of(secretType));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param versionId The version number of the initial version.
         * 
         * @return builder
         * 
         */
        public Builder versionId(Output<String> versionId) {
            $.versionId = versionId;
            return this;
        }

        /**
         * @param versionId The version number of the initial version.
         * 
         * @return builder
         * 
         */
        public Builder versionId(String versionId) {
            return versionId(Output.of(versionId));
        }

        /**
         * @param versionStages The stage label that is used to mark the new version.
         * 
         * @return builder
         * 
         */
        public Builder versionStages(@Nullable Output<List<String>> versionStages) {
            $.versionStages = versionStages;
            return this;
        }

        /**
         * @param versionStages The stage label that is used to mark the new version.
         * 
         * @return builder
         * 
         */
        public Builder versionStages(List<String> versionStages) {
            return versionStages(Output.of(versionStages));
        }

        /**
         * @param versionStages The stage label that is used to mark the new version.
         * 
         * @return builder
         * 
         */
        public Builder versionStages(String... versionStages) {
            return versionStages(List.of(versionStages));
        }

        public SecretArgs build() {
            if ($.secretData == null) {
                throw new MissingRequiredPropertyException("SecretArgs", "secretData");
            }
            if ($.secretName == null) {
                throw new MissingRequiredPropertyException("SecretArgs", "secretName");
            }
            if ($.versionId == null) {
                throw new MissingRequiredPropertyException("SecretArgs", "versionId");
            }
            return $;
        }
    }

}
