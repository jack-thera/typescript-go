//// [tests/cases/compiler/cjsModuleExportsReexportPanic.ts] ////

//// [index.d.ts]
export interface Options {
    pattern: string;
}
export declare function match(list: readonly string[], patterns: string | readonly string[], options?: Options): string[];

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
export declare var extpkg: typeof import("extpkg") | undefined;
