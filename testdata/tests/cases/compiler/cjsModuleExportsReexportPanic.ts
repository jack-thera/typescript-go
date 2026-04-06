// @module: commonjs
// @declaration: true
// @allowJs: true
// @checkJs: true
// @outDir: ./out

// @filename: /node_modules/@types/extpkg/index.d.ts
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

// @filename: /node_modules/@types/extpkg/package.json
{ "name": "@types/extpkg", "version": "1.0.0" }

// @filename: /test.cjs
/** @type {typeof import("extpkg") | undefined} */
let extpkg;

module.exports.extpkg = extpkg;
