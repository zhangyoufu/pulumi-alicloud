// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EcsKeyPairArgs extends com.pulumi.resources.ResourceArgs {

    public static final EcsKeyPairArgs Empty = new EcsKeyPairArgs();

    /**
     * The key file.
     * 
     */
    @Import(name="keyFile")
    private @Nullable Output<String> keyFile;

    /**
     * @return The key file.
     * 
     */
    public Optional<Output<String>> keyFile() {
        return Optional.ofNullable(this.keyFile);
    }

    /**
     * Field `key_name` has been deprecated from provider version 1.121.0. New field `key_pair_name` instead.
     * 
     * @deprecated
     * Field &#39;key_name&#39; has been deprecated from provider version 1.121.0. New field &#39;key_pair_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead. */
    @Import(name="keyName")
    private @Nullable Output<String> keyName;

    /**
     * @return Field `key_name` has been deprecated from provider version 1.121.0. New field `key_pair_name` instead.
     * 
     * @deprecated
     * Field &#39;key_name&#39; has been deprecated from provider version 1.121.0. New field &#39;key_pair_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead. */
    public Optional<Output<String>> keyName() {
        return Optional.ofNullable(this.keyName);
    }

    @Import(name="keyNamePrefix")
    private @Nullable Output<String> keyNamePrefix;

    public Optional<Output<String>> keyNamePrefix() {
        return Optional.ofNullable(this.keyNamePrefix);
    }

    /**
     * The key pair&#39;s name. It is the only in one Alicloud account, the key pair&#39;s name. must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    @Import(name="keyPairName")
    private @Nullable Output<String> keyPairName;

    /**
     * @return The key pair&#39;s name. It is the only in one Alicloud account, the key pair&#39;s name. must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    public Optional<Output<String>> keyPairName() {
        return Optional.ofNullable(this.keyPairName);
    }

    /**
     * You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
     * 
     */
    @Import(name="publicKey")
    private @Nullable Output<String> publicKey;

    /**
     * @return You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
     * 
     */
    public Optional<Output<String>> publicKey() {
        return Optional.ofNullable(this.publicKey);
    }

    /**
     * The Id of resource group which the key pair belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the key pair belongs.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private EcsKeyPairArgs() {}

    private EcsKeyPairArgs(EcsKeyPairArgs $) {
        this.keyFile = $.keyFile;
        this.keyName = $.keyName;
        this.keyNamePrefix = $.keyNamePrefix;
        this.keyPairName = $.keyPairName;
        this.publicKey = $.publicKey;
        this.resourceGroupId = $.resourceGroupId;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EcsKeyPairArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EcsKeyPairArgs $;

        public Builder() {
            $ = new EcsKeyPairArgs();
        }

        public Builder(EcsKeyPairArgs defaults) {
            $ = new EcsKeyPairArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param keyFile The key file.
         * 
         * @return builder
         * 
         */
        public Builder keyFile(@Nullable Output<String> keyFile) {
            $.keyFile = keyFile;
            return this;
        }

        /**
         * @param keyFile The key file.
         * 
         * @return builder
         * 
         */
        public Builder keyFile(String keyFile) {
            return keyFile(Output.of(keyFile));
        }

        /**
         * @param keyName Field `key_name` has been deprecated from provider version 1.121.0. New field `key_pair_name` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;key_name&#39; has been deprecated from provider version 1.121.0. New field &#39;key_pair_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead. */
        public Builder keyName(@Nullable Output<String> keyName) {
            $.keyName = keyName;
            return this;
        }

        /**
         * @param keyName Field `key_name` has been deprecated from provider version 1.121.0. New field `key_pair_name` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;key_name&#39; has been deprecated from provider version 1.121.0. New field &#39;key_pair_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead. */
        public Builder keyName(String keyName) {
            return keyName(Output.of(keyName));
        }

        public Builder keyNamePrefix(@Nullable Output<String> keyNamePrefix) {
            $.keyNamePrefix = keyNamePrefix;
            return this;
        }

        public Builder keyNamePrefix(String keyNamePrefix) {
            return keyNamePrefix(Output.of(keyNamePrefix));
        }

        /**
         * @param keyPairName The key pair&#39;s name. It is the only in one Alicloud account, the key pair&#39;s name. must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder keyPairName(@Nullable Output<String> keyPairName) {
            $.keyPairName = keyPairName;
            return this;
        }

        /**
         * @param keyPairName The key pair&#39;s name. It is the only in one Alicloud account, the key pair&#39;s name. must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder keyPairName(String keyPairName) {
            return keyPairName(Output.of(keyPairName));
        }

        /**
         * @param publicKey You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(@Nullable Output<String> publicKey) {
            $.publicKey = publicKey;
            return this;
        }

        /**
         * @param publicKey You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, `resource_group_id` is the key pair belongs.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(String publicKey) {
            return publicKey(Output.of(publicKey));
        }

        /**
         * @param resourceGroupId The Id of resource group which the key pair belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The Id of resource group which the key pair belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public EcsKeyPairArgs build() {
            return $;
        }
    }

}
