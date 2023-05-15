// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds.inputs;

import com.pulumi.alicloud.rds.inputs.GetCollationTimeZonesCollationTimeZone;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCollationTimeZonesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCollationTimeZonesPlainArgs Empty = new GetCollationTimeZonesPlainArgs();

    /**
     * An array that consists of the character set collations and time zones that are available for
     * use in ApsaraDB RDS.
     * 
     */
    @Import(name="collationTimeZones")
    private @Nullable List<GetCollationTimeZonesCollationTimeZone> collationTimeZones;

    /**
     * @return An array that consists of the character set collations and time zones that are available for
     * use in ApsaraDB RDS.
     * 
     */
    public Optional<List<GetCollationTimeZonesCollationTimeZone>> collationTimeZones() {
        return Optional.ofNullable(this.collationTimeZones);
    }

    /**
     * File name where to save data source results (after running `pulumi up`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi up`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetCollationTimeZonesPlainArgs() {}

    private GetCollationTimeZonesPlainArgs(GetCollationTimeZonesPlainArgs $) {
        this.collationTimeZones = $.collationTimeZones;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCollationTimeZonesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCollationTimeZonesPlainArgs $;

        public Builder() {
            $ = new GetCollationTimeZonesPlainArgs();
        }

        public Builder(GetCollationTimeZonesPlainArgs defaults) {
            $ = new GetCollationTimeZonesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param collationTimeZones An array that consists of the character set collations and time zones that are available for
         * use in ApsaraDB RDS.
         * 
         * @return builder
         * 
         */
        public Builder collationTimeZones(@Nullable List<GetCollationTimeZonesCollationTimeZone> collationTimeZones) {
            $.collationTimeZones = collationTimeZones;
            return this;
        }

        /**
         * @param collationTimeZones An array that consists of the character set collations and time zones that are available for
         * use in ApsaraDB RDS.
         * 
         * @return builder
         * 
         */
        public Builder collationTimeZones(GetCollationTimeZonesCollationTimeZone... collationTimeZones) {
            return collationTimeZones(List.of(collationTimeZones));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi up`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public GetCollationTimeZonesPlainArgs build() {
            return $;
        }
    }

}
