// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AccountArgs, AccountState } from "./account";
export type Account = import("./account").Account;
export const Account: typeof import("./account").Account = null as any;
utilities.lazyLoad(exports, ["Account"], () => require("./account"));

export { AccountPrivilegeArgs, AccountPrivilegeState } from "./accountPrivilege";
export type AccountPrivilege = import("./accountPrivilege").AccountPrivilege;
export const AccountPrivilege: typeof import("./accountPrivilege").AccountPrivilege = null as any;
utilities.lazyLoad(exports, ["AccountPrivilege"], () => require("./accountPrivilege"));

export { BackupPolicyArgs, BackupPolicyState } from "./backupPolicy";
export type BackupPolicy = import("./backupPolicy").BackupPolicy;
export const BackupPolicy: typeof import("./backupPolicy").BackupPolicy = null as any;
utilities.lazyLoad(exports, ["BackupPolicy"], () => require("./backupPolicy"));

export { ConnectionArgs, ConnectionState } from "./connection";
export type Connection = import("./connection").Connection;
export const Connection: typeof import("./connection").Connection = null as any;
utilities.lazyLoad(exports, ["Connection"], () => require("./connection"));

export { DatabaseArgs, DatabaseState } from "./database";
export type Database = import("./database").Database;
export const Database: typeof import("./database").Database = null as any;
utilities.lazyLoad(exports, ["Database"], () => require("./database"));

export { DbNodeArgs, DbNodeState } from "./dbNode";
export type DbNode = import("./dbNode").DbNode;
export const DbNode: typeof import("./dbNode").DbNode = null as any;
utilities.lazyLoad(exports, ["DbNode"], () => require("./dbNode"));

export { DdrInstanceArgs, DdrInstanceState } from "./ddrInstance";
export type DdrInstance = import("./ddrInstance").DdrInstance;
export const DdrInstance: typeof import("./ddrInstance").DdrInstance = null as any;
utilities.lazyLoad(exports, ["DdrInstance"], () => require("./ddrInstance"));

export { GetAccountsArgs, GetAccountsResult, GetAccountsOutputArgs } from "./getAccounts";
export const getAccounts: typeof import("./getAccounts").getAccounts = null as any;
export const getAccountsOutput: typeof import("./getAccounts").getAccountsOutput = null as any;
utilities.lazyLoad(exports, ["getAccounts","getAccountsOutput"], () => require("./getAccounts"));

export { GetCharacterSetNamesArgs, GetCharacterSetNamesResult, GetCharacterSetNamesOutputArgs } from "./getCharacterSetNames";
export const getCharacterSetNames: typeof import("./getCharacterSetNames").getCharacterSetNames = null as any;
export const getCharacterSetNamesOutput: typeof import("./getCharacterSetNames").getCharacterSetNamesOutput = null as any;
utilities.lazyLoad(exports, ["getCharacterSetNames","getCharacterSetNamesOutput"], () => require("./getCharacterSetNames"));

export { GetCollationTimeZonesArgs, GetCollationTimeZonesResult, GetCollationTimeZonesOutputArgs } from "./getCollationTimeZones";
export const getCollationTimeZones: typeof import("./getCollationTimeZones").getCollationTimeZones = null as any;
export const getCollationTimeZonesOutput: typeof import("./getCollationTimeZones").getCollationTimeZonesOutput = null as any;
utilities.lazyLoad(exports, ["getCollationTimeZones","getCollationTimeZonesOutput"], () => require("./getCollationTimeZones"));

export { GetCrossRegionBackupsArgs, GetCrossRegionBackupsResult, GetCrossRegionBackupsOutputArgs } from "./getCrossRegionBackups";
export const getCrossRegionBackups: typeof import("./getCrossRegionBackups").getCrossRegionBackups = null as any;
export const getCrossRegionBackupsOutput: typeof import("./getCrossRegionBackups").getCrossRegionBackupsOutput = null as any;
utilities.lazyLoad(exports, ["getCrossRegionBackups","getCrossRegionBackupsOutput"], () => require("./getCrossRegionBackups"));

