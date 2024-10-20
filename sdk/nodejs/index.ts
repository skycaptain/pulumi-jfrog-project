// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { EnvironmentArgs, EnvironmentState } from "./environment";
export type Environment = import("./environment").Environment;
export const Environment: typeof import("./environment").Environment = null as any;
utilities.lazyLoad(exports, ["Environment"], () => require("./environment"));

export { GroupArgs, GroupState } from "./group";
export type Group = import("./group").Group;
export const Group: typeof import("./group").Group = null as any;
utilities.lazyLoad(exports, ["Group"], () => require("./group"));

export { ProjectArgs, ProjectState } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;
utilities.lazyLoad(exports, ["Project"], () => require("./project"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { RepositoryArgs, RepositoryState } from "./repository";
export type Repository = import("./repository").Repository;
export const Repository: typeof import("./repository").Repository = null as any;
utilities.lazyLoad(exports, ["Repository"], () => require("./repository"));

export { RoleArgs, RoleState } from "./role";
export type Role = import("./role").Role;
export const Role: typeof import("./role").Role = null as any;
utilities.lazyLoad(exports, ["Role"], () => require("./role"));

export { ShareRepositoryArgs, ShareRepositoryState } from "./shareRepository";
export type ShareRepository = import("./shareRepository").ShareRepository;
export const ShareRepository: typeof import("./shareRepository").ShareRepository = null as any;
utilities.lazyLoad(exports, ["ShareRepository"], () => require("./shareRepository"));

export { ShareRepositoryWithAllArgs, ShareRepositoryWithAllState } from "./shareRepositoryWithAll";
export type ShareRepositoryWithAll = import("./shareRepositoryWithAll").ShareRepositoryWithAll;
export const ShareRepositoryWithAll: typeof import("./shareRepositoryWithAll").ShareRepositoryWithAll = null as any;
utilities.lazyLoad(exports, ["ShareRepositoryWithAll"], () => require("./shareRepositoryWithAll"));

export { UserArgs, UserState } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "jfrog-project:index/environment:Environment":
                return new Environment(name, <any>undefined, { urn })
            case "jfrog-project:index/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "jfrog-project:index/project:Project":
                return new Project(name, <any>undefined, { urn })
            case "jfrog-project:index/repository:Repository":
                return new Repository(name, <any>undefined, { urn })
            case "jfrog-project:index/role:Role":
                return new Role(name, <any>undefined, { urn })
            case "jfrog-project:index/shareRepository:ShareRepository":
                return new ShareRepository(name, <any>undefined, { urn })
            case "jfrog-project:index/shareRepositoryWithAll:ShareRepositoryWithAll":
                return new ShareRepositoryWithAll(name, <any>undefined, { urn })
            case "jfrog-project:index/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("jfrog-project", "index/environment", _module)
pulumi.runtime.registerResourceModule("jfrog-project", "index/group", _module)
pulumi.runtime.registerResourceModule("jfrog-project", "index/project", _module)
pulumi.runtime.registerResourceModule("jfrog-project", "index/repository", _module)
pulumi.runtime.registerResourceModule("jfrog-project", "index/role", _module)
pulumi.runtime.registerResourceModule("jfrog-project", "index/shareRepository", _module)
pulumi.runtime.registerResourceModule("jfrog-project", "index/shareRepositoryWithAll", _module)
pulumi.runtime.registerResourceModule("jfrog-project", "index/user", _module)
pulumi.runtime.registerResourcePackage("jfrog-project", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:jfrog-project") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
