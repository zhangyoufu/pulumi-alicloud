// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAlidnsDomainsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAlidnsDomainsPlainArgs Empty = new GetAlidnsDomainsPlainArgs();

    /**
     * Specifies whether the domain is from Alibaba Cloud or not.
     * 
     */
    @Import(name="aliDomain")
    private @Nullable Boolean aliDomain;

    /**
     * @return Specifies whether the domain is from Alibaba Cloud or not.
     * 
     */
    public Optional<Boolean> aliDomain() {
        return Optional.ofNullable(this.aliDomain);
    }

    /**
     * A regex string to filter results by the domain name.
     * 
     */
    @Import(name="domainNameRegex")
    private @Nullable String domainNameRegex;

    /**
     * @return A regex string to filter results by the domain name.
     * 
     */
    public Optional<String> domainNameRegex() {
        return Optional.ofNullable(this.domainNameRegex);
    }

    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * Domain group ID, if not filled, the default is all groups.
     * 
     */
    @Import(name="groupId")
    private @Nullable String groupId;

    /**
     * @return Domain group ID, if not filled, the default is all groups.
     * 
     */
    public Optional<String> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * A regex string to filter results by the group name.
     * 
     */
    @Import(name="groupNameRegex")
    private @Nullable String groupNameRegex;

    /**
     * @return A regex string to filter results by the group name.
     * 
     */
    public Optional<String> groupNameRegex() {
        return Optional.ofNullable(this.groupNameRegex);
    }

    /**
     * A list of domain IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of domain IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * Cloud analysis product ID.
     * 
     */
    @Import(name="instanceId")
    private @Nullable String instanceId;

    /**
     * @return Cloud analysis product ID.
     * 
     */
    public Optional<String> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The keywords are searched according to the `%KeyWord%` mode, which is not case sensitive.
     * 
     */
    @Import(name="keyWord")
    private @Nullable String keyWord;

    /**
     * @return The keywords are searched according to the `%KeyWord%` mode, which is not case sensitive.
     * 
     */
    public Optional<String> keyWord() {
        return Optional.ofNullable(this.keyWord);
    }

    /**
     * User language.
     * 
     */
    @Import(name="lang")
    private @Nullable String lang;

    /**
     * @return User language.
     * 
     */
    public Optional<String> lang() {
        return Optional.ofNullable(this.lang);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The Id of resource group which the dns belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable String resourceGroupId;

    /**
     * @return The Id of resource group which the dns belongs.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * Search mode, `LIKE` fuzzy search, `EXACT` exact search.
     * 
     */
    @Import(name="searchMode")
    private @Nullable String searchMode;

    /**
     * @return Search mode, `LIKE` fuzzy search, `EXACT` exact search.
     * 
     */
    public Optional<String> searchMode() {
        return Optional.ofNullable(this.searchMode);
    }

    /**
     * Whether to query the domain name star.
     * 
     */
    @Import(name="starmark")
    private @Nullable Boolean starmark;

    /**
     * @return Whether to query the domain name star.
     * 
     */
    public Optional<Boolean> starmark() {
        return Optional.ofNullable(this.starmark);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,Object> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Map<String,Object>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Cloud analysis version code.
     * 
     */
    @Import(name="versionCode")
    private @Nullable String versionCode;

    /**
     * @return Cloud analysis version code.
     * 
     */
    public Optional<String> versionCode() {
        return Optional.ofNullable(this.versionCode);
    }

    private GetAlidnsDomainsPlainArgs() {}

    private GetAlidnsDomainsPlainArgs(GetAlidnsDomainsPlainArgs $) {
        this.aliDomain = $.aliDomain;
        this.domainNameRegex = $.domainNameRegex;
        this.enableDetails = $.enableDetails;
        this.groupId = $.groupId;
        this.groupNameRegex = $.groupNameRegex;
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.keyWord = $.keyWord;
        this.lang = $.lang;
        this.outputFile = $.outputFile;
        this.resourceGroupId = $.resourceGroupId;
        this.searchMode = $.searchMode;
        this.starmark = $.starmark;
        this.tags = $.tags;
        this.versionCode = $.versionCode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAlidnsDomainsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAlidnsDomainsPlainArgs $;

        public Builder() {
            $ = new GetAlidnsDomainsPlainArgs();
        }

        public Builder(GetAlidnsDomainsPlainArgs defaults) {
            $ = new GetAlidnsDomainsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param aliDomain Specifies whether the domain is from Alibaba Cloud or not.
         * 
         * @return builder
         * 
         */
        public Builder aliDomain(@Nullable Boolean aliDomain) {
            $.aliDomain = aliDomain;
            return this;
        }

        /**
         * @param domainNameRegex A regex string to filter results by the domain name.
         * 
         * @return builder
         * 
         */
        public Builder domainNameRegex(@Nullable String domainNameRegex) {
            $.domainNameRegex = domainNameRegex;
            return this;
        }

        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param groupId Domain group ID, if not filled, the default is all groups.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable String groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupNameRegex A regex string to filter results by the group name.
         * 
         * @return builder
         * 
         */
        public Builder groupNameRegex(@Nullable String groupNameRegex) {
            $.groupNameRegex = groupNameRegex;
            return this;
        }

        /**
         * @param ids A list of domain IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of domain IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId Cloud analysis product ID.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable String instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param keyWord The keywords are searched according to the `%KeyWord%` mode, which is not case sensitive.
         * 
         * @return builder
         * 
         */
        public Builder keyWord(@Nullable String keyWord) {
            $.keyWord = keyWord;
            return this;
        }

        /**
         * @param lang User language.
         * 
         * @return builder
         * 
         */
        public Builder lang(@Nullable String lang) {
            $.lang = lang;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param resourceGroupId The Id of resource group which the dns belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param searchMode Search mode, `LIKE` fuzzy search, `EXACT` exact search.
         * 
         * @return builder
         * 
         */
        public Builder searchMode(@Nullable String searchMode) {
            $.searchMode = searchMode;
            return this;
        }

        /**
         * @param starmark Whether to query the domain name star.
         * 
         * @return builder
         * 
         */
        public Builder starmark(@Nullable Boolean starmark) {
            $.starmark = starmark;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,Object> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param versionCode Cloud analysis version code.
         * 
         * @return builder
         * 
         */
        public Builder versionCode(@Nullable String versionCode) {
            $.versionCode = versionCode;
            return this;
        }

        public GetAlidnsDomainsPlainArgs build() {
            return $;
        }
    }

}
