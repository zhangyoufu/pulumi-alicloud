// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ots.Outputs
{

    [OutputType]
    public sealed class SearchIndexSchemaFieldSchema
    {
        /// <summary>
        /// Specifies the type of the analyzer that you want to use. If fieldType is set to Text, you can configure this parameter. Otherwise, the default analyzer type single-word tokenization is used.
        /// </summary>
        public readonly string? Analyzer;
        /// <summary>
        /// Specifies whether to enable sorting and aggregation. Type: Boolean. Sorting can be enabled only for fields for which enable_sort_and_agg is set to true.
        /// </summary>
        public readonly bool? EnableSortAndAgg;
        /// <summary>
        /// The name of the field that is used to sort data. only required if sorter_type is FieldSort.
        /// </summary>
        public readonly string FieldName;
        /// <summary>
        /// Specifies the type of the field. Use FieldType.XXX to set the type.
        /// </summary>
        public readonly string FieldType;
        /// <summary>
        /// Specifies whether to enable indexing for the column. Type: Boolean.
        /// </summary>
        public readonly bool? Index;
        /// <summary>
        /// Specifies whether the value is an array. Type: Boolean.
        /// </summary>
        public readonly bool? IsArray;
        /// <summary>
        /// Specifies whether to store the value of the field in the search index. Type: Boolean. If you set store to true, you can read the value of the field from the search index without querying the data table. This improves query performance.
        /// </summary>
        public readonly bool? Store;

        [OutputConstructor]
        private SearchIndexSchemaFieldSchema(
            string? analyzer,

            bool? enableSortAndAgg,

            string fieldName,

            string fieldType,

            bool? index,

            bool? isArray,

            bool? store)
        {
            Analyzer = analyzer;
            EnableSortAndAgg = enableSortAndAgg;
            FieldName = fieldName;
            FieldType = fieldType;
            Index = index;
            IsArray = isArray;
            Store = store;
        }
    }
}