export { GetCrossRegionsArgs, GetCrossRegionsResult, GetCrossRegionsOutputArgs } from "./getCrossRegions";
export const getCrossRegions: typeof import("./getCrossRegions").getCrossRegions = null as any;
export const getCrossRegionsOutput: typeof import("./getCrossRegions").getCrossRegionsOutput = null as any;
utilities.lazyLoad(exports, ["getCrossRegions","getCrossRegionsOutput"], () => require("./getCrossRegions"));

export { GetInstanceClassInfosArgs, GetInstanceClassInfosResult, GetInstanceClassInfosOutputArgs } from "./getInstanceClassInfos";
export const getInstanceClassInfos: typeof import("./getInstanceClassInfos").getInstanceClassInfos = null as any;
export const getInstanceClassInfosOutput: typeof import("./getInstanceClassInfos").getInstanceClassInfosOutput = null as any;
utilities.lazyLoad(exports, ["getInstanceClassInfos","getInstanceClassInfosOutput"], () => require("./getInstanceClassInfos"));

export { GetInstanceClassesArgs, GetInstanceClassesResult, GetInstanceClassesOutputArgs } from "./getInstanceClasses";
export const getInstanceClasses: typeof import("./getInstanceClasses").getInstanceClasses = null as any;
export const getInstanceClassesOutput: typeof import("./getInstanceClasses").getInstanceClassesOutput = null as any;
utilities.lazyLoad(exports, ["getInstanceClasses","getInstanceClassesOutput"], () => require("./getInstanceClasses"));

export { GetInstanceEnginesArgs, GetInstanceEnginesResult, GetInstanceEnginesOutputArgs } from "./getInstanceEngines";
export const getInstanceEngines: typeof import("./getInstanceEngines").getInstanceEngines = null as any;
export const getInstanceEnginesOutput: typeof import("./getInstanceEngines").getInstanceEnginesOutput = null as any;
utilities.lazyLoad(exports, ["getInstanceEngines","getInstanceEnginesOutput"], () => require("./getInstanceEngines"));

export { GetInstancesArgs, GetInstancesResult, GetInstancesOutputArgs } from "./getInstances";
export const getInstances: typeof import("./getInstances").getInstances = null as any;
export const getInstancesOutput: typeof import("./getInstances").getInstancesOutput = null as any;
utilities.lazyLoad(exports, ["getInstances","getInstancesOutput"], () => require("./getInstances"));

export { GetModifyParameterLogsArgs, GetModifyParameterLogsResult, GetModifyParameterLogsOutputArgs } from "./getModifyParameterLogs";
export const getModifyParameterLogs: typeof import("./getModifyParameterLogs").getModifyParameterLogs = null as any;
export const getModifyParameterLogsOutput: typeof import("./getModifyParameterLogs").getModifyParameterLogsOutput = null as any;
utilities.lazyLoad(exports, ["getModifyParameterLogs","getModifyParameterLogsOutput"], () => require("./getModifyParameterLogs"));

export { GetRdsBackupsArgs, GetRdsBackupsResult, GetRdsBackupsOutputArgs } from "./getRdsBackups";
export const getRdsBackups: typeof import("./getRdsBackups").getRdsBackups = null as any;
export const getRdsBackupsOutput: typeof import("./getRdsBackups").getRdsBackupsOutput = null as any;
utilities.lazyLoad(exports, ["getRdsBackups","getRdsBackupsOutput"], () => require("./getRdsBackups"));

export { GetRdsParameterGroupsArgs, GetRdsParameterGroupsResult, GetRdsParameterGroupsOutputArgs } from "./getRdsParameterGroups";
export const getRdsParameterGroups: typeof import("./getRdsParameterGroups").getRdsParameterGroups = null as any;
export const getRdsParameterGroupsOutput: typeof import("./getRdsParameterGroups").getRdsParameterGroupsOutput = null as any;
utilities.lazyLoad(exports, ["getRdsParameterGroups","getRdsParameterGroupsOutput"], () => require("./getRdsParameterGroups"));

