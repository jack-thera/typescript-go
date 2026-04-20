// @allowJs: true
// @checkJs: true
// @noEmit: true
// @module: nodenext
// @target: esnext

// @filename: /runner.js
function Runner() {}
Runner.constants = {
    EVENT_RUN_BEGIN: "begin",
};
module.exports = Runner;

// @filename: /stats-collector.mjs
/**
 * @typedef {import('./runner.js')} Runner
 */

import Runner from "./runner.js";

const { EVENT_RUN_BEGIN } = Runner.constants;

function createStatsCollector(runner) {}

export { createStatsCollector };
