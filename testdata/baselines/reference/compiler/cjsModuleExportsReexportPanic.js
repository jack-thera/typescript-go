//// [tests/cases/compiler/cjsModuleExportsReexportPanic.ts] ////

//// [index.d.ts]
declare namespace extpkg {
    interface Options {
        pattern: string;
    }
}
interface ExtPkg {
    (list: readonly string[], patterns: string | readonly string[], options?: extpkg.Options): string[];
}
declare const extpkg: ExtPkg;
export = extpkg;

//// [package.json]
{ "name": "@types/extpkg", "version": "1.0.0" }

//// [test.cjs]
/** @type {typeof import("extpkg") | undefined} */
let extpkg;

module.exports.extpkg = extpkg;


//// [test.cjs]
"use strict";
/** @type {typeof import("extpkg") | undefined} */
let extpkg;
module.exports.extpkg = extpkg;


//// [test.d.cts]
export declare var extpkg: ExtPkg | undefined;