export { GetZonesArgs, GetZonesResult, GetZonesOutputArgs } from "./getZones";
export const getZones: typeof import("./getZones").getZones = null as any;
export const getZonesOutput: typeof import("./getZones").getZonesOutput = null as any;
utilities.lazyLoad(exports, ["getZones","getZonesOutput"], () => require("./getZones"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { RdsAccountArgs, RdsAccountState } from "./rdsAccount";
export type RdsAccount = import("./rdsAccount").RdsAccount;
export const RdsAccount: typeof import("./rdsAccount").RdsAccount = null as any;
utilities.lazyLoad(exports, ["RdsAccount"], () => require("./rdsAccount"));

export { RdsBackupArgs, RdsBackupState } from "./rdsBackup";
export type RdsBackup = import("./rdsBackup").RdsBackup;
export const RdsBackup: typeof import("./rdsBackup").RdsBackup = null as any;
utilities.lazyLoad(exports, ["RdsBackup"], () => require("./rdsBackup"));

export { RdsCloneDbInstanceArgs, RdsCloneDbInstanceState } from "./rdsCloneDbInstance";
export type RdsCloneDbInstance = import("./rdsCloneDbInstance").RdsCloneDbInstance;
export const RdsCloneDbInstance: typeof import("./rdsCloneDbInstance").RdsCloneDbInstance = null as any;
utilities.lazyLoad(exports, ["RdsCloneDbInstance"], () => require("./rdsCloneDbInstance"));

export { RdsDbProxyArgs, RdsDbProxyState } from "./rdsDbProxy";
export type RdsDbProxy = import("./rdsDbProxy").RdsDbProxy;
export const RdsDbProxy: typeof import("./rdsDbProxy").RdsDbProxy = null as any;
utilities.lazyLoad(exports, ["RdsDbProxy"], () => require("./rdsDbProxy"));

export { RdsInstanceCrossBackupPolicyArgs, RdsInstanceCrossBackupPolicyState } from "./rdsInstanceCrossBackupPolicy";
export type RdsInstanceCrossBackupPolicy = import("./rdsInstanceCrossBackupPolicy").RdsInstanceCrossBackupPolicy;
export const RdsInstanceCrossBackupPolicy: typeof import("./rdsInstanceCrossBackupPolicy").RdsInstanceCrossBackupPolicy = null as any;
utilities.lazyLoad(exports, ["RdsInstanceCrossBackupPolicy"], () => require("./rdsInstanceCrossBackupPolicy"));

export { RdsParameterGroupArgs, RdsParameterGroupState } from "./rdsParameterGroup";
export type RdsParameterGroup = import("./rdsParameterGroup").RdsParameterGroup;
export const RdsParameterGroup: typeof import("./rdsParameterGroup").RdsParameterGroup = null as any;
utilities.lazyLoad(exports, ["RdsParameterGroup"], () => require("./rdsParameterGroup"));

export { RdsServiceLinkedRoleArgs, RdsServiceLinkedRoleState } from "./rdsServiceLinkedRole";
export type RdsServiceLinkedRole = import("./rdsServiceLinkedRole").RdsServiceLinkedRole;
export const RdsServiceLinkedRole: typeof import("./rdsServiceLinkedRole").RdsServiceLinkedRole = null as any;
utilities.lazyLoad(exports, ["RdsServiceLinkedRole"], () => require("./rdsServiceLinkedRole"));

export { RdsUpgradeDbInstanceArgs, RdsUpgradeDbInstanceState } from "./rdsUpgradeDbInstance";
export type RdsUpgradeDbInstance = import("./rdsUpgradeDbInstance").RdsUpgradeDbInstance;
export const RdsUpgradeDbInstance: typeof import("./rdsUpgradeDbInstance").RdsUpgradeDbInstance = null as any;
utilities.lazyLoad(exports, ["RdsUpgradeDbInstance"], () => require("./rdsUpgradeDbInstance"));

export { ReadOnlyInstanceArgs, ReadOnlyInstanceState } from "./readOnlyInstance";
export type ReadOnlyInstance = import("./readOnlyInstance").ReadOnlyInstance;
export const ReadOnlyInstance: typeof import("./readOnlyInstance").ReadOnlyInstance = null as any;
utilities.lazyLoad(exports, ["ReadOnlyInstance"], () => require("./readOnlyInstance"));

export { ReadWriteSplittingConnectionArgs, ReadWriteSplittingConnectionState } from "./readWriteSplittingConnection";
export type ReadWriteSplittingConnection = import("./readWriteSplittingConnection").ReadWriteSplittingConnection;
export const ReadWriteSplittingConnection: typeof import("./readWriteSplittingConnection").ReadWriteSplittingConnection = null as any;
utilities.lazyLoad(exports, ["ReadWriteSplittingConnection"], () => require("./readWriteSplittingConnection"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:rds/account:Account":
                return new Account(name, <any>undefined, { urn })
            case "alicloud:rds/accountPrivilege:AccountPrivilege":
                return new AccountPrivilege(name, <any>undefined, { urn })
            case "alicloud:rds/backupPolicy:BackupPolicy":
                return new BackupPolicy(name, <any>undefined, { urn })
            case "alicloud:rds/connection:Connection":
                return new Connection(name, <any>undefined, { urn })
            case "alicloud:rds/database:Database":
                return new Database(name, <any>undefined, { urn })
            case "alicloud:rds/dbNode:DbNode":
                return new DbNode(name, <any>undefined, { urn })
            case "alicloud:rds/ddrInstance:DdrInstance":
                return new DdrInstance(name, <any>undefined, { urn })
            case "alicloud:rds/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "alicloud:rds/rdsAccount:RdsAccount":
                return new RdsAccount(name, <any>undefined, { urn })
            case "alicloud:rds/rdsBackup:RdsBackup":
                return new RdsBackup(name, <any>undefined, { urn })
            case "alicloud:rds/rdsCloneDbInstance:RdsCloneDbInstance":
                return new RdsCloneDbInstance(name, <any>undefined, { urn })
            case "alicloud:rds/rdsDbProxy:RdsDbProxy":
                return new RdsDbProxy(name, <any>undefined, { urn })
            case "alicloud:rds/rdsInstanceCrossBackupPolicy:RdsInstanceCrossBackupPolicy":
                return new RdsInstanceCrossBackupPolicy(name, <any>undefined, { urn })
            case "alicloud:rds/rdsParameterGroup:RdsParameterGroup":
                return new RdsParameterGroup(name, <any>undefined, { urn })
            case "alicloud:rds/rdsServiceLinkedRole:RdsServiceLinkedRole":
                return new RdsServiceLinkedRole(name, <any>undefined, { urn })
            case "alicloud:rds/rdsUpgradeDbInstance:RdsUpgradeDbInstance":
                return new RdsUpgradeDbInstance(name, <any>undefined, { urn })
            case "alicloud:rds/readOnlyInstance:ReadOnlyInstance":
                return new ReadOnlyInstance(name, <any>undefined, { urn })
            case "alicloud:rds/readWriteSplittingConnection:ReadWriteSplittingConnection":
                return new ReadWriteSplittingConnection(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "rds/account", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/accountPrivilege", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/backupPolicy", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/connection", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/database", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/dbNode", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/ddrInstance", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/instance", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/rdsAccount", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/rdsBackup", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/rdsCloneDbInstance", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/rdsDbProxy", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/rdsInstanceCrossBackupPolicy", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/rdsParameterGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/rdsServiceLinkedRole", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/rdsUpgradeDbInstance", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/readOnlyInstance", _module)
pulumi.runtime.registerResourceModule("alicloud", "rds/readWriteSplittingConnection", _module)
