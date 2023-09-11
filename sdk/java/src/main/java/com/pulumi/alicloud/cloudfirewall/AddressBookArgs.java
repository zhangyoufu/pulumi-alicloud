// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall;

import com.pulumi.alicloud.cloudfirewall.inputs.AddressBookEcsTagArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AddressBookArgs extends com.pulumi.resources.ResourceArgs {

    public static final AddressBookArgs Empty = new AddressBookArgs();

    /**
     * The list of addresses.
     * 
     */
    @Import(name="addressLists")
    private @Nullable Output<List<String>> addressLists;

    /**
     * @return The list of addresses.
     * 
     */
    public Optional<Output<List<String>>> addressLists() {
        return Optional.ofNullable(this.addressLists);
    }

    /**
     * Whether you want to automatically add new matching tags of the ECS IP address to the Address Book. Valid values: `0`, `1`.
     * 
     */
    @Import(name="autoAddTagEcs")
    private @Nullable Output<Integer> autoAddTagEcs;

    /**
     * @return Whether you want to automatically add new matching tags of the ECS IP address to the Address Book. Valid values: `0`, `1`.
     * 
     */
    public Optional<Output<Integer>> autoAddTagEcs() {
        return Optional.ofNullable(this.autoAddTagEcs);
    }

    /**
     * The description of the Address Book.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return The description of the Address Book.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * A list of ECS tags. See `ecs_tags` below.
     * 
     */
    @Import(name="ecsTags")
    private @Nullable Output<List<AddressBookEcsTagArgs>> ecsTags;

    /**
     * @return A list of ECS tags. See `ecs_tags` below.
     * 
     */
    public Optional<Output<List<AddressBookEcsTagArgs>>> ecsTags() {
        return Optional.ofNullable(this.ecsTags);
    }

    /**
     * The name of the Address Book.
     * 
     */
    @Import(name="groupName", required=true)
    private Output<String> groupName;

    /**
     * @return The name of the Address Book.
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
    }

    /**
     * The type of the Address Book. Valid values: `ip`, `tag`.
     * 
     */
    @Import(name="groupType", required=true)
    private Output<String> groupType;

    /**
     * @return The type of the Address Book. Valid values: `ip`, `tag`.
     * 
     */
    public Output<String> groupType() {
        return this.groupType;
    }

    /**
     * The language of the content within the request and response. Valid values: `zh`, `en`.
     * 
     */
    @Import(name="lang")
    private @Nullable Output<String> lang;

    /**
     * @return The language of the content within the request and response. Valid values: `zh`, `en`.
     * 
     */
    public Optional<Output<String>> lang() {
        return Optional.ofNullable(this.lang);
    }

    /**
     * The logical relation among the ECS tags that to be matched. Default value: `and`. Valid values:
     * 
     */
    @Import(name="tagRelation")
    private @Nullable Output<String> tagRelation;

    /**
     * @return The logical relation among the ECS tags that to be matched. Default value: `and`. Valid values:
     * 
     */
    public Optional<Output<String>> tagRelation() {
        return Optional.ofNullable(this.tagRelation);
    }

    private AddressBookArgs() {}

    private AddressBookArgs(AddressBookArgs $) {
        this.addressLists = $.addressLists;
        this.autoAddTagEcs = $.autoAddTagEcs;
        this.description = $.description;
        this.ecsTags = $.ecsTags;
        this.groupName = $.groupName;
        this.groupType = $.groupType;
        this.lang = $.lang;
        this.tagRelation = $.tagRelation;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AddressBookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AddressBookArgs $;

        public Builder() {
            $ = new AddressBookArgs();
        }

        public Builder(AddressBookArgs defaults) {
            $ = new AddressBookArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressLists The list of addresses.
         * 
         * @return builder
         * 
         */
        public Builder addressLists(@Nullable Output<List<String>> addressLists) {
            $.addressLists = addressLists;
            return this;
        }

        /**
         * @param addressLists The list of addresses.
         * 
         * @return builder
         * 
         */
        public Builder addressLists(List<String> addressLists) {
            return addressLists(Output.of(addressLists));
        }

        /**
         * @param addressLists The list of addresses.
         * 
         * @return builder
         * 
         */
        public Builder addressLists(String... addressLists) {
            return addressLists(List.of(addressLists));
        }

        /**
         * @param autoAddTagEcs Whether you want to automatically add new matching tags of the ECS IP address to the Address Book. Valid values: `0`, `1`.
         * 
         * @return builder
         * 
         */
        public Builder autoAddTagEcs(@Nullable Output<Integer> autoAddTagEcs) {
            $.autoAddTagEcs = autoAddTagEcs;
            return this;
        }

        /**
         * @param autoAddTagEcs Whether you want to automatically add new matching tags of the ECS IP address to the Address Book. Valid values: `0`, `1`.
         * 
         * @return builder
         * 
         */
        public Builder autoAddTagEcs(Integer autoAddTagEcs) {
            return autoAddTagEcs(Output.of(autoAddTagEcs));
        }

        /**
         * @param description The description of the Address Book.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the Address Book.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ecsTags A list of ECS tags. See `ecs_tags` below.
         * 
         * @return builder
         * 
         */
        public Builder ecsTags(@Nullable Output<List<AddressBookEcsTagArgs>> ecsTags) {
            $.ecsTags = ecsTags;
            return this;
        }

        /**
         * @param ecsTags A list of ECS tags. See `ecs_tags` below.
         * 
         * @return builder
         * 
         */
        public Builder ecsTags(List<AddressBookEcsTagArgs> ecsTags) {
            return ecsTags(Output.of(ecsTags));
        }

        /**
         * @param ecsTags A list of ECS tags. See `ecs_tags` below.
         * 
         * @return builder
         * 
         */
        public Builder ecsTags(AddressBookEcsTagArgs... ecsTags) {
            return ecsTags(List.of(ecsTags));
        }

        /**
         * @param groupName The name of the Address Book.
         * 
         * @return builder
         * 
         */
        public Builder groupName(Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        /**
         * @param groupName The name of the Address Book.
         * 
         * @return builder
         * 
         */
        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
        }

        /**
         * @param groupType The type of the Address Book. Valid values: `ip`, `tag`.
         * 
         * @return builder
         * 
         */
        public Builder groupType(Output<String> groupType) {
            $.groupType = groupType;
            return this;
        }

        /**
         * @param groupType The type of the Address Book. Valid values: `ip`, `tag`.
         * 
         * @return builder
         * 
         */
        public Builder groupType(String groupType) {
            return groupType(Output.of(groupType));
        }

        /**
         * @param lang The language of the content within the request and response. Valid values: `zh`, `en`.
         * 
         * @return builder
         * 
         */
        public Builder lang(@Nullable Output<String> lang) {
            $.lang = lang;
            return this;
        }

        /**
         * @param lang The language of the content within the request and response. Valid values: `zh`, `en`.
         * 
         * @return builder
         * 
         */
        public Builder lang(String lang) {
            return lang(Output.of(lang));
        }

        /**
         * @param tagRelation The logical relation among the ECS tags that to be matched. Default value: `and`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder tagRelation(@Nullable Output<String> tagRelation) {
            $.tagRelation = tagRelation;
            return this;
        }

        /**
         * @param tagRelation The logical relation among the ECS tags that to be matched. Default value: `and`. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder tagRelation(String tagRelation) {
            return tagRelation(Output.of(tagRelation));
        }

        public AddressBookArgs build() {
            $.description = Objects.requireNonNull($.description, "expected parameter 'description' to be non-null");
            $.groupName = Objects.requireNonNull($.groupName, "expected parameter 'groupName' to be non-null");
            $.groupType = Objects.requireNonNull($.groupType, "expected parameter 'groupType' to be non-null");
            return $;
        }
    }

}
